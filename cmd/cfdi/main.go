// Comando cfdi lee CFDI 3.3/4.0 y de Retenciones 1.0/2.0 y los escribe como
// JSON. Todos los subcomandos reconocen una Retenciones por su raíz
// (cfdi.Detectar) y la leen, verifican o validan con su API.
//
//	cfdi parse archivo.xml [-sello]
//	cfdi verificar archivo.xml [-certs-sat DIR]
//	cfdi validar archivo.xml
//	cfdi lote RUTA... [-o salida.jsonl] [-resumen resumen.json] [-rutas] [-j N] [-sello | -verificar] [-certs-sat DIR] [-validar]
//
// parse escribe el JSON indentado a stdout y sale con 1 si el archivo no es
// un CFDI. verificar escribe la cadena original y el resultado del sello (y
// del timbre, con los certificados del SAT de -certs-sat) y sale con 0 solo
// si todo lo verificado es válido. validar escribe los hallazgos del esquema
// XSD y de las reglas aritméticas y sale con 1 si el documento es inválido
// (un complemento sin esquema, incompleto, sale con 0). lote recorre archivos .xml, carpetas
// (recursivo) y .zip (también zips dentro de carpetas y de otros zips),
// escribe una línea JSONL por CFDI en el orden en que los encontró y, si se
// pide, un resumen con conteos; -sello incluye los sellos y verifica el del
// emisor, -verificar solo lo verifica, -certs-sat verifica el timbre y
// -validar agrega la validación (y sus conteos por código al resumen). En
// lote, la línea de una Retenciones lleva "tipo_documento":"retenciones" y la
// clave "retenciones" en vez de "cfdi"; las de CFDI no cambian.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

const usoLote = "cfdi lote RUTA... [-o salida.jsonl] [-resumen resumen.json] [-rutas] [-j N] [-sello | -verificar] [-certs-sat DIR] [-validar]"

const uso = `uso:
  cfdi parse archivo.xml [-sello]
  cfdi verificar archivo.xml [-certs-sat DIR]
  cfdi validar archivo.xml
  ` + usoLote + "\n"

func run(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprint(stderr, uso)
		return 2
	}
	switch args[0] {
	case "parse":
		return cmdParse(args[1:], stdout, stderr)
	case "lote":
		return cmdLote(args[1:], stdout, stderr)
	case "verificar":
		return cmdVerificar(args[1:], stdout, stderr)
	case "validar":
		return cmdValidar(args[1:], stdout, stderr)
	case "-h", "-help", "--help", "help":
		fmt.Fprint(stdout, uso)
		return 0
	}
	fmt.Fprintf(stderr, "comando desconocido %q\n%s", args[0], uso)
	return 2
}

// parsearIntercalado deja poner las banderas antes o después de los
// argumentos posicionales (flag se detiene en el primero que no es bandera).
func parsearIntercalado(fs *flag.FlagSet, args []string) ([]string, error) {
	var pos []string
	for {
		if err := fs.Parse(args); err != nil {
			return nil, err
		}
		args = fs.Args()
		if len(args) == 0 {
			return pos, nil
		}
		pos = append(pos, args[0])
		args = args[1:]
	}
}
