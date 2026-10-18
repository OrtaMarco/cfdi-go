package cfdi

import "github.com/OrtaMarco/cfdi-go/decimal"

// Convenciones del modelo:
//
//   - Los campos Go llevan el nombre del atributo o nodo del SAT; en JSON van
//     en snake_case ("TipoDeComprobante" → "tipo_de_comprobante"). Las letras
//     ñ/ü se transliteran: "Año" → Anio/"anio", "Antigüedad" → Antiguedad.
//   - Un nodo contenedor sin atributos (Conceptos, Traslados, OtrosPagos…) se
//     aplana: el campo es la lista de sus hijos con el nombre del contenedor.
//     Un nodo con atributos propios (Percepciones, Deducciones, Pagos) se
//     conserva como objeto y su lista lleva el nombre del hijo (Percepcion,
//     Deduccion, Pago).
//   - Montos y cantidades: *decimal.Decimal. nil = el atributo no venía.
//   - Enteros del SAT (NumParcialidad, Dias, Año…) se guardan como texto para
//     no perder lo que venía si no es un entero válido.
//   - Claves de catálogo (FormaPago, UsoCFDI…) se guardan tal cual.

// Comprobante es un CFDI 3.3 o 4.0 leído y normalizado.
type Comprobante struct {
	Version           string           `json:"version,omitempty"`
	Serie             string           `json:"serie,omitempty"`
	Folio             string           `json:"folio,omitempty"`
	Fecha             Fecha            `json:"fecha,omitzero"`
	Sello             string           `json:"sello,omitempty"`
	FormaPago         string           `json:"forma_pago,omitempty"`
	NoCertificado     string           `json:"no_certificado,omitempty"`
	Certificado       string           `json:"certificado,omitempty"`
	CondicionesDePago string           `json:"condiciones_de_pago,omitempty"`
	SubTotal          *decimal.Decimal `json:"sub_total,omitempty"`
	Descuento         *decimal.Decimal `json:"descuento,omitempty"`
	Moneda            string           `json:"moneda,omitempty"`
	TipoCambio        *decimal.Decimal `json:"tipo_cambio,omitempty"`
	Total             *decimal.Decimal `json:"total,omitempty"`
	TipoDeComprobante string           `json:"tipo_de_comprobante,omitempty"`
	Exportacion       string           `json:"exportacion,omitempty"`
	MetodoPago        string           `json:"metodo_pago,omitempty"`
	LugarExpedicion   string           `json:"lugar_expedicion,omitempty"`
	Confirmacion      string           `json:"confirmacion,omitempty"`

	InformacionGlobal *InformacionGlobal `json:"informacion_global,omitempty"`
	// CfdiRelacionados: en 3.3 hay a lo más uno; en 4.0 puede haber varios.
	// Siempre se normaliza a lista.
	CfdiRelacionados []CfdiRelacionados `json:"cfdi_relacionados,omitempty"`
	Emisor           *Emisor            `json:"emisor,omitempty"`
	Receptor         *Receptor          `json:"receptor,omitempty"`
	Conceptos        []Concepto         `json:"conceptos,omitempty"`
	Impuestos        *Impuestos         `json:"impuestos,omitempty"`

	// Complementos soportados.
	TimbreFiscalDigital       *TimbreFiscalDigital       `json:"timbre_fiscal_digital,omitempty"`
	Pagos                     *Pagos                     `json:"pagos,omitempty"`
	Nomina                    []Nomina                   `json:"nomina,omitempty"`
	ImpuestosLocales          *ImpuestosLocales          `json:"impuestos_locales,omitempty"`
	CartaPorte                *CartaPorte                `json:"carta_porte,omitempty"`
	ComercioExterior          *ComercioExterior          `json:"comercio_exterior,omitempty"`
	Divisas                   *Divisas                   `json:"divisas,omitempty"`
	CFDIRegistroFiscal        *CFDIRegistroFiscal        `json:"cfdi_registro_fiscal,omitempty"`
	LeyendasFiscales          *LeyendasFiscales          `json:"leyendas_fiscales,omitempty"`
	EstadoDeCuentaCombustible *EstadoDeCuentaCombustible `json:"estado_de_cuenta_combustible,omitempty"`
	// HidroYPetro es complemento de concepto (Concepto.HidroYPetro); aquí van
	// los que vienen fuera de lugar, en el Complemento del comprobante.
	HidroYPetro                     []HidroYPetro                    `json:"hidro_y_petro,omitempty"`
	Aerolineas                      *Aerolineas                      `json:"aerolineas,omitempty"`
	Donatarias                      *Donatarias                      `json:"donatarias,omitempty"`
	INE                             *INE                             `json:"ine,omitempty"`
	ValesDeDespensa                 *ValesDeDespensa                 `json:"vales_de_despensa,omitempty"`
	ConsumoDeCombustibles           *ConsumoDeCombustibles           `json:"consumo_de_combustibles,omitempty"`
	PagoEnEspecie                   *PagoEnEspecie                   `json:"pago_en_especie,omitempty"`
	NotariosPublicos                *NotariosPublicos                `json:"notarios_publicos,omitempty"`
	ParcialesConstruccion           *ParcialesConstruccion           `json:"parcialesconstruccion,omitempty"`
	VehiculoUsado                   *VehiculoUsado                   `json:"vehiculo_usado,omitempty"`
	TuristaPasajeroExtranjero       *TuristaPasajeroExtranjero       `json:"turista_pasajero_extranjero,omitempty"`
	ComplementoSPEI                 []SPEITercero                    `json:"complemento_spei,omitempty"` // Complemento_SPEI/SPEI_Tercero
	RenovacionySustitucionVehiculos *RenovacionySustitucionVehiculos `json:"renovacionysustitucionvehiculos,omitempty"`
	CertificadoDeDestruccion        *CertificadoDeDestruccion        `json:"certificadodedestruccion,omitempty"`
	ObrasArteAntiguedades           *ObrasArteAntiguedades           `json:"obrasarteantiguedades,omitempty"`
	PFintegranteCoordinado          *PFintegranteCoordinado          `json:"pf_integrante_coordinado,omitempty"`
	GastosHidrocarburos             *GastosHidrocarburos             `json:"gastos_hidrocarburos,omitempty"`
	IngresosHidrocarburos           *IngresosHidrocarburos           `json:"ingresos_hidrocarburos,omitempty"`
	EstadoDeCuentaBancario          *EstadoDeCuentaBancario          `json:"estado_de_cuenta_bancario,omitempty"`
	// ComplementosNoSoportados lista los nodos de Complemento (y de
	// ComplementoConcepto) que no se leen, como "{namespace}NombreLocal",
	// sin repetir.
	ComplementosNoSoportados []string `json:"complementos_no_soportados,omitempty"`

	// Advertencias de la lectura. En JSON van fuera de "cfdi", en el sobre.
	Advertencias []Advertencia `json:"-"`
}

