package cfdi

import (
	"fmt"
	"math/big"
	"strings"
	"unicode/utf8"

	"github.com/OrtaMarco/cfdi-go/decimal"
)

// Reglas de presencia y coherencia de Comercio Exterior 2.0 (matriz CCE
// 2.0, CCE102-172) que se pueden comprobar sin catálogos externos: lo que
// pide c_Pais (Unión Europea, patrones de NumRegIdTrib y de código postal),
// c_Estado, c_Municipio, c_Localidad, c_Colonia, c_CodigoPostal,
// c_FraccionArancelaria (vigencia y UMT), la lista de RFC y el tipo de
// cambio publicado queda fuera (ver el README).
const (
	ReglaCCEExportacion    = "REGLA_CCE_EXPORTACION"
	ReglaCCENodosCFDI      = "REGLA_CCE_NODOS_CFDI"
	ReglaCCEMotivoTraslado = "REGLA_CCE_MOTIVO_TRASLADO"
	ReglaCCETraslado       = "REGLA_CCE_TRASLADO"
	ReglaCCEEmisor         = "REGLA_CCE_EMISOR"
	ReglaCCENumRegIdTrib   = "REGLA_CCE_NUMREGIDTRIB"
	ReglaCCEFraccion       = "REGLA_CCE_FRACCION"
	ReglaCCEValorImporte   = "REGLA_CCE_VALOR_IMPORTE"
	ReglaCCEMuestras       = "REGLA_CCE_MUESTRAS"
)

// DocEstCCE20 es el estándar del complemento; sus páginas son las del PDF
// (no las imprime).
const DocEstCCE20 = "Estándar Comercio Exterior 2.0"

