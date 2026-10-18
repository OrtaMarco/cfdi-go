package cfdi

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"slices"
	"strings"
	"testing"
	"time"
)

func leerRetenciones(t *testing.T, nombre string) *Retenciones {
	t.Helper()
	b, err := os.ReadFile(filepath.Join("testdata", "sinteticos", "retenciones", nombre))
	if err != nil {
		t.Fatal(err)
	}
	rt, err := ParseRetencionesBytes(b)
	if err != nil {
		t.Fatalf("%s: %v", nombre, err)
	}
	return rt
}

// rutasRet junta las advertencias como "CODIGO ruta".
func rutasRet(rt *Retenciones) []string {
	var out []string
	for _, a := range rt.Advertencias {
		out = append(out, a.Codigo+" "+a.Ruta)
	}
	return out
}

func TestDetectar(t *testing.T) {
	casos := []struct {
		nombre string
		xml    string
		quiero TipoDocumento
	}{
		{"CFDI 4.0", string(leerSintetico(t, "ingreso40-pue.xml")), TipoComprobante},
		{"CFDI 3.3", `<c:Comprobante xmlns:c="http://www.sat.gob.mx/cfd/3"/>`, TipoComprobante},
		{"Retenciones 2.0", `<r:Retenciones xmlns:r="http://www.sat.gob.mx/esquemas/retencionpago/2"/>`, TipoRetenciones},
		{"Retenciones 1.0 con BOM", "\xef\xbb\xbf<?xml version=\"1.0\"?><Retenciones xmlns=\"http://www.sat.gob.mx/esquemas/retencionpago/1\"/>", TipoRetenciones},
		{"Retenciones no estricto", `<r:Retenciones xmlns:r="http://www.sat.gob.mx/esquemas/retencionpago/2" a="x & y"/>`, TipoRetenciones},
		{"Retenciones de otro namespace", `<Retenciones xmlns="http://www.sat.gob.mx/esquemas/retencionpago/3"/>`, TipoOtro},
		{"Comprobante sin namespace", `<Comprobante/>`, TipoOtro},
		{"otra raíz", `<a><r:Retenciones xmlns:r="http://www.sat.gob.mx/esquemas/retencionpago/2"/></a>`, TipoOtro},
		{"no es XML", "hola", TipoOtro},
		{"vacío", "", TipoOtro},
	}
	for _, c := range casos {
		if got := Detectar([]byte(c.xml)); got != c.quiero {
			t.Errorf("%s: %s, quiero %s", c.nombre, got, c.quiero)
		}
	}
}

