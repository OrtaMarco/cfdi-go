package cfdi

import (
	"sort"
	"strings"
	"testing"

	"github.com/OrtaMarco/cfdi-go/decimal"
)

// Reglas del CFDI de pago y de Pagos 1.0 y 2.0 que llegaron en T7: cada
// una con un caso que la dispara y uno que no.
func TestValidarReglasPagos(t *testing.T) {
	pago := string(leerFixture(t, "testdata/sinteticos/pago20.xml"))
	pago10 := string(leerFixture(t, "testdata/sinteticos/pago10.xml"))
	const (
		pagoP     = `FormaDePagoP="03" MonedaP="MXN" TipoCambioP="1" Monto="1660.00" NumOperacion="OP-0001"`
		dr2       = `ImpSaldoAnt="800.00" ImpPagado="500.00" ImpSaldoInsoluto="300.00"`
		pago10P1  = `FormaDePagoP="03" MonedaP="USD" TipoCambioP="20.1234" Monto="150.00"`
		pago10P2  = `FormaDePagoP="03" MonedaP="MXN" Monto="152.16"`
		pago10DR1 = `MonedaDR="USD" MetodoDePagoDR="PPD" NumParcialidad="1"`
		pago10DR2 = `MonedaDR="USD" TipoCambioDR="20.1" MetodoDePagoDR="PPD" NumParcialidad="2" ImpSaldoAnt="7.57" ImpPagado="7.57" ImpSaldoInsoluto="0.00"/>`
		cadena    = ` CertPago="Q0VSVElGSUNBRE8gREUgUFJVRUJB" CadPago="Q0FERU5BIERFIFBSVUVCQQ==" SelloPago="U0VMTE8gREUgUFJVRUJB" TipoCadPago="01"`
	)
	conPago := func(nuevo string) []cambio { return []cambio{{pagoP, nuevo}} }
	casos := []struct {
		nombre string
		base   string
		cs     []cambio
		quiero []string
	}{
		// Pagos 2.0.
		{"forma de pago 99", pago, conPago(`FormaDePagoP="99" MonedaP="MXN" TipoCambioP="1" Monto="1660.00" NumOperacion="OP-0001"`), []string{ReglaPagoFormaPago}},
		{"MonedaP XXX", pago, conPago(`FormaDePagoP="03" MonedaP="XXX" Monto="1660" NumOperacion="OP-0001"`), []string{ReglaPagoMonedaXXX}},
		{"MonedaDR XXX", pago, []cambio{{`MonedaDR="MXN" EquivalenciaDR="1" NumParcialidad="1"`, `MonedaDR="XXX" EquivalenciaDR="1" NumParcialidad="1"`}}, []string{ReglaPagoMonedaXXX, ReglaPagoDecimales}},
		{"MXN sin TipoCambioP", pago, conPago(`FormaDePagoP="03" MonedaP="MXN" Monto="1660.00" NumOperacion="OP-0001"`), []string{ReglaPagoTipoCambio}},
		{"USD sin TipoCambioP", pago, conPago(`FormaDePagoP="03" MonedaP="USD" Monto="1660.00" NumOperacion="OP-0001"`), []string{ReglaPagoTipoCambio}},
		{"USD con TipoCambioP", pago, []cambio{{pagoP, `FormaDePagoP="03" MonedaP="USD" TipoCambioP="1" Monto="1660.00" NumOperacion="OP-0001"`}}, nil},
		{"monto cero", pago, []cambio{{pagoP, `FormaDePagoP="03" MonedaP="MXN" TipoCambioP="1" Monto="0.00" NumOperacion="OP-0001"`}, {`MontoTotalPagos="1660.00"`, `MontoTotalPagos="0.00"`}},
			[]string{ReglaPagoMayorCero, ReglaPagoMonto}},
		{"ImpPagado cero", pago, []cambio{{dr2, `ImpSaldoAnt="800.00" ImpPagado="0.00" ImpSaldoInsoluto="800.00"`}}, []string{ReglaPagoMayorCero}},
		{"ImpSaldoAnt cero", pago, []cambio{{dr2, `ImpSaldoAnt="0.00" ImpPagado="500.00" ImpSaldoInsoluto="300.00"`}}, []string{ReglaPagoMayorCero, ReglaPagoSaldoInsoluto}},
		{"decimales del monto", pago, conPago(`FormaDePagoP="03" MonedaP="MXN" TipoCambioP="1" Monto="1660.000" NumOperacion="OP-0001"`), []string{ReglaPagoDecimales}},
		{"decimales del saldo", pago, []cambio{{dr2, `ImpSaldoAnt="800.00" ImpPagado="500.00" ImpSaldoInsoluto="300.000"`}}, []string{ReglaPagoDecimales}},
		{"cuenta ordenante con transferencia", pago, conPago(pagoP + ` CtaOrdenante="1234567890"`), nil},
		{"cuenta ordenante con efectivo", pago, conPago(strings.Replace(pagoP, `"03"`, `"01"`, 1) + ` CtaOrdenante="1234567890"`), []string{ReglaPagoCuentas}},
		{"RFC ordenante con efectivo", pago, conPago(strings.Replace(pagoP, `"03"`, `"01"`, 1) + ` RfcEmisorCtaOrd="EKU9003173C9"`), []string{ReglaPagoCuentas}},
		{"cuenta ordenante con 06", pago, conPago(strings.Replace(pagoP, `"03"`, `"06"`, 1) + ` CtaOrdenante="1234567890"`), nil},
		{"cuenta beneficiaria con 06", pago, conPago(strings.Replace(pagoP, `"03"`, `"06"`, 1) + ` CtaBeneficiario="1234567890"`), []string{ReglaPagoCuentas}},
		{"RFC beneficiario con 06", pago, conPago(strings.Replace(pagoP, `"03"`, `"06"`, 1) + ` RfcEmisorCtaBen="EKU9003173C9"`), []string{ReglaPagoCuentas}},
		{"RFC extranjero sin banco", pago, conPago(pagoP + ` RfcEmisorCtaOrd="XEXX010101000"`), []string{ReglaPagoCuentas}},
		{"RFC extranjero con banco", pago, conPago(pagoP + ` RfcEmisorCtaOrd="XEXX010101000" NomBancoOrdExt="BANCO DE PRUEBA"`), nil},
		{"cadena de pago con transferencia", pago, conPago(pagoP + cadena), nil},
		{"cadena de pago con cheque", pago, conPago(strings.Replace(pagoP, `"03"`, `"02"`, 1) + cadena), []string{ReglaPagoCuentas}},
		{"TipoCadPago sin sello", pago, conPago(pagoP + strings.Replace(cadena, ` SelloPago="U0VMTE8gREUgUFJVRUJB"`, ``, 1)), []string{ReglaPagoCadena}},
		{"sello sin TipoCadPago", pago, conPago(pagoP + ` SelloPago="U0VMTE8gREUgUFJVRUJB"`), []string{ReglaPagoCadena}},
		// El CFDI de pago 4.0.
		{"pago con UsoCFDI P01 en 4.0", pago, []cambio{{`UsoCFDI="CP01"`, `UsoCFDI="P01"`}}, []string{ReglaPagoComprobante}},
		{"pago con Exportacion 02", pago, []cambio{{`Exportacion="01"`, `Exportacion="02"`}}, []string{ReglaPagoComprobante, ReglaCCEExportacion}},
		{"pago con NoIdentificacion", pago, []cambio{{`ClaveProdServ="84111506"`, `ClaveProdServ="84111506" NoIdentificacion="P-1"`}}, []string{ReglaPagoComprobante}},
		{"pago 4.0 con ACuentaTerceros", pago, []cambio{{`ObjetoImp="01"/>`, `ObjetoImp="01"><cfdi:ACuentaTerceros RfcACuentaTerceros="XOJI740919U48" NombreACuentaTerceros="TERCERO DE PRUEBA" RegimenFiscalACuentaTerceros="612" DomicilioFiscalACuentaTerceros="76028"/></cfdi:Concepto>`}}, nil},
		{"pago 4.0 con CuentaPredial", pago, []cambio{{`ObjetoImp="01"/>`, `ObjetoImp="01"><cfdi:CuentaPredial Numero="CP-1"/></cfdi:Concepto>`}}, []string{ReglaPagoComprobante}},
		// Pagos 1.0 (CFDI 3.3).
		{"1.0 TipoCambioDR con la misma moneda", pago10, []cambio{{pago10DR1, `MonedaDR="USD" TipoCambioDR="1" MetodoDePagoDR="PPD" NumParcialidad="1"`}}, []string{ReglaPagoTipoCambioDR}},
		{"1.0 sin TipoCambioDR con otra moneda", pago10, []cambio{{pago10DR2, strings.Replace(pago10DR2, ` TipoCambioDR="20.1"`, ``, 1)}}, []string{ReglaPagoTipoCambioDR}},
		{"1.0 PPD sin parcialidad", pago10, []cambio{{pago10DR1, `MonedaDR="USD" MetodoDePagoDR="PPD"`}}, []string{ReglaPago10Parcialidades}},
		{"1.0 PUE sin parcialidad", pago10, []cambio{{pago10DR1, `MonedaDR="USD" MetodoDePagoDR="PUE"`}}, nil},
		{"1.0 sin ImpPagado con TipoCambioDR", pago10, []cambio{{pago10DR2, strings.Replace(pago10DR2, ` ImpPagado="7.57"`, ``, 1)}}, []string{ReglaPago10Parcialidades}},
		{"1.0 con Impuestos en el pago", pago10, []cambio{{pago10DR2, pago10DR2 + `<pago10:Impuestos TotalImpuestosTrasladados="0.00"/>`}}, []string{ReglaPago10Impuestos}},
		{"1.0 monto menor que lo pagado", pago10, []cambio{{pago10P1, strings.Replace(pago10P1, `"150.00"`, `"149.99"`, 1)}}, []string{ReglaPagoMonto}},
		{"1.0 saldo insoluto", pago10, []cambio{{`ImpSaldoInsoluto="7.57"`, `ImpSaldoInsoluto="7.58"`}}, []string{ReglaPagoSaldoInsoluto}},
		{"1.0 TipoCambioP 1 en MXN", pago10, []cambio{{pago10P2, `FormaDePagoP="03" MonedaP="MXN" TipoCambioP="1" Monto="152.16"`}}, nil},
		{"1.0 TipoCambioP 1.5 en MXN", pago10, []cambio{{pago10P2, `FormaDePagoP="03" MonedaP="MXN" TipoCambioP="1.5" Monto="152.16"`}}, []string{ReglaPagoTipoCambio}},
		{"1.0 USD sin TipoCambioP", pago10, []cambio{{pago10P1, `FormaDePagoP="03" MonedaP="USD" Monto="150.00"`}}, []string{ReglaPagoTipoCambio}},
		{"1.0 forma de pago 99", pago10, []cambio{{pago10P2, `FormaDePagoP="99" MonedaP="MXN" Monto="152.16"`}}, []string{ReglaPagoFormaPago}},
		// TipoCambioDR 0.0497 (USD por MXN): 7.57 / 0.0497 = 152.31…; el
		// límite inferior es 7.565 / 0.0497499999 → 152.06.
		{"1.0 en el límite inferior", pago10, []cambio{{pago10DR2, strings.Replace(pago10DR2, `"20.1"`, `"0.0497"`, 1)}, {pago10P2, `FormaDePagoP="03" MonedaP="MXN" Monto="152.06"`}}, nil},
		{"1.0 bajo el límite inferior", pago10, []cambio{{pago10DR2, strings.Replace(pago10DR2, `"20.1"`, `"0.0497"`, 1)}, {pago10P2, `FormaDePagoP="03" MonedaP="MXN" Monto="152.05"`}}, []string{ReglaPagoMonto}},
		// El CFDI de pago 3.3.
		{"3.3 pago con UsoCFDI G03", pago10, []cambio{{`UsoCFDI="P01"`, `UsoCFDI="G03"`}}, []string{ReglaPagoComprobante}},
		{"3.3 pago con Unidad", pago10, []cambio{{`ClaveUnidad="ACT"`, `ClaveUnidad="ACT" Unidad="Actividad"`}}, []string{ReglaPagoComprobante}},
		{"3.3 pago en MXN", pago10, []cambio{{`Moneda="XXX"`, `Moneda="MXN"`}}, []string{ReglaPagoComprobante}},
		{"3.3 pago con impuestos en el concepto", pago10, []cambio{{`Importe="0"/>`, `Importe="0"><cfdi:Impuestos><cfdi:Traslados><cfdi:Traslado Base="1" Impuesto="002" TipoFactor="Exento"/></cfdi:Traslados></cfdi:Impuestos></cfdi:Concepto>`}},
			[]string{ReglaPagoComprobante}},
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
			if h.Nivel != reglaDe[h.Codigo].Nivel || h.Regla == "" || h.Mensaje == "" || !strings.HasPrefix(h.Ruta, "/Comprobante") {
				t.Errorf("%s: hallazgo incompleto %+v", c.nombre, h)
			}
		}
	}
}

