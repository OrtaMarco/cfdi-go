package cfdi

import (
	"encoding/json"
	"testing"
)

// G4 (T10): decretos de vehículos, obras de arte, PFIC, hidrocarburos y los
// estados de cuenta de CFDI 3.2. Públicos de satcfdi (sin timbre); ECB y ECC
// sin versión no tienen ejemplo público: se agregan a un 3.2 sintético.

// cumpleSchema valida el JSON de c contra schema/cfdi.v1.schema.json (lo
// que TestSchemaFixtures hace con los fixtures de testdata).
func cumpleSchema(t *testing.T, c *Comprobante) {
	t.Helper()
	out, err := json.Marshal(c.Documento(OpcionesJSON{IncluirSello: true}))
	if err != nil {
		t.Fatal(err)
	}
	for _, e := range validarJSON(t, cargarEsquema(t), out) {
		t.Error(e)
	}
}

func TestRenovacionySustitucionVehiculos(t *testing.T) {
	// El ejemplo trae dos nodos (renovación y sustitución): se usa el primero.
	p := leerRuta(t, satcfdi40+"renovacionysustitucionvehiculos.xml")
	if cs := codigos(p); cs[AdvNodoDuplicado] != 1 || cs[AdvSinTimbre] != 1 || len(cs) != 2 {
		t.Errorf("advertencias = %v", p.Advertencias)
	}
	soportado(t, p, "renovacionysustitucionvehiculos 1.0")
	x := p.RenovacionySustitucionVehiculos
	if x == nil || x.DecretoRenovVehicular == nil || x.DecretoSustitVehicular != nil || len(x.DecretoRenovVehicular.VehiculosUsadosEnajenadoPermAlFab) != 1 {
		t.Fatalf("renovacionysustitucionvehiculos = %+v", x)
	}
	igual(t, "TipoDeDecreto", x.TipoDeDecreto, "01")
	u := x.DecretoRenovVehicular.VehiculosUsadosEnajenadoPermAlFab[0]
	monto(t, "PrecioVehUsado", u.PrecioVehUsado, "10001.56")
	igual(t, "Año", u.Anio, "2001")
	igual(t, "FechaRegulVeh", u.FechaRegulVeh.Original, "1999-12-31")
	igual(t, "RFC del nuevo", x.DecretoRenovVehicular.VehiculoNuvoSemEnajenadoFabAlPerm.RFC, "PEPJ8001019Q8")

	// Sin el primero queda la sustitución, que exige además NumFolAvisoint.
	sinRenov := []string{`<decreto:renovacionysustitucionvehiculos Version="1.0" TipoDeDecreto="01">`, `<!--`,
		`</decreto:renovacionysustitucionvehiculos>`, `-->`}
	s := leerMutado(t, satcfdi40+"renovacionysustitucionvehiculos.xml", sinRenov...)
	sinMasQue(t, s, AdvSinTimbre)
	d := s.RenovacionySustitucionVehiculos.DecretoSustitVehicular
	if d == nil || d.VehiculoUsadoEnajenadoPermAlFab == nil || d.VehiculoNuvoSemEnajenadoFabAlPerm == nil {
		t.Fatalf("DecretoSustitVehicular = %+v", d)
	}
	igual(t, "NumFolAvisoint", d.VehiculoUsadoEnajenadoPermAlFab.NumFolAvisoint, "ISO12345")
	igual(t, "VehEnaj", d.VehEnaj, "01")

	m := leerMutado(t, satcfdi40+"renovacionysustitucionvehiculos.xml", append(sinRenov, ` NumFolAvisoint="ISO12345"`, ``)...)
	if cs := codigos(m); cs[AdvAtributoFaltante] != 1 || len(cs) != 2 {
		t.Errorf("advertencias = %v", m.Advertencias)
	}
}

