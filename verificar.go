package cfdi

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"
)

// Estado del resultado de verificar un sello.
const (
	SelloValido        = "valido"         // la firma corresponde a la cadena y al certificado
	SelloInvalido      = "invalido"       // el documento dice algo que su sello no respalda
	SelloNoVerificable = "no_verificable" // no se pudo decidir (cadena no garantizable, certificado ilegible o ausente…)
)

// Motivos de un resultado de sello, además de los de ErrorCadena (que dan
// SelloNoVerificable). Son estables.
const (
	MotivoSinSello                 = "SIN_SELLO"                     // inválido: no hay Sello (o SelloSAT)
	MotivoSelloIlegible            = "SELLO_ILEGIBLE"                // inválido: el sello no es base64
	MotivoSelloNoCoincide          = "SELLO_NO_COINCIDE"             // inválido: la firma no es de esa cadena con esa llave
	MotivoNoCertificadoDistinto    = "NOCERTIFICADO_DISTINTO"        // inválido: NoCertificado no es el número del certificado
	MotivoSelloCFDDistinto         = "SELLO_CFD_DISTINTO"            // inválido: el SelloCFD del TFD no es el Sello del comprobante
	MotivoSinCertificado           = "SIN_CERTIFICADO"               // no verificable: no viene el Certificado
	MotivoCertificadoIlegible      = "CERTIFICADO_ILEGIBLE"          // no verificable: no es base64 de un X.509 DER
	MotivoLlaveNoRSA               = "LLAVE_NO_RSA"                  // no verificable: el certificado no trae llave RSA
	MotivoCertificadoSATFaltante   = "CERTIFICADO_SAT_NO_DISPONIBLE" // no verificable: el proveedor no dio el certificado del SAT
	MotivoNoCertificadoSATDistinto = "NOCERTIFICADO_SAT_DISTINTO"    // no verificable: el proveedor dio otro certificado
)

// Motivos lista todos los motivos de cadena y de sello, en el orden de
// arriba.
var Motivos = []string{
	MotivoXMLIlegible, MotivoNoEsComprobante, MotivoNoEsRetenciones, MotivoVersionNoSoportada, MotivoComplementoSinCadena, MotivoSinTimbre,
	MotivoSinSello, MotivoSelloIlegible, MotivoSelloNoCoincide, MotivoNoCertificadoDistinto, MotivoSelloCFDDistinto,
	MotivoSinCertificado, MotivoCertificadoIlegible, MotivoLlaveNoRSA, MotivoCertificadoSATFaltante, MotivoNoCertificadoSATDistinto,
}

// ResultadoSello es la verificación de un sello (el del emisor o el del
// SAT en el timbre).
type ResultadoSello struct {
	Estado  string `json:"estado"`            // SelloValido, SelloInvalido o SelloNoVerificable
	Motivo  string `json:"motivo,omitempty"`  // código estable; vacío si es válido
	Mensaje string `json:"mensaje,omitempty"` // para personas; puede cambiar
	// Algoritmo con que verificó la firma ("SHA256withRSA", "SHA1withRSA").
	Algoritmo string `json:"algoritmo,omitempty"`
	// NoCertificado del certificado con que se verificó (ver NoCertificado).
	NoCertificado string `json:"no_certificado,omitempty"`
}

// Verificacion es el resultado de Verificar.
type Verificacion struct {
	// CadenaOriginal del comprobante; vacía si no se pudo garantizar.
	CadenaOriginal string `json:"cadena_original,omitempty"`
	// Sello del emisor sobre CadenaOriginal, con el Certificado del XML.
	Sello ResultadoSello `json:"sello"`
	// CadenaOriginalTFD del timbre, si lo hay y se pudo armar.
	CadenaOriginalTFD string `json:"cadena_original_tfd,omitempty"`
	// TFD es el SelloSAT del timbre; nil si no se pidió (certificados nil).
	TFD *ResultadoSello `json:"tfd,omitempty"`
	// Certificado son los chequeos del certificado del emisor (vigencia,
	// RFC y número), aparte del sello: no cambian Sello.Estado.
	Certificado *ChequeosCertificado `json:"certificado,omitempty"`
}

