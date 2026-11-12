package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/OrtaMarco/cfdi-go"
)

// salidaValidar es el JSON de `cfdi validar`.
type salidaValidar struct {
	Archivo string `json:"archivo"`
	// Error dice por qué no se evaluaron las reglas: el documento no se
	// pudo leer como CFDI (clase y mensaje de Parse).
	Error *errorLectura `json:"error,omitempty"`
	*cfdi.Validacion
}

type errorLectura struct {
	Clase   string `json:"clase"`
	Mensaje string `json:"mensaje"`
}

func cmdValidar(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("validar", flag.ContinueOnError)
	fs.SetOutput(stderr)
	pos, err := parsearIntercalado(fs, args)
	if err != nil {
		return 2
	}
	if len(pos) != 1 {
		fmt.Fprint(stderr, "uso: cfdi validar archivo.xml\n")
		return 2
	}
	b, err := os.ReadFile(pos[0])
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	s := salidaValidar{Archivo: pos[0]}
	if cfdi.Detectar(b) == cfdi.TipoRetenciones {
		// Esquema y reglas de Retenciones.
		s.Validacion = cfdi.ValidarRetenciones(b)
	} else {
		c, err := cfdi.ParseBytes(b)
		if err != nil {
			s.Error = &errorLectura{Clase: cfdi.ClaseDe(err), Mensaje: err.Error()}
		}
		s.Validacion = cfdi.Validar(b, c)
	}
	e := json.NewEncoder(stdout)
	e.SetEscapeHTML(false)
	e.SetIndent("", "  ")
	if err := e.Encode(s); err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	if s.Estado == cfdi.ValidacionInvalida {
		return 1
	}
	return 0
}
