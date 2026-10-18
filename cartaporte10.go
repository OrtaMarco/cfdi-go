package cfdi

import "fmt"

// cartaPorte10 lee CartaPorte 1.0 (namespace .../CartaPorte) en el struct
// de 2.0 y 3.x. Tiene su propio lector porque 2.0 cambió la forma de varios
// nodos y qué atributos exige: aquí solo se exige lo que pide el XSD de 1.0.
func (l *lector) cartaPorte10(n *nodo, r string) *CartaPorte {
	ns := n.ns
	cp := &CartaPorte{
		Version:           l.versionComplemento(n, r, "1.0"),
		TranspInternac:    l.str(n, "TranspInternac"),
		EntradaSalidaMerc: l.str(n, "EntradaSalidaMerc"),
		ViaEntradaSalida:  l.str(n, "ViaEntradaSalida"),
		TotalDistRec:      l.dec(n, r, "TotalDistRec"),
	}
	l.requerido(n, r, "TranspInternac")

	if us := l.unico(n, ns, "Ubicaciones", r); us != nil {
		ru := r + "/Ubicaciones"
		for i, h := range l.hijos(us, ns, "Ubicacion", ru) {
			cp.Ubicaciones = append(cp.Ubicaciones, l.ubicacion10(h, ns, idx(ru, "Ubicacion", i)))
		}
		if len(cp.Ubicaciones) == 0 {
			l.advertir(AdvNodoFaltante, ru+"/Ubicacion", "Ubicaciones sin ninguna Ubicacion")
		}
	} else {
		l.advertir(AdvNodoFaltante, r+"/Ubicaciones", "CartaPorte sin Ubicaciones")
	}

	if m := l.unico(n, ns, "Mercancias", r); m != nil {
		cp.Mercancias = l.mercancias10(m, ns, r+"/Mercancias")
	} else {
		l.advertir(AdvNodoFaltante, r+"/Mercancias", "CartaPorte sin Mercancias")
	}

	if f := l.unico(n, ns, "FiguraTransporte", r); f != nil {
		cp.FiguraTransporte10 = l.figuraTransporte10(f, ns, r+"/FiguraTransporte")
	}
	return cp
}

func (l *lector) ubicacion10(n *nodo, ns, r string) Ubicacion {
	u := Ubicacion{
		TipoEstacion:       l.str(n, "TipoEstacion"),
		DistanciaRecorrida: l.dec(n, r, "DistanciaRecorrida"),
	}
	if o := l.unico(n, ns, "Origen", r); o != nil {
		ro := r + "/Origen"
		u.Origen = &CartaPorte10Origen{
			IDOrigen:          l.str(o, "IDOrigen"),
			RFCRemitente:      rfc(l.str(o, "RFCRemitente")),
			NombreRemitente:   l.str(o, "NombreRemitente"),
			NumRegIdTrib:      l.str(o, "NumRegIdTrib"),
			ResidenciaFiscal:  l.str(o, "ResidenciaFiscal"),
			NumEstacion:       l.str(o, "NumEstacion"),
			NombreEstacion:    l.str(o, "NombreEstacion"),
			NavegacionTrafico: l.str(o, "NavegacionTrafico"),
			FechaHoraSalida:   l.fecha(o, ro, "FechaHoraSalida", false),
		}
		l.requerido(o, ro, "FechaHoraSalida")
	}
	if d := l.unico(n, ns, "Destino", r); d != nil {
		rd := r + "/Destino"
		u.Destino = &CartaPorte10Destino{
			IDDestino:            l.str(d, "IDDestino"),
			RFCDestinatario:      rfc(l.str(d, "RFCDestinatario")),
			NombreDestinatario:   l.str(d, "NombreDestinatario"),
			NumRegIdTrib:         l.str(d, "NumRegIdTrib"),
			ResidenciaFiscal:     l.str(d, "ResidenciaFiscal"),
			NumEstacion:          l.str(d, "NumEstacion"),
			NombreEstacion:       l.str(d, "NombreEstacion"),
			NavegacionTrafico:    l.str(d, "NavegacionTrafico"),
			FechaHoraProgLlegada: l.fecha(d, rd, "FechaHoraProgLlegada", false),
		}
		l.requerido(d, rd, "FechaHoraProgLlegada")
	}
	if d := l.unico(n, ns, "Domicilio", r); d != nil {
		u.Domicilio = l.domicilio(d, r+"/Domicilio", true)
	}
	return u
}

