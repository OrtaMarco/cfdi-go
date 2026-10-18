package cfdi

import (
	"fmt"
	"math/big"
	"sort"

	"github.com/OrtaMarco/cfdi-go/decimal"
)

// Reglas del CFDI de pago y de los complementos Pagos 1.0 (en CFDI 3.3) y
// 2.0 (en 4.0). Cada regla cita el código de las dos matrices cuando
// existe en ambas.
var reglasPagos = []Regla{
	regla(ReglaPagoComprobante, NivelError,
		"tipo P: Moneda XXX, UsoCFDI CP01 (4.0) o P01 (3.3), Exportacion 01 (4.0) y un solo concepto 84111506, Cantidad 1, ClaveUnidad ACT, Descripcion Pago, ValorUnitario e Importe 0, ObjetoImp 01 (4.0), sin NoIdentificacion ni Unidad ni nodos hijo (en 4.0 solo ACuentaTerceros)",
		citas(cod(DocMatrizPagos20, "CRP20102", "CRP20104", "CRP20111", "CRP20112", "CRP20113", "CRP20114", "CRP20115", "CRP20116", "CRP20117", "CRP20118", "CRP20119", "CRP20120", "CRP20122"),
			pag(DocGuiaPagos20, 9, 17, 18),
			cod(DocMatrizPagos10, "CRP103", "CRP110", "CRP111", "CRP112", "CRP113", "CRP114", "CRP115", "CRP116", "CRP117", "CRP118", "CRP119", "CRP120"),
			pag(DocGuiaPagos10, 8, 13, 14, 15)), ""),
	regla(ReglaPagoSaldoInsoluto, NivelError, "ImpSaldoInsoluto = ImpSaldoAnt − ImpPagado",
		citas(cod(DocMatrizPagos20, "CRP20244"), pag(DocGuiaPagos20, 42), cod(DocMatrizPagos10, "CRP226"), pag(DocGuiaPagos10, 28)), ""),
	regla(ReglaPagoEquivalencia, NivelError, "Pagos 2.0: EquivalenciaDR 1 si MonedaDR es MonedaP; requerida si son distintas",
		citas(cod(DocMatrizPagos20, "CRP20237", "CRP20238"), pag(DocGuiaPagos20, 40)), ""),
	regla(ReglaPagoTipoCambioDR, NivelError, "Pagos 1.0: TipoCambioDR requerido si MonedaDR es distinta de MonedaP; ausente si es la misma",
		citas(cod(DocMatrizPagos10, "CRP218", "CRP219"), pag(DocGuiaPagos10, 26)),
		"CRP220 (MonedaDR MXN, TipoCambioDR 1) no se aplica: la guía vigente lo quitó (control de cambios, p. 79)"),
	regla(ReglaPagoMonto, NivelError,
		"Monto mayor o igual que la suma de ImpPagado / EquivalenciaDR (TipoCambioDR en 1.0) truncada a los decimales de MonedaP; con documentos en otra moneda basta con que sea mayor o igual que la suma de los límites inferiores (ImpPagado − ½·10^-dI) / (EquivalenciaDR + ½·10^-dE − 10^-10), cada uno truncado (con la misma moneda, EquivalenciaDR 1 con 10 decimales)",
		citas(cod(DocMatrizPagos20, "CRP20217", "CRP20275", "CRP20276", "CRP20277"), pag(DocGuiaPagos20, 31, 32, 33, 34), cod(DocMatrizPagos10, "CRP206"), pag(DocGuiaPagos10, 20)),
		"la fórmula de los límites es una imagen en la guía 2.0: se toma el texto de la guía 1.0 p. 20 con EquivalenciaDR por TipoCambioDR"),
	regla(ReglaPagoTipoCambio, NivelError, "MonedaP distinta de MXN con TipoCambioP; MonedaP MXN con TipoCambioP 1 (en 1.0 también puede omitirse)",
		citas(cod(DocMatrizPagos20, "CRP20214", "CRP20215"), pag(DocGuiaPagos20, 30), cod(DocMatrizPagos10, "CRP203", "CRP204"), pag(DocGuiaPagos10, 18, 43)),
		"CRP204 pide omitirlo y la pregunta 7 de la guía 1.0 admite el valor 1: se aceptan ambos"),
	conMargen(regla(ReglaPagosTotales, NivelError, "Pagos 2.0: cada atributo de Totales igual al redondeo a 2 decimales de la suma de los ImpuestosP (y Monto) por TipoCambioP; con ImpuestosP de IVA o retenciones, el atributo de Totales que les corresponde",
		citas(cod(DocMatrizPagos20, "CRP20201", "CRP20202", "CRP20203", "CRP20204", "CRP20205", "CRP20206", "CRP20207", "CRP20208", "CRP20209", "CRP20210", "CRP20211", "CRP20264", "CRP20267", "CRP20273"),
			pag(DocEstPagos20, 52, 53, 54, 55, 56, 57), pag(DocGuiaPagos20, 21, 27)), "se redondea la suma, no cada producto; el SAT no da tolerancia"),
		"dos niveles: si no es el redondeo de la suma pero cae entre la suma de cada producto truncado y la de cada uno redondeado hacia arriba a 2 decimales, advertencia: el margen de redondeo por pago (redondear cada producto antes de sumar), que la letra no admite"),
	regla(ReglaPagoMonedaXXX, NivelError, "MonedaP y MonedaDR distintas de XXX",
		citas(cod(DocMatrizPagos20, "CRP20213", "CRP20236"), cod(DocMatrizPagos10, "CRP202", "CRP217"), pag(DocGuiaPagos10, 18, 25)), ""),
	regla(ReglaPagoMayorCero, NivelError, "Monto, ImpSaldoAnt e ImpPagado mayores que cero",
		citas(cod(DocMatrizPagos20, "CRP20218", "CRP20239", "CRP20241"), cod(DocMatrizPagos10, "CRP207", "CRP221", "CRP223"), pag(DocGuiaPagos10, 20, 27, 28)), ""),
	regla(ReglaPagoDecimales, NivelError, "Monto con no más decimales de los que admite MonedaP; ImpSaldoAnt, ImpPagado e ImpSaldoInsoluto, los de MonedaDR",
		citas(cod(DocMatrizPagos20, "CRP20219", "CRP20240", "CRP20242", "CRP20243"), cod(DocMatrizPagos10, "CRP208", "CRP222", "CRP224", "CRP225"), pag(DocGuiaPagos10, 18, 26)), ""),
	regla(ReglaPagoFormaPago, NivelError, "FormaDePagoP distinta de 99",
		citas(cod(DocMatrizPagos20, "CRP20212"), cod(DocMatrizPagos10, "CRP201"), pag(DocGuiaPagos10, 17)), ""),
	regla(ReglaPagoCuentas, NivelError,
		"RfcEmisorCtaOrd y CtaOrdenante solo con FormaDePagoP 02, 03, 04, 05, 06, 28 o 29; RfcEmisorCtaBen y CtaBeneficiario solo con 02, 03, 04, 05, 28 o 29; TipoCadPago solo con 03; NomBancoOrdExt requerido con RfcEmisorCtaOrd XEXX010101000",
		citas(cod(DocMatrizPagos20, "CRP20221", "CRP20224", "CRP20227", "CRP20228", "CRP20229", "CRP20223"),
			cod(DocMatrizPagos10, "CRP238", "CRP212", "CRP214", "CRP215", "CRP216", "CRP211"), pag(DocGuiaPagos10, 22, 23)),
		"los patrones de cuenta de c_FormaPago (CRP213, CRP239, CRP20225, CRP20226) no se revisan: el catálogo no está en la librería"),
	regla(ReglaPagoCadena, NivelError, "CertPago, CadPago y SelloPago van los tres si hay TipoCadPago y ninguno si no",
		citas(cod(DocMatrizPagos20, "CRP20230", "CRP20231", "CRP20232", "CRP20233", "CRP20234", "CRP20235"),
			cod(DocMatrizPagos10, "CRP227", "CRP228", "CRP229", "CRP230", "CRP231", "CRP232"), pag(DocGuiaPagos10, 23, 24)), ""),
	regla(ReglaPago10Parcialidades, NivelError, "Pagos 1.0: con MetodoDePagoDR PPD, NumParcialidad, ImpSaldoAnt e ImpSaldoInsoluto; ImpPagado requerido con más de un documento o con TipoCambioDR",
		citas(cod(DocMatrizPagos10, "CRP233", "CRP234", "CRP236", "CRP235"), pag(DocGuiaPagos10, 27, 28)), ""),
	regla(ReglaPago10Impuestos, NivelError, "Pagos 1.0: el Pago no lleva nodo Impuestos",
		citas(cod(DocMatrizPagos10, "CRP237"), pag(DocGuiaPagos10, 29)), ""),
}

