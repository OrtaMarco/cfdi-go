package cfdi

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/OrtaMarco/cfdi-go/decimal"
)

// Reglas de los impuestos de Pagos 2.0: los ImpuestosDR de cada documento
// relacionado y los ImpuestosP de cada pago (matriz de Pagos 2.0 Rev B,
// CRP20246-20279; estándar Rev B p. 61-66).
const (
	ReglaPagoObjetoImpDR    = "REGLA_PAGO_OBJETO_IMP_DR"
	ReglaPagoImpuestoDR     = "REGLA_PAGO_IMPUESTO_DR"
	ReglaPagoImporteDR      = "REGLA_PAGO_IMPORTE_DR"
	ReglaPagoImpuestosP     = "REGLA_PAGO_IMPUESTOS_P"
	ReglaPagoImpuestosPSuma = "REGLA_PAGO_IMPUESTOS_P_SUMA"
)

// conceptosMargenDR es cuántos conceptos de la factura original cubre el
// segundo nivel de ReglaPagoImporteDR: cada uno, con su Importe redondeado a
// los decimales de la moneda, aleja la suma a lo más ½ unidad de
// BaseDR × TasaOCuotaDR, y el CFDI de pago no dice cuántos eran. 20 cubre al
// 99.5 % de las facturas PPD del corpus de regresión; los ImporteDR fuera de
// la letra que se encontraron ahí son todos el impuesto de su factura (3 a 15
// conceptos), a lo más 3.24 centavos de BaseDR × TasaOCuotaDR.
const conceptosMargenDR = 20

// margenImporteDR es el segundo nivel de ReglaPagoImporteDR.
const margenImporteDR = "dos niveles: si no cumple pero cae entre los mismos límites a los decimales de MonedaDR, ampliados en 20 × ½ unidad de esos decimales (±10 unidades: 0.10 en MXN), advertencia: el redondeo por concepto de la factura original, de hasta 20 conceptos, cuyo impuesto (o su proporción del pago) se copia al ImporteDR"

// margenImpuestosP es el segundo nivel de ReglaPagoImpuestosPSuma.
const margenImpuestosP = "dos niveles: si no es la suma pero cae entre la suma de cada ImporteDR (BaseDR) / EquivalenciaDR truncado y la de cada uno redondeado hacia arriba a los decimales de MonedaP (en ImporteP, cada sumando también como BaseDR × TasaOCuotaDR / EquivalenciaDR sin redondear), advertencia: el margen de redondeo por documento, como en los traslados del CFDI"