// InformacionGlobal (4.0) para CFDI a público en general.
type InformacionGlobal struct {
	Periodicidad string `json:"periodicidad,omitempty"`
	Meses        string `json:"meses,omitempty"`
	Anio         string `json:"anio,omitempty"` // atributo "Año"
}

// CfdiRelacionados agrupa los UUID relacionados con un mismo TipoRelacion.
type CfdiRelacionados struct {
	TipoRelacion    string            `json:"tipo_relacion,omitempty"`
	CfdiRelacionado []CfdiRelacionado `json:"cfdi_relacionado,omitempty"`
}

// CfdiRelacionado es un UUID relacionado, normalizado como el del timbre.
type CfdiRelacionado struct {
	UUID         string `json:"uuid,omitempty"`
	UUIDOriginal string `json:"uuid_original,omitempty"` // solo si difiere de UUID
}

// Emisor es el nodo cfdi:Emisor del comprobante.
type Emisor struct {
	Rfc              string `json:"rfc,omitempty"`
	Nombre           string `json:"nombre,omitempty"`
	RegimenFiscal    string `json:"regimen_fiscal,omitempty"`
	FacAtrAdquirente string `json:"fac_atr_adquirente,omitempty"`
}

// Receptor es el nodo cfdi:Receptor del comprobante.
type Receptor struct {
	Rfc                     string `json:"rfc,omitempty"`
	Nombre                  string `json:"nombre,omitempty"`
	DomicilioFiscalReceptor string `json:"domicilio_fiscal_receptor,omitempty"`
	ResidenciaFiscal        string `json:"residencia_fiscal,omitempty"`
	NumRegIdTrib            string `json:"num_reg_id_trib,omitempty"`
	RegimenFiscalReceptor   string `json:"regimen_fiscal_receptor,omitempty"`
	UsoCFDI                 string `json:"uso_cfdi,omitempty"`
}

