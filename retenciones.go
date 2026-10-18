package cfdi

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"

	"github.com/OrtaMarco/cfdi-go/decimal"
)

// Namespaces del CFDI de Retenciones e información de pagos y de sus
// complementos.
const (
	NSRetenciones10                 = "http://www.sat.gob.mx/esquemas/retencionpago/1"
	NSRetenciones20                 = "http://www.sat.gob.mx/esquemas/retencionpago/2"
	NSRetPlataformasTecnologicas10  = "http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10"
	NSRetIntereses                  = "http://www.sat.gob.mx/esquemas/retencionpago/1/intereses"
	NSRetPagosAExtranjeros          = "http://www.sat.gob.mx/esquemas/retencionpago/1/pagosaextranjeros"
	NSRetDividendos                 = "http://www.sat.gob.mx/esquemas/retencionpago/1/dividendos"
	NSRetEnajenacionDeAcciones      = "http://www.sat.gob.mx/esquemas/retencionpago/1/enajenaciondeacciones"
	NSRetArrendamientoEnFideicomiso = "http://www.sat.gob.mx/esquemas/retencionpago/1/arrendamientoenfideicomiso"
	NSRetOperacionesConDerivados    = "http://www.sat.gob.mx/esquemas/retencionpago/1/operacionesconderivados"
	NSRetInteresesHipotecarios      = "http://www.sat.gob.mx/esquemas/retencionpago/1/intereseshipotecarios"
	NSRetFideicomisoNoEmpresarial   = "http://www.sat.gob.mx/esquemas/retencionpago/1/fideicomisonoempresarial"
	NSRetPlanesDeRetiro10           = "http://www.sat.gob.mx/esquemas/retencionpago/1/planesderetiro"
	NSRetPlanesDeRetiro11           = "http://www.sat.gob.mx/esquemas/retencionpago/1/planesderetiro11"
	NSRetPremios                    = "http://www.sat.gob.mx/esquemas/retencionpago/1/premios"
	NSRetSectorFinanciero           = "http://www.sat.gob.mx/esquemas/retencionpago/1/sectorfinanciero"
)

// SchemaVersionRetenciones es la versión del formato JSON de Retenciones
// (schema/retenciones.v1.schema.json). Sube por su cuenta, como
// SchemaVersion la del CFDI.
const SchemaVersionRetenciones = "1"

// TipoDocumento es lo que dice Detectar de un XML.
type TipoDocumento string

// Tipos de documento de Detectar.
const (
	TipoComprobante TipoDocumento = "comprobante" // raíz Comprobante de cfd/3 o cfd/4: se lee con Parse
	TipoRetenciones TipoDocumento = "retenciones" // raíz Retenciones de retencionpago/1 o /2: se lee con ParseRetenciones
	TipoOtro        TipoDocumento = "otro"        // otra raíz, o no es XML
)

// Detectar mira solo el elemento raíz para decidir con qué función leer el
// documento. Lee tolerante como Parse (encoding corregido, XML no
// estricto): lo que Parse o ParseRetenciones rechazarían por no ser XML sale
// como TipoOtro.
func Detectar(b []byte) TipoDocumento {
	b = normalizarCodificacion(b, func(string, string) {})
	d := xml.NewDecoder(bytes.NewReader(b))
	d.CharsetReader = charsetIdentidad
	d.Strict = false
	for {
		tok, err := d.Token()
		if err != nil {
			return TipoOtro
		}
		if t, ok := tok.(xml.StartElement); ok {
			switch {
			case t.Name.Local == "Comprobante" && (t.Name.Space == NSCFDI33 || t.Name.Space == NSCFDI40):
				return TipoComprobante
			case t.Name.Local == "Retenciones" && (t.Name.Space == NSRetenciones10 || t.Name.Space == NSRetenciones20):
				return TipoRetenciones
			}
			return TipoOtro
		}
	}
}