// CertificadosSAT da el certificado del SAT (o del PAC) con que se timbró,
// por su número (NoCertificadoSAT del timbre). El certificado no viene en
// el XML y la librería no usa la red: quien llama decide de dónde sale
// (un directorio, una caché, una descarga). Para no encontrado, devolver
// un error.
type CertificadosSAT func(noCertificadoSAT string) (*x509.Certificate, error)

// ErrCertificadoNoEncontrado lo devuelve el proveedor de CertificadosEn.
var ErrCertificadoNoEncontrado = errors.New("certificado no encontrado")

// CertificadosEn es un proveedor sobre certificados ya cargados, indexados
// por NoCertificado.
func CertificadosEn(certs ...*x509.Certificate) CertificadosSAT {
	m := make(map[string]*x509.Certificate, len(certs))
	for _, c := range certs {
		m[NoCertificado(c)] = c
	}
	return func(no string) (*x509.Certificate, error) {
		if c, ok := m[no]; ok {
			return c, nil
		}
		return nil, fmt.Errorf("%w: %s", ErrCertificadoNoEncontrado, no)
	}
}

// NoCertificado es el número de certificado en el formato del SAT: el
// número de serie del certificado, cuyos bytes son los dígitos en ASCII
// (serie 0x3330303031… → "30001…"). Si los bytes no son ASCII imprimible
// (un certificado que no es del SAT), devuelve la serie en hexadecimal.
func NoCertificado(c *x509.Certificate) string {
	b := c.SerialNumber.Bytes()
	for _, x := range b {
		if x < 0x20 || x > 0x7e {
			return strings.ToUpper(hex.EncodeToString(b))
		}
	}
	return string(b)
}

// Verificar arma la cadena original y verifica el sello del emisor con el
// Certificado que trae el XML: RSA PKCS#1 v1.5 con SHA-256 (CFDI 3.3 y 4.0;
// en 3.2 también SHA-1), y que NoCertificado sea el número de ese
// certificado. No revisa la cadena de confianza del certificado ni si estaba
// revocado (eso necesita red).
//
// Aparte, sin cambiar el veredicto del sello, Verificacion.Certificado dice
// si la Fecha cae en la vigencia del certificado, si su RFC es el del Emisor
// y si su número es el NoCertificado (ver ChequeosCertificado).
//
// Si certs no es nil, verifica también el timbre: SelloSAT sobre la cadena
// del TFD con el certificado que dé certs para NoCertificadoSAT (SHA-256 en
// TFD 1.1; en 1.0 también SHA-1), y que SelloCFD sea el Sello del
// comprobante.
//
// Nunca falla: lo que no se puede decidir sale como SelloNoVerificable con
// su motivo.
func Verificar(xml []byte, certs CertificadosSAT) *Verificacion {
	return verificar(xml, certs, perfilComprobante)
}

func verificar(xml []byte, certs CertificadosSAT, perfil func(*nodoC) perfilSello) *Verificacion {
	v := &Verificacion{}
	raiz, correccion, err := arbolCadena(xml)
	if err != nil {
		v.Sello = noVerificable(err)
		if certs != nil {
			r := noVerificable(err)
			v.TFD = &r
		}
		v.Certificado = chequeosNoVerificables(err)
		return v
	}
	p := perfil(raiz)
	v.Sello, v.CadenaOriginal = verificarSelloEmisor(raiz, correccion, p)
	v.Certificado = chequearCertificado(raiz, p)
	if certs != nil {
		r := verificarTimbre(raiz, correccion, certs, &v.CadenaOriginalTFD)
		v.TFD = &r
	} else if tfd := timbreDe(raiz); tfd != nil && correccion == "" {
		v.CadenaOriginalTFD, _ = cadenaTimbre(tfd)
	}
	return v
}

func noVerificable(err error) ResultadoSello {
	r := ResultadoSello{Estado: SelloNoVerificable, Mensaje: err.Error()}
	var ec *ErrorCadena
	if errors.As(err, &ec) {
		r.Motivo = ec.Motivo
	}
	return r
}

func resultado(estado, motivo, msg string) ResultadoSello {
	return ResultadoSello{Estado: estado, Motivo: motivo, Mensaje: msg}
}

