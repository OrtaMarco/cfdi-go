package cfdi

import (
	"crypto"
	"encoding/base64"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// sellarRetenciones es sellar para Retenciones: NoCertificado (NumCert en
// 1.0), Certificado (Cert) y Sello con la cadena de
// CadenaOriginalRetenciones, que las pruebas de cadena comparan contra
// libxslt; y el timbre (1.1 en 2.0, 1.0 en minúsculas en 1.0) si pac no es
// nil.
func sellarRetenciones(t *testing.T, xml string, emisor llave, h crypto.Hash, pac *llave) string {
	t.Helper()
	cert, noCert := "Certificado", "NoCertificado"
	selloCFD, noCertSAT, selloSAT := "SelloCFD", "NoCertificadoSAT", "SelloSAT"
	if strings.Contains(xml, `"`+NSRetenciones10+`"`) {
		cert, noCert = "Cert", "NumCert"
		selloCFD, noCertSAT, selloSAT = "selloCFD", "noCertificadoSAT", "selloSAT"
	}
	xml = conAttr(t, xml, noCert, NoCertificado(emisor.cert))
	xml = conAttr(t, xml, cert, base64.StdEncoding.EncodeToString(emisor.der))
	cadena, err := CadenaOriginalRetenciones([]byte(xml))
	if err != nil {
		t.Fatal(err)
	}
	s := emisor.firmar(t, h, cadena)
	xml = conAttr(t, xml, "Sello", s)
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

// Retenciones selladas de punta a punta con los CSD de pruebas (emisor
// EKU9003173C9, PAC CSD01_AAA010101AAA); se regeneran con -update como
// TestSellados. La 1.0 con SHA-1, como los ejemplos del SAT.
var selladosRetenciones = []struct {
	origen string
	hash   crypto.Hash
}{
	{"retenciones/ret20-plataformas.xml", crypto.SHA256},
	{"retenciones/ret10-extranjero.xml", crypto.SHA1},
}

func TestSelladosRetenciones(t *testing.T) {
	emisor, pac := cargarLlave(t, "EKU9003173C9"), cargarLlave(t, "CSD01_AAA010101AAA")
	certs := CertificadosEn(pac.cert)
	for _, s := range selladosRetenciones {
		destino := filepath.Join("testdata", "sinteticos", "sellados", filepath.Base(s.origen))
		if *update {
			xml := sellarRetenciones(t, string(leerSintetico(t, s.origen)), emisor, s.hash, &pac)
			if err := os.WriteFile(destino, []byte(xml), 0o644); err != nil {
				t.Fatal(err)
			}
		}
		b, err := os.ReadFile(destino)
		if err != nil {
			t.Fatalf("%v (corre con -update)", err)
		}
		v := VerificarRetenciones(b, certs)
		algoritmo := nombreAlgoritmo(s.hash)
		if v.Sello.Estado != SelloValido || v.Sello.Algoritmo != algoritmo || v.Sello.NoCertificado != "30001000000500003416" {
			t.Errorf("%s: sello %+v", destino, v.Sello)
		}
		if v.TFD == nil || v.TFD.Estado != SelloValido || v.TFD.Algoritmo != algoritmo || v.TFD.NoCertificado != "30001000000300023708" {
			t.Errorf("%s: TFD %+v", destino, v.TFD)
		}
		cad, _ := CadenaOriginalRetenciones(b)
		cadT, _ := CadenaOriginalTFD(b)
		if v.CadenaOriginal != cad || v.CadenaOriginalTFD != cadT || cad == "" || cadT == "" {
			t.Errorf("%s: VerificarRetenciones no devolvió las cadenas", destino)
		}
		// Verificar (el de CFDI) no la reconoce.
		if w := Verificar(b, nil); w.Sello.Motivo != MotivoNoEsComprobante {
			t.Errorf("%s: Verificar de una Retenciones: %+v", destino, w.Sello)
		}
	}
}

// Sellos genuinos de Retenciones que firmaron otros (el SAT en los ejemplos
// de satcfdi, CfdiUtils, Lunasoft) sobre la cadena de sus XSLT.
func TestSellosPublicosRetenciones(t *testing.T) {
	sat := "SAT-CFDI-python-satcfdi/tests__cfdi_ejemplos__retenciones10__retenciones10-"
	for archivo, algoritmo := range map[string]string{
		sat + "intereses.xml":                                                       "SHA1withRSA",
		sat + "dividendos.xml":                                                      "SHA1withRSA",
		sat + "operaciones-con-derivados.xml":                                       "SHA1withRSA",
		sat + "pagos-a-extranjeros.xml":                                             "SHA1withRSA",
		sat + "planes-de-retiro.xml":                                                "SHA1withRSA",
		sat + "sector-financiero.xml":                                               "SHA1withRSA",
		sat + "arrendamiento-en-fideicomiso.xml":                                    "SHA1withRSA",
		sat + "enajenacion_de_acciones.xml":                                         "SHA1withRSA",
		sat + "intereses-hipotecarios.xml":                                          "SHA1withRSA",
		sat + "premios.xml":                                                         "SHA1withRSA",
		sat + "ejemplo.xml":                                                         "SHA1withRSA",
		"eclipxe13-CfdiUtils/tests__assets__retenciones__retenciones10.xml":         "SHA1withRSA",
		"eclipxe13-CfdiUtils/tests__assets__retenciones__retenciones20.xml":         "SHA256withRSA",
		"lunasoft-sw-sdk-php/Test__Resources__retencion20_sellado.xml":              "SHA256withRSA",
		"nodecfdi-cfdi-validator/tests___files__retenciones__sample-before-tfd.xml": "SHA1withRSA",
	} {
		b, err := os.ReadFile(filepath.Join("testdata", "publicos", archivo))
		if err != nil {
			t.Fatal(err)
		}
		v := VerificarRetenciones(b, nil)
		if v.Sello.Estado != SelloValido || v.Sello.Algoritmo != algoritmo {
			t.Errorf("%s: %+v", archivo, v.Sello)
		}
	}
}

func TestVerificarRetencionesNegativos(t *testing.T) {
	bueno := string(leerSintetico(t, "sellados/ret20-plataformas.xml"))
	bueno10 := string(leerSintetico(t, "sellados/ret10-extranjero.xml"))
	emisor := cargarLlave(t, "EKU9003173C9")
	casos := []struct {
		nombre, xml, estado, motivo string
	}{
		{"monto modificado tras firmar", strings.Replace(bueno, `MontoRet="80.00"`, `MontoRet="81.00"`, 1), SelloInvalido, MotivoSelloNoCoincide},
		{"complemento modificado", strings.Replace(bueno, `NumServ="2"`, `NumServ="3"`, 1), SelloInvalido, MotivoSelloNoCoincide},
		{"receptor modificado (1.0)", strings.Replace(bueno10, `NumRegIdTrib="998877665544"`, `NumRegIdTrib="998877665545"`, 1), SelloInvalido, MotivoSelloNoCoincide},
		{"NumCert de otro (1.0)", strings.Replace(bueno10, `NumCert="30001000000500003416"`, `NumCert="30001000000500003417"`, 1), SelloInvalido, MotivoSelloNoCoincide},
		{"sin Cert (1.0)", conAttr(t, bueno10, "Cert", ""), SelloNoVerificable, MotivoSinCertificado},
		{"complemento desconocido", strings.Replace(bueno, "</retenciones:Complemento>", `<x:Raro xmlns:x="urn:raro"/></retenciones:Complemento>`, 1), SelloNoVerificable, MotivoComplementoSinCadena},
		{"un CFDI", string(leerSintetico(t, "sellados/ingreso40-pue.xml")), SelloNoVerificable, MotivoNoEsRetenciones},
		{"versión 3.0", strings.Replace(bueno, `Version="2.0"`, `Version="3.0"`, 1), SelloNoVerificable, MotivoVersionNoSoportada},
	}
	for _, c := range casos {
		v := VerificarRetenciones([]byte(c.xml), nil)
		if v.Sello.Estado != c.estado || v.Sello.Motivo != c.motivo {
			t.Errorf("%s: %s %s (%s), quiero %s %s", c.nombre, v.Sello.Estado, v.Sello.Motivo, v.Sello.Mensaje, c.estado, c.motivo)
		}
	}
	// 2.0 firmada con SHA-1 no vale; 1.0 con SHA-256 sí.
	x := sellarRetenciones(t, string(leerSintetico(t, "retenciones/ret20-plataformas.xml")), emisor, crypto.SHA1, nil)
	if v := VerificarRetenciones([]byte(x), nil); v.Sello.Estado != SelloInvalido || v.Sello.Motivo != MotivoSelloNoCoincide {
		t.Errorf("2.0 con SHA-1: %+v", v.Sello)
	}
	x = sellarRetenciones(t, string(leerSintetico(t, "retenciones/ret10-extranjero.xml")), emisor, crypto.SHA256, nil)
	if v := VerificarRetenciones([]byte(x), nil); v.Sello.Estado != SelloValido || v.Sello.Algoritmo != "SHA256withRSA" {
		t.Errorf("1.0 con SHA-256: %+v", v.Sello)
	}
}

func TestCadenaRetencionesErrores(t *testing.T) {
	base := string(leerSintetico(t, "retenciones/ret20-plataformas.xml"))
	casos := []struct {
		nombre, xml, motivo, nodo string
	}{
		{"un CFDI", string(leerSintetico(t, "ingreso40-pue.xml")), MotivoNoEsRetenciones, "{http://www.sat.gob.mx/cfd/4}Comprobante"},
		{"2.0 con Version 1.0", strings.Replace(base, `Version="2.0"`, `Version="1.0"`, 1), MotivoVersionNoSoportada, ""},
		{"complemento desconocido", strings.Replace(base, "</retenciones:Complemento>", `<x:Raro xmlns:x="urn:raro"/></retenciones:Complemento>`, 1), MotivoComplementoSinCadena, "{urn:raro}Raro"},
		{"XML roto", base[:200], MotivoXMLIlegible, ""},
		{"encoding que miente", strings.Replace(base, `encoding="UTF-8"`, `encoding="ISO-8859-1"`, 1), MotivoXMLIlegible, ""},
	}
	for _, c := range casos {
		_, err := CadenaOriginalRetenciones([]byte(c.xml))
		var ec *ErrorCadena
		if !errors.As(err, &ec) || ec.Motivo != c.motivo || ec.Nodo != c.nodo {
			t.Errorf("%s: %v, quiero %s %s", c.nombre, err, c.motivo, c.nodo)
		}
	}
	// CadenaOriginal (la de CFDI) sigue rechazando una Retenciones.
	if _, err := CadenaOriginal([]byte(base)); !esMotivo(err, MotivoNoEsComprobante) {
		t.Errorf("CadenaOriginal de una Retenciones: %v", err)
	}
	// La cadena del TFD de una Retenciones es la de siempre.
	got, err := CadenaOriginalTFD([]byte(base))
	if err != nil || got != "||1.1|0F1E2D3C-4B5A-4978-8695-A4B3C2D1E0F9|2026-01-15T10:31:12|AAA010101AAA||30001000000300023708||" {
		t.Errorf("TFD de una Retenciones: %q, %v", got, err)
	}
}
