package cfdi

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/OrtaMarco/cfdi-go/decimal"
)

// codigosDe lista los códigos de hs sin repetir. Un hallazgo con un nivel
// distinto del de su regla (el segundo nivel de Regla.Margen) sale como
// adv(codigo).
func codigosDe(hs []Hallazgo) []string {
	var out []string
	vistos := map[string]bool{}
	for _, h := range hs {
		c := h.Codigo
		if h.Nivel != reglaDe[h.Codigo].Nivel {
			c += "/" + h.Nivel
		}
		if !vistos[c] {
			vistos[c] = true
			out = append(out, c)
		}
	}
	sort.Strings(out)
	return out
}

// adv es el código de una regla de dos niveles cuando sale como
// advertencia.
func adv(codigo string) string { return codigo + "/" + NivelAdvertencia }

func reglasDe(t *testing.T, xml string) []Hallazgo {
	t.Helper()
	c, err := ParseBytes([]byte(xml))
	if err != nil {
		t.Fatal(err)
	}
	return ValidarReglas(c)
}

// Casos positivos: documentos que sus fuentes declaran válidos (CfdiUtils y
// nodecfdi los usan como "valid" en sus pruebas) y los sintéticos
// coherentes. Ninguna regla puede saltar.
func TestValidarReglasPositivos(t *testing.T) {
	archivos := []string{
		"testdata/publicos/eclipxe13-CfdiUtils/tests__assets__cfdi33-valid.xml",
		"testdata/publicos/eclipxe13-CfdiUtils/tests__assets__cfdi40-valid.xml",
		"testdata/publicos/nodecfdi-cfdi-validator/tests___files__cfdi40-valid.xml",
		"testdata/sinteticos/ingreso40-pue.xml",
		"testdata/sinteticos/ingreso33.xml",
		"testdata/sinteticos/egreso40-relacionados.xml",
		"testdata/sinteticos/traslado40.xml",
		"testdata/sinteticos/pago20.xml",
		"testdata/sinteticos/pago10.xml",
		"testdata/sinteticos/nomina12.xml",
		"testdata/sinteticos/cce20.xml",
		"testdata/sinteticos/implocal40.xml",
		"testdata/sinteticos/cartaporte31.xml",
		"testdata/sinteticos/sellados/ingreso40-pue.xml",
	}
	for _, p := range archivos {
		if hs := reglasDe(t, string(leerFixture(t, p))); len(hs) != 0 {
			t.Errorf("%s: %+v", p, hs)
		}
	}
}

type cambio struct{ viejo, nuevo string }

func aplicarCambios(t *testing.T, base string, cs []cambio) string {
	t.Helper()
	for _, c := range cs {
		if strings.Count(base, c.viejo) != 1 {
			t.Fatalf("%q aparece %d veces en la base", c.viejo, strings.Count(base, c.viejo))
		}
		base = strings.Replace(base, c.viejo, c.nuevo, 1)
	}
	return base
}

