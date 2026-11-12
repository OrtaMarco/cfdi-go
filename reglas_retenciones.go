package cfdi

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/OrtaMarco/cfdi-go/decimal"
)

// Códigos de hallazgo de ValidarReglasRetenciones. Son parte de la API
// pública, como los de ValidarReglas: no se renombran.
const (
	ReglaRetMontoOperacion = "REGLA_RET_MONTO_OPERACION"
	ReglaRetMontoMayor     = "REGLA_RET_MONTO_MAYOR_OPERACION"
	ReglaRetMontoTotRet    = "REGLA_RET_MONTO_TOT_RET"
	ReglaRet10MontoTotRet  = "REGLA_RET10_MONTO_TOT_RET"
	ReglaRetPeriodo        = "REGLA_RET_PERIODO"
	ReglaRetEjercicio      = "REGLA_RET_EJERCICIO"
	ReglaRetTipoPago       = "REGLA_RET_TIPO_PAGO"
	ReglaRetDescRetenc     = "REGLA_RET_DESC_RETENC"
	ReglaRetCveComplemento = "REGLA_RET_CVE_COMPLEMENTO"
	ReglaRetCveRIF         = "REGLA_RET_CVE_RIF"
	ReglaRetReceptor       = "REGLA_RET_RECEPTOR_NACIONAL"
)

// Documentos del SAT del CFDI de Retenciones e información de pagos.
const (
	// DocMatrizRet20 es la matriz de errores de Retenciones 2.0 (códigos
	// Reten20xxx), versión del portal del 2025-04-09. El SAT no publica
	// matriz de errores de Retenciones 1.0.
	DocMatrizRet20 = "Matriz de errores Retenciones 2.0 (2025-04-09)"
	// DocEstRet20 es la sección II del Anexo 20 publicado en el DOF el
	// 2022-01-13 (el mismo PDF de DocEstandar40): el estándar de Retenciones
	// 2.0 empieza en la página 83 del PDF y sus validaciones están en las
	// 116 a 119.
	DocEstRet20 = "Anexo 20 Retenciones 2.0 (DOF 2022-01-13)"
	// DocEstRet10 es la sección II del Anexo 20 publicado en el DOF el
	// 2017-07-28 (el mismo PDF de DocEstandar33): el estándar de
	// Retenciones 1.0 empieza en la página 73 del PDF.
	DocEstRet10 = "Anexo 20 Retenciones 1.0 (DOF 2017-07-28)"
	// DocCatRet20 son los catálogos de Retenciones 2.0 (xls del portal,
	// vigencia 2022-01-01). Se cita por hoja: lo que no vienen en el XSD son
	// la columna «Nombre del complemento» de c_CveRetenc y la columna «Tipo
	// impuesto» de c_TipoPagoRet.
	DocCatRet20 = "Catálogos Retenciones 2.0 (2022-01-01)"
)

