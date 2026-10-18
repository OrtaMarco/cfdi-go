package cfdi

import (
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/OrtaMarco/cfdi-go/decimal"
	"github.com/OrtaMarco/cfdi-go/internal/genxsd"
)

// esquema_reglas.go y esquema_catalogos.go tienen que ser lo que genera
// internal/genxsd con los XSD del repo: nadie los edita a mano ni se quedan
// viejos.
func TestEsquemaAlDia(t *testing.T) {
	reglas, catalogos, err := genxsd.Generar("internal/genxsd/xsd/sitio_internet")
	if err != nil {
		t.Fatal(err)
	}
	for _, a := range []struct {
		ruta string
		b    []byte
	}{{"esquema_reglas.go", reglas}, {"esquema_catalogos.go", catalogos}} {
		want, err := os.ReadFile(a.ruta)
		if err != nil {
			t.Fatal(err)
		}
		if string(a.b) != string(want) {
			t.Errorf("%s no es lo que genera internal/genxsd: corre `go generate .`", a.ruta)
		}
	}
}

type manifiestoMutantes struct {
	Mutantes []struct {
		ID        string   `json:"id"`
		Mutacion  string   `json:"mutacion"`
		Veredicto string   `json:"veredicto"`
		Codigos   []string `json:"codigos"`
	} `json:"mutantes"`
}

// Cada mutante da el código de hallazgo de la regla que rompe (además del
// veredicto, que compara TestValidarEsquemaContraOraculo).
func TestValidarEsquemaMutantes(t *testing.T) {
	b, err := os.ReadFile("testdata/mutantes/MANIFIESTO.json")
	if err != nil {
		t.Fatal(err)
	}
	var m manifiestoMutantes
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatal(err)
	}
	if len(m.Mutantes) < 50 {
		t.Fatalf("solo %d mutantes", len(m.Mutantes))
	}
	for _, mu := range m.Mutantes {
		x, err := os.ReadFile(filepath.Join("testdata/mutantes", mu.ID+".xml"))
		if err != nil {
			t.Fatal(err)
		}
		r := ValidarEsquema(x)
		if r.Estado != mu.Veredicto {
			t.Errorf("%s (%s): estado %s, quiero %s: %+v", mu.ID, mu.Mutacion, r.Estado, mu.Veredicto, r.Hallazgos)
		}
		for _, c := range mu.Codigos {
			if !TieneHallazgo(r.Hallazgos, c) {
				t.Errorf("%s (%s): falta el hallazgo %s: %+v", mu.ID, mu.Mutacion, c, r.Hallazgos)
			}
		}
		// Nada más que lo esperado: el mutante tiene un solo error a propósito.
		for _, h := range r.Hallazgos {
			ok := false
			for _, c := range mu.Codigos {
				ok = ok || h.Codigo == c
			}
			if !ok {
				t.Errorf("%s (%s): hallazgo de más %+v", mu.ID, mu.Mutacion, h)
			}
		}
	}
}

func leerFixture(t *testing.T, p string) []byte {
	t.Helper()
	b, err := os.ReadFile(p)
	if err != nil {
		t.Fatal(err)
	}
	return b
}

func buscarHallazgo(hs []Hallazgo, codigo string) *Hallazgo {
	for i := range hs {
		if hs[i].Codigo == codigo {
			return &hs[i]
		}
	}
	return nil
}