// Totales de Pagos 2.0: «el redondeo de la suma» de los productos
// (CRP20211), no la suma de los productos redondeados. Dos pagos de 0.01
// USD a 17.255: 0.17255 + 0.17255 = 0.3451 → 0.35. Redondeando cada uno
// daría 0.34: no es la letra, pero cae en el margen por pago [0.34, 0.36]
// (cada producto truncado o hacia arriba): advertencia. 0.33 y 0.37, error.
func TestValidarReglasTotalesPagosRedondeo(t *testing.T) {
	for declarado, nivel := range map[string]string{"0.35": "", "0.34": NivelAdvertencia, "0.36": NivelAdvertencia, "0.33": NivelError, "0.37": NivelError} {
		pg := Pago{FormaDePagoP: "03", MonedaP: "USD", TipoCambioP: ptrDec("17.255"), Monto: ptrDec("0.01")}
		c := &Comprobante{Version: "4.0", Pagos: &Pagos{Version: "2.0", Totales: &TotalesPagos{MontoTotalPagos: ptrDec(declarado)}, Pago: []Pago{pg, pg}}}
		hs := ValidarReglas(c)
		h := buscarHallazgo(hs, ReglaPagosTotales)
		if (nivel == "" && h != nil) || (nivel != "" && (h == nil || h.Nivel != nivel || len(hs) != 1)) {
			t.Errorf("MontoTotalPagos %s, quiero %q: %+v", declarado, nivel, hs)
		}
	}
}

