package cfdi

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/OrtaMarco/cfdi-go/decimal"
)

func leer(t *testing.T, nombre string) *Comprobante {
	t.Helper()
	b, err := os.ReadFile(filepath.Join("testdata", "sinteticos", nombre))
	if err != nil {
		t.Fatal(err)
	}
	c, err := ParseBytes(b)
	if err != nil {
		t.Fatalf("%s: %v", nombre, err)
	}
	return c
}

// monto exige que d exista, valga want y conserve el texto exacto.
func monto(t *testing.T, campo string, d *decimal.Decimal, want string) {
	t.Helper()
	if d == nil {
		t.Errorf("%s: nil, quiero %s", campo, want)
		return
	}
	if d.Cmp(decimal.MustParse(want)) != 0 || d.Original() != want {
		t.Errorf("%s = %q, quiero %q", campo, d.Original(), want)
	}
}

func igual[T comparable](t *testing.T, campo string, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("%s = %v, quiero %v", campo, got, want)
	}
}

// codigos cuenta las advertencias por código.
func codigos(c *Comprobante) map[string]int {
	m := map[string]int{}
	for _, a := range c.Advertencias {
		m[a.Codigo]++
	}
	return m
}

func sinAdvertencias(t *testing.T, c *Comprobante) {
	t.Helper()
	if len(c.Advertencias) > 0 {
		t.Errorf("advertencias inesperadas: %+v", c.Advertencias)
	}
}

func sumar(ds ...*decimal.Decimal) decimal.Decimal {
	var s decimal.Decimal
	for _, d := range ds {
		if d != nil {
			s = s.Add(*d)
		}
	}
	return s
}

func TestIngreso40PUE(t *testing.T) {
	c := leer(t, "ingreso40-pue.xml")
	sinAdvertencias(t, c)
	igual(t, "Version", c.Version, "4.0")
	igual(t, "TipoDeComprobante", c.TipoDeComprobante, "I")
	igual(t, "MetodoPago", c.MetodoPago, "PUE")
	igual(t, "Exportacion", c.Exportacion, "01")
	monto(t, "SubTotal", c.SubTotal, "1500.00")
	monto(t, "Total", c.Total, "1640.00")
	if c.Descuento != nil {
		t.Errorf("Descuento = %v, quiero nil", c.Descuento)
	}
	igual(t, "Fecha.Original", c.Fecha.Original, "2026-01-15T10:30:00")
	igual(t, "Fecha.Time", c.Fecha.Time, time.Date(2026, 1, 15, 10, 30, 0, 0, time.UTC))
	igual(t, "Emisor.Rfc", c.Emisor.Rfc, "EKU9003173C9")
	igual(t, "Receptor.UsoCFDI", c.Receptor.UsoCFDI, "S01")

	if len(c.Conceptos) != 2 {
		t.Fatalf("conceptos = %d, quiero 2", len(c.Conceptos))
	}
	c1, c2 := c.Conceptos[0], c.Conceptos[1]
	monto(t, "Concepto[1].Importe", c1.Importe, "1000.00")
	monto(t, "Concepto[2].Cantidad", c2.Cantidad, "2.5")
	if len(c1.Impuestos.Traslados) != 1 || len(c1.Impuestos.Retenciones) != 1 {
		t.Fatalf("concepto 1: %d traslados, %d retenciones", len(c1.Impuestos.Traslados), len(c1.Impuestos.Retenciones))
	}
	monto(t, "Concepto[1] IVA", c1.Impuestos.Traslados[0].Importe, "160.00")
	monto(t, "Concepto[1] TasaOCuota", c1.Impuestos.Traslados[0].TasaOCuota, "0.160000")
	monto(t, "Concepto[1] ret ISR", c1.Impuestos.Retenciones[0].Importe, "100.00")
	igual(t, "ret Impuesto", c1.Impuestos.Retenciones[0].Impuesto, "001")
	igual(t, "Pedimento", c2.InformacionAduanera[0].NumeroPedimento, "26  47  3807  6000001")

	monto(t, "TotalImpuestosTrasladados", c.Impuestos.TotalImpuestosTrasladados, "240.00")
	monto(t, "TotalImpuestosRetenidos", c.Impuestos.TotalImpuestosRetenidos, "100.00")
	// Los traslados de los conceptos suman lo que dice el total.
	tras := sumar(c1.Impuestos.Traslados[0].Importe, c2.Impuestos.Traslados[0].Importe)
	if tras.Cmp(*c.Impuestos.TotalImpuestosTrasladados) != 0 {
		t.Errorf("suma de traslados = %s", tras)
	}
	// SubTotal + traslados - retenciones = Total.
	calc := c.SubTotal.Add(*c.Impuestos.TotalImpuestosTrasladados).Sub(*c.Impuestos.TotalImpuestosRetenidos)
	if calc.Cmp(*c.Total) != 0 {
		t.Errorf("total calculado = %s", calc)
	}

	igual(t, "UUID", c.UUID(), "ABCDEF01-2345-4678-9ABC-DEF012345678")
	igual(t, "UUIDOriginal", c.TimbreFiscalDigital.UUIDOriginal, "abcdef01-2345-4678-9abc-def012345678")
	igual(t, "TFD.Version", c.TimbreFiscalDigital.Version, "1.1")
	if !slices.Equal(c.Complementos(), []string{"TimbreFiscalDigital 1.1"}) {
		t.Errorf("Complementos = %v", c.Complementos())
	}
}

