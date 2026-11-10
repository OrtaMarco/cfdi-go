// Package gencadena traduce las hojas XSLT del SAT que arman la cadena
// original (CFDI 3.2, 3.3 y 4.0, TFD 1.0 y 1.1, y los complementos que lee
// la librería) a reglas Go: el archivo cadena_reglas.go del paquete cfdi.
//
// Go no tiene XSLT, pero esas hojas usan un subconjunto mínimo: plantillas
// por nombre de elemento, for-each, apply-templates, if, y las plantillas
// Requerido / Opcional de utilerias.xslt (normalize-space del atributo, o
// del texto de un elemento en Detallista, con "|" delante). Este paquete
// entiende exactamente eso y falla con cualquier
// otra construcción, para que un XSLT nuevo del SAT no se traduzca mal en
// silencio.
//
// Las hojas están copiadas sin cambios en xslt/sitio_internet (de
// phpcfdi/resources-sat-xml, que las baja del SAT). Para regenerar:
//
//	go run ./internal/gencadena/generar
package gencadena

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"go/format"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
)

const nsXSL = "http://www.w3.org/1999/XSL/Transform"

// Hojas que se generan: nombre de la variable Go y hoja raíz. Las rutas son
// relativas a sitio_internet; las de Retenciones están en
// www.sat.gob.mx/esquemas, al lado (../esquemas/…).
var Hojas = []struct{ Var, Raiz string }{
	{"hojaCFDI32", "cfd/3/cadenaoriginal_3_2/cadenaoriginal_3_2.xslt"},
	{"hojaCFDI33", "cfd/3/cadenaoriginal_3_3/cadenaoriginal_3_3.xslt"},
	{"hojaCFDI40", "cfd/4/cadenaoriginal_4_0/cadenaoriginal_4_0.xslt"},
	{"hojaTFD10", "timbrefiscaldigital/cadenaoriginal_TFD_1_0.xslt"},
	{"hojaTFD11", "cfd/TimbreFiscalDigital/cadenaoriginal_TFD_1_1.xslt"},
	{"hojaRetenciones10", "../esquemas/retencionpago/1/retenciones.xslt"},
	{"hojaRetenciones20", "../esquemas/retencionpago/2/retenciones.xslt"},
}

// retencion1 es el directorio de las hojas de Retenciones 1.0 y de sus
// complementos (que la 2.0 también incluye).
const retencion1 = "../esquemas/retencionpago/1/"