// Impuestos de Pagos 2.0 (CRP20246-20279): ImpuestosDR de cada documento e
// ImpuestosP de cada pago, con los casos que pasan y los que no. pago20.xml
// paga una factura con IVA 16 % (Base 1000.00, Importe 160.00) y otra
// exenta (Base 500.00), las dos en MXN.
func TestValidarReglasPagosImpuestos(t *testing.T) {
	pago := string(leerFixture(t, "testdata/sinteticos/pago20.xml"))
	const (
		dr1      = `ImpSaldoInsoluto="0.00" ObjetoImpDR="02"`
		dr2      = `ImpSaldoInsoluto="300.00" ObjetoImpDR="02"`
		tDR1     = `<pago20:TrasladoDR BaseDR="1000.00" ImpuestoDR="002" TipoFactorDR="Tasa" TasaOCuotaDR="0.160000" ImporteDR="160.00"/>`
		tDR2     = `<pago20:TrasladoDR BaseDR="500.00" ImpuestoDR="002" TipoFactorDR="Exento"/>`
		tP1      = `<pago20:TrasladoP BaseP="1000.00" ImpuestoP="002" TipoFactorP="Tasa" TasaOCuotaP="0.160000" ImporteP="160.00"/>`
		tP2      = `<pago20:TrasladoP BaseP="500.00" ImpuestoP="002" TipoFactorP="Exento"/>`
		totIVA16 = `TotalTrasladosBaseIVA16="1000.00" TotalTrasladosImpuestoIVA16="160.00"`
		impDR2   = `<pago20:ImpuestosDR>
            <pago20:TrasladosDR>
              ` + tDR2 + `
            </pago20:TrasladosDR>
          </pago20:ImpuestosDR>`
		impuestosP = `<pago20:ImpuestosP>
          <pago20:TrasladosP>
            ` + tP1 + `
            ` + tP2 + `
          </pago20:TrasladosP>
        </pago20:ImpuestosP>`
	)
	// iva16 cambia el traslado de la primera factura, el del pago y los
	// totales a la vez.
	iva16 := func(baseDR, importeDR, baseP, importeP string) []cambio {
		return []cambio{
			{tDR1, `<pago20:TrasladoDR BaseDR="` + baseDR + `" ImpuestoDR="002" TipoFactorDR="Tasa" TasaOCuotaDR="0.160000" ImporteDR="` + importeDR + `"/>`},
			{tP1, `<pago20:TrasladoP BaseP="` + baseP + `" ImpuestoP="002" TipoFactorP="Tasa" TasaOCuotaP="0.160000" ImporteP="` + importeP + `"/>`},
			{totIVA16, `TotalTrasladosBaseIVA16="` + baseP + `" TotalTrasladosImpuestoIVA16="` + decimal.MustParse(importeP).StringFixed(2) + `"`},
		}
	}
	retencion := `<pago20:RetencionesDR><pago20:RetencionDR BaseDR="1000.00" ImpuestoDR="001" TipoFactorDR="Tasa" TasaOCuotaDR="0.100000" ImporteDR="100.00"/></pago20:RetencionesDR>
              <pago20:TrasladosDR>
              ` + tDR1
	// isr agrega a la primera factura una retención de ISR a 1.25 % con la
	// base y el importe dados, su RetencionP y su total.
	isr := func(base, importe string) []cambio {
		return []cambio{
			{"<pago20:TrasladosDR>\n              " + tDR1, `<pago20:RetencionesDR><pago20:RetencionDR BaseDR="` + base + `" ImpuestoDR="001" TipoFactorDR="Tasa" TasaOCuotaDR="0.012500" ImporteDR="` + importe + `"/></pago20:RetencionesDR>
              <pago20:TrasladosDR>
              ` + tDR1},
			{"<pago20:TrasladosP>\n            " + tP1, `<pago20:RetencionesP><pago20:RetencionP ImpuestoP="001" ImporteP="` + importe + `"/></pago20:RetencionesP><pago20:TrasladosP>` + "\n            " + tP1},
			{`<pago20:Totales `, `<pago20:Totales TotalRetencionesISR="` + importe + `" `},
		}
	}
	casos := []struct {
		nombre string
		cs     []cambio
		quiero []string
	}{
		// ObjetoImpDR (CRP20246, CRP20247, CRP20278, CRP20279).
		{"objeto 01 con impuestos", []cambio{{dr1, `ImpSaldoInsoluto="0.00" ObjetoImpDR="01"`}}, []string{ReglaPagoObjetoImpDR}},
		{"objeto 04 con impuestos", []cambio{{dr1, `ImpSaldoInsoluto="0.00" ObjetoImpDR="04"`}}, []string{ReglaPagoObjetoImpDR}},
		{"objeto 02 sin impuestos", []cambio{{impDR2, ``}}, []string{ReglaPagoObjetoImpDR, ReglaPagoImpuestosP}},
		{"objeto 01 sin impuestos", []cambio{{impDR2, ``}, {dr2, `ImpSaldoInsoluto="300.00" ObjetoImpDR="01"`}, {"\n            " + tP2, ``},
			{` TotalTrasladosBaseIVAExento="500.00"`, ``}}, nil},
		{"objeto 06 con IVA", []cambio{{dr2, `ImpSaldoInsoluto="300.00" ObjetoImpDR="06"`}}, []string{ReglaPagoObjetoImpDR}},
		{"objeto 08 con retención de ISR", []cambio{{dr1, `ImpSaldoInsoluto="0.00" ObjetoImpDR="08"`}, {"<pago20:TrasladosDR>\n              " + tDR1, retencion},
			{"<pago20:TrasladosP>\n            " + tP1, `<pago20:RetencionesP><pago20:RetencionP ImpuestoP="001" ImporteP="100.00"/></pago20:RetencionesP><pago20:TrasladosP>` + "\n            " + tP1},
			{`<pago20:Totales `, `<pago20:Totales TotalRetencionesISR="100.00" `}}, []string{ReglaPagoObjetoImpDR}}, // el IVA de la primera
		{"objeto 07 con IVA", []cambio{{dr2, `ImpSaldoInsoluto="300.00" ObjetoImpDR="07"`}}, []string{ReglaPagoObjetoImpDR}},
		// ImpuestosDR (CRP20248, CRP20249, CRP20252, CRP20255, CRP20258,
		// CRP20259).
		{"ImpuestosDR vacío", []cambio{{impDR2, `<pago20:ImpuestosDR/>`}}, []string{ReglaPagoImpuestoDR, ReglaPagoImpuestosP}},
		{"exento con importe", []cambio{{tDR2, `<pago20:TrasladoDR BaseDR="500.00" ImpuestoDR="002" TipoFactorDR="Exento" ImporteDR="0.00"/>`}}, []string{ReglaPagoImpuestoDR}},
		{"tasa sin importe", []cambio{{tDR1, `<pago20:TrasladoDR BaseDR="1000.00" ImpuestoDR="002" TipoFactorDR="Tasa" TasaOCuotaDR="0.160000"/>`}}, []string{ReglaPagoImpuestoDR}},
		{"base cero", []cambio{{tDR2, `<pago20:TrasladoDR BaseDR="0.00" ImpuestoDR="002" TipoFactorDR="Exento"/>`}}, []string{ReglaPagoImpuestoDR, ReglaPagoImpuestosPSuma}},
		{"retención exenta", []cambio{{"<pago20:TrasladosDR>\n              " + tDR1, strings.Replace(retencion, `TipoFactorDR="Tasa" TasaOCuotaDR="0.100000"`, `TipoFactorDR="Exento" TasaOCuotaDR="0.100000"`, 1)}},
			[]string{ReglaPagoImpuestoDR, ReglaPagoImpuestosP}}, // y falta su RetencionP
		// ImporteDR (CRP20254, CRP20261): Base 1000.00 × 0.16 admite
		// [159.99, 160.01] a 2 decimales; con 6, [159.999200, 160.000800].
		// El segundo nivel amplía los de 2 decimales en ±0.10 (20 conceptos
		// × ½ centavo): [159.89, 160.11].
		{"ImporteDR en el límite", iva16("1000.00", "160.01", "1000.00", "160.01"), nil},
		{"ImporteDR sobre el límite", iva16("1000.00", "160.02", "1000.00", "160.02"), []string{adv(ReglaPagoImporteDR)}},
		{"ImporteDR en el borde superior del margen", iva16("1000.00", "160.11", "1000.00", "160.11"), []string{adv(ReglaPagoImporteDR)}},
		{"ImporteDR fuera del margen por arriba", iva16("1000.00", "160.12", "1000.00", "160.12"), []string{ReglaPagoImporteDR}},
		{"ImporteDR en el borde inferior del margen", iva16("1000.00", "159.89", "1000.00", "159.89"), []string{adv(ReglaPagoImporteDR)}},
		{"ImporteDR fuera del margen por abajo", iva16("1000.00", "159.88", "1000.00", "159.88"), []string{ReglaPagoImporteDR}},
		// ISR de RESICO (1.25 %) copiado de una factura de varios conceptos:
		// 3790.00 × 0.0125 = 47.375, la letra admite [47.37, 47.38].
		{"ISR 1.25 % de la factura", isr("3790.00", "47.39"), []string{adv(ReglaPagoImporteDR)}},
		{"ISR 1.25 % en la letra", isr("3790.00", "47.38"), nil},
		{"ISR 1.25 % en el borde del margen", isr("3790.00", "47.48"), []string{adv(ReglaPagoImporteDR)}},
		{"ISR 1.25 % fuera del margen", isr("3790.00", "47.49"), []string{ReglaPagoImporteDR}},
		{"ImporteDR de 6 decimales en el margen", iva16("1000.00", "160.010000", "1000.00", "160.01"), []string{adv(ReglaPagoImporteDR)}},
		{"ImporteDR de 6 decimales en su límite", iva16("1000.00", "160.000800", "1000.00", "160.000800"), nil},
		// ImpuestosP (CRP20263, CRP20266, CRP20271).
		{"TrasladoP repetido", []cambio{{tP1, tP1 + tP1}}, []string{ReglaPagoImpuestosP, ReglaPagosTotales}},
		{"TrasladoP exento con importe", []cambio{{tP2, `<pago20:TrasladoP BaseP="500.00" ImpuestoP="002" TipoFactorP="Exento" ImporteP="0.00"/>`}}, []string{ReglaPagoImpuestosP}},
		{"TrasladoP sin documento", []cambio{{tP2, tP2 + `<pago20:TrasladoP BaseP="10.00" ImpuestoP="003" TipoFactorP="Tasa" TasaOCuotaP="0.080000" ImporteP="0.80"/>`}}, []string{ReglaPagoImpuestosP}},
		{"sin ImpuestosP", []cambio{{impuestosP, ``}}, []string{ReglaPagoImpuestosP, ReglaPagosTotales}},
		{"tasa distinta en el pago", []cambio{{`TasaOCuotaP="0.160000"`, `TasaOCuotaP="0.080000"`}}, []string{ReglaPagoImpuestosP, ReglaPagosTotales}},
		// Sumas (CRP20265, CRP20268, CRP20269, CRP20274). Base 1000.03 ×
		// 0.16 = 160.0048, escrito 160.00: un ImporteP de 160.01 no es la
		// suma pero sí el margen por documento; 160.02, error.
		{"ImporteP distinto de la suma", iva16("1000.00", "160.00", "1000.00", "160.01"), []string{ReglaPagoImpuestosPSuma}},
		{"BaseP distinta de la suma", iva16("1000.00", "160.00", "1000.01", "160.00"), []string{ReglaPagoImpuestosPSuma}},
		{"ImporteP con Base × Tasa sin redondear", iva16("1000.03", "160.00", "1000.03", "160.01"), []string{adv(ReglaPagoImpuestosPSuma)}},
		{"ImporteP fuera del margen", iva16("1000.03", "160.00", "1000.03", "160.02"), []string{ReglaPagoImpuestosPSuma}},
		{"ImporteP con 6 decimales", iva16("1000.03", "160.00", "1000.03", "160.000000"), nil},
		// En otra moneda la suma es BaseDR / EquivalenciaDR: 247.78 /
		// 0.055333 = 4477.978…, redondeada 4477.98 o truncada 4477.97.
		{"BaseP de otra moneda truncada", pagoOtraMoneda("4477.97"), nil},
		{"BaseP de otra moneda redondeada", pagoOtraMoneda("4477.98"), nil},
		{"BaseP de otra moneda fuera", pagoOtraMoneda("4477.99"), []string{ReglaPagoImpuestosPSuma}},
		{"BaseP de otra moneda sin convertir", pagoOtraMoneda("247.78"), []string{ReglaPagoImpuestosPSuma}},
	}
	for _, c := range casos {
		hs := reglasDe(t, aplicarCambios(t, pago, c.cs))
		got := codigosDe(hs)
		quiero := append([]string(nil), c.quiero...)
		sort.Strings(quiero)
		if strings.Join(got, ",") != strings.Join(quiero, ",") {
			t.Errorf("%s: %v, quiero %v\n%+v", c.nombre, got, quiero, hs)
		}
		for _, h := range hs {
			r := reglaDe[h.Codigo]
			if (h.Nivel != r.Nivel && !(r.Margen != "" && h.Nivel == NivelAdvertencia)) || h.Regla == "" || h.Mensaje == "" || !strings.HasPrefix(h.Ruta, "/Comprobante") {
				t.Errorf("%s: hallazgo incompleto %+v", c.nombre, h)
			}
		}
	}
	// Sin EquivalenciaDR en otra moneda no se puede convertir: solo el
	// hallazgo de la equivalencia, no el de la suma.
	hs := reglasDe(t, aplicarCambios(t, pago, []cambio{{`MonedaDR="MXN" EquivalenciaDR="1" NumParcialidad="2"`, `MonedaDR="USD" NumParcialidad="2"`}}))
	if got := codigosDe(hs); strings.Join(got, ",") != ReglaPagoEquivalencia {
		t.Errorf("sin equivalencia: %v", got)
	}
}