var reglasPagosImpuestos = []Regla{
	regla(ReglaPagoObjetoImpDR, NivelError,
		"Pagos 2.0: ObjetoImpDR 02 con ImpuestosDR; 01, 03, 04 y 05 sin él; 06 y 08 sin RetencionDR ni TrasladoDR de IVA (002) o IEPS (003); 07 sin IVA y con un TrasladoDR de IEPS",
		citas(cod(DocMatrizPagos20, "CRP20246", "CRP20247", "CRP20278", "CRP20279"), pag(DocEstPagos20, 61), pag(DocGuiaPagos20, 43)),
		"la guía solo menciona 01 y 03; 04, 05, 06, 07 y 08 salen de la matriz Rev B"),
	regla(ReglaPagoImpuestoDR, NivelError,
		"Pagos 2.0: ImpuestosDR con al menos un traslado o una retención; BaseDR mayor que cero; RetencionDR con TipoFactorDR distinto de Exento; TrasladoDR Exento sin TasaOCuotaDR ni ImporteDR, y Tasa o Cuota con los dos",
		citas(cod(DocMatrizPagos20, "CRP20248", "CRP20249", "CRP20252", "CRP20255", "CRP20258", "CRP20259"), pag(DocEstPagos20, 61, 62, 63), pag(DocGuiaPagos20, 43, 44, 46, 47, 48)), ""),
	conMargen(regla(ReglaPagoImporteDR, NivelError,
		"Pagos 2.0: ImporteDR entre (BaseDR − ½·10^-dB)·TasaOCuotaDR truncado y (BaseDR + ½·10^-dB − 10^-12)·TasaOCuotaDR redondeado hacia arriba, a los decimales del propio ImporteDR",
		citas(cod(DocMatrizPagos20, "CRP20254", "CRP20261"), pag(DocEstPagos20, 62, 63), pag(DocGuiaPagos20, 4, 46), pag(DocEstandar40, 77, 78)),
		"el pago «se aplicará proporcionalmente a los conceptos integrados en el comprobante» (guía p. 4): el ImporteDR puede ser el impuesto de la factura, redondeado concepto por concepto (Anexo 20 p. 77, 78), o su proporción del pago"),
		margenImporteDR),
	regla(ReglaPagoImpuestosP, NivelError,
		"Pagos 2.0: con impuestos en los documentos relacionados, ImpuestosP con una RetencionP por cada ImpuestoDR retenido y un TrasladoP por cada ImpuestoDR, TipoFactorDR y TasaOCuotaDR trasladado; sin repetir y sin grupos que no estén en los documentos; TrasladoP Exento sin TasaOCuotaP ni ImporteP",
		citas(cod(DocMatrizPagos20, "CRP20263", "CRP20266", "CRP20271"), pag(DocEstPagos20, 25, 27, 63, 64, 65), pag(DocGuiaPagos20, 49, 50, 51)),
		"RetencionesP y TrasladosP: «Es requerido cuando en los documentos relacionados se registre algún impuesto retenido» («un impuesto trasladado»), estándar p. 25, 27 y guía p. 49, 50; un grupo sin documentos no puede ser «la suma» de CRP20265 y CRP20274"),
	conMargen(regla(ReglaPagoImpuestosPSuma, NivelError,
		"Pagos 2.0: ImporteP de cada RetencionP, y BaseP e ImporteP de cada TrasladoP, iguales a la suma de ImporteDR (BaseDR) / EquivalenciaDR de su grupo en los documentos relacionados del pago (EquivalenciaDR 1 si no viene y MonedaDR es MonedaP): la suma exacta, o redondeada o truncada a los decimales del propio atributo (al menos los de MonedaP)",
		citas(cod(DocMatrizPagos20, "CRP20265", "CRP20268", "CRP20269", "CRP20274"), pag(DocEstPagos20, 64, 65, 66), pag(DocGuiaPagos20, 50, 51, 53)),
		"el SAT no da tolerancia ni dice cómo redondear el cociente («Calcular la suma como: (ImporteDR/EquivalenciaDR)+…», estándar p. 64): se acepta la suma exacta, redondeada o truncada"),
		margenImpuestosP),
}

// techoRat redondea r (positivo) hacia arriba a n decimales.
func techoRat(r *big.Rat, n int32) decimal.Decimal {
	return redondearRatCon(r, n, func(rem, den *big.Int) bool { return rem.Sign() > 0 })
}

// redondearRat redondea r a n decimales, la mitad lejos de cero.
func redondearRat(r *big.Rat, n int32) decimal.Decimal {
	return redondearRatCon(r, n, func(rem, den *big.Int) bool {
		return new(big.Int).Lsh(new(big.Int).Abs(rem), 1).Cmp(den) >= 0
	})
}

// redondearRatCon corta r a n decimales hacia cero y le suma (o resta, si es
// negativo) una unidad del último decimal si sube(resto, denominador).
func redondearRatCon(r *big.Rat, n int32, sube func(rem, den *big.Int) bool) decimal.Decimal {
	x := new(big.Rat).Mul(r, new(big.Rat).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(n)), nil)))
	q, rem := new(big.Int).QuoRem(x.Num(), x.Denom(), new(big.Int))
	if rem.Sign() != 0 && sube(rem, x.Denom()) {
		q.Add(q, big.NewInt(int64(rem.Sign())))
	}
	d, _ := decimal.Parse(q.String())
	return d.Mul(decimal.NewFromInt64(1, n))
}

// terminoRat es un sumando de ImpuestosP en MonedaP: ImporteDR (BaseDR) /
// EquivalenciaDR y, en el ImporteDR, BaseDR × TasaOCuotaDR / EquivalenciaDR
// sin redondear (calculado).
type terminoRat struct {
	escrito, calculado *big.Rat
}

// grupoDR junta los impuestos de los documentos relacionados de un pago por
// (ImpuestoDR, TipoFactorDR, TasaOCuotaDR) en traslados y por ImpuestoDR en
// retenciones, ya convertidos a MonedaP.
type grupoDR struct {
	base, importe   *big.Rat
	bases, importes []terminoRat
}

