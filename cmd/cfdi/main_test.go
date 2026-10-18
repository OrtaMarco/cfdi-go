package main

import (
	"archive/zip"
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

const sinteticos = "../../testdata/sinteticos"

func copiar(t *testing.T, origen, destino string) {
	t.Helper()
	b, err := os.ReadFile(filepath.Join(sinteticos, origen))
	if err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Dir(destino), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(destino, b, 0o644); err != nil {
		t.Fatal(err)
	}
}

func zipBytes(t *testing.T, entradas map[string][]byte, orden []string) []byte {
	t.Helper()
	var buf bytes.Buffer
	zw := zip.NewWriter(&buf)
	for _, n := range orden {
		w, err := zw.Create(n)
		if err != nil {
			t.Fatal(err)
		}
		w.Write(entradas[n])
	}
	if err := zw.Close(); err != nil {
		t.Fatal(err)
	}
	return buf.Bytes()
}

func leerFixture(t *testing.T, n string) []byte {
	t.Helper()
	b, err := os.ReadFile(filepath.Join(sinteticos, n))
	if err != nil {
		t.Fatal(err)
	}
	return b
}

func TestParse(t *testing.T) {
	var out, errb bytes.Buffer
	if code := run([]string{"parse", filepath.Join(sinteticos, "ingreso40-pue.xml")}, &out, &errb); code != 0 {
		t.Fatalf("código %d: %s", code, errb.String())
	}
	var doc map[string]any
	if err := json.Unmarshal(out.Bytes(), &doc); err != nil {
		t.Fatal(err)
	}
	c := doc["cfdi"].(map[string]any)
	if c["total"] != "1640.00" || doc["schema_version"] != "1" {
		t.Errorf("salida inesperada: %s", out.String())
	}
	if _, ok := c["sello"]; ok {
		t.Error("sello sin -sello")
	}
	if !strings.Contains(out.String(), "\n  \"cfdi\"") {
		t.Error("la salida no está indentada")
	}

	out.Reset()
	// La bandera puede ir después del archivo.
	if code := run([]string{"parse", filepath.Join(sinteticos, "ingreso40-pue.xml"), "-sello"}, &out, &errb); code != 0 {
		t.Fatalf("código %d", code)
	}
	if !strings.Contains(out.String(), `"sello": "U0VMTE8gREUgUFJVRUJB"`) {
		t.Error("-sello no incluyó el sello")
	}

	out.Reset()
	errb.Reset()
	otra := filepath.Join(t.TempDir(), "otra.xml")
	os.WriteFile(otra, []byte(`<x:Otro xmlns:x="urn:otro"/>`), 0o644)
	if code := run([]string{"parse", otra}, &out, &errb); code != 1 {
		t.Errorf("fatal: código %d, quiero 1", code)
	}
	if out.Len() != 0 || !strings.Contains(errb.String(), "no_es_comprobante") {
		t.Errorf("fatal: stdout=%q stderr=%q", out.String(), errb.String())
	}

	// Una Retenciones se lee con ParseRetenciones.
	out.Reset()
	errb.Reset()
	if code := run([]string{"parse", filepath.Join(sinteticos, "retenciones", "ret20-plataformas.xml")}, &out, &errb); code != 0 {
		t.Fatalf("Retenciones: código %d: %s", code, errb.String())
	}
	doc = nil
	if err := json.Unmarshal(out.Bytes(), &doc); err != nil {
		t.Fatal(err)
	}
	rt, _ := doc["retenciones"].(map[string]any)
	if doc["tipo_documento"] != "retenciones" || rt == nil || rt["cve_retenc"] != "26" || doc["cfdi"] != nil {
		t.Errorf("Retenciones: %s", out.String())
	}
	// Truncada sigue siendo fatal.
	out.Reset()
	errb.Reset()
	trunca := filepath.Join(t.TempDir(), "trunca.xml")
	ret := leerFixture(t, "retenciones/ret20-plataformas.xml")
	os.WriteFile(trunca, ret[:len(ret)-60], 0o644)
	if code := run([]string{"parse", trunca}, &out, &errb); code != 1 || !strings.Contains(errb.String(), "xml_mal_formado") {
		t.Errorf("Retenciones truncada: código %d, stderr %q", code, errb.String())
	}
	if code := run([]string{"parse"}, &out, &errb); code != 2 {
		t.Errorf("sin archivo: código %d, quiero 2", code)
	}
	if code := run([]string{"parse", "/no/existe.xml"}, &out, &errb); code != 1 {
		t.Errorf("archivo inexistente: código %d, quiero 1", code)
	}
	if code := run([]string{"otro"}, &out, &errb); code != 2 {
		t.Errorf("comando desconocido: código %d", code)
	}
}

func TestLote(t *testing.T) {
	dir := t.TempDir()
	raiz := filepath.Join(dir, "entrada")
	copiar(t, "ingreso40-pue.xml", filepath.Join(raiz, "a", "uno.xml"))
	copiar(t, "anomalias.xml", filepath.Join(raiz, "a", "dos.XML"))
	copiar(t, "fatales/no-es-xml.xml", filepath.Join(raiz, "b", "roto.xml"))
	copiar(t, "pago20.xml", filepath.Join(raiz, "b", "ignorado.txt"))
	copiar(t, "nomina12.xml", filepath.Join(raiz, "__MACOSX", "._basura.xml"))

	interno := zipBytes(t, map[string][]byte{"p10.xml": leerFixture(t, "pago10.xml")}, []string{"p10.xml"})
	externo := zipBytes(t, map[string][]byte{
		"dir/":                 nil,
		"dir/pago.xml":         leerFixture(t, "pago20.xml"),
		"interno.zip":          interno,
		"sin-tfd.xml":          leerFixture(t, "sin-tfd.xml"),
		"leeme.txt":            []byte("no"),
		"__MACOSX/._pago.xml":  []byte("basura"),
		"fatales/truncado.xml": leerFixture(t, "fatales/truncado.xml"),
	}, []string{"dir/", "dir/pago.xml", "interno.zip", "sin-tfd.xml", "leeme.txt", "__MACOSX/._pago.xml", "fatales/truncado.xml"})
	if err := os.WriteFile(filepath.Join(raiz, "c.zip"), externo, 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(raiz, "dañado.zip"), []byte("PK no es zip"), 0o644); err != nil {
		t.Fatal(err)
	}
	suelto := filepath.Join(dir, "suelto.xml")
	copiar(t, "nomina12.xml", suelto)

	jsonl := filepath.Join(dir, "salida.jsonl")
	resumen := filepath.Join(dir, "resumen.json")
	var out, errb bytes.Buffer
	code := run([]string{"lote", raiz, "-o", jsonl, suelto, "-resumen", resumen, "-j", "3"}, &out, &errb)
	if code != 0 {
		t.Fatalf("código %d: %s", code, errb.String())
	}

	f, err := os.Open(jsonl)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	var origenes []string
	sc := bufio.NewScanner(f)
	sc.Buffer(nil, 1<<20)
	for sc.Scan() {
		var l map[string]any
		if err := json.Unmarshal(sc.Bytes(), &l); err != nil {
			t.Fatalf("línea no es JSON: %v", err)
		}
		o, _ := l["origen"].(string)
		origenes = append(origenes, strings.TrimPrefix(filepath.ToSlash(o), filepath.ToSlash(dir)+"/"))
		if _, fatal := l["error"]; !fatal && l["schema_version"] != "1" {
			t.Errorf("línea sin schema_version: %s", sc.Text())
		}
	}
	// Orden de WalkDir (léxico) y, dentro del zip, el orden de sus entradas.
	want := []string{
		"entrada/a/dos.XML",
		"entrada/a/uno.xml",
		"entrada/b/roto.xml",
		"entrada/c.zip!dir/pago.xml",
		"entrada/c.zip!interno.zip!p10.xml",
		"entrada/c.zip!sin-tfd.xml",
		"entrada/c.zip!fatales/truncado.xml",
		"entrada/dañado.zip",
		"suelto.xml",
	}
	if strings.Join(origenes, "\n") != strings.Join(want, "\n") {
		t.Errorf("orígenes:\n%s\nquiero:\n%s", strings.Join(origenes, "\n"), strings.Join(want, "\n"))
	}

	b, err := os.ReadFile(resumen)
	if err != nil {
		t.Fatal(err)
	}
	var r Resumen
	if err := json.Unmarshal(b, &r); err != nil {
		t.Fatal(err)
	}
	if r.Total != 9 || r.Leidos != 6 || r.Fatales != 3 || r.Advertidos != 2 {
		t.Errorf("totales: %+v", r)
	}
	if c := r.FatalesPorClase["no_es_xml"]; c == nil || c.CFDI != 1 {
		t.Errorf("no_es_xml: %+v", c)
	}
	// El ejemplo es la huella del contenido + la extensión del contenedor.
	h := sha256.Sum256(leerFixture(t, "fatales/truncado.xml"))
	if c := r.FatalesPorClase["xml_mal_formado"]; c == nil || c.CFDI != 1 || c.Ejemplos[0] != hex.EncodeToString(h[:6])+".zip" {
		t.Errorf("xml_mal_formado: %+v", c)
	}
	if c := r.FatalesPorClase["lectura"]; c == nil || c.CFDI != 1 {
		t.Errorf("lectura: %+v", c)
	}
	if c := r.AdvertenciasPorCodigo["SIN_TIMBRE"]; c == nil || c.CFDI != 1 || c.Ocurrencias != 1 {
		t.Errorf("SIN_TIMBRE: %+v", c)
	}
	if c := r.AdvertenciasPorCodigo["DECIMAL_INVALIDO"]; c == nil || c.CFDI != 1 || c.Ocurrencias != 2 {
		t.Errorf("DECIMAL_INVALIDO: %+v", c)
	}
	if r.PorVersion["4.0"] != 5 || r.PorVersion["3.3"] != 1 {
		t.Errorf("por versión: %v", r.PorVersion)
	}
	if r.PorTipo["P"] != 2 || r.PorTipo["N"] != 1 || r.PorTipo["I"] != 3 {
		t.Errorf("por tipo: %v", r.PorTipo)
	}
	// anomalias.xml declara Pagos Version="2.0" (en el namespace de 1.0): se
	// cuenta por la versión declarada, con su advertencia COMPLEMENTO_VERSION.
	if r.PorComplemento["Pagos 1.0"] != 1 || r.PorComplemento["Pagos 2.0"] != 2 || r.PorComplemento["Nomina 1.2"] != 1 || r.PorComplemento["TimbreFiscalDigital 1.1"] != 5 {
		t.Errorf("por complemento: %v", r.PorComplemento)
	}
	// Nada de RFC ni nombres en el resumen.
	for _, prohibido := range []string{"EKU9003173C9", "XAXX010101000", "XOJI740919U48", "ESCUELA", "PRUEBA", "PUBLICO"} {
		if bytes.Contains(b, []byte(prohibido)) {
			t.Errorf("el resumen contiene %q", prohibido)
		}
	}
	if !strings.Contains(errb.String(), "9 documentos: 6 leídos (2 con advertencias), 3 fatales") {
		t.Errorf("stderr: %q", errb.String())
	}
}

// Con el primer documento atorado, el productor debe detenerse al llenar la
// ventana en vez de seguir acumulando resultados detrás de él; al soltarlo,
// todo sale en el orden de emisión.
func TestCanalizarAcotadoYEnOrden(t *testing.T) {
	const n, workers, ventana = 5000, 4, 16
	var emitidos atomic.Int64
	soltar := make(chan struct{})
	producir := func(emitir func(trabajo)) {
		for i := range n {
			emitidos.Add(1)
			emitir(trabajo{origen: strconv.Itoa(i)})
		}
	}
	procesar := func(tr trabajo) resultado {
		if tr.origen == "0" {
			<-soltar
		}
		return resultado{origen: tr.origen}
	}
	var salida []string
	listo := make(chan struct{})
	go func() {
		defer close(listo)
		canalizar(producir, workers, ventana, procesar, func(r resultado) { salida = append(salida, r.origen) })
	}()

	// El 0 está en el consumidor, ventana canales esperan en orden y el
	// productor queda bloqueado con uno más: ventana+2 emitidos.
	limite := int64(ventana + 2)
	for deadline := time.Now().Add(5 * time.Second); emitidos.Load() < limite && time.Now().Before(deadline); {
		time.Sleep(time.Millisecond)
	}
	time.Sleep(50 * time.Millisecond)
	if e := emitidos.Load(); e != limite {
		t.Errorf("con la cabeza atorada se emitieron %d trabajos, quiero %d", e, limite)
	}
	close(soltar)
	<-listo

	if len(salida) != n {
		t.Fatalf("%d resultados, quiero %d", len(salida), n)
	}
	for i, o := range salida {
		if o != strconv.Itoa(i) {
			t.Fatalf("resultado %d es %q: el orden de salida no es el de emisión", i, o)
		}
	}
}

// El resumen se comparte: sin -rutas ningún pedazo de ruta (carpeta, nombre
// de archivo, entrada de zip) puede aparecer en él, aunque lleve un RFC.
func TestResumenSinRutas(t *testing.T) {
	dir := t.TempDir()
	raiz := filepath.Join(dir, "CLIENTE-EKU9003173C9")
	copiar(t, "anomalias.xml", filepath.Join(raiz, "XOJI740919U48-factura.xml"))
	copiar(t, "fatales/no-es-xml.xml", filepath.Join(raiz, "XAXX010101000-roto.xml"))
	z := zipBytes(t, map[string][]byte{
		"carpeta-interna/sin-timbre-AAA010101AAA.xml": leerFixture(t, "sin-tfd.xml"),
		"carpeta-interna/truncado-CACX7605101P8.xml":  leerFixture(t, "fatales/truncado.xml"),
	}, []string{"carpeta-interna/sin-timbre-AAA010101AAA.xml", "carpeta-interna/truncado-CACX7605101P8.xml"})
	if err := os.WriteFile(filepath.Join(raiz, "descarga-XIA190128J61.zip"), z, 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(raiz, "dañado-ZZZ.zip"), []byte("PK no"), 0o644); err != nil {
		t.Fatal(err)
	}
	prohibidos := []string{dir, "CLIENTE", "EKU9003173C9", "XOJI740919U48", "factura", "XAXX010101000", "roto",
		"carpeta-interna", "AAA010101AAA", "sin-timbre", "CACX7605101P8", "truncado", "descarga", "XIA190128J61", "dañado", "ZZZ", "!"}

	resumen := func(extra ...string) (string, Resumen) {
		t.Helper()
		p := filepath.Join(dir, "resumen.json")
		var out, errb bytes.Buffer
		if code := run(append([]string{"lote", raiz, "-o", filepath.Join(dir, "salida.jsonl"), "-resumen", p}, extra...), &out, &errb); code != 0 {
			t.Fatalf("código %d: %s", code, errb.String())
		}
		b, err := os.ReadFile(p)
		if err != nil {
			t.Fatal(err)
		}
		var r Resumen
		if err := json.Unmarshal(b, &r); err != nil {
			t.Fatal(err)
		}
		return string(b), r
	}

	b, r := resumen()
	if r.Fatales != 3 || r.Advertidos != 2 {
		t.Fatalf("el lote de prueba no tiene los fatales y advertidos esperados: %+v", r)
	}
	for _, p := range prohibidos {
		if strings.Contains(b, p) {
			t.Errorf("sin -rutas, el resumen contiene %q:\n%s", p, b)
		}
	}
	formato := regexp.MustCompile(`^[0-9a-f]{12}\.(xml|zip)$`)
	ejemplos := 0
	for _, m := range []map[string]*Conteo{r.FatalesPorClase, r.AdvertenciasPorCodigo} {
		for _, c := range m {
			for _, e := range c.Ejemplos {
				ejemplos++
				if !formato.MatchString(e) {
					t.Errorf("ejemplo %q no es huella.ext", e)
				}
			}
		}
	}
	if ejemplos == 0 {
		t.Error("el resumen no trae ejemplos")
	}

	// Con -rutas sí van completas, y el JSONL siempre lleva la ruta.
	b, _ = resumen("-rutas")
	for _, p := range []string{"XOJI740919U48-factura.xml", "descarga-XIA190128J61.zip!carpeta-interna/truncado-CACX7605101P8.xml", "dañado-ZZZ.zip"} {
		if !strings.Contains(b, p) {
			t.Errorf("con -rutas, falta %q en el resumen", p)
		}
	}
	jsonl, _ := os.ReadFile(filepath.Join(dir, "salida.jsonl"))
	if !strings.Contains(string(jsonl), `"origen":"`+filepath.Join(raiz, "XOJI740919U48-factura.xml")+`"`) {
		t.Errorf("el JSONL no lleva el origen completo")
	}
}

// Una entrada de zip con el CRC roto es fatal de lectura aunque se haya leído
// con el tamaño declarado.
func TestLoteZipCorrupto(t *testing.T) {
	var buf bytes.Buffer
	zw := zip.NewWriter(&buf)
	w, err := zw.CreateHeader(&zip.FileHeader{Name: "a.xml", Method: zip.Store})
	if err != nil {
		t.Fatal(err)
	}
	w.Write(leerFixture(t, "ingreso40-pue.xml"))
	zw.Close()
	b := buf.Bytes()
	i := bytes.Index(b, []byte("Comprobante"))
	b[i] = 'X' // mismo tamaño, otro CRC
	p := filepath.Join(t.TempDir(), "corrupto.zip")
	if err := os.WriteFile(p, b, 0o644); err != nil {
		t.Fatal(err)
	}
	var out, errb bytes.Buffer
	run([]string{"lote", p}, &out, &errb)
	if !strings.Contains(out.String(), `"clase":"lectura"`) || !strings.Contains(out.String(), "checksum") {
		t.Errorf("salida: %s", out.String())
	}
}

func TestLoteStdoutYErrores(t *testing.T) {
	var out, errb bytes.Buffer
	code := run([]string{"lote", "-j", "1", filepath.Join(sinteticos, "pago10.xml"), "/no/existe"}, &out, &errb)
	if code != 1 {
		t.Errorf("ruta inexistente: código %d, quiero 1", code)
	}
	if n := strings.Count(out.String(), "\n"); n != 1 {
		t.Errorf("%d líneas en stdout, quiero 1", n)
	}
	if !strings.Contains(errb.String(), "/no/existe") {
		t.Errorf("stderr no menciona la ruta: %q", errb.String())
	}
	if strings.Contains(out.String(), "sello_pago") {
		t.Error("sello_pago sin -sello")
	}
	if code := run([]string{"lote"}, &out, &errb); code != 2 {
		t.Errorf("sin rutas: código %d", code)
	}
}