// pagoOtraMoneda pasa la factura exenta de pago20.xml a USD (EquivalenciaDR
// 0.055333, BaseDR 247.78) con la BaseP dada.
func pagoOtraMoneda(baseP string) []cambio {
	return []cambio{
		{`MonedaDR="MXN" EquivalenciaDR="1" NumParcialidad="2" ImpSaldoAnt="800.00" ImpPagado="500.00" ImpSaldoInsoluto="300.00"`,
			`MonedaDR="USD" EquivalenciaDR="0.055333" NumParcialidad="2" ImpSaldoAnt="247.78" ImpPagado="247.78" ImpSaldoInsoluto="0.00"`},
		{`Monto="1660.00"`, `Monto="5637.97"`}, {`MontoTotalPagos="1660.00"`, `MontoTotalPagos="5637.97"`},
		{`BaseDR="500.00"`, `BaseDR="247.78"`}, {`BaseP="500.00"`, `BaseP="` + baseP + `"`},
		{`TotalTrasladosBaseIVAExento="500.00"`, `TotalTrasladosBaseIVAExento="` + baseP + `"`},
	}
}

// ObjetoImp contra el nodo Impuestos del concepto (CFDI40171, CFDI40172,
// CFDI140226, CFDI140227), y FormaPago en nómina según la versión.
func TestValidarReglasObjetoImp(t *testing.T) {
	tr := func(impuesto string) Impuesto { return Impuesto{Impuesto: impuesto, TipoFactor: "Exento"} }
	ret := func(impuesto string) Impuesto { return Impuesto{Impuesto: impuesto, TipoFactor: "Tasa"} }
	casos := []struct {
		obj         string
		traslados   []Impuesto
		retenciones []Impuesto
		sinNodo     bool
		mal         bool
	}{
		{"01", nil, nil, true, false},
		{"01", []Impuesto{tr("002")}, nil, false, true},
		{"02", nil, nil, true, true},
		{"02", []Impuesto{tr("002")}, nil, false, false},
		{"03", nil, nil, true, false},
		{"03", []Impuesto{tr("002")}, nil, false, true},
		{"04", []Impuesto{tr("002")}, nil, false, true},
		{"05", []Impuesto{tr("002")}, nil, false, true},
		{"06", nil, []Impuesto{ret("001")}, false, false},
		{"06", []Impuesto{tr("002")}, nil, false, true},
		{"08", nil, []Impuesto{ret("003")}, false, true},
		{"08", nil, nil, true, false},
		{"07", []Impuesto{tr("003")}, []Impuesto{ret("001")}, false, false},
		{"07", nil, []Impuesto{ret("001")}, false, true},
		{"07", []Impuesto{tr("003")}, []Impuesto{ret("002")}, false, true},
		{"07", nil, []Impuesto{ret("003")}, false, true}, // el IEPS tiene que ser traslado
	}
	for _, c := range casos {
		cp := Concepto{ObjetoImp: c.obj}
		if !c.sinNodo {
			cp.Impuestos = &ImpuestosConcepto{Traslados: c.traslados, Retenciones: c.retenciones}
		}
		hs := ValidarReglas(&Comprobante{Version: "4.0", Conceptos: []Concepto{cp}})
		if TieneHallazgo(hs, ReglaObjetoImp) != c.mal {
			t.Errorf("ObjetoImp %s %+v: %+v", c.obj, cp.Impuestos, hs)
		}
		// En 3.3 no hay ObjetoImp.
		if hs := ValidarReglas(&Comprobante{Version: "3.3", Conceptos: []Concepto{cp}}); TieneHallazgo(hs, ReglaObjetoImp) {
			t.Errorf("3.3 ObjetoImp %s: %+v", c.obj, hs)
		}
	}
	for v, mal := range map[string]bool{"4.0": true, "3.3": false} {
		hs := ValidarReglas(&Comprobante{Version: v, TipoDeComprobante: "N", FormaPago: "99"})
		if TieneHallazgo(hs, ReglaTipoCompFormaPago) != mal {
			t.Errorf("nómina %s con FormaPago: %+v", v, hs)
		}
	}
	// Base de un traslado del comprobante con decimales de más: error en 4.0.
	hs := ValidarReglas(&Comprobante{Version: "4.0", Moneda: "MXN", Impuestos: &Impuestos{Traslados: []Impuesto{{Base: ptrDec("1.000"), Impuesto: "002", TipoFactor: "Exento"}}}})
	if !TieneHallazgo(hs, ReglaDecimalesImpuesto) {
		t.Errorf("Base con 3 decimales: %+v", hs)
	}
}
