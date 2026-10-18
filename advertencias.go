package cfdi

import "strings"

// Advertencia describe algo raro que se encontró al leer un CFDI sin que
// impida leerlo. Codigo es estable entre versiones (sirve para agrupar en
// lote); Ruta señala el nodo o atributo; Mensaje es para personas y puede
// cambiar.
//
// # Formato de Ruta (contrato estable)
//
// Ruta es una ruta de nodos por nombre local, sin prefijos de namespace,
// desde la raíz: "/Comprobante/Conceptos/Concepto[2]/Impuestos/Traslados/Traslado[1]".
// Si la advertencia es de un atributo, termina en "/@Nombre" y Atributo
// lleva ese Nombre; Nodo() devuelve la ruta sin él.
//
//   - Los nodos que el esquema permite repetir llevan su posición, desde 1,
//     entre corchetes y siempre, también el primero: Concepto[1],
//     CfdiRelacionados[1], Pago[1], Traslado[1]…
//   - Los nodos únicos (Emisor, Impuestos, Totales…) van sin índice. Si un
//     nodo único viene repetido, la ruta es la del primero.
//   - Excepción: Complemento y los complementos que se acumulan (Nomina,
//     HidroYPetro) van sin índice el primero y con [n] desde el segundo
//     ("/Comprobante/Complemento/Nomina", ".../Nomina[2]").
//   - Las advertencias del documento entero (encoding, XML no estricto,
//     contenido tras la raíz) van con Ruta "".
//
// El nombre del atributo es el que pide el SAT ("Version", "UUID"), aunque
// el documento lo traiga en minúsculas (TFD 1.0, CFDI 3.2).
type Advertencia struct {
	Codigo  string `json:"codigo"`
	Ruta    string `json:"ruta,omitempty"`
	Mensaje string `json:"mensaje,omitempty"`
	// Atributo es el nombre del atributo de la advertencia ("" si es de un
	// nodo o del documento). Es lo que va después de "/@" en Ruta.
	Atributo string `json:"atributo,omitempty"`
	// Valor es el texto original del atributo, tal como venía, en las
	// advertencias de valores: DECIMAL_INVALIDO (el campo quedó vacío y este
	// es el único lugar con el texto), FECHA_INVALIDA, FECHA_FORMATO y
	// UUID_INVALIDO. En el JSON se omite si es "".
	Valor string `json:"valor,omitempty"`
}

// Nodo es la ruta del nodo de la advertencia: Ruta sin el "/@Atributo" final.
func (a Advertencia) Nodo() string {
	if a.Atributo == "" {
		return a.Ruta
	}
	return strings.TrimSuffix(a.Ruta, "/@"+a.Atributo)
}

// AdvertenciasDe devuelve las advertencias del nodo con esa ruta (ver el
// formato en Advertencia) y las de sus atributos, sin las de sus hijos. Con
// rutaNodo "" devuelve las del documento entero.
func (c *Comprobante) AdvertenciasDe(rutaNodo string) []Advertencia {
	var out []Advertencia
	for _, a := range c.Advertencias {
		if a.Nodo() == rutaNodo {
			out = append(out, a)
		}
	}
	return out
}

// AdvertenciaEn devuelve la primera advertencia del atributo de ese nodo:
// AdvertenciaEn("/Comprobante", "Fecha") en vez de buscar el sufijo
// "/@Fecha" en Ruta.
func (c *Comprobante) AdvertenciaEn(rutaNodo, atributo string) (Advertencia, bool) {
	for _, a := range c.Advertencias {
		if a.Atributo == atributo && a.Nodo() == rutaNodo {
			return a, true
		}
	}
	return Advertencia{}, false
}

// Códigos de advertencia. Son parte de la API pública: no se renombran.
const (
	// Codificación y forma del XML.
	AdvEncodingCorregido   = "ENCODING_CORREGIDO"   // los bytes no coinciden con el encoding declarado; se convirtieron
	AdvEncodingDesconocido = "ENCODING_DESCONOCIDO" // encoding declarado que no se reconoce; se leyó como UTF-8 o windows-1252
	AdvCaracterInvalido    = "CARACTER_INVALIDO"    // caracteres de control prohibidos en XML 1.0; se quitaron
	AdvXMLNoEstricto       = "XML_NO_ESTRICTO"      // el XML solo se pudo leer en modo no estricto (entidades o & sueltos)
	AdvContenidoTrasRaiz   = "CONTENIDO_TRAS_RAIZ"  // texto o basura después del cierre de cfdi:Comprobante; se ignoró

	// Estructura.
	AdvVersionNoSoportada     = "VERSION_NO_SOPORTADA" // Version que no es 3.0, 3.2, 3.3 ni 4.0
	AdvVersionAntigua         = "VERSION_ANTIGUA"      // CFDI 3.0 o 3.2: lectura parcial, solo lo que coincide con 3.3
	AdvVersionNamespace       = "VERSION_NAMESPACE"    // Version no corresponde al namespace (p. ej. 3.3 en cfd/4)
	AdvNamespaceInesperado    = "NAMESPACE_INESPERADO" // nodo con el nombre esperado pero en otro namespace; se leyó
	AdvNodoFaltante           = "NODO_FALTANTE"        // falta un nodo requerido (Emisor, Receptor, Conceptos…)
	AdvNodoDuplicado          = "NODO_DUPLICADO"       // nodo que debe ser único aparece varias veces; se usa el primero
	AdvAtributoFaltante       = "ATRIBUTO_FALTANTE"    // falta un atributo requerido
	AdvSinTimbre              = "SIN_TIMBRE"           // no hay TimbreFiscalDigital
	AdvComplementoNoSoportado = "COMPLEMENTO_NO_SOPORTADO"
	AdvComplementoVersion     = "COMPLEMENTO_VERSION" // versión de complemento desconocida o que no corresponde a su namespace

	// Valores.
	AdvDecimalInvalido  = "DECIMAL_INVALIDO"   // monto o cantidad que no es xs:decimal; el campo queda vacío
	AdvFechaInvalida    = "FECHA_INVALIDA"     // fecha ilegible; queda solo el texto original
	AdvFechaFormato     = "FECHA_FORMATO"      // fecha legible pero fuera del formato SAT (zona, fracción, sin hora…)
	AdvUUIDInvalido     = "UUID_INVALIDO"      // UUID que no tiene la forma 8-4-4-4-12 hexadecimal
	AdvExentoConImporte = "EXENTO_CON_IMPORTE" // traslado Exento con TasaOCuota o Importe; se conservan, pero no deberían venir
)

// Codigos lista todos los códigos de advertencia, en el orden de arriba.
var Codigos = []string{
	AdvEncodingCorregido, AdvEncodingDesconocido, AdvCaracterInvalido, AdvXMLNoEstricto, AdvContenidoTrasRaiz,
	AdvVersionNoSoportada, AdvVersionAntigua, AdvVersionNamespace, AdvNamespaceInesperado, AdvNodoFaltante, AdvNodoDuplicado,
	AdvAtributoFaltante, AdvSinTimbre, AdvComplementoNoSoportado, AdvComplementoVersion,
	AdvDecimalInvalido, AdvFechaInvalida, AdvFechaFormato, AdvUUIDInvalido, AdvExentoConImporte,
}

// TieneAdvertencia dice si la lectura produjo al menos una advertencia con
// ese código.
func (c *Comprobante) TieneAdvertencia(codigo string) bool {
	for _, a := range c.Advertencias {
		if a.Codigo == codigo {
			return true
		}
	}
	return false
}