var reglasRetenciones = []Regla{
	regla(ReglaRetMontoOperacion, NivelError, "Retenciones 2.0: MontoTotOperacion igual a MontoTotGrav + MontoTotExent",
		citas(cod(DocMatrizRet20, "Reten20127"), pag(DocEstRet20, 118)), ""),
	regla(ReglaRetMontoMayor, NivelError, "Retenciones 2.0: MontoTotGrav y MontoTotExent menores o iguales que MontoTotOperacion",
		citas(cod(DocMatrizRet20, "Reten20128", "Reten20129"), pag(DocEstRet20, 118)), ""),
	regla(ReglaRetMontoTotRet, NivelError, "Retenciones 2.0: con MontoTotRet mayor que cero, al menos un ImpRetenidos y MontoTotRet igual a la suma de sus MontoRet",
		citas(cod(DocMatrizRet20, "Reten20130", "Reten20131"), pag(DocEstRet20, 97, 118)),
		"con MontoTotRet en cero el SAT no pide nada: «si el valor es mayor que cero»"),
	regla(ReglaRet10MontoTotRet, NivelAdvertencia, "Retenciones 1.0: montoTotRet igual a la suma de los montoRet de los ImpRetenidos",
		citas(pag(DocEstRet10, 84)),
		"advertencia: el estándar 1.0 lo describe («Sumatoria de los montos de retención del nodo ImpRetenidos») pero no tiene matriz de errores ni regla de validación, y los ejemplos públicos 1.0 de 2014-2015 del repo (satcfdi) no lo cumplen"),
	regla(ReglaRetPeriodo, NivelError, "Retenciones 2.0: MesIni menor o igual que MesFin",
		citas(cod(DocMatrizRet20, "Reten20122", "Reten20124"), pag(DocEstRet20, 118)), ""),
	regla(ReglaRetEjercicio, NivelError, "Retenciones 2.0: Ejercicio no posterior al año de FechaExp ni anterior en más de 5 años",
		citas(cod(DocMatrizRet20, "Reten20126"), pag(DocEstRet20, 118)),
		"el estándar (2022) pide el año de FechaExp o el inmediato anterior; la matriz vigente (2025) admite hasta 5 ejercicios anteriores y se aplica la matriz; «el año en curso» es el de FechaExp, como dice el estándar"),
	regla(ReglaRetTipoPago, NivelError,
		"Retenciones 2.0: con ImpuestoRet 001 (ISR), TipoPagoRet 03 o 04; con 002 (IVA), 01; con 003 (IEPS), 02",
		citas(cod(DocMatrizRet20, "Reten20135"), pag(DocEstRet20, 119), cod(DocCatRet20, "c_TipoPagoRet")),
		"la columna «Tipo impuesto» de c_TipoPagoRet no viene en el XSD: son sus cuatro filas"),
	regla(ReglaRetDescRetenc, NivelError, "Retenciones 2.0: con CveRetenc 25 (otro tipo de retenciones), DescRetenc",
		citas(cod(DocMatrizRet20, "Reten20106"), pag(DocEstRet20, 117)), ""),
	regla(ReglaRetCveComplemento, NivelError,
		"Retenciones 2.0: el complemento que c_CveRetenc asocia a la clave: 06 y 19 EnajenaciondeAcciones; 14 y 15 Dividendos; 16 Intereses; 17 Arrendamientoenfideicomiso; 18 Pagosaextranjeros; 20 Premios; 21 Fideicomisonoempresarial; 22 Planesderetiro; 23 Intereseshipotecarios; 24 Operacionesconderivados; 26 ServiciosPlataformasTecnologicas",
		citas(cod(DocMatrizRet20, "Reten20107"), pag(DocEstRet20, 117), cod(DocCatRet20, "c_CveRetenc")),
		"las claves 01, 02, 04 y 05 (Pagosaextranjeros «en caso de operaciones realizadas con extranjeros») y 27 (SectorFinanciero «solo cuando» el integrante del sistema financiero sea fiduciario) dependen de una condición que ningún atributo expresa y no se revisan; la columna «Nombre del complemento» no viene en el XSD"),
	regla(ReglaRetCveRIF, NivelError,
		"Retenciones 2.0: con CveRetenc 28 (RIF), MontoTotGrav y MontoTotExent en cero, y UtilidadBimestral e ISRCorrespondiente mayores que cero",
		citas(cod(DocMatrizRet20, "Reten20108", "Reten20132", "Reten20133"), pag(DocEstRet20, 118, 119)), ""),
	regla(ReglaRetReceptor, NivelError,
		"Retenciones 2.0, Receptor/Nacional: sin CurpR si RfcR tiene 12 posiciones; NomDenRazSocR «PÚBLICO EN GENERAL» solo con RfcR XAXX010101000; con XAXX010101000, DomicilioFiscalR igual a LugarExpRetenc",
		citas(cod(DocMatrizRet20, "Reten20116", "Reten20117", "Reten20120"), pag(DocEstRet20, 117, 118)),
		"Reten20116 y Reten20120 son solo de la matriz: el estándar (p. 117) pide únicamente que NomDenRazSocR y DomicilioFiscalR estén asociados al RFC en la lista del SAT, que necesita red"),
}

// ReglasRetenciones son las reglas de ValidarReglasRetenciones, en el orden
// de Reglas (donde van al final).
var ReglasRetenciones = reglasRetenciones

// complementoCve es el complemento que c_CveRetenc asocia a una clave: el
// nombre del complemento (el nombre local de su elemento) y si el modelo lo
// trae.
type complementoCve struct {
	nombre    string
	presente  func(*Retenciones) bool
	nombreCat string // «Nombre del complemento» del catálogo
}