// Concepto es el nodo cfdi:Concepto del comprobante, con sus impuestos y complementos de concepto.
type Concepto struct {
	ClaveProdServ    string           `json:"clave_prod_serv,omitempty"`
	NoIdentificacion string           `json:"no_identificacion,omitempty"`
	Cantidad         *decimal.Decimal `json:"cantidad,omitempty"`
	ClaveUnidad      string           `json:"clave_unidad,omitempty"`
	Unidad           string           `json:"unidad,omitempty"`
	Descripcion      string           `json:"descripcion,omitempty"`
	ValorUnitario    *decimal.Decimal `json:"valor_unitario,omitempty"`
	Importe          *decimal.Decimal `json:"importe,omitempty"`
	Descuento        *decimal.Decimal `json:"descuento,omitempty"`
	ObjetoImp        string           `json:"objeto_imp,omitempty"`

	Impuestos           *ImpuestosConcepto    `json:"impuestos,omitempty"`
	ACuentaTerceros     *ACuentaTerceros      `json:"a_cuenta_terceros,omitempty"`
	InformacionAduanera []InformacionAduanera `json:"informacion_aduanera,omitempty"`
	// CuentaPredial: 3.3 admite una, 4.0 varias. Siempre lista.
	CuentaPredial []CuentaPredial `json:"cuenta_predial,omitempty"`
	// Complementos de concepto soportados.
	PorCuentadeTerceros *PorCuentadeTerceros `json:"por_cuentade_terceros,omitempty"`
	VentaVehiculos      *VentaVehiculos      `json:"venta_vehiculos,omitempty"`
	HidroYPetro         *HidroYPetro         `json:"hidro_y_petro,omitempty"`
	InstEducativas      *InstEducativas      `json:"inst_educativas,omitempty"`
	AcreditamientoIEPS  *AcreditamientoIEPS  `json:"acreditamiento_ieps,omitempty"`
	// ParcialesConstruccion fuera de lugar: el XSD lo pone en el Complemento
	// del comprobante (Comprobante.ParcialesConstruccion).
	ParcialesConstruccion *ParcialesConstruccion `json:"parcialesconstruccion,omitempty"`
	// ComplementoConcepto lista los complementos del concepto que no se
	// leen, como "{namespace}NombreLocal".
	ComplementoConcepto []string `json:"complemento_concepto,omitempty"`
	Parte               []Parte  `json:"parte,omitempty"`
}

// ImpuestosConcepto son los traslados y retenciones de un concepto.
type ImpuestosConcepto struct {
	Traslados   []Impuesto `json:"traslados,omitempty"`
	Retenciones []Impuesto `json:"retenciones,omitempty"`
}

// Impuesto es un Traslado o una Retencion, del concepto o del comprobante.
// Con TipoFactor "Exento" no hay TasaOCuota ni Importe (nil).
type Impuesto struct {
	Base       *decimal.Decimal `json:"base,omitempty"`
	Impuesto   string           `json:"impuesto,omitempty"`
	TipoFactor string           `json:"tipo_factor,omitempty"`
	TasaOCuota *decimal.Decimal `json:"tasa_o_cuota,omitempty"`
	// Tasa en porcentaje (16.00): solo en los traslados del
	// EstadoDeCuentaCombustible sin versión (ecc). El tasa de los traslados
	// de CFDI 3.2 no se lee.
	Tasa    *decimal.Decimal `json:"tasa,omitempty"`
	Importe *decimal.Decimal `json:"importe,omitempty"`
}

// ACuentaTerceros es el nodo cfdi:ACuentaTerceros de un concepto (CFDI 4.0).
type ACuentaTerceros struct {
	RfcACuentaTerceros             string `json:"rfc_a_cuenta_terceros,omitempty"`
	NombreACuentaTerceros          string `json:"nombre_a_cuenta_terceros,omitempty"`
	RegimenFiscalACuentaTerceros   string `json:"regimen_fiscal_a_cuenta_terceros,omitempty"`
	DomicilioFiscalACuentaTerceros string `json:"domicilio_fiscal_a_cuenta_terceros,omitempty"`
}

// InformacionAduanera es el nodo cfdi:InformacionAduanera de un concepto o de una parte.
type InformacionAduanera struct {
	NumeroPedimento string `json:"numero_pedimento,omitempty"`
}

// CuentaPredial es el nodo cfdi:CuentaPredial de un concepto.
type CuentaPredial struct {
	Numero string `json:"numero,omitempty"`
}

