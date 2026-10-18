package cfdi

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"math/big"
	"slices"
	"strings"
	"sync"
	"testing"
	"time"
)

var (
	llavePruebaOnce sync.Once
	llavePrueba     *rsa.PrivateKey
)

// certPrueba genera en la prueba un certificado (no real) con la serie, el
// x500UniqueIdentifier y la vigencia dados, autofirmado con una llave RSA
// que se genera una sola vez.
func certPrueba(t *testing.T, serie, rfcCert string, desde, hasta time.Time) llave {
	t.Helper()
	llavePruebaOnce.Do(func() {
		var err error
		if llavePrueba, err = rsa.GenerateKey(rand.Reader, 2048); err != nil {
			panic(err)
		}
	})
	sujeto := pkix.Name{CommonName: "CERTIFICADO DE PRUEBA"}
	if rfcCert != "" {
		sujeto.ExtraNames = []pkix.AttributeTypeAndValue{{Type: asn1.ObjectIdentifier{2, 5, 4, 45}, Value: rfcCert}}
	}
	tpl := &x509.Certificate{SerialNumber: new(big.Int).SetBytes([]byte(serie)), Subject: sujeto, NotBefore: desde, NotAfter: hasta}
	der, err := x509.CreateCertificate(rand.Reader, tpl, tpl, &llavePrueba.PublicKey, llavePrueba)
	if err != nil {
		t.Fatal(err)
	}
	c, err := x509.ParseCertificate(der)
	if err != nil {
		t.Fatal(err)
	}
	return llave{cert: c, der: der, priv: llavePrueba}
}

var (
	vigenteDesde = time.Date(2026, 1, 10, 12, 0, 0, 0, time.UTC)
	vigenteHasta = time.Date(2026, 7, 10, 12, 0, 0, 0, time.UTC)
)

func chequeo(r ResultadoChequeo) string {
	return strings.TrimSpace(r.Estado + " " + r.Motivo)
}

// La Fecha del CFDI es hora civil sin zona: en México, de UTC-5 a UTC-8. En
// la vigencia con todos los husos, cumple; fuera con todos, no; si depende,
// ambigua. Con zona, el instante exacto.
func TestChequeoVigenciaBordes(t *testing.T) {
	cert := certPrueba(t, "20001000000300099999", "EKU9003173C9 / VADA800927DJ3", vigenteDesde, vigenteHasta).cert
	casos := []struct{ fecha, quiero string }{
		{"2026-03-01T10:00:00", "cumple"},
		// NotBefore 12:00 UTC.
		{"2026-01-10T07:00:00", "cumple"},                               // +5h = 12:00, el primer instante posible
		{"2026-01-10T06:59:59", "no_verificable VIGENCIA_AMBIGUA"},      // +5h antes, +8h después
		{"2026-01-10T04:00:00", "no_verificable VIGENCIA_AMBIGUA"},      // +8h = 12:00: con UTC-8 ya vale
		{"2026-01-10T03:59:59", "no_cumple CERTIFICADO_NO_VIGENTE_AUN"}, // +8h = 11:59:59
		// NotAfter 12:00 UTC.
		{"2026-07-10T04:00:00", "cumple"},                          // +8h = 12:00, el último instante posible
		{"2026-07-10T04:00:01", "no_verificable VIGENCIA_AMBIGUA"}, // +8h ya pasó, +5h no
		{"2026-07-10T07:00:00", "no_verificable VIGENCIA_AMBIGUA"}, // +5h = 12:00: con UTC-5 todavía vale
		{"2026-07-10T07:00:01", "no_cumple CERTIFICADO_VENCIDO"},   // +5h = 12:00:01
		// Con zona (FechaExp de Retenciones 1.0): el instante exacto.
		{"2026-07-10T06:00:00-06:00", "cumple"},
		{"2026-07-10T06:00:01-06:00", "no_cumple CERTIFICADO_VENCIDO"},
		{"2026-01-10T06:00:00-06:00", "cumple"},
		{"2026-01-10T05:59:59-06:00", "no_cumple CERTIFICADO_NO_VIGENTE_AUN"},
		{"2026-01-10T12:00:00Z", "cumple"},
		// Fuera de formato pero legible: la hora civil.
		{" 2026-03-01T10:00:00 ", "cumple"},
		{"ayer", "no_verificable FECHA_ILEGIBLE"},
		{"", "no_verificable FECHA_ILEGIBLE"},
	}
	for _, c := range casos {
		if got := chequeo(chequearVigencia(cert, c.fecha, c.fecha != "")); got != c.quiero {
			t.Errorf("%q: %s, quiero %s", c.fecha, got, c.quiero)
		}
	}
	if got := chequeo(chequearVigencia(cert, "", false)); got != "no_verificable FECHA_ILEGIBLE" {
		t.Errorf("sin fecha: %s", got)
	}
}

