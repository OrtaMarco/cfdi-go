// Package genxsd traduce los esquemas XSD del SAT (CFDI 4.0 y 3.3, TFD 1.1
// y 1.0 y los complementos de Raices y Variantes, con los catálogos y tipos de datos que
// importan) a tablas Go: los archivos esquema_reglas.go y
// esquema_catalogos.go del paquete cfdi, que usa ValidarEsquema.
//
// Los XSD del SAT usan un subconjunto chico de XML Schema: elementos con
// complexType anónimo o con nombre (por type, o por una xs:extension sin
// contenido propio), una xs:sequence sin anidar que puede llevar un
// xs:choice de elementos que no se repite, un xs:all de elementos que van
// a lo más una vez, xs:any estricto, atributos con tipo simple, y tipos
// simples por restricción de tipos integrados con las facetas pattern,
// enumeration (también en xs:int, por valor), minInclusive, maxInclusive
// (también en xs:date y xs:dateTime), fractionDigits, totalDigits, length,
// minLength, maxLength y whiteSpace. Este paquete entiende exactamente eso y
// falla con cualquier otra construcción (xs:group, choice repetible o
// anidado, extensiones con contenido, listas, uniones, minExclusive,
// atributos desconocidos…) en vez de ignorarla, para que un XSD nuevo del
// SAT no se traduzca mal en silencio.
//
// Los patrones XSD se traducen a RE2 (ver traducirPatron); lo que no se
// sabe traducir también detiene la generación.
//
// Los XSD están copiados sin cambios en xsd/sitio_internet (de
// phpcfdi/resources-sat-xml, que los baja del SAT). Para regenerar:
//
//	go run ./internal/genxsd/generar
package genxsd

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"go/format"
	"os"
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"unicode"

	"github.com/OrtaMarco/cfdi-go/decimal"
)

const nsXS = "http://www.w3.org/2001/XMLSchema"

// reFechaHoraLimite es la forma que se acepta en el valor de minInclusive
// de un xs:dateTime: con zona (FechaExp de Retenciones 1.0) o sin ella (ECC
// 1.1); ValidarEsquema compara con el orden parcial de XML Schema.
var reFechaHoraLimite = regexp.MustCompile(`^[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}(Z|[+-][0-9]{2}:[0-9]{2})?$`)

// reFechaLimite es lo mismo para un xs:date (FechaOperacion de SPEI).
var reFechaLimite = regexp.MustCompile(`^[0-9]{4}-[0-9]{2}-[0-9]{2}(Z|[+-][0-9]{2}:[0-9]{2})?$`)

// Raices son los esquemas que se generan, con el prefijo de sus variables
// Go. Los que importan (catálogos, tipos de datos) se siguen solos.
var Raices = []struct{ Prefijo, Archivo string }{
	{"CFDI40", "cfd/4/cfdv40.xsd"},
	{"CFDI33", "cfd/3/cfdv33.xsd"},
	{"TFD11", "cfd/TimbreFiscalDigital/TimbreFiscalDigitalv11.xsd"},
	{"Pagos20", "cfd/Pagos/Pagos20.xsd"},
	{"Nomina12", "cfd/nomina/nomina12.xsd"},
	{"ImpLocal", "cfd/implocal/implocal.xsd"},
	{"CartaPorte31", "cfd/CartaPorte/CartaPorte31.xsd"},
	{"Pagos10", "cfd/Pagos/Pagos10.xsd"},
	{"CartaPorte30", "cfd/CartaPorte/CartaPorte30.xsd"},
	{"CartaPorte20", "cfd/CartaPorte/CartaPorte20.xsd"},
	{"ComExt20", "cfd/ComercioExterior20/ComercioExterior20.xsd"},
	{"ComExt11", "cfd/ComercioExterior11/ComercioExterior11.xsd"},
	{"HidroYPetro", "cfd/hidrocarburospetroliferos.xsd"},
	{"RegistroFiscal", "cfd/cfdiregistrofiscal/cfdiregistrofiscal.xsd"},
	{"Divisas", "cfd/divisas/divisas.xsd"},
	{"ECC12", "cfd/EstadoDeCuentaCombustible/ecc12.xsd"},
	{"Terceros11", "cfd/terceros/terceros11.xsd"},
	{"VentaVehiculos11", "cfd/ventavehiculos/ventavehiculos11.xsd"},
	{"Aerolineas", "cfd/aerolineas/aerolineas.xsd"},
	{"LeyendasFisc", "cfd/leyendasFiscales/leyendasFisc.xsd"},
	{"IEDU", "cfd/iedu/iedu.xsd"},
	{"Donat11", "cfd/donat/donat11.xsd"},
	{"CartaPorte10", "cfd/CartaPorte/CartaPorte.xsd"},
	{"ComExt10", "cfd/ComercioExterior/ComercioExterior10.xsd"},
	{"ECC11", "cfd/EstadoDeCuentaCombustible/ecc11.xsd"},
	{"Nomina11", "cfd/nomina/nomina11.xsd"},
	{"INE11", "cfd/ine/ine11.xsd"},
	{"ValesDespensa", "cfd/valesdedespensa/valesdedespensa.xsd"},
	{"ConsumoComb10", "cfd/consumodecombustibles/consumodecombustibles.xsd"},
	{"ConsumoComb11", "cfd/consumodecombustibles/consumodeCombustibles11.xsd"},
	{"PagoEnEspecie", "cfd/pagoenespecie/pagoenespecie.xsd"},
	{"Notarios", "cfd/notariospublicos/notariospublicos.xsd"},
	{"ParcialesConstruccion", "cfd/servicioparcialconstruccion/servicioparcialconstruccion.xsd"},
	{"VehiculoUsado", "cfd/vehiculousado/vehiculousado.xsd"},
	{"TuristaPasajero", "cfd/TuristaPasajeroExtranjero/TuristaPasajeroExtranjero.xsd"},
	{"AcreditamientoIEPS", "cfd/acreditamiento/AcreditamientoIEPS10.xsd"},
	{"SPEI", "cfd/spei/spei.xsd"},
	{"Renovacion", "cfd/renovacionysustitucionvehiculos/renovacionysustitucionvehiculos.xsd"},
	{"CertDestruccion", "cfd/certificadodestruccion/certificadodedestruccion.xsd"},
	{"ObrasArte", "cfd/arteantiguedades/obrasarteantiguedades.xsd"},
	{"PFIC", "cfd/pfic/pfic.xsd"},
	{"GastosHidro", "cfd/GastosHidrocarburos10/GastosHidrocarburos10.xsd"},
	{"IngresosHidro", "cfd/IngresosHidrocarburos10/IngresosHidrocarburos.xsd"},
	{"ECB", "cfd/ecb/ecb.xsd"},
	{"ECC10", "cfd/ecc/ecc.xsd"},

	// Retenciones e información de pagos y sus complementos, en
	// www.sat.gob.mx/esquemas, al lado de sitio_internet.
	{"Ret20", retencion + "2/retencionpagov2.xsd"},
	{"Ret10", retencion + "1/retencionpagov1.xsd"},
	{"RetPlataformas10", retencion + "1/PlataformasTecnologicas10/ServiciosPlataformasTecnologicas10.xsd"},
	{"RetIntereses", retencion + "1/intereses/intereses.xsd"},
	{"RetPagosExtranjeros", retencion + "1/pagosaextranjeros/pagosaextranjeros.xsd"},
	{"RetDividendos", retencion + "1/dividendos/dividendos.xsd"},
	{"RetEnajenacion", retencion + "1/enajenaciondeacciones/enajenaciondeacciones.xsd"},
	{"RetArrendamiento", retencion + "1/arrendamientoenfideicomiso/arrendamientoenfideicomiso.xsd"},
	{"RetDerivados", retencion + "1/operacionesconderivados/operacionesconderivados.xsd"},
	{"RetHipotecarios", retencion + "1/intereseshipotecarios/intereseshipotecarios.xsd"},
	{"RetFideicomiso", retencion + "1/fideicomisonoempresarial/fideicomisonoempresarial.xsd"},
	{"RetPlanesRetiro10", retencion + "1/planesderetiro/planesderetiro.xsd"},
	{"RetPlanesRetiro11", retencion + "1/planesderetiro11/planesderetiro11.xsd"},
	{"RetPremios", retencion + "1/premios/premios.xsd"},
	{"RetSectorFinanciero", retencion + "1/sectorfinanciero/sectorfinanciero.xsd"},
}

