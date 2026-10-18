package cfdi

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Complementos que entraron en T10: lectura de un fixture público o
// sintético por complemento, más los casos tolerantes.

// leerRuta lee un fixture por su ruta bajo testdata/.
func leerRuta(t *testing.T, ruta string) *Comprobante {
	t.Helper()
	b, err := os.ReadFile(filepath.Join("testdata", ruta))
	if err != nil {
		t.Fatal(err)
	}
	c, err := ParseBytes(b)
	if err != nil {
		t.Fatalf("%s: %v", ruta, err)
	}
	return c
}

// leerMutado lee un fixture después de reemplazar viejo por nuevo (una vez).
func leerMutado(t *testing.T, ruta string, cambios ...string) *Comprobante {
	t.Helper()
	b, err := os.ReadFile(filepath.Join("testdata", ruta))
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i+1 < len(cambios); i += 2 {
		b = reemplazar(t, b, cambios[i], cambios[i+1])
	}
	c, err := ParseBytes(b)
	if err != nil {
		t.Fatal(err)
	}
	return c
}

const satcfdi40 = "publicos/SAT-CFDI-python-satcfdi/tests__cfdi_ejemplos__comprobante40__cfdv40-"

func TestAerolineas(t *testing.T) {
	c := leerRuta(t, "mutantes/base-aerolineas.xml")
	sinAdvertencias(t, c)
	soportado(t, c, "Aerolineas 1.0")
	a := c.Aerolineas
	if a == nil || a.OtrosCargos == nil || len(a.OtrosCargos.Cargo) != 1 {
		t.Fatalf("Aerolineas = %+v", a)
	}
	igual(t, "Version", a.Version, "1.0")
	monto(t, "TUA", a.TUA, "150.00")
	monto(t, "TotalCargos", a.OtrosCargos.TotalCargos, "80.00")
	igual(t, "CodigoCargo", a.OtrosCargos.Cargo[0].CodigoCargo, "XV")
	monto(t, "Importe", a.OtrosCargos.Cargo[0].Importe, "80.00")

	p := leerRuta(t, satcfdi40+"aerolineas.xml")
	soportado(t, p, "Aerolineas 1.0")
	monto(t, "público TUA", p.Aerolineas.TUA, "0")

	// Tolerante: TUA ilegible, OtrosCargos sin Cargo.
	m := leerMutado(t, "mutantes/base-aerolineas.xml",
		`TUA="150.00"`, `TUA="mucho"`,
		`<aerolineas:Cargo CodigoCargo="XV" Importe="80.00"/>`, ``)
	if m.Aerolineas == nil || m.Aerolineas.TUA != nil {
		t.Fatalf("Aerolineas = %+v", m.Aerolineas)
	}
	if cs := codigos(m); cs[AdvDecimalInvalido] != 1 || cs[AdvNodoFaltante] != 1 || len(cs) != 2 {
		t.Errorf("advertencias = %v", m.Advertencias)
	}
	monto(t, "lo demás se lee", m.Aerolineas.OtrosCargos.TotalCargos, "80.00")
}

