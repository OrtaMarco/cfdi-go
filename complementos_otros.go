package cfdi

import "github.com/OrtaMarco/cfdi-go/decimal"

// Complementos de T10 que no estaban en las tandas anteriores ni son de uso
// común: decretos de vehículos, obras de arte, hidrocarburos, estado de
// cuenta bancario. Mismas convenciones de modelo.go.

// RenovacionySustitucionVehiculos 1.0 (nodo renovacionysustitucionvehiculos):
// decretos de renovación (varios vehículos usados) o sustitución (uno) del
// parque vehicular.
type RenovacionySustitucionVehiculos struct {
	Version                string                            `json:"version,omitempty"`
	TipoDeDecreto          string                            `json:"tipo_de_decreto,omitempty"`
	DecretoRenovVehicular  *RenovacionDecretoRenovVehicular  `json:"decreto_renov_vehicular,omitempty"`
	DecretoSustitVehicular *RenovacionDecretoSustitVehicular `json:"decreto_sustit_vehicular,omitempty"`
}

// RenovacionDecretoRenovVehicular es el nodo decreto:DecretoRenovVehicular del complemento renovacionysustitucionvehiculos.
type RenovacionDecretoRenovVehicular struct {
	VehEnaj                           string                    `json:"veh_enaj,omitempty"`
	VehiculosUsadosEnajenadoPermAlFab []RenovacionVehiculoUsado `json:"vehiculos_usados_enajenado_perm_al_fab,omitempty"`
	VehiculoNuvoSemEnajenadoFabAlPerm *RenovacionVehiculoNuevo  `json:"vehiculo_nuvo_sem_enajenado_fab_al_perm,omitempty"`
}

// RenovacionDecretoSustitVehicular es el nodo decreto:DecretoSustitVehicular del complemento renovacionysustitucionvehiculos.
type RenovacionDecretoSustitVehicular struct {
	VehEnaj                           string                   `json:"veh_enaj,omitempty"`
	VehiculoUsadoEnajenadoPermAlFab   *RenovacionVehiculoUsado `json:"vehiculo_usado_enajenado_perm_al_fab,omitempty"`
	VehiculoNuvoSemEnajenadoFabAlPerm *RenovacionVehiculoNuevo `json:"vehiculo_nuvo_sem_enajenado_fab_al_perm,omitempty"`
}

// RenovacionVehiculoUsado es un vehículo usado enajenado al fabricante: el
// de un decreto de sustitución trae además NumFolAvisoint y exige más.
type RenovacionVehiculoUsado struct {
	PrecioVehUsado *decimal.Decimal `json:"precio_veh_usado,omitempty"`
	TipoVeh        string           `json:"tipo_veh,omitempty"`
	Marca          string           `json:"marca,omitempty"`
	TipooClase     string           `json:"tipoo_clase,omitempty"`
	Anio           string           `json:"anio,omitempty"` // atributo "Año"
	Modelo         string           `json:"modelo,omitempty"`
	NIV            string           `json:"niv,omitempty"`
	NumSerie       string           `json:"num_serie,omitempty"`
	NumPlacas      string           `json:"num_placas,omitempty"`
	NumMotor       string           `json:"num_motor,omitempty"`
	NumFolTarjCir  string           `json:"num_fol_tarj_cir,omitempty"`
	NumFolAvisoint string           `json:"num_fol_avisoint,omitempty"` // solo sustitución
	NumPedIm       string           `json:"num_ped_im,omitempty"`
	Aduana         string           `json:"aduana,omitempty"`
	FechaRegulVeh  Fecha            `json:"fecha_regul_veh,omitzero"`
	Foliofiscal    string           `json:"foliofiscal,omitempty"`
}

// RenovacionVehiculoNuevo es el VehiculoNuvoSemEnajenadoFabAlPerm (así, sin
// e) de los dos decretos.
type RenovacionVehiculoNuevo struct {
	Anio      string `json:"anio,omitempty"` // atributo "Año"
	Modelo    string `json:"modelo,omitempty"`
	NumPlacas string `json:"num_placas,omitempty"`
	RFC       string `json:"rfc,omitempty"`
}

// CertificadoDeDestruccion 1.0 (nodo certificadodedestruccion).
type CertificadoDeDestruccion struct {
	Version             string                              `json:"version,omitempty"`
	Serie               string                              `json:"serie,omitempty"`
	NumFolDesVeh        string                              `json:"num_fol_des_veh,omitempty"`
	VehiculoDestruido   *CertificadoDestruccionVehiculo     `json:"vehiculo_destruido,omitempty"`
	InformacionAduanera *CertificadoDestruccionInfoAduanera `json:"informacion_aduanera,omitempty"`
}