// comprobantePago revisa lo fijo de un CFDI de pago (TipoDeComprobante P).
// Lo que ya revisan las reglas por TipoDeComprobante (SubTotal y Total 0,
// sin FormaPago, MetodoPago, CondicionesDePago, Descuento ni Impuestos) y
// la de TipoCambio con XXX no se repite aquí.
func (v *validadorReglas) comprobantePago() {
	c := v.c
	if v.moneda != "" && v.moneda != "XXX" {
		v.hallazgoTexto(ReglaPagoComprobante, "/Comprobante", "Moneda", c.Moneda, "un comprobante de pago va con Moneda XXX")
	}
	uso := "P01"
	if v.v40 {
		uso = "CP01"
		if e := clave(c.Exportacion); e != "" && e != "01" {
			v.hallazgoTexto(ReglaPagoComprobante, "/Comprobante", "Exportacion", c.Exportacion, "un comprobante de pago va con Exportacion 01")
		}
	}
	if c.Receptor != nil && clave(c.Receptor.UsoCFDI) != "" && clave(c.Receptor.UsoCFDI) != uso {
		v.hallazgoTexto(ReglaPagoComprobante, "/Comprobante/Receptor", "UsoCFDI", c.Receptor.UsoCFDI, "un comprobante de pago va con UsoCFDI "+uso)
	}
	if len(c.Conceptos) != 1 {
		v.hallazgo(ReglaPagoComprobante, "/Comprobante/Conceptos", "", nil, fmt.Sprintf("un comprobante de pago lleva un solo concepto y trae %d", len(c.Conceptos)))
		return
	}
	cp := c.Conceptos[0]
	r := "/Comprobante/Conceptos/Concepto[1]"
	texto := []struct{ attr, valor, quiero string }{
		{"ClaveProdServ", cp.ClaveProdServ, "84111506"},
		{"ClaveUnidad", cp.ClaveUnidad, "ACT"},
		{"Descripcion", cp.Descripcion, "Pago"},
		{"NoIdentificacion", cp.NoIdentificacion, ""},
		{"Unidad", cp.Unidad, ""},
	}
	if v.v40 {
		texto = append(texto, struct{ attr, valor, quiero string }{"ObjetoImp", clave(cp.ObjetoImp), "01"})
	}
	for _, x := range texto {
		switch {
		case x.valor == x.quiero:
		case x.quiero == "":
			v.hallazgoTexto(ReglaPagoComprobante, r, x.attr, x.valor, "en un comprobante de pago el concepto no lleva "+x.attr)
		default:
			v.hallazgoTexto(ReglaPagoComprobante, r, x.attr, x.valor, fmt.Sprintf("en un comprobante de pago %s es %q", x.attr, x.quiero))
		}
	}
	num := []struct {
		attr  string
		valor *decimal.Decimal
		n     int64
	}{{"Cantidad", cp.Cantidad, 1}, {"ValorUnitario", cp.ValorUnitario, 0}, {"Importe", cp.Importe, 0}}
	for _, x := range num {
		if x.valor != nil && !x.valor.Equal(decimal.NewFromInt64(x.n, 0)) {
			v.hallazgo(ReglaPagoComprobante, r, x.attr, x.valor, fmt.Sprintf("en un comprobante de pago %s es %d", x.attr, x.n))
		}
	}
	// Nodos hijo del concepto: en 4.0 solo ACuentaTerceros (CRP20112); el
	// nodo Impuestos lo revisa ObjetoImp. En 3.3 ninguno (CRP112).
	var hijos []string
	if !v.v40 && cp.Impuestos != nil {
		hijos = append(hijos, "Impuestos")
	}
	if !v.v40 && cp.ACuentaTerceros != nil {
		hijos = append(hijos, "ACuentaTerceros")
	}
	if len(cp.InformacionAduanera) > 0 {
		hijos = append(hijos, "InformacionAduanera")
	}
	if len(cp.CuentaPredial) > 0 {
		hijos = append(hijos, "CuentaPredial")
	}
	if len(cp.Parte) > 0 {
		hijos = append(hijos, "Parte")
	}
	if cp.PorCuentadeTerceros != nil || cp.VentaVehiculos != nil || cp.HidroYPetro != nil || len(cp.ComplementoConcepto) > 0 {
		hijos = append(hijos, "ComplementoConcepto")
	}
	for _, h := range hijos {
		v.hallazgo(ReglaPagoComprobante, r+"/"+h, "", nil, "en un comprobante de pago el concepto no lleva el nodo "+h)
	}
}

