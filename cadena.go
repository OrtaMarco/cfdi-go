package cfdi

//go:generate go run ./internal/gencadena/generar

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

// Motivos por los que no se puede armar una cadena original. Son estables,
// como los códigos de advertencia, y son también motivos de un sello
// NO_VERIFICABLE (ver Verificar).
const (
	MotivoXMLIlegible          = "XML_ILEGIBLE"           // XML que un procesador XSLT no leería igual (mal formado, encoding que miente, caracteres prohibidos)
	MotivoNoEsComprobante      = "NO_ES_COMPROBANTE"      // la raíz no es un Comprobante cfd/3 o cfd/4
	MotivoVersionNoSoportada   = "VERSION_NO_SOPORTADA"   // versión del CFDI o del TFD sin hoja de cadena (p. ej. 3.0)
	MotivoComplementoSinCadena = "COMPLEMENTO_SIN_CADENA" // complemento que la librería no sabe llevar a la cadena
	MotivoSinTimbre            = "SIN_TIMBRE"             // no hay TimbreFiscalDigital
)

// ErrorCadena explica por qué no se pudo armar una cadena original. La
// librería no devuelve nunca una cadena que no pueda garantizar.
type ErrorCadena struct {
	Motivo string // uno de los Motivo* de arriba
	Nodo   string // "{namespace}Nombre" del nodo que la impide, si aplica
	Err    error
}

func (e *ErrorCadena) Error() string {
	s := "cfdi: cadena original: " + e.Motivo
	if e.Nodo != "" {
		s += " " + e.Nodo
	}
	if e.Err != nil {
		s += ": " + e.Err.Error()
	}
	return s
}

func (e *ErrorCadena) Unwrap() error { return e.Err }

// CadenaOriginal arma la cadena original del comprobante (CFDI 3.2, 3.3 o
// 4.0) tal como la arman los XSLT del SAT: el orden y los atributos los
// fijan esas hojas (traducidas a Go en cadena_reglas.go), los nodos que se
// repiten van en el orden del documento y cada valor pasa por
// normalize-space.
//
// Trabaja sobre los bytes y no sobre el Comprobante porque la cadena depende
// del documento tal cual (atributos vacíos, espacios, nodos en otro
// namespace) y no de lo que el modelo normaliza. Devuelve *ErrorCadena si el
// XML no es estrictamente válido, si la versión no tiene hoja o si trae un
// complemento que la librería no sabe llevar a la cadena: en ese caso el
// SAT lo incluiría y una cadena sin él saldría mal.
//
// Tampoco arma la cadena (MotivoXMLIlegible) si los bytes no son del
// encoding que declaran o traen caracteres prohibidos en XML: Parse los lee
// corrigiéndolos, pero un procesador XSLT los rechazaría o leería otros
// caracteres, y no hay forma de saber cuáles firmó el emisor.
func CadenaOriginal(xml []byte) (string, error) {
	raiz, err := arbolCadenaEstricto(xml)
	if err != nil {
		return "", err
	}
	return cadenaComprobante(raiz)
}

// CadenaOriginalTFD arma la cadena original del TimbreFiscalDigital (1.0 o
// 1.1) del CFDI, que es lo que firma el SelloSAT.
func CadenaOriginalTFD(xml []byte) (string, error) {
	raiz, err := arbolCadenaEstricto(xml)
	if err != nil {
		return "", err
	}
	tfd := timbreDe(raiz)
	if tfd == nil {
		return "", &ErrorCadena{Motivo: MotivoSinTimbre}
	}
	return cadenaTimbre(tfd)
}

func arbolCadenaEstricto(b []byte) (*nodoC, error) {
	raiz, correccion, err := arbolCadena(b)
	if err == nil && correccion != "" {
		err = &ErrorCadena{Motivo: MotivoXMLIlegible, Err: errors.New(correccion)}
	}
	return raiz, err
}

