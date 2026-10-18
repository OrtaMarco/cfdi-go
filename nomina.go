package cfdi

func (l *lector) nomina(n *nodo, r string) Nomina {
	ns := n.ns
	nm := Nomina{
		Version:           l.versionComplemento(n, r, "1.2"),
		TipoNomina:        l.str(n, "TipoNomina"),
		FechaPago:         l.fecha(n, r, "FechaPago", true),
		FechaInicialPago:  l.fecha(n, r, "FechaInicialPago", true),
		FechaFinalPago:    l.fecha(n, r, "FechaFinalPago", true),
		NumDiasPagados:    l.dec(n, r, "NumDiasPagados"),
		TotalPercepciones: l.dec(n, r, "TotalPercepciones"),
		TotalDeducciones:  l.dec(n, r, "TotalDeducciones"),
		TotalOtrosPagos:   l.dec(n, r, "TotalOtrosPagos"),
	}
	l.requerido(n, r, "TipoNomina", "FechaPago", "FechaInicialPago", "FechaFinalPago", "NumDiasPagados")

	if e := l.unico(n, ns, "Emisor", r); e != nil {
		re := r + "/Emisor"
		ne := &NominaEmisor{
			Curp:             l.str(e, "Curp"),
			RegistroPatronal: l.str(e, "RegistroPatronal"),
			RfcPatronOrigen:  rfc(l.str(e, "RfcPatronOrigen")),
		}
		if s := l.unico(e, ns, "EntidadSNCF", re); s != nil {
			ne.EntidadSNCF = &EntidadSNCF{
				OrigenRecurso:      l.str(s, "OrigenRecurso"),
				MontoRecursoPropio: l.dec(s, re+"/EntidadSNCF", "MontoRecursoPropio"),
			}
		}
		nm.Emisor = ne
	}

	if e := l.unico(n, ns, "Receptor", r); e != nil {
		rr := r + "/Receptor"
		nr := &NominaReceptor{
			Curp:                   l.str(e, "Curp"),
			NumSeguridadSocial:     l.str(e, "NumSeguridadSocial"),
			FechaInicioRelLaboral:  l.fecha(e, rr, "FechaInicioRelLaboral", true),
			Antiguedad:             l.str(e, "Antigüedad"),
			TipoContrato:           l.str(e, "TipoContrato"),
			Sindicalizado:          l.str(e, "Sindicalizado"),
			TipoJornada:            l.str(e, "TipoJornada"),
			TipoRegimen:            l.str(e, "TipoRegimen"),
			NumEmpleado:            l.str(e, "NumEmpleado"),
			Departamento:           l.str(e, "Departamento"),
			Puesto:                 l.str(e, "Puesto"),
			RiesgoPuesto:           l.str(e, "RiesgoPuesto"),
			PeriodicidadPago:       l.str(e, "PeriodicidadPago"),
			Banco:                  l.str(e, "Banco"),
			CuentaBancaria:         l.str(e, "CuentaBancaria"),
			SalarioBaseCotApor:     l.dec(e, rr, "SalarioBaseCotApor"),
			SalarioDiarioIntegrado: l.dec(e, rr, "SalarioDiarioIntegrado"),
			ClaveEntFed:            l.str(e, "ClaveEntFed"),
		}
		l.requerido(e, rr, "Curp", "TipoContrato", "TipoRegimen", "NumEmpleado", "PeriodicidadPago", "ClaveEntFed")
		for i, s := range l.hijos(e, ns, "SubContratacion", rr) {
			nr.SubContratacion = append(nr.SubContratacion, SubContratacion{
				RfcLabora:        rfc(l.str(s, "RfcLabora")),
				PorcentajeTiempo: l.dec(s, idx(rr, "SubContratacion", i), "PorcentajeTiempo"),
			})
		}
		nm.Receptor = nr
	} else {
		l.advertir(AdvNodoFaltante, r+"/Receptor", "Nomina sin Receptor")
	}

	if p := l.unico(n, ns, "Percepciones", r); p != nil {
		nm.Percepciones = l.percepciones(p, ns, r+"/Percepciones")
	}

	if d := l.unico(n, ns, "Deducciones", r); d != nil {
		rd := r + "/Deducciones"
		dd := &Deducciones{
			TotalOtrasDeducciones:   l.dec(d, rd, "TotalOtrasDeducciones"),
			TotalImpuestosRetenidos: l.dec(d, rd, "TotalImpuestosRetenidos"),
		}
		for i, h := range l.hijos(d, ns, "Deduccion", rd) {
			ri := idx(rd, "Deduccion", i)
			dd.Deduccion = append(dd.Deduccion, Deduccion{
				TipoDeduccion: l.str(h, "TipoDeduccion"),
				Clave:         l.str(h, "Clave"),
				Concepto:      l.str(h, "Concepto"),
				Importe:       l.dec(h, ri, "Importe"),
			})
			l.requerido(h, ri, "TipoDeduccion", "Importe")
		}
		nm.Deducciones = dd
	}

	if o := l.unico(n, ns, "OtrosPagos", r); o != nil {
		ro := r + "/OtrosPagos"
		for i, h := range l.hijos(o, ns, "OtroPago", ro) {
			ri := idx(ro, "OtroPago", i)
			op := OtroPago{
				TipoOtroPago: l.str(h, "TipoOtroPago"),
				Clave:        l.str(h, "Clave"),
				Concepto:     l.str(h, "Concepto"),
				Importe:      l.dec(h, ri, "Importe"),
			}
			if s := l.unico(h, ns, "SubsidioAlEmpleo", ri); s != nil {
				op.SubsidioAlEmpleo = &SubsidioAlEmpleo{SubsidioCausado: l.dec(s, ri+"/SubsidioAlEmpleo", "SubsidioCausado")}
			}
			if s := l.unico(h, ns, "CompensacionSaldosAFavor", ri); s != nil {
				rs := ri + "/CompensacionSaldosAFavor"
				op.CompensacionSaldosAFavor = &CompensacionSaldosAFavor{
					SaldoAFavor:     l.dec(s, rs, "SaldoAFavor"),
					Anio:            l.str(s, "Año"),
					RemanenteSalFav: l.dec(s, rs, "RemanenteSalFav"),
				}
			}
			nm.OtrosPagos = append(nm.OtrosPagos, op)
		}
	}

	if in := l.unico(n, ns, "Incapacidades", r); in != nil {
		ri := r + "/Incapacidades"
		for i, h := range l.hijos(in, ns, "Incapacidad", ri) {
			nm.Incapacidades = append(nm.Incapacidades, Incapacidad{
				DiasIncapacidad:  l.str(h, "DiasIncapacidad"),
				TipoIncapacidad:  l.str(h, "TipoIncapacidad"),
				ImporteMonetario: l.dec(h, idx(ri, "Incapacidad", i), "ImporteMonetario"),
			})
		}
	}
	return nm
}