func TestParseRetenciones20(t *testing.T) {
	rt := leerRetenciones(t, "ret20-plataformas.xml")
	if len(rt.Advertencias) > 0 {
		t.Errorf("advertencias: %v", rutasRet(rt))
	}
	igual(t, "Version", rt.Version, "2.0")
	igual(t, "FolioInt", rt.FolioInt, "R100")
	igual(t, "NoCertificado", rt.NoCertificado, "30001000000500003416")
	igual(t, "FechaExp", rt.FechaExp.Time, time.Date(2026, 1, 15, 10, 30, 0, 0, time.UTC))
	igual(t, "LugarExpRetenc", rt.LugarExpRetenc, "42501")
	igual(t, "CveRetenc", rt.CveRetenc, "26")
	igual(t, "DescRetenc", rt.DescRetenc, "Plataformas tecnológicas")
	igual(t, "relacionado", *rt.CfdiRetenRelacionados, CfdiRetenRelacionados{TipoRelacion: "04",
		UUID: "ABCDEF01-2345-4789-8ABC-DEF012345678", UUIDOriginal: "abcdef01-2345-4789-8abc-def012345678"})
	igual(t, "Emisor", *rt.Emisor, EmisorRetenciones{RfcE: "EKU9003173C9", NomDenRazSocE: "ESCUELA KEMPER URGATE", RegimenFiscalE: "601"})
	igual(t, "NacionalidadR", rt.Receptor.NacionalidadR, "Nacional")
	igual(t, "Nacional", *rt.Receptor.Nacional, ReceptorNacionalRetenciones{RfcR: "XAXX010101000", NomDenRazSocR: "PUBLICO EN GENERAL", DomicilioFiscalR: "42501"})
	if rt.Receptor.Extranjero != nil {
		t.Error("Extranjero no debía leerse")
	}
	igual(t, "Periodo", *rt.Periodo, PeriodoRetenciones{MesIni: "12", MesFin: "12", Ejercicio: "2025"})
	tot := rt.Totales
	monto(t, "MontoTotOperacion", tot.MontoTotOperacion, "1000.00")
	monto(t, "MontoTotGrav", tot.MontoTotGrav, "1000.00")
	monto(t, "MontoTotExent", tot.MontoTotExent, "0.00")
	monto(t, "MontoTotRet", tot.MontoTotRet, "100.00")
	if tot.UtilidadBimestral != nil || tot.ISRCorrespondiente != nil {
		t.Error("UtilidadBimestral e ISRCorrespondiente no venían")
	}
	if len(tot.ImpRetenidos) != 2 {
		t.Fatalf("%d ImpRetenidos", len(tot.ImpRetenidos))
	}
	monto(t, "ImpRetenidos[1].BaseRet", tot.ImpRetenidos[1].BaseRet, "160.00")
	monto(t, "ImpRetenidos[1].MontoRet", tot.ImpRetenidos[1].MontoRet, "80.00")
	igual(t, "ImpRetenidos[1].ImpuestoRet", tot.ImpRetenidos[1].ImpuestoRet, "002")
	igual(t, "ImpRetenidos[1].TipoPagoRet", tot.ImpRetenidos[1].TipoPagoRet, "01")

	p := rt.ServiciosPlataformasTecnologicas
	igual(t, "Plataformas.Version", p.Version, "1.0")
	igual(t, "NumServ", p.NumServ, "2")
	monto(t, "MonTotServSIVA", p.MonTotServSIVA, "1000.00")
	monto(t, "DifIVAEntregadoPrestServ", p.DifIVAEntregadoPrestServ, "80.00")
	monto(t, "MonTotalContribucionGubernamental", p.MonTotalContribucionGubernamental, "10.00")
	if len(p.Servicios) != 2 {
		t.Fatalf("%d servicios", len(p.Servicios))
	}
	s0, s1 := p.Servicios[0], p.Servicios[1]
	igual(t, "FechaServ", s0.FechaServ.Time, time.Date(2025, 12, 10, 0, 0, 0, 0, time.UTC))
	monto(t, "TasaCuota", s0.ImpuestosTrasladadosdelServicio.TasaCuota, "0.160000")
	monto(t, "ImpContrib", s0.ContribucionGubernamental.ImpContrib, "10.00")
	monto(t, "Porcentaje", s0.ComisionDelServicio.Porcentaje, "0.05")
	monto(t, "TasaCuota 2", s1.ImpuestosTrasladadosdelServicio.TasaCuota, "0.16")
	monto(t, "Comisión 2", s1.ComisionDelServicio.Importe, "20.00")
	if s1.ContribucionGubernamental != nil || s1.ComisionDelServicio.Base != nil {
		t.Error("el segundo servicio no trae contribución ni base de comisión")
	}
	igual(t, "UUID", rt.UUID(), "0F1E2D3C-4B5A-4978-8695-A4B3C2D1E0F9")
	igual(t, "Complementos", strings.Join(rt.Complementos(), ","), "TimbreFiscalDigital 1.1,ServiciosPlataformasTecnologicas 1.0")
}

