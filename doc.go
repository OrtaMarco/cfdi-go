// Package cfdi lee CFDI 3.3 y 4.0 (facturas electrónicas del SAT, México) y
// los normaliza a structs Go y a un JSON con esquema versionado.
//
// Lectura tolerante: Parse solo falla (con un *Error de Clase estable) si el
// documento no es XML o si la raíz no es un Comprobante de los namespaces
// http://www.sat.gob.mx/cfd/3 o /cfd/4. Lo demás —encodings que mienten,
// atributos con formato raro, complementos desconocidos, falta de timbre— se
// lee como se pueda y se reporta en Comprobante.Advertencias con códigos
// estables (las constantes Adv*).
//
// Montos: decimal.Decimal exacto, nunca float; en JSON van como string con el
// texto original del XML.
//
// Leer no valida: eso es aparte y sobre los bytes o el Comprobante ya leído.
// CadenaOriginal y Verificar arman la cadena y verifican los sellos;
// ValidarEsquema valida contra los XSD del SAT (el CFDI, el timbre y los
// complementos que lista el README), ValidarReglas revisa las reglas aritméticas y de
// coherencia, y Validar junta las dos. Cada regla de Reglas cita su
// fundamento en la matriz de errores, el estándar o la guía de llenado del
// SAT (Regla.Citas). Nada consulta al SAT ni usa la red.
//
// El CFDI de Retenciones e información de pagos (1.0 y 2.0) tiene su propia
// API: Detectar dice qué es un XML, ParseRetenciones lo lee (con su JSON y su
// SchemaRetenciones), CadenaOriginalRetenciones y VerificarRetenciones arman
// la cadena y verifican los sellos, y ValidarRetenciones lo valida contra los
// XSD. Parse sigue rechazando una Retenciones (no_es_comprobante).
//
//	c, err := cfdi.Parse(f)
//	if err != nil { ... }            // no es un CFDI
//	for _, a := range c.Advertencias { ... }
//	b, _ := c.JSON()                 // {"schema_version":"1","cfdi":{...},"advertencias":[...]}
package cfdi
