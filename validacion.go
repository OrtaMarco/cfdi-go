package cfdi

// Validacion junta las dos validaciones de un documento: la del esquema
// (sobre los bytes) y la de las reglas (sobre el Comprobante). Es lo que
// escriben `cfdi validar` y `cfdi lote -validar`.
type Validacion struct {
	// Estado es el peor de los dos: invalido si alguno lo es; incompleto
	// si alguno lo es y ninguno es inválido; si no, valido.
	Estado  string               `json:"estado"`
	Esquema *ResultadoValidacion `json:"esquema"`
	// Reglas es nil si el documento no se pudo leer como CFDI (Parse
	// falló).
	Reglas *ResultadoValidacion `json:"reglas,omitempty"`
}

// Validar corre ValidarEsquema sobre xml y ValidarReglas sobre c. Si c es
// nil lee xml con ParseBytes; si tampoco se puede leer, Reglas queda nil y
// el estado es el del esquema (que entonces es inválido). Nunca falla.
func Validar(xml []byte, c *Comprobante) *Validacion {
	v := &Validacion{Esquema: ValidarEsquema(xml)}
	if c == nil {
		c, _ = ParseBytes(xml)
	}
	if c != nil {
		hs := ValidarReglas(c)
		v.Reglas = &ResultadoValidacion{Estado: EstadoDe(hs), Hallazgos: hs}
	}
	v.Estado = v.Esquema.Estado
	if v.Reglas != nil {
		v.Estado = peorEstado(v.Estado, v.Reglas.Estado)
	}
	return v
}

func peorEstado(a, b string) string {
	rango := map[string]int{ValidacionValida: 0, ValidacionIncompleta: 1, ValidacionInvalida: 2}
	if rango[b] > rango[a] {
		return b
	}
	return a
}

// Hallazgos devuelve los hallazgos del esquema y luego los de las reglas.
func (v *Validacion) Hallazgos() []Hallazgo {
	out := append([]Hallazgo(nil), v.Esquema.Hallazgos...)
	if v.Reglas != nil {
		out = append(out, v.Reglas.Hallazgos...)
	}
	return out
}
