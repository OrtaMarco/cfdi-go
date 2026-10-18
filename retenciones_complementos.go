package cfdi

import "github.com/OrtaMarco/cfdi-go/decimal"

// Complementos de Retenciones. Todos son chicos y se modelan completos,
// con los nombres de su XSD.

// ServiciosPlataformasTecnologicas 1.0: servicios de plataformas
// tecnológicas de intermediación.
type ServiciosPlataformasTecnologicas struct {
	Version                           string                `json:"version,omitempty"`
	Periodicidad                      string                `json:"periodicidad,omitempty"`
	NumServ                           string                `json:"num_serv,omitempty"`
	MonTotServSIVA                    *decimal.Decimal      `json:"mon_tot_serv_siva,omitempty"`
	TotalIVATrasladado                *decimal.Decimal      `json:"total_iva_trasladado,omitempty"`
	TotalIVARetenido                  *decimal.Decimal      `json:"total_iva_retenido,omitempty"`
	TotalISRRetenido                  *decimal.Decimal      `json:"total_isr_retenido,omitempty"`
	DifIVAEntregadoPrestServ          *decimal.Decimal      `json:"dif_iva_entregado_prest_serv,omitempty"`
	MonTotalporUsoPlataforma          *decimal.Decimal      `json:"mon_totalpor_uso_plataforma,omitempty"`
	MonTotalContribucionGubernamental *decimal.Decimal      `json:"mon_total_contribucion_gubernamental,omitempty"`
	Servicios                         []DetallesDelServicio `json:"servicios,omitempty"` // Servicios/DetallesDelServicio
}

// DetallesDelServicio es el nodo plataformasTecnologicas:DetallesDelServicio del complemento de Retenciones ServiciosPlataformasTecnologicas.
type DetallesDelServicio struct {
	FormaPagoServ                   string                           `json:"forma_pago_serv,omitempty"`
	TipoDeServ                      string                           `json:"tipo_de_serv,omitempty"`
	SubTipServ                      string                           `json:"sub_tip_serv,omitempty"`
	RFCTerceroAutorizado            string                           `json:"rfc_tercero_autorizado,omitempty"`
	FechaServ                       Fecha                            `json:"fecha_serv,omitzero"`
	PrecioServSinIVA                *decimal.Decimal                 `json:"precio_serv_sin_iva,omitempty"`
	ImpuestosTrasladadosdelServicio *ImpuestosTrasladadosdelServicio `json:"impuestos_trasladadosdel_servicio,omitempty"`
	ContribucionGubernamental       *ContribucionGubernamental       `json:"contribucion_gubernamental,omitempty"`
	ComisionDelServicio             *ComisionDelServicio             `json:"comision_del_servicio,omitempty"`
}

// ImpuestosTrasladadosdelServicio es el nodo plataformasTecnologicas:ImpuestosTrasladadosdelServicio de un DetallesDelServicio.
type ImpuestosTrasladadosdelServicio struct {
	Base       *decimal.Decimal `json:"base,omitempty"`
	Impuesto   string           `json:"impuesto,omitempty"`
	TipoFactor string           `json:"tipo_factor,omitempty"`
	TasaCuota  *decimal.Decimal `json:"tasa_cuota,omitempty"`
	Importe    *decimal.Decimal `json:"importe,omitempty"`
}

// ContribucionGubernamental es el nodo plataformasTecnologicas:ContribucionGubernamental de un DetallesDelServicio.
type ContribucionGubernamental struct {
	ImpContrib                     *decimal.Decimal `json:"imp_contrib,omitempty"`
	EntidadDondePagaLaContribucion string           `json:"entidad_donde_paga_la_contribucion,omitempty"`
}

// ComisionDelServicio es el nodo plataformasTecnologicas:ComisionDelServicio de un DetallesDelServicio.
type ComisionDelServicio struct {
	Base       *decimal.Decimal `json:"base,omitempty"`
	Porcentaje *decimal.Decimal `json:"porcentaje,omitempty"`
	Importe    *decimal.Decimal `json:"importe,omitempty"`
}

// Intereses es el complemento de Retenciones intereses:Intereses.
type Intereses struct {
	Version           string           `json:"version,omitempty"`
	SistFinanciero    string           `json:"sist_financiero,omitempty"`
	RetiroAORESRetInt string           `json:"retiro_aores_ret_int,omitempty"`
	OperFinancDerivad string           `json:"oper_financ_derivad,omitempty"`
	MontIntNominal    *decimal.Decimal `json:"mont_int_nominal,omitempty"`
	MontIntReal       *decimal.Decimal `json:"mont_int_real,omitempty"`
	Perdida           *decimal.Decimal `json:"perdida,omitempty"`
}