// CertificadoDestruccionVehiculo es el nodo destruccion:VehiculoDestruido del complemento certificadodedestruccion.
type CertificadoDestruccionVehiculo struct {
	Marca         string `json:"marca,omitempty"`
	TipooClase    string `json:"tipoo_clase,omitempty"`
	Anio          string `json:"anio,omitempty"` // atributo "Año"
	Modelo        string `json:"modelo,omitempty"`
	NIV           string `json:"niv,omitempty"`
	NumSerie      string `json:"num_serie,omitempty"`
	NumPlacas     string `json:"num_placas,omitempty"`
	NumMotor      string `json:"num_motor,omitempty"`
	NumFolTarjCir string `json:"num_fol_tarj_cir,omitempty"`
}

// CertificadoDestruccionInfoAduanera es el nodo destruccion:InformacionAduanera del complemento certificadodedestruccion.
type CertificadoDestruccionInfoAduanera struct {
	NumPedImp string `json:"num_ped_imp,omitempty"`
	Fecha     Fecha  `json:"fecha,omitzero"`
	Aduana    string `json:"aduana,omitempty"`
}

// ObrasArteAntiguedades 1.0 (nodo obrasarteantiguedades).
type ObrasArteAntiguedades struct {
	Version                     string           `json:"version,omitempty"`
	TipoBien                    string           `json:"tipo_bien,omitempty"`
	OtrosTipoBien               string           `json:"otros_tipo_bien,omitempty"`
	TituloAdquirido             string           `json:"titulo_adquirido,omitempty"`
	OtrosTituloAdquirido        string           `json:"otros_titulo_adquirido,omitempty"`
	Subtotal                    *decimal.Decimal `json:"subtotal,omitempty"`
	IVA                         *decimal.Decimal `json:"iva,omitempty"`
	FechaAdquisicion            Fecha            `json:"fecha_adquisicion,omitzero"`
	CaracteristicasDeObraoPieza string           `json:"caracteristicas_de_obrao_pieza,omitempty"` // atributo "CaracterísticasDeObraoPieza"
}

// PFintegranteCoordinado 1.0 (pfic): persona física integrante de una
// coordinación de autotransporte.
type PFintegranteCoordinado struct {
	Version        string `json:"version,omitempty"`
	ClaveVehicular string `json:"clave_vehicular,omitempty"`
	Placa          string `json:"placa,omitempty"`
	RFCPF          string `json:"rfcpf,omitempty"`
}

// GastosHidrocarburos 1.0 (GastosHidrocarburos10).
type GastosHidrocarburos struct {
	Version         string                         `json:"version,omitempty"`
	NumeroContrato  string                         `json:"numero_contrato,omitempty"`
	AreaContractual string                         `json:"area_contractual,omitempty"`
	Erogacion       []GastosHidrocarburosErogacion `json:"erogacion,omitempty"`
}

// GastosHidrocarburosErogacion es el nodo gceh:Erogacion del complemento GastosHidrocarburos.
type GastosHidrocarburosErogacion struct {
	TipoErogacion        string                                    `json:"tipo_erogacion,omitempty"`
	MontocuErogacion     *decimal.Decimal                          `json:"montocu_erogacion,omitempty"`
	Porcentaje           *decimal.Decimal                          `json:"porcentaje,omitempty"`
	DocumentoRelacionado []GastosHidrocarburosDocumentoRelacionado `json:"documento_relacionado,omitempty"`
	Actividades          []GastosHidrocarburosActividades          `json:"actividades,omitempty"`
	CentroCostos         []GastosHidrocarburosCentroCostos         `json:"centro_costos,omitempty"`
}

