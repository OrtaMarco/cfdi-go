package cfdi

import "errors"

// Clases de error fatal. Son estables: el modo lote agrupa por ellas.
const (
	ClaseNoEsXML         = "no_es_xml"         // no hay un elemento raíz legible
	ClaseXMLMalFormado   = "xml_mal_formado"   // empieza como XML pero se rompe (truncado, etiquetas cruzadas)
	ClaseNoEsComprobante = "no_es_comprobante" // la raíz no es Comprobante de cfd/3 o cfd/4
	ClaseNoEsRetenciones = "no_es_retenciones" // ParseRetenciones: la raíz no es Retenciones de retencionpago/1 o /2
	ClaseLectura         = "lectura"           // falló la lectura del io.Reader
)

// Error es un error fatal de Parse: el documento no se pudo leer como CFDI.
type Error struct {
	Clase string
	Err   error
}

func (e *Error) Error() string { return "cfdi: " + e.Clase + ": " + e.Err.Error() }
func (e *Error) Unwrap() error { return e.Err }

// ClaseDe devuelve la clase de un error de Parse, o "" si no es un *Error.
func ClaseDe(err error) string {
	var e *Error
	if errors.As(err, &e) {
		return e.Clase
	}
	return ""
}