// Incluidos son los xsl:include que se siguen: las utilerías y los
// complementos que la librería lee. Los demás includes de las hojas raíz se
// omiten; si uno de esos complementos viene en un CFDI, la cadena devuelve
// error en vez de una cadena incompleta.
var Incluidos = map[string]bool{
	"cfd/2/cadenaoriginal_2_0/utilerias.xslt":                                  true,
	"cfd/Pagos/Pagos10.xslt":                                                   true,
	"cfd/Pagos/Pagos20.xslt":                                                   true,
	"cfd/nomina/nomina12.xslt":                                                 true,
	"cfd/implocal/implocal.xslt":                                               true,
	"cfd/CartaPorte/CartaPorte20.xslt":                                         true,
	"cfd/CartaPorte/CartaPorte30.xslt":                                         true,
	"cfd/CartaPorte/CartaPorte31.xslt":                                         true,
	"cfd/ComercioExterior11/ComercioExterior11.xslt":                           true,
	"cfd/ComercioExterior20/ComercioExterior20.xslt":                           true,
	"cfd/divisas/divisas.xslt":                                                 true,
	"cfd/cfdiregistrofiscal/cfdiregistrofiscal.xslt":                           true,
	"cfd/leyendasFiscales/leyendasFisc.xslt":                                   true,
	"cfd/EstadoDeCuentaCombustible/ecc12.xslt":                                 true,
	"cfd/terceros/terceros11.xslt":                                             true,
	"cfd/ventavehiculos/ventavehiculos11.xslt":                                 true,
	"cfd/HidrocarburosPetro/hidrocarburospetroliferos.xslt":                    true,
	"cfd/aerolineas/aerolineas.xslt":                                           true,
	"cfd/donat/donat11.xslt":                                                   true,
	"cfd/iedu/iedu.xslt":                                                       true,
	"cfd/CartaPorte/CartaPorte.xslt":                                           true,
	"cfd/ComercioExterior/ComercioExterior10.xslt":                             true,
	"cfd/EstadoDeCuentaCombustible/ecc11.xslt":                                 true,
	"cfd/nomina/nomina11.xslt":                                                 true,
	"cfd/ine/ine10.xslt":                                                       true,
	"cfd/ine/ine11.xslt":                                                       true,
	"cfd/valesdedespensa/valesdedespensa.xslt":                                 true,
	"cfd/consumodecombustibles/consumodecombustibles.xslt":                     true,
	"cfd/consumodecombustibles/consumodeCombustibles11.xslt":                   true,
	"cfd/pagoenespecie/pagoenespecie.xslt":                                     true,
	"cfd/notariospublicos/notariospublicos.xslt":                               true,
	"cfd/servicioparcialconstruccion/servicioparcialconstruccion.xslt":         true,
	"cfd/vehiculousado/vehiculousado.xslt":                                     true,
	"cfd/TuristaPasajeroExtranjero/TuristaPasajeroExtranjero.xslt":             true,
	"cfd/acreditamiento/AcreditamientoIEPS10.xslt":                             true,
	"cfd/spei/spei.xslt":                                                       true,
	"cfd/renovacionysustitucionvehiculos/renovacionysustitucionvehiculos.xslt": true,
	"cfd/certificadodestruccion/certificadodedestruccion.xslt":                 true,
	"cfd/arteantiguedades/obrasarteantiguedades.xslt":                          true,
	"cfd/pfic/pfic.xslt":                                                       true,
	"cfd/GastosHidrocarburos10/GastosHidrocarburos10.xslt":                     true,
	"cfd/IngresosHidrocarburos10/IngresosHidrocarburos.xslt":                   true,
	"cfd/ecb/ecb.xslt":                                                         true,
	"cfd/ecc/ecc.xslt":                                                         true,
	"cfd/detallista/detallista.xslt":                                           true,

	// Complementos de Retenciones: todos los que incluyen las hojas 1.0 y
	// 2.0 son chicos y se llevan a la cadena.
	retencion1 + "arrendamientoenfideicomiso/arrendamientoenfideicomiso.xslt":        true,
	retencion1 + "dividendos/dividendos.xslt":                                        true,
	retencion1 + "enajenaciondeacciones/enajenaciondeacciones.xslt":                  true,
	retencion1 + "fideicomisonoempresarial/fideicomisonoempresarial.xslt":            true,
	retencion1 + "intereses/intereses.xslt":                                          true,
	retencion1 + "intereseshipotecarios/intereseshipotecarios.xslt":                  true,
	retencion1 + "pagosaextranjeros/pagosaextranjeros.xslt":                          true,
	retencion1 + "planesderetiro/planesderetiro.xslt":                                true,
	retencion1 + "premios/premios.xslt":                                              true,
	retencion1 + "operacionesconderivados/operacionesconderivados.xslt":              true,
	retencion1 + "sectorfinanciero/sectorfinanciero.xslt":                            true,
	retencion1 + "planesderetiro11/planesderetiro11.xslt":                            true,
	retencion1 + "PlataformasTecnologicas10/ServiciosPlataformasTecnologicas10.xslt": true,
}

// normalizarRuta deja una ruta relativa a sitio_internet: las hojas de
// Retenciones incluyen las utilerías como ../../../sitio_internet/cfd/2/…,
// que desde ../esquemas/retencionpago/2 es ../sitio_internet/cfd/2/…, el
// mismo archivo que cfd/2/….
func normalizarRuta(p string) string {
	p = path.Clean(p)
	if r, ok := strings.CutPrefix(p, "../sitio_internet/"); ok {
		return r
	}
	return p
}

// Firmas esperadas de las plantillas de utilerías: si el SAT las cambia, el
// generador falla en vez de seguir traduciendo Requerido y Opcional como
// antes.
var firmasUtilerias = map[string]string{
	"Requerido":      `template{name=Requerido}[param{name=valor}[]"|"call-template{name=ManejaEspacios}[with-param{name=s,select=$valor}[]]]`,
	"Opcional":       `template{name=Opcional}[param{name=valor}[]if{test=$valor}["|"call-template{name=ManejaEspacios}[with-param{name=s,select=$valor}[]]]]`,
	"ManejaEspacios": `template{name=ManejaEspacios}[param{name=s}[]value-of{select=normalize-space(string($s))}[]]`,
}

// elem es un nodo del XSLT con los namespaces en alcance.
type elem struct {
	ns, local string
	attrs     map[string]string // solo los atributos sin namespace
	hijos     []any             // *elem o string
	scope     map[string]string // prefijo → URI
	archivo   string
}

