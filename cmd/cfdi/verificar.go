package main

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/OrtaMarco/cfdi-go"
)

// salidaVerificar es el JSON de `cfdi verificar`.
type salidaVerificar struct {
	Archivo string `json:"archivo"`
	*cfdi.Verificacion
}

func cmdVerificar(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("verificar", flag.ContinueOnError)
	fs.SetOutput(stderr)
	dir := fs.String("certs-sat", "", "directorio con los certificados (.cer) del SAT para verificar el timbre")
	pos, err := parsearIntercalado(fs, args)
	if err != nil {
		return 2
	}
	if len(pos) != 1 {
		fmt.Fprint(stderr, "uso: cfdi verificar archivo.xml [-certs-sat DIR]\n")
		return 2
	}
	b, err := os.ReadFile(pos[0])
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	var certs cfdi.CertificadosSAT
	if *dir != "" {
		if certs, err = cargarCertificados(*dir, stderr); err != nil {
			fmt.Fprintln(stderr, err)
			return 1
		}
	}
	verificar := cfdi.Verificar
	if cfdi.Detectar(b) == cfdi.TipoRetenciones {
		verificar = cfdi.VerificarRetenciones
	}
	v := verificar(b, certs)
	e := json.NewEncoder(stdout)
	e.SetEscapeHTML(false)
	e.SetIndent("", "  ")
	if err := e.Encode(salidaVerificar{Archivo: pos[0], Verificacion: v}); err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	if v.Sello.Estado != cfdi.SelloValido || (v.TFD != nil && v.TFD.Estado != cfdi.SelloValido) {
		return 1
	}
	return 0
}

// cargarCertificados lee los .cer (DER o PEM) de dir. Los que no se pueden
// leer se avisan por stderr y se saltan.
func cargarCertificados(dir string, stderr io.Writer) (cfdi.CertificadosSAT, error) {
	entradas, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var certs []*x509.Certificate
	for _, e := range entradas {
		ext := strings.ToLower(filepath.Ext(e.Name()))
		if e.IsDir() || (ext != ".cer" && ext != ".crt" && ext != ".pem") {
			continue
		}
		p := filepath.Join(dir, e.Name())
		b, err := os.ReadFile(p)
		if err != nil {
			fmt.Fprintln(stderr, err)
			continue
		}
		if bloque, _ := pem.Decode(b); bloque != nil {
			if bloque.Type != "CERTIFICATE" {
				continue // una llave u otra cosa en PEM: no es un certificado
			}
			b = bloque.Bytes
		}
		c, err := x509.ParseCertificate(b)
		if err != nil {
			fmt.Fprintf(stderr, "%s: %v\n", p, err)
			continue
		}
		certs = append(certs, c)
	}
	if len(certs) == 0 {
		fmt.Fprintf(stderr, "%s: no hay certificados legibles\n", dir)
	}
	return cfdi.CertificadosEn(certs...), nil
}
