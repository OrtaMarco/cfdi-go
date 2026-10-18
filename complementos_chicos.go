package cfdi

// Complementos chicos: un nodo y pocos atributos.

func (l *lector) divisas(n *nodo, r string) *Divisas {
	d := &Divisas{
		Version:       l.versionComplemento(n, r, "1.0"),
		TipoOperacion: l.str(n, "tipoOperacion"),
	}
	l.requerido(n, r, "tipoOperacion")
	return d
}

func (l *lector) registroFiscal(n *nodo, r string) *CFDIRegistroFiscal {
	rf := &CFDIRegistroFiscal{
		Version: l.versionComplemento(n, r, "1.0"),
		Folio:   l.str(n, "Folio"),
	}
	l.requerido(n, r, "Folio")
	return rf
}

// estadoDeCuentaCombustible lee ECC 1.2 o 1.1 (la versión la dice el
// namespace). En 1.1 el concepto trae TAR y NoIdentificacion en vez de
// TipoCombustible, y el traslado TasaoCuota (con o minúscula).
func (l *lector) estadoDeCuentaCombustible(n *nodo, r string) *EstadoDeCuentaCombustible {
	ns := n.ns
	v11 := ns == NSEstadoDeCuentaCombustible11
	esperada, tasa := "1.2", "TasaOCuota"
	if v11 {
		esperada, tasa = "1.1", "TasaoCuota"
	}
	e := &EstadoDeCuentaCombustible{
		Version:        l.versionComplemento(n, r, esperada),
		TipoOperacion:  l.str(n, "TipoOperacion"),
		NumeroDeCuenta: l.str(n, "NumeroDeCuenta"),
		SubTotal:       l.dec(n, r, "SubTotal"),
		Total:          l.dec(n, r, "Total"),
	}
	l.requerido(n, r, "TipoOperacion", "NumeroDeCuenta", "SubTotal", "Total")
	cs := l.unico(n, ns, "Conceptos", r)
	if cs == nil {
		l.advertir(AdvNodoFaltante, r+"/Conceptos", "EstadoDeCuentaCombustible sin Conceptos")
		return e
	}
	rc := r + "/Conceptos"
	for i, h := range l.hijos(cs, ns, "ConceptoEstadoDeCuentaCombustible", rc) {
		ri := idx(rc, "ConceptoEstadoDeCuentaCombustible", i)
		cp := ConceptoEstadoDeCuentaCombustible{
			Identificador:     l.str(h, "Identificador"),
			Fecha:             l.fecha(h, ri, "Fecha", false),
			Rfc:               rfc(l.str(h, "Rfc")),
			ClaveEstacion:     l.str(h, "ClaveEstacion"),
			TAR:               l.str(h, "TAR"),
			Cantidad:          l.dec(h, ri, "Cantidad"),
			TipoCombustible:   l.str(h, "TipoCombustible"),
			NoIdentificacion:  l.str(h, "NoIdentificacion"),
			Unidad:            l.str(h, "Unidad"),
			NombreCombustible: l.str(h, "NombreCombustible"),
			FolioOperacion:    l.str(h, "FolioOperacion"),
			ValorUnitario:     l.dec(h, ri, "ValorUnitario"),
			Importe:           l.dec(h, ri, "Importe"),
		}
		l.requerido(h, ri, "Identificador", "Fecha", "Rfc", "ClaveEstacion", "Cantidad",
			"NombreCombustible", "FolioOperacion", "ValorUnitario", "Importe")
		if v11 {
			l.requerido(h, ri, "NoIdentificacion")
		} else {
			l.requerido(h, ri, "TipoCombustible")
		}
		if ts := l.unico(h, ns, "Traslados", ri); ts != nil {
			rt := ri + "/Traslados"
			for j, t := range l.hijos(ts, ns, "Traslado", rt) {
				rj := idx(rt, "Traslado", j)
				cp.Traslados = append(cp.Traslados, Impuesto{
					Impuesto:   l.str(t, "Impuesto"),
					TasaOCuota: l.dec(t, rj, tasa),
					Importe:    l.dec(t, rj, "Importe"),
				})
				l.requerido(t, rj, "Impuesto", tasa, "Importe")
			}
		} else {
			l.advertir(AdvNodoFaltante, ri+"/Traslados", "concepto de estado de cuenta sin Traslados")
		}
		e.Conceptos = append(e.Conceptos, cp)
	}
	if len(e.Conceptos) == 0 {
		l.advertir(AdvNodoFaltante, rc+"/ConceptoEstadoDeCuentaCombustible", "Conceptos sin ningún concepto")
	}
	return e
}

func (l *lector) aerolineas(n *nodo, r string) *Aerolineas {
	ns := n.ns
	a := &Aerolineas{
		Version: l.versionComplemento(n, r, "1.0"),
		TUA:     l.dec(n, r, "TUA"),
	}
	l.requerido(n, r, "TUA")
	if oc := l.unico(n, ns, "OtrosCargos", r); oc != nil {
		ro := r + "/OtrosCargos"
		a.OtrosCargos = &AerolineasOtrosCargos{TotalCargos: l.dec(oc, ro, "TotalCargos")}
		l.requerido(oc, ro, "TotalCargos")
		for i, h := range l.hijos(oc, ns, "Cargo", ro) {
			ri := idx(ro, "Cargo", i)
			a.OtrosCargos.Cargo = append(a.OtrosCargos.Cargo, AerolineasCargo{
				CodigoCargo: l.str(h, "CodigoCargo"),
				Importe:     l.dec(h, ri, "Importe"),
			})
			l.requerido(h, ri, "CodigoCargo", "Importe")
		}
		if len(a.OtrosCargos.Cargo) == 0 {
			l.advertir(AdvNodoFaltante, ro+"/Cargo", "OtrosCargos sin ningún Cargo")
		}
	}
	return a
}

func (l *lector) donatarias(n *nodo, r string) *Donatarias {
	d := &Donatarias{
		Version:           l.versionComplemento(n, r, "1.1", "1.0"),
		NoAutorizacion:    l.str(n, "noAutorizacion"),
		FechaAutorizacion: l.fecha(n, r, "fechaAutorizacion", true),
		Leyenda:           l.str(n, "leyenda"),
	}
	l.requerido(n, r, "noAutorizacion", "fechaAutorizacion", "leyenda")
	return d
}

func (l *lector) leyendasFiscales(n *nodo, r string) *LeyendasFiscales {
	lf := &LeyendasFiscales{Version: l.versionComplemento(n, r, "1.0")}
	for i, h := range l.hijos(n, n.ns, "Leyenda", r) {
		lf.Leyenda = append(lf.Leyenda, Leyenda{
			DisposicionFiscal: l.str(h, "disposicionFiscal"),
			Norma:             l.str(h, "norma"),
			TextoLeyenda:      l.str(h, "textoLeyenda"),
		})
		l.requerido(h, idx(r, "Leyenda", i), "textoLeyenda")
	}
	if len(lf.Leyenda) == 0 {
		l.advertir(AdvNodoFaltante, r+"/Leyenda", "LeyendasFiscales sin ninguna Leyenda")
	}
	return lf
}
