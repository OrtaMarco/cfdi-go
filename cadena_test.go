package cfdi

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/OrtaMarco/cfdi-go/internal/gencadena"
)

// esperado lee el golden de una cadena: el texto exacto (.txt) o el motivo
// por el que debe fallar (.error). Los generó scripts/generar_cadenas.py con
// libxslt y los XSLT del SAT, sin usar código de esta librería.
func esperado(t *testing.T, base, sufijo string) (cadena, motivo, nodo string, hay bool) {
	t.Helper()
	if b, err := os.ReadFile(base + sufijo + ".txt"); err == nil {
		return string(b), "", "", true
	}
	b, err := os.ReadFile(base + sufijo + ".error")
	if err != nil {
		return "", "", "", false
	}
	motivo, nodo, _ = strings.Cut(strings.TrimSpace(string(b)), " ")
	return "", motivo, nodo, true
}

func compararCadena(t *testing.T, que, base, sufijo string, got string, err error) {
	t.Helper()
	want, motivo, nodo, hay := esperado(t, base, sufijo)
	if !hay {
		if err == nil {
			t.Errorf("%s: no hay golden %s.{txt,error} y se armó %q", que, sufijo, got)
		}
		return
	}
	if motivo == "" {
		if err != nil {
			t.Errorf("%s: error %v, quiero\n%s", que, err, want)
		} else if got != want {
			t.Errorf("%s difiere del golden:\n%s", que, diferenciaCadena(want, got))
		}
		return
	}
	var ec *ErrorCadena
	if !errors.As(err, &ec) {
		t.Errorf("%s: quiero error %s %s y obtuve %q (err %v)", que, motivo, nodo, got, err)
		return
	}
	if ec.Motivo != motivo || ec.Nodo != nodo {
		t.Errorf("%s: error %s %s, quiero %s %s", que, ec.Motivo, ec.Nodo, motivo, nodo)
	}
}

// diferenciaCadena señala el primer campo distinto entre dos cadenas.
func diferenciaCadena(want, got string) string {
	w, g := strings.Split(want, "|"), strings.Split(got, "|")
	for i := 0; i < len(w) || i < len(g); i++ {
		var a, b string
		if i < len(w) {
			a = w[i]
		}
		if i < len(g) {
			b = g[i]
		}
		if a != b {
			return "campo " + strconv.Itoa(i) + ": quiero " + strconv.Quote(a) + ", obtuve " + strconv.Quote(b) +
				"\n  quiero: " + want + "\n  obtuve: " + got
		}
	}
	return "(mismos campos)\n  quiero: " + want + "\n  obtuve: " + got
}

// fixturesCadena son los fixtures que entran a las pruebas de cadena y
// sello: no los fatales ni los *-real* (CFDI reales que publicaron otros
// proyectos: ya estaban en el repo para la lectura y no se usan para nada
// nuevo).
func fixturesCadena(t *testing.T) []string {
	var out []string
	for _, p := range fixtures(t) {
		if !esFatal(p) && !strings.Contains(filepath.Base(p), "-real") {
			out = append(out, p)
		}
	}
	return out
}

func TestCadenaGolden(t *testing.T) {
	conTxt, conError, retenciones := 0, 0, 0
	for _, p := range fixturesCadena(t) {
		base := strings.TrimSuffix(p, filepath.Ext(p))
		t.Run(filepath.ToSlash(p), func(t *testing.T) {
			b, err := os.ReadFile(p)
			if err != nil {
				t.Fatal(err)
			}
			cadena := CadenaOriginal
			if Detectar(b) == TipoRetenciones {
				cadena = CadenaOriginalRetenciones
				if _, err := os.Stat(base + ".cadena.txt"); err == nil {
					retenciones++
				}
			}
			got, err := cadena(b)
			compararCadena(t, "cadena", base, ".cadena", got, err)
			gotT, errT := CadenaOriginalTFD(b)
			if _, _, _, hay := esperado(t, base, ".cadena-tfd"); !hay {
				// Sin TFD (o el documento no se lee): tiene que fallar.
				if errT == nil {
					t.Errorf("TFD: sin golden y se armó %q", gotT)
				}
				return
			}
			compararCadena(t, "cadena TFD", base, ".cadena-tfd", gotT, errT)
		})
		if _, err := os.Stat(base + ".cadena.txt"); err == nil {
			conTxt++
		} else {
			conError++
		}
	}
	// Que la prueba no pase en vacío si alguien borra los goldens.
	if conTxt < 240 || conError < 40 || retenciones < 50 {
		t.Errorf("pocos goldens: %d cadenas (%d de Retenciones) y %d errores", conTxt, retenciones, conError)
	}
}