// Variantes son esquemas cuyo elemento global se llama igual (mismo
// namespace y nombre) que el de una raíz: TFD 1.0 y 1.1 comparten
// {http://www.sat.gob.mx/TimbreFiscalDigital}TimbreFiscalDigital. libxml2
// carga un solo esquema por namespace y no puede elegir; ValidarEsquema sí:
// usa la variante cuando el elemento trae Atributo con el valor fixed que
// le da su XSD (version="1.0"), y si no, el de la raíz. El atributo tiene
// que ser requerido y fixed en la variante, y en la raíz no existir (TFD
// 1.1 lo llama Version) o ser fixed con otro valor (Donatarias y
// VentaVehiculos 1.1: version="1.1").
var Variantes = []struct{ Prefijo, Archivo, Atributo string }{
	{"TFD10", "cfd/TimbreFiscalDigital/TimbreFiscalDigital.xsd", "version"},
	{"Donat10", "cfd/donat/donat.xsd", "version"},
	{"VentaVehiculos10", "cfd/ventavehiculos/ventavehiculos.xsd", "version"},
	{"INE10", "cfd/ine/ine10.xsd", "Version"},
}

// retencion es el directorio de los XSD de Retenciones, relativo a
// sitio_internet.
const retencion = "../esquemas/retencionpago/"

// normalizarRuta deja una ruta relativa a sitio_internet: Retenciones 2.0
// importa catCFDI como ../../../sitio_internet/cfd/catalogos/catCFDI.xsd, que
// desde ../esquemas/retencionpago/2 es ../sitio_internet/cfd/…, el mismo
// archivo que cfd/….
func normalizarRuta(p string) string {
	p = path.Clean(p)
	if r, ok := strings.CutPrefix(p, "../sitio_internet/"); ok {
		return r
	}
	return p
}

// Importados son los esquemas que pueden importar las raíces, con su
// prefijo. Un xs:import a otro archivo detiene la generación: así un
// catálogo nuevo entra a propósito y no de rebote.
var Importados = map[string]string{
	"cfd/catalogos/catCFDI.xsd":                  "CatCFDI",
	"cfd/tipoDatos/tdCFDI/tdCFDI.xsd":            "TdCFDI",
	"cfd/catalogos/Pagos/catPagos.xsd":           "CatPagos",
	"cfd/catalogos/Nomina/catNomina.xsd":         "CatNomina",
	"cfd/catalogos/CartaPorte/catCartaPorte.xsd": "CatCartaPorte",
	"cfd/catalogos/ComExt/catComExt.xsd":         "CatComExt",

	"cfd/catalogos/HidrocarburosPetro/catHidroYPetro.xsd": "CatHidroYPetro",
	"cfd/catalogos/Combustible/catCombustible.xsd":        "CatCombustible",
	"cfd/catalogos/hidrocarburos/catHidrocarburos.xsd":    "CatHidrocarburos",

	retencion + "1/catalogos/catRetenciones.xsd":                                       "CatRetenciones",
	retencion + "1/PlataformasTecnologicas10/catalogos/CatPlataformasTecnologicas.xsd": "CatPlataformas",
	retencion + "1/planesderetiro11/catalogos/CatPlanesderetiro.xsd":                   "CatPlanesRetiro",
}

// nodo es un elemento del XSD con los namespaces en alcance.
type nodo struct {
	ns, local string
	attrs     map[string]string
	hijos     []*nodo
	texto     bool // trae texto que no es espacio
	scope     map[string]string
	archivo   string
	linea     int
}

func (n *nodo) xs(local string) bool { return n.ns == nsXS && n.local == local }

func (n *nodo) donde() string {
	s := fmt.Sprintf("%s:%d: <xs:%s", n.archivo, n.linea, n.local)
	if v, ok := n.attrs["name"]; ok {
		s += " name=" + strconv.Quote(v)
	}
	return s + ">"
}

func (n *nodo) errorf(formato string, args ...any) error {
	return fmt.Errorf("%s: %s", n.donde(), fmt.Sprintf(formato, args...))
}

// soloAttrs falla si n trae atributos fuera de los permitidos.
func (n *nodo) soloAttrs(permitidos ...string) error {
	for k := range n.attrs {
		ok := false
		for _, p := range permitidos {
			ok = ok || k == p
		}
		if !ok {
			return n.errorf("atributo %q no soportado", k)
		}
	}
	return nil
}

// hijosXS devuelve los hijos de n sin las anotaciones y falla si alguno no
// es de los permitidos.
func (n *nodo) hijosXS(permitidos ...string) ([]*nodo, error) {
	var out []*nodo
	for _, h := range n.hijos {
		if h.xs("annotation") {
			if err := revisarAnotacion(h); err != nil {
				return nil, err
			}
			continue
		}
		ok := false
		for _, p := range permitidos {
			ok = ok || h.xs(p)
		}
		if !ok {
			return nil, h.errorf("construcción no soportada dentro de <xs:%s>", n.local)
		}
		out = append(out, h)
	}
	if n.texto {
		return nil, n.errorf("texto suelto no soportado")
	}
	return out, nil
}

func revisarAnotacion(a *nodo) error {
	if err := a.soloAttrs(); err != nil {
		return err
	}
	for _, h := range a.hijos {
		if !h.xs("documentation") {
			return h.errorf("construcción no soportada dentro de <xs:annotation>")
		}
	}
	return nil
}

// cache de árboles ya leídos, por archivo y contenido: los catálogos pesan
// megas y las pruebas generan muchas veces. Los árboles no se modifican.
var (
	cacheMu sync.Mutex
	cache   = map[[32]byte]*nodo{}
)

func leer(dir, rel string) (*nodo, []byte, error) {
	b, err := os.ReadFile(path.Join(dir, rel))
	if err != nil {
		return nil, nil, err
	}
	clave := sha256.Sum256(append([]byte(rel+"\x00"), b...))
	cacheMu.Lock()
	n := cache[clave]
	cacheMu.Unlock()
	if n != nil {
		return n, b, nil
	}
	n, err = leerXSD(rel, b)
	if err != nil {
		return nil, nil, err
	}
	cacheMu.Lock()
	cache[clave] = n
	cacheMu.Unlock()
	return n, b, nil
}