// rat convierte un decimal en racional exacto (para dividir entre
// EquivalenciaDR sin float).
func rat(d decimal.Decimal) *big.Rat {
	r, _ := new(big.Rat).SetString(d.String())
	return r
}

// truncarRat corta r a n decimales hacia cero y lo devuelve como decimal.
func truncarRat(r *big.Rat, n int32) decimal.Decimal {
	escala := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(n)), nil)
	x := new(big.Rat).Mul(r, new(big.Rat).SetInt(escala))
	q := new(big.Int).Quo(x.Num(), x.Denom())
	d, _ := decimal.Parse(q.String())
	return d.Mul(decimal.NewFromInt64(1, n))
}

// diezALaMenos10 es lo que la guía de Pagos 1.0 (p. 20) resta al tipo de
// cambio del límite inferior.
var diezALaMenos10 = decimal.NewFromInt64(1, 10)

// limiteInferiorPagado es el límite inferior de ImpPagado en la moneda del
// pago, guía de Pagos 1.0 p. 20: (ImpPagado − ½·10^-dI) / (eq + ½·10^-dE −
// 10^-10), truncado a decP (guía de Pagos 2.0 p. 33, punto v). dE son los
// decimales de eq; con la misma moneda, eq es 1 con 10 decimales (CRP20277).
func limiteInferiorPagado(pagado, eq decimal.Decimal, dE, decP int32) decimal.Decimal {
	den := eq.Add(mitadUlp(dE)).Sub(diezALaMenos10)
	return truncarRat(new(big.Rat).Quo(rat(pagado.Sub(mitadUlp(pagado.Scale()))), rat(den)), decP)
}

