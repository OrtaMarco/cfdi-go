package cfdi

// complementoConcepto despacha un hijo de Concepto/ComplementoConcepto. Los no
// soportados quedan en Concepto.ComplementoConcepto y en
// Comprobante.ComplementosNoSoportados.
func (l *lector) complementoConcepto(cp *Concepto, c *Comprobante, h *nodo, rc string) {
	r := rc + "/" + h.local
	switch {
	case h.ns == NSTerceros && h.local == "PorCuentadeTerceros":
		if !l.yaLeido(cp.PorCuentadeTerceros != nil, r, h.local) {
			cp.PorCuentadeTerceros = l.porCuentadeTerceros(h, r)
		}
	case h.ns == NSVentaVehiculos && h.local == "VentaVehiculos":
		if !l.yaLeido(cp.VentaVehiculos != nil, r, h.local) {
			cp.VentaVehiculos = l.ventaVehiculos(h, r)
		}
	case h.ns == NSHidroYPetro && h.local == "HidroYPetro":
		if !l.yaLeido(cp.HidroYPetro != nil, r, h.local) {
			cp.HidroYPetro = l.hidroYPetro(h, r)
		}
	case h.ns == NSIEDU && h.local == "instEducativas":
		if !l.yaLeido(cp.InstEducativas != nil, r, h.local) {
			cp.InstEducativas = l.instEducativas(h, r)
		}
	case h.ns == NSAcreditamientoIEPS && h.local == "acreditamientoIEPS":
		if !l.yaLeido(cp.AcreditamientoIEPS != nil, r, h.local) {
			cp.AcreditamientoIEPS = l.acreditamientoIEPS(h, r)
		}
	case h.ns == NSServicioParcialConstruccion && h.local == "parcialesconstruccion":
		if !l.yaLeido(cp.ParcialesConstruccion != nil, r, h.local) {
			cp.ParcialesConstruccion = l.parcialesConstruccion(h, r)
		}
	default:
		nom := h.nombre()
		cp.ComplementoConcepto = append(cp.ComplementoConcepto, nom)
		agregarUnico(&c.ComplementosNoSoportados, nom)
		l.advertir(AdvComplementoNoSoportado, r, "complemento de concepto no soportado: "+nom)
	}
}

func (l *lector) porCuentadeTerceros(n *nodo, r string) *PorCuentadeTerceros {
	ns := n.ns
	t := &PorCuentadeTerceros{
		Version: l.versionComplemento(n, r, "1.1"),
		Rfc:     rfc(l.str(n, "rfc")),
		Nombre:  l.str(n, "nombre"),
	}
	l.requerido(n, r, "rfc")

	if u := l.unico(n, ns, "InformacionFiscalTercero", r); u != nil {
		ru := r + "/InformacionFiscalTercero"
		t.InformacionFiscalTercero = &UbicacionFiscal{
			Calle:        l.str(u, "calle"),
			NoExterior:   l.str(u, "noExterior"),
			NoInterior:   l.str(u, "noInterior"),
			Colonia:      l.str(u, "colonia"),
			Localidad:    l.str(u, "localidad"),
			Referencia:   l.str(u, "referencia"),
			Municipio:    l.str(u, "municipio"),
			Estado:       l.str(u, "estado"),
			Pais:         l.str(u, "pais"),
			CodigoPostal: l.str(u, "codigoPostal"),
		}
		l.requerido(u, ru, "calle", "municipio", "estado", "pais", "codigoPostal")
	}
	if a := l.unico(n, ns, "InformacionAduanera", r); a != nil {
		ia := l.aduaneraComplemento(a, r+"/InformacionAduanera")
		t.InformacionAduanera = &ia
	}
	t.Parte = l.partesComplemento(n, ns, r)
	if cu := l.unico(n, ns, "CuentaPredial", r); cu != nil {
		t.CuentaPredial = &CuentaPredial{Numero: l.str(cu, "numero")}
		l.requerido(cu, r+"/CuentaPredial", "numero")
	}

	im := l.unico(n, ns, "Impuestos", r)
	if im == nil {
		l.advertir(AdvNodoFaltante, r+"/Impuestos", "PorCuentadeTerceros sin Impuestos")
		return t
	}
	ri := r + "/Impuestos"
	t.Impuestos = &ImpuestosTerceros{}
	lista := func(cont, local string, conTasa bool) []ImpuestoTerceros {
		cn := l.unico(im, ns, cont, ri)
		if cn == nil {
			return nil
		}
		rc := ri + "/" + cont
		var out []ImpuestoTerceros
		for i, h := range l.hijos(cn, ns, local, rc) {
			rh := idx(rc, local, i)
			it := ImpuestoTerceros{
				Impuesto: l.str(h, "impuesto"),
				Tasa:     l.dec(h, rh, "tasa"),
				Importe:  l.dec(h, rh, "importe"),
			}
			l.requerido(h, rh, "impuesto", "importe")
			if conTasa {
				l.requerido(h, rh, "tasa")
			}
			out = append(out, it)
		}
		return out
	}
	t.Impuestos.Retenciones = lista("Retenciones", "Retencion", false)
	t.Impuestos.Traslados = lista("Traslados", "Traslado", true)
	return t
}

