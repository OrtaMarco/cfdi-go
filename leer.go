package cfdi

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"strings"
)

// Namespaces del SAT que reconoce la librería.
const (
	NSCFDI33                      = "http://www.sat.gob.mx/cfd/3"
	NSCFDI40                      = "http://www.sat.gob.mx/cfd/4"
	NSTimbre                      = "http://www.sat.gob.mx/TimbreFiscalDigital"
	NSPagos10                     = "http://www.sat.gob.mx/Pagos"
	NSPagos20                     = "http://www.sat.gob.mx/Pagos20"
	NSNomina11                    = "http://www.sat.gob.mx/nomina"
	NSNomina12                    = "http://www.sat.gob.mx/nomina12"
	NSImpLocal                    = "http://www.sat.gob.mx/implocal"
	NSCartaPorte10                = "http://www.sat.gob.mx/CartaPorte"
	NSCartaPorte20                = "http://www.sat.gob.mx/CartaPorte20"
	NSCartaPorte30                = "http://www.sat.gob.mx/CartaPorte30"
	NSCartaPorte31                = "http://www.sat.gob.mx/CartaPorte31"
	NSComercioExterior10          = "http://www.sat.gob.mx/ComercioExterior"
	NSComercioExterior11          = "http://www.sat.gob.mx/ComercioExterior11"
	NSComercioExterior20          = "http://www.sat.gob.mx/ComercioExterior20"
	NSDivisas                     = "http://www.sat.gob.mx/divisas"
	NSRegistroFiscal              = "http://www.sat.gob.mx/registrofiscal"
	NSLeyendasFiscales            = "http://www.sat.gob.mx/leyendasFiscales"
	NSEstadoDeCuentaCombustible11 = "http://www.sat.gob.mx/EstadoDeCuentaCombustible"
	NSEstadoDeCuentaCombustible12 = "http://www.sat.gob.mx/EstadoDeCuentaCombustible12"
	NSHidroYPetro                 = "http://www.sat.gob.mx/hidrocarburospetroliferos"
	NSTerceros                    = "http://www.sat.gob.mx/terceros"
	NSVentaVehiculos              = "http://www.sat.gob.mx/ventavehiculos"
	NSAerolineas                  = "http://www.sat.gob.mx/aerolineas"
	NSDonatarias                  = "http://www.sat.gob.mx/donat"
	NSIEDU                        = "http://www.sat.gob.mx/iedu"
	NSINE                         = "http://www.sat.gob.mx/ine"
	NSValesDeDespensa             = "http://www.sat.gob.mx/valesdedespensa"
	NSConsumoDeCombustibles10     = "http://www.sat.gob.mx/consumodecombustibles"
	NSConsumoDeCombustibles11     = "http://www.sat.gob.mx/ConsumoDeCombustibles11"
	NSPagoEnEspecie               = "http://www.sat.gob.mx/pagoenespecie"
	NSNotariosPublicos            = "http://www.sat.gob.mx/notariospublicos"
	NSServicioParcialConstruccion = "http://www.sat.gob.mx/servicioparcialconstruccion"
	NSVehiculoUsado               = "http://www.sat.gob.mx/vehiculousado"
	NSTuristaPasajeroExtranjero   = "http://www.sat.gob.mx/TuristaPasajeroExtranjero"
	NSAcreditamientoIEPS          = "http://www.sat.gob.mx/acreditamiento"
	NSSPEI                        = "http://www.sat.gob.mx/spei"
	NSRenovacionVehiculos         = "http://www.sat.gob.mx/renovacionysustitucionvehiculos"
	NSCertificadoDestruccion      = "http://www.sat.gob.mx/certificadodestruccion"
	NSObrasArteAntiguedades       = "http://www.sat.gob.mx/arteantiguedades"
	NSPFIC                        = "http://www.sat.gob.mx/pfic"
	NSGastosHidrocarburos         = "http://www.sat.gob.mx/GastosHidrocarburos10"
	NSIngresosHidrocarburos       = "http://www.sat.gob.mx/IngresosHidrocarburos10"
	NSEstadoDeCuentaBancario      = "http://www.sat.gob.mx/ecb"
	NSEstadoDeCuentaCombustible10 = "http://www.sat.gob.mx/ecc"
	NSDetallista                  = "http://www.sat.gob.mx/detallista"
	nsXSI                         = "http://www.w3.org/2001/XMLSchema-instance"
	nsXMLNSAttrs                  = "xmlns"
)

// Parse lee un CFDI 3.3 o 4.0.
//
// Solo devuelve error (un *Error) si el documento no es XML o si su raíz no es
// un Comprobante de los namespaces cfd/3 o cfd/4. Cualquier otra rareza se
// lee como se pueda y queda en Comprobante.Advertencias.
func Parse(r io.Reader) (*Comprobante, error) { return ParseOpciones(r, Opciones{}) }

// ParseBytes es Parse sobre un []byte. No modifica b.
func ParseBytes(b []byte) (*Comprobante, error) { return ParseBytesOpciones(b, Opciones{}) }

// Opciones ajusta la lectura. El valor cero es la lectura tolerante de
// Parse y ParseBytes.
//
// Es un struct (y no opciones funcionales) por lo mismo que OpcionesJSON:
// son interruptores independientes, el valor cero es el comportamiento de
// siempre y un campo nuevo no rompe a nadie.
type Opciones struct {
	// FechaEstricta solo interpreta las fechas que vienen exactamente en el
	// formato del SAT (AAAA-MM-DDThh:mm:ss, o AAAA-MM-DD en los campos de
	// solo fecha). Las demás se quedan con Time en cero (Fecha.Valida() da
	// false), Fecha.FueraDeFormato en true, y la advertencia es
	// FECHA_INVALIDA en vez de FECHA_FORMATO.
	FechaEstricta bool
}