func TestIngreso33(t *testing.T) {
	c := leer(t, "ingreso33.xml")
	sinAdvertencias(t, c)
	igual(t, "Version", c.Version, "3.3")
	igual(t, "MetodoPago", c.MetodoPago, "PPD")
	igual(t, "Moneda", c.Moneda, "USD")
	monto(t, "TipoCambio", c.TipoCambio, "19.8765")
	monto(t, "Descuento", c.Descuento, "300.00")
	monto(t, "Total", c.Total, "3132.00")
	// El RFC venía en minúsculas.
	igual(t, "Emisor.Rfc", c.Emisor.Rfc, "EKU9003173C9")
	igual(t, "ResidenciaFiscal", c.Receptor.ResidenciaFiscal, "USA")
	igual(t, "conceptos", len(c.Conceptos), 1)
	igual(t, "CuentaPredial", c.Conceptos[0].CuentaPredial[0].Numero, "ABC123")
	monto(t, "Concepto Descuento", c.Conceptos[0].Descuento, "300.00")
	monto(t, "IVA", c.Conceptos[0].Impuestos.Traslados[0].Importe, "432.00")
	igual(t, "UUID", c.UUID(), "11111111-2222-3333-4444-555555555555")
	igual(t, "UUIDOriginal", c.TimbreFiscalDigital.UUIDOriginal, "")
}

func TestEgresoRelacionados(t *testing.T) {
	c := leer(t, "egreso40-relacionados.xml")
	sinAdvertencias(t, c)
	igual(t, "TipoDeComprobante", c.TipoDeComprobante, "E")
	if len(c.CfdiRelacionados) != 2 {
		t.Fatalf("CfdiRelacionados = %d grupos, quiero 2", len(c.CfdiRelacionados))
	}
	g1, g2 := c.CfdiRelacionados[0], c.CfdiRelacionados[1]
	igual(t, "TipoRelacion 1", g1.TipoRelacion, "01")
	igual(t, "TipoRelacion 2", g2.TipoRelacion, "07")
	igual(t, "uuids grupo 1", len(g1.CfdiRelacionado), 2)
	igual(t, "uuid 1.1", g1.CfdiRelacionado[0].UUID, "ABCDEF01-2345-4678-9ABC-DEF012345678")
	igual(t, "uuid 1.1 original", g1.CfdiRelacionado[0].UUIDOriginal, "")
	igual(t, "uuid 1.2", g1.CfdiRelacionado[1].UUID, "22222222-3333-4444-5555-66666666666A")
	igual(t, "uuid 1.2 original", g1.CfdiRelacionado[1].UUIDOriginal, " 22222222-3333-4444-5555-66666666666a ")
	igual(t, "uuid 2.1", g2.CfdiRelacionado[0].UUID, "33333333-4444-5555-6666-777777777777")
	monto(t, "Total", c.Total, "116.00")
}