// Pagosaextranjeros es el complemento de Retenciones pagosaextranjeros:Pagosaextranjeros.
type Pagosaextranjeros struct {
	Version              string                    `json:"version,omitempty"`
	EsBenefEfectDelCobro string                    `json:"es_benef_efect_del_cobro,omitempty"`
	NoBeneficiario       *NoBeneficiarioExtranjero `json:"no_beneficiario,omitempty"`
	Beneficiario         *BeneficiarioExtranjero   `json:"beneficiario,omitempty"`
}

// NoBeneficiarioExtranjero es el nodo pagosaextranjeros:NoBeneficiario del complemento Pagosaextranjeros.
type NoBeneficiarioExtranjero struct {
	PaisDeResidParaEfecFisc string `json:"pais_de_resid_para_efec_fisc,omitempty"`
	ConceptoPago            string `json:"concepto_pago,omitempty"`
	DescripcionConcepto     string `json:"descripcion_concepto,omitempty"`
}

// BeneficiarioExtranjero es el nodo pagosaextranjeros:Beneficiario del complemento Pagosaextranjeros.
type BeneficiarioExtranjero struct {
	RFC                 string `json:"rfc,omitempty"`
	CURP                string `json:"curp,omitempty"`
	NomDenRazSocB       string `json:"nom_den_raz_soc_b,omitempty"`
	ConceptoPago        string `json:"concepto_pago,omitempty"`
	DescripcionConcepto string `json:"descripcion_concepto,omitempty"`
}

// Dividendos es el complemento de Retenciones dividendos:Dividendos.
type Dividendos struct {
	Version    string      `json:"version,omitempty"`
	DividOUtil *DividOUtil `json:"divid_o_util,omitempty"`
	Remanente  *Remanente  `json:"remanente,omitempty"`
}

// DividOUtil es el nodo dividendos:DividOUtil del complemento Dividendos.
type DividOUtil struct {
	CveTipDivOUtil            string           `json:"cve_tip_div_o_util,omitempty"`
	MontISRAcredRetMexico     *decimal.Decimal `json:"mont_isr_acred_ret_mexico,omitempty"`
	MontISRAcredRetExtranjero *decimal.Decimal `json:"mont_isr_acred_ret_extranjero,omitempty"`
	MontRetExtDivExt          *decimal.Decimal `json:"mont_ret_ext_div_ext,omitempty"`
	TipoSocDistrDiv           string           `json:"tipo_soc_distr_div,omitempty"`
	MontISRAcredNal           *decimal.Decimal `json:"mont_isr_acred_nal,omitempty"`
	MontDivAcumNal            *decimal.Decimal `json:"mont_div_acum_nal,omitempty"`
	MontDivAcumExt            *decimal.Decimal `json:"mont_div_acum_ext,omitempty"`
}

// Remanente es el nodo dividendos:Remanente del complemento Dividendos.
type Remanente struct {
	ProporcionRem *decimal.Decimal `json:"proporcion_rem,omitempty"`
}

// EnajenaciondeAcciones es el complemento de Retenciones enajenaciondeacciones:EnajenaciondeAcciones.
type EnajenaciondeAcciones struct {
	Version                string           `json:"version,omitempty"`
	ContratoIntermediacion string           `json:"contrato_intermediacion,omitempty"`
	Ganancia               *decimal.Decimal `json:"ganancia,omitempty"`
	Perdida                *decimal.Decimal `json:"perdida,omitempty"`
}

// Arrendamientoenfideicomiso es el complemento de Retenciones arrendamientoenfideicomiso:Arrendamientoenfideicomiso.
type Arrendamientoenfideicomiso struct {
	Version                    string           `json:"version,omitempty"`
	PagProvEfecPorFiduc        *decimal.Decimal `json:"pag_prov_efec_por_fiduc,omitempty"`
	RendimFideicom             *decimal.Decimal `json:"rendim_fideicom,omitempty"`
	DeduccCorresp              *decimal.Decimal `json:"deducc_corresp,omitempty"`
	MontTotRet                 *decimal.Decimal `json:"mont_tot_ret,omitempty"`
	MontResFiscDistFibras      *decimal.Decimal `json:"mont_res_fisc_dist_fibras,omitempty"`
	MontOtrosConceptDistr      *decimal.Decimal `json:"mont_otros_concept_distr,omitempty"`
	DescrMontOtrosConceptDistr string           `json:"descr_mont_otros_concept_distr,omitempty"`
}

