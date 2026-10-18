package cfdi

import _ "embed"

// Schema es el JSON Schema (draft 2020-12) del Documento: el mismo archivo
// que schema/cfdi.v1.schema.json, embebido para quien quiera validar o
// publicar la salida.
//
//go:embed schema/cfdi.v1.schema.json
var Schema []byte

// SchemaRetenciones es el JSON Schema del DocumentoRetenciones:
// schema/retenciones.v1.schema.json.
//
//go:embed schema/retenciones.v1.schema.json
var SchemaRetenciones []byte
