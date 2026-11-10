# cfdi-go — guía para contribuir (personas y agentes)

Librería en Go para leer, verificar y validar CFDI del SAT (México). Módulo
`github.com/OrtaMarco/cfdi-go`, paquete `cfdi`. El README explica el alcance y la API.

## Reglas que no se negocian

- **Cero datos reales en el repo**: ni XML, ni RFC, ni CURP, ni nombres, ni UUID de contribuyentes
  reales, en código, pruebas, fixtures, commits ni mensajes de commit. Los fixtures son ejemplos públicos
  (los del SAT, o los de phpcfdi/nodecfdi/CfdiUtils/satcfdi con su aviso MIT en `NOTICE`) o sintéticos
  con RFC genéricos o de prueba del SAT (`EKU9003173C9`, `XAXX010101000`, `XIQB891116QE4`…). Si un
  ejemplo público trae un RFC o un nombre que no es de prueba, se sustituye y se anota en `NOTICE`.
- Si pruebas la librería contra facturas reales, hazlo fuera del repo; a lo más se versionan conteos y
  clases de error, nunca extractos. `corpus/` y `corpus-reportes/` están en `.gitignore` para eso.
- **Montos sin float**: tipo decimal exacto propio (`decimal.Decimal`), que conserva el texto original;
  en JSON van como string.
- **Sin dependencias externas** en la librería ni en la CLI: solo la biblioteca estándar.
- Errores tolerantes: un CFDI raro devuelve lo que se pudo leer más advertencias; no aborta.
- `schema_version` en el JSON; el JSON Schema publicado vive en `schema/`.
- Directiva `go` conservadora (1.24) en `go.mod`, para que la puedan importar proyectos no tan nuevos.
- Las pruebas tienen que fallar si se rompe el código: comprueba con mutaciones a mano las reglas
  importantes (montos, versiones, complementos), no solo que «pasa».
- Commits en español, en el estilo del historial (`área: qué cambia`).

## Comandos

```bash
go test ./...
go vet ./...
gofmt -l .
go run ./cmd/cfdi parse factura.xml        # JSON a stdout
go run ./cmd/cfdi lote carpeta-o-zip/       # JSONL + resumen de fallos
```

Los archivos generados (`cadena_reglas.go`, `esquema_*.go`) salen de `go generate .`; los oráculos de
`testdata/` (cadenas, veredictos de esquema, goldens) se regeneran como dice la sección «Pruebas» del
README.

## Fuentes del SAT

`docs/fuentes.json` es la foto de las fuentes oficiales: los documentos que citan las reglas, las
páginas del portal que los listan y el commit de la réplica de la que se copian XSD y XSLT.
`python3 scripts/vigilar_fuentes.py` la compara con lo publicado hoy y el workflow `Fuentes del SAT` lo
corre cada lunes: si algo cambió o no se pudo comprobar, deja un issue con el informe. Las páginas del
portal son las que detectan una matriz o un catálogo nuevos, porque el SAT les pone la fecha en el
nombre y la URL vieja sigue respondiendo igual.

Que el informe diga «cambió» no arregla nada: hay que leer el documento nuevo, llevar a las reglas lo
que toque (con su cita) y solo entonces rehacer la foto con `--actualizar`. Al citar un documento nuevo
en `docs/reglas.md`, se añade también a `docs/fuentes.json`.
