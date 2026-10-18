package cfdi

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/OrtaMarco/cfdi-go/decimal"
)

// Códigos de hallazgo de ValidarReglas. Son parte de la API pública: no se
// renombran. La referencia de cada uno (de dónde sale la regla) está en
// Reglas.
const (
	ReglaVersionNoSoportada = "REGLA_VERSION_NO_SOPORTADA" // aviso: versión sin reglas (3.2, 3.0…)

	ReglaImporteConcepto       = "REGLA_IMPORTE_CONCEPTO"
	ReglaDescuentoConcepto     = "REGLA_DESCUENTO_CONCEPTO"
	ReglaImporteImpuesto       = "REGLA_IMPORTE_IMPUESTO"
	ReglaExentoConImporte      = "REGLA_EXENTO_CON_IMPORTE"
	ReglaTasaSinImporte        = "REGLA_TASA_SIN_IMPORTE"
	ReglaRetencionExenta       = "REGLA_RETENCION_EXENTA"
	ReglaObjetoImp             = "REGLA_OBJETO_IMP"
	ReglaSubTotal              = "REGLA_SUBTOTAL"
	ReglaDescuento             = "REGLA_DESCUENTO"
	ReglaDescuentoMayor        = "REGLA_DESCUENTO_MAYOR_SUBTOTAL"
	ReglaTotal                 = "REGLA_TOTAL"
	ReglaTotalTraslados        = "REGLA_TOTAL_TRASLADOS"
	ReglaTotalRetenciones      = "REGLA_TOTAL_RETENCIONES"
	ReglaImpuestosSinTotal     = "REGLA_IMPUESTOS_SIN_TOTAL"
	ReglaTrasladoSuma          = "REGLA_TRASLADO_SUMA"
	ReglaTrasladoFaltante      = "REGLA_TRASLADO_FALTANTE"
	ReglaTrasladoSobrante      = "REGLA_TRASLADO_SOBRANTE"
	ReglaRetencionSuma         = "REGLA_RETENCION_SUMA"
	ReglaRetencionFaltante     = "REGLA_RETENCION_FALTANTE"
	ReglaRetencionSobrante     = "REGLA_RETENCION_SOBRANTE"
	ReglaImpLocalTotales       = "REGLA_IMPLOCAL_TOTALES"
	ReglaDecimalesMoneda       = "REGLA_DECIMALES_MONEDA"
	ReglaDecimalesImpuesto     = "REGLA_DECIMALES_IMPUESTO"
	ReglaDecimalesImpuesto33   = "REGLA_DECIMALES_IMPUESTO_33"
	ReglaTipoCambioMXN         = "REGLA_TIPOCAMBIO_MXN"
	ReglaTipoCambioXXX         = "REGLA_TIPOCAMBIO_XXX"
	ReglaTipoCambioRequerido   = "REGLA_TIPOCAMBIO_REQUERIDO"
	ReglaTipoCompCondiciones   = "REGLA_TIPOCOMP_CONDICIONES"
	ReglaTipoCompImpuestos     = "REGLA_TIPOCOMP_IMPUESTOS"
	ReglaTipoCompFormaPago     = "REGLA_TIPOCOMP_FORMAPAGO"
	ReglaTipoCompMetodoPago    = "REGLA_TIPOCOMP_METODOPAGO"
	ReglaTipoCompDescuento     = "REGLA_TIPOCOMP_DESCUENTO"
	ReglaTipoCompSubTotalCero  = "REGLA_TIPOCOMP_SUBTOTAL_CERO"
	ReglaTipoCompTotalCero     = "REGLA_TIPOCOMP_TOTAL_CERO"
	ReglaTipoCompValorUnitario = "REGLA_TIPOCOMP_VALOR_UNITARIO"
	ReglaTipoCompMonedaNomina  = "REGLA_TIPOCOMP_MONEDA_NOMINA"
	ReglaPagoComprobante       = "REGLA_PAGO_COMPROBANTE"
	ReglaPagoSaldoInsoluto     = "REGLA_PAGO_SALDO_INSOLUTO"
	ReglaPagoEquivalencia      = "REGLA_PAGO_EQUIVALENCIA"
	ReglaPagoTipoCambioDR      = "REGLA_PAGO_TIPOCAMBIODR"
	ReglaPagoMonto             = "REGLA_PAGO_MONTO"
	ReglaPagoTipoCambio        = "REGLA_PAGO_TIPOCAMBIO"
	ReglaPagosTotales          = "REGLA_PAGOS_TOTALES"
	ReglaPagoMonedaXXX         = "REGLA_PAGO_MONEDA_XXX"
	ReglaPagoMayorCero         = "REGLA_PAGO_MAYOR_CERO"
	ReglaPagoDecimales         = "REGLA_PAGO_DECIMALES"
	ReglaPagoFormaPago         = "REGLA_PAGO_FORMAPAGO"
	ReglaPagoCuentas           = "REGLA_PAGO_CUENTAS"
	ReglaPagoCadena            = "REGLA_PAGO_CADENA"
	ReglaPago10Parcialidades   = "REGLA_PAGO10_PARCIALIDADES"
	ReglaPago10Impuestos       = "REGLA_PAGO10_IMPUESTOS"
)

// Regla describe una regla de ValidarReglas: su código, el nivel de sus
// hallazgos, de dónde sale (la referencia que va en Hallazgo.Regla) y qué
// comprueba.
type Regla struct {
	Codigo string `json:"codigo"`
	// Nivel es NivelError si la regla tiene fundamento en un documento
	// oficial del SAT (Citas: matriz de errores, estándar o guía de
	// llenado); NivelAdvertencia si no lo tiene o si el SAT pide no
	// aplicarla.
	Nivel string `json:"nivel"`
	// Referencia es el texto de Citas más las fuentes secundarias
	// (CfdiUtils, satcfdi) cuando ayudan a entender la regla.
	Referencia  string    `json:"referencia"`
	Descripcion string    `json:"descripcion"`
	Citas       []CitaSAT `json:"citas,omitempty"`
	// Margen, si no está vacío, hace de la regla una de dos niveles: lo
	// que no cumple la letra del SAT (Descripcion) pero cae dentro de
	// Margen sale con NivelAdvertencia, con el mismo código y la misma
	// referencia; lo que queda fuera también del margen, con Nivel. Lo
	// llevan el Importe del impuesto del concepto y las sumas de impuestos:
	// los PAC timbran documentos con un centavo de diferencia por el
	// redondeo de cada concepto, y marcarlos como error engañaría a quien
	// usa la librería.
	Margen string `json:"margen,omitempty"`
	// Vigencia, si no está vacía, dice qué partes de la regla son error
	// solo desde una fecha: en un comprobante con Fecha anterior, el
	// hallazgo de esas partes sale con NivelAdvertencia y el mensaje lo
	// dice. Lo llevan las reglas de Nómina que el SAT agregó después de la
	// primera versión del estándar.
	Vigencia string `json:"vigencia,omitempty"`
	// SoloGuia, si no está vacío, dice qué partes de la regla solo salen de
	// la guía de llenado, sin código en la matriz de errores ni texto en el
	// estándar: sus hallazgos salen con NivelAdvertencia.
	SoloGuia string `json:"solo_guia,omitempty"`
}