// ParseOpciones es Parse con opciones.
func ParseOpciones(r io.Reader, o Opciones) (*Comprobante, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, &Error{Clase: ClaseLectura, Err: err}
	}
	return ParseBytesOpciones(b, o)
}

// ParseBytesOpciones es ParseBytes con opciones.
func ParseBytesOpciones(b []byte, o Opciones) (*Comprobante, error) {
	l := &lector{op: o}
	raiz, err := l.raizTolerante(b, "del Comprobante")
	if err != nil {
		return nil, err
	}
	if raiz.local != "Comprobante" || (raiz.ns != NSCFDI33 && raiz.ns != NSCFDI40) {
		return nil, &Error{Clase: ClaseNoEsComprobante, Err: fmt.Errorf("la raíz es {%s}%s", raiz.ns, raiz.local)}
	}
	c := l.comprobante(raiz)
	c.Advertencias = l.adv
	return c, nil
}

// nodo es un elemento XML ya resuelto a namespace + nombre local.
//
// El CFDI lleva casi todo en atributos, así que el texto solo se guarda
// cuando trae algo: texto es la concatenación de los trozos de texto directos
// del elemento (no de sus hijos) desde el primero que no es puro espacio XML
// (" \t\r\n"). La indentación no se guarda, para no gastar memoria en CFDI
// grandes: un elemento con puro espacio queda con texto "". Solo se guarda
// en los elementos del namespace de Detallista, el único que lo lee y que sí
// lleva datos en el texto; en los demás texto queda "".
type nodo struct {
	ns, local string
	attrs     []xml.Attr
	hijos     []*nodo
	texto     string
}

// leerArbol arma el árbol del primer elemento raíz y devuelve los bytes que
// quedan después de su cierre.
func leerArbol(b []byte, estricto bool) (*nodo, []byte, error) {
	d := xml.NewDecoder(bytes.NewReader(b))
	d.CharsetReader = charsetIdentidad
	d.Strict = estricto
	var pila []*nodo
	var raiz *nodo
	// buf es el texto acumulado del elemento abierto de cada nivel de la pila
	// (solo los de Detallista lo guardan): se convierte a string una vez, al
	// cerrar el elemento. Concatenar string a string recopiaba todo lo anterior
	// en cada trozo que parte un comentario (cuadrático).
	var buf []*bytes.Buffer
	for {
		tok, err := d.Token()
		if err != nil {
			if raiz == nil {
				if errors.Is(err, io.EOF) {
					err = errors.New("no hay elemento raíz")
				}
				return nil, nil, &Error{Clase: ClaseNoEsXML, Err: err}
			}
			if errors.Is(err, io.EOF) {
				err = io.ErrUnexpectedEOF
			}
			return nil, nil, &Error{Clase: ClaseXMLMalFormado, Err: err}
		}
		switch t := tok.(type) {
		case xml.StartElement:
			// La Addenda no se lee y puede pesar megas (catálogos, PDF en
			// base64): se recorre sin guardarla.
			if len(pila) == 1 && t.Name.Local == "Addenda" {
				if err := d.Skip(); err != nil {
					return nil, nil, &Error{Clase: ClaseXMLMalFormado, Err: err}
				}
				continue
			}
			n := &nodo{ns: t.Name.Space, local: t.Name.Local, attrs: t.Attr}
			if len(pila) == 0 {
				raiz = n
			} else {
				p := pila[len(pila)-1]
				p.hijos = append(p.hijos, n)
			}
			pila = append(pila, n)
			buf = append(buf, nil)
		case xml.EndElement:
			if b := buf[len(buf)-1]; b != nil {
				pila[len(pila)-1].texto = b.String()
			}
			pila = pila[:len(pila)-1]
			buf = buf[:len(buf)-1]
			if len(pila) == 0 {
				return raiz, b[d.InputOffset():], nil
			}
		case xml.CharData:
			if raiz == nil && len(bytes.TrimSpace(t)) > 0 {
				return nil, nil, &Error{Clase: ClaseNoEsXML, Err: errors.New("texto antes del elemento raíz")}
			}
			if len(pila) > 0 {
				i := len(pila) - 1
				// Solo Detallista lee el texto; en el resto se descarta.
				if pila[i].ns == NSDetallista && (buf[i] != nil || len(bytes.Trim(t, " \t\r\n")) > 0) {
					if buf[i] == nil {
						buf[i] = new(bytes.Buffer)
					}
					buf[i].Write(t) // t solo vale hasta el siguiente Token
				}
			}
		}
	}
}

// attr devuelve el valor de un atributo sin namespace por nombre local. Si no
// lo encuentra exacto, prueba sin distinguir mayúsculas (TFD 1.0 y CFDI 3.2
// usan minúsculas).
func (n *nodo) attr(nombre string) (string, bool) {
	for _, a := range n.attrs {
		if a.Name.Local == nombre && a.Name.Space != nsXMLNSAttrs {
			return a.Value, true
		}
	}
	for _, a := range n.attrs {
		if strings.EqualFold(a.Name.Local, nombre) && a.Name.Space != nsXMLNSAttrs && a.Name.Space != nsXSI {
			return a.Value, true
		}
	}
	return "", false
}

func (n *nodo) nombre() string {
	if n.ns == "" {
		return n.local
	}
	return "{" + n.ns + "}" + n.local
}