var reglasCCE = []Regla{
	regla(ReglaCCEExportacion, NivelError,
		"CFDI 4.0: Exportacion 02 con Comercio Exterior; con Comercio Exterior 2.0, Exportacion distinta de 03, y con 02, Incoterm",
		citas(cod(DocMatrizCCE20, "CCE104", "CCE105", "CCE120"), pag(DocEstCCE20, 37), pag(DocGuiaCCE20, 14, 45)), ""),
	conSoloGuia(regla(ReglaCCENodosCFDI, NivelError,
		"Comercio Exterior 2.0: sin InformacionGlobal, ni ACuentaTerceros, InformacionAduanera o CuentaPredial en los conceptos; en T, ObjetoImp 01 en cada concepto; en el Complemento, de los complementos que lee la librería, solo TimbreFiscalDigital, ImpuestosLocales, LeyendasFiscales, CFDIRegistroFiscal y CartaPorte",
		citas(cod(DocMatrizCCE20, "CCE106", "CCE110", "CCE111", "CCE112", "CCE115"), pag(DocEstCCE20, 35, 36), pag(DocGuiaCCE20, 17, 28, 32, 42)),
		"«Otros derechos e impuestos» es ImpuestosLocales; los complementos que la librería no lee no se revisan; InformacionAduanera no está en la matriz ni en el estándar (p. 35-36), solo en la guía (p. 32: «Este nodo no debe existir»)"),
		"sin InformacionAduanera en los conceptos (guía p. 32)"),
	regla(ReglaCCEMotivoTraslado, NivelError,
		"Comercio Exterior 2.0: MotivoTraslado en T y solo en T; con MotivoTraslado 05, al menos un Propietario, y sin 05 ninguno; con 01, CfdiRelacionados con TipoRelacion y UUID",
		citas(cod(DocMatrizCCE20, "CCE102", "CCE103", "CCE116", "CCE117"), pag(DocEstCCE20, 35, 36), pag(DocGuiaCCE20, 13, 14, 42, 43)),
		"con 01 la guía (p. 43) pide TipoRelacion 05; la matriz solo que haya TipoRelacion y UUID: se aplica la matriz"),
	regla(ReglaCCETraslado, NivelError,
		"Comercio Exterior 2.0 en T: Rfc del Receptor igual al del Emisor; un solo Destinatario; si una mercancía no trae CantidadAduana, sus conceptos con Cantidad de al menos 0.001 y hasta 3 decimales, Unidad del catálogo c_UnidadAduana y ValorUnitario de hasta 4 decimales y no más de los de la moneda",
		citas(cod(DocMatrizCCE20, "CCE108", "CCE142", "CCE154", "CCE155", "CCE156"), pag(DocEstCCE20, 40, 41, 42), pag(DocGuiaCCE20, 54, 55, 70)), ""),
	regla(ReglaCCEEmisor, NivelError,
		"Comercio Exterior 2.0: Emisor/Curp solo si el Rfc del emisor del CFDI tiene 13 caracteres (persona física), y entonces requerida; Emisor/Domicilio/Pais MEX",
		citas(cod(DocMatrizCCE20, "CCE123", "CCE124", "CCE125"), pag(DocEstCCE20, 37), pag(DocGuiaCCE20, 46, 47)), ""),
	regla(ReglaCCENumRegIdTrib, NivelError,
		"Comercio Exterior 2.0: Receptor/NumRegIdTrib requerido si el Rfc del receptor del CFDI es XEXX010101000 y ausente si no; Receptor y Destinatario con NumRegIdTrib tienen domicilio con Pais distinto de MEX",
		citas(cod(DocMatrizCCE20, "CCE133", "CCE134", "CCE143"), pag(DocEstCCE20, 38, 40), pag(DocGuiaCCE20, 48, 51)),
		"si el NumRegIdTrib existe en el registro del país o cumple su patrón (CCE131, CCE132, CCE135, CCE144) necesita c_Pais y no se revisa"),
	regla(ReglaCCEFraccion, NivelError,
		"Comercio Exterior 2.0: FraccionArancelaria ausente si UnidadAduana es 99 o un concepto de la mercancía tiene ClaveUnidad E48 (servicios); requerida en otro caso",
		citas(cod(DocMatrizCCE20, "CCE159", "CCE160", "CCE170"), pag(DocEstCCE20, 43, 44), pag(DocGuiaCCE20, 56, 57)), ""),
	regla(ReglaCCEValorImporte, NivelError,
		"Comercio Exterior 2.0 en I y E: por cada NoIdentificacion sin ValorDolares 0 ni 1, la suma de sus ValorDolares entre (Σ Importe − ½·10^-dI)·(TipoCambio − ½·10^-dT)/(TipoCambioUSD + ½·10^-dU − 10^-12) truncado y (Σ Importe + ½·10^-dI − 10^-12)·(TipoCambio + ½·10^-dT − 10^-12)/(TipoCambioUSD − ½·10^-dU) redondeado hacia arriba, a 2 decimales, con Σ Importe la de los conceptos de ese NoIdentificacion; con Moneda MXN el TipoCambio es 1 exacto, y con USD vale también Σ Importe ± ½·10^-dI",
		citas(cod(DocMatrizCCE20, "CCE157", "CCE158", "CCE172"), pag(DocEstCCE20, 42, 43, 44), pag(DocGuiaCCE20, 55, 56, 58, 59)),
		"el estándar escribe NumDecimalesTipoCambio en los dos tipos de cambio: se usan los decimales de cada uno; con USD la guía (p. 58-59) pide ValorDolares igual al importe aunque el TipoCambio pactado difiera del TipoCambioUSD; en T la conversión no está definida (CCE157 es de I y E) y no se revisa"),
	regla(ReglaCCEMuestras, NivelError,
		"Comercio Exterior 2.0 en I y E: con FraccionArancelaria 9801000100 (muestras), la suma de sus ValorDolares por NoIdentificacion no mayor que la suma de los Descuento de los conceptos de ese NoIdentificacion en dólares (el límite superior de la conversión de REGLA_CCE_VALOR_IMPORTE)",
		citas(cod(DocMatrizCCE20, "CCE163", "CCE164"), pag(DocEstCCE20, 43), pag(DocGuiaCCE20, 56, 57)),
		"el SAT no dice con qué tolerancia se convierte el Descuento: se toma el límite superior de la conversión del estándar p. 42"),
}

