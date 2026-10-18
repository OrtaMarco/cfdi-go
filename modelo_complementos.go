package cfdi

import "github.com/OrtaMarco/cfdi-go/decimal"

// Complementos de la v0.2. Siguen las convenciones de modelo.go; los nombres
// salen de los XSD del SAT tal cual, aunque el SAT junte palabras
// ("TotaldeTraslados" → "totalde_traslados"), para que la clave JSON se pueda
// deducir del atributo sin tabla. Los atributos que el XSD escribe en
// minúsculas ("version", "tipoOperacion") van con mayúscula inicial en Go y se
// leen igual.

// ImpuestosLocales 1.0 (implocal). Los importes locales no están en
// Comprobante.Impuestos pero sí en el Total:
//
//	Total = SubTotal - Descuento + TotalImpuestosTrasladados - TotalImpuestosRetenidos
//	        + TotaldeTraslados - TotaldeRetenciones
type ImpuestosLocales struct {
	Version            string           `json:"version,omitempty"`
	TotaldeRetenciones *decimal.Decimal `json:"totalde_retenciones,omitempty"`
	TotaldeTraslados   *decimal.Decimal `json:"totalde_traslados,omitempty"`
	// El XSD repite la secuencia (RetencionesLocales?, TrasladosLocales?):
	// cada uno puede venir varias veces.
	RetencionesLocales []RetencionesLocales `json:"retenciones_locales,omitempty"`
	TrasladosLocales   []TrasladosLocales   `json:"traslados_locales,omitempty"`
}

// RetencionesLocales es el nodo implocal:RetencionesLocales del complemento ImpuestosLocales.
type RetencionesLocales struct {
	ImpLocRetenido  string           `json:"imp_loc_retenido,omitempty"`
	TasadeRetencion *decimal.Decimal `json:"tasade_retencion,omitempty"`
	Importe         *decimal.Decimal `json:"importe,omitempty"`
}

// TrasladosLocales es el nodo implocal:TrasladosLocales del complemento ImpuestosLocales.
type TrasladosLocales struct {
	ImpLocTrasladado string           `json:"imp_loc_trasladado,omitempty"`
	TasadeTraslado   *decimal.Decimal `json:"tasade_traslado,omitempty"`
	Importe          *decimal.Decimal `json:"importe,omitempty"`
}

// CartaPorte 2.0, 3.0 y 3.1 normalizadas en un solo struct: los XSD solo
// agregan, quitan o mueven nodos y atributos, sin cambiar el tipo de ninguno
// que compartan, así que la unión no pierde datos. Version dice cuál era y lo
// exclusivo de una versión queda vacío en las otras:
//
//   - 2.0: Mercancia.Pedimentos, TransporteMaritimo.NumCertITC.
//   - 3.0 y 3.1: IdCCP, RegistroISTMO, UbicacionPoloOrigen/Destino,
//     Mercancia.DocumentacionAduanera y los datos de COFEPRIS de Mercancia,
//     IdentificacionVehicular.PesoBrutoVehicular, TransporteMaritimo.Puntal y
//     PermisoTempNavegacion, los datos CCP de los contenedores marítimos.
//   - 3.0: RegimenAduanero (atributo) y TransporteMaritimo.RemolquesCCP.
//   - 3.1: RegimenesAduaneros (lista) y Contenedor.RemolquesCCP (los remolques
//     pasaron del transporte marítimo a cada contenedor).
//
// CartaPorte 1.0 (namespace .../CartaPorte) entra en el mismo struct con lo
// que comparte (Ubicacion con su Domicilio, Mercancia, CantidadTransporta,
// DetalleMercancia, TransporteMaritimo, TransporteAereo,
// TransporteFerroviario) y campos propios para lo que 2.0 cambió de forma:
// Ubicacion.Origen y Destino, Mercancias.AutotransporteFederal y
// FiguraTransporte10 (en 1.0 FiguraTransporte es un nodo con CveTransporte y
// sus Operadores, Propietario, Arrendatario y Notificado, y no cabe en la
// lista de TiposFigura de 2.0).
type CartaPorte struct {
	Version              string           `json:"version,omitempty"`
	IdCCP                string           `json:"id_ccp,omitempty"`
	TranspInternac       string           `json:"transp_internac,omitempty"`
	RegimenAduanero      string           `json:"regimen_aduanero,omitempty"` // 3.0
	EntradaSalidaMerc    string           `json:"entrada_salida_merc,omitempty"`
	PaisOrigenDestino    string           `json:"pais_origen_destino,omitempty"`
	ViaEntradaSalida     string           `json:"via_entrada_salida,omitempty"`
	TotalDistRec         *decimal.Decimal `json:"total_dist_rec,omitempty"`
	RegistroISTMO        string           `json:"registro_istmo,omitempty"`
	UbicacionPoloOrigen  string           `json:"ubicacion_polo_origen,omitempty"`
	UbicacionPoloDestino string           `json:"ubicacion_polo_destino,omitempty"`

	RegimenesAduaneros []RegimenAduaneroCCP          `json:"regimenes_aduaneros,omitempty"` // 3.1
	Ubicaciones        []Ubicacion                   `json:"ubicaciones,omitempty"`
	Mercancias         *CartaPorteMercancias         `json:"mercancias,omitempty"`
	FiguraTransporte   []TiposFigura                 `json:"figura_transporte,omitempty"`
	FiguraTransporte10 *CartaPorte10FiguraTransporte `json:"figura_transporte10,omitempty"` // 1.0
}