// CitaSAT es el fundamento de una regla en un documento oficial del SAT: un
// código de su matriz de errores o una página de un estándar o de una guía
// de llenado (la página impresa del PDF). La lista de documentos, con su
// URL oficial, está en el README.
type CitaSAT struct {
	Documento string `json:"documento"`
	Codigo    string `json:"codigo,omitempty"`
	Pagina    int    `json:"pagina,omitempty"`
}

func (c CitaSAT) String() string {
	if c.Codigo != "" {
		return c.Documento + " " + c.Codigo
	}
	return c.Documento + " p. " + strconv.Itoa(c.Pagina)
}

// Documentos oficiales del SAT que citan las reglas.
const (
	DocMatriz40      = "Matriz de errores CFDI 4.0 (2026-03-25)"
	DocEstandar40    = "Anexo 20 CFDI 4.0 (DOF 2022-01-13)"
	DocGuia40        = "Guía de llenado CFDI 4.0 (2026-01-16)"
	DocMatriz33      = "Matriz de errores CFDI 3.3"
	DocEstandar33    = "Anexo 20 CFDI 3.3 (DOF 2017-07-28)"
	DocGuia33        = "Guía de llenado CFDI 3.3 (2019-10-14)"
	DocMatrizPagos20 = "Matriz de errores Pagos 2.0 Rev B"
	DocGuiaPagos20   = "Guía de llenado Pagos 2.0 (2026-01-16)"
	DocEstPagos20    = "Estándar Pagos 2.0 Rev B"
	DocMatrizPagos10 = "Matriz de errores Pagos 1.0"
	DocGuiaPagos10   = "Guía de llenado Pagos 1.0 (2018-08-31)"
	DocImpLocal      = "Estándar ImpuestosLocales"
)

func cod(doc string, codigos ...string) []CitaSAT {
	out := make([]CitaSAT, len(codigos))
	for i, c := range codigos {
		out[i] = CitaSAT{Documento: doc, Codigo: c}
	}
	return out
}

func pag(doc string, paginas ...int) []CitaSAT {
	out := make([]CitaSAT, len(paginas))
	for i, p := range paginas {
		out[i] = CitaSAT{Documento: doc, Pagina: p}
	}
	return out
}

func citas(grupos ...[]CitaSAT) []CitaSAT {
	var out []CitaSAT
	for _, g := range grupos {
		out = append(out, g...)
	}
	return out
}

// regla arma una Regla: la Referencia es el texto de las citas y, si hay,
// la nota.
func regla(codigo, nivel, descripcion string, cs []CitaSAT, nota string) Regla {
	partes := make([]string, 0, len(cs)+1)
	for _, c := range cs {
		partes = append(partes, c.String())
	}
	if nota != "" {
		partes = append(partes, nota)
	}
	return Regla{Codigo: codigo, Nivel: nivel, Referencia: strings.Join(partes, "; "), Descripcion: descripcion, Citas: cs}
}

// conMargen le da a r su segundo nivel (ver Regla.Margen).
func conMargen(r Regla, margen string) Regla {
	r.Margen = margen
	return r
}

// conVigencia documenta las partes de r que son error solo desde una fecha
// (ver Regla.Vigencia).
func conVigencia(r Regla, vigencia string) Regla {
	r.Vigencia = vigencia
	return r
}

// conSoloGuia documenta las partes de r que solo salen de la guía de
// llenado (ver Regla.SoloGuia).
func conSoloGuia(r Regla, partes string) Regla {
	r.SoloGuia = partes
	return r
}

// vigencia es la fecha desde la que una parte de una regla es error (ver
// Regla.Vigencia): desde es el primer día, en hora civil como Fecha.Time, y
// porque lo que dice el mensaje de un comprobante anterior.
type vigencia struct {
	desde  time.Time
	porque string
}

// Márgenes de las reglas de dos niveles.
const (
	margenSumaConceptos = "dos niveles: si no es la suma ni su redondeo pero cae entre la suma de cada sumando truncado y la de cada uno redondeado hacia arriba a los decimales de la moneda (cada sumando, el Importe escrito en el concepto o Base × TasaOCuota sin redondear), advertencia: el margen de redondeo por concepto, que un PAC timbra"
	margenSumaNodos     = "dos niveles: si no es la suma ni su redondeo pero cae entre la suma de cada Importe truncado y la de cada uno redondeado hacia arriba a los decimales de la moneda (en 3.3 pueden traer 6 decimales), advertencia: el margen de redondeo por sumando, que un PAC timbra"
)

// Reglas lista las reglas que evalúa ValidarReglas, en el orden de los
// códigos. Cada una con su fundamento en los documentos del SAT (T7: se
// contrastaron todas contra la matriz de errores, el estándar y la guía de
// llenado de cada versión; lo que no coincidía se corrigió).
var Reglas = concatenar(reglasCFDI, reglasPagos, reglasPagosImpuestos, reglasComplementos, reglasCCE, reglasNomina)

func concatenar(listas ...[]Regla) []Regla {
	var out []Regla
	for _, l := range listas {
		out = append(out, l...)
	}
	return out
}

