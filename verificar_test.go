package cfdi

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"math/big"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"
)

// llave es un CSD de pruebas del SAT (testdata/certs, ver NOTICE).
type llave struct {
	cert *x509.Certificate
	der  []byte
	priv *rsa.PrivateKey
}

func cargarLlave(t *testing.T, nombre string) llave {
	t.Helper()
	der, err := os.ReadFile("testdata/certs/" + nombre + ".cer")
	if err != nil {
		t.Fatal(err)
	}
	cert, err := x509.ParseCertificate(der)
	if err != nil {
		t.Fatal(err)
	}
	k := llave{cert: cert, der: der}
	p, err := os.ReadFile("testdata/certs/" + nombre + ".key.pem")
	if err == nil {
		bloque, _ := pem.Decode(p)
		priv, err := x509.ParsePKCS8PrivateKey(bloque.Bytes)
		if err != nil {
			t.Fatal(err)
		}
		k.priv = priv.(*rsa.PrivateKey)
	}
	return k
}

func (k llave) firmar(t *testing.T, h crypto.Hash, cadena string) string {
	t.Helper()
	d := digerir(h, cadena)
	firma, err := rsa.SignPKCS1v15(nil, k.priv, h, d)
	if err != nil {
		t.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(firma)
}

// conAttr pone el valor de un atributo (la primera aparición con ese nombre
// exacto).
func conAttr(t *testing.T, xml, nombre, valor string) string {
	t.Helper()
	re := regexp.MustCompile(`(\s)` + nombre + `="[^"]*"`)
	loc := re.FindStringIndex(xml)
	if loc == nil {
		t.Fatalf("no hay atributo %s", nombre)
	}
	return xml[:loc[0]+1] + nombre + `="` + valor + `"` + xml[loc[1]:]
}

// sellar llena NoCertificado, Certificado y Sello del comprobante con la
// llave del emisor y, si pac no es nil, SelloCFD, NoCertificadoSAT y
// SelloSAT del timbre. La cadena sale de CadenaOriginal, que las pruebas de
// cadena comparan contra libxslt en estos mismos fixtures.
func sellar(t *testing.T, xml string, emisor llave, h crypto.Hash, pac *llave) string {
	t.Helper()
	sello, cert, noCert := "Sello", "Certificado", "NoCertificado"
	selloCFD, noCertSAT, selloSAT := "SelloCFD", "NoCertificadoSAT", "SelloSAT"
	if strings.Contains(xml, `version="3.2"`) {
		sello, cert, noCert = "sello", "certificado", "noCertificado"
		selloCFD, noCertSAT, selloSAT = "selloCFD", "noCertificadoSAT", "selloSAT"
	}
	xml = conAttr(t, xml, noCert, NoCertificado(emisor.cert))
	xml = conAttr(t, xml, cert, base64.StdEncoding.EncodeToString(emisor.der))
	cadena, err := CadenaOriginal([]byte(xml))
	if err != nil {
		t.Fatal(err)
	}
	s := emisor.firmar(t, h, cadena)
	xml = conAttr(t, xml, sello, s)
	if pac == nil {
		return xml
	}
	xml = conAttr(t, xml, selloCFD, s)
	xml = conAttr(t, xml, noCertSAT, NoCertificado(pac.cert))
	cadenaTFD, err := CadenaOriginalTFD([]byte(xml))
	if err != nil {
		t.Fatal(err)
	}
	return conAttr(t, xml, selloSAT, pac.firmar(t, h, cadenaTFD))
}

// Fixtures sellados de punta a punta con los CSD de pruebas: el emisor
// EKU9003173C9 y, como PAC, CSD01_AAA010101AAA. Se regeneran con
//
//	go test -run TestSellados -update . && python3 scripts/generar_cadenas.py && go test -run TestGolden -update .
//
// (la firma PKCS#1 v1.5 es determinista: regenerar sin cambios no cambia
// los archivos).
var sellados = []struct {
	origen string
	hash   crypto.Hash
}{
	{"ingreso40-pue.xml", crypto.SHA256},
	{"ingreso33.xml", crypto.SHA256},
	{"pago20.xml", crypto.SHA256},
	{"nomina12.xml", crypto.SHA256},
	{"cadena/domicilios32.xml", crypto.SHA1},
}

func TestSellados(t *testing.T) {
	emisor, pac := cargarLlave(t, "EKU9003173C9"), cargarLlave(t, "CSD01_AAA010101AAA")
	certs := CertificadosEn(pac.cert)
	for _, s := range sellados {
		destino := filepath.Join("testdata", "sinteticos", "sellados", filepath.Base(s.origen))
		if *update {
			xml := sellar(t, string(leerSintetico(t, s.origen)), emisor, s.hash, &pac)
			if err := os.WriteFile(destino, []byte(xml), 0o644); err != nil {
				t.Fatal(err)
			}
		}
		b, err := os.ReadFile(destino)
		if err != nil {
			t.Fatalf("%v (corre con -update)", err)
		}
		v := Verificar(b, certs)
		algoritmo := nombreAlgoritmo(s.hash)
		if v.Sello.Estado != SelloValido || v.Sello.Algoritmo != algoritmo || v.Sello.NoCertificado != "30001000000500003416" {
			t.Errorf("%s: sello %+v", destino, v.Sello)
		}
		if v.TFD == nil || v.TFD.Estado != SelloValido || v.TFD.Algoritmo != algoritmo || v.TFD.NoCertificado != "30001000000300023708" {
			t.Errorf("%s: TFD %+v", destino, v.TFD)
		}
		cad, _ := CadenaOriginal(b)
		cadT, _ := CadenaOriginalTFD(b)
		if v.CadenaOriginal != cad || v.CadenaOriginalTFD != cadT || cad == "" || cadT == "" {
			t.Errorf("%s: Verificar no devolvió las cadenas", destino)
		}
	}
}

// Sellos genuinos de otros proyectos: los firmaron ellos (CfdiUtils,
// nodecfdi, satcfdi) sobre la cadena que armaron sus XSLT, así que
// verifican la cadena y la firma de punta a punta sin usar nada de aquí.
func TestSellosPublicosValidos(t *testing.T) {
	for archivo, no := range map[string]string{
		"eclipxe13-CfdiUtils/tests__assets__cfdi33-valid.xml":                                        "30001000000400002434",
		"eclipxe13-CfdiUtils/tests__assets__cfdi40-valid.xml":                                        "30001000000500003416",
		"nodecfdi-cfdi-validator/tests___files__cfdi40-valid.xml":                                    "30001000000400002434",
		"SAT-CFDI-python-satcfdi/tests__cfdi_ejemplos__comprobante40__cfdv40-ejemplo-signed-tfd.xml": "",
		"SAT-CFDI-python-satcfdi/tests__cfdi_ejemplos__comprobante33__cfdv33-signed-tfd.xml":         "",
	} {
		b, err := os.ReadFile(filepath.Join("testdata", "publicos", archivo))
		if err != nil {
			t.Fatal(err)
		}
		v := Verificar(b, nil)
		if v.Sello.Estado != SelloValido || v.Sello.Algoritmo != "SHA256withRSA" || (no != "" && v.Sello.NoCertificado != no) {
			t.Errorf("%s: %+v", archivo, v.Sello)
		}
		if v.TFD != nil {
			t.Errorf("%s: sin proveedor de certificados no debía verificar el TFD", archivo)
		}
		if v.CadenaOriginalTFD == "" {
			t.Errorf("%s: falta la cadena del TFD", archivo)
		}
	}

	// En cfdi33-valid, CfdiUtils dejó a propósito un timbre que no
	// corresponde (lo dice el comentario del XML): con el certificado del
	// SAT que lo firmó, el SelloSAT no verifica.
	b, _ := os.ReadFile("testdata/publicos/eclipxe13-CfdiUtils/tests__assets__cfdi33-valid.xml")
	sat := cargarLlave(t, "20001000000300022779")
	v := Verificar(b, CertificadosEn(sat.cert))
	if v.TFD.Estado != SelloInvalido || v.TFD.Motivo != MotivoSelloNoCoincide {
		t.Errorf("TFD de cfdi33-valid: %+v", v.TFD)
	}
}

func TestVerificarNegativos(t *testing.T) {
	emisor, pac := cargarLlave(t, "EKU9003173C9"), cargarLlave(t, "CSD01_AAA010101AAA")
	bueno := string(leerSintetico(t, "sellados/ingreso40-pue.xml"))
	certs := CertificadosEn(pac.cert)
	selloBueno := regexp.MustCompile(` Sello="([^"]*)"`).FindStringSubmatch(bueno)[1]
	otroCert := base64.StdEncoding.EncodeToString(pac.der)

	casos := []struct {
		nombre     string
		xml        string
		estado     string
		motivo     string
		tfdEstado  string
		tfdMotivo  string
		sinCadenas bool
	}{
		{"sello alterado", strings.Replace(bueno, ` Sello="`+selloBueno[:10], ` Sello="`+invertirCaso(selloBueno[:10]), 1),
			SelloInvalido, MotivoSelloNoCoincide, SelloInvalido, MotivoSelloCFDDistinto, false},
		{"Total modificado tras firmar", strings.Replace(bueno, `Total="1640.00"`, `Total="1641.00"`, 1),
			SelloInvalido, MotivoSelloNoCoincide, SelloValido, "", false},
		{"espacio de más en un atributo: normalize-space lo absorbe", strings.Replace(bueno, `Total="1640.00"`, `Total=" 1640.00 "`, 1),
			SelloValido, "", SelloValido, "", false},
		{"NoCertificado que no es del certificado", strings.Replace(bueno, `NoCertificado="30001000000500003416"`, `NoCertificado="30001000000500003417"`, 1),
			SelloInvalido, MotivoSelloNoCoincide, SelloValido, "", false},
		{"certificado de otro con la firma del emisor", conAttr(t, bueno, "Certificado", otroCert),
			SelloInvalido, MotivoSelloNoCoincide, SelloValido, "", false},
		{"certificado ilegible", conAttr(t, bueno, "Certificado", "Q0VSVElGSUNBRE8gREUgUFJVRUJB"),
			SelloNoVerificable, MotivoCertificadoIlegible, SelloValido, "", false},
		{"certificado que no es base64", conAttr(t, bueno, "Certificado", "@@@"),
			SelloNoVerificable, MotivoCertificadoIlegible, SelloValido, "", false},
		{"sin certificado", conAttr(t, bueno, "Certificado", " "),
			SelloNoVerificable, MotivoSinCertificado, SelloValido, "", false},
		{"sello que no es base64", conAttr(t, bueno, "Sello", "no*es*base64"),
			SelloInvalido, MotivoSelloIlegible, SelloInvalido, MotivoSelloCFDDistinto, false},
		{"sin sello", conAttr(t, bueno, "Sello", ""),
			SelloInvalido, MotivoSinSello, SelloInvalido, MotivoSelloCFDDistinto, false},
		{"SelloSAT alterado", strings.Replace(bueno, `SelloSAT="`, `SelloSAT="AAAA`, 1),
			SelloValido, "", SelloInvalido, MotivoSelloNoCoincide, false},
		{"timbre modificado", strings.Replace(bueno, `FechaTimbrado="2026-01-15T10:31:12"`, `FechaTimbrado="2026-01-15T10:31:13"`, 1),
			SelloValido, "", SelloInvalido, MotivoSelloNoCoincide, false},
		{"NoCertificadoSAT desconocido", strings.Replace(bueno, `NoCertificadoSAT="30001000000300023708"`, `NoCertificadoSAT="99999999999999999999"`, 1),
			SelloValido, "", SelloNoVerificable, MotivoCertificadoSATFaltante, false},
		{"complemento sin cadena", strings.Replace(bueno, "</cfdi:Complemento>", `<x:Raro xmlns:x="urn:raro"/></cfdi:Complemento>`, 1),
			SelloNoVerificable, MotivoComplementoSinCadena, SelloValido, "", true},
		{"XML roto", bueno[:200], SelloNoVerificable, MotivoXMLIlegible, SelloNoVerificable, MotivoXMLIlegible, true},
		{"certificado con saltos de línea", conAttr(t, bueno, "Certificado", partir(base64.StdEncoding.EncodeToString(emisor.der), 76)),
			SelloValido, "", SelloValido, "", false},
		{"certificado PEM en base64", conAttr(t, bueno, "Certificado", base64.StdEncoding.EncodeToString(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: emisor.der}))),
			SelloValido, "", SelloValido, "", false},
	}
	for _, c := range casos {
		v := Verificar([]byte(c.xml), certs)
		if v.Sello.Estado != c.estado || v.Sello.Motivo != c.motivo {
			t.Errorf("%s: sello %s %s (%s), quiero %s %s", c.nombre, v.Sello.Estado, v.Sello.Motivo, v.Sello.Mensaje, c.estado, c.motivo)
		}
		if v.TFD == nil || v.TFD.Estado != c.tfdEstado || v.TFD.Motivo != c.tfdMotivo {
			t.Errorf("%s: TFD %+v, quiero %s %s", c.nombre, v.TFD, c.tfdEstado, c.tfdMotivo)
		}
		if c.estado != SelloValido && c.estado != SelloInvalido && v.Sello.Mensaje == "" {
			t.Errorf("%s: no verificable sin mensaje", c.nombre)
		}
		if c.sinCadenas != (v.CadenaOriginal == "") {
			t.Errorf("%s: CadenaOriginal = %q", c.nombre, v.CadenaOriginal)
		}
	}

	// Firma buena sobre un NoCertificado que no es el del certificado: la
	// cadena coincide, el documento miente.
	xml := conAttr(t, string(leerSintetico(t, "ingreso40-pue.xml")), "NoCertificado", "30001000000500009999")
	xml = conAttr(t, xml, "Certificado", base64.StdEncoding.EncodeToString(emisor.der))
	cadena, err := CadenaOriginal([]byte(xml))
	if err != nil {
		t.Fatal(err)
	}
	xml = conAttr(t, xml, "Sello", emisor.firmar(t, crypto.SHA256, cadena))
	if v := Verificar([]byte(xml), nil); v.Sello.Estado != SelloInvalido || v.Sello.Motivo != MotivoNoCertificadoDistinto {
		t.Errorf("NoCertificado distinto: %+v", v.Sello)
	}

	// Un CFDI 3.3 firmado con SHA-1 no vale: SHA-1 solo se acepta en 3.2.
	xml33 := sellar(t, string(leerSintetico(t, "ingreso33.xml")), emisor, crypto.SHA1, nil)
	if v := Verificar([]byte(xml33), nil); v.Sello.Estado != SelloInvalido || v.Sello.Motivo != MotivoSelloNoCoincide {
		t.Errorf("3.3 con SHA-1: %+v", v.Sello)
	}
	// El mismo con SHA-256 sí.
	xml33 = sellar(t, string(leerSintetico(t, "ingreso33.xml")), emisor, crypto.SHA256, nil)
	if v := Verificar([]byte(xml33), nil); v.Sello.Estado != SelloValido {
		t.Errorf("3.3 con SHA-256: %+v", v.Sello)
	}
	// 3.2 acepta los dos.
	xml32 := sellar(t, string(leerSintetico(t, "cadena/domicilios32.xml")), emisor, crypto.SHA256, &pac)
	if v := Verificar([]byte(xml32), certs); v.Sello.Algoritmo != "SHA256withRSA" || v.TFD.Algoritmo != "SHA256withRSA" || v.TFD.Estado != SelloValido {
		t.Errorf("3.2 con SHA-256: %+v %+v", v.Sello, v.TFD)
	}
}