func TestDonatarias(t *testing.T) {
	c := leerRuta(t, "mutantes/base-donat11.xml")
	sinAdvertencias(t, c)
	soportado(t, c, "Donatarias 1.1")
	d := c.Donatarias
	if d == nil {
		t.Fatal("sin Donatarias")
	}
	igual(t, "NoAutorizacion", d.NoAutorizacion, "123")
	igual(t, "FechaAutorizacion", d.FechaAutorizacion.Original, "2026-01-02")
	igual(t, "Leyenda", d.Leyenda, "Leyenda de prueba")

	p := leerRuta(t, satcfdi40+"donat11.xml")
	soportado(t, p, "Donatarias 1.1")
	igual(t, "público NoAutorizacion", p.Donatarias.NoAutorizacion, "123456ABC")
	igual(t, "público FechaAutorizacion", p.Donatarias.FechaAutorizacion.Original, "1973-04-04")

	// 1.0 comparte namespace: se lee sin advertencia; otra versión, con
	// COMPLEMENTO_VERSION; sin leyenda, ATRIBUTO_FALTANTE.
	v10 := leerMutado(t, "mutantes/base-donat11.xml", `version="1.1"`, `version="1.0"`)
	sinAdvertencias(t, v10)
	soportado(t, v10, "Donatarias 1.0")
	mal := leerMutado(t, "mutantes/base-donat11.xml", `version="1.1"`, `version="2.0"`, ` leyenda="Leyenda de prueba"`, ``)
	if cs := codigos(mal); cs[AdvComplementoVersion] != 1 || cs[AdvAtributoFaltante] != 1 || len(cs) != 2 {
		t.Errorf("advertencias = %v", mal.Advertencias)
	}
}

func TestInstEducativas(t *testing.T) {
	c := leerRuta(t, "mutantes/base-iedu.xml")
	sinAdvertencias(t, c)
	soportado(t, c, "instEducativas 1.0")
	ie := c.Conceptos[0].InstEducativas
	if ie == nil {
		t.Fatal("sin instEducativas en el primer concepto")
	}
	if c.Conceptos[0].ComplementoConcepto != nil || c.Conceptos[1].InstEducativas != nil {
		t.Errorf("conceptos = %+v", c.Conceptos)
	}
	igual(t, "NombreAlumno", ie.NombreAlumno, "ALUMNO DE PRUEBA")
	igual(t, "CURP", ie.CURP, "XEXX010101HNEXXXA4")
	igual(t, "NivelEducativo", ie.NivelEducativo, "Primaria")
	igual(t, "AutRVOE", ie.AutRVOE, "RVOE0001")
	igual(t, "RfcPago", ie.RfcPago, "XOJI740919U48")

	p := leerRuta(t, satcfdi40+"iedu.xml")
	soportado(t, p, "instEducativas 1.0")
	igual(t, "público CURP", p.Conceptos[0].InstEducativas.CURP, "JUAN010101MAZBCÑ90")

	// rfcPago se normaliza como los demás RFC; sin CURP, ATRIBUTO_FALTANTE.
	m := leerMutado(t, "mutantes/base-iedu.xml", `rfcPago="XOJI740919U48"`, `rfcPago=" xoji740919u48 "`, ` CURP="XEXX010101HNEXXXA4"`, ``)
	igual(t, "RfcPago normalizado", m.Conceptos[0].InstEducativas.RfcPago, "XOJI740919U48")
	if cs := codigos(m); cs[AdvAtributoFaltante] != 1 || len(cs) != 1 {
		t.Errorf("advertencias = %v", m.Advertencias)
	}
}

// G2: versiones viejas de complementos que ya se leían.