var reglasCFDI = []Regla{
	{Codigo: ReglaVersionNoSoportada, Nivel: NivelAviso, Descripcion: "aviso: la versión del CFDI no tiene reglas en la librería (solo 3.3 y 4.0)"},
	regla(ReglaImporteConcepto, NivelError,
		"Importe del concepto entre (Cantidad − ½·10^-dC)·(ValorUnitario − ½·10^-dV) truncado y (Cantidad + ½·10^-dC − 10^-12)·(ValorUnitario + ½·10^-dV − 10^-12) redondeado hacia arriba, a los decimales del propio Importe",
		citas(cod(DocMatriz40, "CFDI40167"), pag(DocEstandar40, 76), cod(DocMatriz33, "CFDI33149"), pag(DocEstandar33, 68)), ""),
	regla(ReglaDescuentoConcepto, NivelError, "Descuento del concepto menor o igual que su Importe",
		citas(cod(DocMatriz40, "CFDI40169"), pag(DocGuia40, 24), cod(DocMatriz33, "CFDI33151")), ""),
	conMargen(regla(ReglaImporteImpuesto, NivelError,
		"Importe del impuesto del concepto entre (Base − ½·10^-dB)·TasaOCuota truncado y (Base + ½·10^-dB − 10^-12)·TasaOCuota redondeado hacia arriba, a los decimales del propio Importe",
		citas(cod(DocMatriz40, "CFDI40180", "CFDI40186"), pag(DocEstandar40, 77, 78), cod(DocMatriz33, "CFDI33161", "CFDI33169"), pag(DocEstandar33, 69, 70)), ""),
		"dos niveles: si no cumple pero cae entre los mismos límites truncado y redondeado hacia arriba a los decimales de la moneda (un Importe de 6 decimales en MXN: 52.620000 con Base 328.90 × 0.160000), advertencia: el margen de redondeo por concepto, que un PAC timbra"),
	regla(ReglaExentoConImporte, NivelError, "traslado Exento sin TasaOCuota ni Importe",
		citas(cod(DocMatriz40, "CFDI40177"), pag(DocEstandar40, 77), cod(DocMatriz33, "CFDI33157")), ""),
	regla(ReglaTasaSinImporte, NivelError, "traslado Tasa o Cuota con TasaOCuota e Importe",
		citas(cod(DocMatriz40, "CFDI40178"), pag(DocEstandar40, 77), cod(DocMatriz33, "CFDI33158")), ""),
	regla(ReglaRetencionExenta, NivelError, "retención con TipoFactor distinto de Exento",
		citas(cod(DocMatriz40, "CFDI40184"), pag(DocEstandar40, 78), cod(DocMatriz33, "CFDI33166")), ""),
	regla(ReglaObjetoImp, NivelError,
		"4.0: ObjetoImp 02 con nodo Impuestos en el concepto; 01, 03, 04 y 05 sin él; 06 y 08 sin traslados ni retenciones de IVA (002) o IEPS (003); 07 sin IVA y con traslado de IEPS",
		citas(cod(DocMatriz40, "CFDI40171", "CFDI40172", "CFDI140226", "CFDI140227"), pag(DocEstandar40, 76)), ""),
	regla(ReglaSubTotal, NivelError, "I, E y N: SubTotal igual a la suma de los Importe de los conceptos o a su redondeo a los decimales de la moneda",
		citas(cod(DocMatriz40, "CFDI40108"), pag(DocGuia40, 8), cod(DocMatriz33, "CFDI33107"), pag(DocEstandar33, 64)), ""),
	regla(ReglaDescuento, NivelError, "Descuento igual a la suma de los descuentos de los conceptos o a su redondeo; sin descuentos en los conceptos, sin Descuento",
		citas(cod(DocMatriz40, "CFDI40111"), pag(DocGuia40, 8), cod(DocMatriz33, "CFDI33110")), ""),
	regla(ReglaDescuentoMayor, NivelError, "Descuento menor o igual que SubTotal",
		citas(cod(DocMatriz40, "CFDI40110"), pag(DocGuia40, 8), cod(DocMatriz33, "CFDI33109")), ""),
	regla(ReglaTotal, NivelError, "Total = SubTotal − Descuento + TotalImpuestosTrasladados − TotalImpuestosRetenidos + traslados locales − retenciones locales (implocal), redondeado a la moneda",
		citas(cod(DocMatriz40, "CFDI40119"), pag(DocEstandar40, 72), pag(DocGuia40, 10), cod(DocMatriz33, "CFDI33118")),
		"«impuestos trasladados - federales y/o locales»; sin tolerancia en ningún documento"),
	conMargen(regla(ReglaTotalTraslados, NivelError, "TotalImpuestosTrasladados igual a la suma de los Importe de Impuestos/Traslados o a su redondeo a la moneda (en 3.3 pueden traer 6 decimales)",
		citas(cod(DocMatriz40, "CFDI40205"), pag(DocEstandar40, 81), cod(DocMatriz33, "CFDI33183")), ""), margenSumaNodos),
	conMargen(regla(ReglaTotalRetenciones, NivelError, "TotalImpuestosRetenidos igual a la suma de los Importe de Impuestos/Retenciones o a su redondeo a la moneda",
		citas(cod(DocMatriz40, "CFDI40203"), pag(DocEstandar40, 81), cod(DocMatriz33, "CFDI33181")), ""), margenSumaNodos),
	regla(ReglaImpuestosSinTotal, NivelError, "con traslados (no exentos) debe venir TotalImpuestosTrasladados; con retenciones, TotalImpuestosRetenidos",
		citas(cod(DocMatriz40, "CFDI40212", "CFDI40206", "CFDI40209"), cod(DocMatriz33, "CFDI33190", "CFDI33184", "CFDI33187")), ""),
	conMargen(regla(ReglaTrasladoSuma, NivelError,
		"cada Impuestos/Traslados/Traslado (Impuesto, TipoFactor, TasaOCuota) con Importe (y Base en 4.0) igual a la suma de los de los conceptos o a su redondeo a la moneda: se redondea al final, no cada concepto",
		citas(cod(DocMatriz40, "CFDI40221", "CFDI40215", "CFDI40216"), pag(DocGuia40, 90), cod(DocMatriz33, "CFDI33195"), pag(DocGuia33, 80)),
		"pregunta 39 de la guía 4.0 y 44 de la 3.3: «Redondear al final del cálculo y no antes»"), margenSumaConceptos),
	regla(ReglaTrasladoFaltante, NivelError, "cada grupo de traslados de los conceptos está en Impuestos/Traslados; los Exento solo en 4.0 (en 3.3 el nodo del comprobante no los lleva)",
		citas(cod(DocMatriz40, "CFDI40221", "CFDI40213"), pag(DocGuia40, 90), cod(DocMatriz33, "CFDI33195"), pag(DocGuia33, 32)), ""),
	regla(ReglaTrasladoSobrante, NivelError, "Impuestos/Traslados no tiene grupos que no vengan de los conceptos",
		citas(cod(DocMatriz40, "CFDI40221"), cod(DocMatriz33, "CFDI33195")), ""),
	conMargen(regla(ReglaRetencionSuma, NivelError, "cada Impuestos/Retenciones/Retencion (Impuesto) con Importe igual a la suma de los de los conceptos o a su redondeo a la moneda",
		citas(cod(DocMatriz40, "CFDI40211"), cod(DocMatriz33, "CFDI33189")), ""), margenSumaConceptos),
	regla(ReglaRetencionFaltante, NivelError, "cada impuesto retenido en los conceptos está en Impuestos/Retenciones",
		citas(cod(DocMatriz40, "CFDI40211"), cod(DocMatriz33, "CFDI33189")), ""),
	regla(ReglaRetencionSobrante, NivelError, "Impuestos/Retenciones no tiene impuestos que no vengan de los conceptos",
		citas(cod(DocMatriz40, "CFDI40211"), cod(DocMatriz33, "CFDI33189")), ""),
	regla(ReglaImpLocalTotales, NivelError, "TotaldeTraslados y TotaldeRetenciones iguales a la suma de los Importe locales",
		pag(DocImpLocal, 2), "«la suma total de traslados (retenciones) aplicables»; el SAT no publica matriz de errores de implocal"),
	regla(ReglaDecimalesMoneda, NivelError,
		"SubTotal, Descuento, Total y totales de impuestos con no más decimales de los que admite la moneda (XXX admite 0: un Total 0.00 en XXX no pasa)",
		citas(cod(DocMatriz40, "CFDI40107", "CFDI40112", "CFDI40202", "CFDI40204"), pag(DocEstandar40, 72), cod(DocMatriz33, "CFDI33106", "CFDI33111", "CFDI33180", "CFDI33182")), ""),
	regla(ReglaDecimalesImpuesto, NivelError, "4.0: Base e Importe de Impuestos/Traslados/Traslado e Importe de Retenciones/Retencion con no más decimales de los que admite la moneda",
		citas(cod(DocMatriz40, "CFDI40210", "CFDI40214", "CFDI40220"), pag(DocEstandar40, 81, 82), pag(DocGuia40, 38, 40)), ""),
	regla(ReglaDecimalesImpuesto33, NivelAdvertencia, "3.3: Importe de Impuestos/Traslados/Traslado y Retenciones/Retencion con no más decimales de los que admite la moneda",
		citas(cod(DocMatriz33, "CFDI33188", "CFDI33194"), pag(DocGuia33, 32, 34)),
		"advertencia: la matriz 3.3 dice «NO DEBEN CONSIDERAR EN SU APLICATIVO LOS ERRORES» CFDI33188 y CFDI33194"),
	regla(ReglaTipoCambioMXN, NivelError, "con Moneda MXN, TipoCambio ausente o 1",
		citas(cod(DocMatriz40, "CFDI40114"), cod(DocMatriz33, "CFDI33113")), ""),
	regla(ReglaTipoCambioXXX, NivelError, "con Moneda XXX, sin TipoCambio",
		citas(cod(DocMatriz40, "CFDI40116"), cod(DocMatriz33, "CFDI33115")), ""),
	regla(ReglaTipoCambioRequerido, NivelError, "con otra moneda, TipoCambio requerido",
		citas(cod(DocMatriz40, "CFDI40115"), cod(DocMatriz33, "CFDI33114")), ""),
	regla(ReglaTipoCompCondiciones, NivelError, "T, P y N sin CondicionesDePago",
		citas(pag(DocEstandar40, 73), pag(DocGuia40, 11), pag(DocEstandar33, 66)), ""),
	regla(ReglaTipoCompImpuestos, NivelError, "T, P y N sin nodo Impuestos",
		citas(cod(DocMatriz40, "CFDI40201"), pag(DocEstandar40, 73), cod(DocMatriz33, "CFDI33179"), pag(DocEstandar33, 66)),
		"la matriz dice T o P; el estándar y la guía agregan N"),
	regla(ReglaTipoCompFormaPago, NivelError, "T y P sin FormaPago; en 4.0 también N",
		citas(cod(DocMatriz40, "CFDI40103"), pag(DocEstandar40, 71), pag(DocEstandar33, 66)), ""),
	regla(ReglaTipoCompMetodoPago, NivelError, "T y P sin MetodoPago",
		citas(cod(DocMatriz40, "CFDI40125"), cod(DocMatriz33, "CFDI33123")), ""),
	regla(ReglaTipoCompDescuento, NivelError, "T y P sin Descuento, ni en el comprobante ni en los conceptos",
		citas(cod(DocMatriz40, "CFDI40111"), pag(DocEstandar40, 73), cod(DocMatriz33, "CFDI33110")), ""),
	regla(ReglaTipoCompSubTotalCero, NivelError, "T y P con SubTotal 0",
		citas(cod(DocMatriz40, "CFDI40109"), cod(DocMatriz33, "CFDI33108")), ""),
	regla(ReglaTipoCompTotalCero, NivelError, "T y P con Total 0",
		citas(pag(DocEstandar40, 72), pag(DocGuia40, 10)), ""),
	regla(ReglaTipoCompValorUnitario, NivelError, "I, E y N con ValorUnitario mayor que cero en cada concepto",
		citas(cod(DocMatriz40, "CFDI40166"), pag(DocGuia40, 23), cod(DocMatriz33, "CFDI33147")), ""),
	regla(ReglaTipoCompMonedaNomina, NivelError, "N con Moneda MXN",
		citas(pag(DocEstandar40, 71), pag(DocEstandar33, 65)), ""),
}