func TestChequeoRFC(t *testing.T) {
	moral := certPrueba(t, "20001000000300099999", "EKU9003173C9 / VADA800927DJ3", vigenteDesde, vigenteHasta).cert
	fisica := certPrueba(t, "20001000000300099999", "XOJI740919U48", vigenteDesde, vigenteHasta).cert
	sinRFC := certPrueba(t, "20001000000300099999", "", vigenteDesde, vigenteHasta).cert
	// Servicio gratuito de facturación del SAT: titular el SAT, representante
	// el contribuyente.
	delSAT := certPrueba(t, "20001000000300099999", "SAT970701NN3 / XOJI740919U48", vigenteDesde, vigenteHasta).cert
	satRepresentante := certPrueba(t, "20001000000300099999", "AAA010101AAA / SAT970701NN3", vigenteDesde, vigenteHasta).cert
	casos := []struct {
		nombre string
		cert   *x509.Certificate
		emisor string
		quiero string
	}{
		{"certificado del SAT", delSAT, "EKU9003173C9", "cumple CERTIFICADO_DEL_SAT"},
		{"certificado del SAT, emisor en minúsculas", delSAT, " xoji740919u48 ", "cumple CERTIFICADO_DEL_SAT"},
		{"el SAT emisor con su certificado", delSAT, "SAT970701NN3", "cumple"},
		{"el SAT como representante no cuenta", satRepresentante, "EKU9003173C9", "no_cumple RFC_CERTIFICADO_DISTINTO"},
		{"el SAT como representante y emisor", satRepresentante, "SAT970701NN3", "no_cumple RFC_CERTIFICADO_DISTINTO"},
		{"certificado del SAT sin RFC del emisor", delSAT, "", "no_verificable SIN_RFC_EMISOR"},
		{"persona moral: cuenta el titular", moral, "EKU9003173C9", "cumple"},
		{"el RFC del representante no es el emisor", moral, "VADA800927DJ3", "no_cumple RFC_CERTIFICADO_DISTINTO"},
		{"otro emisor", moral, "AAA010101AAA", "no_cumple RFC_CERTIFICADO_DISTINTO"},
		{"emisor en minúsculas y con espacios", moral, " eku9003173c9 ", "cumple"},
		{"persona física", fisica, "XOJI740919U48", "cumple"},
		{"persona física, otra", fisica, "XOJI740919U49", "no_cumple RFC_CERTIFICADO_DISTINTO"},
		{"certificado sin x500UniqueIdentifier", sinRFC, "EKU9003173C9", "no_verificable CERTIFICADO_SIN_RFC"},
		{"sin RFC del emisor", moral, "", "no_verificable SIN_RFC_EMISOR"},
	}
	for _, c := range casos {
		if got := chequeo(chequearRFC(c.cert, c.emisor, "")); got != c.quiero {
			t.Errorf("%s: %s, quiero %s", c.nombre, got, c.quiero)
		}
	}
	// RfcProvCertif del SAT solo cambia el mensaje; no es requisito, y no
	// rescata un certificado de otro tercero.
	sinProv := chequearRFC(delSAT, "EKU9003173C9", "")
	conProv := chequearRFC(delSAT, "EKU9003173C9", "SAT970701NN3")
	if chequeo(conProv) != "cumple CERTIFICADO_DEL_SAT" || !strings.Contains(conProv.Mensaje, "RfcProvCertif") ||
		strings.Contains(sinProv.Mensaje, "RfcProvCertif") || !strings.Contains(sinProv.Mensaje, "EKU9003173C9") {
		t.Errorf("mensajes: %+v / %+v", sinProv, conProv)
	}
	if got := chequeo(chequearRFC(moral, "AAA010101AAA", "SAT970701NN3")); got != "no_cumple RFC_CERTIFICADO_DISTINTO" {
		t.Errorf("otro tercero con timbre del SAT: %s", got)
	}
	// Los CSD de prueba del SAT.
	igual(t, "RFC del CSD EKU", RFCDelCertificado(cargarLlave(t, "EKU9003173C9").cert), "EKU9003173C9")
	igual(t, "RFC del CSD AAA", RFCDelCertificado(cargarLlave(t, "CSD01_AAA010101AAA").cert), "AAA010101AAA")
}