func leerXSD(rel string, b []byte) (*nodo, error) {
	d := xml.NewDecoder(bytes.NewReader(b))
	var pila []*nodo
	var raiz *nodo
	for {
		off := d.InputOffset()
		tok, err := d.Token()
		if err != nil {
			if raiz != nil && len(pila) == 0 {
				return raiz, nil
			}
			return nil, fmt.Errorf("%s: %w", rel, err)
		}
		switch t := tok.(type) {
		case xml.StartElement:
			n := &nodo{ns: t.Name.Space, local: t.Name.Local, attrs: map[string]string{}, scope: map[string]string{},
				archivo: rel, linea: 1 + bytes.Count(b[:off], []byte("\n"))}
			if len(pila) > 0 {
				p := pila[len(pila)-1]
				for k, v := range p.scope {
					n.scope[k] = v
				}
				// El contenido de la documentación no interesa.
				if !(p.ns == nsXS && p.local == "documentation") {
					p.hijos = append(p.hijos, n)
				}
			} else {
				raiz = n
			}
			for _, a := range t.Attr {
				switch {
				case a.Name.Space == "xmlns":
					n.scope[a.Name.Local] = a.Value
				case a.Name.Space == "" && a.Name.Local == "xmlns":
					n.scope[""] = a.Value
				case a.Name.Space == "":
					n.attrs[a.Name.Local] = a.Value
				default:
					return nil, fmt.Errorf("%s:%d: atributo con namespace %s:%s no soportado", rel, n.linea, a.Name.Space, a.Name.Local)
				}
			}
			pila = append(pila, n)
		case xml.EndElement:
			pila = pila[:len(pila)-1]
		case xml.CharData:
			if len(pila) > 0 && len(bytes.TrimSpace(t)) > 0 {
				p := pila[len(pila)-1]
				if !(p.ns == nsXS && p.local == "documentation") {
					p.texto = true
				}
			}
		}
	}
}

// qname resuelve "prefijo:local" con los namespaces en alcance de n.
func (n *nodo) qname(v string) (qn, error) {
	pre, local, ok := strings.Cut(v, ":")
	if !ok {
		pre, local = "", v
	}
	ns, ok := n.scope[pre]
	if !ok {
		return qn{}, n.errorf("prefijo %q sin declarar en %q", pre, v)
	}
	return qn{ns, local}, nil
}

type qn struct{ ns, local string }

func (q qn) String() string { return "{" + q.ns + "}" + q.local }

// esquema es un archivo XSD cargado.
type esquema struct {
	archivo, prefijo, tns string
	raiz                  *nodo
	sha                   string
}

// generador junta todos los esquemas y traduce lo que se alcanza desde los
// elementos globales de las raíces.
type generador struct {
	dir       string
	esquemas  map[string]*esquema // por archivo
	orden     []*esquema
	tipos     map[qn]*nodo // tipos simples globales
	tipoArch  map[qn]*esquema
	complejos map[qn]*nodo // tipos complejos globales
	usados    map[qn]bool  // tipos complejos que algún elemento usa
	globales  map[qn]*nodo // elementos globales
	globArch  map[qn]*esquema
	tiposGen  map[qn]*tipo // ya traducidos
	tiposOrd  []*tipo      // en orden de salida
	elems     []*elemento  // en orden de salida
	patrones  map[string]string
	patOrden  []string
	nombres   map[string]bool
	variantes map[string]bool // archivos de Variantes
}

// tipo es un tipo simple traducido.
type tipo struct {
	varGo   string
	nombre  string // para mensajes
	prim    string // primitivo XSD: string, decimal, int, short, integer, date, dateTime, base64Binary
	ws      string // preserve, replace, collapse
	facetas []faceta
}

type faceta struct {
	tipo     string // nombre XSD de la faceta
	n        int
	valor    string
	patrones []string // varGo de cada alternativa (pattern)
	catVar   string   // enumeration: variable del catálogo
	catNom   string
	valores  []string
}

type elemento struct {
	varGo      string
	ns, local  string
	attrs      []atributo
	vacio      bool
	todos      bool // xs:all: las partículas en cualquier orden
	seqMin     int
	seqMax     int // -1 = unbounded
	particulas []particula
	archivo    string
}

type atributo struct {
	nombre    string
	tipoVar   string
	requerido bool
	fijo      *string
}

type particula struct {
	elemVar  string // "" si es comodín o xs:choice
	comodin  bool
	opciones []particula // xs:choice: una partícula de elemento por opción
	min, max int
}

