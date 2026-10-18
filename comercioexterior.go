package cfdi

// comercioExterior lee ComercioExterior 1.0, 1.1 o 2.0; la versión la dice
// el namespace y solo cambia qué atributos se exigen.
func (l *lector) comercioExterior(n *nodo, r string) *ComercioExterior {
	ns := n.ns
	v20 := ns == NSComercioExterior20
	v10 := ns == NSComercioExterior10
	esperada := "1.1"
	switch {
	case v20:
		esperada = "2.0"
	case v10:
		esperada = "1.0"
	}
	ce := &ComercioExterior{
		Version:                   l.versionComplemento(n, r, esperada),
		MotivoTraslado:            l.str(n, "MotivoTraslado"),
		TipoOperacion:             l.str(n, "TipoOperacion"),
		ClaveDePedimento:          l.str(n, "ClaveDePedimento"),
		CertificadoOrigen:         l.str(n, "CertificadoOrigen"),
		NumCertificadoOrigen:      l.str(n, "NumCertificadoOrigen"),
		NumeroExportadorConfiable: l.str(n, "NumeroExportadorConfiable"),
		Incoterm:                  l.str(n, "Incoterm"),
		Subdivision:               l.str(n, "Subdivision"),
		Observaciones:             l.str(n, "Observaciones"),
		TipoCambioUSD:             l.dec(n, r, "TipoCambioUSD"),
		TotalUSD:                  l.dec(n, r, "TotalUSD"),
	}
	if v20 {
		l.requerido(n, r, "ClaveDePedimento", "CertificadoOrigen", "TipoCambioUSD", "TotalUSD")
	} else {
		l.requerido(n, r, "TipoOperacion")
	}

	if e := l.unico(n, ns, "Emisor", r); e != nil {
		re := r + "/Emisor"
		ce.Emisor = &ComercioExteriorEmisor{Curp: l.str(e, "Curp")}
		if d := l.unico(e, ns, "Domicilio", re); d != nil {
			ce.Emisor.Domicilio = l.domicilio(d, re+"/Domicilio", true)
		} else if v20 {
			l.advertir(AdvNodoFaltante, re+"/Domicilio", "ComercioExterior 2.0: Emisor sin Domicilio")
		}
	}
	for i, p := range l.hijos(n, ns, "Propietario", r) {
		ce.Propietario = append(ce.Propietario, ComercioExteriorPropietario{
			NumRegIdTrib:     l.str(p, "NumRegIdTrib"),
			ResidenciaFiscal: l.str(p, "ResidenciaFiscal"),
		})
		l.requerido(p, idx(r, "Propietario", i), "NumRegIdTrib", "ResidenciaFiscal")
	}
	if e := l.unico(n, ns, "Receptor", r); e != nil {
		re := r + "/Receptor"
		ce.Receptor = &ComercioExteriorReceptor{Curp: l.str(e, "Curp"), NumRegIdTrib: l.str(e, "NumRegIdTrib")}
		if v10 {
			l.requerido(e, re, "NumRegIdTrib")
		}
		if d := l.unico(e, ns, "Domicilio", re); d != nil {
			ce.Receptor.Domicilio = l.domicilio(d, re+"/Domicilio", true)
		}
	} else if v10 {
		l.advertir(AdvNodoFaltante, r+"/Receptor", "ComercioExterior 1.0 sin Receptor")
	}
	for i, d := range l.hijos(n, ns, "Destinatario", r) {
		rd := idx(r, "Destinatario", i)
		de := ComercioExteriorDestinatario{
			NumRegIdTrib: l.str(d, "NumRegIdTrib"),
			Rfc:          rfc(l.str(d, "Rfc")),
			Curp:         l.str(d, "Curp"),
			Nombre:       l.str(d, "Nombre"),
		}
		for j, dom := range l.hijos(d, ns, "Domicilio", rd) {
			de.Domicilio = append(de.Domicilio, *l.domicilio(dom, idx(rd, "Domicilio", j), true))
		}
		if len(de.Domicilio) == 0 {
			l.advertir(AdvNodoFaltante, rd+"/Domicilio", "Destinatario sin Domicilio")
		}
		ce.Destinatario = append(ce.Destinatario, de)
	}

	if ms := l.unico(n, ns, "Mercancias", r); ms != nil {
		rm := r + "/Mercancias"
		for i, m := range l.hijos(ms, ns, "Mercancia", rm) {
			ri := idx(rm, "Mercancia", i)
			me := ComercioExteriorMercancia{
				NoIdentificacion:    l.str(m, "NoIdentificacion"),
				FraccionArancelaria: l.str(m, "FraccionArancelaria"),
				CantidadAduana:      l.dec(m, ri, "CantidadAduana"),
				UnidadAduana:        l.str(m, "UnidadAduana"),
				ValorUnitarioAduana: l.dec(m, ri, "ValorUnitarioAduana"),
				ValorDolares:        l.dec(m, ri, "ValorDolares"),
			}
			l.requerido(m, ri, "NoIdentificacion", "ValorDolares")
			for j, de := range l.hijos(m, ns, "DescripcionesEspecificas", ri) {
				me.DescripcionesEspecificas = append(me.DescripcionesEspecificas, DescripcionesEspecificas{
					Marca:       l.str(de, "Marca"),
					Modelo:      l.str(de, "Modelo"),
					SubModelo:   l.str(de, "SubModelo"),
					NumeroSerie: l.str(de, "NumeroSerie"),
				})
				l.requerido(de, idx(ri, "DescripcionesEspecificas", j), "Marca")
			}
			ce.Mercancias = append(ce.Mercancias, me)
		}
		if len(ce.Mercancias) == 0 {
			l.advertir(AdvNodoFaltante, rm+"/Mercancia", "Mercancias sin ninguna Mercancia")
		}
	} else if v20 {
		l.advertir(AdvNodoFaltante, r+"/Mercancias", "ComercioExterior 2.0 sin Mercancias")
	}
	return ce
}