func TestPago20(t *testing.T) {
	c := leer(t, "pago20.xml")
	sinAdvertencias(t, c)
	igual(t, "TipoDeComprobante", c.TipoDeComprobante, "P")
	monto(t, "Total", c.Total, "0")
	p := c.Pagos
	if p == nil {
		t.Fatal("sin Pagos")
	}
	igual(t, "Pagos.Version", p.Version, "2.0")
	monto(t, "MontoTotalPagos", p.Totales.MontoTotalPagos, "1660.00")
	monto(t, "TotalTrasladosImpuestoIVA16", p.Totales.TotalTrasladosImpuestoIVA16, "160.00")
	monto(t, "TotalTrasladosBaseIVAExento", p.Totales.TotalTrasladosBaseIVAExento, "500.00")
	if len(p.Pago) != 1 {
		t.Fatalf("pagos = %d", len(p.Pago))
	}
	pg := p.Pago[0]
	monto(t, "Monto", pg.Monto, "1660.00")
	igual(t, "FechaPago", pg.FechaPago.Time, time.Date(2026, 3, 9, 0, 0, 0, 0, time.UTC))
	if len(pg.DoctoRelacionado) != 2 {
		t.Fatalf("doctos = %d, quiero 2", len(pg.DoctoRelacionado))
	}
	d1, d2 := pg.DoctoRelacionado[0], pg.DoctoRelacionado[1]
	igual(t, "IdDocumento 1", d1.IdDocumento, "ABCDEF01-2345-4678-9ABC-DEF012345678")
	igual(t, "IdDocumento 2", d2.IdDocumento, "55555555-6666-7777-8888-99999999999F")
	igual(t, "IdDocumentoOriginal 2", d2.IdDocumentoOriginal, "55555555-6666-7777-8888-99999999999f")
	igual(t, "NumParcialidad 2", d2.NumParcialidad, "2")
	monto(t, "EquivalenciaDR", d1.EquivalenciaDR, "1")
	monto(t, "ImpSaldoInsoluto 2", d2.ImpSaldoInsoluto, "300.00")
	if s := sumar(d1.ImpPagado, d2.ImpPagado); s.Cmp(*pg.Monto) != 0 {
		t.Errorf("suma ImpPagado = %s, Monto = %s", s, pg.Monto)
	}
	monto(t, "ImporteDR 1", d1.ImpuestosDR.TrasladosDR[0].ImporteDR, "160.00")
	ex := d2.ImpuestosDR.TrasladosDR[0]
	igual(t, "TipoFactorDR 2", ex.TipoFactorDR, "Exento")
	if ex.ImporteDR != nil || ex.TasaOCuotaDR != nil {
		t.Errorf("TrasladoDR exento con importe/tasa: %+v", ex)
	}
	monto(t, "BaseDR exento", ex.BaseDR, "500.00")
	if len(pg.ImpuestosP.TrasladosP) != 2 {
		t.Fatalf("TrasladosP = %d", len(pg.ImpuestosP.TrasladosP))
	}
	monto(t, "ImporteP", pg.ImpuestosP.TrasladosP[0].ImporteP, "160.00")
	if pg.Impuestos != nil {
		t.Errorf("Impuestos (1.0) en un pago 2.0: %+v", pg.Impuestos)
	}
	igual(t, "UUID", c.UUID(), "66666666-7777-8888-9999-AAAAAAAAAAAA")
	if !slices.Contains(c.Complementos(), "Pagos 2.0") {
		t.Errorf("Complementos = %v", c.Complementos())
	}
}

func TestPago10(t *testing.T) {
	c := leer(t, "pago10.xml")
	sinAdvertencias(t, c)
	p := c.Pagos
	if p == nil {
		t.Fatal("pagos 1.0 no se leyó")
	}
	igual(t, "Pagos.Version", p.Version, "1.0")
	if p.Totales != nil {
		t.Errorf("Totales en pagos 1.0")
	}
	if len(p.Pago) != 2 {
		t.Fatalf("pagos = %d, quiero 2", len(p.Pago))
	}
	p1, p2 := p.Pago[0], p.Pago[1]
	monto(t, "Monto 1", p1.Monto, "150.00")
	monto(t, "TipoCambioP 1", p1.TipoCambioP, "20.1234")
	igual(t, "MonedaP 1", p1.MonedaP, "USD")
	igual(t, "RfcEmisorCtaOrd", p1.RfcEmisorCtaOrd, "XEXX010101000")
	igual(t, "SelloPago", p1.SelloPago, "U0VMTE8gREUgUFJVRUJB")
	d1, d2 := p1.DoctoRelacionado[0], p2.DoctoRelacionado[0]
	igual(t, "IdDocumento", d1.IdDocumento, "11111111-2222-3333-4444-555555555555")
	igual(t, "MetodoDePagoDR", d1.MetodoDePagoDR, "PPD")
	monto(t, "TipoCambioDR 2", d2.TipoCambioDR, "20.1")
	if d2.EquivalenciaDR != nil {
		t.Errorf("EquivalenciaDR en pagos 1.0")
	}
	if s := sumar(d1.ImpPagado, d2.ImpPagado); s.Cmp(*d1.ImpSaldoAnt) != 0 {
		t.Errorf("parcialidades suman %s, saldo inicial %s", s, d1.ImpSaldoAnt)
	}
	monto(t, "ImpSaldoInsoluto final", d2.ImpSaldoInsoluto, "0.00")
}

