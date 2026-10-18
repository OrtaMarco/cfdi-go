package cfdi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"slices"
	"sort"
	"strings"
	"testing"

	"github.com/OrtaMarco/cfdi-go/decimal"
)

// Validador mínimo de JSON Schema: solo las palabras que usa
// schema/cfdi.v1.schema.json ($ref local, type, const, enum, pattern,
// properties, additionalProperties: false, required, items). Una palabra
// desconocida en el esquema es un error, para que el validador no se quede
// atrás del esquema en silencio.

type esquema map[string]any

func cargarEsquema(t *testing.T) esquema {
	t.Helper()
	return cargarEsquemaDe(t, "schema/cfdi.v1.schema.json", Schema)
}

func cargarEsquemaDe(t *testing.T, ruta string, embebido []byte) esquema {
	t.Helper()
	b, err := os.ReadFile(ruta)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(b, embebido) {
		t.Fatal("el esquema embebido no es " + ruta)
	}
	var s esquema
	if err := json.Unmarshal(b, &s); err != nil {
		t.Fatal(err)
	}
	return s
}

var palabrasConocidas = map[string]bool{
	"$schema": true, "$id": true, "$defs": true, "$ref": true, "title": true, "description": true,
	"type": true, "const": true, "enum": true, "pattern": true, "properties": true,
	"additionalProperties": true, "required": true, "items": true,
}

func (raiz esquema) resolver(n map[string]any) (map[string]any, error) {
	for {
		ref, ok := n["$ref"].(string)
		if !ok {
			return n, nil
		}
		nombre, ok := strings.CutPrefix(ref, "#/$defs/")
		if !ok {
			return nil, fmt.Errorf("$ref no local: %s", ref)
		}
		def, ok := raiz["$defs"].(map[string]any)[nombre].(map[string]any)
		if !ok {
			return nil, fmt.Errorf("$ref a una definición que no existe: %s", ref)
		}
		n = def
	}
}

func (raiz esquema) validar(n map[string]any, v any, ruta string) []string {
	n, err := raiz.resolver(n)
	if err != nil {
		return []string{ruta + ": " + err.Error()}
	}
	var errs []string
	for k := range n {
		if !palabrasConocidas[k] {
			errs = append(errs, fmt.Sprintf("%s: el validador no entiende la palabra %q del esquema", ruta, k))
		}
	}
	if c, ok := n["const"]; ok && !reflect.DeepEqual(c, v) {
		errs = append(errs, fmt.Sprintf("%s: %v ≠ const %v", ruta, v, c))
	}
	if en, ok := n["enum"].([]any); ok && !slices.Contains(en, v) {
		errs = append(errs, fmt.Sprintf("%s: %v no está en enum", ruta, v))
	}
	switch tipo, _ := n["type"].(string); tipo {
	case "":
	case "string":
		s, ok := v.(string)
		if !ok {
			return append(errs, fmt.Sprintf("%s: se esperaba string y vino %T (%v)", ruta, v, v))
		}
		if p, ok := n["pattern"].(string); ok && !regexp.MustCompile(p).MatchString(s) {
			errs = append(errs, fmt.Sprintf("%s: %q no cumple %s", ruta, s, p))
		}
	case "array":
		a, ok := v.([]any)
		if !ok {
			return append(errs, fmt.Sprintf("%s: se esperaba array y vino %T", ruta, v))
		}
		items, _ := n["items"].(map[string]any)
		for i, e := range a {
			errs = append(errs, raiz.validar(items, e, fmt.Sprintf("%s[%d]", ruta, i))...)
		}
	case "object":
		o, ok := v.(map[string]any)
		if !ok {
			return append(errs, fmt.Sprintf("%s: se esperaba object y vino %T", ruta, v))
		}
		props, _ := n["properties"].(map[string]any)
		req, _ := n["required"].([]any)
		for _, r := range req {
			if _, ok := o[r.(string)]; !ok {
				errs = append(errs, fmt.Sprintf("%s: falta la clave requerida %q", ruta, r))
			}
		}
		claves := make([]string, 0, len(o))
		for k := range o {
			claves = append(claves, k)
		}
		sort.Strings(claves)
		for _, k := range claves {
			p, ok := props[k].(map[string]any)
			if !ok {
				if ap, _ := n["additionalProperties"].(bool); !ap {
					errs = append(errs, fmt.Sprintf("%s: clave %q no declarada en el esquema", ruta, k))
				}
				continue
			}
			errs = append(errs, raiz.validar(p, o[k], ruta+"."+k)...)
		}
	default:
		errs = append(errs, fmt.Sprintf("%s: type %q no soportado", ruta, tipo))
	}
	return errs
}

