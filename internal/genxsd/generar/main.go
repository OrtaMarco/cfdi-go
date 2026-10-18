// Comando generar escribe esquema_reglas.go y esquema_catalogos.go a partir
// de los XSD del SAT. Se corre desde la raíz del repo:
//
//	go run ./internal/genxsd/generar
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/OrtaMarco/cfdi-go/internal/genxsd"
)

func main() {
	dir := flag.String("xsd", "internal/genxsd/xsd/sitio_internet", "directorio sitio_internet con los XSD del SAT")
	reglas := flag.String("o", "esquema_reglas.go", "archivo Go con elementos, atributos y tipos")
	catalogos := flag.String("catalogos", "esquema_catalogos.go", "archivo Go con los catálogos")
	flag.Parse()
	r, c, err := genxsd.Generar(*dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for _, s := range []struct {
		ruta string
		b    []byte
	}{{*reglas, r}, {*catalogos, c}} {
		if err := os.WriteFile(s.ruta, s.b, 0o644); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
