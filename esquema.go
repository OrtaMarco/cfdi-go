package cfdi

//go:generate go run ./internal/genxsd/generar

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"math/big"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/OrtaMarco/cfdi-go/decimal"
)

// Códigos de hallazgo de ValidarEsquema. Son parte de la API pública: no se
// renombran.
const (
	EsquemaXMLIlegible         = "ESQUEMA_XML_ILEGIBLE"          // un procesador XML conforme no lo leería (mal formado, encoding que miente…)
	EsquemaNoSoportado         = "ESQUEMA_NO_SOPORTADO"          // aviso: raíz o complemento sin esquema en la librería; esa parte no se validó
	EsquemaElementoNoPermitido = "ESQUEMA_ELEMENTO_NO_PERMITIDO" // elemento que el esquema no admite en ese lugar
	EsquemaElementoFaltante    = "ESQUEMA_ELEMENTO_FALTANTE"     // falta un elemento requerido (minOccurs)
	EsquemaElementoExcedente   = "ESQUEMA_ELEMENTO_EXCEDENTE"    // un elemento se repite más de lo que permite maxOccurs
	EsquemaOrden               = "ESQUEMA_ORDEN"                 // elemento fuera del orden de la secuencia
	EsquemaContenido           = "ESQUEMA_CONTENIDO"             // texto donde el esquema no admite texto
	EsquemaAtributoFaltante    = "ESQUEMA_ATRIBUTO_FALTANTE"     // falta un atributo requerido
	EsquemaAtributoNoPermitido = "ESQUEMA_ATRIBUTO_NO_PERMITIDO" // atributo que el esquema no declara
	EsquemaTipo                = "ESQUEMA_TIPO"                  // valor que no es del tipo XSD (decimal, fecha, entero…)
	EsquemaPatron              = "ESQUEMA_PATRON"                // valor que no cumple el pattern
	EsquemaCatalogo            = "ESQUEMA_CATALOGO"              // valor fuera del catálogo (enumeration)
	EsquemaRango               = "ESQUEMA_RANGO"                 // valor fuera de minInclusive/maxInclusive
	EsquemaDigitos             = "ESQUEMA_DIGITOS"               // más decimales (fractionDigits) o dígitos (totalDigits) de los permitidos
	EsquemaLongitud            = "ESQUEMA_LONGITUD"              // longitud fuera de length/minLength/maxLength
	EsquemaValorFijo           = "ESQUEMA_VALOR_FIJO"            // valor distinto del fixed del esquema
)

// CodigosEsquema lista los códigos de ValidarEsquema, en el orden de
// arriba.
var CodigosEsquema = []string{
	EsquemaXMLIlegible, EsquemaNoSoportado, EsquemaElementoNoPermitido, EsquemaElementoFaltante,
	EsquemaElementoExcedente, EsquemaOrden, EsquemaContenido, EsquemaAtributoFaltante,
	EsquemaAtributoNoPermitido, EsquemaTipo, EsquemaPatron, EsquemaCatalogo, EsquemaRango,
	EsquemaDigitos, EsquemaLongitud, EsquemaValorFijo,
}

// codigoFaceta da el código de hallazgo de cada faceta XSD.
var codigoFaceta = map[string]string{
	"pattern":        EsquemaPatron,
	"enumeration":    EsquemaCatalogo,
	"minInclusive":   EsquemaRango,
	"maxInclusive":   EsquemaRango,
	"fractionDigits": EsquemaDigitos,
	"totalDigits":    EsquemaDigitos,
	"length":         EsquemaLongitud,
	"minLength":      EsquemaLongitud,
	"maxLength":      EsquemaLongitud,
}

// ValidarEsquema valida el documento contra los esquemas XSD del SAT que
// trae la librería (traducidos a Go en esquema_reglas.go y
// esquema_catalogos.go): CFDI 4.0 y 3.3, TimbreFiscalDigital 1.1, y los
// complementos Pagos 2.0 y 1.0, Nómina 1.2, ImpuestosLocales, CartaPorte 3.1,
// 3.0 y 2.0, ComercioExterior 2.0 y 1.1, HidroYPetro, CFDIRegistroFiscal,
// Divisas, EstadoDeCuentaCombustible 1.2, PorCuentadeTerceros 1.1,
// VentaVehiculos 1.1, Aerolineas, LeyendasFiscales, instEducativas y
// Donatarias 1.1, con sus catálogos.
//
// Valida como lo haría un validador XSD (libxml2, p. ej.) con esos esquemas
// cargados: estructura (orden, minOccurs, maxOccurs, elementos y atributos
// no declarados, texto), atributos requeridos y fijos, tipos y facetas
// (pattern, catálogos, rangos, dígitos, longitudes), después de aplicar el
// whiteSpace del tipo. Los xs:any del SAT (Complemento,
// ComplementoConcepto, Addenda) son estrictos: un elemento ahí sin esquema
// en la librería no se valida y queda como aviso EsquemaNoSoportado (el
// resultado es ValidacionIncompleta si no hay errores), igual que la raíz si
// no es un Comprobante 3.3 o 4.0.
//
// Trabaja sobre los bytes, como CadenaOriginal: el esquema habla del
// documento tal cual y no de lo que el modelo normaliza. Por lo mismo, lee
// el XML como un procesador conforme: con el encoding que declara y sin
// corregir nada. Un documento que Parse lee con advertencias de encoding,
// caracteres prohibidos o XML no estricto aquí es EsquemaXMLIlegible.
//
// Nunca falla: todo lo que encuentra va en los hallazgos.
func ValidarEsquema(xml []byte) *ResultadoValidacion {
	v := &validadorXSD{}
	v.validar(xml)
	if v.hs == nil {
		v.hs = []Hallazgo{}
	}
	return &ResultadoValidacion{Estado: EstadoDe(v.hs), Hallazgos: v.hs}
}

