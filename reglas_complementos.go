package cfdi

import (
	"fmt"

	"github.com/OrtaMarco/cfdi-go/decimal"
)

// Códigos de las reglas de Comercio Exterior 2.0 y Nómina 1.2.
const (
	ReglaCCETipoComprobante  = "REGLA_CCE_TIPOCOMPROBANTE"
	ReglaCCECertificado      = "REGLA_CCE_CERTIFICADO"
	ReglaCCETotalUSD         = "REGLA_CCE_TOTALUSD"
	ReglaCCENoIdentificacion = "REGLA_CCE_NOIDENTIFICACION"
	ReglaCCEDatosAduana      = "REGLA_CCE_DATOS_ADUANA"
	ReglaCCEValorDolares     = "REGLA_CCE_VALOR_DOLARES"

	ReglaNominaTotales         = "REGLA_NOMINA_TOTALES"
	ReglaNominaPercepciones    = "REGLA_NOMINA_PERCEPCIONES"
	ReglaNominaDeducciones     = "REGLA_NOMINA_DEDUCCIONES"
	ReglaNominaImportes        = "REGLA_NOMINA_IMPORTES"
	ReglaNominaConcepto        = "REGLA_NOMINA_CONCEPTO"
	ReglaNominaFechas          = "REGLA_NOMINA_FECHAS"
	ReglaNominaSubContratacion = "REGLA_NOMINA_SUBCONTRATACION"
	ReglaNominaOtrosPagos      = "REGLA_NOMINA_OTROS_PAGOS"
)

// Documentos del SAT de los complementos.
const (
	DocMatrizCCE20    = "Matriz de errores Comercio Exterior 2.0"
	DocGuiaCCE20      = "Guía de llenado Comercio Exterior 2.0"
	DocMatrizNomina12 = "Matriz de errores Nómina 1.2 Rev D (2025-02-21)"
	DocGuiaNomina12   = "Guía de llenado Nómina 1.2"
	// DocMatrizNomina12B es la matriz de la Rev B (publicada el
	// 2019-12-17), la primera que trae lo que agregó la Rev B del estándar;
	// fija su vigencia (ver Regla.Vigencia). DocMatrizNomina12C es la de la
	// Rev C, con la numeración actual de esos códigos.
	DocMatrizNomina12B = "Matriz de errores Nómina 1.2 Rev B (2019-12-17)"
	DocMatrizNomina12C = "Matriz de errores Nómina 1.2 Rev C (2021-12-29)"
	// DocEstNomina12 es el estándar del complemento; sus páginas son las
	// del PDF (no las imprime). Su control de cambios (p. 56-57) dice qué
	// agregó cada revisión.
	DocEstNomina12 = "Estándar Nómina 1.2 Rev D (2025-02-21)"
)

