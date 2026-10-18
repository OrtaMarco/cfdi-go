package cfdi

import "fmt"

// complemento despacha un hijo de cfdi:Complemento por namespace + nombre
// local. Pago existe en Pagos 1.0 y 2.0: se distingue por el namespace del
// nodo Pagos y sus hijos se buscan en ese mismo namespace.
func (l *lector) complemento(c *Comprobante, h *nodo, rc string) {
	r := rc + "/" + h.local
	switch {
	case h.ns == NSTimbre && h.local == "TimbreFiscalDigital":
		if c.TimbreFiscalDigital != nil {
			l.advertir(AdvNodoDuplicado, r, "más de un TimbreFiscalDigital; se usa el primero")
			return
		}
		c.TimbreFiscalDigital = l.timbre(h, r)
	case (h.ns == NSPagos10 || h.ns == NSPagos20) && h.local == "Pagos":
		if c.Pagos != nil {
			l.advertir(AdvNodoDuplicado, r, "más de un complemento Pagos; se usa el primero")
			return
		}
		c.Pagos = l.pagos(h, r)
	case h.ns == NSNomina12 && h.local == "Nomina":
		if len(c.Nomina) > 0 {
			r = idx(rc, "Nomina", len(c.Nomina))
		}
		c.Nomina = append(c.Nomina, l.nomina(h, r))
	case h.ns == NSNomina11 && h.local == "Nomina":
		if len(c.Nomina) > 0 {
			r = idx(rc, "Nomina", len(c.Nomina))
		}
		c.Nomina = append(c.Nomina, l.nomina11(h, r))
	case h.ns == NSImpLocal && h.local == "ImpuestosLocales":
		if !l.yaLeido(c.ImpuestosLocales != nil, r, h.local) {
			c.ImpuestosLocales = l.impuestosLocales(h, r)
		}
	case (h.ns == NSCartaPorte20 || h.ns == NSCartaPorte30 || h.ns == NSCartaPorte31) && h.local == "CartaPorte":
		if !l.yaLeido(c.CartaPorte != nil, r, h.local) {
			c.CartaPorte = l.cartaPorte(h, r)
		}
	case h.ns == NSCartaPorte10 && h.local == "CartaPorte":
		if !l.yaLeido(c.CartaPorte != nil, r, h.local) {
			c.CartaPorte = l.cartaPorte10(h, r)
		}
	case (h.ns == NSComercioExterior10 || h.ns == NSComercioExterior11 || h.ns == NSComercioExterior20) && h.local == "ComercioExterior":
		if !l.yaLeido(c.ComercioExterior != nil, r, h.local) {
			c.ComercioExterior = l.comercioExterior(h, r)
		}
	case h.ns == NSDivisas && h.local == "Divisas":
		if !l.yaLeido(c.Divisas != nil, r, h.local) {
			c.Divisas = l.divisas(h, r)
		}
	case h.ns == NSRegistroFiscal && h.local == "CFDIRegistroFiscal":
		if !l.yaLeido(c.CFDIRegistroFiscal != nil, r, h.local) {
			c.CFDIRegistroFiscal = l.registroFiscal(h, r)
		}
	case h.ns == NSLeyendasFiscales && h.local == "LeyendasFiscales":
		if !l.yaLeido(c.LeyendasFiscales != nil, r, h.local) {
			c.LeyendasFiscales = l.leyendasFiscales(h, r)
		}
	case (h.ns == NSEstadoDeCuentaCombustible11 || h.ns == NSEstadoDeCuentaCombustible12) && h.local == "EstadoDeCuentaCombustible":
		if !l.yaLeido(c.EstadoDeCuentaCombustible != nil, r, h.local) {
			c.EstadoDeCuentaCombustible = l.estadoDeCuentaCombustible(h, r)
		}
	case h.ns == NSHidroYPetro && h.local == "HidroYPetro":
		if len(c.HidroYPetro) > 0 {
			r = idx(rc, h.local, len(c.HidroYPetro))
		}
		c.HidroYPetro = append(c.HidroYPetro, *l.hidroYPetro(h, r))
	case h.ns == NSAerolineas && h.local == "Aerolineas":
		if !l.yaLeido(c.Aerolineas != nil, r, h.local) {
			c.Aerolineas = l.aerolineas(h, r)
		}
	case h.ns == NSDonatarias && h.local == "Donatarias":
		if !l.yaLeido(c.Donatarias != nil, r, h.local) {
			c.Donatarias = l.donatarias(h, r)
		}
	case h.ns == NSINE && h.local == "INE":
		if !l.yaLeido(c.INE != nil, r, h.local) {
			c.INE = l.ine(h, r)
		}
	case h.ns == NSValesDeDespensa && h.local == "ValesDeDespensa":
		if !l.yaLeido(c.ValesDeDespensa != nil, r, h.local) {
			c.ValesDeDespensa = l.valesDeDespensa(h, r)
		}
	case (h.ns == NSConsumoDeCombustibles10 || h.ns == NSConsumoDeCombustibles11) && h.local == "ConsumoDeCombustibles":
		if !l.yaLeido(c.ConsumoDeCombustibles != nil, r, h.local) {
			c.ConsumoDeCombustibles = l.consumoDeCombustibles(h, r)
		}
	case h.ns == NSPagoEnEspecie && h.local == "PagoEnEspecie":
		if !l.yaLeido(c.PagoEnEspecie != nil, r, h.local) {
			c.PagoEnEspecie = l.pagoEnEspecie(h, r)
		}
	case h.ns == NSNotariosPublicos && h.local == "NotariosPublicos":
		if !l.yaLeido(c.NotariosPublicos != nil, r, h.local) {
			c.NotariosPublicos = l.notariosPublicos(h, r)
		}
	case h.ns == NSServicioParcialConstruccion && h.local == "parcialesconstruccion":
		if !l.yaLeido(c.ParcialesConstruccion != nil, r, h.local) {
			c.ParcialesConstruccion = l.parcialesConstruccion(h, r)
		}
	case h.ns == NSVehiculoUsado && h.local == "VehiculoUsado":
		if !l.yaLeido(c.VehiculoUsado != nil, r, h.local) {
			c.VehiculoUsado = l.vehiculoUsado(h, r)
		}
	case h.ns == NSTuristaPasajeroExtranjero && h.local == "TuristaPasajeroExtranjero":
		if !l.yaLeido(c.TuristaPasajeroExtranjero != nil, r, h.local) {
			c.TuristaPasajeroExtranjero = l.turistaPasajeroExtranjero(h, r)
		}
	case h.ns == NSSPEI && h.local == "Complemento_SPEI":
		if !l.yaLeido(c.ComplementoSPEI != nil, r, h.local) {
			c.ComplementoSPEI = l.complementoSPEI(h, r)
		}
	case h.ns == NSRenovacionVehiculos && h.local == "renovacionysustitucionvehiculos":
		if !l.yaLeido(c.RenovacionySustitucionVehiculos != nil, r, h.local) {
			c.RenovacionySustitucionVehiculos = l.renovacionySustitucionVehiculos(h, r)
		}
	case h.ns == NSCertificadoDestruccion && h.local == "certificadodedestruccion":
		if !l.yaLeido(c.CertificadoDeDestruccion != nil, r, h.local) {
			c.CertificadoDeDestruccion = l.certificadoDeDestruccion(h, r)
		}
	case h.ns == NSObrasArteAntiguedades && h.local == "obrasarteantiguedades":
		if !l.yaLeido(c.ObrasArteAntiguedades != nil, r, h.local) {
			c.ObrasArteAntiguedades = l.obrasArteAntiguedades(h, r)
		}
	case h.ns == NSPFIC && h.local == "PFintegranteCoordinado":
		if !l.yaLeido(c.PFintegranteCoordinado != nil, r, h.local) {
			c.PFintegranteCoordinado = l.pfIntegranteCoordinado(h, r)
		}
	case h.ns == NSGastosHidrocarburos && h.local == "GastosHidrocarburos":
		if !l.yaLeido(c.GastosHidrocarburos != nil, r, h.local) {
			c.GastosHidrocarburos = l.gastosHidrocarburos(h, r)
		}
	case h.ns == NSIngresosHidrocarburos && h.local == "IngresosHidrocarburos":
		if !l.yaLeido(c.IngresosHidrocarburos != nil, r, h.local) {
			c.IngresosHidrocarburos = l.ingresosHidrocarburos(h, r)
		}
	case h.ns == NSEstadoDeCuentaBancario && h.local == "EstadoDeCuentaBancario":
		if !l.yaLeido(c.EstadoDeCuentaBancario != nil, r, h.local) {
			c.EstadoDeCuentaBancario = l.estadoDeCuentaBancario(h, r)
		}
	case h.ns == NSEstadoDeCuentaCombustible10 && h.local == "EstadoDeCuentaCombustible":
		if !l.yaLeido(c.EstadoDeCuentaCombustible != nil, r, h.local) {
			c.EstadoDeCuentaCombustible = l.estadoDeCuentaCombustible10(h, r)
		}
	default:
		nom := h.nombre()
		agregarUnico(&c.ComplementosNoSoportados, nom)
		l.advertir(AdvComplementoNoSoportado, r, "complemento no soportado: "+nom)
	}
}