func leer(dir, rel string) (*elem, []byte, error) {
	b, err := os.ReadFile(path.Join(dir, rel))
	if err != nil {
		return nil, nil, err
	}
	d := xml.NewDecoder(bytes.NewReader(b))
	var pila []*elem
	var raiz *elem
	for {
		tok, err := d.Token()
		if err != nil {
			if raiz != nil && len(pila) == 0 {
				return raiz, b, nil
			}
			return nil, nil, fmt.Errorf("%s: %w", rel, err)
		}
		switch t := tok.(type) {
		case xml.StartElement:
			e := &elem{ns: t.Name.Space, local: t.Name.Local, attrs: map[string]string{}, scope: map[string]string{}, archivo: rel}
			if len(pila) > 0 {
				p := pila[len(pila)-1]
				for k, v := range p.scope {
					e.scope[k] = v
				}
				p.hijos = append(p.hijos, e)
			} else {
				raiz = e
			}
			for _, a := range t.Attr {
				switch {
				case a.Name.Space == "xmlns":
					e.scope[a.Name.Local] = a.Value
				case a.Name.Space == "" && a.Name.Local == "xmlns":
					e.scope[""] = a.Value
				case a.Name.Space == "":
					e.attrs[a.Name.Local] = a.Value
				default:
					return nil, nil, fmt.Errorf("%s: atributo con namespace %s:%s no soportado", rel, a.Name.Space, a.Name.Local)
				}
			}
			pila = append(pila, e)
		case xml.EndElement:
			pila = pila[:len(pila)-1]
		case xml.CharData:
			if len(pila) > 0 {
				p := pila[len(pila)-1]
				p.hijos = append(p.hijos, string(t))
			}
		}
	}
}

func (e *elem) xsl(local string) bool { return e.ns == nsXSL && e.local == local }

func (e *elem) donde() string { return e.archivo + ": <" + e.local + ">" }

// soloAttrs falla si e tiene atributos fuera de los permitidos.
func (e *elem) soloAttrs(permitidos ...string) error {
	for k := range e.attrs {
		ok := false
		for _, p := range permitidos {
			ok = ok || k == p
		}
		if !ok {
			return fmt.Errorf("%s: atributo %q no soportado", e.donde(), k)
		}
	}
	return nil
}

// elementos devuelve los hijos elemento y falla si hay texto que no sea
// espacio en blanco.
func (e *elem) elementos() ([]*elem, error) {
	var out []*elem
	for _, h := range e.hijos {
		switch v := h.(type) {
		case *elem:
			out = append(out, v)
		case string:
			if strings.TrimSpace(v) != "" {
				return nil, fmt.Errorf("%s: texto %q no esperado", e.donde(), v)
			}
		}
	}
	return out, nil
}

// firma serializa un elemento XSLT de forma canónica (atributos ordenados,
// sin el texto que es solo espacio, que XSLT descarta).
func firma(e *elem) string {
	var b strings.Builder
	b.WriteString(e.local)
	if e.ns != nsXSL {
		b.WriteString("!no-xsl")
	}
	claves := make([]string, 0, len(e.attrs))
	for k := range e.attrs {
		claves = append(claves, k)
	}
	sort.Strings(claves)
	b.WriteString("{")
	for i, k := range claves {
		if i > 0 {
			b.WriteString(",")
		}
		b.WriteString(k + "=" + e.attrs[k])
	}
	b.WriteString("}[")
	for _, h := range e.hijos {
		switch v := h.(type) {
		case *elem:
			b.WriteString(firma(v))
		case string:
			if strings.TrimSpace(v) != "" {
				b.WriteString(strconv.Quote(v))
			}
		}
	}
	b.WriteString("]")
	return b.String()
}

// Paso es una instrucción ya traducida.
type Paso struct {
	Op     string // texto, requerido, opcional, requerido_nodo, opcional_nodo, cada, aplicar, si, si_igual, conjunto
	Texto  string // texto literal, o nombre del atributo
	Valor  string // si_igual
	Abs    bool
	Ruta   []PasoRuta
	Cuerpo []Paso
}

// PasoRuta es un paso de una ruta XPath: hijo, o descendiente si Desc, con
// ese nombre ("*" = cualquier elemento).
type PasoRuta struct {
	NS, Local string
	Desc      bool
}

type plantilla struct {
	ns, local string
	e         *elem
	pasos     []Paso
}

type hoja struct {
	variable, raizArchivo string
	archivos              []string
	omitidos              []string
	raiz                  *elem
	pasosRaiz             []Paso
	plantillas            []*plantilla
	nombradas             map[string]*elem
	integradas            []PasoRuta // ver buscarIntegradas
	sombreadas            []string   // hojas incluidas con su propio match="/" (ver registrar)
}