// GastosHidrocarburosDocumentoRelacionado es el nodo gceh:DocumentoRelacionado de una Erogacion del complemento GastosHidrocarburos.
type GastosHidrocarburosDocumentoRelacionado struct {
	OrigenErogacion                string           `json:"origen_erogacion,omitempty"`
	FolioFiscalVinculado           string           `json:"folio_fiscal_vinculado,omitempty"`
	RFCProveedor                   string           `json:"rfc_proveedor,omitempty"`
	MontoTotalIVA                  *decimal.Decimal `json:"monto_total_iva,omitempty"`
	MontoRetencionISR              *decimal.Decimal `json:"monto_retencion_isr,omitempty"`
	MontoRetencionIVA              *decimal.Decimal `json:"monto_retencion_iva,omitempty"`
	MontoRetencionOtrosImpuestos   *decimal.Decimal `json:"monto_retencion_otros_impuestos,omitempty"`
	NumeroPedimentoVinculado       string           `json:"numero_pedimento_vinculado,omitempty"`
	ClavePedimentoVinculado        string           `json:"clave_pedimento_vinculado,omitempty"`
	ClavePagoPedimentoVinculado    string           `json:"clave_pago_pedimento_vinculado,omitempty"`
	MontoIVAPedimento              *decimal.Decimal `json:"monto_iva_pedimento,omitempty"`
	OtrosImpuestosPagadosPedimento *decimal.Decimal `json:"otros_impuestos_pagados_pedimento,omitempty"`
	FechaFolioFiscalVinculado      Fecha            `json:"fecha_folio_fiscal_vinculado,omitzero"`
	Mes                            string           `json:"mes,omitempty"`
	MontoTotalErogaciones          *decimal.Decimal `json:"monto_total_erogaciones,omitempty"`
}

// GastosHidrocarburosActividades es el nodo gceh:Actividades de una Erogacion del complemento GastosHidrocarburos.
type GastosHidrocarburosActividades struct {
	ActividadRelacionada string                              `json:"actividad_relacionada,omitempty"`
	SubActividades       []GastosHidrocarburosSubActividades `json:"sub_actividades,omitempty"`
}

// GastosHidrocarburosSubActividades es el nodo gceh:SubActividades del complemento GastosHidrocarburos.
type GastosHidrocarburosSubActividades struct {
	SubActividadRelacionada string                      `json:"sub_actividad_relacionada,omitempty"`
	Tareas                  []GastosHidrocarburosTareas `json:"tareas,omitempty"`
}

// GastosHidrocarburosTareas es el nodo gceh:Tareas del complemento GastosHidrocarburos.
type GastosHidrocarburosTareas struct {
	TareaRelacionada string `json:"tarea_relacionada,omitempty"`
}

// GastosHidrocarburosCentroCostos es el nodo gceh:CentroCostos de una Erogacion del complemento GastosHidrocarburos.
type GastosHidrocarburosCentroCostos struct {
	Campo       string                           `json:"campo,omitempty"`
	Yacimientos []GastosHidrocarburosYacimientos `json:"yacimientos,omitempty"`
}

// GastosHidrocarburosYacimientos es el nodo gceh:Yacimientos del complemento GastosHidrocarburos.
type GastosHidrocarburosYacimientos struct {
	Yacimiento string                     `json:"yacimiento,omitempty"`
	Pozos      []GastosHidrocarburosPozos `json:"pozos,omitempty"`
}

// GastosHidrocarburosPozos es el nodo gceh:Pozos del complemento GastosHidrocarburos.
type GastosHidrocarburosPozos struct {
	Pozo string `json:"pozo,omitempty"`
}

// IngresosHidrocarburos 1.0 (IngresosHidrocarburos10).
type IngresosHidrocarburos struct {
	Version                        string                                      `json:"version,omitempty"`
	NumeroContrato                 string                                      `json:"numero_contrato,omitempty"`
	ContraprestacionPagadaOperador *decimal.Decimal                            `json:"contraprestacion_pagada_operador,omitempty"`
	Porcentaje                     *decimal.Decimal                            `json:"porcentaje,omitempty"`
	DocumentoRelacionado           []IngresosHidrocarburosDocumentoRelacionado `json:"documento_relacionado,omitempty"`
}

// IngresosHidrocarburosDocumentoRelacionado es el nodo ieeh:DocumentoRelacionado del complemento IngresosHidrocarburos.
type IngresosHidrocarburosDocumentoRelacionado struct {
	FolioFiscalVinculado      string `json:"folio_fiscal_vinculado,omitempty"`
	FechaFolioFiscalVinculado Fecha  `json:"fecha_folio_fiscal_vinculado,omitzero"`
	Mes                       string `json:"mes,omitempty"`
}

// EstadoDeCuentaBancario 1.0 (ecb), de CFDI 3.2. Movimientos no tiene
// atributos pero sí dos clases de hijos: se conserva como objeto con las dos
// listas.
type EstadoDeCuentaBancario struct {
	Version       string          `json:"version,omitempty"`
	NumeroCuenta  string          `json:"numero_cuenta,omitempty"`
	NombreCliente string          `json:"nombre_cliente,omitempty"`
	Periodo       string          `json:"periodo,omitempty"`
	Sucursal      string          `json:"sucursal,omitempty"`
	Movimientos   *ECBMovimientos `json:"movimientos,omitempty"`
}