// Ruta, Atributo, Valor y Regla de los hallazgos siguen el contrato de
// Advertencia.
func TestValidarEsquemaRutas(t *testing.T) {
	base := string(leerFixture(t, "testdata/sinteticos/ingreso40-pue.xml"))
	casos := []struct {
		nombre, viejo, nuevo             string
		codigo, ruta, attr, valor, regla string
	}{
		{"atributo de concepto repetible", `ClaveProdServ="43211500"`, `ClaveProdServ="99999999"`,
			EsquemaCatalogo, "/Comprobante/Conceptos/Concepto[2]/@ClaveProdServ", "ClaveProdServ", "99999999", "enumeration"},
		{"traslado del primer concepto", `TasaOCuota="0.160000" Importe="160.00"/>
        </cfdi:Traslados>
        <cfdi:Retenciones>`, `TasaOCuota="0.160000" Importe="160.00" X="1"/>
        </cfdi:Traslados>
        <cfdi:Retenciones>`,
			EsquemaAtributoNoPermitido, "/Comprobante/Conceptos/Concepto[1]/Impuestos/Traslados/Traslado[1]/@X", "X", "1", "atributo"},
		{"nodo único sin índice", `RegimenFiscal="601"/>`, `RegimenFiscal="999"/>`,
			EsquemaCatalogo, "/Comprobante/Emisor/@RegimenFiscal", "RegimenFiscal", "999", "enumeration"},
		{"valor original sin colapsar", `Moneda="MXN"`, `Moneda=" MX N "`,
			EsquemaCatalogo, "/Comprobante/@Moneda", "Moneda", " MX N ", "enumeration"},
		{"complemento sin índice", `Version="1.1" UUID`, `Version="1.1" UUID="X" UUIDX`,
			EsquemaAtributoNoPermitido, "/Comprobante/Complemento/TimbreFiscalDigital/@UUIDX", "UUIDX", "abcdef01-2345-4678-9abc-def012345678", "atributo"},
		{"segundo nodo único con índice", `<cfdi:Receptor `, `<cfdi:Emisor Rfc="EKU9003173C9" Nombre="X" RegimenFiscal="601"/><cfdi:Receptor `,
			EsquemaElementoExcedente, "/Comprobante/Emisor[2]", "", "", "maxOccurs"},
		{"complemento sin esquema", `</cfdi:Complemento>`, `<x:Otro xmlns:x="urn:x"/><x:Otro xmlns:x="urn:x"/></cfdi:Complemento>`,
			EsquemaNoSoportado, "/Comprobante/Complemento/Otro", "", "", "any"},
		{"xsi:schemaLocation se acepta en cualquier nodo", `<cfdi:Receptor `, `<cfdi:Receptor xsi:schemaLocation="a b" Otro="1" `,
			EsquemaAtributoNoPermitido, "/Comprobante/Receptor/@Otro", "Otro", "1", "atributo"},
	}
	for _, c := range casos {
		if !strings.Contains(base, c.viejo) {
			t.Fatalf("%s: la base no trae %q", c.nombre, c.viejo)
		}
		r := ValidarEsquema([]byte(strings.Replace(base, c.viejo, c.nuevo, 1)))
		var h *Hallazgo
		for i := range r.Hallazgos {
			if r.Hallazgos[i].Codigo == c.codigo && r.Hallazgos[i].Ruta == c.ruta {
				h = &r.Hallazgos[i]
			}
		}
		if h == nil {
			t.Errorf("%s: no hay %s en %s: %+v", c.nombre, c.codigo, c.ruta, r.Hallazgos)
			continue
		}
		if h.Atributo != c.attr || h.Valor != c.valor || h.Regla != c.regla || h.Mensaje == "" {
			t.Errorf("%s: %+v; quiero atributo %q valor %q regla %q", c.nombre, *h, c.attr, c.valor, c.regla)
		}
		if h.Nodo() != strings.TrimSuffix(c.ruta, "/@"+c.attr) {
			t.Errorf("%s: Nodo() = %q", c.nombre, h.Nodo())
		}
	}
	// Un nodo repetible que falta lleva índice.
	i, j := strings.Index(base, "<cfdi:Conceptos>"), strings.Index(base, "</cfdi:Conceptos>")
	r := ValidarEsquema([]byte(base[:i] + "<cfdi:Conceptos/>" + base[j+len("</cfdi:Conceptos>"):]))
	if len(r.Hallazgos) != 1 || r.Hallazgos[0].Codigo != EsquemaElementoFaltante || r.Hallazgos[0].Ruta != "/Comprobante/Conceptos/Concepto[1]" {
		t.Errorf("Conceptos vacío: %+v", r.Hallazgos)
	}
	// El segundo complemento sin esquema lleva índice.
	r = ValidarEsquema([]byte(strings.Replace(base, `</cfdi:Complemento>`, `<x:Otro xmlns:x="urn:x"/><x:Otro xmlns:x="urn:x"/></cfdi:Complemento>`, 1)))
	if r.Estado != ValidacionIncompleta || len(r.Hallazgos) != 2 || r.Hallazgos[1].Ruta != "/Comprobante/Complemento/Otro[2]" {
		t.Errorf("dos complementos sin esquema: %+v", r)
	}
}

