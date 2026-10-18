package cfdi

import (
	"sort"
	"strings"
	"testing"

	"github.com/OrtaMarco/cfdi-go/decimal"
)

// Reglas de presencia y coherencia de Comercio Exterior 2.0 (matriz CCE
// 2.0, CCE102-172) que llegaron en T8: cada una con casos que la disparan y
// casos que no. cce20.xml es un ingreso en USD (TipoCambio y TipoCambioUSD
// 18.5) con dos mercancías; trasladoCCE lo convierte en un traslado.
func TestValidarReglasCCE(t *testing.T) {
	cce := string(leerFixture(t, "testdata/sinteticos/cce20.xml"))
	ingreso := string(leerFixture(t, "testdata/sinteticos/ingreso40-pue.xml"))
	const (
		nodo     = `<cce20:ComercioExterior Version="2.0"`
		merc2    = `<cce20:Mercancia NoIdentificacion="EXP-002" FraccionArancelaria="9403600100" CantidadAduana="1" UnidadAduana="06" ValorUnitarioAduana="250.500000" ValorDolares="250.5000"/>`
		aduana1  = ` CantidadAduana="2" UnidadAduana="06" ValorUnitarioAduana="500.000000"`
		aduana2  = ` CantidadAduana="1" UnidadAduana="06" ValorUnitarioAduana="250.500000"`
		concepto = `NoIdentificacion="EXP-002" Cantidad="1" ClaveUnidad="H87"`
		fin2     = `Importe="250.50" ObjetoImp="01"/>`
		receptor = `<cce20:Receptor NumRegIdTrib="121585958">`
		dest     = `</cce20:Destinatario>`
		propiet  = `</cce20:Emisor>`
		relacion = `<cfdi:CfdiRelacionados TipoRelacion="05"><cfdi:CfdiRelacionado UUID="11111111-2222-4333-8444-555555555503"/></cfdi:CfdiRelacionados><cfdi:Emisor`
	)
	// trasladoCCE es cce20.xml como traslado (T): SubTotal y Total 0, sin
	// forma ni método de pago, Receptor igual al Emisor y MotivoTraslado.
	trasladoCCE := aplicarCambios(t, cce, []cambio{
		{`FormaPago="03" `, ``}, {` SubTotal="1250.50"`, ` SubTotal="0"`}, {` Total="1250.50"`, ` Total="0"`},
		{`TipoDeComprobante="I"`, `TipoDeComprobante="T"`}, {` MetodoPago="PUE"`, ``},
		{`<cfdi:Receptor Rfc="XEXX010101000"`, `<cfdi:Receptor Rfc="EKU9003173C9"`}, {receptor, `<cce20:Receptor>`},
		{nodo, nodo + ` MotivoTraslado="02"`},
	})
	sinAduana := []cambio{{aduana1, ``}, {aduana2, ``}}
	conVD := func(vd string) string {
		return strings.Replace(merc2, `ValorDolares="250.5000"`, `ValorDolares="`+vd+`"`, 1)
	}
	casos := []struct {
		nombre string
		base   string
		cs     []cambio
		quiero []string
	}{
		{"base de traslado", trasladoCCE, nil, nil},
		// Exportacion e Incoterm (CCE104, CCE105, CCE120).
		{"Exportacion 02 sin el complemento", ingreso, []cambio{{`Exportacion="01"`, `Exportacion="02"`}}, []string{ReglaCCEExportacion}},
		{"Exportacion 03 con el complemento", cce, []cambio{{`Exportacion="02"`, `Exportacion="03"`}}, []string{ReglaCCEExportacion}},
		{"Exportacion 04 sin Incoterm", cce, []cambio{{`Exportacion="02"`, `Exportacion="04"`}, {` Incoterm="FCA"`, ``}}, nil},
		{"Exportacion 02 sin Incoterm", cce, []cambio{{` Incoterm="FCA"`, ``}}, []string{ReglaCCEExportacion}},
		// Nodos del CFDI (CCE106, CCE110, CCE111, CCE112, CCE115).
		{"InformacionGlobal", cce, []cambio{{`<cfdi:Emisor`, `<cfdi:InformacionGlobal Periodicidad="01" Meses="01" Año="2026"/><cfdi:Emisor`}}, []string{ReglaCCENodosCFDI}},
		{"ACuentaTerceros", cce, []cambio{{fin2, `Importe="250.50" ObjetoImp="01"><cfdi:ACuentaTerceros RfcACuentaTerceros="XOJI740919U48" NombreACuentaTerceros="TERCERO DE PRUEBA" RegimenFiscalACuentaTerceros="612" DomicilioFiscalACuentaTerceros="76028"/></cfdi:Concepto>`}}, []string{ReglaCCENodosCFDI}},
		{"InformacionAduanera", cce, []cambio{{fin2, `Importe="250.50" ObjetoImp="01"><cfdi:InformacionAduanera NumeroPedimento="21  47  3807  8003832"/></cfdi:Concepto>`}}, []string{ReglaCCENodosCFDI + "/" + NivelAdvertencia}},
		{"CuentaPredial", cce, []cambio{{fin2, `Importe="250.50" ObjetoImp="01"><cfdi:CuentaPredial Numero="CP-1"/></cfdi:Concepto>`}}, []string{ReglaCCENodosCFDI}},
		{"con Divisas", cce, []cambio{{`<tfd:`, `<divisas:Divisas xmlns:divisas="http://www.sat.gob.mx/divisas" version="1.0" tipoOperacion="venta"/><tfd:`}}, []string{ReglaCCENodosCFDI}},
		{"traslado con ObjetoImp 02", trasladoCCE, []cambio{{fin2, `Importe="250.50" ObjetoImp="02"/>`}}, []string{ReglaCCENodosCFDI, ReglaObjetoImp}},
		// MotivoTraslado y Propietario (CCE102, CCE103, CCE116, CCE117).
		{"ingreso con MotivoTraslado", cce, []cambio{{nodo, nodo + ` MotivoTraslado="02"`}}, []string{ReglaCCEMotivoTraslado}},
		{"traslado sin MotivoTraslado", trasladoCCE, []cambio{{` MotivoTraslado="02"`, ``}}, []string{ReglaCCEMotivoTraslado}},
		{"Propietario sin motivo 05", cce, []cambio{{propiet, propiet + `<cce20:Propietario NumRegIdTrib="123456789" ResidenciaFiscal="USA"/>`}}, []string{ReglaCCEMotivoTraslado}},
		{"motivo 05 sin Propietario", trasladoCCE, []cambio{{` MotivoTraslado="02"`, ` MotivoTraslado="05"`}}, []string{ReglaCCEMotivoTraslado}},
		{"motivo 05 con Propietario", trasladoCCE, []cambio{{` MotivoTraslado="02"`, ` MotivoTraslado="05"`}, {propiet, propiet + `<cce20:Propietario NumRegIdTrib="123456789" ResidenciaFiscal="USA"/>`}}, nil},
		{"motivo 01 sin relacionados", trasladoCCE, []cambio{{` MotivoTraslado="02"`, ` MotivoTraslado="01"`}}, []string{ReglaCCEMotivoTraslado}},
		{"motivo 01 con relacionados", trasladoCCE, []cambio{{` MotivoTraslado="02"`, ` MotivoTraslado="01"`}, {`<cfdi:Emisor`, relacion}}, nil},
		// Traslados (CCE108, CCE142, CCE154-156).
		{"traslado a otro Rfc", trasladoCCE, []cambio{{`<cfdi:Receptor Rfc="EKU9003173C9"`, `<cfdi:Receptor Rfc="XOJI740919U48"`}}, []string{ReglaCCETraslado}},
		{"traslado con dos destinatarios", trasladoCCE, []cambio{{dest, dest + `<cce20:Destinatario Nombre="OTRO DESTINATARIO"><cce20:Domicilio Calle="Third Rd" Estado="TX" Pais="USA" CodigoPostal="78042"/></cce20:Destinatario>`}}, []string{ReglaCCETraslado}},
		{"ingreso con dos destinatarios", cce, []cambio{{dest, dest + `<cce20:Destinatario Nombre="OTRO DESTINATARIO"><cce20:Domicilio Calle="Third Rd" Estado="TX" Pais="USA" CodigoPostal="78042"/></cce20:Destinatario>`}}, nil},
		{"traslado sin aduana y sin Unidad", trasladoCCE, sinAduana, []string{ReglaCCETraslado}},
		{"traslado sin aduana con Unidad de aduana", trasladoCCE, append(sinAduana, cambio{`ClaveUnidad="H87" Descripcion="Equipo`, `ClaveUnidad="H87" Unidad="06" Descripcion="Equipo`}, cambio{`ClaveUnidad="H87" Descripcion="Mueble`, `ClaveUnidad="H87" Unidad="06" Descripcion="Mueble`}), nil},
		{"traslado sin aduana con Unidad libre", trasladoCCE, append(sinAduana, cambio{`ClaveUnidad="H87" Descripcion="Equipo`, `ClaveUnidad="H87" Unidad="Pieza" Descripcion="Equipo`}, cambio{`ClaveUnidad="H87" Descripcion="Mueble`, `ClaveUnidad="H87" Unidad="06" Descripcion="Mueble`}), []string{ReglaCCETraslado}},
		{"traslado sin aduana con Cantidad de 4 decimales", trasladoCCE, append(sinAduana, cambio{`ClaveUnidad="H87" Descripcion="Equipo`, `ClaveUnidad="H87" Unidad="06" Descripcion="Equipo`}, cambio{`ClaveUnidad="H87" Descripcion="Mueble`, `ClaveUnidad="H87" Unidad="06" Descripcion="Mueble`},
			cambio{concepto, `NoIdentificacion="EXP-002" Cantidad="1.0000" ClaveUnidad="H87"`}), []string{ReglaCCETraslado}},
		{"traslado sin aduana con ValorUnitario de 3 decimales", trasladoCCE, append(sinAduana, cambio{`ClaveUnidad="H87" Descripcion="Equipo`, `ClaveUnidad="H87" Unidad="06" Descripcion="Equipo`}, cambio{`ClaveUnidad="H87" Descripcion="Mueble`, `ClaveUnidad="H87" Unidad="06" Descripcion="Mueble`},
			cambio{`ValorUnitario="250.50" Importe="250.50"`, `ValorUnitario="250.500" Importe="250.50"`}), []string{ReglaCCETraslado}},
		// Emisor (CCE123-125).
		{"Curp de persona moral", cce, []cambio{{`<cce20:Emisor>`, `<cce20:Emisor Curp="XEXX010101HNEXXXA4">`}}, []string{ReglaCCEEmisor}},
		{"persona física sin Curp", cce, []cambio{{`<cfdi:Emisor Rfc="EKU9003173C9"`, `<cfdi:Emisor Rfc="XOJI740919U48"`}}, []string{ReglaCCEEmisor}},
		{"persona física con Curp", cce, []cambio{{`<cfdi:Emisor Rfc="EKU9003173C9"`, `<cfdi:Emisor Rfc="XOJI740919U48"`}, {`<cce20:Emisor>`, `<cce20:Emisor Curp="XEXX010101HNEXXXA4">`}}, nil},
		{"emisor fuera de México", cce, []cambio{{`Pais="MEX"`, `Pais="USA"`}}, []string{ReglaCCEEmisor}},
		// NumRegIdTrib (CCE133, CCE134, CCE143).
		{"receptor genérico sin NumRegIdTrib", cce, []cambio{{receptor, `<cce20:Receptor>`}}, []string{ReglaCCENumRegIdTrib}},
		{"receptor no genérico con NumRegIdTrib", cce, []cambio{{`<cfdi:Receptor Rfc="XEXX010101000"`, `<cfdi:Receptor Rfc="XOJI740919U48"`}}, []string{ReglaCCENumRegIdTrib}},
		{"receptor con NumRegIdTrib en México", cce, []cambio{{`Estado="TX" Pais="USA" CodigoPostal="78040"`, `Estado="TX" Pais="MEX" CodigoPostal="78040"`}}, []string{ReglaCCENumRegIdTrib}},
		{"destinatario con NumRegIdTrib en México", cce, []cambio{{`Estado="CA" Pais="USA"`, `Estado="CA" Pais="MEX"`}}, []string{ReglaCCENumRegIdTrib}},
		{"destinatario sin NumRegIdTrib en México", cce, []cambio{{`Estado="CA" Pais="USA"`, `Estado="CA" Pais="MEX"`}, {`<cce20:Destinatario NumRegIdTrib="987654321"`, `<cce20:Destinatario`}}, nil},
		// FraccionArancelaria (CCE159, CCE160, CCE170).
		{"servicio con fracción", cce, []cambio{{aduana1 + ` ValorDolares="1000.0000"`, ` CantidadAduana="2" UnidadAduana="99" ValorUnitarioAduana="0" ValorDolares="1"`}, {`TotalUSD="1250.5000"`, `TotalUSD="251.50"`}}, []string{ReglaCCEFraccion}},
		{"concepto E48 con fracción", cce, []cambio{{concepto, `NoIdentificacion="EXP-002" Cantidad="1" ClaveUnidad="E48"`}}, []string{ReglaCCEFraccion}},
		{"mercancía sin fracción", cce, []cambio{{` FraccionArancelaria="9403600100"`, ``}}, []string{ReglaCCEFraccion}},
		{"servicio E48 sin fracción", cce, []cambio{{concepto, `NoIdentificacion="EXP-002" Cantidad="1" ClaveUnidad="E48"`}, {` FraccionArancelaria="9403600100"`, ``}}, nil},
		// Datos de aduana en todas (CCE165, CCE166).
		{"aduana en una sola mercancía", cce, []cambio{{aduana2, ``}}, []string{ReglaCCEDatosAduana}},
		{"sin aduana en ninguna", cce, sinAduana, nil},
		{"NoIdentificacion repetido sin aduana", cce, append(sinAduana, cambio{`NoIdentificacion="EXP-002" FraccionArancelaria`, `NoIdentificacion="EXP-001" FraccionArancelaria`}), []string{ReglaCCEDatosAduana, ReglaCCEValorImporte}},
		// ValorDolares contra el Importe (CCE157, CCE158).
		// Importe 250.50 USD: [250.49, 250.51] (± ½ centavo, guía p. 58-59, y
		// con TipoCambio = TipoCambioUSD = 18.5000, lo mismo).
		{"ValorDolares de otro importe", cce, []cambio{{merc2, conVD("250.5200")}, {`TotalUSD="1250.5000"`, `TotalUSD="1250.52"`}}, []string{ReglaCCEValorImporte}},
		{"ValorDolares en el límite del importe", cce, []cambio{{merc2, conVD("250.5100")}, {`TotalUSD="1250.5000"`, `TotalUSD="1250.51"`}}, nil},
		{"importe en pesos", cce, []cambio{{`Moneda="USD" TipoCambio="18.5000"`, `Moneda="MXN"`}}, []string{ReglaCCEValorImporte}},
		{"ValorDolares 0 no se compara con el importe", cce, []cambio{{merc2, conVD("0")}, {`TotalUSD="1250.5000"`, `TotalUSD="1000.00"`}}, nil},
		// Muestras (CCE163, CCE164).
		{"muestras sin descuento", cce, []cambio{{`FraccionArancelaria="9403600100"`, `FraccionArancelaria="9801000100"`}}, []string{ReglaCCEMuestras}},
		{"muestras con descuento", cce, []cambio{{`FraccionArancelaria="9403600100"`, `FraccionArancelaria="9801000100"`}, {fin2, `Importe="250.50" Descuento="250.50" ObjetoImp="01"/>`},
			{` SubTotal="1250.50"`, ` SubTotal="1250.50" Descuento="250.50"`}, {` Total="1250.50"`, ` Total="1000.00"`}}, nil},
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
			if (h.Nivel != r.Nivel && !((r.Margen != "" || r.SoloGuia != "") && h.Nivel == NivelAdvertencia)) || h.Regla == "" || h.Mensaje == "" || !strings.HasPrefix(h.Ruta, "/Comprobante") {
				t.Errorf("%s: hallazgo incompleto %+v", c.nombre, h)
			}
		}
	}
}

