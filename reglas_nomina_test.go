package cfdi

import (
	"sort"
	"strings"
	"testing"
	"time"
)

// Reglas de Nómina 1.2 de T8 (matriz Rev D: NOM3-29, NOM31-106): cada una
// con casos que la disparan y casos que no. nomina12.xml es una nómina
// ordinaria 4.0 de un trabajador con TipoContrato 01, TipoRegimen 02,
// RegistroPatronal, Antigüedad P315W desde 2020-01-01 al 2026-01-15, sueldo
// (001), horas extra (019) y subsidio (002).
func TestValidarReglasNominaT8(t *testing.T) {
	nomina := string(leerFixture(t, "testdata/sinteticos/nomina12.xml"))
	const (
		emisor    = `<nomina12:Emisor RegistroPatronal="Z0000001234"/>`
		inicio    = `FechaInicioRelLaboral="2020-01-01" Antigüedad="P315W"`
		claveEnt  = `ClaveEntFed="QUE"/>`
		sueldo    = `<nomina12:Percepcion TipoPercepcion="001" Clave="P001" Concepto="Sueldo" ImporteGravado="5000.00" ImporteExento="0.00"/>`
		horas     = `<nomina12:HorasExtra Dias="2" TipoHoras="01" HorasExtra="4" ImportePagado="600.00"/>`
		percs     = `<nomina12:Percepciones TotalSueldos="5600.00" TotalGravado="5300.00" TotalExento="300.00">`
		finPercs  = `</nomina12:Percepciones>`
		seguridad = `<nomina12:Deduccion TipoDeduccion="001" Clave="D001" Concepto="Seguridad social" Importe="125.50"/>`
		incap     = `<nomina12:Incapacidad DiasIncapacidad="1" TipoIncapacidad="02"/>`
		subsidio  = `<nomina12:SubsidioAlEmpleo SubsidioCausado="10.00"/>`
		finOtros  = `</nomina12:OtroPago>`
		concepto  = `<cfdi:Concepto ClaveProdServ="84111505" Cantidad="1" ClaveUnidad="ACT" Descripcion="Pago de nómina"`
	)
	cuenta := func(attrs string) []cambio { return []cambio{{claveEnt, `ClaveEntFed="QUE" ` + attrs + `/>`}} }
	antiguedad := func(desde, a string) []cambio {
		return []cambio{{inicio, `FechaInicioRelLaboral="` + desde + `" Antigüedad="` + a + `"`}}
	}
	// jubilacion cambia el sueldo por una jubilación tipo (039 o 044) con
	// el nodo JubilacionPensionRetiro dado.
	jubilacion := func(tipo, nodo string) []cambio {
		return []cambio{
			{sueldo, strings.Replace(sueldo, `TipoPercepcion="001"`, `TipoPercepcion="`+tipo+`"`, 1)},
			{percs, `<nomina12:Percepciones TotalSueldos="600.00" TotalJubilacionPensionRetiro="5000.00" TotalGravado="5300.00" TotalExento="300.00">`},
			{finPercs, nodo + finPercs},
		}
	}
	// saldoAFavor agrega un OtroPago 004 de 5.00 con lo que se le pase
	// adentro.
	saldoAFavor := func(dentro string) []cambio {
		return []cambio{
			{finOtros, finOtros + `<nomina12:OtroPago TipoOtroPago="004" Clave="O004" Concepto="Saldo a favor" Importe="5.00">` + dentro + `</nomina12:OtroPago>`},
			{`TotalOtrosPagos="10.00"`, `TotalOtrosPagos="15.00"`}, {`ValorUnitario="5610.00" Importe="5610.00"`, `ValorUnitario="5615.00" Importe="5615.00"`},
			{`SubTotal="5610.00"`, `SubTotal="5615.00"`}, {`Total="5034.50"`, `Total="5039.50"`},
		}
	}
	// Un asimilado (TipoContrato 09, TipoRegimen 09): sin registro
	// patronal ni datos de seguridad social, y sin subsidio.
	asimilado := []cambio{
		{emisor, `<nomina12:Emisor/>`},
		{`<nomina12:Receptor Curp="XEXX010101HNEXXXA4" NumSeguridadSocial="00000000000" FechaInicioRelLaboral="2020-01-01" Antigüedad="P315W" TipoContrato="01" Sindicalizado="No" TipoJornada="01" TipoRegimen="02" NumEmpleado="E-001" Departamento="Pruebas" Puesto="Analista" RiesgoPuesto="1" PeriodicidadPago="04" SalarioBaseCotApor="400.00" SalarioDiarioIntegrado="380.50" ClaveEntFed="QUE"/>`,
			`<nomina12:Receptor Curp="XEXX010101HNEXXXA4" TipoContrato="09" TipoRegimen="09" NumEmpleado="E-001" PeriodicidadPago="04" ClaveEntFed="QUE"/>`},
		{`<nomina12:OtrosPagos>
        <nomina12:OtroPago TipoOtroPago="002" Clave="O002" Concepto="Subsidio para el empleo" Importe="10.00">
          ` + subsidio + `
        ` + finOtros + `
      </nomina12:OtrosPagos>`, ``},
		{` TotalOtrosPagos="10.00"`, ``}, {`ValorUnitario="5610.00" Importe="5610.00"`, `ValorUnitario="5600.00" Importe="5600.00"`},
		{`SubTotal="5610.00"`, `SubTotal="5600.00"`}, {`Total="5034.50"`, `Total="5024.50"`},
	}
	casos := []struct {
		nombre string
		cs     []cambio
		quiero []string
	}{
		// El CFDI de nómina 4.0 (NOM3-29).
		{"Exportacion 03", []cambio{{`Exportacion="01"`, `Exportacion="03"`}}, []string{ReglaNominaComprobante}},
		// Lo que solo pide la guía (p. 10, 16) es advertencia.
		{"MetodoPago PPD", []cambio{{`MetodoPago="PUE"`, `MetodoPago="PPD"`}}, []string{ReglaNominaComprobante + "/" + NivelAdvertencia}},
		{"InformacionGlobal", []cambio{{`<cfdi:Emisor`, `<cfdi:InformacionGlobal Periodicidad="04" Meses="01" Año="2026"/><cfdi:Emisor`}}, []string{ReglaNominaComprobante}},
		{"emisor persona física sin Curp", []cambio{{`<cfdi:Emisor Rfc="EKU9003173C9"`, `<cfdi:Emisor Rfc="CACX7605101P8"`}}, []string{ReglaNominaComprobante}},
		{"emisor persona física con Curp", []cambio{{`<cfdi:Emisor Rfc="EKU9003173C9"`, `<cfdi:Emisor Rfc="CACX7605101P8"`}, {emisor, `<nomina12:Emisor Curp="XEXX010101HNEXXXA4" RegistroPatronal="Z0000001234"/>`}}, nil},
		{"emisor persona moral con Curp", []cambio{{emisor, `<nomina12:Emisor Curp="XEXX010101HNEXXXA4" RegistroPatronal="Z0000001234"/>`}}, []string{ReglaNominaComprobante}},
		{"FacAtrAdquirente", []cambio{{`RegimenFiscal="601"/>`, `RegimenFiscal="601" FacAtrAdquirente="0123456789"/>`}}, []string{ReglaNominaComprobante}},
		{"receptor persona moral", []cambio{{`<cfdi:Receptor Rfc="XOJI740919U48"`, `<cfdi:Receptor Rfc="EKU9003173C9"`}}, []string{ReglaNominaComprobante}},
		{"régimen del receptor", []cambio{{`RegimenFiscalReceptor="605"`, `RegimenFiscalReceptor="612"`}}, []string{ReglaNominaComprobante}},
		{"UsoCFDI", []cambio{{`UsoCFDI="CN01"`, `UsoCFDI="G03"`}}, []string{ReglaNominaComprobante}},
		{"receptor con residencia fiscal", []cambio{{`UsoCFDI="CN01"`, `UsoCFDI="CN01" ResidenciaFiscal="USA"`}}, []string{ReglaNominaComprobante + "/" + NivelAdvertencia}},
		{"receptor con NumRegIdTrib", []cambio{{`UsoCFDI="CN01"`, `UsoCFDI="CN01" NumRegIdTrib="123456789"`}}, []string{ReglaNominaComprobante + "/" + NivelAdvertencia}},
		{"PPD y UsoCFDI G03", []cambio{{`MetodoPago="PUE"`, `MetodoPago="PPD"`}, {`UsoCFDI="CN01"`, `UsoCFDI="G03"`}}, []string{ReglaNominaComprobante + "/" + NivelAdvertencia, ReglaNominaComprobante}},
		{"ClaveProdServ", []cambio{{`ClaveProdServ="84111505"`, `ClaveProdServ="84111506"`}}, []string{ReglaNominaComprobante}},
		{"NoIdentificacion", []cambio{{`ClaveProdServ="84111505"`, `ClaveProdServ="84111505" NoIdentificacion="N-1"`}}, []string{ReglaNominaComprobante}},
		{"Cantidad con decimales", []cambio{{`Cantidad="1" ClaveUnidad="ACT"`, `Cantidad="1.00" ClaveUnidad="ACT"`}}, []string{ReglaNominaComprobante}},
		{"Unidad", []cambio{{`ClaveUnidad="ACT"`, `ClaveUnidad="ACT" Unidad="Actividad"`}}, []string{ReglaNominaComprobante}},
		{"Descripcion sin acento", []cambio{{`Descripcion="Pago de nómina"`, `Descripcion="Pago de nomina"`}}, []string{ReglaNominaComprobante}},
		{"ObjetoImp 02", []cambio{{`Descuento="575.50" ObjetoImp="01"`, `Descuento="575.50" ObjetoImp="02"`}}, []string{ReglaNominaComprobante, ReglaObjetoImp}},
		{"concepto con CuentaPredial", []cambio{{`Descuento="575.50" ObjetoImp="01"/>`, `Descuento="575.50" ObjetoImp="01"><cfdi:CuentaPredial Numero="CP-1"/></cfdi:Concepto>`}}, []string{ReglaNominaComprobante}},
		{"dos conceptos", []cambio{{concepto, concepto + ` ValorUnitario="0.01" Importe="0.01" ObjetoImp="01"/>` + concepto}, {`SubTotal="5610.00"`, `SubTotal="5610.01"`}, {`Total="5034.50"`, `Total="5034.51"`}}, []string{ReglaNominaComprobante}},
		// Totales (NOM36, NOM38).
		{"TotalDeducciones sin Deducciones", []cambio{{`<nomina12:Deducciones TotalOtrasDeducciones="125.50" TotalImpuestosRetenidos="450.00">
        <nomina12:Deduccion TipoDeduccion="002" Clave="D002" Concepto="ISR" Importe="450.00"/>
        ` + seguridad + `
      </nomina12:Deducciones>`, ``}}, []string{ReglaNominaTotales}},
		// Periodicidad (NOM33, NOM34).
		{"extraordinaria quincenal", []cambio{{`TipoNomina="O"`, `TipoNomina="E"`}}, []string{ReglaNominaPeriodicidad}},
		{"extraordinaria con 99", []cambio{{`TipoNomina="O"`, `TipoNomina="E"`}, {`PeriodicidadPago="04"`, `PeriodicidadPago="99"`}}, nil},
		{"ordinaria con 99", []cambio{{`PeriodicidadPago="04"`, `PeriodicidadPago="99"`}}, []string{ReglaNominaPeriodicidad}},
		// Registro patronal (NOM42-44).
		{"asimilado sin registro patronal", asimilado, nil},
		{"contrato 01 sin registro patronal con sus datos", []cambio{{emisor, `<nomina12:Emisor/>`}}, []string{ReglaNominaRegistroPatronal, ReglaNominaRegistroPatronal + "/" + NivelAdvertencia}},
		{"contrato 09 con registro patronal", []cambio{{`TipoContrato="01"`, `TipoContrato="09"`}}, []string{ReglaNominaRegistroPatronal, ReglaNominaTipoRegimen}},
		{"registro patronal sin RiesgoPuesto", []cambio{{` RiesgoPuesto="1"`, ``}}, []string{ReglaNominaRegistroPatronal}},
		// Sin RegistroPatronal, que no vayan los demás solo lo pide la guía (p. 29).
		{"asimilado con número de seguridad social", append(append([]cambio(nil), asimilado...), cambio{`TipoRegimen="09"`, `TipoRegimen="09" NumSeguridadSocial="00000000000"`}), []string{ReglaNominaRegistroPatronal + "/" + NivelAdvertencia}},
		// EntidadSNCF (NOM48-50).
		{"SNCF con ingresos propios", []cambio{{emisor, `<nomina12:Emisor RegistroPatronal="Z0000001234"><nomina12:EntidadSNCF OrigenRecurso="IP"/></nomina12:Emisor>`}}, nil},
		{"SNCF mixto sin monto", []cambio{{emisor, `<nomina12:Emisor RegistroPatronal="Z0000001234"><nomina12:EntidadSNCF OrigenRecurso="IM"/></nomina12:Emisor>`}}, []string{ReglaNominaEntidadSNCF}},
		{"SNCF mixto con monto", []cambio{{emisor, `<nomina12:Emisor RegistroPatronal="Z0000001234"><nomina12:EntidadSNCF OrigenRecurso="IM" MontoRecursoPropio="5609.99"/></nomina12:Emisor>`}}, nil},
		{"SNCF mixto con monto igual al total", []cambio{{emisor, `<nomina12:Emisor RegistroPatronal="Z0000001234"><nomina12:EntidadSNCF OrigenRecurso="IM" MontoRecursoPropio="5610.00"/></nomina12:Emisor>`}}, []string{ReglaNominaEntidadSNCF}},
		{"SNCF propio con monto", []cambio{{emisor, `<nomina12:Emisor RegistroPatronal="Z0000001234"><nomina12:EntidadSNCF OrigenRecurso="IP" MontoRecursoPropio="1.00"/></nomina12:Emisor>`}}, []string{ReglaNominaEntidadSNCF}},
		// Antigüedad (NOM53-55): del 2020-01-01 al 2026-01-15 hay 2207 días
		// contando el final, 315 semanas.
		{"316 semanas", antiguedad("2020-01-01", "P316W"), []string{ReglaNominaAntiguedad}},
		{"inicio después del final", antiguedad("2026-01-16", "P1D"), []string{ReglaNominaAntiguedad}},
		{"años y días al límite", antiguedad("2020-01-01", "P6Y15D"), nil},
		{"años y días de más", antiguedad("2020-01-01", "P6Y16D"), []string{ReglaNominaAntiguedad}},
		{"un mes y un día", antiguedad("2025-12-15", "P1M1D"), nil},
		{"un mes y dos días", antiguedad("2025-12-15", "P1M2D"), []string{ReglaNominaAntiguedad}},
		{"fin de mes", antiguedad("2025-12-31", "P15D"), nil},
		// TipoRegimen (NOM57, NOM58, NOM105, NOM106).
		{"contrato 01 con régimen 09", []cambio{{`TipoRegimen="02"`, `TipoRegimen="09"`}}, []string{ReglaNominaTipoRegimen}},
		{"régimen 02 con 007 en vez de 002", []cambio{{`TipoOtroPago="002"`, `TipoOtroPago="007"`}}, nil},
		{"régimen 02 sin otros pagos", asimilado[2:], []string{ReglaNominaTipoRegimen}},
		{"asimilado con subsidio", asimilado[:2], []string{ReglaNominaTipoRegimen}},
		// CuentaBancaria (NOM62-65).
		{"cuenta de 10 con banco", cuenta(`Banco="002" CuentaBancaria="1234567890"`), nil},
		{"cuenta de 10 sin banco", cuenta(`CuentaBancaria="1234567890"`), []string{ReglaNominaCuentaBancaria}},
		{"tarjeta de 15 con banco", cuenta(`Banco="002" CuentaBancaria="123456789012345"`), nil},
		{"cuenta de 12", cuenta(`Banco="002" CuentaBancaria="123456789012"`), []string{ReglaNominaCuentaBancaria}},
		{"CLABE", cuenta(`CuentaBancaria="999180000000000015"`), nil},
		{"CLABE con banco", cuenta(`Banco="002" CuentaBancaria="999180000000000015"`), []string{ReglaNominaCuentaBancaria}},
		{"CLABE con dígito mal", cuenta(`CuentaBancaria="999180000000000016"`), []string{ReglaNominaCuentaBancaria}},
		// Nodos por tipo de percepción (NOM77-85, NOM89, NOM90).
		{"horas extra sin nodo", []cambio{{"\n          " + horas, ``}}, []string{ReglaNominaNodosPercepcion}},
		{"sueldo con horas extra", []cambio{{sueldo, strings.Replace(sueldo, `/>`, `>`+horas+`</nomina12:Percepcion>`, 1)}}, []string{ReglaNominaNodosPercepcion}},
		{"acciones sin nodo", []cambio{{`TipoPercepcion="001"`, `TipoPercepcion="045"`}}, []string{ReglaNominaNodosPercepcion}},
		{"acciones con nodo", []cambio{{sueldo, `<nomina12:Percepcion TipoPercepcion="045" Clave="P045" Concepto="Acciones" ImporteGravado="5000.00" ImporteExento="0.00"><nomina12:AccionesOTitulos ValorMercado="10.00" PrecioAlOtorgarse="5.00"/></nomina12:Percepcion>`}}, nil},
		{"sin TotalSueldos", []cambio{{percs, `<nomina12:Percepciones TotalSeparacionIndemnizacion="5600.00" TotalGravado="5300.00" TotalExento="300.00">`}}, []string{ReglaNominaNodosPercepcion, ReglaNominaPercepciones}},
		{"jubilación en una exhibición", jubilacion("039", `<nomina12:JubilacionPensionRetiro TotalUnaExhibicion="5000.00" IngresoAcumulable="5000.00" IngresoNoAcumulable="0.00"/>`), nil},
		{"jubilación en una exhibición con parcialidad", jubilacion("039", `<nomina12:JubilacionPensionRetiro TotalUnaExhibicion="5000.00" TotalParcialidad="5000.00" IngresoAcumulable="5000.00" IngresoNoAcumulable="0.00"/>`), []string{ReglaNominaNodosPercepcion}},
		{"jubilación en parcialidades", jubilacion("044", `<nomina12:JubilacionPensionRetiro TotalParcialidad="5000.00" MontoDiario="166.67" IngresoAcumulable="5000.00" IngresoNoAcumulable="0.00"/>`), nil},
		{"jubilación en parcialidades sin monto diario", jubilacion("044", `<nomina12:JubilacionPensionRetiro TotalParcialidad="5000.00" IngresoAcumulable="5000.00" IngresoNoAcumulable="0.00"/>`), []string{ReglaNominaNodosPercepcion}},
		{"jubilación en parcialidades con exhibición", jubilacion("044", `<nomina12:JubilacionPensionRetiro TotalUnaExhibicion="5000.00" IngresoAcumulable="5000.00" IngresoNoAcumulable="0.00"/>`), []string{ReglaNominaNodosPercepcion}},
		{"jubilación sin nodo", jubilacion("039", ``), []string{ReglaNominaNodosPercepcion}},
		{"nodo de jubilación sin jubilación", []cambio{{finPercs, `<nomina12:JubilacionPensionRetiro TotalUnaExhibicion="5000.00" IngresoAcumulable="5000.00" IngresoNoAcumulable="0.00"/>` + finPercs}}, []string{ReglaNominaNodosPercepcion}},
		// Incapacidades (NOM86, NOM87, NOM94, NOM95).
		{"subsidio por incapacidad sin importe", []cambio{{`TipoPercepcion="001"`, `TipoPercepcion="014"`}}, []string{ReglaNominaIncapacidades}},
		{"subsidio por incapacidad con importe", []cambio{{`TipoPercepcion="001"`, `TipoPercepcion="014"`}, {incap, `<nomina12:Incapacidad DiasIncapacidad="1" TipoIncapacidad="02" ImporteMonetario="5000.00"/>`}}, nil},
		{"descuento por incapacidad", []cambio{{seguridad, strings.Replace(seguridad, `"001"`, `"006"`, 1)}}, []string{ReglaNominaIncapacidades}},
		{"descuento por incapacidad con importe", []cambio{{seguridad, strings.Replace(seguridad, `"001"`, `"006"`, 1)}, {incap, `<nomina12:Incapacidad DiasIncapacidad="1" TipoIncapacidad="02" ImporteMonetario="125.50"/>`}}, nil},
		{"descuento por incapacidad sin nodo", []cambio{{seguridad, strings.Replace(seguridad, `"001"`, `"006"`, 1)}, {`<nomina12:Incapacidades>
        ` + incap + `
      </nomina12:Incapacidades>`, ``}}, []string{ReglaNominaIncapacidades}},
		// Otros pagos (NOM98, NOM99, NOM103): FechaPago 2026-01-15.
		{"saldo a favor sin nodo", saldoAFavor(``), []string{ReglaNominaCompensacionSaldo}},
		{"saldo a favor del año anterior", saldoAFavor(`<nomina12:CompensacionSaldosAFavor SaldoAFavor="100.00" Año="2025" RemanenteSalFav="95.00"/>`), nil},
		{"saldo a favor de hace dos años", saldoAFavor(`<nomina12:CompensacionSaldosAFavor SaldoAFavor="100.00" Año="2024" RemanenteSalFav="95.00"/>`), []string{ReglaNominaCompensacionSaldo}},
		{"saldo a favor del año en curso en enero", saldoAFavor(`<nomina12:CompensacionSaldosAFavor SaldoAFavor="100.00" Año="2026" RemanenteSalFav="95.00"/>`), []string{ReglaNominaCompensacionSaldo}},
		{"subsidio sin nodo", []cambio{{"\n          " + subsidio, ``}}, []string{ReglaNominaCompensacionSaldo}},
	}
	for _, c := range casos {
		hs := reglasDe(t, aplicarCambios(t, nomina, c.cs))
		got := codigosDe(hs)
		quiero := append([]string(nil), c.quiero...)
		sort.Strings(quiero)
		if strings.Join(got, ",") != strings.Join(quiero, ",") {
			t.Errorf("%s: %v, quiero %v\n%+v", c.nombre, got, quiero, hs)
		}
		for _, h := range hs {
			// Una parte que baja a advertencia (vigencia, solo guía) lo dice.
			bajada := h.Nivel != reglaDe[h.Codigo].Nivel
			if (bajada && (h.Nivel != NivelAdvertencia || !strings.Contains(h.Mensaje, "; advertencia: "))) || h.Regla == "" || h.Mensaje == "" || !strings.HasPrefix(h.Ruta, "/Comprobante") {
				t.Errorf("%s: hallazgo incompleto %+v", c.nombre, h)
			}
		}
	}
	// NOM31: sin TotalPercepciones ni TotalOtrosPagos. Las reglas del CFDI
	// de nómina 4.0 no corren en 3.3.
	if hs := ValidarReglas(&Comprobante{Version: "4.0", Nomina: []Nomina{{Version: "1.2"}}}); !TieneHallazgo(hs, ReglaNominaTotales) {
		t.Errorf("sin totales: %+v", hs)
	}
	c33 := &Comprobante{Version: "3.3", TipoDeComprobante: "N", Exportacion: "02", Nomina: []Nomina{{Version: "1.2", TotalOtrosPagos: ptrDec("0")}}}
	if hs := ValidarReglas(c33); TieneHallazgo(hs, ReglaNominaComprobante) {
		t.Errorf("3.3: %+v", hs)
	}
}