var reglaDe = func() map[string]Regla {
	m := map[string]Regla{}
	for _, r := range Reglas {
		m[r.Codigo] = r
	}
	return m
}()

// CodigosReglas lista los códigos de ValidarReglas, en el orden de Reglas.
var CodigosReglas = func() []string {
	out := make([]string, len(Reglas))
	for i, r := range Reglas {
		out[i] = r.Codigo
	}
	return out
}()

// ValidarReglas evalúa las reglas aritméticas y de coherencia que se pueden
// comprobar sin red ni catálogos externos (ver Reglas): importes de los
// conceptos y de sus impuestos con los límites del Anexo 20, SubTotal,
// Descuento, Total, totales de impuestos agrupados por impuesto, tipo de
// factor y tasa, decimales y tipo de cambio según la moneda, reglas por
// TipoDeComprobante, ImpuestosLocales, Pagos 1.0 y 2.0, Comercio Exterior
// 2.0 y Nómina 1.2. Todo con decimal, sin float.
//
// Es aparte de ValidarEsquema y trabaja sobre el Comprobante, no sobre los
// bytes: las reglas son de valores (montos ya leídos como decimales) y no
// de la forma del documento; y quien ya hizo Parse no relee el XML. Usan el
// mismo tipo Hallazgo con el mismo contrato de Ruta, así que se pueden
// juntar en una lista.
//
// Cada hallazgo lleva el nivel de su regla (Regla.Nivel): error si la
// regla tiene fundamento en un documento oficial del SAT, advertencia si
// no; las advertencias no hacen inválido el documento (EstadoDe). Las
// reglas con Regla.Margen tienen dos niveles: lo que no cumple la letra
// del SAT pero cae en el margen de redondeo por concepto sale como
// advertencia.
//
// Una regla que necesita un valor ausente o ilegible no se evalúa: de eso
// avisan la lectura (Advertencias) y ValidarEsquema. CFDI 3.3 y 4.0; otra
// versión da un solo aviso ReglaVersionNoSoportada. Las reglas propias de
// 4.0 (ObjetoImp, Base de los traslados del comprobante, Pagos 2.0,
// Comercio Exterior 2.0) solo se evalúan en 4.0; las de Pagos 1.0, solo en
// 3.3.
//
// Nunca devuelve nil: sin hallazgos es una lista vacía.
func ValidarReglas(c *Comprobante) []Hallazgo {
	v := &validadorReglas{c: c, hs: []Hallazgo{}}
	if c.Version != "3.3" && c.Version != "4.0" {
		v.hs = append(v.hs, Hallazgo{Codigo: ReglaVersionNoSoportada, Nivel: NivelAviso, Ruta: "/Comprobante", Regla: "",
			Mensaje: fmt.Sprintf("CFDI %q: la librería solo tiene reglas para 3.3 y 4.0", c.Version)})
		return v.hs
	}
	v.v40 = c.Version == "4.0"
	v.moneda, v.tipo = clave(c.Moneda), clave(c.TipoDeComprobante)
	v.dec, v.monedaOK = decimalesMoneda[v.moneda]
	v.conceptos()
	v.sumas()
	v.decimales()
	v.tipoCambio()
	v.tipoComprobante()
	v.impLocal()
	if p := c.Pagos; p != nil && ((v.v40 && p.Version == "2.0") || (!v.v40 && p.Version == "1.0")) {
		v.pagos()
	}
	v.comercioExterior()
	v.nomina()
	return v.hs
}

type validadorReglas struct {
	c        *Comprobante
	hs       []Hallazgo
	v40      bool
	dec      int32 // decimales de la moneda
	monedaOK bool  // la moneda está en c_Moneda
	moneda   string
	tipo     string // TipoDeComprobante
}

// clave normaliza una clave de catálogo como lo hace el XSD
// (whiteSpace collapse): " MXN " es MXN. El modelo la guarda tal cual.
func clave(s string) string { return strings.TrimSpace(s) }

