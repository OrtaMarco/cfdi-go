package cfdi

import "strings"

// Hallazgo es una violación de una norma encontrada al validar: del esquema
// XSD del SAT (ValidarEsquema) o de una regla aritmética o de coherencia
// (ValidarReglas).
//
// Es un tipo propio y no Advertencia porque dicen cosas distintas: una
// Advertencia cuenta qué tuvo que hacer la lectura tolerante para leer el
// documento (se lee igual), un Hallazgo dice qué norma incumple (el SAT lo
// rechazaría). Un CFDI puede leerse sin una sola advertencia y tener
// hallazgos, y al revés. Lo que sí comparten es la forma de señalar el
// lugar: Ruta y Atributo siguen exactamente el contrato de Advertencia
// (nombres locales, índices [n] en los nodos que el esquema deja repetir,
// "/@Atributo" al final), así que un mismo nodo se busca igual en las dos
// listas.
type Hallazgo struct {
	// Codigo es estable entre versiones (sirve para agrupar en lote): uno
	// de los Esquema* o Regla*.
	Codigo string `json:"codigo"`
	// Nivel es NivelError (el documento incumple), NivelAdvertencia
	// (incumple una regla sin fuente citable que la haga obligatoria: puede
	// estar mal, no se afirma) o NivelAviso (una parte del documento no se
	// pudo validar; no dice que esté mal).
	Nivel    string `json:"nivel"`
	Ruta     string `json:"ruta,omitempty"`
	Atributo string `json:"atributo,omitempty"`
	// Valor es el texto original del atributo en los hallazgos de valores,
	// o el del elemento si es de contenido simple (Detallista; Atributo
	// vacío).
	Valor string `json:"valor,omitempty"`
	// Regla es la norma concreta: en el esquema, la faceta o restricción
	// XSD ("pattern", "enumeration", "minOccurs", "orden"…); en las reglas,
	// la referencia a la fuente (ver ValidarReglas).
	Regla string `json:"regla"`
	// Mensaje es para personas y puede cambiar.
	Mensaje string `json:"mensaje,omitempty"`
}

// Niveles de un Hallazgo.
const (
	NivelError       = "error"
	NivelAdvertencia = "advertencia"
	NivelAviso       = "aviso"
)

// Nodo es la ruta del nodo del hallazgo: Ruta sin el "/@Atributo" final.
func (h Hallazgo) Nodo() string {
	if h.Atributo == "" {
		return h.Ruta
	}
	return strings.TrimSuffix(h.Ruta, "/@"+h.Atributo)
}

// Estados de ResultadoValidacion.
const (
	ValidacionValida     = "valido"     // sin errores ni avisos (puede traer advertencias)
	ValidacionInvalida   = "invalido"   // al menos un hallazgo de nivel error
	ValidacionIncompleta = "incompleto" // sin errores, pero con partes que no se pudieron validar (avisos)
)

// ResultadoValidacion es el resultado de ValidarEsquema.
type ResultadoValidacion struct {
	Estado    string     `json:"estado"`
	Hallazgos []Hallazgo `json:"hallazgos"`
}

// EstadoDe resume una lista de hallazgos en uno de los Validacion*. Solo
// los errores hacen inválido un documento; las advertencias no cambian el
// estado.
func EstadoDe(hs []Hallazgo) string {
	estado := ValidacionValida
	for _, h := range hs {
		switch h.Nivel {
		case NivelError:
			return ValidacionInvalida
		case NivelAviso:
			estado = ValidacionIncompleta
		}
	}
	return estado
}

// TieneHallazgo dice si hay al menos un hallazgo con ese código.
func TieneHallazgo(hs []Hallazgo, codigo string) bool {
	for _, h := range hs {
		if h.Codigo == codigo {
			return true
		}
	}
	return false
}
