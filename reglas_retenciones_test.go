package cfdi

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
	"time"
)

// retBase20 es una Retenciones 2.0 coherente: ninguna regla da hallazgo.
// Cada prueba la muta en un solo punto.
func retBase20() *Retenciones {
	return &Retenciones{
		Version:        "2.0",
		FechaExp:       Fecha{Original: "2025-06-01T10:00:00", Time: time.Date(2025, 6, 1, 10, 0, 0, 0, time.UTC)},
		LugarExpRetenc: "42501",
		CveRetenc:      "01",
		Emisor:         &EmisorRetenciones{RfcE: "EKU9003173C9", NomDenRazSocE: "ESCUELA KEMPER URGATE", RegimenFiscalE: "601"},
		Receptor: &ReceptorRetenciones{NacionalidadR: "Nacional", Nacional: &ReceptorNacionalRetenciones{
			RfcR: "XIQB891116QE4", NomDenRazSocR: "XENON INDUSTRIAL ARTICLES", DomicilioFiscalR: "42501"}},
		Periodo: &PeriodoRetenciones{MesIni: "01", MesFin: "12", Ejercicio: "2025"},
		Totales: &TotalesRetenciones{
			MontoTotOperacion: ptrDec("1000.00"), MontoTotGrav: ptrDec("800.00"), MontoTotExent: ptrDec("200.00"), MontoTotRet: ptrDec("580.00"),
			ImpRetenidos: []ImpRetenidos{
				{BaseRet: ptrDec("1000.00"), ImpuestoRet: "001", MontoRet: ptrDec("400.00"), TipoPagoRet: "03"},
				{BaseRet: ptrDec("1000.00"), ImpuestoRet: "002", MontoRet: ptrDec("180.00"), TipoPagoRet: "01"},
			},
		},
	}
}

func TestReglasRetencionesBaseCoherente(t *testing.T) {
	if hs := ValidarReglasRetenciones(retBase20()); len(hs) != 0 {
		t.Fatalf("la base debía ser coherente: %+v", hs)
	}
	// El estado de una lista vacía es válido.
	if EstadoDe(ValidarReglasRetenciones(retBase20())) != ValidacionValida {
		t.Error("estado")
	}
}

