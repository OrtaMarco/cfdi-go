package cfdi

import (
	"crypto/x509"
	"encoding/asn1"
	"fmt"
	"strings"
	"time"
)

// Estados de un chequeo del certificado.
const (
	ChequeoCumple        = "cumple"         // el documento y el certificado coinciden
	ChequeoNoCumple      = "no_cumple"      // no coinciden
	ChequeoNoVerificable = "no_verificable" // falta un dato para decidir
)

// Motivos de los chequeos del certificado, además de MotivoSinCertificado,
// MotivoCertificadoIlegible y MotivoXMLIlegible (no verificable). Son
// estables.
const (
	MotivoCertificadoNoVigenteAun = "CERTIFICADO_NO_VIGENTE_AUN" // no cumple: la fecha es anterior a NotBefore
	MotivoCertificadoVencido      = "CERTIFICADO_VENCIDO"        // no cumple: la fecha es posterior a NotAfter
	MotivoVigenciaAmbigua         = "VIGENCIA_AMBIGUA"           // no verificable: la fecha sin zona cae a menos de la diferencia de husos del límite
	MotivoFechaIlegible           = "FECHA_ILEGIBLE"             // no verificable: no hay fecha o no se puede leer
	MotivoRFCCertificadoDistinto  = "RFC_CERTIFICADO_DISTINTO"   // no cumple: el RFC del certificado no es el del emisor
	MotivoCertificadoSinRFC       = "CERTIFICADO_SIN_RFC"        // no verificable: el certificado no trae x500UniqueIdentifier
	MotivoSinRFCEmisor            = "SIN_RFC_EMISOR"             // no verificable: el documento no trae el RFC del emisor
	MotivoSinNoCertificado        = "SIN_NOCERTIFICADO"          // no cumple: el documento no trae NoCertificado
	MotivoCertificadoDelSAT       = "CERTIFICADO_DEL_SAT"        // cumple: el titular del certificado es el SAT, que sella a nombre del emisor
	// NoCertificado distinto usa MotivoNoCertificadoDistinto.
)

// rfcSAT es el RFC del Servicio de Administración Tributaria: el titular de
// los certificados con que su servicio gratuito de facturación sella a
// nombre del contribuyente, y el RfcProvCertif del timbre que pone el SAT.
const rfcSAT = "SAT970701NN3"

// MotivosCertificado lista los motivos que pueden salir en
// ChequeosCertificado, en el orden de arriba y luego los compartidos.
var MotivosCertificado = []string{
	MotivoCertificadoNoVigenteAun, MotivoCertificadoVencido, MotivoVigenciaAmbigua, MotivoFechaIlegible,
	MotivoRFCCertificadoDistinto, MotivoCertificadoSinRFC, MotivoSinRFCEmisor, MotivoSinNoCertificado,
	MotivoCertificadoDelSAT,
	MotivoNoCertificadoDistinto, MotivoSinCertificado, MotivoCertificadoIlegible, MotivoXMLIlegible,
}

// ResultadoChequeo es uno de los chequeos del certificado.
type ResultadoChequeo struct {
	Estado  string `json:"estado"`            // ChequeoCumple, ChequeoNoCumple o ChequeoNoVerificable
	Motivo  string `json:"motivo,omitempty"`  // código estable; vacío si cumple, salvo CERTIFICADO_DEL_SAT
	Mensaje string `json:"mensaje,omitempty"` // para personas; puede cambiar
}