// Generar lee los XSD de dir (un directorio sitio_internet) y devuelve los
// dos archivos Go: el de estructura y el de catálogos.
func Generar(dir string) (reglas, catalogos []byte, err error) {
	g := &generador{
		dir: dir, esquemas: map[string]*esquema{}, tipos: map[qn]*nodo{}, tipoArch: map[qn]*esquema{},
		complejos: map[qn]*nodo{}, usados: map[qn]bool{},
		globales: map[qn]*nodo{}, globArch: map[qn]*esquema{}, tiposGen: map[qn]*tipo{},
		patrones: map[string]string{}, nombres: map[string]bool{}, variantes: map[string]bool{},
	}
	for _, r := range Raices {
		if err := g.cargar(r.Archivo, r.Prefijo, false); err != nil {
			return nil, nil, err
		}
	}
	for _, v := range Variantes {
		if g.esquemas[v.Archivo] != nil {
			return nil, nil, fmt.Errorf("%s: la variante ya es raíz o importado", v.Archivo)
		}
		if err := g.cargar(v.Archivo, v.Prefijo, true); err != nil {
			return nil, nil, err
		}
	}
	type raizGlobal struct {
		q    qn
		var_ string
	}
	var globales []raizGlobal
	porNombre := map[qn]*elemento{}
	for _, r := range Raices {
		e := g.esquemas[r.Archivo]
		for _, h := range e.raiz.hijos {
			if !h.xs("element") {
				continue
			}
			q := qn{e.tns, h.attrs["name"]}
			el, err := g.elemento(h, e, r.Prefijo+"_"+identificador(q.local), true)
			if err != nil {
				return nil, nil, err
			}
			globales = append(globales, raizGlobal{q, el.varGo})
			porNombre[q] = el
		}
	}
	type variante struct {
		q                qn
		atributo, valor  string
		var_, porDefecto string
	}
	var variantes []variante
	for _, v := range Variantes {
		e := g.esquemas[v.Archivo]
		for _, h := range e.raiz.hijos {
			if !h.xs("element") {
				continue
			}
			q := qn{e.tns, h.attrs["name"]}
			base := porNombre[q]
			if base == nil {
				return nil, nil, h.errorf("variante sin elemento global %s en las raíces", q)
			}
			el, err := g.elemento(h, e, v.Prefijo+"_"+identificador(q.local), true)
			if err != nil {
				return nil, nil, err
			}
			var sel *atributo
			for i := range el.attrs {
				if el.attrs[i].nombre == v.Atributo {
					sel = &el.attrs[i]
				}
			}
			if sel == nil || !sel.requerido || sel.fijo == nil {
				return nil, nil, h.errorf("la variante no tiene %q requerido y fixed", v.Atributo)
			}
			for _, a := range base.attrs {
				if a.nombre == v.Atributo && (a.fijo == nil || *a.fijo == *sel.fijo) {
					return nil, nil, h.errorf("%q también está en %s sin otro valor fixed: no distingue la variante", v.Atributo, base.archivo)
				}
			}
			variantes = append(variantes, variante{q, v.Atributo, *sel.fijo, el.varGo, base.varGo})
		}
	}
	if err := g.revisarNoAlcanzados(); err != nil {
		return nil, nil, err
	}
	// Un tipo complejo que nadie usa no se traduce: que no haya, para no
	// dejar sin revisar una construcción que mañana sí se use.
	for q, n := range g.complejos {
		if !g.usados[q] {
			return nil, nil, n.errorf("tipo complejo sin usar no soportado")
		}
	}
	// Los elementos globales de los importados no se alcanzan desde un
	// comodín si no son raíces: no se generan. Que no haya.
	for q, n := range g.globales {
		if e := g.globArch[q]; e != nil && Importados[e.archivo] != "" {
			return nil, nil, n.errorf("elemento global en un esquema importado no soportado")
		}
	}

	var b bytes.Buffer
	b.WriteString("// Código generado por `go run ./internal/genxsd/generar` a partir de los\n")
	b.WriteString("// XSD del SAT en internal/genxsd/xsd. NO EDITAR: cambiar el generador.\n\n")
	b.WriteString("package cfdi\n\n")
	b.WriteString("import \"regexp\"\n\n")
	b.WriteString("// Esquemas de los que salen estas tablas (sha256):\n")
	for _, e := range g.orden {
		fmt.Fprintf(&b, "//   - %s %s\n", e.archivo, e.sha)
	}
	b.WriteString("\n// globalesXSD son los elementos globales: la raíz del documento y lo que\n")
	b.WriteString("// puede venir en un xs:any (Complemento, ComplementoConcepto, Addenda).\n")
	b.WriteString("var globalesXSD = map[nombreXML]*elementoXSD{\n")
	for _, gl := range globales {
		fmt.Fprintf(&b, "\t{%q, %q}: %s,\n", gl.q.ns, gl.q.local, gl.var_)
	}
	b.WriteString("}\n\n")
	b.WriteString("// variantesXSD son elementos globales con el mismo namespace y nombre que\n")
	b.WriteString("// uno de globalesXSD (ver internal/genxsd.Variantes): se usa la variante\n")
	b.WriteString("// cuando el elemento trae su atributo con ese valor, y si no, el global.\n")
	b.WriteString("var variantesXSD = map[nombreXML][]varianteXSD{\n")
	var vistosQ []qn
	porQ := map[qn][]variante{}
	for _, v := range variantes {
		if porQ[v.q] == nil {
			vistosQ = append(vistosQ, v.q)
		}
		porQ[v.q] = append(porQ[v.q], v)
	}
	for _, q := range vistosQ {
		fmt.Fprintf(&b, "\t{%q, %q}: {\n", q.ns, q.local)
		for _, v := range porQ[q] {
			fmt.Fprintf(&b, "\t\t{atributo: %q, valor: %q, elem: %s}, // si no, %s\n", v.atributo, v.valor, v.var_, v.porDefecto)
		}
		b.WriteString("\t},\n")
	}
	b.WriteString("}\n\n")
	b.WriteString("// Patrones XSD traducidos a RE2 (ver internal/genxsd.traducirPatron).\n")
	b.WriteString("var (\n")
	for _, p := range g.patOrden {
		re, _ := traducirPatron(p)
		fmt.Fprintf(&b, "\t%s = patronXSD{xsd: %s, re: regexp.MustCompile(%s)}\n", g.patrones[p], literalGo(p), literalGo(re))
	}
	b.WriteString(")\n\n")
	for _, t := range g.tiposOrd {
		fmt.Fprintf(&b, "var %s = &tipoXSD{nombre: %q, prim: %q, ws: %q", t.varGo, t.nombre, t.prim, t.ws)
		if len(t.facetas) > 0 {
			b.WriteString(", facetas: []facetaXSD{\n")
			for _, f := range t.facetas {
				fmt.Fprintf(&b, "\t{tipo: %q", f.tipo)
				switch f.tipo {
				case "pattern":
					fmt.Fprintf(&b, ", patrones: []patronXSD{%s}", strings.Join(f.patrones, ", "))
				case "enumeration":
					fmt.Fprintf(&b, ", catalogo: %s", f.catVar)
				case "minInclusive", "maxInclusive":
					fmt.Fprintf(&b, ", valor: %q", f.valor)
				default:
					fmt.Fprintf(&b, ", n: %d", f.n)
				}
				b.WriteString("},\n")
			}
			b.WriteString("}")
		}
		b.WriteString("}\n")
	}
	b.WriteString("\n")
	for _, el := range g.elems {
		fmt.Fprintf(&b, "// %s: {%s}%s\n", el.archivo, el.ns, el.local)
		fmt.Fprintf(&b, "var %s = &elementoXSD{ns: %q, nombre: %q", el.varGo, el.ns, el.local)
		if el.vacio {
			b.WriteString(", vacio: true")
		} else {
			fmt.Fprintf(&b, ", seqMin: %d, seqMax: %d", el.seqMin, el.seqMax)
		}
		if el.todos {
			b.WriteString(", todos: true")
		}
		if len(el.attrs) > 0 {
			b.WriteString(",\n\tattrs: []atributoXSD{\n")
			for _, a := range el.attrs {
				fmt.Fprintf(&b, "\t\t{nombre: %q, tipo: %s", a.nombre, a.tipoVar)
				if a.requerido {
					b.WriteString(", requerido: true")
				}
				if a.fijo != nil {
					fmt.Fprintf(&b, ", fijo: %q, conFijo: true", *a.fijo)
				}
				b.WriteString("},\n")
			}
			b.WriteString("\t}")
		}
		if len(el.particulas) > 0 {
			b.WriteString(",\n\tparticulas: []particulaXSD{\n")
			for _, p := range el.particulas {
				switch {
				case p.comodin:
					fmt.Fprintf(&b, "\t\t{comodin: true, min: %d, max: %d},\n", p.min, p.max)
				case len(p.opciones) > 0:
					fmt.Fprintf(&b, "\t\t{min: %d, max: %d, opciones: []particulaXSD{\n", p.min, p.max)
					for _, o := range p.opciones {
						fmt.Fprintf(&b, "\t\t\t{elem: %s, min: %d, max: %d},\n", o.elemVar, o.min, o.max)
					}
					b.WriteString("\t\t}},\n")
				default:
					fmt.Fprintf(&b, "\t\t{elem: %s, min: %d, max: %d},\n", p.elemVar, p.min, p.max)
				}
			}
			b.WriteString("\t}")
		}
		b.WriteString("}\n\n")
	}
	reglas, err = format.Source(b.Bytes())
	if err != nil {
		return nil, nil, fmt.Errorf("gofmt de lo generado: %w", err)
	}

	var c bytes.Buffer
	c.WriteString("// Código generado por `go run ./internal/genxsd/generar` a partir de los\n")
	c.WriteString("// XSD del SAT en internal/genxsd/xsd. NO EDITAR: cambiar el generador.\n\n")
	c.WriteString("package cfdi\n\n")
	c.WriteString("// Catálogos (xs:enumeration) de los tipos que alcanzan los esquemas, con\n")
	c.WriteString("// un valor por línea, ya normalizados según el whiteSpace del tipo y\n")
	c.WriteString("// ordenados por bytes para la búsqueda binaria.\n")
	var vistos = map[string]bool{}
	for _, t := range g.tiposOrd {
		for _, f := range t.facetas {
			if f.tipo != "enumeration" || vistos[f.catVar] {
				continue
			}
			vistos[f.catVar] = true
			fmt.Fprintf(&c, "\n// %s: %d valores.\n", f.catNom, len(f.valores))
			fmt.Fprintf(&c, "var %s = &catalogoXSD{nombre: %q, valores: `\n%s\n`}\n", f.catVar, f.catNom, strings.Join(f.valores, "\n"))
		}
	}
	return reglas, c.Bytes(), nil
}

