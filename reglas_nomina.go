package cfdi

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/OrtaMarco/cfdi-go/decimal"
)

// Reglas de Nómina 1.2 que llegaron en T8 (matriz Rev D, NOM3-29 del
// comprobante y NOM31-106 del complemento): presencia de nodos por tipo de
// percepción, deducción y otro pago, relación laboral, cuenta bancaria y el
// CFDI de nómina 4.0. Lo que necesita la lista de RFC, la marca del SNCF o
// los topes del subsidio vigentes queda fuera (ver el README).
const (
	ReglaNominaComprobante       = "REGLA_NOMINA_COMPROBANTE"
	ReglaNominaPeriodicidad      = "REGLA_NOMINA_PERIODICIDAD"
	ReglaNominaRegistroPatronal  = "REGLA_NOMINA_REGISTRO_PATRONAL"
	ReglaNominaEntidadSNCF       = "REGLA_NOMINA_ENTIDAD_SNCF"
	ReglaNominaAntiguedad        = "REGLA_NOMINA_ANTIGUEDAD"
	ReglaNominaTipoRegimen       = "REGLA_NOMINA_TIPO_REGIMEN"
	ReglaNominaCuentaBancaria    = "REGLA_NOMINA_CUENTA_BANCARIA"
	ReglaNominaNodosPercepcion   = "REGLA_NOMINA_NODOS_PERCEPCION"
	ReglaNominaIncapacidades     = "REGLA_NOMINA_INCAPACIDADES"
	ReglaNominaCompensacionSaldo = "REGLA_NOMINA_COMPENSACION_SALDO"
)