// ECBMovimientos es el nodo ecb:Movimientos del complemento EstadoDeCuentaBancario.
type ECBMovimientos struct {
	MovimientoECB       []ECBMovimiento       `json:"movimiento_ecb,omitempty"`
	MovimientoECBFiscal []ECBMovimientoFiscal `json:"movimiento_ecb_fiscal,omitempty"`
}

// ECBMovimiento es el nodo ecb:MovimientoECB de los Movimientos del complemento EstadoDeCuentaBancario.
type ECBMovimiento struct {
	Fecha        Fecha            `json:"fecha,omitzero"`
	Referencia   string           `json:"referencia,omitempty"`
	Descripcion  string           `json:"descripcion,omitempty"`
	Importe      *decimal.Decimal `json:"importe,omitempty"`
	Moneda       string           `json:"moneda,omitempty"`
	SaldoInicial *decimal.Decimal `json:"saldo_inicial,omitempty"`
	SaldoAlCorte *decimal.Decimal `json:"saldo_al_corte,omitempty"`
}

// ECBMovimientoFiscal: el XSD escribe RFCenajenante (e minúscula) e Importe
// con mayúscula, a diferencia de MovimientoECB.
type ECBMovimientoFiscal struct {
	Fecha         Fecha            `json:"fecha,omitzero"`
	Referencia    string           `json:"referencia,omitempty"`
	Descripcion   string           `json:"descripcion,omitempty"`
	RFCEnajenante string           `json:"rfc_enajenante,omitempty"` // atributo "RFCenajenante"
	Importe       *decimal.Decimal `json:"importe,omitempty"`
	Moneda        string           `json:"moneda,omitempty"`
	SaldoInicial  *decimal.Decimal `json:"saldo_inicial,omitempty"`
	SaldoAlCorte  *decimal.Decimal `json:"saldo_al_corte,omitempty"`
}

func (l *lector) renovacionySustitucionVehiculos(n *nodo, r string) *RenovacionySustitucionVehiculos {
	ns := n.ns
	x := &RenovacionySustitucionVehiculos{
		Version:       l.versionComplemento(n, r, "1.0"),
		TipoDeDecreto: l.str(n, "TipoDeDecreto"),
	}
	l.requerido(n, r, "TipoDeDecreto")
	if h := l.unico(n, ns, "DecretoRenovVehicular", r); h != nil {
		rh := r + "/DecretoRenovVehicular"
		d := &RenovacionDecretoRenovVehicular{VehEnaj: l.str(h, "VehEnaj")}
		l.requerido(h, rh, "VehEnaj")
		for i, u := range l.hijos(h, ns, "VehiculosUsadosEnajenadoPermAlFab", rh) {
			d.VehiculosUsadosEnajenadoPermAlFab = append(d.VehiculosUsadosEnajenadoPermAlFab,
				*l.renovacionVehiculoUsado(u, idx(rh, "VehiculosUsadosEnajenadoPermAlFab", i), false))
		}
		if len(d.VehiculosUsadosEnajenadoPermAlFab) == 0 {
			l.advertir(AdvNodoFaltante, rh+"/VehiculosUsadosEnajenadoPermAlFab", "DecretoRenovVehicular sin ningún VehiculosUsadosEnajenadoPermAlFab")
		}
		d.VehiculoNuvoSemEnajenadoFabAlPerm = l.renovacionVehiculoNuevo(h, ns, rh)
		x.DecretoRenovVehicular = d
	}
	if h := l.unico(n, ns, "DecretoSustitVehicular", r); h != nil {
		rh := r + "/DecretoSustitVehicular"
		d := &RenovacionDecretoSustitVehicular{VehEnaj: l.str(h, "VehEnaj")}
		l.requerido(h, rh, "VehEnaj")
		if u := l.unico(h, ns, "VehiculoUsadoEnajenadoPermAlFab", rh); u != nil {
			d.VehiculoUsadoEnajenadoPermAlFab = l.renovacionVehiculoUsado(u, rh+"/VehiculoUsadoEnajenadoPermAlFab", true)
		} else {
			l.advertir(AdvNodoFaltante, rh+"/VehiculoUsadoEnajenadoPermAlFab", "DecretoSustitVehicular sin VehiculoUsadoEnajenadoPermAlFab")
		}
		d.VehiculoNuvoSemEnajenadoFabAlPerm = l.renovacionVehiculoNuevo(h, ns, rh)
		x.DecretoSustitVehicular = d
	}
	return x
}