// Retenciones es un CFDI de Retenciones e información de pagos (1.0 o 2.0)
// leído y normalizado.
//
// Las dos versiones van en el mismo modelo con los nombres de la 2.0; en
// 1.0 cambian estos (el nombre de 1.0 entre paréntesis): NoCertificado
// (NumCert), Certificado (Cert), Emisor.RfcE (RFCEmisor), Emisor.CurpE
// (CURPE, solo 1.0), Receptor.NacionalidadR (Nacionalidad), Nacional.RfcR
// (RFCRecep), Nacional.CurpR (CURPR), Extranjero.NumRegIdTribR
// (NumRegIdTrib), Periodo.Ejercicio (Ejerc), Totales.MontoTot* y MontoRet
// (montoTot*, montoRet) e ImpRetenidos.ImpuestoRet (Impuesto). Las rutas de
// las advertencias usan el nombre que trae el documento. Lo exclusivo de 2.0
// (LugarExpRetenc, CfdiRetenRelacionados, RegimenFiscalE, DomicilioFiscalR,
// UtilidadBimestral, ISRCorrespondiente) viene vacío en 1.0.
type Retenciones struct {
	Version       string `json:"version,omitempty"`
	FolioInt      string `json:"folio_int,omitempty"`
	Sello         string `json:"sello,omitempty"`
	NoCertificado string `json:"no_certificado,omitempty"`
	Certificado   string `json:"certificado,omitempty"`
	// FechaExp: en 2.0, hora civil sin zona como la Fecha del CFDI; en 1.0
	// trae su zona («2015-02-24T11:02:46-06:00»). Fecha.Time es la hora tal
	// cual viene escrita, sin aplicar la zona; el texto la conserva.
	FechaExp       Fecha  `json:"fecha_exp,omitzero"`
	LugarExpRetenc string `json:"lugar_exp_retenc,omitempty"`
	CveRetenc      string `json:"cve_retenc,omitempty"`
	DescRetenc     string `json:"desc_retenc,omitempty"`

	CfdiRetenRelacionados *CfdiRetenRelacionados `json:"cfdi_reten_relacionados,omitempty"`
	Emisor                *EmisorRetenciones     `json:"emisor,omitempty"`
	Receptor              *ReceptorRetenciones   `json:"receptor,omitempty"`
	Periodo               *PeriodoRetenciones    `json:"periodo,omitempty"`
	Totales               *TotalesRetenciones    `json:"totales,omitempty"`

	// Complementos soportados.
	TimbreFiscalDigital              *TimbreFiscalDigital              `json:"timbre_fiscal_digital,omitempty"`
	ServiciosPlataformasTecnologicas *ServiciosPlataformasTecnologicas `json:"servicios_plataformas_tecnologicas,omitempty"`
	Intereses                        *Intereses                        `json:"intereses,omitempty"`
	Pagosaextranjeros                *Pagosaextranjeros                `json:"pagosaextranjeros,omitempty"`
	Dividendos                       *Dividendos                       `json:"dividendos,omitempty"`
	EnajenaciondeAcciones            *EnajenaciondeAcciones            `json:"enajenacionde_acciones,omitempty"`
	Arrendamientoenfideicomiso       *Arrendamientoenfideicomiso       `json:"arrendamientoenfideicomiso,omitempty"`
	Operacionesconderivados          *Operacionesconderivados          `json:"operacionesconderivados,omitempty"`
	Intereseshipotecarios            *Intereseshipotecarios            `json:"intereseshipotecarios,omitempty"`
	Fideicomisonoempresarial         *Fideicomisonoempresarial         `json:"fideicomisonoempresarial,omitempty"`
	Planesderetiro                   *Planesderetiro                   `json:"planesderetiro,omitempty"`
	Premios                          *Premios                          `json:"premios,omitempty"`
	SectorFinanciero                 *SectorFinanciero                 `json:"sector_financiero,omitempty"`
	// ComplementosNoSoportados lista los nodos de Complemento que no se
	// leen, como "{namespace}NombreLocal", sin repetir.
	ComplementosNoSoportados []string `json:"complementos_no_soportados,omitempty"`

	// Advertencias de la lectura. En JSON van fuera de "retenciones", en el
	// sobre.
	Advertencias []Advertencia `json:"-"`
}

