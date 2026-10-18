package cfdi

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/OrtaMarco/cfdi-go/decimal"
)

// leerPublico lee un fixture de testdata/publicos/<dir>/<nombre>.
func leerPublico(t *testing.T, dir, nombre string) *Comprobante {
	t.Helper()
	b, err := os.ReadFile(filepath.Join("testdata", "publicos", dir, nombre))
	if err != nil {
		t.Fatal(err)
	}
	c, err := ParseBytes(b)
	if err != nil {
		t.Fatalf("%s: %v", nombre, err)
	}
	return c
}

func soportado(t *testing.T, c *Comprobante, nombre string) {
	t.Helper()
	if !slices.Contains(c.Complementos(), nombre) {
		t.Errorf("Complementos = %v, falta %q", c.Complementos(), nombre)
	}
	if len(c.ComplementosNoSoportados) > 0 || c.TieneAdvertencia(AdvComplementoNoSoportado) {
		t.Errorf("no soportados: %v", c.ComplementosNoSoportados)
	}
}

func TestImpuestosLocales(t *testing.T) {
	c := leer(t, "implocal40.xml")
	sinAdvertencias(t, c)
	soportado(t, c, "ImpuestosLocales 1.0")
	il := c.ImpuestosLocales
	if il == nil {
		t.Fatal("sin ImpuestosLocales")
	}
	igual(t, "Version", il.Version, "1.0")
	monto(t, "TotaldeTraslados", il.TotaldeTraslados, "50.00")
	monto(t, "TotaldeRetenciones", il.TotaldeRetenciones, "5.00")
	if il.TotaldeTraslados == nil || il.TotaldeRetenciones == nil {
		t.FailNow()
	}
	if len(il.TrasladosLocales) != 2 || len(il.RetencionesLocales) != 1 {
		t.Fatalf("%d traslados y %d retenciones locales", len(il.TrasladosLocales), len(il.RetencionesLocales))
	}
	igual(t, "ImpLocTrasladado", il.TrasladosLocales[0].ImpLocTrasladado, "ISH")
	monto(t, "TasadeTraslado", il.TrasladosLocales[0].TasadeTraslado, "3.00")
	monto(t, "Importe traslado 2", il.TrasladosLocales[1].Importe, "20.00")
	igual(t, "ImpLocRetenido", il.RetencionesLocales[0].ImpLocRetenido, "CINCO AL MILLAR")
	monto(t, "TasadeRetencion", il.RetencionesLocales[0].TasadeRetencion, "0.50")

	// Los traslados locales suman su total…
	var tras decimal.Decimal
	for _, tl := range il.TrasladosLocales {
		tras = tras.Add(*tl.Importe)
	}
	if !tras.Equal(*il.TotaldeTraslados) {
		t.Errorf("suma de traslados locales = %s", tras)
	}
	// …y cuadran el Total del comprobante.
	calc := c.SubTotal.Add(*c.Impuestos.TotalImpuestosTrasladados).
		Add(*il.TotaldeTraslados).Sub(*il.TotaldeRetenciones)
	if !calc.Equal(*c.Total) {
		t.Errorf("SubTotal + IVA + locales = %s, Total = %s", calc, c.Total)
	}

	// Público (CFDI 4.0 de satcfdi): dos traslados y dos retenciones locales.
	p := leerPublico(t, "SAT-CFDI-python-satcfdi", "tests__cfdi_ejemplos__comprobante40__cfdv40-implocal.xml")
	soportado(t, p, "ImpuestosLocales 1.0")
	pl := p.ImpuestosLocales
	if pl == nil {
		t.Fatal("público: sin ImpuestosLocales")
	}
	igual(t, "público Version", pl.Version, "1.0")
	monto(t, "público TotaldeTraslados", pl.TotaldeTraslados, "1234.56")
	monto(t, "público TotaldeRetenciones", pl.TotaldeRetenciones, "5678.90")
	if len(pl.TrasladosLocales) != 2 || len(pl.RetencionesLocales) != 2 {
		t.Fatalf("público: %d traslados y %d retenciones locales", len(pl.TrasladosLocales), len(pl.RetencionesLocales))
	}
	for i, tl := range pl.TrasladosLocales {
		igual(t, fmt.Sprintf("público ImpLocTrasladado %d", i), tl.ImpLocTrasladado, "234.56")
		monto(t, fmt.Sprintf("público TasadeTraslado %d", i), tl.TasadeTraslado, "567.89")
		monto(t, fmt.Sprintf("público Importe traslado %d", i), tl.Importe, "789.12")
	}
	for i, rl := range pl.RetencionesLocales {
		igual(t, fmt.Sprintf("público ImpLocRetenido %d", i), rl.ImpLocRetenido, "123.45")
		monto(t, fmt.Sprintf("público TasadeRetencion %d", i), rl.TasadeRetencion, "456.78")
		monto(t, fmt.Sprintf("público Importe retención %d", i), rl.Importe, "678.90")
	}
}