var reglasNomina = []Regla{
	conSoloGuia(regla(ReglaNominaComprobante, NivelError,
		"CFDI 4.0 de nómina: Exportacion 01, MetodoPago PUE, sin InformacionGlobal ni FacAtrAdquirente; Nomina/Emisor/Curp solo y siempre con Rfc del emisor de 13 caracteres; Receptor con Rfc de 13 caracteres, RegimenFiscalReceptor 605, UsoCFDI CN01 y sin ResidenciaFiscal ni NumRegIdTrib; un solo concepto 84111505, Cantidad 1 sin decimales, ACT, «Pago de nómina», ObjetoImp 01, sin NoIdentificacion, Unidad ni nodos hijo",
		citas(cod(DocMatrizNomina12, "NOM3", "NOM4", "NOM5", "NOM6", "NOM7", "NOM8", "NOM11", "NOM12", "NOM13", "NOM14", "NOM15", "NOM16", "NOM17", "NOM18", "NOM19", "NOM23", "NOM24", "NOM25", "NOM26", "NOM27", "NOM28", "NOM29"),
			pag(DocGuiaNomina12, 10, 12, 13, 15, 16, 17, 18, 19)),
		"MetodoPago PUE y sin ResidenciaFiscal ni NumRegIdTrib salen de la guía (p. 10, 16), no de la matriz ni del estándar (p. 48-50); CFDI40151 y CFDI40154 los prohíben con un Rfc inscrito en el SAT, pero eso necesita la lista de RFC"),
		"MetodoPago PUE (guía p. 10) y sin ResidenciaFiscal ni NumRegIdTrib en el Receptor (guía p. 16)"),
	regla(ReglaNominaPeriodicidad, NivelError, "Nómina 1.2: TipoNomina O con PeriodicidadPago distinta de 99; E con 99",
		citas(cod(DocMatrizNomina12, "NOM33", "NOM34"), pag(DocGuiaNomina12, 22, 42)), ""),
	conSoloGuia(regla(ReglaNominaRegistroPatronal, NivelError,
		"Nómina 1.2: con TipoContrato 01 a 08, Emisor/RegistroPatronal; con 09, 10 o 99, sin él; con RegistroPatronal, NumSeguridadSocial, FechaInicioRelLaboral, Antigüedad, RiesgoPuesto y SalarioDiarioIntegrado, y sin él ninguno de ellos",
		citas(cod(DocMatrizNomina12, "NOM42", "NOM43", "NOM44"), pag(DocEstNomina12, 51), pag(DocGuiaNomina12, 28, 29)),
		"«en caso contrario, no debe existir ninguno de los campos anteriormente señalados» es de la guía (p. 29); la matriz y el estándar (p. 51) solo piden que existan"),
		"sin RegistroPatronal, ninguno de NumSeguridadSocial, FechaInicioRelLaboral, Antigüedad, RiesgoPuesto y SalarioDiarioIntegrado (guía p. 29)"),
	regla(ReglaNominaEntidadSNCF, NivelError,
		"Nómina 1.2: con OrigenRecurso IM, MontoRecursoPropio; sin IM, sin él; MontoRecursoPropio menor que TotalPercepciones + TotalOtrosPagos",
		citas(cod(DocMatrizNomina12, "NOM48", "NOM49", "NOM50"), pag(DocGuiaNomina12, 31, 32)),
		"si el emisor está adherido al SNCF (NOM45, NOM46) necesita la lista de RFC y no se revisa"),
	regla(ReglaNominaAntiguedad, NivelError,
		"Nómina 1.2: FechaInicioRelLaboral menor o igual que FechaFinalPago; Antigüedad PnW con n × 7 ≤ días de FechaInicioRelLaboral a FechaFinalPago + 1; PaYmMdD con FechaInicioRelLaboral + a años + m meses + d días ≤ FechaFinalPago + 1 día",
		citas(cod(DocMatrizNomina12, "NOM53", "NOM54", "NOM55"), pag(DocGuiaNomina12, 34, 35, 36)),
		"NOM55 pide que «corresponda»; la guía (p. 35) que sea «menor o igual que el tiempo transcurrido» y cuenta los días como calendario (un mes de marzo es P31D): se aplica menor o igual, contando el día final"),
	conVigencia(regla(ReglaNominaTipoRegimen, NivelError,
		"Nómina 1.2: con TipoContrato 01 a 08, TipoRegimen 02, 03 o 04; con 09 o más, 05 a 99; con TipoRegimen 02, un OtroPago 002 salvo que haya 007 o 008; con otro TipoRegimen, sin OtroPago 002, 007 ni 008",
		citas(cod(DocMatrizNomina12, "NOM57", "NOM58", "NOM105", "NOM106"), cod(DocMatrizNomina12C, "NOM105", "NOM106"), cod(DocMatrizNomina12B, "NOM226", "NOM227"), pag(DocEstNomina12, 52, 57), pag(DocGuiaNomina12, 39)), ""),
		"el OtroPago 002, 007 y 008 según TipoRegimen (NOM105, NOM106) es error desde el 2020-02-01: "+porqueNOM105),
	regla(ReglaNominaCuentaBancaria, NivelError,
		"Nómina 1.2: CuentaBancaria de 10, 11, 15, 16 o 18 posiciones; con 18 (CLABE), sin Banco y con dígito de control correcto; con 10, 11, 15 o 16, con Banco",
		citas(cod(DocMatrizNomina12, "NOM62", "NOM63", "NOM64", "NOM65"), pag(DocGuiaNomina12, 43)),
		"la matriz admite 10, 11, 16 y 18; la guía (p. 43) y el XSD también 15 (tarjetas): se aceptan; el dígito de control es el de Banxico (pesos 3, 7, 1); el ejemplo de la guía, 002215911558451272, no lo cumple"),
	regla(ReglaNominaNodosPercepcion, NivelError,
		"Nómina 1.2: TotalSueldos con percepciones distintas de 022, 023, 025, 039 y 044; TotalSeparacionIndemnizacion y SeparacionIndemnizacion con 022, 023 o 025; TotalJubilacionPensionRetiro y JubilacionPensionRetiro con 039 o 044, y sin ellas ninguno; con 039, TotalUnaExhibicion sin TotalParcialidad ni MontoDiario; con 044 al revés; AccionesOTitulos solo y siempre en 045; HorasExtra solo y siempre en 019; TotalUnaExhibicion sin MontoDiario ni TotalParcialidad; TotalParcialidad con MontoDiario y sin TotalUnaExhibicion",
		citas(cod(DocMatrizNomina12, "NOM77", "NOM78", "NOM79", "NOM80", "NOM81", "NOM82", "NOM83", "NOM84", "NOM85", "NOM89", "NOM90"), pag(DocGuiaNomina12, 53, 54, 56, 57, 59, 60)), ""),
	regla(ReglaNominaIncapacidades, NivelError,
		"Nómina 1.2: con percepción 014 o deducción 006, nodo Incapacidades; la suma de ImporteMonetario igual a ImporteGravado + ImporteExento de las percepciones 014 y al Importe de las deducciones 006",
		citas(cod(DocMatrizNomina12, "NOM86", "NOM87", "NOM94", "NOM95"), pag(DocGuiaNomina12, 54, 69)), ""),
	conVigencia(regla(ReglaNominaCompensacionSaldo, NivelError,
		"Nómina 1.2: OtroPago 004 con CompensacionSaldosAFavor y 002 con SubsidioAlEmpleo; el Año de CompensacionSaldosAFavor es el anterior al de FechaPago, o el mismo si el periodo de pago es diciembre",
		citas(cod(DocMatrizNomina12, "NOM98", "NOM99", "NOM103"), cod(DocMatrizNomina12C, "NOM103"), cod(DocMatrizNomina12B, "NOM223"), pag(DocEstNomina12, 56, 57), pag(DocGuiaNomina12, 73, 78, 123)),
		"«el periodo de pago sea diciembre»: FechaPago o FechaFinalPago en diciembre"),
		"el Año de CompensacionSaldosAFavor (NOM103) es error desde el 2020-02-01: "+porqueNOM103),
}