// Casos negativos sintéticos, uno o pocos por regla, con los límites
// exactos de la tolerancia. Cada caso lista exactamente los códigos que
// tienen que salir: ni uno de más.
func TestValidarReglasSinteticos(t *testing.T) {
	const (
		concepto2   = `ValorUnitario="200.00" Importe="500.00" ObjetoImp="02">`
		trasladoC2  = `<cfdi:Traslado Base="500.00" Impuesto="002" TipoFactor="Tasa" TasaOCuota="0.160000" Importe="80.00"/>`
		trasladoCom = `<cfdi:Traslado Base="1500.00" Impuesto="002" TipoFactor="Tasa" TasaOCuota="0.160000" Importe="240.00"/>`
		retencionC  = `<cfdi:Retencion Impuesto="001" Importe="100.00"/>`
		totalTras   = `TotalImpuestosTrasladados="240.00"`
		totalRet    = `TotalImpuestosRetenidos="100.00"`
		subTotal    = `SubTotal="1500.00"`
		total       = ` Total="1640.00"`
	)
	conImporte2 := func(importe, sub, tot string) []cambio {
		return []cambio{{concepto2, `ValorUnitario="200.00" Importe="` + importe + `" ObjetoImp="02">`}, {subTotal, `SubTotal="` + sub + `"`}, {total, ` Total="` + tot + `"`}}
	}
	conTraslado2 := func(importe, comp, tot string) []cambio {
		return []cambio{
			{trasladoC2, strings.Replace(trasladoC2, `Importe="80.00"`, `Importe="`+importe+`"`, 1)},
			{trasladoCom, strings.Replace(trasladoCom, `Importe="240.00"`, `Importe="`+comp+`"`, 1)},
			{totalTras, `TotalImpuestosTrasladados="` + comp + `"`}, {total, ` Total="` + tot + `"`},
		}
	}
	// El segundo concepto pasa a IVA exento; comprobanteExento es lo que
	// se agrega en el comprobante junto al IVA 16 %.
	exento40 := func(comprobanteExento string) []cambio {
		return []cambio{
			{trasladoC2, `<cfdi:Traslado Base="500.00" Impuesto="002" TipoFactor="Exento"/>`},
			{totalTras, `TotalImpuestosTrasladados="160.00"`}, {total, ` Total="1560.00"`},
			{trasladoCom, `<cfdi:Traslado Base="1000.00" Impuesto="002" TipoFactor="Tasa" TasaOCuota="0.160000" Importe="160.00"/>` + comprobanteExento},
		}
	}
	// Un segundo concepto de 3.3 con IVA exento.
	exento33 := []cambio{
		{`</cfdi:Concepto>
  </cfdi:Conceptos>`, `</cfdi:Concepto>
    <cfdi:Concepto ClaveProdServ="81112100" Cantidad="1" ClaveUnidad="E48" Descripcion="Servicio exento de prueba" ValorUnitario="100.00" Importe="100.00">
      <cfdi:Impuestos>
        <cfdi:Traslados>
          <cfdi:Traslado Base="100.00" Impuesto="002" TipoFactor="Exento"/>
        </cfdi:Traslados>
      </cfdi:Impuestos>
    </cfdi:Concepto>
  </cfdi:Conceptos>`},
		{`SubTotal="3000.00"`, `SubTotal="3100.00"`}, {`Total="3132.00"`, `Total="3232.00"`},
	}
	centavos := func(comp, tot string) []cambio {
		return []cambio{
			{trasladoC2, `<cfdi:Traslado Base="500.03" Impuesto="002" TipoFactor="Tasa" TasaOCuota="0.160000" Importe="80.00"/>`},
			{trasladoCom, `<cfdi:Traslado Base="1500.03" Impuesto="002" TipoFactor="Tasa" TasaOCuota="0.160000" Importe="` + comp + `"/>`},
			{totalTras, `TotalImpuestosTrasladados="` + comp + `"`}, {total, ` Total="` + tot + `"`},
		}
	}
	// Retención con Base 1000.04: 100.004, escrito 100.00.
	retencionCentavos := func(comp, tot string) []cambio {
		return []cambio{
			{`<cfdi:Retencion Base="1000.00" Impuesto="001" TipoFactor="Tasa" TasaOCuota="0.100000" Importe="100.00"/>`, `<cfdi:Retencion Base="1000.04" Impuesto="001" TipoFactor="Tasa" TasaOCuota="0.100000" Importe="100.00"/>`},
			{retencionC, `<cfdi:Retencion Impuesto="001" Importe="` + comp + `"/>`},
			{totalRet, `TotalImpuestosRetenidos="` + comp + `"`}, {total, ` Total="` + tot + `"`},
		}
	}
	seisDecimales := func(concepto, grupo, totalTraslados, tot string) []cambio {
		return []cambio{
			{`TasaOCuota="0.160000" Importe="432.00"/>
        </cfdi:Traslados>`, `TasaOCuota="0.160000" Importe="` + concepto + `"/>
        </cfdi:Traslados>`},
			{`TasaOCuota="0.160000" Importe="432.00"/>
    </cfdi:Traslados>`, `TasaOCuota="0.160000" Importe="` + grupo + `"/>
    </cfdi:Traslados>`},
			{`TotalImpuestosTrasladados="432.00"`, `TotalImpuestosTrasladados="` + totalTraslados + `"`},
			{`Base="2700.00"`, `Base="2700.03"`}, // 2700.03 × 0.16 = 432.0048
			{`Total="3132.00"`, `Total="` + tot + `"`},
		}
	}
	// La segunda factura en USD, exenta: su BaseDR 247.78 es 247.78 /
	// 0.055333 = 4477.978… en MXN (BaseP, CRP20268).
	pagoUSD := func(monto string) []cambio {
		return []cambio{
			{`MonedaDR="MXN" EquivalenciaDR="1" NumParcialidad="2" ImpSaldoAnt="800.00" ImpPagado="500.00" ImpSaldoInsoluto="300.00"`,
				`MonedaDR="USD" EquivalenciaDR="0.055333" NumParcialidad="2" ImpSaldoAnt="247.78" ImpPagado="247.78" ImpSaldoInsoluto="0.00"`},
			{`Monto="1660.00"`, `Monto="` + monto + `"`}, {`MontoTotalPagos="1660.00"`, `MontoTotalPagos="` + monto + `"`},
			{`BaseDR="500.00"`, `BaseDR="247.78"`}, {`BaseP="500.00"`, `BaseP="4477.98"`},
			{`TotalTrasladosBaseIVAExento="500.00"`, `TotalTrasladosBaseIVAExento="4477.98"`},
		}
	}
	// Pago en USD de las dos facturas en MXN, con EquivalenciaDR 20.
	pagoUSDdeMXN := func(monto, totalPagos string) []cambio {
		return []cambio{
			{`MonedaP="MXN" TipoCambioP="1" Monto="1660.00"`, `MonedaP="USD" TipoCambioP="20" Monto="` + monto + `"`},
			{`MonedaDR="MXN" EquivalenciaDR="1" NumParcialidad="1"`, `MonedaDR="MXN" EquivalenciaDR="20" NumParcialidad="1"`},
			{`MonedaDR="MXN" EquivalenciaDR="1" NumParcialidad="2"`, `MonedaDR="MXN" EquivalenciaDR="20" NumParcialidad="2"`},
			{`MontoTotalPagos="1660.00"`, `MontoTotalPagos="` + totalPagos + `"`},
			{`<pago20:TrasladoP BaseP="1000.00" ImpuestoP="002" TipoFactorP="Tasa" TasaOCuotaP="0.160000" ImporteP="160.00"/>`, `<pago20:TrasladoP BaseP="50.00" ImpuestoP="002" TipoFactorP="Tasa" TasaOCuotaP="0.160000" ImporteP="8.00"/>`},
			{`<pago20:TrasladoP BaseP="500.00" ImpuestoP="002" TipoFactorP="Exento"/>`, `<pago20:TrasladoP BaseP="25.00" ImpuestoP="002" TipoFactorP="Exento"/>`},
		}
	}
	ingreso := string(leerFixture(t, "testdata/sinteticos/ingreso40-pue.xml"))
	pago := string(leerFixture(t, "testdata/sinteticos/pago20.xml"))
	implocal := string(leerFixture(t, "testdata/sinteticos/implocal40.xml"))
	nomina := string(leerFixture(t, "testdata/sinteticos/nomina12.xml"))
	traslado := string(leerFixture(t, "testdata/sinteticos/traslado40.xml"))
	ingreso33 := string(leerFixture(t, "testdata/sinteticos/ingreso33.xml"))

	casos := []struct {
		nombre string
		base   string
		cs     []cambio
		quiero []string
	}{
		// Concepto/@Importe: Cantidad 2.5 (±0.05) × ValorUnitario 200.00
		// (±0.005) da [489.98775, 510.01275): truncado 489.98, hacia arriba
		// 510.02.
		{"importe en el límite superior", ingreso, conImporte2("510.02", "1510.02", "1650.02"), nil},
		{"importe sobre el límite superior", ingreso, conImporte2("510.03", "1510.03", "1650.03"), []string{ReglaImporteConcepto}},
		{"importe en el límite inferior", ingreso, conImporte2("489.98", "1489.98", "1629.98"), nil},
		{"importe bajo el límite inferior", ingreso, conImporte2("489.97", "1489.97", "1629.97"), []string{ReglaImporteConcepto}},
		// Traslado/@Importe: Base 500.00 (±0.005) × 0.16 da [79.9992, 80.0008).
		{"impuesto en el límite superior", ingreso, conTraslado2("80.01", "240.01", "1640.01"), nil},
		{"impuesto sobre el límite superior", ingreso, conTraslado2("80.02", "240.02", "1640.02"), []string{ReglaImporteImpuesto}},
		{"impuesto en el límite inferior", ingreso, conTraslado2("79.99", "239.99", "1639.99"), nil},
		{"impuesto bajo el límite inferior", ingreso, conTraslado2("79.98", "239.98", "1639.98"), []string{ReglaImporteImpuesto}},
		{"subtotal", ingreso, []cambio{{subTotal, `SubTotal="1500.01"`}, {total, ` Total="1640.01"`}}, []string{ReglaSubTotal}},
		{"total", ingreso, []cambio{{total, ` Total="1640.01"`}}, []string{ReglaTotal}},
		{"total que suma las retenciones", ingreso, []cambio{{total, ` Total="1840.00"`}}, []string{ReglaTotal}},
		{"total de traslados", ingreso, []cambio{{totalTras, `TotalImpuestosTrasladados="240.01"`}, {total, ` Total="1640.01"`}}, []string{ReglaTotalTraslados}},
		{"total de retenciones", ingreso, []cambio{{totalRet, `TotalImpuestosRetenidos="100.01"`}, {total, ` Total="1639.99"`}}, []string{ReglaTotalRetenciones}},
		{"suma de un traslado", ingreso, []cambio{{trasladoCom, strings.Replace(trasladoCom, `Importe="240.00"`, `Importe="240.01"`, 1)}, {totalTras, `TotalImpuestosTrasladados="240.01"`}, {total, ` Total="1640.01"`}}, []string{ReglaTrasladoSuma}},
		{"base de un traslado", ingreso, []cambio{{trasladoCom, strings.Replace(trasladoCom, `Base="1500.00"`, `Base="1500.01"`, 1)}}, []string{ReglaTrasladoSuma}},
		{"tasa distinta en un concepto", ingreso, []cambio{{trasladoC2, `<cfdi:Traslado Base="500.00" Impuesto="002" TipoFactor="Tasa" TasaOCuota="0.080000" Importe="40.00"/>`}}, []string{ReglaTrasladoFaltante, ReglaTrasladoSuma}},
		{"tasa escrita distinto se agrupa igual", ingreso, []cambio{{trasladoC2, strings.Replace(trasladoC2, `TasaOCuota="0.160000"`, `TasaOCuota="0.16"`, 1)}}, nil},
		{"traslado sobrante", ingreso, []cambio{{trasladoCom, trasladoCom + `<cfdi:Traslado Base="0.00" Impuesto="003" TipoFactor="Tasa" TasaOCuota="0.080000" Importe="0.00"/>`}}, []string{ReglaTrasladoSobrante}},
		// IEPS 3 % en el comprobante sin ningún concepto que lo traiga:
		// CfdiUtils lo marca (SUMAS07, «No existen más nodos de impuestos
		// trasladados en el comprobante de los que se han calculado»).
		{"IEPS en el comprobante sin concepto", ingreso, []cambio{
			{trasladoCom, trasladoCom + `<cfdi:Traslado Base="100.00" Impuesto="003" TipoFactor="Tasa" TasaOCuota="0.030000" Importe="3.00"/>`},
			{totalTras, `TotalImpuestosTrasladados="243.00"`}, {total, ` Total="1643.00"`},
		}, []string{ReglaTrasladoSobrante}},
		// Con el IEPS en un concepto, aunque la tasa se escriba distinto
		// (CfdiUtils agrupa por el texto; aquí, por valor).
		{"IEPS en un concepto y en el comprobante", ingreso, []cambio{
			{trasladoC2, trasladoC2 + `<cfdi:Traslado Base="100.00" Impuesto="003" TipoFactor="Tasa" TasaOCuota="0.03" Importe="3.00"/>`},
			{trasladoCom, trasladoCom + `<cfdi:Traslado Base="100.00" Impuesto="003" TipoFactor="Tasa" TasaOCuota="0.030000" Importe="3.00"/>`},
			{totalTras, `TotalImpuestosTrasladados="243.00"`}, {total, ` Total="1643.00"`},
		}, nil},
		{"suma de una retención", ingreso, []cambio{{retencionC, `<cfdi:Retencion Impuesto="001" Importe="100.01"/>`}, {totalRet, `TotalImpuestosRetenidos="100.01"`}, {total, ` Total="1639.99"`}}, []string{ReglaRetencionSuma}},
		{"retención sobrante", ingreso, []cambio{{retencionC, retencionC + `<cfdi:Retencion Impuesto="002" Importe="0.00"/>`}}, []string{ReglaRetencionSobrante}},
		{"retención faltante", ingreso, []cambio{{"<cfdi:Retenciones>\n      " + retencionC + "\n    </cfdi:Retenciones>", ""}, {totalRet + " ", ""}, {total, ` Total="1740.00"`}}, []string{ReglaRetencionFaltante}},
		{"traslados sin total", ingreso, []cambio{{" " + totalTras, ""}, {total, ` Total="1400.00"`}}, []string{ReglaImpuestosSinTotal}},
		{"descuento del concepto mayor que su importe", ingreso, []cambio{{concepto2, `ValorUnitario="200.00" Importe="500.00" Descuento="500.01" ObjetoImp="02">`}}, []string{ReglaDescuentoConcepto, ReglaDescuento}},
		{"descuento que no es la suma", ingreso, []cambio{{concepto2, `ValorUnitario="200.00" Importe="500.00" Descuento="100.00" ObjetoImp="02">`}, {subTotal, subTotal + ` Descuento="100.01"`}, {total, ` Total="1539.99"`}}, []string{ReglaDescuento}},
		{"descuento que sí es la suma", ingreso, []cambio{{concepto2, `ValorUnitario="200.00" Importe="500.00" Descuento="100.00" ObjetoImp="02">`}, {subTotal, subTotal + ` Descuento="100.00"`}, {total, ` Total="1540.00"`}}, nil},
		{"decimales de la moneda", ingreso, []cambio{{subTotal, `SubTotal="1500.000"`}}, []string{ReglaDecimalesMoneda}},
		{"decimales de un traslado", ingreso, []cambio{{trasladoCom, strings.Replace(trasladoCom, `Importe="240.00"`, `Importe="240.000"`, 1)}}, []string{ReglaDecimalesImpuesto}},
		{"decimales de una retención", ingreso, []cambio{{retencionC, `<cfdi:Retencion Impuesto="001" Importe="100.000"/>`}}, []string{ReglaDecimalesImpuesto}},
		{"3.3 traslado del comprobante con 6 decimales", ingreso33, []cambio{{`TasaOCuota="0.160000" Importe="432.00"/>
    </cfdi:Traslados>`, `TasaOCuota="0.160000" Importe="432.000000"/>
    </cfdi:Traslados>`}}, []string{ReglaDecimalesImpuesto33}},
		{"total de traslados con decimales de más", ingreso, []cambio{{totalTras, `TotalImpuestosTrasladados="240.000"`}}, []string{ReglaDecimalesMoneda}},
		{"XXX con total 0.00", traslado, []cambio{{` Total="0"`, ` Total="0.00"`}}, []string{ReglaDecimalesMoneda}},
		{"XXX con subtotal 0.0", pago, []cambio{{`SubTotal="0"`, `SubTotal="0.0"`}}, []string{ReglaDecimalesMoneda}},
		{"moneda con espacios", ingreso, []cambio{{`Moneda="MXN"`, `Moneda=" MXN "`}}, nil},
		{"tipo de cambio en MXN", ingreso, []cambio{{`Moneda="MXN"`, `Moneda="MXN" TipoCambio="1.000001"`}}, []string{ReglaTipoCambioMXN}},
		{"tipo de cambio 1 en MXN", ingreso, []cambio{{`Moneda="MXN"`, `Moneda="MXN" TipoCambio="1.000000"`}}, nil},
		{"USD sin tipo de cambio", ingreso, []cambio{{`Moneda="MXN"`, `Moneda="USD"`}}, []string{ReglaTipoCambioRequerido}},
		{"moneda de 3 decimales", ingreso, []cambio{{`Moneda="MXN"`, `Moneda="BHD" TipoCambio="45.5"`}, {subTotal, `SubTotal="1500.000"`}}, nil},
		{"moneda sin decimales", ingreso, []cambio{{`Moneda="MXN"`, `Moneda="JPY" TipoCambio="0.12"`}}, []string{ReglaDecimalesMoneda, ReglaDecimalesImpuesto}},
		{"objeto de impuesto 01 con impuestos", ingreso, []cambio{{concepto2, `ValorUnitario="200.00" Importe="500.00" ObjetoImp="01">`}}, []string{ReglaObjetoImp}},
		{"objeto de impuesto 02 sin impuestos", traslado, []cambio{{`ObjetoImp="01"`, `ObjetoImp="02"`}}, []string{ReglaObjetoImp}},
		{"exento con importe", ingreso, []cambio{{trasladoC2, `<cfdi:Traslado Base="500.00" Impuesto="002" TipoFactor="Exento" Importe="0.00"/>`}}, []string{ReglaExentoConImporte, ReglaTrasladoFaltante, ReglaTrasladoSuma}},
		{"tasa sin importe", ingreso, []cambio{{trasladoC2, `<cfdi:Traslado Base="500.00" Impuesto="002" TipoFactor="Tasa" TasaOCuota="0.160000"/>`}}, []string{ReglaTasaSinImporte, ReglaTrasladoSuma}},
		{"retención exenta", ingreso, []cambio{{`<cfdi:Retencion Base="1000.00" Impuesto="001" TipoFactor="Tasa" TasaOCuota="0.100000" Importe="100.00"/>`, `<cfdi:Retencion Base="1000.00" Impuesto="001" TipoFactor="Exento" TasaOCuota="0.100000" Importe="100.00"/>`}}, []string{ReglaRetencionExenta}},
		// Por TipoDeComprobante.
		{"traslado con subtotal", traslado, []cambio{{`SubTotal="0"`, `SubTotal="1"`}}, []string{ReglaTipoCompSubTotalCero, ReglaTotal}}, // CFDI40108 solo en I, E y N,
		{"traslado con total", traslado, []cambio{{` Total="0"`, ` Total="1"`}}, []string{ReglaTipoCompTotalCero, ReglaTotal}},
		{"traslado con forma, método y condiciones", traslado, []cambio{{`SubTotal="0"`, `CondicionesDePago="X" FormaPago="01" MetodoPago="PUE" SubTotal="0"`}}, []string{ReglaTipoCompCondiciones, ReglaTipoCompFormaPago, ReglaTipoCompMetodoPago}},
		{"traslado con descuento", traslado, []cambio{{`ValorUnitario="0.00" Importe="0.00" ObjetoImp="01">`, `ValorUnitario="0.00" Importe="0.00" Descuento="0.00" ObjetoImp="01">`}, {`SubTotal="0"`, `SubTotal="0" Descuento="0"`}}, []string{ReglaTipoCompDescuento}},
		{"traslado con impuestos", traslado, []cambio{{`</cfdi:Conceptos>`, `</cfdi:Conceptos><cfdi:Impuestos/>`}}, []string{ReglaTipoCompImpuestos}},
		{"ingreso con valor unitario cero", ingreso, []cambio{{`ValorUnitario="200.00" Importe="500.00"`, `ValorUnitario="0" Importe="0"`}, {subTotal, `SubTotal="1000.00"`}, {total, ` Total="1140.00"`}}, []string{ReglaTipoCompValorUnitario}},
		{"nómina en USD", nomina, []cambio{{`Moneda="MXN"`, `Moneda="USD" TipoCambio="17.5"`}}, []string{ReglaTipoCompMonedaNomina}},
		// Impuestos locales.
		{"total de traslados locales", implocal, []cambio{{`TotaldeTraslados="50.00"`, `TotaldeTraslados="51.00"`}}, []string{ReglaImpLocalTotales, ReglaTotal}},
		{"total sin los impuestos locales", implocal, []cambio{{`Total="1205.00"`, `Total="1160.00"`}}, []string{ReglaTotal}},
		// Pagos 2.0.
		{"pago con moneda", pago, []cambio{{`Moneda="XXX"`, `Moneda="MXN"`}}, []string{ReglaPagoComprobante}},
		{"pago con otra descripción", pago, []cambio{{`Descripcion="Pago"`, `Descripcion="Abono"`}}, []string{ReglaPagoComprobante}},
		{"pago con importe", pago, []cambio{{`ValorUnitario="0" Importe="0"`, `ValorUnitario="0" Importe="1"`}}, []string{ReglaPagoComprobante}}, // 1 × 0, escritos sin decimales, admite [0, 1]
		{"pago con importe fuera de tolerancia", pago, []cambio{{`ValorUnitario="0" Importe="0"`, `ValorUnitario="0" Importe="2"`}}, []string{ReglaPagoComprobante, ReglaImporteConcepto}},
		{"saldo insoluto", pago, []cambio{{`ImpSaldoInsoluto="300.00"`, `ImpSaldoInsoluto="300.01"`}}, []string{ReglaPagoSaldoInsoluto}},
		{"equivalencia con la misma moneda", pago, []cambio{{`MonedaDR="MXN" EquivalenciaDR="1" NumParcialidad="1"`, `MonedaDR="MXN" EquivalenciaDR="2" NumParcialidad="1"`}}, []string{ReglaPagoEquivalencia, ReglaPagoImpuestosPSuma}}, // ImporteP se convierte con la EquivalenciaDR escrita
		{"otra moneda sin equivalencia", pago, []cambio{{`MonedaDR="MXN" EquivalenciaDR="1" NumParcialidad="1"`, `MonedaDR="USD" NumParcialidad="1"`}}, []string{ReglaPagoEquivalencia}},
		{"monto menor que lo pagado", pago, []cambio{{`Monto="1660.00"`, `Monto="1659.99"`}, {`MontoTotalPagos="1660.00"`, `MontoTotalPagos="1659.99"`}}, []string{ReglaPagoMonto}},
		{"tipo de cambio del pago en MXN", pago, []cambio{{`TipoCambioP="1"`, `TipoCambioP="1.5"`}}, []string{ReglaPagoTipoCambio, ReglaPagosTotales}},
		{"monto total de pagos", pago, []cambio{{`MontoTotalPagos="1660.00"`, `MontoTotalPagos="1660.01"`}}, []string{ReglaPagosTotales}},
		{"falta un total", pago, []cambio{{` TotalTrasladosBaseIVAExento="500.00"`, ``}}, []string{ReglaPagosTotales}},
		{"total que ningún pago tiene", pago, []cambio{{`MontoTotalPagos="1660.00"`, `MontoTotalPagos="1660.00" TotalRetencionesIVA="1.00"`}}, []string{ReglaPagosTotales}},
		{"pago en USD de facturas en MXN", pago, []cambio{
			{`MonedaP="MXN" TipoCambioP="1" Monto="1660.00"`, `MonedaP="USD" TipoCambioP="20" Monto="83.00"`},
			{`MonedaDR="MXN" EquivalenciaDR="1" NumParcialidad="1"`, `MonedaDR="MXN" EquivalenciaDR="20" NumParcialidad="1"`},
			{`MonedaDR="MXN" EquivalenciaDR="1" NumParcialidad="2"`, `MonedaDR="MXN" EquivalenciaDR="20" NumParcialidad="2"`},
			{`<pago20:TrasladoP BaseP="1000.00" ImpuestoP="002" TipoFactorP="Tasa" TasaOCuotaP="0.160000" ImporteP="160.00"/>`, `<pago20:TrasladoP BaseP="50.00" ImpuestoP="002" TipoFactorP="Tasa" TasaOCuotaP="0.160000" ImporteP="8.00"/>`},
			{`<pago20:TrasladoP BaseP="500.00" ImpuestoP="002" TipoFactorP="Exento"/>`, `<pago20:TrasladoP BaseP="25.00" ImpuestoP="002" TipoFactorP="Exento"/>`},
		}, nil},
		// Con EquivalenciaDR 20 (sin decimales) el margen es ±0.5: los
		// límites inferiores son 1159.995 / 20.4999999999 y 499.995 /
		// 20.4999999999, truncados 56.58 y 24.39.
		{"pago en USD en el límite inferior", pago, pagoUSDdeMXN("80.97", "1619.40"), nil},
		{"pago en USD bajo el límite inferior", pago, pagoUSDdeMXN("80.96", "1619.20"), []string{ReglaPagoMonto}},
		// Sumas al centavo: el grupo es el redondeo de la suma de los
		// Importe escritos (CFDI40221, «Redondear al final del cálculo y
		// no antes»). Con Base 500.03 el concepto escribe 80.00 (80.0048):
		// el grupo es 240.00 y no 240.01, aunque Σ Base × Tasa dé 240.0048.
		// 240.01 cae en el margen de redondeo por concepto (80.0048 hacia
		// arriba es 80.01): advertencia.
		{"traslado que suma Base × Tasa sin redondear", ingreso, centavos("240.01", "1640.01"), []string{adv(ReglaTrasladoSuma)}},
		{"traslado con la suma de lo escrito", ingreso, centavos("240.00", "1640.00"), nil},
		{"traslado bajo la suma", ingreso, centavos("239.99", "1639.99"), []string{ReglaTrasladoSuma}},
		{"retención que suma Base × Tasa sin redondear", ingreso, retencionCentavos("100.01", "1639.99"), []string{adv(ReglaRetencionSuma)}},
		{"retención con la suma de lo escrito", ingreso, retencionCentavos("100.00", "1640.00"), nil},
		// 3.3 con 6 decimales en los conceptos y en el comprobante: el
		// grupo es la suma exacta o su redondeo; TotalImpuestosTrasladados,
		// el redondeo de la suma de los grupos. 432.01 es 432.0048 hacia
		// arriba, el margen: advertencia; 432.02, error.
		{"3.3 con 6 decimales y totales redondeados", ingreso33, seisDecimales("432.004800", "432.004800", "432.00", "3132.00"), []string{ReglaDecimalesImpuesto33}},
		{"3.3 grupo redondeado", ingreso33, seisDecimales("432.004800", "432.00", "432.00", "3132.00"), nil},
		{"3.3 grupo redondeado hacia arriba", ingreso33, seisDecimales("432.004800", "432.01", "432.01", "3132.01"), []string{adv(ReglaTrasladoSuma)}},
		{"3.3 total de traslados redondeado hacia arriba", ingreso33, seisDecimales("432.004800", "432.004800", "432.01", "3132.01"), []string{ReglaDecimalesImpuesto33, adv(ReglaTotalTraslados)}},
		{"3.3 total de traslados a 2 centavos", ingreso33, seisDecimales("432.004800", "432.004800", "432.02", "3132.02"), []string{ReglaDecimalesImpuesto33, ReglaTotalTraslados}},
		{"3.3 grupo a 2 centavos", ingreso33, seisDecimales("432.004800", "432.02", "432.02", "3132.02"), []string{ReglaTrasladoSuma}},
		// Límites a los decimales del Importe (estándar 4.0 p. 76): 2.5 ×
		// 200.00 escrito sin decimales admite [489, 511].
		{"importe sin decimales en su límite", ingreso, conImporte2("511", "1511.00", "1651.00"), nil},
		{"importe sin decimales sobre su límite", ingreso, conImporte2("512", "1512.00", "1652.00"), []string{ReglaImporteConcepto}},
		// Base 500.00 × 0.16 con el Importe escrito sin decimales admite
		// [79, 81].
		{"impuesto sin decimales en su límite", ingreso, conTraslado2("81", "241.00", "1641.00"), nil},
		{"impuesto sin decimales sobre su límite", ingreso, conTraslado2("82", "242.00", "1642.00"), []string{ReglaImporteImpuesto}},
		// Exento. En 3.3 el Traslado del comprobante exige TasaOCuota e
		// Importe (cfdv33.xsd): un concepto exento no tiene grupo en el
		// comprobante. En 4.0 sí (CfdiUtils 2.23.3).
		{"3.3 con un concepto exento", ingreso33, exento33, nil},
		{"3.3 con el exento también en el comprobante", ingreso33, append(exento33[:len(exento33):len(exento33)],
			cambio{`TasaOCuota="0.160000" Importe="432.00"/>
    </cfdi:Traslados>`, `TasaOCuota="0.160000" Importe="432.00"/>
      <cfdi:Traslado Impuesto="002" TipoFactor="Exento"/>
    </cfdi:Traslados>`}), []string{ReglaTrasladoSobrante}},
		{"4.0 con el exento en el comprobante", ingreso, exento40(`<cfdi:Traslado Base="500.00" Impuesto="002" TipoFactor="Exento"/>`), nil},
		{"4.0 sin el exento en el comprobante", ingreso, exento40(""), []string{ReglaTrasladoFaltante}},
		{"4.0 con la base del exento mal", ingreso, exento40(`<cfdi:Traslado Base="500.01" Impuesto="002" TipoFactor="Exento"/>`), []string{ReglaTrasladoSuma}},
		// Pago en MXN de una factura en MXN (1160.00) y otra en USD: la
		// suma 1160 + 247.78 / 0.055333 = 5637.978… se trunca a 5637.97
		// (guía Pagos 2.0 p. 32, ii). Límites inferiores (p. 33, CRP20277:
		// la de MXN con EquivalenciaDR 1 de 10 decimales), truncados:
		// 1159.995 / 0.99999999995 → 1159.99 y 247.775 / 0.0553334999 →
		// 4477.84; suman 5637.83.
		{"pago en la suma truncada", pago, pagoUSD("5637.97"), nil},
		{"pago bajo la suma pero dentro del intervalo", pago, pagoUSD("5637.96"), nil},
		{"pago en el límite inferior", pago, pagoUSD("5637.83"), nil},
		{"pago bajo el límite inferior", pago, pagoUSD("5637.82"), []string{ReglaPagoMonto}},
		{"pago sin monto no se evalúa", pago, []cambio{{` Monto="1660.00"`, ``}}, nil},
		// CFDI 3.3.
		{"3.3 sin tipo de cambio", ingreso33, []cambio{{` TipoCambio="19.8765"`, ``}}, []string{ReglaTipoCambioRequerido}},
		{"3.3 sin objeto de impuesto", ingreso33, []cambio{{`Descuento="300.00">`, `Descuento="300.00" ObjetoImp="01">`}}, nil},
	}
	for _, c := range casos {
		hs := reglasDe(t, aplicarCambios(t, c.base, c.cs))
		got := codigosDe(hs)
		quiero := append([]string(nil), c.quiero...)
		sort.Strings(quiero)
		if strings.Join(got, ",") != strings.Join(quiero, ",") {
			t.Errorf("%s: %v, quiero %v\n%+v", c.nombre, got, quiero, hs)
		}
		for _, h := range hs {
			r := reglaDe[h.Codigo]
			segundoNivel := r.Margen != "" && h.Nivel == NivelAdvertencia
			if (h.Nivel != r.Nivel && !segundoNivel) || h.Regla == "" || h.Mensaje == "" || h.Ruta == "" || !strings.HasPrefix(h.Ruta, "/Comprobante") {
				t.Errorf("%s: hallazgo incompleto %+v", c.nombre, h)
			}
			if h.Atributo != "" && !strings.HasSuffix(h.Ruta, "/@"+h.Atributo) {
				t.Errorf("%s: Ruta %q y Atributo %q no coinciden", c.nombre, h.Ruta, h.Atributo)
			}
		}
	}
}

