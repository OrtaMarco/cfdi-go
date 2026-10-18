package cfdi

import (
	"strings"
	"testing"
)

// G3 (T10): complementos de uso común. Públicos de satcfdi (sin timbre) y
// sintéticos para lo que no tiene ejemplo público.

// sinMasQue exige que las únicas advertencias sean las de los códigos dados
// (p. ej. SIN_TIMBRE en los ejemplos de satcfdi).
func sinMasQue(t *testing.T, c *Comprobante, permitidos ...string) {
	t.Helper()
	for _, a := range c.Advertencias {
		ok := false
		for _, p := range permitidos {
			ok = ok || a.Codigo == p
		}
		if !ok {
			t.Errorf("advertencia inesperada: %+v", a)
		}
	}
}

func TestINE(t *testing.T) {
	p := leerRuta(t, satcfdi40+"ine11.xml")
	sinMasQue(t, p, AdvSinTimbre)
	soportado(t, p, "INE 1.1")
	x := p.INE
	if x == nil || len(x.Entidad) != 1 || len(x.Entidad[0].Contabilidad) != 1 {
		t.Fatalf("INE = %+v", x)
	}
	igual(t, "TipoProceso", x.TipoProceso, "Campaña")
	igual(t, "TipoComite", x.TipoComite, "Directivo Estatal")
	igual(t, "IdContabilidad", x.IdContabilidad, "123456")
	igual(t, "ClaveEntidad", x.Entidad[0].ClaveEntidad, "ZAC")
	igual(t, "Ambito", x.Entidad[0].Ambito, "Federal")
	igual(t, "Contabilidad", x.Entidad[0].Contabilidad[0].IdContabilidad, "123456")

	// 1.0 comparte namespace: se lee sin advertencia.
	c := leerRuta(t, "sinteticos/t10/ine10.xml")
	sinAdvertencias(t, c)
	soportado(t, c, "INE 1.0")
	if len(c.INE.Entidad) != 2 || len(c.INE.Entidad[0].Contabilidad) != 2 || c.INE.Entidad[1].Contabilidad != nil {
		t.Fatalf("INE 1.0 = %+v", c.INE)
	}
	igual(t, "ClaveEntidad 1.0", c.INE.Entidad[0].ClaveEntidad, "QTO")

	m := leerMutado(t, "sinteticos/t10/ine10.xml", `Version="1.0" TipoProceso`, `Version="2.0" TipoProceso`, ` ClaveEntidad="HID"`, ``)
	if cs := codigos(m); cs[AdvComplementoVersion] != 1 || cs[AdvAtributoFaltante] != 1 || len(cs) != 2 {
		t.Errorf("advertencias = %v", m.Advertencias)
	}
}

func TestValesDeDespensa(t *testing.T) {
	p := leerRuta(t, satcfdi40+"valesdedespensa.xml")
	sinMasQue(t, p, AdvSinTimbre)
	soportado(t, p, "ValesDeDespensa 1.0")
	v := p.ValesDeDespensa
	if v == nil || len(v.Conceptos) != 2 {
		t.Fatalf("ValesDeDespensa = %+v", v)
	}
	igual(t, "TipoOperacion", v.TipoOperacion, "monedero electrónico")
	igual(t, "RegistroPatronal", v.RegistroPatronal, "ABCXYZ")
	monto(t, "Total", v.Total, "302.00")
	igual(t, "NumSeguridadSocial", v.Conceptos[0].NumSeguridadSocial, "XYZ1234")
	igual(t, "NumSeguridadSocial opcional", v.Conceptos[1].NumSeguridadSocial, "")
	monto(t, "Importe", v.Conceptos[1].Importe, "201.99")
	igual(t, "Fecha", v.Conceptos[0].Fecha.Original, "2014-07-18T12:37:17")

	m := leerMutado(t, "mutantes/base-vales.xml", `<valesdedespensa:Concepto identificador="ID-01" fecha="2026-01-15T10:00:00" rfc="XOJI740919U48" curp="XEXX010101HNEXXXA4" nombre="TRABAJADOR DE PRUEBA" importe="100.00"/>`, ``)
	if cs := codigos(m); cs[AdvNodoFaltante] != 1 || len(cs) != 1 {
		t.Errorf("advertencias = %v", m.Advertencias)
	}
}