// Una prueba por regla: el caso que pasa (sin hallazgo con ese código) y el
// que falla (con el nivel, el atributo y la ruta que se esperan).
func TestReglasRetenciones20(t *testing.T) {
	const er, ad = NivelError, NivelAdvertencia
	_ = ad
	casos := []struct {
		nombre string
		muta   func(*Retenciones)
		codigo string
		ruta   string // de los hallazgos de ese código, en orden; vacío: ninguno
		attr   string
	}{
		// REGLA_RET_MONTO_OPERACION (Reten20127)
		{"operación = gravado + exento", func(r *Retenciones) {}, ReglaRetMontoOperacion, "", ""},
		{"operación con otra escala", func(r *Retenciones) { r.Totales.MontoTotOperacion = ptrDec("1000.000000") }, ReglaRetMontoOperacion, "", ""},
		{"operación con escalas mezcladas", func(r *Retenciones) {
			r.Totales.MontoTotOperacion, r.Totales.MontoTotGrav, r.Totales.MontoTotExent = ptrDec("0.3"), ptrDec("0.1"), ptrDec("0.2")
		}, ReglaRetMontoOperacion, "", ""},
		{"operación un centavo menos", func(r *Retenciones) { r.Totales.MontoTotOperacion = ptrDec("999.99") }, ReglaRetMontoOperacion, "/Retenciones/Totales", "MontoTotOperacion"},
		{"operación un centavo más", func(r *Retenciones) { r.Totales.MontoTotOperacion = ptrDec("1000.01") }, ReglaRetMontoOperacion, "/Retenciones/Totales", "MontoTotOperacion"},
		{"operación una millonésima más", func(r *Retenciones) { r.Totales.MontoTotOperacion = ptrDec("1000.000001") }, ReglaRetMontoOperacion, "/Retenciones/Totales", "MontoTotOperacion"},
		{"operación sin gravado", func(r *Retenciones) { r.Totales.MontoTotGrav = nil; r.Totales.MontoTotOperacion = ptrDec("5") }, ReglaRetMontoOperacion, "", ""},

		// REGLA_RET_MONTO_MAYOR_OPERACION (Reten20128, Reten20129)
		{"gravado igual a la operación", func(r *Retenciones) {
			r.Totales.MontoTotGrav, r.Totales.MontoTotExent = ptrDec("1000.00"), ptrDec("0")
		}, ReglaRetMontoMayor, "", ""},
		{"gravado mayor que la operación", func(r *Retenciones) { r.Totales.MontoTotGrav = ptrDec("1000.01") }, ReglaRetMontoMayor, "/Retenciones/Totales", "MontoTotGrav"},
		{"exento mayor que la operación", func(r *Retenciones) { r.Totales.MontoTotExent = ptrDec("1000.01") }, ReglaRetMontoMayor, "/Retenciones/Totales", "MontoTotExent"},

		// REGLA_RET_MONTO_TOT_RET (Reten20130, Reten20131)
		{"total retenido = suma", func(r *Retenciones) {}, ReglaRetMontoTotRet, "", ""},
		{"total retenido con decimales sueltos", func(r *Retenciones) {
			r.Totales.MontoTotRet = ptrDec("0.3")
			r.Totales.ImpRetenidos = []ImpRetenidos{{MontoRet: ptrDec("0.1")}, {MontoRet: ptrDec("0.20")}}
		}, ReglaRetMontoTotRet, "", ""},
		{"total retenido un centavo más", func(r *Retenciones) { r.Totales.MontoTotRet = ptrDec("580.01") }, ReglaRetMontoTotRet, "/Retenciones/Totales", "MontoTotRet"},
		{"total retenido un centavo menos", func(r *Retenciones) { r.Totales.MontoTotRet = ptrDec("579.99") }, ReglaRetMontoTotRet, "/Retenciones/Totales", "MontoTotRet"},
		{"total retenido sin ImpRetenidos", func(r *Retenciones) { r.Totales.ImpRetenidos = nil }, ReglaRetMontoTotRet, "/Retenciones/Totales", "MontoTotRet"},
		{"total retenido cero sin ImpRetenidos", func(r *Retenciones) {
			r.Totales.MontoTotRet, r.Totales.ImpRetenidos = ptrDec("0"), nil
		}, ReglaRetMontoTotRet, "", ""},
		{"total retenido cero con ImpRetenidos (el SAT solo pide con mayor que cero)", func(r *Retenciones) { r.Totales.MontoTotRet = ptrDec("0.00") }, ReglaRetMontoTotRet, "", ""},
		{"total retenido con un MontoRet ilegible no se compara", func(r *Retenciones) {
			r.Totales.MontoTotRet, r.Totales.ImpRetenidos[0].MontoRet = ptrDec("1"), nil
		}, ReglaRetMontoTotRet, "", ""},

		// REGLA_RET_PERIODO (Reten20122, Reten20124)
		{"mes inicial = final", func(r *Retenciones) { r.Periodo.MesIni, r.Periodo.MesFin = "06", "06" }, ReglaRetPeriodo, "", ""},
		{"mes inicial después del final", func(r *Retenciones) { r.Periodo.MesIni, r.Periodo.MesFin = "07", "06" }, ReglaRetPeriodo, "/Retenciones/Periodo", "MesIni"},
		{"meses comparados como número, no como texto", func(r *Retenciones) { r.Periodo.MesIni, r.Periodo.MesFin = "9", "10" }, ReglaRetPeriodo, "", ""},
		{"mes ilegible no se compara", func(r *Retenciones) { r.Periodo.MesIni, r.Periodo.MesFin = "xx", "06" }, ReglaRetPeriodo, "", ""},

		// REGLA_RET_EJERCICIO (Reten20126)
		{"ejercicio del año de FechaExp", func(r *Retenciones) {}, ReglaRetEjercicio, "", ""},
		{"ejercicio del año anterior", func(r *Retenciones) { r.Periodo.Ejercicio = "2024" }, ReglaRetEjercicio, "", ""},
		{"ejercicio 5 años antes", func(r *Retenciones) { r.Periodo.Ejercicio = "2020" }, ReglaRetEjercicio, "", ""},
		{"ejercicio 6 años antes", func(r *Retenciones) { r.Periodo.Ejercicio = "2019" }, ReglaRetEjercicio, "/Retenciones/Periodo", "Ejercicio"},
		{"ejercicio del año siguiente", func(r *Retenciones) { r.Periodo.Ejercicio = "2026" }, ReglaRetEjercicio, "/Retenciones/Periodo", "Ejercicio"},
		{"ejercicio sin FechaExp legible", func(r *Retenciones) { r.FechaExp = Fecha{}; r.Periodo.Ejercicio = "2030" }, ReglaRetEjercicio, "", ""},
		{"ejercicio y fin de año: FechaExp el 31 de diciembre", func(r *Retenciones) {
			r.FechaExp = Fecha{Original: "2025-12-31T23:59:59", Time: time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC)}
		}, ReglaRetEjercicio, "", ""},

		// REGLA_RET_TIPO_PAGO (Reten20135)
		{"ISR definitivo", func(r *Retenciones) {}, ReglaRetTipoPago, "", ""},
		{"ISR provisional", func(r *Retenciones) { r.Totales.ImpRetenidos[0].TipoPagoRet = "04" }, ReglaRetTipoPago, "", ""},
		{"ISR con pago definitivo de IVA", func(r *Retenciones) { r.Totales.ImpRetenidos[0].TipoPagoRet = "01" }, ReglaRetTipoPago, "/Retenciones/Totales/ImpRetenidos[1]", "TipoPagoRet"},
		{"IVA con pago de ISR", func(r *Retenciones) { r.Totales.ImpRetenidos[1].TipoPagoRet = "03" }, ReglaRetTipoPago, "/Retenciones/Totales/ImpRetenidos[2]", "TipoPagoRet"},
		{"IEPS definitivo", func(r *Retenciones) {
			r.Totales.ImpRetenidos[1].ImpuestoRet, r.Totales.ImpRetenidos[1].TipoPagoRet = "003", "02"
		}, ReglaRetTipoPago, "", ""},
		{"IEPS con pago de IVA", func(r *Retenciones) {
			r.Totales.ImpRetenidos[1].ImpuestoRet, r.Totales.ImpRetenidos[1].TipoPagoRet = "003", "01"
		}, ReglaRetTipoPago, "/Retenciones/Totales/ImpRetenidos[2]", "TipoPagoRet"},
		{"ImpuestoRet ausente no se compara", func(r *Retenciones) { r.Totales.ImpRetenidos[0].ImpuestoRet = "" }, ReglaRetTipoPago, "", ""},
		{"TipoPagoRet fuera del catálogo lo reporta el esquema", func(r *Retenciones) { r.Totales.ImpRetenidos[0].TipoPagoRet = "09" }, ReglaRetTipoPago, "", ""},

		// REGLA_RET_DESC_RETENC (Reten20106)
		{"clave 25 con descripción", func(r *Retenciones) { r.CveRetenc, r.DescRetenc = "25", "Otro" }, ReglaRetDescRetenc, "", ""},
		{"clave 25 sin descripción", func(r *Retenciones) { r.CveRetenc = "25" }, ReglaRetDescRetenc, "/Retenciones", "DescRetenc"},
		{"clave 25 con descripción en blanco", func(r *Retenciones) { r.CveRetenc, r.DescRetenc = "25", "  " }, ReglaRetDescRetenc, "/Retenciones", "DescRetenc"},
		{"otra clave sin descripción", func(r *Retenciones) { r.CveRetenc = "03" }, ReglaRetDescRetenc, "", ""},

		// REGLA_RET_CVE_COMPLEMENTO (Reten20107)
		{"clave 26 con Plataformas", func(r *Retenciones) {
			r.CveRetenc, r.ServiciosPlataformasTecnologicas = "26", &ServiciosPlataformasTecnologicas{}
		}, ReglaRetCveComplemento, "", ""},
		{"clave 26 sin Plataformas", func(r *Retenciones) { r.CveRetenc = "26" }, ReglaRetCveComplemento, "/Retenciones", "CveRetenc"},
		{"clave 26 con Plataformas de una versión que la lectura no modela", func(r *Retenciones) {
			r.CveRetenc, r.ComplementosNoSoportados = "26", []string{"{http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas11}ServiciosPlataformasTecnologicas"}
		}, ReglaRetCveComplemento, "", ""},
		{"clave 26 con otro complemento no soportado", func(r *Retenciones) {
			r.CveRetenc, r.ComplementosNoSoportados = "26", []string{"{urn:raro}Raro"}
		}, ReglaRetCveComplemento, "/Retenciones", "CveRetenc"},
		{"clave 14 con Dividendos", func(r *Retenciones) { r.CveRetenc, r.Dividendos = "14", &Dividendos{} }, ReglaRetCveComplemento, "", ""},
		{"clave 15 sin Dividendos", func(r *Retenciones) { r.CveRetenc = "15" }, ReglaRetCveComplemento, "/Retenciones", "CveRetenc"},
		{"clave 06 con Enajenación", func(r *Retenciones) { r.CveRetenc, r.EnajenaciondeAcciones = "06", &EnajenaciondeAcciones{} }, ReglaRetCveComplemento, "", ""},
		{"clave 19 sin Enajenación", func(r *Retenciones) { r.CveRetenc = "19" }, ReglaRetCveComplemento, "/Retenciones", "CveRetenc"},
		{"clave 22 con Planes de retiro", func(r *Retenciones) { r.CveRetenc, r.Planesderetiro = "22", &Planesderetiro{} }, ReglaRetCveComplemento, "", ""},
		{"clave 21 con Planes de retiro en vez de Fideicomiso", func(r *Retenciones) { r.CveRetenc, r.Planesderetiro = "21", &Planesderetiro{} }, ReglaRetCveComplemento, "/Retenciones", "CveRetenc"},
		{"clave 01 sin complemento (condicional: no se revisa)", func(r *Retenciones) {
			r.CveRetenc, r.Receptor = "01", &ReceptorRetenciones{NacionalidadR: "Extranjero", Extranjero: &ReceptorExtranjeroRetenciones{}}
		}, ReglaRetCveComplemento, "", ""},
		{"clave 27 sin complemento (condicional: no se revisa)", func(r *Retenciones) { r.CveRetenc = "27" }, ReglaRetCveComplemento, "", ""},
		{"clave 03 sin complemento", func(r *Retenciones) { r.CveRetenc = "03" }, ReglaRetCveComplemento, "", ""},

		// REGLA_RET_CVE_RIF (Reten20108, Reten20132, Reten20133)
		{"clave 28 completa", func(r *Retenciones) { retRIF(r) }, ReglaRetCveRIF, "", ""},
		{"clave 28 con gravado", func(r *Retenciones) { retRIF(r); r.Totales.MontoTotGrav = ptrDec("0.01") }, ReglaRetCveRIF, "/Retenciones/Totales", "MontoTotGrav"},
		{"clave 28 con exento", func(r *Retenciones) { retRIF(r); r.Totales.MontoTotExent = ptrDec("1") }, ReglaRetCveRIF, "/Retenciones/Totales", "MontoTotExent"},
		{"clave 28 sin utilidad", func(r *Retenciones) { retRIF(r); r.Totales.UtilidadBimestral = nil }, ReglaRetCveRIF, "/Retenciones/Totales", "UtilidadBimestral"},
		{"clave 28 con utilidad en cero", func(r *Retenciones) { retRIF(r); r.Totales.UtilidadBimestral = ptrDec("0.00") }, ReglaRetCveRIF, "/Retenciones/Totales", "UtilidadBimestral"},
		{"clave 28 sin ISR correspondiente", func(r *Retenciones) { retRIF(r); r.Totales.ISRCorrespondiente = nil }, ReglaRetCveRIF, "/Retenciones/Totales", "ISRCorrespondiente"},
		{"clave 28 con ISR en cero", func(r *Retenciones) { retRIF(r); r.Totales.ISRCorrespondiente = ptrDec("0") }, ReglaRetCveRIF, "/Retenciones/Totales", "ISRCorrespondiente"},
		{"otra clave sin utilidad ni ISR", func(r *Retenciones) {}, ReglaRetCveRIF, "", ""},

		// REGLA_RET_RECEPTOR_NACIONAL (Reten20116, Reten20117, Reten20120)
		{"persona física con CURP", func(r *Retenciones) { r.Receptor.Nacional.CurpR = "XEXX010101HNEXXXA4" }, ReglaRetReceptor, "", ""},
		{"persona moral con CURP", func(r *Retenciones) {
			r.Receptor.Nacional.RfcR, r.Receptor.Nacional.CurpR = "EKU9003173C9", "XEXX010101HNEXXXA4"
		}, ReglaRetReceptor, "/Retenciones/Receptor/Nacional", "CurpR"},
		{"persona moral sin CURP", func(r *Retenciones) { r.Receptor.Nacional.RfcR = "EKU9003173C9" }, ReglaRetReceptor, "", ""},
		{"público en general con RFC genérico y su domicilio", func(r *Retenciones) {
			n := r.Receptor.Nacional
			n.RfcR, n.NomDenRazSocR = "XAXX010101000", "PÚBLICO EN GENERAL"
		}, ReglaRetReceptor, "", ""},
		{"público en general sin acento con RFC propio", func(r *Retenciones) { r.Receptor.Nacional.NomDenRazSocR = "Publico en General" }, ReglaRetReceptor, "/Retenciones/Receptor/Nacional", "RfcR"},
		{"público en general con RFC propio", func(r *Retenciones) { r.Receptor.Nacional.NomDenRazSocR = "PÚBLICO EN GENERAL" }, ReglaRetReceptor, "/Retenciones/Receptor/Nacional", "RfcR"},
		{"RFC genérico con otro domicilio", func(r *Retenciones) {
			n := r.Receptor.Nacional
			n.RfcR, n.NomDenRazSocR, n.DomicilioFiscalR = "XAXX010101000", "PUBLICO EN GENERAL", "06600"
		}, ReglaRetReceptor, "/Retenciones/Receptor/Nacional", "DomicilioFiscalR"},
		{"RFC propio con otro domicilio", func(r *Retenciones) { r.Receptor.Nacional.DomicilioFiscalR = "06600" }, ReglaRetReceptor, "", ""},
		{"receptor extranjero no revisa Nacional", func(r *Retenciones) {
			r.Receptor = &ReceptorRetenciones{NacionalidadR: "Extranjero", Extranjero: &ReceptorExtranjeroRetenciones{NumRegIdTribR: "1"}}
		}, ReglaRetReceptor, "", ""},
	}
	for _, c := range casos {
		rt := retBase20()
		c.muta(rt)
		var hs []Hallazgo
		for _, h := range ValidarReglasRetenciones(rt) {
			if h.Codigo == c.codigo {
				hs = append(hs, h)
			}
		}
		if c.ruta == "" {
			if len(hs) != 0 {
				t.Errorf("%s: %+v, quiero ninguno", c.nombre, hs)
			}
			continue
		}
		if len(hs) != 1 {
			t.Errorf("%s: %d hallazgos %+v, quiero 1", c.nombre, len(hs), hs)
			continue
		}
		h := hs[0]
		if h.Nivel != er || h.Nodo() != c.ruta || h.Atributo != c.attr || h.Regla == "" || h.Mensaje == "" || h.Nivel != reglaDe[h.Codigo].Nivel {
			t.Errorf("%s: %+v, quiero nivel %s, nodo %s, atributo %s", c.nombre, h, er, c.ruta, c.attr)
		}
	}
}