func TestValidarEsquemaRaices(t *testing.T) {
	casos := []struct {
		nombre, xml, estado, codigo string
	}{
		{"no es XML", "hola", ValidacionInvalida, EsquemaXMLIlegible},
		{"vacío", "", ValidacionInvalida, EsquemaXMLIlegible},
		{"dos raíces", `<a/><b/>`, ValidacionInvalida, EsquemaXMLIlegible},
		{"texto tras la raíz", `<a/>x`, ValidacionInvalida, EsquemaXMLIlegible},
		{"raíz sin esquema", `<r:Otro xmlns:r="urn:otro"><x/></r:Otro>`, ValidacionIncompleta, EsquemaNoSoportado},
		{"Retenciones tiene esquema", `<r:Retenciones xmlns:r="http://www.sat.gob.mx/esquemas/retencionpago/2"><x/></r:Retenciones>`, ValidacionInvalida, EsquemaAtributoFaltante},
		{"CFDI 3.2", `<cfdi:Comprobante xmlns:cfdi="http://www.sat.gob.mx/cfd/3" version="3.2"/>`, ValidacionIncompleta, EsquemaNoSoportado},
		{"3.3 sin versión se valida", `<cfdi:Comprobante xmlns:cfdi="http://www.sat.gob.mx/cfd/3"/>`, ValidacionInvalida, EsquemaAtributoFaltante},
		{"namespace con espacios es solo advertencia de libxml2", `<a xmlns:x="urn: x"/>`, ValidacionIncompleta, EsquemaNoSoportado},
		{"prefijo sin declarar dentro de lo que no se valida", `<a><x:b/></a>`, ValidacionInvalida, EsquemaXMLIlegible},
		{"atributo repetido dentro de lo que no se valida", `<a><b c="1" c="2"/></a>`, ValidacionInvalida, EsquemaXMLIlegible},
		{"mal formado dentro de lo que no se valida", `<a><b></c></a>`, ValidacionInvalida, EsquemaXMLIlegible},
		{"prefijo de atributo sin declarar", `<a z:b="1"/>`, ValidacionInvalida, EsquemaXMLIlegible},
		{"entidad desconocida", `<a>&nada;</a>`, ValidacionInvalida, EsquemaXMLIlegible},
		{"control prohibido", "<a>\x01</a>", ValidacionInvalida, EsquemaXMLIlegible},
		{"UTF-8 inválido", "<a>\xff</a>", ValidacionInvalida, EsquemaXMLIlegible},
		{"encoding desconocido", `<?xml version="1.0" encoding="EBCDIC"?><a/>`, ValidacionInvalida, EsquemaXMLIlegible},
		{"ASCII con bytes altos", "<?xml version=\"1.0\" encoding=\"US-ASCII\"?><a>\xe9</a>", ValidacionInvalida, EsquemaXMLIlegible},
		{"windows-1252 con un byte que no existe", "<?xml version=\"1.0\" encoding=\"windows-1252\"?><a>\x81</a>", ValidacionInvalida, EsquemaXMLIlegible},
		{"latin1 se lee", "<?xml version=\"1.0\" encoding=\"ISO-8859-1\"?><a>\xe9</a>", ValidacionIncompleta, EsquemaNoSoportado},
		{"windows-1252 se lee", "<?xml version=\"1.0\" encoding=\"windows-1252\"?><a>\x80</a>", ValidacionIncompleta, EsquemaNoSoportado},
		{"BOM UTF-16", "\xff\xfe<\x00a\x00/\x00>\x00", ValidacionIncompleta, EsquemaNoSoportado},
		{"BOM UTF-16BE", "\xfe\xff\x00<\x00a\x00/\x00>", ValidacionIncompleta, EsquemaNoSoportado},
		{"BOM UTF-8", "\xef\xbb\xbf<a/>", ValidacionIncompleta, EsquemaNoSoportado},
	}
	for _, c := range casos {
		r := ValidarEsquema([]byte(c.xml))
		if r.Estado != c.estado || buscarHallazgo(r.Hallazgos, c.codigo) == nil {
			t.Errorf("%s: %+v; quiero %s con %s", c.nombre, r, c.estado, c.codigo)
		}
	}
	if r := ValidarEsquema(leerFixture(t, "testdata/sinteticos/ingreso40-pue.xml")); r.Estado != ValidacionValida || r.Hallazgos == nil || len(r.Hallazgos) != 0 {
		t.Errorf("válido: %+v (Hallazgos tiene que ser [] y no nil)", r)
	}
}