func TestImpuestosLocalesTolerante(t *testing.T) {
	b, err := os.ReadFile(filepath.Join("testdata", "sinteticos", "implocal40.xml"))
	if err != nil {
		t.Fatal(err)
	}
	mal := reemplazar(t, b, `TotaldeTraslados="50.00"`, `TotaldeTraslados=""`)
	mal = reemplazar(t, mal, `version="1.0" Total`, `version="9.9" Total`)
	mal = reemplazar(t, mal, `ImpLocRetenido="CINCO AL MILLAR" `, ``)
	c, err := ParseBytes(mal)
	if err != nil {
		t.Fatal(err)
	}
	if c.ImpuestosLocales == nil {
		t.Fatal("sin ImpuestosLocales")
	}
	if c.ImpuestosLocales.TotaldeTraslados != nil {
		t.Errorf("TotaldeTraslados vacío se leyó como %v", c.ImpuestosLocales.TotaldeTraslados)
	}
	m := codigos(c)
	if m[AdvDecimalInvalido] != 1 || m[AdvComplementoVersion] != 1 || m[AdvAtributoFaltante] != 1 || len(m) != 3 {
		t.Errorf("advertencias = %v", c.Advertencias)
	}
	monto(t, "lo demás se lee", c.ImpuestosLocales.TrasladosLocales[1].Importe, "20.00")
}

func TestComplementoDuplicado(t *testing.T) {
	b := sintetico(t, "implocal40.xml")
	otro := `<implocal:ImpuestosLocales version="1.0" TotaldeRetenciones="0.00" TotaldeTraslados="1.00"/>`
	mal := reemplazar(t, b, `<tfd:TimbreFiscalDigital`, otro+`<tfd:TimbreFiscalDigital`)
	c, err := ParseBytes(mal)
	if err != nil {
		t.Fatal(err)
	}
	if m := codigos(c); m[AdvNodoDuplicado] != 1 || len(m) != 1 {
		t.Errorf("advertencias = %+v", c.Advertencias)
	}
	monto(t, "se queda el primero", c.ImpuestosLocales.TotaldeTraslados, "50.00")
}

// reemplazar cambia la primera aparición de viejo; falla si no está.
func reemplazar(t *testing.T, b []byte, viejo, nuevo string) []byte {
	t.Helper()
	if !bytes.Contains(b, []byte(viejo)) {
		t.Fatalf("no se encontró %q", viejo)
	}
	return bytes.Replace(b, []byte(viejo), []byte(nuevo), 1)
}

// quitarEntre borra desde desde (incluido) hasta hasta (excluido).
func quitarEntre(t *testing.T, b []byte, desde, hasta string) []byte {
	t.Helper()
	i, j := bytes.Index(b, []byte(desde)), bytes.Index(b, []byte(hasta))
	if i < 0 || j < i {
		t.Fatalf("no se encontró %q … %q", desde, hasta)
	}
	return append(b[:i:i], b[j:]...)
}

func sintetico(t *testing.T, nombre string) []byte {
	t.Helper()
	b, err := os.ReadFile(filepath.Join("testdata", "sinteticos", nombre))
	if err != nil {
		t.Fatal(err)
	}
	return b
}

// sinAdvertenciasEn exige que ninguna advertencia caiga bajo la ruta dada.
func sinAdvertenciasEn(t *testing.T, c *Comprobante, ruta string) {
	t.Helper()
	for _, a := range c.Advertencias {
		if strings.HasPrefix(a.Ruta, ruta) {
			t.Errorf("advertencia inesperada: %+v", a)
		}
	}
}

// distancia suma DistanciaRecorrida de las ubicaciones.
func distancia(us []Ubicacion) decimal.Decimal {
	var d decimal.Decimal
	for _, u := range us {
		if u.DistanciaRecorrida != nil {
			d = d.Add(*u.DistanciaRecorrida)
		}
	}
	return d
}