func claveGrupoDR(impuesto, factor string, tasa *decimal.Decimal, traslado bool) string {
	if !traslado {
		return clave(impuesto)
	}
	t := ""
	if tasa != nil {
		t = tasa.String()
	}
	return clave(impuesto) + "|" + clave(factor) + "|" + t
}

// equivalenciaDR es la EquivalenciaDR con que se convierte el documento a
// MonedaP: la escrita, o 1 si no viene y la moneda es la del pago. nil si no
// se puede convertir.
func equivalenciaDR(pg Pago, d DoctoRelacionado) *big.Rat {
	eq := d.EquivalenciaDR
	if eq == nil && clave(d.MonedaDR) == clave(pg.MonedaP) {
		eq = &uno
	}
	if eq == nil || eq.Sign() <= 0 {
		return nil
	}
	return rat(*eq)
}

// impuestosDR revisa los impuestos de un documento relacionado de Pagos 2.0.
func (v *validadorReglas) impuestosDR(d DoctoRelacionado, rd string) {
	v.objetoImpDR(d, rd)
	imp := d.ImpuestosDR
	if imp == nil {
		return
	}
	ri := rd + "/ImpuestosDR"
	if len(imp.TrasladosDR) == 0 && len(imp.RetencionesDR) == 0 {
		v.hallazgo(ReglaPagoImpuestoDR, ri, "", nil, "ImpuestosDR sin traslados ni retenciones")
	}
	decDR, ok := decimalesMoneda[clave(d.MonedaDR)]
	for j, t := range imp.RetencionesDR {
		r := idx(ri+"/RetencionesDR", "RetencionDR", j)
		if clave(t.TipoFactorDR) == "Exento" {
			v.hallazgoTexto(ReglaPagoImpuestoDR, r, "TipoFactorDR", t.TipoFactorDR, "una RetencionDR no puede ser Exento")
		}
		v.impuestoDR(t, r, decDR, ok)
	}
	for j, t := range imp.TrasladosDR {
		r := idx(ri+"/TrasladosDR", "TrasladoDR", j)
		switch f := clave(t.TipoFactorDR); {
		case f == "Exento" && (t.TasaOCuotaDR != nil || t.ImporteDR != nil):
			v.hallazgoTexto(ReglaPagoImpuestoDR, r, "TipoFactorDR", t.TipoFactorDR, "TrasladoDR Exento con TasaOCuotaDR o ImporteDR")
		case (f == "Tasa" || f == "Cuota") && (t.TasaOCuotaDR == nil || t.ImporteDR == nil):
			v.hallazgoTexto(ReglaPagoImpuestoDR, r, "TipoFactorDR", t.TipoFactorDR, "TrasladoDR "+f+" sin TasaOCuotaDR o sin ImporteDR")
		}
		v.impuestoDR(t, r, decDR, ok)
	}
}

// impuestoDR revisa la BaseDR y los límites del ImporteDR (estándar Rev B
// p. 62-63, la misma fórmula que el Importe del impuesto del concepto).
func (v *validadorReglas) impuestoDR(t ImpuestoDR, r string, decDR int32, decOK bool) {
	if t.BaseDR != nil && t.BaseDR.Sign() <= 0 {
		v.hallazgo(ReglaPagoImpuestoDR, r, "BaseDR", t.BaseDR, "BaseDR tiene que ser mayor que cero")
	}
	if t.BaseDR == nil || t.TasaOCuotaDR == nil || t.ImporteDR == nil {
		return
	}
	inf, sup := limites(*t.BaseDR, *t.TasaOCuotaDR, true, t.ImporteDR.Scale())
	if t.ImporteDR.Cmp(inf) >= 0 && t.ImporteDR.Cmp(sup) <= 0 {
		return
	}
	if !decOK {
		decDR = t.ImporteDR.Scale()
	}
	// Segundo nivel: los límites a los decimales de MonedaDR, ampliados en
	// conceptosMargenDR medias unidades de esos decimales.
	mInf, mSup := limites(*t.BaseDR, *t.TasaOCuotaDR, true, decDR)
	holgura := decimal.NewFromInt64(conceptosMargenDR*5, decDR+1)
	mInf, mSup = mInf.Sub(holgura), mSup.Add(holgura)
	v.hallazgoDosNiveles(ReglaPagoImporteDR, r, "ImporteDR", t.ImporteDR,
		fmt.Sprintf("ImporteDR %s fuera de [%s, %s] (BaseDR %s × TasaOCuotaDR %s)", t.ImporteDR.Original(), inf, sup, t.BaseDR.Original(), t.TasaOCuotaDR.Original()),
		t.ImporteDR.Cmp(mInf) >= 0 && t.ImporteDR.Cmp(mSup) <= 0,
		fmt.Sprintf("margen de redondeo por concepto de la factura (hasta %d conceptos), [%s, %s] a los decimales de MonedaDR", conceptosMargenDR, mInf, mSup))
}