// renovacionVehiculoUsado lee un vehículo usado; el de la sustitución
// (sustit) exige además NumFolAvisoint, NumPedIm, Aduana y FechaRegulVeh.
func (l *lector) renovacionVehiculoUsado(n *nodo, r string, sustit bool) *RenovacionVehiculoUsado {
	x := &RenovacionVehiculoUsado{
		PrecioVehUsado: l.dec(n, r, "PrecioVehUsado"),
		TipoVeh:        l.str(n, "TipoVeh"),
		Marca:          l.str(n, "Marca"),
		TipooClase:     l.str(n, "TipooClase"),
		Anio:           l.str(n, "Año"),
		Modelo:         l.str(n, "Modelo"),
		NIV:            l.str(n, "NIV"),
		NumSerie:       l.str(n, "NumSerie"),
		NumPlacas:      l.str(n, "NumPlacas"),
		NumMotor:       l.str(n, "NumMotor"),
		NumFolTarjCir:  l.str(n, "NumFolTarjCir"),
		NumFolAvisoint: l.str(n, "NumFolAvisoint"),
		NumPedIm:       l.str(n, "NumPedIm"),
		Aduana:         l.str(n, "Aduana"),
		FechaRegulVeh:  l.fecha(n, r, "FechaRegulVeh", true),
		Foliofiscal:    l.str(n, "Foliofiscal"),
	}
	l.requerido(n, r, "PrecioVehUsado", "TipoVeh", "Marca", "TipooClase", "Año", "NumPlacas", "NumFolTarjCir", "Foliofiscal")
	if sustit {
		l.requerido(n, r, "NumFolAvisoint", "NumPedIm", "Aduana", "FechaRegulVeh")
	}
	return x
}

// renovacionVehiculoNuevo lee el VehiculoNuvoSemEnajenadoFabAlPerm de un
// decreto (requerido en los dos).
func (l *lector) renovacionVehiculoNuevo(n *nodo, ns, r string) *RenovacionVehiculoNuevo {
	h := l.unico(n, ns, "VehiculoNuvoSemEnajenadoFabAlPerm", r)
	if h == nil {
		l.advertir(AdvNodoFaltante, r+"/VehiculoNuvoSemEnajenadoFabAlPerm", n.local+" sin VehiculoNuvoSemEnajenadoFabAlPerm")
		return nil
	}
	x := &RenovacionVehiculoNuevo{
		Anio:      l.str(h, "Año"),
		Modelo:    l.str(h, "Modelo"),
		NumPlacas: l.str(h, "NumPlacas"),
		RFC:       rfc(l.str(h, "RFC")),
	}
	l.requerido(h, r+"/VehiculoNuvoSemEnajenadoFabAlPerm", "Año", "NumPlacas")
	return x
}

func (l *lector) certificadoDeDestruccion(n *nodo, r string) *CertificadoDeDestruccion {
	ns := n.ns
	x := &CertificadoDeDestruccion{
		Version:      l.versionComplemento(n, r, "1.0"),
		Serie:        l.str(n, "Serie"),
		NumFolDesVeh: l.str(n, "NumFolDesVeh"),
	}
	l.requerido(n, r, "Serie", "NumFolDesVeh")
	if h := l.unico(n, ns, "VehiculoDestruido", r); h != nil {
		x.VehiculoDestruido = &CertificadoDestruccionVehiculo{
			Marca:         l.str(h, "Marca"),
			TipooClase:    l.str(h, "TipooClase"),
			Anio:          l.str(h, "Año"),
			Modelo:        l.str(h, "Modelo"),
			NIV:           l.str(h, "NIV"),
			NumSerie:      l.str(h, "NumSerie"),
			NumPlacas:     l.str(h, "NumPlacas"),
			NumMotor:      l.str(h, "NumMotor"),
			NumFolTarjCir: l.str(h, "NumFolTarjCir"),
		}
		l.requerido(h, r+"/VehiculoDestruido", "Marca", "TipooClase", "Año", "NumPlacas", "NumFolTarjCir")
	} else {
		l.advertir(AdvNodoFaltante, r+"/VehiculoDestruido", "certificadodedestruccion sin VehiculoDestruido")
	}
	if h := l.unico(n, ns, "InformacionAduanera", r); h != nil {
		ri := r + "/InformacionAduanera"
		x.InformacionAduanera = &CertificadoDestruccionInfoAduanera{
			NumPedImp: l.str(h, "NumPedImp"),
			Fecha:     l.fecha(h, ri, "Fecha", true),
			Aduana:    l.str(h, "Aduana"),
		}
		l.requerido(h, ri, "NumPedImp", "Fecha", "Aduana")
	}
	return x
}