func TestCartaPorte31(t *testing.T) {
	c := leer(t, "cartaporte31.xml")
	sinAdvertencias(t, c)
	soportado(t, c, "CartaPorte 3.1")
	cp := c.CartaPorte
	if cp == nil {
		t.Fatal("sin CartaPorte")
	}
	igual(t, "Version", cp.Version, "3.1")
	igual(t, "IdCCP", cp.IdCCP, "CCC00001-0000-4000-8000-000000000031")
	igual(t, "RegistroISTMO", cp.RegistroISTMO, "Sí")
	if len(cp.RegimenesAduaneros) != 2 || cp.RegimenesAduaneros[1].RegimenAduanero != "ITR" || cp.RegimenAduanero != "" {
		t.Errorf("regímenes = %+v / %q", cp.RegimenesAduaneros, cp.RegimenAduanero)
	}

	if len(cp.Ubicaciones) != 3 {
		t.Fatalf("ubicaciones = %d, quiero 3", len(cp.Ubicaciones))
	}
	monto(t, "TotalDistRec", cp.TotalDistRec, "230.75")
	if d := distancia(cp.Ubicaciones); !d.Equal(*cp.TotalDistRec) {
		t.Errorf("suma de DistanciaRecorrida = %s", d)
	}
	o := cp.Ubicaciones[0]
	igual(t, "origen TipoUbicacion", o.TipoUbicacion, "Origen")
	igual(t, "RFC normalizado", o.RFCRemitenteDestinatario, "EKU9003173C9")
	igual(t, "FechaHoraSalidaLlegada", o.FechaHoraSalidaLlegada.Original, "2026-04-06T08:00:00")
	igual(t, "origen CP", o.Domicilio.CodigoPostal, "42501")
	igual(t, "destino extranjero", cp.Ubicaciones[2].ResidenciaFiscal, "USA")

	ms := cp.Mercancias
	monto(t, "PesoBrutoTotal", ms.PesoBrutoTotal, "3000.000")
	igual(t, "NumTotalMercancias", ms.NumTotalMercancias, "2")
	if len(ms.Mercancia) != 2 {
		t.Fatalf("mercancías = %d", len(ms.Mercancia))
	}
	m := ms.Mercancia[0]
	igual(t, "BienesTransp", m.BienesTransp, "31181701")
	monto(t, "PesoEnKg", m.PesoEnKg, "2000.000")
	igual(t, "FraccionArancelaria", m.FraccionArancelaria, "8708999999")
	igual(t, "NumPedimento", m.DocumentacionAduanera[0].NumPedimento, "26  47  3807  6000031")
	if len(m.CantidadTransporta) != 2 || m.CantidadTransporta[1].IDDestino != "DE000003" {
		t.Errorf("CantidadTransporta = %+v", m.CantidadTransporta)
	}
	monto(t, "CantidadTransporta[2]", m.CantidadTransporta[1].Cantidad, "40")
	igual(t, "IdentDocAduanero", ms.Mercancia[1].DocumentacionAduanera[0].IdentDocAduanero, "DOC-0031")
	// Los pesos de las mercancías suman el total.
	if p := sumar(ms.Mercancia[0].PesoEnKg, ms.Mercancia[1].PesoEnKg); !p.Equal(*ms.PesoBrutoTotal) {
		t.Errorf("suma de PesoEnKg = %s", p)
	}

	at := ms.Autotransporte
	igual(t, "PlacaVM", at.IdentificacionVehicular.PlacaVM, "XYZ987")
	monto(t, "PesoBrutoVehicular", at.IdentificacionVehicular.PesoBrutoVehicular, "12.00")
	igual(t, "PolizaRespCivil", at.Seguros.PolizaRespCivil, "POL-31")

	// En 3.1 los RemolquesCCP van dentro de cada contenedor.
	tm := ms.TransporteMaritimo
	if tm.RemolquesCCP != nil || len(tm.Contenedor) != 1 || len(tm.Contenedor[0].RemolquesCCP) != 2 {
		t.Fatalf("remolques CCP: marítimo %v, contenedores %+v", tm.RemolquesCCP, tm.Contenedor)
	}
	igual(t, "PlacaCCP", tm.Contenedor[0].RemolquesCCP[1].PlacaCCP, "RCCP32")

	if len(cp.FiguraTransporte) != 2 {
		t.Fatalf("figuras = %d", len(cp.FiguraTransporte))
	}
	igual(t, "RFCFigura", cp.FiguraTransporte[0].RFCFigura, "CACX7605101P8")
	igual(t, "ParteTransporte", cp.FiguraTransporte[1].PartesTransporte[0].ParteTransporte, "PT01")
	igual(t, "Domicilio figura", cp.FiguraTransporte[1].Domicilio.Estado, "HID")
}

func TestCartaPorte30(t *testing.T) {
	c := leer(t, "cartaporte30.xml")
	sinAdvertencias(t, c)
	soportado(t, c, "CartaPorte 3.0")
	cp := c.CartaPorte
	if cp == nil {
		t.Fatal("sin CartaPorte")
	}
	igual(t, "Version", cp.Version, "3.0")
	igual(t, "RegimenAduanero", cp.RegimenAduanero, "IMD")
	if cp.RegimenesAduaneros != nil {
		t.Errorf("RegimenesAduaneros en 3.0: %v", cp.RegimenesAduaneros)
	}
	if len(cp.Ubicaciones) != 2 {
		t.Fatalf("ubicaciones = %d", len(cp.Ubicaciones))
	}
	monto(t, "TotalDistRec", cp.TotalDistRec, "310.40")
	if d := distancia(cp.Ubicaciones); !d.Equal(*cp.TotalDistRec) {
		t.Errorf("suma de DistanciaRecorrida = %s", d)
	}
	m := cp.Mercancias.Mercancia[0]
	igual(t, "SectorCOFEPRIS", m.SectorCOFEPRIS, "01")
	igual(t, "FechaCaducidad", m.FechaCaducidad.Original, "2027-12-31")
	if !m.FechaCaducidad.Valida() {
		t.Error("FechaCaducidad no se interpretó")
	}
	monto(t, "ValorMercancia", m.ValorMercancia, "98765.43")
	igual(t, "UUIDComercioExt", m.UUIDComercioExt, "11111111-2222-4333-8444-555555555530")
	monto(t, "PesoTara", m.DetalleMercancia.PesoTara, "50.500")
	igual(t, "RFCImpo", m.DocumentacionAduanera[0].RFCImpo, "EKU9003173C9")
	igual(t, "Remolque", cp.Mercancias.Autotransporte.Remolques[0].Placa, "REM0030")
	// En 3.0 los RemolquesCCP cuelgan del transporte marítimo.
	tm := cp.Mercancias.TransporteMaritimo
	if len(tm.RemolquesCCP) != 1 || tm.Contenedor[0].RemolquesCCP != nil {
		t.Errorf("remolques CCP: %+v / %+v", tm.RemolquesCCP, tm.Contenedor)
	}
	monto(t, "Puntal", tm.Puntal, "5.25")
	igual(t, "FechaCertificacionCCP", tm.Contenedor[0].FechaCertificacionCCP.Original, "2026-03-01T12:00:00")
	igual(t, "NombreFigura", cp.FiguraTransporte[0].NombreFigura, "OPERADOR DE PRUEBA")
}