// Operacionesconderivados es el complemento de Retenciones operacionesconderivados:Operacionesconderivados.
type Operacionesconderivados struct {
	Version     string           `json:"version,omitempty"`
	MontGanAcum *decimal.Decimal `json:"mont_gan_acum,omitempty"`
	MontPerdDed *decimal.Decimal `json:"mont_perd_ded,omitempty"`
}

// Intereseshipotecarios es el complemento de Retenciones intereseshipotecarios:Intereseshipotecarios.
type Intereseshipotecarios struct {
	Version                    string           `json:"version,omitempty"`
	CreditoDeInstFinanc        string           `json:"credito_de_inst_financ,omitempty"`
	SaldoInsoluto              *decimal.Decimal `json:"saldo_insoluto,omitempty"`
	PropDeducDelCredit         *decimal.Decimal `json:"prop_deduc_del_credit,omitempty"`
	MontTotIntNominalesDev     *decimal.Decimal `json:"mont_tot_int_nominales_dev,omitempty"`
	MontTotIntNominalesDevYPag *decimal.Decimal `json:"mont_tot_int_nominales_dev_y_pag,omitempty"`
	MontTotIntRealPagDeduc     *decimal.Decimal `json:"mont_tot_int_real_pag_deduc,omitempty"`
	NumContrato                string           `json:"num_contrato,omitempty"`
}

// Fideicomisonoempresarial es el complemento de Retenciones fideicomisonoempresarial:Fideicomisonoempresarial.
type Fideicomisonoempresarial struct {
	Version             string               `json:"version,omitempty"`
	IngresosOEntradas   *IngresosOEntradas   `json:"ingresos_o_entradas,omitempty"`
	DeduccOSalidas      *DeduccOSalidas      `json:"deducc_o_salidas,omitempty"`
	RetEfectFideicomiso *RetEfectFideicomiso `json:"ret_efect_fideicomiso,omitempty"`
}

// IngresosOEntradas es el nodo fideicomisonoempresarial:IngresosOEntradas del complemento Fideicomisonoempresarial.
type IngresosOEntradas struct {
	MontTotEntradasPeriodo  *decimal.Decimal `json:"mont_tot_entradas_periodo,omitempty"`
	PartPropAcumDelFideicom *decimal.Decimal `json:"part_prop_acum_del_fideicom,omitempty"`
	PropDelMontTot          *decimal.Decimal `json:"prop_del_mont_tot,omitempty"`
	// IntegracIngresos/@Concepto.
	IntegracIngresos string `json:"integrac_ingresos,omitempty"`
}

// DeduccOSalidas es el nodo fideicomisonoempresarial:DeduccOSalidas del complemento Fideicomisonoempresarial.
type DeduccOSalidas struct {
	MontTotEgresPeriodo *decimal.Decimal `json:"mont_tot_egres_periodo,omitempty"`
	PartPropDelFideicom *decimal.Decimal `json:"part_prop_del_fideicom,omitempty"`
	PropDelMontTot      *decimal.Decimal `json:"prop_del_mont_tot,omitempty"`
	// IntegracEgresos/@ConceptoS.
	IntegracEgresos string `json:"integrac_egresos,omitempty"`
}

// RetEfectFideicomiso es el nodo fideicomisonoempresarial:RetEfectFideicomiso del complemento Fideicomisonoempresarial.
type RetEfectFideicomiso struct {
	MontRetRelPagFideic *decimal.Decimal `json:"mont_ret_rel_pag_fideic,omitempty"`
	DescRetRelPagFideic string           `json:"desc_ret_rel_pag_fideic,omitempty"`
}

// Planesderetiro 1.0 y 1.1 en un objeto con version: la 1.1 agrega
// AportacionesODepositos y NumReferencia, que vienen vacíos en la 1.0.
type Planesderetiro struct {
	Version                        string                   `json:"version,omitempty"`
	SistemaFinanc                  string                   `json:"sistema_financ,omitempty"`
	MontTotAportAnioInmAnterior    *decimal.Decimal         `json:"mont_tot_aport_anio_inm_anterior,omitempty"`
	MontIntRealesDevengAniooInmAnt *decimal.Decimal         `json:"mont_int_reales_deveng_anioo_inm_ant,omitempty"`
	HuboRetirosAnioInmAntPer       string                   `json:"hubo_retiros_anio_inm_ant_per,omitempty"`
	MontTotRetiradoAnioInmAntPer   *decimal.Decimal         `json:"mont_tot_retirado_anio_inm_ant_per,omitempty"`
	MontTotExentRetiradoAnioInmAnt *decimal.Decimal         `json:"mont_tot_exent_retirado_anio_inm_ant,omitempty"`
	MontTotExedenteAnioInmAnt      *decimal.Decimal         `json:"mont_tot_exedente_anio_inm_ant,omitempty"`
	HuboRetirosAnioInmAnt          string                   `json:"hubo_retiros_anio_inm_ant,omitempty"`
	MontTotRetiradoAnioInmAnt      *decimal.Decimal         `json:"mont_tot_retirado_anio_inm_ant,omitempty"`
	NumReferencia                  string                   `json:"num_referencia,omitempty"`
	AportacionesODepositos         []AportacionesODepositos `json:"aportaciones_o_depositos,omitempty"`
}