// CartaPorte10Origen y CartaPorte10Destino son los nodos Origen y Destino de
// la Ubicacion de CartaPorte 1.0 (2.0 los juntó en atributos de Ubicacion).
type CartaPorte10Origen struct {
	IDOrigen          string `json:"id_origen,omitempty"`
	RFCRemitente      string `json:"rfc_remitente,omitempty"`
	NombreRemitente   string `json:"nombre_remitente,omitempty"`
	NumRegIdTrib      string `json:"num_reg_id_trib,omitempty"`
	ResidenciaFiscal  string `json:"residencia_fiscal,omitempty"`
	NumEstacion       string `json:"num_estacion,omitempty"`
	NombreEstacion    string `json:"nombre_estacion,omitempty"`
	NavegacionTrafico string `json:"navegacion_trafico,omitempty"`
	FechaHoraSalida   Fecha  `json:"fecha_hora_salida,omitzero"`
}

// CartaPorte10Destino es el nodo Destino de una Ubicacion de CartaPorte 1.0.
type CartaPorte10Destino struct {
	IDDestino            string `json:"id_destino,omitempty"`
	RFCDestinatario      string `json:"rfc_destinatario,omitempty"`
	NombreDestinatario   string `json:"nombre_destinatario,omitempty"`
	NumRegIdTrib         string `json:"num_reg_id_trib,omitempty"`
	ResidenciaFiscal     string `json:"residencia_fiscal,omitempty"`
	NumEstacion          string `json:"num_estacion,omitempty"`
	NombreEstacion       string `json:"nombre_estacion,omitempty"`
	NavegacionTrafico    string `json:"navegacion_trafico,omitempty"`
	FechaHoraProgLlegada Fecha  `json:"fecha_hora_prog_llegada,omitzero"`
}

// AutotransporteFederal es el autotransporte de CartaPorte 1.0: los seguros
// van en atributos (en 2.0, nodo Seguros).
type AutotransporteFederal struct {
	PermSCT                 string                   `json:"perm_sct,omitempty"`
	NumPermisoSCT           string                   `json:"num_permiso_sct,omitempty"`
	NombreAseg              string                   `json:"nombre_aseg,omitempty"`
	NumPolizaSeguro         string                   `json:"num_poliza_seguro,omitempty"`
	IdentificacionVehicular *IdentificacionVehicular `json:"identificacion_vehicular,omitempty"`
	Remolques               []Remolque               `json:"remolques,omitempty"`
}

// CartaPorte10FiguraTransporte es la FiguraTransporte de 1.0. Operadores
// puede repetirse y no tiene atributos: sus Operador van en una sola lista,
// en orden de documento.
type CartaPorte10FiguraTransporte struct {
	CveTransporte string                     `json:"cve_transporte,omitempty"`
	Operadores    []CartaPorte10Operador     `json:"operadores,omitempty"`
	Propietario   []CartaPorte10Propietario  `json:"propietario,omitempty"`
	Arrendatario  []CartaPorte10Arrendatario `json:"arrendatario,omitempty"`
	Notificado    []CartaPorte10Notificado   `json:"notificado,omitempty"`
}

// CartaPorte10Operador es el nodo Operador de la FiguraTransporte de CartaPorte 1.0.
type CartaPorte10Operador struct {
	RFCOperador              string     `json:"rfc_operador,omitempty"`
	NumLicencia              string     `json:"num_licencia,omitempty"`
	NombreOperador           string     `json:"nombre_operador,omitempty"`
	NumRegIdTribOperador     string     `json:"num_reg_id_trib_operador,omitempty"`
	ResidenciaFiscalOperador string     `json:"residencia_fiscal_operador,omitempty"`
	Domicilio                *Domicilio `json:"domicilio,omitempty"`
}

// CartaPorte10Propietario es el nodo Propietario de la FiguraTransporte de CartaPorte 1.0.
type CartaPorte10Propietario struct {
	RFCPropietario              string     `json:"rfc_propietario,omitempty"`
	NombrePropietario           string     `json:"nombre_propietario,omitempty"`
	NumRegIdTribPropietario     string     `json:"num_reg_id_trib_propietario,omitempty"`
	ResidenciaFiscalPropietario string     `json:"residencia_fiscal_propietario,omitempty"`
	Domicilio                   *Domicilio `json:"domicilio,omitempty"`
}

// CartaPorte10Arrendatario es el nodo Arrendatario de la FiguraTransporte de CartaPorte 1.0.
type CartaPorte10Arrendatario struct {
	RFCArrendatario              string     `json:"rfc_arrendatario,omitempty"`
	NombreArrendatario           string     `json:"nombre_arrendatario,omitempty"`
	NumRegIdTribArrendatario     string     `json:"num_reg_id_trib_arrendatario,omitempty"`
	ResidenciaFiscalArrendatario string     `json:"residencia_fiscal_arrendatario,omitempty"`
	Domicilio                    *Domicilio `json:"domicilio,omitempty"`
}

// CartaPorte10Notificado es el nodo Notificado de la FiguraTransporte de CartaPorte 1.0.
type CartaPorte10Notificado struct {
	RFCNotificado              string     `json:"rfc_notificado,omitempty"`
	NombreNotificado           string     `json:"nombre_notificado,omitempty"`
	NumRegIdTribNotificado     string     `json:"num_reg_id_trib_notificado,omitempty"`
	ResidenciaFiscalNotificado string     `json:"residencia_fiscal_notificado,omitempty"`
	Domicilio                  *Domicilio `json:"domicilio,omitempty"`
}

// RegimenAduaneroCCP es el nodo RegimenAduaneroCCP de CartaPorte 3.1.
type RegimenAduaneroCCP struct {
	RegimenAduanero string `json:"regimen_aduanero,omitempty"`
}