// En 1.0 los atributos se llaman distinto; el modelo usa los nombres de 2.0.
func TestParseRetenciones10(t *testing.T) {
	rt := leerRetenciones(t, "ret10-extranjero.xml")
	if len(rt.Advertencias) > 0 {
		t.Errorf("advertencias: %v", rutasRet(rt))
	}
	igual(t, "Version", rt.Version, "1.0")
	igual(t, "NoCertificado (NumCert)", rt.NoCertificado, "30001000000500003416")
	// La zona se queda en el texto; Time es la hora tal cual.
	igual(t, "FechaExp.Original", rt.FechaExp.Original, "2020-06-30T18:00:00-06:00")
	igual(t, "FechaExp.Time", rt.FechaExp.Time, time.Date(2020, 6, 30, 18, 0, 0, 0, time.UTC))
	igual(t, "FechaExp.FueraDeFormato", rt.FechaExp.FueraDeFormato, false)
	igual(t, "Emisor", *rt.Emisor, EmisorRetenciones{RfcE: "EKU9003173C9", NomDenRazSocE: "ESCUELA KEMPER URGATE"})
	igual(t, "NacionalidadR (Nacionalidad)", rt.Receptor.NacionalidadR, "Extranjero")
	igual(t, "Extranjero", *rt.Receptor.Extranjero, ReceptorExtranjeroRetenciones{NumRegIdTribR: "998877665544", NomDenRazSocR: "EMPRESA DEL EXTRANJERO DE PRUEBA"})
	igual(t, "Periodo", *rt.Periodo, PeriodoRetenciones{MesIni: "1", MesFin: "6", Ejercicio: "2020"})
	monto(t, "MontoTotOperacion (montoTotOperacion)", rt.Totales.MontoTotOperacion, "25000.00")
	monto(t, "MontoTotRet (montoTotRet)", rt.Totales.MontoTotRet, "6250.000000")
	monto(t, "MontoRet (montoRet)", rt.Totales.ImpRetenidos[0].MontoRet, "6250.000000")
	igual(t, "ImpuestoRet (Impuesto)", rt.Totales.ImpRetenidos[0].ImpuestoRet, "01")
	igual(t, "TipoPagoRet", rt.Totales.ImpRetenidos[0].TipoPagoRet, "Pago definitivo")
	igual(t, "NoBeneficiario", *rt.Pagosaextranjeros.NoBeneficiario, NoBeneficiarioExtranjero{PaisDeResidParaEfecFisc: "US", ConceptoPago: "2", DescripcionConcepto: "Regalías de prueba"})
	pr := rt.Planesderetiro
	igual(t, "Planesderetiro.Version", pr.Version, "1.1")
	igual(t, "NumReferencia", pr.NumReferencia, "REF1")
	monto(t, "MontAportODep", pr.AportacionesODepositos[0].MontAportODep, "1000.00")
	igual(t, "TFD 1.0", rt.TimbreFiscalDigital.Version, "1.0")
	igual(t, "Complementos", strings.Join(rt.Complementos(), ","), "TimbreFiscalDigital 1.0,Pagosaextranjeros 1.0,Planesderetiro 1.1")
}

func TestParseRetencionesAnomalias(t *testing.T) {
	rt := leerRetenciones(t, "ret20-anomalias.xml")
	quiero := []string{
		"FECHA_FORMATO /Retenciones/@FechaExp",
		"NODO_FALTANTE /Retenciones/Receptor/Nacional",
		"ATRIBUTO_FALTANTE /Retenciones/Periodo/@Ejercicio",
		"DECIMAL_INVALIDO /Retenciones/Totales/@MontoTotGrav",
		"ATRIBUTO_FALTANTE /Retenciones/Totales/ImpRetenidos[1]/@TipoPagoRet",
		"NODO_DUPLICADO /Retenciones/Complemento/Intereses",
		"COMPLEMENTO_NO_SOPORTADO /Retenciones/Complemento/Raro",
		"SIN_TIMBRE /Retenciones/Complemento",
	}
	if got := rutasRet(rt); !slices.Equal(got, quiero) {
		t.Errorf("advertencias\n%q\nquiero\n%q", got, quiero)
	}
	a, ok := rt.AdvertenciaEn("/Retenciones/Totales", "MontoTotGrav")
	if !ok || a.Valor != "1,000.00" || rt.Totales.MontoTotGrav != nil {
		t.Errorf("MontoTotGrav ilegible: %+v %v", a, rt.Totales.MontoTotGrav)
	}
	igual(t, "RfcE normalizado", rt.Emisor.RfcE, "EKU9003173C9")
	igual(t, "Intereses: el primero", rt.Intereses.SistFinanciero, "SI")
	igual(t, "FechaExp con zona en 2.0", rt.FechaExp.FueraDeFormato, true)
	igual(t, "no soportados", strings.Join(rt.ComplementosNoSoportados, ","), "{urn:raro}Raro")
	if !rt.TieneAdvertencia(AdvSinTimbre) || rt.TieneAdvertencia(AdvVersionNamespace) {
		t.Error("TieneAdvertencia")
	}
}

