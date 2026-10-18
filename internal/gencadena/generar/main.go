// Comando generar escribe cadena_reglas.go a partir de los XSLT del SAT.
// Se corre desde la raíz del repo:
//
//	go run ./internal/gencadena/generar
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/OrtaMarco/cfdi-go/internal/gencadena"
)

func main() {
	dir := flag.String("xslt", "internal/gencadena/xslt/sitio_internet", "directorio sitio_internet con las hojas del SAT")
	salida := flag.String("o", "cadena_reglas.go", "archivo Go que se escribe")
	flag.Parse()
	b, err := gencadena.Generar(*dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := os.WriteFile(*salida, b, 0o644); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