// Domicilio de carta porte y de comercio exterior (mismos atributos).
type Domicilio struct {
	Calle          string `json:"calle,omitempty"`
	NumeroExterior string `json:"numero_exterior,omitempty"`
	NumeroInterior string `json:"numero_interior,omitempty"`
	Colonia        string `json:"colonia,omitempty"`
	Localidad      string `json:"localidad,omitempty"`
	Referencia     string `json:"referencia,omitempty"`
	Municipio      string `json:"municipio,omitempty"`
	Estado         string `json:"estado,omitempty"`
	Pais           string `json:"pais,omitempty"`
	CodigoPostal   string `json:"codigo_postal,omitempty"`
}

// Ubicacion es el nodo Ubicacion del complemento CartaPorte.
type Ubicacion struct {
	TipoUbicacion               string               `json:"tipo_ubicacion,omitempty"`
	IDUbicacion                 string               `json:"id_ubicacion,omitempty"`
	RFCRemitenteDestinatario    string               `json:"rfc_remitente_destinatario,omitempty"`
	NombreRemitenteDestinatario string               `json:"nombre_remitente_destinatario,omitempty"`
	NumRegIdTrib                string               `json:"num_reg_id_trib,omitempty"`
	ResidenciaFiscal            string               `json:"residencia_fiscal,omitempty"`
	NumEstacion                 string               `json:"num_estacion,omitempty"`
	NombreEstacion              string               `json:"nombre_estacion,omitempty"`
	NavegacionTrafico           string               `json:"navegacion_trafico,omitempty"`
	FechaHoraSalidaLlegada      Fecha                `json:"fecha_hora_salida_llegada,omitzero"`
	TipoEstacion                string               `json:"tipo_estacion,omitempty"`
	DistanciaRecorrida          *decimal.Decimal     `json:"distancia_recorrida,omitempty"`
	Origen                      *CartaPorte10Origen  `json:"origen,omitempty"`  // 1.0
	Destino                     *CartaPorte10Destino `json:"destino,omitempty"` // 1.0
	Domicilio                   *Domicilio           `json:"domicilio,omitempty"`
}

// CartaPorteMercancias es el nodo Mercancias del complemento CartaPorte.
type CartaPorteMercancias struct {
	PesoBrutoTotal                        *decimal.Decimal `json:"peso_bruto_total,omitempty"`
	UnidadPeso                            string           `json:"unidad_peso,omitempty"`
	PesoNetoTotal                         *decimal.Decimal `json:"peso_neto_total,omitempty"`
	NumTotalMercancias                    string           `json:"num_total_mercancias,omitempty"`
	CargoPorTasacion                      *decimal.Decimal `json:"cargo_por_tasacion,omitempty"`
	LogisticaInversaRecoleccionDevolucion string           `json:"logistica_inversa_recoleccion_devolucion,omitempty"` // 3.x

	Mercancia             []CartaPorteMercancia  `json:"mercancia,omitempty"`
	AutotransporteFederal *AutotransporteFederal `json:"autotransporte_federal,omitempty"` // 1.0
	Autotransporte        *Autotransporte        `json:"autotransporte,omitempty"`
	TransporteMaritimo    *TransporteMaritimo    `json:"transporte_maritimo,omitempty"`
	TransporteAereo       *TransporteAereo       `json:"transporte_aereo,omitempty"`
	TransporteFerroviario *TransporteFerroviario `json:"transporte_ferroviario,omitempty"`
}

// CartaPorteMercancia es el nodo Mercancia del complemento CartaPorte.
type CartaPorteMercancia struct {
	BienesTransp                       string           `json:"bienes_transp,omitempty"`
	ClaveSTCC                          string           `json:"clave_stcc,omitempty"`
	Descripcion                        string           `json:"descripcion,omitempty"`
	Cantidad                           *decimal.Decimal `json:"cantidad,omitempty"`
	ClaveUnidad                        string           `json:"clave_unidad,omitempty"`
	Unidad                             string           `json:"unidad,omitempty"`
	Dimensiones                        string           `json:"dimensiones,omitempty"`
	MaterialPeligroso                  string           `json:"material_peligroso,omitempty"`
	CveMaterialPeligroso               string           `json:"cve_material_peligroso,omitempty"`
	Embalaje                           string           `json:"embalaje,omitempty"`
	DescripEmbalaje                    string           `json:"descrip_embalaje,omitempty"`
	SectorCOFEPRIS                     string           `json:"sector_cofepris,omitempty"`
	NombreIngredienteActivo            string           `json:"nombre_ingrediente_activo,omitempty"`
	NomQuimico                         string           `json:"nom_quimico,omitempty"`
	DenominacionGenericaProd           string           `json:"denominacion_generica_prod,omitempty"`
	DenominacionDistintivaProd         string           `json:"denominacion_distintiva_prod,omitempty"`
	Fabricante                         string           `json:"fabricante,omitempty"`
	FechaCaducidad                     Fecha            `json:"fecha_caducidad,omitzero"`
	LoteMedicamento                    string           `json:"lote_medicamento,omitempty"`
	FormaFarmaceutica                  string           `json:"forma_farmaceutica,omitempty"`
	CondicionesEspTransp               string           `json:"condiciones_esp_transp,omitempty"`
	RegistroSanitarioFolioAutorizacion string           `json:"registro_sanitario_folio_autorizacion,omitempty"`
	PermisoImportacion                 string           `json:"permiso_importacion,omitempty"`
	FolioImpoVUCEM                     string           `json:"folio_impo_vucem,omitempty"`
	NumCAS                             string           `json:"num_cas,omitempty"`
	RazonSocialEmpImp                  string           `json:"razon_social_emp_imp,omitempty"`
	NumRegSanPlagCOFEPRIS              string           `json:"num_reg_san_plag_cofepris,omitempty"`
	DatosFabricante                    string           `json:"datos_fabricante,omitempty"`
	DatosFormulador                    string           `json:"datos_formulador,omitempty"`
	DatosMaquilador                    string           `json:"datos_maquilador,omitempty"`
	UsoAutorizado                      string           `json:"uso_autorizado,omitempty"`
	PesoEnKg                           *decimal.Decimal `json:"peso_en_kg,omitempty"`
	ValorMercancia                     *decimal.Decimal `json:"valor_mercancia,omitempty"`
	Moneda                             string           `json:"moneda,omitempty"`
	FraccionArancelaria                string           `json:"fraccion_arancelaria,omitempty"`
	UUIDComercioExt                    string           `json:"uuid_comercio_ext,omitempty"`
	TipoMateria                        string           `json:"tipo_materia,omitempty"`
	DescripcionMateria                 string           `json:"descripcion_materia,omitempty"`

	Pedimentos            []Pedimentos            `json:"pedimentos,omitempty"`             // 2.0
	DocumentacionAduanera []DocumentacionAduanera `json:"documentacion_aduanera,omitempty"` // 3.x
	GuiasIdentificacion   []GuiasIdentificacion   `json:"guias_identificacion,omitempty"`
	CantidadTransporta    []CantidadTransporta    `json:"cantidad_transporta,omitempty"`
	DetalleMercancia      *DetalleMercancia       `json:"detalle_mercancia,omitempty"`
}