// atributosSello son los nombres de Sello, Certificado y NoCertificado por
// versión: 3.2 los escribe en minúsculas.
func atributosSello(raiz *nodoC) (sello, cert, noCert string) {
	if raiz.attr("Version") == "" && raiz.attr("version") != "" {
		return "sello", "certificado", "noCertificado"
	}
	return "Sello", "Certificado", "NoCertificado"
}

// perfilSello dice cómo se sella un tipo de documento: su cadena, los
// nombres de Sello, Certificado y NoCertificado, los algoritmos que se
// aceptan, y dónde están la fecha y el RFC del emisor (para los chequeos del
// certificado).
type perfilSello struct {
	cadena              func(*nodoC) (string, error)
	sello, cert, noCert string
	algoritmos          []crypto.Hash
	fecha, rfcEmisor    string // atributo de la raíz y de su hijo Emisor
}

// perfilComprobante es el del CFDI: 3.2 en minúsculas y con SHA-1.
func perfilComprobante(raiz *nodoC) perfilSello {
	p := perfilSello{cadena: cadenaComprobante, algoritmos: []crypto.Hash{crypto.SHA256}, fecha: "Fecha", rfcEmisor: "Rfc"}
	p.sello, p.cert, p.noCert = atributosSello(raiz)
	if raiz.attr("version") == "3.2" {
		p.algoritmos = append(p.algoritmos, crypto.SHA1)
	}
	if p.sello == "sello" {
		p.fecha, p.rfcEmisor = "fecha", "rfc"
	}
	return p
}

func verificarSelloEmisor(raiz *nodoC, correccion string, p perfilSello) (ResultadoSello, string) {
	cadena, err := p.cadena(raiz)
	if err != nil {
		return noVerificable(err), ""
	}
	aSello, aNoCert := p.sello, p.noCert
	certB64, ok := raiz.attrOk(p.cert)
	if !ok || strings.TrimSpace(certB64) == "" {
		return resultado(SelloNoVerificable, MotivoSinCertificado, "el comprobante no trae Certificado"), cadena
	}
	cert, pemAdentro, err := leerCertificado(certB64)
	if err != nil {
		return resultado(SelloNoVerificable, MotivoCertificadoIlegible, err.Error()), cadena
	}
	r := verificarFirma(cert, raiz.attr(aSello), cadena, p.algoritmos)
	if r.Estado == SelloValido {
		if no := raiz.attr(aNoCert); no != r.NoCertificado {
			r = resultado(SelloInvalido, MotivoNoCertificadoDistinto,
				fmt.Sprintf("NoCertificado %q y el certificado es %q", no, r.NoCertificado))
		} else if pemAdentro {
			r.Mensaje = "el Certificado viene como base64 de un PEM y no de un DER"
		}
	}
	if correccion != "" {
		// Hubo que corregir el encoding para leerlo: si aun así la firma
		// verifica, la lectura es la que se firmó; si no, no se sabe.
		if r.Estado != SelloValido {
			return resultado(SelloNoVerificable, MotivoXMLIlegible, correccion), ""
		}
		r.Mensaje = "el sello verifica con el XML leído corrigiendo su codificación: " + correccion
	}
	return r, cadena
}

