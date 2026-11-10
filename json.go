package cfdi

import (
	"bytes"
	"encoding/json"
	"io"
)

// SchemaVersion es la versión del formato JSON. Sube cuando un cambio deja de
// ser compatible (renombrar o quitar una clave, cambiar un tipo); agregar
// claves opcionales no la sube.
const SchemaVersion = "1"

// Documento es el sobre del JSON: versión del esquema, el CFDI y las
// advertencias de la lectura (siempre presente, [] si no hubo).
type Documento struct {
	// Origen es de dónde salió el CFDI (lo llena quien lee en lote).
	Origen        string        `json:"origen,omitempty"`
	SchemaVersion string        `json:"schema_version"`
	CFDI          *Comprobante  `json:"cfdi"`
	Advertencias  []Advertencia `json:"advertencias"`
	// VerificacionSello es el resultado de verificar el sello del emisor
	// (Verificar(...).Sello). Documento no lo llena: lo agrega quien
	// verifica, como `cfdi lote -sello`.
	VerificacionSello *ResultadoSello `json:"verificacion_sello,omitempty"`
	// VerificacionTFD es el resultado de verificar el SelloSAT del timbre
	// (Verificar(...).TFD). Lo agrega quien verifica, como
	// `cfdi lote -certs-sat`.
	VerificacionTFD *ResultadoSello `json:"verificacion_tfd,omitempty"`
	// Validacion es el resultado de Validar. Lo agrega quien valida, como
	// `cfdi lote -validar`.
	Validacion *Validacion `json:"validacion,omitempty"`
}

// OpcionesJSON ajusta la salida JSON.
type OpcionesJSON struct {
	// IncluirSello conserva Sello, Certificado, SelloCFD, SelloSAT, los
	// sellos/certificados/cadenas de los pagos y el sello y la cadena de
	// cada SPEI_Tercero. Por omisión se quitan: pesan varios KB y no aportan
	// a quien solo quiere los datos.
	IncluirSello bool
	// Indentar con dos espacios.
	Indentar bool
}

// JSON devuelve el Documento sin sellos y sin indentar.
func (c *Comprobante) JSON() ([]byte, error) { return c.JSONOpciones(OpcionesJSON{}) }

// JSONOpciones devuelve el Documento con las opciones dadas.
func (c *Comprobante) JSONOpciones(o OpcionesJSON) ([]byte, error) {
	var b bytes.Buffer
	if err := c.Documento(o).Escribir(&b, o.Indentar); err != nil {
		return nil, err
	}
	return bytes.TrimSuffix(b.Bytes(), []byte("\n")), nil
}

// Escribir manda el Documento como JSON a w, terminado en salto de línea (una
// línea JSONL si no se indenta). No escapa <, > ni & como hace json.Marshal:
// un RFC como "E&Ñ…" sale tal cual y no como "E\u0026Ñ…".
func (d Documento) Escribir(w io.Writer, indentar bool) error {
	return escribirJSON(w, d, indentar)
}

func escribirJSON(w io.Writer, d any, indentar bool) error {
	e := json.NewEncoder(w)
	e.SetEscapeHTML(false)
	if indentar {
		e.SetIndent("", "  ")
	}
	return e.Encode(d)
}

// Documento arma el sobre. Sin IncluirSello trabaja sobre una copia: el
// Comprobante original no se toca.
func (c *Comprobante) Documento(o OpcionesJSON) Documento {
	cc := c
	if !o.IncluirSello {
		cc = c.sinSellos()
	}
	adv := c.Advertencias
	if adv == nil {
		adv = []Advertencia{}
	}
	return Documento{SchemaVersion: SchemaVersion, CFDI: cc, Advertencias: adv}
}

func (c *Comprobante) sinSellos() *Comprobante {
	cp := *c
	cp.Sello, cp.Certificado = "", ""
	if c.TimbreFiscalDigital != nil {
		t := *c.TimbreFiscalDigital
		t.SelloCFD, t.SelloSAT = "", ""
		cp.TimbreFiscalDigital = &t
	}
	if c.Pagos != nil {
		p := *c.Pagos
		p.Pago = make([]Pago, len(c.Pagos.Pago))
		for i, pg := range c.Pagos.Pago {
			pg.CertPago, pg.CadPago, pg.SelloPago = "", "", ""
			p.Pago[i] = pg
		}
		cp.Pagos = &p
	}
	if c.ComplementoSPEI != nil {
		cp.ComplementoSPEI = make([]SPEITercero, len(c.ComplementoSPEI))
		for i, t := range c.ComplementoSPEI {
			t.Sello, t.CadenaCDA = "", ""
			cp.ComplementoSPEI[i] = t
		}
	}
	return &cp
}

// UUID del timbre, normalizado; "" si no hay timbre.
func (c *Comprobante) UUID() string {
	if c.TimbreFiscalDigital == nil {
		return ""
	}
	return c.TimbreFiscalDigital.UUID
}