func TestCertificadoDeDestruccion(t *testing.T) {
	p := leerRuta(t, satcfdi40+"certdestruccion.xml")
	sinMasQue(t, p, AdvSinTimbre)
	soportado(t, p, "certificadodedestruccion 1.0")
	x := p.CertificadoDeDestruccion
	if x == nil || x.VehiculoDestruido == nil || x.InformacionAduanera == nil {
		t.Fatalf("certificadodedestruccion = %+v", x)
	}
	igual(t, "Serie", x.Serie, "SERIE E")
	igual(t, "NumFolDesVeh", x.NumFolDesVeh, "123ABC")
	igual(t, "Año", x.VehiculoDestruido.Anio, "1999")
	igual(t, "NumFolTarjCir", x.VehiculoDestruido.NumFolTarjCir, "3456")
	igual(t, "Fecha aduana", x.InformacionAduanera.Fecha.Original, "2015-05-31")

	m := leerMutado(t, satcfdi40+"certdestruccion.xml", `<destruccion:VehiculoDestruido`, `<!--`, `TipooClase="TipooClase" />`, `-->`)
	if cs := codigos(m); cs[AdvNodoFaltante] != 1 || len(cs) != 2 {
		t.Errorf("advertencias = %v", m.Advertencias)
	}
}

func TestObrasArteAntiguedades(t *testing.T) {
	p := leerRuta(t, satcfdi40+"obrasarte.xml")
	sinMasQue(t, p, AdvSinTimbre)
	soportado(t, p, "obrasarteantiguedades 1.0")
	x := p.ObrasArteAntiguedades
	igual(t, "TipoBien", x.TipoBien, "04")
	igual(t, "OtrosTituloAdquirido", x.OtrosTituloAdquirido, "OtrosTituloAdquirido")
	igual(t, "CaracterísticasDeObraoPieza", x.CaracteristicasDeObraoPieza, "07")
	monto(t, "Subtotal", x.Subtotal, "100.999999")
	monto(t, "IVA", x.IVA, "10.01")
	igual(t, "FechaAdquisicion", x.FechaAdquisicion.Original, "2015-08-10")

	m := leerMutado(t, satcfdi40+"obrasarte.xml", `IVA="10.01"`, `IVA="diez"`, ` CaracterísticasDeObraoPieza="07"`, ``)
	if cs := codigos(m); cs[AdvDecimalInvalido] != 1 || cs[AdvAtributoFaltante] != 1 || len(cs) != 3 {
		t.Errorf("advertencias = %v", m.Advertencias)
	}
}

func TestPFintegranteCoordinado(t *testing.T) {
	// El atributo es version, en minúscula.
	p := leerRuta(t, satcfdi40+"pfic.xml")
	sinMasQue(t, p, AdvSinTimbre)
	soportado(t, p, "PFintegranteCoordinado 1.0")
	x := p.PFintegranteCoordinado
	igual(t, "ClaveVehicular", x.ClaveVehicular, "A12345")
	igual(t, "Placa", x.Placa, "Placa")

	m := leerMutado(t, satcfdi40+"pfic.xml", `RFCPF="AAA010101AAA"`, `RFCPF=" aaa010101aaa "`, ` Placa="Placa"`, ``)
	igual(t, "RFCPF normalizado", m.PFintegranteCoordinado.RFCPF, "AAA010101AAA")
	if cs := codigos(m); cs[AdvAtributoFaltante] != 1 || len(cs) != 2 {
		t.Errorf("advertencias = %v", m.Advertencias)
	}
}