// yaLeido advierte NODO_DUPLICADO si el complemento ya se leyó: solo cuenta
// el primero.
func (l *lector) yaLeido(ya bool, r, nombre string) bool {
	if ya {
		l.advertir(AdvNodoDuplicado, r, "más de un complemento "+nombre+"; se usa el primero")
	}
	return ya
}

func (l *lector) versionComplemento(n *nodo, r string, validas ...string) string {
	v := l.str(n, "Version")
	for _, ok := range validas {
		if v == ok {
			return v
		}
	}
	l.advertir(AdvComplementoVersion, r+"/@Version", fmt.Sprintf("versión %q no esperada para {%s}%s", v, n.ns, n.local))
	return v
}

func (l *lector) timbre(n *nodo, r string) *TimbreFiscalDigital {
	t := &TimbreFiscalDigital{
		Version:          l.versionComplemento(n, r, "1.0", "1.1"),
		FechaTimbrado:    l.fecha(n, r, "FechaTimbrado", false),
		RfcProvCertif:    rfc(l.str(n, "RfcProvCertif")),
		Leyenda:          l.str(n, "Leyenda"),
		SelloCFD:         l.str(n, "SelloCFD"),
		NoCertificadoSAT: l.str(n, "NoCertificadoSAT"),
		SelloSAT:         l.str(n, "SelloSAT"),
	}
	t.UUID, t.UUIDOriginal = l.uuid(n, r, "UUID")
	l.requerido(n, r, "UUID", "FechaTimbrado")
	return t
}