// Vigencia de lo que el SAT agregó a Nómina 1.2 después de la primera
// versión del estándar (Regla.Vigencia). El control de cambios del estándar
// (p. 56-57) dice en qué revisión entró cada validación, pero no desde
// cuándo es obligatoria. La matriz de la Rev B (publicada el 17-12-2019) es
// la primera que las trae (NOM226, NOM227 y NOM223; la Rev C las renumeró
// NOM105, NOM106 y NOM103), y el portal del SAT da los documentos
// anteriores como «vigentes hasta el 31 de enero del 2020»: son error desde
// el 2020-02-01 y advertencia antes. El resto de las
// validaciones del complemento está en el estándar desde su primera versión
// (Rev A, 2016) y no tiene fecha: Nómina 1.2 es la única versión válida
// desde el 1 de abril de 2017 (guía p. 86).
const (
	porqueNOM105 = "la Rev B del estándar (2019-12-11, control de cambios p. 57) agregó la validación de la clave 002 en TipoRegimen y la matriz de la Rev B (2019-12-17, NOM226 y NOM227) la trae; los documentos anteriores fueron vigentes hasta el 2020-01-31"
	porqueNOM103 = "la guía la agregó en su revisión 2 (2017-11-10, control de cambios p. 123) con una nota de entrada en vigor que ya no publica, la Rev B del estándar (2019-12-11, p. 57) la actualizó y la matriz de la Rev B (2019-12-17, NOM223) la trae; la matriz anterior (2017) solo pedía un Año menor que el año en curso, y fue vigente hasta el 2020-01-31"
)

var (
	matrizNomina12B = time.Date(2020, 2, 1, 0, 0, 0, 0, time.UTC)
	vigenciaNOM105  = vigencia{matrizNomina12B, porqueNOM105}
	vigenciaNOM103  = vigencia{matrizNomina12B, porqueNOM103}
)

// nominaT8 revisa las reglas de T8 de un complemento de Nómina 1.2.
func (v *validadorReglas) nominaT8(n Nomina, r string) {
	// Sin Receptor (lo reporta el esquema) no se evalúa lo que depende de él.
	if rc := n.Receptor; rc != nil {
		v.periodicidad(n, r)
		v.registroPatronal(n, rc, r)
		v.antiguedad(n, rc, r+"/Receptor")
		v.tipoRegimen(n, rc, r)
		v.cuentaBancaria(rc, r+"/Receptor")
	}
	v.entidadSNCF(n, r)
	if p := n.Percepciones; p != nil {
		v.nodosPercepcion(p, r+"/Percepciones")
	}
	v.incapacidades(n, r)
	v.compensacionSaldo(n, r)
}

func (v *validadorReglas) periodicidad(n Nomina, r string) {
	tipo := clave(n.TipoNomina)
	if tipo == "" {
		return
	}
	p := clave(n.Receptor.PeriodicidadPago)
	switch {
	case tipo == "O" && p == "99":
		v.hallazgoTexto(ReglaNominaPeriodicidad, r+"/Receptor", "PeriodicidadPago", n.Receptor.PeriodicidadPago, "una nómina ordinaria no lleva PeriodicidadPago 99")
	case tipo == "E" && p != "" && p != "99":
		v.hallazgoTexto(ReglaNominaPeriodicidad, r+"/Receptor", "PeriodicidadPago", n.Receptor.PeriodicidadPago, "una nómina extraordinaria lleva PeriodicidadPago 99")
	}
}

// claveNum lee una clave de catálogo numérica ("01", "99"); -1 si no lo es.
func claveNum(s string) int {
	n, err := strconv.Atoi(clave(s))
	if err != nil {
		return -1
	}
	return n
}