// CfdiRetenRelacionados (2.0) es el comprobante de retenciones que este
// sustituye u otra relación.
type CfdiRetenRelacionados struct {
	TipoRelacion string `json:"tipo_relacion,omitempty"`
	UUID         string `json:"uuid,omitempty"`
	UUIDOriginal string `json:"uuid_original,omitempty"` // solo si difiere de UUID
}

// EmisorRetenciones es el nodo retenciones:Emisor del CFDI de Retenciones.
type EmisorRetenciones struct {
	RfcE           string `json:"rfc_e,omitempty"`
	NomDenRazSocE  string `json:"nom_den_raz_soc_e,omitempty"`
	RegimenFiscalE string `json:"regimen_fiscal_e,omitempty"` // 2.0
	CurpE          string `json:"curp_e,omitempty"`           // 1.0 (CURPE)
}

// ReceptorRetenciones trae Nacional o Extranjero según NacionalidadR.
type ReceptorRetenciones struct {
	NacionalidadR string                         `json:"nacionalidad_r,omitempty"`
	Nacional      *ReceptorNacionalRetenciones   `json:"nacional,omitempty"`
	Extranjero    *ReceptorExtranjeroRetenciones `json:"extranjero,omitempty"`
}

// ReceptorNacionalRetenciones es el nodo retenciones:Nacional del Receptor del CFDI de Retenciones.
type ReceptorNacionalRetenciones struct {
	RfcR             string `json:"rfc_r,omitempty"`
	NomDenRazSocR    string `json:"nom_den_raz_soc_r,omitempty"`
	CurpR            string `json:"curp_r,omitempty"`
	DomicilioFiscalR string `json:"domicilio_fiscal_r,omitempty"` // 2.0
}

// ReceptorExtranjeroRetenciones es el nodo retenciones:Extranjero del Receptor del CFDI de Retenciones.
type ReceptorExtranjeroRetenciones struct {
	NumRegIdTribR string `json:"num_reg_id_trib_r,omitempty"`
	NomDenRazSocR string `json:"nom_den_raz_soc_r,omitempty"`
}

// PeriodoRetenciones: meses y ejercicio como texto, como los demás enteros
// del SAT.
type PeriodoRetenciones struct {
	MesIni    string `json:"mes_ini,omitempty"`
	MesFin    string `json:"mes_fin,omitempty"`
	Ejercicio string `json:"ejercicio,omitempty"`
}

// TotalesRetenciones es el nodo retenciones:Totales del CFDI de Retenciones.
type TotalesRetenciones struct {
	MontoTotOperacion  *decimal.Decimal `json:"monto_tot_operacion,omitempty"`
	MontoTotGrav       *decimal.Decimal `json:"monto_tot_grav,omitempty"`
	MontoTotExent      *decimal.Decimal `json:"monto_tot_exent,omitempty"`
	MontoTotRet        *decimal.Decimal `json:"monto_tot_ret,omitempty"`
	UtilidadBimestral  *decimal.Decimal `json:"utilidad_bimestral,omitempty"`  // 2.0
	ISRCorrespondiente *decimal.Decimal `json:"isr_correspondiente,omitempty"` // 2.0
	ImpRetenidos       []ImpRetenidos   `json:"imp_retenidos,omitempty"`
}

// ImpRetenidos es el nodo retenciones:ImpRetenidos de los Totales del CFDI de Retenciones.
type ImpRetenidos struct {
	BaseRet     *decimal.Decimal `json:"base_ret,omitempty"`
	ImpuestoRet string           `json:"impuesto_ret,omitempty"`
	MontoRet    *decimal.Decimal `json:"monto_ret,omitempty"`
	TipoPagoRet string           `json:"tipo_pago_ret,omitempty"`
}