// Tablas que genera internal/genxsd.

type patronXSD struct {
	xsd string // el patrón como viene en el XSD, para los mensajes
	re  *regexp.Regexp
}

type facetaXSD struct {
	tipo     string // nombre XSD de la faceta
	n        int    // fractionDigits, totalDigits, length, minLength, maxLength
	valor    string // minInclusive, maxInclusive
	patrones []patronXSD
	catalogo *catalogoXSD
}

type tipoXSD struct {
	nombre  string // para los mensajes
	prim    string // tipo integrado del que deriva: string, decimal, int…
	ws      string // whiteSpace efectivo: preserve, replace o collapse
	facetas []facetaXSD
}

// catalogoXSD es una enumeration: valores, uno por línea y ordenados, que se
// parten la primera vez que se consultan.
type catalogoXSD struct {
	nombre  string
	valores string
	once    sync.Once
	lista   []string
}

func (c *catalogoXSD) contiene(v string) bool {
	c.once.Do(func() { c.lista = strings.Split(strings.Trim(c.valores, "\n"), "\n") })
	i := sort.SearchStrings(c.lista, v)
	return i < len(c.lista) && c.lista[i] == v
}

type atributoXSD struct {
	nombre    string
	tipo      *tipoXSD
	requerido bool
	fijo      string
	conFijo   bool
}

type particulaXSD struct {
	elem     *elementoXSD   // nil en un comodín o un xs:choice
	comodin  bool           // xs:any estricto de cualquier namespace
	opciones []particulaXSD // xs:choice (maxOccurs 1) de elementos
	min, max int            // max -1 = unbounded
}

// opcion devuelve la opción del xs:choice p que acepta n, o -1.
func (p *particulaXSD) opcion(n nombreXML) int {
	for i, o := range p.opciones {
		if o.elem.ns == n.ns && o.elem.nombre == n.local {
			return i
		}
	}
	return -1
}

// minimo es lo mínimo que exige la partícula: en un xs:choice, cero si
// alguna opción puede ir vacía.
func (p *particulaXSD) minimo() int {
	if p.opciones == nil || p.min == 0 {
		return p.min
	}
	m := p.opciones[0].min
	for _, o := range p.opciones[1:] {
		m = min(m, o.min)
	}
	return m
}

type elementoXSD struct {
	ns, nombre     string
	vacio          bool     // sin secuencia: ni texto ni hijos
	texto          *tipoXSD // contenido simple: sin hijos, el texto es de este tipo
	mixto          bool     // mixed: admite texto entre los hijos
	todos          bool     // xs:all: las partículas en cualquier orden, cada una a lo más una vez
	seqMin, seqMax int      // ocurrencias de la secuencia (max -1 = unbounded)
	attrs          []atributoXSD
	particulas     []particulaXSD
}

// varianteXSD es otra declaración de un elemento global (TFD 1.0 frente a
// 1.1), que se elige cuando el elemento trae atributo = valor.
type varianteXSD struct {
	atributo, valor string
	elem            *elementoXSD
}

// declGlobal es la declaración global de un elemento con nombre n y
// atributos attrs: la de su variante si trae el atributo que la elige
// (version="1.0" en el TFD), si no la de globalesXSD. Un solo esquema por
// namespace, como libxml2, validaría un TFD 1.0 contra el XSD de 1.1.
func declGlobal(n nombreXML, attrs []xml.Attr) *elementoXSD {
	for _, v := range variantesXSD[n] {
		for _, a := range attrs {
			if a.Name.Space == "" && a.Name.Local == v.atributo && a.Value == v.valor {
				return v.elem
			}
		}
	}
	return globalesXSD[n]
}

func (e *elementoXSD) atributo(nombre string) *atributoXSD {
	for i := range e.attrs {
		if e.attrs[i].nombre == nombre {
			return &e.attrs[i]
		}
	}
	return nil
}

// particula devuelve el índice de la partícula de la secuencia que acepta
// un hijo con ese nombre (por nombre, sin importar la posición), o -1.
func (e *elementoXSD) particula(n nombreXML) int {
	for i := range e.particulas {
		p := &e.particulas[i]
		if p.comodin || (p.elem != nil && p.elem.ns == n.ns && p.elem.nombre == n.local) || (p.opciones != nil && p.opcion(n) >= 0) {
			return i
		}
	}
	return -1
}

