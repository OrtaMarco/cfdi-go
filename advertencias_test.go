package cfdi

import (
	"bytes"
	"encoding/json"
	"os"
	"strings"
	"testing"
)

// Toda advertencia de atributo trae Atributo igual al sufijo de Ruta, y
// Nodo() + "/@" + Atributo reconstruye Ruta: el contrato que documenta
// Advertencia, comprobado sobre todos los fixtures.
func TestAdvertenciaAtributoCoincideConRuta(t *testing.T) {
	vistas := 0
	for _, p := range fixtures(t) {
		if esFatal(p) || esRetenciones(t, p) {
			continue
		}
		b, _ := os.ReadFile(p)
		c, err := ParseBytes(b)
		if err != nil {
			t.Fatal(err)
		}
		for _, a := range c.Advertencias {
			_, sufijo, conAttr := strings.Cut(a.Ruta, "/@")
			switch {
			case conAttr && a.Atributo != sufijo:
				t.Errorf("%s: Ruta %q y Atributo %q", p, a.Ruta, a.Atributo)
			case !conAttr && a.Atributo != "":
				t.Errorf("%s: Ruta %q sin /@ y Atributo %q", p, a.Ruta, a.Atributo)
			case conAttr && a.Nodo()+"/@"+a.Atributo != a.Ruta:
				t.Errorf("%s: Nodo() %q no reconstruye %q", p, a.Nodo(), a.Ruta)
			}
			if conAttr {
				vistas++
			}
		}
	}
	if vistas == 0 {
		t.Fatal("ningún fixture tiene advertencias de atributo: la prueba no prueba nada")
	}
}

func TestAdvertenciaValorYPorCampo(t *testing.T) {
	c := leer(t, "anomalias.xml")
	casos := []struct {
		nodo, attr, codigo, valor string
	}{
		{"/Comprobante", "SubTotal", AdvDecimalInvalido, "1,000.00"},
		{"/Comprobante/Conceptos/Concepto[1]", "Cantidad", AdvDecimalInvalido, "1e3"},
		{"/Comprobante", "Fecha", AdvFechaFormato, "2026-07-01T10:00:00-06:00"},
		{"/Comprobante/Complemento/TimbreFiscalDigital", "FechaTimbrado", AdvFechaInvalida, "ayer"},
		{"/Comprobante/Complemento/TimbreFiscalDigital", "UUID", AdvUUIDInvalido, "no-es-un-uuid"},
		{"/Comprobante/Conceptos/Concepto[1]/Impuestos/Traslados/Traslado[2]", "Importe", AdvAtributoFaltante, ""},
	}
	for _, cs := range casos {
		a, ok := c.AdvertenciaEn(cs.nodo, cs.attr)
		if !ok {
			t.Errorf("AdvertenciaEn(%q, %q): no la encontró", cs.nodo, cs.attr)
			continue
		}
		igual(t, cs.attr+" código", a.Codigo, cs.codigo)
		igual(t, cs.attr+" valor", a.Valor, cs.valor)
		igual(t, cs.attr+" atributo", a.Atributo, cs.attr)
	}
	if _, ok := c.AdvertenciaEn("/Comprobante", "Total"); ok {
		t.Error("Total no tiene advertencia y AdvertenciaEn dijo que sí")
	}
	// El nodo y el atributo tienen que coincidir los dos.
	if _, ok := c.AdvertenciaEn("/Comprobante/Conceptos/Concepto[1]", "SubTotal"); ok {
		t.Error("AdvertenciaEn encontró SubTotal en el Concepto")
	}

	// AdvertenciasDe trae las del nodo y sus atributos, no las de sus hijos.
	de := c.AdvertenciasDe("/Comprobante/Conceptos/Concepto[1]")
	if len(de) != 1 || de[0].Atributo != "Cantidad" {
		t.Errorf("AdvertenciasDe(Concepto[1]) = %+v, quiero solo la de Cantidad", de)
	}
	tras := c.AdvertenciasDe("/Comprobante/Conceptos/Concepto[1]/Impuestos/Traslados/Traslado[1]")
	if len(tras) != 1 || tras[0].Codigo != AdvExentoConImporte || tras[0].Atributo != "" {
		t.Errorf("AdvertenciasDe(Traslado[1]) = %+v, quiero EXENTO_CON_IMPORTE del nodo", tras)
	}
	doc := c.AdvertenciasDe("")
	if len(doc) != 1 || doc[0].Codigo != AdvContenidoTrasRaiz {
		t.Errorf("AdvertenciasDe(\"\") = %+v, quiero CONTENIDO_TRAS_RAIZ", doc)
	}

	// En el JSON salen atributo y valor; valor vacío se omite.
	b, err := c.JSON()
	if err != nil {
		t.Fatal(err)
	}
	for _, s := range []string{`"atributo":"SubTotal","valor":"1,000.00"`, `"atributo":"Importe"}`} {
		if !bytes.Contains(b, []byte(s)) {
			t.Errorf("el JSON no trae %s", s)
		}
	}
}