func TestParseRetencionesVersiones(t *testing.T) {
	base := string(leerSintetico(t, "retenciones/ret20-plataformas.xml"))
	base10 := string(leerSintetico(t, "retenciones/ret10-extranjero.xml"))
	casos := []struct {
		nombre string
		xml    string
		quiero []string
	}{
		{"Version 1.0 en el namespace 2", strings.Replace(base, `Version="2.0"`, `Version="1.0"`, 1), []string{"VERSION_NAMESPACE /Retenciones/@Version"}},
		{"Version desconocida", strings.Replace(base, `Version="2.0"`, `Version="3.0"`, 1), []string{"VERSION_NO_SOPORTADA /Retenciones/@Version"}},
		{"1.0 sin zona", strings.Replace(base10, `-06:00"`, `"`, 1), []string{"FECHA_FORMATO /Retenciones/@FechaExp"}},
		{"1.0 con fracción", strings.Replace(base10, `18:00:00-06:00"`, `18:00:00.5-06:00"`, 1), []string{"FECHA_FORMATO /Retenciones/@FechaExp"}},
		{"1.0 con Z", strings.Replace(base10, `-06:00"`, `Z"`, 1), []string{"FECHA_FORMATO /Retenciones/@FechaExp"}},
		{"1.0 ilegible", strings.Replace(base10, `2020-06-30T18:00:00-06:00`, `ayer`, 1), []string{"FECHA_INVALIDA /Retenciones/@FechaExp"}},
		{"1.0 sin RFCEmisor", strings.Replace(base10, ` RFCEmisor="EKU9003173C9"`, ``, 1), []string{"ATRIBUTO_FALTANTE /Retenciones/Emisor/@RFCEmisor"}},
		{"1.0 sin montoRet", strings.Replace(base10, ` montoRet="6250.000000"`, ``, 1), []string{"ATRIBUTO_FALTANTE /Retenciones/Totales/ImpRetenidos[1]/@montoRet"}},
		{"2.0 sin LugarExpRetenc", strings.Replace(base, ` LugarExpRetenc="42501"`, ``, 1), []string{"ATRIBUTO_FALTANTE /Retenciones/@LugarExpRetenc"}},
		{"contenido tras la raíz", base + "basura", []string{"CONTENIDO_TRAS_RAIZ "}},
		{"XML no estricto", strings.Replace(base, `NomDenRazSocE="ESCUELA KEMPER URGATE"`, `NomDenRazSocE="ESCUELA & KEMPER"`, 1), []string{"XML_NO_ESTRICTO "}},
		{"Totales duplicado", strings.Replace(base, `<retenciones:Complemento>`, `<retenciones:Totales MontoTotOperacion="1" MontoTotGrav="1" MontoTotExent="0" MontoTotRet="0"/><retenciones:Complemento>`, 1),
			[]string{"NODO_DUPLICADO /Retenciones/Totales"}},
	}
	for _, c := range casos {
		rt, err := ParseRetencionesBytes([]byte(c.xml))
		if err != nil {
			t.Errorf("%s: %v", c.nombre, err)
			continue
		}
		if got := rutasRet(rt); !slices.Equal(got, c.quiero) {
			t.Errorf("%s: %q, quiero %q", c.nombre, got, c.quiero)
		}
	}
	// Totales duplicado: se queda el primero.
	rt, _ := ParseRetencionesBytes([]byte(casos[len(casos)-1].xml))
	monto(t, "Totales del primero", rt.Totales.MontoTotOperacion, "1000.00")
}

func TestParseRetencionesErrores(t *testing.T) {
	cfdi := leerSintetico(t, "ingreso40-pue.xml")
	ret := leerSintetico(t, "retenciones/ret20-plataformas.xml")
	casos := []struct {
		nombre string
		xml    []byte
		clase  string
	}{
		{"un CFDI no es Retenciones", cfdi, ClaseNoEsRetenciones},
		{"otra raíz", []byte(`<a/>`), ClaseNoEsRetenciones},
		{"no es XML", []byte("hola"), ClaseNoEsXML},
		{"truncado", ret[:len(ret)-40], ClaseXMLMalFormado},
	}
	for _, c := range casos {
		if _, err := ParseRetencionesBytes(c.xml); ClaseDe(err) != c.clase {
			t.Errorf("%s: %v, quiero %s", c.nombre, err, c.clase)
		}
	}
	// Parse no cambia: una Retenciones sigue siendo no_es_comprobante.
	if _, err := ParseBytes(ret); ClaseDe(err) != ClaseNoEsComprobante {
		t.Errorf("Parse de una Retenciones: %v", err)
	}
	if _, err := ParseRetenciones(lectorRoto{}); ClaseDe(err) != ClaseLectura {
		t.Errorf("lector roto: %v", err)
	}
	rt, err := ParseRetenciones(bytes.NewReader(ret))
	if err != nil || rt.Version != "2.0" {
		t.Errorf("ParseRetenciones: %v", err)
	}
}