func TestChequeoNoCertificado(t *testing.T) {
	cert := certPrueba(t, "20001000000300099999", "EKU9003173C9", vigenteDesde, vigenteHasta).cert
	casos := []struct {
		valor    string
		presente bool
		quiero   string
	}{
		{"20001000000300099999", true, "cumple"},
		{" 20001000000300099999 ", true, "cumple"},
		{"20001000000300099998", true, "no_cumple NOCERTIFICADO_DISTINTO"},
		{"", true, "no_cumple SIN_NOCERTIFICADO"},
		{"", false, "no_cumple SIN_NOCERTIFICADO"},
	}
	for _, c := range casos {
		if got := chequeo(chequearNoCertificado(cert, "NoCertificado", c.valor, c.presente)); got != c.quiero {
			t.Errorf("%q: %s, quiero %s", c.valor, got, c.quiero)
		}
	}
}

// De punta a punta: los chequeos salen en Verificar y VerificarRetenciones
// y no cambian el veredicto del sello.
func TestVerificarChequeosCertificado(t *testing.T) {
	todos := func(c *ChequeosCertificado) string {
		if c == nil {
			return "nil"
		}
		return chequeo(c.Vigencia) + " | " + chequeo(c.Rfc) + " | " + chequeo(c.NoCertificado)
	}
	// Los sellados con el CSD de prueba (vigente de 2023-05-18 a 2027-05-18).
	// La Retenciones 1.0 es de 2020: el sello verifica, pero el certificado
	// todavía no existía.
	for f, quiero := range map[string]string{
		"sellados/ingreso40-pue.xml":     "cumple | cumple | cumple",
		"sellados/ret20-plataformas.xml": "cumple | cumple | cumple",
		"sellados/ret10-extranjero.xml":  "no_cumple CERTIFICADO_NO_VIGENTE_AUN | cumple | cumple",
	} {
		b := leerSintetico(t, f)
		verificar := Verificar
		if Detectar(b) == TipoRetenciones {
			verificar = VerificarRetenciones
		}
		v := verificar(b, nil)
		if got := todos(v.Certificado); v.Sello.Estado != SelloValido || got != quiero {
			t.Errorf("%s: sello %s, certificado %s", f, v.Sello.Estado, got)
		}
	}

	// Un CFDI sellado con un certificado ya vencido y de otro RFC: el sello
	// verifica igual.
	vencido := certPrueba(t, "20001000000300099999", "XOJI740919U48", time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC))
	xml := sellar(t, string(leerSintetico(t, "ingreso40-pue.xml")), vencido, crypto.SHA256, nil)
	v := Verificar([]byte(xml), nil)
	if v.Sello.Estado != SelloValido || todos(v.Certificado) != "no_cumple CERTIFICADO_VENCIDO | no_cumple RFC_CERTIFICADO_DISTINTO | cumple" {
		t.Errorf("certificado vencido: sello %+v, certificado %s", v.Sello, todos(v.Certificado))
	}
	// Sellado con un certificado del SAT a nombre del emisor: el RFC cumple
	// con CERTIFICADO_DEL_SAT, y el mensaje dice si el timbre también es del
	// SAT (el timbre no entra en la cadena del emisor: el sello sigue igual).
	delSAT := certPrueba(t, "20001000000300099999", "SAT970701NN3 / XOJI740919U48", vigenteDesde, vigenteHasta)
	xmlSAT := sellar(t, string(leerSintetico(t, "ingreso40-pue.xml")), delSAT, crypto.SHA256, nil)
	if !strings.Contains(xmlSAT, `RfcProvCertif="AAA010101AAA"`) {
		t.Fatal("el fixture no trae el RfcProvCertif esperado")
	}
	for prov, conMensaje := range map[string]bool{"AAA010101AAA": false, "SAT970701NN3": true} {
		v := Verificar([]byte(strings.Replace(xmlSAT, `RfcProvCertif="AAA010101AAA"`, `RfcProvCertif="`+prov+`"`, 1)), nil)
		if v.Sello.Estado != SelloValido || chequeo(v.Certificado.Rfc) != "cumple CERTIFICADO_DEL_SAT" ||
			strings.Contains(v.Certificado.Rfc.Mensaje, "RfcProvCertif") != conMensaje {
			t.Errorf("certificado del SAT, RfcProvCertif %s: sello %s, %+v", prov, v.Sello.Estado, v.Certificado.Rfc)
		}
	}

	// NoCertificado de otro: el sello ya lo dice; el chequeo también.
	otro := conAttr(t, xml, "NoCertificado", "20001000000300099998")
	v = Verificar([]byte(otro), nil)
	if v.Sello.Motivo != MotivoSelloNoCoincide || chequeo(v.Certificado.NoCertificado) != "no_cumple NOCERTIFICADO_DISTINTO" {
		t.Errorf("NoCertificado de otro: %+v %s", v.Sello, todos(v.Certificado))
	}

	// Retenciones 1.0: FechaExp con zona, instante exacto en el borde.
	ret10 := string(leerSintetico(t, "retenciones/ret10-extranjero.xml"))
	borde := certPrueba(t, "20001000000300099999", "EKU9003173C9 / VADA800927DJ3", time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2020, 7, 1, 0, 0, 0, 0, time.UTC))
	for fecha, quiero := range map[string]string{
		"2020-06-30T18:00:00-06:00": "cumple",                        // = NotAfter
		"2020-06-30T18:00:01-06:00": "no_cumple CERTIFICADO_VENCIDO", // un segundo después
	} {
		x := sellarRetenciones(t, strings.Replace(ret10, "2020-06-30T18:00:00-06:00", fecha, 1), borde, crypto.SHA1, nil)
		v := VerificarRetenciones([]byte(x), nil)
		if v.Sello.Estado != SelloValido || chequeo(v.Certificado.Vigencia) != quiero || chequeo(v.Certificado.Rfc) != "cumple" {
			t.Errorf("Retenciones 1.0 %s: sello %s, %s", fecha, v.Sello.Estado, todos(v.Certificado))
		}
	}

	// Sin certificado, ilegible o XML roto: no verificables.
	bueno := string(leerSintetico(t, "sellados/ingreso40-pue.xml"))
	for nombre, c := range map[string]struct{ xml, motivo string }{
		"sin certificado":      {conAttr(t, bueno, "Certificado", ""), MotivoSinCertificado},
		"certificado ilegible": {conAttr(t, bueno, "Certificado", "QUJD"), MotivoCertificadoIlegible},
		"XML roto":             {bueno[:200], MotivoXMLIlegible},
	} {
		v := Verificar([]byte(c.xml), nil)
		r := "no_verificable " + c.motivo
		if todos(v.Certificado) != r+" | "+r+" | "+r {
			t.Errorf("%s: %s", nombre, todos(v.Certificado))
		}
	}
	// Sin Emisor ni Fecha.
	sinEmisor := strings.Replace(conAttr(t, bueno, "Fecha", ""), `Rfc="EKU9003173C9"`, `X="1"`, 1)
	if v := Verificar([]byte(sinEmisor), nil); todos(v.Certificado) != "no_verificable FECHA_ILEGIBLE | no_verificable SIN_RFC_EMISOR | cumple" {
		t.Errorf("sin fecha ni RFC: %s", todos(v.Certificado))
	}
	// La lista de motivos tiene todos los que salen.
	for _, m := range []string{MotivoCertificadoNoVigenteAun, MotivoCertificadoVencido, MotivoVigenciaAmbigua, MotivoFechaIlegible,
		MotivoRFCCertificadoDistinto, MotivoCertificadoSinRFC, MotivoSinRFCEmisor, MotivoSinNoCertificado, MotivoCertificadoDelSAT} {
		if !slices.Contains(MotivosCertificado, m) {
			t.Errorf("MotivosCertificado no tiene %s", m)
		}
	}
}