// La conversión del Importe a dólares del estándar de Comercio Exterior
// 2.0 (p. 42), con los ejemplos de la guía: el escenario 4 (p. 77-78, EUR)
// da ValorDolares 1022424.8112 y el escenario 2 (p. 73-74, USD con
// TipoCambio pactado 18 y TipoCambioUSD 19.1202) pide ValorDolares igual al
// Importe.
func TestIntervalosUSD(t *testing.T) {
	casos := []struct {
		moneda, tc, usd, importe string
		dentro, fuera            []string
	}{
		{"EUR", "23.09903", "19.1202", "846311.16", []string{"1022424.8112", "1022421.91", "1022427.72"}, []string{"1022421.90", "1022427.73"}},
		{"MXN", "", "18.5000", "1000.00", []string{"54.05", "54.06"}, []string{"54.04", "54.07"}},
		{"USD", "18.00000", "19.1202", "1022424.81", []string{"1022424.81", "1022424.82"}, []string{"1022424.79", "1022424.83"}},
		{"USD", "18.00000", "19.1202", "1000.00", []string{"1000.00", "941.40", "941.43"}, []string{"941.39", "999.98"}},
	}
	for _, c := range casos {
		comp := &Comprobante{Version: "4.0", Moneda: c.moneda, ComercioExterior: &ComercioExterior{TipoCambioUSD: ptrDec(c.usd)}}
		if c.tc != "" {
			comp.TipoCambio = ptrDec(c.tc)
		}
		v := &validadorReglas{c: comp, moneda: c.moneda}
		is := v.intervalosUSD(decimal.MustParse(c.importe))
		for _, x := range c.dentro {
			if !enIntervalos(decimal.MustParse(x), is) {
				t.Errorf("%s %s: %s fuera de %s", c.moneda, c.importe, x, describirIntervalos(is))
			}
		}
		for _, x := range c.fuera {
			if enIntervalos(decimal.MustParse(x), is) {
				t.Errorf("%s %s: %s dentro de %s", c.moneda, c.importe, x, describirIntervalos(is))
			}
		}
	}
	// Sin TipoCambioUSD no hay conversión.
	v := &validadorReglas{c: &Comprobante{Moneda: "MXN", ComercioExterior: &ComercioExterior{}}, moneda: "MXN"}
	if is := v.intervalosUSD(decimal.MustParse("1")); is != nil {
		t.Errorf("sin TipoCambioUSD: %v", is)
	}
}