// complementosIncompatiblesCCE son los complementos del Comprobante (de los
// que lee la librería) que no pueden ir junto a Comercio Exterior 2.0
// (CCE115).
func (v *validadorReglas) complementosIncompatiblesCCE() []string {
	c := v.c
	var out []string
	if c.Pagos != nil {
		out = append(out, "Pagos")
	}
	if len(c.Nomina) > 0 {
		out = append(out, "Nomina")
	}
	if c.Divisas != nil {
		out = append(out, "Divisas")
	}
	if c.EstadoDeCuentaCombustible != nil {
		out = append(out, "EstadoDeCuentaCombustible")
	}
	if len(c.HidroYPetro) > 0 {
		out = append(out, "HidroYPetro")
	}
	return out
}

// exportacionSinCCE revisa CCE104: Exportacion 02 sin el complemento. Si
// viene un ComercioExterior que la librería no lee (1.0), no se sabe.
func (v *validadorReglas) exportacionSinCCE() {
	c := v.c
	if !v.v40 || c.ComercioExterior != nil || clave(c.Exportacion) != "02" {
		return
	}
	for _, n := range c.ComplementosNoSoportados {
		if strings.HasSuffix(n, "}ComercioExterior") {
			return
		}
	}
	v.hallazgoTexto(ReglaCCEExportacion, "/Comprobante", "Exportacion", c.Exportacion, "con Exportacion 02 (definitiva con clave A1) va el complemento de Comercio Exterior")
}