func TestConsumoDeCombustibles(t *testing.T) {
	p := leerRuta(t, satcfdi40+"consumodecombustibles11.xml")
	sinMasQue(t, p, AdvSinTimbre)
	soportado(t, p, "ConsumoDeCombustibles 1.1")
	cc := p.ConsumoDeCombustibles
	if cc == nil || len(cc.Conceptos) != 1 || len(cc.Conceptos[0].Determinados) != 2 {
		t.Fatalf("ConsumoDeCombustibles = %+v", cc)
	}
	igual(t, "TipoCombustible", cc.Conceptos[0].TipoCombustible, "1")
	d := cc.Conceptos[0].Determinados[1]
	igual(t, "Impuesto", d.Impuesto, "IEPS")
	monto(t, "TasaOCuota", d.TasaOCuota, "10.0")
	if d.Tasa != nil {
		t.Errorf("1.1 con Tasa: %+v", d)
	}

	// 1.0: tasa en vez de tasaOCuota, sin tipoCombustible, y no lo exige.
	c := leerRuta(t, "sinteticos/t10/consumo10.xml")
	sinAdvertencias(t, c)
	soportado(t, c, "ConsumoDeCombustibles 1.0")
	d10 := c.ConsumoDeCombustibles.Conceptos[0].Determinados[0]
	monto(t, "Tasa", d10.Tasa, "16.00")
	if d10.TasaOCuota != nil {
		t.Errorf("1.0 con TasaOCuota: %+v", d10)
	}

	// 1.1 sí exige tipoCombustible y tasaOCuota.
	m := leerMutado(t, "mutantes/base-consumo11.xml", ` tipoCombustible="1"`, ``, ` tasaOCuota="15.5"`, ``)
	if cs := codigos(m); cs[AdvAtributoFaltante] != 2 || len(cs) != 1 {
		t.Errorf("advertencias = %v", m.Advertencias)
	}
}

func TestPagoEnEspecie(t *testing.T) {
	p := leerRuta(t, satcfdi40+"pagoenespecie.xml")
	sinMasQue(t, p, AdvSinTimbre)
	soportado(t, p, "PagoEnEspecie 1.0")
	pe := p.PagoEnEspecie
	igual(t, "CvePIC", pe.CvePIC, "A&Ñ091939AZ0-20891030-999")
	igual(t, "FolioSolDon", pe.FolioSolDon, "PE-09-12345")
	igual(t, "PzaArtAProd", pe.PzaArtAProd, "1901")
	igual(t, "PzaArtDim", pe.PzaArtDim, "260x987")

	m := leerMutado(t, "mutantes/base-pagoenespecie.xml", ` PzaArtDim="260x987"`, ``)
	if cs := codigos(m); cs[AdvAtributoFaltante] != 1 || len(cs) != 1 {
		t.Errorf("advertencias = %v", m.Advertencias)
	}
}

func TestNotariosPublicos(t *testing.T) {
	p := leerRuta(t, satcfdi40+"notariospublicos.xml")
	sinMasQue(t, p, AdvSinTimbre)
	soportado(t, p, "NotariosPublicos 1.0")
	np := p.NotariosPublicos
	if np == nil || len(np.DescInmuebles) != 1 || np.DatosOperacion == nil || np.DatosNotario == nil ||
		np.DatosEnajenante == nil || np.DatosEnajenante.DatosUnEnajenante == nil ||
		np.DatosAdquiriente == nil || np.DatosAdquiriente.DatosUnAdquiriente == nil {
		t.Fatalf("NotariosPublicos = %+v", np)
	}
	igual(t, "DescInmueble.Estado", np.DescInmuebles[0].Estado, "19")
	igual(t, "DescInmueble.NoInterior", np.DescInmuebles[0].NoInterior, "D")
	igual(t, "FechaInstNotarial", np.DatosOperacion.FechaInstNotarial.Original, "2014-05-02")
	monto(t, "IVA", np.DatosOperacion.IVA, "10")
	igual(t, "NumNotaria", np.DatosNotario.NumNotaria, "777")
	igual(t, "CoproSocConyugalE", np.DatosEnajenante.CoproSocConyugalE, "No")

	// Copropietarios (sintético): la lista aplanada.
	c := leerRuta(t, "mutantes/base-notarios.xml")
	sinAdvertencias(t, c)
	a := c.NotariosPublicos.DatosAdquiriente
	if a.DatosUnAdquiriente != nil || len(a.DatosAdquirientesCopSC) != 2 {
		t.Fatalf("DatosAdquiriente = %+v", a)
	}
	igual(t, "CopSC.RFC", a.DatosAdquirientesCopSC[1].RFC, "CACX7605101P8")
	monto(t, "CopSC.Porcentaje", a.DatosAdquirientesCopSC[1].Porcentaje, "50.00")

	// El XSD exige CURP en DatosUnEnajenante y no en DatosUnAdquiriente.
	m := leerMutado(t, "mutantes/base-notarios.xml", `RFC="XOJI740919U48" CURP="XEXX010101HNEXXXA4"`, `RFC="XOJI740919U48"`)
	if _, ok := m.AdvertenciaEn("/Comprobante/Complemento/NotariosPublicos/DatosEnajenante/DatosUnEnajenante", "CURP"); !ok || len(m.Advertencias) != 1 {
		t.Errorf("advertencias = %v", m.Advertencias)
	}
	m = leerMutado(t, "mutantes/base-notarios.xml", `<np:DatosAdquiriente CoproSocConyugalE="Si">`,
		`<np:DatosAdquiriente CoproSocConyugalE="Si"><np:DatosUnAdquiriente Nombre="UNO" RFC="XIQB891116QE4"/>`)
	sinAdvertencias(t, m)
	igual(t, "DatosUnAdquiriente sin CURP", m.NotariosPublicos.DatosAdquiriente.DatosUnAdquiriente.RFC, "XIQB891116QE4")
}