func validarJSON(t *testing.T, s esquema, b []byte) []string {
	t.Helper()
	var v any
	if err := json.Unmarshal(b, &v); err != nil {
		t.Fatal(err)
	}
	return s.validar(s, v, "$")
}

func TestSchemaFixtures(t *testing.T) {
	s := cargarEsquema(t)
	for _, p := range fixtures(t) {
		if esFatal(p) || esRetenciones(t, p) {
			continue
		}
		b, _ := os.ReadFile(p)
		c, err := ParseBytes(b)
		if err != nil {
			t.Fatal(err)
		}
		for _, sello := range []bool{false, true} {
			d := c.Documento(OpcionesJSON{IncluirSello: sello})
			d.Origen = p
			if sello {
				d.Validacion = Validar(b, c)
				d.VerificacionTFD = &ResultadoSello{Estado: SelloNoVerificable, Motivo: MotivoCertificadoSATFaltante}
			}
			out, _ := json.Marshal(d)
			for _, e := range validarJSON(t, s, out) {
				t.Errorf("%s (sello=%v): %s", p, sello, e)
			}
		}
	}
}

// El validador tiene que rechazar lo que debe: si no, la prueba de arriba no
// prueba nada.
func TestSchemaRechaza(t *testing.T) {
	s := cargarEsquema(t)
	c := leer(t, "ingreso40-pue.xml")
	b, _ := c.JSON()
	if errs := validarJSON(t, s, b); len(errs) > 0 {
		t.Fatalf("el documento base no valida: %v", errs)
	}
	casos := map[string][2]string{
		"monto como número":    {`"total":"1640.00"`, `"total":1640.00`},
		"monto con exponente":  {`"total":"1640.00"`, `"total":"1.64e3"`},
		"clave no declarada":   {`"moneda":"MXN"`, `"moneda":"MXN","divisa":"MXN"`},
		"clave anidada extra":  {`"uso_cfdi":"S01"`, `"uso_cfdi":"S01","correo":"x"`},
		"código desconocido":   {`"advertencias":[]`, `"advertencias":[{"codigo":"INVENTADO"}]`},
		"schema_version otro":  {`"schema_version":"1"`, `"schema_version":"2"`},
		"cantidad como número": {`"cantidad":"2.5"`, `"cantidad":2.5`},
	}
	for nombre, rep := range casos {
		mal := bytes.Replace(b, []byte(rep[0]), []byte(rep[1]), 1)
		if bytes.Equal(mal, b) {
			t.Fatalf("%s: el reemplazo no aplicó", nombre)
		}
		if errs := validarJSON(t, s, mal); len(errs) == 0 {
			t.Errorf("%s: el validador lo aceptó", nombre)
		}
	}
}