// repetible dice si el esquema deja repetir el hijo n, de la partícula i:
// entonces la ruta lleva siempre su índice.
func (e *elementoXSD) repetible(i int, n nombreXML) bool {
	p := &e.particulas[i]
	max := p.max
	if p.opciones != nil {
		max = p.opciones[p.opcion(n)].max
	}
	return !p.comodin && (max != 1 || e.seqMax != 1)
}

// nodoV es un elemento del documento para validar.
type nodoV struct {
	nombre     nombreXML
	attrs      []xml.Attr
	hijos      []*nodoV
	decl       *elementoXSD // nil si no tiene declaración (su contenido no se leyó)
	texto      bool         // trae texto (también solo espacios)
	textoNoEsp bool         // trae texto que no es espacio
	valor      []byte       // el texto, solo si la declaración es de contenido simple
}

type validadorXSD struct {
	hs []Hallazgo
}

func (v *validadorXSD) hallazgo(codigo, nivel, ruta, attr, valor, regla, msg string) {
	if attr != "" {
		ruta += "/@" + attr
	}
	v.hs = append(v.hs, Hallazgo{Codigo: codigo, Nivel: nivel, Ruta: ruta, Atributo: attr, Valor: valor, Regla: regla, Mensaje: msg})
}

func (v *validadorXSD) error(codigo, ruta, regla, msg string) {
	v.hallazgo(codigo, NivelError, ruta, "", "", regla, msg)
}

func (v *validadorXSD) validar(b []byte) {
	b, err := decodificarXML(b)
	if err == nil {
		var raiz *nodoV
		if raiz, err = arbolXSD(b); err == nil {
			v.raiz(raiz)
			return
		}
	}
	v.error(EsquemaXMLIlegible, "", "xml", err.Error())
}

func (v *validadorXSD) raiz(n *nodoV) {
	ruta := "/" + n.nombre.local
	if n.nombre.local == "Comprobante" && n.nombre.ns == NSCFDI33 {
		// cfd/3 es el namespace de 3.3 y también de 3.2 y 3.0, que se
		// distinguen por @version en minúsculas.
		if _, mayus := attrSinNS(n.attrs, "Version"); !mayus {
			if ver, ok := attrSinNS(n.attrs, "version"); ok {
				v.hallazgo(EsquemaNoSoportado, NivelAviso, ruta, "", "", "raiz",
					fmt.Sprintf("CFDI %s: la librería no trae su esquema", ver))
				return
			}
		}
	}
	if n.decl == nil {
		v.hallazgo(EsquemaNoSoportado, NivelAviso, ruta, "", "", "raiz",
			"la raíz "+n.nombre.String()+" no tiene esquema en la librería")
		return
	}
	v.elemento(n, ruta)
}

func attrSinNS(attrs []xml.Attr, nombre string) (string, bool) {
	for _, a := range attrs {
		if a.Name.Space == "" && a.Name.Local == nombre {
			return a.Value, true
		}
	}
	return "", false
}

// elemento valida n, que tiene declaración, y su contenido.
func (v *validadorXSD) elemento(n *nodoV, ruta string) {
	v.atributos(n, ruta)
	d := n.decl
	if d.texto != nil {
		// Contenido simple (Detallista): el texto, ya con el whiteSpace del
		// tipo, contra el tipo; ningún hijo.
		cuenta := map[string]int{}
		for _, h := range n.hijos {
			cuenta[h.nombre.local]++
			v.error(EsquemaElementoNoPermitido, rutaHijo(ruta, h.nombre.local, cuenta[h.nombre.local], false), "texto",
				"el elemento "+d.nombre+" es de contenido simple y no admite elementos hijos")
		}
		valor := string(n.valor)
		for _, f := range validarValor(d.texto, valor) {
			v.hallazgo(f.codigo, NivelError, ruta, "", valor, f.regla, f.msg)
		}
		return
	}
	if d.vacio {
		if n.texto {
			v.error(EsquemaContenido, ruta, "vacio", "el elemento no admite contenido y trae texto")
		}
		cuenta := map[string]int{}
		for _, h := range n.hijos {
			cuenta[h.nombre.local]++
			v.error(EsquemaElementoNoPermitido, rutaHijo(ruta, h.nombre.local, cuenta[h.nombre.local], false), "vacio",
				"el elemento "+d.nombre+" no admite elementos hijos")
		}
		return
	}
	if n.textoNoEsp && !d.mixto {
		v.error(EsquemaContenido, ruta, "solo-elementos", "el elemento solo admite elementos hijos y trae texto")
	}
	rutas := v.secuencia(n, ruta)
	for i, h := range n.hijos {
		switch {
		case h.decl != nil:
			v.elemento(h, rutas[i])
		case d.particula(h.nombre) >= 0:
			// Por un xs:any sin esquema para ese elemento.
			v.hallazgo(EsquemaNoSoportado, NivelAviso, rutas[i], "", "", "any",
				"el elemento "+h.nombre.String()+" no tiene esquema en la librería")
		}
	}
}

// rutaHijo arma la ruta de un hijo con el contrato de Advertencia: índice
// siempre en los nodos repetibles; en los demás, sin índice el primero y
// con [n] desde el segundo.
func rutaHijo(padre, local string, n int, repetible bool) string {
	if repetible || n > 1 {
		return padre + "/" + local + "[" + strconv.Itoa(n) + "]"
	}
	return padre + "/" + local
}