// cargar lee un esquema y lo que importa. Los elementos globales de una
// variante no se registran en g.globales: se llaman como los de una raíz.
func (g *generador) cargar(archivo, prefijo string, variante bool) error {
	if g.esquemas[archivo] != nil {
		return nil
	}
	raiz, b, err := leer(g.dir, archivo)
	if err != nil {
		return err
	}
	if !raiz.xs("schema") {
		return fmt.Errorf("%s: la raíz no es xs:schema", archivo)
	}
	if err := raiz.soloAttrs("targetNamespace", "elementFormDefault", "attributeFormDefault"); err != nil {
		return err
	}
	// Sin el atributo vale el default de XML Schema, que es unqualified
	// (catCombustible.xsd no lo declara).
	if v, ok := raiz.attrs["attributeFormDefault"]; ok && v != "unqualified" {
		return raiz.errorf("se esperaba attributeFormDefault=unqualified")
	}
	tns, ok := raiz.attrs["targetNamespace"]
	if !ok {
		return raiz.errorf("esquema sin targetNamespace no soportado")
	}
	h := sha256.Sum256(b)
	e := &esquema{archivo: archivo, prefijo: prefijo, tns: tns, raiz: raiz, sha: hex.EncodeToString(h[:])}
	g.esquemas[archivo] = e
	g.orden = append(g.orden, e)
	hijos, err := raiz.hijosXS("import", "element", "simpleType", "complexType")
	if err != nil {
		return err
	}
	// Los tipos con nombre de una variante son suyos: VentaVehiculos 1.0 y
	// 1.1 definen los dos t_InformacionAduanera en el mismo namespace. Se
	// registran con una clave propia (ver alcance).
	clave := func(local string) qn { return qn{tns, local} }
	if variante {
		g.variantes[archivo] = true
		clave = func(local string) qn { return qn{tns + "#" + archivo, local} }
	}
	for _, n := range hijos {
		switch n.local {
		case "import":
			if err := n.soloAttrs("namespace", "schemaLocation"); err != nil {
				return err
			}
			if _, err := n.hijosXS(); err != nil {
				return err
			}
			destino := normalizarRuta(path.Join(path.Dir(archivo), n.attrs["schemaLocation"]))
			pre, ok := Importados[destino]
			if !ok {
				return n.errorf("import de %s no previsto (agregarlo a Importados)", destino)
			}
			if err := g.cargar(destino, pre, false); err != nil {
				return err
			}
			if g.esquemas[destino].tns != n.attrs["namespace"] {
				return n.errorf("el namespace del import no es el targetNamespace de %s", destino)
			}
		case "simpleType":
			q := clave(n.attrs["name"])
			if q.local == "" {
				return n.errorf("tipo global sin nombre")
			}
			if g.tipos[q] != nil {
				return n.errorf("tipo %s repetido", q)
			}
			g.tipos[q], g.tipoArch[q] = n, e
		case "complexType":
			// Tipos complejos con nombre (terceros 1.1, ventavehiculos 1.1): se
			// traducen en cada elemento que los usa, como si fueran anónimos.
			q := clave(n.attrs["name"])
			if q.local == "" {
				return n.errorf("tipo complejo global sin nombre")
			}
			if g.complejos[q] != nil {
				return n.errorf("tipo complejo %s repetido", q)
			}
			g.complejos[q] = n
		case "element":
			// Los elementos locales van en el namespace del esquema.
			if raiz.attrs["elementFormDefault"] != "qualified" {
				return raiz.errorf("esquema con elementos y sin elementFormDefault=qualified")
			}
			if variante {
				continue
			}
			q := qn{tns, n.attrs["name"]}
			if g.globales[q] != nil {
				return n.errorf("elemento global %s repetido", q)
			}
			g.globales[q], g.globArch[q] = n, e
		}
	}
	return nil
}

// revisarNoAlcanzados traduce también los tipos globales que ningún
// elemento usa, sin generarlos: una construcción desconocida en un catálogo
// detiene la generación aunque hoy no se use, porque mañana sí.
func (g *generador) revisarNoAlcanzados() error {
	var faltan []qn
	for q := range g.tipos {
		if g.tiposGen[q] == nil {
			faltan = append(faltan, q)
		}
	}
	sort.Slice(faltan, func(i, j int) bool { return faltan[i].String() < faltan[j].String() })
	nTipos, nPat := len(g.tiposOrd), len(g.patOrden)
	for _, q := range faltan {
		if _, err := g.tipoPorNombre(q, g.tipos[q]); err != nil {
			return err
		}
	}
	for _, p := range g.patOrden[nPat:] {
		delete(g.patrones, p)
	}
	g.tiposOrd, g.patOrden = g.tiposOrd[:nTipos], g.patOrden[:nPat]
	return nil
}

// ocurrencias lee minOccurs y maxOccurs (-1 = unbounded).
func ocurrencias(n *nodo) (min, max int, err error) {
	min, max = 1, 1
	if v, ok := n.attrs["minOccurs"]; ok {
		if min, err = strconv.Atoi(v); err != nil || min < 0 {
			return 0, 0, n.errorf("minOccurs %q no soportado", v)
		}
	}
	if v, ok := n.attrs["maxOccurs"]; ok {
		if v == "unbounded" {
			max = -1
		} else if max, err = strconv.Atoi(v); err != nil || max < 1 {
			return 0, 0, n.errorf("maxOccurs %q no soportado", v)
		}
	}
	if max != -1 && max < min {
		return 0, 0, n.errorf("maxOccurs menor que minOccurs")
	}
	return min, max, nil
}