// TestSchemaCubreModelo recorre los tipos Go por reflexión y exige que el
// esquema declare exactamente sus claves JSON, con el tipo que corresponde:
// así una clave nueva en el modelo sin esquema (o al revés) falla aunque
// ningún fixture la use.
func TestSchemaCubreModelo(t *testing.T) {
	s := cargarEsquema(t)
	var docto Documento
	compararTipo(t, s, map[string]any{"type": "object", "properties": s["properties"]}, reflect.TypeOf(docto), "$", map[reflect.Type]bool{})

	// El enum de códigos es la lista de constantes.
	adv := s["$defs"].(map[string]any)["advertencia"].(map[string]any)
	enum := adv["properties"].(map[string]any)["codigo"].(map[string]any)["enum"].([]any)
	var got []string
	for _, e := range enum {
		got = append(got, e.(string))
	}
	if !slices.Equal(got, Codigos) {
		t.Errorf("enum de códigos del esquema ≠ Codigos:\n%v\n%v", got, Codigos)
	}

	// El de hallazgos, los de ValidarEsquema y ValidarReglas.
	hal := s["$defs"].(map[string]any)["hallazgo"].(map[string]any)
	got = nil
	for _, e := range hal["properties"].(map[string]any)["codigo"].(map[string]any)["enum"].([]any) {
		got = append(got, e.(string))
	}
	if quiero := append(slices.Clone(CodigosEsquema), CodigosReglas...); !slices.Equal(got, quiero) {
		t.Errorf("enum de hallazgos del esquema ≠ CodigosEsquema + CodigosReglas:\n%v\n%v", got, quiero)
	}
}

var (
	tipoDecimal = reflect.TypeOf(decimal.Decimal{})
	tipoFecha   = reflect.TypeOf(Fecha{})
)

func compararTipo(t *testing.T, s esquema, n map[string]any, typ reflect.Type, ruta string, vistos map[reflect.Type]bool) {
	t.Helper()
	ref, _ := n["$ref"].(string)
	for typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}
	switch {
	case typ == tipoDecimal:
		if ref != "#/$defs/decimal" {
			t.Errorf("%s: es decimal.Decimal y el esquema no usa #/$defs/decimal", ruta)
		}
		return
	case typ == tipoFecha:
		if ref != "#/$defs/fecha" {
			t.Errorf("%s: es Fecha y el esquema no usa #/$defs/fecha", ruta)
		}
		return
	}
	r, err := s.resolver(n)
	if err != nil {
		t.Errorf("%s: %v", ruta, err)
		return
	}
	switch typ.Kind() {
	case reflect.String:
		if r["type"] != "string" && r["const"] == nil {
			t.Errorf("%s: es string y el esquema dice %v", ruta, r["type"])
		}
	case reflect.Slice:
		if r["type"] != "array" {
			t.Errorf("%s: es lista y el esquema dice %v", ruta, r["type"])
			return
		}
		items, _ := r["items"].(map[string]any)
		compararTipo(t, s, items, typ.Elem(), ruta+"[]", vistos)
	case reflect.Struct:
		if r["type"] != "object" {
			t.Errorf("%s: es struct y el esquema dice %v", ruta, r["type"])
			return
		}
		if vistos[typ] && ref != "" {
			return
		}
		vistos[typ] = true
		props, _ := r["properties"].(map[string]any)
		declaradas := map[string]bool{}
		for k := range props {
			declaradas[k] = true
		}
		for i := 0; i < typ.NumField(); i++ {
			f := typ.Field(i)
			nombre, _, _ := strings.Cut(f.Tag.Get("json"), ",")
			if nombre == "-" || !f.IsExported() {
				continue
			}
			if nombre == "" {
				t.Errorf("%s.%s: campo sin etiqueta json", ruta, f.Name)
				continue
			}
			p, ok := props[nombre].(map[string]any)
			if !ok {
				t.Errorf("%s: la clave %q (%s.%s) no está en el esquema", ruta, nombre, typ.Name(), f.Name)
				continue
			}
			delete(declaradas, nombre)
			compararTipo(t, s, p, f.Type, ruta+"."+nombre, vistos)
		}
		for k := range declaradas {
			t.Errorf("%s: el esquema declara %q y %s no tiene ese campo", ruta, k, typ.Name())
		}
	default:
		t.Errorf("%s: tipo Go %s sin regla en la prueba", ruta, typ)
	}
}