// secuencia empareja los hijos de n con la secuencia de su declaración y
// devuelve la ruta de cada hijo. Los esquemas del SAT cumplen la regla de
// atribución única (y el generador exige nombres distintos en cada
// secuencia), así que el emparejamiento ávido es el del esquema.
func (v *validadorXSD) secuencia(n *nodoV, ruta string) []string {
	d := n.decl
	if d.todos {
		return v.todos(n, ruta)
	}
	ps := d.particulas
	rutas := make([]string, len(n.hijos))
	idx := make([]int, len(n.hijos))
	op := make([]int, len(n.hijos)) // opción del xs:choice de cada hijo
	cuenta := map[string]int{}
	for i, h := range n.hijos {
		idx[i] = d.particula(h.nombre)
		if idx[i] >= 0 && ps[idx[i]].opciones != nil {
			op[i] = ps[idx[i]].opcion(h.nombre)
		}
		cuenta[h.nombre.local]++
		rutas[i] = rutaHijo(ruta, h.nombre.local, cuenta[h.nombre.local], idx[i] >= 0 && d.repetible(idx[i], h.nombre))
	}
	cnt := make([]int, len(ps))     // ocurrencias de cada partícula en la vuelta actual
	elegida := make([]int, len(ps)) // opción elegida de cada xs:choice en la vuelta actual
	for j := range elegida {
		elegida[j] = -1
	}
	// minDe es el mínimo de la partícula j: en un xs:choice, el de la opción
	// elegida o, si no hay, el de la más laxa.
	minDe := func(j int) int {
		if ps[j].opciones != nil && elegida[j] >= 0 {
			return ps[j].opciones[elegida[j]].min
		}
		return ps[j].minimo()
	}
	faltante := func(j int) {
		p := ps[j]
		switch {
		case p.comodin:
			v.error(EsquemaElementoFaltante, ruta, "minOccurs", fmt.Sprintf("%s requiere al menos %d elemento hijo", d.nombre, p.min))
			return
		case p.opciones != nil && elegida[j] < 0:
			var noms []string
			for _, o := range p.opciones {
				noms = append(noms, o.elem.nombre)
			}
			v.error(EsquemaElementoFaltante, ruta, "minOccurs", "falta uno de los elementos "+strings.Join(noms, ", "))
			return
		case p.opciones != nil:
			p = p.opciones[elegida[j]]
		}
		v.error(EsquemaElementoFaltante, rutaHijo(ruta, p.elem.nombre, cnt[j]+1, d.repetible(j, nombreXML{p.elem.ns, p.elem.nombre})), "minOccurs",
			"falta el elemento "+p.elem.nombre)
	}
	// aparece dice si un elemento de la partícula j viene entre los hijos
	// desde la posición k.
	aparece := func(j, k int) bool {
		for _, x := range idx[k:] {
			if x == j {
				return true
			}
		}
		return false
	}
	saltados := map[int]bool{} // partículas que vinieron fuera de orden y ya se reportaron
	pi, iter := 0, 1
	// cumplidas dice si las partículas de desde a hasta-1 ya tienen su mínimo.
	cumplidas := func(desde, hasta int) bool {
		for j := desde; j < hasta; j++ {
			if cnt[j] < minDe(j) && !saltados[j] {
				return false
			}
		}
		return true
	}
	// contar suma el hijo k a su partícula j.
	contar := func(j, k int) {
		cnt[j]++
		if ps[j].opciones != nil {
			elegida[j] = op[k]
		}
	}
	for k, h := range n.hijos {
		j := idx[k]
		if j < 0 {
			v.error(EsquemaElementoNoPermitido, rutas[k], "elemento",
				"el elemento "+h.nombre.String()+" no está permitido en "+d.nombre)
			continue
		}
		maxOK := ps[j].max == -1 || cnt[j] < ps[j].max
		otra := false // en un xs:choice, viene otra opción que la elegida
		if ps[j].opciones != nil {
			m := ps[j].opciones[op[k]].max
			otra = elegida[j] >= 0 && elegida[j] != op[k]
			maxOK = !otra && (m == -1 || cnt[j] < m)
		}
		switch {
		case j == pi && maxOK:
			contar(j, k)
		case j > pi:
			if cumplidas(pi, j) {
				pi = j
				contar(j, k)
				continue
			}
			// Faltan partículas requeridas entre la actual y esta. Si alguna
			// viene después, esta está fuera de orden; si no, faltan.
			fuera := false
			for u := pi; u < j; u++ {
				if cnt[u] < minDe(u) && !saltados[u] && aparece(u, k+1) {
					fuera = true
				}
			}
			if fuera {
				saltados[j] = true
				v.error(EsquemaOrden, rutas[k], "orden", "el elemento "+h.nombre.local+" está fuera del orden de la secuencia")
				continue
			}
			for u := pi; u < j; u++ {
				if cnt[u] < minDe(u) && !saltados[u] {
					faltante(u)
				}
			}
			pi = j
			contar(j, k)
		default: // j < pi, o j == pi sin lugar
			// ¿Empieza otra vuelta de la secuencia?
			if (d.seqMax == -1 || iter < d.seqMax) && cumplidas(pi, len(ps)) && cumplidas(0, j) && reiniciable(ps, j) {
				iter++
				for u := range cnt {
					cnt[u], elegida[u] = 0, -1
				}
				pi = j
				contar(j, k)
				continue
			}
			if otra {
				v.error(EsquemaElementoNoPermitido, rutas[k], "choice", fmt.Sprintf("el elemento %s no puede ir junto con %s (xs:choice)",
					h.nombre.local, ps[j].opciones[elegida[j]].elem.nombre))
			} else if !maxOK {
				v.error(EsquemaElementoExcedente, rutas[k], "maxOccurs",
					fmt.Sprintf("el elemento %s se repite más de %d veces", h.nombre.local, ps[j].max))
			} else {
				saltados[j] = true
				v.error(EsquemaOrden, rutas[k], "orden", "el elemento "+h.nombre.local+" está fuera del orden de la secuencia")
			}
		}
	}
	for u := pi; u < len(ps); u++ {
		if cnt[u] < minDe(u) && !saltados[u] {
			faltante(u)
		}
	}
	return rutas
}