// La cadena que publica CfdiUtils para cfdi33-valid.xml (fuente distinta de
// libxslt: la armó ese proyecto). Corresponde a una versión anterior del
// fixture, firmada con otro certificado: difiere solo en NoCertificado. En el
// fixture y en esta cadena el Receptor se cambió por uno de prueba (ver NOTICE).
func TestCadenaPublicadaCfdiUtils(t *testing.T) {
	dir := "testdata/publicos/eclipxe13-CfdiUtils/"
	pub, err := os.ReadFile(dir + "tests__assets__cfdi33-valid-cadenaorigen.txt")
	if err != nil {
		t.Fatal(err)
	}
	want := strings.TrimSpace(string(pub))
	if strings.Count(want, "|30001000000500003416|") != 1 {
		t.Fatal("la cadena publicada ya no trae el NoCertificado anterior: revisar la prueba")
	}
	want = strings.Replace(want, "|30001000000500003416|", "|30001000000400002434|", 1)
	b, err := os.ReadFile(dir + "tests__assets__cfdi33-valid.xml")
	if err != nil {
		t.Fatal(err)
	}
	got, err := CadenaOriginal(b)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("difiere de la cadena de CfdiUtils:\n%s", diferenciaCadena(want, got))
	}
}

// cadena_reglas.go tiene que ser lo que genera el generador con los XSLT
// del repo: nadie lo edita a mano ni se queda viejo.
func TestReglasAlDia(t *testing.T) {
	got, err := gencadena.Generar("internal/gencadena/xslt/sitio_internet")
	if err != nil {
		t.Fatal(err)
	}
	want, err := os.ReadFile("cadena_reglas.go")
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != string(want) {
		t.Error("cadena_reglas.go no es lo que genera internal/gencadena: corre `go generate .`")
	}
}

