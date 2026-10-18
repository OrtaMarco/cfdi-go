package cfdi

import (
	"fmt"
	"slices"
)

func (l *lector) comprobante(n *nodo) *Comprobante {
	const r = "/Comprobante"
	ns := n.ns
	c := &Comprobante{
		Version:           l.str(n, "Version"),
		Serie:             l.str(n, "Serie"),
		Folio:             l.str(n, "Folio"),
		Fecha:             l.fecha(n, r, "Fecha", false),
		Sello:             l.str(n, "Sello"),
		FormaPago:         l.str(n, "FormaPago"),
		NoCertificado:     l.str(n, "NoCertificado"),
		Certificado:       l.str(n, "Certificado"),
		CondicionesDePago: l.str(n, "CondicionesDePago"),
		SubTotal:          l.dec(n, r, "SubTotal"),
		Descuento:         l.dec(n, r, "Descuento"),
		Moneda:            l.str(n, "Moneda"),
		TipoCambio:        l.dec(n, r, "TipoCambio"),
		Total:             l.dec(n, r, "Total"),
		TipoDeComprobante: l.str(n, "TipoDeComprobante"),
		Exportacion:       l.str(n, "Exportacion"),
		MetodoPago:        l.str(n, "MetodoPago"),
		LugarExpedicion:   l.str(n, "LugarExpedicion"),
		Confirmacion:      l.str(n, "Confirmacion"),
	}
	// Las reglas que cambian entre versiones (atributos requeridos) siguen al
	// namespace, que es lo que valida el PAC.
	v4 := ns == NSCFDI40
	// CFDI 3.0 y 3.2 comparten namespace con 3.3. Se leen los atributos que
	// coinciden con 3.3 (sin distinguir mayúsculas: "total", "fecha"…) y no se
	// exige lo que 3.2 no tenía; lo exclusivo de 3.2 (formaDePago,
	// metodoDePago, Emisor/RegimenFiscal como nodo…) no se modela.
	l.antigua = !v4 && (c.Version == "3.2" || c.Version == "3.0")
	switch {
	case l.antigua:
		l.advertir(AdvVersionAntigua, r+"/@Version", fmt.Sprintf("CFDI %s: lectura parcial, solo los campos que coinciden con 3.3", c.Version))
	case c.Version != "3.3" && c.Version != "4.0":
		l.advertir(AdvVersionNoSoportada, r+"/@Version", fmt.Sprintf("versión %q; se leyó según el namespace %s", c.Version, ns))
	case (c.Version == "4.0") != v4:
		l.advertir(AdvVersionNamespace, r+"/@Version", fmt.Sprintf("Version %s en el namespace %s", c.Version, ns))
	}
	l.requerido(n, r, "Version", "Fecha", "SubTotal", "Total", "TipoDeComprobante")
	if !l.antigua {
		l.requerido(n, r, "Moneda", "LugarExpedicion")
	}
	if v4 {
		l.requerido(n, r, "Exportacion")
	}

	if g := l.unico(n, ns, "InformacionGlobal", r); g != nil {
		c.InformacionGlobal = &InformacionGlobal{
			Periodicidad: l.str(g, "Periodicidad"),
			Meses:        l.str(g, "Meses"),
			Anio:         l.str(g, "Año"),
		}
	}

	for i, rel := range l.hijos(n, ns, "CfdiRelacionados", r) {
		rr := idx(r, "CfdiRelacionados", i)
		grupo := CfdiRelacionados{TipoRelacion: l.str(rel, "TipoRelacion")}
		for j, u := range l.hijos(rel, ns, "CfdiRelacionado", rr) {
			uuid, orig := l.uuid(u, idx(rr, "CfdiRelacionado", j), "UUID")
			grupo.CfdiRelacionado = append(grupo.CfdiRelacionado, CfdiRelacionado{UUID: uuid, UUIDOriginal: orig})
		}
		c.CfdiRelacionados = append(c.CfdiRelacionados, grupo)
	}

	if e := l.unico(n, ns, "Emisor", r); e != nil {
		c.Emisor = &Emisor{
			Rfc:              rfc(l.str(e, "Rfc")),
			Nombre:           l.str(e, "Nombre"),
			RegimenFiscal:    l.str(e, "RegimenFiscal"),
			FacAtrAdquirente: l.str(e, "FacAtrAdquirente"),
		}
		l.requerido(e, r+"/Emisor", "Rfc")
		if !l.antigua {
			l.requerido(e, r+"/Emisor", "RegimenFiscal")
		}
		if v4 {
			l.requerido(e, r+"/Emisor", "Nombre")
		}
	} else {
		l.advertir(AdvNodoFaltante, r+"/Emisor", "falta el nodo Emisor")
	}

	if e := l.unico(n, ns, "Receptor", r); e != nil {
		c.Receptor = &Receptor{
			Rfc:                     rfc(l.str(e, "Rfc")),
			Nombre:                  l.str(e, "Nombre"),
			DomicilioFiscalReceptor: l.str(e, "DomicilioFiscalReceptor"),
			ResidenciaFiscal:        l.str(e, "ResidenciaFiscal"),
			NumRegIdTrib:            l.str(e, "NumRegIdTrib"),
			RegimenFiscalReceptor:   l.str(e, "RegimenFiscalReceptor"),
			UsoCFDI:                 l.str(e, "UsoCFDI"),
		}
		l.requerido(e, r+"/Receptor", "Rfc")
		if !l.antigua {
			l.requerido(e, r+"/Receptor", "UsoCFDI")
		}
		if v4 {
			l.requerido(e, r+"/Receptor", "Nombre", "DomicilioFiscalReceptor", "RegimenFiscalReceptor")
		}
	} else {
		l.advertir(AdvNodoFaltante, r+"/Receptor", "falta el nodo Receptor")
	}

	if cs := l.unico(n, ns, "Conceptos", r); cs != nil {
		rc := r + "/Conceptos"
		for i, cn := range l.hijos(cs, ns, "Concepto", rc) {
			c.Conceptos = append(c.Conceptos, l.concepto(cn, ns, idx(rc, "Concepto", i), v4, c))
		}
		if len(c.Conceptos) == 0 {
			l.advertir(AdvNodoFaltante, rc+"/Concepto", "Conceptos no tiene ningún Concepto")
		}
	} else {
		l.advertir(AdvNodoFaltante, r+"/Conceptos", "falta el nodo Conceptos")
	}

	if im := l.unico(n, ns, "Impuestos", r); im != nil {
		c.Impuestos = l.impuestos(im, ns, r+"/Impuestos")
	}

	for i, comp := range l.hijos(n, ns, "Complemento", r) {
		rr := r + "/Complemento"
		if i > 0 {
			rr = idx(r, "Complemento", i)
		}
		for _, h := range comp.hijos {
			l.complemento(c, h, rr)
		}
	}
	if c.TimbreFiscalDigital == nil {
		l.advertir(AdvSinTimbre, r+"/Complemento", "no hay TimbreFiscalDigital: el CFDI no está timbrado o se le quitó el timbre")
	}
	return c
}