func (v *validadorReglas) registroPatronal(n Nomina, rc *NominaReceptor, r string) {
	registro := ""
	if n.Emisor != nil {
		registro = clave(n.Emisor.RegistroPatronal)
	}
	switch c := claveNum(rc.TipoContrato); {
	case c >= 1 && c <= 8 && registro == "":
		v.hallazgo(ReglaNominaRegistroPatronal, r+"/Emisor", "RegistroPatronal", nil, "con TipoContrato "+rc.TipoContrato+" se requiere RegistroPatronal")
	case (c == 9 || c == 10 || c == 99) && registro != "":
		v.hallazgoTexto(ReglaNominaRegistroPatronal, r+"/Emisor", "RegistroPatronal", n.Emisor.RegistroPatronal, "con TipoContrato "+rc.TipoContrato+" no va RegistroPatronal")
	}
	campos := []struct {
		attr  string
		tiene bool
	}{
		{"NumSeguridadSocial", rc.NumSeguridadSocial != ""}, {"FechaInicioRelLaboral", rc.FechaInicioRelLaboral.Original != ""},
		{"Antigüedad", rc.Antiguedad != ""}, {"RiesgoPuesto", rc.RiesgoPuesto != ""}, {"SalarioDiarioIntegrado", rc.SalarioDiarioIntegrado != nil},
	}
	for _, x := range campos {
		switch {
		case registro != "" && !x.tiene:
			v.hallazgo(ReglaNominaRegistroPatronal, r+"/Receptor", x.attr, nil, "con RegistroPatronal se requiere "+x.attr)
		case registro == "" && x.tiene:
			v.hallazgo(ReglaNominaRegistroPatronal, r+"/Receptor", x.attr, nil, "sin RegistroPatronal no va "+x.attr)
			v.soloGuia("guía de Nómina p. 29")
		}
	}
}

func (v *validadorReglas) entidadSNCF(n Nomina, r string) {
	if n.Emisor == nil || n.Emisor.EntidadSNCF == nil {
		return
	}
	e := n.Emisor.EntidadSNCF
	re := r + "/Emisor/EntidadSNCF"
	im := clave(e.OrigenRecurso) == "IM"
	switch {
	case im && e.MontoRecursoPropio == nil:
		v.hallazgo(ReglaNominaEntidadSNCF, re, "MontoRecursoPropio", nil, "con OrigenRecurso IM (ingresos mixtos) se requiere MontoRecursoPropio")
	case !im && e.MontoRecursoPropio != nil:
		v.hallazgo(ReglaNominaEntidadSNCF, re, "MontoRecursoPropio", e.MontoRecursoPropio, "sin OrigenRecurso IM no va MontoRecursoPropio")
	}
	if m := e.MontoRecursoPropio; m != nil && (n.TotalPercepciones != nil || n.TotalOtrosPagos != nil) {
		s := valorDe(n.TotalPercepciones).Add(valorDe(n.TotalOtrosPagos))
		if m.Cmp(s) >= 0 {
			v.hallazgo(ReglaNominaEntidadSNCF, re, "MontoRecursoPropio", m, fmt.Sprintf("MontoRecursoPropio %s no es menor que TotalPercepciones + TotalOtrosPagos, %s", m.Original(), s))
		}
	}
}

var (
	antiguedadSemanas = regexp.MustCompile(`^P([0-9]+)W$`)
	antiguedadFecha   = regexp.MustCompile(`^P(?:([0-9]+)Y)?(?:([0-9]+)M)?([0-9]+)D$`)
)

// sumarMeses suma años y meses a t sin desbordar el mes: el 31 de enero más
// un mes es el 28 (o 29) de febrero, no el 3 de marzo.
func sumarMeses(t time.Time, anios, meses int) time.Time {
	m := int(t.Month()) - 1 + anios*12 + meses
	y, mes := t.Year()+m/12, time.Month(m%12+1)
	d := t.Day()
	if ultimo := time.Date(y, mes+1, 0, 0, 0, 0, 0, time.UTC).Day(); d > ultimo {
		d = ultimo
	}
	return time.Date(y, mes, d, 0, 0, 0, 0, time.UTC)
}

func (v *validadorReglas) antiguedad(n Nomina, rc *NominaReceptor, r string) {
	inicio, fin := rc.FechaInicioRelLaboral, n.FechaFinalPago
	if !inicio.Valida() || !fin.Valida() {
		return
	}
	i := time.Date(inicio.Time.Year(), inicio.Time.Month(), inicio.Time.Day(), 0, 0, 0, 0, time.UTC)
	f := time.Date(fin.Time.Year(), fin.Time.Month(), fin.Time.Day(), 0, 0, 0, 0, time.UTC)
	if i.After(f) {
		v.hallazgoTexto(ReglaNominaAntiguedad, r, "FechaInicioRelLaboral", inicio.Original, "FechaInicioRelLaboral posterior a FechaFinalPago "+fin.Original)
		return
	}
	a := clave(rc.Antiguedad)
	dias := int(f.Sub(i).Hours()/24) + 1 // contando el día final
	if m := antiguedadSemanas.FindStringSubmatch(a); m != nil {
		if s, _ := strconv.Atoi(m[1]); s*7 > dias {
			v.hallazgoTexto(ReglaNominaAntiguedad, r, "Antigüedad", rc.Antiguedad, fmt.Sprintf("%d semanas y de FechaInicioRelLaboral a FechaFinalPago hay %d días (%d semanas)", s, dias, dias/7))
		}
		return
	}
	m := antiguedadFecha.FindStringSubmatch(a)
	if m == nil {
		return // el patrón lo revisa el esquema
	}
	anios, _ := strconv.Atoi("0" + m[1])
	meses, _ := strconv.Atoi("0" + m[2])
	d, _ := strconv.Atoi(m[3])
	if hasta := sumarMeses(i, anios, meses).AddDate(0, 0, d); hasta.After(f.AddDate(0, 0, 1)) {
		v.hallazgoTexto(ReglaNominaAntiguedad, r, "Antigüedad", rc.Antiguedad, "Antigüedad mayor que el tiempo de FechaInicioRelLaboral "+inicio.Original+" a FechaFinalPago "+fin.Original)
	}
}