// Pedimentos es el nodo Pedimentos de una Mercancia del complemento CartaPorte.
type Pedimentos struct {
	Pedimento string `json:"pedimento,omitempty"`
}

// DocumentacionAduanera es el nodo DocumentacionAduanera de una Mercancia del complemento CartaPorte (3.x).
type DocumentacionAduanera struct {
	TipoDocumento    string `json:"tipo_documento,omitempty"`
	NumPedimento     string `json:"num_pedimento,omitempty"`
	IdentDocAduanero string `json:"ident_doc_aduanero,omitempty"`
	RFCImpo          string `json:"rfc_impo,omitempty"`
}

// GuiasIdentificacion es el nodo GuiasIdentificacion de una Mercancia del complemento CartaPorte.
type GuiasIdentificacion struct {
	NumeroGuiaIdentificacion  string           `json:"numero_guia_identificacion,omitempty"`
	DescripGuiaIdentificacion string           `json:"descrip_guia_identificacion,omitempty"`
	PesoGuiaIdentificacion    *decimal.Decimal `json:"peso_guia_identificacion,omitempty"`
}

// CantidadTransporta es el nodo CantidadTransporta de una Mercancia del complemento CartaPorte.
type CantidadTransporta struct {
	Cantidad       *decimal.Decimal `json:"cantidad,omitempty"`
	IDOrigen       string           `json:"id_origen,omitempty"`
	IDDestino      string           `json:"id_destino,omitempty"`
	CvesTransporte string           `json:"cves_transporte,omitempty"`
}

// DetalleMercancia es el nodo DetalleMercancia de una Mercancia del complemento CartaPorte.
type DetalleMercancia struct {
	UnidadPeso     string           `json:"unidad_peso,omitempty"` // 1.0 (2.0 lo llamó UnidadPesoMerc)
	UnidadPesoMerc string           `json:"unidad_peso_merc,omitempty"`
	PesoBruto      *decimal.Decimal `json:"peso_bruto,omitempty"`
	PesoNeto       *decimal.Decimal `json:"peso_neto,omitempty"`
	PesoTara       *decimal.Decimal `json:"peso_tara,omitempty"`
	NumPiezas      string           `json:"num_piezas,omitempty"`
}

// Autotransporte es el nodo Autotransporte del complemento CartaPorte (2.0 y 3.x).
type Autotransporte struct {
	PermSCT                 string                   `json:"perm_sct,omitempty"`
	NumPermisoSCT           string                   `json:"num_permiso_sct,omitempty"`
	IdentificacionVehicular *IdentificacionVehicular `json:"identificacion_vehicular,omitempty"`
	Seguros                 *Seguros                 `json:"seguros,omitempty"`
	Remolques               []Remolque               `json:"remolques,omitempty"`
}

// IdentificacionVehicular es el nodo IdentificacionVehicular del Autotransporte del complemento CartaPorte.
type IdentificacionVehicular struct {
	ConfigVehicular    string           `json:"config_vehicular,omitempty"`
	PesoBrutoVehicular *decimal.Decimal `json:"peso_bruto_vehicular,omitempty"` // 3.x
	PlacaVM            string           `json:"placa_vm,omitempty"`
	AnioModeloVM       string           `json:"anio_modelo_vm,omitempty"`
}

// Seguros es el nodo Seguros del Autotransporte del complemento CartaPorte.
type Seguros struct {
	AseguraRespCivil   string           `json:"asegura_resp_civil,omitempty"`
	PolizaRespCivil    string           `json:"poliza_resp_civil,omitempty"`
	AseguraMedAmbiente string           `json:"asegura_med_ambiente,omitempty"`
	PolizaMedAmbiente  string           `json:"poliza_med_ambiente,omitempty"`
	AseguraCarga       string           `json:"asegura_carga,omitempty"`
	PolizaCarga        string           `json:"poliza_carga,omitempty"`
	PrimaSeguro        *decimal.Decimal `json:"prima_seguro,omitempty"`
}

