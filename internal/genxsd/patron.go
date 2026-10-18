package genxsd

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// traducirPatron convierte una expresión regular de XML Schema (XSD 1.0,
// apéndice F) en una de RE2 (regexp de Go) con el mismo lenguaje.
//
// Traducciones:
//
//   - Anclaje implícito: en XSD el patrón tiene que cubrir el valor entero;
//     se envuelve en ^(?:…)$. En RE2 sin la bandera m, $ solo coincide al
//     final del texto.
//   - ^ y $ son literales en XSD: se escapan.
//   - . en XSD es cualquier carácter salvo \n y \r: [^\n\r] (en RE2 . sí
//     acepta \r).
//   - ( abre un grupo sin captura, (?:…), para no pagar capturas.
//   - \s y \S son [ \t\n\r] y su complemento (RE2 incluye \f en \s).
//   - \d y \D son \p{Nd} y \P{Nd} (en RE2 \d es solo ASCII).
//   - \w y \W son [^\p{P}\p{Z}\p{C}] y [\p{P}\p{Z}\p{C}] fuera de una clase.
//   - \p{X} y \P{X} con una categoría general Unicode pasan tal cual.
//   - Escapes de un carácter (\n \r \t \\ \| \. \- \^ \? \* \+ \{ \} \( \)
//     \[ \]) pasan tal cual.
//   - En una clase, - al principio o al final es literal y se escapa, igual
//     que ^ fuera de la primera posición.
//
// Falla (en vez de traducir mal) con: resta de clases ([a-z-[aeiou]]),
// bloques Unicode (\p{IsBasicLatin}), \i \I \c \C, \w \W \S \D dentro de
// una clase negada, cuantificadores de más de 1000 (el límite de RE2) y
// cualquier escape o sintaxis que no sea de XSD.
func traducirPatron(p string) (string, error) {
	t := &traductor{src: []rune(p)}
	var b strings.Builder
	b.WriteString("^(?:")
	if err := t.expr(&b); err != nil {
		return "", fmt.Errorf("patrón %q: %w", p, err)
	}
	b.WriteString(")$")
	return b.String(), nil
}

type traductor struct {
	src   []rune
	i     int
	nivel int
}

func (t *traductor) fin() bool { return t.i >= len(t.src) }

func (t *traductor) expr(b *strings.Builder) error {
	for !t.fin() {
		c := t.src[t.i]
		switch c {
		case '(':
			t.i++
			t.nivel++
			b.WriteString("(?:")
		case ')':
			if t.nivel == 0 {
				return fmt.Errorf("paréntesis sin abrir en la posición %d", t.i)
			}
			t.i++
			t.nivel--
			b.WriteByte(')')
			if err := t.cuantificador(b); err != nil {
				return err
			}
			continue
		case '|':
			t.i++
			b.WriteByte('|')
		case '*', '+', '?', '{':
			return fmt.Errorf("cuantificador sin átomo en la posición %d", t.i)
		case '}', ']':
			return fmt.Errorf("%q suelto en la posición %d", c, t.i)
		case '[':
			if err := t.clase(b); err != nil {
				return err
			}
			if err := t.cuantificador(b); err != nil {
				return err
			}
			continue
		case '.':
			t.i++
			b.WriteString(`[^\n\r]`)
			if err := t.cuantificador(b); err != nil {
				return err
			}
			continue
		case '\\':
			s, err := t.escape(false)
			if err != nil {
				return err
			}
			b.WriteString(s)
			if err := t.cuantificador(b); err != nil {
				return err
			}
			continue
		default:
			t.i++
			b.WriteString(literal(c))
			if err := t.cuantificador(b); err != nil {
				return err
			}
			continue
		}
	}
	if t.nivel != 0 {
		return fmt.Errorf("paréntesis sin cerrar")
	}
	return nil
}