func (l *lector) mercancias10(n *nodo, ns, r string) *CartaPorteMercancias {
	ms := &CartaPorteMercancias{
		PesoBrutoTotal:     l.dec(n, r, "PesoBrutoTotal"),
		UnidadPeso:         l.str(n, "UnidadPeso"),
		PesoNetoTotal:      l.dec(n, r, "PesoNetoTotal"),
		NumTotalMercancias: l.str(n, "NumTotalMercancias"),
		CargoPorTasacion:   l.dec(n, r, "CargoPorTasacion"),
	}
	l.requerido(n, r, "NumTotalMercancias")

	for i, h := range l.hijos(n, ns, "Mercancia", r) {
		rm := idx(r, "Mercancia", i)
		m := CartaPorteMercancia{
			BienesTransp:         l.str(h, "BienesTransp"),
			ClaveSTCC:            l.str(h, "ClaveSTCC"),
			Descripcion:          l.str(h, "Descripcion"),
			Cantidad:             l.dec(h, rm, "Cantidad"),
			ClaveUnidad:          l.str(h, "ClaveUnidad"),
			Unidad:               l.str(h, "Unidad"),
			Dimensiones:          l.str(h, "Dimensiones"),
			MaterialPeligroso:    l.str(h, "MaterialPeligroso"),
			CveMaterialPeligroso: l.str(h, "CveMaterialPeligroso"),
			Embalaje:             l.str(h, "Embalaje"),
			DescripEmbalaje:      l.str(h, "DescripEmbalaje"),
			PesoEnKg:             l.dec(h, rm, "PesoEnKg"),
			ValorMercancia:       l.dec(h, rm, "ValorMercancia"),
			Moneda:               l.str(h, "Moneda"),
			FraccionArancelaria:  l.str(h, "FraccionArancelaria"),
			UUIDComercioExt:      l.str(h, "UUIDComercioExt"),
		}
		l.requerido(h, rm, "PesoEnKg")
		for j, c := range l.hijos(h, ns, "CantidadTransporta", rm) {
			rc := idx(rm, "CantidadTransporta", j)
			m.CantidadTransporta = append(m.CantidadTransporta, CantidadTransporta{
				Cantidad:       l.dec(c, rc, "Cantidad"),
				IDOrigen:       l.str(c, "IDOrigen"),
				IDDestino:      l.str(c, "IDDestino"),
				CvesTransporte: l.str(c, "CvesTransporte"),
			})
			l.requerido(c, rc, "Cantidad", "IDOrigen", "IDDestino")
		}
		if d := l.unico(h, ns, "DetalleMercancia", rm); d != nil {
			rd := rm + "/DetalleMercancia"
			m.DetalleMercancia = &DetalleMercancia{
				UnidadPeso: l.str(d, "UnidadPeso"),
				PesoBruto:  l.dec(d, rd, "PesoBruto"),
				PesoNeto:   l.dec(d, rd, "PesoNeto"),
				PesoTara:   l.dec(d, rd, "PesoTara"),
				NumPiezas:  l.str(d, "NumPiezas"),
			}
			l.requerido(d, rd, "UnidadPeso", "PesoBruto", "PesoNeto", "PesoTara")
		}
		ms.Mercancia = append(ms.Mercancia, m)
	}
	if len(ms.Mercancia) == 0 {
		l.advertir(AdvNodoFaltante, r+"/Mercancia", "Mercancias sin ninguna Mercancia")
	}

	if a := l.unico(n, ns, "AutotransporteFederal", r); a != nil {
		ra := r + "/AutotransporteFederal"
		af := &AutotransporteFederal{
			PermSCT:         l.str(a, "PermSCT"),
			NumPermisoSCT:   l.str(a, "NumPermisoSCT"),
			NombreAseg:      l.str(a, "NombreAseg"),
			NumPolizaSeguro: l.str(a, "NumPolizaSeguro"),
		}
		l.requerido(a, ra, "PermSCT", "NumPermisoSCT", "NombreAseg", "NumPolizaSeguro")
		if iv := l.unico(a, ns, "IdentificacionVehicular", ra); iv != nil {
			ri := ra + "/IdentificacionVehicular"
			af.IdentificacionVehicular = &IdentificacionVehicular{
				ConfigVehicular: l.str(iv, "ConfigVehicular"),
				PlacaVM:         l.str(iv, "PlacaVM"),
				AnioModeloVM:    l.str(iv, "AnioModeloVM"),
			}
			l.requerido(iv, ri, "ConfigVehicular", "PlacaVM", "AnioModeloVM")
		} else {
			l.advertir(AdvNodoFaltante, ra+"/IdentificacionVehicular", "AutotransporteFederal sin IdentificacionVehicular")
		}
		if rem := l.unico(a, ns, "Remolques", ra); rem != nil {
			rr := ra + "/Remolques"
			for i, h := range l.hijos(rem, ns, "Remolque", rr) {
				af.Remolques = append(af.Remolques, Remolque{SubTipoRem: l.str(h, "SubTipoRem"), Placa: l.str(h, "Placa")})
				l.requerido(h, idx(rr, "Remolque", i), "SubTipoRem", "Placa")
			}
		}
		ms.AutotransporteFederal = af
	}

	// TransporteMaritimo de 1.0 tiene los atributos de 2.0 (y exige los
	// mismos, NumCertITC y MatriculaContenedor incluidos).
	if m := l.unico(n, ns, "TransporteMaritimo", r); m != nil {
		ms.TransporteMaritimo = l.transporteMaritimo(m, ns, r+"/TransporteMaritimo", true)
		if len(ms.TransporteMaritimo.Contenedor) == 0 {
			l.advertir(AdvNodoFaltante, r+"/TransporteMaritimo/Contenedor", "TransporteMaritimo sin ningún Contenedor")
		}
	}

	if a := l.unico(n, ns, "TransporteAereo", r); a != nil {
		ra := r + "/TransporteAereo"
		ms.TransporteAereo = &TransporteAereo{
			PermSCT:                  l.str(a, "PermSCT"),
			NumPermisoSCT:            l.str(a, "NumPermisoSCT"),
			MatriculaAeronave:        l.str(a, "MatriculaAeronave"),
			NombreAseg:               l.str(a, "NombreAseg"),
			NumPolizaSeguro:          l.str(a, "NumPolizaSeguro"),
			NumeroGuia:               l.str(a, "NumeroGuia"),
			LugarContrato:            l.str(a, "LugarContrato"),
			RFCTransportista:         rfc(l.str(a, "RFCTransportista")),
			CodigoTransportista:      l.str(a, "CodigoTransportista"),
			NumRegIdTribTranspor:     l.str(a, "NumRegIdTribTranspor"),
			ResidenciaFiscalTranspor: l.str(a, "ResidenciaFiscalTranspor"),
			NombreTransportista:      l.str(a, "NombreTransportista"),
			RFCEmbarcador:            rfc(l.str(a, "RFCEmbarcador")),
			NumRegIdTribEmbarc:       l.str(a, "NumRegIdTribEmbarc"),
			ResidenciaFiscalEmbarc:   l.str(a, "ResidenciaFiscalEmbarc"),
			NombreEmbarcador:         l.str(a, "NombreEmbarcador"),
		}
		l.requerido(a, ra, "PermSCT", "NumPermisoSCT", "MatriculaAeronave", "NumeroGuia", "CodigoTransportista")
	}

	if f := l.unico(n, ns, "TransporteFerroviario", r); f != nil {
		rf := r + "/TransporteFerroviario"
		tf := &TransporteFerroviario{
			TipoDeServicio:  l.str(f, "TipoDeServicio"),
			NombreAseg:      l.str(f, "NombreAseg"),
			NumPolizaSeguro: l.str(f, "NumPolizaSeguro"),
			Concesionario:   l.str(f, "Concesionario"),
		}
		l.requerido(f, rf, "TipoDeServicio")
		for i, h := range l.hijos(f, ns, "DerechosDePaso", rf) {
			rd := idx(rf, "DerechosDePaso", i)
			tf.DerechosDePaso = append(tf.DerechosDePaso, DerechosDePaso{
				TipoDerechoDePaso: l.str(h, "TipoDerechoDePaso"),
				KilometrajePagado: l.dec(h, rd, "KilometrajePagado"),
			})
			l.requerido(h, rd, "TipoDerechoDePaso", "KilometrajePagado")
		}
		for i, h := range l.hijos(f, ns, "Carro", rf) {
			rc := idx(rf, "Carro", i)
			ca := Carro{
				TipoCarro:           l.str(h, "TipoCarro"),
				MatriculaCarro:      l.str(h, "MatriculaCarro"),
				GuiaCarro:           l.str(h, "GuiaCarro"),
				ToneladasNetasCarro: l.dec(h, rc, "ToneladasNetasCarro"),
			}
			l.requerido(h, rc, "TipoCarro", "MatriculaCarro", "GuiaCarro", "ToneladasNetasCarro")
			for j, co := range l.hijos(h, ns, "Contenedor", rc) {
				rco := idx(rc, "Contenedor", j)
				ca.Contenedor = append(ca.Contenedor, ContenedorFerroviario{
					TipoContenedor:      l.str(co, "TipoContenedor"),
					PesoContenedorVacio: l.dec(co, rco, "PesoContenedorVacio"),
					PesoNetoMercancia:   l.dec(co, rco, "PesoNetoMercancia"),
				})
				l.requerido(co, rco, "TipoContenedor", "PesoContenedorVacio", "PesoNetoMercancia")
			}
			tf.Carro = append(tf.Carro, ca)
		}
		if len(tf.Carro) == 0 {
			l.advertir(AdvNodoFaltante, rf+"/Carro", "TransporteFerroviario sin ningún Carro")
		}
		ms.TransporteFerroviario = tf
	}
	return ms
}