// Sin IncluirSello, el JSON no lleva sellos ni certificado, y el original no
// se toca.
func TestRetencionesJSONSinSellos(t *testing.T) {
	rt := leerRetenciones(t, "ret20-plataformas.xml")
	rt.Sello, rt.Certificado = "SELLO", "CERT"
	rt.TimbreFiscalDigital.SelloCFD, rt.TimbreFiscalDigital.SelloSAT = "SCFD", "SSAT"
	b, err := rt.JSON()
	if err != nil {
		t.Fatal(err)
	}
	for _, s := range []string{"SELLO", "CERT", "SCFD", "SSAT"} {
		if bytes.Contains(b, []byte(s)) {
			t.Errorf("el JSON trae %s", s)
		}
	}
	if !bytes.HasPrefix(b, []byte(`{"tipo_documento":"retenciones","schema_version":"1","retenciones":{`)) || bytes.HasSuffix(b, []byte("\n")) {
		t.Errorf("sobre: %.80s", b)
	}
	if rt.Sello != "SELLO" || rt.TimbreFiscalDigital.SelloSAT != "SSAT" {
		t.Error("JSON() tocó el original")
	}
	con, _ := rt.JSONOpciones(OpcionesJSON{IncluirSello: true, Indentar: true})
	if !bytes.Contains(con, []byte(`"sello": "SELLO"`)) || !bytes.Contains(con, []byte(`"sello_sat": "SSAT"`)) {
		t.Error("con IncluirSello faltan los sellos")
	}
	// Sin advertencias, [] y no null.
	if !bytes.Contains(b, []byte(`"advertencias":[]`)) {
		t.Error("advertencias no es []")
	}
}

func TestRetencionesJSONIdaYVuelta(t *testing.T) {
	n := 0
	for _, p := range fixtures(t) {
		if !esRetenciones(t, p) {
			continue
		}
		n++
		b, _ := os.ReadFile(p)
		rt, err := ParseRetencionesBytes(b)
		if err != nil {
			t.Fatal(err)
		}
		uno, _ := rt.JSONOpciones(OpcionesJSON{IncluirSello: true})
		var d DocumentoRetenciones
		if err := json.Unmarshal(uno, &d); err != nil {
			t.Fatalf("%s: %v", p, err)
		}
		d.Retenciones.Advertencias = d.Advertencias
		dos, _ := d.Retenciones.JSONOpciones(OpcionesJSON{IncluirSello: true})
		if !bytes.Equal(uno, dos) {
			t.Errorf("%s: el JSON no sobrevive ida y vuelta", p)
		}
	}
	if n < 50 {
		t.Errorf("solo %d fixtures de Retenciones", n)
	}
}

// El contrato de Ruta y Atributo de las advertencias, en Retenciones.
func TestRetencionesAdvertenciasRuta(t *testing.T) {
	for _, p := range fixtures(t) {
		if !esRetenciones(t, p) {
			continue
		}
		b, _ := os.ReadFile(p)
		rt, _ := ParseRetencionesBytes(b)
		for _, a := range rt.Advertencias {
			_, sufijo, conAttr := strings.Cut(a.Ruta, "/@")
			if (conAttr && a.Atributo != sufijo) || (!conAttr && a.Atributo != "") || !slices.Contains(Codigos, a.Codigo) {
				t.Errorf("%s: %+v", p, a)
			}
			if a.Ruta != "" && !strings.HasPrefix(a.Ruta, "/Retenciones") {
				t.Errorf("%s: ruta %q", p, a.Ruta)
			}
		}
	}
}

func cargarEsquemaRetenciones(t *testing.T) esquema {
	t.Helper()
	return cargarEsquemaDe(t, "schema/retenciones.v1.schema.json", SchemaRetenciones)
}

