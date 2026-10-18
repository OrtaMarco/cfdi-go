package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const certs = "../../testdata/certs"

func verificarJSON(t *testing.T, args ...string) (int, map[string]any, string) {
	t.Helper()
	var out, errb bytes.Buffer
	code := run(append([]string{"verificar"}, args...), &out, &errb)
	var m map[string]any
	if out.Len() > 0 {
		if err := json.Unmarshal(out.Bytes(), &m); err != nil {
			t.Fatalf("salida no es JSON: %v\n%s", err, out.String())
		}
	}
	return code, m, errb.String()
}

func estado(m map[string]any, clave string) string {
	r, _ := m[clave].(map[string]any)
	e, _ := r["estado"].(string)
	if mo, _ := r["motivo"].(string); mo != "" {
		e += " " + mo
	}
	return e
}

func TestVerificar(t *testing.T) {
	sellado := filepath.Join(sinteticos, "sellados", "ingreso40-pue.xml")

	// testdata/certs también tiene llaves .key.pem: se saltan sin aviso.
	code, m, errs := verificarJSON(t, sellado, "-certs-sat", certs)
	if code != 0 || estado(m, "sello") != "valido" || estado(m, "tfd") != "valido" || errs != "" {
		t.Errorf("sellado con certificados: código %d, %v, stderr %q", code, m, errs)
	}
	if c, _ := m["cadena_original"].(string); !strings.HasPrefix(c, "||4.0|A|1001|") || m["archivo"] != sellado {
		t.Errorf("cadena o archivo: %v", m)
	}
	if c, _ := m["cadena_original_tfd"].(string); !strings.HasPrefix(c, "||1.1|") {
		t.Errorf("cadena del TFD: %v", m["cadena_original_tfd"])
	}

	// Sin -certs-sat no se verifica el timbre.
	code, m, _ = verificarJSON(t, sellado)
	if _, ok := m["tfd"]; code != 0 || ok {
		t.Errorf("sin certificados: código %d, %v", code, m)
	}

	// cfdi33-valid trae un timbre que no corresponde: sale 1 pero escribe el JSON.
	code, m, _ = verificarJSON(t, "../../testdata/publicos/eclipxe13-CfdiUtils/tests__assets__cfdi33-valid.xml", "-certs-sat", certs)
	if code != 1 || estado(m, "sello") != "valido" || estado(m, "tfd") != "invalido SELLO_NO_COINCIDE" {
		t.Errorf("timbre inválido: código %d, %v", code, m)
	}

	// Sello que no se puede verificar: sale 1.
	code, m, _ = verificarJSON(t, filepath.Join(sinteticos, "ingreso40-pue.xml"))
	if code != 1 || estado(m, "sello") != "no_verificable CERTIFICADO_ILEGIBLE" {
		t.Errorf("certificado de mentira: código %d, %v", code, m)
	}

	// Un directorio con un .cer roto y otro bueno: avisa y sigue.
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "roto.cer"), []byte("no es certificado"), 0o644)
	os.WriteFile(filepath.Join(dir, "leeme.txt"), []byte("se ignora"), 0o644)
	os.Mkdir(filepath.Join(dir, "sub.cer"), 0o755)
	pac, _ := os.ReadFile(filepath.Join(certs, "CSD01_AAA010101AAA.cer"))
	os.WriteFile(filepath.Join(dir, "pac.cer"), pac, 0o644)
	code, m, errs = verificarJSON(t, sellado, "-certs-sat", dir)
	if code != 0 || estado(m, "tfd") != "valido" || !strings.Contains(errs, "roto.cer") || strings.Contains(errs, "leeme") {
		t.Errorf("directorio mixto: código %d, %v, stderr %q", code, m, errs)
	}
	// Sin certificados legibles: avisa y el timbre queda no verificable.
	vacio := t.TempDir()
	code, m, errs = verificarJSON(t, sellado, "-certs-sat", vacio)
	if code != 1 || estado(m, "tfd") != "no_verificable CERTIFICADO_SAT_NO_DISPONIBLE" || !strings.Contains(errs, "no hay certificados") {
		t.Errorf("directorio vacío: código %d, %v, stderr %q", code, m, errs)
	}

	if code, _, _ := verificarJSON(t, sellado, "-certs-sat", "/no/existe"); code != 1 {
		t.Errorf("directorio inexistente: código %d", code)
	}
	if code, _, _ := verificarJSON(t, "/no/existe.xml"); code != 1 {
		t.Errorf("archivo inexistente: código %d", code)
	}
	if code, _, _ := verificarJSON(t); code != 2 {
		t.Errorf("sin archivo: código %d", code)
	}
	if code, _, _ := verificarJSON(t, "-x"); code != 2 {
		t.Errorf("bandera desconocida: código %d", code)
	}
}