// AportacionesODepositos es el nodo planesderetiro11:AportacionesODepositos del complemento de Retenciones PlanesDeRetiro 1.1.
type AportacionesODepositos struct {
	TipoAportacionODeposito string           `json:"tipo_aportacion_o_deposito,omitempty"`
	MontAportODep           *decimal.Decimal `json:"mont_aport_o_dep,omitempty"`
	RFCFiduciaria           string           `json:"rfc_fiduciaria,omitempty"`
}

// Premios es el complemento de Retenciones premios:Premios.
type Premios struct {
	Version           string           `json:"version,omitempty"`
	EntidadFederativa string           `json:"entidad_federativa,omitempty"`
	MontTotPago       *decimal.Decimal `json:"mont_tot_pago,omitempty"`
	MontTotPagoGrav   *decimal.Decimal `json:"mont_tot_pago_grav,omitempty"`
	MontTotPagoExent  *decimal.Decimal `json:"mont_tot_pago_exent,omitempty"`
}

// SectorFinanciero es el complemento de Retenciones sectorfinanciero:SectorFinanciero.
type SectorFinanciero struct {
	Version         string `json:"version,omitempty"`
	IdFideicom      string `json:"id_fideicom,omitempty"`
	NomFideicom     string `json:"nom_fideicom,omitempty"`
	DescripFideicom string `json:"descrip_fideicom,omitempty"`
}

