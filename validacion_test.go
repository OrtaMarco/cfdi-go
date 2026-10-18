package cfdi

import (
	"strings"
	"testing"
)

func TestEstadoDe(t *testing.T) {
	e, a, w := Hallazgo{Nivel: NivelError}, Hallazgo{Nivel: NivelAviso}, Hallazgo{Nivel: NivelAdvertencia}
	casos := []struct {
		hs     []Hallazgo
		quiero string
	}{
		{nil, ValidacionValida},
		{[]Hallazgo{w}, ValidacionValida},
		{[]Hallazgo{w, a}, ValidacionIncompleta},
		{[]Hallazgo{a, w}, ValidacionIncompleta},
		{[]Hallazgo{w, e}, ValidacionInvalida},
		{[]Hallazgo{a, e, w}, ValidacionInvalida},
	}
	for i, c := range casos {
		if got := EstadoDe(c.hs); got != c.quiero {
			t.Errorf("caso %d: %s, quiero %s", i, got, c.quiero)
		}
	}
}

func TestValidar(t *testing.T) {
	ingreso := string(leerFixture(t, "testdata/sinteticos/ingreso40-pue.xml"))
	terceros33 := string(leerFixture(t, "testdata/mutantes/base-terceros33.xml")) // 3.3 válido contra el XSD
	casos := []struct {
		nombre, xml, estado, esquema, reglas string
	}{
		{"válido", ingreso, ValidacionValida, ValidacionValida, ValidacionValida},
		{"esquema inválido, reglas válidas", strings.Replace(ingreso, `RegimenFiscal="601"/>`, `RegimenFiscal="601" X="1"/>`, 1), ValidacionInvalida, ValidacionInvalida, ValidacionValida},
		{"reglas inválidas", strings.Replace(ingreso, ` Total="1640.00"`, ` Total="1640.01"`, 1), ValidacionInvalida, ValidacionValida, ValidacionInvalida},
		// En 3.3 los decimales de los traslados del comprobante son
		// advertencia (la matriz 3.3 pide no aplicar CFDI33194).
		{"solo advertencias: válido", strings.Replace(terceros33, `TasaOCuota="0.160000" Importe="208.00"/>`, `TasaOCuota="0.160000" Importe="208.000000"/>`, 1),
			ValidacionValida, ValidacionValida, ValidacionValida},
		{"incompleto", strings.Replace(ingreso, `</cfdi:Complemento>`, `<x:Y xmlns:x="urn:x"/></cfdi:Complemento>`, 1), ValidacionIncompleta, ValidacionIncompleta, ValidacionValida},
		{"incompleto e inválido", strings.Replace(strings.Replace(ingreso, `</cfdi:Complemento>`, `<x:Y xmlns:x="urn:x"/></cfdi:Complemento>`, 1), ` Total="1640.00"`, ` Total="1"`, 1), ValidacionInvalida, ValidacionIncompleta, ValidacionInvalida},
		{"CFDI 3.2: los dos avisos", string(leerFixture(t, "testdata/sinteticos/sellados/domicilios32.xml")), ValidacionIncompleta, ValidacionIncompleta, ValidacionIncompleta},
		{"no es CFDI", "<a/>", ValidacionIncompleta, ValidacionIncompleta, ""},
		{"no es XML", "<a", ValidacionInvalida, ValidacionInvalida, ""},
	}
	for _, c := range casos {
		v := Validar([]byte(c.xml), nil)
		reglas := ""
		if v.Reglas != nil {
			reglas = v.Reglas.Estado
		}
		if v.Estado != c.estado || v.Esquema.Estado != c.esquema || reglas != c.reglas {
			t.Errorf("%s: %s (esquema %s, reglas %q); quiero %s (%s, %q)", c.nombre, v.Estado, v.Esquema.Estado, reglas, c.estado, c.esquema, c.reglas)
		}
		if n := len(v.Esquema.Hallazgos); v.Reglas != nil {
			n += len(v.Reglas.Hallazgos)
			if len(v.Hallazgos()) != n {
				t.Errorf("%s: Hallazgos() no junta los dos", c.nombre)
			}
		}
	}
	if v := Validar([]byte(casos[3].xml), nil); len(v.Reglas.Hallazgos) != 1 || v.Reglas.Hallazgos[0].Nivel != NivelAdvertencia {
		t.Errorf("solo advertencias: %+v", v.Reglas)
	}
	// Con el Comprobante ya leído no se vuelve a leer.
	c, _ := ParseBytes([]byte(ingreso))
	c.Total = nil
	if v := Validar([]byte(ingreso), c); v.Reglas.Estado != ValidacionValida || v.Estado != ValidacionValida {
		t.Errorf("con Comprobante: %+v", v)
	}
}