// Parte es el nodo cfdi:Parte de un concepto.
type Parte struct {
	ClaveProdServ       string                `json:"clave_prod_serv,omitempty"`
	NoIdentificacion    string                `json:"no_identificacion,omitempty"`
	Cantidad            *decimal.Decimal      `json:"cantidad,omitempty"`
	Unidad              string                `json:"unidad,omitempty"`
	Descripcion         string                `json:"descripcion,omitempty"`
	ValorUnitario       *decimal.Decimal      `json:"valor_unitario,omitempty"`
	Importe             *decimal.Decimal      `json:"importe,omitempty"`
	InformacionAduanera []InformacionAduanera `json:"informacion_aduanera,omitempty"`
}

// Impuestos son los totales de impuestos del comprobante. El mismo struct sirve
// para el nodo Impuestos de un Pago 1.0.
type Impuestos struct {
	TotalImpuestosRetenidos   *decimal.Decimal `json:"total_impuestos_retenidos,omitempty"`
	TotalImpuestosTrasladados *decimal.Decimal `json:"total_impuestos_trasladados,omitempty"`
	Retenciones               []Impuesto       `json:"retenciones,omitempty"`
	Traslados                 []Impuesto       `json:"traslados,omitempty"`
}

// TimbreFiscalDigital 1.0 o 1.1. En 1.0 los atributos van en minúsculas
// (version, selloCFD…) y no existe RfcProvCertif; se leen igual.
type TimbreFiscalDigital struct {
	Version string `json:"version,omitempty"`
	// UUID normalizado: sin espacios y en mayúsculas.
	UUID string `json:"uuid,omitempty"`
	// UUIDOriginal es el texto como venía, solo si difiere de UUID.
	UUIDOriginal     string `json:"uuid_original,omitempty"`
	FechaTimbrado    Fecha  `json:"fecha_timbrado,omitzero"`
	RfcProvCertif    string `json:"rfc_prov_certif,omitempty"`
	Leyenda          string `json:"leyenda,omitempty"`
	SelloCFD         string `json:"sello_cfd,omitempty"`
	NoCertificadoSAT string `json:"no_certificado_sat,omitempty"`
	SelloSAT         string `json:"sello_sat,omitempty"`
}

// Pagos es el complemento de recepción de pagos, 1.0 o 2.0 normalizados en un
// solo struct. Los campos que solo existen en una versión quedan vacíos en la
// otra (Totales, ImpuestosP, EquivalenciaDR, ObjetoImpDR, ImpuestosDR son de
// 2.0; TipoCambioDR, MetodoDePagoDR e Impuestos del pago son de 1.0).
type Pagos struct {
	Version string        `json:"version,omitempty"`
	Totales *TotalesPagos `json:"totales,omitempty"`
	Pago    []Pago        `json:"pago,omitempty"`
}

// TotalesPagos es el nodo pago20:Totales del complemento Pagos 2.0.
type TotalesPagos struct {
	TotalRetencionesIVA         *decimal.Decimal `json:"total_retenciones_iva,omitempty"`
	TotalRetencionesISR         *decimal.Decimal `json:"total_retenciones_isr,omitempty"`
	TotalRetencionesIEPS        *decimal.Decimal `json:"total_retenciones_ieps,omitempty"`
	TotalTrasladosBaseIVA16     *decimal.Decimal `json:"total_traslados_base_iva16,omitempty"`
	TotalTrasladosImpuestoIVA16 *decimal.Decimal `json:"total_traslados_impuesto_iva16,omitempty"`
	TotalTrasladosBaseIVA8      *decimal.Decimal `json:"total_traslados_base_iva8,omitempty"`
	TotalTrasladosImpuestoIVA8  *decimal.Decimal `json:"total_traslados_impuesto_iva8,omitempty"`
	TotalTrasladosBaseIVA0      *decimal.Decimal `json:"total_traslados_base_iva0,omitempty"`
	TotalTrasladosImpuestoIVA0  *decimal.Decimal `json:"total_traslados_impuesto_iva0,omitempty"`
	TotalTrasladosBaseIVAExento *decimal.Decimal `json:"total_traslados_base_iva_exento,omitempty"`
	MontoTotalPagos             *decimal.Decimal `json:"monto_total_pagos,omitempty"`
}

