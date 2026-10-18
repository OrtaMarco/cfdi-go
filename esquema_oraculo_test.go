package cfdi

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

// goldenEsquema es lo que escribe scripts/validar_xsd.py (libxml2 con los
// XSD oficiales) para cada fixture.
type goldenEsquema struct {
	Veredicto string `json:"veredicto"`
	Errores   []struct {
		Linea   int    `json:"linea"`
		Ruta    string `json:"ruta"`
		Mensaje string `json:"mensaje"`
	} `json:"errores"`
}

// fixturesEsquema devuelve los .xml de testdata que tienen golden de esquema.
func fixturesEsquema(t *testing.T) []string {
	t.Helper()
	var out []string
	err := filepath.WalkDir("testdata", func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(p, ".xml") {
			out = append(out, p)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	sort.Strings(out)
	return out
}

// El veredicto (valido, invalido, incompleto) de ValidarEsquema coincide
// con el de libxml2 en cada fixture: públicos, sintéticos y mutantes.
func TestValidarEsquemaContraOraculo(t *testing.T) {
	xmls := fixturesEsquema(t)
	coinciden, total := 0, 0
	porVeredicto := map[string]int{}
	for _, p := range xmls {
		gb, err := os.ReadFile(strings.TrimSuffix(p, ".xml") + ".esquema.json")
		if err != nil {
			t.Errorf("%s: sin golden de esquema (correr scripts/validar_xsd.py): %v", p, err)
			continue
		}
		var g goldenEsquema
		if err := json.Unmarshal(gb, &g); err != nil {
			t.Fatalf("%s: %v", p, err)
		}
		b, err := os.ReadFile(p)
		if err != nil {
			t.Fatal(err)
		}
		r := ValidarEsquema(b)
		total++
		porVeredicto[g.Veredicto]++
		if r.Estado == g.Veredicto {
			coinciden++
			continue
		}
		var lib, gohs []string
		for _, e := range g.Errores {
			lib = append(lib, e.Ruta+": "+e.Mensaje)
		}
		for _, h := range r.Hallazgos {
			gohs = append(gohs, h.Codigo+" "+h.Ruta+" "+h.Regla+" "+h.Mensaje)
		}
		t.Errorf("%s: Go dice %s y libxml2 %s\n  libxml2:\n    %s\n  Go:\n    %s", p, r.Estado, g.Veredicto,
			strings.Join(lib, "\n    "), strings.Join(gohs, "\n    "))
	}
	t.Logf("coincide el veredicto en %d de %d documentos (libxml2: %v)", coinciden, total, porVeredicto)
	if total < 300 {
		t.Errorf("solo %d fixtures con golden de esquema", total)
	}
}