func (v *validadorReglas) tipoRegimen(n Nomina, rc *NominaReceptor, r string) {
	reg := claveNum(rc.TipoRegimen)
	switch c := claveNum(rc.TipoContrato); {
	case c >= 1 && c <= 8 && reg != -1 && (reg < 2 || reg > 4):
		v.hallazgoTexto(ReglaNominaTipoRegimen, r+"/Receptor", "TipoRegimen", rc.TipoRegimen, "con TipoContrato "+rc.TipoContrato+" el TipoRegimen es 02, 03 o 04")
	case c >= 9 && reg != -1 && reg < 5:
		v.hallazgoTexto(ReglaNominaTipoRegimen, r+"/Receptor", "TipoRegimen", rc.TipoRegimen, "con TipoContrato "+rc.TipoContrato+" el TipoRegimen es de 05 a 99")
	}
	if clave(rc.TipoRegimen) == "" {
		return
	}
	tipos := map[string]int{}
	for j, o := range n.OtrosPagos {
		if _, ok := tipos[clave(o.TipoOtroPago)]; !ok {
			tipos[clave(o.TipoOtroPago)] = j
		}
	}
	_, con002 := tipos["002"]
	_, con007 := tipos["007"]
	_, con008 := tipos["008"]
	if reg == 2 {
		if !con002 && !con007 && !con008 {
			v.hallazgo(ReglaNominaTipoRegimen, r+"/OtrosPagos", "", nil, "con TipoRegimen 02 va un OtroPago 002 (subsidio para el empleo), salvo que haya 007 u 008")
			v.segunVigencia(vigenciaNOM105)
		}
		return
	}
	for _, t := range []string{"002", "007", "008"} {
		if j, ok := tipos[t]; ok {
			v.hallazgoTexto(ReglaNominaTipoRegimen, idx(r+"/OtrosPagos", "OtroPago", j), "TipoOtroPago", n.OtrosPagos[j].TipoOtroPago, "con TipoRegimen "+rc.TipoRegimen+" no va el OtroPago "+t)
			v.segunVigencia(vigenciaNOM105)
		}
	}
}

// digitoCLABE es el dígito de control de una CLABE (Banxico): pesos 3, 7 y
// 1 sobre los 17 primeros dígitos, cada producto módulo 10, y 10 menos la
// suma módulo 10.
func digitoCLABE(s string) int {
	pesos := [3]int{3, 7, 1}
	suma := 0
	for i := 0; i < 17; i++ {
		suma += int(s[i]-'0') * pesos[i%3] % 10
	}
	return (10 - suma%10) % 10
}

func soloDigitos(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return s != ""
}

func (v *validadorReglas) cuentaBancaria(rc *NominaReceptor, r string) {
	cuenta := clave(rc.CuentaBancaria)
	if cuenta == "" {
		return
	}
	banco := clave(rc.Banco) != ""
	switch utf8.RuneCountInString(cuenta) {
	case 18:
		if banco {
			v.hallazgoTexto(ReglaNominaCuentaBancaria, r, "Banco", rc.Banco, "con una CLABE (18 posiciones) no va Banco")
		}
		if soloDigitos(cuenta) && digitoCLABE(cuenta) != int(cuenta[17]-'0') {
			v.hallazgoTexto(ReglaNominaCuentaBancaria, r, "CuentaBancaria", rc.CuentaBancaria, fmt.Sprintf("el dígito de control de la CLABE es %d", digitoCLABE(cuenta)))
		}
	case 10, 11, 15, 16:
		if !banco {
			v.hallazgo(ReglaNominaCuentaBancaria, r, "Banco", nil, "con una cuenta, tarjeta o celular (10, 11, 15 o 16 posiciones) se requiere Banco")
		}
	default:
		v.hallazgoTexto(ReglaNominaCuentaBancaria, r, "CuentaBancaria", rc.CuentaBancaria, fmt.Sprintf("CuentaBancaria de %d posiciones: van 10, 11, 15, 16 o 18", utf8.RuneCountInString(cuenta)))
	}
}