func TestCartaPorte20(t *testing.T) {
	c := leerPublico(t, "SAT-CFDI-python-satcfdi", "tests__cfdi_ejemplos__comprobante40__cfdv40-cartaporte20.xml")
	soportado(t, c, "CartaPorte 2.0")
	sinAdvertenciasEn(t, c, "/Comprobante/Complemento/CartaPorte")
	cp := c.CartaPorte
	if cp == nil {
		t.Fatal("sin CartaPorte")
	}
	igual(t, "Version", cp.Version, "2.0")
	igual(t, "IdCCP (no existe en 2.0)", cp.IdCCP, "")
	if len(cp.Ubicaciones) != 2 {
		t.Fatalf("ubicaciones = %d", len(cp.Ubicaciones))
	}
	monto(t, "TotalDistRec", cp.TotalDistRec, "123.9")
	monto(t, "DistanciaRecorrida", cp.Ubicaciones[1].DistanciaRecorrida, "123.45")
	ms := cp.Mercancias
	monto(t, "CargoPorTasacion", ms.CargoPorTasacion, "10.123456")
	m := ms.Mercancia[0]
	igual(t, "Pedimento", m.Pedimentos[0].Pedimento, "12  34  5678  1234567")
	if m.DocumentacionAduanera != nil {
		t.Errorf("DocumentacionAduanera en 2.0: %v", m.DocumentacionAduanera)
	}
	igual(t, "NumPiezas", m.DetalleMercancia.NumPiezas, "42")
	igual(t, "NumCertITC", ms.TransporteMaritimo.NumCertITC, "NumCertITC")
	igual(t, "MatriculaContenedor", ms.TransporteMaritimo.Contenedor[0].MatriculaContenedor, "ABC45678901")
	igual(t, "CodigoTransportista", ms.TransporteAereo.CodigoTransportista, "CA156")
	tf := ms.TransporteFerroviario
	monto(t, "KilometrajePagado", tf.DerechosDePaso[0].KilometrajePagado, "12.3")
	monto(t, "PesoContenedorVacio", tf.Carro[0].Contenedor[0].PesoContenedorVacio, "9.998")
	igual(t, "Remolque", ms.Autotransporte.Remolques[0].SubTipoRem, "CTR032")

	// El mínimo tampoco advierte nada en la carta porte.
	min := leerPublico(t, "SAT-CFDI-python-satcfdi", "tests__cfdi_ejemplos__comprobante40__cfdv40-cartaporte20-min.xml")
	soportado(t, min, "CartaPorte 2.0")
	sinAdvertenciasEn(t, min, "/Comprobante/Complemento/CartaPorte")
}

