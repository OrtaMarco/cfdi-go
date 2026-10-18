package cfdi

import "fmt"

// cartaPorte lee CartaPorte 2.0, 3.0 o 3.1. La versión la dice el namespace,
// como en Pagos; los nodos y atributos de otra versión que vengan se leen
// igual (el struct es la unión), y solo se exige lo que pide el XSD de la
// versión del namespace.
func (l *lector) cartaPorte(n *nodo, r string) *CartaPorte {
	ns := n.ns
	v2 := ns == NSCartaPorte20
	esperada := map[string]string{NSCartaPorte20: "2.0", NSCartaPorte30: "3.0", NSCartaPorte31: "3.1"}[ns]
	cp := &CartaPorte{
		Version:              l.versionComplemento(n, r, esperada),
		IdCCP:                l.str(n, "IdCCP"),
		TranspInternac:       l.str(n, "TranspInternac"),
		RegimenAduanero:      l.str(n, "RegimenAduanero"),
		EntradaSalidaMerc:    l.str(n, "EntradaSalidaMerc"),
		PaisOrigenDestino:    l.str(n, "PaisOrigenDestino"),
		ViaEntradaSalida:     l.str(n, "ViaEntradaSalida"),
		TotalDistRec:         l.dec(n, r, "TotalDistRec"),
		RegistroISTMO:        l.str(n, "RegistroISTMO"),
		UbicacionPoloOrigen:  l.str(n, "UbicacionPoloOrigen"),
		UbicacionPoloDestino: l.str(n, "UbicacionPoloDestino"),
	}
	l.requerido(n, r, "TranspInternac")
	if !v2 {
		l.requerido(n, r, "IdCCP")
	}

	if ra := l.unico(n, ns, "RegimenesAduaneros", r); ra != nil {
		rr := r + "/RegimenesAduaneros"
		for i, h := range l.hijos(ra, ns, "RegimenAduaneroCCP", rr) {
			cp.RegimenesAduaneros = append(cp.RegimenesAduaneros, RegimenAduaneroCCP{RegimenAduanero: l.str(h, "RegimenAduanero")})
			l.requerido(h, idx(rr, "RegimenAduaneroCCP", i), "RegimenAduanero")
		}
	}

	if us := l.unico(n, ns, "Ubicaciones", r); us != nil {
		ru := r + "/Ubicaciones"
		for i, h := range l.hijos(us, ns, "Ubicacion", ru) {
			cp.Ubicaciones = append(cp.Ubicaciones, l.ubicacion(h, ns, idx(ru, "Ubicacion", i)))
		}
		if len(cp.Ubicaciones) < 2 {
			l.advertir(AdvNodoFaltante, ru+"/Ubicacion", fmt.Sprintf("CartaPorte con %d Ubicacion; el mínimo es 2 (origen y destino)", len(cp.Ubicaciones)))
		}
	} else {
		l.advertir(AdvNodoFaltante, r+"/Ubicaciones", "CartaPorte sin Ubicaciones")
	}

	if m := l.unico(n, ns, "Mercancias", r); m != nil {
		cp.Mercancias = l.cartaPorteMercancias(m, ns, r+"/Mercancias", v2)
	} else {
		l.advertir(AdvNodoFaltante, r+"/Mercancias", "CartaPorte sin Mercancias")
	}

	if f := l.unico(n, ns, "FiguraTransporte", r); f != nil {
		rf := r + "/FiguraTransporte"
		for i, h := range l.hijos(f, ns, "TiposFigura", rf) {
			ri := idx(rf, "TiposFigura", i)
			tf := TiposFigura{
				TipoFigura:             l.str(h, "TipoFigura"),
				RFCFigura:              rfc(l.str(h, "RFCFigura")),
				NumLicencia:            l.str(h, "NumLicencia"),
				NombreFigura:           l.str(h, "NombreFigura"),
				NumRegIdTribFigura:     l.str(h, "NumRegIdTribFigura"),
				ResidenciaFiscalFigura: l.str(h, "ResidenciaFiscalFigura"),
			}
			l.requerido(h, ri, "TipoFigura")
			if !v2 {
				l.requerido(h, ri, "NombreFigura")
			}
			for j, p := range l.hijos(h, ns, "PartesTransporte", ri) {
				tf.PartesTransporte = append(tf.PartesTransporte, PartesTransporte{ParteTransporte: l.str(p, "ParteTransporte")})
				l.requerido(p, idx(ri, "PartesTransporte", j), "ParteTransporte")
			}
			if d := l.unico(h, ns, "Domicilio", ri); d != nil {
				tf.Domicilio = l.domicilio(d, ri+"/Domicilio", false)
			}
			cp.FiguraTransporte = append(cp.FiguraTransporte, tf)
		}
	}
	return cp
}

