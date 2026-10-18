package gencadena

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const dirSAT = "xslt/sitio_internet"

// copiaSAT copia las hojas a un directorio temporal y aplica un reemplazo a
// una de ellas. Copia también esquemas/ (Retenciones), al lado de
// sitio_internet; archivo es relativo a sitio_internet, como en Hojas.
func copiaSAT(t *testing.T, archivo, viejo, nuevo string) string {
	t.Helper()
	dst := t.TempDir()
	base := filepath.Dir(dirSAT)
	err := filepath.WalkDir(base, func(p string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		rel, _ := filepath.Rel(base, p)
		relSI, _ := filepath.Rel(dirSAT, p)
		b, err := os.ReadFile(p)
		if err != nil {
			return err
		}
		if filepath.ToSlash(relSI) == archivo {
			s := string(b)
			if !strings.Contains(s, viejo) {
				t.Fatalf("%s no contiene %q", archivo, viejo)
			}
			b = []byte(strings.Replace(s, viejo, nuevo, 1))
		}
		if err := os.MkdirAll(filepath.Join(dst, filepath.Dir(rel)), 0o755); err != nil {
			return err
		}
		return os.WriteFile(filepath.Join(dst, rel), b, 0o644)
	})
	if err != nil {
		t.Fatal(err)
	}
	return filepath.Join(dst, filepath.Base(dirSAT))
}

func TestGenerarHojasDelSAT(t *testing.T) {
	b, err := Generar(dirSAT)
	if err != nil {
		t.Fatal(err)
	}
	for _, h := range Hojas {
		if !strings.Contains(string(b), "var "+h.Var+" = &hojaCadena{") {
			t.Errorf("falta %s", h.Var)
		}
	}
	const ret2 = "http://www.sat.gob.mx/esquemas/retencionpago/2"
	for _, s := range []string{
		// Retenciones: ./retenciones:Nacional/@RfcR es el atributo del primer
		// Nacional, como $param/@attr.
		"conjunto(rel(hijo(\"" + ret2 + "\", \"Nacional\")),\n\t\t\t\t\treq(\"RfcR\"),\n\t\t\t\t),",
		// apply-templates sobre ImpRetenidos, que no tiene plantilla.
		"integradas: map[nombreXML]bool{\n\t\t{\"" + ret2 + "\", \"ImpRetenidos\"}: true,\n\t},",
		// Las utilerías que incluye como ../../../sitio_internet/… son las mismas.
		"//   - ../esquemas/retencionpago/2/retenciones.xslt ",
		// La hoja de 3.2 lleva ine10.xslt e ine11.xslt, idénticas: una sola
		// plantilla.
		"//   - cfd/ine/ine10.xslt ",
		"//   - cfd/ine/ine11.xslt ",
		// vehiculousado.xslt trae su propio match="/", que la raíz sustituye.
		"// Plantillas match=\"/\" de hojas incluidas, que la de la raíz sustituye:\n//   - cfd/vehiculousado/vehiculousado.xslt\n",
		// Nómina 1.1: el xsl:choose por Version es un siIgual por cada when.
		"siIgual(\"Version\", \"1.0\",\n\t\t\t\treq(\"Version\"),\n\t\t\t\treq(\"RegistroPatronal\"),",
		"siIgual(\"Version\", \"1.1\",\n\t\t\t\treq(\"Version\"),\n\t\t\t\topc(\"RegistroPatronal\"),",
	} {
		if !strings.Contains(string(b), s) {
			t.Errorf("falta en lo generado:\n%s", s)
		}
	}
	// Las hojas de CFDI no tienen elementos sin plantilla.
	if strings.Count(string(b), "integradas:") != 2 {
		t.Errorf("hay %d hojas con integradas, quiero 2 (las de Retenciones)", strings.Count(string(b), "integradas:"))
	}
}