func TestSchemaRetencionesFixtures(t *testing.T) {
	s := cargarEsquemaRetenciones(t)
	for _, p := range fixtures(t) {
		if !esRetenciones(t, p) {
			continue
		}
		b, _ := os.ReadFile(p)
		rt, err := ParseRetencionesBytes(b)
		if err != nil {
			t.Fatal(err)
		}
		for _, sello := range []bool{false, true} {
			d := rt.Documento(OpcionesJSON{IncluirSello: sello})
			d.Origen = p
			if sello {
				v := VerificarRetenciones(b, CertificadosEn())
				d.VerificacionSello, d.VerificacionTFD = &v.Sello, v.TFD
				d.Validacion = ValidarRetenciones(b)
			}
			out, _ := json.Marshal(d)
			for _, e := range validarJSON(t, s, out) {
				t.Errorf("%s (sello=%v): %s", p, sello, e)
			}
		}
	}
}

func TestSchemaRetencionesRechaza(t *testing.T) {
	s := cargarEsquemaRetenciones(t)
	b, _ := leerRetenciones(t, "ret20-plataformas.xml").JSON()
	if errs := validarJSON(t, s, b); len(errs) > 0 {
		t.Fatalf("el documento base no valida: %v", errs)
	}
	casos := map[string][2]string{
		"monto como número":   {`"monto_tot_ret":"100.00"`, `"monto_tot_ret":100.00`},
		"clave no declarada":  {`"cve_retenc":"26"`, `"cve_retenc":"26","clave":"x"`},
		"clave anidada extra": {`"rfc_r":"XAXX010101000"`, `"rfc_r":"XAXX010101000","correo":"x"`},
		"tipo_documento otro": {`"tipo_documento":"retenciones"`, `"tipo_documento":"comprobante"`},
		"schema_version otro": {`"schema_version":"1"`, `"schema_version":"2"`},
		"código desconocido":  {`"advertencias":[]`, `"advertencias":[{"codigo":"INVENTADO"}]`},
		"tasa como número":    {`"tasa_cuota":"0.16"`, `"tasa_cuota":0.16`},
		"clave como número":   {`"tipo_pago_ret":"01"`, `"tipo_pago_ret":1`},
	}
	for nombre, rep := range casos {
		mal := bytes.Replace(b, []byte(rep[0]), []byte(rep[1]), 1)
		if bytes.Equal(mal, b) {
			t.Fatalf("%s: el reemplazo no aplicó", nombre)
		}
		if errs := validarJSON(t, s, mal); len(errs) == 0 {
			t.Errorf("%s: el validador lo aceptó", nombre)
		}
	}
}

// Como TestSchemaCubreModelo: el esquema declara exactamente las claves del
// DocumentoRetenciones y sus enums son las listas de códigos.
func TestSchemaRetencionesCubreModelo(t *testing.T) {
	s := cargarEsquemaRetenciones(t)
	var d DocumentoRetenciones
	compararTipo(t, s, map[string]any{"type": "object", "properties": s["properties"]}, reflect.TypeOf(d), "$", map[reflect.Type]bool{})
	defs := s["$defs"].(map[string]any)
	enum := func(def, prop string) []string {
		var out []string
		for _, e := range defs[def].(map[string]any)["properties"].(map[string]any)[prop].(map[string]any)["enum"].([]any) {
			out = append(out, e.(string))
		}
		return out
	}
	if got := enum("advertencia", "codigo"); !slices.Equal(got, Codigos) {
		t.Errorf("enum de advertencias ≠ Codigos:\n%v\n%v", got, Codigos)
	}
	if got, quiero := enum("hallazgo", "codigo"), append(slices.Clone(CodigosEsquema), CodigosReglas...); !slices.Equal(got, quiero) {
		t.Errorf("enum de hallazgos ≠ CodigosEsquema + CodigosReglas:\n%v\n%v", got, quiero)
	}
	if got := enum("resultado_sello", "motivo"); !slices.Equal(got, Motivos) {
		t.Errorf("enum de motivos ≠ Motivos:\n%v\n%v", got, Motivos)
	}
	// Las definiciones compartidas son las mismas que en el esquema del CFDI.
	cfdi := cargarEsquema(t)["$defs"].(map[string]any)
	for _, k := range []string{"validacion", "estado_validacion", "resultado_validacion", "hallazgo", "resultado_sello", "decimal", "fecha", "uuid", "advertencia", "timbre_fiscal_digital"} {
		if !reflect.DeepEqual(defs[k], cfdi[k]) {
			t.Errorf("$defs/%s difiere de la del esquema del CFDI", k)
		}
	}
}