// ParseRetenciones lee un CFDI de Retenciones e información de pagos 1.0 o
// 2.0, con la misma tolerancia que Parse.
//
// Solo devuelve error (un *Error) si el documento no es XML o si su raíz no
// es un Retenciones de retencionpago/1 o retencionpago/2
// (ClaseNoEsRetenciones). Cualquier otra rareza se lee como se pueda y queda
// en Retenciones.Advertencias.
func ParseRetenciones(r io.Reader) (*Retenciones, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, &Error{Clase: ClaseLectura, Err: err}
	}
	return ParseRetencionesBytes(b)
}

// ParseRetencionesBytes es ParseRetenciones sobre un []byte. No modifica b.
func ParseRetencionesBytes(b []byte) (*Retenciones, error) {
	l := &lector{}
	raiz, err := l.raizTolerante(b, "de Retenciones")
	if err != nil {
		return nil, err
	}
	if raiz.local != "Retenciones" || (raiz.ns != NSRetenciones10 && raiz.ns != NSRetenciones20) {
		return nil, &Error{Clase: ClaseNoEsRetenciones, Err: fmt.Errorf("la raíz es {%s}%s", raiz.ns, raiz.local)}
	}
	rt := l.retenciones(raiz)
	rt.Advertencias = l.adv
	return rt, nil
}

// raizTolerante lee el árbol tolerante: encoding corregido, y en modo no
// estricto si el estricto falla (con advertencia). nombre completa el
// mensaje de CONTENIDO_TRAS_RAIZ ("del Comprobante").
func (l *lector) raizTolerante(b []byte, nombre string) (*nodo, error) {
	b = normalizarCodificacion(b, func(codigo, msg string) { l.advertir(codigo, "", msg) })
	raiz, resto, err := leerArbol(b, true)
	if err != nil && ClaseDe(err) == ClaseXMLMalFormado {
		if r2, resto2, err2 := leerArbol(b, false); err2 == nil {
			raiz, resto, err = r2, resto2, nil
			l.advertir(AdvXMLNoEstricto, "", "el XML no es estrictamente válido (entidades o '&' sueltos); se leyó en modo tolerante")
		}
	}
	if err != nil {
		return nil, err
	}
	if len(bytes.Trim(resto, " \t\r\n\x00")) > 0 {
		l.advertir(AdvContenidoTrasRaiz, "", fmt.Sprintf("%d bytes después del cierre %s", len(resto), nombre))
	}
	return raiz, nil
}

// DocumentoRetenciones es el sobre del JSON de Retenciones: como Documento,
// con tipo_documento para distinguirlo de un CFDI en la misma salida (una
// línea de `cfdi lote`).
type DocumentoRetenciones struct {
	Origen string `json:"origen,omitempty"`
	// TipoDocumento es siempre "retenciones".
	TipoDocumento     TipoDocumento   `json:"tipo_documento"`
	SchemaVersion     string          `json:"schema_version"`
	Retenciones       *Retenciones    `json:"retenciones"`
	Advertencias      []Advertencia   `json:"advertencias"`
	VerificacionSello *ResultadoSello `json:"verificacion_sello,omitempty"`
	VerificacionTFD   *ResultadoSello `json:"verificacion_tfd,omitempty"`
	Validacion        *Validacion     `json:"validacion,omitempty"`
}

// JSON devuelve el DocumentoRetenciones sin sellos y sin indentar.
func (rt *Retenciones) JSON() ([]byte, error) { return rt.JSONOpciones(OpcionesJSON{}) }

// JSONOpciones devuelve el DocumentoRetenciones con las opciones dadas.
func (rt *Retenciones) JSONOpciones(o OpcionesJSON) ([]byte, error) {
	var b bytes.Buffer
	if err := rt.Documento(o).Escribir(&b, o.Indentar); err != nil {
		return nil, err
	}
	return bytes.TrimSuffix(b.Bytes(), []byte("\n")), nil
}

