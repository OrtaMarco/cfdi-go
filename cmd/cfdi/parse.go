package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/OrtaMarco/cfdi-go"
)

func cmdParse(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("parse", flag.ContinueOnError)
	fs.SetOutput(stderr)
	sello := fs.Bool("sello", false, "incluir sellos, certificados y cadenas")
	pos, err := parsearIntercalado(fs, args)
	if err != nil {
		return 2
	}
	if len(pos) != 1 {
		fmt.Fprint(stderr, "uso: cfdi parse archivo.xml [-sello]\n")
		return 2
	}
	datos, err := os.ReadFile(pos[0])
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	o := cfdi.OpcionesJSON{IncluirSello: *sello, Indentar: true}
	var b []byte
	if cfdi.Detectar(datos) == cfdi.TipoRetenciones {
		var rt *cfdi.Retenciones
		if rt, err = cfdi.ParseRetencionesBytes(datos); err == nil {
			b, err = rt.JSONOpciones(o)
		}
	} else {
		var c *cfdi.Comprobante
		if c, err = cfdi.ParseBytes(datos); err == nil {
			b, err = c.JSONOpciones(o)
		}
	}
	if cfdi.ClaseDe(err) != "" {
		fmt.Fprintf(stderr, "%s: %v\n", pos[0], err)
		return 1
	}
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	if _, err := stdout.Write(append(b, '\n')); err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	return 0
}