func TestNomina12(t *testing.T) {
	c := leer(t, "nomina12.xml")
	sinAdvertencias(t, c)
	igual(t, "TipoDeComprobante", c.TipoDeComprobante, "N")
	if len(c.Nomina) != 1 {
		t.Fatalf("Nomina = %d", len(c.Nomina))
	}
	n := c.Nomina[0]
	igual(t, "Version", n.Version, "1.2")
	igual(t, "TipoNomina", n.TipoNomina, "O")
	igual(t, "FechaPago", n.FechaPago.Time, time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC))
	monto(t, "NumDiasPagados", n.NumDiasPagados, "15.000")
	// Escalas distintas: 15.000 días = 15; 5600.00 + 10.00 + 0.5 ≠ 5610.
	if n.NumDiasPagados.Cmp(decimal.MustParse("15")) != 0 || n.NumDiasPagados.Cmp(decimal.MustParse("14.9999")) != 1 {
		t.Errorf("NumDiasPagados no compara por valor: %s", n.NumDiasPagados)
	}
	if s := n.TotalPercepciones.Add(decimal.MustParse("0.5")); s.String() != "5600.5" {
		t.Errorf("5600.00 + 0.5 = %s", s)
	}
	monto(t, "TotalPercepciones", n.TotalPercepciones, "5600.00")
	monto(t, "TotalDeducciones", n.TotalDeducciones, "575.50")
	monto(t, "TotalOtrosPagos", n.TotalOtrosPagos, "10.00")
	igual(t, "RegistroPatronal", n.Emisor.RegistroPatronal, "Z0000001234")
	igual(t, "Antiguedad", n.Receptor.Antiguedad, "P315W")
	igual(t, "ClaveEntFed", n.Receptor.ClaveEntFed, "QUE")
	monto(t, "SalarioDiarioIntegrado", n.Receptor.SalarioDiarioIntegrado, "380.50")

	pc := n.Percepciones
	igual(t, "percepciones", len(pc.Percepcion), 2)
	monto(t, "TotalSueldos", pc.TotalSueldos, "5600.00")
	var grav, exen decimal.Decimal
	for _, p := range pc.Percepcion {
		grav = grav.Add(*p.ImporteGravado)
		exen = exen.Add(*p.ImporteExento)
	}
	if grav.Cmp(*pc.TotalGravado) != 0 || exen.Cmp(*pc.TotalExento) != 0 {
		t.Errorf("gravado %s / exento %s no cuadran con los totales", grav, exen)
	}
	igual(t, "HorasExtra.Dias", pc.Percepcion[1].HorasExtra[0].Dias, "2")
	monto(t, "HorasExtra.ImportePagado", pc.Percepcion[1].HorasExtra[0].ImportePagado, "600.00")

	igual(t, "deducciones", len(n.Deducciones.Deduccion), 2)
	monto(t, "TotalImpuestosRetenidos", n.Deducciones.TotalImpuestosRetenidos, "450.00")
	monto(t, "ISR", n.Deducciones.Deduccion[0].Importe, "450.00")
	igual(t, "otros pagos", len(n.OtrosPagos), 1)
	monto(t, "SubsidioCausado", n.OtrosPagos[0].SubsidioAlEmpleo.SubsidioCausado, "10.00")
	igual(t, "incapacidades", len(n.Incapacidades), 1)
	igual(t, "DiasIncapacidad", n.Incapacidades[0].DiasIncapacidad, "1")

	// Descuento del comprobante = deducciones; SubTotal = percepciones + otros pagos.
	if s := n.TotalPercepciones.Add(*n.TotalOtrosPagos); s.Cmp(*c.SubTotal) != 0 {
		t.Errorf("SubTotal %s ≠ percepciones + otros pagos %s", c.SubTotal, s)
	}
	if !slices.Equal(c.Complementos(), []string{"TimbreFiscalDigital 1.1", "Nomina 1.2"}) {
		t.Errorf("Complementos = %v", c.Complementos())
	}
}

func TestTraslado(t *testing.T) {
	c := leer(t, "traslado40.xml")
	sinAdvertencias(t, c)
	igual(t, "TipoDeComprobante", c.TipoDeComprobante, "T")
	igual(t, "partes", len(c.Conceptos[0].Parte), 1)
	monto(t, "Parte.Cantidad", c.Conceptos[0].Parte[0].Cantidad, "10")
	if !c.Total.IsZero() {
		t.Errorf("Total = %s", c.Total)
	}
}

func TestCodificaciones(t *testing.T) {
	const desc = "Café “orgánico” — año pingüino €"
	casos := []struct {
		archivo    string
		advertidas []string
	}{
		{"windows1252.xml", nil},
		{"utf8-miente.xml", []string{AdvEncodingCorregido}},
		{"latin1-miente.xml", []string{AdvEncodingCorregido}},
		{"bom.xml", nil},
	}
	for _, cs := range casos {
		c := leer(t, cs.archivo)
		igual(t, cs.archivo+" Descripcion", c.Conceptos[0].Descripcion, desc)
		igual(t, cs.archivo+" Serie", c.Serie, "Ñ")
		igual(t, cs.archivo+" Receptor.Nombre", c.Receptor.Nombre, "PÚBLICO EN GENERAL")
		var got []string
		for _, a := range c.Advertencias {
			got = append(got, a.Codigo)
		}
		if !slices.Equal(got, cs.advertidas) {
			t.Errorf("%s: advertencias %v, quiero %v", cs.archivo, got, cs.advertidas)
		}
	}
}

func TestComplementoDesconocidoYPrefijos(t *testing.T) {
	c := leer(t, "complemento-desconocido.xml")
	// Prefijos x: y t: en lugar de cfdi: y tfd:.
	igual(t, "Emisor.Rfc", c.Emisor.Rfc, "EKU9003173C9")
	igual(t, "UUID", c.UUID(), "AAAAAAAA-BBBB-CCCC-DDDD-EEEEEEEEEEEE")
	want := []string{"{urn:ejemplo:concepto}DeConcepto", "{urn:ejemplo:comprobante}DeComprobante"}
	if !slices.Equal(c.ComplementosNoSoportados, want) {
		t.Errorf("ComplementosNoSoportados = %v", c.ComplementosNoSoportados)
	}
	igual(t, "ComplementoConcepto", c.Conceptos[0].ComplementoConcepto[0], want[0])
	if m := codigos(c); m[AdvComplementoNoSoportado] != 2 || len(m) != 1 {
		t.Errorf("advertencias = %v", m)
	}
	// Exento bien formado: sin tasa ni importe y sin advertencia.
	tr := c.Conceptos[0].Impuestos.Traslados[0]
	igual(t, "TipoFactor", tr.TipoFactor, "Exento")
	if tr.Importe != nil || tr.TasaOCuota != nil {
		t.Errorf("Exento con importe: %+v", tr)
	}
	monto(t, "Base exento", tr.Base, "2000.00")
}