func (l *lector) ventaVehiculos(n *nodo, r string) *VentaVehiculos {
	ns := n.ns
	v := &VentaVehiculos{
		Version:        l.versionComplemento(n, r, "1.0", "1.1"),
		ClaveVehicular: l.str(n, "ClaveVehicular"),
		Niv:            l.str(n, "Niv"),
	}
	l.requerido(n, r, "ClaveVehicular")
	if v.Version != "1.0" {
		l.requerido(n, r, "Niv")
	}
	for i, a := range l.hijos(n, ns, "InformacionAduanera", r) {
		v.InformacionAduanera = append(v.InformacionAduanera, l.aduaneraComplemento(a, idx(r, "InformacionAduanera", i)))
	}
	v.Parte = l.partesComplemento(n, ns, r)
	return v
}

func (l *lector) hidroYPetro(n *nodo, r string) *HidroYPetro {
	h := &HidroYPetro{
		Version:        l.versionComplemento(n, r, "1.0"),
		TipoPermiso:    l.str(n, "TipoPermiso"),
		NumeroPermiso:  l.str(n, "NumeroPermiso"),
		ClaveHYP:       l.str(n, "ClaveHYP"),
		SubProductoHYP: l.str(n, "SubProductoHYP"),
	}
	l.requerido(n, r, "TipoPermiso", "NumeroPermiso", "ClaveHYP", "SubProductoHYP")
	return h
}

func (l *lector) instEducativas(n *nodo, r string) *InstEducativas {
	ie := &InstEducativas{
		Version:        l.versionComplemento(n, r, "1.0"),
		NombreAlumno:   l.str(n, "nombreAlumno"),
		CURP:           l.str(n, "CURP"),
		NivelEducativo: l.str(n, "nivelEducativo"),
		AutRVOE:        l.str(n, "autRVOE"),
		RfcPago:        rfc(l.str(n, "rfcPago")),
	}
	l.requerido(n, r, "nombreAlumno", "CURP", "nivelEducativo", "autRVOE")
	return ie
}

func (l *lector) aduaneraComplemento(n *nodo, r string) InformacionAduaneraComplemento {
	a := InformacionAduaneraComplemento{
		Numero: l.str(n, "numero"),
		Fecha:  l.fecha(n, r, "fecha", true),
		Aduana: l.str(n, "aduana"),
	}
	l.requerido(n, r, "numero", "fecha")
	return a
}

func (l *lector) partesComplemento(n *nodo, ns, r string) []ParteComplemento {
	var out []ParteComplemento
	for i, p := range l.hijos(n, ns, "Parte", r) {
		rp := idx(r, "Parte", i)
		pc := ParteComplemento{
			Cantidad:         l.dec(p, rp, "cantidad"),
			Unidad:           l.str(p, "unidad"),
			NoIdentificacion: l.str(p, "noIdentificacion"),
			Descripcion:      l.str(p, "descripcion"),
			ValorUnitario:    l.dec(p, rp, "valorUnitario"),
			Importe:          l.dec(p, rp, "importe"),
		}
		l.requerido(p, rp, "cantidad", "descripcion")
		for j, a := range l.hijos(p, ns, "InformacionAduanera", rp) {
			pc.InformacionAduanera = append(pc.InformacionAduanera, l.aduaneraComplemento(a, idx(rp, "InformacionAduanera", j)))
		}
		out = append(out, pc)
	}
	return out
}