func (v *validadorReglas) nodosPercepcion(p *Percepciones, r string) {
	var sueldos, separacion, jubilacion, t039, t044 bool
	for j, pc := range p.Percepcion {
		rp := idx(r, "Percepcion", j)
		tipo := clave(pc.TipoPercepcion)
		switch {
		case percepcionesSeparacion[tipo]:
			separacion = true
		case percepcionesJubilacion[tipo]:
			jubilacion = true
			t039, t044 = t039 || tipo == "039", t044 || tipo == "044"
		default:
			sueldos = true
		}
		switch {
		case tipo == "045" && pc.AccionesOTitulos == nil:
			v.hallazgo(ReglaNominaNodosPercepcion, rp, "", nil, "la percepción 045 lleva AccionesOTitulos")
		case tipo != "045" && pc.AccionesOTitulos != nil:
			v.hallazgo(ReglaNominaNodosPercepcion, rp+"/AccionesOTitulos", "", nil, "solo la percepción 045 lleva AccionesOTitulos")
		}
		switch {
		case tipo == "019" && len(pc.HorasExtra) == 0:
			v.hallazgo(ReglaNominaNodosPercepcion, rp, "", nil, "la percepción 019 lleva HorasExtra")
		case tipo != "019" && len(pc.HorasExtra) > 0:
			v.hallazgo(ReglaNominaNodosPercepcion, idx(rp, "HorasExtra", 0), "", nil, "solo la percepción 019 lleva HorasExtra")
		}
	}
	if sueldos && p.TotalSueldos == nil {
		v.hallazgo(ReglaNominaNodosPercepcion, r, "TotalSueldos", nil, "hay percepciones distintas de 022, 023, 025, 039 y 044: se requiere TotalSueldos")
	}
	if separacion {
		if p.TotalSeparacionIndemnizacion == nil {
			v.hallazgo(ReglaNominaNodosPercepcion, r, "TotalSeparacionIndemnizacion", nil, "con percepciones 022, 023 o 025 se requiere TotalSeparacionIndemnizacion")
		}
		if p.SeparacionIndemnizacion == nil {
			v.hallazgo(ReglaNominaNodosPercepcion, r, "", nil, "con percepciones 022, 023 o 025 se requiere el nodo SeparacionIndemnizacion")
		}
	}
	j := p.JubilacionPensionRetiro
	switch {
	case jubilacion && p.TotalJubilacionPensionRetiro == nil:
		v.hallazgo(ReglaNominaNodosPercepcion, r, "TotalJubilacionPensionRetiro", nil, "con percepciones 039 o 044 se requiere TotalJubilacionPensionRetiro")
	case !jubilacion && p.TotalJubilacionPensionRetiro != nil:
		v.hallazgo(ReglaNominaNodosPercepcion, r, "TotalJubilacionPensionRetiro", p.TotalJubilacionPensionRetiro, "sin percepciones 039 ni 044 no va TotalJubilacionPensionRetiro")
	}
	switch {
	case jubilacion && j == nil:
		v.hallazgo(ReglaNominaNodosPercepcion, r, "", nil, "con percepciones 039 o 044 se requiere el nodo JubilacionPensionRetiro")
	case !jubilacion && j != nil:
		v.hallazgo(ReglaNominaNodosPercepcion, r+"/JubilacionPensionRetiro", "", nil, "sin percepciones 039 ni 044 no va JubilacionPensionRetiro")
	}
	if j == nil {
		return
	}
	rj := r + "/JubilacionPensionRetiro"
	unaExhibicion := t039 || (j.TotalUnaExhibicion != nil && !t044)
	parcialidades := t044 || (j.TotalParcialidad != nil && !t039)
	if unaExhibicion {
		if j.TotalUnaExhibicion == nil {
			v.hallazgo(ReglaNominaNodosPercepcion, rj, "TotalUnaExhibicion", nil, "con la percepción 039 (una exhibición) se requiere TotalUnaExhibicion")
		}
		for _, x := range []struct {
			attr string
			d    *decimal.Decimal
		}{{"TotalParcialidad", j.TotalParcialidad}, {"MontoDiario", j.MontoDiario}} {
			if x.d != nil {
				v.hallazgo(ReglaNominaNodosPercepcion, rj, x.attr, x.d, "en una sola exhibición no va "+x.attr)
			}
		}
	}
	if parcialidades {
		if j.TotalUnaExhibicion != nil && !unaExhibicion {
			v.hallazgo(ReglaNominaNodosPercepcion, rj, "TotalUnaExhibicion", j.TotalUnaExhibicion, "en parcialidades no va TotalUnaExhibicion")
		}
		for _, x := range []struct {
			attr string
			d    *decimal.Decimal
		}{{"TotalParcialidad", j.TotalParcialidad}, {"MontoDiario", j.MontoDiario}} {
			if x.d == nil {
				v.hallazgo(ReglaNominaNodosPercepcion, rj, x.attr, nil, "en parcialidades se requiere "+x.attr)
			}
		}
	}
}