// TFD 1.0 y 1.1 comparten namespace y nombre: version="1.0" elige el XSD de
// 1.0 (como raíz o en un Complemento); cualquier otra cosa, el de 1.1.
func TestValidarEsquemaTFD10(t *testing.T) {
	ret10 := string(leerFixture(t, "testdata/sinteticos/retenciones/ret10-extranjero.xml"))
	const timbre = `<tfd:TimbreFiscalDigital xmlns:tfd="http://www.sat.gob.mx/TimbreFiscalDigital" version="1.0" UUID="1A2B3C4D-5E6F-4A7B-8C9D-0E1F2A3B4C5D" FechaTimbrado="2020-06-30T18:01:00" selloCFD="" noCertificadoSAT="30001000000300023708" selloSAT=""/>`
	if !strings.Contains(ret10, timbre) || strings.Count(ret10, `version="1.0" UUID`) != 1 {
		t.Fatal("el fixture no trae el timbre esperado")
	}
	cfdi40 := string(leerFixture(t, "testdata/sinteticos/ingreso40-pue.xml"))
	casos := []struct {
		nombre, xml, estado, codigo, attr string
	}{
		{"TFD 1.0 en Retenciones 1.0", ret10, ValidacionValida, "", ""},
		{"TFD 1.0 como raíz", timbre, ValidacionValida, "", ""},
		{"TFD 1.0 sin selloSAT", strings.Replace(ret10, ` selloSAT=""`, "", 1), ValidacionInvalida, EsquemaAtributoFaltante, "selloSAT"},
		{"TFD 1.0 con un atributo de 1.1", strings.Replace(ret10, ` selloSAT=""`, ` selloSAT="" RfcProvCertif="SAT970701NN3"`, 1),
			ValidacionInvalida, EsquemaAtributoNoPermitido, "RfcProvCertif"},
		{"noCertificadoSAT de 1.0 con su longitud", strings.Replace(ret10, `"30001000000300023708"`, `"3000100000030002370"`, 1),
			ValidacionInvalida, EsquemaLongitud, "noCertificadoSAT"},
		{"Version=\"1.1\" va contra 1.1", strings.Replace(ret10, `version="1.0" UUID`, `Version="1.1" UUID`, 1), ValidacionInvalida, EsquemaAtributoNoPermitido, "selloSAT"},
		{"version=\"1.1\" en minúsculas no es 1.0", strings.Replace(ret10, `version="1.0" UUID`, `version="1.1" UUID`, 1), ValidacionInvalida, EsquemaAtributoFaltante, "Version"},
		{"sin versión va contra 1.1", strings.Replace(ret10, ` version="1.0" UUID`, " UUID", 1), ValidacionInvalida, EsquemaAtributoFaltante, "SelloSAT"},
		{"TFD 1.1 válido", cfdi40, ValidacionValida, "", ""},
		{"TFD 1.1 sin SelloSAT", regexp.MustCompile(` SelloSAT="[^"]*"`).ReplaceAllString(cfdi40, ""), ValidacionInvalida, EsquemaAtributoFaltante, "SelloSAT"},
		{"TFD 1.1 con version=\"1.0\" va contra 1.0", strings.Replace(cfdi40, `Version="1.1" UUID`, `Version="1.1" version="1.0" UUID`, 1),
			ValidacionInvalida, EsquemaAtributoNoPermitido, "Version"},
	}
	for _, c := range casos {
		r := ValidarEsquema([]byte(c.xml))
		if r.Estado != c.estado {
			t.Errorf("%s: %s, quiero %s: %+v", c.nombre, r.Estado, c.estado, r.Hallazgos)
			continue
		}
		if c.codigo == "" {
			continue
		}
		ok := false
		for _, h := range r.Hallazgos {
			ok = ok || h.Codigo == c.codigo && h.Atributo == c.attr && strings.Contains(h.Ruta, "TimbreFiscalDigital")
		}
		if !ok {
			t.Errorf("%s: no hay %s en @%s del timbre: %+v", c.nombre, c.codigo, c.attr, r.Hallazgos)
		}
	}
}