func (v *validadorReglas) hallazgo(codigo, ruta, attr string, valor *decimal.Decimal, msg string) {
	r := reglaDe[codigo]
	h := Hallazgo{Codigo: codigo, Nivel: r.Nivel, Ruta: ruta, Atributo: attr, Regla: r.Referencia, Mensaje: msg}
	if attr != "" {
		h.Ruta += "/@" + attr
	}
	if valor != nil {
		h.Valor = valor.Original()
	}
	v.hs = append(v.hs, h)
}

// hallazgoDosNiveles registra el hallazgo de una regla con Regla.Margen
// que no cumple la letra del SAT: advertencia si enMargen, con su nivel si
// no. margen describe el intervalo del margen para el mensaje.
func (v *validadorReglas) hallazgoDosNiveles(codigo, ruta, attr string, valor *decimal.Decimal, msg string, enMargen bool, margen string) {
	if !enMargen {
		v.hallazgo(codigo, ruta, attr, valor, msg+"; fuera también del "+margen)
		return
	}
	v.hallazgo(codigo, ruta, attr, valor, msg+": no cumple la letra del SAT, pero sí el "+margen)
	v.hs[len(v.hs)-1].Nivel = NivelAdvertencia
}

// aAdvertencia baja el último hallazgo a advertencia y agrega por qué al
// mensaje.
func (v *validadorReglas) aAdvertencia(porque string) {
	h := &v.hs[len(v.hs)-1]
	h.Nivel = NivelAdvertencia
	h.Mensaje += "; advertencia: " + porque
}

// segunVigencia baja el último hallazgo a advertencia si la Fecha del
// comprobante es anterior a vg (Regla.Vigencia). Sin Fecha legible la regla
// se aplica: no hay cómo saber que es anterior.
func (v *validadorReglas) segunVigencia(vg vigencia) {
	if f := v.c.Fecha; f.Valida() && f.Time.Before(vg.desde) {
		v.aAdvertencia(fmt.Sprintf("el comprobante (%s) es anterior al %s; %s", f.Original, vg.desde.Format("2006-01-02"), vg.porque))
	}
}

// soloGuia baja el último hallazgo a advertencia: esa parte de la regla
// solo sale de la guía de llenado (Regla.SoloGuia). cita es la página.
func (v *validadorReglas) soloGuia(cita string) {
	v.aAdvertencia("solo lo pide la " + cita + ", sin código en la matriz de errores")
}

func (v *validadorReglas) hallazgoTexto(codigo, ruta, attr, valor, msg string) {
	v.hallazgo(codigo, ruta, attr, nil, msg)
	v.hs[len(v.hs)-1].Valor = valor
}

// decimalesPara son los decimales a los que se redondea una suma que se
// compara con valor: los de la moneda, o los del propio valor si la moneda
// no está en el catálogo.
func (v *validadorReglas) decimalesPara(valor decimal.Decimal) int32 {
	if v.monedaOK {
		return v.dec
	}
	return valor.Scale()
}

var cero = decimal.Decimal{}

func valorDe(d *decimal.Decimal) decimal.Decimal {
	if d == nil {
		return cero
	}
	return *d
}

// mitadUlp es ½·10^-n.
func mitadUlp(n int32) decimal.Decimal { return decimal.NewFromInt64(5, n+1) }

// casiCero es 10^-12, lo que el Anexo 20 resta al límite superior.
var casiCero = decimal.NewFromInt64(1, 12)

// truncar corta d a n decimales hacia cero; techo redondea hacia +∞.
func truncar(d decimal.Decimal, n int32) decimal.Decimal {
	r := d.Round(n)
	switch {
	case d.Sign() >= 0 && r.Cmp(d) > 0:
		return r.Sub(decimal.NewFromInt64(1, n))
	case d.Sign() < 0 && r.Cmp(d) < 0:
		return r.Add(decimal.NewFromInt64(1, n))
	}
	return r
}

func techo(d decimal.Decimal, n int32) decimal.Decimal {
	r := d.Round(n)
	if r.Cmp(d) < 0 {
		return r.Add(decimal.NewFromInt64(1, n))
	}
	return r
}

// limites calcula el intervalo del Anexo 20 (estándar 4.0 p. 76-78, 3.3 p.
// 68-70) para a×b, con a y b leídos con sus decimales: [(a − ½·10^-da)(b −
// ½·10^-db) truncado, (a + ½·10^-da − 10^-12)(b + ½·10^-db − 10^-12) hacia
// arriba], a dec decimales: «con la cantidad de decimales que tenga
// registrado este atributo», el Importe. Con exacto, b no varía (TasaOCuota
// es exacta).
func limites(a, b decimal.Decimal, exacto bool, dec int32) (inf, sup decimal.Decimal) {
	da, db := mitadUlp(a.Scale()), mitadUlp(b.Scale())
	bInf, bSup := b.Sub(db), b.Add(db).Sub(casiCero)
	if exacto {
		bInf, bSup = b, b
	}
	inf = truncar(a.Sub(da).Mul(bInf), dec)
	sup = techo(a.Add(da).Sub(casiCero).Mul(bSup), dec)
	return inf, sup
}

func (v *validadorReglas) conceptos() {
	for i, cp := range v.c.Conceptos {
		r := idx("/Comprobante/Conceptos", "Concepto", i)
		if cp.Cantidad != nil && cp.ValorUnitario != nil && cp.Importe != nil {
			inf, sup := limites(*cp.Cantidad, *cp.ValorUnitario, false, cp.Importe.Scale())
			if cp.Importe.Cmp(inf) < 0 || cp.Importe.Cmp(sup) > 0 {
				v.hallazgo(ReglaImporteConcepto, r, "Importe", cp.Importe,
					fmt.Sprintf("Importe %s fuera de [%s, %s] (Cantidad %s × ValorUnitario %s)", cp.Importe.Original(), inf, sup, cp.Cantidad.Original(), cp.ValorUnitario.Original()))
			}
		}
		if cp.Descuento != nil && cp.Importe != nil && cp.Descuento.Cmp(*cp.Importe) > 0 {
			v.hallazgo(ReglaDescuentoConcepto, r, "Descuento", cp.Descuento,
				fmt.Sprintf("Descuento %s mayor que el Importe %s", cp.Descuento.Original(), cp.Importe.Original()))
		}
		if v.v40 {
			v.objetoImp(cp, r)
		}
		if cp.Impuestos == nil {
			continue
		}
		for j, t := range cp.Impuestos.Traslados {
			v.impuestoConcepto(t, idx(r+"/Impuestos/Traslados", "Traslado", j), true)
		}
		for j, t := range cp.Impuestos.Retenciones {
			v.impuestoConcepto(t, idx(r+"/Impuestos/Retenciones", "Retencion", j), false)
		}
	}
}