func TestGastosHidrocarburos(t *testing.T) {
	p := leerRuta(t, satcfdi40+"gceh.xml")
	sinMasQue(t, p, AdvSinTimbre)
	soportado(t, p, "GastosHidrocarburos 1.0")
	x := p.GastosHidrocarburos
	if x == nil || len(x.Erogacion) != 2 || len(x.Erogacion[0].Actividades) != 1 || len(x.Erogacion[0].CentroCostos) != 1 || x.Erogacion[1].Actividades != nil {
		t.Fatalf("GastosHidrocarburos = %+v", x)
	}
	igual(t, "NumeroContrato", x.NumeroContrato, "12345678")
	e := x.Erogacion[0]
	monto(t, "MontocuErogacion", e.MontocuErogacion, "100.999")
	igual(t, "Tarea", e.Actividades[0].SubActividades[0].Tareas[0].TareaRelacionada, "0155")
	igual(t, "Pozo", e.CentroCostos[0].Yacimientos[0].Pozos[0].Pozo, "Pozo")
	dr := e.DocumentoRelacionado[0]
	igual(t, "OrigenErogacion", dr.OrigenErogacion, "Extranjero")
	monto(t, "MontoIVAPedimento", dr.MontoIVAPedimento, "90.000001")
	igual(t, "FechaFolioFiscalVinculado", dr.FechaFolioFiscalVinculado.Original, "2018-10-22")

	m := leerMutado(t, satcfdi40+"gceh.xml", ` Mes="10" MontoTotalErogaciones="0.123"`, ` MontoTotalErogaciones="0.123"`,
		`MontocuErogacion="100.999"`, `MontocuErogacion="1,0"`)
	if cs := codigos(m); cs[AdvAtributoFaltante] != 1 || cs[AdvDecimalInvalido] != 1 || len(cs) != 3 {
		t.Errorf("advertencias = %v", m.Advertencias)
	}
}

func TestIngresosHidrocarburos(t *testing.T) {
	p := leerRuta(t, satcfdi40+"ieeh.xml")
	sinMasQue(t, p, AdvSinTimbre)
	soportado(t, p, "IngresosHidrocarburos 1.0")
	x := p.IngresosHidrocarburos
	if x == nil || len(x.DocumentoRelacionado) != 2 {
		t.Fatalf("IngresosHidrocarburos = %+v", x)
	}
	monto(t, "ContraprestacionPagadaOperador", x.ContraprestacionPagadaOperador, "1000.999999")
	monto(t, "Porcentaje", x.Porcentaje, "80.001")
	igual(t, "Mes", x.DocumentoRelacionado[1].Mes, "09")

	m := leerMutado(t, satcfdi40+"ieeh.xml", ` Version="1.0" Porcentaje="80.001"`, ` Version="1.1"`)
	if cs := codigos(m); cs[AdvComplementoVersion] != 1 || cs[AdvAtributoFaltante] != 1 || len(cs) != 3 {
		t.Errorf("advertencias = %v", m.Advertencias)
	}
}

// ecb32 es un EstadoDeCuentaBancario con los dos tipos de movimiento
// intercalados (el xs:choice repetible de Movimientos), RFC de prueba.
const ecb32 = `<ecb:EstadoDeCuentaBancario xmlns:ecb="http://www.sat.gob.mx/ecb" version="1.0" numeroCuenta="0123456789" nombreCliente="CLIENTE DE PRUEBA" periodo="JUNIO 2016">
      <ecb:Movimientos>
        <ecb:MovimientoECB fecha="2016-06-01T09:00:00" descripcion="DEPOSITO" importe="500.00" moneda="MXN"/>
        <ecb:MovimientoECBFiscal fecha="2016-06-02T10:00:00" descripcion="COMISION" RFCenajenante="eku9003173c9" Importe="10.00"/>
        <ecb:MovimientoECB fecha="2016-06-03T11:00:00" referencia="R-2" descripcion="RETIRO" importe="-100.00" saldoInicial="1000.00" saldoAlCorte="1390.00"/>
      </ecb:Movimientos>
    </ecb:EstadoDeCuentaBancario>
    `