func TestCartaPorte10(t *testing.T) {
	c := leerRuta(t, "sinteticos/t10/cartaporte10.xml")
	sinAdvertencias(t, c)
	soportado(t, c, "CartaPorte 1.0")
	cp := c.CartaPorte
	if cp == nil || len(cp.Ubicaciones) != 2 || cp.Mercancias == nil || cp.FiguraTransporte10 == nil {
		t.Fatalf("CartaPorte = %+v", cp)
	}
	igual(t, "TranspInternac", cp.TranspInternac, "No")
	monto(t, "TotalDistRec", cp.TotalDistRec, "150.50")
	o, d := cp.Ubicaciones[0].Origen, cp.Ubicaciones[1].Destino
	if o == nil || d == nil || cp.Ubicaciones[0].Destino != nil || cp.Ubicaciones[1].Origen != nil {
		t.Fatalf("Ubicaciones = %+v", cp.Ubicaciones)
	}
	igual(t, "IDOrigen", o.IDOrigen, "OR000001")
	igual(t, "RFCRemitente", o.RFCRemitente, "EKU9003173C9")
	igual(t, "FechaHoraSalida", o.FechaHoraSalida.Original, "2021-12-15T08:00:00")
	igual(t, "IDDestino", d.IDDestino, "DE000002")
	igual(t, "FechaHoraProgLlegada", d.FechaHoraProgLlegada.Original, "2021-12-15T11:00:00")
	monto(t, "DistanciaRecorrida", cp.Ubicaciones[1].DistanciaRecorrida, "150.50")
	igual(t, "Domicilio.Calle", cp.Ubicaciones[0].Domicilio.Calle, "Calle de prueba")

	ms := cp.Mercancias
	igual(t, "NumTotalMercancias", ms.NumTotalMercancias, "1")
	if len(ms.Mercancia) != 1 || ms.Mercancia[0].DetalleMercancia == nil || len(ms.Mercancia[0].CantidadTransporta) != 1 {
		t.Fatalf("Mercancia = %+v", ms.Mercancia)
	}
	monto(t, "PesoEnKg", ms.Mercancia[0].PesoEnKg, "2000.000")
	igual(t, "DetalleMercancia.UnidadPeso", ms.Mercancia[0].DetalleMercancia.UnidadPeso, "KGM")
	igual(t, "DetalleMercancia.UnidadPesoMerc", ms.Mercancia[0].DetalleMercancia.UnidadPesoMerc, "")
	af := ms.AutotransporteFederal
	if af == nil || af.IdentificacionVehicular == nil || len(af.Remolques) != 1 || ms.Autotransporte != nil {
		t.Fatalf("AutotransporteFederal = %+v, Autotransporte = %+v", af, ms.Autotransporte)
	}
	igual(t, "NombreAseg", af.NombreAseg, "ASEGURADORA DE PRUEBA")
	igual(t, "ConfigVehicular", af.IdentificacionVehicular.ConfigVehicular, "C2")
	igual(t, "Remolque.Placa", af.Remolques[0].Placa, "REM010")

	f := cp.FiguraTransporte10
	if cp.FiguraTransporte != nil || len(f.Operadores) != 1 || len(f.Propietario) != 1 || f.Operadores[0].Domicilio == nil {
		t.Fatalf("FiguraTransporte10 = %+v", f)
	}
	igual(t, "CveTransporte", f.CveTransporte, "01")
	igual(t, "RFCOperador", f.Operadores[0].RFCOperador, "CACX7605101P8")
	igual(t, "Operador.Domicilio.Calle", f.Operadores[0].Domicilio.Calle, "Calle del operador")
	igual(t, "NombrePropietario", f.Propietario[0].NombrePropietario, "PROPIETARIO DE PRUEBA")

	// Tolerante y con los requeridos de 1.0, no los de 2.0: sin PesoEnKg ni
	// CveTransporte, dos advertencias; el RFC se normaliza.
	m := leerMutado(t, "sinteticos/t10/cartaporte10.xml",
		` PesoEnKg="2000.000"`, ``, ` CveTransporte="01"`, ``, `RFCRemitente="EKU9003173C9"`, `RFCRemitente=" eku9003173c9"`)
	if cs := codigos(m); cs[AdvAtributoFaltante] != 2 || len(cs) != 1 {
		t.Errorf("advertencias = %v", m.Advertencias)
	}
	igual(t, "RFCRemitente normalizado", m.CartaPorte.Ubicaciones[0].Origen.RFCRemitente, "EKU9003173C9")
	if _, ok := m.AdvertenciaEn("/Comprobante/Complemento/CartaPorte/FiguraTransporte", "CveTransporte"); !ok {
		t.Errorf("falta la advertencia de CveTransporte: %v", m.Advertencias)
	}
}