// objetoImp revisa el nodo Impuestos del concepto contra su ObjetoImp
// (matriz 4.0 CFDI40171, CFDI40172, CFDI140226 y CFDI140227).
func (v *validadorReglas) objetoImp(cp Concepto, r string) {
	obj := clave(cp.ObjetoImp)
	tiene := func(impuesto string, traslados bool, retenciones bool) bool {
		if cp.Impuestos == nil {
			return false
		}
		var lista []Impuesto
		if traslados {
			lista = append(lista, cp.Impuestos.Traslados...)
		}
		if retenciones {
			lista = append(lista, cp.Impuestos.Retenciones...)
		}
		for _, t := range lista {
			if clave(t.Impuesto) == impuesto {
				return true
			}
		}
		return false
	}
	msg := ""
	switch obj {
	case "02":
		if cp.Impuestos == nil {
			msg = "ObjetoImp 02 (sí objeto de impuesto) y el concepto no trae Impuestos"
		}
	case "01", "03", "04", "05":
		if cp.Impuestos != nil {
			msg = "ObjetoImp " + obj + " y el concepto trae Impuestos"
		}
	case "06", "08":
		if tiene("002", true, true) || tiene("003", true, true) {
			msg = "ObjetoImp " + obj + " y el concepto trae impuestos de IVA o IEPS"
		}
	case "07":
		switch {
		case tiene("002", true, true):
			msg = "ObjetoImp 07 y el concepto trae IVA"
		case !tiene("003", true, false):
			msg = "ObjetoImp 07 y el concepto no trae el traslado de IEPS"
		}
	}
	if msg != "" {
		v.hallazgoTexto(ReglaObjetoImp, r, "ObjetoImp", cp.ObjetoImp, msg)
	}
}

func (v *validadorReglas) impuestoConcepto(t Impuesto, r string, traslado bool) {
	factor := clave(t.TipoFactor)
	switch {
	case !traslado && factor == "Exento":
		v.hallazgoTexto(ReglaRetencionExenta, r, "TipoFactor", t.TipoFactor, "una retención no puede ser Exento")
		return
	case factor == "Exento":
		if t.TasaOCuota != nil || t.Importe != nil {
			v.hallazgoTexto(ReglaExentoConImporte, r, "TipoFactor", t.TipoFactor, "traslado Exento con TasaOCuota o Importe")
		}
		return
	case traslado && (factor == "Tasa" || factor == "Cuota") && (t.TasaOCuota == nil || t.Importe == nil):
		v.hallazgoTexto(ReglaTasaSinImporte, r, "TipoFactor", t.TipoFactor, "traslado "+t.TipoFactor+" sin TasaOCuota o sin Importe")
		return
	}
	if t.Base == nil || t.TasaOCuota == nil || t.Importe == nil {
		return
	}
	inf, sup := limites(*t.Base, *t.TasaOCuota, true, t.Importe.Scale())
	if t.Importe.Cmp(inf) < 0 || t.Importe.Cmp(sup) > 0 {
		// Segundo nivel: los mismos límites a los decimales de la moneda.
		mInf, mSup := limites(*t.Base, *t.TasaOCuota, true, v.decimalesPara(*t.Importe))
		v.hallazgoDosNiveles(ReglaImporteImpuesto, r, "Importe", t.Importe,
			fmt.Sprintf("Importe %s fuera de [%s, %s] (Base %s × TasaOCuota %s)", t.Importe.Original(), inf, sup, t.Base.Original(), t.TasaOCuota.Original()),
			t.Importe.Cmp(mInf) >= 0 && t.Importe.Cmp(mSup) <= 0,
			fmt.Sprintf("margen de redondeo por concepto, [%s, %s] a los decimales de la moneda", mInf, mSup))
	}
}

// grupoImpuesto junta los impuestos de los conceptos por (Impuesto,
// TipoFactor, TasaOCuota) en traslados y por Impuesto en retenciones. La
// tasa se agrupa por valor ("0.16" y "0.160000" son la misma).
type grupoImpuesto struct {
	base, importe   decimal.Decimal // sumas exactas de lo escrito
	bases, importes []termino       // sumandos, para el margen
}

// termino es un sumando de un total del comprobante: el valor escrito y,
// en el Importe de un impuesto del concepto, Base × TasaOCuota sin
// redondear (calculado), que es lo que algunos emisores suman en vez del
// Importe ya redondeado del concepto.
type termino struct {
	escrito   decimal.Decimal
	calculado *decimal.Decimal
}

func terminoImpuesto(t Impuesto) termino {
	tm := termino{escrito: valorDe(t.Importe)}
	if t.Importe != nil && t.Base != nil && t.TasaOCuota != nil {
		bt := t.Base.Mul(*t.TasaOCuota)
		tm.calculado = &bt
	}
	return tm
}

func claveGrupo(t Impuesto, traslado bool) string {
	if !traslado {
		return clave(t.Impuesto)
	}
	tasa := ""
	if t.TasaOCuota != nil {
		tasa = t.TasaOCuota.String()
	}
	return clave(t.Impuesto) + "|" + clave(t.TipoFactor) + "|" + tasa
}

// En 3.3 los traslados Exento no se agrupan: el Traslado del comprobante
// exige TasaOCuota e Importe (cfdv33.xsd) y la guía 3.3 (p. 32) dice que
// con solo conceptos exentos el nodo no debe existir. En 4.0 sí van, con su
// Base (matriz 4.0 CFDI40213 y CFDI40216).
func agrupar(cs []Concepto, traslado, v40 bool) (map[string]*grupoImpuesto, []string) {
	m := map[string]*grupoImpuesto{}
	var orden []string
	for _, cp := range cs {
		if cp.Impuestos == nil {
			continue
		}
		lista := cp.Impuestos.Retenciones
		if traslado {
			lista = cp.Impuestos.Traslados
		}
		for _, t := range lista {
			if traslado && !v40 && clave(t.TipoFactor) == "Exento" {
				continue
			}
			k := claveGrupo(t, traslado)
			g := m[k]
			if g == nil {
				g = &grupoImpuesto{}
				m[k] = g
				orden = append(orden, k)
			}
			g.base = g.base.Add(valorDe(t.Base))
			g.importe = g.importe.Add(valorDe(t.Importe))
			g.bases = append(g.bases, termino{escrito: valorDe(t.Base)})
			g.importes = append(g.importes, terminoImpuesto(t))
		}
	}
	return m, orden
}

// sumaCoincide dice si declarado es la suma exacta o su redondeo a los
// decimales de la moneda. La matriz 4.0 pide «el redondeo de la suma»
// (CFDI40108, CFDI40211, CFDI40215, CFDI40221) y la 3.3 «la suma» con
// importes de hasta 6 decimales; las guías piden «Redondear al final del
// cálculo y no antes» (4.0 pregunta 39, 3.3 pregunta 44): en la letra no
// hay tolerancia por redondeo de cada concepto; esa es el segundo nivel
// (intervaloSuma).
func (v *validadorReglas) sumaCoincide(declarado, suma decimal.Decimal) bool {
	return declarado.Equal(suma) || declarado.Equal(suma.Round(v.decimalesPara(declarado)))
}