// complementoRetenciones despacha un hijo de retenciones:Complemento por
// namespace + nombre local; solo cuenta el primero de cada uno.
func (l *lector) complementoRetenciones(rt *Retenciones, h *nodo, rc string) {
	r := rc + "/" + h.local
	switch {
	case h.ns == NSTimbre && h.local == "TimbreFiscalDigital":
		if !l.yaLeido(rt.TimbreFiscalDigital != nil, r, h.local) {
			rt.TimbreFiscalDigital = l.timbre(h, r)
		}
	case h.ns == NSRetPlataformasTecnologicas10 && h.local == "ServiciosPlataformasTecnologicas":
		if !l.yaLeido(rt.ServiciosPlataformasTecnologicas != nil, r, h.local) {
			rt.ServiciosPlataformasTecnologicas = l.plataformas(h, r)
		}
	case h.ns == NSRetIntereses && h.local == "Intereses":
		if !l.yaLeido(rt.Intereses != nil, r, h.local) {
			rt.Intereses = &Intereses{
				Version:           l.versionComplemento(h, r, "1.0"),
				SistFinanciero:    l.str(h, "SistFinanciero"),
				RetiroAORESRetInt: l.str(h, "RetiroAORESRetInt"),
				OperFinancDerivad: l.str(h, "OperFinancDerivad"),
				MontIntNominal:    l.dec(h, r, "MontIntNominal"),
				MontIntReal:       l.dec(h, r, "MontIntReal"),
				Perdida:           l.dec(h, r, "Perdida"),
			}
			l.requerido(h, r, "SistFinanciero", "RetiroAORESRetInt", "OperFinancDerivad", "MontIntNominal", "MontIntReal", "Perdida")
		}
	case h.ns == NSRetPagosAExtranjeros && h.local == "Pagosaextranjeros":
		if !l.yaLeido(rt.Pagosaextranjeros != nil, r, h.local) {
			rt.Pagosaextranjeros = l.pagosAExtranjeros(h, r)
		}
	case h.ns == NSRetDividendos && h.local == "Dividendos":
		if !l.yaLeido(rt.Dividendos != nil, r, h.local) {
			rt.Dividendos = l.dividendos(h, r)
		}
	case h.ns == NSRetEnajenacionDeAcciones && h.local == "EnajenaciondeAcciones":
		if !l.yaLeido(rt.EnajenaciondeAcciones != nil, r, h.local) {
			rt.EnajenaciondeAcciones = &EnajenaciondeAcciones{
				Version:                l.versionComplemento(h, r, "1.0"),
				ContratoIntermediacion: l.str(h, "ContratoIntermediacion"),
				Ganancia:               l.dec(h, r, "Ganancia"),
				Perdida:                l.dec(h, r, "Perdida"),
			}
			l.requerido(h, r, "ContratoIntermediacion", "Ganancia", "Perdida")
		}
	case h.ns == NSRetArrendamientoEnFideicomiso && h.local == "Arrendamientoenfideicomiso":
		if !l.yaLeido(rt.Arrendamientoenfideicomiso != nil, r, h.local) {
			rt.Arrendamientoenfideicomiso = &Arrendamientoenfideicomiso{
				Version:                    l.versionComplemento(h, r, "1.0"),
				PagProvEfecPorFiduc:        l.dec(h, r, "PagProvEfecPorFiduc"),
				RendimFideicom:             l.dec(h, r, "RendimFideicom"),
				DeduccCorresp:              l.dec(h, r, "DeduccCorresp"),
				MontTotRet:                 l.dec(h, r, "MontTotRet"),
				MontResFiscDistFibras:      l.dec(h, r, "MontResFiscDistFibras"),
				MontOtrosConceptDistr:      l.dec(h, r, "MontOtrosConceptDistr"),
				DescrMontOtrosConceptDistr: l.str(h, "DescrMontOtrosConceptDistr"),
			}
			l.requerido(h, r, "PagProvEfecPorFiduc", "RendimFideicom", "DeduccCorresp")
		}
	case h.ns == NSRetOperacionesConDerivados && h.local == "Operacionesconderivados":
		if !l.yaLeido(rt.Operacionesconderivados != nil, r, h.local) {
			rt.Operacionesconderivados = &Operacionesconderivados{
				Version:     l.versionComplemento(h, r, "1.0"),
				MontGanAcum: l.dec(h, r, "MontGanAcum"),
				MontPerdDed: l.dec(h, r, "MontPerdDed"),
			}
			l.requerido(h, r, "MontGanAcum", "MontPerdDed")
		}
	case h.ns == NSRetInteresesHipotecarios && h.local == "Intereseshipotecarios":
		if !l.yaLeido(rt.Intereseshipotecarios != nil, r, h.local) {
			rt.Intereseshipotecarios = &Intereseshipotecarios{
				Version:                    l.versionComplemento(h, r, "1.0"),
				CreditoDeInstFinanc:        l.str(h, "CreditoDeInstFinanc"),
				SaldoInsoluto:              l.dec(h, r, "SaldoInsoluto"),
				PropDeducDelCredit:         l.dec(h, r, "PropDeducDelCredit"),
				MontTotIntNominalesDev:     l.dec(h, r, "MontTotIntNominalesDev"),
				MontTotIntNominalesDevYPag: l.dec(h, r, "MontTotIntNominalesDevYPag"),
				MontTotIntRealPagDeduc:     l.dec(h, r, "MontTotIntRealPagDeduc"),
				NumContrato:                l.str(h, "NumContrato"),
			}
			l.requerido(h, r, "CreditoDeInstFinanc", "SaldoInsoluto")
		}
	case h.ns == NSRetFideicomisoNoEmpresarial && h.local == "Fideicomisonoempresarial":
		if !l.yaLeido(rt.Fideicomisonoempresarial != nil, r, h.local) {
			rt.Fideicomisonoempresarial = l.fideicomiso(h, r)
		}
	case (h.ns == NSRetPlanesDeRetiro10 || h.ns == NSRetPlanesDeRetiro11) && h.local == "Planesderetiro":
		if !l.yaLeido(rt.Planesderetiro != nil, r, h.local) {
			rt.Planesderetiro = l.planesDeRetiro(h, r)
		}
	case h.ns == NSRetPremios && h.local == "Premios":
		if !l.yaLeido(rt.Premios != nil, r, h.local) {
			rt.Premios = &Premios{
				Version:           l.versionComplemento(h, r, "1.0"),
				EntidadFederativa: l.str(h, "EntidadFederativa"),
				MontTotPago:       l.dec(h, r, "MontTotPago"),
				MontTotPagoGrav:   l.dec(h, r, "MontTotPagoGrav"),
				MontTotPagoExent:  l.dec(h, r, "MontTotPagoExent"),
			}
			l.requerido(h, r, "EntidadFederativa", "MontTotPago", "MontTotPagoGrav", "MontTotPagoExent")
		}
	case h.ns == NSRetSectorFinanciero && h.local == "SectorFinanciero":
		if !l.yaLeido(rt.SectorFinanciero != nil, r, h.local) {
			rt.SectorFinanciero = &SectorFinanciero{
				Version:         l.versionComplemento(h, r, "1.0"),
				IdFideicom:      l.str(h, "IdFideicom"),
				NomFideicom:     l.str(h, "NomFideicom"),
				DescripFideicom: l.str(h, "DescripFideicom"),
			}
			l.requerido(h, r, "IdFideicom", "DescripFideicom")
		}
	default:
		nom := h.nombre()
		agregarUnico(&rt.ComplementosNoSoportados, nom)
		l.advertir(AdvComplementoNoSoportado, r, "complemento no soportado: "+nom)
	}
}