func (v *validadorReglas) incapacidades(n Nomina, r string) {
	var monetario, percepcion014, deduccion006 decimal.Decimal
	for _, in := range n.Incapacidades {
		monetario = monetario.Add(valorDe(in.ImporteMonetario))
	}
	con014, con006 := false, false
	if p := n.Percepciones; p != nil {
		for _, pc := range p.Percepcion {
			if clave(pc.TipoPercepcion) == "014" {
				con014 = true
				percepcion014 = percepcion014.Add(valorDe(pc.ImporteGravado)).Add(valorDe(pc.ImporteExento))
			}
		}
	}
	if d := n.Deducciones; d != nil {
		for _, dd := range d.Deduccion {
			if clave(dd.TipoDeduccion) == "006" {
				con006 = true
				deduccion006 = deduccion006.Add(valorDe(dd.Importe))
			}
		}
	}
	for _, x := range []struct {
		con        bool
		nodo, cual string
		suma       decimal.Decimal
	}{{con014, "/Percepciones", "la percepción 014 (subsidio por incapacidad)", percepcion014}, {con006, "/Deducciones", "la deducción 006 (descuento por incapacidad)", deduccion006}} {
		switch {
		case !x.con:
		case len(n.Incapacidades) == 0:
			v.hallazgo(ReglaNominaIncapacidades, r+x.nodo, "", nil, "con "+x.cual+" se requiere el nodo Incapacidades")
		case !monetario.Equal(x.suma):
			v.hallazgo(ReglaNominaIncapacidades, r+"/Incapacidades", "", nil, fmt.Sprintf("los ImporteMonetario de Incapacidades suman %s y %s, %s", monetario, x.cual, x.suma))
		}
	}
}

func (v *validadorReglas) compensacionSaldo(n Nomina, r string) {
	for j, o := range n.OtrosPagos {
		ro := idx(r+"/OtrosPagos", "OtroPago", j)
		switch tipo := clave(o.TipoOtroPago); {
		case tipo == "004" && o.CompensacionSaldosAFavor == nil:
			v.hallazgo(ReglaNominaCompensacionSaldo, ro, "", nil, "el OtroPago 004 lleva CompensacionSaldosAFavor")
		case tipo == "002" && o.SubsidioAlEmpleo == nil:
			v.hallazgo(ReglaNominaCompensacionSaldo, ro, "", nil, "el OtroPago 002 lleva SubsidioAlEmpleo")
		}
		c := o.CompensacionSaldosAFavor
		if c == nil || !n.FechaPago.Valida() {
			continue
		}
		anio, err := strconv.Atoi(clave(c.Anio))
		if err != nil {
			continue
		}
		pago := n.FechaPago.Time.Year()
		diciembre := n.FechaPago.Time.Month() == time.December || (n.FechaFinalPago.Valida() && n.FechaFinalPago.Time.Month() == time.December)
		if anio != pago-1 && !(anio == pago && diciembre) {
			v.hallazgoTexto(ReglaNominaCompensacionSaldo, ro+"/CompensacionSaldosAFavor", "Año", c.Anio,
				fmt.Sprintf("con FechaPago %s el Año es %d (o %d si el periodo de pago es diciembre)", n.FechaPago.Original, pago-1, pago))
			v.segunVigencia(vigenciaNOM103)
		}
	}
}