// cce revisa las reglas de presencia y coherencia de Comercio Exterior 2.0.
// conceptos son los del comprobante por NoIdentificacion.
func (v *validadorReglas) cce(ce *ComercioExterior, conceptos map[string][]Concepto) {
	c := v.c
	r := "/Comprobante/Complemento/ComercioExterior"
	exp := clave(c.Exportacion)
	if exp == "03" {
		v.hallazgoTexto(ReglaCCEExportacion, "/Comprobante", "Exportacion", c.Exportacion, "con Exportacion 03 (temporal) no va el complemento de Comercio Exterior")
	}
	if exp == "02" && clave(ce.Incoterm) == "" {
		v.hallazgo(ReglaCCEExportacion, r, "Incoterm", nil, "con Exportacion 02 se requiere Incoterm")
	}

	// Nodos del CFDI.
	if c.InformacionGlobal != nil {
		v.hallazgo(ReglaCCENodosCFDI, "/Comprobante/InformacionGlobal", "", nil, "con Comercio Exterior no va InformacionGlobal")
	}
	for i, cp := range c.Conceptos {
		rc := idx("/Comprobante/Conceptos", "Concepto", i)
		if cp.ACuentaTerceros != nil {
			v.hallazgo(ReglaCCENodosCFDI, rc+"/ACuentaTerceros", "", nil, "con Comercio Exterior el concepto no lleva ACuentaTerceros")
		}
		if len(cp.CuentaPredial) > 0 {
			v.hallazgo(ReglaCCENodosCFDI, rc+"/CuentaPredial", "", nil, "con Comercio Exterior el concepto no lleva CuentaPredial")
		}
		if len(cp.InformacionAduanera) > 0 {
			v.hallazgo(ReglaCCENodosCFDI, rc+"/InformacionAduanera", "", nil, "con Comercio Exterior el concepto no lleva InformacionAduanera")
			v.soloGuia("guía de Comercio Exterior p. 32")
		}
		if o := clave(cp.ObjetoImp); v.tipo == "T" && o != "" && o != "01" {
			v.hallazgoTexto(ReglaCCENodosCFDI, rc, "ObjetoImp", cp.ObjetoImp, "en un traslado con Comercio Exterior el ObjetoImp es 01")
		}
	}
	for _, n := range v.complementosIncompatiblesCCE() {
		v.hallazgo(ReglaCCENodosCFDI, "/Comprobante/Complemento", "", nil, "Comercio Exterior no puede ir con el complemento "+n)
	}

	// MotivoTraslado y Propietario.
	motivo := clave(ce.MotivoTraslado)
	switch {
	case v.tipo == "T" && motivo == "":
		v.hallazgo(ReglaCCEMotivoTraslado, r, "MotivoTraslado", nil, "en un traslado se requiere MotivoTraslado")
	case v.tipo != "T" && v.tipo != "" && motivo != "":
		v.hallazgoTexto(ReglaCCEMotivoTraslado, r, "MotivoTraslado", ce.MotivoTraslado, "MotivoTraslado solo va en un comprobante T")
	}
	switch {
	case motivo == "05" && len(ce.Propietario) == 0:
		v.hallazgo(ReglaCCEMotivoTraslado, r, "MotivoTraslado", nil, "con MotivoTraslado 05 se requiere al menos un Propietario")
	case motivo != "05" && len(ce.Propietario) > 0:
		v.hallazgo(ReglaCCEMotivoTraslado, r+"/Propietario[1]", "", nil, "sin MotivoTraslado 05 no va Propietario")
	}
	if motivo == "01" && !conUUIDRelacionado(c) {
		v.hallazgoTexto(ReglaCCEMotivoTraslado, r, "MotivoTraslado", ce.MotivoTraslado, "con MotivoTraslado 01 se requiere CfdiRelacionados con TipoRelacion y el UUID de la factura")
	}

	// Traslados.
	if v.tipo == "T" {
		if c.Emisor != nil && c.Receptor != nil && clave(c.Receptor.Rfc) != clave(c.Emisor.Rfc) {
			v.hallazgoTexto(ReglaCCETraslado, "/Comprobante/Receptor", "Rfc", c.Receptor.Rfc, "en un traslado con Comercio Exterior el Rfc del receptor es el del emisor")
		}
		if len(ce.Destinatario) > 1 {
			v.hallazgo(ReglaCCETraslado, idx(r, "Destinatario", 1), "", nil, fmt.Sprintf("en un traslado va un solo Destinatario y hay %d", len(ce.Destinatario)))
		}
		revisados := map[string]bool{}
		for _, m := range ce.Mercancias {
			if m.CantidadAduana != nil || revisados[m.NoIdentificacion] {
				continue
			}
			revisados[m.NoIdentificacion] = true
			v.conceptosSinCantidadAduana(m.NoIdentificacion)
		}
	}

	// Emisor.
	if c.Emisor != nil && ce.Emisor != nil {
		re := r + "/Emisor"
		switch n := utf8.RuneCountInString(clave(c.Emisor.Rfc)); {
		case n == 12 && ce.Emisor.Curp != "":
			v.hallazgoTexto(ReglaCCEEmisor, re, "Curp", ce.Emisor.Curp, "el emisor es persona moral (Rfc de 12): no va Curp")
		case n == 13 && ce.Emisor.Curp == "":
			v.hallazgo(ReglaCCEEmisor, re, "Curp", nil, "el emisor es persona física (Rfc de 13): se requiere Curp")
		}
		if d := ce.Emisor.Domicilio; d != nil && clave(d.Pais) != "" && clave(d.Pais) != "MEX" {
			v.hallazgoTexto(ReglaCCEEmisor, re+"/Domicilio", "Pais", d.Pais, "el domicilio del emisor es MEX")
		}
	}

	// NumRegIdTrib.
	if rc := ce.Receptor; c.Receptor != nil {
		generico := clave(c.Receptor.Rfc) == "XEXX010101000"
		con := rc != nil && clave(rc.NumRegIdTrib) != ""
		switch {
		case generico && !con:
			v.hallazgo(ReglaCCENumRegIdTrib, r+"/Receptor", "NumRegIdTrib", nil, "con el Rfc genérico XEXX010101000 se requiere NumRegIdTrib")
		case !generico && con:
			v.hallazgoTexto(ReglaCCENumRegIdTrib, r+"/Receptor", "NumRegIdTrib", rc.NumRegIdTrib, "sin el Rfc genérico XEXX010101000 no va NumRegIdTrib")
		}
		if con && rc.Domicilio != nil && clave(rc.Domicilio.Pais) == "MEX" {
			v.hallazgoTexto(ReglaCCENumRegIdTrib, r+"/Receptor/Domicilio", "Pais", rc.Domicilio.Pais, "con NumRegIdTrib el domicilio del receptor no es MEX")
		}
	}
	for i, d := range ce.Destinatario {
		if clave(d.NumRegIdTrib) == "" {
			continue
		}
		for j, dom := range d.Domicilio {
			if clave(dom.Pais) == "MEX" {
				v.hallazgoTexto(ReglaCCENumRegIdTrib, idx(idx(r, "Destinatario", i), "Domicilio", j), "Pais", dom.Pais, "con NumRegIdTrib el domicilio del destinatario no es MEX")
			}
		}
	}

	// FraccionArancelaria.
	for i, m := range ce.Mercancias {
		rm := idx(r+"/Mercancias", "Mercancia", i)
		servicio := clave(m.UnidadAduana) == "99"
		for _, cp := range conceptos[m.NoIdentificacion] {
			servicio = servicio || clave(cp.ClaveUnidad) == "E48"
		}
		switch f := clave(m.FraccionArancelaria); {
		case servicio && f != "":
			v.hallazgoTexto(ReglaCCEFraccion, rm, "FraccionArancelaria", m.FraccionArancelaria, "un servicio (UnidadAduana 99 o ClaveUnidad E48) no lleva FraccionArancelaria")
		case !servicio && f == "":
			v.hallazgo(ReglaCCEFraccion, rm, "FraccionArancelaria", nil, "una mercancía que no es servicio lleva FraccionArancelaria")
		}
	}

	if v.tipo == "I" || v.tipo == "E" {
		v.cceValorImporte(ce, conceptos)
	}
}

