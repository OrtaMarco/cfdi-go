// Package decimal es un número decimal exacto para los montos de un CFDI.
//
// Un CFDI escribe sus montos como texto xs:decimal ("1160.00", "0.160000") y
// sumar o comparar esos textos con float64 pierde centavos. Decimal guarda el
// valor como entero escalado (coeficiente × 10^-escala) y además conserva el
// texto tal como venía en el XML, para poder devolverlo sin tocarlo.
//
// Formato aceptado (léxico de xs:decimal): signo opcional, dígitos y punto
// decimal opcional ("1", "-0.50", "+3.", ".5"). Se rechazan notación
// científica, NaN, Inf, comas, guiones bajos y cadenas vacías.
//
// Espacios: el XSD declara whiteSpace="collapse" para los decimales, así que
// Parse quita los espacios en blanco XML (espacio, tabulador, CR, LF) de los
// extremos y Original devuelve el texto ya recortado. Un espacio en medio del
// número es un error.
//
// El valor cero de Decimal es el número 0 y se puede usar como acumulador.
package decimal

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
)

// Decimal es inmutable: las operaciones devuelven valores nuevos.
type Decimal struct {
	coef  *big.Int // nil equivale a 0
	scale int32    // número de dígitos después del punto
	orig  string   // texto recibido (recortado); "" si el valor se calculó
}

// ErrFormato es la causa de todos los errores de Parse.
var ErrFormato = errors.New("decimal: formato inválido")

// Parse lee un decimal en formato xs:decimal.
func Parse(s string) (Decimal, error) {
	t := strings.Trim(s, " \t\r\n")
	if t == "" {
		return Decimal{}, fmt.Errorf("%w: cadena vacía", ErrFormato)
	}
	body := t
	neg := false
	switch body[0] {
	case '-':
		neg = true
		body = body[1:]
	case '+':
		body = body[1:]
	}
	// "3." y ".5" son válidos en xs:decimal; un segundo punto cae en fracPart y
	// lo rechaza soloDigitos.
	intPart, fracPart, _ := strings.Cut(body, ".")
	if intPart == "" && fracPart == "" {
		return Decimal{}, fmt.Errorf("%w: %q no tiene dígitos", ErrFormato, s)
	}
	if !soloDigitos(intPart) || !soloDigitos(fracPart) {
		return Decimal{}, fmt.Errorf("%w: %q", ErrFormato, s)
	}
	coef, ok := new(big.Int).SetString(intPart+fracPart, 10)
	if !ok {
		// Solo llega aquí con intPart+fracPart vacío, ya descartado arriba.
		return Decimal{}, fmt.Errorf("%w: %q", ErrFormato, s)
	}
	if neg {
		coef.Neg(coef)
	}
	return Decimal{coef: coef, scale: int32(len(fracPart)), orig: t}, nil
}

// MustParse es Parse para literales en pruebas y tablas; entra en pánico si el
// texto no es válido.
func MustParse(s string) Decimal {
	d, err := Parse(s)
	if err != nil {
		panic(err)
	}
	return d
}

// NewFromInt64 construye coef × 10^-scale.
func NewFromInt64(coef int64, scale int32) Decimal {
	if scale < 0 {
		panic("decimal: escala negativa")
	}
	return Decimal{coef: big.NewInt(coef), scale: scale}
}