func TestLoteSello(t *testing.T) {
	dir := t.TempDir()
	raiz := filepath.Join(dir, "entrada")
	copiar(t, "sellados/ingreso40-pue.xml", filepath.Join(raiz, "1.xml"))
	copiar(t, "sellados/pago20.xml", filepath.Join(raiz, "2.xml"))
	copiar(t, "ingreso40-pue.xml", filepath.Join(raiz, "3.xml"))        // certificado de mentira
	copiar(t, "cadena/pagos20-en-33.xml", filepath.Join(raiz, "4.xml")) // sin cadena posible
	copiar(t, "fatales/no-es-xml.xml", filepath.Join(raiz, "5.xml"))    // fatal: no cuenta en sellos
	alterado := strings.Replace(string(leerFixture(t, "sellados/nomina12.xml")), `Total="`, `Total="1`, 1)
	os.WriteFile(filepath.Join(raiz, "6.xml"), []byte(alterado), 0o644)

	lote := func(extra ...string) ([]map[string]any, Resumen, string) {
		t.Helper()
		jsonl, resumen := filepath.Join(dir, "s.jsonl"), filepath.Join(dir, "r.json")
		var out, errb bytes.Buffer
		if code := run(append([]string{"lote", raiz, "-o", jsonl, "-resumen", resumen, "-j", "2"}, extra...), &out, &errb); code != 0 {
			t.Fatalf("código %d: %s", code, errb.String())
		}
		f, _ := os.Open(jsonl)
		defer f.Close()
		var lineas []map[string]any
		sc := bufio.NewScanner(f)
		sc.Buffer(nil, 1<<20)
		for sc.Scan() {
			var l map[string]any
			if err := json.Unmarshal(sc.Bytes(), &l); err != nil {
				t.Fatal(err)
			}
			lineas = append(lineas, l)
		}
		b, _ := os.ReadFile(resumen)
		var r Resumen
		if err := json.Unmarshal(b, &r); err != nil {
			t.Fatal(err)
		}
		return lineas, r, errb.String()
	}

	quiero := []string{"valido", "valido", "no_verificable CERTIFICADO_ILEGIBLE", "no_verificable COMPLEMENTO_SIN_CADENA", "", "invalido SELLO_NO_COINCIDE"}
	for _, bandera := range []string{"-sello", "-verificar"} {
		lineas, r, errs := lote(bandera)
		if len(lineas) != len(quiero) {
			t.Fatalf("%s: %d líneas", bandera, len(lineas))
		}
		for i, l := range lineas {
			if got := estado(l, "verificacion_sello"); got != quiero[i] {
				t.Errorf("%s línea %d: %q, quiero %q", bandera, i+1, got, quiero[i])
			}
		}
		_, conSello := lineas[0]["cfdi"].(map[string]any)["sello"]
		if conSello != (bandera == "-sello") {
			t.Errorf("%s: sello en la línea = %v", bandera, conSello)
		}
		s := r.Sellos
		if s == nil || s.Valido != 2 || s.Invalido["SELLO_NO_COINCIDE"] != 1 ||
			s.NoVerificable["CERTIFICADO_ILEGIBLE"] != 1 || s.NoVerificable["COMPLEMENTO_SIN_CADENA"] != 1 {
			t.Errorf("%s: resumen de sellos %+v", bandera, s)
		}
		if !strings.Contains(errs, "sellos: 2 válidos, 1 inválidos, 2 no verificables") {
			t.Errorf("%s: stderr %q", bandera, errs)
		}
	}

	// Sin bandera no se verifica nada.
	lineas, r, errs := lote()
	if _, ok := lineas[0]["verificacion_sello"]; ok || r.Sellos != nil || strings.Contains(errs, "sellos:") {
		t.Errorf("sin -sello: %v %+v %q", lineas[0]["verificacion_sello"], r.Sellos, errs)
	}
}