// nomina11 lee Nómina 1.1 (namespace .../nomina, de CFDI 3.2) en el mismo
// struct que 1.2: lo que 1.1 pone en el nodo Nomina (los datos del
// trabajador, HorasExtras) va en sus campos de 1.1, no en Receptor.
func (l *lector) nomina11(n *nodo, r string) Nomina {
	ns := n.ns
	nm := Nomina{
		Version:                l.versionComplemento(n, r, "1.1"),
		RegistroPatronal:       l.str(n, "RegistroPatronal"),
		NumEmpleado:            l.str(n, "NumEmpleado"),
		CURP:                   l.str(n, "CURP"),
		TipoRegimen:            l.str(n, "TipoRegimen"),
		NumSeguridadSocial:     l.str(n, "NumSeguridadSocial"),
		FechaPago:              l.fecha(n, r, "FechaPago", true),
		FechaInicialPago:       l.fecha(n, r, "FechaInicialPago", true),
		FechaFinalPago:         l.fecha(n, r, "FechaFinalPago", true),
		NumDiasPagados:         l.dec(n, r, "NumDiasPagados"),
		Departamento:           l.str(n, "Departamento"),
		CLABE:                  l.str(n, "CLABE"),
		Banco:                  l.str(n, "Banco"),
		FechaInicioRelLaboral:  l.fecha(n, r, "FechaInicioRelLaboral", true),
		Antiguedad:             l.str(n, "Antiguedad"),
		Puesto:                 l.str(n, "Puesto"),
		TipoContrato:           l.str(n, "TipoContrato"),
		TipoJornada:            l.str(n, "TipoJornada"),
		PeriodicidadPago:       l.str(n, "PeriodicidadPago"),
		SalarioBaseCotApor:     l.dec(n, r, "SalarioBaseCotApor"),
		RiesgoPuesto:           l.str(n, "RiesgoPuesto"),
		SalarioDiarioIntegrado: l.dec(n, r, "SalarioDiarioIntegrado"),
	}
	l.requerido(n, r, "NumEmpleado", "CURP", "TipoRegimen", "FechaPago", "FechaInicialPago", "FechaFinalPago",
		"NumDiasPagados", "PeriodicidadPago")

	if p := l.unico(n, ns, "Percepciones", r); p != nil {
		rp := r + "/Percepciones"
		nm.Percepciones = l.percepciones(p, ns, rp)
		l.requerido(p, rp, "TotalGravado", "TotalExento")
		for i, h := range l.hijos(p, ns, "Percepcion", rp) {
			l.requerido(h, idx(rp, "Percepcion", i), "Clave", "Concepto")
		}
	}

	if d := l.unico(n, ns, "Deducciones", r); d != nil {
		rd := r + "/Deducciones"
		dd := &Deducciones{
			TotalGravado: l.dec(d, rd, "TotalGravado"),
			TotalExento:  l.dec(d, rd, "TotalExento"),
		}
		l.requerido(d, rd, "TotalGravado", "TotalExento")
		for i, h := range l.hijos(d, ns, "Deduccion", rd) {
			ri := idx(rd, "Deduccion", i)
			dd.Deduccion = append(dd.Deduccion, Deduccion{
				TipoDeduccion:  l.str(h, "TipoDeduccion"),
				Clave:          l.str(h, "Clave"),
				Concepto:       l.str(h, "Concepto"),
				ImporteGravado: l.dec(h, ri, "ImporteGravado"),
				ImporteExento:  l.dec(h, ri, "ImporteExento"),
			})
			l.requerido(h, ri, "TipoDeduccion", "Clave", "Concepto", "ImporteGravado", "ImporteExento")
		}
		nm.Deducciones = dd
	}

	if in := l.unico(n, ns, "Incapacidades", r); in != nil {
		ri := r + "/Incapacidades"
		for i, h := range l.hijos(in, ns, "Incapacidad", ri) {
			rh := idx(ri, "Incapacidad", i)
			nm.Incapacidades = append(nm.Incapacidades, Incapacidad{
				DiasIncapacidad: l.str(h, "DiasIncapacidad"),
				TipoIncapacidad: l.str(h, "TipoIncapacidad"),
				Descuento:       l.dec(h, rh, "Descuento"),
			})
			l.requerido(h, rh, "DiasIncapacidad", "TipoIncapacidad", "Descuento")
		}
	}

	if hs := l.unico(n, ns, "HorasExtras", r); hs != nil {
		rh := r + "/HorasExtras"
		for i, h := range l.hijos(hs, ns, "HorasExtra", rh) {
			ri := idx(rh, "HorasExtra", i)
			nm.HorasExtras = append(nm.HorasExtras, HorasExtra{
				Dias:          l.str(h, "Dias"),
				TipoHoras:     l.str(h, "TipoHoras"),
				HorasExtra:    l.str(h, "HorasExtra"),
				ImportePagado: l.dec(h, ri, "ImportePagado"),
			})
			l.requerido(h, ri, "Dias", "TipoHoras", "HorasExtra", "ImportePagado")
		}
	}
	return nm
}

