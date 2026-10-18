package genxsd

import (
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

const dirSAT = "xsd/sitio_internet"

// copiaSAT arma un directorio temporal con los XSD (enlaces simbólicos, para
// no copiar los catálogos grandes) y aplica un reemplazo a uno de ellos.
// Copia también esquemas/ (Retenciones), al lado de sitio_internet; archivo
// es relativo a sitio_internet, como en Raices.
func copiaSAT(t *testing.T, archivo, viejo, nuevo string) string {
	t.Helper()
	dst := t.TempDir()
	abs, err := filepath.Abs(filepath.Dir(dirSAT))
	if err != nil {
		t.Fatal(err)
	}
	absSI := filepath.Join(abs, filepath.Base(dirSAT))
	encontrado := false
	err = filepath.WalkDir(abs, func(p string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		rel, _ := filepath.Rel(abs, p)
		relSI, _ := filepath.Rel(absSI, p)
		if err := os.MkdirAll(filepath.Join(dst, filepath.Dir(rel)), 0o755); err != nil {
			return err
		}
		if filepath.ToSlash(relSI) != archivo {
			return os.Symlink(p, filepath.Join(dst, rel))
		}
		encontrado = true
		b, err := os.ReadFile(p)
		if err != nil {
			return err
		}
		s := string(b)
		if !strings.Contains(s, viejo) {
			t.Fatalf("%s no contiene %q", archivo, viejo)
		}
		return os.WriteFile(filepath.Join(dst, rel), []byte(strings.Replace(s, viejo, nuevo, 1)), 0o644)
	})
	if err != nil {
		t.Fatal(err)
	}
	if !encontrado {
		t.Fatalf("no existe %s", archivo)
	}
	return filepath.Join(dst, filepath.Base(dirSAT))
}

func TestGenerarEsquemasDelSAT(t *testing.T) {
	r, c, err := Generar(dirSAT)
	if err != nil {
		t.Fatal(err)
	}
	for _, s := range []string{"eCFDI40_Comprobante", "eCFDI33_Comprobante", "eTFD11_TimbreFiscalDigital", "ePagos20_Pagos",
		"eNomina12_Nomina", "eImpLocal_ImpuestosLocales", "eCartaPorte31_CartaPorte", "ePagos10_Pagos",
		"eCartaPorte30_CartaPorte", "eCartaPorte20_CartaPorte", "eComExt20_ComercioExterior", "eComExt11_ComercioExterior",
		"eHidroYPetro_HidroYPetro", "eRegistroFiscal_CFDIRegistroFiscal", "eDivisas_Divisas", "eECC12_EstadoDeCuentaCombustible",
		"eTerceros11_PorCuentadeTerceros", "eVentaVehiculos11_VentaVehiculos", "eAerolineas_Aerolineas",
		"eLeyendasFisc_LeyendasFiscales", "eIEDU_instEducativas", "eDonat11_Donatarias",
		"eRet20_Retenciones", "eRet10_Retenciones", "eRetPlataformas10_ServiciosPlataformasTecnologicas", "eRetIntereses_Intereses",
		"eRetPagosExtranjeros_Pagosaextranjeros", "eRetDividendos_Dividendos", "eRetEnajenacion_EnajenaciondeAcciones",
		"eRetArrendamiento_Arrendamientoenfideicomiso", "eRetDerivados_Operacionesconderivados", "eRetHipotecarios_Intereseshipotecarios",
		"eRetFideicomiso_Fideicomisonoempresarial", "eRetPlanesRetiro10_Planesderetiro", "eRetPlanesRetiro11_Planesderetiro",
		"eRetPremios_Premios", "eRetSectorFinanciero_SectorFinanciero"} {
		if !strings.Contains(string(r), "var "+s+" = &elementoXSD{") {
			t.Errorf("falta %s", s)
		}
	}
	// TFD 1.0: variante del elemento de TFD 1.1, elegida por version="1.0".
	if !strings.Contains(string(r), "var eTFD10_TimbreFiscalDigital = &elementoXSD{") ||
		!strings.Contains(string(r), "{atributo: \"version\", valor: \"1.0\", elem: eTFD10_TimbreFiscalDigital}, // si no, eTFD11_TimbreFiscalDigital") {
		t.Error("falta la variante TFD 1.0")
	}
	// Donatarias y VentaVehiculos 1.0: variantes por version="1.0", que en la
	// raíz es fixed="1.1".
	for _, s := range []string{
		"{atributo: \"version\", valor: \"1.0\", elem: eDonat10_Donatarias}, // si no, eDonat11_Donatarias",
		"{atributo: \"version\", valor: \"1.0\", elem: eVentaVehiculos10_VentaVehiculos}, // si no, eVentaVehiculos11_VentaVehiculos",
	} {
		if !strings.Contains(string(r), s) {
			t.Errorf("falta la variante:\n%s", s)
		}
	}
	for _, s := range []string{"eCartaPorte10_CartaPorte", "eComExt10_ComercioExterior", "eECC11_EstadoDeCuentaCombustible", "eNomina11_Nomina"} {
		if !strings.Contains(string(r), "var "+s+" = &elementoXSD{") {
			t.Errorf("falta %s", s)
		}
	}
	if !strings.Contains(string(c), "var catCatCFDI_cMoneda = &catalogoXSD{") {
		t.Error("falta el catálogo c_Moneda")
	}
	// terceros 1.1: el xs:choice dentro de la secuencia, y el tipo complejo
	// con nombre traducido en cada elemento que lo usa.
	for _, s := range []string{
		"{min: 1, max: 1, opciones: []particulaXSD{\n\t\t\t{elem: eTerceros11_PorCuentadeTerceros_InformacionAduanera, min: 0, max: 1},\n\t\t\t{elem: eTerceros11_PorCuentadeTerceros_Parte, min: 0, max: -1},\n\t\t\t{elem: eTerceros11_PorCuentadeTerceros_CuentaPredial, min: 0, max: 1},\n\t\t}},\n\t\t{elem: eTerceros11_PorCuentadeTerceros_Impuestos, min: 1, max: 1},",
		"var eTerceros11_PorCuentadeTerceros_Parte_InformacionAduanera = &elementoXSD{ns: \"http://www.sat.gob.mx/terceros\", nombre: \"InformacionAduanera\", vacio: true,\n\tattrs: []atributoXSD{\n\t\t{nombre: \"numero\"",
		"var eTerceros11_PorCuentadeTerceros_InformacionFiscalTercero = &elementoXSD{ns: \"http://www.sat.gob.mx/terceros\", nombre: \"InformacionFiscalTercero\", vacio: true,\n\tattrs: []atributoXSD{\n\t\t{nombre: \"calle\"",
		// ventavehiculos 1.1: un xs:choice como todo el contenido.
		"var eVentaVehiculos11_VentaVehiculos = &elementoXSD{ns: \"http://www.sat.gob.mx/ventavehiculos\", nombre: \"VentaVehiculos\", seqMin: 1, seqMax: 1,",
	} {
		if !strings.Contains(string(r), s) {
			t.Errorf("falta en lo generado:\n%s", s)
		}
	}
	// enumeration en un xs:int (CertificadoOrigen de Comercio Exterior).
	if !strings.Contains(string(c), "var catComExt20_ComercioExterior_CertificadoOrigen = &catalogoXSD{nombre: \"ComExt20 ComercioExterior/@CertificadoOrigen\", valores: `\n0\n1\n`}") {
		t.Error("falta el catálogo entero de CertificadoOrigen")
	}
	// enumeration en un xs:decimal (c_TasaCuota de Plataformas): por valor.
	if !strings.Contains(string(c), "nombre: \"c_TasaCuota\", valores: `\n0\n0.08\n0.16\n0.5\n`}") {
		t.Error("falta el catálogo decimal c_TasaCuota en forma canónica")
	}
	// FechaExp de Retenciones 1.0: minInclusive de un xs:dateTime con zona.
	if !strings.Contains(string(r), "{tipo: \"minInclusive\", valor: \"2014-01-01T00:00:00-06:00\"}") {
		t.Error("falta el minInclusive de FechaExp")
	}
	// SPEI: xs:all, xs:time y el rango de un xs:date.
	for _, s := range []string{
		"var eSPEI_ComplementoSPEI_SPEITercero = &elementoXSD{ns: \"http://www.sat.gob.mx/spei\", nombre: \"SPEI_Tercero\", seqMin: 1, seqMax: 1, todos: true,",
		"prim: \"time\"",
		"{tipo: \"maxInclusive\", valor: \"9999-12-31\"}",
	} {
		if !strings.Contains(string(r), s) {
			t.Errorf("falta en lo generado:\n%s", s)
		}
	}
	// Fecha de ECC 1.1: minInclusive de un xs:dateTime sin zona.
	if !strings.Contains(string(r), "{tipo: \"minInclusive\", valor: \"2015-01-01T00:00:00\"}") {
		t.Error("falta el minInclusive sin zona de ECC 1.1")
	}
}

// Un xs:choice repetible como todo el contenido (así viene Movimientos del
// ECB) son vueltas de una secuencia de una sola opción: la partícula sigue
// con max 1 y la repetición va en seqMax. Se prueba sobre VentaVehiculos 1.1,
// cuyo contenido es un xs:choice.
func TestGenerarChoiceRepetible(t *testing.T) {
	const vv = "cfd/ventavehiculos/ventavehiculos11.xsd"
	for _, c := range []struct{ choice, secuencia, particula string }{
		{`<xs:choice maxOccurs="unbounded">`, "seqMin: 1, seqMax: -1", "{min: 1, max: 1, opciones:"},
		{`<xs:choice minOccurs="0" maxOccurs="3">`, "seqMin: 0, seqMax: 3", "{min: 0, max: 1, opciones:"},
	} {
		r, _, err := Generar(copiaSAT(t, vv, `<xs:choice>`, c.choice))
		if err != nil {
			t.Fatalf("%s: %v", c.choice, err)
		}
		for _, s := range []string{
			`nombre: "VentaVehiculos", ` + c.secuencia + ",\n",
			"particulas: []particulaXSD{\n\t\t" + c.particula + " []particulaXSD{\n\t\t\t{elem: eVentaVehiculos11_VentaVehiculos_InformacionAduanera",
		} {
			if !strings.Contains(string(r), s) {
				t.Errorf("%s: falta en lo generado:\n%s", c.choice, s)
			}
		}
	}
}

// Cualquier cosa fuera del subconjunto conocido detiene la generación.
func TestGenerarRechazaLoDesconocido(t *testing.T) {
	const (
		cfdi40 = "cfd/4/cfdv40.xsd"
		tfd    = "cfd/TimbreFiscalDigital/TimbreFiscalDigitalv11.xsd"
		tfd10  = "cfd/TimbreFiscalDigital/TimbreFiscalDigital.xsd"
		td     = "cfd/tipoDatos/tdCFDI/tdCFDI.xsd"
		cat    = "cfd/catalogos/Pagos/catPagos.xsd"
		vv     = "cfd/ventavehiculos/ventavehiculos11.xsd"
		ter    = "cfd/terceros/terceros11.xsd"
		cee    = "cfd/ComercioExterior20/ComercioExterior20.xsd"
		ret10  = retencion + "1/retencionpagov1.xsd"
		ret20  = retencion + "2/retencionpagov2.xsd"
		fechaH = `<xs:pattern value="(20[1-9][0-9])-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])T(([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9])"/>`

		catPlat = retencion + "1/PlataformasTecnologicas10/catalogos/CatPlataformasTecnologicas.xsd"
	)
	seqEmisor := `<xs:element name="Emisor">`
	casos := []struct {
		nombre, archivo, viejo, nuevo, quiero string
	}{
		{"xs:choice de una opción", cfdi40, seqEmisor, `<xs:choice><xs:element name="Z"><xs:complexType/></xs:element></xs:choice>` + seqEmisor, "menos de dos opciones"},
		{"xs:choice repetible en una secuencia", ter, `<xs:choice>`, `<xs:choice maxOccurs="2">`, "maxOccurs"},
		{"xs:choice de contenido con minOccurs 2", vv, `<xs:choice>`, `<xs:choice minOccurs="2" maxOccurs="unbounded">`, "minOccurs"},
		{"xs:choice con secuencia", vv, `<xs:choice>`, `<xs:choice><xs:sequence/>`, "construcción no soportada"},
		{"xs:choice con atributo raro", vv, `<xs:choice>`, `<xs:choice id="x">`, "no soportado"},
		{"xs:choice después de atributos", vv, `<xs:choice>`, `<xs:attribute name="a"/><xs:choice>`, "después de los atributos"},
		{"xs:choice con nombre repetido", ter, `<xs:element name="CuentaPredial" minOccurs="0">`, `<xs:element name="InformacionFiscalTercero" minOccurs="0">`, "repetido"},
		{"secuencia anidada", cfdi40, seqEmisor, `<xs:sequence/>` + seqEmisor, "construcción no soportada"},
		{"complexType con nombre sin usar", cfdi40, `<xs:import namespace="http://www.sat.gob.mx/sitio_internet/cfd/catalogos"`, `<xs:complexType name="T"/><xs:import namespace="http://www.sat.gob.mx/sitio_internet/cfd/catalogos"`, "sin usar"},
		{"complexType con nombre repetido", vv, `<xs:complexType name="t_InformacionAduanera">`, `<xs:complexType name="t_InformacionAduanera"/><xs:complexType name="t_InformacionAduanera">`, "repetido"},
		{"complexType global sin nombre", vv, `<xs:complexType name="t_InformacionAduanera">`, `<xs:complexType/><xs:complexType name="t_InformacionAduanera">`, "sin nombre"},
		{"complexType con nombre y mixed", vv, `<xs:complexType name="t_InformacionAduanera">`, `<xs:complexType name="t_InformacionAduanera" mixed="true">`, "no soportado"},
		{"type y complexType anónimo", vv, `type="ventavehiculos:t_InformacionAduanera" minOccurs="0" maxOccurs="unbounded">`, `type="ventavehiculos:t_InformacionAduanera" minOccurs="0" maxOccurs="unbounded"><xs:complexType/>`, "se esperaba exactamente"},
		{"type de tipo complejo sin definir", vv, `type="ventavehiculos:t_InformacionAduanera" minOccurs="0" maxOccurs="unbounded">`, `type="ventavehiculos:t_Nada" minOccurs="0" maxOccurs="unbounded">`, "no definido"},
		{"type con prefijo sin declarar", vv, `type="ventavehiculos:t_InformacionAduanera" minOccurs="0" maxOccurs="unbounded">`, `type="zz:t_InformacionAduanera" minOccurs="0" maxOccurs="unbounded">`, "prefijo"},
		{"extension con contenido", ter, `<xs:extension base="terceros:t_UbicacionFiscal"/>`, `<xs:extension base="terceros:t_UbicacionFiscal"><xs:attribute name="x"/></xs:extension>`, "contenido propio"},
		{"extension con atributo raro", ter, `<xs:extension base="terceros:t_UbicacionFiscal"/>`, `<xs:extension base="terceros:t_UbicacionFiscal" id="x"/>`, "no soportado"},
		{"extension de tipo sin definir", ter, `<xs:extension base="terceros:t_UbicacionFiscal"/>`, `<xs:extension base="terceros:t_Nada"/>`, "no definido"},
		{"restriction de complexContent", ter, `<xs:extension base="terceros:t_UbicacionFiscal"/>`, `<xs:restriction base="terceros:t_UbicacionFiscal"/>`, "construcción no soportada"},
		{"complexContent vacío", ter, `<xs:extension base="terceros:t_UbicacionFiscal"/>`, ``, "exactamente una xs:extension"},
		{"complexContent con atributo raro", ter, `<xs:complexContent>
							<xs:extension base="terceros:t_UbicacionFiscal"/>`, `<xs:complexContent id="x">
							<xs:extension base="terceros:t_UbicacionFiscal"/>`, "no soportado"},
		{"complexContent junto con atributos", ter, `<xs:extension base="terceros:t_UbicacionFiscal"/>
						</xs:complexContent>`, `<xs:extension base="terceros:t_UbicacionFiscal"/>
						</xs:complexContent><xs:attribute name="x"/>`, "complexContent junto"},
		{"base con complexContent", ter, `<xs:complexType name="t_UbicacionFiscal">`, `<xs:complexType name="t_UbicacionFiscal"><xs:complexContent/>`, "construcción no soportada"},
		{"enumeration de un int que no es entero", cee, `<xs:enumeration value="0"/>`, `<xs:enumeration value="0.5"/>`, "no es un entero"},
		{"attributeFormDefault qualified", tfd, `attributeFormDefault="unqualified"`, `attributeFormDefault="qualified"`, "attributeFormDefault"},
		{"elemento con ref", cfdi40, seqEmisor, `<xs:element ref="cfdi:Emisor"/>` + seqEmisor, "no soportado"},
		{"elemento con type", cfdi40, seqEmisor, `<xs:element name="X" type="xs:string"/>` + seqEmisor, "no soportado"},
		{"complexType mixed", cfdi40, `<xs:element name="Emisor">
					<xs:annotation>
						<xs:documentation>Nodo requerido para expresar la información del contribuyente emisor del comprobante.</xs:documentation>
					</xs:annotation>
					<xs:complexType>`, `<xs:element name="Emisor"><xs:complexType mixed="true">`, "no soportado"},
		{"anyAttribute", tfd, `<xs:attribute name="Version"`, `<xs:anyAttribute/><xs:attribute name="Version"`, "construcción no soportada"},
		{"any lax", cfdi40, `<xs:any minOccurs="0" maxOccurs="unbounded"/>`, `<xs:any minOccurs="0" maxOccurs="unbounded" processContents="lax"/>`, "processContents"},
		{"any de otro namespace", cfdi40, `<xs:any minOccurs="0" maxOccurs="unbounded"/>`, `<xs:any minOccurs="0" maxOccurs="unbounded" namespace="##other"/>`, "namespace"},
		{"atributo con default", tfd, `<xs:attribute name="UUID" use="required" id="UUID">`, `<xs:attribute name="UUID" default="x" id="UUID">`, "no soportado"},
		{"use prohibited", tfd, `<xs:attribute name="UUID" use="required" id="UUID">`, `<xs:attribute name="UUID" use="prohibited" id="UUID">`, "no soportado"},
		{"minExclusive", td, `<xs:minInclusive value="0.000000"/>`, `<xs:minExclusive value="0"/>`, "faceta no soportada"},
		{"lista", td, `<xs:simpleType name="t_Importe">`, `<xs:simpleType name="t_Lista"><xs:list itemType="xs:string"/></xs:simpleType><xs:simpleType name="t_Importe">`, "construcción no soportada"},
		{"tipo integrado desconocido", td, `<xs:restriction base="xs:decimal">`, `<xs:restriction base="xs:long">`, "no soportado"},
		{"patrón con resta de clases", td, `<xs:pattern value="[0-9]{1,18}(.[0-9]{1,6})?"/>`, `<xs:pattern value="[a-z-[aeiou]]"/>`, "resta de clases"},
		{"patrón con bloque Unicode", td, `<xs:pattern value="[0-9]{1,18}(.[0-9]{1,6})?"/>`, `<xs:pattern value="\p{IsBasicLatin}+"/>`, "bloque Unicode"},
		{"enumeration en una fecha", td, fechaH, `<xs:enumeration value="2020-01-01T00:00:00"/>`, "faceta no soportada"},
		{"maxInclusive en una fecha", td, fechaH, fechaH + `<xs:maxInclusive value="2030-01-01T00:00:00Z"/>`, "faceta no soportada"},
		{"minInclusive de una fecha sin hora", ret10, `<xs:minInclusive value="2014-01-01T00:00:00-06:00"/>`, `<xs:minInclusive value="2014-01-01"/>`, "no es un xs:dateTime"},
		{"minInclusive de una fecha con zona mal escrita", ret10, `<xs:minInclusive value="2014-01-01T00:00:00-06:00"/>`, `<xs:minInclusive value="2014-01-01T00:00:00-6"/>`, "no es un xs:dateTime"},
		{"enumeration de un decimal que no es decimal", catPlat, `<xs:enumeration value="0.160000"/>`, `<xs:enumeration value="dieciseis"/>`, "no es un decimal"},
		{"import de Retenciones no previsto", ret20, `schemaLocation="../1/catalogos/catRetenciones.xsd"`, `schemaLocation="../1/catalogos/otro.xsd"`, "no previsto"},
		{"length en un decimal", td, `<xs:pattern value="[0-9]{1,18}(.[0-9]{1,6})?"/>`, `<xs:length value="1"/>`, "faceta no soportada"},
		{"whiteSpace preserve en un decimal", td, `<xs:minInclusive value="0.000000"/>
			<xs:whiteSpace value="collapse"/>`, `<xs:minInclusive value="0.000000"/>
			<xs:whiteSpace value="preserve"/>`, "whiteSpace"},
		{"whiteSpace desconocido", td, `<xs:whiteSpace value="collapse"/>`, `<xs:whiteSpace value="todo"/>`, "whiteSpace"},
		{"faceta con fixed", td, `<xs:pattern value="[0-9]{1,18}(.[0-9]{1,6})?"/>`, `<xs:fractionDigits value="6" fixed="true"/>`, "no soportado"},
		// xs:all (SPEI): elementos con maxOccurs 1 y nada más.
		{"xs:all con maxOccurs", "cfd/spei/spei.xsd", `<xs:element name="Beneficiario">`, `<xs:element name="Beneficiario" maxOccurs="2">`, "maxOccurs distinto de 1"},
		{"xs:all con secuencia", "cfd/spei/spei.xsd", `<xs:all>`, `<xs:all><xs:sequence/>`, "construcción no soportada"},
		{"xs:all repetible", "cfd/spei/spei.xsd", `<xs:all>`, `<xs:all maxOccurs="2">`, "no soportado"},
		{"xs:all vacío", "cfd/spei/spei.xsd", `<xs:all>`, `<xs:all/><xs:all>`, "vacío"},
		{"maxInclusive de una fecha mal escrito", "cfd/spei/spei.xsd", `<xs:maxInclusive value="9999-12-31"/>`, `<xs:maxInclusive value="9999-12-31T00:00:00"/>`, "no es un xs:date"},
		{"enumeration en un xs:time", "cfd/spei/spei.xsd", `<xs:restriction base="xs:time">`, `<xs:restriction base="xs:time"><xs:enumeration value="12:00:00"/>`, "faceta no soportada"},
		{"fractionDigits distinto de 0 en un entero", "cfd/ine/ine11.xsd", `<xs:fractionDigits value="0"/>`, `<xs:fractionDigits value="2"/>`, "en un tipo entero"},
		{"fractionDigits negativo", td, `<xs:pattern value="[0-9]{1,18}(.[0-9]{1,6})?"/>`, `<xs:fractionDigits value="-1"/>`, "no negativo"},
		{"minInclusive que no es decimal", td, `<xs:minInclusive value="0.000000"/>`, `<xs:minInclusive value="cero"/>`, "no es decimal"},
		{"enumeration con salto de línea", cat, `value="01"`, `value="0&#96;1"`, "no se puede guardar"},
		{"import no previsto", cfdi40, `schemaLocation="../catalogos/catCFDI.xsd"`, `schemaLocation="../catalogos/otro.xsd"`, "no previsto"},
		{"tipo sin definir", tfd, `type="tdCFDI:t_FechaH"`, `type="tdCFDI:t_Nada"`, "no definido"},
		{"prefijo sin declarar", tfd, `type="tdCFDI:t_FechaH"`, `type="zz:t_FechaH"`, "prefijo"},
		{"maxOccurs raro", cfdi40, `<xs:element name="CfdiRelacionados" minOccurs="0" maxOccurs="unbounded">`, `<xs:element name="CfdiRelacionados" minOccurs="0" maxOccurs="muchos">`, "maxOccurs"},
		{"maxOccurs menor que minOccurs", cfdi40, `<xs:element name="CfdiRelacionados" minOccurs="0" maxOccurs="unbounded">`, `<xs:element name="CfdiRelacionados" minOccurs="3" maxOccurs="2">`, "menor"},
		{"elemento repetido en la secuencia", cfdi40, seqEmisor, `<xs:element name="Emisor"><xs:complexType/></xs:element>` + seqEmisor, "repetido"},
		{"elementFormDefault", tfd, `elementFormDefault="qualified"`, `elementFormDefault="unqualified"`, "elementFormDefault"},
		{"fixed en un decimal", tfd, `<xs:attribute name="FechaTimbrado" use="required" type="tdCFDI:t_FechaH">`, `<xs:attribute name="FechaTimbrado" use="required" type="xs:decimal" fixed="1">`, "fixed"},
		{"texto suelto", tfd, `<xs:attribute name="Version"`, `hola<xs:attribute name="Version"`, "texto suelto"},
		{"appinfo", tfd, `<xs:documentation>Complemento requerido`, `<xs:appinfo/><xs:documentation>Complemento requerido`, "construcción no soportada"},
		// La variante TFD 1.0: su atributo selector tiene que ser requerido y
		// fixed, no estar en la raíz, y el elemento tiene que existir en ella.
		{"variante sin fixed", tfd10, `use="required" fixed="1.0"`, `use="required"`, "requerido y fixed"},
		{"variante opcional", tfd10, `use="required" fixed="1.0"`, `fixed="1.0"`, "requerido y fixed"},
		{"selector también en la raíz", tfd, `<xs:attribute name="Version"`, `<xs:attribute name="version"/><xs:attribute name="Version"`, "no distingue"},
		{"selector fixed con el mismo valor en la raíz", "cfd/donat/donat11.xsd", `use="required" fixed="1.1"`, `use="required" fixed="1.0"`, "no distingue"},
		{"variante sin elemento en las raíces", tfd10, `<xs:element name="TimbreFiscalDigital">`, `<xs:element name="Otro">`, "sin elemento global"},
	}
	for _, c := range casos {
		_, _, err := Generar(copiaSAT(t, c.archivo, c.viejo, c.nuevo))
		if err == nil || !strings.Contains(err.Error(), c.quiero) {
			t.Errorf("%s: error %v, quiero que diga %q", c.nombre, err, c.quiero)
		}
	}
	if _, _, err := Generar(t.TempDir()); err == nil {
		t.Error("un directorio vacío no dio error")
	}
}

func TestTraducirPatron(t *testing.T) {
	casos := []struct {
		xsd, re string
		si, no  []string
	}{
		{`[0-9]{5}`, `^(?:[0-9]{5})$`, []string{"42501"}, []string{"425010", "4250", "x42501"}},
		{`(.[0-9]{3})?`, `^(?:(?:[^\n\r][0-9]{3})?)$`, []string{"", ".123", "x123"}, []string{"\n123", "\r123"}},
		{`a^b$`, `^(?:a\^b\$)$`, []string{"a^b$"}, []string{"ab"}},
		{`[^(?!.*\s)-]{5,7}`, `^(?:[^(?!.*\x20\t\n\r)\-]{5,7})$`, []string{"ABCDE", "A1B2C3"}, []string{"AB DE", "AB-DE", "AB.DE", "ABCD"}},
		{`\d+`, `^(?:\p{Nd}+)$`, []string{"12", "١٢"}, []string{"a"}},
		{`\s\S`, `^(?:[\x20\t\n\r][^\x20\t\n\r])$`, []string{" a"}, []string{"\fa", "  "}},
		{`\w\W`, `^(?:[^\p{P}\p{Z}\p{C}][\p{P}\p{Z}\p{C}])$`, []string{"a."}, []string{".a"}},
		{`\p{Lu}\P{Lu}\D`, `^(?:\p{Lu}\P{Lu}\P{Nd})$`, []string{"Aaa"}, []string{"AAa", "Aa1"}},
		{`[\d\-a]`, `^(?:[\p{Nd}\-a])$`, []string{"-", "5", "a"}, []string{"b"}},
		{`[-a]`, `^(?:[\-a])$`, []string{"-"}, []string{"b"}},
		{`[a^]`, `^(?:[a\^])$`, []string{"^"}, []string{"b"}},
		{`\{\}\.\n\r\t\\\|\?\*\+\(\)\[\]\^\-`, `^(?:\{\}\.\n\r\t\\\|\?\*\+\(\)\[\]\^\-)$`, []string{"{}.\n\r\t\\|?*+()[]^-"}, nil},
		{`a|b*|c+?`, "", nil, nil}, // cuantificador doble: error
		{`XEXX010101000|[A-Z&Ñ]{3}`, `^(?:XEXX010101000|[A-Z&Ñ]{3})$`, []string{"XEXX010101000", "ÑA&"}, []string{"XEXX0101010000"}},
		{`(a)(b){2,}`, `^(?:(?:a)(?:b){2,})$`, []string{"abb", "abbb"}, []string{"ab"}},
	}
	for _, c := range casos {
		got, err := traducirPatron(c.xsd)
		if c.re == "" {
			if err == nil {
				t.Errorf("%q: se tradujo a %q y debía fallar", c.xsd, got)
			}
			continue
		}
		if err != nil || got != c.re {
			t.Errorf("%q: %q, %v; quiero %q", c.xsd, got, err, c.re)
			continue
		}
		re := regexp.MustCompile(got)
		for _, s := range c.si {
			if !re.MatchString(s) {
				t.Errorf("%q no acepta %q", c.xsd, s)
			}
		}
		for _, s := range c.no {
			if re.MatchString(s) {
				t.Errorf("%q acepta %q", c.xsd, s)
			}
		}
	}
	errores := []struct{ xsd, quiero string }{
		{`[a-z-[aeiou]]`, "resta de clases"},
		{`\p{IsBasicLatin}`, "bloque Unicode"},
		{`\p{Xx}`, "desconocida"},
		{`\p`, "sin {"},
		{`\p{L`, "sin cerrar"},
		{`\i`, "no soportado"},
		{`\c`, "no soportado"},
		{`[\S]`, `\S dentro`},
		{`[\w]`, `\w dentro`},
		{`[\W]`, `\W dentro`},
		{`a{1001}`, "límite"},
		{`a{2,1}`, "no soportado"},
		{`a{x}`, "no soportado"},
		{`a{1`, "sin cerrar"},
		{`(a`, "sin cerrar"},
		{`a)`, "sin abrir"},
		{`*a`, "sin átomo"},
		{`a]`, "suelto"},
		{`[a`, "sin cerrar"},
		{`[]`, "vacía"},
		{`[a[b]]`, "dentro de una clase"},
		{`a\`, "al final"},
		{`a**`, "doble"},
	}
	for _, e := range errores {
		if got, err := traducirPatron(e.xsd); err == nil || !strings.Contains(err.Error(), e.quiero) {
			t.Errorf("%q: %q, %v; quiero un error que diga %q", e.xsd, got, err, e.quiero)
		}
	}
}

func TestIdentificador(t *testing.T) {
	for in, want := range map[string]string{"Año": "Anio", "Antigüedad": "Antiguedad", "NumAñosServicio": "NumAniosServicio", "c_Moneda": "cMoneda", "ÑÁÉÍÓÚ": "NiAEIOU"} {
		if got := identificador(in); got != want {
			t.Errorf("identificador(%q) = %q, quiero %q", in, got, want)
		}
	}
}