// Pago es el nodo Pago del complemento Pagos (pago10:Pago o pago20:Pago).
type Pago struct {
	FechaPago        Fecha              `json:"fecha_pago,omitzero"`
	FormaDePagoP     string             `json:"forma_de_pago_p,omitempty"`
	MonedaP          string             `json:"moneda_p,omitempty"`
	TipoCambioP      *decimal.Decimal   `json:"tipo_cambio_p,omitempty"`
	Monto            *decimal.Decimal   `json:"monto,omitempty"`
	NumOperacion     string             `json:"num_operacion,omitempty"`
	RfcEmisorCtaOrd  string             `json:"rfc_emisor_cta_ord,omitempty"`
	NomBancoOrdExt   string             `json:"nom_banco_ord_ext,omitempty"`
	CtaOrdenante     string             `json:"cta_ordenante,omitempty"`
	RfcEmisorCtaBen  string             `json:"rfc_emisor_cta_ben,omitempty"`
	CtaBeneficiario  string             `json:"cta_beneficiario,omitempty"`
	TipoCadPago      string             `json:"tipo_cad_pago,omitempty"`
	CertPago         string             `json:"cert_pago,omitempty"`
	CadPago          string             `json:"cad_pago,omitempty"`
	SelloPago        string             `json:"sello_pago,omitempty"`
	DoctoRelacionado []DoctoRelacionado `json:"docto_relacionado,omitempty"`
	ImpuestosP       *ImpuestosP        `json:"impuestos_p,omitempty"` // 2.0
	Impuestos        []Impuestos        `json:"impuestos,omitempty"`   // 1.0
}

// DoctoRelacionado es el nodo DoctoRelacionado de un Pago del complemento Pagos.
type DoctoRelacionado struct {
	// IdDocumento normalizado como el UUID del timbre (mayúsculas), para
	// cruzarlo contra él.
	IdDocumento         string           `json:"id_documento,omitempty"`
	IdDocumentoOriginal string           `json:"id_documento_original,omitempty"` // solo si difiere
	Serie               string           `json:"serie,omitempty"`
	Folio               string           `json:"folio,omitempty"`
	MonedaDR            string           `json:"moneda_dr,omitempty"`
	EquivalenciaDR      *decimal.Decimal `json:"equivalencia_dr,omitempty"`   // 2.0
	TipoCambioDR        *decimal.Decimal `json:"tipo_cambio_dr,omitempty"`    // 1.0
	MetodoDePagoDR      string           `json:"metodo_de_pago_dr,omitempty"` // 1.0
	NumParcialidad      string           `json:"num_parcialidad,omitempty"`
	ImpSaldoAnt         *decimal.Decimal `json:"imp_saldo_ant,omitempty"`
	ImpPagado           *decimal.Decimal `json:"imp_pagado,omitempty"`
	ImpSaldoInsoluto    *decimal.Decimal `json:"imp_saldo_insoluto,omitempty"`
	ObjetoImpDR         string           `json:"objeto_imp_dr,omitempty"` // 2.0
	ImpuestosDR         *ImpuestosDR     `json:"impuestos_dr,omitempty"`  // 2.0
}

// ImpuestosDR es el nodo pago20:ImpuestosDR de un DoctoRelacionado del complemento Pagos 2.0.
type ImpuestosDR struct {
	RetencionesDR []ImpuestoDR `json:"retenciones_dr,omitempty"`
	TrasladosDR   []ImpuestoDR `json:"traslados_dr,omitempty"`
}

// ImpuestoDR es un nodo pago20:RetencionDR o pago20:TrasladoDR de ImpuestosDR, en el complemento Pagos 2.0.
type ImpuestoDR struct {
	BaseDR       *decimal.Decimal `json:"base_dr,omitempty"`
	ImpuestoDR   string           `json:"impuesto_dr,omitempty"`
	TipoFactorDR string           `json:"tipo_factor_dr,omitempty"`
	TasaOCuotaDR *decimal.Decimal `json:"tasa_o_cuota_dr,omitempty"`
	ImporteDR    *decimal.Decimal `json:"importe_dr,omitempty"`
}

// ImpuestosP es el nodo pago20:ImpuestosP de un Pago del complemento Pagos 2.0.
type ImpuestosP struct {
	RetencionesP []ImpuestoP `json:"retenciones_p,omitempty"`
	TrasladosP   []ImpuestoP `json:"traslados_p,omitempty"`
}