var complementosPorCve = func() map[string]complementoCve {
	enaj := complementoCve{"EnajenaciondeAcciones", func(r *Retenciones) bool { return r.EnajenaciondeAcciones != nil }, "Enajenación de Acciones"}
	divi := complementoCve{"Dividendos", func(r *Retenciones) bool { return r.Dividendos != nil }, "Dividendos"}
	return map[string]complementoCve{
		"06": enaj, "19": enaj, "14": divi, "15": divi,
		"16": {"Intereses", func(r *Retenciones) bool { return r.Intereses != nil }, "Intereses"},
		"17": {"Arrendamientoenfideicomiso", func(r *Retenciones) bool { return r.Arrendamientoenfideicomiso != nil }, "Arrendamiento en fideicomiso"},
		"18": {"Pagosaextranjeros", func(r *Retenciones) bool { return r.Pagosaextranjeros != nil }, "Pagos a Extranjeros"},
		"20": {"Premios", func(r *Retenciones) bool { return r.Premios != nil }, "Premios"},
		"21": {"Fideicomisonoempresarial", func(r *Retenciones) bool { return r.Fideicomisonoempresarial != nil }, "Fideicomiso no Empresarial"},
		"22": {"Planesderetiro", func(r *Retenciones) bool { return r.Planesderetiro != nil }, "Planes de Retiro"},
		"23": {"Intereseshipotecarios", func(r *Retenciones) bool { return r.Intereseshipotecarios != nil }, "Intereses Hipotecarios"},
		"24": {"Operacionesconderivados", func(r *Retenciones) bool { return r.Operacionesconderivados != nil }, "Operaciones con Derivados"},
		"26": {"ServiciosPlataformasTecnologicas", func(r *Retenciones) bool { return r.ServiciosPlataformasTecnologicas != nil }, "Servicios mediante Plataformas Tecnológicas"},
	}
}()

// tipoPagoPorImpuesto es la columna «Tipo impuesto» de c_TipoPagoRet: los
// ImpuestoRet (c_Impuesto) que admite cada TipoPagoRet.
var tipoPagoPorImpuesto = map[string][]string{
	"001": {"03", "04"}, // ISR: pago definitivo o provisional
	"002": {"01"},       // IVA: pago definitivo
	"003": {"02"},       // IEPS: pago definitivo
}

const rutaRet = "/Retenciones"

// ValidarReglasRetenciones evalúa las reglas aritméticas y de coherencia de
// un CFDI de Retenciones e información de pagos que se pueden comprobar sin
// red ni catálogos externos (ver ReglasRetenciones): los totales contra
// MontoRet, los periodos, CveRetenc contra el complemento y los atributos
// que la matriz de errores de Retenciones 2.0 (Reten20xxx) amarra entre sí.
// Cada regla cita su fundamento en Regla.Citas.
//
// Un atributo ausente o ilegible no se evalúa: lo reporta ValidarEsquema.
// En 1.0 solo corre REGLA_RET10_MONTO_TOT_RET, porque el SAT no publica
// matriz de errores de esa versión. Una Version que no sea 1.0 ni 2.0 da un
// aviso REGLA_VERSION_NO_SOPORTADA. Nunca falla.
func ValidarReglasRetenciones(rt *Retenciones) []Hallazgo {
	v := &validadorRet{rt: rt, hs: []Hallazgo{}}
	switch rt.Version {
	case "1.0":
		v.montoTotRet10()
	case "2.0":
		v.cve()
		v.totales()
		v.periodo()
		v.receptor()
	default:
		v.hs = append(v.hs, Hallazgo{Codigo: ReglaVersionNoSoportada, Nivel: NivelAviso, Ruta: rutaRet, Regla: "",
			Mensaje: fmt.Sprintf("Retenciones %q: la librería solo tiene reglas para 1.0 y 2.0", rt.Version)})
	}
	return v.hs
}

type validadorRet struct {
	rt *Retenciones
	hs []Hallazgo
}

func (v *validadorRet) hallazgo(codigo, ruta, attr, valor, msg string) {
	r := reglaDe[codigo]
	h := Hallazgo{Codigo: codigo, Nivel: r.Nivel, Ruta: ruta, Atributo: attr, Valor: valor, Regla: r.Referencia, Mensaje: msg}
	if attr != "" {
		h.Ruta += "/@" + attr
	}
	v.hs = append(v.hs, h)
}

func (v *validadorRet) hallazgoDec(codigo, ruta, attr string, valor *decimal.Decimal, msg string) {
	var texto string
	if valor != nil {
		texto = valor.Original()
	}
	v.hallazgo(codigo, ruta, attr, texto, msg)
}

// sumaMontoRet suma los MontoRet de los ImpRetenidos; ok es false si
// alguno no se pudo leer.
func sumaMontoRet(t *TotalesRetenciones) (decimal.Decimal, bool) {
	var suma decimal.Decimal
	for _, im := range t.ImpRetenidos {
		if im.MontoRet == nil {
			return suma, false
		}
		suma = suma.Add(*im.MontoRet)
	}
	return suma, true
}