// domicilio lee un Domicilio de carta porte o comercio exterior. Estado, Pais
// y CodigoPostal siempre son requeridos; Calle solo en comercio exterior.
func (l *lector) domicilio(n *nodo, r string, conCalle bool) *Domicilio {
	d := &Domicilio{
		Calle:          l.str(n, "Calle"),
		NumeroExterior: l.str(n, "NumeroExterior"),
		NumeroInterior: l.str(n, "NumeroInterior"),
		Colonia:        l.str(n, "Colonia"),
		Localidad:      l.str(n, "Localidad"),
		Referencia:     l.str(n, "Referencia"),
		Municipio:      l.str(n, "Municipio"),
		Estado:         l.str(n, "Estado"),
		Pais:           l.str(n, "Pais"),
		CodigoPostal:   l.str(n, "CodigoPostal"),
	}
	if conCalle {
		l.requerido(n, r, "Calle")
	}
	l.requerido(n, r, "Estado", "Pais", "CodigoPostal")
	return d
}

func (l *lector) ubicacion(n *nodo, ns, r string) Ubicacion {
	u := Ubicacion{
		TipoUbicacion:               l.str(n, "TipoUbicacion"),
		IDUbicacion:                 l.str(n, "IDUbicacion"),
		RFCRemitenteDestinatario:    rfc(l.str(n, "RFCRemitenteDestinatario")),
		NombreRemitenteDestinatario: l.str(n, "NombreRemitenteDestinatario"),
		NumRegIdTrib:                l.str(n, "NumRegIdTrib"),
		ResidenciaFiscal:            l.str(n, "ResidenciaFiscal"),
		NumEstacion:                 l.str(n, "NumEstacion"),
		NombreEstacion:              l.str(n, "NombreEstacion"),
		NavegacionTrafico:           l.str(n, "NavegacionTrafico"),
		FechaHoraSalidaLlegada:      l.fecha(n, r, "FechaHoraSalidaLlegada", false),
		TipoEstacion:                l.str(n, "TipoEstacion"),
		DistanciaRecorrida:          l.dec(n, r, "DistanciaRecorrida"),
	}
	l.requerido(n, r, "TipoUbicacion", "RFCRemitenteDestinatario", "FechaHoraSalidaLlegada")
	if d := l.unico(n, ns, "Domicilio", r); d != nil {
		u.Domicilio = l.domicilio(d, r+"/Domicilio", false)
	}
	return u
}