// Generar lee las hojas bajo dir (el directorio sitio_internet) y devuelve
// el código Go de cadena_reglas.go ya formateado.
func Generar(dir string) ([]byte, error) {
	var b bytes.Buffer
	b.WriteString("// Código generado por `go run ./internal/gencadena/generar` a partir de los\n")
	b.WriteString("// XSLT del SAT en internal/gencadena/xslt. NO EDITAR: cambiar el generador.\n\n")
	b.WriteString("package cfdi\n")
	for _, hv := range Hojas {
		h := &hoja{variable: hv.Var, raizArchivo: hv.Raiz, nombradas: map[string]*elem{}}
		hashes := map[string]string{}
		if err := h.cargar(dir, hv.Raiz, true, hashes, map[string]bool{}); err != nil {
			return nil, err
		}
		if err := h.compilarTodo(); err != nil {
			return nil, err
		}
		h.escribir(&b, hashes)
	}
	out, err := format.Source(b.Bytes())
	if err != nil {
		return nil, fmt.Errorf("gofmt del código generado: %w", err)
	}
	return out, nil
}

// cargar lee rel y, recursivamente, sus includes permitidos.
func (h *hoja) cargar(dir, rel string, esRaiz bool, hashes map[string]string, vistos map[string]bool) error {
	if vistos[rel] {
		return nil
	}
	vistos[rel] = true
	raiz, crudo, err := leer(dir, rel)
	if err != nil {
		return err
	}
	sum := sha256.Sum256(crudo)
	hashes[rel] = hex.EncodeToString(sum[:])
	h.archivos = append(h.archivos, rel)
	if !raiz.xsl("stylesheet") {
		return fmt.Errorf("%s: la raíz no es xsl:stylesheet", rel)
	}
	hijos, err := raiz.elementos()
	if err != nil {
		return err
	}
	for _, e := range hijos {
		switch {
		case e.xsl("include"):
			inc := normalizarRuta(path.Join(path.Dir(rel), e.attrs["href"]))
			if !Incluidos[inc] {
				h.omitidos = append(h.omitidos, inc)
				continue
			}
			if err := h.cargar(dir, inc, false, hashes, vistos); err != nil {
				return err
			}
		case e.xsl("output"):
			if e.attrs["method"] != "text" {
				return fmt.Errorf("%s: xsl:output method=%q, se esperaba text", rel, e.attrs["method"])
			}
		case e.xsl("template"):
			if err := h.registrar(e, esRaiz); err != nil {
				return err
			}
		default:
			return fmt.Errorf("%s: elemento de primer nivel no soportado", e.donde())
		}
	}
	return nil
}

func (h *hoja) registrar(e *elem, esRaiz bool) error {
	if err := e.soloAttrs("match", "name"); err != nil {
		return err
	}
	m, conMatch := e.attrs["match"]
	n, conNombre := e.attrs["name"]
	switch {
	case conMatch && conNombre:
		return fmt.Errorf("%s: plantilla con match y name", e.donde())
	case conNombre:
		if h.nombradas[n] != nil {
			return fmt.Errorf("%s: plantilla %q repetida", e.donde(), n)
		}
		if want, ok := firmasUtilerias[n]; ok && firma(e) != want {
			return fmt.Errorf("%s: la plantilla %s cambió:\n  %s\nse esperaba\n  %s", e.donde(), n, firma(e), want)
		}
		h.nombradas[n] = e
		return nil
	}
	m = strings.TrimSpace(m)
	if m == "/" {
		if esRaiz {
			if h.raiz != nil {
				return fmt.Errorf("%s: match=\"/\" repetido", e.donde())
			}
			h.raiz = e
			return nil
		}
		// vehiculousado.xslt trae su propio match="/" (para usarla sola).
		// Con dos plantillas para "/", XSLT usa la última en orden de
		// documento: la de la hoja raíz, que va después de sus includes. Si
		// la raíz ya tenía la suya, ganaría la incluida: no se soporta.
		if h.raiz != nil {
			return fmt.Errorf("%s: match=\"/\" en una hoja incluida después del de la raíz", e.donde())
		}
		h.sombreadas = append(h.sombreadas, e.archivo)
		return nil
	}
	ns, local, err := nombre(e, m)
	if err != nil {
		return err
	}
	if local == "*" {
		return fmt.Errorf("%s: match=%q no soportado", e.donde(), m)
	}
	for _, p := range h.plantillas {
		if p.ns == ns && p.local == local {
			// La hoja de 3.2 incluye ine10.xslt e ine11.xslt, que son
			// idénticas: libxslt usa la última, y da lo mismo. Una repetida
			// distinta sí es un error.
			if firma(p.e) == firma(e) {
				return nil
			}
			return fmt.Errorf("%s: plantilla para {%s}%s repetida y distinta (ya en %s)", e.donde(), ns, local, p.e.archivo)
		}
	}
	h.plantillas = append(h.plantillas, &plantilla{ns: ns, local: local, e: e})
	return nil
}