func TestSinTimbre(t *testing.T) {
	c := leer(t, "sin-tfd.xml")
	if m := codigos(c); m[AdvSinTimbre] != 1 || len(m) != 1 {
		t.Errorf("advertencias = %v", m)
	}
	igual(t, "UUID", c.UUID(), "")
	// Namespace por omisión, sin prefijo.
	igual(t, "Anio", c.InformacionGlobal.Anio, "2026")
	igual(t, "conceptos", len(c.Conceptos), 1)
	monto(t, "Total", c.Total, "11.60")
}

func TestTFD10(t *testing.T) {
	c := leer(t, "tfd10-minusculas.xml")
	sinAdvertencias(t, c)
	tfd := c.TimbreFiscalDigital
	igual(t, "Version", tfd.Version, "1.0")
	igual(t, "SelloCFD", tfd.SelloCFD, "U0VMTE8gREUgUFJVRUJB")
	igual(t, "NoCertificadoSAT", tfd.NoCertificadoSAT, "00001000000300000000")
	igual(t, "UUID", tfd.UUID, "CCCCCCCC-DDDD-EEEE-FFFF-000000000000")
}

func TestAnomalias(t *testing.T) {
	c := leer(t, "anomalias.xml")
	want := map[string]int{
		AdvContenidoTrasRaiz:   1,
		AdvFechaFormato:        1,
		AdvDecimalInvalido:     2,
		AdvNodoDuplicado:       2,
		AdvNamespaceInesperado: 1,
		AdvExentoConImporte:    1,
		AdvAtributoFaltante:    1,
		AdvFechaInvalida:       1,
		AdvUUIDInvalido:        1,
		AdvComplementoVersion:  1,
		AdvNodoFaltante:        1,
	}
	got := codigos(c)
	for k, v := range want {
		if got[k] != v {
			t.Errorf("%s: %d advertencias, quiero %d", k, got[k], v)
		}
	}
	for k := range got {
		if _, ok := want[k]; !ok {
			t.Errorf("advertencia inesperada %s", k)
		}
	}
	rutas := map[string]bool{}
	for _, a := range c.Advertencias {
		rutas[a.Codigo+" "+a.Ruta] = true
	}
	for _, r := range []string{
		"DECIMAL_INVALIDO /Comprobante/@SubTotal",
		"DECIMAL_INVALIDO /Comprobante/Conceptos/Concepto[1]/@Cantidad",
		"EXENTO_CON_IMPORTE /Comprobante/Conceptos/Concepto[1]/Impuestos/Traslados/Traslado[1]",
		"ATRIBUTO_FALTANTE /Comprobante/Conceptos/Concepto[1]/Impuestos/Traslados/Traslado[2]/@Importe",
		"NAMESPACE_INESPERADO /Comprobante/Receptor",
	} {
		if !rutas[r] {
			t.Errorf("falta la advertencia %q", r)
		}
	}

	// Lo ilegible queda vacío; lo demás se lee.
	if c.SubTotal != nil {
		t.Errorf("SubTotal ilegible debía quedar nil: %v", c.SubTotal)
	}
	monto(t, "Total", c.Total, "1160.00")
	igual(t, "Fecha civil", c.Fecha.Time, time.Date(2026, 7, 1, 10, 0, 0, 0, time.UTC))
	igual(t, "Emisor (el primero)", c.Emisor.Nombre, "ESCUELA KEMPER URGATE")
	igual(t, "Emisor.Rfc normalizado", c.Emisor.Rfc, "EKU9003173C9")
	igual(t, "Receptor de otro namespace", c.Receptor.Rfc, "XAXX010101000")
	// El Exento con importe conserva lo que venía.
	monto(t, "Exento.Importe", c.Conceptos[0].Impuestos.Traslados[0].Importe, "0.00")
	igual(t, "UUID ilegible normalizado", c.UUID(), "NO-ES-UN-UUID")
	if c.TimbreFiscalDigital.FechaTimbrado.Valida() {
		t.Errorf("FechaTimbrado 'ayer' no debía ser válida")
	}
	igual(t, "FechaTimbrado.Original", c.TimbreFiscalDigital.FechaTimbrado.Original, "ayer")
}