var reglasComplementos = []Regla{
	regla(ReglaCCETipoComprobante, NivelError, "Comercio Exterior 2.0: TipoDeComprobante I, E o T",
		citas(cod(DocMatrizCCE20, "CCE101"), pag(DocGuiaCCE20, 13)), ""),
	regla(ReglaCCECertificado, NivelError, "Comercio Exterior 2.0: NumCertificadoOrigen solo con CertificadoOrigen 1",
		citas(cod(DocMatrizCCE20, "CCE118"), pag(DocGuiaCCE20, 43, 44)), ""),
	regla(ReglaCCETotalUSD, NivelError, "Comercio Exterior 2.0: TotalUSD igual a la suma de ValorDolares redondeada a 2 decimales",
		citas(cod(DocMatrizCCE20, "CCE122"), pag(DocGuiaCCE20, 45, 46)), ""),
	regla(ReglaCCENoIdentificacion, NivelError,
		"Comercio Exterior 2.0: cada Concepto con NoIdentificacion, cada Mercancia con un Concepto del mismo NoIdentificacion, sin repetir NoIdentificacion y FraccionArancelaria",
		citas(cod(DocMatrizCCE20, "CCE151", "CCE152", "CCE153"), pag(DocGuiaCCE20, 55)), ""),
	regla(ReglaCCEDatosAduana, NivelError,
		"Comercio Exterior 2.0: CantidadAduana, UnidadAduana y ValorUnitarioAduana van los tres o ninguno, y los tres en todas las mercancías si alguna los trae o si su NoIdentificacion se repite en conceptos o mercancías; con UnidadAduana distinta de 99, ValorUnitarioAduana mayor que cero",
		citas(cod(DocMatrizCCE20, "CCE165", "CCE166", "CCE168"), pag(DocEstCCE20, 43, 44), pag(DocGuiaCCE20, 56, 57)), ""),
	conMargen(regla(ReglaCCEValorDolares, NivelError,
		"Comercio Exterior 2.0: con CantidadAduana, ValorDolares entre (CantidadAduana − ½·10^-dC)·(ValorUnitarioAduana − ½·10^-dV) truncado y (CantidadAduana + ½·10^-dC − 10^-12)·(ValorUnitarioAduana + ½·10^-dV − 10^-12) redondeado hacia arriba, a 2 decimales, con dC y dV los decimales de Cantidad y ValorUnitario del concepto del mismo NoIdentificacion; ValorDolares 0 o 1 no se revisa",
		citas(cod(DocMatrizCCE20, "CCE169", "CCE171"), pag(DocEstCCE20, 44), pag(DocGuiaCCE20, 59, 60, 72)),
		"«NumDecimalesCantidad» y «NumDecimalesValorUnitario» (estándar p. 44) son los del concepto: el ejemplo de la guía p. 72 usa 10^-5 de la Cantidad 482.60594 y no 10^-3 de la CantidadAduana 482,605.940"),
		"dos niveles: en I y E, si ValorDolares cae fuera de esos límites pero dentro de los de REGLA_CCE_VALOR_IMPORTE con el Importe de su concepto (un solo concepto y una sola mercancía con ese NoIdentificacion), advertencia: la guía (p. 72) dice que ValorDolares es el Importe entre TipoCambioUSD y que multiplicar CantidadAduana por ValorUnitarioAduana es «el cálculo incorrecto»; lo que no cuadra son los datos de aduana"),
	regla(ReglaNominaTotales, NivelError,
		"Nómina 1.2: TotalPercepciones o TotalOtrosPagos; TotalPercepciones solo con Percepciones y = TotalSueldos + TotalSeparacionIndemnizacion + TotalJubilacionPensionRetiro; TotalDeducciones solo con Deducciones y = TotalOtrasDeducciones + TotalImpuestosRetenidos; con OtrosPagos, TotalOtrosPagos = Σ Importe",
		citas(cod(DocMatrizNomina12, "NOM31", "NOM36", "NOM37", "NOM38", "NOM39", "NOM40"), pag(DocGuiaNomina12, 26, 27)), ""),
	regla(ReglaNominaPercepciones, NivelError,
		"Nómina 1.2: TotalSueldos, TotalSeparacionIndemnizacion (022, 023, 025) y TotalJubilacionPensionRetiro (039, 044) suman ImporteGravado + ImporteExento de sus percepciones; TotalGravado y TotalExento, sus importes; y los tres primeros suman TotalGravado + TotalExento",
		citas(cod(DocMatrizNomina12, "NOM69", "NOM70", "NOM71", "NOM72", "NOM73", "NOM74"), pag(DocGuiaNomina12, 48, 49, 50, 51)), ""),
	regla(ReglaNominaDeducciones, NivelError, "Nómina 1.2: TotalImpuestosRetenidos igual a la suma de las deducciones 002; sin deducciones 002, ausente",
		citas(cod(DocMatrizNomina12, "NOM91", "NOM92"), pag(DocGuiaNomina12, 66)), ""),
	regla(ReglaNominaImportes, NivelError,
		"Nómina 1.2: ImporteGravado + ImporteExento de cada percepción, Importe de cada deducción e Importe de cada otro pago distinto de 002, mayores que cero",
		citas(cod(DocMatrizNomina12, "NOM75", "NOM96", "NOM100"), pag(DocGuiaNomina12, 56, 69, 76)), ""),
	regla(ReglaNominaConcepto, NivelError,
		"Nómina 1.2: el concepto con ValorUnitario e Importe = TotalPercepciones + TotalOtrosPagos y Descuento = TotalDeducciones",
		citas(cod(DocMatrizNomina12, "NOM20", "NOM21", "NOM22"), pag(DocGuiaNomina12, 19)), ""),
	regla(ReglaNominaFechas, NivelError, "Nómina 1.2: FechaInicialPago menor o igual que FechaFinalPago",
		citas(cod(DocMatrizNomina12, "NOM35"), pag(DocGuiaNomina12, 24)), ""),
	regla(ReglaNominaSubContratacion, NivelError, "Nómina 1.2: la suma de PorcentajeTiempo de SubContratacion es 100",
		citas(cod(DocMatrizNomina12, "NOM68"), pag(DocGuiaNomina12, 48)), ""),
	regla(ReglaNominaOtrosPagos, NivelError, "Nómina 1.2: SaldoAFavor mayor o igual que RemanenteSalFav; en el otro pago 002, Importe menor o igual que SubsidioCausado",
		citas(cod(DocMatrizNomina12, "NOM102", "NOM107"), pag(DocGuiaNomina12, 77)), ""),
}

