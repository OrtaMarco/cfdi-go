package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const mutantes = "../../testdata/mutantes"

func validarJSON(t *testing.T, args ...string) (int, map[string]any, string) {
	t.Helper()
	var out, errb bytes.Buffer
	code := run(append([]string{"validar"}, args...), &out, &errb)
	var m map[string]any
	if out.Len() > 0 {
		if err := json.Unmarshal(out.Bytes(), &m); err != nil {
			t.Fatalf("salida no es JSON: %v\n%s", err, out.String())
		}
	}
	return code, m, errb.String()
}

func codigosHallazgos(m map[string]any, parte string) []string {
	r, _ := m[parte].(map[string]any)
	hs, _ := r["hallazgos"].([]any)
	var out []string
	for _, h := range hs {
		out = append(out, h.(map[string]any)["codigo"].(string))
	}
	return out
}

func TestValidar(t *testing.T) {
	escribir := func(nombre, contenido string) string {
		p := filepath.Join(t.TempDir(), nombre)
		os.WriteFile(p, []byte(contenido), 0o644)
		return p
	}
	base := string(leerFixture(t, "ingreso40-pue.xml"))
	casos := []struct {
		nombre, archivo string
		codigo          int
		estado          string
		esquema, reglas []string // códigos esperados; nil = ninguno
		sinReglas       bool
	}{
		{"válido", filepath.Join(sinteticos, "ingreso40-pue.xml"), 0, "valido", nil, nil, false},
		{"esquema inválido", filepath.Join(mutantes, "catalogo-moneda.xml"), 1, "invalido", []string{"ESQUEMA_CATALOGO"}, []string{"REGLA_TIPOCAMBIO_REQUERIDO"}, false},
		{"complemento sin esquema", filepath.Join(mutantes, "complemento-no-soportado.xml"), 0, "incompleto", []string{"ESQUEMA_NO_SOPORTADO"}, nil, false},
		{"regla rota", escribir("total.xml", strings.Replace(base, ` Total="1640.00"`, ` Total="1640.01"`, 1)), 1, "invalido", nil, []string{"REGLA_TOTAL"}, false},
		{"no es comprobante", escribir("otro.xml", `<x:Otro xmlns:x="urn:otro"/>`), 0, "incompleto", []string{"ESQUEMA_NO_SOPORTADO"}, nil, true},
		{"no es XML", escribir("x.xml", "hola"), 1, "invalido", []string{"ESQUEMA_XML_ILEGIBLE"}, nil, true},
		{"CFDI 3.2", filepath.Join(sinteticos, "sellados", "domicilios32.xml"), 0, "incompleto", []string{"ESQUEMA_NO_SOPORTADO"}, []string{"REGLA_VERSION_NO_SOPORTADA"}, false},
	}
	for _, c := range casos {
		code, m, errs := validarJSON(t, c.archivo)
		if code != c.codigo || m["estado"] != c.estado || m["archivo"] != c.archivo || errs != "" {
			t.Errorf("%s: código %d, estado %v, stderr %q", c.nombre, code, m["estado"], errs)
		}
		if got := codigosHallazgos(m, "esquema"); strings.Join(got, ",") != strings.Join(c.esquema, ",") {
			t.Errorf("%s: esquema %v, quiero %v", c.nombre, got, c.esquema)
		}
		_, hayReglas := m["reglas"]
		_, hayError := m["error"]
		if hayReglas == c.sinReglas || hayError != c.sinReglas {
			t.Errorf("%s: reglas %v, error %v", c.nombre, m["reglas"], m["error"])
		}
		if got := codigosHallazgos(m, "reglas"); strings.Join(got, ",") != strings.Join(c.reglas, ",") {
			t.Errorf("%s: reglas %v, quiero %v", c.nombre, got, c.reglas)
		}
	}
	// Retenciones: solo el esquema, sin reglas ni error de lectura.
	for archivo, quiero := range map[string]string{"ret20-plataformas.xml": "valido", "ret20-anomalias.xml": "invalido"} {
		code, m, errs := validarJSON(t, filepath.Join(sinteticos, "retenciones", archivo))
		_, hayReglas := m["reglas"]
		_, hayError := m["error"]
		if m["estado"] != quiero || (code == 1) != (quiero == "invalido") || hayReglas || hayError || errs != "" {
			t.Errorf("Retenciones %s: código %d, %v", archivo, code, m)
		}
	}
	if code, _, _ := validarJSON(t, "/no/existe.xml"); code != 1 {
		t.Errorf("archivo inexistente: código %d", code)
	}
	if code, _, _ := validarJSON(t); code != 2 {
		t.Errorf("sin archivo: código %d", code)
	}
	if code, _, _ := validarJSON(t, "-x"); code != 2 {
		t.Errorf("bandera desconocida: código %d", code)
	}
	var out bytes.Buffer
	if run([]string{"help"}, &out, &out); !strings.Contains(out.String(), "cfdi validar archivo.xml") || !strings.Contains(out.String(), "-validar") {
		t.Errorf("la ayuda no menciona validar: %s", out.String())
	}
}