func TestValidarValor(t *testing.T) {
	re := func(s string) []patronXSD { return []patronXSD{{xsd: s, re: regexp.MustCompile(s)}} }
	cat := &catalogoXSD{nombre: "c_Prueba", valores: "\nA\nB B\nC\n"}
	catInt := &catalogoXSD{nombre: "c_Entero", valores: "\n0\n1\n"}
	catDec := &catalogoXSD{nombre: "c_Decimal", valores: "\n0\n0.16\n"}
	fechaMin := &tipoXSD{prim: "dateTime", ws: "collapse", facetas: []facetaXSD{{tipo: "minInclusive", valor: "2014-01-01T00:00:00-06:00"}}}
	fechaMinSinZona := &tipoXSD{prim: "dateTime", ws: "collapse", facetas: []facetaXSD{{tipo: "minInclusive", valor: "2015-01-01T00:00:00"}}}
	fechaRango := &tipoXSD{prim: "date", ws: "collapse", facetas: []facetaXSD{{tipo: "minInclusive", valor: "2016-01-01"}, {tipo: "maxInclusive", valor: "2016-12-31"}}}
	casos := []struct {
		nombre string
		tipo   *tipoXSD
		valor  string
		quiero []string // reglas que fallan
	}{
		{"decimal ok", &tipoXSD{prim: "decimal", ws: "collapse"}, " 1.50 ", nil},
		{"decimal con signo y punto final", &tipoXSD{prim: "decimal", ws: "collapse"}, "+3.", nil},
		{"decimal inválido", &tipoXSD{prim: "decimal", ws: "collapse"}, "1e3", []string{"tipo"}},
		{"decimal con espacio en medio", &tipoXSD{prim: "decimal", ws: "collapse"}, "1 5", []string{"tipo"}},
		{"int fuera de rango", &tipoXSD{prim: "int", ws: "collapse"}, "2147483648", []string{"tipo"}},
		{"int en el límite", &tipoXSD{prim: "int", ws: "collapse"}, "-2147483648", nil},
		{"short fuera de rango", &tipoXSD{prim: "short", ws: "collapse"}, "40000", []string{"tipo"}},
		{"integer grande", &tipoXSD{prim: "integer", ws: "collapse"}, "+123456789012345678901234567890", nil},
		{"integer con punto", &tipoXSD{prim: "integer", ws: "collapse"}, "1.0", []string{"tipo"}},
		{"fecha bisiesta", &tipoXSD{prim: "date", ws: "collapse"}, "2024-02-29", nil},
		{"fecha no bisiesta", &tipoXSD{prim: "date", ws: "collapse"}, "2023-02-29", []string{"tipo"}},
		{"siglo no bisiesto", &tipoXSD{prim: "date", ws: "collapse"}, "1900-02-29", []string{"tipo"}},
		{"siglo bisiesto", &tipoXSD{prim: "date", ws: "collapse"}, "2000-02-29", nil},
		{"fecha con zona", &tipoXSD{prim: "date", ws: "collapse"}, "2024-01-31-06:00", nil},
		{"zona fuera de rango", &tipoXSD{prim: "date", ws: "collapse"}, "2024-01-31+15:00", []string{"tipo"}},
		{"zona 14:30", &tipoXSD{prim: "date", ws: "collapse"}, "2024-01-31+14:30", []string{"tipo"}},
		{"año cero", &tipoXSD{prim: "date", ws: "collapse"}, "0000-01-01", []string{"tipo"}},
		{"año de 5 cifras con cero", &tipoXSD{prim: "date", ws: "collapse"}, "01999-01-01", []string{"tipo"}},
		{"mes 13", &tipoXSD{prim: "date", ws: "collapse"}, "2024-13-01", []string{"tipo"}},
		{"día 0", &tipoXSD{prim: "date", ws: "collapse"}, "2024-01-00", []string{"tipo"}},
		{"fecha y hora", &tipoXSD{prim: "dateTime", ws: "collapse"}, "2024-01-31T23:59:59.5Z", nil},
		{"24:00:00", &tipoXSD{prim: "dateTime", ws: "collapse"}, "2024-01-31T24:00:00", nil},
		{"24:00:01", &tipoXSD{prim: "dateTime", ws: "collapse"}, "2024-01-31T24:00:01", []string{"tipo"}},
		{"minuto 60", &tipoXSD{prim: "dateTime", ws: "collapse"}, "2024-01-31T10:60:00", []string{"tipo"}},
		{"sin hora", &tipoXSD{prim: "dateTime", ws: "collapse"}, "2024-01-31", []string{"tipo"}},
		{"base64", &tipoXSD{prim: "base64Binary", ws: "collapse"}, "QUJD RA==", nil},
		{"base64 mal", &tipoXSD{prim: "base64Binary", ws: "collapse"}, "QUJ", []string{"tipo"}},
		{"base64 con caracter ajeno", &tipoXSD{prim: "base64Binary", ws: "collapse"}, "QU*D", []string{"tipo"}},
		{"patrón", &tipoXSD{prim: "string", ws: "collapse", facetas: []facetaXSD{{tipo: "pattern", patrones: re(`^(?:[0-9]{2})$`)}}}, "123", []string{"pattern"}},
		{"patrones alternativos", &tipoXSD{prim: "string", ws: "collapse", facetas: []facetaXSD{{tipo: "pattern", patrones: append(re(`^(?:a)$`), re(`^(?:b)$`)...)}}}, "b", nil},
		{"catálogo con espacio interno", &tipoXSD{prim: "string", ws: "collapse", facetas: []facetaXSD{{tipo: "enumeration", catalogo: cat}}}, "  B   B ", nil},
		{"catálogo", &tipoXSD{prim: "string", ws: "collapse", facetas: []facetaXSD{{tipo: "enumeration", catalogo: cat}}}, "D", []string{"enumeration"}},
		{"catálogo antes del primero", &tipoXSD{prim: "string", ws: "collapse", facetas: []facetaXSD{{tipo: "enumeration", catalogo: cat}}}, "0", []string{"enumeration"}},
		{"preserve no colapsa", &tipoXSD{prim: "string", ws: "preserve", facetas: []facetaXSD{{tipo: "enumeration", catalogo: cat}}}, " A", []string{"enumeration"}},
		{"replace", &tipoXSD{prim: "string", ws: "replace", facetas: []facetaXSD{{tipo: "length", n: 3}}}, "a\tb", nil},
		{"length en caracteres", &tipoXSD{prim: "string", ws: "collapse", facetas: []facetaXSD{{tipo: "length", n: 3}}}, "ÑÑÑ", nil},
		{"minLength", &tipoXSD{prim: "string", ws: "collapse", facetas: []facetaXSD{{tipo: "minLength", n: 1}}}, "   ", []string{"minLength"}},
		{"maxLength", &tipoXSD{prim: "string", ws: "collapse", facetas: []facetaXSD{{tipo: "maxLength", n: 2}}}, "abc", []string{"maxLength"}},
		{"minInclusive", &tipoXSD{prim: "decimal", ws: "collapse", facetas: []facetaXSD{{tipo: "minInclusive", valor: "0.000001"}}}, "0.0000001", []string{"minInclusive"}},
		{"maxInclusive", &tipoXSD{prim: "decimal", ws: "collapse", facetas: []facetaXSD{{tipo: "maxInclusive", valor: "100"}}}, "100.0", nil},
		{"maxInclusive excedido", &tipoXSD{prim: "decimal", ws: "collapse", facetas: []facetaXSD{{tipo: "maxInclusive", valor: "100"}}}, "100.01", []string{"maxInclusive"}},
		{"rango no se evalúa si no es decimal", &tipoXSD{prim: "decimal", ws: "collapse", facetas: []facetaXSD{{tipo: "minInclusive", valor: "1"}, {tipo: "fractionDigits", n: 1}}}, "x", []string{"tipo"}},
		{"fractionDigits cuenta el valor", &tipoXSD{prim: "decimal", ws: "collapse", facetas: []facetaXSD{{tipo: "fractionDigits", n: 2}}}, "1.2500000", nil},
		{"fractionDigits excedido", &tipoXSD{prim: "decimal", ws: "collapse", facetas: []facetaXSD{{tipo: "fractionDigits", n: 2}}}, "1.255", []string{"fractionDigits"}},
		{"totalDigits sin ceros a la izquierda", &tipoXSD{prim: "decimal", ws: "collapse", facetas: []facetaXSD{{tipo: "totalDigits", n: 3}}}, "-000123.000", nil},
		{"totalDigits excedido", &tipoXSD{prim: "decimal", ws: "collapse", facetas: []facetaXSD{{tipo: "totalDigits", n: 3}}}, "12.34", []string{"totalDigits"}},
		{"catálogo de un int compara valores", &tipoXSD{prim: "int", ws: "collapse", facetas: []facetaXSD{{tipo: "enumeration", catalogo: catInt}}}, "+01", nil},
		{"catálogo de un int", &tipoXSD{prim: "int", ws: "collapse", facetas: []facetaXSD{{tipo: "enumeration", catalogo: catInt}}}, "2", []string{"enumeration"}},
		{"catálogo de un int con valor ilegible", &tipoXSD{prim: "int", ws: "collapse", facetas: []facetaXSD{{tipo: "enumeration", catalogo: catInt}}}, "uno", []string{"tipo", "enumeration"}},
		{"catálogo de texto no canoniza", &tipoXSD{prim: "string", ws: "collapse", facetas: []facetaXSD{{tipo: "enumeration", catalogo: catInt}}}, "01", []string{"enumeration"}},
		{"catálogo de un decimal compara valores", &tipoXSD{prim: "decimal", ws: "collapse", facetas: []facetaXSD{{tipo: "enumeration", catalogo: catDec}}}, "0.160000", nil},
		{"catálogo de un decimal", &tipoXSD{prim: "decimal", ws: "collapse", facetas: []facetaXSD{{tipo: "enumeration", catalogo: catDec}}}, "0.1600001", []string{"enumeration"}},
		// minInclusive de un xs:dateTime (FechaExp de Retenciones 1.0): como
		// instante, con la zona.
		{"fecha en el límite", fechaMin, "2014-01-01T00:00:00-06:00", nil},
		{"fecha en el límite en otra zona", fechaMin, "2014-01-01T06:00:00Z", nil},
		{"fecha un segundo antes", fechaMin, "2013-12-31T23:59:59-06:00", []string{"minInclusive"}},
		{"fecha antes por la zona", fechaMin, "2014-01-01T00:00:00-05:00", []string{"minInclusive"}},
		{"fracción antes del límite", fechaMin, "2014-01-01T05:59:59.999Z", []string{"minInclusive"}},
		{"sin zona, lejos después", fechaMin, "2014-01-02T00:00:00", nil},
		{"sin zona, lejos antes", fechaMin, "2013-12-30T00:00:00", []string{"minInclusive"}},
		{"sin zona, indeterminada", fechaMin, "2014-01-01T00:00:00", []string{"minInclusive"}},
		{"fecha ilegible no evalúa el rango", fechaMin, "2014-13-01T00:00:00Z", []string{"tipo"}},
		// Límite sin zona (Fecha de ECC 1.1): contra otra sin zona se compara
		// directo; contra una con zona, solo fuera del margen de ±14 horas.
		{"límite sin zona, igual", fechaMinSinZona, "2015-01-01T00:00:00", nil},
		{"límite sin zona, un segundo antes", fechaMinSinZona, "2014-12-31T23:59:59", []string{"minInclusive"}},
		{"límite sin zona, con zona indeterminada", fechaMinSinZona, "2015-01-01T00:00:00-06:00", []string{"minInclusive"}},
		{"límite sin zona, con zona lejos después", fechaMinSinZona, "2015-01-01T14:00:01Z", nil},
		{"límite sin zona, con zona justo en el margen", fechaMinSinZona, "2015-01-01T14:00:00Z", []string{"minInclusive"}},
		{"límite sin zona, con zona lejos antes", fechaMinSinZona, "2014-12-31T09:59:59Z", []string{"minInclusive"}},
		// xs:time y el rango de un xs:date (SPEI).
		{"hora", &tipoXSD{prim: "time", ws: "collapse"}, "11:59:30", nil},
		{"hora con fracción y zona", &tipoXSD{prim: "time", ws: "collapse"}, "11:59:30.25-06:00", nil},
		{"hora 24:00:00", &tipoXSD{prim: "time", ws: "collapse"}, "24:00:00", nil},
		{"hora 24:30:00", &tipoXSD{prim: "time", ws: "collapse"}, "24:30:00", []string{"tipo"}},
		{"hora sin segundos", &tipoXSD{prim: "time", ws: "collapse"}, "11:59", []string{"tipo"}},
		{"hora con fecha", &tipoXSD{prim: "time", ws: "collapse"}, "2024-01-01T11:59:30", []string{"tipo"}},
		{"fecha en el rango", fechaRango, "2016-06-01", nil},
		{"fecha en el máximo", fechaRango, "2016-12-31", nil},
		{"fecha después del máximo", fechaRango, "2017-01-01", []string{"maxInclusive"}},
		{"fecha antes del mínimo", fechaRango, "2015-12-31", []string{"minInclusive"}},
		{"fecha con zona lejos del límite", fechaRango, "2016-06-01-06:00", nil},
		{"fecha con zona junto al límite", fechaRango, "2016-01-01+05:00", []string{"minInclusive"}},
	}
	for _, c := range casos {
		var got []string
		for _, f := range validarValor(c.tipo, c.valor) {
			got = append(got, f.regla)
			if f.msg == "" || f.codigo == "" {
				t.Errorf("%s: falla sin código o mensaje: %+v", c.nombre, f)
			}
		}
		if strings.Join(got, ",") != strings.Join(c.quiero, ",") {
			t.Errorf("%s: %q falla en %v, quiero %v", c.nombre, c.valor, got, c.quiero)
		}
	}
}