func verificarTimbre(raiz *nodoC, correccion string, certs CertificadosSAT, cadenaTFD *string) ResultadoSello {
	tfd := timbreDe(raiz)
	if tfd == nil {
		return resultado(SelloNoVerificable, MotivoSinTimbre, "no hay TimbreFiscalDigital")
	}
	if correccion != "" {
		return resultado(SelloNoVerificable, MotivoXMLIlegible, correccion)
	}
	cadena, err := cadenaTimbre(tfd)
	if err != nil {
		return noVerificable(err)
	}
	*cadenaTFD = cadena
	v11 := tfd.attr("Version") == "1.1"
	aSello, aNoCert, aSelloCFD := "SelloSAT", "NoCertificadoSAT", "SelloCFD"
	algoritmos := []crypto.Hash{crypto.SHA256}
	if !v11 {
		aSello, aNoCert, aSelloCFD = "selloSAT", "noCertificadoSAT", "selloCFD"
		algoritmos = append(algoritmos, crypto.SHA1)
	}
	no := tfd.attr(aNoCert)
	cert, err := certs(no)
	switch {
	case err != nil:
		return resultado(SelloNoVerificable, MotivoCertificadoSATFaltante, err.Error())
	case cert == nil:
		return resultado(SelloNoVerificable, MotivoCertificadoSATFaltante, "el proveedor no dio certificado para "+no)
	case NoCertificado(cert) != no:
		return resultado(SelloNoVerificable, MotivoNoCertificadoSATDistinto,
			fmt.Sprintf("se pidió %q y el proveedor dio %q", no, NoCertificado(cert)))
	}
	r := verificarFirma(cert, tfd.attr(aSello), cadena, algoritmos)
	if r.Estado != SelloValido {
		return r
	}
	aSelloComprobante, _, _ := atributosSello(raiz)
	if sinEspacios(tfd.attr(aSelloCFD)) != sinEspacios(raiz.attr(aSelloComprobante)) {
		return ResultadoSello{Estado: SelloInvalido, Motivo: MotivoSelloCFDDistinto, Algoritmo: r.Algoritmo, NoCertificado: r.NoCertificado,
			Mensaje: "el SAT selló el timbre, pero su SelloCFD no es el Sello del comprobante"}
	}
	return r
}

// verificarFirma comprueba un sello en base64 sobre la cadena. Prueba los
// algoritmos en orden y se queda con el primero que verifique.
func verificarFirma(cert *x509.Certificate, selloB64, cadena string, algoritmos []crypto.Hash) ResultadoSello {
	if strings.TrimSpace(selloB64) == "" {
		return resultado(SelloInvalido, MotivoSinSello, "el sello está vacío o no viene")
	}
	firma, err := base64.StdEncoding.DecodeString(sinEspacios(selloB64))
	if err != nil {
		return resultado(SelloInvalido, MotivoSelloIlegible, "el sello no es base64: "+err.Error())
	}
	pub, ok := cert.PublicKey.(*rsa.PublicKey)
	if !ok {
		return resultado(SelloNoVerificable, MotivoLlaveNoRSA, fmt.Sprintf("llave %T", cert.PublicKey))
	}
	for _, h := range algoritmos {
		if rsa.VerifyPKCS1v15(pub, h, digerir(h, cadena), firma) == nil {
			return ResultadoSello{Estado: SelloValido, Algoritmo: nombreAlgoritmo(h), NoCertificado: NoCertificado(cert)}
		}
	}
	r := resultado(SelloInvalido, MotivoSelloNoCoincide, "la firma no corresponde a la cadena original con la llave del certificado")
	r.NoCertificado = NoCertificado(cert)
	return r
}

func digerir(h crypto.Hash, s string) []byte {
	if h == crypto.SHA1 {
		d := sha1.Sum([]byte(s))
		return d[:]
	}
	d := sha256.Sum256([]byte(s))
	return d[:]
}

func nombreAlgoritmo(h crypto.Hash) string {
	if h == crypto.SHA1 {
		return "SHA1withRSA"
	}
	return "SHA256withRSA"
}

// leerCertificado decodifica el atributo Certificado: base64 del DER. Hay
// emisores que ponen el base64 de un PEM; se acepta y pem dice que fue así.
func leerCertificado(b64 string) (c *x509.Certificate, pemAdentro bool, err error) {
	der, err := base64.StdEncoding.DecodeString(sinEspacios(b64))
	if err != nil {
		return nil, false, fmt.Errorf("el Certificado no es base64: %w", err)
	}
	if bloque, _ := pem.Decode(der); bloque != nil && bloque.Type == "CERTIFICATE" {
		der, pemAdentro = bloque.Bytes, true
	}
	c, err = x509.ParseCertificate(der)
	if err != nil {
		return nil, false, fmt.Errorf("el Certificado no es un X.509 legible: %w", err)
	}
	return c, pemAdentro, nil
}

// sinEspacios quita los espacios y saltos que algunos emisores meten en el
// base64 de sellos y certificados.
func sinEspacios(s string) string {
	return strings.Map(func(r rune) rune {
		if r == ' ' || r == '\t' || r == '\r' || r == '\n' {
			return -1
		}
		return r
	}, s)
}