func hojaComprobante(raiz *nodoC) (*hojaCadena, error) {
	switch {
	case raiz.nombre == nombreXML{NSCFDI40, "Comprobante"} && raiz.attr("Version") == "4.0":
		return hojaCFDI40, nil
	case raiz.nombre == nombreXML{NSCFDI33, "Comprobante"} && raiz.attr("Version") == "3.3":
		return hojaCFDI33, nil
	case raiz.nombre == nombreXML{NSCFDI33, "Comprobante"} && raiz.attr("version") == "3.2":
		return hojaCFDI32, nil
	case raiz.nombre.local != "Comprobante" || (raiz.nombre.ns != NSCFDI33 && raiz.nombre.ns != NSCFDI40):
		return nil, &ErrorCadena{Motivo: MotivoNoEsComprobante, Nodo: raiz.nombre.String()}
	}
	v, ok := raiz.attrOk("Version")
	if !ok {
		v = raiz.attr("version")
	}
	return nil, &ErrorCadena{Motivo: MotivoVersionNoSoportada, Err: fmt.Errorf("version %q en %s", v, raiz.nombre.ns)}
}

func cadenaComprobante(raiz *nodoC) (string, error) {
	h, err := hojaComprobante(raiz)
	if err != nil {
		return "", err
	}
	return h.evaluar(raiz)
}

func cadenaTimbre(tfd *nodoC) (string, error) {
	switch {
	case tfd.attr("Version") == "1.1":
		return hojaTFD11.evaluar(tfd)
	case tfd.attr("version") == "1.0":
		return hojaTFD10.evaluar(tfd)
	}
	return "", &ErrorCadena{Motivo: MotivoVersionNoSoportada, Err: fmt.Errorf("TimbreFiscalDigital Version %q", tfd.attr("Version")+tfd.attr("version"))}
}

// timbreDe busca el primer TimbreFiscalDigital en Comprobante/Complemento.
func timbreDe(raiz *nodoC) *nodoC {
	for _, c := range raiz.hijos {
		if c.nombre.local != "Complemento" || c.nombre.ns != raiz.nombre.ns {
			continue
		}
		for _, t := range c.hijos {
			if t.nombre == (nombreXML{NSTimbre, "TimbreFiscalDigital"}) {
				return t
			}
		}
	}
	return nil
}

// sinCadena son los elementos que pueden venir en un Complemento y que las
// hojas del SAT no llevan a la cadena a propósito: el timbre lo agrega el PAC
// después de que el emisor sella. XSLT les aplica la plantilla integrada,
// que solo copia texto (y el TFD no tiene).
var sinCadena = map[nombreXML]bool{
	{NSTimbre, "TimbreFiscalDigital"}: true,
}

// nombreXML es un nombre de elemento resuelto: namespace y nombre local.
type nombreXML struct{ ns, local string }

func (n nombreXML) String() string { return "{" + n.ns + "}" + n.local }

type opCadena uint8

const (
	opTexto     opCadena = iota // texto literal del XSLT ("|", "||")
	opRequerido                 // plantilla Requerido: "|" + normalize-space(@attr), aunque falte
	opOpcional                  // plantilla Opcional: lo mismo, solo si el atributo existe (aunque esté vacío)
	opCada                      // xsl:for-each
	opAplicar                   // xsl:apply-templates
	opSi                        // xsl:if con una ruta: si selecciona algo
	opSiIgual                   // xsl:if test="./@attr='valor'"
	opConjunto                  // plantilla con nombre y parámetro nodo: su cuerpo (solo Requerido/Opcional) lee $param/@attr
)

type pasoRuta struct {
	nombre nombreXML // local "*" = cualquier elemento
	desc   bool      // descendiente (//) en vez de hijo (/)
}

type rutaXP struct {
	abs   bool // desde la raíz del documento
	pasos []pasoRuta
}