// Las tablas generadas son coherentes: catálogos ordenados y sin repetir
// (la búsqueda binaria lo necesita), patrones anclados, y los decimales
// de las facetas legibles.
func TestTablasEsquema(t *testing.T) {
	vistos := map[*tipoXSD]bool{}
	var revisar func(e *elementoXSD, ruta string)
	revisar = func(e *elementoXSD, ruta string) {
		for _, a := range e.attrs {
			if vistos[a.tipo] {
				continue
			}
			vistos[a.tipo] = true
			for _, f := range a.tipo.facetas {
				switch f.tipo {
				case "enumeration":
					f.catalogo.contiene("")
					l := f.catalogo.lista
					for i := 1; i < len(l); i++ {
						if l[i-1] >= l[i] {
							t.Errorf("%s: catálogo %s desordenado en %q, %q", ruta, f.catalogo.nombre, l[i-1], l[i])
						}
					}
				case "pattern":
					for _, p := range f.patrones {
						if s := p.re.String(); !strings.HasPrefix(s, "^(?:") || !strings.HasSuffix(s, ")$") {
							t.Errorf("%s: patrón sin anclar %s", ruta, s)
						}
					}
				case "minInclusive", "maxInclusive":
					if a.tipo.prim == "dateTime" || a.tipo.prim == "date" {
						if _, lexOK := lexicoXSD(a.tipo.prim, f.valor); !lexOK {
							t.Errorf("%s: %s no es un xs:%s", ruta, f.valor, a.tipo.prim)
						}
					} else if _, err := decimal.Parse(f.valor); err != nil {
						t.Errorf("%s: %v", ruta, err)
					}
				}
			}
		}
		for _, p := range e.particulas {
			if p.elem != nil {
				revisar(p.elem, ruta+"/"+p.elem.nombre)
			}
			for _, o := range p.opciones {
				revisar(o.elem, ruta+"/"+o.elem.nombre)
			}
		}
	}
	for n, e := range globalesXSD {
		revisar(e, n.String())
	}
	if len(vistos) < 200 {
		t.Errorf("solo %d tipos alcanzables", len(vistos))
	}
	// Catálogos clave presentes.
	for _, c := range []struct {
		cat *catalogoXSD
		v   string
	}{{catCatCFDI_cMoneda, "MXN"}, {catCatCFDI_cMoneda, "XXX"}, {catCatCFDI_cClaveProdServ, "01010101"}, {catCatCFDI_cCodigoPostal, "42501"}} {
		if !c.cat.contiene(c.v) {
			t.Errorf("%s no tiene %s", c.cat.nombre, c.v)
		}
	}
}