// comercioExterior evalúa las reglas de Comercio Exterior 2.0 que no
// necesitan red ni catálogos.
func (v *validadorReglas) comercioExterior() {
	v.exportacionSinCCE()
	ce := v.c.ComercioExterior
	if ce == nil || ce.Version != "2.0" || !v.v40 {
		return
	}
	r := "/Comprobante/Complemento/ComercioExterior"
	if v.tipo != "" && v.tipo != "I" && v.tipo != "E" && v.tipo != "T" {
		v.hallazgoTexto(ReglaCCETipoComprobante, "/Comprobante", "TipoDeComprobante", v.c.TipoDeComprobante, "Comercio Exterior va en un comprobante I, E o T")
	}
	if clave(ce.CertificadoOrigen) != "1" && ce.NumCertificadoOrigen != "" {
		v.hallazgoTexto(ReglaCCECertificado, r, "NumCertificadoOrigen", ce.NumCertificadoOrigen, "sin CertificadoOrigen 1 no va NumCertificadoOrigen")
	}
	// Conceptos por NoIdentificacion.
	conceptos := map[string][]Concepto{}
	for i, cp := range v.c.Conceptos {
		if cp.NoIdentificacion == "" {
			v.hallazgo(ReglaCCENoIdentificacion, idx("/Comprobante/Conceptos", "Concepto", i), "NoIdentificacion", nil, "con Comercio Exterior cada concepto lleva NoIdentificacion")
			continue
		}
		conceptos[cp.NoIdentificacion] = append(conceptos[cp.NoIdentificacion], cp)
	}
	// CCE165 y CCE166: los datos de aduana son obligatorios en todas las
	// mercancías si alguna los trae, y en las de un NoIdentificacion que se
	// repite en conceptos o en mercancías.
	conAduana := false
	mercanciasPor := map[string]int{}
	for _, m := range ce.Mercancias {
		conAduana = conAduana || m.CantidadAduana != nil || m.UnidadAduana != "" || m.ValorUnitarioAduana != nil
		mercanciasPor[m.NoIdentificacion]++
	}
	var suma decimal.Decimal
	todos := true
	vistos := map[string]bool{}
	for i, m := range ce.Mercancias {
		rm := idx(r+"/Mercancias", "Mercancia", i)
		if m.ValorDolares == nil {
			todos = false
		} else {
			suma = suma.Add(*m.ValorDolares)
		}
		if k := m.NoIdentificacion + "|" + m.FraccionArancelaria; vistos[k] {
			v.hallazgoTexto(ReglaCCENoIdentificacion, rm, "NoIdentificacion", m.NoIdentificacion, "se repite la mercancía con NoIdentificacion "+m.NoIdentificacion+" y FraccionArancelaria "+m.FraccionArancelaria)
		} else {
			vistos[k] = true
		}
		cs := conceptos[m.NoIdentificacion]
		if len(cs) == 0 {
			v.hallazgoTexto(ReglaCCENoIdentificacion, rm, "NoIdentificacion", m.NoIdentificacion, "ningún concepto tiene el NoIdentificacion "+m.NoIdentificacion)
		}
		presentes := 0
		for _, ok := range []bool{m.CantidadAduana != nil, m.UnidadAduana != "", m.ValorUnitarioAduana != nil} {
			if ok {
				presentes++
			}
		}
		switch {
		case presentes != 0 && presentes != 3:
			v.hallazgo(ReglaCCEDatosAduana, rm, "", nil, "CantidadAduana, UnidadAduana y ValorUnitarioAduana van los tres o ninguno")
		case presentes == 0 && conAduana:
			v.hallazgo(ReglaCCEDatosAduana, rm, "", nil, "otra mercancía trae CantidadAduana, UnidadAduana y ValorUnitarioAduana: van en todas")
		case presentes == 0 && (len(cs) > 1 || mercanciasPor[m.NoIdentificacion] > 1):
			v.hallazgo(ReglaCCEDatosAduana, rm, "", nil, "el NoIdentificacion "+m.NoIdentificacion+" se repite en conceptos o mercancías: van CantidadAduana, UnidadAduana y ValorUnitarioAduana")
		}
		if u := clave(m.UnidadAduana); u != "" && u != "99" && m.ValorUnitarioAduana != nil && m.ValorUnitarioAduana.Sign() <= 0 {
			v.hallazgo(ReglaCCEDatosAduana, rm, "ValorUnitarioAduana", m.ValorUnitarioAduana, "con UnidadAduana "+u+" el ValorUnitarioAduana tiene que ser mayor que cero")
		}
		v.valorDolares(m, rm, cs, mercanciasPor[m.NoIdentificacion] == 1)
	}
	if ce.TotalUSD != nil && todos && len(ce.Mercancias) > 0 && !ce.TotalUSD.Equal(suma.Round(2)) {
		v.hallazgo(ReglaCCETotalUSD, r, "TotalUSD", ce.TotalUSD, fmt.Sprintf("TotalUSD %s y la suma de ValorDolares redondeada es %s", ce.TotalUSD.Original(), suma.Round(2)))
	}
	v.cce(ce, conceptos)
}