// Documento arma el sobre. Sin IncluirSello trabaja sobre una copia.
func (rt *Retenciones) Documento(o OpcionesJSON) DocumentoRetenciones {
	cc := rt
	if !o.IncluirSello {
		cp := *rt
		cp.Sello, cp.Certificado = "", ""
		if rt.TimbreFiscalDigital != nil {
			t := *rt.TimbreFiscalDigital
			t.SelloCFD, t.SelloSAT = "", ""
			cp.TimbreFiscalDigital = &t
		}
		cc = &cp
	}
	adv := rt.Advertencias
	if adv == nil {
		adv = []Advertencia{}
	}
	return DocumentoRetenciones{TipoDocumento: TipoRetenciones, SchemaVersion: SchemaVersionRetenciones, Retenciones: cc, Advertencias: adv}
}

// Escribir manda el documento como JSON a w, como Documento.Escribir.
func (d DocumentoRetenciones) Escribir(w io.Writer, indentar bool) error {
	return escribirJSON(w, d, indentar)
}

// UUID del timbre, normalizado; "" si no hay timbre.
func (rt *Retenciones) UUID() string {
	if rt.TimbreFiscalDigital == nil {
		return ""
	}
	return rt.TimbreFiscalDigital.UUID
}

// Complementos lista los complementos presentes sin repetir: los soportados
// como "Nombre versión" ("TimbreFiscalDigital 1.1",
// "ServiciosPlataformasTecnologicas 1.0", "Planesderetiro 1.1") y los no
// soportados como "{namespace}NombreLocal".
func (rt *Retenciones) Complementos() []string {
	var out []string
	add := func(presente bool, nombre, version string) {
		if presente {
			out = append(out, nombre+" "+version)
		}
	}
	if t := rt.TimbreFiscalDigital; t != nil {
		add(true, "TimbreFiscalDigital", t.Version)
	}
	if c := rt.ServiciosPlataformasTecnologicas; c != nil {
		add(true, "ServiciosPlataformasTecnologicas", c.Version)
	}
	if c := rt.Intereses; c != nil {
		add(true, "Intereses", c.Version)
	}
	if c := rt.Pagosaextranjeros; c != nil {
		add(true, "Pagosaextranjeros", c.Version)
	}
	if c := rt.Dividendos; c != nil {
		add(true, "Dividendos", c.Version)
	}
	if c := rt.EnajenaciondeAcciones; c != nil {
		add(true, "EnajenaciondeAcciones", c.Version)
	}
	if c := rt.Arrendamientoenfideicomiso; c != nil {
		add(true, "Arrendamientoenfideicomiso", c.Version)
	}
	if c := rt.Operacionesconderivados; c != nil {
		add(true, "Operacionesconderivados", c.Version)
	}
	if c := rt.Intereseshipotecarios; c != nil {
		add(true, "Intereseshipotecarios", c.Version)
	}
	if c := rt.Fideicomisonoempresarial; c != nil {
		add(true, "Fideicomisonoempresarial", c.Version)
	}
	if c := rt.Planesderetiro; c != nil {
		add(true, "Planesderetiro", c.Version)
	}
	if c := rt.Premios; c != nil {
		add(true, "Premios", c.Version)
	}
	if c := rt.SectorFinanciero; c != nil {
		add(true, "SectorFinanciero", c.Version)
	}
	return append(out, rt.ComplementosNoSoportados...)
}

// TieneAdvertencia dice si la lectura produjo al menos una advertencia con
// ese código.
func (rt *Retenciones) TieneAdvertencia(codigo string) bool {
	for _, a := range rt.Advertencias {
		if a.Codigo == codigo {
			return true
		}
	}
	return false
}

// AdvertenciaEn es Comprobante.AdvertenciaEn para Retenciones.
func (rt *Retenciones) AdvertenciaEn(rutaNodo, atributo string) (Advertencia, bool) {
	for _, a := range rt.Advertencias {
		if a.Atributo == atributo && a.Nodo() == rutaNodo {
			return a, true
		}
	}
	return Advertencia{}, false
}