// elemento traduce un xs:element (global o local) con su complexType anónimo.
func (g *generador) elemento(n *nodo, e *esquema, nombre string, global bool) (*elemento, error) {
	permitidos := []string{"name", "minOccurs", "maxOccurs", "type"}
	if global {
		permitidos = []string{"name"}
	}
	if err := n.soloAttrs(permitidos...); err != nil {
		return nil, err
	}
	if n.attrs["name"] == "" {
		return nil, n.errorf("elemento sin nombre (ref no soportado)")
	}
	hijos, err := n.hijosXS("complexType")
	if err != nil {
		return nil, err
	}
	var ct *nodo
	switch tq, conTipo := n.attrs["type"]; {
	case conTipo && len(hijos) == 0:
		// Solo tipos complejos con nombre: un elemento de tipo simple
		// tendría texto, y eso no se soporta.
		if ct, err = g.complejo(n, tq); err != nil {
			return nil, err
		}
	case !conTipo && len(hijos) == 1:
		ct = hijos[0]
		if err := ct.soloAttrs(); err != nil {
			return nil, err
		}
	default:
		return nil, n.errorf("se esperaba exactamente un complexType anónimo o un type de tipo complejo")
	}
	partes, err := ct.hijosXS("sequence", "choice", "all", "attribute", "complexContent")
	if err != nil {
		return nil, err
	}
	if len(partes) == 1 && partes[0].local == "complexContent" {
		// <xs:complexContent><xs:extension base="T"/></xs:complexContent>
		// sin nada más es el mismo contenido de T.
		cc := partes[0]
		if err := cc.soloAttrs(); err != nil {
			return nil, err
		}
		ext, err := cc.hijosXS("extension")
		if err != nil {
			return nil, err
		}
		if len(ext) != 1 {
			return nil, cc.errorf("se esperaba exactamente una xs:extension (restriction no soportada)")
		}
		if err := ext[0].soloAttrs("base"); err != nil {
			return nil, err
		}
		if extra, err := ext[0].hijosXS(); err != nil || len(extra) > 0 {
			return nil, ext[0].errorf("xs:extension con contenido propio no soportada")
		}
		base, err := g.complejo(ext[0], ext[0].attrs["base"])
		if err != nil {
			return nil, err
		}
		if partes, err = base.hijosXS("sequence", "choice", "attribute"); err != nil {
			return nil, err
		}
	}
	el := &elemento{varGo: g.nombreUnico("e" + nombre), ns: e.tns, local: n.attrs["name"], vacio: true, archivo: e.archivo}
	vistosAttr := map[string]bool{}
	for i, p := range partes {
		switch p.local {
		case "complexContent":
			return nil, p.errorf("xs:complexContent junto con otras partes no soportado")
		case "choice":
			// Un xs:choice como contenido es una secuencia de una sola
			// partícula (ventavehiculos 1.1). Si se repite (Movimientos del
			// ECB), sus ocurrencias son las vueltas de esa secuencia: en cada
			// una va una opción.
			if i != 0 {
				return nil, p.errorf("xs:choice después de los atributos")
			}
			el.vacio = false
			if el.seqMin, el.seqMax, err = ocurrencias(p); err != nil {
				return nil, err
			}
			if el.seqMin > 1 {
				// El validador no cuenta vueltas mínimas: solo la partícula.
				return nil, p.errorf("xs:choice con minOccurs %q no soportado", p.attrs["minOccurs"])
			}
			vistos := map[string]bool{}
			pa, err := g.eleccion(p, e, nombre, vistos, true)
			if err != nil {
				return nil, err
			}
			el.particulas = append(el.particulas, pa)
		case "all":
			// xs:all (SPEI_Tercero): sus elementos en cualquier orden, cada
			// uno a lo más una vez (XML Schema 1.0 no admite otra cosa).
			if i != 0 {
				return nil, p.errorf("xs:all después de los atributos")
			}
			if err := p.soloAttrs("minOccurs"); err != nil {
				return nil, err
			}
			grupoMin, _, err := ocurrencias(p)
			if err != nil {
				return nil, err
			}
			items, err := p.hijosXS("element")
			if err != nil {
				return nil, err
			}
			if len(items) == 0 {
				return nil, p.errorf("xs:all vacío no soportado")
			}
			el.vacio, el.todos = false, true
			el.seqMin, el.seqMax = 1, 1
			vistos := map[string]bool{}
			for _, it := range items {
				pa, err := g.particulaElemento(it, e, nombre, vistos)
				if err != nil {
					return nil, err
				}
				if pa.max != 1 {
					return nil, it.errorf("elemento de xs:all con maxOccurs distinto de 1")
				}
				if grupoMin == 0 {
					pa.min = 0
				}
				el.particulas = append(el.particulas, pa)
			}
		case "sequence":
			if i != 0 {
				return nil, p.errorf("xs:sequence después de los atributos")
			}
			if err := p.soloAttrs("minOccurs", "maxOccurs"); err != nil {
				return nil, err
			}
			el.vacio = false
			if el.seqMin, el.seqMax, err = ocurrencias(p); err != nil {
				return nil, err
			}
			items, err := p.hijosXS("element", "any", "choice")
			if err != nil {
				return nil, err
			}
			if len(items) == 0 {
				return nil, p.errorf("xs:sequence vacía no soportada")
			}
			vistos := map[string]bool{}
			for _, it := range items {
				if it.local == "choice" {
					pa, err := g.eleccion(it, e, nombre, vistos, false)
					if err != nil {
						return nil, err
					}
					el.particulas = append(el.particulas, pa)
					continue
				}
				min, max, err := ocurrencias(it)
				if err != nil {
					return nil, err
				}
				if it.local == "any" {
					if err := it.soloAttrs("minOccurs", "maxOccurs", "namespace", "processContents"); err != nil {
						return nil, err
					}
					if v, ok := it.attrs["namespace"]; ok && v != "##any" {
						return nil, it.errorf("xs:any con namespace %q no soportado", v)
					}
					if v, ok := it.attrs["processContents"]; ok && v != "strict" {
						return nil, it.errorf("xs:any con processContents %q no soportado", v)
					}
					if _, err := it.hijosXS(); err != nil {
						return nil, err
					}
					if len(items) != 1 {
						return nil, it.errorf("xs:any junto con otros elementos no soportado")
					}
					el.particulas = append(el.particulas, particula{comodin: true, min: min, max: max})
					continue
				}
				pa, err := g.particulaElemento(it, e, nombre, vistos)
				if err != nil {
					return nil, err
				}
				el.particulas = append(el.particulas, pa)
			}
		case "attribute":
			a, err := g.atributo(p, e, nombre)
			if err != nil {
				return nil, err
			}
			if vistosAttr[a.nombre] {
				return nil, p.errorf("atributo repetido")
			}
			vistosAttr[a.nombre] = true
			el.attrs = append(el.attrs, a)
		}
	}
	g.elems = append(g.elems, el)
	return el, nil
}

// alcance da la clave de un tipo con nombre q visto desde un nodo: la de la
// variante, si desde está en una y ella lo define; si no, la de siempre.
func (g *generador) alcance(q qn, desde *nodo) qn {
	if desde == nil || !g.variantes[desde.archivo] {
		return q
	}
	s := qn{q.ns + "#" + desde.archivo, q.local}
	if g.tipos[s] != nil || g.complejos[s] != nil {
		return s
	}
	return q
}

// complejo busca el tipo complejo con nombre v (un qname) desde n.
func (g *generador) complejo(n *nodo, v string) (*nodo, error) {
	q, err := n.qname(v)
	if err != nil {
		return nil, err
	}
	q = g.alcance(q, n)
	ct := g.complejos[q]
	if ct == nil {
		return nil, n.errorf("tipo complejo %s no definido (elementos de tipo simple no soportados)", q)
	}
	if err := ct.soloAttrs("name"); err != nil {
		return nil, err
	}
	g.usados[q] = true
	return ct, nil
}

// particulaElemento traduce un xs:element dentro de una secuencia o de un
// xs:choice. vistos son los nombres ya usados en la secuencia.
func (g *generador) particulaElemento(it *nodo, e *esquema, nombre string, vistos map[string]bool) (particula, error) {
	min, max, err := ocurrencias(it)
	if err != nil {
		return particula{}, err
	}
	local := it.attrs["name"]
	if vistos[local] {
		// Con nombres repetidos el emparejamiento por nombre
		// dejaría de ser determinista.
		return particula{}, it.errorf("elemento %q repetido en la secuencia", local)
	}
	vistos[local] = true
	hijo, err := g.elemento(it, e, nombre+"_"+identificador(local), false)
	if err != nil {
		return particula{}, err
	}
	return particula{elemVar: hijo.varGo, min: min, max: max}, nil
}