// partir mete saltos de línea cada n caracteres, como hacen algunos emisores.
func partir(s string, n int) string {
	var b strings.Builder
	for i := 0; i < len(s); i += n {
		b.WriteString(s[i:min(i+n, len(s))])
		b.WriteString("\n")
	}
	return b.String()
}

func invertirCaso(s string) string {
	return strings.Map(func(r rune) rune {
		switch {
		case r >= 'a' && r <= 'z':
			return r - 32
		case r >= 'A' && r <= 'Z':
			return r + 32
		}
		return r
	}, s)
}

func TestVerificarProveedor(t *testing.T) {
	pac := cargarLlave(t, "CSD01_AAA010101AAA")
	bueno := leerSintetico(t, "sellados/ingreso40-pue.xml")
	otro := cargarLlave(t, "EKU9003173C9")
	casos := map[string]struct {
		certs  CertificadosSAT
		motivo string
	}{
		"error del proveedor":   {func(string) (*x509.Certificate, error) { return nil, errors.New("sin red") }, MotivoCertificadoSATFaltante},
		"nil sin error":         {func(string) (*x509.Certificate, error) { return nil, nil }, MotivoCertificadoSATFaltante},
		"otro certificado":      {func(string) (*x509.Certificate, error) { return otro.cert, nil }, MotivoNoCertificadoSATDistinto},
		"el que es":             {CertificadosEn(pac.cert), ""},
		"el que es entre otros": {CertificadosEn(otro.cert, pac.cert), ""},
	}
	for nombre, c := range casos {
		v := Verificar(bueno, c.certs)
		if v.TFD.Motivo != c.motivo || (c.motivo == "" && v.TFD.Estado != SelloValido) {
			t.Errorf("%s: %+v", nombre, v.TFD)
		}
	}
	_, err := CertificadosEn()("123")
	if !errors.Is(err, ErrCertificadoNoEncontrado) {
		t.Errorf("CertificadosEn vacío: %v", err)
	}
	// Sin timbre.
	v := Verificar(leerSintetico(t, "sin-tfd.xml"), CertificadosEn(pac.cert))
	if v.TFD.Motivo != MotivoSinTimbre {
		t.Errorf("sin TFD: %+v", v.TFD)
	}
	// Timbre de versión desconocida.
	v = Verificar([]byte(strings.Replace(string(bueno), `Version="1.1"`, `Version="9.9"`, 1)), CertificadosEn(pac.cert))
	if v.TFD.Motivo != MotivoVersionNoSoportada || v.CadenaOriginalTFD != "" {
		t.Errorf("TFD 9.9: %+v", v.TFD)
	}
}