// objetoImpDR revisa ImpuestosDR contra ObjetoImpDR (CRP20246, CRP20247,
// CRP20278, CRP20279), como ObjetoImp en el concepto.
func (v *validadorReglas) objetoImpDR(d DoctoRelacionado, rd string) {
	imp := d.ImpuestosDR
	tiene := func(impuesto string, retenciones bool) bool {
		if imp == nil {
			return false
		}
		lista := imp.TrasladosDR
		if retenciones {
			lista = append(append([]ImpuestoDR(nil), lista...), imp.RetencionesDR...)
		}
		for _, t := range lista {
			if clave(t.ImpuestoDR) == impuesto {
				return true
			}
		}
		return false
	}
	obj := clave(d.ObjetoImpDR)
	msg := ""
	switch obj {
	case "02":
		if imp == nil {
			msg = "ObjetoImpDR 02 (sí objeto de impuesto) y el documento no trae ImpuestosDR"
		}
	case "01", "03", "04", "05":
		if imp != nil {
			msg = "ObjetoImpDR " + obj + " y el documento trae ImpuestosDR"
		}
	case "06", "08":
		if tiene("002", true) || tiene("003", true) {
			msg = "ObjetoImpDR " + obj + " y el documento trae impuestos de IVA o IEPS"
		}
	case "07":
		switch {
		case tiene("002", true):
			msg = "ObjetoImpDR 07 y el documento trae IVA"
		case !tiene("003", false):
			msg = "ObjetoImpDR 07 y el documento no trae el TrasladoDR de IEPS"
		}
	}
	if msg != "" {
		v.hallazgoTexto(ReglaPagoObjetoImpDR, rd, "ObjetoImpDR", d.ObjetoImpDR, msg)
	}
}

// agruparDR suma los impuestos de los documentos relacionados del pago en
// MonedaP. ok es false si algún documento con impuestos no se puede
// convertir (sin EquivalenciaDR en otra moneda): entonces no se comparan.
func agruparDR(pg Pago, traslado bool) (grupos map[string]*grupoDR, orden []string, ok bool) {
	grupos = map[string]*grupoDR{}
	for _, d := range pg.DoctoRelacionado {
		if d.ImpuestosDR == nil {
			continue
		}
		lista := d.ImpuestosDR.RetencionesDR
		if traslado {
			lista = d.ImpuestosDR.TrasladosDR
		}
		if len(lista) == 0 {
			continue
		}
		eq := equivalenciaDR(pg, d)
		if eq == nil {
			return nil, nil, false
		}
		for _, t := range lista {
			k := claveGrupoDR(t.ImpuestoDR, t.TipoFactorDR, t.TasaOCuotaDR, traslado)
			g := grupos[k]
			if g == nil {
				g = &grupoDR{base: new(big.Rat), importe: new(big.Rat)}
				grupos[k] = g
				orden = append(orden, k)
			}
			if t.BaseDR != nil {
				b := new(big.Rat).Quo(rat(*t.BaseDR), eq)
				g.base.Add(g.base, b)
				g.bases = append(g.bases, terminoRat{escrito: b})
			}
			if t.ImporteDR != nil {
				tm := terminoRat{escrito: new(big.Rat).Quo(rat(*t.ImporteDR), eq)}
				if t.BaseDR != nil && t.TasaOCuotaDR != nil {
					tm.calculado = new(big.Rat).Quo(rat(t.BaseDR.Mul(*t.TasaOCuotaDR)), eq)
				}
				g.importe.Add(g.importe, tm.escrito)
				g.importes = append(g.importes, tm)
			}
		}
	}
	return grupos, orden, true
}