// Vigencia de NOM103, NOM105 y NOM106 (Regla.Vigencia): error desde el
// 2020-02-01 (matriz Rev B), advertencia antes; el borde es el primer
// segundo del día, en hora civil. Lo demás de la regla no cambia con la
// fecha.
func TestValidarReglasNominaVigencia(t *testing.T) {
	nomina := string(leerFixture(t, "testdata/sinteticos/nomina12.xml"))
	const fecha = `Fecha="2026-01-15T08:00:00"`
	sinOtrosPagos := []cambio{
		{`<nomina12:OtrosPagos>
        <nomina12:OtroPago TipoOtroPago="002" Clave="O002" Concepto="Subsidio para el empleo" Importe="10.00">
          <nomina12:SubsidioAlEmpleo SubsidioCausado="10.00"/>
        </nomina12:OtroPago>
      </nomina12:OtrosPagos>`, ``},
		{` TotalOtrosPagos="10.00"`, ``}, {`ValorUnitario="5610.00" Importe="5610.00"`, `ValorUnitario="5600.00" Importe="5600.00"`},
		{`SubTotal="5610.00"`, `SubTotal="5600.00"`}, {`Total="5034.50"`, `Total="5024.50"`},
	}
	regimen09 := []cambio{{`TipoRegimen="02"`, `TipoRegimen="09"`}, {`TipoContrato="01"`, `TipoContrato="99"`}, {`<nomina12:Emisor RegistroPatronal="Z0000001234"/>`, `<nomina12:Emisor/>`},
		{` NumSeguridadSocial="00000000000" FechaInicioRelLaboral="2020-01-01" Antigüedad="P315W"`, ``}, {` RiesgoPuesto="1"`, ``}, {` SalarioDiarioIntegrado="380.50"`, ``}}
	saldo2024 := []cambio{
		{`</nomina12:OtroPago>`, `</nomina12:OtroPago><nomina12:OtroPago TipoOtroPago="004" Clave="O004" Concepto="Saldo a favor" Importe="5.00"><nomina12:CompensacionSaldosAFavor SaldoAFavor="100.00" Año="2024" RemanenteSalFav="95.00"/></nomina12:OtroPago>`},
		{`TotalOtrosPagos="10.00"`, `TotalOtrosPagos="15.00"`}, {`ValorUnitario="5610.00" Importe="5610.00"`, `ValorUnitario="5615.00" Importe="5615.00"`},
		{`SubTotal="5610.00"`, `SubTotal="5615.00"`}, {`Total="5034.50"`, `Total="5039.50"`},
	}
	adv := func(c string) string { return c + "/" + NivelAdvertencia }
	casos := []struct {
		nombre, fecha string
		cs            []cambio
		quiero        string
	}{
		{"NOM105 el día anterior", "2020-01-31T23:59:59", sinOtrosPagos, adv(ReglaNominaTipoRegimen)},
		{"NOM105 el primer segundo", "2020-02-01T00:00:00", sinOtrosPagos, ReglaNominaTipoRegimen},
		{"NOM105 en 2019", "2019-06-30T12:00:00", sinOtrosPagos, adv(ReglaNominaTipoRegimen)},
		{"NOM106 el día anterior", "2020-01-31T23:59:59", regimen09, adv(ReglaNominaTipoRegimen)},
		{"NOM106 el primer segundo", "2020-02-01T00:00:00", regimen09, ReglaNominaTipoRegimen},
		{"NOM103 el día anterior", "2020-01-31T23:59:59", saldo2024, adv(ReglaNominaCompensacionSaldo)},
		{"NOM103 el primer segundo", "2020-02-01T00:00:00", saldo2024, ReglaNominaCompensacionSaldo},
		// Entre la matriz Rev B y la Rev C ya es error (antes se tomaba la
		// Rev C como inicio).
		{"NOM105 en 2021", "2021-06-30T12:00:00", sinOtrosPagos, ReglaNominaTipoRegimen},
		{"NOM103 en 2021", "2021-06-30T12:00:00", saldo2024, ReglaNominaCompensacionSaldo},
		// NOM57 (TipoContrato 01 con TipoRegimen 09) está desde la primera
		// versión: error en cualquier fecha. Con TipoRegimen 09 también
		// salta NOM106 por el OtroPago 002, que sí baja.
		{"NOM57 en 2019", "2019-06-30T12:00:00", []cambio{{`TipoRegimen="02"`, `TipoRegimen="09"`}}, ReglaNominaTipoRegimen + "," + adv(ReglaNominaTipoRegimen)},
		// Sin Fecha legible la regla se aplica.
		{"NOM105 con fecha ilegible", "2021-13-40T00:00:00", sinOtrosPagos, ReglaNominaTipoRegimen},
	}
	for _, c := range casos {
		cs := append([]cambio{{fecha, `Fecha="` + c.fecha + `"`}}, c.cs...)
		hs := reglasDe(t, aplicarCambios(t, nomina, cs))
		got := strings.Join(codigosDe(hs), ",")
		if got != c.quiero {
			t.Errorf("%s: %s, quiero %s\n%+v", c.nombre, got, c.quiero, hs)
		}
		for _, h := range hs {
			if h.Nivel == NivelAdvertencia && !strings.Contains(h.Mensaje, "es anterior al 2020-02-01") {
				t.Errorf("%s: la advertencia no dice la vigencia: %s", c.nombre, h.Mensaje)
			}
		}
	}
	// En 3.3, como los 178 del corpus: la vigencia es la misma.
	c33 := func(fecha string) []Hallazgo {
		f, _, _ := parseFecha(fecha, false)
		return ValidarReglas(&Comprobante{Version: "3.3", Fecha: f, TipoDeComprobante: "N", Nomina: []Nomina{{Version: "1.2", TotalPercepciones: ptrDec("0"),
			Percepciones: &Percepciones{}, Receptor: &NominaReceptor{TipoRegimen: "02"}}}})
	}
	for fecha, nivel := range map[string]string{"2020-01-15T10:00:00": NivelAdvertencia, "2020-03-15T10:00:00": NivelError} {
		n := 0
		for _, h := range c33(fecha) {
			if h.Codigo == ReglaNominaTipoRegimen {
				n++
				if h.Nivel != nivel {
					t.Errorf("3.3 %s: %s, quiero %s", fecha, h.Nivel, nivel)
				}
			}
		}
		if n != 1 {
			t.Errorf("3.3 %s: %d hallazgos de TipoRegimen", fecha, n)
		}
	}
}