// nombre resuelve "prefijo:Local" con los namespaces de e. "*" es
// cualquier elemento. Un nombre sin prefijo (namespace vacío) no se usa en
// estas hojas y se rechaza.
func nombre(e *elem, s string) (ns, local string, err error) {
	if s == "*" {
		return "", "*", nil
	}
	pre, loc, ok := strings.Cut(s, ":")
	if !ok || pre == "" || loc == "" || strings.ContainsAny(loc, "[]()@$=':/|* ") {
		return "", "", fmt.Errorf("%s: nombre %q no soportado", e.donde(), s)
	}
	uri, ok := e.scope[pre]
	if !ok {
		return "", "", fmt.Errorf("%s: prefijo %q sin declarar", e.donde(), pre)
	}
	return uri, loc, nil
}

// ruta traduce una expresión de nodos: ".", "./a:B/a:C", ".//a:B",
// "a:B", "/a:B", "./*".
func ruta(e *elem, sel string) (abs bool, pasos []PasoRuta, err error) {
	s := strings.TrimSpace(sel)
	if s == "." {
		return false, nil, nil
	}
	desc := false
	switch {
	case strings.HasPrefix(s, ".//"):
		s, desc = s[3:], true
	case strings.HasPrefix(s, "./"):
		s = s[2:]
	case strings.HasPrefix(s, "/"):
		s, abs = s[1:], true
	}
	for _, parte := range strings.Split(s, "/") {
		if parte == "" {
			if desc {
				return false, nil, fmt.Errorf("%s: ruta %q no soportada", e.donde(), sel)
			}
			desc = true
			continue
		}
		ns, local, err := nombre(e, parte)
		if err != nil {
			return false, nil, fmt.Errorf("ruta %q: %w", sel, err)
		}
		pasos = append(pasos, PasoRuta{NS: ns, Local: local, Desc: desc})
		desc = false
	}
	if desc || len(pasos) == 0 {
		return false, nil, fmt.Errorf("%s: ruta %q no soportada", e.donde(), sel)
	}
	return abs, pasos, nil
}

// atributo traduce "./@Nombre" (o "@Nombre"), o "$param/@Nombre" dentro de
// una plantilla con nombre.
func atributo(e *elem, sel, param string) (string, error) {
	s := strings.TrimSpace(sel)
	var resto string
	var ok bool
	if param != "" {
		resto, ok = strings.CutPrefix(s, "$"+param+"/@")
	} else if resto, ok = strings.CutPrefix(s, "./@"); !ok {
		resto, ok = strings.CutPrefix(s, "@")
	}
	if !ok || resto == "" || strings.ContainsAny(resto, "/[]()$=': ") {
		return "", fmt.Errorf("%s: valor %q no soportado (se esperaba un atributo)", e.donde(), sel)
	}
	return resto, nil
}

func (h *hoja) compilarTodo() error {
	if h.raiz == nil {
		return fmt.Errorf("%s: sin plantilla match=\"/\"", h.raizArchivo)
	}
	var err error
	if h.pasosRaiz, err = h.compilar(h.raiz, ""); err != nil {
		return err
	}
	for _, p := range h.plantillas {
		if p.pasos, err = h.compilar(p.e, ""); err != nil {
			return err
		}
	}
	h.buscarIntegradas()
	return nil
}

// buscarIntegradas junta los elementos a los que la hoja aplica plantillas
// por su nombre sin tener plantilla para ellos: XSLT les aplica la integrada,
// que copia el texto de sus descendientes. Pasa en Retenciones, que hace
// apply-templates sobre cada ImpRetenidos y lee sus atributos fuera. Los que
// se alcanzan con "*" (el contenido de Complemento) no cuentan: sin
// plantilla son complementos que la librería no conoce.
func (h *hoja) buscarIntegradas() {
	con := map[PasoRuta]bool{}
	for _, p := range h.plantillas {
		con[PasoRuta{NS: p.ns, Local: p.local}] = true
	}
	vistos := map[PasoRuta]bool{}
	ultimo := func(r []PasoRuta, ctx *PasoRuta) *PasoRuta {
		if len(r) == 0 {
			return ctx
		}
		u := r[len(r)-1]
		u.Desc = false
		return &u
	}
	var recorrer func(pasos []Paso, ctx *PasoRuta)
	recorrer = func(pasos []Paso, ctx *PasoRuta) {
		for _, p := range pasos {
			switch p.Op {
			case "cada":
				recorrer(p.Cuerpo, ultimo(p.Ruta, ctx))
			case "si", "si_igual":
				recorrer(p.Cuerpo, ctx)
			case "aplicar":
				if t := ultimo(p.Ruta, ctx); t != nil && t.Local != "*" && !con[*t] && !vistos[*t] {
					vistos[*t] = true
					h.integradas = append(h.integradas, *t)
				}
			}
		}
	}
	recorrer(h.pasosRaiz, nil)
	for _, p := range h.plantillas {
		recorrer(p.pasos, &PasoRuta{NS: p.ns, Local: p.local})
	}
}