var (
	formasOrdenante    = map[string]bool{"02": true, "03": true, "04": true, "05": true, "06": true, "28": true, "29": true}
	formasBeneficiario = map[string]bool{"02": true, "03": true, "04": true, "05": true, "28": true, "29": true}
)

func (v *validadorReglas) pagos() {
	p := v.c.Pagos
	v10 := p.Version == "1.0"
	r := "/Comprobante/Complemento/Pagos"
	for i, pg := range p.Pago {
		rp := idx(r, "Pago", i)
		v.pago(pg, rp, v10)
		for j, d := range pg.DoctoRelacionado {
			rd := idx(rp, "DoctoRelacionado", j)
			v.doctoRelacionado(pg, d, rd, v10, len(pg.DoctoRelacionado))
			if !v10 {
				v.impuestosDR(d, rd)
			}
		}
		v.montoPago(pg, rp, v10)
		if !v10 {
			v.impuestosP(pg, rp)
		}
	}
	if !v10 {
		v.totalesPagos()
	}
}

// pago revisa los atributos del nodo Pago.
func (v *validadorReglas) pago(pg Pago, rp string, v10 bool) {
	monedaP := clave(pg.MonedaP)
	forma := clave(pg.FormaDePagoP)
	if forma == "99" {
		v.hallazgoTexto(ReglaPagoFormaPago, rp, "FormaDePagoP", pg.FormaDePagoP, "FormaDePagoP no puede ser 99 (por definir)")
	}
	if monedaP == "XXX" {
		v.hallazgoTexto(ReglaPagoMonedaXXX, rp, "MonedaP", pg.MonedaP, "MonedaP no puede ser XXX")
	}
	switch {
	case monedaP == "MXN" && pg.TipoCambioP != nil && !pg.TipoCambioP.Equal(uno):
		v.hallazgo(ReglaPagoTipoCambio, rp, "TipoCambioP", pg.TipoCambioP, "con MonedaP MXN el TipoCambioP es 1")
	case monedaP == "MXN" && pg.TipoCambioP == nil && !v10:
		v.hallazgo(ReglaPagoTipoCambio, rp, "TipoCambioP", nil, "con MonedaP MXN se registra TipoCambioP 1")
	case monedaP != "" && monedaP != "MXN" && monedaP != "XXX" && pg.TipoCambioP == nil:
		v.hallazgo(ReglaPagoTipoCambio, rp, "TipoCambioP", nil, "con MonedaP "+monedaP+" se requiere TipoCambioP")
	}
	if pg.Monto != nil && pg.Monto.Sign() <= 0 {
		v.hallazgo(ReglaPagoMayorCero, rp, "Monto", pg.Monto, "el Monto tiene que ser mayor que cero")
	}
	if dec, ok := decimalesMoneda[monedaP]; ok && excedeDecimales(pg.Monto, dec) {
		v.hallazgo(ReglaPagoDecimales, rp, "Monto", pg.Monto, fmt.Sprintf("Monto %s tiene %d decimales y la moneda %s admite %d", pg.Monto.Original(), pg.Monto.Scale(), monedaP, dec))
	}
	if forma != "" {
		cuentas := []struct {
			attr, valor string
			permitidas  map[string]bool
		}{
			{"RfcEmisorCtaOrd", pg.RfcEmisorCtaOrd, formasOrdenante}, {"CtaOrdenante", pg.CtaOrdenante, formasOrdenante},
			{"RfcEmisorCtaBen", pg.RfcEmisorCtaBen, formasBeneficiario}, {"CtaBeneficiario", pg.CtaBeneficiario, formasBeneficiario},
			{"TipoCadPago", pg.TipoCadPago, map[string]bool{"03": true}},
		}
		for _, x := range cuentas {
			if x.valor != "" && !x.permitidas[forma] {
				v.hallazgoTexto(ReglaPagoCuentas, rp, x.attr, x.valor, "con FormaDePagoP "+forma+" no va "+x.attr)
			}
		}
	}
	if clave(pg.RfcEmisorCtaOrd) == "XEXX010101000" && pg.NomBancoOrdExt == "" {
		v.hallazgo(ReglaPagoCuentas, rp, "NomBancoOrdExt", nil, "con RfcEmisorCtaOrd XEXX010101000 se requiere NomBancoOrdExt")
	}
	conTipo := pg.TipoCadPago != ""
	for _, x := range []struct{ attr, valor string }{{"CertPago", pg.CertPago}, {"CadPago", pg.CadPago}, {"SelloPago", pg.SelloPago}} {
		switch {
		case conTipo && x.valor == "":
			v.hallazgo(ReglaPagoCadena, rp, x.attr, nil, "con TipoCadPago se requiere "+x.attr)
		case !conTipo && x.valor != "":
			v.hallazgo(ReglaPagoCadena, rp, x.attr, nil, "sin TipoCadPago no va "+x.attr)
		}
	}
	if v10 && len(pg.Impuestos) > 0 {
		v.hallazgo(ReglaPago10Impuestos, rp+"/Impuestos", "", nil, "un Pago 1.0 no lleva el nodo Impuestos")
	}
}