func (l *lector) obrasArteAntiguedades(n *nodo, r string) *ObrasArteAntiguedades {
	x := &ObrasArteAntiguedades{
		Version:                     l.versionComplemento(n, r, "1.0"),
		TipoBien:                    l.str(n, "TipoBien"),
		OtrosTipoBien:               l.str(n, "OtrosTipoBien"),
		TituloAdquirido:             l.str(n, "TituloAdquirido"),
		OtrosTituloAdquirido:        l.str(n, "OtrosTituloAdquirido"),
		Subtotal:                    l.dec(n, r, "Subtotal"),
		IVA:                         l.dec(n, r, "IVA"),
		FechaAdquisicion:            l.fecha(n, r, "FechaAdquisicion", true),
		CaracteristicasDeObraoPieza: l.str(n, "CaracterísticasDeObraoPieza"),
	}
	l.requerido(n, r, "TipoBien", "TituloAdquirido", "FechaAdquisicion", "CaracterísticasDeObraoPieza")
	return x
}

func (l *lector) pfIntegranteCoordinado(n *nodo, r string) *PFintegranteCoordinado {
	x := &PFintegranteCoordinado{
		Version:        l.versionComplemento(n, r, "1.0"),
		ClaveVehicular: l.str(n, "ClaveVehicular"),
		Placa:          l.str(n, "Placa"),
		RFCPF:          rfc(l.str(n, "RFCPF")),
	}
	l.requerido(n, r, "ClaveVehicular", "Placa")
	return x
}

func (l *lector) gastosHidrocarburos(n *nodo, r string) *GastosHidrocarburos {
	ns := n.ns
	x := &GastosHidrocarburos{
		Version:         l.versionComplemento(n, r, "1.0"),
		NumeroContrato:  l.str(n, "NumeroContrato"),
		AreaContractual: l.str(n, "AreaContractual"),
	}
	l.requerido(n, r, "NumeroContrato")
	for i, h := range l.hijos(n, ns, "Erogacion", r) {
		re := idx(r, "Erogacion", i)
		e := GastosHidrocarburosErogacion{
			TipoErogacion:    l.str(h, "TipoErogacion"),
			MontocuErogacion: l.dec(h, re, "MontocuErogacion"),
			Porcentaje:       l.dec(h, re, "Porcentaje"),
		}
		l.requerido(h, re, "TipoErogacion", "MontocuErogacion", "Porcentaje")
		for j, d := range l.hijos(h, ns, "DocumentoRelacionado", re) {
			rd := idx(re, "DocumentoRelacionado", j)
			e.DocumentoRelacionado = append(e.DocumentoRelacionado, GastosHidrocarburosDocumentoRelacionado{
				OrigenErogacion:                l.str(d, "OrigenErogacion"),
				FolioFiscalVinculado:           l.str(d, "FolioFiscalVinculado"),
				RFCProveedor:                   rfc(l.str(d, "RFCProveedor")),
				MontoTotalIVA:                  l.dec(d, rd, "MontoTotalIVA"),
				MontoRetencionISR:              l.dec(d, rd, "MontoRetencionISR"),
				MontoRetencionIVA:              l.dec(d, rd, "MontoRetencionIVA"),
				MontoRetencionOtrosImpuestos:   l.dec(d, rd, "MontoRetencionOtrosImpuestos"),
				NumeroPedimentoVinculado:       l.str(d, "NumeroPedimentoVinculado"),
				ClavePedimentoVinculado:        l.str(d, "ClavePedimentoVinculado"),
				ClavePagoPedimentoVinculado:    l.str(d, "ClavePagoPedimentoVinculado"),
				MontoIVAPedimento:              l.dec(d, rd, "MontoIVAPedimento"),
				OtrosImpuestosPagadosPedimento: l.dec(d, rd, "OtrosImpuestosPagadosPedimento"),
				FechaFolioFiscalVinculado:      l.fecha(d, rd, "FechaFolioFiscalVinculado", true),
				Mes:                            l.str(d, "Mes"),
				MontoTotalErogaciones:          l.dec(d, rd, "MontoTotalErogaciones"),
			})
			l.requerido(d, rd, "OrigenErogacion", "Mes", "MontoTotalErogaciones")
		}
		if len(e.DocumentoRelacionado) == 0 {
			l.advertir(AdvNodoFaltante, re+"/DocumentoRelacionado", "Erogacion sin ningún DocumentoRelacionado")
		}
		for j, a := range l.hijos(h, ns, "Actividades", re) {
			ra := idx(re, "Actividades", j)
			ac := GastosHidrocarburosActividades{ActividadRelacionada: l.str(a, "ActividadRelacionada")}
			for k, s := range l.hijos(a, ns, "SubActividades", ra) {
				rs := idx(ra, "SubActividades", k)
				sa := GastosHidrocarburosSubActividades{SubActividadRelacionada: l.str(s, "SubActividadRelacionada")}
				for _, t := range l.hijos(s, ns, "Tareas", rs) {
					sa.Tareas = append(sa.Tareas, GastosHidrocarburosTareas{TareaRelacionada: l.str(t, "TareaRelacionada")})
				}
				ac.SubActividades = append(ac.SubActividades, sa)
			}
			e.Actividades = append(e.Actividades, ac)
		}
		for j, c := range l.hijos(h, ns, "CentroCostos", re) {
			rc := idx(re, "CentroCostos", j)
			cc := GastosHidrocarburosCentroCostos{Campo: l.str(c, "Campo")}
			for k, y := range l.hijos(c, ns, "Yacimientos", rc) {
				ry := idx(rc, "Yacimientos", k)
				ya := GastosHidrocarburosYacimientos{Yacimiento: l.str(y, "Yacimiento")}
				for _, p := range l.hijos(y, ns, "Pozos", ry) {
					ya.Pozos = append(ya.Pozos, GastosHidrocarburosPozos{Pozo: l.str(p, "Pozo")})
				}
				cc.Yacimientos = append(cc.Yacimientos, ya)
			}
			e.CentroCostos = append(e.CentroCostos, cc)
		}
		x.Erogacion = append(x.Erogacion, e)
	}
	if len(x.Erogacion) == 0 {
		l.advertir(AdvNodoFaltante, r+"/Erogacion", "GastosHidrocarburos sin ninguna Erogacion")
	}
	return x
}

