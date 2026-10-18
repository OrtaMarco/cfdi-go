package cfdi

import (
	"strings"
	"testing"
)

// Las reglas de dos niveles (Regla.Margen), con números como los de
// documentos timbrados por un PAC: sin hallazgo si cumple la letra del SAT,
// advertencia si no la cumple pero cae en el margen de redondeo por
// concepto, error si queda fuera también del margen.
func TestValidarReglasDosNiveles(t *testing.T) {
	tras := func(base, importe string) Impuesto {
		return Impuesto{Base: ptrDec(base), Impuesto: "002", TipoFactor: "Tasa", TasaOCuota: ptrDec("0.160000"), Importe: ptrDec(importe)}
	}
	ret := func(base, importe string) Impuesto {
		return Impuesto{Base: ptrDec(base), Impuesto: "001", TipoFactor: "Tasa", TasaOCuota: ptrDec("0.100000"), Importe: ptrDec(importe)}
	}
	comprobante := func(version string, imp *Impuestos, cs ...ImpuestosConcepto) *Comprobante {
		c := &Comprobante{Version: version, TipoDeComprobante: "I", Moneda: "MXN", Impuestos: imp,
			Emisor: &Emisor{Rfc: "EKU9003173C9"}, Receptor: &Receptor{Rfc: "XAXX010101000"}}
		for i := range cs {
			c.Conceptos = append(c.Conceptos, Concepto{ObjetoImp: "02", Impuestos: &cs[i]})
		}
		return c
	}
	// Un concepto con IVA 16 % sobre base, sin nodo Impuestos en el
	// comprobante.
	importeImpuesto := func(base, importe string) *Comprobante {
		return comprobante("4.0", nil, ImpuestosConcepto{Traslados: []Impuesto{tras(base, importe)}})
	}
	// Dos conceptos que suman 1633.7952 (1000.0001 + 633.7951): la letra
	// admite 1633.7952 o 1633.80; truncando y redondeando hacia arriba cada
	// concepto, [1633.79, 1633.81]. Las bases suman 10211.22 y con 6
	// decimales admiten [10211.21, 10211.23].
	trasladoSuma := func(base, importe string) *Comprobante {
		return comprobante("4.0", &Impuestos{TotalImpuestosTrasladados: ptrDec(importe), Traslados: []Impuesto{tras(base, importe)}},
			ImpuestosConcepto{Traslados: []Impuesto{tras("6250.000625", "1000.0001")}},
			ImpuestosConcepto{Traslados: []Impuesto{tras("3961.219375", "633.7951")}})
	}
	// 500.0099 + 477.2317 = 977.2416: la letra, 977.24; el margen, [977.23,
	// 977.25].
	trasladoSumaBaja := func(importe string) *Comprobante {
		return comprobante("4.0", &Impuestos{TotalImpuestosTrasladados: ptrDec(importe), Traslados: []Impuesto{tras("6107.76", importe)}},
			ImpuestosConcepto{Traslados: []Impuesto{tras("3125.061875", "500.0099")}},
			ImpuestosConcepto{Traslados: []Impuesto{tras("2982.698125", "477.2317")}})
	}
	// Retenciones de ISR 10 % de 3.94375 cada una.
	retencionSuma := func(importe string, conceptos int) *Comprobante {
		cs := make([]ImpuestosConcepto, conceptos)
		for i := range cs {
			cs[i] = ImpuestosConcepto{Retenciones: []Impuesto{ret("39.4375", "3.94375")}}
		}
		return comprobante("4.0", &Impuestos{TotalImpuestosRetenidos: ptrDec(importe), Retenciones: []Impuesto{{Impuesto: "001", Importe: ptrDec(importe)}}}, cs...)
	}
	// 3.3: los traslados y retenciones del comprobante con 6 decimales.
	totalTraslados := func(total string) *Comprobante {
		return comprobante("3.3", &Impuestos{TotalImpuestosTrasladados: ptrDec(total), Traslados: []Impuesto{tras("2700.03", "432.004800")}})
	}
	totalRetenciones := func(total string) *Comprobante {
		return comprobante("3.3", &Impuestos{TotalImpuestosRetenidos: ptrDec(total), Retenciones: []Impuesto{{Impuesto: "001", Importe: ptrDec("3.943750")}}})
	}

	const no, ad, er = "", NivelAdvertencia, NivelError
	casos := []struct {
		nombre string
		c      *Comprobante
		codigo string
		attr   string // del hallazgo, si hay
		nivel  string
	}{
		// Base 328.90 × 0.160000 = 52.624: a los 6 decimales del Importe,
		// [52.623200, 52.624800]; a los de la moneda, [52.62, 52.63].
		{"impuesto en la letra", importeImpuesto("328.90", "52.623200"), ReglaImporteImpuesto, "", no},
		{"impuesto a 2 decimales", importeImpuesto("328.90", "52.62"), ReglaImporteImpuesto, "", no},
		{"impuesto en el margen", importeImpuesto("328.90", "52.620000"), ReglaImporteImpuesto, "Importe", ad},
		{"impuesto en el margen superior", importeImpuesto("328.90", "52.630000"), ReglaImporteImpuesto, "Importe", ad},
		{"impuesto fuera del margen", importeImpuesto("328.90", "52.610000"), ReglaImporteImpuesto, "Importe", er},
		{"impuesto sobre el margen", importeImpuesto("328.90", "52.640000"), ReglaImporteImpuesto, "Importe", er},
		// Base 1726.98: [276.316000, 276.317600]; moneda, [276.31, 276.32].
		{"otro impuesto en el margen", importeImpuesto("1726.98", "276.320000"), ReglaImporteImpuesto, "Importe", ad},
		{"otro impuesto fuera del margen", importeImpuesto("1726.98", "276.330000"), ReglaImporteImpuesto, "Importe", er},

		{"traslado en la letra", trasladoSuma("10211.22", "1633.80"), ReglaTrasladoSuma, "", no},
		{"traslado en el margen", trasladoSuma("10211.22", "1633.81"), ReglaTrasladoSuma, "Importe", ad},
		{"traslado fuera del margen", trasladoSuma("10211.22", "1633.82"), ReglaTrasladoSuma, "Importe", er},
		{"traslado bajo la letra", trasladoSuma("10211.22", "1633.79"), ReglaTrasladoSuma, "Importe", ad},
		{"traslado bajo el margen", trasladoSuma("10211.22", "1633.78"), ReglaTrasladoSuma, "Importe", er},
		{"base del traslado en el margen", trasladoSuma("10211.23", "1633.80"), ReglaTrasladoSuma, "Base", ad},
		{"base del traslado fuera del margen", trasladoSuma("10211.24", "1633.80"), ReglaTrasladoSuma, "Base", er},
		{"traslado redondeado", trasladoSumaBaja("977.24"), ReglaTrasladoSuma, "", no},
		{"traslado truncado por concepto", trasladoSumaBaja("977.23"), ReglaTrasladoSuma, "Importe", ad},
		{"traslado bajo el truncado", trasladoSumaBaja("977.22"), ReglaTrasladoSuma, "Importe", er},

		// 3.94375: la letra, 3.94; el margen, [3.94, 3.95].
		{"retención en la letra", retencionSuma("3.94", 1), ReglaRetencionSuma, "", no},
		{"retención en el margen", retencionSuma("3.95", 1), ReglaRetencionSuma, "Importe", ad},
		{"retención fuera del margen", retencionSuma("3.96", 1), ReglaRetencionSuma, "Importe", er},
		// 2 × 3.94375 = 7.8875: la letra, 7.89; el margen, [7.88, 7.90].
		{"dos retenciones en la letra", retencionSuma("7.89", 2), ReglaRetencionSuma, "", no},
		{"dos retenciones en el margen", retencionSuma("7.90", 2), ReglaRetencionSuma, "Importe", ad},
		{"dos retenciones fuera del margen", retencionSuma("7.91", 2), ReglaRetencionSuma, "Importe", er},

		// 432.004800: la letra, 432.0048 o 432.00; el margen, [432.00, 432.01].
		{"3.3 total de traslados en la letra", totalTraslados("432.00"), ReglaTotalTraslados, "", no},
		{"3.3 total de traslados exacto", totalTraslados("432.004800"), ReglaTotalTraslados, "", no},
		{"3.3 total de traslados en el margen", totalTraslados("432.01"), ReglaTotalTraslados, "TotalImpuestosTrasladados", ad},
		{"3.3 total de traslados fuera del margen", totalTraslados("432.02"), ReglaTotalTraslados, "TotalImpuestosTrasladados", er},
		{"3.3 total de retenciones en la letra", totalRetenciones("3.94"), ReglaTotalRetenciones, "", no},
		{"3.3 total de retenciones en el margen", totalRetenciones("3.95"), ReglaTotalRetenciones, "TotalImpuestosRetenidos", ad},
		{"3.3 total de retenciones fuera del margen", totalRetenciones("3.96"), ReglaTotalRetenciones, "TotalImpuestosRetenidos", er},
	}
	for _, c := range casos {
		var hs []Hallazgo
		for _, h := range ValidarReglas(c.c) {
			if h.Codigo == c.codigo {
				hs = append(hs, h)
			}
		}
		if c.nivel == no {
			if len(hs) != 0 {
				t.Errorf("%s: %+v, quiero ninguno", c.nombre, hs)
			}
			continue
		}
		if len(hs) != 1 {
			t.Errorf("%s: %+v, quiero uno", c.nombre, hs)
			continue
		}
		h := hs[0]
		r := reglaDe[c.codigo]
		if h.Nivel != c.nivel || h.Atributo != c.attr || h.Regla != r.Referencia || r.Margen == "" {
			t.Errorf("%s: %+v, quiero %s en %s", c.nombre, h, c.nivel, c.attr)
		}
		// El mensaje dice en qué nivel cae y cuál es el margen.
		quiero := "no cumple la letra del SAT, pero sí el margen de redondeo por "
		if c.nivel == er {
			quiero = "fuera también del margen de redondeo por "
		}
		if !strings.Contains(h.Mensaje, quiero) {
			t.Errorf("%s: mensaje %q, quiero que diga %q", c.nombre, h.Mensaje, quiero)
		}
	}
}