func (l *lector) concepto(n *nodo, ns, r string, v4 bool, c *Comprobante) Concepto {
	cp := Concepto{
		ClaveProdServ:    l.str(n, "ClaveProdServ"),
		NoIdentificacion: l.str(n, "NoIdentificacion"),
		Cantidad:         l.dec(n, r, "Cantidad"),
		ClaveUnidad:      l.str(n, "ClaveUnidad"),
		Unidad:           l.str(n, "Unidad"),
		Descripcion:      l.str(n, "Descripcion"),
		ValorUnitario:    l.dec(n, r, "ValorUnitario"),
		Importe:          l.dec(n, r, "Importe"),
		Descuento:        l.dec(n, r, "Descuento"),
		ObjetoImp:        l.str(n, "ObjetoImp"),
	}
	if l.antigua {
		l.requerido(n, r, "Cantidad", "Unidad", "Descripcion", "ValorUnitario", "Importe")
	} else {
		l.requerido(n, r, "ClaveProdServ", "Cantidad", "ClaveUnidad", "Descripcion", "ValorUnitario", "Importe")
	}
	if v4 {
		l.requerido(n, r, "ObjetoImp")
	}

	if im := l.unico(n, ns, "Impuestos", r); im != nil {
		ri := r + "/Impuestos"
		ic := &ImpuestosConcepto{}
		if ts := l.unico(im, ns, "Traslados", ri); ts != nil {
			ic.Traslados = l.listaImpuestos(ts, ns, "Traslado", ri+"/Traslados", true)
		}
		if rs := l.unico(im, ns, "Retenciones", ri); rs != nil {
			ic.Retenciones = l.listaImpuestos(rs, ns, "Retencion", ri+"/Retenciones", false)
		}
		cp.Impuestos = ic
	}

	if t := l.unico(n, ns, "ACuentaTerceros", r); t != nil {
		cp.ACuentaTerceros = &ACuentaTerceros{
			RfcACuentaTerceros:             rfc(l.str(t, "RfcACuentaTerceros")),
			NombreACuentaTerceros:          l.str(t, "NombreACuentaTerceros"),
			RegimenFiscalACuentaTerceros:   l.str(t, "RegimenFiscalACuentaTerceros"),
			DomicilioFiscalACuentaTerceros: l.str(t, "DomicilioFiscalACuentaTerceros"),
		}
	}

	cp.InformacionAduanera = l.aduanera(n, ns, r)

	cuentas := l.hijos(n, ns, "CuentaPredial", r)
	if len(cuentas) > 1 && !v4 {
		l.advertir(AdvNodoDuplicado, r+"/CuentaPredial", fmt.Sprintf("CFDI 3.3 admite una CuentaPredial y vienen %d; se conservan todas", len(cuentas)))
	}
	for _, cu := range cuentas {
		cp.CuentaPredial = append(cp.CuentaPredial, CuentaPredial{Numero: l.str(cu, "Numero")})
	}

	if cc := l.unico(n, ns, "ComplementoConcepto", r); cc != nil {
		for _, h := range cc.hijos {
			l.complementoConcepto(&cp, c, h, r+"/ComplementoConcepto")
		}
	}

	for i, p := range l.hijos(n, ns, "Parte", r) {
		rp := idx(r, "Parte", i)
		cp.Parte = append(cp.Parte, Parte{
			ClaveProdServ:       l.str(p, "ClaveProdServ"),
			NoIdentificacion:    l.str(p, "NoIdentificacion"),
			Cantidad:            l.dec(p, rp, "Cantidad"),
			Unidad:              l.str(p, "Unidad"),
			Descripcion:         l.str(p, "Descripcion"),
			ValorUnitario:       l.dec(p, rp, "ValorUnitario"),
			Importe:             l.dec(p, rp, "Importe"),
			InformacionAduanera: l.aduanera(p, ns, rp),
		})
	}
	return cp
}