func (l *lector) ingresosHidrocarburos(n *nodo, r string) *IngresosHidrocarburos {
	x := &IngresosHidrocarburos{
		Version:                        l.versionComplemento(n, r, "1.0"),
		NumeroContrato:                 l.str(n, "NumeroContrato"),
		ContraprestacionPagadaOperador: l.dec(n, r, "ContraprestacionPagadaOperador"),
		Porcentaje:                     l.dec(n, r, "Porcentaje"),
	}
	l.requerido(n, r, "NumeroContrato", "ContraprestacionPagadaOperador", "Porcentaje")
	for i, h := range l.hijos(n, n.ns, "DocumentoRelacionado", r) {
		rd := idx(r, "DocumentoRelacionado", i)
		x.DocumentoRelacionado = append(x.DocumentoRelacionado, IngresosHidrocarburosDocumentoRelacionado{
			FolioFiscalVinculado:      l.str(h, "FolioFiscalVinculado"),
			FechaFolioFiscalVinculado: l.fecha(h, rd, "FechaFolioFiscalVinculado", true),
			Mes:                       l.str(h, "Mes"),
		})
		l.requerido(h, rd, "FolioFiscalVinculado", "FechaFolioFiscalVinculado", "Mes")
	}
	if len(x.DocumentoRelacionado) == 0 {
		l.advertir(AdvNodoFaltante, r+"/DocumentoRelacionado", "IngresosHidrocarburos sin ningún DocumentoRelacionado")
	}
	return x
}