type pasoCadena struct {
	op     opCadena
	texto  string // literal, o nombre del atributo
	valor  string // opSiIgual
	ruta   rutaXP
	cuerpo []pasoCadena
}

// hojaCadena es una hoja XSLT traducida: la plantilla de la raíz del
// documento y una plantilla por elemento.
type hojaCadena struct {
	nombre     string
	raiz       []pasoCadena
	plantillas map[nombreXML][]pasoCadena
	// integradas son los elementos a los que la hoja aplica plantillas por
	// nombre sin tener una (ImpRetenidos en Retenciones): XSLT les aplica la
	// integrada, que solo copia texto.
	integradas map[nombreXML]bool
}

// Constructores que usa cadena_reglas.go.

func txt(s string) pasoCadena { return pasoCadena{op: opTexto, texto: s} }
func req(a string) pasoCadena { return pasoCadena{op: opRequerido, texto: a} }
func opc(a string) pasoCadena { return pasoCadena{op: opOpcional, texto: a} }
func aplicar(r rutaXP) pasoCadena {
	return pasoCadena{op: opAplicar, ruta: r}
}
func cada(r rutaXP, cuerpo ...pasoCadena) pasoCadena {
	return pasoCadena{op: opCada, ruta: r, cuerpo: cuerpo}
}
func si(r rutaXP, cuerpo ...pasoCadena) pasoCadena {
	return pasoCadena{op: opSi, ruta: r, cuerpo: cuerpo}
}
func siIgual(a, v string, cuerpo ...pasoCadena) pasoCadena {
	return pasoCadena{op: opSiIgual, texto: a, valor: v, cuerpo: cuerpo}
}
func conjunto(r rutaXP, cuerpo ...pasoCadena) pasoCadena {
	return pasoCadena{op: opConjunto, ruta: r, cuerpo: cuerpo}
}
func rel(p ...pasoRuta) rutaXP       { return rutaXP{pasos: p} }
func abs(p ...pasoRuta) rutaXP       { return rutaXP{abs: true, pasos: p} }
func hijo(ns, local string) pasoRuta { return pasoRuta{nombre: nombreXML{ns, local}} }
func desc(ns, local string) pasoRuta { return pasoRuta{nombre: nombreXML{ns, local}, desc: true} }

// nodoC es un nodo del documento para la cadena: elemento con sus
// atributos tal cual, o texto (nombre vacío), con su posición en el
// documento.
type nodoC struct {
	nombre nombreXML
	attrs  []xml.Attr
	hijos  []*nodoC
	texto  string
	orden  int
}

func (n *nodoC) esElemento() bool { return n.nombre.local != "" }

// attrOk busca un atributo sin namespace por nombre exacto, como @Nombre en
// XPath (que distingue mayúsculas y no ve las declaraciones xmlns).
func (n *nodoC) attrOk(nombre string) (string, bool) {
	if n == nil {
		return "", false
	}
	for _, a := range n.attrs {
		if a.Name.Space == "" && a.Name.Local == nombre {
			return a.Value, true
		}
	}
	return "", false
}

func (n *nodoC) attr(nombre string) string {
	v, _ := n.attrOk(nombre)
	return v
}