func TestFatales(t *testing.T) {
	casos := map[string]string{
		"no-es-xml.xml":        ClaseNoEsXML,
		"raiz-retenciones.xml": ClaseNoEsComprobante,
		"sin-namespace.xml":    ClaseNoEsComprobante,
		"truncado.xml":         ClaseXMLMalFormado,
	}
	for archivo, clase := range casos {
		b, err := os.ReadFile(filepath.Join("testdata", "sinteticos", "fatales", archivo))
		if err != nil {
			t.Fatal(err)
		}
		c, err := ParseBytes(b)
		if c != nil || err == nil {
			t.Errorf("%s: se leyó, quería error %s", archivo, clase)
			continue
		}
		igual(t, archivo, ClaseDe(err), clase)
	}
	for _, in := range []string{"", "   \n", "<", "{\"json\":true}"} {
		if _, err := ParseBytes([]byte(in)); ClaseDe(err) != ClaseNoEsXML {
			t.Errorf("ParseBytes(%q): %v, quiero %s", in, err, ClaseNoEsXML)
		}
	}
}

type lectorRoto struct{}

func (lectorRoto) Read([]byte) (int, error) { return 0, errors.New("disco roto") }

func TestParseReader(t *testing.T) {
	f, err := os.Open(filepath.Join("testdata", "sinteticos", "ingreso33.xml"))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	c, err := Parse(f)
	if err != nil {
		t.Fatal(err)
	}
	monto(t, "Total", c.Total, "3132.00")

	_, err = Parse(lectorRoto{})
	igual(t, "clase", ClaseDe(err), ClaseLectura)
	if !strings.Contains(err.Error(), "disco roto") {
		t.Errorf("el error no envuelve la causa: %v", err)
	}
	igual(t, "ClaseDe(nil)", ClaseDe(errors.New("x")), "")
}

// minimo es un CFDI 4.0 completo y válido con un hueco para variar.
func minimo(descripcion string) string {
	return `<cfdi:Comprobante xmlns:cfdi="http://www.sat.gob.mx/cfd/4" Version="4.0" Fecha="2026-01-01T00:00:00" SubTotal="1.00" Moneda="MXN" Total="1.00" TipoDeComprobante="I" Exportacion="01" LugarExpedicion="42501">` +
		`<cfdi:Emisor Rfc="EKU9003173C9" Nombre="E" RegimenFiscal="601"/>` +
		`<cfdi:Receptor Rfc="XAXX010101000" Nombre="R" DomicilioFiscalReceptor="42501" RegimenFiscalReceptor="616" UsoCFDI="S01"/>` +
		`<cfdi:Conceptos><cfdi:Concepto ClaveProdServ="01010101" Cantidad="1" ClaveUnidad="ACT" Descripcion="` + descripcion + `" ValorUnitario="1.00" Importe="1.00" ObjetoImp="01"/></cfdi:Conceptos>` +
		`<cfdi:Complemento><tfd:TimbreFiscalDigital xmlns:tfd="http://www.sat.gob.mx/TimbreFiscalDigital" Version="1.1" UUID="11111111-2222-3333-4444-555555555555" FechaTimbrado="2026-01-01T00:00:01"/></cfdi:Complemento>` +
		`</cfdi:Comprobante>`
}

func TestMinimoSinAdvertencias(t *testing.T) {
	c, err := ParseBytes([]byte(minimo("x")))
	if err != nil {
		t.Fatal(err)
	}
	sinAdvertencias(t, c)
}

func TestTolerancias(t *testing.T) {
	casos := []struct {
		nombre string
		doc    []byte
		codigo string
		desc   string
	}{
		{"caracter de control", []byte(minimo("a\x01b\x0bc")), AdvCaracterInvalido, "abc"},
		{"ampersand suelto", []byte(minimo("A & B")), AdvXMLNoEstricto, "A & B"},
		{"entidad HTML", []byte(minimo("A&nbsp;B")), AdvXMLNoEstricto, "A&nbsp;B"},
		{"encoding desconocido", []byte(`<?xml version="1.0" encoding="x-raro"?>` + minimo("ñ")), AdvEncodingDesconocido, "ñ"},
		{"sin declaración y latin1", []byte(strings.Replace(minimo("caf~"), "~", "\xe9", 1)), AdvEncodingCorregido, "café"},
		{"utf-16le con BOM", utf16le(`<?xml version="1.0" encoding="UTF-16"?>` + minimo("año")), "", "año"},
		{"basura nula al final", []byte(minimo("x") + "\x00\x00\n"), "", "x"},
	}
	for _, cs := range casos {
		t.Run(cs.nombre, func(t *testing.T) {
			c, err := ParseBytes(cs.doc)
			if err != nil {
				t.Fatal(err)
			}
			got := codigos(c)
			if cs.codigo == "" {
				sinAdvertencias(t, c)
			} else if got[cs.codigo] != 1 || len(got) != 1 {
				t.Errorf("advertencias %v, quiero solo %s", got, cs.codigo)
			}
			igual(t, "Descripcion", c.Conceptos[0].Descripcion, cs.desc)
		})
	}
}

func utf16le(s string) []byte {
	out := []byte{0xFF, 0xFE}
	for _, r := range s {
		if r > 0xFFFF {
			panic("fuera del BMP")
		}
		out = append(out, byte(r), byte(r>>8))
	}
	return out
}