// doctoRelacionado revisa un DoctoRelacionado; n es cuántos trae el pago.
func (v *validadorReglas) doctoRelacionado(pg Pago, d DoctoRelacionado, rd string, v10 bool, n int) {
	monedaP, monedaDR := clave(pg.MonedaP), clave(d.MonedaDR)
	if monedaDR == "XXX" {
		v.hallazgoTexto(ReglaPagoMonedaXXX, rd, "MonedaDR", d.MonedaDR, "MonedaDR no puede ser XXX")
	}
	importes := []struct {
		attr  string
		valor *decimal.Decimal
	}{{"ImpSaldoAnt", d.ImpSaldoAnt}, {"ImpPagado", d.ImpPagado}, {"ImpSaldoInsoluto", d.ImpSaldoInsoluto}}
	dec, decOK := decimalesMoneda[monedaDR]
	for _, x := range importes {
		if x.attr != "ImpSaldoInsoluto" && x.valor != nil && x.valor.Sign() <= 0 {
			v.hallazgo(ReglaPagoMayorCero, rd, x.attr, x.valor, x.attr+" tiene que ser mayor que cero")
		}
		if decOK && excedeDecimales(x.valor, dec) {
			v.hallazgo(ReglaPagoDecimales, rd, x.attr, x.valor, fmt.Sprintf("%s %s tiene %d decimales y la moneda %s admite %d", x.attr, x.valor.Original(), x.valor.Scale(), monedaDR, dec))
		}
	}
	if d.ImpSaldoAnt != nil && d.ImpPagado != nil && d.ImpSaldoInsoluto != nil {
		if esperado := d.ImpSaldoAnt.Sub(*d.ImpPagado); !d.ImpSaldoInsoluto.Equal(esperado) {
			v.hallazgo(ReglaPagoSaldoInsoluto, rd, "ImpSaldoInsoluto", d.ImpSaldoInsoluto,
				fmt.Sprintf("ImpSaldoInsoluto %s y ImpSaldoAnt − ImpPagado es %s", d.ImpSaldoInsoluto.Original(), esperado))
		}
	}
	misma := monedaDR != "" && monedaDR == monedaP
	distinta := monedaDR != "" && monedaP != "" && monedaDR != monedaP
	if !v10 {
		eq := d.EquivalenciaDR
		switch {
		case misma && eq != nil && !eq.Equal(uno):
			v.hallazgo(ReglaPagoEquivalencia, rd, "EquivalenciaDR", eq, "MonedaDR es la MonedaP: EquivalenciaDR es 1")
		case distinta && eq == nil:
			v.hallazgo(ReglaPagoEquivalencia, rd, "EquivalenciaDR", nil, "MonedaDR "+monedaDR+" distinta de MonedaP "+monedaP+": se requiere EquivalenciaDR")
		}
		return
	}
	switch {
	case misma && d.TipoCambioDR != nil:
		v.hallazgo(ReglaPagoTipoCambioDR, rd, "TipoCambioDR", d.TipoCambioDR, "MonedaDR es la MonedaP: no va TipoCambioDR")
	case distinta && d.TipoCambioDR == nil:
		v.hallazgo(ReglaPagoTipoCambioDR, rd, "TipoCambioDR", nil, "MonedaDR "+monedaDR+" distinta de MonedaP "+monedaP+": se requiere TipoCambioDR")
	}
	if clave(d.MetodoDePagoDR) == "PPD" {
		for _, x := range []struct {
			attr  string
			falta bool
		}{{"NumParcialidad", d.NumParcialidad == ""}, {"ImpSaldoAnt", d.ImpSaldoAnt == nil}, {"ImpSaldoInsoluto", d.ImpSaldoInsoluto == nil}} {
			if x.falta {
				v.hallazgo(ReglaPago10Parcialidades, rd, x.attr, nil, "con MetodoDePagoDR PPD se requiere "+x.attr)
			}
		}
	}
	if d.ImpPagado == nil && (n > 1 || d.TipoCambioDR != nil) {
		v.hallazgo(ReglaPago10Parcialidades, rd, "ImpPagado", nil, "con más de un documento relacionado o con TipoCambioDR se requiere ImpPagado")
	}
}