func (l *lector) aduanera(n *nodo, ns, r string) []InformacionAduanera {
	var out []InformacionAduanera
	for _, a := range l.hijos(n, ns, "InformacionAduanera", r) {
		out = append(out, InformacionAduanera{NumeroPedimento: l.str(a, "NumeroPedimento")})
	}
	return out
}

// impuestos lee un nodo Impuestos de totales (del comprobante o de un Pago 1.0).
func (l *lector) impuestos(n *nodo, ns, r string) *Impuestos {
	im := &Impuestos{
		TotalImpuestosRetenidos:   l.dec(n, r, "TotalImpuestosRetenidos"),
		TotalImpuestosTrasladados: l.dec(n, r, "TotalImpuestosTrasladados"),
	}
	if rs := l.unico(n, ns, "Retenciones", r); rs != nil {
		im.Retenciones = l.listaImpuestos(rs, ns, "Retencion", r+"/Retenciones", false)
	}
	if ts := l.unico(n, ns, "Traslados", r); ts != nil {
		im.Traslados = l.listaImpuestos(ts, ns, "Traslado", r+"/Traslados", true)
	}
	return im
}

// listaImpuestos lee los Traslado o Retencion de un contenedor. Un traslado
// Exento no lleva TasaOCuota ni Importe; si los trae se conservan y se advierte.
// Un traslado Tasa/Cuota sin Importe también se advierte.
func (l *lector) listaImpuestos(n *nodo, ns, local, r string, traslado bool) []Impuesto {
	var out []Impuesto
	for i, h := range l.hijos(n, ns, local, r) {
		ri := idx(r, local, i)
		im := Impuesto{
			Base:       l.dec(h, ri, "Base"),
			Impuesto:   l.str(h, "Impuesto"),
			TipoFactor: l.str(h, "TipoFactor"),
			TasaOCuota: l.dec(h, ri, "TasaOCuota"),
			Importe:    l.dec(h, ri, "Importe"),
		}
		l.requerido(h, ri, "Impuesto")
		if traslado {
			l.revisarExento(h, ri, im.TipoFactor, "TasaOCuota", "Importe")
		} else {
			l.requerido(h, ri, "Importe")
		}
		out = append(out, im)
	}
	return out
}

// revisarExento aplica la regla del TipoFactor Exento a cualquier traslado
// (concepto, comprobante, ImpuestosDR, ImpuestosP): sin tasa ni importe.
func (l *lector) revisarExento(h *nodo, r, tipoFactor, attrTasa, attrImporte string) {
	_, conTasa := h.attr(attrTasa)
	_, conImporte := h.attr(attrImporte)
	if tipoFactor == "Exento" {
		if conTasa || conImporte {
			l.advertir(AdvExentoConImporte, r, "traslado Exento con "+attrTasa+" o "+attrImporte)
		}
		return
	}
	if !conImporte {
		l.advertir(AdvAtributoFaltante, r+"/@"+attrImporte, "traslado "+tipoFactor+" sin "+attrImporte)
	}
}

func agregarUnico(lista *[]string, s string) {
	if !slices.Contains(*lista, s) {
		*lista = append(*lista, s)
	}
}