func TestLoteValidarYTimbre(t *testing.T) {
	dir := t.TempDir()
	raiz := filepath.Join(dir, "entrada")
	copiar(t, "sellados/ingreso40-pue.xml", filepath.Join(raiz, "1.xml"))
	copiar(t, "../mutantes/catalogo-moneda.xml", filepath.Join(raiz, "2.xml"))
	copiar(t, "../mutantes/complemento-no-soportado.xml", filepath.Join(raiz, "3.xml"))
	copiar(t, "fatales/no-es-xml.xml", filepath.Join(raiz, "4.xml"))
	copiar(t, "../publicos/eclipxe13-CfdiUtils/tests__assets__cfdi33-valid.xml", filepath.Join(raiz, "5.xml")) // timbre que no corresponde

	lote := func(extra ...string) ([]map[string]any, Resumen, string) {
		t.Helper()
		jsonl, resumen := filepath.Join(dir, "s.jsonl"), filepath.Join(dir, "r.json")
		var out, errb bytes.Buffer
		if code := run(append([]string{"lote", raiz, "-o", jsonl, "-resumen", resumen, "-j", "2"}, extra...), &out, &errb); code != 0 {
			t.Fatalf("código %d: %s", code, errb.String())
		}
		b, _ := os.ReadFile(jsonl)
		var lineas []map[string]any
		for _, l := range bytes.Split(bytes.TrimSpace(b), []byte("\n")) {
			var m map[string]any
			if err := json.Unmarshal(l, &m); err != nil {
				t.Fatal(err)
			}
			lineas = append(lineas, m)
		}
		rb, _ := os.ReadFile(resumen)
		var r Resumen
		if err := json.Unmarshal(rb, &r); err != nil {
			t.Fatal(err)
		}
		return lineas, r, errb.String()
	}

	lineas, r, errs := lote("-validar")
	quiero := []string{"valido", "invalido", "incompleto", "", "valido"}
	for i, l := range lineas {
		v, _ := l["validacion"].(map[string]any)
		if got, _ := v["estado"].(string); got != quiero[i] {
			t.Errorf("línea %d: validación %q, quiero %q", i+1, got, quiero[i])
		}
	}
	if _, ok := lineas[0]["verificacion_tfd"]; ok {
		t.Error("verificacion_tfd sin -certs-sat")
	}
	v := r.Validacion
	if v == nil || v.PorEstado["valido"] != 2 || v.PorEstado["invalido"] != 1 || v.PorEstado["incompleto"] != 1 {
		t.Fatalf("resumen de validación: %+v", v)
	}
	c := v.HallazgosPorCodigo["ESQUEMA_CATALOGO"]
	if c == nil || c.CFDI != 1 || c.Ocurrencias != 1 || len(c.Ejemplos) != 1 || strings.Contains(c.Ejemplos[0], "/") {
		t.Errorf("conteo ESQUEMA_CATALOGO: %+v", c)
	}
	if c := v.HallazgosPorCodigo["ESQUEMA_NO_SOPORTADO"]; c == nil || c.CFDI != 1 {
		t.Errorf("conteo ESQUEMA_NO_SOPORTADO: %+v", c)
	}
	if !strings.Contains(errs, "validación: 2 válidos, 1 inválidos, 1 incompletos") || r.SellosTFD != nil || r.Sellos != nil {
		t.Errorf("stderr %q, sellos %+v %+v", errs, r.SellosTFD, r.Sellos)
	}

	// -certs-sat verifica el timbre sin el sello del emisor, salvo que se
	// pida.
	lineas, r, errs = lote("-certs-sat", certs)
	quieroTFD := []string{"valido", "no_verificable CERTIFICADO_SAT_NO_DISPONIBLE", "no_verificable CERTIFICADO_SAT_NO_DISPONIBLE", "", "invalido SELLO_NO_COINCIDE"}
	for i, l := range lineas {
		if got := estado(l, "verificacion_tfd"); got != quieroTFD[i] {
			t.Errorf("línea %d: timbre %q, quiero %q", i+1, got, quieroTFD[i])
		}
		if _, ok := l["verificacion_sello"]; ok {
			t.Errorf("línea %d: verificacion_sello sin -verificar", i+1)
		}
		if _, ok := l["validacion"]; ok {
			t.Errorf("línea %d: validacion sin -validar", i+1)
		}
	}
	if s := r.SellosTFD; s == nil || s.Valido != 1 || s.Invalido["SELLO_NO_COINCIDE"] != 1 || s.NoVerificable["CERTIFICADO_SAT_NO_DISPONIBLE"] != 2 || r.Validacion != nil {
		t.Errorf("resumen de timbres: %+v", s)
	}
	if !strings.Contains(errs, "timbres: 1 válidos, 1 inválidos, 2 no verificables") {
		t.Errorf("stderr %q", errs)
	}

	// Con -verificar también, las dos.
	lineas, r, _ = lote("-certs-sat", certs, "-verificar")
	if estado(lineas[0], "verificacion_sello") != "valido" || estado(lineas[0], "verificacion_tfd") != "valido" || r.Sellos == nil || r.SellosTFD == nil {
		t.Errorf("-verificar y -certs-sat: %v %v", lineas[0]["verificacion_sello"], lineas[0]["verificacion_tfd"])
	}

	var out, errb bytes.Buffer
	if code := run([]string{"lote", raiz, "-certs-sat", "/no/existe"}, &out, &errb); code != 1 {
		t.Errorf("-certs-sat inexistente: código %d", code)
	}
}
