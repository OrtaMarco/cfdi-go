package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// lote corre `cfdi lote` sobre rutas y devuelve las líneas crudas y el
// resumen.
func correrLote(t *testing.T, dir string, rutas []string, extra ...string) ([][]byte, Resumen, string) {
	t.Helper()
	jsonl, resumen := filepath.Join(dir, "s.jsonl"), filepath.Join(dir, "r.json")
	var out, errb bytes.Buffer
	args := append(append([]string{"lote", "-o", jsonl, "-resumen", resumen, "-j", "2"}, extra...), rutas...)
	if code := run(args, &out, &errb); code != 0 {
		t.Fatalf("código %d: %s", code, errb.String())
	}
	b, _ := os.ReadFile(jsonl)
	lineas := bytes.SplitAfter(b, []byte("\n"))
	lineas = lineas[:len(lineas)-1]
	rb, _ := os.ReadFile(resumen)
	var r Resumen
	if err := json.Unmarshal(rb, &r); err != nil {
		t.Fatal(err)
	}
	return lineas, r, errb.String()
}

func TestVerificarRetenciones(t *testing.T) {
	var out, errb bytes.Buffer
	code := run([]string{"verificar", filepath.Join(sinteticos, "sellados", "ret20-plataformas.xml"), "-certs-sat", certs}, &out, &errb)
	var m map[string]any
	if err := json.Unmarshal(out.Bytes(), &m); err != nil {
		t.Fatalf("%v: %s", err, out.String())
	}
	s, _ := m["sello"].(map[string]any)
	tfd, _ := m["tfd"].(map[string]any)
	c, _ := m["certificado"].(map[string]any)
	if code != 0 || s["estado"] != "valido" || tfd["estado"] != "valido" || c == nil || !strings.HasPrefix(m["cadena_original"].(string), "||2.0|") {
		t.Errorf("código %d: %s %s", code, out.String(), errb.String())
	}
	// El CFDI también trae los chequeos del certificado.
	out.Reset()
	run([]string{"verificar", filepath.Join(sinteticos, "sellados", "ingreso40-pue.xml")}, &out, &errb)
	if !strings.Contains(out.String(), `"vigencia": {`) {
		t.Errorf("verificar de un CFDI sin chequeos del certificado: %s", out.String())
	}
}

func TestLoteRetenciones(t *testing.T) {
	dir := t.TempDir()
	cfdi := filepath.Join(sinteticos, "ingreso40-pue.xml")
	ret20 := filepath.Join(sinteticos, "retenciones", "ret20-plataformas.xml")
	ret10 := filepath.Join(sinteticos, "retenciones", "ret10-extranjero.xml")
	trunca := filepath.Join(dir, "trunca.xml")
	ret := leerFixture(t, "retenciones/ret20-plataformas.xml")
	os.WriteFile(trunca, ret[:len(ret)-60], 0o644)

	solo, _, _ := correrLote(t, dir, []string{cfdi})
	lineas, r, errs := correrLote(t, dir, []string{ret20, cfdi, ret10, trunca})
	if len(lineas) != 4 {
		t.Fatalf("%d líneas", len(lineas))
	}
	// La línea del CFDI es la misma con o sin Retenciones en el lote.
	if !bytes.Equal(lineas[1], solo[0]) {
		t.Errorf("la línea del CFDI cambió:\n%s\n%s", lineas[1], solo[0])
	}
	if bytes.Contains(lineas[1], []byte("tipo_documento")) {
		t.Error("la línea del CFDI no lleva tipo_documento")
	}
	for i, quiero := range map[int]string{0: "2.0", 2: "1.0"} {
		var m map[string]any
		json.Unmarshal(lineas[i], &m)
		rt, _ := m["retenciones"].(map[string]any)
		if m["tipo_documento"] != "retenciones" || m["schema_version"] != "1" || rt["version"] != quiero || m["cfdi"] != nil {
			t.Errorf("línea %d: %s", i+1, lineas[i])
		}
	}
	if !strings.Contains(string(lineas[3]), `"clase":"xml_mal_formado"`) {
		t.Errorf("Retenciones truncada: %s", lineas[3])
	}

	if r.Total != 4 || r.Leidos != 3 || r.Fatales != 1 {
		t.Errorf("totales: %+v", r)
	}
	// Las versiones y complementos de Retenciones van aparte.
	if r.PorVersion["4.0"] != 1 || len(r.PorVersion) != 1 || r.PorTipo["I"] != 1 || r.PorComplemento["TimbreFiscalDigital 1.1"] != 1 {
		t.Errorf("CFDI: %v %v %v", r.PorVersion, r.PorTipo, r.PorComplemento)
	}
	rr := r.Retenciones
	if rr == nil || rr.Leidos != 2 || rr.PorVersion["2.0"] != 1 || rr.PorVersion["1.0"] != 1 ||
		rr.PorComplemento["ServiciosPlataformasTecnologicas 1.0"] != 1 || rr.PorComplemento["Pagosaextranjeros 1.0"] != 1 ||
		rr.PorComplemento["TimbreFiscalDigital 1.1"] != 1 || rr.PorComplemento["TimbreFiscalDigital 1.0"] != 1 {
		t.Errorf("Retenciones: %+v", rr)
	}
	if !strings.Contains(errs, "de los leídos, 2 son CFDI de Retenciones") {
		t.Errorf("stderr: %q", errs)
	}

	// Sin Retenciones, el resumen no trae la clave.
	_, r, errs = correrLote(t, dir, []string{cfdi})
	if r.Retenciones != nil || strings.Contains(errs, "Retenciones") {
		t.Errorf("sin Retenciones: %+v %q", r.Retenciones, errs)
	}
	b, _ := os.ReadFile(filepath.Join(dir, "r.json"))
	if bytes.Contains(b, []byte(`"retenciones"`)) {
		t.Errorf("resumen sin Retenciones con la clave: %s", b)
	}

	// -verificar y -validar también las leen.
	lineas, r, _ = correrLote(t, dir, []string{ret20, ret10}, "-verificar", "-validar")
	for i, l := range lineas {
		var m map[string]any
		json.Unmarshal(l, &m)
		v, _ := m["verificacion_sello"].(map[string]any)
		val, _ := m["validacion"].(map[string]any)
		if v == nil || val == nil || val["reglas"] != nil {
			t.Errorf("línea %d: %s", i+1, l)
		}
	}
	// La 1.0 trae TFD 1.0, que se valida con su propio XSD.
	if r.Sellos == nil || r.Validacion == nil || r.Validacion.PorEstado["valido"] != 2 || r.Validacion.PorEstado["invalido"] != 0 {
		t.Errorf("resumen: %+v %+v", r.Sellos, r.Validacion)
	}
}
