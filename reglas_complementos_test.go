package cfdi

import (
	"sort"
	"strings"
	"testing"
)

// Reglas de Comercio Exterior 2.0 (matriz CCE 2.0) y Nómina 1.2 (matriz Rev
// D): cada una con un caso que la dispara y uno que no.
func TestValidarReglasComplementos(t *testing.T) {
	cce := string(leerFixture(t, "testdata/sinteticos/cce20.xml"))
	nomina := string(leerFixture(t, "testdata/sinteticos/nomina12.xml"))
	const (
		merc1  = `CantidadAduana="2" UnidadAduana="06" ValorUnitarioAduana="500.000000" ValorDolares="1000.0000"`
		merc2  = `<cce20:Mercancia NoIdentificacion="EXP-002" FraccionArancelaria="9403600100"`
		perc   = `<nomina12:Percepciones TotalSueldos="5600.00" TotalGravado="5300.00" TotalExento="300.00">`
		sueldo = `ImporteGravado="5000.00" ImporteExento="0.00"/>`
		otro   = `TipoOtroPago="002" Clave="O002" Concepto="Subsidio para el empleo" Importe="10.00"`
	)
	// conSueldo cambia el sueldo y todo lo que lo suma.
	conSueldo := func(g, total, gravado string) []cambio {
		return []cambio{
			{sueldo, `ImporteGravado="` + g + `" ImporteExento="0.00"/>`},
			{perc, `<nomina12:Percepciones TotalSueldos="` + total + `" TotalGravado="` + gravado + `" TotalExento="300.00">`},
		}
	}
	casos := []struct {
		nombre string
		base   string
		cs     []cambio
		quiero []string
	}{
		// Comercio Exterior 2.0.
		{"NumCertificadoOrigen sin certificado", cce, []cambio{{`CertificadoOrigen="0"`, `CertificadoOrigen="0" NumCertificadoOrigen="CO-1"`}}, []string{ReglaCCECertificado}},
		{"NumCertificadoOrigen con certificado", cce, []cambio{{`CertificadoOrigen="0"`, `CertificadoOrigen="1" NumCertificadoOrigen="CO-1"`}}, nil},
		{"TotalUSD distinto de la suma", cce, []cambio{{`TotalUSD="1250.5000"`, `TotalUSD="1250.5100"`}}, []string{ReglaCCETotalUSD}},
		{"TotalUSD con dos decimales", cce, []cambio{{`TotalUSD="1250.5000"`, `TotalUSD="1250.50"`}}, nil},
		{"concepto sin NoIdentificacion", cce, []cambio{{` NoIdentificacion="EXP-002" Cantidad="1"`, ` Cantidad="1"`}}, []string{ReglaCCENoIdentificacion}},
		{"mercancía sin concepto", cce, []cambio{{merc2, `<cce20:Mercancia NoIdentificacion="EXP-009" FraccionArancelaria="9403600100"`}}, []string{ReglaCCENoIdentificacion}},
		{"mercancía repetida", cce, []cambio{{merc2, `<cce20:Mercancia NoIdentificacion="EXP-001" FraccionArancelaria="8471300100"`}}, []string{ReglaCCENoIdentificacion, ReglaCCEValorImporte}}, // EXP-001 suma 1250.50 y su concepto 1000.00
		{"datos de aduana incompletos", cce, []cambio{{merc1, `CantidadAduana="2" ValorUnitarioAduana="500.000000" ValorDolares="1000.0000"`}}, []string{ReglaCCEDatosAduana}},
		// Con ValorUnitarioAduana 0 el producto es 0, pero ValorDolares sí
		// es el Importe del concepto en dólares: advertencia.
		{"valor unitario de aduana cero", cce, []cambio{{merc1, `CantidadAduana="2" UnidadAduana="06" ValorUnitarioAduana="0" ValorDolares="1000.0000"`}}, []string{ReglaCCEDatosAduana, adv(ReglaCCEValorDolares)}},
		{"valor unitario de aduana cero en servicios", cce, []cambio{{`FraccionArancelaria="8471300100" ` + merc1, `CantidadAduana="2" UnidadAduana="99" ValorUnitarioAduana="0" ValorDolares="1"`}, {`TotalUSD="1250.5000"`, `TotalUSD="251.50"`}}, nil},
		// Límites de CantidadAduana × ValorUnitarioAduana (estándar p. 44)
		// con Cantidad 2 y ValorUnitario 500.00 del concepto: ±0.5 y
		// ±0.005. Con ValorUnitarioAduana 399.995 el superior es (2.5 −
		// 10^-12)(400 − 10^-12) → 1000.00; con 399.99, 999.9875 → 999.99.
		// Con 666.671666 el inferior es 1.5 × 666.666666 → 999.99; con
		// 666.68, 1000.0125 → 1000.01. ValorDolares 1000 sigue siendo el
		// Importe del concepto (USD): fuera de los límites es advertencia.
		{"ValorDolares en el límite superior", cce, []cambio{{merc1, `CantidadAduana="2" UnidadAduana="06" ValorUnitarioAduana="399.995000" ValorDolares="1000.0000"`}}, nil},
		{"ValorDolares sobre el límite superior", cce, []cambio{{merc1, `CantidadAduana="2" UnidadAduana="06" ValorUnitarioAduana="399.990000" ValorDolares="1000.0000"`}}, []string{adv(ReglaCCEValorDolares)}},
		{"ValorDolares en el límite inferior", cce, []cambio{{merc1, `CantidadAduana="2" UnidadAduana="06" ValorUnitarioAduana="666.671666" ValorDolares="1000.0000"`}}, nil},
		{"ValorDolares bajo el límite inferior", cce, []cambio{{merc1, `CantidadAduana="2" UnidadAduana="06" ValorUnitarioAduana="666.680000" ValorDolares="1000.0000"`}}, []string{adv(ReglaCCEValorDolares)}},
		// Fuera de [749.99, 1250.02] y del Importe: error en las dos.
		{"ValorDolares fuera de los dos", cce, []cambio{{merc1, `CantidadAduana="2" UnidadAduana="06" ValorUnitarioAduana="500.000000" ValorDolares="1250.0300"`}, {`TotalUSD="1250.5000"`, `TotalUSD="1500.53"`}}, []string{ReglaCCEValorDolares, ReglaCCEValorImporte}},
		{"ValorDolares en el límite pero no es el Importe", cce, []cambio{{merc1, `CantidadAduana="2" UnidadAduana="06" ValorUnitarioAduana="500.000000" ValorDolares="1250.0200"`}, {`TotalUSD="1250.5000"`, `TotalUSD="1500.52"`}}, []string{ReglaCCEValorImporte}},
		{"ValorDolares 1 no se revisa", cce, []cambio{{merc1, `CantidadAduana="2" UnidadAduana="06" ValorUnitarioAduana="500.000000" ValorDolares="1"`}, {`TotalUSD="1250.5000"`, `TotalUSD="251.50"`}}, nil},
		// Nómina 1.2.
		{"TotalPercepciones", nomina, []cambio{{`TotalPercepciones="5600.00"`, `TotalPercepciones="5600.01"`}}, []string{ReglaNominaTotales, ReglaNominaConcepto}},
		{"TotalDeducciones", nomina, []cambio{{`TotalDeducciones="575.50"`, `TotalDeducciones="575.51"`}}, []string{ReglaNominaTotales, ReglaNominaConcepto}},
		{"TotalOtrosPagos", nomina, []cambio{{`TotalOtrosPagos="10.00"`, `TotalOtrosPagos="10.01"`}}, []string{ReglaNominaTotales, ReglaNominaConcepto}},
		{"sin TotalOtrosPagos", nomina, []cambio{{` TotalOtrosPagos="10.00"`, ``}}, []string{ReglaNominaTotales, ReglaNominaConcepto}},
		{"TotalSueldos", nomina, conSueldo("5000.01", "5600.00", "5300.01"), []string{ReglaNominaPercepciones}},
		{"TotalGravado", nomina, []cambio{{perc, `<nomina12:Percepciones TotalSueldos="5600.00" TotalGravado="5300.01" TotalExento="299.99">`}}, []string{ReglaNominaPercepciones}},
		{"separación sin su total", nomina, []cambio{{`TipoPercepcion="001"`, `TipoPercepcion="022"`}}, []string{ReglaNominaPercepciones, ReglaNominaNodosPercepcion}}, // y sin SeparacionIndemnizacion (NOM78)
		{"percepción en cero", nomina, append(conSueldo("0.00", "600.00", "300.00"),
			cambio{`TotalPercepciones="5600.00"`, `TotalPercepciones="600.00"`},
			cambio{`ValorUnitario="5610.00" Importe="5610.00"`, `ValorUnitario="610.00" Importe="610.00"`},
			cambio{`SubTotal="5610.00"`, `SubTotal="610.00"`}, cambio{`Total="5034.50"`, `Total="34.50"`}), []string{ReglaNominaImportes}},
		{"impuestos retenidos", nomina, []cambio{{`Concepto="ISR" Importe="450.00"`, `Concepto="ISR" Importe="449.99"`}}, []string{ReglaNominaDeducciones}},
		{"impuestos retenidos sin deducción 002", nomina, []cambio{{`TipoDeduccion="002"`, `TipoDeduccion="004"`}}, []string{ReglaNominaDeducciones}},
		{"deducción en cero", nomina, []cambio{{`Concepto="Seguridad social" Importe="125.50"`, `Concepto="Seguridad social" Importe="0.00"`}}, []string{ReglaNominaImportes}},
		{"otro pago en cero", nomina, []cambio{{otro, `TipoOtroPago="999" Clave="O002" Concepto="Subsidio para el empleo" Importe="0.00"`}, {`TotalOtrosPagos="10.00"`, `TotalOtrosPagos="0.00"`},
			{`ValorUnitario="5610.00" Importe="5610.00"`, `ValorUnitario="5600.00" Importe="5600.00"`}, {`SubTotal="5610.00"`, `SubTotal="5600.00"`}, {`Total="5034.50"`, `Total="5024.50"`}}, []string{ReglaNominaImportes, ReglaNominaTipoRegimen}}, // TipoRegimen 02 sin OtroPago 002 (NOM105)
		{"subsidio mayor que el causado", nomina, []cambio{{`SubsidioCausado="10.00"`, `SubsidioCausado="9.99"`}}, []string{ReglaNominaOtrosPagos}},
		{"saldo a favor menor que el remanente", nomina, []cambio{{`<nomina12:SubsidioAlEmpleo SubsidioCausado="10.00"/>`,
			`<nomina12:SubsidioAlEmpleo SubsidioCausado="10.00"/><nomina12:CompensacionSaldosAFavor SaldoAFavor="100.00" Año="2025" RemanenteSalFav="100.01"/>`}}, []string{ReglaNominaOtrosPagos}},
		{"saldo a favor igual al remanente", nomina, []cambio{{`<nomina12:SubsidioAlEmpleo SubsidioCausado="10.00"/>`,
			`<nomina12:SubsidioAlEmpleo SubsidioCausado="10.00"/><nomina12:CompensacionSaldosAFavor SaldoAFavor="100.00" Año="2025" RemanenteSalFav="100.00"/>`}}, nil},
		{"concepto distinto de las percepciones", nomina, []cambio{{`ValorUnitario="5610.00" Importe="5610.00"`, `ValorUnitario="5610.01" Importe="5610.01"`},
			{`SubTotal="5610.00"`, `SubTotal="5610.01"`}, {`Total="5034.50"`, `Total="5034.51"`}}, []string{ReglaNominaConcepto}},
		{"concepto sin descuento", nomina, []cambio{{` Descuento="575.50" ObjetoImp="01"`, ` ObjetoImp="01"`}, {` Descuento="575.50" Moneda`, ` Moneda`}, {`Total="5034.50"`, `Total="5610.00"`}}, []string{ReglaNominaConcepto}},
		{"fechas al revés", nomina, []cambio{{`FechaInicialPago="2026-01-01"`, `FechaInicialPago="2026-01-16"`}}, []string{ReglaNominaFechas}},
		{"fechas iguales", nomina, []cambio{{`FechaInicialPago="2026-01-01"`, `FechaInicialPago="2026-01-15"`}}, nil},
		{"subcontratación que no suma 100", nomina, []cambio{{`ClaveEntFed="QUE"/>`, `ClaveEntFed="QUE"><nomina12:SubContratacion RfcLabora="EKU9003173C9" PorcentajeTiempo="60"/><nomina12:SubContratacion RfcLabora="XOJI740919U48" PorcentajeTiempo="30.000"/></nomina12:Receptor>`}}, []string{ReglaNominaSubContratacion}},
		{"subcontratación que suma 100", nomina, []cambio{{`ClaveEntFed="QUE"/>`, `ClaveEntFed="QUE"><nomina12:SubContratacion RfcLabora="EKU9003173C9" PorcentajeTiempo="60"/><nomina12:SubContratacion RfcLabora="XOJI740919U48" PorcentajeTiempo="40.000"/></nomina12:Receptor>`}}, nil},
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
			if (h.Nivel != r.Nivel && !(r.Margen != "" && h.Nivel == NivelAdvertencia)) || h.Regla == "" || h.Mensaje == "" || !strings.HasPrefix(h.Ruta, "/Comprobante") {
				t.Errorf("%s: hallazgo incompleto %+v", c.nombre, h)
			}
		}
	}
	// TipoDeComprobante (CCE101), solo en Comercio Exterior 2.0 y CFDI 4.0.
	tipos := []struct {
		version, cce, tipo string
		mal                bool
	}{{"4.0", "2.0", "P", true}, {"4.0", "2.0", "N", true}, {"4.0", "2.0", "T", false}, {"4.0", "2.0", "E", false}, {"4.0", "1.1", "P", false}, {"3.3", "2.0", "P", false}}
	for _, x := range tipos {
		c := &Comprobante{Version: x.version, TipoDeComprobante: x.tipo, ComercioExterior: &ComercioExterior{Version: x.cce}}
		if hs := ValidarReglas(c); TieneHallazgo(hs, ReglaCCETipoComprobante) != x.mal {
			t.Errorf("CFDI %s, CCE %s, tipo %s: %+v", x.version, x.cce, x.tipo, hs)
		}
	}
	// Con dos complementos de nómina cada uno lleva su índice en la ruta y
	// el concepto no se compara.
	n := Nomina{Version: "1.2", TotalPercepciones: ptrDec("1"), Percepciones: &Percepciones{TotalSueldos: ptrDec("2")}}
	hs := ValidarReglas(&Comprobante{Version: "4.0", TipoDeComprobante: "N", Nomina: []Nomina{n, n}, Conceptos: []Concepto{{Importe: ptrDec("9")}}})
	if len(hs) != 4 || hs[0].Ruta != "/Comprobante/Complemento/Nomina[1]/@TotalPercepciones" || TieneHallazgo(hs, ReglaNominaConcepto) {
		t.Errorf("dos nóminas: %+v", hs)
	}
}