func (l *lector) estadoDeCuentaBancario(n *nodo, r string) *EstadoDeCuentaBancario {
	ns := n.ns
	x := &EstadoDeCuentaBancario{
		Version:       l.versionComplemento(n, r, "1.0"),
		NumeroCuenta:  l.str(n, "numeroCuenta"),
		NombreCliente: l.str(n, "nombreCliente"),
		Periodo:       l.str(n, "periodo"),
		Sucursal:      l.str(n, "sucursal"),
	}
	l.requerido(n, r, "numeroCuenta", "nombreCliente", "periodo")
	m := l.unico(n, ns, "Movimientos", r)
	if m == nil {
		l.advertir(AdvNodoFaltante, r+"/Movimientos", "EstadoDeCuentaBancario sin Movimientos")
		return x
	}
	rm := r + "/Movimientos"
	x.Movimientos = &ECBMovimientos{}
	for i, h := range l.hijos(m, ns, "MovimientoECB", rm) {
		ri := idx(rm, "MovimientoECB", i)
		x.Movimientos.MovimientoECB = append(x.Movimientos.MovimientoECB, ECBMovimiento{
			Fecha:        l.fecha(h, ri, "fecha", false),
			Referencia:   l.str(h, "referencia"),
			Descripcion:  l.str(h, "descripcion"),
			Importe:      l.dec(h, ri, "importe"),
			Moneda:       l.str(h, "moneda"),
			SaldoInicial: l.dec(h, ri, "saldoInicial"),
			SaldoAlCorte: l.dec(h, ri, "saldoAlCorte"),
		})
		l.requerido(h, ri, "fecha", "descripcion", "importe")
	}
	for i, h := range l.hijos(m, ns, "MovimientoECBFiscal", rm) {
		ri := idx(rm, "MovimientoECBFiscal", i)
		x.Movimientos.MovimientoECBFiscal = append(x.Movimientos.MovimientoECBFiscal, ECBMovimientoFiscal{
			Fecha:         l.fecha(h, ri, "fecha", false),
			Referencia:    l.str(h, "referencia"),
			Descripcion:   l.str(h, "descripcion"),
			RFCEnajenante: rfc(l.str(h, "RFCenajenante")),
			Importe:       l.dec(h, ri, "Importe"),
			Moneda:        l.str(h, "moneda"),
			SaldoInicial:  l.dec(h, ri, "saldoInicial"),
			SaldoAlCorte:  l.dec(h, ri, "saldoAlCorte"),
		})
		l.requerido(h, ri, "fecha", "descripcion", "RFCenajenante", "Importe")
	}
	return x
}

// estadoDeCuentaCombustible10 lee el EstadoDeCuentaCombustible sin versión
// (namespace .../ecc, de CFDI 3.2) en el struct de 1.1 y 1.2. Su XSD no
// trae atributo de versión: Version es "1.0" porque la dicta el namespace.
// Los atributos van en minúsculas y el traslado trae tasa (porcentaje, con
// dos decimales) en vez de TasaOCuota: va en Impuesto.Tasa.
func (l *lector) estadoDeCuentaCombustible10(n *nodo, r string) *EstadoDeCuentaCombustible {
	ns := n.ns
	e := &EstadoDeCuentaCombustible{
		Version:        "1.0",
		TipoOperacion:  l.str(n, "tipoOperacion"),
		NumeroDeCuenta: l.str(n, "numeroDeCuenta"),
		SubTotal:       l.dec(n, r, "subTotal"),
		Total:          l.dec(n, r, "total"),
	}
	l.requerido(n, r, "tipoOperacion", "numeroDeCuenta", "total")
	cs := l.unico(n, ns, "Conceptos", r)
	if cs == nil {
		l.advertir(AdvNodoFaltante, r+"/Conceptos", "EstadoDeCuentaCombustible sin Conceptos")
		return e
	}
	rc := r + "/Conceptos"
	for i, h := range l.hijos(cs, ns, "ConceptoEstadoDeCuentaCombustible", rc) {
		ri := idx(rc, "ConceptoEstadoDeCuentaCombustible", i)
		cp := ConceptoEstadoDeCuentaCombustible{
			Identificador:     l.str(h, "identificador"),
			Fecha:             l.fecha(h, ri, "fecha", false),
			Rfc:               rfc(l.str(h, "rfc")),
			ClaveEstacion:     l.str(h, "claveEstacion"),
			Cantidad:          l.dec(h, ri, "cantidad"),
			NombreCombustible: l.str(h, "nombreCombustible"),
			FolioOperacion:    l.str(h, "folioOperacion"),
			ValorUnitario:     l.dec(h, ri, "valorUnitario"),
			Importe:           l.dec(h, ri, "importe"),
		}
		l.requerido(h, ri, "identificador", "fecha", "rfc", "claveEstacion", "cantidad", "nombreCombustible",
			"folioOperacion", "valorUnitario", "importe")
		if ts := l.unico(h, ns, "Traslados", ri); ts != nil {
			rt := ri + "/Traslados"
			for j, t := range l.hijos(ts, ns, "Traslado", rt) {
				rj := idx(rt, "Traslado", j)
				cp.Traslados = append(cp.Traslados, Impuesto{
					Impuesto: l.str(t, "impuesto"),
					Tasa:     l.dec(t, rj, "tasa"),
					Importe:  l.dec(t, rj, "importe"),
				})
				l.requerido(t, rj, "impuesto", "tasa", "importe")
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