func TestComercioExterior10(t *testing.T) {
	c := leerRuta(t, "sinteticos/t10/cce10.xml")
	sinAdvertencias(t, c)
	soportado(t, c, "ComercioExterior 1.0")
	ce := c.ComercioExterior
	if ce == nil || ce.Receptor == nil || len(ce.Destinatario) != 1 || len(ce.Mercancias) != 1 {
		t.Fatalf("ComercioExterior = %+v", ce)
	}
	igual(t, "Version", ce.Version, "1.0")
	igual(t, "TipoOperacion", ce.TipoOperacion, "2")
	igual(t, "Receptor.NumRegIdTrib", ce.Receptor.NumRegIdTrib, "123456789")
	igual(t, "Destinatario.NumRegIdTrib", ce.Destinatario[0].NumRegIdTrib, "987654321")
	if len(ce.Destinatario[0].Domicilio) != 1 {
		t.Fatalf("Destinatario.Domicilio = %+v", ce.Destinatario[0].Domicilio)
	}
	igual(t, "FraccionArancelaria", ce.Mercancias[0].FraccionArancelaria, "84713001")

	// En CFDI 3.2 (lo único que lo lleva a la cadena): Curp del Receptor y
	// Rfc del Destinatario, exclusivos de 1.0.
	p := leerRuta(t, "sinteticos/t10/cce10-32.xml")
	soportado(t, p, "ComercioExterior 1.0")
	pe := p.ComercioExterior
	igual(t, "Receptor.Curp", pe.Receptor.Curp, "XEXX010101HNEXXXA4")
	igual(t, "Destinatario.Rfc", pe.Destinatario[0].Rfc, "XEXX010101000")
	igual(t, "Domicilio.Municipio", pe.Destinatario[0].Domicilio[0].Municipio, "LAREDO")
	if len(pe.Mercancias[0].DescripcionesEspecificas) != 2 {
		t.Errorf("DescripcionesEspecificas = %+v", pe.Mercancias[0].DescripcionesEspecificas)
	}

	// 1.0 exige Receptor y su NumRegIdTrib.
	m := leerMutado(t, "sinteticos/t10/cce10.xml", `<cce:Receptor NumRegIdTrib="123456789"/>`, ``)
	if cs := codigos(m); cs[AdvNodoFaltante] != 1 || len(cs) != 1 {
		t.Errorf("advertencias = %v", m.Advertencias)
	}
	m = leerMutado(t, "sinteticos/t10/cce10.xml", `<cce:Receptor NumRegIdTrib="123456789"/>`, `<cce:Receptor/>`)
	if cs := codigos(m); cs[AdvAtributoFaltante] != 1 || len(cs) != 1 {
		t.Errorf("advertencias = %v", m.Advertencias)
	}
}

func TestEstadoDeCuentaCombustible11(t *testing.T) {
	c := leerRuta(t, "sinteticos/t10/ecc11.xml")
	sinAdvertencias(t, c)
	soportado(t, c, "EstadoDeCuentaCombustible 1.1")
	e := c.EstadoDeCuentaCombustible
	if e == nil || len(e.Conceptos) != 2 || len(e.Conceptos[0].Traslados) != 1 {
		t.Fatalf("EstadoDeCuentaCombustible = %+v", e)
	}
	cp := e.Conceptos[0]
	igual(t, "TAR", cp.TAR, "617")
	igual(t, "NoIdentificacion", cp.NoIdentificacion, "32011")
	igual(t, "TipoCombustible", cp.TipoCombustible, "")
	monto(t, "TasaoCuota", cp.Traslados[0].TasaOCuota, "0.160000")
	igual(t, "TAR opcional", e.Conceptos[1].TAR, "")

	// En 1.1 el atributo es TasaoCuota: la ruta de la advertencia lo dice
	// así; y se exige NoIdentificacion, no TipoCombustible.
	m := leerMutado(t, "sinteticos/t10/ecc11.xml", ` TasaoCuota="0.160000" Importe="128.00"`, ` Importe="128.00"`, ` NoIdentificacion="34006"`, ``)
	if cs := codigos(m); cs[AdvAtributoFaltante] != 2 || len(cs) != 1 {
		t.Errorf("advertencias = %v", m.Advertencias)
	}
	r := "/Comprobante/Complemento/EstadoDeCuentaCombustible/Conceptos/ConceptoEstadoDeCuentaCombustible[1]/Traslados/Traslado[1]"
	if _, ok := m.AdvertenciaEn(r, "TasaoCuota"); !ok {
		t.Errorf("falta la advertencia de TasaoCuota: %v", m.Advertencias)
	}
}