func TestEstadoDeCuentaBancario(t *testing.T) {
	c := leerMutado(t, "sinteticos/t10/spei32.xml", `<spei:Complemento_SPEI>`, ecb32+`<spei:Complemento_SPEI>`)
	sinMasQue(t, c, AdvVersionAntigua)
	cumpleSchema(t, c)
	soportado(t, c, "EstadoDeCuentaBancario 1.0")
	x := c.EstadoDeCuentaBancario
	if x == nil || x.Movimientos == nil || len(x.Movimientos.MovimientoECB) != 2 || len(x.Movimientos.MovimientoECBFiscal) != 1 {
		t.Fatalf("EstadoDeCuentaBancario = %+v", x)
	}
	igual(t, "NumeroCuenta", x.NumeroCuenta, "0123456789")
	igual(t, "Periodo", x.Periodo, "JUNIO 2016")
	monto(t, "importe", x.Movimientos.MovimientoECB[1].Importe, "-100.00")
	monto(t, "saldoAlCorte", x.Movimientos.MovimientoECB[1].SaldoAlCorte, "1390.00")
	f := x.Movimientos.MovimientoECBFiscal[0]
	igual(t, "RFCenajenante normalizado", f.RFCEnajenante, "EKU9003173C9")
	monto(t, "Importe (mayúscula)", f.Importe, "10.00")
	igual(t, "fecha", f.Fecha.Original, "2016-06-02T10:00:00")

	m := leerMutado(t, "sinteticos/t10/spei32.xml", `<spei:Complemento_SPEI>`, ecb32+`<spei:Complemento_SPEI>`,
		` RFCenajenante="eku9003173c9"`, ``, ` periodo="JUNIO 2016"`, ``)
	if cs := codigos(m); cs[AdvAtributoFaltante] != 2 || len(cs) != 2 {
		t.Errorf("advertencias = %v", m.Advertencias)
	}
}

// ecc32 es un EstadoDeCuentaCombustible sin versión (namespace .../ecc):
// atributos en minúscula y tasa en el traslado.
const ecc32 = `<ecc:EstadoDeCuentaCombustible xmlns:ecc="http://www.sat.gob.mx/ecc" tipoOperacion="Tarjeta" numeroDeCuenta="0123456789" subTotal="100.00" total="116.00">
      <ecc:Conceptos>
        <ecc:ConceptoEstadoDeCuentaCombustible identificador="T-1" fecha="2016-06-01T08:00:00" rfc="eku9003173c9" claveEstacion="E-1" cantidad="7.5" nombreCombustible="MAGNA" folioOperacion="F-1" valorUnitario="13.333333" importe="100.00">
          <ecc:Traslados>
            <ecc:Traslado impuesto="IVA" tasa="16.00" importe="16.00"/>
          </ecc:Traslados>
        </ecc:ConceptoEstadoDeCuentaCombustible>
      </ecc:Conceptos>
    </ecc:EstadoDeCuentaCombustible>
    `

func TestEstadoDeCuentaCombustible10(t *testing.T) {
	c := leerMutado(t, "sinteticos/t10/spei32.xml", `<spei:Complemento_SPEI>`, ecc32+`<spei:Complemento_SPEI>`)
	sinMasQue(t, c, AdvVersionAntigua)
	cumpleSchema(t, c)
	soportado(t, c, "EstadoDeCuentaCombustible 1.0")
	e := c.EstadoDeCuentaCombustible
	if e == nil || len(e.Conceptos) != 1 || len(e.Conceptos[0].Traslados) != 1 {
		t.Fatalf("EstadoDeCuentaCombustible = %+v", e)
	}
	igual(t, "Version (del namespace)", e.Version, "1.0")
	igual(t, "TipoOperacion", e.TipoOperacion, "Tarjeta")
	monto(t, "Total", e.Total, "116.00")
	cp := e.Conceptos[0]
	igual(t, "rfc normalizado", cp.Rfc, "EKU9003173C9")
	monto(t, "valorUnitario", cp.ValorUnitario, "13.333333")
	tr := cp.Traslados[0]
	monto(t, "tasa", tr.Tasa, "16.00")
	if tr.TasaOCuota != nil {
		t.Errorf("ecc sin versión con TasaOCuota: %+v", tr)
	}

	// Sin tasa en el traslado ni numeroDeCuenta: los dos los exige su XSD.
	m := leerMutado(t, "sinteticos/t10/spei32.xml", `<spei:Complemento_SPEI>`, ecc32+`<spei:Complemento_SPEI>`,
		` tasa="16.00" importe="16.00"`, ` importe="16.00"`, ` numeroDeCuenta="0123456789"`, ``)
	if cs := codigos(m); cs[AdvAtributoFaltante] != 2 || len(cs) != 2 {
		t.Errorf("advertencias = %v", m.Advertencias)
	}
}