// Cualquier cosa fuera del subconjunto conocido detiene la generación.
func TestGenerarRechazaLoDesconocido(t *testing.T) {
	const cfdi40 = "cfd/4/cadenaoriginal_4_0/cadenaoriginal_4_0.xslt"
	const util = "cfd/2/cadenaoriginal_2_0/utilerias.xslt"
	const ret10 = "../esquemas/retencionpago/1/retenciones.xslt"
	const ret20 = "../esquemas/retencionpago/2/retenciones.xslt"
	const nom11 = "cfd/nomina/nomina11.xslt"
	casos := []struct {
		nombre, archivo, viejo, nuevo, quiero string
	}{
		{"value-of suelto", cfdi40, `<xsl:apply-templates select="./cfdi:Emisor"/>`, `<xsl:value-of select="./@Version"/>`, "instrucción no soportada"},
		{"Requerido cambiado", util, `<xsl:param name="valor"/>|`, `<xsl:param name="valor"/>||`, "la plantilla Requerido cambió"},
		{"normalize-space quitado", util, `normalize-space(string($s))`, `string($s)`, "la plantilla ManejaEspacios cambió"},
		{"plantilla repetida", cfdi40, `<xsl:template match="cfdi:Receptor">`, `<xsl:template match="cfdi:Emisor">`, "repetida"},
		{"predicado en una ruta", cfdi40, `select="./cfdi:Conceptos"`, `select="./cfdi:Conceptos[1]"`, "no soportado"},
		{"prefijo sin declarar", cfdi40, `select="./cfdi:Receptor"`, `select="./zz:Receptor"`, "prefijo"},
		{"valor que no es atributo", cfdi40, `select="./@Serie"`, `select="'A'"`, "se esperaba un atributo"},
		{"test con función", cfdi40, `test="./cfdi:CuentaPredial"`, `test="count(./cfdi:CuentaPredial) > 1"`, "no soportad"},
		{"elemento literal", cfdi40, `<xsl:apply-templates select="./cfdi:Emisor"/>`, `<salida/>`, "elemento literal"},
		{"salida xml", cfdi40, `method="text"`, `method="xml"`, "method"},
		{"match con ruta", cfdi40, `match="cfdi:Receptor"`, `match="cfdi:Comprobante/cfdi:Receptor"`, "no soportado"},
		{"atributo de un hijo con predicado", ret20, `select="./retenciones:Nacional/@RfcR"`, `select="./retenciones:Nacional[1]/@RfcR"`, "se esperaba un atributo"},
		{"atributo de un hijo con función", ret20, `select="./retenciones:Nacional/@RfcR"`, `select="./retenciones:Nacional/@RfcR|./@X"`, "se esperaba un atributo"},
		{"Retenciones con value-of", ret10, `<xsl:apply-templates select="./retenciones:Emisor"/>`, `<xsl:value-of select="./@Version"/>`, "instrucción no soportada"},
		{"plantilla inexistente", cfdi40, `name="Opcional">
      <xsl:with-param name="valor" select="./@Serie"/>`, `name="Otra">
      <xsl:with-param name="valor" select="./@Serie"/>`, "no definida"},
		// xsl:choose: solo when que comparan el mismo atributo con valores
		// distintos (nomina11.xslt), sin otherwise.
		{"choose con otherwise", nom11, `</xsl:choose>`, `<xsl:otherwise/></xsl:choose>`, "sin otherwise"},
		{"choose con atributos distintos", nom11, `<xsl:when test="./@Version='1.1'">`, `<xsl:when test="./@Otra='1.1'">`, "atributos distintos"},
		{"choose con valor repetido", nom11, `<xsl:when test="./@Version='1.1'">`, `<xsl:when test="./@Version='1.0'">`, "mismo valor"},
		{"choose con test de nodo", nom11, `<xsl:when test="./@Version='1.1'">`, `<xsl:when test="./nomina:Percepciones">`, "no soportado"},
		{"choose vacío", nom11, `<xsl:choose>`, `<xsl:choose/><xsl:choose>`, "sin when"},
		// ine10.xslt e ine11.xslt (las dos en la hoja de 3.2) son idénticas: una
		// plantilla repetida solo se acepta si es igual.
		{"plantilla repetida distinta", "cfd/ine/ine10.xslt", `select="./@TipoComite"`, `select="./@Otro"`, "repetida y distinta"},
	}
	for _, c := range casos {
		_, err := Generar(copiaSAT(t, c.archivo, c.viejo, c.nuevo))
		if err == nil || !strings.Contains(err.Error(), c.quiero) {
			t.Errorf("%s: error %v, quiero que diga %q", c.nombre, err, c.quiero)
		}
	}

	// Un match="/" en una hoja incluida (vehiculousado.xslt lo trae) solo se
	// acepta si el de la raíz va después: si la raíz lo tiene antes del
	// include, XSLT usaría el de la incluida.
	dir := copiaSAT(t, "cfd/3/cadenaoriginal_3_3/cadenaoriginal_3_3.xslt", `</xsl:stylesheet>`,
		`<xsl:include href="../../ine/ine10.xslt"/></xsl:stylesheet>`)
	p := filepath.Join(dir, "cfd/ine/ine10.xslt")
	b, err := os.ReadFile(p)
	if err != nil {
		t.Fatal(err)
	}
	b = []byte(strings.Replace(string(b), `<xsl:template match="ine:INE">`, `<xsl:template match="/">|</xsl:template><xsl:template match="ine:INE">`, 1))
	if err := os.WriteFile(p, b, 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := Generar(dir); err == nil || !strings.Contains(err.Error(), "después del de la raíz") {
		t.Errorf("match=\"/\" incluido después del de la raíz: error %v", err)
	}
	if _, err := Generar(t.TempDir()); err == nil {
		t.Error("un directorio vacío no dio error")
	}
}