// todos valida los hijos de un xs:all: sin orden, cada partícula entre su
// mínimo y una vez (el generador exige maxOccurs 1).
func (v *validadorXSD) todos(n *nodoV, ruta string) []string {
	d := n.decl
	rutas := make([]string, len(n.hijos))
	cnt := make([]int, len(d.particulas))
	cuenta := map[string]int{}
	for k, h := range n.hijos {
		cuenta[h.nombre.local]++
		rutas[k] = rutaHijo(ruta, h.nombre.local, cuenta[h.nombre.local], false)
		j := d.particula(h.nombre)
		if j < 0 {
			v.error(EsquemaElementoNoPermitido, rutas[k], "elemento", "el elemento "+h.nombre.String()+" no está permitido en "+d.nombre)
			continue
		}
		if cnt[j]++; cnt[j] > d.particulas[j].max {
			v.error(EsquemaElementoExcedente, rutas[k], "maxOccurs",
				fmt.Sprintf("el elemento %s se repite más de %d veces", h.nombre.local, d.particulas[j].max))
		}
	}
	for j, p := range d.particulas {
		if cnt[j] < p.min {
			v.error(EsquemaElementoFaltante, rutaHijo(ruta, p.elem.nombre, 1, false), "minOccurs", "falta el elemento "+p.elem.nombre)
		}
	}
	return rutas
}

// reiniciable dice si una vuelta nueva de la secuencia puede empezar en la
// partícula j: todas las anteriores son opcionales.
func reiniciable(ps []particulaXSD, j int) bool {
	for i := range ps[:j] {
		if ps[i].minimo() > 0 {
			return false
		}
	}
	return true
}

func (v *validadorXSD) atributos(n *nodoV, ruta string) {
	d := n.decl
	presentes := map[string]bool{}
	for _, a := range n.attrs {
		switch {
		case a.Name.Space == nsXMLNSAttrs || (a.Name.Space == "" && a.Name.Local == "xmlns"):
			continue
		case a.Name.Space == nsXSI && (a.Name.Local == "schemaLocation" || a.Name.Local == "noNamespaceSchemaLocation"):
			continue
		case a.Name.Space != "":
			v.hallazgo(EsquemaAtributoNoPermitido, NivelError, ruta, "{"+a.Name.Space+"}"+a.Name.Local, a.Value, "atributo",
				"el atributo {"+a.Name.Space+"}"+a.Name.Local+" no está permitido")
			continue
		}
		decl := d.atributo(a.Name.Local)
		if decl == nil {
			v.hallazgo(EsquemaAtributoNoPermitido, NivelError, ruta, a.Name.Local, a.Value, "atributo",
				"el atributo "+a.Name.Local+" no está declarado en "+d.nombre)
			continue
		}
		presentes[a.Name.Local] = true
		for _, f := range validarValor(decl.tipo, a.Value) {
			v.hallazgo(f.codigo, NivelError, ruta, a.Name.Local, a.Value, f.regla, f.msg)
		}
		if decl.conFijo && normalizarWS(a.Value, decl.tipo.ws) != decl.fijo {
			v.hallazgo(EsquemaValorFijo, NivelError, ruta, a.Name.Local, a.Value, "fixed",
				fmt.Sprintf("el valor debe ser %q", decl.fijo))
		}
	}
	for _, a := range d.attrs {
		if a.requerido && !presentes[a.nombre] {
			v.hallazgo(EsquemaAtributoFaltante, NivelError, ruta, a.nombre, "", "required", "falta el atributo requerido "+a.nombre)
		}
	}
}

type fallaValor struct{ codigo, regla, msg string }