// valorDolares revisa los límites de ValorDolares con CantidadAduana
// (estándar de Comercio Exterior 2.0 p. 44, guía p. 59 y el ejemplo de la
// p. 72): CantidadAduana × ValorUnitarioAduana, cada uno ± ½·10^-d con d
// los decimales de Cantidad y ValorUnitario del concepto (el ejemplo usa
// 10^-5 de la Cantidad 482.60594, no 10^-3 de la CantidadAduana
// 482605.940), el inferior truncado y el superior hacia arriba a
// centésimas. Si hay más de un concepto con el mismo NoIdentificacion no
// se sabe de cuál salen los decimales y no se revisa.
//
// Segundo nivel: la guía (p. 72) dice que ValorDolares es el Importe del
// concepto entre TipoCambioUSD y que multiplicar CantidadAduana por
// ValorUnitarioAduana es «el cálculo incorrecto»; un ValorDolares fuera de
// esos límites pero que sí es el Importe en dólares (REGLA_CCE_VALOR_IMPORTE,
// en I y E, con un solo concepto y una sola mercancía de ese
// NoIdentificacion) es advertencia: lo que no cuadra son los datos de
// aduana. unica dice si la mercancía es la única de su NoIdentificacion.
func (v *validadorReglas) valorDolares(m ComercioExteriorMercancia, rm string, cs []Concepto, unica bool) {
	vd := m.ValorDolares
	if vd == nil || m.CantidadAduana == nil || m.ValorUnitarioAduana == nil || len(cs) != 1 || cs[0].Cantidad == nil || cs[0].ValorUnitario == nil {
		return
	}
	if vd.IsZero() || vd.Equal(uno) {
		return
	}
	dC, dV := mitadUlp(cs[0].Cantidad.Scale()), mitadUlp(cs[0].ValorUnitario.Scale())
	inf := truncar(m.CantidadAduana.Sub(dC).Mul(m.ValorUnitarioAduana.Sub(dV)), 2)
	sup := techo(m.CantidadAduana.Add(dC).Sub(casiCero).Mul(m.ValorUnitarioAduana.Add(dV).Sub(casiCero)), 2)
	if vd.Cmp(inf) >= 0 && vd.Cmp(sup) <= 0 {
		return
	}
	msg := fmt.Sprintf("ValorDolares %s fuera de [%s, %s] (CantidadAduana %s × ValorUnitarioAduana %s, con los decimales de Cantidad %s y ValorUnitario %s del concepto)",
		vd.Original(), inf, sup, m.CantidadAduana.Original(), m.ValorUnitarioAduana.Original(), cs[0].Cantidad.Original(), cs[0].ValorUnitario.Original())
	var is [][2]decimal.Decimal
	if (v.tipo == "I" || v.tipo == "E") && unica && cs[0].Importe != nil {
		is = v.intervalosUSD(*cs[0].Importe)
	}
	if is == nil {
		v.hallazgo(ReglaCCEValorDolares, rm, "ValorDolares", vd, msg)
		return
	}
	v.hallazgoDosNiveles(ReglaCCEValorDolares, rm, "ValorDolares", vd, msg, enIntervalos(*vd, is),
		fmt.Sprintf("Importe %s del concepto en dólares, %s (guía p. 72: ValorDolares es el Importe entre TipoCambioUSD; lo que no cuadra son los datos de aduana)", cs[0].Importe.Original(), describirIntervalos(is)))
}