func (l *lector) pagos(n *nodo, r string) *Pagos {
	ns := n.ns
	v20 := ns == NSPagos20
	esperada := "1.0"
	if v20 {
		esperada = "2.0"
	}
	p := &Pagos{Version: l.versionComplemento(n, r, esperada)}

	if v20 {
		if t := l.unico(n, ns, "Totales", r); t != nil {
			rt := r + "/Totales"
			p.Totales = &TotalesPagos{
				TotalRetencionesIVA:         l.dec(t, rt, "TotalRetencionesIVA"),
				TotalRetencionesISR:         l.dec(t, rt, "TotalRetencionesISR"),
				TotalRetencionesIEPS:        l.dec(t, rt, "TotalRetencionesIEPS"),
				TotalTrasladosBaseIVA16:     l.dec(t, rt, "TotalTrasladosBaseIVA16"),
				TotalTrasladosImpuestoIVA16: l.dec(t, rt, "TotalTrasladosImpuestoIVA16"),
				TotalTrasladosBaseIVA8:      l.dec(t, rt, "TotalTrasladosBaseIVA8"),
				TotalTrasladosImpuestoIVA8:  l.dec(t, rt, "TotalTrasladosImpuestoIVA8"),
				TotalTrasladosBaseIVA0:      l.dec(t, rt, "TotalTrasladosBaseIVA0"),
				TotalTrasladosImpuestoIVA0:  l.dec(t, rt, "TotalTrasladosImpuestoIVA0"),
				TotalTrasladosBaseIVAExento: l.dec(t, rt, "TotalTrasladosBaseIVAExento"),
				MontoTotalPagos:             l.dec(t, rt, "MontoTotalPagos"),
			}
			l.requerido(t, rt, "MontoTotalPagos")
		} else {
			l.advertir(AdvNodoFaltante, r+"/Totales", "Pagos 2.0 sin nodo Totales")
		}
	}

	pagos := l.hijos(n, ns, "Pago", r)
	if len(pagos) == 0 {
		l.advertir(AdvNodoFaltante, r+"/Pago", "Pagos sin ningún Pago")
	}
	for i, pn := range pagos {
		rp := idx(r, "Pago", i)
		pg := Pago{
			FechaPago:       l.fecha(pn, rp, "FechaPago", false),
			FormaDePagoP:    l.str(pn, "FormaDePagoP"),
			MonedaP:         l.str(pn, "MonedaP"),
			TipoCambioP:     l.dec(pn, rp, "TipoCambioP"),
			Monto:           l.dec(pn, rp, "Monto"),
			NumOperacion:    l.str(pn, "NumOperacion"),
			RfcEmisorCtaOrd: rfc(l.str(pn, "RfcEmisorCtaOrd")),
			NomBancoOrdExt:  l.str(pn, "NomBancoOrdExt"),
			CtaOrdenante:    l.str(pn, "CtaOrdenante"),
			RfcEmisorCtaBen: rfc(l.str(pn, "RfcEmisorCtaBen")),
			CtaBeneficiario: l.str(pn, "CtaBeneficiario"),
			TipoCadPago:     l.str(pn, "TipoCadPago"),
			CertPago:        l.str(pn, "CertPago"),
			CadPago:         l.str(pn, "CadPago"),
			SelloPago:       l.str(pn, "SelloPago"),
		}
		l.requerido(pn, rp, "FechaPago", "FormaDePagoP", "MonedaP", "Monto")

		for j, dn := range l.hijos(pn, ns, "DoctoRelacionado", rp) {
			pg.DoctoRelacionado = append(pg.DoctoRelacionado, l.docto(dn, ns, idx(rp, "DoctoRelacionado", j), v20))
		}

		if v20 {
			if ip := l.unico(pn, ns, "ImpuestosP", rp); ip != nil {
				pg.ImpuestosP = l.impuestosP(ip, ns, rp+"/ImpuestosP")
			}
		} else {
			for j, in := range l.hijos(pn, ns, "Impuestos", rp) {
				pg.Impuestos = append(pg.Impuestos, *l.impuestos(in, ns, idx(rp, "Impuestos", j)))
			}
		}
		p.Pago = append(p.Pago, pg)
	}
	return p
}