func (l *lector) plataformas(n *nodo, r string) *ServiciosPlataformasTecnologicas {
	ns := n.ns
	p := &ServiciosPlataformasTecnologicas{
		Version:                           l.versionComplemento(n, r, "1.0"),
		Periodicidad:                      l.str(n, "Periodicidad"),
		NumServ:                           l.str(n, "NumServ"),
		MonTotServSIVA:                    l.dec(n, r, "MonTotServSIVA"),
		TotalIVATrasladado:                l.dec(n, r, "TotalIVATrasladado"),
		TotalIVARetenido:                  l.dec(n, r, "TotalIVARetenido"),
		TotalISRRetenido:                  l.dec(n, r, "TotalISRRetenido"),
		DifIVAEntregadoPrestServ:          l.dec(n, r, "DifIVAEntregadoPrestServ"),
		MonTotalporUsoPlataforma:          l.dec(n, r, "MonTotalporUsoPlataforma"),
		MonTotalContribucionGubernamental: l.dec(n, r, "MonTotalContribucionGubernamental"),
	}
	l.requerido(n, r, "Periodicidad", "NumServ", "MonTotServSIVA", "TotalIVATrasladado", "TotalIVARetenido",
		"TotalISRRetenido", "DifIVAEntregadoPrestServ", "MonTotalporUsoPlataforma")
	s := l.unico(n, ns, "Servicios", r)
	if s == nil {
		l.advertir(AdvNodoFaltante, r+"/Servicios", "ServiciosPlataformasTecnologicas sin Servicios")
		return p
	}
	rs := r + "/Servicios"
	for i, d := range l.hijos(s, ns, "DetallesDelServicio", rs) {
		rd := idx(rs, "DetallesDelServicio", i)
		ds := DetallesDelServicio{
			FormaPagoServ:        l.str(d, "FormaPagoServ"),
			TipoDeServ:           l.str(d, "TipoDeServ"),
			SubTipServ:           l.str(d, "SubTipServ"),
			RFCTerceroAutorizado: rfc(l.str(d, "RFCTerceroAutorizado")),
			FechaServ:            l.fecha(d, rd, "FechaServ", true),
			PrecioServSinIVA:     l.dec(d, rd, "PrecioServSinIVA"),
		}
		l.requerido(d, rd, "FormaPagoServ", "TipoDeServ", "FechaServ", "PrecioServSinIVA")
		if t := l.unico(d, ns, "ImpuestosTrasladadosdelServicio", rd); t != nil {
			rt := rd + "/ImpuestosTrasladadosdelServicio"
			ds.ImpuestosTrasladadosdelServicio = &ImpuestosTrasladadosdelServicio{
				Base:       l.dec(t, rt, "Base"),
				Impuesto:   l.str(t, "Impuesto"),
				TipoFactor: l.str(t, "TipoFactor"),
				TasaCuota:  l.dec(t, rt, "TasaCuota"),
				Importe:    l.dec(t, rt, "Importe"),
			}
			l.requerido(t, rt, "Base", "Impuesto", "TipoFactor", "TasaCuota", "Importe")
		}
		if c := l.unico(d, ns, "ContribucionGubernamental", rd); c != nil {
			rc := rd + "/ContribucionGubernamental"
			ds.ContribucionGubernamental = &ContribucionGubernamental{
				ImpContrib:                     l.dec(c, rc, "ImpContrib"),
				EntidadDondePagaLaContribucion: l.str(c, "EntidadDondePagaLaContribucion"),
			}
			l.requerido(c, rc, "ImpContrib", "EntidadDondePagaLaContribucion")
		}
		if c := l.unico(d, ns, "ComisionDelServicio", rd); c != nil {
			rc := rd + "/ComisionDelServicio"
			ds.ComisionDelServicio = &ComisionDelServicio{
				Base:       l.dec(c, rc, "Base"),
				Porcentaje: l.dec(c, rc, "Porcentaje"),
				Importe:    l.dec(c, rc, "Importe"),
			}
			l.requerido(c, rc, "Importe")
		}
		p.Servicios = append(p.Servicios, ds)
	}
	if len(p.Servicios) == 0 {
		l.advertir(AdvNodoFaltante, rs+"/DetallesDelServicio", "Servicios sin ningún DetallesDelServicio")
	}
	return p
}