// igualSuma dice si d (presente) es la suma s.
func igualSuma(d *decimal.Decimal, s decimal.Decimal) bool { return d == nil || d.Equal(s) }

// nomina evalúa las reglas aritméticas y de coherencia de Nómina 1.2.
func (v *validadorReglas) nomina() {
	for i, n := range v.c.Nomina {
		if n.Version != "1.2" {
			continue
		}
		r := "/Comprobante/Complemento/Nomina"
		if len(v.c.Nomina) > 1 {
			r = idx("/Comprobante/Complemento", "Nomina", i)
		}
		v.nominaTotales(n, r)
		v.nominaT8(n, r)
		if p := n.Percepciones; p != nil {
			v.percepciones(p, r+"/Percepciones")
		}
		if d := n.Deducciones; d != nil {
			v.deducciones(d, r+"/Deducciones")
		}
		for j, o := range n.OtrosPagos {
			ro := idx(r+"/OtrosPagos", "OtroPago", j)
			tipo := clave(o.TipoOtroPago)
			if tipo != "002" && o.Importe != nil && o.Importe.Sign() <= 0 {
				v.hallazgo(ReglaNominaImportes, ro, "Importe", o.Importe, "el Importe del otro pago tiene que ser mayor que cero")
			}
			if s := o.SubsidioAlEmpleo; tipo == "002" && s != nil && s.SubsidioCausado != nil && o.Importe != nil && o.Importe.Cmp(*s.SubsidioCausado) > 0 {
				v.hallazgo(ReglaNominaOtrosPagos, ro, "Importe", o.Importe, fmt.Sprintf("Importe %s mayor que SubsidioCausado %s", o.Importe.Original(), s.SubsidioCausado.Original()))
			}
			if c := o.CompensacionSaldosAFavor; c != nil && c.SaldoAFavor != nil && c.RemanenteSalFav != nil && c.SaldoAFavor.Cmp(*c.RemanenteSalFav) < 0 {
				v.hallazgo(ReglaNominaOtrosPagos, ro+"/CompensacionSaldosAFavor", "SaldoAFavor", c.SaldoAFavor,
					fmt.Sprintf("SaldoAFavor %s menor que RemanenteSalFav %s", c.SaldoAFavor.Original(), c.RemanenteSalFav.Original()))
			}
		}
		if n.FechaInicialPago.Valida() && n.FechaFinalPago.Valida() && n.FechaInicialPago.Time.After(n.FechaFinalPago.Time) {
			v.hallazgoTexto(ReglaNominaFechas, r, "FechaInicialPago", n.FechaInicialPago.Original, "FechaInicialPago posterior a FechaFinalPago "+n.FechaFinalPago.Original)
		}
		if rc := n.Receptor; rc != nil && len(rc.SubContratacion) > 0 {
			var s decimal.Decimal
			for _, sc := range rc.SubContratacion {
				s = s.Add(valorDe(sc.PorcentajeTiempo))
			}
			if !s.Equal(decimal.NewFromInt64(100, 0)) {
				v.hallazgo(ReglaNominaSubContratacion, r+"/Receptor", "", nil, fmt.Sprintf("los PorcentajeTiempo de SubContratacion suman %s y no 100", s))
			}
		}
	}
	v.conceptoNomina()
	v.comprobanteNomina()
}

