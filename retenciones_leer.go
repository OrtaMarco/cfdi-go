package cfdi

import (
	"fmt"
	"strings"
	"time"
)

// nombresRet son los atributos que cambian de nombre entre Retenciones 1.0 y
// 2.0. Las rutas de las advertencias usan el del documento.
type nombresRet struct {
	noCert, cert, rfcE, nacionalidad, rfcR, curpR, numRegIdTrib, ejercicio          string
	montoTotOperacion, montoTotGrav, montoTotExent, montoTotRet, montoRet, impuesto string
}

var (
	nombresRet10 = nombresRet{"NumCert", "Cert", "RFCEmisor", "Nacionalidad", "RFCRecep", "CURPR", "NumRegIdTrib", "Ejerc",
		"montoTotOperacion", "montoTotGrav", "montoTotExent", "montoTotRet", "montoRet", "Impuesto"}
	nombresRet20 = nombresRet{"NoCertificado", "Certificado", "RfcE", "NacionalidadR", "RfcR", "CurpR", "NumRegIdTribR", "Ejercicio",
		"MontoTotOperacion", "MontoTotGrav", "MontoTotExent", "MontoTotRet", "MontoRet", "ImpuestoRet"}
)

func (l *lector) retenciones(n *nodo) *Retenciones {
	const r = "/Retenciones"
	ns := n.ns
	v2 := ns == NSRetenciones20
	nom := nombresRet10
	esperada := "1.0"
	if v2 {
		nom, esperada = nombresRet20, "2.0"
	}
	rt := &Retenciones{
		Version:        l.str(n, "Version"),
		FolioInt:       l.str(n, "FolioInt"),
		Sello:          l.str(n, "Sello"),
		NoCertificado:  l.str(n, nom.noCert),
		Certificado:    l.str(n, nom.cert),
		LugarExpRetenc: l.str(n, "LugarExpRetenc"),
		CveRetenc:      l.str(n, "CveRetenc"),
		DescRetenc:     l.str(n, "DescRetenc"),
	}
	// Las reglas que cambian entre versiones siguen al namespace.
	switch {
	case rt.Version != "1.0" && rt.Version != "2.0":
		l.advertir(AdvVersionNoSoportada, r+"/@Version", fmt.Sprintf("versión %q; se leyó según el namespace %s", rt.Version, ns))
	case rt.Version != esperada:
		l.advertir(AdvVersionNamespace, r+"/@Version", fmt.Sprintf("Version %s en el namespace %s", rt.Version, ns))
	}
	if v2 {
		rt.FechaExp = l.fecha(n, r, "FechaExp", false)
	} else {
		rt.FechaExp = l.fechaConZona(n, r, "FechaExp")
	}
	l.requerido(n, r, "Version", "Sello", nom.noCert, nom.cert, "FechaExp", "CveRetenc")
	if v2 {
		l.requerido(n, r, "LugarExpRetenc")
	}

	if v2 {
		if rel := l.unico(n, ns, "CfdiRetenRelacionados", r); rel != nil {
			rr := r + "/CfdiRetenRelacionados"
			c := &CfdiRetenRelacionados{TipoRelacion: l.str(rel, "TipoRelacion")}
			c.UUID, c.UUIDOriginal = l.uuid(rel, rr, "UUID")
			l.requerido(rel, rr, "TipoRelacion", "UUID")
			rt.CfdiRetenRelacionados = c
		}
	}

	if e := l.unico(n, ns, "Emisor", r); e != nil {
		re := r + "/Emisor"
		rt.Emisor = &EmisorRetenciones{
			RfcE:           rfc(l.str(e, nom.rfcE)),
			NomDenRazSocE:  l.str(e, "NomDenRazSocE"),
			RegimenFiscalE: l.str(e, "RegimenFiscalE"),
			CurpE:          l.str(e, "CURPE"),
		}
		l.requerido(e, re, nom.rfcE)
		if v2 {
			l.requerido(e, re, "NomDenRazSocE", "RegimenFiscalE")
		}
	} else {
		l.advertir(AdvNodoFaltante, r+"/Emisor", "falta el nodo Emisor")
	}

	if e := l.unico(n, ns, "Receptor", r); e != nil {
		rr := r + "/Receptor"
		rec := &ReceptorRetenciones{NacionalidadR: l.str(e, nom.nacionalidad)}
		l.requerido(e, rr, nom.nacionalidad)
		if na := l.unico(e, ns, "Nacional", rr); na != nil {
			rn := rr + "/Nacional"
			rec.Nacional = &ReceptorNacionalRetenciones{
				RfcR:             rfc(l.str(na, nom.rfcR)),
				NomDenRazSocR:    l.str(na, "NomDenRazSocR"),
				CurpR:            l.str(na, nom.curpR),
				DomicilioFiscalR: l.str(na, "DomicilioFiscalR"),
			}
			l.requerido(na, rn, nom.rfcR)
			if v2 {
				l.requerido(na, rn, "NomDenRazSocR", "DomicilioFiscalR")
			}
		}
		if ex := l.unico(e, ns, "Extranjero", rr); ex != nil {
			rec.Extranjero = &ReceptorExtranjeroRetenciones{
				NumRegIdTribR: l.str(ex, nom.numRegIdTrib),
				NomDenRazSocR: l.str(ex, "NomDenRazSocR"),
			}
			l.requerido(ex, rr+"/Extranjero", "NomDenRazSocR")
		}
		if rec.Nacional == nil && rec.Extranjero == nil {
			l.advertir(AdvNodoFaltante, rr+"/Nacional", "el Receptor no trae Nacional ni Extranjero")
		}
		rt.Receptor = rec
	} else {
		l.advertir(AdvNodoFaltante, r+"/Receptor", "falta el nodo Receptor")
	}

	if p := l.unico(n, ns, "Periodo", r); p != nil {
		rt.Periodo = &PeriodoRetenciones{
			MesIni:    l.str(p, "MesIni"),
			MesFin:    l.str(p, "MesFin"),
			Ejercicio: l.str(p, nom.ejercicio),
		}
		l.requerido(p, r+"/Periodo", "MesIni", "MesFin", nom.ejercicio)
	} else {
		l.advertir(AdvNodoFaltante, r+"/Periodo", "falta el nodo Periodo")
	}

	if t := l.unico(n, ns, "Totales", r); t != nil {
		rtot := r + "/Totales"
		tot := &TotalesRetenciones{
			MontoTotOperacion:  l.dec(t, rtot, nom.montoTotOperacion),
			MontoTotGrav:       l.dec(t, rtot, nom.montoTotGrav),
			MontoTotExent:      l.dec(t, rtot, nom.montoTotExent),
			MontoTotRet:        l.dec(t, rtot, nom.montoTotRet),
			UtilidadBimestral:  l.dec(t, rtot, "UtilidadBimestral"),
			ISRCorrespondiente: l.dec(t, rtot, "ISRCorrespondiente"),
		}
		l.requerido(t, rtot, nom.montoTotOperacion, nom.montoTotGrav, nom.montoTotExent, nom.montoTotRet)
		for i, im := range l.hijos(t, ns, "ImpRetenidos", rtot) {
			ri := idx(rtot, "ImpRetenidos", i)
			tot.ImpRetenidos = append(tot.ImpRetenidos, ImpRetenidos{
				BaseRet:     l.dec(im, ri, "BaseRet"),
				ImpuestoRet: l.str(im, nom.impuesto),
				MontoRet:    l.dec(im, ri, nom.montoRet),
				TipoPagoRet: l.str(im, "TipoPagoRet"),
			})
			l.requerido(im, ri, nom.montoRet, "TipoPagoRet")
		}
		rt.Totales = tot
	} else {
		l.advertir(AdvNodoFaltante, r+"/Totales", "falta el nodo Totales")
	}

	for i, comp := range l.hijos(n, ns, "Complemento", r) {
		rc := r + "/Complemento"
		if i > 0 {
			rc = idx(r, "Complemento", i)
		}
		for _, h := range comp.hijos {
			l.complementoRetenciones(rt, h, rc)
		}
	}
	if rt.TimbreFiscalDigital == nil {
		l.advertir(AdvSinTimbre, r+"/Complemento", "no hay TimbreFiscalDigital: el CFDI no está timbrado o se le quitó el timbre")
	}
	return rt
}

// layoutFechaHoraZona es la FechaExp de Retenciones 1.0: con la zona del
// lugar de expedición (el XSD la exige).
const layoutFechaHoraZona = "2006-01-02T15:04:05-07:00"

// fechaConZona lee una fecha que el SAT pide con zona (FechaExp de
// Retenciones 1.0). Time es la hora tal cual viene escrita, como en Fecha;
// la zona se queda en el texto. Sin zona, o en otro formato, FECHA_FORMATO.
func (l *lector) fechaConZona(n *nodo, ruta, nombre string) Fecha {
	v, ok := n.attr(nombre)
	if !ok {
		return Fecha{}
	}
	if t, err := time.Parse(layoutFechaHoraZona, v); err == nil && !strings.Contains(v, ".") {
		return Fecha{Original: v, Time: time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, time.UTC)}
	}
	f, _, err := parseFecha(v, false)
	if err != nil {
		l.advertirValor(AdvFechaInvalida, ruta+"/@"+nombre, err.Error(), v)
		return f
	}
	l.advertirValor(AdvFechaFormato, ruta+"/@"+nombre, fmt.Sprintf("%q no tiene el formato %s", v, layoutFechaHoraZona), v)
	f.FueraDeFormato = true
	return f
}