// Remolque es el nodo Remolque del Autotransporte del complemento CartaPorte.
type Remolque struct {
	SubTipoRem string `json:"sub_tipo_rem,omitempty"`
	Placa      string `json:"placa,omitempty"`
}

// TransporteMaritimo es el nodo TransporteMaritimo del complemento CartaPorte.
type TransporteMaritimo struct {
	PermSCT                string           `json:"perm_sct,omitempty"`
	NumPermisoSCT          string           `json:"num_permiso_sct,omitempty"`
	NombreAseg             string           `json:"nombre_aseg,omitempty"`
	NumPolizaSeguro        string           `json:"num_poliza_seguro,omitempty"`
	TipoEmbarcacion        string           `json:"tipo_embarcacion,omitempty"`
	Matricula              string           `json:"matricula,omitempty"`
	NumeroOMI              string           `json:"numero_omi,omitempty"`
	AnioEmbarcacion        string           `json:"anio_embarcacion,omitempty"`
	NombreEmbarc           string           `json:"nombre_embarc,omitempty"`
	NacionalidadEmbarc     string           `json:"nacionalidad_embarc,omitempty"`
	UnidadesDeArqBruto     *decimal.Decimal `json:"unidades_de_arq_bruto,omitempty"`
	TipoCarga              string           `json:"tipo_carga,omitempty"`
	NumCertITC             string           `json:"num_cert_itc,omitempty"` // 2.0
	Eslora                 *decimal.Decimal `json:"eslora,omitempty"`
	Manga                  *decimal.Decimal `json:"manga,omitempty"`
	Calado                 *decimal.Decimal `json:"calado,omitempty"`
	Puntal                 *decimal.Decimal `json:"puntal,omitempty"` // 3.x
	LineaNaviera           string           `json:"linea_naviera,omitempty"`
	NombreAgenteNaviero    string           `json:"nombre_agente_naviero,omitempty"`
	NumAutorizacionNaviero string           `json:"num_autorizacion_naviero,omitempty"`
	NumViaje               string           `json:"num_viaje,omitempty"`
	NumConocEmbarc         string           `json:"num_conoc_embarc,omitempty"`
	PermisoTempNavegacion  string           `json:"permiso_temp_navegacion,omitempty"` // 3.x

	Contenedor   []ContenedorMaritimo `json:"contenedor,omitempty"`
	RemolquesCCP []RemolqueCCP        `json:"remolques_ccp,omitempty"` // 3.0
}

// ContenedorMaritimo es el nodo Contenedor del TransporteMaritimo del complemento CartaPorte.
type ContenedorMaritimo struct {
	MatriculaContenedor   string        `json:"matricula_contenedor,omitempty"`
	TipoContenedor        string        `json:"tipo_contenedor,omitempty"`
	NumPrecinto           string        `json:"num_precinto,omitempty"`
	IdCCPRelacionado      string        `json:"id_ccp_relacionado,omitempty"`     // 3.x
	PlacaVMCCP            string        `json:"placa_vmccp,omitempty"`            // 3.x
	FechaCertificacionCCP Fecha         `json:"fecha_certificacion_ccp,omitzero"` // 3.x
	RemolquesCCP          []RemolqueCCP `json:"remolques_ccp,omitempty"`          // 3.1
}

// RemolqueCCP es el nodo RemolqueCCP de un Contenedor marítimo de CartaPorte 3.1.
type RemolqueCCP struct {
	SubTipoRemCCP string `json:"sub_tipo_rem_ccp,omitempty"`
	PlacaCCP      string `json:"placa_ccp,omitempty"`
}

// TransporteAereo es el nodo TransporteAereo del complemento CartaPorte.
type TransporteAereo struct {
	PermSCT                  string `json:"perm_sct,omitempty"`
	NumPermisoSCT            string `json:"num_permiso_sct,omitempty"`
	MatriculaAeronave        string `json:"matricula_aeronave,omitempty"`
	NombreAseg               string `json:"nombre_aseg,omitempty"`
	NumPolizaSeguro          string `json:"num_poliza_seguro,omitempty"`
	NumeroGuia               string `json:"numero_guia,omitempty"`
	LugarContrato            string `json:"lugar_contrato,omitempty"`
	RFCTransportista         string `json:"rfc_transportista,omitempty"` // 1.0
	CodigoTransportista      string `json:"codigo_transportista,omitempty"`
	NumRegIdTribTranspor     string `json:"num_reg_id_trib_transpor,omitempty"`   // 1.0
	ResidenciaFiscalTranspor string `json:"residencia_fiscal_transpor,omitempty"` // 1.0
	NombreTransportista      string `json:"nombre_transportista,omitempty"`       // 1.0
	RFCEmbarcador            string `json:"rfc_embarcador,omitempty"`
	NumRegIdTribEmbarc       string `json:"num_reg_id_trib_embarc,omitempty"`
	ResidenciaFiscalEmbarc   string `json:"residencia_fiscal_embarc,omitempty"`
	NombreEmbarcador         string `json:"nombre_embarcador,omitempty"`
}

// TransporteFerroviario es el nodo TransporteFerroviario del complemento CartaPorte.
type TransporteFerroviario struct {
	TipoDeServicio  string           `json:"tipo_de_servicio,omitempty"`
	TipoDeTrafico   string           `json:"tipo_de_trafico,omitempty"`
	NombreAseg      string           `json:"nombre_aseg,omitempty"`
	NumPolizaSeguro string           `json:"num_poliza_seguro,omitempty"`
	Concesionario   string           `json:"concesionario,omitempty"` // 1.0
	DerechosDePaso  []DerechosDePaso `json:"derechos_de_paso,omitempty"`
	Carro           []Carro          `json:"carro,omitempty"`
}