// validarValor revisa un valor contra su tipo: whiteSpace, forma léxica del
// tipo integrado y facetas.
func validarValor(t *tipoXSD, crudo string) []fallaValor {
	var out []fallaValor
	s := normalizarWS(crudo, t.ws)
	num, lexOK := lexicoXSD(t.prim, s)
	if !lexOK {
		out = append(out, fallaValor{EsquemaTipo, "tipo", fmt.Sprintf("%q no es un valor válido de xs:%s", s, t.prim)})
	}
	for _, f := range t.facetas {
		ok := true
		var msg string
		switch f.tipo {
		case "pattern":
			ok = false
			var pats []string
			for _, p := range f.patrones {
				ok = ok || p.re.MatchString(s)
				pats = append(pats, p.xsd)
			}
			msg = fmt.Sprintf("%q no cumple el patrón %s", s, strings.Join(pats, " | "))
		case "enumeration":
			// En los numéricos compara el valor (el catálogo trae la forma
			// canónica); en los demás, el texto.
			clave := s
			if numericoXSD[t.prim] && lexOK {
				clave = num.String()
			}
			ok = f.catalogo.contiene(clave)
			msg = fmt.Sprintf("%q no está en el catálogo %s", s, f.catalogo.nombre)
		case "length", "minLength", "maxLength":
			l := utf8.RuneCountInString(s)
			ok = (f.tipo == "length" && l == f.n) || (f.tipo == "minLength" && l >= f.n) || (f.tipo == "maxLength" && l <= f.n)
			msg = fmt.Sprintf("%q tiene %d caracteres (%s %d)", s, l, f.tipo, f.n)
		case "minInclusive", "maxInclusive":
			if !lexOK {
				continue
			}
			var cmp int
			switch t.prim {
			case "dateTime":
				cmp = compararFechaHoraXSD(s, f.valor)
			case "date":
				cmp = compararFechaHoraXSD(fechaComoFechaHora(s), fechaComoFechaHora(f.valor))
			default:
				cmp = num.Cmp(decimal.MustParse(f.valor))
			}
			ok = (f.tipo == "minInclusive" && (cmp == 0 || cmp == 1)) || (f.tipo == "maxInclusive" && (cmp == 0 || cmp == -1))
			msg = fmt.Sprintf("%s está fuera del rango (%s %s)", s, f.tipo, f.valor)
		case "fractionDigits", "totalDigits":
			if !lexOK {
				continue
			}
			ent, frac := digitos(num)
			ok = (f.tipo == "fractionDigits" && frac <= f.n) || (f.tipo == "totalDigits" && ent+frac <= f.n)
			msg = fmt.Sprintf("%s tiene más dígitos de los permitidos (%s %d)", s, f.tipo, f.n)
		}
		if !ok {
			out = append(out, fallaValor{codigoFaceta[f.tipo], f.tipo, msg})
		}
	}
	return out
}

// digitos cuenta los dígitos enteros (sin ceros a la izquierda) y
// decimales (sin ceros a la derecha) del valor: las facetas totalDigits y
// fractionDigits hablan del valor, no del texto ("2.500" tiene 1 decimal).
func digitos(d decimal.Decimal) (enteros, decimales int) {
	s := strings.TrimPrefix(d.String(), "-")
	ent, frac, _ := strings.Cut(s, ".")
	ent = strings.TrimLeft(ent, "0")
	return len(ent), len(frac)
}

func normalizarWS(s, ws string) string {
	switch ws {
	case "replace":
		return strings.Map(func(r rune) rune {
			if r == '\t' || r == '\n' || r == '\r' {
				return ' '
			}
			return r
		}, s)
	case "collapse":
		return strings.Join(strings.FieldsFunc(s, func(r rune) bool {
			return r == ' ' || r == '\t' || r == '\n' || r == '\r'
		}), " ")
	}
	return s
}

var (
	reEnteroXSD   = regexp.MustCompile(`^[+-]?[0-9]+$`)
	reFechaXSD    = regexp.MustCompile(`^(-?)([0-9]{4,})-([0-9]{2})-([0-9]{2})(Z|[+-][0-9]{2}:[0-9]{2})?$`)
	reFechaHraXSD = regexp.MustCompile(`^(-?)([0-9]{4,})-([0-9]{2})-([0-9]{2})T([0-9]{2}):([0-9]{2}):([0-9]{2})(\.[0-9]+)?(Z|[+-][0-9]{2}:[0-9]{2})?$`)
	reHoraXSD     = regexp.MustCompile(`^([0-9]{2}):([0-9]{2}):([0-9]{2})(\.[0-9]+)?(Z|[+-][0-9]{2}:[0-9]{2})?$`)
	limitesEntero = map[string][2]int64{"int": {-1 << 31, 1<<31 - 1}, "short": {-1 << 15, 1<<15 - 1}}
	numericoXSD   = map[string]bool{"decimal": true, "integer": true, "int": true, "short": true}
)

// esNmtoken dice si s es un xs:NMTOKEN: uno o más NameChar de XML 1.0. En
// ASCII son letras, dígitos, '.', '-', '_' y ':'; fuera de ASCII se aproxima
// con letras, dígitos y marcas combinantes de Unicode, más el punto medio
// (U+00B7). Los valores de los XSD del SAT son ASCII.
func esNmtoken(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		switch {
		case r < 0x80:
			if !(r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' || strings.ContainsRune(".-_:", r)) {
				return false
			}
		case r == 0xB7 || unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.In(r, unicode.Mn, unicode.Mc):
		default:
			return false
		}
	}
	return true
}