// Complementos lista los complementos presentes, también los de concepto y sin
// repetir: los soportados como "Nombre versión" ("TimbreFiscalDigital 1.1",
// "Pagos 2.0", "CartaPorte 3.1", "VentaVehiculos 1.1") y los no soportados
// como "{namespace}NombreLocal".
func (c *Comprobante) Complementos() []string {
	var out []string
	if t := c.TimbreFiscalDigital; t != nil {
		out = append(out, "TimbreFiscalDigital "+t.Version)
	}
	if p := c.Pagos; p != nil {
		out = append(out, "Pagos "+p.Version)
	}
	for _, n := range c.Nomina {
		agregarUnico(&out, "Nomina "+n.Version)
	}
	if il := c.ImpuestosLocales; il != nil {
		out = append(out, "ImpuestosLocales "+il.Version)
	}
	if cp := c.CartaPorte; cp != nil {
		out = append(out, "CartaPorte "+cp.Version)
	}
	if ce := c.ComercioExterior; ce != nil {
		out = append(out, "ComercioExterior "+ce.Version)
	}
	if d := c.Divisas; d != nil {
		out = append(out, "Divisas "+d.Version)
	}
	if rf := c.CFDIRegistroFiscal; rf != nil {
		out = append(out, "CFDIRegistroFiscal "+rf.Version)
	}
	if lf := c.LeyendasFiscales; lf != nil {
		out = append(out, "LeyendasFiscales "+lf.Version)
	}
	if e := c.EstadoDeCuentaCombustible; e != nil {
		out = append(out, "EstadoDeCuentaCombustible "+e.Version)
	}
	for _, h := range c.HidroYPetro {
		agregarUnico(&out, "HidroYPetro "+h.Version)
	}
	if a := c.Aerolineas; a != nil {
		out = append(out, "Aerolineas "+a.Version)
	}
	if d := c.Donatarias; d != nil {
		out = append(out, "Donatarias "+d.Version)
	}
	if x := c.INE; x != nil {
		out = append(out, "INE "+x.Version)
	}
	if x := c.ValesDeDespensa; x != nil {
		out = append(out, "ValesDeDespensa "+x.Version)
	}
	if x := c.ConsumoDeCombustibles; x != nil {
		out = append(out, "ConsumoDeCombustibles "+x.Version)
	}
	if x := c.PagoEnEspecie; x != nil {
		out = append(out, "PagoEnEspecie "+x.Version)
	}
	if x := c.NotariosPublicos; x != nil {
		out = append(out, "NotariosPublicos "+x.Version)
	}
	if x := c.ParcialesConstruccion; x != nil {
		out = append(out, "parcialesconstruccion "+x.Version)
	}
	if x := c.VehiculoUsado; x != nil {
		out = append(out, "VehiculoUsado "+x.Version)
	}
	if x := c.TuristaPasajeroExtranjero; x != nil {
		out = append(out, "TuristaPasajeroExtranjero "+x.Version)
	}
	if c.ComplementoSPEI != nil {
		out = append(out, "Complemento_SPEI") // sin versión
	}
	if x := c.RenovacionySustitucionVehiculos; x != nil {
		out = append(out, "renovacionysustitucionvehiculos "+x.Version)
	}
	if x := c.CertificadoDeDestruccion; x != nil {
		out = append(out, "certificadodedestruccion "+x.Version)
	}
	if x := c.ObrasArteAntiguedades; x != nil {
		out = append(out, "obrasarteantiguedades "+x.Version)
	}
	if x := c.PFintegranteCoordinado; x != nil {
		out = append(out, "PFintegranteCoordinado "+x.Version)
	}
	if x := c.GastosHidrocarburos; x != nil {
		out = append(out, "GastosHidrocarburos "+x.Version)
	}
	if x := c.IngresosHidrocarburos; x != nil {
		out = append(out, "IngresosHidrocarburos "+x.Version)
	}
	if x := c.EstadoDeCuentaBancario; x != nil {
		out = append(out, "EstadoDeCuentaBancario "+x.Version)
	}
	if c.Detallista != nil {
		out = append(out, "detallista") // sin versión
	}
	for _, cp := range c.Conceptos {
		if t := cp.PorCuentadeTerceros; t != nil {
			agregarUnico(&out, "PorCuentadeTerceros "+t.Version)
		}
		if v := cp.VentaVehiculos; v != nil {
			agregarUnico(&out, "VentaVehiculos "+v.Version)
		}
		if h := cp.HidroYPetro; h != nil {
			agregarUnico(&out, "HidroYPetro "+h.Version)
		}
		if ie := cp.InstEducativas; ie != nil {
			agregarUnico(&out, "instEducativas "+ie.Version)
		}
		if a := cp.AcreditamientoIEPS; a != nil {
			agregarUnico(&out, "acreditamientoIEPS "+a.Version)
		}
		if p := cp.ParcialesConstruccion; p != nil {
			agregarUnico(&out, "parcialesconstruccion "+p.Version)
		}
	}
	return append(out, c.ComplementosNoSoportados...)
}