// Las rutas y valores de los hallazgos de reglas siguen el contrato de
// Advertencia.
func TestValidarReglasRutas(t *testing.T) {
	ingreso := string(leerFixture(t, "testdata/sinteticos/ingreso40-pue.xml"))
	hs := reglasDe(t, strings.Replace(ingreso, `Importe="80.00"`, `Importe="80.02"`, 1))
	h := buscarHallazgo(hs, ReglaImporteImpuesto)
	if h == nil || h.Ruta != "/Comprobante/Conceptos/Concepto[2]/Impuestos/Traslados/Traslado[1]/@Importe" ||
		h.Atributo != "Importe" || h.Valor != "80.02" || !strings.Contains(h.Regla, "Anexo 20") || h.Nodo() != "/Comprobante/Conceptos/Concepto[2]/Impuestos/Traslados/Traslado[1]" {
		t.Errorf("%+v", hs)
	}
	pago := string(leerFixture(t, "testdata/sinteticos/pago20.xml"))
	hs = reglasDe(t, strings.Replace(pago, `ImpSaldoInsoluto="300.00"`, `ImpSaldoInsoluto="300.01"`, 1))
	if len(hs) != 1 || hs[0].Ruta != "/Comprobante/Complemento/Pagos/Pago[1]/DoctoRelacionado[2]/@ImpSaldoInsoluto" || hs[0].Valor != "300.01" {
		t.Errorf("%+v", hs)
	}
	hs = reglasDe(t, strings.Replace(pago, `MontoTotalPagos="1660.00"`, `MontoTotalPagos="1660.01"`, 1))
	if len(hs) != 1 || hs[0].Ruta != "/Comprobante/Complemento/Pagos/Totales/@MontoTotalPagos" {
		t.Errorf("%+v", hs)
	}
}