// lexicoXSD revisa la forma léxica de s (ya normalizado) en el tipo
// integrado prim. Para los numéricos devuelve también el valor.
func lexicoXSD(prim, s string) (decimal.Decimal, bool) {
	switch prim {
	case "NMTOKEN":
		return decimal.Decimal{}, esNmtoken(s)
	case "decimal":
		d, err := decimal.Parse(s)
		return d, err == nil && !strings.ContainsAny(s, " \t\r\n")
	case "integer", "int", "short":
		if !reEnteroXSD.MatchString(s) {
			return decimal.Decimal{}, false
		}
		d := decimal.MustParse(s)
		if lim, ok := limitesEntero[prim]; ok {
			n, _ := new(big.Int).SetString(strings.TrimPrefix(s, "+"), 10)
			if n.Cmp(big.NewInt(lim[0])) < 0 || n.Cmp(big.NewInt(lim[1])) > 0 {
				return d, false
			}
		}
		return d, true
	case "date":
		m := reFechaXSD.FindStringSubmatch(s)
		return decimal.Decimal{}, m != nil && fechaValida(m[2], m[3], m[4]) && zonaValida(m[5])
	case "dateTime":
		m := reFechaHraXSD.FindStringSubmatch(s)
		return decimal.Decimal{}, m != nil && fechaValida(m[2], m[3], m[4]) && horaValida(m[5], m[6], m[7], m[8]) && zonaValida(m[9])
	case "time":
		m := reHoraXSD.FindStringSubmatch(s)
		return decimal.Decimal{}, m != nil && horaValida(m[1], m[2], m[3], m[4]) && zonaValida(m[5])
	case "base64Binary":
		sin := strings.ReplaceAll(s, " ", "")
		if len(sin)%4 != 0 {
			return decimal.Decimal{}, false
		}
		_, err := base64.StdEncoding.DecodeString(sin)
		return decimal.Decimal{}, err == nil
	}
	return decimal.Decimal{}, true // string, anySimpleType
}

// compararFechaHoraXSD compara dos xs:dateTime ya válidos como el orden
// parcial de XML Schema que usa libxml2: -1, 0 o 1, y 2 si no se puede
// decidir. Si los dos traen zona (o ninguno) se comparan directo, como
// instantes (u horas). Si solo uno la trae, al otro le cabe cualquiera entre
// -14:00 y +14:00, así que solo se decide si queda del mismo lado con todas.
func compararFechaHoraXSD(s, lim string) int {
	a, conZona := instanteXSD(s)
	b, limConZona := instanteXSD(lim)
	if conZona != limConZona {
		const margen = 14 * time.Hour
		switch {
		case a.Add(margen).Before(b):
			return -1
		case a.Add(-margen).After(b):
			return 1
		}
		return 2
	}
	return a.Compare(b)
}

// fechaComoFechaHora convierte un xs:date válido en el xs:dateTime de las
// 00:00:00 de ese día, con la misma zona (o sin ella), para compararlo
// como tal.
func fechaComoFechaHora(s string) string {
	m := reFechaXSD.FindStringSubmatch(s)
	return m[1] + m[2] + "-" + m[3] + "-" + m[4] + "T00:00:00" + m[5]
}

// instanteXSD lee un xs:dateTime válido como instante en UTC; sin zona, la
// hora tal cual.
func instanteXSD(s string) (t time.Time, conZona bool) {
	m := reFechaHraXSD.FindStringSubmatch(s)
	num := func(x string) int { n, _ := strconv.Atoi(x); return n }
	anio := num(m[2])
	if m[1] == "-" {
		anio = -anio
	}
	var nanos int
	if m[8] != "" {
		frac := (m[8][1:] + "000000000")[:9]
		nanos = num(frac)
	}
	t = time.Date(anio, time.Month(num(m[3])), num(m[4]), num(m[5]), num(m[6]), num(m[7]), nanos, time.UTC)
	switch z := m[9]; z {
	case "":
		return t, false
	case "Z":
		return t, true
	default:
		off := time.Duration(num(z[1:3]))*time.Hour + time.Duration(num(z[4:6]))*time.Minute
		if z[0] == '+' {
			off = -off
		}
		return t.Add(off), true
	}
}

func fechaValida(anio, mes, dia string) bool {
	if len(anio) > 4 && anio[0] == '0' {
		return false
	}
	a, err := strconv.Atoi(anio)
	if err != nil || a == 0 {
		return false
	}
	m, _ := strconv.Atoi(mes)
	d, _ := strconv.Atoi(dia)
	if m < 1 || m > 12 || d < 1 {
		return false
	}
	dias := [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}[m]
	if m == 2 && (a%4 == 0 && (a%100 != 0 || a%400 == 0)) {
		dias = 29
	}
	return d <= dias
}

func horaValida(hh, mm, ss, frac string) bool {
	h, _ := strconv.Atoi(hh)
	m, _ := strconv.Atoi(mm)
	s, _ := strconv.Atoi(ss)
	if h == 24 {
		return m == 0 && s == 0 && strings.Trim(frac, ".0") == ""
	}
	return h < 24 && m < 60 && s < 60
}

