package cfdi

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Fecha guarda una fecha del CFDI como texto original y como hora civil.
//
// El SAT escribe las fechas sin zona horaria ("2026-01-15T10:30:00"): es la
// hora local del lugar de expedición, que puede ser cualquiera de los husos de
// México. Esta librería no inventa la zona: Time lleva la fecha y hora tal
// cual (hora civil) con time.UTC como contenedor, sin que eso signifique que
// el instante sea UTC. Para un instante real use En con la zona que
// corresponda.
//
// Los nodos de nómina usan fechas sin hora ("2026-01-15"); entonces la hora
// es 00:00:00.
type Fecha struct {
	Original string    // texto como venía
	Time     time.Time // hora civil; cero si el texto no se pudo leer
	// FueraDeFormato dice que el texto es una fecha legible pero no en el
	// formato del SAT (con zona, con fracción de segundo, con espacios, sin
	// hora o al revés): es la versión por campo de FECHA_FORMATO. Con
	// Opciones.FechaEstricta esas fechas no se interpretan (Time queda en
	// cero) y FueraDeFormato sigue en true. No va al JSON.
	FueraDeFormato bool
}

// Valida dice si el texto se pudo interpretar.
func (f Fecha) Valida() bool { return !f.Time.IsZero() }

// En interpreta la hora civil en la zona loc (p. ej. America/Mexico_City).
func (f Fecha) En(loc *time.Location) time.Time {
	t := f.Time
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), loc)
}

func (f Fecha) String() string { return f.Original }

// MarshalJSON escribe el texto original.
func (f Fecha) MarshalJSON() ([]byte, error) { return json.Marshal(f.Original) }

// UnmarshalJSON lee el texto y vuelve a interpretar la hora civil; si no se
// puede, deja Time en cero sin error (igual que Parse). Como el JSON no dice
// si el campo era fecha u hora, FueraDeFormato acepta los dos formatos SAT.
func (f *Fecha) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	*f, _, _ = parseFecha(s, false)
	if _, err := time.Parse(layoutFecha, s); err == nil {
		f.FueraDeFormato = false
	}
	return nil
}

const (
	layoutFechaHora = "2006-01-02T15:04:05"
	layoutFecha     = "2006-01-02"
)

// parseFecha lee una fecha SAT. soloFecha indica que el atributo es xs:date.
// Devuelve formatoRaro=true si se pudo leer pero no en el formato que pide el
// SAT (con zona, con fracción de segundo, fecha sin hora o al revés); y err si
// no se pudo leer.
func parseFecha(s string, soloFecha bool) (f Fecha, formatoRaro bool, err error) {
	defer func() { f.FueraDeFormato = formatoRaro }()
	f.Original = s
	t := strings.TrimSpace(s)
	if t != s {
		formatoRaro = true
	}
	esperado, otro := layoutFechaHora, layoutFecha
	if soloFecha {
		esperado, otro = layoutFecha, layoutFechaHora
	}
	if v, e := time.Parse(esperado, t); e == nil {
		// time.Parse acepta fracción de segundo aunque el layout no la tenga.
		f.Time = v
		return f, formatoRaro || strings.Contains(t, "."), nil
	}
	if v, e := time.Parse(otro, t); e == nil {
		f.Time = v
		return f, true, nil
	}
	// Fracción de segundo y/o zona: se conserva la hora de pared tal cual.
	for _, l := range []string{"2006-01-02T15:04:05.999999999", time.RFC3339Nano, "2006-01-02T15:04:05.999999999-0700"} {
		if v, e := time.Parse(l, t); e == nil {
			f.Time = time.Date(v.Year(), v.Month(), v.Day(), v.Hour(), v.Minute(), v.Second(), v.Nanosecond(), time.UTC)
			return f, true, nil
		}
	}
	return Fecha{Original: s}, false, fmt.Errorf("fecha ilegible %q", s)
}