// DerechosDePaso es el nodo DerechosDePaso del TransporteFerroviario del complemento CartaPorte.
type DerechosDePaso struct {
	TipoDerechoDePaso string           `json:"tipo_derecho_de_paso,omitempty"`
	KilometrajePagado *decimal.Decimal `json:"kilometraje_pagado,omitempty"`
}

// Carro es el nodo Carro del TransporteFerroviario del complemento CartaPorte.
type Carro struct {
	TipoCarro           string                  `json:"tipo_carro,omitempty"`
	MatriculaCarro      string                  `json:"matricula_carro,omitempty"`
	GuiaCarro           string                  `json:"guia_carro,omitempty"`
	ToneladasNetasCarro *decimal.Decimal        `json:"toneladas_netas_carro,omitempty"`
	Contenedor          []ContenedorFerroviario `json:"contenedor,omitempty"`
}

// ContenedorFerroviario es el nodo Contenedor de un Carro del complemento CartaPorte.
type ContenedorFerroviario struct {
	TipoContenedor      string           `json:"tipo_contenedor,omitempty"`
	PesoContenedorVacio *decimal.Decimal `json:"peso_contenedor_vacio,omitempty"`
	PesoNetoMercancia   *decimal.Decimal `json:"peso_neto_mercancia,omitempty"`
}

// TiposFigura es el nodo TiposFigura de la FiguraTransporte del complemento CartaPorte (2.0 y 3.x).
type TiposFigura struct {
	TipoFigura             string             `json:"tipo_figura,omitempty"`
	RFCFigura              string             `json:"rfc_figura,omitempty"`
	NumLicencia            string             `json:"num_licencia,omitempty"`
	NombreFigura           string             `json:"nombre_figura,omitempty"`
	NumRegIdTribFigura     string             `json:"num_reg_id_trib_figura,omitempty"`
	ResidenciaFiscalFigura string             `json:"residencia_fiscal_figura,omitempty"`
	PartesTransporte       []PartesTransporte `json:"partes_transporte,omitempty"`
	Domicilio              *Domicilio         `json:"domicilio,omitempty"`
}

// PartesTransporte es el nodo PartesTransporte de TiposFigura del complemento CartaPorte.
type PartesTransporte struct {
	ParteTransporte string `json:"parte_transporte,omitempty"`
}

// ComercioExterior 1.0, 1.1 y 2.0 en un solo struct: 1.1 y 2.0 tienen los
// mismos nodos y atributos salvo TipoOperacion y Subdivision, que 2.0 quitó.
// 2.0 vuelve requeridos ClaveDePedimento, CertificadoOrigen, TipoCambioUSD,
// TotalUSD, Mercancias y el Domicilio del Emisor, y acepta más decimales en
// ValorUnitarioAduana y ValorDolares (se conservan tal cual). 1.0 no tiene
// MotivoTraslado, Propietario ni domicilios del Emisor y el Receptor; su
// Receptor trae Curp y su único Destinatario Rfc y Curp.
type ComercioExterior struct {
	Version                   string           `json:"version,omitempty"`
	MotivoTraslado            string           `json:"motivo_traslado,omitempty"`
	TipoOperacion             string           `json:"tipo_operacion,omitempty"` // 1.1
	ClaveDePedimento          string           `json:"clave_de_pedimento,omitempty"`
	CertificadoOrigen         string           `json:"certificado_origen,omitempty"`
	NumCertificadoOrigen      string           `json:"num_certificado_origen,omitempty"`
	NumeroExportadorConfiable string           `json:"numero_exportador_confiable,omitempty"`
	Incoterm                  string           `json:"incoterm,omitempty"`
	Subdivision               string           `json:"subdivision,omitempty"` // 1.1
	Observaciones             string           `json:"observaciones,omitempty"`
	TipoCambioUSD             *decimal.Decimal `json:"tipo_cambio_usd,omitempty"`
	TotalUSD                  *decimal.Decimal `json:"total_usd,omitempty"`

	Emisor       *ComercioExteriorEmisor        `json:"emisor,omitempty"`
	Propietario  []ComercioExteriorPropietario  `json:"propietario,omitempty"`
	Receptor     *ComercioExteriorReceptor      `json:"receptor,omitempty"`
	Destinatario []ComercioExteriorDestinatario `json:"destinatario,omitempty"`
	Mercancias   []ComercioExteriorMercancia    `json:"mercancias,omitempty"`
}

// ComercioExteriorEmisor es el nodo Emisor del complemento ComercioExterior.
type ComercioExteriorEmisor struct {
	Curp      string     `json:"curp,omitempty"`
	Domicilio *Domicilio `json:"domicilio,omitempty"`
}

// ComercioExteriorPropietario es el nodo Propietario del complemento ComercioExterior.
type ComercioExteriorPropietario struct {
	NumRegIdTrib     string `json:"num_reg_id_trib,omitempty"`
	ResidenciaFiscal string `json:"residencia_fiscal,omitempty"`
}

// ComercioExteriorReceptor es el nodo Receptor del complemento ComercioExterior.
type ComercioExteriorReceptor struct {
	Curp         string     `json:"curp,omitempty"` // 1.0
	NumRegIdTrib string     `json:"num_reg_id_trib,omitempty"`
	Domicilio    *Domicilio `json:"domicilio,omitempty"`
}

// ComercioExteriorDestinatario puede tener varios domicilios (en 1.0, uno
// solo y un solo Destinatario).
type ComercioExteriorDestinatario struct {
	NumRegIdTrib string      `json:"num_reg_id_trib,omitempty"`
	Rfc          string      `json:"rfc,omitempty"`  // 1.0
	Curp         string      `json:"curp,omitempty"` // 1.0
	Nombre       string      `json:"nombre,omitempty"`
	Domicilio    []Domicilio `json:"domicilio,omitempty"`
}