func TestNomina11(t *testing.T) {
	p := leerRuta(t, "publicos/SAT-CFDI-python-satcfdi/tests__cfdi_ejemplos__comprobante32__cfdv32-nomina.xml")
	soportado(t, p, "Nomina 1.1")
	if len(p.Nomina) != 1 {
		t.Fatalf("Nomina = %+v", p.Nomina)
	}
	n := p.Nomina[0]
	igual(t, "Version", n.Version, "1.1")
	igual(t, "NumEmpleado", n.NumEmpleado, "ABC123")
	igual(t, "TipoRegimen", n.TipoRegimen, "2")
	igual(t, "FechaPago", n.FechaPago.Original, "2015-01-14")
	monto(t, "NumDiasPagados", n.NumDiasPagados, "14")
	if n.Receptor != nil || n.Emisor != nil {
		t.Errorf("1.1 no llena Receptor ni Emisor: %+v %+v", n.Receptor, n.Emisor)
	}
	if n.Percepciones == nil || len(n.Percepciones.Percepcion) != 3 || n.Deducciones == nil || len(n.Deducciones.Deduccion) != 4 {
		t.Fatalf("Percepciones = %+v, Deducciones = %+v", n.Percepciones, n.Deducciones)
	}
	monto(t, "Percepciones.TotalGravado", n.Percepciones.TotalGravado, "9590.10")
	monto(t, "Deducciones.TotalExento", n.Deducciones.TotalExento, "167.66")
	monto(t, "Deduccion.ImporteExento", n.Deducciones.Deduccion[1].ImporteExento, "501.18")
	if n.Deducciones.Deduccion[1].Importe != nil {
		t.Errorf("Deduccion 1.1 con Importe: %+v", n.Deducciones.Deduccion[1])
	}

	c := leerRuta(t, "sinteticos/t10/nomina11.xml")
	sinAdvertencias(t, c)
	s := c.Nomina[0]
	igual(t, "CLABE", s.CLABE, "002215911558451272")
	igual(t, "Antiguedad", s.Antiguedad, "106")
	igual(t, "FechaInicioRelLaboral", s.FechaInicioRelLaboral.Original, "2015-01-01")
	if len(s.Incapacidades) != 1 || len(s.HorasExtras) != 1 {
		t.Fatalf("Incapacidades = %+v, HorasExtras = %+v", s.Incapacidades, s.HorasExtras)
	}
	monto(t, "Incapacidad.Descuento", s.Incapacidades[0].Descuento, "50.00")
	igual(t, "HorasExtra.TipoHoras", s.HorasExtras[0].TipoHoras, "Dobles")
	monto(t, "HorasExtra.ImportePagado", s.HorasExtras[0].ImportePagado, "100.00")

	// Las reglas de Nómina son de 1.2: un 1.1 no las dispara.
	for _, h := range ValidarReglas(c) {
		if strings.HasPrefix(h.Codigo, "REGLA_NOMINA") {
			t.Errorf("regla de nómina 1.2 sobre un 1.1: %+v", h)
		}
	}

	// Tolerante: sin CURP, una advertencia.
	m := leerMutado(t, "sinteticos/t10/nomina11.xml", ` CURP="XEXX010101HNEXXXA4"`, ``)
	if cs := codigos(m); cs[AdvAtributoFaltante] != 1 || len(cs) != 1 {
		t.Errorf("advertencias = %v", m.Advertencias)
	}
}
