package cfdi

import "github.com/OrtaMarco/cfdi-go/decimal"

// Complementos de uso común (T10). Siguen las convenciones de modelo.go:
// nombres del XSD del SAT, contenedores sin atributos aplanados, la versión
// la dice el namespace.

// INE 1.1 y 1.0 (ine): proceso electoral. Comparten namespace y nodos; 1.1
// solo admite más entidades.
type INE struct {
	Version        string       `json:"version,omitempty"`
	TipoProceso    string       `json:"tipo_proceso,omitempty"`
	TipoComite     string       `json:"tipo_comite,omitempty"`
	IdContabilidad string       `json:"id_contabilidad,omitempty"`
	Entidad        []INEEntidad `json:"entidad,omitempty"`
}

// INEEntidad es el nodo ine:Entidad del complemento INE.
type INEEntidad struct {
	ClaveEntidad string            `json:"clave_entidad,omitempty"`
	Ambito       string            `json:"ambito,omitempty"`
	Contabilidad []INEContabilidad `json:"contabilidad,omitempty"`
}

// INEContabilidad es el nodo ine:Contabilidad de una Entidad del complemento INE.
type INEContabilidad struct {
	IdContabilidad string `json:"id_contabilidad,omitempty"`
}

// ValesDeDespensa 1.0. Conceptos aplana Conceptos/Concepto.
type ValesDeDespensa struct {
	Version          string                    `json:"version,omitempty"`
	TipoOperacion    string                    `json:"tipo_operacion,omitempty"`
	RegistroPatronal string                    `json:"registro_patronal,omitempty"`
	NumeroDeCuenta   string                    `json:"numero_de_cuenta,omitempty"`
	Total            *decimal.Decimal          `json:"total,omitempty"`
	Conceptos        []ValesDeDespensaConcepto `json:"conceptos,omitempty"`
}

// ValesDeDespensaConcepto es el nodo valesdedespensa:Concepto del complemento ValesDeDespensa.
type ValesDeDespensaConcepto struct {
	Identificador      string           `json:"identificador,omitempty"`
	Fecha              Fecha            `json:"fecha,omitzero"`
	Rfc                string           `json:"rfc,omitempty"`
	Curp               string           `json:"curp,omitempty"`
	Nombre             string           `json:"nombre,omitempty"`
	NumSeguridadSocial string           `json:"num_seguridad_social,omitempty"`
	Importe            *decimal.Decimal `json:"importe,omitempty"`
}

// ConsumoDeCombustibles 1.1 (ConsumoDeCombustibles11) y 1.0
// (consumodecombustibles). 1.1 agregó tipoCombustible al concepto y
// renombró la tasa del Determinado (tasa → tasaOCuota): cada una en su campo.
type ConsumoDeCombustibles struct {
	Version        string                          `json:"version,omitempty"`
	TipoOperacion  string                          `json:"tipo_operacion,omitempty"`
	NumeroDeCuenta string                          `json:"numero_de_cuenta,omitempty"`
	SubTotal       *decimal.Decimal                `json:"sub_total,omitempty"`
	Total          *decimal.Decimal                `json:"total,omitempty"`
	Conceptos      []ConceptoConsumoDeCombustibles `json:"conceptos,omitempty"`
}

// ConceptoConsumoDeCombustibles es el nodo ConceptoConsumoDeCombustibles del complemento ConsumoDeCombustibles.
type ConceptoConsumoDeCombustibles struct {
	Identificador     string                             `json:"identificador,omitempty"`
	Fecha             Fecha                              `json:"fecha,omitzero"`
	Rfc               string                             `json:"rfc,omitempty"`
	ClaveEstacion     string                             `json:"clave_estacion,omitempty"`
	TipoCombustible   string                             `json:"tipo_combustible,omitempty"` // 1.1
	Cantidad          *decimal.Decimal                   `json:"cantidad,omitempty"`
	NombreCombustible string                             `json:"nombre_combustible,omitempty"`
	FolioOperacion    string                             `json:"folio_operacion,omitempty"`
	ValorUnitario     *decimal.Decimal                   `json:"valor_unitario,omitempty"`
	Importe           *decimal.Decimal                   `json:"importe,omitempty"`
	Determinados      []ConsumoDeCombustiblesDeterminado `json:"determinados,omitempty"`
}

// ConsumoDeCombustiblesDeterminado es el nodo Determinado de un concepto del complemento ConsumoDeCombustibles.
type ConsumoDeCombustiblesDeterminado struct {
	Impuesto   string           `json:"impuesto,omitempty"`
	Tasa       *decimal.Decimal `json:"tasa,omitempty"`         // 1.0
	TasaOCuota *decimal.Decimal `json:"tasa_o_cuota,omitempty"` // 1.1
	Importe    *decimal.Decimal `json:"importe,omitempty"`
}