// intervaloSuma es el margen de las sumas de impuestos (Regla.Margen): el
// intervalo en que cae un total formado con cada término, escrito o
// calculado, truncado o redondeado hacia arriba a n decimales antes de
// sumar. Contiene la suma exacta y su redondeo. Los PAC timbran totales
// con 1 o 2 centavos de diferencia que solo se explican así.
func intervaloSuma(ts []termino, n int32) (lo, hi decimal.Decimal) {
	for _, t := range ts {
		a, b := t.escrito, t.escrito
		if c := t.calculado; c != nil {
			if c.Cmp(a) < 0 {
				a = *c
			} else {
				b = *c
			}
		}
		lo, hi = lo.Add(truncar(a, n)), hi.Add(techo(b, n))
	}
	return lo, hi
}

// revisarSuma evalúa una suma de dos niveles: sin hallazgo si declarado es
// la suma o su redondeo (la letra del SAT); advertencia si cae en el
// intervaloSuma de ts; error si no. porSumando dice qué se redondea en el
// margen, para el mensaje.
func (v *validadorReglas) revisarSuma(codigo, ruta, attr string, declarado *decimal.Decimal, suma decimal.Decimal, ts []termino, msg, porSumando string) {
	if v.sumaCoincide(*declarado, suma) {
		return
	}
	lo, hi := intervaloSuma(ts, v.decimalesPara(*declarado))
	v.hallazgoDosNiveles(codigo, ruta, attr, declarado, msg, declarado.Cmp(lo) >= 0 && declarado.Cmp(hi) <= 0,
		fmt.Sprintf("margen de redondeo por %s, [%s, %s]", porSumando, lo, hi))
}

func (v *validadorReglas) sumas() {
	c := v.c
	var importes, descuentos decimal.Decimal
	todos, conDescuento := true, false
	for _, cp := range c.Conceptos {
		if cp.Importe == nil {
			todos = false
		}
		importes = importes.Add(valorDe(cp.Importe))
		if cp.Descuento != nil {
			conDescuento = true
			descuentos = descuentos.Add(*cp.Descuento)
		}
	}
	ien := v.tipo == "I" || v.tipo == "E" || v.tipo == "N"
	if ien && c.SubTotal != nil && todos && !v.sumaCoincide(*c.SubTotal, importes) {
		v.hallazgo(ReglaSubTotal, "/Comprobante", "SubTotal", c.SubTotal,
			fmt.Sprintf("SubTotal %s y la suma de los importes de los conceptos es %s", c.SubTotal.Original(), importes))
	}
	if conDescuento || c.Descuento != nil {
		d := valorDe(c.Descuento)
		if !v.sumaCoincide(d, descuentos) {
			v.hallazgo(ReglaDescuento, "/Comprobante", "Descuento", c.Descuento,
				fmt.Sprintf("Descuento %s y la suma de los descuentos de los conceptos es %s", d.Original(), descuentos))
		}
	}
	if c.Descuento != nil && c.SubTotal != nil && c.Descuento.Cmp(*c.SubTotal) > 0 {
		v.hallazgo(ReglaDescuentoMayor, "/Comprobante", "Descuento", c.Descuento,
			fmt.Sprintf("Descuento %s mayor que SubTotal %s", c.Descuento.Original(), c.SubTotal.Original()))
	}

	imp := c.Impuestos
	if imp == nil {
		imp = &Impuestos{}
	}
	// Total.
	if c.Total != nil && c.SubTotal != nil {
		t := c.SubTotal.Sub(valorDe(c.Descuento)).Add(valorDe(imp.TotalImpuestosTrasladados)).Sub(valorDe(imp.TotalImpuestosRetenidos))
		detalle := ""
		if il := c.ImpuestosLocales; il != nil {
			t = t.Add(valorDe(il.TotaldeTraslados)).Sub(valorDe(il.TotaldeRetenciones))
			detalle = " (con impuestos locales)"
		}
		if !c.Total.Equal(t.Round(v.decimalesPara(*c.Total))) {
			v.hallazgo(ReglaTotal, "/Comprobante", "Total", c.Total,
				fmt.Sprintf("Total %s y SubTotal − Descuento + trasladados − retenidos%s da %s", c.Total.Original(), detalle, t))
		}
	}
	// Totales de impuestos contra sus nodos.
	totalImpuestos := func(total *decimal.Decimal, lista []Impuesto, codigo, attr, nodos string) {
		var s decimal.Decimal
		var ts []termino
		for _, t := range lista {
			if t.Importe != nil {
				s = s.Add(*t.Importe)
				ts = append(ts, termino{escrito: *t.Importe})
			}
		}
		switch {
		case total != nil:
			v.revisarSuma(codigo, "/Comprobante/Impuestos", attr, total, s, ts,
				fmt.Sprintf("%s %s y la suma de Impuestos/%s es %s", attr, total.Original(), nodos, s), "cada Importe de Impuestos/"+nodos)
		case len(ts) > 0:
			v.hallazgo(ReglaImpuestosSinTotal, "/Comprobante/Impuestos", attr, nil, "hay "+strings.ToLower(nodos)+" con importe y falta "+attr)
		}
	}
	totalImpuestos(imp.TotalImpuestosTrasladados, imp.Traslados, ReglaTotalTraslados, "TotalImpuestosTrasladados", "Traslados")
	totalImpuestos(imp.TotalImpuestosRetenidos, imp.Retenciones, ReglaTotalRetenciones, "TotalImpuestosRetenidos", "Retenciones")
	v.grupos(imp.Traslados, true)
	v.grupos(imp.Retenciones, false)
}

// grupos compara los impuestos del comprobante con los de los conceptos
// agrupados.
func (v *validadorReglas) grupos(declarados []Impuesto, traslado bool) {
	calc, orden := agrupar(v.c.Conceptos, traslado, v.v40)
	cSuma, cFalta, cSobra, nodo, base := ReglaRetencionSuma, ReglaRetencionFaltante, ReglaRetencionSobrante, "Retencion", "/Comprobante/Impuestos/Retenciones"
	if traslado {
		cSuma, cFalta, cSobra, nodo, base = ReglaTrasladoSuma, ReglaTrasladoFaltante, ReglaTrasladoSobrante, "Traslado", "/Comprobante/Impuestos/Traslados"
	}
	vistos := map[string]bool{}
	for j, t := range declarados {
		r := idx(base, nodo, j)
		k := claveGrupo(t, traslado)
		g := calc[k]
		if g == nil {
			v.hallazgoTexto(cSobra, r, "Impuesto", t.Impuesto, "el grupo "+describirGrupo(k)+" no está en ningún concepto")
			continue
		}
		vistos[k] = true
		if t.Importe != nil {
			v.revisarSuma(cSuma, r, "Importe", t.Importe, g.importe, g.importes,
				fmt.Sprintf("Importe %s y la suma en los conceptos de %s es %s", t.Importe.Original(), describirGrupo(k), g.importe), "concepto")
		}
		if traslado && v.v40 && t.Base != nil {
			v.revisarSuma(cSuma, r, "Base", t.Base, g.base, g.bases,
				fmt.Sprintf("Base %s y la suma en los conceptos de %s es %s", t.Base.Original(), describirGrupo(k), g.base), "concepto")
		}
	}
	for _, k := range orden {
		if !vistos[k] {
			v.hallazgo(cFalta, base, "", nil, "los conceptos tienen "+describirGrupo(k)+" y no está en el comprobante")
		}
	}
}

