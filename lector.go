package cfdi

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/OrtaMarco/cfdi-go/decimal"
)

// lector acumula advertencias mientras convierte el árbol en el modelo.
type lector struct {
	adv     []Advertencia
	antigua bool // CFDI 3.0/3.2: no se exigen los atributos que introdujo 3.3
	op      Opciones
}

func (l *lector) advertir(codigo, ruta, msg string) {
	l.advertirValor(codigo, ruta, msg, "")
}

// advertirValor es advertir con el texto original del atributo. El Atributo
// sale de la ruta, para que Ruta y Atributo no puedan contradecirse.
func (l *lector) advertirValor(codigo, ruta, msg, valor string) {
	var attr string
	if i := strings.LastIndex(ruta, "/@"); i >= 0 {
		attr = ruta[i+2:]
	}
	l.adv = append(l.adv, Advertencia{Codigo: codigo, Ruta: ruta, Mensaje: msg, Atributo: attr, Valor: valor})
}

// hijos devuelve los hijos con ese nombre local. Los que están en el namespace
// esperado van siempre; los de otro namespace también se aceptan, con
// advertencia, porque hay emisores que mezclan cfd/3 y cfd/4 o pierden el
// prefijo.
func (l *lector) hijos(n *nodo, ns, local, ruta string) []*nodo {
	var out []*nodo
	for _, h := range n.hijos {
		if h.local != local {
			continue
		}
		if h.ns != ns {
			l.advertir(AdvNamespaceInesperado, ruta+"/"+local,
				fmt.Sprintf("se esperaba el namespace %q y vino %q", ns, h.ns))
		}
		out = append(out, h)
	}
	return out
}

// unico devuelve el primer hijo con ese nombre y advierte si hay más de uno.
func (l *lector) unico(n *nodo, ns, local, ruta string) *nodo {
	hs := l.hijos(n, ns, local, ruta)
	if len(hs) == 0 {
		return nil
	}
	if len(hs) > 1 {
		l.advertir(AdvNodoDuplicado, ruta+"/"+local, fmt.Sprintf("%d nodos %s; se usa el primero", len(hs), local))
	}
	return hs[0]
}

// requerido advierte si faltan atributos.
func (l *lector) requerido(n *nodo, ruta string, nombres ...string) {
	for _, nom := range nombres {
		if _, ok := n.attr(nom); !ok {
			l.advertir(AdvAtributoFaltante, ruta+"/@"+nom, "falta el atributo "+nom)
		}
	}
}

func (l *lector) str(n *nodo, nombre string) string {
	v, _ := n.attr(nombre)
	return v
}

// dec lee un atributo decimal. Ausente → nil. Ilegible → nil con advertencia,
// que conserva el texto en Valor.
func (l *lector) dec(n *nodo, ruta, nombre string) *decimal.Decimal {
	v, ok := n.attr(nombre)
	if !ok {
		return nil
	}
	return l.decValor(v, ruta+"/@"+nombre)
}

// decValor interpreta un decimal ya extraído (de un atributo o del texto de un
// elemento); ruta es la del campo, con su "/@Nombre" si es un atributo.
// decimal.Parse ya recorta los espacios XML.
func (l *lector) decValor(v, ruta string) *decimal.Decimal {
	d, err := decimal.Parse(v)
	if err != nil {
		l.advertirValor(AdvDecimalInvalido, ruta, fmt.Sprintf("%q no es un decimal válido", v), v)
		return nil
	}
	return &d
}

func (l *lector) fecha(n *nodo, ruta, nombre string, soloFecha bool) Fecha {
	v, ok := n.attr(nombre)
	if !ok {
		return Fecha{}
	}
	return l.fechaValor(v, ruta+"/@"+nombre, soloFecha)
}

// fechaValor es fecha sobre un valor ya extraído; ruta es la del campo, con su
// "/@Nombre" si es un atributo.
func (l *lector) fechaValor(v, ruta string, soloFecha bool) Fecha {
	f, raro, err := parseFecha(v, soloFecha)
	esperado := layoutFechaHora
	if soloFecha {
		esperado = layoutFecha
	}
	switch {
	case err != nil:
		l.advertirValor(AdvFechaInvalida, ruta, err.Error(), v)
	case raro && l.op.FechaEstricta:
		l.advertirValor(AdvFechaInvalida, ruta, fmt.Sprintf("%q no tiene el formato %s (lectura estricta)", v, esperado), v)
		return Fecha{Original: v, FueraDeFormato: true}
	case raro:
		l.advertirValor(AdvFechaFormato, ruta, fmt.Sprintf("%q no tiene el formato %s", v, esperado), v)
	}
	return f
}

var reUUID = regexp.MustCompile(`^[0-9A-F]{8}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{12}$`)

// uuid normaliza un UUID (sin espacios, mayúsculas) y devuelve el original
// solo si cambió.
func (l *lector) uuid(n *nodo, ruta, nombre string) (norm, original string) {
	v, ok := n.attr(nombre)
	if !ok {
		return "", ""
	}
	norm = strings.ToUpper(strings.TrimSpace(v))
	if !reUUID.MatchString(norm) {
		l.advertirValor(AdvUUIDInvalido, ruta+"/@"+nombre, fmt.Sprintf("%q no es un UUID", v), v)
	}
	if norm != v {
		original = v
	}
	return norm, original
}

func rfc(s string) string { return strings.ToUpper(strings.TrimSpace(s)) }

func idx(ruta, local string, i int) string { return fmt.Sprintf("%s/%s[%d]", ruta, local, i+1) }