func TestParseFecha(t *testing.T) {
	casos := []struct {
		in        string
		soloFecha bool
		valida    bool
		raro      bool
		want      time.Time
	}{
		{"2026-01-15T10:30:00", false, true, false, time.Date(2026, 1, 15, 10, 30, 0, 0, time.UTC)},
		{"2026-01-15", true, true, false, time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC)},
		{"2026-01-15", false, true, true, time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC)},
		{"2026-01-15T10:30:00", true, true, true, time.Date(2026, 1, 15, 10, 30, 0, 0, time.UTC)},
		{"2026-01-15T10:30:00.123", false, true, true, time.Date(2026, 1, 15, 10, 30, 0, 123e6, time.UTC)},
		// Con zona: se conserva la hora de pared, no se convierte.
		{"2026-01-15T10:30:00-06:00", false, true, true, time.Date(2026, 1, 15, 10, 30, 0, 0, time.UTC)},
		{"2026-01-15T10:30:00Z", false, true, true, time.Date(2026, 1, 15, 10, 30, 0, 0, time.UTC)},
		{" 2026-01-15T10:30:00 ", false, true, true, time.Date(2026, 1, 15, 10, 30, 0, 0, time.UTC)},
		{"15/01/2026", false, false, false, time.Time{}},
		{"2026-02-30T00:00:00", false, false, false, time.Time{}},
		{"", false, false, false, time.Time{}},
	}
	for _, c := range casos {
		f, raro, err := parseFecha(c.in, c.soloFecha)
		if (err == nil) != c.valida || raro != c.raro || !f.Time.Equal(c.want) || f.Original != c.in || f.Valida() != c.valida {
			t.Errorf("parseFecha(%q, %v) = %v, raro=%v, err=%v", c.in, c.soloFecha, f.Time, raro, err)
		}
	}
	mx, err := time.LoadLocation("America/Mexico_City")
	if err != nil {
		t.Skip("sin tzdata:", err)
	}
	f, _, _ := parseFecha("2026-01-15T10:30:00", false)
	if got := f.En(mx); got.Hour() != 10 || got.Location() != mx || !got.Equal(time.Date(2026, 1, 15, 16, 30, 0, 0, time.UTC)) {
		t.Errorf("En(México) = %v", got)
	}
}

func TestJSONSobre(t *testing.T) {
	c := leer(t, "pago10.xml")
	b, err := c.JSON()
	if err != nil {
		t.Fatal(err)
	}
	var sobre map[string]json.RawMessage
	if err := json.Unmarshal(b, &sobre); err != nil {
		t.Fatal(err)
	}
	for _, k := range []string{"schema_version", "cfdi", "advertencias"} {
		if _, ok := sobre[k]; !ok {
			t.Errorf("falta %q en el sobre", k)
		}
	}
	igual(t, "schema_version", string(sobre["schema_version"]), `"1"`)
	igual(t, "advertencias", string(sobre["advertencias"]), `[]`)
	if _, ok := sobre["origen"]; ok {
		t.Errorf("origen vacío debía omitirse")
	}
	// Sin sellos por omisión…
	for _, s := range []string{`"sello"`, `"certificado"`, `"sello_cfd"`, `"sello_sat"`, `"sello_pago"`, `"cert_pago"`, `"cad_pago"`} {
		if bytes.Contains(b, []byte(s+":")) {
			t.Errorf("%s salió sin IncluirSello", s)
		}
	}
	// …sin tocar el Comprobante…
	igual(t, "Sello conservado", c.Sello, "U0VMTE8gREUgUFJVRUJB")
	igual(t, "SelloPago conservado", c.Pagos.Pago[0].SelloPago, "U0VMTE8gREUgUFJVRUJB")
	// …y con ellos si se piden.
	b2, _ := c.JSONOpciones(OpcionesJSON{IncluirSello: true})
	for _, s := range []string{`"sello"`, `"certificado"`, `"sello_cfd"`, `"sello_sat"`, `"sello_pago"`, `"cert_pago"`, `"cad_pago"`} {
		if !bytes.Contains(b2, []byte(s+":")) {
			t.Errorf("%s no salió con IncluirSello", s)
		}
	}
	// Montos como string con el texto original.
	if !bytes.Contains(b, []byte(`"imp_saldo_insoluto":"0.00"`)) || !bytes.Contains(b, []byte(`"tipo_cambio_dr":"20.1"`)) {
		t.Errorf("montos no salieron como string original: %s", b)
	}
}

func TestJSONIdaYVuelta(t *testing.T) {
	for _, p := range fixtures(t) {
		if esFatal(p) || esRetenciones(t, p) {
			continue
		}
		b, _ := os.ReadFile(p)
		c, err := ParseBytes(b)
		if err != nil {
			t.Fatal(err)
		}
		uno, _ := c.JSONOpciones(OpcionesJSON{IncluirSello: true})
		var d Documento
		if err := json.Unmarshal(uno, &d); err != nil {
			t.Fatalf("%s: %v", p, err)
		}
		d.CFDI.Advertencias = d.Advertencias
		dos, _ := d.CFDI.JSONOpciones(OpcionesJSON{IncluirSello: true})
		if !bytes.Equal(uno, dos) {
			t.Errorf("%s: el JSON no sobrevive ida y vuelta", p)
		}
	}
}