func TestCadenaErrores(t *testing.T) {
	base := string(leerSintetico(t, "ingreso40-pue.xml"))
	casos := []struct {
		nombre, xml, motivo, nodo string
	}{
		{"no es XML", "hola", MotivoXMLIlegible, ""},
		{"truncado", base[:len(base)-40], MotivoXMLIlegible, ""},
		{"dos raíces", base + "<otra/>", MotivoXMLIlegible, ""},
		{"texto tras la raíz", base + "basura", MotivoXMLIlegible, ""},
		{"entidad desconocida", strings.Replace(base, `Folio="1001"`, `Folio="&nbsp;"`, 1), MotivoXMLIlegible, ""},
		{"utf-8 que no lo es", strings.Replace(base, "retención", "retenci\xf3n", 1), MotivoXMLIlegible, ""},
		{"carácter prohibido", strings.Replace(base, `Folio="1001"`, "Folio=\"10\x0101\"", 1), MotivoXMLIlegible, ""},
		{"Addenda rota", strings.Replace(base, "</cfdi:Complemento>", "</cfdi:Complemento><cfdi:Addenda><x>", 1), MotivoXMLIlegible, ""},
		// encoding/xml los deja pasar; libxml2 (el procesador de los XSLT) no.
		{"atributo repetido", strings.Replace(base, `Serie="A"`, `Serie="A" Serie="B"`, 1), MotivoXMLIlegible, ""},
		{"prefijo sin declarar", strings.Replace(base, `<cfdi:Emisor `, `<zz:Emisor `, 1), MotivoXMLIlegible, ""},
		{"prefijo sin declarar en un atributo", strings.Replace(base, `<cfdi:Emisor `, `<cfdi:Emisor zz:a="1" `, 1), MotivoXMLIlegible, ""},
		{"prefijo sin declarar en la Addenda", strings.Replace(base, "</cfdi:Complemento>", "</cfdi:Complemento><cfdi:Addenda><zz:x/></cfdi:Addenda>", 1), MotivoXMLIlegible, ""},
		{"atributo repetido en la Addenda", strings.Replace(base, "</cfdi:Complemento>", `</cfdi:Complemento><cfdi:Addenda><x a="1" a="2"/></cfdi:Addenda>`, 1), MotivoXMLIlegible, ""},
		{"no es comprobante", `<Retenciones xmlns="http://www.sat.gob.mx/esquemas/retencionpago/2" Version="2.0"/>`, MotivoNoEsComprobante, "{http://www.sat.gob.mx/esquemas/retencionpago/2}Retenciones"},
		{"versión 3.0", strings.Replace(strings.Replace(base, "cfd/4", "cfd/3", 1), `Version="4.0"`, `version="3.0"`, 1), MotivoVersionNoSoportada, ""},
		{"4.0 con Version 3.3", strings.Replace(base, `Version="4.0"`, `Version="3.3"`, 1), MotivoVersionNoSoportada, ""},
		{"complemento desconocido", strings.Replace(base, "</cfdi:Complemento>", `<x:Raro xmlns:x="urn:raro"/></cfdi:Complemento>`, 1), MotivoComplementoSinCadena, "{urn:raro}Raro"},
	}
	for _, c := range casos {
		_, err := CadenaOriginal([]byte(c.xml))
		var ec *ErrorCadena
		if !errors.As(err, &ec) {
			t.Errorf("%s: quiero *ErrorCadena y obtuve %v", c.nombre, err)
			continue
		}
		if ec.Motivo != c.motivo || ec.Nodo != c.nodo {
			t.Errorf("%s: %s %s, quiero %s %s", c.nombre, ec.Motivo, ec.Nodo, c.motivo, c.nodo)
		}
		if !strings.Contains(ec.Error(), c.motivo) || (c.nodo != "" && !strings.Contains(ec.Error(), c.nodo)) {
			t.Errorf("%s: el mensaje no dice el motivo y el nodo: %v", c.nombre, ec)
		}
	}

	// TFD: sin timbre, versión desconocida y el error de lectura.
	if _, err := CadenaOriginalTFD(leerSintetico(t, "sin-tfd.xml")); !esMotivo(err, MotivoSinTimbre) {
		t.Errorf("sin TFD: %v", err)
	}
	if _, err := CadenaOriginalTFD([]byte(strings.Replace(base, `Version="1.1"`, `Version="2.0"`, 1))); !esMotivo(err, MotivoVersionNoSoportada) {
		t.Errorf("TFD 2.0: %v", err)
	}
	if _, err := CadenaOriginalTFD([]byte("<")); !esMotivo(err, MotivoXMLIlegible) {
		t.Errorf("TFD de XML roto: %v", err)
	}
	// El error envuelve la causa.
	if _, err := CadenaOriginal([]byte("<a>")); errors.Unwrap(err) == nil {
		t.Error("ErrorCadena no envuelve la causa")
	}
}

func esMotivo(err error, motivo string) bool {
	var ec *ErrorCadena
	return errors.As(err, &ec) && ec.Motivo == motivo
}

func TestNormalizarEspacios(t *testing.T) {
	casos := map[string]string{
		"":                  "",
		"   ":               "",
		" a ":               "a",
		"a\t\tb":            "a b",
		"a\r\n b":           "a b",
		" a ":               " a ", // el espacio no separable no es espacio XML
		"a b":               "a b",
		"  uno  dos   tres": "uno dos tres",
	}
	for in, want := range casos {
		if got := normalizarEspacios(in); got != want {
			t.Errorf("normalizarEspacios(%q) = %q, quiero %q", in, got, want)
		}
	}
}

// Las hojas del SAT no usan rutas con // desde varios nodos a la vez, pero
// el evaluador las resuelve como XPath: en orden de documento y sin
// repetidos.
func TestSeleccionarDescendientesSinRepetir(t *testing.T) {
	raiz, _, err := arbolCadena([]byte(`<a xmlns="urn:x"><b><b><c i="1"/></b><c i="2"/></b><c i="3"/></a>`))
	if err != nil {
		t.Fatal(err)
	}
	ev := &evaluador{doc: raiz}
	var got []string
	for _, n := range ev.seleccionar(raiz, rel(desc("urn:x", "b"), desc("urn:x", "c"))) {
		got = append(got, n.attr("i"))
	}
	if strings.Join(got, ",") != "1,2" {
		t.Errorf("//b//c = %v, quiero [1 2]", got)
	}
	if n := ev.seleccionar(nil, abs(hijo("urn:otro", "a"))); len(n) != 0 {
		t.Errorf("/otro:a seleccionó %d nodos", len(n))
	}
	if n := ev.seleccionar(raiz, rel(hijo("", "*"))); len(n) != 2 {
		t.Errorf("./* = %d nodos, quiero 2 (sin el texto)", len(n))
	}
}