func TestValidarReglasVersiones(t *testing.T) {
	for _, v := range []string{"3.2", "3.0", "", "5.0"} {
		hs := ValidarReglas(&Comprobante{Version: v})
		if len(hs) != 1 || hs[0].Codigo != ReglaVersionNoSoportada || hs[0].Nivel != NivelAviso {
			t.Errorf("versión %q: %+v", v, hs)
		}
	}
	if hs := ValidarReglas(&Comprobante{Version: "4.0"}); hs == nil || len(hs) != 0 {
		t.Errorf("comprobante vacío: %+v (tiene que ser [] y no nil)", hs)
	}
	// Las reglas de 4.0 no se aplican a 3.3, aunque traiga Pagos 2.0.
	c := &Comprobante{Version: "3.3", Pagos: &Pagos{Version: "2.0", Pago: []Pago{{DoctoRelacionado: []DoctoRelacionado{{
		ImpSaldoAnt: ptrDec("10"), ImpPagado: ptrDec("1"), ImpSaldoInsoluto: ptrDec("1")}}}}}}
	if hs := ValidarReglas(c); len(hs) != 0 {
		t.Errorf("3.3 con Pagos 2.0: %+v", hs)
	}
}

func ptrDec(s string) *decimal.Decimal {
	d := decimal.MustParse(s)
	return &d
}