func (l *lector) pagosAExtranjeros(n *nodo, r string) *Pagosaextranjeros {
	p := &Pagosaextranjeros{
		Version:              l.versionComplemento(n, r, "1.0"),
		EsBenefEfectDelCobro: l.str(n, "EsBenefEfectDelCobro"),
	}
	l.requerido(n, r, "EsBenefEfectDelCobro")
	if nb := l.unico(n, n.ns, "NoBeneficiario", r); nb != nil {
		p.NoBeneficiario = &NoBeneficiarioExtranjero{
			PaisDeResidParaEfecFisc: l.str(nb, "PaisDeResidParaEfecFisc"),
			ConceptoPago:            l.str(nb, "ConceptoPago"),
			DescripcionConcepto:     l.str(nb, "DescripcionConcepto"),
		}
		l.requerido(nb, r+"/NoBeneficiario", "PaisDeResidParaEfecFisc", "ConceptoPago", "DescripcionConcepto")
	}
	if b := l.unico(n, n.ns, "Beneficiario", r); b != nil {
		p.Beneficiario = &BeneficiarioExtranjero{
			RFC:                 rfc(l.str(b, "RFC")),
			CURP:                l.str(b, "CURP"),
			NomDenRazSocB:       l.str(b, "NomDenRazSocB"),
			ConceptoPago:        l.str(b, "ConceptoPago"),
			DescripcionConcepto: l.str(b, "DescripcionConcepto"),
		}
		l.requerido(b, r+"/Beneficiario", "RFC", "CURP", "NomDenRazSocB", "ConceptoPago", "DescripcionConcepto")
	}
	return p
}

func (l *lector) dividendos(n *nodo, r string) *Dividendos {
	d := &Dividendos{Version: l.versionComplemento(n, r, "1.0")}
	if du := l.unico(n, n.ns, "DividOUtil", r); du != nil {
		rd := r + "/DividOUtil"
		d.DividOUtil = &DividOUtil{
			CveTipDivOUtil:            l.str(du, "CveTipDivOUtil"),
			MontISRAcredRetMexico:     l.dec(du, rd, "MontISRAcredRetMexico"),
			MontISRAcredRetExtranjero: l.dec(du, rd, "MontISRAcredRetExtranjero"),
			MontRetExtDivExt:          l.dec(du, rd, "MontRetExtDivExt"),
			TipoSocDistrDiv:           l.str(du, "TipoSocDistrDiv"),
			MontISRAcredNal:           l.dec(du, rd, "MontISRAcredNal"),
			MontDivAcumNal:            l.dec(du, rd, "MontDivAcumNal"),
			MontDivAcumExt:            l.dec(du, rd, "MontDivAcumExt"),
		}
		l.requerido(du, rd, "CveTipDivOUtil", "MontISRAcredRetMexico", "MontISRAcredRetExtranjero", "TipoSocDistrDiv")
	}
	if re := l.unico(n, n.ns, "Remanente", r); re != nil {
		d.Remanente = &Remanente{ProporcionRem: l.dec(re, r+"/Remanente", "ProporcionRem")}
	}
	return d
}