// compilar traduce el contenido de una plantilla, for-each o if. param es
// el parámetro de la plantilla con nombre que se está expandiendo ("" si
// no).
func (h *hoja) compilar(c *elem, param string) ([]Paso, error) {
	var out []Paso
	for _, hijo := range c.hijos {
		if s, ok := hijo.(string); ok {
			if strings.TrimSpace(s) != "" {
				out = append(out, Paso{Op: "texto", Texto: s})
			}
			continue
		}
		e := hijo.(*elem)
		if e.ns != nsXSL {
			return nil, fmt.Errorf("%s: elemento literal no soportado", e.donde())
		}
		if e.local == "choose" {
			ps, err := h.eleccion(e, param)
			if err != nil {
				return nil, err
			}
			out = append(out, ps...)
			continue
		}
		p, err := h.instruccion(e, param)
		if err != nil {
			return nil, err
		}
		if p != nil {
			out = append(out, *p)
		}
	}
	return out, nil
}

func (h *hoja) instruccion(e *elem, param string) (*Paso, error) {
	switch e.local {
	case "param":
		// Solo como primer hijo de una plantilla con nombre, y ya se
		// resolvió al llamarla.
		if param == "" || e.attrs["name"] != param {
			return nil, fmt.Errorf("%s: xsl:param fuera de lugar", e.donde())
		}
		return nil, nil
	case "call-template":
		return h.llamada(e, param)
	case "for-each", "apply-templates":
		if err := e.soloAttrs("select"); err != nil {
			return nil, err
		}
		if param != "" {
			return nil, fmt.Errorf("%s: %s dentro de una plantilla con parámetro", e.donde(), e.local)
		}
		abs, r, err := ruta(e, e.attrs["select"])
		if err != nil {
			return nil, err
		}
		if e.local == "apply-templates" {
			if hs, err := e.elementos(); err != nil || len(hs) > 0 {
				return nil, fmt.Errorf("%s: apply-templates con contenido", e.donde())
			}
			return &Paso{Op: "aplicar", Abs: abs, Ruta: r}, nil
		}
		cuerpo, err := h.compilar(e, "")
		if err != nil {
			return nil, err
		}
		return &Paso{Op: "cada", Abs: abs, Ruta: r, Cuerpo: cuerpo}, nil
	case "if":
		if err := e.soloAttrs("test"); err != nil {
			return nil, err
		}
		if param != "" {
			return nil, fmt.Errorf("%s: if dentro de una plantilla con parámetro", e.donde())
		}
		cuerpo, err := h.compilar(e, "")
		if err != nil {
			return nil, err
		}
		test := strings.TrimSpace(e.attrs["test"])
		if strings.Contains(test, "=") {
			a, v, err := igualdad(e, test)
			if err != nil {
				return nil, err
			}
			return &Paso{Op: "si_igual", Texto: a, Valor: v, Cuerpo: cuerpo}, nil
		}
		abs, r, err := ruta(e, test)
		if err != nil {
			return nil, err
		}
		return &Paso{Op: "si", Abs: abs, Ruta: r, Cuerpo: cuerpo}, nil
	}
	return nil, fmt.Errorf("%s: instrucción no soportada", e.donde())
}

// igualdad traduce un test "./@Atributo='valor'".
func igualdad(e *elem, test string) (attr, valor string, err error) {
	izq, der, _ := strings.Cut(test, "=")
	v, ok := strings.CutPrefix(der, "'")
	v, ok2 := strings.CutSuffix(v, "'")
	a, err := atributo(e, izq, "")
	if !ok || !ok2 || err != nil || strings.Contains(v, "'") {
		return "", "", fmt.Errorf("%s: test %q no soportado", e.donde(), test)
	}
	return a, v, nil
}