// arbolCadena lee el documento entero en modo estricto. La Addenda se salta
// sin guardarla: ninguna hoja la lee. correccion no es "" si hubo que
// convertir bytes que no eran del encoding declarado o quitar caracteres
// prohibidos (lo que en Parse son ENCODING_CORREGIDO, ENCODING_DESCONOCIDO y
// CARACTER_INVALIDO).
func arbolCadena(b []byte) (raiz *nodoC, correccion string, err error) {
	b = normalizarCodificacion(b, func(_, msg string) {
		if correccion == "" {
			correccion = msg
		}
	})
	d := xml.NewDecoder(bytes.NewReader(b))
	d.CharsetReader = charsetIdentidad
	var pila []*nodoC
	var ns alcanceNS
	orden := 0
	for {
		tok, err := d.Token()
		if err == io.EOF && raiz != nil && len(pila) == 0 {
			return raiz, correccion, nil
		}
		if err != nil {
			return nil, "", &ErrorCadena{Motivo: MotivoXMLIlegible, Err: err}
		}
		switch t := tok.(type) {
		case xml.StartElement:
			if raiz != nil && len(pila) == 0 {
				return nil, "", &ErrorCadena{Motivo: MotivoXMLIlegible, Err: errors.New("más de un elemento raíz")}
			}
			if err := ns.entrar(t); err != nil {
				return nil, "", &ErrorCadena{Motivo: MotivoXMLIlegible, Err: err}
			}
			if len(pila) == 1 && t.Name.Local == "Addenda" {
				if err := ns.saltar(d); err != nil {
					return nil, "", &ErrorCadena{Motivo: MotivoXMLIlegible, Err: err}
				}
				continue
			}
			orden++
			n := &nodoC{nombre: nombreXML{t.Name.Space, t.Name.Local}, attrs: t.Attr, orden: orden}
			if len(pila) == 0 {
				raiz = n
			} else {
				p := pila[len(pila)-1]
				p.hijos = append(p.hijos, n)
			}
			pila = append(pila, n)
		case xml.EndElement:
			pila = pila[:len(pila)-1]
			ns.salir()
		case xml.CharData:
			if len(pila) == 0 {
				if len(bytes.Trim(t, " \t\r\n")) > 0 {
					return nil, "", &ErrorCadena{Motivo: MotivoXMLIlegible, Err: errors.New("texto fuera del elemento raíz")}
				}
				continue
			}
			orden++
			p := pila[len(pila)-1]
			p.hijos = append(p.hijos, &nodoC{texto: string(t), orden: orden})
		}
	}
}

// evaluar corre la hoja con doc como elemento raíz del documento.
func (h *hojaCadena) evaluar(doc *nodoC) (string, error) {
	ev := &evaluador{hoja: h, doc: doc}
	if err := ev.ejecutar(h.raiz, nil); err != nil {
		return "", err
	}
	return ev.b.String(), nil
}

type evaluador struct {
	hoja *hojaCadena
	doc  *nodoC
	b    strings.Builder
}

// ejecutar corre los pasos con ctx como nodo de contexto (nil = el nodo
// documento).
func (ev *evaluador) ejecutar(pasos []pasoCadena, ctx *nodoC) error {
	for _, p := range pasos {
		switch p.op {
		case opTexto:
			ev.b.WriteString(p.texto)
		case opRequerido, opOpcional:
			v, ok := ctx.attrOk(p.texto)
			ev.valor(p.op, v, ok)
		case opSiIgual:
			if v, ok := ctx.attrOk(p.texto); ok && v == p.valor {
				if err := ev.ejecutar(p.cuerpo, ctx); err != nil {
					return err
				}
			}
		case opSi:
			if len(ev.seleccionar(ctx, p.ruta)) > 0 {
				if err := ev.ejecutar(p.cuerpo, ctx); err != nil {
					return err
				}
			}
		case opCada:
			for _, n := range ev.seleccionar(ctx, p.ruta) {
				if err := ev.ejecutar(p.cuerpo, n); err != nil {
					return err
				}
			}
		case opAplicar:
			for _, n := range ev.seleccionar(ctx, p.ruta) {
				if err := ev.aplicar(n); err != nil {
					return err
				}
			}
		case opConjunto:
			// $Nodo/@attr es el conjunto de esos atributos de todos los
			// nodos; string() toma el primero en orden de documento, que es
			// el del primer nodo que lo tenga. El generador solo deja
			// Requerido y Opcional en este cuerpo.
			ns := ev.seleccionar(ctx, p.ruta)
			for _, q := range p.cuerpo {
				var v string
				var ok bool
				for _, n := range ns {
					if v, ok = n.attrOk(q.texto); ok {
						break
					}
				}
				ev.valor(q.op, v, ok)
			}
		}
	}
	return nil
}