// CFDI 3.2: se lee lo que coincide con 3.3 (atributos en minúsculas), se
// advierte VERSION_ANTIGUA y no se exige lo que introdujo 3.3.
func TestVersion32(t *testing.T) {
	const doc = `<cfdi:Comprobante xmlns:cfdi="http://www.sat.gob.mx/cfd/3" version="3.2" fecha="2015-02-18T11:49:20" subTotal="100.00" total="116.00" tipoDeComprobante="ingreso" formaDePago="PAGO EN UNA SOLA EXHIBICION" metodoDePago="01" LugarExpedicion="Monterrey">` +
		`<cfdi:Emisor rfc="EKU9003173C9"><cfdi:RegimenFiscal Regimen="General"/></cfdi:Emisor>` +
		`<cfdi:Receptor rfc="XAXX010101000"/>` +
		`<cfdi:Conceptos><cfdi:Concepto cantidad="1" unidad="PZA" descripcion="x" valorUnitario="100.00" importe="100.00"/></cfdi:Conceptos>` +
		`<cfdi:Complemento><tfd:TimbreFiscalDigital xmlns:tfd="http://www.sat.gob.mx/TimbreFiscalDigital" version="1.0" UUID="11111111-2222-3333-4444-555555555555" FechaTimbrado="2015-02-18T11:50:00"/></cfdi:Complemento>` +
		`</cfdi:Comprobante>`
	c, err := ParseBytes([]byte(doc))
	if err != nil {
		t.Fatal(err)
	}
	if got := codigos(c); got[AdvVersionAntigua] != 1 || len(got) != 1 {
		t.Errorf("advertencias %+v, quiero solo %s", c.Advertencias, AdvVersionAntigua)
	}
	igual(t, "Version", c.Version, "3.2")
	monto(t, "Total", c.Total, "116.00")
	monto(t, "Importe", c.Conceptos[0].Importe, "100.00")
	igual(t, "Emisor.Rfc", c.Emisor.Rfc, "EKU9003173C9")
	igual(t, "UUID", c.UUID(), "11111111-2222-3333-4444-555555555555")
	// Lo exclusivo de 3.2 no se modela.
	igual(t, "FormaPago", c.FormaPago, "")
	igual(t, "MetodoPago", c.MetodoPago, "")

	// Una versión que no es 3.0/3.2/3.3/4.0 sigue siendo VERSION_NO_SOPORTADA,
	// con las exigencias de 3.3.
	c, _ = ParseBytes([]byte(strings.Replace(doc, `version="3.2"`, `version="3.1"`, 1)))
	if got := codigos(c); got[AdvVersionNoSoportada] != 1 || got[AdvVersionAntigua] != 0 || got[AdvAtributoFaltante] == 0 {
		t.Errorf("3.1: advertencias %v", got)
	}
}

// La Addenda se recorre sin guardarla: no cambia la lectura, pero si está mal
// formada el documento sigue siendo fatal.
func TestAddendaSeSalta(t *testing.T) {
	addenda := `<cfdi:Addenda><x:Pedido xmlns:x="http://example.org/x"><cfdi:Emisor Rfc="OTRO"/><x:L n="1"/></x:Pedido></cfdi:Addenda>`
	sin, _ := ParseBytes([]byte(minimo("x")))
	con, err := ParseBytes([]byte(strings.Replace(minimo("x"), `</cfdi:Comprobante>`, addenda+`</cfdi:Comprobante>`, 1)))
	if err != nil {
		t.Fatal(err)
	}
	a, _ := sin.JSON()
	b, _ := con.JSON()
	if !bytes.Equal(a, b) {
		t.Errorf("la Addenda cambió la lectura:\n%s\n%s", a, b)
	}
	roto := strings.Replace(minimo("x"), `</cfdi:Comprobante>`, `<cfdi:Addenda><x:A n="1`, 1) // truncado
	if _, err := ParseBytes([]byte(roto)); ClaseDe(err) != ClaseXMLMalFormado {
		t.Errorf("Addenda mal formada: %v", err)
	}
}

// encoding/json escapa & como &; un RFC con & debe salir tal cual.
func TestJSONSinEscaparHTML(t *testing.T) {
	c, err := ParseBytes([]byte(strings.Replace(minimo("x"), `Rfc="XAXX010101000"`, `Rfc="E&amp;Ñ831019M53"`, 1)))
	if err != nil {
		t.Fatal(err)
	}
	b, _ := c.JSON()
	if !bytes.Contains(b, []byte(`"rfc":"E&Ñ831019M53"`)) || bytes.HasSuffix(b, []byte("\n")) {
		t.Errorf("JSON = %s", b)
	}
}