// ImpuestoP es un nodo pago20:RetencionP o pago20:TrasladoP de ImpuestosP, en el complemento Pagos 2.0.
type ImpuestoP struct {
	BaseP       *decimal.Decimal `json:"base_p,omitempty"`
	ImpuestoP   string           `json:"impuesto_p,omitempty"`
	TipoFactorP string           `json:"tipo_factor_p,omitempty"`
	TasaOCuotaP *decimal.Decimal `json:"tasa_o_cuota_p,omitempty"`
	ImporteP    *decimal.Decimal `json:"importe_p,omitempty"`
}

// Nomina es el complemento de nómina 1.2 o 1.1 (CFDI 3.2), normalizados en
// un solo struct; Version dice cuál. 1.1 lleva los datos del trabajador como
// atributos del nodo Nomina (los campos marcados 1.1, vacíos en 1.2) y no
// tiene Emisor, Receptor, OtrosPagos ni los totales de 1.2. Cada atributo
// queda donde lo pone su XSD: un 1.1 no llena Receptor.
type Nomina struct {
	Version            string           `json:"version,omitempty"`
	TipoNomina         string           `json:"tipo_nomina,omitempty"`
	RegistroPatronal   string           `json:"registro_patronal,omitempty"`    // 1.1
	NumEmpleado        string           `json:"num_empleado,omitempty"`         // 1.1
	CURP               string           `json:"curp,omitempty"`                 // 1.1
	TipoRegimen        string           `json:"tipo_regimen,omitempty"`         // 1.1
	NumSeguridadSocial string           `json:"num_seguridad_social,omitempty"` // 1.1
	FechaPago          Fecha            `json:"fecha_pago,omitzero"`
	FechaInicialPago   Fecha            `json:"fecha_inicial_pago,omitzero"`
	FechaFinalPago     Fecha            `json:"fecha_final_pago,omitzero"`
	NumDiasPagados     *decimal.Decimal `json:"num_dias_pagados,omitempty"`
	TotalPercepciones  *decimal.Decimal `json:"total_percepciones,omitempty"`
	TotalDeducciones   *decimal.Decimal `json:"total_deducciones,omitempty"`
	TotalOtrosPagos    *decimal.Decimal `json:"total_otros_pagos,omitempty"`

	// 1.1: el resto de los atributos del trabajador. Antiguedad son semanas
	// (un entero), no la duración P…W de 1.2.
	Departamento           string           `json:"departamento,omitempty"`
	CLABE                  string           `json:"clabe,omitempty"`
	Banco                  string           `json:"banco,omitempty"`
	FechaInicioRelLaboral  Fecha            `json:"fecha_inicio_rel_laboral,omitzero"`
	Antiguedad             string           `json:"antiguedad,omitempty"`
	Puesto                 string           `json:"puesto,omitempty"`
	TipoContrato           string           `json:"tipo_contrato,omitempty"`
	TipoJornada            string           `json:"tipo_jornada,omitempty"`
	PeriodicidadPago       string           `json:"periodicidad_pago,omitempty"`
	SalarioBaseCotApor     *decimal.Decimal `json:"salario_base_cot_apor,omitempty"`
	RiesgoPuesto           string           `json:"riesgo_puesto,omitempty"`
	SalarioDiarioIntegrado *decimal.Decimal `json:"salario_diario_integrado,omitempty"`

	Emisor        *NominaEmisor   `json:"emisor,omitempty"`
	Receptor      *NominaReceptor `json:"receptor,omitempty"`
	Percepciones  *Percepciones   `json:"percepciones,omitempty"`
	Deducciones   *Deducciones    `json:"deducciones,omitempty"`
	OtrosPagos    []OtroPago      `json:"otros_pagos,omitempty"`
	Incapacidades []Incapacidad   `json:"incapacidades,omitempty"`
	HorasExtras   []HorasExtra    `json:"horas_extras,omitempty"` // 1.1: HorasExtras/HorasExtra
}

// NominaEmisor es el nodo nomina12:Emisor del complemento Nómina.
type NominaEmisor struct {
	Curp             string       `json:"curp,omitempty"`
	RegistroPatronal string       `json:"registro_patronal,omitempty"`
	RfcPatronOrigen  string       `json:"rfc_patron_origen,omitempty"`
	EntidadSNCF      *EntidadSNCF `json:"entidad_sncf,omitempty"`
}

// EntidadSNCF es el nodo nomina12:EntidadSNCF del Emisor del complemento Nómina.
type EntidadSNCF struct {
	OrigenRecurso      string           `json:"origen_recurso,omitempty"`
	MontoRecursoPropio *decimal.Decimal `json:"monto_recurso_propio,omitempty"`
}