// Casos de la prueba de CfdiUtils ComprobanteTipoCambioTest
// (tests/CfdiUtilsTests/Validate/Cfdi33/Standard/ComprobanteTipoCambioTest.php,
// providerMonedaWithValidValues y providerMonedaWithInvalidValues). Los que
// en CfdiUtils fallan por el patrón del TipoCambio ("abc", "0.", "0..0",
// más de 18 enteros o de 6 decimales) aquí son del esquema
// (ValidarEsquema) y no se repiten; un TipoCambio ilegible queda nil en el
// modelo, que es el caso "sin tipo de cambio".
func TestValidarReglasTipoCambioCfdiUtils(t *testing.T) {
	casos := []struct {
		moneda, tipoCambio, quiero string
	}{
		{"MXN", "1", ""},
		{"MXN", "1.000000", ""},
		{"MXN", "", ""},
		{"XXX", "", ""},
		{"USD", "10.0", ""},
		{"USD", "20", ""},
		{"USD", "0005.10000", ""},
		{"USD", "123456789012345678.0", ""},
		{"USD", "0.123456", ""},
		{"MXN", "1.000001", ReglaTipoCambioMXN},
		{"MXN", "0.999999", ReglaTipoCambioMXN},
		{"MXN", "10.0", ReglaTipoCambioMXN},
		{"XXX", "10.0", ReglaTipoCambioXXX},
		{"USD", "", ReglaTipoCambioRequerido},
	}
	for _, c := range casos {
		cp := &Comprobante{Version: "4.0", Moneda: c.moneda}
		if c.tipoCambio != "" {
			cp.TipoCambio = ptrDec(c.tipoCambio)
		}
		hs := ValidarReglas(cp)
		var got []string
		for _, h := range hs {
			if strings.HasPrefix(h.Codigo, "REGLA_TIPOCAMBIO") {
				got = append(got, h.Codigo)
			}
		}
		if strings.Join(got, ",") != c.quiero {
			t.Errorf("%s %q: %v, quiero %q", c.moneda, c.tipoCambio, got, c.quiero)
		}
	}
}