// eleccion traduce un xs:choice de elementos. Dentro de una secuencia no
// se repite (maxOccurs 1); como todo el contenido (contenido) puede
// repetirse, y su min/max ya los tomó la secuencia: la partícula es una sola
// opción por vuelta. Sus opciones comparten los nombres de la secuencia.
func (g *generador) eleccion(c *nodo, e *esquema, nombre string, vistos map[string]bool, contenido bool) (particula, error) {
	if err := c.soloAttrs("minOccurs", "maxOccurs"); err != nil {
		return particula{}, err
	}
	min, max, err := ocurrencias(c)
	if err != nil {
		return particula{}, err
	}
	if contenido {
		max = 1
	}
	if max != 1 {
		return particula{}, c.errorf("xs:choice con maxOccurs %q no soportado", c.attrs["maxOccurs"])
	}
	items, err := c.hijosXS("element")
	if err != nil {
		return particula{}, err
	}
	if len(items) < 2 {
		return particula{}, c.errorf("xs:choice con menos de dos opciones no soportado")
	}
	pa := particula{min: min, max: 1}
	for _, it := range items {
		op, err := g.particulaElemento(it, e, nombre, vistos)
		if err != nil {
			return particula{}, err
		}
		pa.opciones = append(pa.opciones, op)
	}
	return pa, nil
}

func (g *generador) atributo(n *nodo, e *esquema, elem string) (atributo, error) {
	if err := n.soloAttrs("name", "type", "use", "fixed", "id"); err != nil {
		return atributo{}, err
	}
	a := atributo{nombre: n.attrs["name"]}
	if a.nombre == "" {
		return a, n.errorf("atributo sin nombre (ref no soportado)")
	}
	switch n.attrs["use"] {
	case "required":
		a.requerido = true
	case "optional", "":
	default:
		return a, n.errorf("use %q no soportado", n.attrs["use"])
	}
	hijos, err := n.hijosXS("simpleType")
	if err != nil {
		return a, err
	}
	var t *tipo
	switch tq, conTipo := n.attrs["type"]; {
	case conTipo && len(hijos) == 0:
		q, err := n.qname(tq)
		if err != nil {
			return a, err
		}
		if t, err = g.tipoPorNombre(q, n); err != nil {
			return a, err
		}
	case !conTipo && len(hijos) == 1:
		if t, err = g.tipoSimple(hijos[0], e, "t"+elem+"_"+identificador(a.nombre), elem2nombre(elem)+"/@"+a.nombre); err != nil {
			return a, err
		}
	case !conTipo && len(hijos) == 0:
		// Sin tipo es xs:anySimpleType (TFD 1.1 @Version, que es fixed).
		if t, err = g.tipoPorNombre(qn{nsXS, "anySimpleType"}, n); err != nil {
			return a, err
		}
	default:
		return a, n.errorf("se esperaba type o un simpleType anónimo")
	}
	a.tipoVar = t.varGo
	if v, ok := n.attrs["fixed"]; ok {
		// En los tipos no textuales fixed compara valores ("1.0" = "1"); no
		// hace falta y no se soporta.
		if t.prim != "string" && t.prim != "anySimpleType" {
			return a, n.errorf("fixed en un tipo derivado de xs:%s no soportado", t.prim)
		}
		f := normalizar(v, t.ws)
		a.fijo = &f
	}
	return a, nil
}

// elem2nombre convierte "CFDI40_Comprobante_Emisor" en "CFDI40 Comprobante/Emisor"
// para los mensajes.
func elem2nombre(s string) string {
	pre, resto, _ := strings.Cut(s, "_")
	return pre + " " + strings.ReplaceAll(resto, "_", "/")
}

// primitivos que se saben validar, con el whiteSpace que imponen.
var primitivos = map[string]string{
	"anySimpleType": "preserve",
	"string":        "preserve",
	"decimal":       "collapse",
	"integer":       "collapse",
	"int":           "collapse",
	"short":         "collapse",
	"date":          "collapse",
	"dateTime":      "collapse",
	"time":          "collapse",
	"base64Binary":  "collapse",
}

// facetas permitidas por primitivo.
var facetasPermitidas = map[string][]string{
	"anySimpleType": {},
	"string":        {"pattern", "enumeration", "length", "minLength", "maxLength", "whiteSpace"},
	"decimal":       {"pattern", "enumeration", "minInclusive", "maxInclusive", "fractionDigits", "totalDigits", "whiteSpace"}, // enumeration: c_TasaCuota de Plataformas
	"integer":       {"pattern", "minInclusive", "maxInclusive", "totalDigits", "fractionDigits", "whiteSpace"},
	"int":           {"pattern", "enumeration", "minInclusive", "maxInclusive", "totalDigits", "fractionDigits", "whiteSpace"}, // fractionDigits: solo 0 (INE)
	"short":         {"pattern", "minInclusive", "maxInclusive", "totalDigits", "fractionDigits", "whiteSpace"},
	"date":          {"pattern", "minInclusive", "maxInclusive", "whiteSpace"}, // rango: FechaOperacion de SPEI
	"dateTime":      {"pattern", "minInclusive", "whiteSpace"},                 // minInclusive: FechaExp de Retenciones 1.0, Fecha de ECC 1.1
	"time":          {"pattern", "whiteSpace"},                                 // Hora de SPEI
	"base64Binary":  {"pattern", "whiteSpace"},
}

func (g *generador) tipoPorNombre(q qn, desde *nodo) (*tipo, error) {
	if q.ns == nsXS {
		ws, ok := primitivos[q.local]
		if !ok {
			return nil, desde.errorf("tipo integrado xs:%s no soportado", q.local)
		}
		varGo := "xs" + identificador(q.local)
		if t := g.tiposGen[q]; t != nil {
			return t, nil
		}
		t := &tipo{varGo: varGo, nombre: "xs:" + q.local, prim: q.local, ws: ws}
		g.tiposGen[q] = t
		g.tiposOrd = append(g.tiposOrd, t)
		return t, nil
	}
	q = g.alcance(q, desde)
	if t := g.tiposGen[q]; t != nil {
		return t, nil
	}
	n := g.tipos[q]
	if n == nil {
		return nil, desde.errorf("tipo %s no definido", q)
	}
	e := g.tipoArch[q]
	t, err := g.tipoSimple(n, e, "t"+e.prefijo+"_"+identificador(q.local), q.local)
	if err != nil {
		return nil, err
	}
	g.tiposGen[q] = t
	return t, nil
}