// NominaReceptor es el nodo nomina12:Receptor del complemento Nómina.
type NominaReceptor struct {
	Curp                   string            `json:"curp,omitempty"`
	NumSeguridadSocial     string            `json:"num_seguridad_social,omitempty"`
	FechaInicioRelLaboral  Fecha             `json:"fecha_inicio_rel_laboral,omitzero"`
	Antiguedad             string            `json:"antiguedad,omitempty"` // atributo "Antigüedad"
	TipoContrato           string            `json:"tipo_contrato,omitempty"`
	Sindicalizado          string            `json:"sindicalizado,omitempty"`
	TipoJornada            string            `json:"tipo_jornada,omitempty"`
	TipoRegimen            string            `json:"tipo_regimen,omitempty"`
	NumEmpleado            string            `json:"num_empleado,omitempty"`
	Departamento           string            `json:"departamento,omitempty"`
	Puesto                 string            `json:"puesto,omitempty"`
	RiesgoPuesto           string            `json:"riesgo_puesto,omitempty"`
	PeriodicidadPago       string            `json:"periodicidad_pago,omitempty"`
	Banco                  string            `json:"banco,omitempty"`
	CuentaBancaria         string            `json:"cuenta_bancaria,omitempty"`
	SalarioBaseCotApor     *decimal.Decimal  `json:"salario_base_cot_apor,omitempty"`
	SalarioDiarioIntegrado *decimal.Decimal  `json:"salario_diario_integrado,omitempty"`
	ClaveEntFed            string            `json:"clave_ent_fed,omitempty"`
	SubContratacion        []SubContratacion `json:"sub_contratacion,omitempty"`
}

// SubContratacion es el nodo nomina12:SubContratacion del Receptor del complemento Nómina.
type SubContratacion struct {
	RfcLabora        string           `json:"rfc_labora,omitempty"`
	PorcentajeTiempo *decimal.Decimal `json:"porcentaje_tiempo,omitempty"`
}

// Percepciones es el nodo nomina12:Percepciones del complemento Nómina.
type Percepciones struct {
	TotalSueldos                 *decimal.Decimal         `json:"total_sueldos,omitempty"`
	TotalSeparacionIndemnizacion *decimal.Decimal         `json:"total_separacion_indemnizacion,omitempty"`
	TotalJubilacionPensionRetiro *decimal.Decimal         `json:"total_jubilacion_pension_retiro,omitempty"`
	TotalGravado                 *decimal.Decimal         `json:"total_gravado,omitempty"`
	TotalExento                  *decimal.Decimal         `json:"total_exento,omitempty"`
	Percepcion                   []Percepcion             `json:"percepcion,omitempty"`
	JubilacionPensionRetiro      *JubilacionPensionRetiro `json:"jubilacion_pension_retiro,omitempty"`
	SeparacionIndemnizacion      *SeparacionIndemnizacion `json:"separacion_indemnizacion,omitempty"`
}

// Percepcion es el nodo nomina12:Percepcion del complemento Nómina.
type Percepcion struct {
	TipoPercepcion   string            `json:"tipo_percepcion,omitempty"`
	Clave            string            `json:"clave,omitempty"`
	Concepto         string            `json:"concepto,omitempty"`
	ImporteGravado   *decimal.Decimal  `json:"importe_gravado,omitempty"`
	ImporteExento    *decimal.Decimal  `json:"importe_exento,omitempty"`
	AccionesOTitulos *AccionesOTitulos `json:"acciones_o_titulos,omitempty"`
	HorasExtra       []HorasExtra      `json:"horas_extra,omitempty"`
}

// AccionesOTitulos es el nodo nomina12:AccionesOTitulos de una Percepcion del complemento Nómina.
type AccionesOTitulos struct {
	ValorMercado      *decimal.Decimal `json:"valor_mercado,omitempty"`
	PrecioAlOtorgarse *decimal.Decimal `json:"precio_al_otorgarse,omitempty"`
}

// HorasExtra es el nodo nomina12:HorasExtra de una Percepcion del complemento Nómina.
type HorasExtra struct {
	Dias          string           `json:"dias,omitempty"`
	TipoHoras     string           `json:"tipo_horas,omitempty"`
	HorasExtra    string           `json:"horas_extra,omitempty"`
	ImportePagado *decimal.Decimal `json:"importe_pagado,omitempty"`
}