func (l *lector) cartaPorteMercancias(n *nodo, ns, r string, v2 bool) *CartaPorteMercancias {
	ms := &CartaPorteMercancias{
		PesoBrutoTotal:                        l.dec(n, r, "PesoBrutoTotal"),
		UnidadPeso:                            l.str(n, "UnidadPeso"),
		PesoNetoTotal:                         l.dec(n, r, "PesoNetoTotal"),
		NumTotalMercancias:                    l.str(n, "NumTotalMercancias"),
		CargoPorTasacion:                      l.dec(n, r, "CargoPorTasacion"),
		LogisticaInversaRecoleccionDevolucion: l.str(n, "LogisticaInversaRecoleccionDevolucion"),
	}
	l.requerido(n, r, "PesoBrutoTotal", "UnidadPeso", "NumTotalMercancias")

	for i, h := range l.hijos(n, ns, "Mercancia", r) {
		ms.Mercancia = append(ms.Mercancia, l.cartaPorteMercancia(h, ns, idx(r, "Mercancia", i)))
	}
	if len(ms.Mercancia) == 0 {
		l.advertir(AdvNodoFaltante, r+"/Mercancia", "Mercancias sin ninguna Mercancia")
	}

	if a := l.unico(n, ns, "Autotransporte", r); a != nil {
		ra := r + "/Autotransporte"
		at := &Autotransporte{PermSCT: l.str(a, "PermSCT"), NumPermisoSCT: l.str(a, "NumPermisoSCT")}
		l.requerido(a, ra, "PermSCT", "NumPermisoSCT")
		if iv := l.unico(a, ns, "IdentificacionVehicular", ra); iv != nil {
			ri := ra + "/IdentificacionVehicular"
			at.IdentificacionVehicular = &IdentificacionVehicular{
				ConfigVehicular:    l.str(iv, "ConfigVehicular"),
				PesoBrutoVehicular: l.dec(iv, ri, "PesoBrutoVehicular"),
				PlacaVM:            l.str(iv, "PlacaVM"),
				AnioModeloVM:       l.str(iv, "AnioModeloVM"),
			}
			l.requerido(iv, ri, "ConfigVehicular", "PlacaVM", "AnioModeloVM")
			if !v2 {
				l.requerido(iv, ri, "PesoBrutoVehicular")
			}
		} else {
			l.advertir(AdvNodoFaltante, ra+"/IdentificacionVehicular", "Autotransporte sin IdentificacionVehicular")
		}
		if s := l.unico(a, ns, "Seguros", ra); s != nil {
			rs := ra + "/Seguros"
			at.Seguros = &Seguros{
				AseguraRespCivil:   l.str(s, "AseguraRespCivil"),
				PolizaRespCivil:    l.str(s, "PolizaRespCivil"),
				AseguraMedAmbiente: l.str(s, "AseguraMedAmbiente"),
				PolizaMedAmbiente:  l.str(s, "PolizaMedAmbiente"),
				AseguraCarga:       l.str(s, "AseguraCarga"),
				PolizaCarga:        l.str(s, "PolizaCarga"),
				PrimaSeguro:        l.dec(s, rs, "PrimaSeguro"),
			}
			l.requerido(s, rs, "AseguraRespCivil", "PolizaRespCivil")
		} else {
			l.advertir(AdvNodoFaltante, ra+"/Seguros", "Autotransporte sin Seguros")
		}
		if rem := l.unico(a, ns, "Remolques", ra); rem != nil {
			rr := ra + "/Remolques"
			for i, h := range l.hijos(rem, ns, "Remolque", rr) {
				at.Remolques = append(at.Remolques, Remolque{SubTipoRem: l.str(h, "SubTipoRem"), Placa: l.str(h, "Placa")})
				l.requerido(h, idx(rr, "Remolque", i), "SubTipoRem", "Placa")
			}
		}
		ms.Autotransporte = at
	}

	if m := l.unico(n, ns, "TransporteMaritimo", r); m != nil {
		ms.TransporteMaritimo = l.transporteMaritimo(m, ns, r+"/TransporteMaritimo", v2)
	}

	if a := l.unico(n, ns, "TransporteAereo", r); a != nil {
		ra := r + "/TransporteAereo"
		ms.TransporteAereo = &TransporteAereo{
			PermSCT:                l.str(a, "PermSCT"),
			NumPermisoSCT:          l.str(a, "NumPermisoSCT"),
			MatriculaAeronave:      l.str(a, "MatriculaAeronave"),
			NombreAseg:             l.str(a, "NombreAseg"),
			NumPolizaSeguro:        l.str(a, "NumPolizaSeguro"),
			NumeroGuia:             l.str(a, "NumeroGuia"),
			LugarContrato:          l.str(a, "LugarContrato"),
			CodigoTransportista:    l.str(a, "CodigoTransportista"),
			RFCEmbarcador:          rfc(l.str(a, "RFCEmbarcador")),
			NumRegIdTribEmbarc:     l.str(a, "NumRegIdTribEmbarc"),
			ResidenciaFiscalEmbarc: l.str(a, "ResidenciaFiscalEmbarc"),
			NombreEmbarcador:       l.str(a, "NombreEmbarcador"),
		}
		l.requerido(a, ra, "PermSCT", "NumPermisoSCT", "NumeroGuia", "CodigoTransportista")
	}

	if f := l.unico(n, ns, "TransporteFerroviario", r); f != nil {
		rf := r + "/TransporteFerroviario"
		tf := &TransporteFerroviario{
			TipoDeServicio:  l.str(f, "TipoDeServicio"),
			TipoDeTrafico:   l.str(f, "TipoDeTrafico"),
			NombreAseg:      l.str(f, "NombreAseg"),
			NumPolizaSeguro: l.str(f, "NumPolizaSeguro"),
		}
		l.requerido(f, rf, "TipoDeServicio", "TipoDeTrafico")
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

func (l *lector) cartaPorteMercancia(n *nodo, ns, r string) CartaPorteMercancia {
	m := CartaPorteMercancia{
		BienesTransp:                       l.str(n, "BienesTransp"),
		ClaveSTCC:                          l.str(n, "ClaveSTCC"),
		Descripcion:                        l.str(n, "Descripcion"),
		Cantidad:                           l.dec(n, r, "Cantidad"),
		ClaveUnidad:                        l.str(n, "ClaveUnidad"),
		Unidad:                             l.str(n, "Unidad"),
		Dimensiones:                        l.str(n, "Dimensiones"),
		MaterialPeligroso:                  l.str(n, "MaterialPeligroso"),
		CveMaterialPeligroso:               l.str(n, "CveMaterialPeligroso"),
		Embalaje:                           l.str(n, "Embalaje"),
		DescripEmbalaje:                    l.str(n, "DescripEmbalaje"),
		SectorCOFEPRIS:                     l.str(n, "SectorCOFEPRIS"),
		NombreIngredienteActivo:            l.str(n, "NombreIngredienteActivo"),
		NomQuimico:                         l.str(n, "NomQuimico"),
		DenominacionGenericaProd:           l.str(n, "DenominacionGenericaProd"),
		DenominacionDistintivaProd:         l.str(n, "DenominacionDistintivaProd"),
		Fabricante:                         l.str(n, "Fabricante"),
		FechaCaducidad:                     l.fecha(n, r, "FechaCaducidad", true),
		LoteMedicamento:                    l.str(n, "LoteMedicamento"),
		FormaFarmaceutica:                  l.str(n, "FormaFarmaceutica"),
		CondicionesEspTransp:               l.str(n, "CondicionesEspTransp"),
		RegistroSanitarioFolioAutorizacion: l.str(n, "RegistroSanitarioFolioAutorizacion"),
		PermisoImportacion:                 l.str(n, "PermisoImportacion"),
		FolioImpoVUCEM:                     l.str(n, "FolioImpoVUCEM"),
		NumCAS:                             l.str(n, "NumCAS"),
		RazonSocialEmpImp:                  l.str(n, "RazonSocialEmpImp"),
		NumRegSanPlagCOFEPRIS:              l.str(n, "NumRegSanPlagCOFEPRIS"),
		DatosFabricante:                    l.str(n, "DatosFabricante"),
		DatosFormulador:                    l.str(n, "DatosFormulador"),
		DatosMaquilador:                    l.str(n, "DatosMaquilador"),
		UsoAutorizado:                      l.str(n, "UsoAutorizado"),
		PesoEnKg:                           l.dec(n, r, "PesoEnKg"),
		ValorMercancia:                     l.dec(n, r, "ValorMercancia"),
		Moneda:                             l.str(n, "Moneda"),
		FraccionArancelaria:                l.str(n, "FraccionArancelaria"),
		UUIDComercioExt:                    l.str(n, "UUIDComercioExt"),
		TipoMateria:                        l.str(n, "TipoMateria"),
		DescripcionMateria:                 l.str(n, "DescripcionMateria"),
	}
	l.requerido(n, r, "BienesTransp", "Descripcion", "Cantidad", "ClaveUnidad", "PesoEnKg")

	for i, h := range l.hijos(n, ns, "Pedimentos", r) {
		m.Pedimentos = append(m.Pedimentos, Pedimentos{Pedimento: l.str(h, "Pedimento")})
		l.requerido(h, idx(r, "Pedimentos", i), "Pedimento")
	}
	for i, h := range l.hijos(n, ns, "DocumentacionAduanera", r) {
		m.DocumentacionAduanera = append(m.DocumentacionAduanera, DocumentacionAduanera{
			TipoDocumento:    l.str(h, "TipoDocumento"),
			NumPedimento:     l.str(h, "NumPedimento"),
			IdentDocAduanero: l.str(h, "IdentDocAduanero"),
			RFCImpo:          rfc(l.str(h, "RFCImpo")),
		})
		l.requerido(h, idx(r, "DocumentacionAduanera", i), "TipoDocumento")
	}
	for i, h := range l.hijos(n, ns, "GuiasIdentificacion", r) {
		rg := idx(r, "GuiasIdentificacion", i)
		m.GuiasIdentificacion = append(m.GuiasIdentificacion, GuiasIdentificacion{
			NumeroGuiaIdentificacion:  l.str(h, "NumeroGuiaIdentificacion"),
			DescripGuiaIdentificacion: l.str(h, "DescripGuiaIdentificacion"),
			PesoGuiaIdentificacion:    l.dec(h, rg, "PesoGuiaIdentificacion"),
		})
		l.requerido(h, rg, "NumeroGuiaIdentificacion", "DescripGuiaIdentificacion", "PesoGuiaIdentificacion")
	}
	for i, h := range l.hijos(n, ns, "CantidadTransporta", r) {
		rc := idx(r, "CantidadTransporta", i)
		m.CantidadTransporta = append(m.CantidadTransporta, CantidadTransporta{
			Cantidad:       l.dec(h, rc, "Cantidad"),
			IDOrigen:       l.str(h, "IDOrigen"),
			IDDestino:      l.str(h, "IDDestino"),
			CvesTransporte: l.str(h, "CvesTransporte"),
		})
		l.requerido(h, rc, "Cantidad", "IDOrigen", "IDDestino")
	}
	if d := l.unico(n, ns, "DetalleMercancia", r); d != nil {
		rd := r + "/DetalleMercancia"
		m.DetalleMercancia = &DetalleMercancia{
			UnidadPesoMerc: l.str(d, "UnidadPesoMerc"),
			PesoBruto:      l.dec(d, rd, "PesoBruto"),
			PesoNeto:       l.dec(d, rd, "PesoNeto"),
			PesoTara:       l.dec(d, rd, "PesoTara"),
			NumPiezas:      l.str(d, "NumPiezas"),
		}
		l.requerido(d, rd, "UnidadPesoMerc", "PesoBruto", "PesoNeto", "PesoTara")
	}
	return m
}

func (l *lector) transporteMaritimo(n *nodo, ns, r string, v2 bool) *TransporteMaritimo {
	tm := &TransporteMaritimo{
		PermSCT:                l.str(n, "PermSCT"),
		NumPermisoSCT:          l.str(n, "NumPermisoSCT"),
		NombreAseg:             l.str(n, "NombreAseg"),
		NumPolizaSeguro:        l.str(n, "NumPolizaSeguro"),
		TipoEmbarcacion:        l.str(n, "TipoEmbarcacion"),
		Matricula:              l.str(n, "Matricula"),
		NumeroOMI:              l.str(n, "NumeroOMI"),
		AnioEmbarcacion:        l.str(n, "AnioEmbarcacion"),
		NombreEmbarc:           l.str(n, "NombreEmbarc"),
		NacionalidadEmbarc:     l.str(n, "NacionalidadEmbarc"),
		UnidadesDeArqBruto:     l.dec(n, r, "UnidadesDeArqBruto"),
		TipoCarga:              l.str(n, "TipoCarga"),
		NumCertITC:             l.str(n, "NumCertITC"),
		Eslora:                 l.dec(n, r, "Eslora"),
		Manga:                  l.dec(n, r, "Manga"),
		Calado:                 l.dec(n, r, "Calado"),
		Puntal:                 l.dec(n, r, "Puntal"),
		LineaNaviera:           l.str(n, "LineaNaviera"),
		NombreAgenteNaviero:    l.str(n, "NombreAgenteNaviero"),
		NumAutorizacionNaviero: l.str(n, "NumAutorizacionNaviero"),
		NumViaje:               l.str(n, "NumViaje"),
		NumConocEmbarc:         l.str(n, "NumConocEmbarc"),
		PermisoTempNavegacion:  l.str(n, "PermisoTempNavegacion"),
	}
	l.requerido(n, r, "TipoEmbarcacion", "Matricula", "NumeroOMI", "NacionalidadEmbarc", "UnidadesDeArqBruto",
		"TipoCarga", "NombreAgenteNaviero", "NumAutorizacionNaviero")
	if v2 {
		l.requerido(n, r, "NumCertITC")
	}
	for i, h := range l.hijos(n, ns, "Contenedor", r) {
		rc := idx(r, "Contenedor", i)
		c := ContenedorMaritimo{
			MatriculaContenedor:   l.str(h, "MatriculaContenedor"),
			TipoContenedor:        l.str(h, "TipoContenedor"),
			NumPrecinto:           l.str(h, "NumPrecinto"),
			IdCCPRelacionado:      l.str(h, "IdCCPRelacionado"),
			PlacaVMCCP:            l.str(h, "PlacaVMCCP"),
			FechaCertificacionCCP: l.fecha(h, rc, "FechaCertificacionCCP", false),
			RemolquesCCP:          l.remolquesCCP(h, ns, rc),
		}
		l.requerido(h, rc, "TipoContenedor")
		if v2 {
			l.requerido(h, rc, "MatriculaContenedor")
		}
		tm.Contenedor = append(tm.Contenedor, c)
	}
	tm.RemolquesCCP = l.remolquesCCP(n, ns, r)
	return tm
}

// remolquesCCP lee el contenedor RemolquesCCP de n: cuelga de
// TransporteMaritimo en 3.0 y de cada Contenedor en 3.1.
func (l *lector) remolquesCCP(n *nodo, ns, r string) []RemolqueCCP {
	rs := l.unico(n, ns, "RemolquesCCP", r)
	if rs == nil {
		return nil
	}
	r += "/RemolquesCCP"
	var out []RemolqueCCP
	for i, h := range l.hijos(rs, ns, "RemolqueCCP", r) {
		out = append(out, RemolqueCCP{SubTipoRemCCP: l.str(h, "SubTipoRemCCP"), PlacaCCP: l.str(h, "PlacaCCP")})
		l.requerido(h, idx(r, "RemolqueCCP", i), "SubTipoRemCCP", "PlacaCCP")
	}
	return out
}