// conUUIDRelacionado dice si el comprobante relaciona algún UUID con un
// TipoRelacion.
func conUUIDRelacionado(c *Comprobante) bool {
	for _, cr := range c.CfdiRelacionados {
		if clave(cr.TipoRelacion) == "" {
			continue
		}
		for _, u := range cr.CfdiRelacionado {
			if u.UUID != "" {
				return true
			}
		}
	}
	return false
}

var milesima = decimal.NewFromInt64(1, 3)

// conceptosSinCantidadAduana revisa CCE154-156 en los conceptos del
// NoIdentificacion de una mercancía de un traslado sin CantidadAduana.
func (v *validadorReglas) conceptosSinCantidadAduana(noIdent string) {
	for i, cp := range v.c.Conceptos {
		if cp.NoIdentificacion != noIdent {
			continue
		}
		rc := idx("/Comprobante/Conceptos", "Concepto", i)
		if q := cp.Cantidad; q != nil && (q.Cmp(milesima) < 0 || q.Scale() > 3) {
			v.hallazgo(ReglaCCETraslado, rc, "Cantidad", q, "sin CantidadAduana en la mercancía, la Cantidad del concepto es al menos 0.001 y tiene hasta 3 decimales")
		}
		if u := clave(cp.Unidad); !catCatComExt_cUnidadAduana.contiene(u) {
			v.hallazgoTexto(ReglaCCETraslado, rc, "Unidad", cp.Unidad, "sin CantidadAduana en la mercancía, la Unidad del concepto es una clave de c_UnidadAduana")
		}
		vu := cp.ValorUnitario
		if vu == nil {
			continue
		}
		maxDec := int32(4)
		if v.monedaOK && v.dec < maxDec {
			maxDec = v.dec
		}
		if vu.Sign() < 0 || vu.Scale() > maxDec {
			v.hallazgo(ReglaCCETraslado, rc, "ValorUnitario", vu, fmt.Sprintf("sin CantidadAduana en la mercancía, el ValorUnitario del concepto no es negativo y tiene hasta %d decimales", maxDec))
		}
	}
}

// intervalosUSD convierte un importe en la moneda del comprobante a dólares
// con los límites del estándar de Comercio Exterior 2.0 (p. 42), a 2
// decimales: [(i − ½·10^-dI)(tc − ½·10^-dT)/(usd + ½·10^-dU − 10^-12)
// truncado, (i + ½·10^-dI − 10^-12)(tc + ½·10^-dT − 10^-12)/(usd − ½·10^-dU)
// hacia arriba]. Con MXN el tipo de cambio es 1 exacto; con USD se agrega
// el intervalo del importe mismo (guía p. 58-59). nil si faltan tipos de
// cambio.
func (v *validadorReglas) intervalosUSD(importe decimal.Decimal) [][2]decimal.Decimal {
	ce := v.c.ComercioExterior
	usd := ce.TipoCambioUSD
	if usd == nil || usd.Sign() <= 0 {
		return nil
	}
	hI := mitadUlp(importe.Scale())
	iLo, iHi := importe.Sub(hI), importe.Add(hI).Sub(casiCero)
	var out [][2]decimal.Decimal
	if v.moneda == "USD" {
		out = append(out, [2]decimal.Decimal{truncar(iLo, 2), techo(iHi, 2)})
	}
	tcLo, tcHi := uno, uno
	switch tc := v.c.TipoCambio; {
	case v.moneda == "MXN":
	case tc == nil || tc.Sign() <= 0:
		return out
	default:
		hT := mitadUlp(tc.Scale())
		tcLo, tcHi = tc.Sub(hT), tc.Add(hT).Sub(casiCero)
	}
	hU := mitadUlp(usd.Scale())
	uLo, uHi := usd.Sub(hU), usd.Add(hU).Sub(casiCero)
	if uLo.Sign() <= 0 {
		return out
	}
	lo := truncarRat(new(big.Rat).Quo(rat(iLo.Mul(tcLo)), rat(uHi)), 2)
	hi := techoRat(new(big.Rat).Quo(rat(iHi.Mul(tcHi)), rat(uLo)), 2)
	return append(out, [2]decimal.Decimal{lo, hi})
}

