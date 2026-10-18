# Reglas de `ValidarReglas`

Detalle de las reglas aritméticas y de coherencia. La tabla con todas las reglas y su fundamento está en
el [README](../README.md#reglas-validarreglas).

Cada regla tiene un nivel (`Regla.Nivel`): `error` si tiene fundamento en un documento del SAT;
`advertencia` si no lo tiene o si el SAT pide no aplicarla (hoy solo `REGLA_DECIMALES_IMPUESTO_33`: la
matriz 3.3 dice «NO DEBEN CONSIDERAR EN SU APLICATIVO» CFDI33188 y CFDI33194). Solo los errores hacen
`invalido` un documento: con advertencias y sin errores sigue `valido` (y `incompleto` si además hay
avisos).

Una regla de nivel `error` puede tener partes que salen como `advertencia`, con el mismo código y un
mensaje que termina en «; advertencia: …» con el porqué:

- **Vigencia** (`Regla.Vigencia`): la parte es error solo en comprobantes con `Fecha` desde el día que
  dice la regla (hora civil, como `Fecha.Time`); antes, advertencia. Sin `Fecha` legible se aplica. Hoy,
  lo que la Rev B del estándar de Nómina agregó sin fecha de entrada en vigor (ver abajo).
- **Solo guía** (`Regla.SoloGuia`): la parte solo sale de la guía de llenado, sin código en la matriz de
  errores ni texto en el estándar; siempre advertencia. Se revisaron todas las citas de las reglas de
  Nómina, Comercio Exterior 2.0 e impuestos de Pagos 2.0 contra la matriz y el estándar; quedan cuatro partes, en tres reglas:

  | Regla | Parte | Única fuente |
  |---|---|---|
  | `REGLA_NOMINA_COMPROBANTE` | MetodoPago PUE | guía de Nómina p. 10 (y su pregunta 2, p. 86); ni la matriz Rev D ni el estándar (p. 48-50) |
  | `REGLA_NOMINA_COMPROBANTE` | sin ResidenciaFiscal ni NumRegIdTrib en el Receptor | guía de Nómina p. 16; CFDI40151 y CFDI40154 los prohíben con un Rfc inscrito, pero eso necesita la lista de RFC |
  | `REGLA_NOMINA_REGISTRO_PATRONAL` | sin RegistroPatronal, ninguno de NumSeguridadSocial, FechaInicioRelLaboral, Antigüedad, RiesgoPuesto y SalarioDiarioIntegrado | guía de Nómina p. 29; NOM44 y el estándar (p. 51) solo piden que existan con él |
  | `REGLA_CCE_NODOS_CFDI` | sin InformacionAduanera en los conceptos | guía de Comercio Exterior p. 32; no está en la matriz ni en el estándar (p. 35-36) |

  En la misma revisión, dos partes que parecían solo de la guía sí tienen texto en el estándar y siguen
  como error: los ImpuestosP de Pagos 2.0 por cada impuesto de los documentos («Es requerido cuando en los
  documentos relacionados se registre algún impuesto», estándar p. 25, 27) y MotivoTraslado solo en T
  (estándar de Comercio Exterior p. 35).

Nueve reglas tienen **dos niveles** (`Regla.Margen`): `REGLA_IMPORTE_IMPUESTO`, `REGLA_TRASLADO_SUMA`,
`REGLA_RETENCION_SUMA`, `REGLA_TOTAL_TRASLADOS` y `REGLA_TOTAL_RETENCIONES`; en Pagos 2.0
`REGLA_PAGO_IMPORTE_DR`, `REGLA_PAGO_IMPUESTOS_P_SUMA` y `REGLA_PAGOS_TOTALES`, y en Comercio Exterior
`REGLA_CCE_VALOR_DOLARES`. Lo que cumple la letra del
SAT no da hallazgo; lo que no la cumple pero cae en el margen de redondeo por concepto da una
`advertencia` con el mismo código, la misma cita y un mensaje que lo dice («no cumple la letra del SAT,
pero sí el margen de redondeo por concepto, [lo, hi]»); lo que queda fuera también del margen es
`error`. El porqué: un PAC timbra estas diferencias de un centavo (un Traslado de 1633.81 cuando los
conceptos suman 1633.7952; una Retencion de 3.95 sobre 3.94375; un Importe de 52.620000 con Base 328.90
× 0.160000), y marcarlas como error engaña a quien usa la librería. El margen:

- En el Importe del impuesto del concepto, los mismos límites del Anexo 20, truncados y redondeados
  hacia arriba a los decimales de la moneda y no a los del Importe.
- En las sumas de traslados y retenciones, la suma de cada sumando truncado y la de cada uno
  redondeado hacia arriba a los decimales de la moneda; cada sumando es el Importe escrito en el
  concepto o Base × TasaOCuota sin redondear (en la Base, solo la escrita).
- En TotalImpuestosTrasladados y TotalImpuestosRetenidos, lo mismo con cada Importe de los nodos del
  comprobante (en 3.3 pueden traer 6 decimales).
- En el ImporteDR de Pagos 2.0, los límites del estándar a los decimales de MonedaDR, ampliados en 20 ×
  ½ unidad de esos decimales (±0.10 en MXN). El pago «se aplicará proporcionalmente a los conceptos
  integrados en el comprobante» (guía de Pagos p. 4), y los sistemas copian al ImporteDR el impuesto de la
  factura (o su proporción), que el Anexo 20 (p. 77, 78) redondea concepto por concepto: cada concepto
  aleja la suma hasta ½ unidad de BaseDR × TasaOCuotaDR. El CFDI de pago no dice cuántos conceptos tenía
  la factura; el margen cubre 20 (calibrado en local: cubre a casi todas las facturas PPD del corpus de
  regresión, y todos los ImporteDR fuera de la letra que se encontraron ahí eran el impuesto exacto de su
  factura). La librería no ve la factura: lo que exceda el margen es error.
- En ImporteP y BaseP de Pagos 2.0, la suma de cada ImporteDR (BaseDR) / EquivalenciaDR truncado y la
  de cada uno redondeado hacia arriba a los decimales de MonedaP (en ImporteP, cada sumando también como
  BaseDR × TasaOCuotaDR / EquivalenciaDR).
- En los Totales de Pagos 2.0, la suma de cada producto ImporteP (BaseP, Monto) × TipoCambioP truncado y
  la de cada uno redondeado hacia arriba a 2 decimales.

- En ValorDolares de Comercio Exterior 2.0 (I y E), los límites de `REGLA_CCE_VALOR_IMPORTE` con el
  Importe de su concepto: ver abajo.

En las tres de Pagos el SAT no da tolerancia: la letra es «el redondeo de la suma» (Totales, estándar p.
52-57) o «la suma» (ImpuestosP, estándar p. 64-66, sin decir cómo redondear el cociente entre
EquivalenciaDR: se acepta exacta, redondeada o truncada a los decimales del atributo, nunca menos que los
de MonedaP). El margen solo baja a advertencia lo que se explica por redondear cada documento o cada pago
antes de sumar.

Todas las reglas se contrastaron con el texto del SAT. Lo que no coincidía se corrigió:

- Los límites del Importe del concepto y de sus impuestos se truncan y redondean a los decimales del
  propio Importe («con la cantidad de decimales que tenga registrado este atributo», estándar 4.0 p.
  76-78), no a los de la moneda. Con los de la moneda queda el segundo nivel (advertencia).
- Los traslados y retenciones del comprobante son la suma de los de los conceptos o su redondeo, sin
  tolerancia por redondeo de cada concepto (CFDI40211, CFDI40221; guía 4.0, pregunta 39: «Redondear al
  final del cálculo y no antes»). Un centavo de diferencia que se explica por el redondeo de cada
  concepto es advertencia (segundo nivel); más allá, error.
- SubTotal solo se compara en I, E y N (CFDI40108); FormaPago está prohibida también en N 4.0
  (CFDI40103); ObjetoImp revisa también 03-08 (CFDI40172, CFDI140226, CFDI140227).
- Pagos 2.0: la suma de ImpPagado / EquivalenciaDR se trunca, y la alternativa por límites suma los
  límites inferiores truncados de cada documento, también con equivalencias sin decimales; con la misma
  moneda la equivalencia es 1 con 10 decimales (guía p. 31-34). Los Totales son el redondeo de la suma
  (CRP20201-20211), no la suma de productos redondeados.



Documentos (las páginas son las impresas en el PDF):

| Documento | URL oficial |
|---|---|
| Matriz de errores CFDI 4.0 (2026-03-25) | http://omawww.sat.gob.mx/tramitesyservicios/Paginas/documentos/MatrizDeErrores_CFDI_v40_20260325.xls |
| Anexo 20 CFDI 4.0 (DOF 2022-01-13) | http://omawww.sat.gob.mx/tramitesyservicios/Paginas/documentos/Anexo20_2022.pdf |
| Guía de llenado CFDI 4.0 (2026-01-16) | http://omawww.sat.gob.mx/tramitesyservicios/Paginas/documentos/Anexo_20_Guia_de_llenado_CFDI.pdf |
| Matriz de errores CFDI 3.3 | http://omawww.sat.gob.mx/tramitesyservicios/Paginas/documentos/MatrizDeErrores_CFDI_v33.xls |
| Anexo 20 CFDI 3.3 (DOF 2017-07-28) | http://omawww.sat.gob.mx/tramitesyservicios/Paginas/documentos/cfdv33.pdf |
| Guía de llenado CFDI 3.3 (2019-10-14) | http://omawww.sat.gob.mx/tramitesyservicios/Paginas/documentos/GuiaAnexo2014102019.pdf |
| Matriz de errores Pagos 2.0 Rev B | http://omawww.sat.gob.mx/tramitesyservicios/Paginas/documentos/MatrizDeErrores_CRP_V20_RevB.xls |
| Guía de llenado Pagos 2.0 (2026-01-16) | http://omawww.sat.gob.mx/tramitesyservicios/Paginas/documentos/Guia_llenado_pagos.pdf |
| Estándar Pagos 2.0 Rev B | http://omawww.sat.gob.mx/tramitesyservicios/Paginas/documentos/Complemento_de_Pagos20_RevB.pdf |
| Matriz de errores Pagos 1.0 | http://omawww.sat.gob.mx/tramitesyservicios/Paginas/documentos/MatrizDeErrores_CRP_V10.xls |
| Guía de llenado Pagos 1.0 (2018-08-31) | http://omawww.sat.gob.mx/tramitesyservicios/Paginas/documentos/Guia_comple_pagos.pdf |
| Estándar ImpuestosLocales | http://omawww.sat.gob.mx/informacion_fiscal/factura_electronica/Documents/Complementoscfdi/implocal.pdf |
| Matriz de errores Comercio Exterior 2.0 | http://omawww.sat.gob.mx/tramitesyservicios/Paginas/documentos/Matriz_de_Errores_CCE_v2.0.xls |
| Guía de llenado Comercio Exterior 2.0 | http://omawww.sat.gob.mx/tramitesyservicios/Paginas/documentos/Guia_complemento_Comercio_Exterior.pdf |
| Estándar Comercio Exterior 2.0 (páginas del PDF: no las imprime) | http://omawww.sat.gob.mx/tramitesyservicios/Paginas/documentos/EstandarComercioExterior_v20.pdf |
| Matriz de errores Nómina 1.2 Rev D (2025-02-21) | enlace del portal del SAT a `21022025_Matriz_Errores_Nomina_v12_Rev_D…xls` |
| Matriz de errores Nómina 1.2 Rev C (2021-12-29) | http://omawww.sat.gob.mx/tramitesyservicios/Paginas/documentos/29122021_Matriz_ErroresNomina_v12_RevC.xls |
| Matriz de errores Nómina 1.2 Rev B (2019-12-17) | http://omawww.sat.gob.mx/tramitesyservicios/Paginas/documentos/matriz_erroresnomina_v12B.xls |
| Estándar Nómina 1.2 Rev D (2025-02-21; páginas del PDF: no las imprime) | enlace del portal del SAT a `Estandar_nomina12_21022025…pdf` |
| Guía de llenado Nómina 1.2 | http://omawww.sat.gob.mx/tramitesyservicios/Paginas/documentos/Guia_llenado_Nomina.pdf |

Las que necesitan red o catálogos externos (lista de RFC, c_TasaOCuota, patrones de c_FormaPago, límites
de TipoCambio y de Monto publicados por el SAT) quedan fuera; el detalle por complemento, abajo.

## Nómina 1.2: decisiones y lo que queda fuera

- NOM55 pide que la Antigüedad «corresponda» con los años, meses y días transcurridos; la guía (p. 35)
  que sea «menor o igual» y cuenta los días como calendario (un mes de marzo es `P31D`, un febrero
  `P28D`). Se aplica menor o igual contando el día final; sumar meses no desborda (31 de enero + 1 mes =
  28 o 29 de febrero).
- NOM62 admite cuentas de 10, 11, 16 y 18 posiciones; la guía (p. 43) y el XSD también 15 (tarjetas): se
  aceptan. El dígito de la CLABE es el de Banxico (pesos 3, 7, 1). El ejemplo de la guía,
  `002215911558451272`, no lo cumple (su dígito sería 4).
- Sin RegistroPatronal, la guía (p. 29) pide que no existan NumSeguridadSocial, FechaInicioRelLaboral,
  Antigüedad, RiesgoPuesto ni SalarioDiarioIntegrado; la matriz y el estándar solo validan el caso
  contrario. Se aplica la guía como advertencia (**Solo guía**, arriba).
- Las del CFDI de nómina (NOM3-29) son de la matriz Rev D, que es de CFDI 4.0: no corren en 3.3. Las del
  complemento corren en las dos.

**Vigencia.** El control de cambios del estándar (p. 56-57) dice qué agregó cada revisión, pero no desde
cuándo es obligatorio. La matriz de la Rev B (publicada el 2019-12-17) es la primera que trae esas
validaciones, con los códigos NOM226, NOM227 y NOM223 (la Rev C las renumeró NOM105, NOM106 y NOM103), y el
portal del SAT (`complemento_nomina.htm`) da los documentos anteriores como «vigentes hasta el 31 de enero
del 2020». Por eso lo que entró después de la primera versión del estándar es error desde el 2020-02-01 y
advertencia antes. La matriz original (2017) pedía para el Año solo que fuera «menor que el año en curso»
(NOM223 de entonces); no se aplica aparte:

| Validación | Cuándo entró | Desde cuándo es error |
|---|---|---|
| NOM105, NOM106: OtroPago 002 (o 007, 008) con TipoRegimen 02 y sin él con otro | Rev B del estándar, 2019-12-11 («Para el atributo TipoRegimen se adiciona validación de la clave 002 en el atributo TipoOtroPago», p. 57), matriz Rev B (2019-12-17): NOM226, NOM227 | 2020-02-01 |
| NOM103: Año de CompensacionSaldosAFavor | guía, revisión 2 (2017-11-10, p. 123) «y una nota al pie para indicar su entrada en vigor» que la guía actual ya no trae; la Rev B del estándar la actualizó (p. 57); matriz Rev B: NOM223 | 2020-02-01 |
| NOM3-29: el CFDI de nómina 4.0 | Rev C del estándar, 2021-12-29 | solo corren en CFDI 4.0, que es posterior |
| Las demás del complemento | primera versión del estándar (Rev A, 2016): el control de cambios no las da como agregadas después | siempre: Nómina 1.2 es la única versión válida desde el 1 de abril de 2017 (guía p. 86) |

| Código | Por qué queda fuera |
|---|---|
| NOM1, NOM2, NOM20-22 | ya los revisan `REGLA_TIPOCOMP_MONEDA_NOMINA` y `REGLA_NOMINA_CONCEPTO` |
| NOM9, NOM41, NOM67 | RFC en la lista de inscritos (l_RFC): red |
| NOM10 | Curp del receptor con XAXX010101000: la exige el XSD siempre |
| NOM30 | el nodo Nomina dentro de ComplementoConcepto: estructura (esquema) |
| NOM32, NOM47, NOM51, NOM52, NOM56, NOM59-61, NOM66, NOM76, NOM88, NOM93, NOM97, NOM104 | claves de catálogo: las revisa `ValidarEsquema` con los catálogos del XSD |
| NOM45, NOM46 | EntidadSNCF según la marca del emisor en la lista de RFC: red |
| NOM101, NOM108 | topes del SubsidioCausado: dependen del decreto vigente en la FechaPago (Rev C: 407.02 y 13.39 por día; Rev D: 475.00 y 15.63) y la matriz no da las fechas de vigencia |
| NOM109 | error no clasificado |

## ValorDolares de Comercio Exterior 2.0

El SAT revisa ValorDolares de dos formas, y las dos están en `ValidarReglas`:

- Por mercancía con CantidadAduana (CCE169, CCE171; estándar p. 44, guía p. 59 y ejemplo p. 72):
  `[(CantidadAduana − ½·10^-dC)·(ValorUnitarioAduana − ½·10^-dV)` truncado,
  `(CantidadAduana + ½·10^-dC − 10^-12)·(ValorUnitarioAduana + ½·10^-dV − 10^-12)` hacia arriba`]`, a
  centésimas, con dC y dV los decimales de **Cantidad y ValorUnitario del concepto** (el ejemplo de la
  guía usa 10^-5 de la Cantidad 482.60594 y no 10^-3 de la CantidadAduana 482,605.940). ValorDolares 0 o 1
  no se revisan. Es `REGLA_CCE_VALOR_DOLARES`.
- Por NoIdentificacion, en I y E (CCE157, CCE158, CCE172; estándar p. 42-44): la suma de ValorDolares
  contra la suma de Importe de los conceptos × TipoCambio / TipoCambioUSD, cada uno ± ½ de su último
  decimal, el mínimo truncado y el máximo hacia arriba a centésimas. Es `REGLA_CCE_VALOR_IMPORTE`.

La guía (p. 72, 74, 76, 78) dice en cada ejemplo que ValorDolares es el Importe entre TipoCambioUSD (o
el Importe mismo en USD) y que «algunos sistemas realizan el cálculo incorrecto, multiplicando la
cantidad en aduana por el valor unitario aduana». Por eso `REGLA_CCE_VALOR_DOLARES` tiene dos niveles:
un ValorDolares fuera de los límites de CantidadAduana × ValorUnitarioAduana que sí es el
Importe de su concepto en dólares es advertencia (lo que no cuadra son los datos de aduana: una
CantidadAduana en otra unidad, un ValorUnitarioAduana redondeado o de otro precio); si tampoco es el
Importe, error. Solo en I y E con un solo concepto y una sola mercancía de ese NoIdentificacion.

## Comercio Exterior 2.0: lo que queda fuera

| Código | Por qué |
|---|---|
| CCE107 | RFC del receptor en la lista de RFC inscritos (l_RFC): red |
| CCE109 | ObjetoImp 02 con Impuestos: ya lo revisa `REGLA_OBJETO_IMP` (CFDI40172) |
| CCE113, CCE114 | un solo nodo, hijo de Complemento: la lectura (`NODO_DUPLICADO`) y el esquema |
| CCE115 (en parte) | coexistencia con complementos que la librería no lee: no se distinguen los del Complemento de los del concepto; algunos que la librería sí lee (Aerolíneas, Donatarias…) la regla todavía no los mira |
| CCE119 | NumeroExportadorConfiable: columna Agrupación (Unión Europea) de c_Pais, que no viene en el XSD |
| CCE121 | TipoCambioUSD del Banco de México (art. 20 CFF): dato externo |
| CCE126-130, CCE136-141, CCE145-150 | domicilios contra c_Estado, c_Municipio, c_Localidad, c_Colonia y c_CodigoPostal: catálogos que no vienen en los XSD, y los patrones de código postal de c_Pais |
| CCE131, CCE132, CCE135, CCE144 | NumRegIdTrib en el registro del país o con el patrón de c_Pais: red o columnas de c_Pais que no vienen en el XSD |
| CCE161, CCE162, CCE167 | vigencia y UMT de c_FraccionArancelaria: el XSD solo trae las claves |
| CCE172 en T | ValorDolares = Importe × TipoCambio / TipoCambioUSD: la conversión solo está definida en I y E (CCE157) |
| CCE999 | error no clasificado |

Por qué dos funciones y un solo tipo `Hallazgo`: el esquema habla del documento (bytes: atributos vacíos,
orden, namespaces) y las reglas de valores ya leídos (`Comprobante`); quien ya hizo `Parse` no relee el
XML. Pero los hallazgos de ambas tienen la misma forma y la misma `Ruta` que las advertencias, así que se
juntan en una lista (`Validacion.Hallazgos()`). `Hallazgo` no es `Advertencia`: una advertencia cuenta lo
que tuvo que hacer la lectura tolerante; un hallazgo, la norma que el documento incumple.