func (l *lector) docto(n *nodo, ns, r string, v20 bool) DoctoRelacionado {
	d := DoctoRelacionado{
		Serie:            l.str(n, "Serie"),
		Folio:            l.str(n, "Folio"),
		MonedaDR:         l.str(n, "MonedaDR"),
		EquivalenciaDR:   l.dec(n, r, "EquivalenciaDR"),
		TipoCambioDR:     l.dec(n, r, "TipoCambioDR"),
		MetodoDePagoDR:   l.str(n, "MetodoDePagoDR"),
		NumParcialidad:   l.str(n, "NumParcialidad"),
		ImpSaldoAnt:      l.dec(n, r, "ImpSaldoAnt"),
		ImpPagado:        l.dec(n, r, "ImpPagado"),
		ImpSaldoInsoluto: l.dec(n, r, "ImpSaldoInsoluto"),
		ObjetoImpDR:      l.str(n, "ObjetoImpDR"),
	}
	d.IdDocumento, d.IdDocumentoOriginal = l.uuid(n, r, "IdDocumento")
	l.requerido(n, r, "IdDocumento", "MonedaDR")
	if !v20 {
		return d
	}
	l.requerido(n, r, "ObjetoImpDR")
	if in := l.unico(n, ns, "ImpuestosDR", r); in != nil {
		ri := r + "/ImpuestosDR"
		imp := &ImpuestosDR{}
		if rs := l.unico(in, ns, "RetencionesDR", ri); rs != nil {
			imp.RetencionesDR = l.listaDR(rs, ns, "RetencionDR", ri+"/RetencionesDR", false)
		}
		if ts := l.unico(in, ns, "TrasladosDR", ri); ts != nil {
			imp.TrasladosDR = l.listaDR(ts, ns, "TrasladoDR", ri+"/TrasladosDR", true)
		}
		d.ImpuestosDR = imp
	}
	return d
}

func (l *lector) listaDR(n *nodo, ns, local, r string, traslado bool) []ImpuestoDR {
	var out []ImpuestoDR
	for i, h := range l.hijos(n, ns, local, r) {
		ri := idx(r, local, i)
		im := ImpuestoDR{
			BaseDR:       l.dec(h, ri, "BaseDR"),
			ImpuestoDR:   l.str(h, "ImpuestoDR"),
			TipoFactorDR: l.str(h, "TipoFactorDR"),
			TasaOCuotaDR: l.dec(h, ri, "TasaOCuotaDR"),
			ImporteDR:    l.dec(h, ri, "ImporteDR"),
		}
		if traslado {
			l.revisarExento(h, ri, im.TipoFactorDR, "TasaOCuotaDR", "ImporteDR")
		} else {
			l.requerido(h, ri, "ImporteDR")
		}
		out = append(out, im)
	}
	return out
}

func (l *lector) impuestosP(n *nodo, ns, r string) *ImpuestosP {
	imp := &ImpuestosP{}
	lista := func(cont, local string, traslado bool) []ImpuestoP {
		c := l.unico(n, ns, cont, r)
		if c == nil {
			return nil
		}
		rc := r + "/" + cont
		var out []ImpuestoP
		for i, h := range l.hijos(c, ns, local, rc) {
			ri := idx(rc, local, i)
			im := ImpuestoP{
				BaseP:       l.dec(h, ri, "BaseP"),
				ImpuestoP:   l.str(h, "ImpuestoP"),
				TipoFactorP: l.str(h, "TipoFactorP"),
				TasaOCuotaP: l.dec(h, ri, "TasaOCuotaP"),
				ImporteP:    l.dec(h, ri, "ImporteP"),
			}
			if traslado {
				l.revisarExento(h, ri, im.TipoFactorP, "TasaOCuotaP", "ImporteP")
			} else {
				l.requerido(h, ri, "ImporteP")
			}
			out = append(out, im)
		}
		return out
	}
	imp.RetencionesP = lista("RetencionesP", "RetencionP", false)
	imp.TrasladosP = lista("TrasladosP", "TrasladoP", true)
	return imp
}