func enIntervalos(x decimal.Decimal, is [][2]decimal.Decimal) bool {
	for _, i := range is {
		if x.Cmp(i[0]) >= 0 && x.Cmp(i[1]) <= 0 {
			return true
		}
	}
	return false
}

func describirIntervalos(is [][2]decimal.Decimal) string {
	partes := make([]string, len(is))
	for i, x := range is {
		partes[i] = fmt.Sprintf("[%s, %s]", x[0], x[1])
	}
	return strings.Join(partes, " o ")
}

// cceValorImporte revisa CCE157/158 (la suma de ValorDolares de cada
// NoIdentificacion contra el importe de sus conceptos en dólares) y
// CCE163/164 (muestras contra el descuento).
func (v *validadorReglas) cceValorImporte(ce *ComercioExterior, conceptos map[string][]Concepto) {
	r := "/Comprobante/Complemento/ComercioExterior/Mercancias"
	type acumulado struct {
		primera                 int
		vd, muestras            decimal.Decimal
		completo, ceroUno, hayM bool
	}
	por := map[string]*acumulado{}
	var orden []string
	for i, m := range ce.Mercancias {
		a := por[m.NoIdentificacion]
		if a == nil {
			a = &acumulado{primera: i, completo: true}
			por[m.NoIdentificacion] = a
			orden = append(orden, m.NoIdentificacion)
		}
		if m.ValorDolares == nil {
			a.completo = false
			continue
		}
		a.vd = a.vd.Add(*m.ValorDolares)
		a.ceroUno = a.ceroUno || m.ValorDolares.IsZero() || m.ValorDolares.Equal(uno)
		if clave(m.FraccionArancelaria) == "9801000100" {
			a.hayM = true
			a.muestras = a.muestras.Add(*m.ValorDolares)
		}
	}
	for _, k := range orden {
		a, cs := por[k], conceptos[k]
		if !a.completo || len(cs) == 0 {
			continue
		}
		var importe, descuento decimal.Decimal
		todos := true
		for _, cp := range cs {
			if cp.Importe == nil {
				todos = false
			}
			importe = importe.Add(valorDe(cp.Importe))
			descuento = descuento.Add(valorDe(cp.Descuento))
		}
		rm := idx(r, "Mercancia", a.primera)
		if todos && !a.ceroUno {
			if is := v.intervalosUSD(importe); is != nil && !enIntervalos(a.vd, is) {
				v.hallazgo(ReglaCCEValorImporte, rm, "ValorDolares", nil,
					fmt.Sprintf("los ValorDolares de NoIdentificacion %s suman %s y el Importe de sus conceptos, %s, en dólares admite %s", k, a.vd, importe, describirIntervalos(is)))
			}
		}
		if a.hayM {
			is := v.intervalosUSD(descuento)
			if is == nil {
				continue
			}
			max := is[0][1]
			for _, x := range is[1:] {
				if x[1].Cmp(max) > 0 {
					max = x[1]
				}
			}
			if a.muestras.Cmp(max) > 0 {
				v.hallazgo(ReglaCCEMuestras, rm, "ValorDolares", nil,
					fmt.Sprintf("las muestras (9801000100) de NoIdentificacion %s suman %s dólares y el Descuento de sus conceptos, %s, a lo más %s", k, a.muestras, descuento, max))
			}
		}
	}
}