func TestComercioExterior20(t *testing.T) {
	c := leer(t, "cce20.xml")
	sinAdvertencias(t, c)
	soportado(t, c, "ComercioExterior 2.0")
	ce := c.ComercioExterior
	if ce == nil {
		t.Fatal("sin ComercioExterior")
	}
	igual(t, "Version", ce.Version, "2.0")
	igual(t, "ClaveDePedimento", ce.ClaveDePedimento, "A1")
	igual(t, "CertificadoOrigen", ce.CertificadoOrigen, "0")
	igual(t, "Incoterm", ce.Incoterm, "FCA")
	igual(t, "TipoOperacion (no existe en 2.0)", ce.TipoOperacion, "")
	monto(t, "TipoCambioUSD", ce.TipoCambioUSD, "18.5000")
	monto(t, "TotalUSD", ce.TotalUSD, "1250.5000")

	if len(ce.Mercancias) != 2 {
		t.Fatalf("mercancías = %d", len(ce.Mercancias))
	}
	m := ce.Mercancias[0]
	igual(t, "NoIdentificacion", m.NoIdentificacion, "EXP-001")
	igual(t, "FraccionArancelaria", m.FraccionArancelaria, "8471300100")
	monto(t, "ValorDolares", m.ValorDolares, "1000.0000")
	monto(t, "ValorUnitarioAduana", m.ValorUnitarioAduana, "500.000000")
	igual(t, "Marca", m.DescripcionesEspecificas[0].Marca, "MARCA")
	igual(t, "NumeroSerie", m.DescripcionesEspecificas[0].NumeroSerie, "SERIE-0001")
	igual(t, "FraccionArancelaria 2", ce.Mercancias[1].FraccionArancelaria, "9403600100")
	// El valor en dólares de las mercancías suma TotalUSD.
	if v := sumar(ce.Mercancias[0].ValorDolares, ce.Mercancias[1].ValorDolares); !v.Equal(*ce.TotalUSD) {
		t.Errorf("suma de ValorDolares = %s", v)
	}
	// Y cada mercancía corresponde a un concepto por NoIdentificacion.
	for i, m := range ce.Mercancias {
		igual(t, "NoIdentificacion = concepto", m.NoIdentificacion, c.Conceptos[i].NoIdentificacion)
	}

	igual(t, "Emisor CP", ce.Emisor.Domicilio.CodigoPostal, "42501")
	igual(t, "Receptor Pais", ce.Receptor.Domicilio.Pais, "USA")
	if len(ce.Destinatario) != 1 || len(ce.Destinatario[0].Domicilio) != 2 {
		t.Fatalf("destinatarios = %+v", ce.Destinatario)
	}
	igual(t, "Destinatario domicilio 2", ce.Destinatario[0].Domicilio[1].Estado, "CA")
	if c.CartaPorte != nil {
		t.Error("ComercioExterior se leyó también como CartaPorte")
	}
}

func TestComercioExterior11(t *testing.T) {
	c := leerPublico(t, "SAT-CFDI-python-satcfdi", "tests__cfdi_ejemplos__comprobante40__cfdv40-cce11.xml")
	soportado(t, c, "ComercioExterior 1.1")
	sinAdvertenciasEn(t, c, "/Comprobante/Complemento/ComercioExterior")
	ce := c.ComercioExterior
	if ce == nil {
		t.Fatal("sin ComercioExterior")
	}
	igual(t, "Version", ce.Version, "1.1")
	igual(t, "TipoOperacion", ce.TipoOperacion, "2")
	igual(t, "Subdivision", ce.Subdivision, "1")
	monto(t, "TotalUSD", ce.TotalUSD, "123.45")
	monto(t, "TipoCambioUSD", ce.TipoCambioUSD, "234.5678")
	igual(t, "Curp", ce.Emisor.Curp, "PEPJ800101HDFBCDQ8")
	igual(t, "Emisor Colonia", ce.Emisor.Domicilio.Colonia, "0001")
	if len(ce.Propietario) != 1 || ce.Propietario[0].NumRegIdTrib != "1234567" {
		t.Errorf("Propietario = %+v", ce.Propietario)
	}
	igual(t, "Receptor NumRegIdTrib", ce.Receptor.NumRegIdTrib, "234567890")
	igual(t, "Destinatario Nombre", ce.Destinatario[0].Nombre, "Nombre")
	if len(ce.Mercancias) != 2 {
		t.Fatalf("mercancías = %d", len(ce.Mercancias))
	}
	igual(t, "FraccionArancelaria", ce.Mercancias[1].FraccionArancelaria, "99999999")
	igual(t, "NoIdentificacion", ce.Mercancias[1].NoIdentificacion, "n101")
	monto(t, "ValorDolares", ce.Mercancias[0].ValorDolares, "123.45")
	monto(t, "CantidadAduana", ce.Mercancias[0].CantidadAduana, "0.001")

	// El mismo en CFDI 3.3.
	c33 := leerPublico(t, "SAT-CFDI-python-satcfdi", "tests__cfdi_ejemplos__comprobante33__cfdv33-cce11.xml")
	soportado(t, c33, "ComercioExterior 1.1")
}

func TestComercioExteriorIncompleto(t *testing.T) {
	b := sintetico(t, "cce20.xml")
	mal := reemplazar(t, b, ` TotalUSD="1250.5000"`, ``)
	mal = reemplazar(t, mal, `ValorDolares="250.5000"`, `ValorDolares="250,50"`)
	mal = quitarEntre(t, mal, `<cce20:Domicilio Calle="Avenida de prueba"`, `</cce20:Emisor>`)
	c, err := ParseBytes(mal)
	if err != nil {
		t.Fatal(err)
	}
	if c.ComercioExterior == nil {
		t.Fatal("sin ComercioExterior")
	}
	m := codigos(c)
	if m[AdvAtributoFaltante] != 1 || m[AdvDecimalInvalido] != 1 || m[AdvNodoFaltante] != 1 || len(m) != 3 {
		t.Errorf("advertencias = %+v", c.Advertencias)
	}
	if c.ComercioExterior.Mercancias[1].ValorDolares != nil || c.ComercioExterior.TotalUSD != nil {
		t.Errorf("se inventaron montos: %+v", c.ComercioExterior)
	}
	monto(t, "lo demás se lee", c.ComercioExterior.Mercancias[0].ValorDolares, "1000.0000")
}