// literal escapa lo que en RE2 es especial y en XSD no.
func literal(c rune) string {
	if strings.ContainsRune(`^$\.[]{}()*+?|`, c) {
		return `\` + string(c)
	}
	return string(c)
}

func (t *traductor) cuantificador(b *strings.Builder) error {
	if t.fin() {
		return nil
	}
	switch c := t.src[t.i]; c {
	case '*', '+', '?':
		t.i++
		b.WriteRune(c)
	case '{':
		j := t.i + 1
		for j < len(t.src) && t.src[j] != '}' {
			j++
		}
		if j >= len(t.src) {
			return fmt.Errorf("cuantificador sin cerrar")
		}
		cuerpo := string(t.src[t.i+1 : j])
		minS, maxS, coma := strings.Cut(cuerpo, ",")
		min, err := strconv.Atoi(minS)
		if err != nil || min < 0 {
			return fmt.Errorf("cuantificador {%s} no soportado", cuerpo)
		}
		limite := min
		if coma && maxS != "" {
			max, err := strconv.Atoi(maxS)
			if err != nil || max < min {
				return fmt.Errorf("cuantificador {%s} no soportado", cuerpo)
			}
			limite = max
		}
		if limite > 1000 {
			return fmt.Errorf("cuantificador {%s} pasa del límite de RE2 (1000)", cuerpo)
		}
		t.i = j + 1
		b.WriteString("{" + cuerpo + "}")
	default:
		return nil
	}
	// XSD no tiene cuantificadores perezosos ni posesivos: un segundo
	// cuantificador es un error de sintaxis.
	if !t.fin() && strings.ContainsRune("*+?{", t.src[t.i]) {
		return fmt.Errorf("cuantificador doble en la posición %d", t.i)
	}
	return nil
}

// escape traduce un escape que empieza en t.i ('\\'). enClase indica si
// está dentro de [...]: ahí lo que se devuelve no lleva corchetes.
func (t *traductor) escape(enClase bool) (string, error) {
	t.i++
	if t.fin() {
		return "", fmt.Errorf("\\ al final")
	}
	c := t.src[t.i]
	t.i++
	switch c {
	case 'n':
		return `\n`, nil
	case 'r':
		return `\r`, nil
	case 't':
		return `\t`, nil
	case '\\', '|', '.', '-', '^', '?', '*', '+', '{', '}', '(', ')', '[', ']':
		return `\` + string(c), nil
	case 's':
		if enClase {
			return `\x20\t\n\r`, nil
		}
		return `[\x20\t\n\r]`, nil
	case 'S':
		if enClase {
			return "", fmt.Errorf(`\S dentro de una clase no soportado`)
		}
		return `[^\x20\t\n\r]`, nil
	case 'd':
		return `\p{Nd}`, nil
	case 'D':
		return `\P{Nd}`, nil
	case 'w':
		if enClase {
			return "", fmt.Errorf(`\w dentro de una clase no soportado`)
		}
		return `[^\p{P}\p{Z}\p{C}]`, nil
	case 'W':
		if enClase {
			return "", fmt.Errorf(`\W dentro de una clase no soportado`)
		}
		return `[\p{P}\p{Z}\p{C}]`, nil
	case 'p', 'P':
		if t.fin() || t.src[t.i] != '{' {
			return "", fmt.Errorf(`\%c sin {`, c)
		}
		j := t.i + 1
		for j < len(t.src) && t.src[j] != '}' {
			j++
		}
		if j >= len(t.src) {
			return "", fmt.Errorf(`\%c{ sin cerrar`, c)
		}
		cat := string(t.src[t.i+1 : j])
		t.i = j + 1
		if strings.HasPrefix(cat, "Is") {
			return "", fmt.Errorf(`bloque Unicode \%c{%s} no soportado`, c, cat)
		}
		if _, ok := unicode.Categories[cat]; !ok {
			return "", fmt.Errorf(`categoría Unicode \%c{%s} desconocida`, c, cat)
		}
		return `\` + string(c) + "{" + cat + "}", nil
	}
	return "", fmt.Errorf(`escape \%c no soportado`, c)
}

// clase traduce [...] o [^...].
func (t *traductor) clase(b *strings.Builder) error {
	t.i++ // [
	var out strings.Builder
	out.WriteByte('[')
	if !t.fin() && t.src[t.i] == '^' {
		out.WriteByte('^')
		t.i++
	}
	inicio := t.i
	for {
		if t.fin() {
			return fmt.Errorf("clase sin cerrar")
		}
		c := t.src[t.i]
		switch {
		case c == ']':
			if t.i == inicio {
				return fmt.Errorf("clase vacía")
			}
			t.i++
			out.WriteByte(']')
			b.WriteString(out.String())
			return nil
		case c == '[':
			return fmt.Errorf("[ dentro de una clase no soportado")
		case c == '-' && !t.fin() && t.i+1 < len(t.src) && t.src[t.i+1] == '[':
			return fmt.Errorf("resta de clases no soportada")
		case c == '-' && (t.i == inicio || (t.i+1 < len(t.src) && t.src[t.i+1] == ']')):
			t.i++
			out.WriteString(`\-`)
		case c == '-':
			// Rango: lo anterior ya se escribió; el guion pasa tal cual.
			t.i++
			out.WriteByte('-')
		case c == '\\':
			s, err := t.escape(true)
			if err != nil {
				return err
			}
			out.WriteString(s)
		case c == '^':
			t.i++
			out.WriteString(`\^`)
		default:
			t.i++
			out.WriteRune(c)
		}
	}
}