// Un xs:choice requerido (ningún XSD del SAT lo trae; los de terceros y
// ventavehiculos tienen todas sus opciones opcionales): falta si no viene
// ninguna opción, y la opción elegida exige su propio minOccurs.
func TestValidarEsquemaChoiceRequerido(t *testing.T) {
	a := &elementoXSD{ns: "urn:p", nombre: "A", vacio: true}
	b := &elementoXSD{ns: "urn:p", nombre: "B", vacio: true}
	c := &elementoXSD{ns: "urn:p", nombre: "C", vacio: true}
	raiz := &elementoXSD{ns: "urn:p", nombre: "R", seqMin: 1, seqMax: 1, particulas: []particulaXSD{
		{min: 1, max: 1, opciones: []particulaXSD{{elem: a, min: 1, max: 1}, {elem: b, min: 2, max: 3}}},
		{elem: c, min: 0, max: 1},
	}}
	nodo := func(hijos ...*elementoXSD) *nodoV {
		n := &nodoV{nombre: nombreXML{"urn:p", "R"}, decl: raiz}
		for _, h := range hijos {
			n.hijos = append(n.hijos, &nodoV{nombre: nombreXML{h.ns, h.nombre}, decl: h})
		}
		return n
	}
	casos := []struct {
		nombre string
		hijos  []*elementoXSD
		quiero []string // código y ruta de cada hallazgo
	}{
		{"una A", []*elementoXSD{a}, nil},
		{"dos B", []*elementoXSD{b, b, c}, nil},
		{"ninguna opción", []*elementoXSD{c}, []string{EsquemaElementoFaltante + " /R"}},
		{"vacío", nil, []string{EsquemaElementoFaltante + " /R"}},
		{"una B de dos", []*elementoXSD{b}, []string{EsquemaElementoFaltante + " /R/B[2]"}},
		{"cuatro B", []*elementoXSD{b, b, b, b}, []string{EsquemaElementoExcedente + " /R/B[4]"}},
		{"A y B", []*elementoXSD{a, b}, []string{EsquemaElementoNoPermitido + " /R/B[1]"}},
		{"C antes de la opción", []*elementoXSD{c, a}, []string{EsquemaOrden + " /R/C"}},
	}
	for _, cs := range casos {
		v := &validadorXSD{}
		v.elemento(nodo(cs.hijos...), "/R")
		var got []string
		for _, h := range v.hs {
			got = append(got, h.Codigo+" "+h.Ruta)
		}
		if strings.Join(got, "|") != strings.Join(cs.quiero, "|") {
			t.Errorf("%s: %q, quiero %q", cs.nombre, got, cs.quiero)
		}
	}
}