func (v *validadorReglas) nominaTotales(n Nomina, r string) {
	if n.TotalPercepciones == nil && n.TotalOtrosPagos == nil {
		v.hallazgo(ReglaNominaTotales, r, "", nil, "la nómina no tiene TotalPercepciones ni TotalOtrosPagos")
	}
	if n.Percepciones == nil && n.TotalPercepciones != nil {
		v.hallazgo(ReglaNominaTotales, r, "TotalPercepciones", n.TotalPercepciones, "sin nodo Percepciones no va TotalPercepciones")
	}
	if n.Deducciones == nil && n.TotalDeducciones != nil {
		v.hallazgo(ReglaNominaTotales, r, "TotalDeducciones", n.TotalDeducciones, "sin nodo Deducciones no va TotalDeducciones")
	}
	if p := n.Percepciones; p != nil && n.TotalPercepciones != nil {
		s := valorDe(p.TotalSueldos).Add(valorDe(p.TotalSeparacionIndemnizacion)).Add(valorDe(p.TotalJubilacionPensionRetiro))
		if !n.TotalPercepciones.Equal(s) {
			v.hallazgo(ReglaNominaTotales, r, "TotalPercepciones", n.TotalPercepciones, fmt.Sprintf("TotalPercepciones %s y los totales de Percepciones suman %s", n.TotalPercepciones.Original(), s))
		}
	}
	if d := n.Deducciones; d != nil && n.TotalDeducciones != nil {
		s := valorDe(d.TotalOtrasDeducciones).Add(valorDe(d.TotalImpuestosRetenidos))
		if !n.TotalDeducciones.Equal(s) {
			v.hallazgo(ReglaNominaTotales, r, "TotalDeducciones", n.TotalDeducciones, fmt.Sprintf("TotalDeducciones %s y los totales de Deducciones suman %s", n.TotalDeducciones.Original(), s))
		}
	}
	if len(n.OtrosPagos) > 0 {
		var s decimal.Decimal
		for _, o := range n.OtrosPagos {
			s = s.Add(valorDe(o.Importe))
		}
		switch {
		case n.TotalOtrosPagos == nil:
			v.hallazgo(ReglaNominaTotales, r, "TotalOtrosPagos", nil, "con OtrosPagos se requiere TotalOtrosPagos: suman "+s.String())
		case !n.TotalOtrosPagos.Equal(s):
			v.hallazgo(ReglaNominaTotales, r, "TotalOtrosPagos", n.TotalOtrosPagos, fmt.Sprintf("TotalOtrosPagos %s y los otros pagos suman %s", n.TotalOtrosPagos.Original(), s))
		}
	}
}

// Claves de TipoPercepcion que van a TotalSeparacionIndemnizacion y a
// TotalJubilacionPensionRetiro (NOM71, NOM72); las demás, a TotalSueldos.
var (
	percepcionesSeparacion = map[string]bool{"022": true, "023": true, "025": true}
	percepcionesJubilacion = map[string]bool{"039": true, "044": true}
)

func (v *validadorReglas) percepciones(p *Percepciones, r string) {
	var sueldos, separacion, jubilacion, gravado, exento decimal.Decimal
	for j, pc := range p.Percepcion {
		g, e := valorDe(pc.ImporteGravado), valorDe(pc.ImporteExento)
		gravado, exento = gravado.Add(g), exento.Add(e)
		ge := g.Add(e)
		switch tipo := clave(pc.TipoPercepcion); {
		case percepcionesSeparacion[tipo]:
			separacion = separacion.Add(ge)
		case percepcionesJubilacion[tipo]:
			jubilacion = jubilacion.Add(ge)
		default:
			sueldos = sueldos.Add(ge)
		}
		if pc.ImporteGravado != nil && pc.ImporteExento != nil && ge.Sign() <= 0 {
			v.hallazgo(ReglaNominaImportes, idx(r, "Percepcion", j), "ImporteGravado", pc.ImporteGravado, "ImporteGravado + ImporteExento tiene que ser mayor que cero")
		}
	}
	totales := []struct {
		attr  string
		valor *decimal.Decimal
		suma  decimal.Decimal
	}{
		{"TotalSueldos", p.TotalSueldos, sueldos},
		{"TotalSeparacionIndemnizacion", p.TotalSeparacionIndemnizacion, separacion},
		{"TotalJubilacionPensionRetiro", p.TotalJubilacionPensionRetiro, jubilacion},
		{"TotalGravado", p.TotalGravado, gravado},
		{"TotalExento", p.TotalExento, exento},
	}
	for _, t := range totales {
		if !igualSuma(t.valor, t.suma) {
			v.hallazgo(ReglaNominaPercepciones, r, t.attr, t.valor, fmt.Sprintf("%s %s y sus percepciones suman %s", t.attr, t.valor.Original(), t.suma))
		}
	}
	if p.TotalGravado != nil && p.TotalExento != nil {
		tres := valorDe(p.TotalSueldos).Add(valorDe(p.TotalSeparacionIndemnizacion)).Add(valorDe(p.TotalJubilacionPensionRetiro))
		if ge := p.TotalGravado.Add(*p.TotalExento); !tres.Equal(ge) {
			v.hallazgo(ReglaNominaPercepciones, r, "", nil, fmt.Sprintf("TotalSueldos + TotalSeparacionIndemnizacion + TotalJubilacionPensionRetiro es %s y TotalGravado + TotalExento %s", tres, ge))
		}
	}
}