// montoPago compara el Monto con lo pagado en los documentos, en la moneda
// del pago (guía de Pagos 2.0 p. 31-34; guía de Pagos 1.0 p. 20): primero
// la suma de ImpPagado / EquivalenciaDR truncada; si no alcanza y hay
// documentos en otra moneda, la suma de los límites inferiores.
func (v *validadorReglas) montoPago(pg Pago, rp string, v10 bool) {
	if pg.Monto == nil || len(pg.DoctoRelacionado) == 0 {
		return
	}
	monedaP := clave(pg.MonedaP)
	decP, ok := decimalesMoneda[monedaP]
	if !ok {
		decP = pg.Monto.Scale()
	}
	suma := new(big.Rat)
	var minimo decimal.Decimal
	otra := false
	for _, d := range pg.DoctoRelacionado {
		eq := d.EquivalenciaDR
		if v10 {
			eq = d.TipoCambioDR
		}
		misma := clave(d.MonedaDR) == monedaP
		if misma && (eq == nil || v10) {
			eq = &uno
		}
		if d.ImpPagado == nil || eq == nil || eq.IsZero() {
			return
		}
		otra = otra || !misma
		suma.Add(suma, new(big.Rat).Quo(rat(*d.ImpPagado), rat(*eq)))
		dE := eq.Scale()
		if misma {
			eq, dE = &uno, 10
		}
		minimo = minimo.Add(limiteInferiorPagado(*d.ImpPagado, *eq, dE, decP))
	}
	truncada := truncarRat(suma, decP)
	if truncada.Cmp(*pg.Monto) <= 0 || (otra && minimo.Cmp(*pg.Monto) <= 0) {
		return
	}
	msg := fmt.Sprintf("Monto %s menor que la suma de ImpPagado convertidos a MonedaP, %s", pg.Monto.Original(), truncada)
	if otra {
		msg += fmt.Sprintf(", y que la suma de sus límites inferiores, %s", minimo)
	}
	v.hallazgo(ReglaPagoMonto, rp, "Monto", pg.Monto, msg)
}