func TestDivisasRegistroFiscalLeyendas(t *testing.T) {
	const dir = "SAT-CFDI-python-satcfdi"
	d := leerPublico(t, dir, "tests__cfdi_ejemplos__comprobante40__cfdv40-divisas.xml")
	soportado(t, d, "Divisas 1.0")
	sinAdvertenciasEn(t, d, "/Comprobante/Complemento/Divisas")
	igual(t, "Divisas.Version", d.Divisas.Version, "1.0")
	igual(t, "TipoOperacion (tipoOperacion)", d.Divisas.TipoOperacion, "venta")

	rf := leerPublico(t, dir, "tests__cfdi_ejemplos__comprobante40__cfdv40-registrofiscal.xml")
	soportado(t, rf, "CFDIRegistroFiscal 1.0")
	igual(t, "Folio", rf.CFDIRegistroFiscal.Folio, "1234567890123456")

	lf := leerPublico(t, dir, "tests__cfdi_ejemplos__comprobante40__cfdv40-leyendasfisc.xml")
	soportado(t, lf, "LeyendasFiscales 1.0")
	sinAdvertenciasEn(t, lf, "/Comprobante/Complemento/LeyendasFiscales")
	ls := lf.LeyendasFiscales.Leyenda
	if len(ls) != 2 {
		t.Fatalf("leyendas = %d", len(ls))
	}
	igual(t, "disposicionFiscal", ls[0].DisposicionFiscal, "disposicionFiscal")
	igual(t, "norma", ls[0].Norma, "norma")
	igual(t, "textoLeyenda", ls[1].TextoLeyenda, "textoLeyendaRequerido")
	igual(t, "sin norma", ls[1].Norma, "")
}

func TestEstadoDeCuentaCombustible12(t *testing.T) {
	c := leer(t, "ecc12.xml")
	sinAdvertencias(t, c)
	soportado(t, c, "EstadoDeCuentaCombustible 1.2")
	e := c.EstadoDeCuentaCombustible
	if e == nil {
		t.Fatal("sin EstadoDeCuentaCombustible")
	}
	igual(t, "TipoOperacion", e.TipoOperacion, "Tarjeta")
	igual(t, "NumeroDeCuenta", e.NumeroDeCuenta, "CUENTA-0012")
	monto(t, "SubTotal", e.SubTotal, "1010.00")
	monto(t, "Total", e.Total, "1171.60")
	if len(e.Conceptos) != 2 {
		t.Fatalf("conceptos = %d", len(e.Conceptos))
	}
	c1 := e.Conceptos[0]
	igual(t, "Rfc normalizado", c1.Rfc, "XIQB891116QE4")
	igual(t, "Fecha", c1.Fecha.Original, "2026-06-03T07:15:00")
	monto(t, "Cantidad", c1.Cantidad, "40.000")
	monto(t, "ValorUnitario", c1.ValorUnitario, "20.000")
	igual(t, "Traslado Impuesto", c1.Traslados[0].Impuesto, "IVA")
	monto(t, "Traslado TasaOCuota", c1.Traslados[0].TasaOCuota, "0.160000")
	// SubTotal = Σ Importe y Total = SubTotal + Σ traslados.
	var sub, tras decimal.Decimal
	for _, cp := range e.Conceptos {
		sub = sub.Add(*cp.Importe)
		for _, t := range cp.Traslados {
			tras = tras.Add(*t.Importe)
		}
	}
	if !sub.Equal(*e.SubTotal) || !sub.Add(tras).Equal(*e.Total) {
		t.Errorf("Σ importes = %s, Σ traslados = %s", sub, tras)
	}
	// El estado de cuenta no se mezcla con los conceptos del CFDI.
	if len(c.Conceptos) != 1 {
		t.Errorf("conceptos del CFDI = %d", len(c.Conceptos))
	}

	p := leerPublico(t, "SAT-CFDI-python-satcfdi", "tests__cfdi_ejemplos__comprobante40__cfdv40-ecc12.xml")
	soportado(t, p, "EstadoDeCuentaCombustible 1.2")
	sinAdvertenciasEn(t, p, "/Comprobante/Complemento/EstadoDeCuentaCombustible")
	tr := p.EstadoDeCuentaCombustible.Conceptos[0].Traslados
	if len(tr) != 2 || tr[1].Impuesto != "IEPS" {
		t.Fatalf("traslados = %+v", tr)
	}
	monto(t, "IEPS TasaOCuota", tr[1].TasaOCuota, "11.00")
}