// Fin de mes y años bisiestos: el 31 de enero más un mes es el último día
// de febrero, no el 3 de marzo; y el 29 de febrero más un año, el 28.
func TestSumarMeses(t *testing.T) {
	d := func(s string) time.Time { x, _ := time.Parse("2006-01-02", s); return x }
	casos := []struct {
		desde       string
		anios, mese int
		quiero      string
	}{
		{"2026-01-31", 0, 1, "2026-02-28"}, {"2024-01-31", 0, 1, "2024-02-29"}, {"2024-02-29", 1, 0, "2025-02-28"},
		{"2025-12-15", 0, 1, "2026-01-15"}, {"2020-01-01", 6, 0, "2026-01-01"}, {"2025-11-30", 0, 3, "2026-02-28"},
	}
	for _, c := range casos {
		if got := sumarMeses(d(c.desde), c.anios, c.mese); !got.Equal(d(c.quiero)) {
			t.Errorf("%s + %dY%dM = %s, quiero %s", c.desde, c.anios, c.mese, got.Format("2006-01-02"), c.quiero)
		}
	}
	// El dígito de control de la CLABE con los pesos de Banxico; el ejemplo
	// de la guía de nómina (p. 43) no lo cumple.
	for s, dc := range map[string]int{"999180000000000015": 5, "123456789012345673": 3, "002215911558451272": 4} {
		if got := digitoCLABE(s); got != dc {
			t.Errorf("digitoCLABE(%s) = %d, quiero %d", s, got, dc)
		}
	}
}