// ChequeosCertificado compara el certificado del emisor (el Certificado del
// XML) con el documento, sin red: no dice si el certificado estaba revocado
// ni si lo emitió el SAT.
//
//   - Vigencia: la Fecha del CFDI (FechaExp en Retenciones) entre NotBefore y
//     NotAfter del certificado, los dos incluidos. La Fecha del CFDI y la
//     FechaExp de Retenciones 2.0 son hora civil sin zona del lugar de
//     expedición, y México tiene husos de UTC−5 a UTC−8 (contando el horario
//     de verano que hubo hasta 2022): la librería no sabe cuál es, así que
//     esa hora es algún instante entre hora+5h y hora+8h UTC. Cumple si cae
//     en la vigencia con todos los husos; no cumple si no cae con ninguno;
//     si depende del huso (la fecha está a menos de tres horas de un
//     límite), no verificable con VIGENCIA_AMBIGUA. Una fecha con zona (la
//     FechaExp de Retenciones 1.0) es un instante exacto.
//   - Rfc: el RFC del certificado es el del emisor. Es el atributo
//     x500UniqueIdentifier (OID 2.5.4.45) del sujeto; en una persona moral
//     trae "RFC / RFC del representante" y cuenta el primero. Si el titular
//     es el SAT (SAT970701NN3) y no el emisor, es un CFDI del servicio gratuito de
//     facturación del SAT, que sella con su certificado a nombre del
//     contribuyente: cumple con motivo CERTIFICADO_DEL_SAT (el único cumple
//     con motivo), para que quien lo lea sepa que el RFC no coincide y por
//     qué está bien. El SAT como representante, después de la barra, no
//     cuenta.
//   - NoCertificado: el del documento es el número de serie del certificado
//     (ver NoCertificado). Es lo mismo que exige el sello, pero aquí se
//     revisa aunque la firma no verifique.
type ChequeosCertificado struct {
	Vigencia      ResultadoChequeo `json:"vigencia"`
	Rfc           ResultadoChequeo `json:"rfc"`
	NoCertificado ResultadoChequeo `json:"no_certificado"`
}

// oidRFC es x500UniqueIdentifier, donde el SAT pone el RFC del titular.
var oidRFC = asn1.ObjectIdentifier{2, 5, 4, 45}

// Husos de México: la hora civil h es, en UTC, algún instante entre
// h+husoMin y h+husoMax.
const (
	husoMin = 5 * time.Hour
	husoMax = 8 * time.Hour
)

func noVerificableCh(motivo, msg string) ResultadoChequeo {
	return ResultadoChequeo{Estado: ChequeoNoVerificable, Motivo: motivo, Mensaje: msg}
}

func chequeosNoVerificables(err error) *ChequeosCertificado {
	r := noVerificableCh(MotivoXMLIlegible, err.Error())
	return &ChequeosCertificado{Vigencia: r, Rfc: r, NoCertificado: r}
}

func chequearCertificado(raiz *nodoC, p perfilSello) *ChequeosCertificado {
	certB64, ok := raiz.attrOk(p.cert)
	if !ok || strings.TrimSpace(certB64) == "" {
		r := noVerificableCh(MotivoSinCertificado, "el documento no trae "+p.cert)
		return &ChequeosCertificado{Vigencia: r, Rfc: r, NoCertificado: r}
	}
	cert, _, err := leerCertificado(certB64)
	if err != nil {
		r := noVerificableCh(MotivoCertificadoIlegible, err.Error())
		return &ChequeosCertificado{Vigencia: r, Rfc: r, NoCertificado: r}
	}
	fecha, conFecha := raiz.attrOk(p.fecha)
	var rfcEmisor string
	for _, h := range raiz.hijos {
		if h.nombre == (nombreXML{raiz.nombre.ns, "Emisor"}) {
			rfcEmisor = h.attr(p.rfcEmisor)
			break
		}
	}
	noCert, conNoCert := raiz.attrOk(p.noCert)
	var rfcProv string
	if tfd := timbreDe(raiz); tfd != nil {
		rfcProv = tfd.attr("RfcProvCertif")
	}
	return &ChequeosCertificado{
		Vigencia:      chequearVigencia(cert, fecha, conFecha),
		Rfc:           chequearRFC(cert, rfcEmisor, rfcProv),
		NoCertificado: chequearNoCertificado(cert, p.noCert, noCert, conNoCert),
	}
}