func describirGrupo(k string) string {
	p := strings.Split(k, "|")
	s := "Impuesto " + p[0]
	if len(p) == 3 {
		s += " " + p[1]
		if p[2] != "" {
			s += " " + p[2]
		}
	}
	return s
}

// excedeDecimales dice si d tiene más decimales que n.
func excedeDecimales(d *decimal.Decimal, n int32) bool { return d != nil && d.Scale() > n }

func (v *validadorReglas) decimales() {
	if !v.monedaOK {
		return
	}
	c := v.c
	revisar := func(codigo, ruta, attr string, d *decimal.Decimal) {
		if excedeDecimales(d, v.dec) {
			v.hallazgo(codigo, ruta, attr, d, fmt.Sprintf("%s tiene %d decimales y la moneda %s admite %d", d.Original(), d.Scale(), v.moneda, v.dec))
		}
	}
	revisar(ReglaDecimalesMoneda, "/Comprobante", "SubTotal", c.SubTotal)
	revisar(ReglaDecimalesMoneda, "/Comprobante", "Descuento", c.Descuento)
	revisar(ReglaDecimalesMoneda, "/Comprobante", "Total", c.Total)
	imp := c.Impuestos
	if imp == nil {
		return
	}
	revisar(ReglaDecimalesMoneda, "/Comprobante/Impuestos", "TotalImpuestosTrasladados", imp.TotalImpuestosTrasladados)
	revisar(ReglaDecimalesMoneda, "/Comprobante/Impuestos", "TotalImpuestosRetenidos", imp.TotalImpuestosRetenidos)
	codigo := ReglaDecimalesImpuesto33
	if v.v40 {
		codigo = ReglaDecimalesImpuesto
	}
	for j, t := range imp.Traslados {
		r := idx("/Comprobante/Impuestos/Traslados", "Traslado", j)
		if v.v40 {
			revisar(codigo, r, "Base", t.Base)
		}
		revisar(codigo, r, "Importe", t.Importe)
	}
	for j, t := range imp.Retenciones {
		revisar(codigo, idx("/Comprobante/Impuestos/Retenciones", "Retencion", j), "Importe", t.Importe)
	}
}

var uno = decimal.NewFromInt64(1, 0)

func (v *validadorReglas) tipoCambio() {
	c := v.c
	switch {
	case v.moneda == "":
	case v.moneda == "MXN":
		if c.TipoCambio != nil && !c.TipoCambio.Equal(uno) {
			v.hallazgo(ReglaTipoCambioMXN, "/Comprobante", "TipoCambio", c.TipoCambio, "con Moneda MXN el TipoCambio tiene que ser 1 o no venir")
		}
	case v.moneda == "XXX":
		if c.TipoCambio != nil {
			v.hallazgo(ReglaTipoCambioXXX, "/Comprobante", "TipoCambio", c.TipoCambio, "con Moneda XXX no va TipoCambio")
		}
	default:
		if c.TipoCambio == nil {
			v.hallazgo(ReglaTipoCambioRequerido, "/Comprobante", "TipoCambio", nil, "con Moneda "+v.moneda+" se requiere TipoCambio")
		}
	}
}

func (v *validadorReglas) tipoComprobante() {
	c := v.c
	t := v.tipo
	tpn := t == "T" || t == "P" || t == "N"
	tp := t == "T" || t == "P"
	if tpn && c.CondicionesDePago != "" {
		v.hallazgoTexto(ReglaTipoCompCondiciones, "/Comprobante", "CondicionesDePago", c.CondicionesDePago, "un comprobante "+t+" no lleva CondicionesDePago")
	}
	if tpn && c.Impuestos != nil {
		v.hallazgo(ReglaTipoCompImpuestos, "/Comprobante/Impuestos", "", nil, "un comprobante "+t+" no lleva el nodo Impuestos")
	}
	if (tp || (v.v40 && t == "N")) && c.FormaPago != "" {
		v.hallazgoTexto(ReglaTipoCompFormaPago, "/Comprobante", "FormaPago", c.FormaPago, "un comprobante "+t+" no lleva FormaPago")
	}
	if tp && c.MetodoPago != "" {
		v.hallazgoTexto(ReglaTipoCompMetodoPago, "/Comprobante", "MetodoPago", c.MetodoPago, "un comprobante "+t+" no lleva MetodoPago")
	}
	if tp && c.Descuento != nil {
		v.hallazgo(ReglaTipoCompDescuento, "/Comprobante", "Descuento", c.Descuento, "un comprobante "+t+" no lleva Descuento")
	}
	for i, cp := range c.Conceptos {
		r := idx("/Comprobante/Conceptos", "Concepto", i)
		if tp && cp.Descuento != nil {
			v.hallazgo(ReglaTipoCompDescuento, r, "Descuento", cp.Descuento, "los conceptos de un comprobante "+t+" no llevan Descuento")
		}
		if (t == "I" || t == "E" || t == "N") && cp.ValorUnitario != nil && cp.ValorUnitario.Sign() <= 0 {
			v.hallazgo(ReglaTipoCompValorUnitario, r, "ValorUnitario", cp.ValorUnitario, "en un comprobante "+t+" el ValorUnitario tiene que ser mayor que cero")
		}
	}
	if tp && c.SubTotal != nil && !c.SubTotal.IsZero() {
		v.hallazgo(ReglaTipoCompSubTotalCero, "/Comprobante", "SubTotal", c.SubTotal, "un comprobante "+t+" va con SubTotal 0")
	}
	if tp && c.Total != nil && !c.Total.IsZero() {
		v.hallazgo(ReglaTipoCompTotalCero, "/Comprobante", "Total", c.Total, "un comprobante "+t+" va con Total 0")
	}
	if t == "N" && v.moneda != "" && v.moneda != "MXN" {
		v.hallazgoTexto(ReglaTipoCompMonedaNomina, "/Comprobante", "Moneda", c.Moneda, "un comprobante de nómina va en MXN")
	}
	if t == "P" {
		v.comprobantePago()
	}
}

func (v *validadorReglas) impLocal() {
	il := v.c.ImpuestosLocales
	if il == nil {
		return
	}
	r := "/Comprobante/Complemento/ImpuestosLocales"
	var tras, ret decimal.Decimal
	for _, t := range il.TrasladosLocales {
		tras = tras.Add(valorDe(t.Importe))
	}
	for _, t := range il.RetencionesLocales {
		ret = ret.Add(valorDe(t.Importe))
	}
	if il.TotaldeTraslados != nil && !il.TotaldeTraslados.Equal(tras) {
		v.hallazgo(ReglaImpLocalTotales, r, "TotaldeTraslados", il.TotaldeTraslados, fmt.Sprintf("TotaldeTraslados %s y la suma de TrasladosLocales es %s", il.TotaldeTraslados.Original(), tras))
	}
	if il.TotaldeRetenciones != nil && !il.TotaldeRetenciones.Equal(ret) {
		v.hallazgo(ReglaImpLocalTotales, r, "TotaldeRetenciones", il.TotaldeRetenciones, fmt.Sprintf("TotaldeRetenciones %s y la suma de RetencionesLocales es %s", il.TotaldeRetenciones.Original(), ret))
	}
}