// tipoSimple traduce un xs:simpleType (con nombre o anónimo). Las facetas
// de la base van primero; todas se exigen (en un esquema válido las del
// tipo derivado restringen las de la base, así que exigir ambas es lo
// mismo que exigir la derivada).
func (g *generador) tipoSimple(n *nodo, e *esquema, varGo, nombre string) (*tipo, error) {
	permitidos := []string{}
	if n.attrs["name"] != "" {
		permitidos = []string{"name"}
	}
	if err := n.soloAttrs(permitidos...); err != nil {
		return nil, err
	}
	hijos, err := n.hijosXS("restriction")
	if err != nil {
		return nil, err
	}
	if len(hijos) != 1 {
		return nil, n.errorf("se esperaba exactamente una xs:restriction (list y union no soportados)")
	}
	r := hijos[0]
	if err := r.soloAttrs("base"); err != nil {
		return nil, err
	}
	bq, err := r.qname(r.attrs["base"])
	if err != nil {
		return nil, err
	}
	base, err := g.tipoPorNombre(bq, r)
	if err != nil {
		return nil, err
	}
	t := &tipo{varGo: g.nombreUnico(varGo), nombre: nombre, prim: base.prim, ws: base.ws}
	t.facetas = append(t.facetas, base.facetas...)
	facetas, err := r.hijosXS("pattern", "enumeration", "minInclusive", "maxInclusive", "fractionDigits",
		"totalDigits", "length", "minLength", "maxLength", "whiteSpace", "minExclusive", "maxExclusive", "simpleType")
	if err != nil {
		return nil, err
	}
	var patrones []string
	var enum []string
	// whiteSpace primero: normaliza los valores de las demás facetas.
	for _, f := range facetas {
		if !f.xs("whiteSpace") {
			continue
		}
		v := f.attrs["value"]
		switch {
		case v != "preserve" && v != "replace" && v != "collapse":
			return nil, f.errorf("whiteSpace %q no soportado", v)
		case t.prim != "string" && v != "collapse":
			return nil, f.errorf("whiteSpace %q en un tipo que no es string", v)
		}
		t.ws = v
	}
	for _, f := range facetas {
		permitida := false
		for _, p := range facetasPermitidas[t.prim] {
			permitida = permitida || f.xs(p)
		}
		if !permitida {
			return nil, f.errorf("faceta no soportada en un tipo derivado de xs:%s", t.prim)
		}
		if err := f.soloAttrs("value"); err != nil {
			return nil, err
		}
		v, ok := f.attrs["value"]
		if !ok {
			return nil, f.errorf("faceta sin value")
		}
		switch f.local {
		case "enumeration":
			// enumeration puede traer su documentación.
			if _, err := f.hijosXS(); err != nil {
				return nil, err
			}
		default:
			if len(f.hijos) > 0 {
				if _, err := f.hijosXS(); err != nil {
					return nil, err
				}
			}
		}
		switch f.local {
		case "whiteSpace":
		case "pattern":
			if _, err := traducirPatron(v); err != nil {
				return nil, f.errorf("%v", err)
			}
			patrones = append(patrones, g.patron(v))
		case "enumeration":
			v = normalizar(v, t.ws)
			if t.prim == "int" {
				// En un tipo numérico enumeration compara valores ("01" = "1"):
				// el catálogo guarda la forma canónica y ValidarEsquema
				// busca la del valor.
				d, err := decimal.Parse(v)
				if err != nil || d.Scale() != 0 {
					return nil, f.errorf("enumeration %q no es un entero", v)
				}
				v = d.String()
			}
			if t.prim == "decimal" {
				// Igual en xs:decimal: "0.16" = "0.160000".
				d, err := decimal.Parse(v)
				if err != nil {
					return nil, f.errorf("enumeration %q no es un decimal", v)
				}
				v = d.String()
			}
			enum = append(enum, v)
		case "minInclusive", "maxInclusive":
			if t.prim == "dateTime" {
				// Se compara como instante (con zona) o como hora sin zona.
				if !reFechaHoraLimite.MatchString(v) {
					return nil, f.errorf("%s %q no es un xs:dateTime", f.local, v)
				}
			} else if t.prim == "date" {
				// Como un xs:dateTime a las 00:00:00 del día.
				if !reFechaLimite.MatchString(v) {
					return nil, f.errorf("%s %q no es un xs:date", f.local, v)
				}
			} else if _, err := decimal.Parse(v); err != nil {
				return nil, f.errorf("%s %q no es decimal", f.local, v)
			}
			t.facetas = append(t.facetas, faceta{tipo: f.local, valor: v})
		default:
			k, err := strconv.Atoi(v)
			if err != nil || k < 0 {
				return nil, f.errorf("%s %q no es un entero no negativo", f.local, v)
			}
			if f.local == "fractionDigits" && t.prim != "decimal" {
				// En los enteros fractionDigits es fixed 0: otro valor es un
				// esquema inválido, y 0 no restringe nada.
				if k != 0 {
					return nil, f.errorf("fractionDigits %d en un tipo entero (xs:%s)", k, t.prim)
				}
				continue
			}
			t.facetas = append(t.facetas, faceta{tipo: f.local, n: k})
		}
	}
	if len(patrones) > 0 {
		// Varios pattern en un mismo paso son alternativas (OR).
		t.facetas = append(t.facetas, faceta{tipo: "pattern", patrones: patrones})
	}
	if len(enum) > 0 {
		f, err := g.catalogo(t, e, nombre, enum, r)
		if err != nil {
			return nil, err
		}
		t.facetas = append(t.facetas, f)
	}
	g.tiposOrd = append(g.tiposOrd, t)
	return t, nil
}

func (g *generador) catalogo(t *tipo, e *esquema, nombre string, enum []string, desde *nodo) (faceta, error) {
	unicos := map[string]bool{}
	var vals []string
	for _, v := range enum {
		if v == "" || strings.ContainsAny(v, "`\n\r") {
			return faceta{}, desde.errorf("valor de enumeration %q no se puede guardar", v)
		}
		if !unicos[v] {
			unicos[v] = true
			vals = append(vals, v)
		}
	}
	sort.Strings(vals)
	cv := g.nombreUnico("cat" + strings.TrimPrefix(t.varGo, "t"))
	return faceta{tipo: "enumeration", catVar: cv, catNom: nombre, valores: vals}, nil
}

func (g *generador) patron(p string) string {
	if v, ok := g.patrones[p]; ok {
		return v
	}
	v := fmt.Sprintf("patron%03d", len(g.patOrden)+1)
	g.patrones[p] = v
	g.patOrden = append(g.patOrden, p)
	return v
}

func (g *generador) nombreUnico(s string) string {
	if !g.nombres[s] {
		g.nombres[s] = true
		return s
	}
	for i := 2; ; i++ {
		c := fmt.Sprintf("%s%d", s, i)
		if !g.nombres[c] {
			g.nombres[c] = true
			return c
		}
	}
}

// normalizar aplica whiteSpace a un valor.
func normalizar(v, ws string) string {
	switch ws {
	case "replace":
		return strings.Map(func(r rune) rune {
			if r == '\t' || r == '\n' || r == '\r' {
				return ' '
			}
			return r
		}, v)
	case "collapse":
		return strings.Join(strings.FieldsFunc(v, func(r rune) bool {
			return r == ' ' || r == '\t' || r == '\n' || r == '\r'
		}), " ")
	}
	return v
}

// literalGo escribe s como literal Go, entre comillas invertidas si se
// puede (los patrones se leen mejor sin escapar las barras).
func literalGo(s string) string {
	if strconv.CanBackquote(s) {
		return "`" + s + "`"
	}
	return strconv.Quote(s)
}

// identificador vuelve un nombre XML un pedazo de identificador Go ASCII:
// "Año" → "Anio", "Antigüedad" → "Antiguedad".
func identificador(s string) string {
	var b strings.Builder
	for _, r := range s {
		if t, ok := translit[r]; ok {
			b.WriteString(t)
		} else if r < 0x80 && (unicode.IsLetter(r) || unicode.IsDigit(r)) {
			b.WriteRune(r)
		}
	}
	return b.String()
}

var translit = map[rune]string{
	'ñ': "ni", 'Ñ': "Ni",
	'á': "a", 'é': "e", 'í': "i", 'ó': "o", 'ú': "u", 'ü': "u",
	'Á': "A", 'É': "E", 'Í': "I", 'Ó': "O", 'Ú': "U", 'Ü': "U",
}