func zonaValida(z string) bool {
	if z == "" || z == "Z" {
		return true
	}
	h, _ := strconv.Atoi(z[1:3])
	m, _ := strconv.Atoi(z[4:6])
	return m < 60 && (h < 14 || (h == 14 && m == 0))
}

// decodificarXML convierte los bytes a UTF-8 como lo haría un procesador
// XML conforme: con el encoding que declaran y sin corregir. Lo que
// normalizarCodificacion arregla con una advertencia (bytes que no son del
// encoding declarado, encoding desconocido, caracteres prohibidos) aquí es
// un error.
func decodificarXML(b []byte) ([]byte, error) {
	switch {
	case bytes.HasPrefix(b, []byte{0xEF, 0xBB, 0xBF}):
		b = b[3:]
	case bytes.HasPrefix(b, []byte{0xFF, 0xFE}):
		b = desdeUTF16(b[2:], false)
	case bytes.HasPrefix(b, []byte{0xFE, 0xFF}):
		b = desdeUTF16(b[2:], true)
	}
	declarado := encodingDeclarado(b)
	switch strings.ToLower(declarado) {
	case "", "utf-8", "utf8", "utf-16", "utf16":
		if !utf8.Valid(b) {
			return nil, errors.New("los bytes no son UTF-8 válido")
		}
	case "iso-8859-1", "iso8859-1", "latin1", "latin-1", "l1":
		out := make([]byte, 0, len(b)+len(b)/8)
		for _, c := range b {
			out = utf8.AppendRune(out, rune(c))
		}
		b = out
	case "windows-1252", "cp1252":
		for _, c := range b {
			if c == 0x81 || c == 0x8D || c == 0x8F || c == 0x90 || c == 0x9D {
				return nil, fmt.Errorf("el byte 0x%02X no existe en windows-1252", c)
			}
		}
		b = desdeWindows1252(b)
	case "us-ascii", "ascii":
		if !esASCII(b) {
			return nil, errors.New("bytes que no son ASCII en un documento US-ASCII")
		}
	default:
		return nil, fmt.Errorf("encoding %q no soportado", declarado)
	}
	if n := contarInvalidos(b); n > 0 {
		return nil, fmt.Errorf("%d caracteres no permitidos en XML 1.0", n)
	}
	return b, nil
}

// arbolXSD lee el documento en modo estricto y le pone a cada elemento su
// declaración: por nombre en la secuencia del padre o, en un xs:any, entre
// los elementos globales. El contenido de lo que no tiene declaración se
// salta sin guardarlo.
func arbolXSD(b []byte) (*nodoV, error) {
	d := xml.NewDecoder(bytes.NewReader(b))
	d.CharsetReader = charsetIdentidad
	var raiz *nodoV
	var pila []*nodoV
	var ns alcanceNS
	for {
		tok, err := d.Token()
		if err == io.EOF && raiz != nil && len(pila) == 0 {
			return raiz, nil
		}
		if err != nil {
			return nil, err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			if raiz != nil && len(pila) == 0 {
				return nil, errors.New("más de un elemento raíz")
			}
			if err := ns.entrar(t); err != nil {
				return nil, err
			}
			n := &nodoV{nombre: nombreXML{t.Name.Space, t.Name.Local}, attrs: t.Attr}
			if len(pila) == 0 {
				raiz = n
				n.decl = declGlobal(n.nombre, t.Attr)
			} else {
				p := pila[len(pila)-1]
				p.hijos = append(p.hijos, n)
				n.decl = declHijo(p.decl, n.nombre, t.Attr)
			}
			if n.decl == nil && len(pila) > 0 {
				// Lo de adentro no se valida, pero tiene que ser XML
				// conforme.
				if err := ns.saltar(d); err != nil {
					return nil, err
				}
				continue
			}
			pila = append(pila, n)
		case xml.EndElement:
			pila = pila[:len(pila)-1]
			ns.salir()
		case xml.CharData:
			if len(pila) == 0 {
				if len(bytes.Trim(t, " \t\r\n")) > 0 {
					return nil, errors.New("texto fuera del elemento raíz")
				}
				continue
			}
			p := pila[len(pila)-1]
			p.texto = p.texto || len(t) > 0
			p.textoNoEsp = p.textoNoEsp || len(bytes.Trim(t, " \t\r\n")) > 0
			if p.decl != nil && p.decl.texto != nil {
				// Los comentarios parten el texto; su valor es todo junto.
				p.valor = append(p.valor, t...)
			}
		}
	}
}

// declHijo es la declaración de un hijo de un elemento con declaración p:
// la de su partícula por nombre o, en un xs:any, la global (ver declGlobal).
// nil si no hay.
func declHijo(p *elementoXSD, n nombreXML, attrs []xml.Attr) *elementoXSD {
	if p == nil || p.vacio || p.texto != nil {
		return nil
	}
	j := p.particula(n)
	switch {
	case j < 0:
		return nil
	case p.particulas[j].comodin:
		return declGlobal(n, attrs)
	case p.particulas[j].opciones != nil:
		pa := &p.particulas[j]
		return pa.opciones[pa.opcion(n)].elem
	}
	return p.particulas[j].elem
}