func retRIF(r *Retenciones) {
	r.CveRetenc = "28"
	t := r.Totales
	t.MontoTotOperacion, t.MontoTotGrav, t.MontoTotExent = ptrDec("0"), ptrDec("0"), ptrDec("0")
	t.UtilidadBimestral, t.ISRCorrespondiente = ptrDec("12"), ptrDec("1.5")
}

// Cada regla de Retenciones 2.0 sale solo con su versión: una 2.0 con todo
// roto no le da nada a la 1.0 y al revés.
func TestReglasRetenciones10(t *testing.T) {
	rt := &Retenciones{Version: "1.0", CveRetenc: "25",
		Periodo: &PeriodoRetenciones{MesIni: "12", MesFin: "1", Ejercicio: "2001"},
		Totales: &TotalesRetenciones{MontoTotOperacion: ptrDec("9"), MontoTotGrav: ptrDec("1"), MontoTotExent: ptrDec("1"), MontoTotRet: ptrDec("1.00"),
			ImpRetenidos: []ImpRetenidos{{ImpuestoRet: "01", MontoRet: ptrDec("0.60"), TipoPagoRet: "Pago definitivo"}, {MontoRet: ptrDec("0.40")}}}}
	if hs := ValidarReglasRetenciones(rt); len(hs) != 0 {
		t.Errorf("1.0 coherente en su única regla: %+v", hs)
	}
	for _, c := range []struct{ total, suma string }{{"1.01", "1.00"}, {"0.99", "1.00"}, {"10", "1"}, {"0", "0.01"}} {
		rt.Totales.MontoTotRet = ptrDec(c.total)
		rt.Totales.ImpRetenidos = []ImpRetenidos{{MontoRet: ptrDec(c.suma)}}
		hs := ValidarReglasRetenciones(rt)
		if len(hs) != 1 || hs[0].Codigo != ReglaRet10MontoTotRet || hs[0].Nivel != NivelAdvertencia || hs[0].Atributo != "montoTotRet" ||
			hs[0].Ruta != "/Retenciones/Totales/@montoTotRet" || hs[0].Valor != c.total {
			t.Errorf("1.0 total %s, suma %s: %+v", c.total, c.suma, hs)
		}
		if EstadoDe(hs) != ValidacionValida {
			t.Errorf("una advertencia no invalida: %s", EstadoDe(hs))
		}
	}
	// Sin ImpRetenidos la suma es 0.
	rt.Totales.MontoTotRet, rt.Totales.ImpRetenidos = ptrDec("0.00"), nil
	if hs := ValidarReglasRetenciones(rt); len(hs) != 0 {
		t.Errorf("1.0 sin retenciones y total 0: %+v", hs)
	}
	// Con una 2.0 incoherente en todo, la regla de 1.0 no corre.
	r2 := retBase20()
	r2.Totales.MontoTotRet = ptrDec("1")
	if TieneHallazgo(ValidarReglasRetenciones(r2), ReglaRet10MontoTotRet) {
		t.Error("la regla de 1.0 corrió en una 2.0")
	}
}