// ComercioExteriorMercancia no es la Mercancia de CartaPorte: aquí lo
// central es la fracción arancelaria y el valor en dólares.
type ComercioExteriorMercancia struct {
	NoIdentificacion         string                     `json:"no_identificacion,omitempty"`
	FraccionArancelaria      string                     `json:"fraccion_arancelaria,omitempty"`
	CantidadAduana           *decimal.Decimal           `json:"cantidad_aduana,omitempty"`
	UnidadAduana             string                     `json:"unidad_aduana,omitempty"`
	ValorUnitarioAduana      *decimal.Decimal           `json:"valor_unitario_aduana,omitempty"`
	ValorDolares             *decimal.Decimal           `json:"valor_dolares,omitempty"`
	DescripcionesEspecificas []DescripcionesEspecificas `json:"descripciones_especificas,omitempty"`
}

// DescripcionesEspecificas es el nodo DescripcionesEspecificas de una Mercancia del complemento ComercioExterior.
type DescripcionesEspecificas struct {
	Marca       string `json:"marca,omitempty"`
	Modelo      string `json:"modelo,omitempty"`
	SubModelo   string `json:"sub_modelo,omitempty"`
	NumeroSerie string `json:"numero_serie,omitempty"`
}

// Divisas 1.0: compra o venta de divisas.
type Divisas struct {
	Version       string `json:"version,omitempty"`
	TipoOperacion string `json:"tipo_operacion,omitempty"` // atributo "tipoOperacion": compra | venta
}

// CFDIRegistroFiscal 1.0 (registrofiscal).
type CFDIRegistroFiscal struct {
	Version string `json:"version,omitempty"`
	Folio   string `json:"folio,omitempty"`
}

// EstadoDeCuentaCombustible 1.2 (ecc12) y 1.1 (namespace
// .../EstadoDeCuentaCombustible). Conceptos aplana
// Conceptos/ConceptoEstadoDeCuentaCombustible.
type EstadoDeCuentaCombustible struct {
	Version        string                              `json:"version,omitempty"`
	TipoOperacion  string                              `json:"tipo_operacion,omitempty"`
	NumeroDeCuenta string                              `json:"numero_de_cuenta,omitempty"`
	SubTotal       *decimal.Decimal                    `json:"sub_total,omitempty"`
	Total          *decimal.Decimal                    `json:"total,omitempty"`
	Conceptos      []ConceptoEstadoDeCuentaCombustible `json:"conceptos,omitempty"`
}

// ConceptoEstadoDeCuentaCombustible es el nodo ConceptoEstadoDeCuentaCombustible del complemento EstadoDeCuentaCombustible.
type ConceptoEstadoDeCuentaCombustible struct {
	Identificador     string           `json:"identificador,omitempty"`
	Fecha             Fecha            `json:"fecha,omitzero"`
	Rfc               string           `json:"rfc,omitempty"`
	ClaveEstacion     string           `json:"clave_estacion,omitempty"`
	TAR               string           `json:"tar,omitempty"` // 1.1
	Cantidad          *decimal.Decimal `json:"cantidad,omitempty"`
	TipoCombustible   string           `json:"tipo_combustible,omitempty"`  // 1.2
	NoIdentificacion  string           `json:"no_identificacion,omitempty"` // 1.1
	Unidad            string           `json:"unidad,omitempty"`
	NombreCombustible string           `json:"nombre_combustible,omitempty"`
	FolioOperacion    string           `json:"folio_operacion,omitempty"`
	ValorUnitario     *decimal.Decimal `json:"valor_unitario,omitempty"`
	Importe           *decimal.Decimal `json:"importe,omitempty"`
	// Traslados usa el mismo struct que los del CFDI: Impuesto ("IVA" o
	// "IEPS", no la clave 002/003), TasaOCuota (TasaoCuota en 1.1) e
	// Importe; no hay Base ni TipoFactor.
	Traslados []Impuesto `json:"traslados,omitempty"`
}

// LeyendasFiscales 1.0.
type LeyendasFiscales struct {
	Version string    `json:"version,omitempty"`
	Leyenda []Leyenda `json:"leyenda,omitempty"`
}

// Leyenda es el nodo leyendasFisc:Leyenda del complemento LeyendasFiscales.
type Leyenda struct {
	DisposicionFiscal string `json:"disposicion_fiscal,omitempty"`
	Norma             string `json:"norma,omitempty"`
	TextoLeyenda      string `json:"texto_leyenda,omitempty"`
}

// Aerolineas 1.0: TUA y otros cargos del boleto de avión.
type Aerolineas struct {
	Version     string                 `json:"version,omitempty"`
	TUA         *decimal.Decimal       `json:"tua,omitempty"`
	OtrosCargos *AerolineasOtrosCargos `json:"otros_cargos,omitempty"`
}

// AerolineasOtrosCargos es el nodo aerolineas:OtrosCargos del complemento Aerolíneas.
type AerolineasOtrosCargos struct {
	TotalCargos *decimal.Decimal  `json:"total_cargos,omitempty"`
	Cargo       []AerolineasCargo `json:"cargo,omitempty"`
}

// AerolineasCargo es el nodo aerolineas:Cargo de OtrosCargos del complemento Aerolíneas.
type AerolineasCargo struct {
	CodigoCargo string           `json:"codigo_cargo,omitempty"`
	Importe     *decimal.Decimal `json:"importe,omitempty"`
}