func (l *lector) percepciones(p *nodo, ns, r string) *Percepciones {
	pp := &Percepciones{
		TotalSueldos:                 l.dec(p, r, "TotalSueldos"),
		TotalSeparacionIndemnizacion: l.dec(p, r, "TotalSeparacionIndemnizacion"),
		TotalJubilacionPensionRetiro: l.dec(p, r, "TotalJubilacionPensionRetiro"),
		TotalGravado:                 l.dec(p, r, "TotalGravado"),
		TotalExento:                  l.dec(p, r, "TotalExento"),
	}
	for i, h := range l.hijos(p, ns, "Percepcion", r) {
		ri := idx(r, "Percepcion", i)
		pc := Percepcion{
			TipoPercepcion: l.str(h, "TipoPercepcion"),
			Clave:          l.str(h, "Clave"),
			Concepto:       l.str(h, "Concepto"),
			ImporteGravado: l.dec(h, ri, "ImporteGravado"),
			ImporteExento:  l.dec(h, ri, "ImporteExento"),
		}
		l.requerido(h, ri, "TipoPercepcion", "ImporteGravado", "ImporteExento")
		if a := l.unico(h, ns, "AccionesOTitulos", ri); a != nil {
			ra := ri + "/AccionesOTitulos"
			pc.AccionesOTitulos = &AccionesOTitulos{
				ValorMercado:      l.dec(a, ra, "ValorMercado"),
				PrecioAlOtorgarse: l.dec(a, ra, "PrecioAlOtorgarse"),
			}
		}
		for j, hx := range l.hijos(h, ns, "HorasExtra", ri) {
			pc.HorasExtra = append(pc.HorasExtra, HorasExtra{
				Dias:          l.str(hx, "Dias"),
				TipoHoras:     l.str(hx, "TipoHoras"),
				HorasExtra:    l.str(hx, "HorasExtra"),
				ImportePagado: l.dec(hx, idx(ri, "HorasExtra", j), "ImportePagado"),
			})
		}
		pp.Percepcion = append(pp.Percepcion, pc)
	}
	if j := l.unico(p, ns, "JubilacionPensionRetiro", r); j != nil {
		rj := r + "/JubilacionPensionRetiro"
		pp.JubilacionPensionRetiro = &JubilacionPensionRetiro{
			TotalUnaExhibicion:  l.dec(j, rj, "TotalUnaExhibicion"),
			TotalParcialidad:    l.dec(j, rj, "TotalParcialidad"),
			MontoDiario:         l.dec(j, rj, "MontoDiario"),
			IngresoAcumulable:   l.dec(j, rj, "IngresoAcumulable"),
			IngresoNoAcumulable: l.dec(j, rj, "IngresoNoAcumulable"),
		}
	}
	if s := l.unico(p, ns, "SeparacionIndemnizacion", r); s != nil {
		rs := r + "/SeparacionIndemnizacion"
		pp.SeparacionIndemnizacion = &SeparacionIndemnizacion{
			TotalPagado:         l.dec(s, rs, "TotalPagado"),
			NumAniosServicio:    l.str(s, "NumAñosServicio"),
			UltimoSueldoMensOrd: l.dec(s, rs, "UltimoSueldoMensOrd"),
			IngresoAcumulable:   l.dec(s, rs, "IngresoAcumulable"),
			IngresoNoAcumulable: l.dec(s, rs, "IngresoNoAcumulable"),
		}
	}
	return pp
}