func soloDigitos(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

func (d Decimal) c() *big.Int {
	if d.coef == nil {
		return new(big.Int)
	}
	return d.coef
}

// Original devuelve el texto tal como vino (sin los espacios de los extremos).
// Para valores calculados con Add/Sub/Mul o el valor cero devuelve String().
func (d Decimal) Original() string {
	if d.orig != "" {
		return d.orig
	}
	return d.String()
}

// Scale es el número de decimales con que se escribió el valor ("1.50" → 2).
func (d Decimal) Scale() int32 { return d.scale }

// String devuelve la forma normalizada: sin ceros a la izquierda, sin ceros
// sobrantes después del punto, sin punto si no hay fracción y sin signo en el
// cero. "0001.500" → "1.5", "-0.00" → "0", "+3." → "3".
func (d Decimal) String() string {
	coef, scale := d.reducido()
	return Decimal{coef: coef, scale: scale}.texto()
}

// StringFixed redondea a places decimales (mitad lejos de cero) y rellena con
// ceros: MustParse("2.5").StringFixed(2) == "2.50".
func (d Decimal) StringFixed(places int32) string {
	if places < 0 {
		places = 0
	}
	r := d.Round(places)
	coef := r.c()
	if r.scale < places {
		coef = new(big.Int).Mul(coef, pow10(places-r.scale))
	}
	return Decimal{coef: coef, scale: places}.texto()
}

// texto escribe coef/scale sin reducir ceros.
func (d Decimal) texto() string {
	coef := d.c()
	neg := coef.Sign() < 0
	digits := new(big.Int).Abs(coef).String()
	if d.scale > 0 {
		if pad := int(d.scale) - len(digits) + 1; pad > 0 {
			digits = strings.Repeat("0", pad) + digits
		}
		cut := len(digits) - int(d.scale)
		digits = digits[:cut] + "." + digits[cut:]
	}
	if neg {
		return "-" + digits
	}
	return digits
}

// Round redondea a places decimales, la mitad lejos de cero (la regla que usa
// el SAT para los importes).
func (d Decimal) Round(places int32) Decimal {
	if places < 0 {
		places = 0
	}
	if d.scale <= places {
		return Decimal{coef: d.c(), scale: d.scale}
	}
	div := pow10(d.scale - places)
	abs := new(big.Int).Abs(d.c())
	q, m := new(big.Int).QuoRem(abs, div, new(big.Int))
	if new(big.Int).Lsh(m, 1).Cmp(div) >= 0 {
		q.Add(q, big.NewInt(1))
	}
	if d.c().Sign() < 0 {
		q.Neg(q)
	}
	return Decimal{coef: q, scale: places}
}

// reducido quita los ceros sobrantes del final de la fracción.
func (d Decimal) reducido() (*big.Int, int32) {
	coef := new(big.Int).Set(d.c())
	scale := d.scale
	if coef.Sign() == 0 {
		return coef, 0
	}
	ten := big.NewInt(10)
	m := new(big.Int)
	for scale > 0 {
		q, r := new(big.Int).QuoRem(coef, ten, m)
		if r.Sign() != 0 {
			break
		}
		coef = q
		scale--
	}
	return coef, scale
}

func pow10(n int32) *big.Int {
	return new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(n)), nil)
}

// alinear lleva a y b a la misma escala.
func alinear(a, b Decimal) (*big.Int, *big.Int, int32) {
	ac, bc := a.c(), b.c()
	switch {
	case a.scale < b.scale:
		return new(big.Int).Mul(ac, pow10(b.scale-a.scale)), bc, b.scale
	case b.scale < a.scale:
		return ac, new(big.Int).Mul(bc, pow10(a.scale-b.scale)), a.scale
	}
	return ac, bc, a.scale
}

// Cmp devuelve -1, 0 o +1 según d sea menor, igual o mayor que o. Compara el
// valor, no el texto: "1.50" y "1.5" son iguales.
func (d Decimal) Cmp(o Decimal) int {
	a, b, _ := alinear(d, o)
	return a.Cmp(b)
}

// Equal es Cmp(o) == 0.
func (d Decimal) Equal(o Decimal) bool { return d.Cmp(o) == 0 }

// Add devuelve d + o con la escala mayor de las dos.
func (d Decimal) Add(o Decimal) Decimal {
	a, b, s := alinear(d, o)
	return Decimal{coef: new(big.Int).Add(a, b), scale: s}
}

// Sub devuelve d - o con la escala mayor de las dos.
func (d Decimal) Sub(o Decimal) Decimal {
	a, b, s := alinear(d, o)
	return Decimal{coef: new(big.Int).Sub(a, b), scale: s}
}

// Mul devuelve d × o exacto (la escala es la suma de las escalas).
func (d Decimal) Mul(o Decimal) Decimal {
	return Decimal{coef: new(big.Int).Mul(d.c(), o.c()), scale: d.scale + o.scale}
}

// Neg devuelve -d.
func (d Decimal) Neg() Decimal {
	return Decimal{coef: new(big.Int).Neg(d.c()), scale: d.scale}
}

// Sign devuelve -1, 0 o +1.
func (d Decimal) Sign() int { return d.c().Sign() }

// IsZero dice si el valor numérico es cero ("0", "0.00", "-0").
func (d Decimal) IsZero() bool { return d.Sign() == 0 }

// MarshalJSON escribe el texto original como cadena JSON: "1160.00" sale
// "1160.00", no 1160.
func (d Decimal) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.Original() + `"`), nil
}

// UnmarshalJSON acepta una cadena JSON con un decimal válido; también un número
// JSON sin exponente. null deja el valor como está.
func (d *Decimal) UnmarshalJSON(b []byte) error {
	s := string(b)
	if s == "null" {
		return nil
	}
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		s = s[1 : len(s)-1]
		if strings.ContainsAny(s, `\"`) {
			return fmt.Errorf("%w: %s", ErrFormato, b)
		}
	}
	v, err := Parse(s)
	if err != nil {
		return err
	}
	*d = v
	return nil
}