func chequearVigencia(cert *x509.Certificate, fecha string, conFecha bool) ResultadoChequeo {
	if !conFecha || strings.TrimSpace(fecha) == "" {
		return noVerificableCh(MotivoFechaIlegible, "el documento no trae fecha")
	}
	desde, hasta, err := instanteFecha(fecha)
	if err != nil {
		return noVerificableCh(MotivoFechaIlegible, err.Error())
	}
	nb, na := cert.NotBefore, cert.NotAfter
	rango := fmt.Sprintf("vigencia del %s al %s UTC", nb.UTC().Format(time.DateTime), na.UTC().Format(time.DateTime))
	switch {
	case hasta.Before(nb):
		return ResultadoChequeo{Estado: ChequeoNoCumple, Motivo: MotivoCertificadoNoVigenteAun, Mensaje: fecha + " es anterior a la " + rango}
	case desde.After(na):
		return ResultadoChequeo{Estado: ChequeoNoCumple, Motivo: MotivoCertificadoVencido, Mensaje: fecha + " es posterior a la " + rango}
	case desde.Before(nb) || hasta.After(na):
		return noVerificableCh(MotivoVigenciaAmbigua, fecha+" sin zona: según el huso del lugar de expedición queda dentro o fuera de la "+rango)
	}
	return ResultadoChequeo{Estado: ChequeoCumple}
}

// instanteFecha lee la fecha del documento como el intervalo de instantes
// que puede ser: uno solo si trae zona; con hora civil sin zona, de
// hora+husoMin a hora+husoMax.
func instanteFecha(s string) (desde, hasta time.Time, err error) {
	t := strings.TrimSpace(s)
	if z, e := time.Parse(time.RFC3339Nano, t); e == nil {
		return z, z, nil
	}
	f, _, e := parseFecha(t, false)
	if e != nil {
		return time.Time{}, time.Time{}, e
	}
	return f.Time.Add(husoMin), f.Time.Add(husoMax), nil
}

// RFCDelCertificado devuelve el RFC del titular de un certificado del SAT:
// el x500UniqueIdentifier (OID 2.5.4.45) del sujeto, que en una persona
// moral trae "RFC / RFC del representante" (cuenta el primero). "" si no lo
// trae.
func RFCDelCertificado(c *x509.Certificate) string {
	for _, a := range c.Subject.Names {
		if !a.Type.Equal(oidRFC) {
			continue
		}
		s, _ := a.Value.(string)
		titular, _, _ := strings.Cut(s, "/")
		return strings.ToUpper(strings.TrimSpace(titular))
	}
	return ""
}

// chequearRFC compara el titular del certificado con el emisor. rfcProv es
// el RfcProvCertif del timbre ("" si no hay): solo cambia el mensaje cuando
// el certificado es del SAT, porque los timbres viejos no lo traen.
func chequearRFC(cert *x509.Certificate, rfcEmisor, rfcProv string) ResultadoChequeo {
	delCert := RFCDelCertificado(cert)
	emisor := rfc(rfcEmisor)
	switch {
	case delCert == "":
		return noVerificableCh(MotivoCertificadoSinRFC, "el certificado no trae x500UniqueIdentifier (2.5.4.45)")
	case emisor == "":
		return noVerificableCh(MotivoSinRFCEmisor, "el documento no trae el RFC del emisor")
	case delCert != emisor && delCert == rfcSAT:
		msg := fmt.Sprintf("el certificado es del SAT (%s), que sella a nombre del emisor %q con su servicio gratuito de facturación", rfcSAT, emisor)
		if rfc(rfcProv) == rfcSAT {
			msg += "; el timbre también es del SAT (RfcProvCertif " + rfcSAT + ")"
		}
		return ResultadoChequeo{Estado: ChequeoCumple, Motivo: MotivoCertificadoDelSAT, Mensaje: msg}
	case delCert != emisor:
		return ResultadoChequeo{Estado: ChequeoNoCumple, Motivo: MotivoRFCCertificadoDistinto,
			Mensaje: fmt.Sprintf("el certificado es de %q y el emisor es %q", delCert, emisor)}
	}
	return ResultadoChequeo{Estado: ChequeoCumple}
}

func chequearNoCertificado(cert *x509.Certificate, nombre, valor string, presente bool) ResultadoChequeo {
	no := NoCertificado(cert)
	switch {
	case !presente || strings.TrimSpace(valor) == "":
		return ResultadoChequeo{Estado: ChequeoNoCumple, Motivo: MotivoSinNoCertificado, Mensaje: "el documento no trae " + nombre}
	case strings.TrimSpace(valor) != no:
		return ResultadoChequeo{Estado: ChequeoNoCumple, Motivo: MotivoNoCertificadoDistinto,
			Mensaje: fmt.Sprintf("%s %q y el certificado es %q", nombre, valor, no)}
	}
	return ResultadoChequeo{Estado: ChequeoCumple}
}