func TestParcialesConstruccion(t *testing.T) {
	p := leerRuta(t, satcfdi40+"servicioparcial.xml")
	sinMasQue(t, p, AdvSinTimbre)
	soportado(t, p, "parcialesconstruccion 1.0")
	pc := p.ParcialesConstruccion
	if pc == nil || pc.Inmueble == nil {
		t.Fatalf("parcialesconstruccion = %+v", pc)
	}
	igual(t, "NumPerLicoAut", pc.NumPerLicoAut, "1234567")
	igual(t, "Calle tal cual", pc.Inmueble.Calle, " AV. RIO MIXCOAC ")
	igual(t, "Estado", pc.Inmueble.Estado, "32")

	// Fuera de lugar, en un ComplementoConcepto: se lee en el concepto.
	m := leerMutado(t, "mutantes/base-acreditamiento.xml",
		`<aieps:acreditamientoIEPS xmlns:aieps="http://www.sat.gob.mx/acreditamiento" Version="1.0" TAR="617"/>`,
		`<spc:parcialesconstruccion xmlns:spc="http://www.sat.gob.mx/servicioparcialconstruccion" Version="1.0" NumPerLicoAut="7"><spc:Inmueble Calle="C" Municipio="M" Estado="13" CodigoPostal="42501"/></spc:parcialesconstruccion>`)
	sinAdvertencias(t, m)
	soportado(t, m, "parcialesconstruccion 1.0")
	if m.ParcialesConstruccion != nil || m.Conceptos[0].ParcialesConstruccion == nil {
		t.Fatalf("parcialesconstruccion de concepto: %+v / %+v", m.ParcialesConstruccion, m.Conceptos[0].ParcialesConstruccion)
	}
	igual(t, "de concepto", m.Conceptos[0].ParcialesConstruccion.NumPerLicoAut, "7")
}

func TestVehiculoUsado(t *testing.T) {
	p := leerRuta(t, satcfdi40+"vehiculousado.xml")
	sinMasQue(t, p, AdvSinTimbre)
	soportado(t, p, "VehiculoUsado 1.0")
	v := p.VehiculoUsado
	if v == nil || len(v.InformacionAduanera) != 2 {
		t.Fatalf("VehiculoUsado = %+v", v)
	}
	monto(t, "MontoAdquisicion", v.MontoAdquisicion, "1000.47")
	monto(t, "Valor", v.Valor, "123.123456")
	igual(t, "NIV", v.NIV, "0")
	igual(t, "Aduana", v.InformacionAduanera[1].Aduana, "Otra oficina")
	igual(t, "Fecha", v.InformacionAduanera[1].Fecha.Original, "2014-10-24")

	m := leerMutado(t, "mutantes/base-vehiculousado.xml", ` fecha="2020-10-23"`, ``)
	if cs := codigos(m); cs[AdvAtributoFaltante] != 1 || len(cs) != 1 {
		t.Errorf("advertencias = %v", m.Advertencias)
	}
}