func (v *validadorReglas) deducciones(d *Deducciones, r string) {
	var impuestos decimal.Decimal
	hay := false
	for j, dd := range d.Deduccion {
		if clave(dd.TipoDeduccion) == "002" {
			hay = true
			impuestos = impuestos.Add(valorDe(dd.Importe))
		}
		if dd.Importe != nil && dd.Importe.Sign() <= 0 {
			v.hallazgo(ReglaNominaImportes, idx(r, "Deduccion", j), "Importe", dd.Importe, "el Importe de la deducción tiene que ser mayor que cero")
		}
	}
	switch {
	case !hay && d.TotalImpuestosRetenidos != nil:
		v.hallazgo(ReglaNominaDeducciones, r, "TotalImpuestosRetenidos", d.TotalImpuestosRetenidos, "sin deducciones 002 no va TotalImpuestosRetenidos")
	case hay && !igualSuma(d.TotalImpuestosRetenidos, impuestos):
		v.hallazgo(ReglaNominaDeducciones, r, "TotalImpuestosRetenidos", d.TotalImpuestosRetenidos,
			fmt.Sprintf("TotalImpuestosRetenidos %s y las deducciones 002 suman %s", d.TotalImpuestosRetenidos.Original(), impuestos))
	}
}

// conceptoNomina compara el concepto único de un CFDI de nómina con los
// totales del complemento (NOM20-22). Solo con un complemento y un
// concepto: con varios, el SAT no dice cómo se reparten.
func (v *validadorReglas) conceptoNomina() {
	c := v.c
	if v.tipo != "N" || len(c.Nomina) != 1 || c.Nomina[0].Version != "1.2" || len(c.Conceptos) != 1 {
		return
	}
	n, cp := c.Nomina[0], c.Conceptos[0]
	r := "/Comprobante/Conceptos/Concepto[1]"
	if n.TotalPercepciones != nil || n.TotalOtrosPagos != nil {
		s := valorDe(n.TotalPercepciones).Add(valorDe(n.TotalOtrosPagos))
		for _, x := range []struct {
			attr  string
			valor *decimal.Decimal
		}{{"ValorUnitario", cp.ValorUnitario}, {"Importe", cp.Importe}} {
			if !igualSuma(x.valor, s) {
				v.hallazgo(ReglaNominaConcepto, r, x.attr, x.valor, fmt.Sprintf("%s %s y TotalPercepciones + TotalOtrosPagos es %s", x.attr, x.valor.Original(), s))
			}
		}
	}
	d := valorDe(n.TotalDeducciones)
	switch {
	case n.TotalDeducciones != nil && cp.Descuento == nil:
		v.hallazgo(ReglaNominaConcepto, r, "Descuento", nil, "falta el Descuento: TotalDeducciones es "+d.String())
	case cp.Descuento != nil && !cp.Descuento.Equal(d):
		v.hallazgo(ReglaNominaConcepto, r, "Descuento", cp.Descuento, fmt.Sprintf("Descuento %s y TotalDeducciones %s", cp.Descuento.Original(), d))
	}
}