// Donatarias 1.1 (donat11) y 1.0 (donat): comparten namespace y atributos;
// 1.1 solo cambió la versión fija. Los atributos van en minúsculas.
type Donatarias struct {
	Version           string `json:"version,omitempty"`
	NoAutorizacion    string `json:"no_autorizacion,omitempty"`
	FechaAutorizacion Fecha  `json:"fecha_autorizacion,omitzero"` // AAAA-MM-DD
	Leyenda           string `json:"leyenda,omitempty"`
}

// Complementos de concepto (Concepto/ComplementoConcepto).

// PorCuentadeTerceros 1.1 (terceros), de CFDI 3.3: el concepto se cobra por
// cuenta de un tercero. En 4.0 lo sustituye el nodo ACuentaTerceros.
type PorCuentadeTerceros struct {
	Version                  string                          `json:"version,omitempty"`
	Rfc                      string                          `json:"rfc,omitempty"`
	Nombre                   string                          `json:"nombre,omitempty"`
	InformacionFiscalTercero *UbicacionFiscal                `json:"informacion_fiscal_tercero,omitempty"`
	InformacionAduanera      *InformacionAduaneraComplemento `json:"informacion_aduanera,omitempty"`
	Parte                    []ParteComplemento              `json:"parte,omitempty"`
	CuentaPredial            *CuentaPredial                  `json:"cuenta_predial,omitempty"`
	Impuestos                *ImpuestosTerceros              `json:"impuestos,omitempty"`
}

// UbicacionFiscal es el t_UbicacionFiscal de terceros (atributos en
// minúsculas: calle, noExterior…).
type UbicacionFiscal struct {
	Calle        string `json:"calle,omitempty"`
	NoExterior   string `json:"no_exterior,omitempty"`
	NoInterior   string `json:"no_interior,omitempty"`
	Colonia      string `json:"colonia,omitempty"`
	Localidad    string `json:"localidad,omitempty"`
	Referencia   string `json:"referencia,omitempty"`
	Municipio    string `json:"municipio,omitempty"`
	Estado       string `json:"estado,omitempty"`
	Pais         string `json:"pais,omitempty"`
	CodigoPostal string `json:"codigo_postal,omitempty"`
}

// InformacionAduaneraComplemento es el t_InformacionAduanera de terceros y
// ventavehiculos: pedimento con fecha y aduana. No es la InformacionAduanera
// del concepto (solo NumeroPedimento).
type InformacionAduaneraComplemento struct {
	Numero string `json:"numero,omitempty"`
	Fecha  Fecha  `json:"fecha,omitzero"`
	Aduana string `json:"aduana,omitempty"`
}

// ParteComplemento es la Parte de terceros y ventavehiculos.
type ParteComplemento struct {
	Cantidad            *decimal.Decimal                 `json:"cantidad,omitempty"`
	Unidad              string                           `json:"unidad,omitempty"`
	NoIdentificacion    string                           `json:"no_identificacion,omitempty"`
	Descripcion         string                           `json:"descripcion,omitempty"`
	ValorUnitario       *decimal.Decimal                 `json:"valor_unitario,omitempty"`
	Importe             *decimal.Decimal                 `json:"importe,omitempty"`
	InformacionAduanera []InformacionAduaneraComplemento `json:"informacion_aduanera,omitempty"`
}

// ImpuestosTerceros es el nodo terceros:Impuestos del complemento de concepto PorCuentadeTerceros.
type ImpuestosTerceros struct {
	Retenciones []ImpuestoTerceros `json:"retenciones,omitempty"`
	Traslados   []ImpuestoTerceros `json:"traslados,omitempty"`
}

// ImpuestoTerceros: Impuesto va por nombre (ISR, IVA, IEPS); las retenciones
// no llevan Tasa.
type ImpuestoTerceros struct {
	Impuesto string           `json:"impuesto,omitempty"`
	Tasa     *decimal.Decimal `json:"tasa,omitempty"`
	Importe  *decimal.Decimal `json:"importe,omitempty"`
}

// VentaVehiculos 1.0 y 1.1 (ventavehiculos). 1.1 agregó Niv.
type VentaVehiculos struct {
	Version             string                           `json:"version,omitempty"`
	ClaveVehicular      string                           `json:"clave_vehicular,omitempty"`
	Niv                 string                           `json:"niv,omitempty"`
	InformacionAduanera []InformacionAduaneraComplemento `json:"informacion_aduanera,omitempty"`
	Parte               []ParteComplemento               `json:"parte,omitempty"`
}

// HidroYPetro 1.0 (hidrocarburospetroliferos): permiso de la operación con
// hidrocarburos o petrolíferos. El XSD lo define como complemento de
// concepto; si viene en el Complemento del comprobante también se lee
// (Comprobante.HidroYPetro).
type HidroYPetro struct {
	Version        string `json:"version,omitempty"`
	TipoPermiso    string `json:"tipo_permiso,omitempty"`
	NumeroPermiso  string `json:"numero_permiso,omitempty"`
	ClaveHYP       string `json:"clave_hyp,omitempty"`
	SubProductoHYP string `json:"sub_producto_hyp,omitempty"`
}

// InstEducativas 1.0 (iedu): colegiatura de una institución educativa
// privada. El nodo del SAT es instEducativas, en minúscula.
type InstEducativas struct {
	Version        string `json:"version,omitempty"`
	NombreAlumno   string `json:"nombre_alumno,omitempty"`
	CURP           string `json:"curp,omitempty"`
	NivelEducativo string `json:"nivel_educativo,omitempty"`
	AutRVOE        string `json:"aut_rvoe,omitempty"`
	RfcPago        string `json:"rfc_pago,omitempty"`
}