// PagoEnEspecie 1.0: donación de obras de arte en pago en especie.
type PagoEnEspecie struct {
	Version      string `json:"version,omitempty"`
	CvePIC       string `json:"cve_pic,omitempty"`
	FolioSolDon  string `json:"folio_sol_don,omitempty"`
	PzaArtNombre string `json:"pza_art_nombre,omitempty"`
	PzaArtTecn   string `json:"pza_art_tecn,omitempty"`
	PzaArtAProd  string `json:"pza_art_a_prod,omitempty"`
	PzaArtDim    string `json:"pza_art_dim,omitempty"`
}

// NotariosPublicos 1.0: enajenación de inmuebles ante notario. DescInmuebles
// aplana DescInmuebles/DescInmueble.
type NotariosPublicos struct {
	Version          string                            `json:"version,omitempty"`
	DescInmuebles    []NotariosPublicosDescInmueble    `json:"desc_inmuebles,omitempty"`
	DatosOperacion   *NotariosPublicosDatosOperacion   `json:"datos_operacion,omitempty"`
	DatosNotario     *NotariosPublicosDatosNotario     `json:"datos_notario,omitempty"`
	DatosEnajenante  *NotariosPublicosDatosEnajenante  `json:"datos_enajenante,omitempty"`
	DatosAdquiriente *NotariosPublicosDatosAdquiriente `json:"datos_adquiriente,omitempty"`
}

