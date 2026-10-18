package cfdi

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strings"
	"unicode/utf16"
	"unicode/utf8"
)

// normalizarCodificacion convierte el documento a UTF-8 antes de dárselo a
// encoding/xml, que solo entiende UTF-8.
//
//   - BOM UTF-8: se quita. BOM UTF-16 (LE/BE): se convierte.
//   - encoding declarado ISO-8859-1/latin1/windows-1252/ASCII: se convierte
//     desde windows-1252 (superconjunto práctico de latin1, como hace WHATWG).
//     Si los bytes ya son UTF-8 válido con caracteres no ASCII, la
//     declaración miente: se dejan como están y se advierte.
//   - encoding UTF-8 (o sin declaración) con bytes que no son UTF-8 válido: se
//     convierte desde windows-1252 y se advierte.
//   - Por último se quitan los caracteres de control que XML 1.0 prohíbe.
func normalizarCodificacion(b []byte, adv func(codigo, msg string)) []byte {
	switch {
	case bytes.HasPrefix(b, []byte{0xEF, 0xBB, 0xBF}):
		b = b[3:]
	case bytes.HasPrefix(b, []byte{0xFF, 0xFE}):
		b = desdeUTF16(b[2:], false)
	case bytes.HasPrefix(b, []byte{0xFE, 0xFF}):
		b = desdeUTF16(b[2:], true)
	}
	// Relleno de NUL al final (archivos copiados de bloques fijos): se quita
	// sin advertir.
	b = bytes.TrimRight(b, "\x00 \t\r\n")

	declarado := encodingDeclarado(b)
	valido := utf8.Valid(b)
	switch etiqueta := strings.ToLower(declarado); etiqueta {
	case "", "utf-8", "utf8", "utf-16", "utf16":
		// utf-16 solo llega aquí ya convertido por el BOM, o mintiendo.
		if !valido {
			b = desdeWindows1252(b)
			msg := "sin declaración de encoding"
			if declarado != "" {
				msg = "declara " + declarado
			}
			adv(AdvEncodingCorregido, msg+" pero los bytes no son UTF-8 válido; se leyó como windows-1252")
		}
	case "iso-8859-1", "iso8859-1", "latin1", "latin-1", "l1", "windows-1252", "cp1252", "iso-8859-15", "us-ascii", "ascii":
		switch {
		case valido && !esASCII(b):
			adv(AdvEncodingCorregido, fmt.Sprintf("declara %s pero los bytes son UTF-8; se leyó como UTF-8", declarado))
		case !valido:
			b = desdeWindows1252(b)
		}
	default:
		if !valido {
			b = desdeWindows1252(b)
			adv(AdvEncodingDesconocido, fmt.Sprintf("encoding %q desconocido; se leyó como windows-1252", declarado))
		} else {
			adv(AdvEncodingDesconocido, fmt.Sprintf("encoding %q desconocido; se leyó como UTF-8", declarado))
		}
	}

	if n := contarInvalidos(b); n > 0 {
		b = quitarInvalidos(b)
		adv(AdvCaracterInvalido, fmt.Sprintf("se quitaron %d caracteres de control no permitidos en XML", n))
	}
	return b
}

var reEncoding = regexp.MustCompile(`^<\?xml[^>]*?\sencoding\s*=\s*["']([^"']*)["']`)

func encodingDeclarado(b []byte) string {
	head := b
	if len(head) > 512 {
		head = head[:512]
	}
	head = bytes.TrimLeft(head, " \t\r\n")
	if m := reEncoding.FindSubmatch(head); m != nil {
		return strings.TrimSpace(string(m[1]))
	}
	return ""
}

func esASCII(b []byte) bool {
	for _, c := range b {
		if c >= 0x80 {
			return false
		}
	}
	return true
}

// cp1252 mapea 0x80–0x9F; los huecos de windows-1252 quedan como el control C1
// del mismo número.
var cp1252 = [32]rune{
	0x20AC, 0x0081, 0x201A, 0x0192, 0x201E, 0x2026, 0x2020, 0x2021,
	0x02C6, 0x2030, 0x0160, 0x2039, 0x0152, 0x008D, 0x017D, 0x008F,
	0x0090, 0x2018, 0x2019, 0x201C, 0x201D, 0x2022, 0x2013, 0x2014,
	0x02DC, 0x2122, 0x0161, 0x203A, 0x0153, 0x009D, 0x017E, 0x0178,
}

func desdeWindows1252(b []byte) []byte {
	out := make([]byte, 0, len(b)+len(b)/8)
	for _, c := range b {
		switch {
		case c < 0x80:
			out = append(out, c)
		case c < 0xA0:
			out = utf8.AppendRune(out, cp1252[c-0x80])
		default:
			out = utf8.AppendRune(out, rune(c))
		}
	}
	return out
}

func desdeUTF16(b []byte, bigEndian bool) []byte {
	u := make([]uint16, len(b)/2)
	for i := range u {
		if bigEndian {
			u[i] = uint16(b[2*i])<<8 | uint16(b[2*i+1])
		} else {
			u[i] = uint16(b[2*i+1])<<8 | uint16(b[2*i])
		}
	}
	var out []byte
	for _, r := range utf16.Decode(u) {
		out = utf8.AppendRune(out, r)
	}
	return out
}

func xmlPermitido(r rune) bool {
	return r == '\t' || r == '\n' || r == '\r' ||
		(r >= 0x20 && r <= 0xD7FF) || (r >= 0xE000 && r <= 0xFFFD) || (r >= 0x10000 && r <= 0x10FFFF)
}

func contarInvalidos(b []byte) int {
	n := 0
	for i := 0; i < len(b); {
		c := b[i]
		if c < 0x80 {
			if c < 0x20 && c != '\t' && c != '\n' && c != '\r' {
				n++
			}
			i++
			continue
		}
		r, size := utf8.DecodeRune(b[i:])
		if !xmlPermitido(r) {
			n++
		}
		i += size
	}
	return n
}

func quitarInvalidos(b []byte) []byte {
	out := make([]byte, 0, len(b))
	for i := 0; i < len(b); {
		r, size := utf8.DecodeRune(b[i:])
		if xmlPermitido(r) {
			out = append(out, b[i:i+size]...)
		}
		i += size
	}
	return out
}

// charsetIdentidad: el documento ya es UTF-8 cuando llega al decoder, sea
// cual sea la etiqueta que declare.
func charsetIdentidad(_ string, input io.Reader) (io.Reader, error) { return input, nil }