func TestReglasRetencionesVersionNoSoportada(t *testing.T) {
	hs := ValidarReglasRetenciones(&Retenciones{Version: "3.0"})
	if len(hs) != 1 || hs[0].Codigo != ReglaVersionNoSoportada || hs[0].Nivel != NivelAviso || EstadoDe(hs) != ValidacionIncompleta {
		t.Errorf("%+v", hs)
	}
}

// Un Retenciones sin nodos (la lectura tolerante los deja nil) no revienta.
func TestReglasRetencionesNodosAusentes(t *testing.T) {
	for _, rt := range []*Retenciones{{Version: "2.0"}, {Version: "1.0"}, {Version: "2.0", CveRetenc: "28", Totales: &TotalesRetenciones{}},
		{Version: "2.0", Receptor: &ReceptorRetenciones{}, Periodo: &PeriodoRetenciones{}, Totales: &TotalesRetenciones{ImpRetenidos: []ImpRetenidos{{}}}}} {
		_ = ValidarReglasRetenciones(rt)
	}
}

// Los hallazgos de los documentos leídos de XML llevan las rutas del XML
// (nombres de 2.0) y ValidarRetenciones los suma al esquema.
func TestValidarRetencionesIncluyeReglas(t *testing.T) {
	b, err := os.ReadFile("testdata/sinteticos/retenciones/ret20-planes10.xml")
	if err != nil {
		t.Fatal(err)
	}
	// CveRetenc 21 pide Fideicomisonoempresarial y el sintético trae Planes de retiro.
	b = []byte(strings.Replace(string(b), `CveRetenc="22"`, `CveRetenc="21"`, 1))
	v := ValidarRetenciones(b)
	if v.Reglas == nil || !TieneHallazgo(v.Reglas.Hallazgos, ReglaRetCveComplemento) || v.Estado != ValidacionInvalida || v.Reglas.Estado != ValidacionInvalida {
		t.Fatalf("%+v", v)
	}
	if hs := v.Hallazgos(); len(hs) < len(v.Esquema.Hallazgos)+1 || hs[len(hs)-1].Codigo != ReglaRetCveComplemento {
		t.Errorf("Hallazgos no junta esquema y reglas: %+v", hs)
	}
	// Lo que no es Retenciones queda sin reglas y con el estado del esquema.
	v = ValidarRetenciones([]byte("hola"))
	if v.Reglas != nil || v.Estado != ValidacionInvalida {
		t.Errorf("%+v", v)
	}
}

