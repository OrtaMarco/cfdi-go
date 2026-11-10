# cfdi-go

Librería en Go que **lee, verifica y valida** CFDI del SAT (México); **no genera ni timbra**. Lee CFDI 3.3
y 4.0 y CFDI de Retenciones e información de pagos 1.0 y 2.0 y los normaliza a structs y JSON con esquema
versionado; arma la cadena original, verifica los sellos y valida contra los XSD y las reglas del SAT.
Sin dependencias externas.

> Estado: v0.x. La API puede cambiar entre versiones menores hasta la v1.0.0.

```bash
go get github.com/OrtaMarco/cfdi-go@latest
go install github.com/OrtaMarco/cfdi-go/cmd/cfdi@latest   # la CLI
```

```go
import "github.com/OrtaMarco/cfdi-go" // el paquete se llama cfdi
```

Requiere Go 1.24 o más nuevo. Solo usa la biblioteca estándar.

**Versiones soportadas**, en resumen (el detalle está en [Alcance](#alcance)):

| Documento | Lectura y JSON | Cadena original y sello | Esquema (XSD) | Reglas |
|---|---|---|---|---|
| CFDI 4.0 | sí | sí | sí | sí |
| CFDI 3.3 | sí | sí | sí | sí |
| CFDI 3.2 / 3.0 | parcial | 3.2 | no | no |
| Timbre Fiscal Digital 1.0 / 1.1 | sí | sí | sí | — |
| Retenciones e información de pagos 1.0 / 2.0 | sí | sí | sí | no |

**Límites:** no timbra, no cancela, no genera ni firma CFDI (solo arma la cadena original y verifica
sellos), no consulta al SAT ni descarga certificados, no revisa revocación ni cadena de confianza de los
certificados y no valida contra catálogos que no vengan en los XSD. Que un CFDI pase `Validar` no
significa que el SAT lo acepte: la librería cubre lo que se puede revisar sin red. Ver
[Fuera de alcance](#alcance) y las reglas en [Validación](#validación).

*English summary:* a dependency-free Go library to parse Mexican CFDI e-invoices (3.3 and 4.0, plus
Retenciones 1.0/2.0 and most SAT complements) into structs and versioned JSON, build the original
string (cadena original), verify issuer and SAT seals, and validate against the SAT XSD schemas and
arithmetic/consistency rules. Documentation is in Spanish.

## Alcance

- CFDI 3.3 / 4.0: tipos I, E, P, N, T; PUE/PPD; CfdiRelacionados; impuestos por concepto y totales.
- Complementos: TimbreFiscalDigital 1.0 y 1.1, Pagos 1.0 y 2.0 (REP), Nómina 1.2 (lectura).
- Más complementos: ImpuestosLocales 1.0, CartaPorte 2.0 / 3.0 / 3.1, ComercioExterior 1.1 / 2.0,
  Divisas 1.0, CFDIRegistroFiscal 1.0, LeyendasFiscales 1.0, EstadoDeCuentaCombustible 1.2; de concepto:
  PorCuentadeTerceros 1.1, VentaVehiculos 1.0 / 1.1, HidroYPetro 1.0. Todos los nodos y atributos del XSD
  del SAT.
- Y también: Aerolíneas 1.0, Donatarias 1.0 / 1.1, INE 1.0 / 1.1, ValesDeDespensa 1.0,
  ConsumoDeCombustibles 1.0 / 1.1, PagoEnEspecie 1.0, NotariosPublicos 1.0, servicios parciales de
  construcción 1.0, VehiculoUsado 1.0, TuristaPasajeroExtranjero 1.0, SPEI de terceros, y las versiones
  viejas CartaPorte 1.0, ComercioExterior 1.0, EstadoDeCuentaCombustible 1.1 y Nómina 1.1 (de CFDI 3.2);
  de concepto: instEducativas (iedu) 1.0 y acreditamientoIEPS 1.0; y renovacionysustitucionvehiculos 1.0, certificadodedestruccion 1.0, obrasarteantiguedades 1.0,
  PFintegranteCoordinado (pfic) 1.0, GastosHidrocarburos 1.0, IngresosHidrocarburos 1.0,
  EstadoDeCuentaBancario (ecb) 1.0 y EstadoDeCuentaCombustible sin versión (ecc, de CFDI 3.2). Lo que no
  se lee sigue como `COMPLEMENTO_NO_SOPORTADO`.
- Detallista (sector retail, AMC8.1, sin versión): el único complemento del SAT que lleva datos en el
  texto de los elementos y no solo en atributos. Lectura, JSON, cadena y esquema, con sus 10 ejemplos
  públicos en CFDI 3.3 y 4.0.
- Sin leer todavía: psgecfd (solo CFDI 3.2, sin XSD ni ejemplos).
- Montos en decimal exacto (string en JSON).
- Cadena original (CFDI 3.2, 3.3, 4.0 y TFD 1.0 / 1.1) y verificación del sello del emisor y del
  timbre (ver abajo).
- Validación contra los XSD del SAT (CFDI 4.0 y 3.3, TFD 1.0 y 1.1 y 46 versiones de complementos, con
  sus catálogos; lista en [Esquema](#esquema-validaresquema)) y
  reglas aritméticas y de coherencia, cada una con su fundamento en la matriz de errores, el estándar o
  la guía de llenado del SAT (ver abajo).
- CLI: `cfdi parse` (un XML → JSON), `cfdi verificar` (cadena y sellos), `cfdi validar` (esquema y
  reglas) y `cfdi lote` (carpeta o ZIP → JSONL + resumen de fallos).
- CFDI de Retenciones e información de pagos 1.0 y 2.0, con API aparte: lectura con sus trece
  complementos, JSON con su propio esquema, cadena original, sellos y XSD (ver [Retenciones](#retenciones)).
- Chequeos del certificado del emisor sin red: vigencia en la fecha del documento, RFC del titular y
  número de serie (ver [Certificado del emisor](#certificado-del-emisor)).

Fuera de alcance por ahora: reglas que necesitan red o catálogos que no vienen en los XSD (RFC en la
lista del SAT, c_TasaOCuota, rangos de TipoCambio), cadena de confianza y revocación de los certificados,
reglas aritméticas de Retenciones, timbrado, cancelación, descarga masiva.

### CFDI 3.2 (y 3.0): lectura parcial

No son fatales: se leen los atributos y nodos que coinciden con 3.3 (los nombres de 3.2 van en
minúsculas, `total`, `fecha`, `rfc`; la lectura no distingue mayúsculas) y se advierte `VERSION_ANTIGUA`.
No se exige lo que introdujo 3.3 (`Moneda`, `UsoCFDI`, `RegimenFiscal` del emisor, `ClaveProdServ`,
`ClaveUnidad`) y no se modela lo exclusivo de 3.2: `formaDePago`, `metodoDePago`, domicilios,
`RegimenFiscal` como nodo, `tasa` de los traslados. Sus complementos sí (Nómina 1.1, por ejemplo). Los valores van tal cual vienen
(`tipo_de_comprobante` puede ser `"ingreso"`, `moneda` puede ser `"Pesos"`). Una `Version` que no sea
3.0, 3.2, 3.3 ni 4.0 se lee igual, según el namespace, con `VERSION_NO_SOPORTADA`.

## Uso

```go
c, err := cfdi.Parse(f) // err solo si no es XML o la raíz no es un Comprobante cfd/3 o cfd/4
if err != nil {
    log.Println(cfdi.ClaseDe(err)) // no_es_xml | xml_mal_formado | no_es_comprobante | lectura
}
for _, a := range c.Advertencias {
    fmt.Println(a.Codigo, a.Ruta, a.Mensaje)
}
fmt.Println(c.UUID(), c.Total.Original()) // "ABCDEF01-…", "1640.00"
b, _ := c.JSON()
```

```bash
cfdi parse factura.xml [-sello]
cfdi verificar factura.xml [-certs-sat dir-con-los-.cer-del-SAT]
cfdi validar factura.xml
cfdi lote carpeta/ otro.zip [-o salida.jsonl] [-resumen resumen.json] [-rutas] [-j N] [-sello | -verificar] [-certs-sat DIR] [-validar]
```

`verificar` escribe `{"archivo", "cadena_original", "sello", "cadena_original_tfd", "tfd", "certificado"}`
y sale con 0 solo si todo lo verificado es válido (los chequeos de `certificado` no cuentan para la salida,
como no cuentan para el sello). `lote -sello` incluye los sellos y agrega a cada línea
`verificacion_sello`; `-verificar` agrega la verificación sin los sellos. Con cualquiera de las dos el
resumen trae `sellos: {valido, invalido: {motivo: n}, no_verificable: {motivo: n}}`. `lote -certs-sat DIR`
verifica además el timbre (el SelloSAT con los certificados de `DIR`, como `verificar`): cada línea lleva
`verificacion_tfd` y el resumen `sellos_tfd`, con la misma forma.

`validar` escribe `{"archivo", "estado", "esquema": {estado, hallazgos}, "reglas": {estado, hallazgos}}`
(más `error` si el documento no se pudo leer como CFDI; entonces no hay `reglas`) y sale con 1 solo si es
inválido: un complemento sin esquema deja el estado `incompleto` y sale con 0. `lote -validar` agrega a
cada línea `validacion` con la misma forma, y al resumen `validacion: {por_estado, hallazgos_por_codigo}`
(conteos y huellas, como las advertencias).

`lote` recorre `.xml`, carpetas (recursivo) y `.zip` (también zips dentro de carpetas o de otros zips;
ignora `__MACOSX/` y `._*`). Escribe una línea por documento en el orden en que los encontró:

- leído: `{"origen":"c.zip!dir/a.xml","schema_version":"1","cfdi":{…},"advertencias":[…]}`
- Retenciones leída: `{"origen":"…","tipo_documento":"retenciones","schema_version":"1","retenciones":{…},"advertencias":[…]}`
- fatal: `{"origen":"…","error":{"clase":"no_es_xml","mensaje":"…"}}`

Las Retenciones se leen por omisión, sin bandera; la línea de un CFDI no lleva `tipo_documento` (es la
que trae `cfdi`). Con `-verificar`, `-certs-sat` y `-validar` se verifican y validan con su API
(`VerificarRetenciones`, `ValidarRetenciones`). En el resumen cuentan en `total`, `leidos`, `advertidos`,
advertencias, sellos y validación; `por_version`, `por_tipo` y `por_complemento` siguen siendo solo de CFDI,
y las Retenciones van aparte en `retenciones: {leidos, por_version, por_complemento}` (la clave no sale si
el lote no trae ninguna).

La memoria depende de `-j` y del tamaño de los XML, no de cuántos haya: cada worker deja de su CFDI
solo la línea JSON y lo que pide el resumen, y en espera del orden de salida hay a lo más `256×j`
líneas; si un documento lento retiene la cabeza, la lectura se detiene en vez de acumular. La Addenda
no se guarda en memoria (no se lee). Para apretar más en máquinas chicas: `GOMEMLIMIT=256MiB cfdi lote …`.

El resumen trae `total`, `leidos`, `fatales`, `advertidos`, conteos por código de advertencia y por clase
de error (con hasta 5 ejemplos), por versión, por tipo y por complemento. Está pensado para compartirse:
no incluye RFC, nombres, montos ni rutas. Cada ejemplo es una huella, los primeros 12 hex del sha256 del
contenido del XML (de la ruta, si ni se pudo leer) más la extensión del archivo de entrada o del zip más
externo: `3f2a9c1b4d5e.zip`. Se usa el contenido y no la ruta para que nadie pueda confirmar si un nombre
de archivo adivinado (con un RFC, p. ej.) está en el lote. Para ubicar un ejemplo: `sha256sum` sobre los
XML, o volver a correr con `-rutas`, que pone el origen completo (solo para uso local).

El JSONL, en cambio, siempre lleva `origen` con la ruta completa: es la salida del usuario, sobre sus
propios archivos, y sin la ruta no hay cómo volver al documento.

## JSON

```json
{"schema_version": "1", "cfdi": {"version": "4.0", "total": "1640.00", "conceptos": […], …}, "advertencias": []}
```

- Esquema: [`schema/cfdi.v1.schema.json`](schema/cfdi.v1.schema.json) (draft 2020-12), embebido como `cfdi.Schema`.
- Claves en snake_case de los nombres del SAT (`tipo_de_comprobante`, `uso_cfdi`); `Año` → `anio`,
  `Antigüedad` → `antiguedad`.
- Montos y cantidades: string con el texto original (`"0.160000"`, no `0.16`).
- Fechas: el texto original; en Go, `Fecha.Time` es la hora civil sin zona (ver `Fecha`).
- Vacíos omitidos. Sello, certificado y cadenas solo con `IncluirSello` / `-sello`.
- Contenedores sin atributos se aplanan (`conceptos`, `traslados`, `otros_pagos`); los que tienen
  atributos se conservan como objeto con la lista adentro (`percepciones.percepcion`, `pagos.pago`).
- UUID, `IdDocumento` y `CfdiRelacionado` en mayúsculas sin espacios; `*_original` solo si cambió.

### Complementos

Cada complemento soportado tiene su clave dentro de `cfdi`, junto a `timbre_fiscal_digital`, `pagos` y
`nomina`: `impuestos_locales`, `carta_porte`, `comercio_exterior`, `divisas`, `cfdi_registro_fiscal`,
`leyendas_fiscales`, `estado_de_cuenta_combustible`, `aerolineas`, `donatarias`, `ine`, `vales_de_despensa`,
`consumo_de_combustibles`, `pago_en_especie`, `notarios_publicos`, `parcialesconstruccion`, `vehiculo_usado`,
`turista_pasajero_extranjero`, `complemento_spei` (lista de SPEI_Tercero), `renovacionysustitucionvehiculos`,
`certificadodedestruccion`, `obrasarteantiguedades`, `pf_integrante_coordinado`, `gastos_hidrocarburos`,
`ingresos_hidrocarburos`, `estado_de_cuenta_bancario`. Los de concepto van en cada
concepto: `por_cuentade_terceros`, `venta_vehiculos`, `hidro_y_petro`, `inst_educativas`,
`acreditamiento_ieps`. Lo que no se lee queda en
`complementos_no_soportados` (y en `complemento_concepto` del concepto). Las claves salen del nombre del
SAT tal cual aunque junte palabras: `TotaldeTraslados` → `totalde_traslados`, `PorCuentadeTerceros` →
`por_cuentade_terceros`.

- Versiones de un mismo complemento van en un solo objeto con `version`: los XSD de CartaPorte 2.0/3.0/3.1
  y de ComercioExterior 1.1/2.0 solo agregan, quitan o mueven nodos, sin cambiar el tipo de los que
  comparten, así que la unión no pierde datos. Lo exclusivo de una versión (p. ej. `pedimentos` en
  CartaPorte 2.0, `documentacion_aduanera` en 3.x, `regimenes_aduaneros` en 3.1) viene vacío en las otras;
  el esquema lo dice en cada clave. La versión la dicta el namespace, como en Pagos.
- Lo mismo con las versiones viejas: ComercioExterior 1.0 agrega `receptor.curp` y
  `destinatario[].rfc`/`curp`; EstadoDeCuentaCombustible 1.1, `tar` y `no_identificacion` en cada concepto
  (su `TasaoCuota` va en `tasa_o_cuota`); Nómina 1.1 lleva los datos del trabajador como atributos de
  `Nomina` (`curp`, `num_empleado`, `tipo_regimen`… y `horas_extras`), no en `receptor`, y sus deducciones
  `importe_gravado`/`importe_exento`. CartaPorte 1.0 comparte casi todo con 2.0, pero 2.0 cambió la forma
  de tres nodos: lo de 1.0 va en `ubicaciones[].origen`/`destino`, `mercancias.autotransporte_federal` y
  `figura_transporte10` (en 1.0 FiguraTransporte es un nodo con `CveTransporte`, y `figura_transporte` es la
  lista de TiposFigura de 2.0+). Cada versión exige solo lo que pide su XSD.
- `impuestos_locales` no se suma a `impuestos`, pero sí está en el `total` del comprobante:
  `total = sub_total - descuento + trasladados - retenidos + totalde_traslados - totalde_retenciones`.
- HidroYPetro es complemento de concepto según su XSD; si viene en el `Complemento` del comprobante
  también se lee, en `cfdi.hidro_y_petro` (lista). Al revés, `parcialesconstruccion` va en el
  `Complemento` del comprobante, y si viene en un `ComplementoConcepto` se lee en el concepto.
- SPEI no tiene versión (`Complemento_SPEI` en `Complementos()`); su `sello` y `cadena_cda` solo salen con
  `IncluirSello`, como los demás sellos. ConsumoDeCombustibles 1.0 y 1.1 van en un objeto: 1.0 trae `tasa`
  en cada determinado y 1.1 `tasa_o_cuota` y `tipo_combustible`.
- Los de CFDI 3.2 sin versión en el atributo: EstadoDeCuentaCombustible del namespace `.../ecc` va en
  `estado_de_cuenta_combustible` con `version` "1.0" (la dicta el namespace) y su traslado trae `tasa`
  (porcentaje) en vez de `tasa_o_cuota`. En `estado_de_cuenta_bancario`, `movimientos` es un objeto con
  `movimiento_ecb` y `movimiento_ecb_fiscal`: el XSD los deja intercalados y cada lista conserva su orden.
- `renovacionysustitucionvehiculos`: un decreto de renovación (`decreto_renov_vehicular`, varios vehículos
  usados) o de sustitución (`decreto_sustit_vehicular`, uno, que exige además `num_fol_avisoint`,
  `num_ped_im`, `aduana` y `fecha_regul_veh`). El atributo `Año` va en `anio`. Si el CFDI trae dos nodos
  (el ejemplo del SAT trae uno de cada decreto) se lee el primero con `NODO_DUPLICADO`, como con los
  demás complementos.
- Detallista (sin versión, `detallista` en `Complementos()`) es el único que lleva datos en el texto de los
  elementos. Cada elemento complejo es un objeto (una lista si se repite) y cada elemento de contenido simple
  un campo del padre, con el nombre del SAT en snake_case. Los de contenido simple que además llevan
  atributos (`referenceIdentification`, `invoicedQuantity`, `lotNumber`…) son un objeto con sus atributos y
  el texto en `valor`. `shipment_detail` es `{}`: el XSD lo declara vacío y solo importa que exista. La
  cantidad adicional de la línea se llama `aditional_quantity` (sic: así viene en el XSD). `ship_to.name_and_address`
  es una lista de direcciones: en el XSD es una secuencia repetible (name, streetAddressOne, city,
  postalCode) y cada vuelta es una dirección; empieza una nueva cuando el campo siguiente no va después del
  anterior en esa secuencia. El texto libre (`name`, `long_text`, `text`…) va sin recortar, tal cual viene; los
  valores `xs:NMTOKEN` y de enumeración (`entity_type`, `currency_function`, `tax_type_description`,
  `tax_category`, `method_of_payment`, `special_services_type`) salen con los espacios colapsados, como los
  valida el esquema. Con XML inválido el modelo y la cadena pueden discrepar: en un elemento de texto que
  trae un hijo, el modelo lee su texto directo y la cadena el de todos sus descendientes (`string()` de
  XPath); `ValidarEsquema` ya marca ese documento.
- `Comprobante.Complementos()` (y el resumen de `lote`) cuenta también los de concepto.

## Advertencias

Códigos estables (constantes `cfdi.Adv*`, lista en `cfdi.Codigos`):

| Código | Qué pasó |
|---|---|
| `ENCODING_CORREGIDO` | el encoding declarado no coincide con los bytes (p. ej. dice UTF-8 y es latin1); se convirtió |
| `ENCODING_DESCONOCIDO` | encoding declarado que no se reconoce |
| `CARACTER_INVALIDO` | caracteres de control prohibidos en XML; se quitaron |
| `XML_NO_ESTRICTO` | `&` suelto o entidad desconocida; se leyó en modo tolerante |
| `CONTENIDO_TRAS_RAIZ` | basura después de `</cfdi:Comprobante>` |
| `VERSION_ANTIGUA` | CFDI 3.0 o 3.2: lectura parcial (ver arriba) |
| `VERSION_NO_SOPORTADA`, `VERSION_NAMESPACE` | `Version` rara o que no corresponde al namespace |
| `NAMESPACE_INESPERADO` | nodo con el nombre esperado pero en otro namespace; se leyó |
| `NODO_FALTANTE`, `NODO_DUPLICADO`, `ATRIBUTO_FALTANTE` | estructura incompleta o repetida |
| `SIN_TIMBRE` | no hay TimbreFiscalDigital |
| `COMPLEMENTO_NO_SOPORTADO`, `COMPLEMENTO_VERSION` | complemento que no se lee, o versión inesperada |
| `DECIMAL_INVALIDO`, `FECHA_INVALIDA`, `FECHA_FORMATO`, `UUID_INVALIDO` | valores ilegibles o fuera de formato |
| `EXENTO_CON_IMPORTE` | traslado Exento con TasaOCuota o Importe (se conservan) |

Por campo, sin parsear `Ruta`:

```go
if a, ok := c.AdvertenciaEn("/Comprobante", "SubTotal"); ok && a.Codigo == cfdi.AdvDecimalInvalido {
    fmt.Println("SubTotal ilegible:", a.Valor) // el texto original; c.SubTotal quedó nil
}
for _, a := range c.AdvertenciasDe("/Comprobante/Conceptos/Concepto[2]") { … } // el nodo y sus atributos
```

- `Atributo` (`atributo` en JSON) es el nombre del atributo cuando `Ruta` termina en `/@Nombre`;
  `Nodo()` es la ruta sin él. El formato de `Ruta` es contrato estable y está documentado en
  `Advertencia`: nombres locales sin prefijo, `[n]` desde 1 en los nodos repetibles (siempre, también el
  primero), sin índice en los únicos; `Complemento`, `Nomina` y `HidroYPetro` llevan `[n]` solo desde el
  segundo.
- `Valor` (`valor`) trae el texto original en `DECIMAL_INVALIDO`, `FECHA_INVALIDA`, `FECHA_FORMATO` y
  `UUID_INVALIDO`. Es la forma de recuperar un monto ilegible, porque el campo queda vacío.
- `Fecha.FueraDeFormato` marca por campo las fechas legibles pero fuera del formato SAT. Con
  `cfdi.ParseBytesOpciones(b, cfdi.Opciones{FechaEstricta: true})` esas fechas no se interpretan
  (`Valida()` da false) y la advertencia es `FECHA_INVALIDA`.

## Cadena original y sellos

```go
cadena, err := cfdi.CadenaOriginal(xml)      // ||4.0|A|1001|…||
tfd, err := cfdi.CadenaOriginalTFD(xml)      // ||1.1|UUID|…||
v := cfdi.Verificar(xml, cfdi.CertificadosEn(certsSAT...)) // nil: sin verificar el timbre
fmt.Println(v.Sello.Estado, v.Sello.Motivo)  // valido | invalido SELLO_NO_COINCIDE | no_verificable SIN_CERTIFICADO…
```

- Go no tiene XSLT: `internal/gencadena` traduce los XSLT oficiales del SAT (copiados sin cambios en
  `internal/gencadena/xslt/`) a `cadena_reglas.go`, que interpreta `cadena.go`. El generador falla ante
  cualquier construcción XSLT que no conoce; `go generate .` lo vuelve a correr y una prueba exige que
  `cadena_reglas.go` esté al día. Entiende `xsl:choose` cuando todos sus `xsl:when` comparan el
  mismo atributo con valores distintos y no hay `xsl:otherwise` (Nómina 1.1 elige por `@Version`): a lo
  más uno se cumple y se traduce como un `if` por cada `when`. Acepta además una plantilla repetida si es
  idéntica (la hoja de 3.2 incluye `ine10.xslt` e `ine11.xslt`, que lo son) y el `match="/"` propio de
  `vehiculousado.xslt`, porque el de la hoja raíz va después y XSLT usa el último; una repetida distinta,
  o un `match="/"` incluido después del de la raíz, siguen deteniendo el generador. Para Detallista
  entiende además `Requerido`/`Opcional` sobre el texto de un elemento (`select="."` dentro de un
  for-each, o una ruta simple como `detallista:buyer/detallista:gln`): como `string()` de XPath 1.0, el
  texto de todos los descendientes del primer nodo en orden de documento; `Opcional` sale si la ruta
  selecciona algo, aunque esté vacío, y `Requerido` deja `|` solo si no.
- La cadena se arma sobre los bytes, no sobre el `Comprobante`: orden de la hoja del SAT, nodos repetidos
  en orden de documento, `normalize-space`, opcionales presentes aunque vacíos.
- Complementos en la cadena, según lo que incluye la hoja del SAT de cada versión:

  | | 3.2 | 3.3 | 4.0 |
  |---|---|---|---|
  | TFD (no aporta), Nómina 1.2, implocal, Divisas, CFDIRegistroFiscal, LeyendasFiscales, ComercioExterior 1.1, VentaVehiculos 1.0/1.1, Aerolíneas, Donatarias 1.0/1.1, iedu, INE 1.0/1.1, ValesDeDespensa, PagoEnEspecie, NotariosPublicos, servicios parciales de construcción, VehiculoUsado, TuristaPasajeroExtranjero, renovacionysustitucionvehiculos, certificadodedestruccion, obrasarteantiguedades, pfic, detallista | sí | sí | sí |
  | PorCuentadeTerceros 1.1 | sí | sí | — |
  | EstadoDeCuentaCombustible 1.1, ConsumoDeCombustibles 1.0 | sí | sí | — |
  | ComercioExterior 1.0, Nómina 1.1, acreditamientoIEPS, SPEI, EstadoDeCuentaBancario, EstadoDeCuentaCombustible sin versión (ecc) | sí | — | — |
  | Pagos 1.0, CartaPorte 1.0 | — | sí | — |
  | CartaPorte 2.0, EstadoDeCuentaCombustible 1.2, ConsumoDeCombustibles 1.1, GastosHidrocarburos 1.0, IngresosHidrocarburos 1.0 | — | sí | sí |
  | Pagos 2.0, CartaPorte 3.0/3.1, ComercioExterior 2.0, HidroYPetro | — | — | sí |

  Cualquier otro complemento (o uno de la tabla en una versión cuya hoja no lo incluye) da
  `*ErrorCadena` con `COMPLEMENTO_SIN_CADENA`, nunca una cadena incompleta. También se niega con XML no
  estricto o cuyo encoding hubo que corregir (`XML_ILEGIBLE`).
- `Verificar`: RSA PKCS#1 v1.5 con SHA-256 (3.2 y TFD 1.0 aceptan también SHA-1) y `NoCertificado` igual
  al número de serie del certificado (`cfdi.NoCertificado`). El certificado del SAT del timbre no viene
  en el XML: lo da quien llama con un `CertificadosSAT` (la librería no usa la red). No revisa la cadena
  de confianza ni la revocación del certificado; la vigencia, el RFC y el número los revisa aparte, sin
  cambiar el veredicto del sello (ver abajo).
- Motivos estables en `cfdi.Motivos` (`NO_ES_RETENCIONES` es de las funciones de Retenciones); en
  el JSON, `verificacion_sello` es opcional (`schema_version` sigue en "1").

### Certificado del emisor

`Verificar` y `VerificarRetenciones` devuelven en `Verificacion.Certificado` (`certificado` en `cfdi
verificar`) tres chequeos del Certificado que trae el XML, cada uno con `estado` (`cumple`, `no_cumple`,
`no_verificable`), `motivo` estable y `mensaje`. Son aparte: un sello válido con un certificado vencido
sigue siendo `sello.estado: valido`. No usan la red: no dicen si el certificado estaba revocado ni si lo
emitió el SAT. `lote` no los escribe (sus líneas de CFDI no cambian).

- **Vigencia**: la `Fecha` del CFDI (`FechaExp` en Retenciones) entre `NotBefore` y `NotAfter`, los dos
  incluidos. La Fecha del CFDI y la FechaExp de Retenciones 2.0 son hora civil sin zona del lugar de
  expedición, y en México hay husos de UTC−5 (Quintana Roo) a UTC−8 (Baja California), contando el horario
  de verano que hubo hasta 2022. La librería no sabe cuál es (necesitaría el catálogo de códigos postales
  con su huso), así que la hora `h` es algún instante entre `h+5h` y `h+8h` UTC: cumple si cae en la
  vigencia con todos los husos, no cumple si no cae con ninguno, y si depende del huso (a menos de tres
  horas de un límite) es `no_verificable` con `VIGENCIA_AMBIGUA`. La FechaExp de Retenciones 1.0 trae su
  zona (`2020-06-30T18:00:00-06:00`): es un instante exacto, sin ambigüedad.
- **RFC**: el `x500UniqueIdentifier` (OID 2.5.4.45) del sujeto es el RFC del emisor (`Emisor/@Rfc`;
  `RfcE` o `RFCEmisor` en Retenciones), sin distinguir mayúsculas ni espacios. En una persona moral el
  atributo trae `RFC / RFC del representante legal` y cuenta el primero (`cfdi.RFCDelCertificado`).
  Excepción: si el titular es el SAT (`SAT970701NN3`, CN «SERVICIO DE ADMINISTRACION TRIBUTARIA») y no
  el emisor, es un CFDI del servicio gratuito de facturación del SAT, que sella con su certificado a
  nombre del contribuyente: `cumple` con motivo `CERTIFICADO_DEL_SAT`, el único `cumple` con motivo,
  para que se vea que el RFC no coincide y por qué no es un problema (un `no_verificable` diría que
  falta un dato, y no falta). Si el timbre trae `RfcProvCertif="SAT970701NN3"` el mensaje lo dice; no
  es requisito, porque los timbres 1.0 no tienen ese atributo. El SAT como representante (después de
  la barra) no cuenta, y un certificado de cualquier otro tercero sigue en `RFC_CERTIFICADO_DISTINTO`.
- **NoCertificado**: el del documento (`NumCert` en Retenciones 1.0) es el número de serie del
  certificado. Es lo mismo que exige el sello, pero aquí sale aunque la firma no verifique.

| Motivo | Chequeo | Estado |
|---|---|---|
| `CERTIFICADO_NO_VIGENTE_AUN` | vigencia | no_cumple: la fecha es anterior a `NotBefore` |
| `CERTIFICADO_VENCIDO` | vigencia | no_cumple: la fecha es posterior a `NotAfter` |
| `VIGENCIA_AMBIGUA` | vigencia | no_verificable: dentro o fuera según el huso |
| `FECHA_ILEGIBLE` | vigencia | no_verificable: sin fecha o ilegible |
| `RFC_CERTIFICADO_DISTINTO` | rfc | no_cumple |
| `CERTIFICADO_DEL_SAT` | rfc | cumple: el titular es el SAT, que sella a nombre del emisor |
| `CERTIFICADO_SIN_RFC` | rfc | no_verificable: el certificado no trae 2.5.4.45 |
| `SIN_RFC_EMISOR` | rfc | no_verificable: el documento no trae el RFC del emisor |
| `SIN_NOCERTIFICADO` | no_certificado | no_cumple: el documento no trae NoCertificado |
| `NOCERTIFICADO_DISTINTO` | no_certificado | no_cumple |
| `SIN_CERTIFICADO`, `CERTIFICADO_ILEGIBLE`, `XML_ILEGIBLE` | los tres | no_verificable |

La lista está en `cfdi.MotivosCertificado`.

## Validación

```go
r := cfdi.ValidarEsquema(xml)    // nunca falla
fmt.Println(r.Estado)            // valido | invalido | incompleto
for _, h := range r.Hallazgos {  // Codigo, Nivel, Ruta, Atributo, Valor, Regla, Mensaje
    fmt.Println(h.Codigo, h.Ruta, h.Regla) // ESQUEMA_CATALOGO /Comprobante/@Moneda enumeration
}
hs := cfdi.ValidarReglas(c)      // sobre el *Comprobante de Parse
v := cfdi.Validar(xml, c)        // las dos juntas (c puede ser nil): v.Estado, v.Esquema, v.Reglas
```

### Esquema (`ValidarEsquema`)

- Go no trae validador XSD: `internal/genxsd` traduce los XSD oficiales (copiados sin cambios en
  `internal/genxsd/xsd/`) a `esquema_reglas.go` (elementos, secuencias, minOccurs/maxOccurs, atributos,
  tipos y facetas) y `esquema_catalogos.go` (los `xs:enumeration`, un valor por línea). Como con la
  cadena, el generador falla ante cualquier construcción que no conoce (xs:group, xs:choice repetible
  dentro de una secuencia o anidado, extensiones con contenido propio, list, union, minExclusive, xs:any
  lax…), revisa también los tipos que hoy no se usan, y una prueba exige que lo generado esté al día
  (`go generate .`). Un `xs:choice` repetible que es todo el contenido (Movimientos del ECB) sí entra:
  sus ocurrencias son las vueltas de la secuencia, con una opción por vuelta.
- Esquemas: CFDI 4.0 y 3.3, TimbreFiscalDigital 1.1 y 1.0 y los complementos Pagos 2.0 y 1.0, Nómina 1.2,
  ImpuestosLocales, CartaPorte 3.1, 3.0 y 2.0, ComercioExterior 2.0 y 1.1, HidroYPetro, CFDIRegistroFiscal,
  Divisas, EstadoDeCuentaCombustible 1.2, PorCuentadeTerceros 1.1, VentaVehiculos 1.1, Aerolíneas,
  LeyendasFiscales, iedu y Donatarias 1.1. Terceros y VentaVehiculos traen un `xs:choice` (que no se
  repite) y tipos complejos con nombre; ComercioExterior, un catálogo sobre `xs:int`, que se compara
  por valor (`00` está en el catálogo `0`, `1`) como lo hace libxml2.
- Versiones viejas: CartaPorte 1.0, ComercioExterior 1.0 (con sus catálogos en línea, de 2016),
  EstadoDeCuentaCombustible 1.1, Nómina 1.1, y Donatarias y VentaVehiculos 1.0 como variantes (abajo). ECC
  1.1 trae un `minInclusive` sobre `xs:dateTime` sin zona (`Fecha` desde `2015-01-01T00:00:00`): sin zona
  contra sin zona se compara directo; con zona contra sin zona, con el margen de ±14 horas de abajo.
- De uso común: INE 1.1 (y 1.0 como variante por `Version="1.0"`: sus catálogos difieren, `QTO` solo
  está en 1.0 y `QUE` en 1.1), ValesDeDespensa, ConsumoDeCombustibles 1.0 y 1.1, PagoEnEspecie,
  NotariosPublicos, servicios parciales de construcción, VehiculoUsado, TuristaPasajeroExtranjero,
  acreditamientoIEPS y SPEI. Construcciones nuevas del generador, cada una con su prueba: `fractionDigits`
  0 en un `xs:int` (INE; en los enteros es fixed 0, otro valor se rechaza), y en SPEI `xs:all` (sus
  elementos en cualquier orden, cada uno a lo más una vez; el generador exige `maxOccurs` 1), `xs:time` y
  `minInclusive`/`maxInclusive` sobre `xs:date` (se compara como el `xs:dateTime` de las 00:00:00 del día).
- Decretos, hidrocarburos y estados de cuenta: renovacionysustitucionvehiculos, certificadodedestruccion, obrasarteantiguedades, pfic,
  GastosHidrocarburos10 e IngresosHidrocarburos10 (con catHidrocarburos), EstadoDeCuentaBancario (ecb) y
  EstadoDeCuentaCombustible sin versión (ecc). El generador acepta para Movimientos del ECB un `xs:choice`
  repetible que es todo el contenido. Las hojas del SAT los incluyen así: pfic, renovación, destrucción y
  obras de arte en 3.2, 3.3 y 4.0; gastos e ingresos de hidrocarburos en 3.3 y 4.0; ecb y ecc solo en 3.2.
  Como la librería no valida CFDI 3.2 contra XSD, el esquema de ecb y ecc solo aplica si vienen en un 3.3
  o 4.0; sus fixtures (`t10/ecb32.xml`, `t10/ecc32.xml`) prueban la lectura y la cadena.
- Detallista (`detallista.xsd`, incluido por las hojas de 3.2, 3.3 y 4.0), el único XSD del SAT con
  elementos de texto. Construcciones nuevas del generador, cada una con su prueba: elementos de tipo
  simple (anónimo o integrado: `xs:date`, `xs:decimal`) y `xs:simpleContent` con atributos, cuyo texto
  se junta (aunque lo partan comentarios o CDATA) y se valida con el whiteSpace y las facetas del tipo,
  sin hijos (`Hallazgo.Valor` trae el texto y `Atributo` va vacío); `mixed` en un contenido de elementos
  (specialInstruction; con simpleContent no cambia nada, como en libxml2); `xs:NMTOKEN` (uno o más
  NameChar, colapsado; su catálogo y su fixed comparan texto); `xs:appinfo` vacío; `default` en un
  atributo opcional, que no cambia lo válido; y una secuencia con `minOccurs="0"` solo si nada en ella es
  requerido (el validador no cuenta vueltas). La secuencia repetible de `shipTo/nameAndAddress` y los
  `xs:all` de la línea ya se soportaban.
- Retenciones 2.0 y 1.0 (`retencionpagov2.xsd`, `retencionpagov1.xsd`, con catRetenciones) y sus trece
  complementos: Plataformas Tecnológicas 1.0, Intereses, Pagos a extranjeros, Dividendos, Enajenación de
  acciones, Arrendamiento en fideicomiso, Operaciones con derivados, Intereses hipotecarios, Fideicomiso
  no empresarial, Planes de retiro 1.0 y 1.1, Premios y Sector financiero. Traen dos construcciones nuevas
  para el generador: un catálogo sobre `xs:decimal` (`c_TasaCuota` de Plataformas, por valor: `0.16` está
  en el catálogo `0.160000`) y un `minInclusive` sobre `xs:dateTime` (FechaExp de 1.0 desde
  `2014-01-01T00:00:00-06:00`), que se compara como instante con la zona; una fecha sin zona contra el
  límite se decide solo si queda del mismo lado con cualquier zona de −14:00 a +14:00, como el orden
  parcial de XML Schema. Plataformas Tecnológicas 1.1 y Intereses de Retenciones 2.0 no están en las
  fuentes: quedan `ESQUEMA_NO_SOPORTADO`.
- **TFD 1.0 y 1.1**: los dos XSD (`TimbreFiscalDigital.xsd` y `TimbreFiscalDigitalv11.xsd`) declaran el
  mismo elemento en el mismo namespace. libxml2 no puede elegir entre ellos (carga un esquema por
  namespace), `ValidarEsquema` sí, elemento por elemento: un `TimbreFiscalDigital` con `version="1.0"`
  (minúsculas, el atributo fixed del XSD de 1.0) se valida con el de 1.0; con `Version="1.1"`, sin
  versión o con cualquier otro valor, con el de 1.1. El TFD 1.0 es una *variante* en
  `internal/genxsd.Variantes`: el generador exige que su atributo selector sea requerido y fixed y que
  no exista en el XSD de 1.1. Así, una Retenciones 1.0 o un CFDI 3.3 con TFD 1.0 bien formado es
  `valido`. El oráculo de libxml2 lo imita con un segundo envoltorio que trae el XSD de 1.0 en ese
  namespace y valida con él los documentos cuyos timbres traen `version="1.0"`.
- También son variantes Donatarias 1.0 (`donat.xsd`) y VentaVehiculos 1.0
  (`ventavehiculos.xsd`), con `version="1.0"`, e INE 1.0 (`ine10.xsd`, con `Version="1.0"`). Ahí el
  selector sí existe en la raíz, pero fixed con otro valor (`1.1`), y eso basta: el generador lo acepta
  solo si los dos valores fixed difieren. Los tipos con
  nombre de una variante son suyos (VentaVehiculos 1.0 y 1.1 definen los dos `t_InformacionAduanera` en
  el mismo namespace). Donatarias 1.0 fija el texto de `leyenda`; un 1.0 con otra leyenda es
  `ESQUEMA_VALOR_FIJO`. El oráculo arma un envoltorio por combinación de variantes, y si un documento
  mezcla versiones (VentaVehiculos 1.0 y 1.1 en dos conceptos) valida con los dos y de cada uno descarta
  los errores dentro de los elementos de la otra versión, como hace Go elemento por elemento.
- Los patrones XSD se traducen a RE2: anclaje implícito (`^(?:…)$`), `^` y `$` literales, `.` sin `\n`
  ni `\r`, `\s` = `[ \t\n\r]`, `\d` = `\p{Nd}`, `\w`/`\W` por categorías Unicode. La resta de clases,
  los bloques `\p{Is…}`, `\i`/`\c` y los cuantificadores de más de 1000 detienen la generación (ninguno
  aparece hoy en los XSD del SAT). Detalle en `internal/genxsd/patron.go`.
- Valida como libxml2: whiteSpace del tipo antes de las facetas, `fractionDigits` y `totalDigits` sobre el
  valor (`2.5000000` tiene un decimal), fechas y enteros con su forma y rango, texto prohibido también en
  los elementos vacíos (un espacio en `<cfdi:Emisor> </cfdi:Emisor>` es error), `xsi:schemaLocation`
  permitido en cualquier nodo.
- Los `xs:any` del SAT (Complemento, ComplementoConcepto, Addenda) son estrictos: un elemento ahí sin
  esquema en la librería no se valida a medias, queda un aviso `ESQUEMA_NO_SOPORTADO` con su ruta y el
  estado es `incompleto` (si no hay errores). Lo mismo con una raíz que no es CFDI 3.3/4.0 (3.2 incluido)
  ni Retenciones 1.0/2.0.
- Lee el XML como un procesador conforme, sin corregir: lo que `Parse` arregla con una advertencia de
  encoding, caracteres prohibidos o XML no estricto aquí es `ESQUEMA_XML_ILEGIBLE`.
- Códigos (`cfdi.CodigosEsquema`): `ESQUEMA_XML_ILEGIBLE`, `ESQUEMA_NO_SOPORTADO` (aviso),
  `ESQUEMA_ELEMENTO_NO_PERMITIDO`, `ESQUEMA_ELEMENTO_FALTANTE` (minOccurs), `ESQUEMA_ELEMENTO_EXCEDENTE`
  (maxOccurs), `ESQUEMA_ORDEN`, `ESQUEMA_CONTENIDO`, `ESQUEMA_ATRIBUTO_FALTANTE`,
  `ESQUEMA_ATRIBUTO_NO_PERMITIDO`, `ESQUEMA_TIPO`, `ESQUEMA_PATRON`, `ESQUEMA_CATALOGO`, `ESQUEMA_RANGO`,
  `ESQUEMA_DIGITOS`, `ESQUEMA_LONGITUD`, `ESQUEMA_VALOR_FIJO`. `Regla` dice la faceta o restricción
  (`pattern`, `enumeration`, `minInclusive`, `fractionDigits`, `minOccurs`, `orden`, `required`…).

### Reglas (`ValidarReglas`)

Lo que se puede comprobar sin red ni catálogos externos, en CFDI 3.3 y 4.0, con decimal exacto: importe
de cada concepto y de cada impuesto con los límites del Anexo 20, SubTotal, Descuento, Total (con
implocal), totales de impuestos y grupos por impuesto/factor/tasa, decimales y TipoCambio según
`c_Moneda`, reglas por TipoDeComprobante, ObjetoImp, el comprobante de pago, Pagos 1.0 y 2.0, Comercio
Exterior 2.0 y Nómina 1.2. La lista está en `cfdi.Reglas`; cada regla lleva en `Regla.Citas` su
fundamento en los documentos oficiales del SAT: el código de la matriz de errores (`CFDI40167`,
`CRP20217`, `NOM37`…) y la página del estándar o de la guía de llenado. `Hallazgo.Regla` trae ese texto.

El detalle está en [docs/reglas.md](docs/reglas.md): niveles (error, advertencia y aviso), las reglas
con dos niveles y su margen, vigencia por fecha, lo que solo pide la guía de llenado, las decisiones de
Nómina 1.2 y Comercio Exterior 2.0, lo que queda fuera y la lista de documentos del SAT citados.

<details>
<summary>Las reglas, con su fundamento</summary>

| Código | Nivel | Qué revisa | Fundamento (SAT) |
|---|---|---|---|
| `REGLA_IMPORTE_CONCEPTO` | error | Importe del concepto entre (Cantidad − ½·10^-dC)·(ValorUnitario − ½·10^-dV) truncado y (Cantidad + ½·10^-dC − 10^-12)·(ValorUnitario + ½·10^-dV − 10^-12) redondeado hacia arriba, a los decimales del propio Importe | Matriz CFDI 4.0: CFDI40167; Anexo 20 4.0: p. 76; Matriz CFDI 3.3: CFDI33149; Anexo 20 3.3: p. 68 |
| `REGLA_DESCUENTO_CONCEPTO` | error | Descuento del concepto menor o igual que su Importe | Matriz CFDI 4.0: CFDI40169; Guía 4.0: p. 24; Matriz CFDI 3.3: CFDI33151 |
| `REGLA_IMPORTE_IMPUESTO` | error | Importe del impuesto del concepto entre (Base − ½·10^-dB)·TasaOCuota truncado y (Base + ½·10^-dB − 10^-12)·TasaOCuota redondeado hacia arriba, a los decimales del propio Importe. **Dos niveles**: fuera de eso pero dentro de los límites a los decimales de la moneda, advertencia (un PAC lo timbra) | Matriz CFDI 4.0: CFDI40180, CFDI40186; Anexo 20 4.0: p. 77, 78; Matriz CFDI 3.3: CFDI33161, CFDI33169; Anexo 20 3.3: p. 69, 70 |
| `REGLA_EXENTO_CON_IMPORTE` | error | traslado Exento sin TasaOCuota ni Importe | Matriz CFDI 4.0: CFDI40177; Anexo 20 4.0: p. 77; Matriz CFDI 3.3: CFDI33157 |
| `REGLA_TASA_SIN_IMPORTE` | error | traslado Tasa o Cuota con TasaOCuota e Importe | Matriz CFDI 4.0: CFDI40178; Anexo 20 4.0: p. 77; Matriz CFDI 3.3: CFDI33158 |
| `REGLA_RETENCION_EXENTA` | error | retención con TipoFactor distinto de Exento | Matriz CFDI 4.0: CFDI40184; Anexo 20 4.0: p. 78; Matriz CFDI 3.3: CFDI33166 |
| `REGLA_OBJETO_IMP` | error | 4.0: ObjetoImp 02 con nodo Impuestos; 01, 03, 04 y 05 sin él; 06 y 08 sin IVA ni IEPS; 07 sin IVA y con traslado de IEPS | Matriz CFDI 4.0: CFDI40171, CFDI40172, CFDI140226, CFDI140227; Anexo 20 4.0: p. 76 |
| `REGLA_SUBTOTAL` | error | I, E y N: SubTotal igual a la suma de los Importe de los conceptos o a su redondeo | Matriz CFDI 4.0: CFDI40108; Guía 4.0: p. 8; Matriz CFDI 3.3: CFDI33107; Anexo 20 3.3: p. 64 |
| `REGLA_DESCUENTO` | error | Descuento igual a la suma de los descuentos de los conceptos o a su redondeo | Matriz CFDI 4.0: CFDI40111; Guía 4.0: p. 8; Matriz CFDI 3.3: CFDI33110 |
| `REGLA_DESCUENTO_MAYOR_SUBTOTAL` | error | Descuento menor o igual que SubTotal | Matriz CFDI 4.0: CFDI40110; Guía 4.0: p. 8; Matriz CFDI 3.3: CFDI33109 |
| `REGLA_TOTAL` | error | Total = SubTotal − Descuento + trasladados − retenidos ± locales (implocal), redondeado a la moneda | Matriz CFDI 4.0: CFDI40119; Anexo 20 4.0: p. 72; Guía 4.0: p. 10; Matriz CFDI 3.3: CFDI33118 |
| `REGLA_TOTAL_TRASLADOS` | error | TotalImpuestosTrasladados igual a la suma de Impuestos/Traslados o a su redondeo. **Dos niveles**: fuera de eso pero entre la suma de cada Importe truncado y la de cada uno redondeado hacia arriba, advertencia (un PAC lo timbra) | Matriz CFDI 4.0: CFDI40205; Anexo 20 4.0: p. 81; Matriz CFDI 3.3: CFDI33183 |
| `REGLA_TOTAL_RETENCIONES` | error | TotalImpuestosRetenidos igual a la suma de Impuestos/Retenciones o a su redondeo. **Dos niveles**: como `REGLA_TOTAL_TRASLADOS` | Matriz CFDI 4.0: CFDI40203; Anexo 20 4.0: p. 81; Matriz CFDI 3.3: CFDI33181 |
| `REGLA_IMPUESTOS_SIN_TOTAL` | error | con traslados (no exentos) va TotalImpuestosTrasladados; con retenciones, TotalImpuestosRetenidos | Matriz CFDI 4.0: CFDI40212, CFDI40206, CFDI40209; Matriz CFDI 3.3: CFDI33190, CFDI33184, CFDI33187 |
| `REGLA_TRASLADO_SUMA` | error | cada Traslado del comprobante (Impuesto, TipoFactor, TasaOCuota) con Importe (y Base en 4.0) igual a la suma de los conceptos o a su redondeo. **Dos niveles**: fuera de eso pero dentro del margen de redondeo por concepto (cada Importe escrito o Base × TasaOCuota, truncado o redondeado hacia arriba antes de sumar), advertencia (un PAC lo timbra) | Matriz CFDI 4.0: CFDI40221, CFDI40215, CFDI40216; Guía 4.0: p. 90; Matriz CFDI 3.3: CFDI33195; Guía 3.3: p. 80 |
| `REGLA_TRASLADO_FALTANTE` | error | cada grupo de traslados de los conceptos está en el comprobante; los Exento solo en 4.0 | Matriz CFDI 4.0: CFDI40221, CFDI40213; Guía 4.0: p. 90; Matriz CFDI 3.3: CFDI33195; Guía 3.3: p. 32 |
| `REGLA_TRASLADO_SOBRANTE` | error | el comprobante no tiene grupos de traslados que no vengan de los conceptos | Matriz CFDI 4.0: CFDI40221; Matriz CFDI 3.3: CFDI33195 |
| `REGLA_RETENCION_SUMA` | error | cada Retencion del comprobante (Impuesto) igual a la suma de los conceptos o a su redondeo. **Dos niveles**: como `REGLA_TRASLADO_SUMA` | Matriz CFDI 4.0: CFDI40211; Matriz CFDI 3.3: CFDI33189 |
| `REGLA_RETENCION_FALTANTE` | error | cada impuesto retenido en los conceptos está en el comprobante | Matriz CFDI 4.0: CFDI40211; Matriz CFDI 3.3: CFDI33189 |
| `REGLA_RETENCION_SOBRANTE` | error | el comprobante no tiene retenciones que no vengan de los conceptos | Matriz CFDI 4.0: CFDI40211; Matriz CFDI 3.3: CFDI33189 |
| `REGLA_IMPLOCAL_TOTALES` | error | TotaldeTraslados y TotaldeRetenciones iguales a la suma de los Importe locales | Estándar ImpuestosLocales: p. 2 (sin matriz de errores) |
| `REGLA_DECIMALES_MONEDA` | error | SubTotal, Descuento, Total y totales de impuestos con no más decimales que la moneda | Matriz CFDI 4.0: CFDI40107, CFDI40112, CFDI40202, CFDI40204; Anexo 20 4.0: p. 72; Matriz CFDI 3.3: CFDI33106, CFDI33111, CFDI33180, CFDI33182 |
| `REGLA_DECIMALES_IMPUESTO` | error | 4.0: Base e Importe de los traslados e Importe de las retenciones del comprobante con no más decimales que la moneda | Matriz CFDI 4.0: CFDI40210, CFDI40214, CFDI40220; Anexo 20 4.0: p. 81, 82; Guía 4.0: p. 38, 40 |
| `REGLA_DECIMALES_IMPUESTO_33` | advertencia | 3.3: lo mismo en los traslados y retenciones del comprobante | Matriz CFDI 3.3: CFDI33188, CFDI33194 (la matriz pide no aplicarlos); Guía 3.3: p. 32, 34 |
| `REGLA_TIPOCAMBIO_MXN` | error | con Moneda MXN, TipoCambio ausente o 1 | Matriz CFDI 4.0: CFDI40114; Matriz CFDI 3.3: CFDI33113 |
| `REGLA_TIPOCAMBIO_XXX` | error | con Moneda XXX, sin TipoCambio | Matriz CFDI 4.0: CFDI40116; Matriz CFDI 3.3: CFDI33115 |
| `REGLA_TIPOCAMBIO_REQUERIDO` | error | con otra moneda, TipoCambio requerido | Matriz CFDI 4.0: CFDI40115; Matriz CFDI 3.3: CFDI33114 |
| `REGLA_TIPOCOMP_CONDICIONES` | error | T, P y N sin CondicionesDePago | Anexo 20 4.0: p. 73; Guía 4.0: p. 11; Anexo 20 3.3: p. 66 |
| `REGLA_TIPOCOMP_IMPUESTOS` | error | T, P y N sin nodo Impuestos | Matriz CFDI 4.0: CFDI40201; Anexo 20 4.0: p. 73; Matriz CFDI 3.3: CFDI33179; Anexo 20 3.3: p. 66 |
| `REGLA_TIPOCOMP_FORMAPAGO` | error | T y P sin FormaPago; en 4.0 también N | Matriz CFDI 4.0: CFDI40103; Anexo 20 4.0: p. 71; Anexo 20 3.3: p. 66 |
| `REGLA_TIPOCOMP_METODOPAGO` | error | T y P sin MetodoPago | Matriz CFDI 4.0: CFDI40125; Matriz CFDI 3.3: CFDI33123 |
| `REGLA_TIPOCOMP_DESCUENTO` | error | T y P sin Descuento, ni en el comprobante ni en los conceptos | Matriz CFDI 4.0: CFDI40111; Anexo 20 4.0: p. 73; Matriz CFDI 3.3: CFDI33110 |
| `REGLA_TIPOCOMP_SUBTOTAL_CERO` | error | T y P con SubTotal 0 | Matriz CFDI 4.0: CFDI40109; Matriz CFDI 3.3: CFDI33108 |
| `REGLA_TIPOCOMP_TOTAL_CERO` | error | T y P con Total 0 | Anexo 20 4.0: p. 72; Guía 4.0: p. 10 |
| `REGLA_TIPOCOMP_VALOR_UNITARIO` | error | I, E y N con ValorUnitario mayor que cero | Matriz CFDI 4.0: CFDI40166; Guía 4.0: p. 23; Matriz CFDI 3.3: CFDI33147 |
| `REGLA_TIPOCOMP_MONEDA_NOMINA` | error | N con Moneda MXN | Anexo 20 4.0: p. 71; Anexo 20 3.3: p. 65 |
| `REGLA_PAGO_COMPROBANTE` | error | tipo P: Moneda XXX, UsoCFDI CP01 (4.0) o P01 (3.3), Exportacion 01, un solo concepto 84111506 / 1 / ACT / Pago / 0 / 0 / ObjetoImp 01, sin NoIdentificacion, Unidad ni nodos hijo | Matriz Pagos 2.0: CRP20102, CRP20104, CRP20111-20120, CRP20122; Guía Pagos 2.0: p. 9, 17, 18; Matriz Pagos 1.0: CRP103, CRP110-120; Guía Pagos 1.0: p. 8, 13-15 |
| `REGLA_PAGO_SALDO_INSOLUTO` | error | ImpSaldoInsoluto = ImpSaldoAnt − ImpPagado | Matriz Pagos 2.0: CRP20244; Guía Pagos 2.0: p. 42; Matriz Pagos 1.0: CRP226; Guía Pagos 1.0: p. 28 |
| `REGLA_PAGO_EQUIVALENCIA` | error | Pagos 2.0: EquivalenciaDR 1 con la misma moneda; requerida con otra | Matriz Pagos 2.0: CRP20237, CRP20238; Guía Pagos 2.0: p. 40 |
| `REGLA_PAGO_TIPOCAMBIODR` | error | Pagos 1.0: TipoCambioDR requerido con otra moneda; ausente con la misma | Matriz Pagos 1.0: CRP218, CRP219; Guía Pagos 1.0: p. 26 |
| `REGLA_PAGO_MONTO` | error | Monto ≥ suma de ImpPagado / EquivalenciaDR (TipoCambioDR en 1.0) truncada; con otra moneda basta Monto ≥ suma de los límites inferiores truncados | Matriz Pagos 2.0: CRP20217, CRP20275-20277; Guía Pagos 2.0: p. 31-34; Matriz Pagos 1.0: CRP206; Guía Pagos 1.0: p. 20 |
| `REGLA_PAGO_TIPOCAMBIO` | error | MonedaP distinta de MXN con TipoCambioP; MXN con TipoCambioP 1 (en 1.0 también puede omitirse) | Matriz Pagos 2.0: CRP20214, CRP20215; Guía Pagos 2.0: p. 30; Matriz Pagos 1.0: CRP203, CRP204; Guía Pagos 1.0: p. 18, 43 |
| `REGLA_PAGOS_TOTALES` | error | Pagos 2.0: cada atributo de Totales igual al redondeo de la suma de los ImpuestosP (y Monto) × TipoCambioP; con ImpuestosP de IVA o retenciones, su atributo en Totales. **Dos niveles**: fuera de eso pero entre la suma de cada producto truncado y la de cada uno redondeado hacia arriba a 2 decimales, advertencia | Matriz Pagos 2.0: CRP20201-20211, CRP20264, CRP20267, CRP20273; Estándar Pagos 2.0: p. 52-57 (sin tolerancia); Guía Pagos 2.0: p. 21, 27 |
| `REGLA_PAGO_MONEDA_XXX` | error | MonedaP y MonedaDR distintas de XXX | Matriz Pagos 2.0: CRP20213, CRP20236; Matriz Pagos 1.0: CRP202, CRP217 |
| `REGLA_PAGO_MAYOR_CERO` | error | Monto, ImpSaldoAnt e ImpPagado mayores que cero | Matriz Pagos 2.0: CRP20218, CRP20239, CRP20241; Matriz Pagos 1.0: CRP207, CRP221, CRP223 |
| `REGLA_PAGO_DECIMALES` | error | Monto con los decimales de MonedaP; saldos e ImpPagado con los de MonedaDR | Matriz Pagos 2.0: CRP20219, CRP20240, CRP20242, CRP20243; Matriz Pagos 1.0: CRP208, CRP222, CRP224, CRP225 |
| `REGLA_PAGO_FORMAPAGO` | error | FormaDePagoP distinta de 99 | Matriz Pagos 2.0: CRP20212; Matriz Pagos 1.0: CRP201 |
| `REGLA_PAGO_CUENTAS` | error | cuentas y RFC ordenante/beneficiario y TipoCadPago según la forma de pago; NomBancoOrdExt con XEXX010101000 | Matriz Pagos 2.0: CRP20221, CRP20223, CRP20224, CRP20227-20229; Matriz Pagos 1.0: CRP211, CRP212, CRP214-216, CRP238 |
| `REGLA_PAGO_CADENA` | error | CertPago, CadPago y SelloPago van los tres con TipoCadPago y ninguno sin él | Matriz Pagos 2.0: CRP20230-20235; Matriz Pagos 1.0: CRP227-232 |
| `REGLA_PAGO10_PARCIALIDADES` | error | Pagos 1.0: con PPD, NumParcialidad, ImpSaldoAnt e ImpSaldoInsoluto; ImpPagado con varios documentos o con TipoCambioDR | Matriz Pagos 1.0: CRP233-236; Guía Pagos 1.0: p. 27, 28 |
| `REGLA_PAGO10_IMPUESTOS` | error | Pagos 1.0: el Pago no lleva nodo Impuestos | Matriz Pagos 1.0: CRP237; Guía Pagos 1.0: p. 29 |
| `REGLA_PAGO_OBJETO_IMP_DR` | error | Pagos 2.0: ObjetoImpDR 02 con ImpuestosDR; 01, 03, 04 y 05 sin él; 06 y 08 sin IVA ni IEPS; 07 sin IVA y con TrasladoDR de IEPS | Matriz Pagos 2.0: CRP20246, CRP20247, CRP20278, CRP20279; Estándar Pagos 2.0: p. 61; Guía Pagos 2.0: p. 43 |
| `REGLA_PAGO_IMPUESTO_DR` | error | Pagos 2.0: ImpuestosDR con al menos un traslado o una retención; BaseDR mayor que cero; RetencionDR distinta de Exento; TrasladoDR Exento sin TasaOCuotaDR ni ImporteDR, Tasa o Cuota con los dos | Matriz Pagos 2.0: CRP20248, CRP20249, CRP20252, CRP20255, CRP20258, CRP20259; Estándar Pagos 2.0: p. 61-63; Guía Pagos 2.0: p. 43-48 |
| `REGLA_PAGO_IMPORTE_DR` | error | Pagos 2.0: ImporteDR entre (BaseDR − ½·10^-dB)·TasaOCuotaDR truncado y (BaseDR + ½·10^-dB − 10^-12)·TasaOCuotaDR redondeado hacia arriba, a los decimales del propio ImporteDR. **Dos niveles**: fuera de eso pero dentro de los límites a los decimales de MonedaDR ampliados en ±10 unidades (20 conceptos × ½; 0.10 en MXN), advertencia: el impuesto de la factura redondeado concepto por concepto | Matriz Pagos 2.0: CRP20254, CRP20261; Estándar Pagos 2.0: p. 62, 63; Guía Pagos 2.0: p. 4, 46; Anexo 20 4.0: p. 77, 78 |
| `REGLA_PAGO_IMPUESTOS_P` | error | Pagos 2.0: con impuestos en los documentos, ImpuestosP con una RetencionP por cada ImpuestoDR retenido y un TrasladoP por cada ImpuestoDR, TipoFactorDR y TasaOCuotaDR trasladado, sin repetir ni sobrar; TrasladoP Exento sin TasaOCuotaP ni ImporteP | Matriz Pagos 2.0: CRP20263, CRP20266, CRP20271; Estándar Pagos 2.0: p. 25, 27, 63-65; Guía Pagos 2.0: p. 49-51 |
| `REGLA_PAGO_IMPUESTOS_P_SUMA` | error | Pagos 2.0: ImporteP de cada RetencionP, y BaseP e ImporteP de cada TrasladoP, iguales a la suma de ImporteDR (BaseDR) / EquivalenciaDR de su grupo en los documentos del pago: exacta, redondeada o truncada a los decimales del atributo (al menos los de MonedaP). **Dos niveles**: fuera de eso pero entre la suma de cada sumando truncado y la de cada uno redondeado hacia arriba a los decimales de MonedaP, advertencia | Matriz Pagos 2.0: CRP20265, CRP20268, CRP20269, CRP20274; Estándar Pagos 2.0: p. 64-66 (sin tolerancia ni redondeo); Guía Pagos 2.0: p. 50, 51, 53 |
| `REGLA_CCE_TIPOCOMPROBANTE` | error | Comercio Exterior 2.0 en I, E o T | Matriz CCE 2.0: CCE101; Guía CCE 2.0: p. 13 |
| `REGLA_CCE_CERTIFICADO` | error | NumCertificadoOrigen solo con CertificadoOrigen 1 | Matriz CCE 2.0: CCE118; Guía CCE 2.0: p. 43, 44 |
| `REGLA_CCE_TOTALUSD` | error | TotalUSD = suma de ValorDolares redondeada a 2 decimales | Matriz CCE 2.0: CCE122; Guía CCE 2.0: p. 45, 46 |
| `REGLA_CCE_NOIDENTIFICACION` | error | NoIdentificacion en cada concepto, un concepto por mercancía, sin repetir NoIdentificacion + FraccionArancelaria | Matriz CCE 2.0: CCE151-153; Guía CCE 2.0: p. 55 |
| `REGLA_CCE_DATOS_ADUANA` | error | CantidadAduana, UnidadAduana y ValorUnitarioAduana todos o ninguno, y en todas las mercancías si alguna los trae o si su NoIdentificacion se repite; ValorUnitarioAduana > 0 salvo UnidadAduana 99 | Matriz CCE 2.0: CCE165, CCE166, CCE168; Estándar CCE 2.0: p. 43, 44; Guía CCE 2.0: p. 56, 57 |
| `REGLA_CCE_VALOR_DOLARES` | error | ValorDolares entre los límites de CantidadAduana × ValorUnitarioAduana con los decimales de Cantidad y ValorUnitario del concepto, a 2 decimales (0 y 1 no se revisan). **Dos niveles**: en I y E, fuera de eso pero dentro de los límites del Importe de su concepto en dólares (`REGLA_CCE_VALOR_IMPORTE`), advertencia | Matriz CCE 2.0: CCE169, CCE171; Estándar CCE 2.0: p. 44; Guía CCE 2.0: p. 59, 60, 72 |
| `REGLA_NOMINA_TOTALES` | error | TotalPercepciones o TotalOtrosPagos; TotalPercepciones, TotalDeducciones y TotalOtrosPagos solo con sus nodos y contra ellos | Matriz Nómina 1.2 Rev D: NOM31, NOM36-40; Guía Nómina: p. 26, 27 |
| `REGLA_NOMINA_PERCEPCIONES` | error | totales de Percepciones por tipo de percepción y TotalGravado + TotalExento | Matriz Nómina 1.2 Rev D: NOM69-74; Guía Nómina: p. 48-51 |
| `REGLA_NOMINA_DEDUCCIONES` | error | TotalImpuestosRetenidos = deducciones 002; sin 002, ausente | Matriz Nómina 1.2 Rev D: NOM91, NOM92; Guía Nómina: p. 66 |
| `REGLA_NOMINA_IMPORTES` | error | percepciones, deducciones y otros pagos (salvo 002) mayores que cero | Matriz Nómina 1.2 Rev D: NOM75, NOM96, NOM100; Guía Nómina: p. 56, 69, 76 |
| `REGLA_NOMINA_CONCEPTO` | error | ValorUnitario e Importe = TotalPercepciones + TotalOtrosPagos; Descuento = TotalDeducciones | Matriz Nómina 1.2 Rev D: NOM20-22; Guía Nómina: p. 19 |
| `REGLA_NOMINA_FECHAS` | error | FechaInicialPago ≤ FechaFinalPago | Matriz Nómina 1.2 Rev D: NOM35; Guía Nómina: p. 24 |
| `REGLA_NOMINA_SUBCONTRATACION` | error | PorcentajeTiempo de SubContratacion suma 100 | Matriz Nómina 1.2 Rev D: NOM68; Guía Nómina: p. 48 |
| `REGLA_NOMINA_OTROS_PAGOS` | error | SaldoAFavor ≥ RemanenteSalFav; subsidio ≤ SubsidioCausado | Matriz Nómina 1.2 Rev D: NOM102, NOM107; Guía Nómina: p. 77 |
| `REGLA_NOMINA_COMPROBANTE` | error | CFDI 4.0 de nómina: Exportacion 01, MetodoPago PUE, sin InformacionGlobal ni FacAtrAdquirente; Nomina/Emisor/Curp solo y siempre con emisor persona física; receptor persona física (Rfc de 13), 605, CN01, sin ResidenciaFiscal ni NumRegIdTrib; un concepto 84111505, Cantidad 1 sin decimales, ACT, «Pago de nómina», ObjetoImp 01, sin NoIdentificacion, Unidad ni nodos hijo. **Solo guía**: MetodoPago PUE (p. 10), ResidenciaFiscal y NumRegIdTrib (p. 16), advertencia | Matriz Nómina 1.2 Rev D: NOM3-8, NOM11-19, NOM23-29; Guía Nómina: p. 10, 12, 13, 15-19 |
| `REGLA_NOMINA_PERIODICIDAD` | error | TipoNomina O sin PeriodicidadPago 99; E con 99 | Matriz Nómina 1.2 Rev D: NOM33, NOM34; Guía Nómina: p. 22, 42 |
| `REGLA_NOMINA_REGISTRO_PATRONAL` | error | RegistroPatronal con TipoContrato 01-08 y sin él con 09, 10 y 99; con él, NumSeguridadSocial, FechaInicioRelLaboral, Antigüedad, RiesgoPuesto y SalarioDiarioIntegrado, y sin él ninguno. **Solo guía**: «sin él ninguno» (p. 29), advertencia | Matriz Nómina 1.2 Rev D: NOM42-44; Estándar Nómina 1.2 Rev D: p. 51; Guía Nómina: p. 28, 29 |
| `REGLA_NOMINA_ENTIDAD_SNCF` | error | MontoRecursoPropio solo y siempre con OrigenRecurso IM, y menor que TotalPercepciones + TotalOtrosPagos | Matriz Nómina 1.2 Rev D: NOM48-50; Guía Nómina: p. 31, 32 |
| `REGLA_NOMINA_ANTIGUEDAD` | error | FechaInicioRelLaboral ≤ FechaFinalPago; Antigüedad no mayor que el tiempo entre las dos, contando el día final (PnW: n × 7 ≤ días + 1; PaYmMdD: inicio + a años + m meses + d días ≤ final + 1 día) | Matriz Nómina 1.2 Rev D: NOM53-55; Guía Nómina: p. 34-36 |
| `REGLA_NOMINA_TIPO_REGIMEN` | error | TipoRegimen 02-04 con TipoContrato 01-08 y 05-99 con 09 o más; TipoRegimen 02 con OtroPago 002 (salvo 007 u 008); otro TipoRegimen sin 002, 007 ni 008. **Vigencia**: lo del OtroPago (NOM105, NOM106) es error desde el 2020-02-01; antes, advertencia | Matriz Nómina 1.2 Rev D: NOM57, NOM58, NOM105, NOM106; Matriz Nómina 1.2 Rev C: NOM105, NOM106; Matriz Nómina 1.2 Rev B: NOM226, NOM227; Estándar Nómina 1.2 Rev D: p. 52, 57; Guía Nómina: p. 39 |
| `REGLA_NOMINA_CUENTA_BANCARIA` | error | CuentaBancaria de 10, 11, 15, 16 o 18 posiciones; CLABE (18) sin Banco y con dígito de control correcto; las demás con Banco | Matriz Nómina 1.2 Rev D: NOM62-65; Guía Nómina: p. 43 |
| `REGLA_NOMINA_NODOS_PERCEPCION` | error | TotalSueldos, TotalSeparacionIndemnizacion + SeparacionIndemnizacion y TotalJubilacionPensionRetiro + JubilacionPensionRetiro según los tipos de percepción; 039 en una exhibición, 044 en parcialidades; AccionesOTitulos solo y siempre en 045; HorasExtra solo y siempre en 019 | Matriz Nómina 1.2 Rev D: NOM77-85, NOM89, NOM90; Guía Nómina: p. 53, 54, 56, 57, 59, 60 |
| `REGLA_NOMINA_INCAPACIDADES` | error | percepción 014 o deducción 006 con Incapacidades, y la suma de ImporteMonetario igual a su importe | Matriz Nómina 1.2 Rev D: NOM86, NOM87, NOM94, NOM95; Guía Nómina: p. 54, 69 |
| `REGLA_NOMINA_COMPENSACION_SALDO` | error | OtroPago 004 con CompensacionSaldosAFavor, 002 con SubsidioAlEmpleo; Año el anterior al de FechaPago, o el mismo si el periodo es diciembre. **Vigencia**: el Año (NOM103) es error desde el 2020-02-01; antes, advertencia | Matriz Nómina 1.2 Rev D: NOM98, NOM99, NOM103; Matriz Nómina 1.2 Rev C: NOM103; Matriz Nómina 1.2 Rev B: NOM223; Estándar Nómina 1.2 Rev D: p. 56, 57; Guía Nómina: p. 73, 78, 123 |
| `REGLA_CCE_EXPORTACION` | error | 4.0: Exportacion 02 con Comercio Exterior; con CCE 2.0, Exportacion distinta de 03 y, con 02, Incoterm | Matriz CCE 2.0: CCE104, CCE105, CCE120; Estándar CCE 2.0: p. 37; Guía CCE 2.0: p. 14, 45 |
| `REGLA_CCE_NODOS_CFDI` | error | CCE 2.0: sin InformacionGlobal, ni ACuentaTerceros, InformacionAduanera o CuentaPredial en los conceptos; en T, ObjetoImp 01; solo con TimbreFiscalDigital, ImpuestosLocales, LeyendasFiscales, CFDIRegistroFiscal y CartaPorte (de los complementos que lee la librería). **Solo guía**: InformacionAduanera, advertencia | Matriz CCE 2.0: CCE106, CCE110-112, CCE115; Estándar CCE 2.0: p. 35, 36; Guía CCE 2.0: p. 17, 28, 32, 42 |
| `REGLA_CCE_MOTIVO_TRASLADO` | error | CCE 2.0: MotivoTraslado en T y solo en T; con 05, al menos un Propietario, y sin 05 ninguno; con 01, CfdiRelacionados con TipoRelacion y UUID | Matriz CCE 2.0: CCE102, CCE103, CCE116, CCE117; Estándar CCE 2.0: p. 35, 36; Guía CCE 2.0: p. 13, 14, 42, 43 |
| `REGLA_CCE_TRASLADO` | error | CCE 2.0 en T: Rfc del receptor igual al del emisor; un solo Destinatario; sin CantidadAduana, los conceptos de la mercancía con Cantidad ≥ 0.001 de hasta 3 decimales, Unidad de c_UnidadAduana y ValorUnitario de hasta 4 decimales (y los de la moneda) | Matriz CCE 2.0: CCE108, CCE142, CCE154-156; Estándar CCE 2.0: p. 40-42; Guía CCE 2.0: p. 54, 55, 70 |
| `REGLA_CCE_EMISOR` | error | CCE 2.0: Emisor/Curp solo y siempre con Rfc del emisor de 13 caracteres; domicilio del emisor en MEX | Matriz CCE 2.0: CCE123-125; Estándar CCE 2.0: p. 37; Guía CCE 2.0: p. 46, 47 |
| `REGLA_CCE_NUMREGIDTRIB` | error | CCE 2.0: Receptor/NumRegIdTrib solo y siempre con el Rfc genérico XEXX010101000; con NumRegIdTrib, Receptor y Destinatario fuera de MEX | Matriz CCE 2.0: CCE133, CCE134, CCE143; Estándar CCE 2.0: p. 38, 40; Guía CCE 2.0: p. 48, 51 |
| `REGLA_CCE_FRACCION` | error | CCE 2.0: FraccionArancelaria ausente en servicios (UnidadAduana 99 o ClaveUnidad E48 en un concepto de la mercancía) y requerida en lo demás | Matriz CCE 2.0: CCE159, CCE160, CCE170; Estándar CCE 2.0: p. 43, 44; Guía CCE 2.0: p. 56, 57 |
| `REGLA_CCE_VALOR_IMPORTE` | error | CCE 2.0 en I y E: por NoIdentificacion sin ValorDolares 0 ni 1, Σ ValorDolares entre los límites de Σ Importe de sus conceptos × TipoCambio / TipoCambioUSD (cada valor ± ½ de su último decimal, inferior truncado y superior hacia arriba a centésimas); MXN con TipoCambio 1 exacto; USD también Σ Importe ± ½ de su último decimal | Matriz CCE 2.0: CCE157, CCE158, CCE172; Estándar CCE 2.0: p. 42-44; Guía CCE 2.0: p. 55, 56, 58, 59 |
| `REGLA_CCE_MUESTRAS` | error | CCE 2.0 en I y E: las mercancías 9801000100 (muestras) suman en ValorDolares no más que el Descuento de sus conceptos en dólares (límite superior de la conversión de arriba) | Matriz CCE 2.0: CCE163, CCE164; Estándar CCE 2.0: p. 43; Guía CCE 2.0: p. 56, 57 |

</details>

## Retenciones

El CFDI de Retenciones e información de pagos (1.0 de 2014, 2.0 desde 2022) tiene su propia API y su
propio JSON. `Parse` no cambia: una raíz Retenciones sigue siendo `no_es_comprobante`.

```go
switch cfdi.Detectar(xml) {        // TipoComprobante | TipoRetenciones | TipoOtro (solo mira la raíz)
case cfdi.TipoRetenciones:
    rt, err := cfdi.ParseRetencionesBytes(xml) // err: no_es_xml | xml_mal_formado | no_es_retenciones | lectura
    fmt.Println(rt.Version, rt.Totales.MontoTotRet.Original(), rt.Complementos())
    b, _ := rt.JSON()                          // {"tipo_documento":"retenciones","schema_version":"1","retenciones":{…},…}
    cadena, _ := cfdi.CadenaOriginalRetenciones(xml)
    v := cfdi.VerificarRetenciones(xml, certsSAT) // como Verificar
    val := cfdi.ValidarRetenciones(xml)           // solo XSD
}
```

**Modelo.** Las dos versiones van en un solo `Retenciones` con los nombres de la 2.0; la lectura es tan
tolerante como la del CFDI (mismas advertencias, rutas desde `/Retenciones`, con el nombre del atributo que
trae el documento). Lo que cambia de nombre en 1.0:

| Modelo (2.0) | 1.0 |
|---|---|
| `NoCertificado`, `Certificado` | `NumCert`, `Cert` |
| `Emisor.RfcE` | `RFCEmisor` (y `CURPE` → `Emisor.CurpE`, solo 1.0) |
| `Receptor.NacionalidadR` | `Nacionalidad` |
| `Nacional.RfcR`, `Nacional.CurpR` | `RFCRecep`, `CURPR` |
| `Extranjero.NumRegIdTribR` | `NumRegIdTrib` |
| `Periodo.Ejercicio` | `Ejerc` |
| `Totales.MontoTotOperacion`… `MontoTotRet`, `ImpRetenidos.MontoRet` | `montoTotOperacion`… `montoTotRet`, `montoRet` |
| `ImpRetenidos.ImpuestoRet` | `Impuesto` |

Lo exclusivo de 2.0 (`LugarExpRetenc`, `CfdiRetenRelacionados`, `RegimenFiscalE`, `DomicilioFiscalR`,
`UtilidadBimestral`, `ISRCorrespondiente`) viene vacío en 1.0. La `FechaExp` de 1.0 trae zona (el XSD la
exige): se lee sin advertencia, `Fecha.Time` es la hora escrita y el texto conserva la zona; sin zona es
`FECHA_FORMATO`. En 2.0 es como la Fecha del CFDI.

**Complementos**, todos modelados completos (claves de `retenciones` en el JSON): `timbre_fiscal_digital`,
`servicios_plataformas_tecnologicas` (1.0), `intereses`, `pagosaextranjeros`, `dividendos`,
`enajenacionde_acciones`, `arrendamientoenfideicomiso`, `operacionesconderivados`, `intereseshipotecarios`,
`fideicomisonoempresarial`, `planesderetiro` (1.0 y 1.1 en un objeto con `version`), `premios`,
`sector_financiero`. Lo demás (Plataformas Tecnológicas 1.1, Intereses de Retenciones 2.0) queda en
`complementos_no_soportados` con `COMPLEMENTO_NO_SOPORTADO`: no hay XSD ni XSLT en las fuentes.

**JSON.** `schema_version` propio (`SchemaVersionRetenciones`, "1") y
[`schema/retenciones.v1.schema.json`](schema/retenciones.v1.schema.json), embebido como
`cfdi.SchemaRetenciones`. Las definiciones compartidas (decimal, fecha, advertencia, hallazgo, sellos,
validación, timbre) son las mismas del esquema del CFDI; una prueba lo exige, y otra por reflexión que el
esquema y el modelo tengan exactamente las mismas claves.

**Cadena original.** `internal/gencadena` traduce también `esquemas/retencionpago/1/retenciones.xslt` y
`2/retenciones.xslt` con los trece complementos que incluyen (la 2.0 no incluye Planes de retiro 1.0: en
una 2.0 es `COMPLEMENTO_SIN_CADENA`). Esas hojas traen dos cosas que el generador no conocía: el valor de
un atributo de un hijo (`./retenciones:Nacional/@RfcR`, el del primer Nacional, como en XPath) y
`apply-templates` sobre `ImpRetenidos`, que no tiene plantilla (XSLT le aplica la integrada, que no
escribe nada salvo texto; el generador lo detecta y lo deja anotado en la hoja). Las 52 cadenas de
Retenciones de los fixtures coinciden con libxslt, y los dos rechazos son los que dice el oráculo. La del timbre es `CadenaOriginalTFD`, la misma del CFDI.

**Sellos.** 1.0 se sella con SHA-1 (todos los ejemplos públicos del SAT, de 2014 a 2021, verifican con
SHA-1; se acepta también SHA-256); 2.0 con SHA-256. `NoCertificado` (`NumCert`) tiene que ser el del
certificado, como en el CFDI.

**Validación.** `ValidarRetenciones` es solo el esquema. No hay reglas aritméticas: el SAT no publica para
Retenciones matriz de errores ni guía con fórmulas (no están entre los documentos del repo), y la librería
no inventa reglas sin fuente. Que `MontoTotRet` sea la suma de los `MontoRet` (lo dice la documentación del
XSD 2.0) o los totales de Plataformas quedan para cuando haya un documento del SAT que citar.

## Pruebas

```bash
go test ./...
go test -run TestGolden -update .   # regenera los .golden.json de testdata/
python3 scripts/generar_cadenas.py  # regenera las cadenas esperadas con libxslt (necesita lxml)
go test -run TestSellados -update . # vuelve a sellar testdata/sinteticos/sellados/ con los CSD de prueba
python3 scripts/generar_mutantes.py # regenera testdata/mutantes/ (comprueba cada uno con libxml2)
python3 scripts/validar_xsd.py      # regenera los .esquema.json con libxml2 (necesita lxml)
go generate .                       # regenera cadena_reglas.go y esquema_*.go
```

Los dos oráculos cubren también Retenciones (`esquemas/retencionpago`, al lado de `sitio_internet`), con
la lista de complementos que deben llevarse a la cadena escrita a mano en `generar_cadenas.py`. Las
Retenciones selladas de `testdata/sinteticos/sellados/` (`ret20-plataformas.xml` con SHA-256,
`ret10-extranjero.xml` con SHA-1) se regeneran con `go test -run TestSelladosRetenciones -update .`.

El oráculo del esquema es `scripts/validar_xsd.py`: valida cada `.xml` de `testdata/` con
`lxml.etree.XMLSchema` y los XSD oficiales (resolver local, sin red) y guarda `X.esquema.json` con el
veredicto (`valido`, `invalido`, `incompleto`) y los errores de libxml2. La prueba exige que
`ValidarEsquema` dé el mismo veredicto en todos. `testdata/mutantes/` son sintéticos inválidos a
propósito, uno por regla, y `MANIFIESTO.json` dice qué código debe dar cada uno.

Las cadenas esperadas (`*.cadena.txt`, `*.cadena-tfd.txt`, o `*.error` con el motivo del rechazo) las
genera `scripts/generar_cadenas.py` corriendo los mismos XSLT del SAT con libxslt: son un oráculo
independiente del código Go. Además, la cadena que publica CfdiUtils para `cfdi33-valid.xml` y los sellos
genuinos de los fixtures de CfdiUtils, nodecfdi y satcfdi.

Cada `.xml` bajo `testdata/` se compara contra el `.golden.json` de al lado y se valida contra el esquema;
los que están en un directorio `fatales/` deben fallar. Agregar fixtures = soltar un directorio y correr
`-update`.

- `testdata/sinteticos/`: escritos a mano para este repo.
- `testdata/publicos/<org>-<repo>/`: ejemplos de pruebas de proyectos MIT (SAT-CFDI/python-satcfdi, phpcfdi,
  nodecfdi, eclipxe13, lunasoft), sin modificar salvo los RFC, CURP y nombres que no eran de prueba, que se
  sustituyeron por identidades de prueba del SAT (la lista está en `NOTICE`); licencias en `NOTICE`. `INVENTARIO.json` dice de dónde
  sale cada uno (repo, commit, ruta) y su forma (versión, complementos, BOM…), sin RFC ni nombres. Se
  regenera con `python3 scripts/generar_inventario.py`, que además revisa los RFC: aborta si un archivo
  trae certificado de producción (`NoCertificado` 00001…) y un RFC de persona física que no sea de prueba.

## Licencia

MIT (ver `LICENSE`). El código es propio: no se copió código de phpcfdi, nodecfdi ni CfdiUtils. Lo que
viene de terceros conserva su aviso en `NOTICE`: fixtures y certificados de prueba de proyectos MIT, los
casos portados de CfdiUtils en `reglas_test.go`, el catálogo de monedas de satcfdi, y los XSD y XSLT del SAT
tomados de la réplica de phpcfdi/resources-sat-xml (The Unlicense).