// comprobanteNomina revisa lo fijo de un CFDI de nómina 4.0 (NOM3-29); lo
// que ya revisan las reglas por TipoDeComprobante (Moneda MXN, sin
// FormaPago, CondicionesDePago ni Impuestos) y REGLA_NOMINA_CONCEPTO
// (NOM20-22) no se repite.
func (v *validadorReglas) comprobanteNomina() {
	c := v.c
	if !v.v40 || v.tipo != "N" || len(c.Nomina) == 0 {
		return
	}
	// Un atributo ausente lo reporta el esquema.
	texto := func(ruta, attr, valor, quiero string) {
		switch {
		case clave(valor) == quiero || (valor == "" && quiero != ""):
		case quiero == "":
			v.hallazgoTexto(ReglaNominaComprobante, ruta, attr, valor, "en un CFDI de nómina no va "+attr)
		default:
			v.hallazgoTexto(ReglaNominaComprobante, ruta, attr, valor, fmt.Sprintf("en un CFDI de nómina %s es %q", attr, quiero))
		}
	}
	// textoGuia es texto para lo que solo pide la guía (Regla.SoloGuia).
	textoGuia := func(ruta, attr, valor, quiero, cita string) {
		n := len(v.hs)
		texto(ruta, attr, valor, quiero)
		if len(v.hs) > n {
			v.soloGuia(cita)
		}
	}
	texto("/Comprobante", "Exportacion", c.Exportacion, "01")
	textoGuia("/Comprobante", "MetodoPago", c.MetodoPago, "PUE", "guía de Nómina p. 10")
	if c.InformacionGlobal != nil {
		v.hallazgo(ReglaNominaComprobante, "/Comprobante/InformacionGlobal", "", nil, "un CFDI de nómina no lleva InformacionGlobal")
	}
	if e := c.Emisor; e != nil {
		texto("/Comprobante/Emisor", "FacAtrAdquirente", e.FacAtrAdquirente, "")
		largo := utf8.RuneCountInString(clave(e.Rfc))
		for i, n := range c.Nomina {
			r := "/Comprobante/Complemento/Nomina/Emisor"
			if len(c.Nomina) > 1 {
				r = idx("/Comprobante/Complemento", "Nomina", i) + "/Emisor"
			}
			curp := ""
			if n.Emisor != nil {
				curp = n.Emisor.Curp
			}
			switch {
			case largo == 12 && curp != "":
				v.hallazgoTexto(ReglaNominaComprobante, r, "Curp", curp, "el emisor es persona moral (Rfc de 12): no va Curp")
			case largo == 13 && curp == "":
				v.hallazgo(ReglaNominaComprobante, r, "Curp", nil, "el emisor es persona física (Rfc de 13): se requiere Curp")
			}
		}
	}
	if rc := c.Receptor; rc != nil {
		r := "/Comprobante/Receptor"
		if l := utf8.RuneCountInString(clave(rc.Rfc)); rc.Rfc != "" && l != 13 {
			v.hallazgoTexto(ReglaNominaComprobante, r, "Rfc", rc.Rfc, "el receptor de un CFDI de nómina es persona física (Rfc de 13)")
		}
		texto(r, "RegimenFiscalReceptor", rc.RegimenFiscalReceptor, "605")
		texto(r, "UsoCFDI", rc.UsoCFDI, "CN01")
		textoGuia(r, "ResidenciaFiscal", rc.ResidenciaFiscal, "", "guía de Nómina p. 16")
		textoGuia(r, "NumRegIdTrib", rc.NumRegIdTrib, "", "guía de Nómina p. 16")
	}
	if len(c.Conceptos) != 1 {
		v.hallazgo(ReglaNominaComprobante, "/Comprobante/Conceptos", "", nil, fmt.Sprintf("un CFDI de nómina lleva un solo concepto y trae %d", len(c.Conceptos)))
		return
	}
	cp := c.Conceptos[0]
	r := "/Comprobante/Conceptos/Concepto[1]"
	texto(r, "ClaveProdServ", cp.ClaveProdServ, "84111505")
	texto(r, "NoIdentificacion", cp.NoIdentificacion, "")
	texto(r, "ClaveUnidad", cp.ClaveUnidad, "ACT")
	texto(r, "Unidad", cp.Unidad, "")
	texto(r, "ObjetoImp", cp.ObjetoImp, "01")
	if cp.Descripcion != "" && cp.Descripcion != "Pago de nómina" {
		v.hallazgoTexto(ReglaNominaComprobante, r, "Descripcion", cp.Descripcion, "en un CFDI de nómina Descripcion es \"Pago de nómina\"")
	}
	if q := cp.Cantidad; q != nil && (!q.Equal(uno) || q.Scale() != 0) {
		v.hallazgo(ReglaNominaComprobante, r, "Cantidad", q, "en un CFDI de nómina Cantidad es 1, sin decimales")
	}
	var hijos []string
	if cp.Impuestos != nil {
		hijos = append(hijos, "Impuestos")
	}
	if cp.ACuentaTerceros != nil {
		hijos = append(hijos, "ACuentaTerceros")
	}
	if len(cp.InformacionAduanera) > 0 {
		hijos = append(hijos, "InformacionAduanera")
	}
	if len(cp.CuentaPredial) > 0 {
		hijos = append(hijos, "CuentaPredial")
	}
	if cp.PorCuentadeTerceros != nil || cp.VentaVehiculos != nil || cp.HidroYPetro != nil || len(cp.ComplementoConcepto) > 0 {
		hijos = append(hijos, "ComplementoConcepto")
	}
	if len(cp.Parte) > 0 {
		hijos = append(hijos, "Parte")
	}
	for _, h := range hijos {
		v.hallazgo(ReglaNominaComprobante, r+"/"+h, "", nil, "en un CFDI de nómina el concepto no lleva el nodo "+h)
	}
}
