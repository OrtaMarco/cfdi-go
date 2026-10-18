package cfdi

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"testing"
)

var update = flag.Bool("update", false, "regenera los .golden.json de testdata")

// fixtures devuelve todos los .xml bajo testdata/, a cualquier profundidad.
// Agregar un directorio de fixtures (p. ej. testdata/publicos/) basta para
// que entre en las pruebas; sus goldens se crean con -update.
//
// Convención: un .xml dentro de un directorio llamado "fatales" debe fallar
// con error fatal; cualquier otro debe leerse.
func fixtures(t *testing.T) []string {
	t.Helper()
	var out []string
	err := filepath.WalkDir("testdata", func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.EqualFold(filepath.Ext(p), ".xml") {
			out = append(out, p)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(out) == 0 {
		t.Fatal("no hay fixtures en testdata/")
	}
	return out
}

func esFatal(p string) bool {
	return slices.Contains(strings.Split(filepath.ToSlash(filepath.Dir(p)), "/"), "fatales")
}

// esRetenciones dice si un fixture que no es fatal es un CFDI de Retenciones:
// se lee con ParseRetenciones y no entra a las pruebas del Comprobante. Los
// de fatales/ se leen siempre con Parse (raiz-retenciones.xml comprueba que
// Parse los sigue rechazando).
func esRetenciones(t *testing.T, p string) bool {
	t.Helper()
	b, err := os.ReadFile(p)
	if err != nil {
		t.Fatal(err)
	}
	return !esFatal(p) && Detectar(b) == TipoRetenciones
}

// leerFixture lee un fixture con Parse o, si es Retenciones,
// ParseRetenciones, y devuelve su JSON con las opciones dadas.
func leerFixtureJSON(t *testing.T, p string, o OpcionesJSON) ([]byte, error) {
	t.Helper()
	b, err := os.ReadFile(p)
	if err != nil {
		t.Fatal(err)
	}
	if esRetenciones(t, p) {
		rt, err := ParseRetencionesBytes(b)
		if err != nil {
			return nil, err
		}
		return rt.JSONOpciones(o)
	}
	c, err := ParseBytes(b)
	if err != nil {
		return nil, err
	}
	return c.JSONOpciones(o)
}

// salidaGolden es lo que se compara: el Documento indentado, o la clase del
// error fatal (el mensaje no: incluye textos de encoding/xml que cambian entre
// versiones de Go).
func salidaGolden(t *testing.T, p string) []byte {
	t.Helper()
	out, err := leerFixtureJSON(t, p, OpcionesJSON{Indentar: true})
	if err != nil {
		if ClaseDe(err) == "" {
			t.Fatal(err)
		}
		out, _ = json.MarshalIndent(map[string]string{"error_clase": ClaseDe(err)}, "", "  ")
	}
	return append(out, '\n')
}

func TestGolden(t *testing.T) {
	for _, p := range fixtures(t) {
		t.Run(filepath.ToSlash(p), func(t *testing.T) {
			_, perr := leerFixtureJSON(t, p, OpcionesJSON{})
			if esFatal(p) && perr == nil {
				t.Errorf("está en fatales/ pero se leyó sin error")
			}
			if !esFatal(p) && perr != nil {
				t.Errorf("error fatal inesperado: %v", perr)
			}

			got := salidaGolden(t, p)
			golden := strings.TrimSuffix(p, filepath.Ext(p)) + ".golden.json"
			if *update {
				if err := os.WriteFile(golden, got, 0o644); err != nil {
					t.Fatal(err)
				}
				return
			}
			want, err := os.ReadFile(golden)
			if err != nil {
				t.Fatalf("falta %s: corre `go test -run TestGolden -update .`", golden)
			}
			if !bytes.Equal(got, want) {
				t.Errorf("la salida difiere de %s (revisa y corre con -update si el cambio es correcto):\n%s", golden, primeraDiferencia(want, got))
			}
		})
	}
}

func primeraDiferencia(want, got []byte) string {
	wl := strings.Split(string(want), "\n")
	gl := strings.Split(string(got), "\n")
	for i := 0; i < len(wl) || i < len(gl); i++ {
		var w, g string
		if i < len(wl) {
			w = wl[i]
		}
		if i < len(gl) {
			g = gl[i]
		}
		if w != g {
			return "línea " + strconv.Itoa(i+1) + ":\n  quiero: " + w + "\n  obtuve: " + g
		}
	}
	return "(sin diferencia por líneas)"
}