// impuestosP compara los ImpuestosP de un pago con los ImpuestosDR de sus
// documentos.
func (v *validadorReglas) impuestosP(pg Pago, rp string) {
	var trasladosP, retencionesP []ImpuestoP
	if pg.ImpuestosP != nil {
		trasladosP, retencionesP = pg.ImpuestosP.TrasladosP, pg.ImpuestosP.RetencionesP
	}
	decP, ok := decimalesMoneda[clave(pg.MonedaP)]
	if !ok {
		decP = -1 // los del propio atributo
	}
	for _, traslado := range []bool{false, true} {
		grupos, orden, convertible := agruparDR(pg, traslado)
		declarados, base, nodo := retencionesP, rp+"/ImpuestosP/RetencionesP", "RetencionP"
		if traslado {
			declarados, base, nodo = trasladosP, rp+"/ImpuestosP/TrasladosP", "TrasladoP"
		}
		vistos := map[string]bool{}
		for j, t := range declarados {
			r := idx(base, nodo, j)
			k := claveGrupoDR(t.ImpuestoP, t.TipoFactorP, t.TasaOCuotaP, traslado)
			if vistos[k] {
				v.hallazgoTexto(ReglaPagoImpuestosP, r, "ImpuestoP", t.ImpuestoP, "se repite el grupo "+describirGrupo(k)+" en "+nodo)
				continue
			}
			vistos[k] = true
			if traslado && clave(t.TipoFactorP) == "Exento" && (t.TasaOCuotaP != nil || t.ImporteP != nil) {
				v.hallazgoTexto(ReglaPagoImpuestosP, r, "TipoFactorP", t.TipoFactorP, "TrasladoP Exento con TasaOCuotaP o ImporteP")
			}
			if !convertible {
				continue
			}
			g := grupos[k]
			if g == nil {
				v.hallazgoTexto(ReglaPagoImpuestosP, r, "ImpuestoP", t.ImpuestoP, "el grupo "+describirGrupo(k)+" no está en ningún documento relacionado del pago")
				continue
			}
			if t.ImporteP != nil && len(g.importes) > 0 {
				v.revisarSumaP(r, "ImporteP", t.ImporteP, g.importe, g.importes, decP, describirGrupo(k), "ImporteDR")
			}
			if traslado && t.BaseP != nil && len(g.bases) > 0 {
				v.revisarSumaP(r, "BaseP", t.BaseP, g.base, g.bases, decP, describirGrupo(k), "BaseDR")
			}
		}
		for _, k := range orden {
			if !vistos[k] {
				v.hallazgo(ReglaPagoImpuestosP, rp, "", nil, "los documentos relacionados tienen "+describirGrupo(k)+" y no está en ImpuestosP/"+strings.TrimPrefix(base, rp+"/ImpuestosP/"))
			}
		}
	}
}

// revisarSumaP evalúa un atributo de ImpuestosP contra la suma de sus
// documentos en MonedaP (ReglaPagoImpuestosPSuma, dos niveles). decP son
// los decimales de MonedaP (-1: los del propio atributo).
func (v *validadorReglas) revisarSumaP(r, attr string, declarado *decimal.Decimal, suma *big.Rat, ts []terminoRat, decP int32, grupo, campoDR string) {
	n := declarado.Scale()
	if decP > n {
		n = decP
	}
	if rat(*declarado).Cmp(suma) == 0 || declarado.Equal(redondearRat(suma, n)) || declarado.Equal(truncarRat(suma, n)) {
		return
	}
	if decP < 0 {
		decP = declarado.Scale()
	}
	var lo, hi decimal.Decimal
	for _, t := range ts {
		a, b := t.escrito, t.escrito
		if c := t.calculado; c != nil {
			if c.Cmp(a) < 0 {
				a = c
			} else {
				b = c
			}
		}
		lo, hi = lo.Add(truncarRat(a, decP)), hi.Add(techoRat(b, decP))
	}
	v.hallazgoDosNiveles(ReglaPagoImpuestosPSuma, r, attr, declarado,
		fmt.Sprintf("%s %s y la suma de %s / EquivalenciaDR de %s en los documentos es %s", attr, declarado.Original(), campoDR, grupo, redondearRat(suma, n+2)),
		declarado.Cmp(lo) >= 0 && declarado.Cmp(hi) <= 0,
		fmt.Sprintf("margen de redondeo por documento, [%s, %s]", lo, hi))
}