// Casos de CfdiUtils ComprobanteDescuentoTest
// (tests/CfdiUtilsTests/Validate/Cfdi33/Standard/ComprobanteDescuentoTest.php):
// pares (Descuento, SubTotal). Los negativos y vacíos de esa prueba los
// rechaza el esquema.
func TestValidarReglasDescuentoCfdiUtils(t *testing.T) {
	casos := []struct {
		descuento, subtotal string
		mal                 bool
	}{
		{"0", "1", false}, {"1", "1", false}, {"0.000000", "0.000001", false}, {"0", "0", false}, {"1.00", "1.01", false},
		{"1", "0", true}, {"0.000001", "0.000000", true},
	}
	for _, c := range casos {
		hs := ValidarReglas(&Comprobante{Version: "4.0", Descuento: ptrDec(c.descuento), SubTotal: ptrDec(c.subtotal)})
		if TieneHallazgo(hs, ReglaDescuentoMayor) != c.mal {
			t.Errorf("Descuento %s, SubTotal %s: %+v", c.descuento, c.subtotal, hs)
		}
	}
}

var documentosSAT = map[string]bool{
	DocMatriz40: true, DocEstandar40: true, DocGuia40: true, DocMatriz33: true, DocEstandar33: true, DocGuia33: true,
	DocMatrizPagos20: true, DocGuiaPagos20: true, DocEstPagos20: true, DocMatrizPagos10: true, DocGuiaPagos10: true, DocImpLocal: true,
	DocMatrizCCE20: true, DocGuiaCCE20: true, DocEstCCE20: true, DocMatrizNomina12: true, DocGuiaNomina12: true,
	DocMatrizNomina12B: true, DocMatrizNomina12C: true, DocEstNomina12: true,
	DocMatrizRet20: true, DocEstRet20: true, DocEstRet10: true, DocCatRet20: true,
}