func TestTercerosYVentaVehiculos(t *testing.T) {
	c := leer(t, "terceros33.xml")
	sinAdvertencias(t, c)
	soportado(t, c, "PorCuentadeTerceros 1.1")
	soportado(t, c, "VentaVehiculos 1.0")
	if len(c.Conceptos) != 2 {
		t.Fatalf("conceptos = %d", len(c.Conceptos))
	}
	if c.Conceptos[0].ComplementoConcepto != nil || c.Conceptos[1].ComplementoConcepto != nil {
		t.Errorf("complementos de concepto no soportados: %v", c.Conceptos[0].ComplementoConcepto)
	}

	pt := c.Conceptos[0].PorCuentadeTerceros
	if pt == nil {
		t.Fatal("sin PorCuentadeTerceros")
	}
	igual(t, "Version", pt.Version, "1.1")
	igual(t, "rfc normalizado", pt.Rfc, "XIQB891116QE4")
	igual(t, "nombre", pt.Nombre, "TERCERO DE PRUEBA")
	igual(t, "noExterior", pt.InformacionFiscalTercero.NoExterior, "5")
	igual(t, "codigoPostal", pt.InformacionFiscalTercero.CodigoPostal, "42501")
	if pt.InformacionAduanera != nil || len(pt.Parte) != 1 {
		t.Fatalf("choice InformacionAduanera/Parte: %+v / %+v", pt.InformacionAduanera, pt.Parte)
	}
	p := pt.Parte[0]
	monto(t, "Parte cantidad", p.Cantidad, "2")
	monto(t, "Parte importe", p.Importe, "500.000000")
	igual(t, "Parte aduana", p.InformacionAduanera[0].Aduana, "PACHUCA")
	igual(t, "Parte fecha", p.InformacionAduanera[0].Fecha.Original, "2021-06-15")
	igual(t, "CuentaPredial", pt.CuentaPredial.Numero, "CP-0033")
	if len(pt.Impuestos.Retenciones) != 1 || len(pt.Impuestos.Traslados) != 1 {
		t.Fatalf("impuestos = %+v", pt.Impuestos)
	}
	igual(t, "Retencion impuesto", pt.Impuestos.Retenciones[0].Impuesto, "ISR")
	if pt.Impuestos.Retenciones[0].Tasa != nil {
		t.Error("retención con tasa")
	}
	monto(t, "Traslado tasa", pt.Impuestos.Traslados[0].Tasa, "16.00")
	monto(t, "Traslado importe", pt.Impuestos.Traslados[0].Importe, "160.00")
	// El traslado del tercero es el del concepto.
	if !pt.Impuestos.Traslados[0].Importe.Equal(*c.Conceptos[0].Impuestos.Traslados[0].Importe) {
		t.Error("el IVA del tercero no es el del concepto")
	}

	vv := c.Conceptos[1].VentaVehiculos
	if vv == nil {
		t.Fatal("sin VentaVehiculos")
	}
	igual(t, "VentaVehiculos 1.0", vv.Version, "1.0")
	igual(t, "ClaveVehicular", vv.ClaveVehicular, "0000001")
	igual(t, "Niv (no existe en 1.0)", vv.Niv, "")
	monto(t, "Parte importe", vv.Parte[0].Importe, "100.00")
	igual(t, "Parte pedimento", vv.Parte[0].InformacionAduanera[0].Numero, "21  47  3807  6000034")

	// Público 1.1, con Niv e InformacionAduanera directa.
	pub := leerPublico(t, "SAT-CFDI-python-satcfdi", "tests__cfdi_ejemplos__comprobante40__cfdv40-venta11.xml")
	soportado(t, pub, "VentaVehiculos 1.1")
	sinAdvertenciasEn(t, pub, "/Comprobante/Conceptos/Concepto[1]/ComplementoConcepto")
	v11 := pub.Conceptos[0].VentaVehiculos
	igual(t, "Niv", v11.Niv, "0588888")
	igual(t, "aduana", v11.InformacionAduanera[0].Aduana, "VERACRUZ")
	if !v11.InformacionAduanera[0].Fecha.Valida() {
		t.Error("fecha aduanera no se interpretó")
	}
	if pub.Conceptos[1].VentaVehiculos != nil {
		t.Error("el segundo concepto no trae VentaVehiculos")
	}
}

