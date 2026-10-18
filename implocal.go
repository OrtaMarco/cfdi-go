package cfdi

func (l *lector) impuestosLocales(n *nodo, r string) *ImpuestosLocales {
	ns := n.ns
	il := &ImpuestosLocales{
		Version:            l.versionComplemento(n, r, "1.0"),
		TotaldeRetenciones: l.dec(n, r, "TotaldeRetenciones"),
		TotaldeTraslados:   l.dec(n, r, "TotaldeTraslados"),
	}
	l.requerido(n, r, "TotaldeRetenciones", "TotaldeTraslados")
	for i, h := range l.hijos(n, ns, "RetencionesLocales", r) {
		ri := idx(r, "RetencionesLocales", i)
		il.RetencionesLocales = append(il.RetencionesLocales, RetencionesLocales{
			ImpLocRetenido:  l.str(h, "ImpLocRetenido"),
			TasadeRetencion: l.dec(h, ri, "TasadeRetencion"),
			Importe:         l.dec(h, ri, "Importe"),
		})
		l.requerido(h, ri, "ImpLocRetenido", "TasadeRetencion", "Importe")
	}
	for i, h := range l.hijos(n, ns, "TrasladosLocales", r) {
		ri := idx(r, "TrasladosLocales", i)
		il.TrasladosLocales = append(il.TrasladosLocales, TrasladosLocales{
			ImpLocTrasladado: l.str(h, "ImpLocTrasladado"),
			TasadeTraslado:   l.dec(h, ri, "TasadeTraslado"),
			Importe:          l.dec(h, ri, "Importe"),
		})
		l.requerido(h, ri, "ImpLocTrasladado", "TasadeTraslado", "Importe")
	}
	return il
}