func (v *validadorReglas) totalesPagos() {
	p := v.c.Pagos
	tot := map[string]*decimal.Decimal{}
	productos := map[string][]termino{} // para el margen (Regla.Margen)
	sumar := func(attr string, x decimal.Decimal) {
		s := valorDe(tot[attr]).Add(x)
		tot[attr] = &s
		productos[attr] = append(productos[attr], termino{escrito: x})
	}
	retenciones := map[string]string{"001": "TotalRetencionesISR", "002": "TotalRetencionesIVA", "003": "TotalRetencionesIEPS"}
	tasasIVA := map[string]string{"0.16": "IVA16", "0.08": "IVA8", "0": "IVA0"}
	for _, pg := range p.Pago {
		tc := pg.TipoCambioP
		if tc == nil && clave(pg.MonedaP) == "MXN" {
			tc = &uno
		}
		if tc == nil || pg.Monto == nil {
			return // sin tipo de cambio no se puede llevar a MXN
		}
		sumar("MontoTotalPagos", pg.Monto.Mul(*tc))
		if pg.ImpuestosP == nil {
			continue
		}
		for _, r := range pg.ImpuestosP.RetencionesP {
			if a, ok := retenciones[clave(r.ImpuestoP)]; ok && r.ImporteP != nil {
				sumar(a, r.ImporteP.Mul(*tc))
			}
		}
		for _, t := range pg.ImpuestosP.TrasladosP {
			if clave(t.ImpuestoP) != "002" || t.BaseP == nil {
				continue
			}
			if clave(t.TipoFactorP) == "Exento" {
				sumar("TotalTrasladosBaseIVAExento", t.BaseP.Mul(*tc))
				continue
			}
			if t.TasaOCuotaP == nil || clave(t.TipoFactorP) != "Tasa" {
				continue
			}
			sufijo, ok := tasasIVA[t.TasaOCuotaP.String()]
			if !ok {
				continue
			}
			sumar("TotalTrasladosBase"+sufijo, t.BaseP.Mul(*tc))
			if t.ImporteP != nil {
				sumar("TotalTrasladosImpuesto"+sufijo, t.ImporteP.Mul(*tc))
			}
		}
	}
	tp := p.Totales
	if tp == nil {
		return // la falta de Totales la reporta la lectura
	}
	declarados := map[string]*decimal.Decimal{
		"TotalRetencionesIVA": tp.TotalRetencionesIVA, "TotalRetencionesISR": tp.TotalRetencionesISR, "TotalRetencionesIEPS": tp.TotalRetencionesIEPS,
		"TotalTrasladosBaseIVA16": tp.TotalTrasladosBaseIVA16, "TotalTrasladosImpuestoIVA16": tp.TotalTrasladosImpuestoIVA16,
		"TotalTrasladosBaseIVA8": tp.TotalTrasladosBaseIVA8, "TotalTrasladosImpuestoIVA8": tp.TotalTrasladosImpuestoIVA8,
		"TotalTrasladosBaseIVA0": tp.TotalTrasladosBaseIVA0, "TotalTrasladosImpuestoIVA0": tp.TotalTrasladosImpuestoIVA0,
		"TotalTrasladosBaseIVAExento": tp.TotalTrasladosBaseIVAExento, "MontoTotalPagos": tp.MontoTotalPagos,
	}
	attrs := make([]string, 0, len(declarados))
	for a := range declarados {
		attrs = append(attrs, a)
	}
	sort.Strings(attrs)
	r := "/Comprobante/Complemento/Pagos/Totales"
	for _, a := range attrs {
		d, t := declarados[a], tot[a]
		switch {
		case t == nil && d != nil && !d.IsZero():
			v.hallazgo(ReglaPagosTotales, r, a, d, fmt.Sprintf("%s %s y ningún pago lo tiene", a, d.Original()))
		case t != nil && d == nil:
			v.hallazgo(ReglaPagosTotales, r, a, nil, fmt.Sprintf("falta %s: los pagos suman %s", a, t.Round(2)))
		case t != nil && !d.Equal(t.Round(2)):
			lo, hi := intervaloSuma(productos[a], 2)
			v.hallazgoDosNiveles(ReglaPagosTotales, r, a, d, fmt.Sprintf("%s %s y los pagos suman %s", a, d.Original(), t.Round(2)),
				d.Cmp(lo) >= 0 && d.Cmp(hi) <= 0, fmt.Sprintf("margen de redondeo por pago, [%s, %s]", lo, hi))
		}
	}
}