// valor escribe un atributo como las plantillas Requerido (siempre "|", y
// el valor si lo hay) y Opcional (solo si el atributo existe, aunque esté
// vacío: test="$valor" de un node-set no vacío es verdadero).
func (ev *evaluador) valor(op opCadena, v string, ok bool) {
	if op == opOpcional && !ok {
		return
	}
	ev.b.WriteByte('|')
	ev.b.WriteString(normalizarEspacios(v))
}

// aplicar es xsl:apply-templates sobre un nodo: su plantilla, o la
// integrada de XSLT (copiar el texto) si es un elemento que las hojas dejan
// fuera a propósito. Cualquier otro elemento sin plantilla es un
// complemento que la librería no conoce: el SAT lo llevaría a la cadena, así
// que no se puede garantizar.
func (ev *evaluador) aplicar(n *nodoC) error {
	if !n.esElemento() {
		ev.b.WriteString(n.texto)
		return nil
	}
	if pasos, ok := ev.hoja.plantillas[n.nombre]; ok {
		return ev.ejecutar(pasos, n)
	}
	if !sinCadena[n.nombre] && !ev.hoja.integradas[n.nombre] {
		return &ErrorCadena{Motivo: MotivoComplementoSinCadena, Nodo: n.nombre.String(),
			Err: fmt.Errorf("la hoja %s no tiene plantilla para este nodo", ev.hoja.nombre)}
	}
	for _, h := range n.hijos {
		if err := ev.aplicar(h); err != nil {
			return err
		}
	}
	return nil
}

// seleccionar evalúa una ruta desde ctx, en orden de documento y sin
// repetidos, como un node-set de XPath.
func (ev *evaluador) seleccionar(ctx *nodoC, r rutaXP) []*nodoC {
	pasos := r.pasos
	var actual []*nodoC
	switch {
	case r.abs || ctx == nil:
		// El nodo documento tiene un solo hijo: el elemento raíz.
		if len(pasos) == 0 || !coincide(ev.doc, pasos[0]) {
			return nil
		}
		actual, pasos = []*nodoC{ev.doc}, pasos[1:]
	default:
		actual = []*nodoC{ctx}
	}
	for _, p := range pasos {
		var sig []*nodoC
		for _, n := range actual {
			if p.desc {
				sig = descendientes(n, p, sig)
				continue
			}
			for _, h := range n.hijos {
				if coincide(h, p) {
					sig = append(sig, h)
				}
			}
		}
		if p.desc && len(actual) > 1 {
			sig = ordenarUnicos(sig)
		}
		actual = sig
	}
	return actual
}

func coincide(n *nodoC, p pasoRuta) bool {
	return n.esElemento() && (p.nombre.local == "*" || n.nombre == p.nombre)
}

func descendientes(n *nodoC, p pasoRuta, out []*nodoC) []*nodoC {
	for _, h := range n.hijos {
		if coincide(h, p) {
			out = append(out, h)
		}
		out = descendientes(h, p, out)
	}
	return out
}

func ordenarUnicos(ns []*nodoC) []*nodoC {
	sort.Slice(ns, func(i, j int) bool { return ns[i].orden < ns[j].orden })
	out := ns[:0]
	for i, n := range ns {
		if i == 0 || n != ns[i-1] {
			out = append(out, n)
		}
	}
	return out
}

// normalizarEspacios es normalize-space de XPath 1.0: quita los espacios
// del principio y del final y junta las rachas en uno. Espacio es solo
// espacio, tabulador, CR y LF; un espacio no separable (U+00A0) se queda.
func normalizarEspacios(s string) string {
	return strings.Join(strings.FieldsFunc(s, func(r rune) bool {
		return r == ' ' || r == '\t' || r == '\r' || r == '\n'
	}), " ")
}