func TestDecimalVacioConservaTexto(t *testing.T) {
	xml := strings.Replace(string(leerSintetico(t, "ingreso40-pue.xml")), `SubTotal="1500.00"`, `SubTotal=" "`, 1)
	c, err := ParseBytes([]byte(xml))
	if err != nil {
		t.Fatal(err)
	}
	a, ok := c.AdvertenciaEn("/Comprobante", "SubTotal")
	if !ok || a.Codigo != AdvDecimalInvalido || a.Valor != " " {
		t.Errorf("SubTotal=\" \": %+v %v", a, ok)
	}
}

func TestFechaEstricta(t *testing.T) {
	b := leerSintetico(t, "anomalias.xml")
	tolerante, err := ParseBytes(b)
	if err != nil {
		t.Fatal(err)
	}
	estricta, err := ParseBytesOpciones(b, Opciones{FechaEstricta: true})
	if err != nil {
		t.Fatal(err)
	}
	// Tolerante: se lee y se marca.
	if !tolerante.Fecha.Valida() || !tolerante.Fecha.FueraDeFormato {
		t.Errorf("tolerante: %+v, quiero válida y FueraDeFormato", tolerante.Fecha)
	}
	// Estricta: no se interpreta, sigue marcada y la advertencia cambia.
	if estricta.Fecha.Valida() || !estricta.Fecha.FueraDeFormato || estricta.Fecha.Original != "2026-07-01T10:00:00-06:00" {
		t.Errorf("estricta: %+v, quiero inválida con el texto y FueraDeFormato", estricta.Fecha)
	}
	if a, _ := estricta.AdvertenciaEn("/Comprobante", "Fecha"); a.Codigo != AdvFechaInvalida || a.Valor != "2026-07-01T10:00:00-06:00" {
		t.Errorf("estricta: advertencia %+v, quiero FECHA_INVALIDA con el valor", a)
	}
	if estricta.TieneAdvertencia(AdvFechaFormato) {
		t.Error("estricta no debe dejar FECHA_FORMATO")
	}
	// Una fecha bien formada no cambia con la opción, ni las de solo fecha
	// de la nómina.
	for _, f := range []string{"ingreso40-pue.xml", "nomina12.xml"} {
		c, err := ParseBytesOpciones(leerSintetico(t, f), Opciones{FechaEstricta: true})
		if err != nil {
			t.Fatal(err)
		}
		if c.TieneAdvertencia(AdvFechaInvalida) || c.TieneAdvertencia(AdvFechaFormato) || c.Fecha.FueraDeFormato || !c.Fecha.Valida() {
			t.Errorf("%s estricta: %+v %+v", f, c.Fecha, c.Advertencias)
		}
	}

	r, err := ParseOpciones(bytes.NewReader(b), Opciones{FechaEstricta: true})
	if err != nil || r.Fecha.Valida() {
		t.Errorf("ParseOpciones no aplicó la opción: %v %+v", err, r.Fecha)
	}
	if _, err := ParseOpciones(lectorRoto{}, Opciones{}); ClaseDe(err) != ClaseLectura {
		t.Errorf("ParseOpciones con lector roto: %v", err)
	}
}

func TestFechaFueraDeFormatoJSON(t *testing.T) {
	casos := map[string]bool{
		`"2026-01-15T10:30:00"`:       false,
		`"2026-01-15"`:                false,
		`"2026-01-15T10:30:00Z"`:      true,
		`"2026-01-15T10:30:00.5"`:     true,
		`" 2026-01-15T10:30:00"`:      true,
		`"no es fecha"`:               false,
		`"2026-01-15T10:30:00-06:00"`: true,
	}
	for in, want := range casos {
		var f Fecha
		if err := json.Unmarshal([]byte(in), &f); err != nil {
			t.Fatal(err)
		}
		if f.FueraDeFormato != want {
			t.Errorf("%s: FueraDeFormato = %v, quiero %v", in, f.FueraDeFormato, want)
		}
	}
}

func leerSintetico(t *testing.T, nombre string) []byte {
	t.Helper()
	b, err := os.ReadFile("testdata/sinteticos/" + nombre)
	if err != nil {
		t.Fatal(err)
	}
	return b
}