func TestHidroYPetro(t *testing.T) {
	b := sintetico(t, "hidroypetro40.xml")
	c, err := ParseBytes(b)
	if err != nil {
		t.Fatal(err)
	}
	sinAdvertencias(t, c)
	soportado(t, c, "HidroYPetro 1.0")
	if n := strings.Count(strings.Join(c.Complementos(), ","), "HidroYPetro"); n != 1 {
		t.Errorf("HidroYPetro repetido en Complementos: %v", c.Complementos())
	}
	for i, want := range []string{"SP16", "SP19"} {
		h := c.Conceptos[i].HidroYPetro
		if h == nil {
			t.Fatalf("concepto %d sin HidroYPetro", i+1)
		}
		igual(t, "SubProductoHYP", h.SubProductoHYP, want)
		igual(t, "ClaveHYP = ClaveProdServ", h.ClaveHYP, c.Conceptos[i].ClaveProdServ)
		igual(t, "TipoPermiso", h.TipoPermiso, "PER51")
		igual(t, "NumeroPermiso", h.NumeroPermiso, "PL/00000/EXP/ES/2026")
	}
	if c.HidroYPetro != nil {
		t.Errorf("HidroYPetro a nivel comprobante: %+v", c.HidroYPetro)
	}

	// Fuera de lugar, en el Complemento del comprobante, también se lee.
	mal := reemplazar(t, b, `<cfdi:Complemento>`, `<cfdi:Complemento><hyp:HidroYPetro Version="1.0" TipoPermiso="PER51" NumeroPermiso="P2" ClaveHYP="15101514" SubProductoHYP="SP16"/>`)
	c2, err := ParseBytes(mal)
	if err != nil {
		t.Fatal(err)
	}
	sinAdvertencias(t, c2)
	if len(c2.HidroYPetro) != 1 || c2.HidroYPetro[0].NumeroPermiso != "P2" {
		t.Errorf("HidroYPetro del comprobante = %+v", c2.HidroYPetro)
	}
}

func TestComplementoConceptoIncompleto(t *testing.T) {
	b := sintetico(t, "terceros33.xml")
	mal := reemplazar(t, b, ` rfc="xiqb891116qe4"`, ``)
	mal = reemplazar(t, mal, `tasa="16.00"`, `tasa="16 %"`)
	c, err := ParseBytes(mal)
	if err != nil {
		t.Fatal(err)
	}
	m := codigos(c)
	if m[AdvAtributoFaltante] != 1 || m[AdvDecimalInvalido] != 1 || len(m) != 2 {
		t.Errorf("advertencias = %+v", c.Advertencias)
	}
	pt := c.Conceptos[0].PorCuentadeTerceros
	if pt.Impuestos.Traslados[0].Tasa != nil {
		t.Errorf("tasa ilegible se leyó como %v", pt.Impuestos.Traslados[0].Tasa)
	}
	monto(t, "lo demás se lee", pt.Impuestos.Traslados[0].Importe, "160.00")

	// Sin el nodo Impuestos, que es obligatorio.
	sinImp := quitarEntre(t, b, `<terceros:Impuestos>`, `        </terceros:PorCuentadeTerceros>`)
	c2, err := ParseBytes(sinImp)
	if err != nil {
		t.Fatal(err)
	}
	if m := codigos(c2); m[AdvNodoFaltante] != 1 || len(m) != 1 {
		t.Errorf("advertencias = %+v", c2.Advertencias)
	}
	if pt := c2.Conceptos[0].PorCuentadeTerceros; pt.Impuestos != nil || pt.CuentaPredial == nil {
		t.Errorf("PorCuentadeTerceros = %+v", pt)
	}

	// Un complemento de concepto desconocido sigue reportándose.
	desc := reemplazar(t, b, `<ventavehiculos:VentaVehiculos version="1.0"`, `<ventavehiculos:Otro version="1.0"`)
	desc = reemplazar(t, desc, `</ventavehiculos:VentaVehiculos>`, `</ventavehiculos:Otro>`)
	c3, err := ParseBytes(desc)
	if err != nil {
		t.Fatal(err)
	}
	want := "{" + NSVentaVehiculos + "}Otro"
	if c3.Conceptos[1].VentaVehiculos != nil || !slices.Equal(c3.Conceptos[1].ComplementoConcepto, []string{want}) ||
		!slices.Equal(c3.ComplementosNoSoportados, []string{want}) || codigos(c3)[AdvComplementoNoSoportado] != 1 {
		t.Errorf("desconocido: %+v / %v", c3.Conceptos[1], c3.ComplementosNoSoportados)
	}
}

func TestCartaPorteIncompleta(t *testing.T) {
	b := sintetico(t, "cartaporte31.xml")
	// Sin IdCCP, con una sola ubicación y en el namespace 3.1 diciendo 3.0.
	mal := reemplazar(t, b, `Version="3.1" IdCCP="CCC00001-0000-4000-8000-000000000031" `, `Version="3.0" `)
	mal = quitarEntre(t, mal, `<cartaporte31:Ubicacion TipoUbicacion="Destino" IDUbicacion="DE000002"`, `</cartaporte31:Ubicaciones>`)
	c, err := ParseBytes(mal)
	if err != nil {
		t.Fatal(err)
	}
	if c.CartaPorte == nil {
		t.Fatal("sin CartaPorte")
	}
	igual(t, "Version tal cual", c.CartaPorte.Version, "3.0")
	m := codigos(c)
	if m[AdvComplementoVersion] != 1 || m[AdvAtributoFaltante] != 1 || m[AdvNodoFaltante] != 1 || len(m) != 3 {
		t.Errorf("advertencias = %+v", c.Advertencias)
	}
	if len(c.CartaPorte.Ubicaciones) != 1 || c.CartaPorte.Mercancias == nil {
		t.Errorf("lo demás no se leyó: %+v", c.CartaPorte)
	}
}