// NotariosPublicosDescInmueble es el nodo notariospublicos:DescInmueble del complemento NotariosPublicos.
type NotariosPublicosDescInmueble struct {
	TipoInmueble string `json:"tipo_inmueble,omitempty"`
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

// NotariosPublicosDatosOperacion es el nodo notariospublicos:DatosOperacion del complemento NotariosPublicos.
type NotariosPublicosDatosOperacion struct {
	NumInstrumentoNotarial string           `json:"num_instrumento_notarial,omitempty"`
	FechaInstNotarial      Fecha            `json:"fecha_inst_notarial,omitzero"`
	MontoOperacion         *decimal.Decimal `json:"monto_operacion,omitempty"`
	Subtotal               *decimal.Decimal `json:"subtotal,omitempty"`
	IVA                    *decimal.Decimal `json:"iva,omitempty"`
}

// NotariosPublicosDatosNotario es el nodo notariospublicos:DatosNotario del complemento NotariosPublicos.
type NotariosPublicosDatosNotario struct {
	CURP              string `json:"curp,omitempty"`
	NumNotaria        string `json:"num_notaria,omitempty"`
	EntidadFederativa string `json:"entidad_federativa,omitempty"`
	Adscripcion       string `json:"adscripcion,omitempty"`
}

// NotariosPublicosDatosEnajenante trae uno de dos: DatosUnEnajenante, o la
// lista de copropietarios DatosEnajenantesCopSC (aplanada).
type NotariosPublicosDatosEnajenante struct {
	CoproSocConyugalE     string                               `json:"copro_soc_conyugal_e,omitempty"`
	DatosUnEnajenante     *NotariosPublicosDatosPersona        `json:"datos_un_enajenante,omitempty"`
	DatosEnajenantesCopSC []NotariosPublicosDatosCopropietario `json:"datos_enajenantes_cop_sc,omitempty"`
}

// NotariosPublicosDatosAdquiriente es como DatosEnajenante, con los
// adquirientes.
type NotariosPublicosDatosAdquiriente struct {
	CoproSocConyugalE      string                               `json:"copro_soc_conyugal_e,omitempty"`
	DatosUnAdquiriente     *NotariosPublicosDatosPersona        `json:"datos_un_adquiriente,omitempty"`
	DatosAdquirientesCopSC []NotariosPublicosDatosCopropietario `json:"datos_adquirientes_cop_sc,omitempty"`
}

// NotariosPublicosDatosPersona es DatosUnEnajenante o DatosUnAdquiriente
// (mismos atributos; el XSD exige más en el enajenante).
type NotariosPublicosDatosPersona struct {
	Nombre          string `json:"nombre,omitempty"`
	ApellidoPaterno string `json:"apellido_paterno,omitempty"`
	ApellidoMaterno string `json:"apellido_materno,omitempty"`
	RFC             string `json:"rfc,omitempty"`
	CURP            string `json:"curp,omitempty"`
}

// NotariosPublicosDatosCopropietario es DatosEnajenanteCopSC o
// DatosAdquirienteCopSC.
type NotariosPublicosDatosCopropietario struct {
	Nombre          string           `json:"nombre,omitempty"`
	ApellidoPaterno string           `json:"apellido_paterno,omitempty"`
	ApellidoMaterno string           `json:"apellido_materno,omitempty"`
	RFC             string           `json:"rfc,omitempty"`
	CURP            string           `json:"curp,omitempty"`
	Porcentaje      *decimal.Decimal `json:"porcentaje,omitempty"`
}

// ParcialesConstruccion 1.0 (servicioparcialconstruccion, nodo
// parcialesconstruccion): servicios parciales de construcción de inmuebles
// destinados a casa habitación. El XSD lo pone en el Complemento del
// comprobante; si viene en un ComplementoConcepto también se lee, en el
// concepto.
type ParcialesConstruccion struct {
	Version       string                         `json:"version,omitempty"`
	NumPerLicoAut string                         `json:"num_per_lico_aut,omitempty"`
	Inmueble      *ParcialesConstruccionInmueble `json:"inmueble,omitempty"`
}

// ParcialesConstruccionInmueble es el nodo servicioparcial:Inmueble del complemento de servicios parciales de construcción.
type ParcialesConstruccionInmueble struct {
	Calle        string `json:"calle,omitempty"`
	NoExterior   string `json:"no_exterior,omitempty"`
	NoInterior   string `json:"no_interior,omitempty"`
	Colonia      string `json:"colonia,omitempty"`
	Localidad    string `json:"localidad,omitempty"`
	Referencia   string `json:"referencia,omitempty"`
	Municipio    string `json:"municipio,omitempty"`
	Estado       string `json:"estado,omitempty"`
	CodigoPostal string `json:"codigo_postal,omitempty"`
}

// VehiculoUsado 1.0: enajenación de un vehículo usado por una persona
// física. InformacionAduanera es la de terceros y ventavehiculos.
type VehiculoUsado struct {
	Version             string                           `json:"version,omitempty"`
	MontoAdquisicion    *decimal.Decimal                 `json:"monto_adquisicion,omitempty"`
	MontoEnajenacion    *decimal.Decimal                 `json:"monto_enajenacion,omitempty"`
	ClaveVehicular      string                           `json:"clave_vehicular,omitempty"`
	Marca               string                           `json:"marca,omitempty"`
	Tipo                string                           `json:"tipo,omitempty"`
	Modelo              string                           `json:"modelo,omitempty"`
	NumeroMotor         string                           `json:"numero_motor,omitempty"`
	NumeroSerie         string                           `json:"numero_serie,omitempty"`
	NIV                 string                           `json:"niv,omitempty"`
	Valor               *decimal.Decimal                 `json:"valor,omitempty"`
	InformacionAduanera []InformacionAduaneraComplemento `json:"informacion_aduanera,omitempty"`
}

// TuristaPasajeroExtranjero 1.0.
type TuristaPasajeroExtranjero struct {
	Version         string                                  `json:"version,omitempty"`
	FechadeTransito Fecha                                   `json:"fechade_transito,omitzero"`
	TipoTransito    string                                  `json:"tipo_transito,omitempty"`
	DatosTransito   *TuristaPasajeroExtranjeroDatosTransito `json:"datos_transito,omitempty"`
}

// TuristaPasajeroExtranjeroDatosTransito es el nodo tpe:datosTransito del complemento TuristaPasajeroExtranjero.
type TuristaPasajeroExtranjeroDatosTransito struct {
	Via               string `json:"via,omitempty"`
	TipoId            string `json:"tipo_id,omitempty"`
	NumeroId          string `json:"numero_id,omitempty"`
	Nacionalidad      string `json:"nacionalidad,omitempty"`
	EmpresaTransporte string `json:"empresa_transporte,omitempty"`
	IdTransporte      string `json:"id_transporte,omitempty"`
}

// AcreditamientoIEPS 1.0 (acreditamiento, nodo acreditamientoIEPS),
// complemento de concepto: TAR de la enajenación de combustible.
type AcreditamientoIEPS struct {
	Version string `json:"version,omitempty"`
	TAR     string `json:"tar,omitempty"`
}

// SPEITercero es un SPEI_Tercero del complemento SPEI (spei), que no tiene
// versión: el comprobante de una transferencia SPEI a terceros que emite
// Banxico. Complemento_SPEI no tiene atributos y se aplana
// (Comprobante.ComplementoSPEI). Sello y CadenaCDA son de Banxico y, como
// los demás sellos, solo salen en JSON con IncluirSello.
type SPEITercero struct {
	FechaOperacion    Fecha             `json:"fecha_operacion,omitzero"`
	Hora              string            `json:"hora,omitempty"` // xs:time, tal cual
	ClaveSPEI         string            `json:"clave_spei,omitempty"`
	Sello             string            `json:"sello,omitempty"`
	NumeroCertificado string            `json:"numero_certificado,omitempty"`
	CadenaCDA         string            `json:"cadena_cda,omitempty"`
	Ordenante         *SPEIOrdenante    `json:"ordenante,omitempty"`
	Beneficiario      *SPEIBeneficiario `json:"beneficiario,omitempty"`
}

// SPEIOrdenante y SPEIBeneficiario: TipoCuenta y Cuenta son xs:decimal en
// el XSD; el decimal conserva el texto (los ceros a la izquierda de una
// CLABE incluidos).
type SPEIOrdenante struct {
	BancoEmisor string           `json:"banco_emisor,omitempty"`
	Nombre      string           `json:"nombre,omitempty"`
	TipoCuenta  *decimal.Decimal `json:"tipo_cuenta,omitempty"`
	Cuenta      *decimal.Decimal `json:"cuenta,omitempty"`
	RFC         string           `json:"rfc,omitempty"`
}

// SPEIBeneficiario es el nodo spei:Beneficiario del complemento SPEI de terceros.
type SPEIBeneficiario struct {
	BancoReceptor string           `json:"banco_receptor,omitempty"`
	Nombre        string           `json:"nombre,omitempty"`
	TipoCuenta    *decimal.Decimal `json:"tipo_cuenta,omitempty"`
	Cuenta        *decimal.Decimal `json:"cuenta,omitempty"`
	RFC           string           `json:"rfc,omitempty"`
	Concepto      string           `json:"concepto,omitempty"`
	IVA           *decimal.Decimal `json:"iva,omitempty"`
	MontoPago     *decimal.Decimal `json:"monto_pago,omitempty"`
}

func (l *lector) ine(n *nodo, r string) *INE {
	ns := n.ns
	x := &INE{
		Version:        l.versionComplemento(n, r, "1.1", "1.0"),
		TipoProceso:    l.str(n, "TipoProceso"),
		TipoComite:     l.str(n, "TipoComite"),
		IdContabilidad: l.str(n, "IdContabilidad"),
	}
	l.requerido(n, r, "TipoProceso")
	for i, h := range l.hijos(n, ns, "Entidad", r) {
		re := idx(r, "Entidad", i)
		e := INEEntidad{
			ClaveEntidad: l.str(h, "ClaveEntidad"),
			Ambito:       l.str(h, "Ambito"),
		}
		l.requerido(h, re, "ClaveEntidad")
		for j, c := range l.hijos(h, ns, "Contabilidad", re) {
			e.Contabilidad = append(e.Contabilidad, INEContabilidad{IdContabilidad: l.str(c, "IdContabilidad")})
			l.requerido(c, idx(re, "Contabilidad", j), "IdContabilidad")
		}
		x.Entidad = append(x.Entidad, e)
	}
	return x
}

func (l *lector) valesDeDespensa(n *nodo, r string) *ValesDeDespensa {
	ns := n.ns
	x := &ValesDeDespensa{
		Version:          l.versionComplemento(n, r, "1.0"),
		TipoOperacion:    l.str(n, "tipoOperacion"),
		RegistroPatronal: l.str(n, "registroPatronal"),
		NumeroDeCuenta:   l.str(n, "numeroDeCuenta"),
		Total:            l.dec(n, r, "total"),
	}
	l.requerido(n, r, "tipoOperacion", "numeroDeCuenta", "total")
	if c := l.unico(n, ns, "Conceptos", r); c != nil {
		rc := r + "/Conceptos"
		for i, h := range l.hijos(c, ns, "Concepto", rc) {
			ri := idx(rc, "Concepto", i)
			x.Conceptos = append(x.Conceptos, ValesDeDespensaConcepto{
				Identificador:      l.str(h, "identificador"),
				Fecha:              l.fecha(h, ri, "fecha", false),
				Rfc:                rfc(l.str(h, "rfc")),
				Curp:               l.str(h, "curp"),
				Nombre:             l.str(h, "nombre"),
				NumSeguridadSocial: l.str(h, "numSeguridadSocial"),
				Importe:            l.dec(h, ri, "importe"),
			})
			l.requerido(h, ri, "identificador", "fecha", "rfc", "curp", "nombre", "importe")
		}
		if len(x.Conceptos) == 0 {
			l.advertir(AdvNodoFaltante, rc+"/Concepto", "Conceptos sin ningún Concepto")
		}
	} else {
		l.advertir(AdvNodoFaltante, r+"/Conceptos", "ValesDeDespensa sin Conceptos")
	}
	return x
}

// consumoDeCombustibles lee la 1.1 o la 1.0 según el namespace; cada una
// exige lo suyo (tipoCombustible y tasaOCuota en 1.1, tasa en 1.0).
func (l *lector) consumoDeCombustibles(n *nodo, r string) *ConsumoDeCombustibles {
	ns := n.ns
	v11 := ns == NSConsumoDeCombustibles11
	esperada := "1.0"
	if v11 {
		esperada = "1.1"
	}
	x := &ConsumoDeCombustibles{
		Version:        l.versionComplemento(n, r, esperada),
		TipoOperacion:  l.str(n, "tipoOperacion"),
		NumeroDeCuenta: l.str(n, "numeroDeCuenta"),
		SubTotal:       l.dec(n, r, "subTotal"),
		Total:          l.dec(n, r, "total"),
	}
	l.requerido(n, r, "tipoOperacion", "numeroDeCuenta", "total")
	c := l.unico(n, ns, "Conceptos", r)
	if c == nil {
		l.advertir(AdvNodoFaltante, r+"/Conceptos", "ConsumoDeCombustibles sin Conceptos")
		return x
	}
	rc := r + "/Conceptos"
	for i, h := range l.hijos(c, ns, "ConceptoConsumoDeCombustibles", rc) {
		ri := idx(rc, "ConceptoConsumoDeCombustibles", i)
		cp := ConceptoConsumoDeCombustibles{
			Identificador:     l.str(h, "identificador"),
			Fecha:             l.fecha(h, ri, "fecha", false),
			Rfc:               rfc(l.str(h, "rfc")),
			ClaveEstacion:     l.str(h, "claveEstacion"),
			TipoCombustible:   l.str(h, "tipoCombustible"),
			Cantidad:          l.dec(h, ri, "cantidad"),
			NombreCombustible: l.str(h, "nombreCombustible"),
			FolioOperacion:    l.str(h, "folioOperacion"),
			ValorUnitario:     l.dec(h, ri, "valorUnitario"),
			Importe:           l.dec(h, ri, "importe"),
		}
		l.requerido(h, ri, "identificador", "fecha", "rfc", "claveEstacion", "cantidad", "nombreCombustible",
			"folioOperacion", "valorUnitario", "importe")
		if v11 {
			l.requerido(h, ri, "tipoCombustible")
		}
		if ds := l.unico(h, ns, "Determinados", ri); ds != nil {
			rd := ri + "/Determinados"
			for j, d := range l.hijos(ds, ns, "Determinado", rd) {
				rj := idx(rd, "Determinado", j)
				de := ConsumoDeCombustiblesDeterminado{
					Impuesto: l.str(d, "impuesto"),
					Importe:  l.dec(d, rj, "importe"),
				}
				if v11 {
					de.TasaOCuota = l.dec(d, rj, "tasaOCuota")
					l.requerido(d, rj, "impuesto", "tasaOCuota", "importe")
				} else {
					de.Tasa = l.dec(d, rj, "tasa")
					l.requerido(d, rj, "impuesto", "tasa", "importe")
				}
				cp.Determinados = append(cp.Determinados, de)
			}
			if len(cp.Determinados) == 0 {
				l.advertir(AdvNodoFaltante, rd+"/Determinado", "Determinados sin ningún Determinado")
			}
		} else {
			l.advertir(AdvNodoFaltante, ri+"/Determinados", "ConceptoConsumoDeCombustibles sin Determinados")
		}
		x.Conceptos = append(x.Conceptos, cp)
	}
	if len(x.Conceptos) == 0 {
		l.advertir(AdvNodoFaltante, rc+"/ConceptoConsumoDeCombustibles", "Conceptos sin ningún ConceptoConsumoDeCombustibles")
	}
	return x
}

func (l *lector) pagoEnEspecie(n *nodo, r string) *PagoEnEspecie {
	x := &PagoEnEspecie{
		Version:      l.versionComplemento(n, r, "1.0"),
		CvePIC:       l.str(n, "CvePIC"),
		FolioSolDon:  l.str(n, "FolioSolDon"),
		PzaArtNombre: l.str(n, "PzaArtNombre"),
		PzaArtTecn:   l.str(n, "PzaArtTecn"),
		PzaArtAProd:  l.str(n, "PzaArtAProd"),
		PzaArtDim:    l.str(n, "PzaArtDim"),
	}
	l.requerido(n, r, "CvePIC", "FolioSolDon", "PzaArtNombre", "PzaArtTecn", "PzaArtAProd", "PzaArtDim")
	return x
}

func (l *lector) notariosPublicos(n *nodo, r string) *NotariosPublicos {
	ns := n.ns
	x := &NotariosPublicos{Version: l.versionComplemento(n, r, "1.0")}
	if c := l.unico(n, ns, "DescInmuebles", r); c != nil {
		rc := r + "/DescInmuebles"
		for i, h := range l.hijos(c, ns, "DescInmueble", rc) {
			ri := idx(rc, "DescInmueble", i)
			x.DescInmuebles = append(x.DescInmuebles, NotariosPublicosDescInmueble{
				TipoInmueble: l.str(h, "TipoInmueble"),
				Calle:        l.str(h, "Calle"),
				NoExterior:   l.str(h, "NoExterior"),
				NoInterior:   l.str(h, "NoInterior"),
				Colonia:      l.str(h, "Colonia"),
				Localidad:    l.str(h, "Localidad"),
				Referencia:   l.str(h, "Referencia"),
				Municipio:    l.str(h, "Municipio"),
				Estado:       l.str(h, "Estado"),
				Pais:         l.str(h, "Pais"),
				CodigoPostal: l.str(h, "CodigoPostal"),
			})
			l.requerido(h, ri, "TipoInmueble", "Calle", "Municipio", "Estado", "Pais", "CodigoPostal")
		}
		if len(x.DescInmuebles) == 0 {
			l.advertir(AdvNodoFaltante, rc+"/DescInmueble", "DescInmuebles sin ningún DescInmueble")
		}
	} else {
		l.advertir(AdvNodoFaltante, r+"/DescInmuebles", "NotariosPublicos sin DescInmuebles")
	}
	if h := l.unico(n, ns, "DatosOperacion", r); h != nil {
		rh := r + "/DatosOperacion"
		x.DatosOperacion = &NotariosPublicosDatosOperacion{
			NumInstrumentoNotarial: l.str(h, "NumInstrumentoNotarial"),
			FechaInstNotarial:      l.fecha(h, rh, "FechaInstNotarial", true),
			MontoOperacion:         l.dec(h, rh, "MontoOperacion"),
			Subtotal:               l.dec(h, rh, "Subtotal"),
			IVA:                    l.dec(h, rh, "IVA"),
		}
		l.requerido(h, rh, "NumInstrumentoNotarial", "FechaInstNotarial", "MontoOperacion", "Subtotal", "IVA")
	} else {
		l.advertir(AdvNodoFaltante, r+"/DatosOperacion", "NotariosPublicos sin DatosOperacion")
	}
	if h := l.unico(n, ns, "DatosNotario", r); h != nil {
		x.DatosNotario = &NotariosPublicosDatosNotario{
			CURP:              l.str(h, "CURP"),
			NumNotaria:        l.str(h, "NumNotaria"),
			EntidadFederativa: l.str(h, "EntidadFederativa"),
			Adscripcion:       l.str(h, "Adscripcion"),
		}
		l.requerido(h, r+"/DatosNotario", "CURP", "NumNotaria", "EntidadFederativa")
	} else {
		l.advertir(AdvNodoFaltante, r+"/DatosNotario", "NotariosPublicos sin DatosNotario")
	}
	if h := l.unico(n, ns, "DatosEnajenante", r); h != nil {
		rh := r + "/DatosEnajenante"
		e := &NotariosPublicosDatosEnajenante{CoproSocConyugalE: l.str(h, "CoproSocConyugalE")}
		l.requerido(h, rh, "CoproSocConyugalE")
		e.DatosUnEnajenante, e.DatosEnajenantesCopSC = l.notariosPartes(h, ns, rh, "Enajenante", "ApellidoPaterno", "CURP")
		x.DatosEnajenante = e
	} else {
		l.advertir(AdvNodoFaltante, r+"/DatosEnajenante", "NotariosPublicos sin DatosEnajenante")
	}
	if h := l.unico(n, ns, "DatosAdquiriente", r); h != nil {
		rh := r + "/DatosAdquiriente"
		a := &NotariosPublicosDatosAdquiriente{CoproSocConyugalE: l.str(h, "CoproSocConyugalE")}
		l.requerido(h, rh, "CoproSocConyugalE")
		a.DatosUnAdquiriente, a.DatosAdquirientesCopSC = l.notariosPartes(h, ns, rh, "Adquiriente")
		x.DatosAdquiriente = a
	} else {
		l.advertir(AdvNodoFaltante, r+"/DatosAdquiriente", "NotariosPublicos sin DatosAdquiriente")
	}
	return x
}

// notariosPartes lee DatosUn<parte> y DatosEnajenantesCopSC (o
// DatosAdquirientesCopSC) de n. El XSD exige en DatosUnEnajenante, además
// de Nombre y RFC, ApellidoPaterno y CURP (masReq).
func (l *lector) notariosPartes(n *nodo, ns, r, parte string, masReq ...string) (*NotariosPublicosDatosPersona, []NotariosPublicosDatosCopropietario) {
	var un *NotariosPublicosDatosPersona
	if h := l.unico(n, ns, "DatosUn"+parte, r); h != nil {
		un = &NotariosPublicosDatosPersona{
			Nombre:          l.str(h, "Nombre"),
			ApellidoPaterno: l.str(h, "ApellidoPaterno"),
			ApellidoMaterno: l.str(h, "ApellidoMaterno"),
			RFC:             rfc(l.str(h, "RFC")),
			CURP:            l.str(h, "CURP"),
		}
		l.requerido(h, r+"/DatosUn"+parte, append([]string{"Nombre", "RFC"}, masReq...)...)
	}
	cont, local := "Datos"+parte+"sCopSC", "Datos"+parte+"CopSC"
	c := l.unico(n, ns, cont, r)
	if c == nil {
		return un, nil
	}
	rc := r + "/" + cont
	var cop []NotariosPublicosDatosCopropietario
	for i, h := range l.hijos(c, ns, local, rc) {
		ri := idx(rc, local, i)
		cop = append(cop, NotariosPublicosDatosCopropietario{
			Nombre:          l.str(h, "Nombre"),
			ApellidoPaterno: l.str(h, "ApellidoPaterno"),
			ApellidoMaterno: l.str(h, "ApellidoMaterno"),
			RFC:             rfc(l.str(h, "RFC")),
			CURP:            l.str(h, "CURP"),
			Porcentaje:      l.dec(h, ri, "Porcentaje"),
		})
		l.requerido(h, ri, "Nombre", "RFC", "Porcentaje")
	}
	if len(cop) == 0 {
		l.advertir(AdvNodoFaltante, rc+"/"+local, cont+" sin ningún "+local)
	}
	return un, cop
}

func (l *lector) parcialesConstruccion(n *nodo, r string) *ParcialesConstruccion {
	x := &ParcialesConstruccion{
		Version:       l.versionComplemento(n, r, "1.0"),
		NumPerLicoAut: l.str(n, "NumPerLicoAut"),
	}
	l.requerido(n, r, "NumPerLicoAut")
	if h := l.unico(n, n.ns, "Inmueble", r); h != nil {
		ri := r + "/Inmueble"
		x.Inmueble = &ParcialesConstruccionInmueble{
			Calle:        l.str(h, "Calle"),
			NoExterior:   l.str(h, "NoExterior"),
			NoInterior:   l.str(h, "NoInterior"),
			Colonia:      l.str(h, "Colonia"),
			Localidad:    l.str(h, "Localidad"),
			Referencia:   l.str(h, "Referencia"),
			Municipio:    l.str(h, "Municipio"),
			Estado:       l.str(h, "Estado"),
			CodigoPostal: l.str(h, "CodigoPostal"),
		}
		l.requerido(h, ri, "Calle", "Municipio", "Estado", "CodigoPostal")
	} else {
		l.advertir(AdvNodoFaltante, r+"/Inmueble", "parcialesconstruccion sin Inmueble")
	}
	return x
}

func (l *lector) vehiculoUsado(n *nodo, r string) *VehiculoUsado {
	x := &VehiculoUsado{
		Version:          l.versionComplemento(n, r, "1.0"),
		MontoAdquisicion: l.dec(n, r, "montoAdquisicion"),
		MontoEnajenacion: l.dec(n, r, "montoEnajenacion"),
		ClaveVehicular:   l.str(n, "claveVehicular"),
		Marca:            l.str(n, "marca"),
		Tipo:             l.str(n, "tipo"),
		Modelo:           l.str(n, "modelo"),
		NumeroMotor:      l.str(n, "numeroMotor"),
		NumeroSerie:      l.str(n, "numeroSerie"),
		NIV:              l.str(n, "NIV"),
		Valor:            l.dec(n, r, "valor"),
	}
	l.requerido(n, r, "montoAdquisicion", "montoEnajenacion", "claveVehicular", "marca", "tipo", "modelo", "valor")
	for i, h := range l.hijos(n, n.ns, "InformacionAduanera", r) {
		x.InformacionAduanera = append(x.InformacionAduanera, l.aduaneraComplemento(h, idx(r, "InformacionAduanera", i)))
	}
	return x
}

func (l *lector) turistaPasajeroExtranjero(n *nodo, r string) *TuristaPasajeroExtranjero {
	x := &TuristaPasajeroExtranjero{
		Version:         l.versionComplemento(n, r, "1.0"),
		FechadeTransito: l.fecha(n, r, "fechadeTransito", false),
		TipoTransito:    l.str(n, "tipoTransito"),
	}
	l.requerido(n, r, "fechadeTransito", "tipoTransito")
	if h := l.unico(n, n.ns, "datosTransito", r); h != nil {
		x.DatosTransito = &TuristaPasajeroExtranjeroDatosTransito{
			Via:               l.str(h, "Via"),
			TipoId:            l.str(h, "TipoId"),
			NumeroId:          l.str(h, "NumeroId"),
			Nacionalidad:      l.str(h, "Nacionalidad"),
			EmpresaTransporte: l.str(h, "EmpresaTransporte"),
			IdTransporte:      l.str(h, "IdTransporte"),
		}
		l.requerido(h, r+"/datosTransito", "Via", "TipoId", "NumeroId", "Nacionalidad", "EmpresaTransporte")
	} else {
		l.advertir(AdvNodoFaltante, r+"/datosTransito", "TuristaPasajeroExtranjero sin datosTransito")
	}
	return x
}

func (l *lector) acreditamientoIEPS(n *nodo, r string) *AcreditamientoIEPS {
	x := &AcreditamientoIEPS{
		Version: l.versionComplemento(n, r, "1.0"),
		TAR:     l.str(n, "TAR"),
	}
	l.requerido(n, r, "TAR")
	return x
}

// complementoSPEI lee los SPEI_Tercero de un Complemento_SPEI.
func (l *lector) complementoSPEI(n *nodo, r string) []SPEITercero {
	ns := n.ns
	var out []SPEITercero
	for i, h := range l.hijos(n, ns, "SPEI_Tercero", r) {
		rt := idx(r, "SPEI_Tercero", i)
		t := SPEITercero{
			FechaOperacion:    l.fecha(h, rt, "FechaOperacion", true),
			Hora:              l.str(h, "Hora"),
			ClaveSPEI:         l.str(h, "ClaveSPEI"),
			Sello:             l.str(h, "sello"),
			NumeroCertificado: l.str(h, "numeroCertificado"),
			CadenaCDA:         l.str(h, "cadenaCDA"),
		}
		l.requerido(h, rt, "FechaOperacion", "Hora", "ClaveSPEI", "sello", "numeroCertificado", "cadenaCDA")
		if o := l.unico(h, ns, "Ordenante", rt); o != nil {
			ro := rt + "/Ordenante"
			t.Ordenante = &SPEIOrdenante{
				BancoEmisor: l.str(o, "BancoEmisor"),
				Nombre:      l.str(o, "Nombre"),
				TipoCuenta:  l.dec(o, ro, "TipoCuenta"),
				Cuenta:      l.dec(o, ro, "Cuenta"),
				RFC:         rfc(l.str(o, "RFC")),
			}
			l.requerido(o, ro, "BancoEmisor", "Nombre", "TipoCuenta", "Cuenta", "RFC")
		} else {
			l.advertir(AdvNodoFaltante, rt+"/Ordenante", "SPEI_Tercero sin Ordenante")
		}
		if b := l.unico(h, ns, "Beneficiario", rt); b != nil {
			rb := rt + "/Beneficiario"
			t.Beneficiario = &SPEIBeneficiario{
				BancoReceptor: l.str(b, "BancoReceptor"),
				Nombre:        l.str(b, "Nombre"),
				TipoCuenta:    l.dec(b, rb, "TipoCuenta"),
				Cuenta:        l.dec(b, rb, "Cuenta"),
				RFC:           rfc(l.str(b, "RFC")),
				Concepto:      l.str(b, "Concepto"),
				IVA:           l.dec(b, rb, "IVA"),
				MontoPago:     l.dec(b, rb, "MontoPago"),
			}
			l.requerido(b, rb, "BancoReceptor", "Nombre", "TipoCuenta", "Cuenta", "RFC", "Concepto", "MontoPago")
		} else {
			l.advertir(AdvNodoFaltante, rt+"/Beneficiario", "SPEI_Tercero sin Beneficiario")
		}
		out = append(out, t)
	}
	if len(out) == 0 {
		l.advertir(AdvNodoFaltante, r+"/SPEI_Tercero", "Complemento_SPEI sin ningún SPEI_Tercero")
	}
	return out
}