// eleccion traduce un xsl:choose cuyos xsl:when comparan todos el mismo
// atributo con valores distintos (nomina11.xslt: ./@Version='1.0' y '1.1'),
// sin xsl:otherwise. Así a lo más uno se cumple, y el choose es lo mismo que
// un if por cada when, en orden. Cualquier otra forma se rechaza.
func (h *hoja) eleccion(e *elem, param string) ([]Paso, error) {
	if err := e.soloAttrs(); err != nil {
		return nil, err
	}
	if param != "" {
		return nil, fmt.Errorf("%s: choose dentro de una plantilla con parámetro", e.donde())
	}
	whens, err := e.elementos()
	if err != nil {
		return nil, err
	}
	if len(whens) == 0 {
		return nil, fmt.Errorf("%s: choose sin when", e.donde())
	}
	var out []Paso
	vistos := map[string]bool{}
	for _, w := range whens {
		if !w.xsl("when") {
			return nil, fmt.Errorf("%s: dentro de choose solo se soporta when (sin otherwise)", w.donde())
		}
		if err := w.soloAttrs("test"); err != nil {
			return nil, err
		}
		a, v, err := igualdad(w, strings.TrimSpace(w.attrs["test"]))
		if err != nil {
			return nil, err
		}
		if len(out) > 0 && a != out[0].Texto {
			return nil, fmt.Errorf("%s: los when de un choose comparan atributos distintos (%s y %s)", w.donde(), out[0].Texto, a)
		}
		if vistos[v] {
			return nil, fmt.Errorf("%s: dos when con el mismo valor %q", w.donde(), v)
		}
		vistos[v] = true
		cuerpo, err := h.compilar(w, "")
		if err != nil {
			return nil, err
		}
		out = append(out, Paso{Op: "si_igual", Texto: a, Valor: v, Cuerpo: cuerpo})
	}
	return out, nil
}

func (h *hoja) llamada(e *elem, param string) (*Paso, error) {
	if err := e.soloAttrs("name"); err != nil {
		return nil, err
	}
	args, err := e.elementos()
	if err != nil {
		return nil, err
	}
	if len(args) != 1 || !args[0].xsl("with-param") {
		return nil, fmt.Errorf("%s: se esperaba un solo xsl:with-param", e.donde())
	}
	arg := args[0]
	if err := arg.soloAttrs("name", "select"); err != nil {
		return nil, err
	}
	n := e.attrs["name"]
	t := h.nombradas[n]
	if t == nil {
		return nil, fmt.Errorf("%s: plantilla %q no definida", e.donde(), n)
	}
	switch n {
	case "Requerido", "Opcional":
		if arg.attrs["name"] != "valor" {
			return nil, fmt.Errorf("%s: parámetro %q", e.donde(), arg.attrs["name"])
		}
		op := "requerido"
		if n == "Opcional" {
			op = "opcional"
		}
		sel := arg.attrs["select"]
		a, err := atributo(arg, sel, param)
		if err == nil {
			return &Paso{Op: op, Texto: a}, nil
		}
		i := strings.LastIndex(sel, "/@")
		if param == "" && !strings.Contains(sel, "@") {
			// "." o "a:B/a:C" (Detallista): el valor es el texto de un
			// elemento. string() de un node-set es el valor de texto del
			// primero en orden de documento: todo el texto de sus
			// descendientes.
			if abs, r, err2 := ruta(arg, sel); err2 == nil {
				return &Paso{Op: op + "_nodo", Abs: abs, Ruta: r}, nil
			}
			return nil, err
		}
		// "./a:B/@Attr" (el Receptor de Retenciones): el valor es el conjunto
		// de esos atributos de los nodos B, y string() toma el primero en
		// orden de documento, como $param/@attr en una plantilla con nombre.
		if param != "" || i < 0 {
			return nil, err
		}
		abs, r, err2 := ruta(arg, sel[:i])
		if err2 != nil {
			return nil, err
		}
		if a, err = atributo(arg, "@"+sel[i+2:], ""); err != nil {
			return nil, err
		}
		return &Paso{Op: "conjunto", Abs: abs, Ruta: r, Cuerpo: []Paso{{Op: op, Texto: a}}}, nil
	case "ManejaEspacios":
		return nil, fmt.Errorf("%s: ManejaEspacios fuera de Requerido/Opcional", e.donde())
	}
	// Plantilla con nombre y un parámetro que es un conjunto de nodos
	// (Domicilio de CFDI 3.2): su cuerpo solo puede leer $param/@attr, que
	// en XPath 1.0 es el atributo del primer nodo del conjunto que lo tenga
	// (string() de un node-set toma el primero en orden de documento).
	if param != "" {
		return nil, fmt.Errorf("%s: llamada anidada a %q", e.donde(), n)
	}
	hs, err := t.elementos()
	if err != nil {
		return nil, err
	}
	if len(hs) == 0 || !hs[0].xsl("param") || hs[0].attrs["name"] != arg.attrs["name"] {
		return nil, fmt.Errorf("%s: la plantilla %q no declara el parámetro %q", t.donde(), n, arg.attrs["name"])
	}
	abs, r, err := ruta(arg, arg.attrs["select"])
	if err != nil {
		return nil, err
	}
	cuerpo, err := h.compilar(t, arg.attrs["name"])
	if err != nil {
		return nil, err
	}
	return &Paso{Op: "conjunto", Abs: abs, Ruta: r, Cuerpo: cuerpo}, nil
}