func (l *lector) fideicomiso(n *nodo, r string) *Fideicomisonoempresarial {
	ns := n.ns
	f := &Fideicomisonoempresarial{Version: l.versionComplemento(n, r, "1.0")}
	if in := l.unico(n, ns, "IngresosOEntradas", r); in != nil {
		ri := r + "/IngresosOEntradas"
		f.IngresosOEntradas = &IngresosOEntradas{
			MontTotEntradasPeriodo:  l.dec(in, ri, "MontTotEntradasPeriodo"),
			PartPropAcumDelFideicom: l.dec(in, ri, "PartPropAcumDelFideicom"),
			PropDelMontTot:          l.dec(in, ri, "PropDelMontTot"),
		}
		l.requerido(in, ri, "MontTotEntradasPeriodo", "PartPropAcumDelFideicom", "PropDelMontTot")
		if ig := l.unico(in, ns, "IntegracIngresos", ri); ig != nil {
			f.IngresosOEntradas.IntegracIngresos = l.str(ig, "Concepto")
			l.requerido(ig, ri+"/IntegracIngresos", "Concepto")
		} else {
			l.advertir(AdvNodoFaltante, ri+"/IntegracIngresos", "falta el nodo IntegracIngresos")
		}
	} else {
		l.advertir(AdvNodoFaltante, r+"/IngresosOEntradas", "falta el nodo IngresosOEntradas")
	}
	if de := l.unico(n, ns, "DeduccOSalidas", r); de != nil {
		rd := r + "/DeduccOSalidas"
		f.DeduccOSalidas = &DeduccOSalidas{
			MontTotEgresPeriodo: l.dec(de, rd, "MontTotEgresPeriodo"),
			PartPropDelFideicom: l.dec(de, rd, "PartPropDelFideicom"),
			PropDelMontTot:      l.dec(de, rd, "PropDelMontTot"),
		}
		l.requerido(de, rd, "MontTotEgresPeriodo", "PartPropDelFideicom", "PropDelMontTot")
		if ie := l.unico(de, ns, "IntegracEgresos", rd); ie != nil {
			f.DeduccOSalidas.IntegracEgresos = l.str(ie, "ConceptoS")
			l.requerido(ie, rd+"/IntegracEgresos", "ConceptoS")
		} else {
			l.advertir(AdvNodoFaltante, rd+"/IntegracEgresos", "falta el nodo IntegracEgresos")
		}
	} else {
		l.advertir(AdvNodoFaltante, r+"/DeduccOSalidas", "falta el nodo DeduccOSalidas")
	}
	if re := l.unico(n, ns, "RetEfectFideicomiso", r); re != nil {
		rr := r + "/RetEfectFideicomiso"
		f.RetEfectFideicomiso = &RetEfectFideicomiso{
			MontRetRelPagFideic: l.dec(re, rr, "MontRetRelPagFideic"),
			DescRetRelPagFideic: l.str(re, "DescRetRelPagFideic"),
		}
		l.requerido(re, rr, "MontRetRelPagFideic", "DescRetRelPagFideic")
	} else {
		l.advertir(AdvNodoFaltante, r+"/RetEfectFideicomiso", "falta el nodo RetEfectFideicomiso")
	}
	return f
}

func (l *lector) planesDeRetiro(n *nodo, r string) *Planesderetiro {
	esperada := "1.0"
	if n.ns == NSRetPlanesDeRetiro11 {
		esperada = "1.1"
	}
	p := &Planesderetiro{
		Version:                        l.versionComplemento(n, r, esperada),
		SistemaFinanc:                  l.str(n, "SistemaFinanc"),
		MontTotAportAnioInmAnterior:    l.dec(n, r, "MontTotAportAnioInmAnterior"),
		MontIntRealesDevengAniooInmAnt: l.dec(n, r, "MontIntRealesDevengAniooInmAnt"),
		HuboRetirosAnioInmAntPer:       l.str(n, "HuboRetirosAnioInmAntPer"),
		MontTotRetiradoAnioInmAntPer:   l.dec(n, r, "MontTotRetiradoAnioInmAntPer"),
		MontTotExentRetiradoAnioInmAnt: l.dec(n, r, "MontTotExentRetiradoAnioInmAnt"),
		MontTotExedenteAnioInmAnt:      l.dec(n, r, "MontTotExedenteAnioInmAnt"),
		HuboRetirosAnioInmAnt:          l.str(n, "HuboRetirosAnioInmAnt"),
		MontTotRetiradoAnioInmAnt:      l.dec(n, r, "MontTotRetiradoAnioInmAnt"),
		NumReferencia:                  l.str(n, "NumReferencia"),
	}
	l.requerido(n, r, "SistemaFinanc", "MontIntRealesDevengAniooInmAnt", "HuboRetirosAnioInmAntPer", "HuboRetirosAnioInmAnt")
	for i, a := range l.hijos(n, n.ns, "AportacionesODepositos", r) {
		ra := idx(r, "AportacionesODepositos", i)
		p.AportacionesODepositos = append(p.AportacionesODepositos, AportacionesODepositos{
			TipoAportacionODeposito: l.str(a, "TipoAportacionODeposito"),
			MontAportODep:           l.dec(a, ra, "MontAportODep"),
			RFCFiduciaria:           rfc(l.str(a, "RFCFiduciaria")),
		})
		l.requerido(a, ra, "TipoAportacionODeposito", "MontAportODep")
	}
	return p
}