// JubilacionPensionRetiro es el nodo nomina12:JubilacionPensionRetiro de las Percepciones del complemento Nómina.
type JubilacionPensionRetiro struct {
	TotalUnaExhibicion  *decimal.Decimal `json:"total_una_exhibicion,omitempty"`
	TotalParcialidad    *decimal.Decimal `json:"total_parcialidad,omitempty"`
	MontoDiario         *decimal.Decimal `json:"monto_diario,omitempty"`
	IngresoAcumulable   *decimal.Decimal `json:"ingreso_acumulable,omitempty"`
	IngresoNoAcumulable *decimal.Decimal `json:"ingreso_no_acumulable,omitempty"`
}

// SeparacionIndemnizacion es el nodo nomina12:SeparacionIndemnizacion de las Percepciones del complemento Nómina.
type SeparacionIndemnizacion struct {
	TotalPagado         *decimal.Decimal `json:"total_pagado,omitempty"`
	NumAniosServicio    string           `json:"num_anios_servicio,omitempty"` // atributo "NumAñosServicio"
	UltimoSueldoMensOrd *decimal.Decimal `json:"ultimo_sueldo_mens_ord,omitempty"`
	IngresoAcumulable   *decimal.Decimal `json:"ingreso_acumulable,omitempty"`
	IngresoNoAcumulable *decimal.Decimal `json:"ingreso_no_acumulable,omitempty"`
}

// Deducciones es el nodo nomina12:Deducciones del complemento Nómina.
type Deducciones struct {
	TotalOtrasDeducciones   *decimal.Decimal `json:"total_otras_deducciones,omitempty"`
	TotalImpuestosRetenidos *decimal.Decimal `json:"total_impuestos_retenidos,omitempty"`
	TotalGravado            *decimal.Decimal `json:"total_gravado,omitempty"` // 1.1
	TotalExento             *decimal.Decimal `json:"total_exento,omitempty"`  // 1.1
	Deduccion               []Deduccion      `json:"deduccion,omitempty"`
}

// Deduccion es el nodo nomina12:Deduccion del complemento Nómina.
type Deduccion struct {
	TipoDeduccion  string           `json:"tipo_deduccion,omitempty"`
	Clave          string           `json:"clave,omitempty"`
	Concepto       string           `json:"concepto,omitempty"`
	Importe        *decimal.Decimal `json:"importe,omitempty"`
	ImporteGravado *decimal.Decimal `json:"importe_gravado,omitempty"` // 1.1
	ImporteExento  *decimal.Decimal `json:"importe_exento,omitempty"`  // 1.1
}

// OtroPago es el nodo nomina12:OtroPago del complemento Nómina.
type OtroPago struct {
	TipoOtroPago             string                    `json:"tipo_otro_pago,omitempty"`
	Clave                    string                    `json:"clave,omitempty"`
	Concepto                 string                    `json:"concepto,omitempty"`
	Importe                  *decimal.Decimal          `json:"importe,omitempty"`
	SubsidioAlEmpleo         *SubsidioAlEmpleo         `json:"subsidio_al_empleo,omitempty"`
	CompensacionSaldosAFavor *CompensacionSaldosAFavor `json:"compensacion_saldos_a_favor,omitempty"`
}

// SubsidioAlEmpleo es el nodo nomina12:SubsidioAlEmpleo de un OtroPago del complemento Nómina.
type SubsidioAlEmpleo struct {
	SubsidioCausado *decimal.Decimal `json:"subsidio_causado,omitempty"`
}

// CompensacionSaldosAFavor es el nodo nomina12:CompensacionSaldosAFavor de un OtroPago del complemento Nómina.
type CompensacionSaldosAFavor struct {
	SaldoAFavor     *decimal.Decimal `json:"saldo_a_favor,omitempty"`
	Anio            string           `json:"anio,omitempty"` // atributo "Año"
	RemanenteSalFav *decimal.Decimal `json:"remanente_sal_fav,omitempty"`
}

// Incapacidad es el nodo nomina12:Incapacidad del complemento Nómina.
type Incapacidad struct {
	DiasIncapacidad  string           `json:"dias_incapacidad,omitempty"`
	TipoIncapacidad  string           `json:"tipo_incapacidad,omitempty"`
	ImporteMonetario *decimal.Decimal `json:"importe_monetario,omitempty"`
	Descuento        *decimal.Decimal `json:"descuento,omitempty"` // 1.1
}