// Un documento que Parse lee corrigiendo su codificación: si el sello
// verifica con esa lectura, es válido (con aviso); si no, no se sabe.
func TestVerificarCodificacionCorregida(t *testing.T) {
	emisor := cargarLlave(t, "EKU9003173C9")
	xml := sellar(t, string(leerSintetico(t, "ingreso40-pue.xml")), emisor, crypto.SHA256, nil)
	if !strings.Contains(xml, "retención") {
		t.Fatal("el fixture ya no trae caracteres no ASCII")
	}
	// Declara latin1 pero los bytes son UTF-8: la lectura corregida es la
	// que se firmó.
	miente := strings.Replace(xml, `encoding="UTF-8"`, `encoding="ISO-8859-1"`, 1)
	if _, err := CadenaOriginal([]byte(miente)); !esMotivo(err, MotivoXMLIlegible) {
		t.Fatalf("CadenaOriginal debía negarse: %v", err)
	}
	v := Verificar([]byte(miente), nil)
	if v.Sello.Estado != SelloValido || !strings.Contains(v.Sello.Mensaje, "codificación") {
		t.Errorf("encoding que miente, sello bueno: %+v", v.Sello)
	}
	// Si además se tocó un valor, no se puede saber si fue la corrección.
	tocado := strings.Replace(miente, `Total="1640.00"`, `Total="1.00"`, 1)
	v = Verificar([]byte(tocado), CertificadosEn(emisor.cert))
	if v.Sello.Estado != SelloNoVerificable || v.Sello.Motivo != MotivoXMLIlegible || v.CadenaOriginal != "" {
		t.Errorf("encoding que miente, sello malo: %+v", v.Sello)
	}
	if v.TFD.Motivo != MotivoXMLIlegible {
		t.Errorf("TFD con encoding corregido: %+v", v.TFD)
	}
}