// montoTotRet10 revisa montoTotRet de 1.0 contra la suma de los montoRet.
func (v *validadorRet) montoTotRet10() {
	t := v.rt.Totales
	if t == nil || t.MontoTotRet == nil {
		return
	}
	if suma, ok := sumaMontoRet(t); ok && !t.MontoTotRet.Equal(suma) {
		v.hallazgoDec(ReglaRet10MontoTotRet, rutaRet+"/Totales", nombresRet10.montoTotRet, t.MontoTotRet,
			fmt.Sprintf("montoTotRet %s no es la suma de los montoRet de ImpRetenidos (%s)", t.MontoTotRet.Original(), suma.String()))
	}
}

// cve revisa CveRetenc contra DescRetenc, el complemento y la clave 28.
func (v *validadorRet) cve() {
	rt := v.rt
	cve := clave(rt.CveRetenc)
	if cve == "25" && strings.TrimSpace(rt.DescRetenc) == "" {
		v.hallazgo(ReglaRetDescRetenc, rutaRet, "DescRetenc", rt.DescRetenc, "con CveRetenc 25 (otro tipo de retenciones) debe venir DescRetenc")
	}
	if c, ok := complementosPorCve[cve]; ok && !c.presente(rt) && !tieneNoSoportado(rt, c.nombre) {
		v.hallazgo(ReglaRetCveComplemento, rutaRet, "CveRetenc", rt.CveRetenc,
			fmt.Sprintf("CveRetenc %s lleva el complemento %s («%s» en c_CveRetenc) y el comprobante no lo trae", cve, c.nombre, c.nombreCat))
	}
	if cve == "28" && rt.Totales != nil {
		t, rtot := rt.Totales, rutaRet+"/Totales"
		for _, x := range []struct {
			attr string
			val  *decimal.Decimal
		}{{"MontoTotGrav", t.MontoTotGrav}, {"MontoTotExent", t.MontoTotExent}} {
			if x.val != nil && !x.val.IsZero() {
				v.hallazgoDec(ReglaRetCveRIF, rtot, x.attr, x.val, "con CveRetenc 28 "+x.attr+" debe ser cero")
			}
		}
		for _, x := range []struct {
			attr string
			val  *decimal.Decimal
		}{{"UtilidadBimestral", t.UtilidadBimestral}, {"ISRCorrespondiente", t.ISRCorrespondiente}} {
			switch {
			case x.val == nil:
				v.hallazgo(ReglaRetCveRIF, rtot, x.attr, "", "con CveRetenc 28 debe venir "+x.attr+" mayor que cero")
			case x.val.Sign() <= 0:
				v.hallazgoDec(ReglaRetCveRIF, rtot, x.attr, x.val, "con CveRetenc 28 "+x.attr+" debe ser mayor que cero")
			}
		}
	}
}

// tieneNoSoportado dice si entre los complementos que la lectura no modela
// hay uno con ese nombre local (p. ej. una versión nueva del complemento).
func tieneNoSoportado(rt *Retenciones, nombre string) bool {
	for _, n := range rt.ComplementosNoSoportados {
		if i := strings.IndexByte(n, '}'); i >= 0 && strings.EqualFold(n[i+1:], nombre) {
			return true
		}
	}
	return false
}

// totales revisa MontoTotOperacion, MontoTotGrav, MontoTotExent, MontoTotRet
// y el TipoPagoRet de cada ImpRetenidos.
func (v *validadorRet) totales() {
	t := v.rt.Totales
	if t == nil {
		return
	}
	const rtot = rutaRet + "/Totales"
	op, grav, exent := t.MontoTotOperacion, t.MontoTotGrav, t.MontoTotExent
	if op != nil && grav != nil && exent != nil {
		if suma := grav.Add(*exent); !op.Equal(suma) {
			v.hallazgoDec(ReglaRetMontoOperacion, rtot, "MontoTotOperacion", op,
				fmt.Sprintf("MontoTotOperacion %s no es MontoTotGrav + MontoTotExent (%s)", op.Original(), suma.String()))
		}
	}
	if op != nil {
		for _, x := range []struct {
			attr string
			val  *decimal.Decimal
		}{{"MontoTotGrav", grav}, {"MontoTotExent", exent}} {
			if x.val != nil && x.val.Cmp(*op) > 0 {
				v.hallazgoDec(ReglaRetMontoMayor, rtot, x.attr, x.val,
					fmt.Sprintf("%s %s es mayor que MontoTotOperacion %s", x.attr, x.val.Original(), op.Original()))
			}
		}
	}
	if tot := t.MontoTotRet; tot != nil && tot.Sign() > 0 {
		if len(t.ImpRetenidos) == 0 {
			v.hallazgoDec(ReglaRetMontoTotRet, rtot, "MontoTotRet", tot, "MontoTotRet es mayor que cero y no hay ningún ImpRetenidos")
		} else if suma, ok := sumaMontoRet(t); ok && !tot.Equal(suma) {
			v.hallazgoDec(ReglaRetMontoTotRet, rtot, "MontoTotRet", tot,
				fmt.Sprintf("MontoTotRet %s no es la suma de los MontoRet de ImpRetenidos (%s)", tot.Original(), suma.String()))
		}
	}
	for i, im := range t.ImpRetenidos {
		permitidos, ok := tipoPagoPorImpuesto[clave(im.ImpuestoRet)]
		tp := clave(im.TipoPagoRet)
		if !ok || !esTipoPagoRet(tp) {
			continue // clave fuera del catálogo: lo reporta el esquema
		}
		if !contiene(permitidos, tp) {
			v.hallazgo(ReglaRetTipoPago, idx(rtot, "ImpRetenidos", i), "TipoPagoRet", im.TipoPagoRet,
				fmt.Sprintf("TipoPagoRet %s no corresponde a ImpuestoRet %s (admite %s)", tp, clave(im.ImpuestoRet), strings.Join(permitidos, " o ")))
		}
	}
}