// Las reglas nuevas se aplican a todos los ejemplos de Retenciones de
// testdata. Lo único que sale son hallazgos explicados:
//
//   - los 1.0 de satcfdi traen montoTotRet que no es la suma de los montoRet
//     (3500 contra 500 o 1000, o 10 sin ImpRetenidos): advertencia;
//   - retenciones20-{ejemplo,min,min-comps} de satcfdi son pruebas de
//     generación con Ejercicio 2027 en un comprobante de 2021 y
//     MontoTotOperacion 1 = Grav 1 = Exent 1; el «ejemplo» además combina
//     cada impuesto con un TipoPagoRet de otro y MontoTotRet 1 contra
//     1 + 1 + 1;
//   - retenciones20 de CfdiUtils trae ISR (001) con TipoPagoRet 01 (IVA);
//
// Los sintéticos de este repo salían con hallazgos al estrenar las reglas, y
// eran incoherentes de verdad: ret20-planes10 con CveRetenc 21 (Fideicomisos
// no empresariales) y el complemento de Planes de retiro, que es de la 22, y
// ret20-plataformas (y su copia sellada) con ISR y TipoPagoRet 01, que es el
// pago definitivo de IVA. Se corrigieron (22 y 03) y se regeneraron sus
// goldens, cadenas y sello.
func TestReglasRetencionesFixtures(t *testing.T) {
	const ejercicio27 = "REGLA_RET_EJERCICIO /Retenciones/Periodo/@Ejercicio"
	const operacion = "REGLA_RET_MONTO_OPERACION /Retenciones/Totales/@MontoTotOperacion"
	const tipoPago = "REGLA_RET_TIPO_PAGO /Retenciones/Totales/ImpRetenidos[%d]/@TipoPagoRet"
	sat10 := []string{"REGLA_RET10_MONTO_TOT_RET /Retenciones/Totales/@montoTotRet"}
	satEjemplo := []string{operacion, "REGLA_RET_MONTO_TOT_RET /Retenciones/Totales/@MontoTotRet", ejercicio27,
		strings.Replace(tipoPago, "%d", "1", 1), strings.Replace(tipoPago, "%d", "2", 1), strings.Replace(tipoPago, "%d", "3", 1)}
	satMin := []string{operacion, ejercicio27}
	esperado := func(p string) []string {
		base := filepath.Base(p)
		switch {
		case strings.Contains(p, "satcfdi") && strings.Contains(base, "retenciones10__"):
			return sat10
		case strings.Contains(p, "satcfdi") && strings.Contains(base, "retenciones20-ejemplo"):
			return satEjemplo
		case strings.Contains(p, "satcfdi") && strings.Contains(base, "retenciones20-min"): // min y min-comps
			return satMin
		case strings.Contains(p, "eclipxe13") && strings.HasSuffix(base, "retenciones20.xml"):
			return []string{strings.Replace(tipoPago, "%d", "1", 1)}
		}
		return nil
	}
	vistos := 0
	for _, p := range fixtures(t) {
		if esFatal(p) {
			continue
		}
		b, err := os.ReadFile(p)
		if err != nil {
			t.Fatal(err)
		}
		if Detectar(b) != TipoRetenciones {
			continue
		}
		rt, err := ParseRetencionesBytes(b)
		if err != nil {
			t.Fatalf("%s: %v", p, err)
		}
		vistos++
		quiero := slices.Clone(esperado(p))
		var got []string
		for _, h := range ValidarReglasRetenciones(rt) {
			got = append(got, h.Codigo+" "+h.Ruta)
		}
		slices.Sort(got)
		slices.Sort(quiero)
		if !slices.Equal(got, quiero) {
			t.Errorf("%s:\n  obtuve %v\n  quiero %v", p, got, quiero)
		}
	}
	if vistos < 54 {
		t.Errorf("solo %d ejemplos de Retenciones, quiero al menos 54", vistos)
	}
}

// Las reglas de Retenciones citan documentos de Retenciones, están en Reglas
// al final y en ReglasRetenciones, y cada código del que habla un hallazgo
// existe.
func TestReglasRetencionesDocumentadas(t *testing.T) {
	if len(ReglasRetenciones) != 11 {
		t.Errorf("%d reglas de Retenciones, quiero 11", len(ReglasRetenciones))
	}
	n := len(Reglas)
	if !slices.Equal(CodigosReglas[n-len(ReglasRetenciones):], func() []string {
		var out []string
		for _, r := range ReglasRetenciones {
			out = append(out, r.Codigo)
		}
		return out
	}()) {
		t.Error("las reglas de Retenciones no van al final de Reglas")
	}
	for _, r := range ReglasRetenciones {
		if !strings.HasPrefix(r.Codigo, "REGLA_RET") || len(r.Citas) == 0 {
			t.Errorf("%+v", r)
		}
		for _, c := range r.Citas {
			if !strings.Contains(c.Documento, "Retenciones") {
				t.Errorf("%s cita %q", r.Codigo, c.Documento)
			}
		}
	}
}