func TestVerificarLlaveNoRSA(t *testing.T) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	tpl := &x509.Certificate{SerialNumber: big.NewInt(1), Subject: pkix.Name{CommonName: "prueba"},
		NotBefore: time.Now(), NotAfter: time.Now().Add(time.Hour)}
	der, err := x509.CreateCertificate(rand.Reader, tpl, tpl, &priv.PublicKey, priv)
	if err != nil {
		t.Fatal(err)
	}
	bueno := string(leerSintetico(t, "sellados/ingreso40-pue.xml"))
	v := Verificar([]byte(conAttr(t, bueno, "Certificado", base64.StdEncoding.EncodeToString(der))), nil)
	if v.Sello.Estado != SelloNoVerificable || v.Sello.Motivo != MotivoLlaveNoRSA {
		t.Errorf("certificado ECDSA: %+v", v.Sello)
	}
}

// La clave verificacion_sello del JSON valida contra el esquema en todos
// los fixtures, y el enum de motivos del esquema es la lista Motivos.
func TestSchemaVerificacionSello(t *testing.T) {
	s := cargarEsquema(t)
	vistos := map[string]bool{}
	for _, p := range fixturesCadena(t) {
		if esRetenciones(t, p) {
			continue
		}
		b, _ := os.ReadFile(p)
		c, err := ParseBytes(b)
		if err != nil {
			t.Fatal(err)
		}
		d := c.Documento(OpcionesJSON{})
		r := Verificar(b, nil).Sello
		d.VerificacionSello = &r
		vistos[r.Estado] = true
		out, _ := json.Marshal(d)
		for _, e := range validarJSON(t, s, out) {
			t.Errorf("%s: %s", p, e)
		}
	}
	if len(vistos) != 3 {
		t.Errorf("los fixtures no cubren los tres estados: %v", vistos)
	}
	def := s["$defs"].(map[string]any)["resultado_sello"].(map[string]any)["properties"].(map[string]any)
	var enum []string
	for _, e := range def["motivo"].(map[string]any)["enum"].([]any) {
		enum = append(enum, e.(string))
	}
	if strings.Join(enum, ",") != strings.Join(Motivos, ",") {
		t.Errorf("enum de motivos del esquema ≠ Motivos:\n%v\n%v", enum, Motivos)
	}
}

func TestNoCertificado(t *testing.T) {
	k := cargarLlave(t, "EKU9003173C9")
	igual(t, "EKU", NoCertificado(k.cert), "30001000000500003416")
	c := &x509.Certificate{SerialNumber: big.NewInt(0x01ff)}
	igual(t, "serie no ASCII", NoCertificado(c), "01FF")
}