func esTipoPagoRet(s string) bool { return s == "01" || s == "02" || s == "03" || s == "04" }

func contiene(l []string, s string) bool {
	for _, x := range l {
		if x == s {
			return true
		}
	}
	return false
}

// periodo revisa MesIni contra MesFin y Ejercicio contra FechaExp.
func (v *validadorRet) periodo() {
	p := v.rt.Periodo
	if p == nil {
		return
	}
	const rp = rutaRet + "/Periodo"
	ini, err1 := strconv.Atoi(clave(p.MesIni))
	fin, err2 := strconv.Atoi(clave(p.MesFin))
	if err1 == nil && err2 == nil && ini > fin {
		v.hallazgo(ReglaRetPeriodo, rp, "MesIni", p.MesIni, fmt.Sprintf("MesIni %s es posterior a MesFin %s", clave(p.MesIni), clave(p.MesFin)))
	}
	ej, err := strconv.Atoi(clave(p.Ejercicio))
	if err != nil || !v.rt.FechaExp.Valida() {
		return
	}
	anio := v.rt.FechaExp.Time.Year()
	switch {
	case ej > anio:
		v.hallazgo(ReglaRetEjercicio, rp, "Ejercicio", p.Ejercicio, fmt.Sprintf("Ejercicio %d es posterior al año de FechaExp (%d)", ej, anio))
	case ej < anio-5:
		v.hallazgo(ReglaRetEjercicio, rp, "Ejercicio", p.Ejercicio, fmt.Sprintf("Ejercicio %d es anterior en más de 5 años al de FechaExp (%d)", ej, anio))
	}
}

// receptor revisa el Receptor/Nacional.
func (v *validadorRet) receptor() {
	rec := v.rt.Receptor
	if rec == nil || rec.Nacional == nil {
		return
	}
	n := rec.Nacional
	const rn = rutaRet + "/Receptor/Nacional"
	rfcR := strings.ToUpper(clave(n.RfcR))
	if n.CurpR != "" && len(rfcR) == 12 {
		v.hallazgo(ReglaRetReceptor, rn, "CurpR", n.CurpR, "RfcR tiene 12 posiciones (persona moral): no va CurpR")
	}
	generico := rfcR == "XAXX010101000"
	if esPublicoEnGeneral(n.NomDenRazSocR) && rfcR != "" && !generico {
		v.hallazgo(ReglaRetReceptor, rn, "RfcR", n.RfcR, "NomDenRazSocR es «PÚBLICO EN GENERAL» y RfcR no es XAXX010101000")
	}
	if generico && clave(n.DomicilioFiscalR) != "" && clave(v.rt.LugarExpRetenc) != "" && clave(n.DomicilioFiscalR) != clave(v.rt.LugarExpRetenc) {
		v.hallazgo(ReglaRetReceptor, rn, "DomicilioFiscalR", n.DomicilioFiscalR,
			fmt.Sprintf("con RfcR XAXX010101000, DomicilioFiscalR debe ser igual a LugarExpRetenc (%s)", clave(v.rt.LugarExpRetenc)))
	}
}

// esPublicoEnGeneral compara con «PÚBLICO EN GENERAL» sin distinguir
// mayúsculas ni el acento.
func esPublicoEnGeneral(s string) bool {
	return strings.NewReplacer("Ú", "U").Replace(strings.ToUpper(strings.Join(strings.Fields(s), " "))) == "PUBLICO EN GENERAL"
}