func TestTuristaPasajeroExtranjero(t *testing.T) {
	p := leerRuta(t, satcfdi40+"tpe.xml")
	sinMasQue(t, p, AdvSinTimbre)
	soportado(t, p, "TuristaPasajeroExtranjero 1.0")
	x := p.TuristaPasajeroExtranjero
	if x == nil || x.DatosTransito == nil {
		t.Fatalf("TuristaPasajeroExtranjero = %+v", x)
	}
	igual(t, "FechadeTransito", x.FechadeTransito.Original, "2016-12-31T12:00:59")
	igual(t, "TipoTransito", x.TipoTransito, "Salida")
	igual(t, "Via", x.DatosTransito.Via, "Terrestre")
	igual(t, "IdTransporte", x.DatosTransito.IdTransporte, "IDT")

	m := leerMutado(t, "mutantes/base-tpe.xml", `<tpe:datosTransito Via="Terrestre" TipoId="P123456" NumeroId="N123456" Nacionalidad="MEX" EmpresaTransporte="EMPRESA DE PRUEBA"/>`, ``)
	if cs := codigos(m); cs[AdvNodoFaltante] != 1 || len(cs) != 1 {
		t.Errorf("advertencias = %v", m.Advertencias)
	}
}

func TestAcreditamientoIEPS(t *testing.T) {
	c := leerRuta(t, "sinteticos/t10/ine10.xml")
	soportado(t, c, "acreditamientoIEPS 1.0")
	a := c.Conceptos[0].AcreditamientoIEPS
	if a == nil {
		t.Fatal("sin acreditamientoIEPS")
	}
	igual(t, "TAR", a.TAR, "617")
	s := leerRuta(t, "sinteticos/t10/spei32.xml")
	igual(t, "TAR en 3.2", s.Conceptos[0].AcreditamientoIEPS.TAR, "619")
}

func TestComplementoSPEI(t *testing.T) {
	c := leerRuta(t, "sinteticos/t10/spei32.xml")
	sinMasQue(t, c, AdvVersionAntigua)
	soportado(t, c, "Complemento_SPEI")
	if len(c.ComplementoSPEI) != 2 {
		t.Fatalf("ComplementoSPEI = %+v", c.ComplementoSPEI)
	}
	s := c.ComplementoSPEI[0]
	if s.Ordenante == nil || s.Beneficiario == nil {
		t.Fatalf("SPEI_Tercero = %+v", s)
	}
	igual(t, "FechaOperacion", s.FechaOperacion.Original, "2016-06-01")
	igual(t, "Hora", s.Hora, "11:59:30")
	igual(t, "ClaveSPEI", s.ClaveSPEI, "40012")
	monto(t, "Cuenta con ceros a la izquierda", s.Ordenante.Cuenta, "002180000000000018")
	monto(t, "MontoPago", s.Beneficiario.MontoPago, "1160.00")
	if c.ComplementoSPEI[1].Beneficiario.IVA != nil {
		t.Errorf("IVA opcional: %+v", c.ComplementoSPEI[1].Beneficiario)
	}

	// Sello y cadenaCDA solo con IncluirSello, como los demás sellos.
	b, _ := c.JSON()
	if strings.Contains(string(b), "cadena_cda") || strings.Contains(string(b), `"sello"`) {
		t.Errorf("JSON sin IncluirSello trae el sello del SPEI: %s", b)
	}
	if c.ComplementoSPEI[0].Sello == "" {
		t.Error("el JSON sin sellos modificó el Comprobante")
	}
	b, _ = c.JSONOpciones(OpcionesJSON{IncluirSello: true})
	if !strings.Contains(string(b), `"cadena_cda":"||1|01062016|01062016|115930|40012|BANCO DE PRUEBA||"`) {
		t.Errorf("JSON con IncluirSello sin cadena_cda: %s", b)
	}

	m := leerMutado(t, "sinteticos/t10/spei32.xml", `<spei:Beneficiario BancoReceptor="OTRO BANCO" Nombre="ESCUELA KEMPER URGATE" TipoCuenta="40" Cuenta="012180000000000016" RFC="EKU9003173C9" Concepto="SEGUNDO PAGO" MontoPago="1.00"/>`, ``)
	if cs := codigos(m); cs[AdvNodoFaltante] != 1 || cs[AdvVersionAntigua] != 1 || len(cs) != 2 {
		t.Errorf("advertencias = %v", m.Advertencias)
	}
}