func (h *hoja) escribir(b *bytes.Buffer, hashes map[string]string) {
	fmt.Fprintf(b, "\n// %s sale de %s y de lo que incluye.\n//\n// Archivos (sha256):\n", h.variable, h.raizArchivo)
	for _, a := range h.archivos {
		fmt.Fprintf(b, "//   - %s %s\n", a, hashes[a])
	}
	if len(h.omitidos) > 0 {
		b.WriteString("//\n// Includes omitidos (complementos que la librería no lee):\n")
		for _, a := range h.omitidos {
			fmt.Fprintf(b, "//   - %s\n", a)
		}
	}
	if len(h.sombreadas) > 0 {
		b.WriteString("//\n// Plantillas match=\"/\" de hojas incluidas, que la de la raíz sustituye:\n")
		for _, a := range h.sombreadas {
			fmt.Fprintf(b, "//   - %s\n", a)
		}
	}
	fmt.Fprintf(b, "var %s = &hojaCadena{\n", h.variable)
	fmt.Fprintf(b, "nombre: %q,\n", strings.TrimSuffix(path.Base(h.raizArchivo), ".xslt"))
	b.WriteString("raiz: []pasoCadena{\n")
	escribirPasos(b, h.pasosRaiz)
	b.WriteString("},\n")
	b.WriteString("plantillas: map[nombreXML][]pasoCadena{\n")
	for _, p := range h.plantillas {
		fmt.Fprintf(b, "// %s\n{%q, %q}: {\n", p.e.archivo, p.ns, p.local)
		escribirPasos(b, p.pasos)
		b.WriteString("},\n")
	}
	b.WriteString("},\n")
	if len(h.integradas) > 0 {
		b.WriteString("integradas: map[nombreXML]bool{\n")
		for _, t := range h.integradas {
			fmt.Fprintf(b, "{%q, %q}: true,\n", t.NS, t.Local)
		}
		b.WriteString("},\n")
	}
	b.WriteString("}\n")
}

func escribirPasos(b *bytes.Buffer, pasos []Paso) {
	for _, p := range pasos {
		escribirPaso(b, p)
		b.WriteString(",\n")
	}
}

func escribirPaso(b *bytes.Buffer, p Paso) {
	switch p.Op {
	case "texto":
		fmt.Fprintf(b, "txt(%q)", p.Texto)
	case "requerido":
		fmt.Fprintf(b, "req(%q)", p.Texto)
	case "opcional":
		fmt.Fprintf(b, "opc(%q)", p.Texto)
	case "requerido_nodo", "opcional_nodo":
		f := "reqNodo("
		if p.Op == "opcional_nodo" {
			f = "opcNodo("
		}
		b.WriteString(f)
		escribirRuta(b, p)
		b.WriteString(")")
	case "aplicar":
		b.WriteString("aplicar(")
		escribirRuta(b, p)
		b.WriteString(")")
	case "si_igual":
		fmt.Fprintf(b, "siIgual(%q, %q", p.Texto, p.Valor)
		escribirCuerpo(b, p.Cuerpo)
	default: // cada, si, conjunto
		b.WriteString(p.Op + "(")
		escribirRuta(b, p)
		escribirCuerpo(b, p.Cuerpo)
	}
}

func escribirCuerpo(b *bytes.Buffer, cuerpo []Paso) {
	if len(cuerpo) == 0 {
		b.WriteString(")")
		return
	}
	b.WriteString(",\n")
	escribirPasos(b, cuerpo)
	b.WriteString(")")
}

func escribirRuta(b *bytes.Buffer, p Paso) {
	f := "rel"
	if p.Abs {
		f = "abs"
	}
	b.WriteString(f + "(")
	for i, r := range p.Ruta {
		if i > 0 {
			b.WriteString(", ")
		}
		paso := "hijo"
		if r.Desc {
			paso = "desc"
		}
		fmt.Fprintf(b, "%s(%q, %q)", paso, r.NS, r.Local)
	}
	b.WriteString(")")
}