func (l *lector) figuraTransporte10(n *nodo, ns, r string) *CartaPorte10FiguraTransporte {
	f := &CartaPorte10FiguraTransporte{CveTransporte: l.str(n, "CveTransporte")}
	l.requerido(n, r, "CveTransporte")
	domicilio := func(h *nodo, rh string) *Domicilio {
		if d := l.unico(h, ns, "Domicilio", rh); d != nil {
			return l.domicilio(d, rh+"/Domicilio", true)
		}
		return nil
	}
	for i, os := range l.hijos(n, ns, "Operadores", r) {
		ro := idx(r, "Operadores", i)
		hs := l.hijos(os, ns, "Operador", ro)
		if len(hs) == 0 {
			l.advertir(AdvNodoFaltante, ro+"/Operador", fmt.Sprintf("Operadores[%d] sin ningún Operador", i+1))
		}
		for j, h := range hs {
			rh := idx(ro, "Operador", j)
			f.Operadores = append(f.Operadores, CartaPorte10Operador{
				RFCOperador:              rfc(l.str(h, "RFCOperador")),
				NumLicencia:              l.str(h, "NumLicencia"),
				NombreOperador:           l.str(h, "NombreOperador"),
				NumRegIdTribOperador:     l.str(h, "NumRegIdTribOperador"),
				ResidenciaFiscalOperador: l.str(h, "ResidenciaFiscalOperador"),
				Domicilio:                domicilio(h, rh),
			})
		}
	}
	for i, h := range l.hijos(n, ns, "Propietario", r) {
		rh := idx(r, "Propietario", i)
		f.Propietario = append(f.Propietario, CartaPorte10Propietario{
			RFCPropietario:              rfc(l.str(h, "RFCPropietario")),
			NombrePropietario:           l.str(h, "NombrePropietario"),
			NumRegIdTribPropietario:     l.str(h, "NumRegIdTribPropietario"),
			ResidenciaFiscalPropietario: l.str(h, "ResidenciaFiscalPropietario"),
			Domicilio:                   domicilio(h, rh),
		})
	}
	for i, h := range l.hijos(n, ns, "Arrendatario", r) {
		rh := idx(r, "Arrendatario", i)
		f.Arrendatario = append(f.Arrendatario, CartaPorte10Arrendatario{
			RFCArrendatario:              rfc(l.str(h, "RFCArrendatario")),
			NombreArrendatario:           l.str(h, "NombreArrendatario"),
			NumRegIdTribArrendatario:     l.str(h, "NumRegIdTribArrendatario"),
			ResidenciaFiscalArrendatario: l.str(h, "ResidenciaFiscalArrendatario"),
			Domicilio:                    domicilio(h, rh),
		})
	}
	for i, h := range l.hijos(n, ns, "Notificado", r) {
		rh := idx(r, "Notificado", i)
		f.Notificado = append(f.Notificado, CartaPorte10Notificado{
			RFCNotificado:              rfc(l.str(h, "RFCNotificado")),
			NombreNotificado:           l.str(h, "NombreNotificado"),
			NumRegIdTribNotificado:     l.str(h, "NumRegIdTribNotificado"),
			ResidenciaFiscalNotificado: l.str(h, "ResidenciaFiscalNotificado"),
			Domicilio:                  domicilio(h, rh),
		})
	}
	return f
}