// Los códigos de las reglas son únicos, cada uno con referencia y
// descripción, y en ningún fixture sale un código que no esté en la lista.
func TestReglasDocumentadas(t *testing.T) {
	vistos := map[string]bool{}
	for _, r := range Reglas {
		if vistos[r.Codigo] || !strings.HasPrefix(r.Codigo, "REGLA_") || r.Descripcion == "" || (r.Referencia == "" && r.Codigo != ReglaVersionNoSoportada) {
			t.Errorf("regla mal documentada: %+v", r)
		}
		// Solo el aviso de versión es aviso; una regla sin fundamento en un
		// documento del SAT es advertencia; un error tiene que citarlo.
		quiero := r.Nivel
		switch {
		case r.Codigo == ReglaVersionNoSoportada:
			quiero = NivelAviso
		case len(r.Citas) == 0:
			quiero = NivelAdvertencia
		case r.Nivel != NivelAdvertencia:
			quiero = NivelError
		}
		if r.Nivel != quiero {
			t.Errorf("%s: nivel %q, quiero %q", r.Codigo, r.Nivel, quiero)
		}
		for _, c := range r.Citas {
			if !documentosSAT[c.Documento] || (c.Codigo == "") == (c.Pagina == 0) || !strings.Contains(r.Referencia, c.String()) {
				t.Errorf("%s: cita mal formada %+v", r.Codigo, c)
			}
		}
		vistos[r.Codigo] = true
	}
	if len(CodigosReglas) != len(Reglas) {
		t.Error("CodigosReglas no es la lista de Reglas")
	}
	// La tabla del README lista cada regla.
	readme, err := os.ReadFile("README.md")
	if err != nil {
		t.Fatal(err)
	}
	dosNiveles := 0
	for _, r := range Reglas[1:] {
		fila := "| `" + r.Codigo + "` | " + r.Nivel + " |"
		i := strings.Index(string(readme), fila)
		if i < 0 {
			t.Errorf("%s (%s) no está en la tabla del README", r.Codigo, r.Nivel)
			continue
		}
		// Las de dos niveles son errores y lo dicen en su fila.
		fila = string(readme[i:])
		fila = fila[:strings.IndexByte(fila, '\n')]
		if (r.Margen != "") != strings.Contains(fila, "**Dos niveles**") || (r.Margen != "" && r.Nivel != NivelError) {
			t.Errorf("%s: Margen %q y la fila del README no coinciden: %s", r.Codigo, r.Margen, fila)
		}
		if r.Margen != "" {
			dosNiveles++
		}
		// Lo que baja a advertencia por fecha o por salir solo de la guía
		// también lo dice la fila.
		if (r.Vigencia != "") != strings.Contains(fila, "**Vigencia**") || (r.SoloGuia != "") != strings.Contains(fila, "**Solo guía**") {
			t.Errorf("%s: Vigencia %q, SoloGuia %q y la fila del README no coinciden: %s", r.Codigo, r.Vigencia, r.SoloGuia, fila)
		}
	}
	if dosNiveles != 9 {
		t.Errorf("%d reglas de dos niveles, quiero 9", dosNiveles)
	}
	err = filepath.WalkDir("testdata", func(p string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(p, ".xml") {
			return err
		}
		b, _ := os.ReadFile(p)
		c, err := ParseBytes(b)
		if err != nil {
			return nil
		}
		for _, h := range ValidarReglas(c) {
			if !vistos[h.Codigo] {
				t.Errorf("%s: código sin documentar %s", p, h.Codigo)
			}
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

// Las monedas con decimales son exactamente las del catálogo c_Moneda del
// XSD del SAT.
func TestMonedasCatalogo(t *testing.T) {
	catCatCFDI_cMoneda.contiene("")
	if len(catCatCFDI_cMoneda.lista) != len(decimalesMoneda) {
		t.Errorf("c_Moneda tiene %d claves y decimalesMoneda %d", len(catCatCFDI_cMoneda.lista), len(decimalesMoneda))
	}
	for _, m := range catCatCFDI_cMoneda.lista {
		if _, ok := decimalesMoneda[m]; !ok {
			t.Errorf("%s sin decimales", m)
		}
	}
	for m, d := range map[string]int32{"MXN": 2, "USD": 2, "XXX": 0, "JPY": 0, "BHD": 3, "CLF": 4} {
		if decimalesMoneda[m] != d {
			t.Errorf("%s: %d decimales, quiero %d", m, decimalesMoneda[m], d)
		}
	}
}

func TestTruncarTecho(t *testing.T) {
	casos := []struct{ x, trunc, techo string }{
		{"1.239", "1.23", "1.24"}, {"1.230", "1.23", "1.23"}, {"1.235", "1.23", "1.24"},
		{"-1.239", "-1.23", "-1.23"}, {"0", "0", "0"}, {"0.001", "0", "0.01"},
	}
	for _, c := range casos {
		x := decimal.MustParse(c.x)
		if got := truncar(x, 2); !got.Equal(decimal.MustParse(c.trunc)) {
			t.Errorf("truncar(%s) = %s", c.x, got)
		}
		if got := techo(x, 2); !got.Equal(decimal.MustParse(c.techo)) {
			t.Errorf("techo(%s) = %s", c.x, got)
		}
	}
}
