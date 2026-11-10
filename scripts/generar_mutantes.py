#!/usr/bin/env python3
"""Genera los mutantes de esquema: CFDI sintéticos inválidos a propósito, uno por regla.

Uso (desde la raíz del repo; necesita lxml):

    python3 scripts/generar_mutantes.py
    python3 scripts/validar_xsd.py        # después: sus goldens

Cada mutante sale de un fixture sintético de testdata/sinteticos (RFC de prueba del SAT, nada
real). Algunas bases se "reparan" antes (sellos de relleno, RFC en mayúsculas) para que la base
reparada sea válida contra los XSD; luego se aplica una sola mutación. Así el único error del
mutante es el que se quiere probar.

Escribe testdata/mutantes/<id>.xml y testdata/mutantes/MANIFIESTO.json con, por mutante, la base,
la mutación, el veredicto que se busca y los códigos de hallazgo que ValidarEsquema debe dar.

Antes de escribir nada comprueba con libxml2 (el mismo envoltorio de scripts/validar_xsd.py) que
cada base reparada es válida y que cada mutante tiene el veredicto buscado: si un mutante no
rompe lo que dice romper, el script falla en vez de guardar una prueba que no prueba nada.
"""
import json
import os
import shutil
import sys

from lxml import etree

sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))
import validar_xsd  # noqa: E402

RAIZ = validar_xsd.RAIZ
SINTETICOS = os.path.join(RAIZ, 'testdata', 'sinteticos')
DESTINO = os.path.join(RAIZ, 'testdata', 'mutantes')

SELLO = 'Sello="U0VMTE8gREUgUFJVRUJB"'
CERT = 'Certificado="Q0VSVElGSUNBRE8gREUgUFJVRUJB"'
TFD_SELLOS = 'SelloCFD="U0VMTE8gREUgUFJVRUJB" NoCertificadoSAT="30001000000500003456" SelloSAT="U0VMTE8gU0FUIERFIFBSVUVCQQ=="'

# Reparaciones por base: lista de (viejo, nuevo), cada viejo debe aparecer una vez.
BASES = {
    'ingreso40': ('ingreso40-pue.xml', []),
    'pago20': ('pago20.xml', []),
    'nomina12': ('nomina12.xml', []),
    'implocal40': ('implocal40.xml', [
        ('FormaPago="03" NoCertificado="30001000000500003416"', f'{SELLO} FormaPago="03" NoCertificado="30001000000500003416" {CERT}'),
        ('NoCertificadoSAT="30001000000500003456"', TFD_SELLOS),
    ]),
    'cartaporte31': ('cartaporte31.xml', [
        ('FormaPago="03" NoCertificado="30001000000500003416"', f'{SELLO} FormaPago="03" NoCertificado="30001000000500003416" {CERT}'),
        ('RFCRemitenteDestinatario="eku9003173c9"', 'RFCRemitenteDestinatario="EKU9003173C9"'),
        ('NoCertificadoSAT="30001000000500003456"', TFD_SELLOS),
    ]),
    'ingreso33': ('ingreso33.xml', [
        ('Rfc="eku9003173c9"', 'Rfc="EKU9003173C9"'),
        ('<cfdi:CuentaPredial Numero="ABC123"/>', '<cfdi:CuentaPredial Numero="123456"/>'),
    ]),
}

EMISOR40 = '<cfdi:Emisor Rfc="EKU9003173C9" Nombre="ESCUELA KEMPER URGATE" RegimenFiscal="601"/>'
RECEPTOR40 = '<cfdi:Receptor Rfc="XAXX010101000" Nombre="PUBLICO EN GENERAL" DomicilioFiscalReceptor="42501" RegimenFiscalReceptor="616" UsoCFDI="S01"/>'
INFO_GLOBAL = '<cfdi:InformacionGlobal Periodicidad="01" Meses="01" Año="2026"/>'
FIN_COMPLEMENTO = '</cfdi:Complemento>'
TFD40 = 'Version="1.1" UUID="abcdef01-2345-4678-9abc-def012345678"'

# Complementos de T7: cada uno con una base válida y al menos dos mutantes inválidos. Los que no
# tienen sintético propio se insertan en ingreso40-pue.xml (al XSD no le importa el tipo de
# comprobante; las reglas de negocio no se evalúan aquí).
SELLOS40 = ('FormaPago="03" NoCertificado="30001000000500003416"', f'{SELLO} FormaPago="03" NoCertificado="30001000000500003416" {CERT}')
SELLOS_T = ('NoCertificado="30001000000500003416" SubTotal="0"', f'{SELLO} NoCertificado="30001000000500003416" {CERT} SubTotal="0"')
SELLOS_TFD = ('NoCertificadoSAT="30001000000500003456"', TFD_SELLOS)
CONCEPTO1_FIN = '      </cfdi:Impuestos>\n    </cfdi:Concepto>\n    <cfdi:Concepto ClaveProdServ="43211500"'


def en_complemento(xml):
    return [(FIN_COMPLEMENTO, '  ' + xml + '\n  ' + FIN_COMPLEMENTO)]


def en_concepto(xml):
    return [(CONCEPTO1_FIN, '      </cfdi:Impuestos>\n      <cfdi:ComplementoConcepto>' + xml + '</cfdi:ComplementoConcepto>\n'
             '    </cfdi:Concepto>\n    <cfdi:Concepto ClaveProdServ="43211500"')]


CCE20 = [SELLOS40, SELLOS_TFD, ('FraccionArancelaria="8471300100"', 'FraccionArancelaria="9403600100"')]
CP20 = ('<cartaporte20:CartaPorte xmlns:cartaporte20="http://www.sat.gob.mx/CartaPorte20" Version="2.0" TranspInternac="No">'
        '<cartaporte20:Ubicaciones>'
        '<cartaporte20:Ubicacion TipoUbicacion="Origen" IDUbicacion="OR000001" RFCRemitenteDestinatario="EKU9003173C9" FechaHoraSalidaLlegada="2026-03-02T08:00:00"/>'
        '<cartaporte20:Ubicacion TipoUbicacion="Destino" IDUbicacion="DE000002" RFCRemitenteDestinatario="XOJI740919U48" FechaHoraSalidaLlegada="2026-03-02T18:00:00" DistanciaRecorrida="120"/>'
        '</cartaporte20:Ubicaciones>'
        '<cartaporte20:Mercancias PesoBrutoTotal="10.000" UnidadPeso="KGM" NumTotalMercancias="1">'
        '<cartaporte20:Mercancia BienesTransp="43211500" Descripcion="Equipo de prueba" Cantidad="1" ClaveUnidad="H87" PesoEnKg="10.000"/>'
        '</cartaporte20:Mercancias>'
        '</cartaporte20:CartaPorte>')
DIVISAS = '<divisas:Divisas xmlns:divisas="http://www.sat.gob.mx/divisas" version="1.0" tipoOperacion="venta"/>'
DONAT = ('<donat:Donatarias xmlns:donat="http://www.sat.gob.mx/donat" version="1.1" noAutorizacion="123" '
         'fechaAutorizacion="2026-01-02" leyenda="Leyenda de prueba"/>')
LEYENDAS = ('<leyendasFisc:LeyendasFiscales xmlns:leyendasFisc="http://www.sat.gob.mx/leyendasFiscales" version="1.0">'
            '<leyendasFisc:Leyenda disposicionFiscal="RESDERAUTH" norma="Artículo de prueba" textoLeyenda="Texto de prueba"/>'
            '</leyendasFisc:LeyendasFiscales>')
REGFISCAL = '<registrofiscal:CFDIRegistroFiscal xmlns:registrofiscal="http://www.sat.gob.mx/registrofiscal" Version="1.0" Folio="0000000000000001"/>'
AEROLINEAS = ('<aerolineas:Aerolineas xmlns:aerolineas="http://www.sat.gob.mx/aerolineas" Version="1.0" TUA="150.00">'
              '<aerolineas:OtrosCargos TotalCargos="80.00"><aerolineas:Cargo CodigoCargo="XV" Importe="80.00"/></aerolineas:OtrosCargos>'
              '</aerolineas:Aerolineas>')
IEDU = ('<iedu:instEducativas xmlns:iedu="http://www.sat.gob.mx/iedu" version="1.0" nombreAlumno="ALUMNO DE PRUEBA" '
        'CURP="XEXX010101HNEXXXA4" nivelEducativo="Primaria" autRVOE="RVOE0001" rfcPago="XOJI740919U48"/>')
TERCEROS33 = [SELLOS40, SELLOS_TFD,
              ('rfc="xiqb891116qe4"', 'rfc="XIQB891116QE4"'),
              ('          <terceros:CuentaPredial numero="CP-0033"/>\n', ''),
              ('version="1.0" ClaveVehicular="0000001"', 'version="1.1" ClaveVehicular="0000001" Niv="NIV-DE-PRUEBA-0001"')]
VV_PARTE = ('<ventavehiculos:Parte cantidad="1" descripcion="Motor de prueba" importe="100.00">\n'
            '            <ventavehiculos:InformacionAduanera numero="21  47  3807  6000034" fecha="2021-05-20"/>\n'
            '          </ventavehiculos:Parte>')
VV_ADUANA = '<ventavehiculos:InformacionAduanera numero="21  47  3807  6000035" fecha="2021-05-21"/>'

BASES.update({
    'pago10': ('pago10.xml', []),
    'cce20': ('cce20.xml', CCE20),
    'cce11': ('cce20.xml', CCE20 + [
        ('xmlns:cce20="http://www.sat.gob.mx/ComercioExterior20"', 'xmlns:cce20="http://www.sat.gob.mx/ComercioExterior11"'),
        ('Version="2.0" ClaveDePedimento', 'Version="1.1" TipoOperacion="2" ClaveDePedimento'),
        # En 1.1 los importes en dólares van con 2 decimales.
        ('ValorUnitarioAduana="500.000000" ValorDolares="1000.0000"', 'ValorUnitarioAduana="500.00" ValorDolares="1000.00"'),
        ('ValorUnitarioAduana="250.500000" ValorDolares="250.5000"', 'ValorUnitarioAduana="250.50" ValorDolares="250.50"'),
        ('TotalUSD="1250.5000"', 'TotalUSD="1250.50"')]),
    'ecc12': ('ecc12.xml', [SELLOS40, SELLOS_TFD, ('Rfc="xiqb891116qe4"', 'Rfc="XIQB891116QE4"')]),
    'cartaporte30': ('cartaporte30.xml', [SELLOS_T, SELLOS_TFD]),
    'cartaporte20': ('ingreso40-pue.xml', en_complemento(CP20)),
    'hidroypetro': ('hidroypetro40.xml', [SELLOS40, SELLOS_TFD,
                                          ('TipoPermiso="PER51" NumeroPermiso="PL/00000/EXP/ES/2026" ClaveHYP="15101514"',
                                           'TipoPermiso="PER01" NumeroPermiso="PL/00000/EXP/ES/2026" ClaveHYP="15101514"'),
                                          ('TipoPermiso="PER51" NumeroPermiso="PL/00000/EXP/ES/2026" ClaveHYP="15101505"',
                                           'TipoPermiso="PER01" NumeroPermiso="PL/00000/EXP/ES/2026" ClaveHYP="15101505"')]),
    'registrofiscal': ('ingreso40-pue.xml', en_complemento(REGFISCAL)),
    'divisas': ('ingreso40-pue.xml', en_complemento(DIVISAS)),
    'donat11': ('ingreso40-pue.xml', en_complemento(DONAT)),
    'leyendas': ('ingreso40-pue.xml', en_complemento(LEYENDAS)),
    'aerolineas': ('ingreso40-pue.xml', en_complemento(AEROLINEAS)),
    'iedu': ('ingreso40-pue.xml', en_concepto(IEDU)),
    'terceros33': ('terceros33.xml', TERCEROS33),
    # T10: versiones viejas (G2).
    'cartaporte10': ('t10/cartaporte10.xml', []),
    'cce10': ('t10/cce10.xml', []),
    'ecc11': ('t10/ecc11.xml', []),
    'nomina11': ('t10/nomina11.xml', []),
})

# T10, G3: complementos de uso común, insertados en ingreso40-pue.xml (datos de prueba).
VALES = ('<valesdedespensa:ValesDeDespensa xmlns:valesdedespensa="http://www.sat.gob.mx/valesdedespensa" version="1.0" '
         'tipoOperacion="monedero electrónico" numeroDeCuenta="123456" total="100.00"><valesdedespensa:Conceptos>'
         '<valesdedespensa:Concepto identificador="ID-01" fecha="2026-01-15T10:00:00" rfc="XOJI740919U48" curp="XEXX010101HNEXXXA4" '
         'nombre="TRABAJADOR DE PRUEBA" importe="100.00"/></valesdedespensa:Conceptos></valesdedespensa:ValesDeDespensa>')
INE = ('<ine:INE xmlns:ine="http://www.sat.gob.mx/ine" Version="1.1" TipoProceso="Campaña">'
       '<ine:Entidad ClaveEntidad="QUE" Ambito="Local"><ine:Contabilidad IdContabilidad="123456"/></ine:Entidad></ine:INE>')
CONSUMO11 = ('<cc11:ConsumoDeCombustibles xmlns:cc11="http://www.sat.gob.mx/ConsumoDeCombustibles11" version="1.1" '
             'tipoOperacion="monedero electrónico" numeroDeCuenta="123" subTotal="100.00" total="100.00"><cc11:Conceptos>'
             '<cc11:ConceptoConsumoDeCombustibles identificador="ID-01" fecha="2026-01-15T10:00:00" rfc="XIQB891116QE4" '
             'claveEstacion="AB12345678" tipoCombustible="1" cantidad="12.999" nombreCombustible="Gasolina de prueba" folioOperacion="A" '
             'valorUnitario="12.7" importe="0.8"><cc11:Determinados><cc11:Determinado impuesto="IVA" tasaOCuota="15.5" importe="88.99"/>'
             '</cc11:Determinados></cc11:ConceptoConsumoDeCombustibles></cc11:Conceptos></cc11:ConsumoDeCombustibles>')
NOTARIOS = ('<np:NotariosPublicos xmlns:np="http://www.sat.gob.mx/notariospublicos" Version="1.0"><np:DescInmuebles>'
            '<np:DescInmueble TipoInmueble="01" Calle="Calle de prueba" Municipio="PACHUCA" Pais="MEX" Estado="13" CodigoPostal="42501"/>'
            '</np:DescInmuebles><np:DatosOperacion NumInstrumentoNotarial="999999" FechaInstNotarial="2026-01-02" MontoOperacion="100" '
            'Subtotal="110" IVA="10"/><np:DatosNotario CURP="XEXX010101HNEXXXA4" NumNotaria="777" EntidadFederativa="13"/>'
            '<np:DatosEnajenante CoproSocConyugalE="No"><np:DatosUnEnajenante Nombre="ENAJENANTE" ApellidoPaterno="DE PRUEBA" '
            'RFC="XOJI740919U48" CURP="XEXX010101HNEXXXA4"/></np:DatosEnajenante><np:DatosAdquiriente CoproSocConyugalE="Si">'
            '<np:DatosAdquirientesCopSC><np:DatosAdquirienteCopSC Nombre="ADQUIRIENTE UNO" RFC="XIQB891116QE4" Porcentaje="50.00"/>'
            '<np:DatosAdquirienteCopSC Nombre="ADQUIRIENTE DOS" RFC="CACX7605101P8" Porcentaje="50.00"/></np:DatosAdquirientesCopSC>'
            '</np:DatosAdquiriente></np:NotariosPublicos>')
PARCIALES = ('<spc:parcialesconstruccion xmlns:spc="http://www.sat.gob.mx/servicioparcialconstruccion" Version="1.0" NumPerLicoAut="1234567">'
             '<spc:Inmueble Calle="Calle de prueba" Municipio="PACHUCA" Estado="13" CodigoPostal="42501"/></spc:parcialesconstruccion>')
VEHICULO = ('<vu:VehiculoUsado xmlns:vu="http://www.sat.gob.mx/vehiculousado" Version="1.0" montoAdquisicion="1000.47" '
            'montoEnajenacion="99.91" claveVehicular="ABC1" marca="MARCA" tipo="TIPO" modelo="2020" numeroMotor="ABC123" '
            'numeroSerie="12345678901234567" NIV="0" valor="123.12"><vu:InformacionAduanera numero="1" fecha="2020-10-23" aduana="ADUANA"/>'
            '</vu:VehiculoUsado>')
TPE = ('<tpe:TuristaPasajeroExtranjero xmlns:tpe="http://www.sat.gob.mx/TuristaPasajeroExtranjero" version="1.0" '
       'fechadeTransito="2026-01-15T12:00:59" tipoTransito="Salida"><tpe:datosTransito Via="Terrestre" TipoId="P123456" '
       'NumeroId="N123456" Nacionalidad="MEX" EmpresaTransporte="EMPRESA DE PRUEBA"/></tpe:TuristaPasajeroExtranjero>')
PAGOESPECIE = ('<pee:PagoEnEspecie xmlns:pee="http://www.sat.gob.mx/pagoenespecie" Version="1.0" CvePIC="A&amp;Ñ091939AZ0-20891030-999" '
               'FolioSolDon="PE-09-12345" PzaArtNombre="Obra de prueba" PzaArtTecn="Óleo" PzaArtAProd="1901" PzaArtDim="260x987"/>')
ACREDITAMIENTO = '<aieps:acreditamientoIEPS xmlns:aieps="http://www.sat.gob.mx/acreditamiento" Version="1.0" TAR="617"/>'

BASES.update({
    'vales': ('ingreso40-pue.xml', en_complemento(VALES)),
    'ine11': ('ingreso40-pue.xml', en_complemento(INE)),
    'consumo11': ('ingreso40-pue.xml', en_complemento(CONSUMO11)),
    'notarios': ('ingreso40-pue.xml', en_complemento(NOTARIOS)),
    'parciales': ('ingreso40-pue.xml', en_complemento(PARCIALES)),
    'vehiculousado': ('ingreso40-pue.xml', en_complemento(VEHICULO)),
    'tpe': ('ingreso40-pue.xml', en_complemento(TPE)),
    'pagoenespecie': ('ingreso40-pue.xml', en_complemento(PAGOESPECIE)),
    'acreditamiento': ('ingreso40-pue.xml', en_concepto(ACREDITAMIENTO)),
    'ine10': ('t10/ine10.xml', []),
    'consumo10': ('t10/consumo10.xml', []),
    'spei': ('t10/spei33.xml', []),
})
SPEI_ORD = ('<spei:Ordenante BancoEmisor="BANCO DE PRUEBA" Nombre="ORDENANTE DE PRUEBA" TipoCuenta="40" '
            'Cuenta="002180000000000018" RFC="XOJI740919U48"/>')
SPEI_BEN = ('<spei:Beneficiario BancoReceptor="OTRO BANCO" Nombre="ESCUELA KEMPER URGATE" TipoCuenta="40" '
            'Cuenta="012180000000000016" RFC="EKU9003173C9" Concepto="PAGO DE PRUEBA" IVA="160.00" MontoPago="1160.00"/>')

# (id, base, descripción, veredicto, códigos esperados, [(viejo, nuevo)...])
MUTANTES = [
    # Bases reparadas: deben ser válidas.
    ('base-implocal40', 'implocal40', 'base reparada', 'valido', [], []),
    ('base-cartaporte31', 'cartaporte31', 'base reparada', 'valido', [], []),
    ('base-ingreso33', 'ingreso33', 'base reparada', 'valido', [], []),
    # Atributos.
    ('atributo-requerido-ausente', 'ingreso40', 'sin Comprobante/@Moneda', 'invalido', ['ESQUEMA_ATRIBUTO_FALTANTE'],
     [(' Moneda="MXN"', '')]),
    ('atributo-no-declarado', 'ingreso40', 'Emisor/@Otro', 'invalido', ['ESQUEMA_ATRIBUTO_NO_PERMITIDO'],
     [('RegimenFiscal="601"/>', 'RegimenFiscal="601" Otro="1"/>')]),
    ('atributo-xsi-type', 'ingreso40', 'xsi:type en Emisor', 'invalido', ['ESQUEMA_ATRIBUTO_NO_PERMITIDO'],
     [('RegimenFiscal="601"/>', 'RegimenFiscal="601" xsi:type="cfdi:Otro"/>')]),
    ('atributo-otro-namespace', 'ingreso40', 'xml:lang en Emisor', 'invalido', ['ESQUEMA_ATRIBUTO_NO_PERMITIDO'],
     [('RegimenFiscal="601"/>', 'RegimenFiscal="601" xml:lang="es"/>')]),
    ('valor-fijo', 'ingreso40', 'Version 4.1', 'invalido', ['ESQUEMA_VALOR_FIJO'],
     [('Version="4.0"', 'Version="4.1"')]),
    ('tfd-version-fija', 'ingreso40', 'TFD Version 1.0', 'invalido', ['ESQUEMA_VALOR_FIJO'],
     [('Version="1.1"', 'Version="1.0"')]),
    ('tfd-uuid-ausente', 'ingreso40', 'TFD sin UUID', 'invalido', ['ESQUEMA_ATRIBUTO_FALTANTE'],
     [(' UUID="abcdef01-2345-4678-9abc-def012345678"', '')]),
    # Valores.
    ('catalogo-moneda', 'ingreso40', 'Moneda fuera de c_Moneda', 'invalido', ['ESQUEMA_CATALOGO'],
     [('Moneda="MXN"', 'Moneda="MXX"')]),
    ('catalogo-claveprodserv', 'ingreso40', 'ClaveProdServ fuera de c_ClaveProdServ', 'invalido', ['ESQUEMA_CATALOGO'],
     [('ClaveProdServ="84111506"', 'ClaveProdServ="99999999"')]),
    ('catalogo-codigopostal', 'ingreso40', 'LugarExpedicion fuera de c_CodigoPostal', 'invalido', ['ESQUEMA_CATALOGO'],
     [('LugarExpedicion="42501"', 'LugarExpedicion="00001"')]),
    ('patron-rfc', 'ingreso40', 'Emisor/@Rfc con el dígito verificador fuera del patrón', 'invalido', ['ESQUEMA_PATRON'],
     [('Rfc="EKU9003173C9"', 'Rfc="EKU9003173CZ"')]),
    ('patron-barra', 'ingreso40', 'Emisor/@Nombre con |', 'invalido', ['ESQUEMA_PATRON'],
     [('Nombre="ESCUELA KEMPER URGATE"', 'Nombre="ESCUELA | URGATE"')]),
    ('patron-anclado', 'ingreso40', 'Emisor/@Rfc con una letra al final: cabe en maxLength 13 y solo pasa el patrón si no se ancla', 'invalido',
     ['ESQUEMA_PATRON'],
     [('Rfc="EKU9003173C9"', 'Rfc="EKU9003173C9X"')]),
    ('decimales-de-mas', 'ingreso40', 'Concepto/@Cantidad con 7 decimales significativos', 'invalido', ['ESQUEMA_DIGITOS'],
     [('Cantidad="2.5"', 'Cantidad="2.5000001"')]),
    ('decimales-ceros-valido', 'ingreso40', 'Cantidad con ceros de más: fractionDigits cuenta el valor', 'valido', [],
     [('Cantidad="2.5"', 'Cantidad="2.5000000"')]),
    ('importe-decimales', 'ingreso40', 'Traslado/@Importe con 7 decimales', 'invalido', ['ESQUEMA_PATRON', 'ESQUEMA_DIGITOS'],
     [('Importe="80.00"', 'Importe="80.0000001"')]),
    ('rango-minimo', 'ingreso40', 'Cantidad 0 (minInclusive 0.000001)', 'invalido', ['ESQUEMA_RANGO'],
     [('Cantidad="2.5"', 'Cantidad="0"')]),
    ('tipo-fecha', 'ingreso40', 'Fecha 30 de febrero', 'invalido', ['ESQUEMA_TIPO'],
     [('Fecha="2026-01-15T10:30:00"', 'Fecha="2026-02-30T10:30:00"')]),
    ('tipo-decimal', 'ingreso40', 'Cantidad que no es decimal', 'invalido', ['ESQUEMA_TIPO'],
     [('Cantidad="2.5"', 'Cantidad="2,5"')]),
    ('longitud', 'ingreso40', 'NoCertificado de 19 dígitos', 'invalido', ['ESQUEMA_LONGITUD', 'ESQUEMA_PATRON'],
     [('NoCertificado="30001000000500003416"', 'NoCertificado="3000100000050000341"')]),
    ('espacios-colapsan-valido', 'ingreso40', 'Moneda y Fecha con espacios: whiteSpace collapse', 'valido', [],
     [('Moneda="MXN"', 'Moneda=" MXN "'), ('Fecha="2026-01-15T10:30:00"', 'Fecha="\n 2026-01-15T10:30:00\t"')]),
    # Estructura.
    ('elemento-requerido-ausente', 'ingreso40', 'sin Receptor', 'invalido', ['ESQUEMA_ELEMENTO_FALTANTE'],
     [('  ' + RECEPTOR40 + '\n', '')]),
    ('orden-cambiado', 'ingreso40', 'Receptor antes que Emisor', 'invalido', ['ESQUEMA_ORDEN'],
     [(EMISOR40 + '\n  ' + RECEPTOR40, RECEPTOR40 + '\n  ' + EMISOR40)]),
    ('orden-opcional', 'ingreso40', 'InformacionGlobal después de Emisor', 'invalido', ['ESQUEMA_ORDEN'],
     [(EMISOR40, EMISOR40 + '\n  ' + INFO_GLOBAL)]),
    ('opcional-en-orden-valido', 'ingreso40', 'InformacionGlobal antes de Emisor', 'valido', [],
     [(EMISOR40, INFO_GLOBAL + '\n  ' + EMISOR40)]),
    ('max-ocurrencias', 'ingreso40', 'Emisor dos veces', 'invalido', ['ESQUEMA_ELEMENTO_EXCEDENTE'],
     [(EMISOR40, EMISOR40 + '\n  ' + EMISOR40)]),
    ('elemento-no-declarado', 'ingreso40', 'cfdi:Extra dentro de Conceptos', 'invalido', ['ESQUEMA_ELEMENTO_NO_PERMITIDO'],
     [('  </cfdi:Conceptos>', '    <cfdi:Extra/>\n  </cfdi:Conceptos>')]),
    ('hijo-en-vacio', 'ingreso40', 'un hijo dentro de Emisor', 'invalido', ['ESQUEMA_ELEMENTO_NO_PERMITIDO'],
     [('RegimenFiscal="601"/>', 'RegimenFiscal="601"><cfdi:Receptor/></cfdi:Emisor>')]),
    ('espacios-en-vacio', 'ingreso40', 'Emisor con un espacio dentro', 'invalido', ['ESQUEMA_CONTENIDO'],
     [('RegimenFiscal="601"/>', 'RegimenFiscal="601"> </cfdi:Emisor>')]),
    ('comentario-en-vacio-valido', 'ingreso40', 'Emisor con un comentario dentro', 'valido', [],
     [('RegimenFiscal="601"/>', 'RegimenFiscal="601"><!-- nota --></cfdi:Emisor>')]),
    ('texto-entre-elementos', 'ingreso40', 'texto dentro de Conceptos', 'invalido', ['ESQUEMA_CONTENIDO'],
     [('  </cfdi:Conceptos>', '    texto\n  </cfdi:Conceptos>')]),
    ('cfdirelacionados-repetibles-valido', 'ingreso40', 'dos CfdiRelacionados (maxOccurs unbounded)', 'valido', [],
     [(EMISOR40, '<cfdi:CfdiRelacionados TipoRelacion="01"><cfdi:CfdiRelacionado UUID="abcdef01-2345-4678-9abc-def012345670"/></cfdi:CfdiRelacionados>'
       '<cfdi:CfdiRelacionados TipoRelacion="04"><cfdi:CfdiRelacionado UUID="abcdef01-2345-4678-9abc-def012345671"/></cfdi:CfdiRelacionados>\n  ' + EMISOR40)]),
    ('complemento-no-soportado', 'ingreso40', 'complemento sin esquema en la librería', 'incompleto', ['ESQUEMA_NO_SOPORTADO'],
     [(FIN_COMPLEMENTO, '  <ext:Otro xmlns:ext="http://example.com/ext" Version="1.0"/>\n  ' + FIN_COMPLEMENTO)]),
    ('addenda-no-soportada', 'ingreso40', 'Addenda con un elemento sin esquema', 'incompleto', ['ESQUEMA_NO_SOPORTADO'],
     [(FIN_COMPLEMENTO + '\n', FIN_COMPLEMENTO + '\n  <cfdi:Addenda><ext:Pedido xmlns:ext="http://example.com/ext" Numero="1"/></cfdi:Addenda>\n')]),
    ('no-soportado-y-error', 'ingreso40', 'complemento sin esquema y un error', 'invalido', ['ESQUEMA_NO_SOPORTADO', 'ESQUEMA_CATALOGO'],
     [(FIN_COMPLEMENTO, '  <ext:Otro xmlns:ext="http://example.com/ext"/>\n  ' + FIN_COMPLEMENTO), ('Moneda="MXN"', 'Moneda="MXX"')]),
    ('complemento-vacio-40-valido', 'ingreso40', 'Complemento vacío (en 4.0 su xs:any es minOccurs=0)', 'valido', [],
     [('<cfdi:Complemento>', '<cfdi:Complemento/><!--'), (FIN_COMPLEMENTO, '-->')]),
    # XML.
    ('xml-mal-formado', 'ingreso40', 'etiqueta sin cerrar', 'invalido', ['ESQUEMA_XML_ILEGIBLE'],
     [('  </cfdi:Conceptos>', '')]),
    ('prefijo-sin-declarar', 'ingreso40', 'prefijo zz sin declarar', 'invalido', ['ESQUEMA_XML_ILEGIBLE'],
     [('<cfdi:Emisor ', '<zz:Emisor ')]),
    ('atributo-repetido', 'ingreso40', 'atributo repetido', 'invalido', ['ESQUEMA_XML_ILEGIBLE'],
     [('Serie="A"', 'Serie="A" Serie="B"')]),
    # Pagos 2.0.
    ('pagos-atributo-ausente', 'pago20', 'Pago sin Monto', 'invalido', ['ESQUEMA_ATRIBUTO_FALTANTE'],
     [(' Monto="1660.00"', '')]),
    ('pagos-orden', 'pago20', 'Totales después de Pago', 'invalido', ['ESQUEMA_ORDEN'],
     [('      <pago20:Totales TotalTrasladosBaseIVA16="1000.00" TotalTrasladosImpuestoIVA16="160.00" TotalTrasladosBaseIVAExento="500.00" MontoTotalPagos="1660.00"/>\n', ''),
      ('    </pago20:Pagos>', '      <pago20:Totales MontoTotalPagos="1660.00"/>\n    </pago20:Pagos>')]),
    ('pagos-catalogo', 'pago20', 'FormaDePagoP fuera de c_FormaPago', 'invalido', ['ESQUEMA_CATALOGO'],
     [('FormaDePagoP="03"', 'FormaDePagoP="33"')]),
    ('pagos-sin-docto', 'pago20', 'ImpuestosDR vacío', 'invalido', ['ESQUEMA_ELEMENTO_FALTANTE'],
     [('<pago20:TrasladoDR BaseDR="500.00" ImpuestoDR="002" TipoFactorDR="Exento"/>', '')]),
    # Nómina 1.2.
    ('nomina-catalogo', 'nomina12', 'TipoNomina fuera de c_TipoNomina', 'invalido', ['ESQUEMA_CATALOGO'],
     [('TipoNomina="O"', 'TipoNomina="X"')]),
    ('nomina-receptor-ausente', 'nomina12', 'Nomina sin Receptor', 'invalido', ['ESQUEMA_ELEMENTO_FALTANTE'],
     [('<nomina12:Receptor Curp="XEXX010101HNEXXXA4"', '<!-- <nomina12:Receptor Curp="XEXX010101HNEXXXA4"'), ('ClaveEntFed="QUE"/>', 'ClaveEntFed="QUE"/> -->')]),
    ('nomina-patron-curp', 'nomina12', 'Curp en minúsculas', 'invalido', ['ESQUEMA_PATRON'],
     [('Curp="XEXX010101HNEXXXA4"', 'Curp="xexx010101hnexxxa4"')]),
    # implocal: la secuencia se repite (maxOccurs="unbounded" en xs:sequence).
    ('implocal-secuencia-repetida-valido', 'implocal40', 'Retención después de los traslados', 'valido', [],
     [('    </implocal:ImpuestosLocales>', '      <implocal:RetencionesLocales ImpLocRetenido="OTRA" TasadeRetencion="1.00" Importe="10.00"/>\n    </implocal:ImpuestosLocales>')]),
    ('implocal-tipo', 'implocal40', 'TasadeRetencion que no es decimal', 'invalido', ['ESQUEMA_TIPO'],
     [('TasadeRetencion="0.50"', 'TasadeRetencion="medio"')]),
    ('implocal-elemento-ajeno', 'implocal40', 'un elemento de cfd/4 dentro de ImpuestosLocales', 'invalido', ['ESQUEMA_ELEMENTO_NO_PERMITIDO'],
     [('    </implocal:ImpuestosLocales>', '      <cfdi:Concepto/>\n    </implocal:ImpuestosLocales>')]),
    # CartaPorte 3.1.
    ('cartaporte-catalogo', 'cartaporte31', 'ConfigVehicular fuera de c_ConfigAutotransporte', 'invalido', ['ESQUEMA_CATALOGO'],
     [('ConfigVehicular="C2"', 'ConfigVehicular="Z9"')]),
    ('cartaporte-mercancias-ausente', 'cartaporte31', 'CartaPorte sin Mercancias', 'invalido', ['ESQUEMA_ELEMENTO_FALTANTE'],
     [('<cartaporte31:Mercancias ', '<!-- <cartaporte31:Mercancias '), ('</cartaporte31:Mercancias>', '</cartaporte31:Mercancias> -->')]),
    ('cartaporte-max', 'cartaporte31', 'dos Autotransporte', 'invalido', ['ESQUEMA_ELEMENTO_EXCEDENTE'],
     [('        <cartaporte31:TransporteMaritimo ', '        <cartaporte31:Autotransporte PermSCT="TPAF01" NumPermisoSCT="PERMISO32"><cartaporte31:IdentificacionVehicular ConfigVehicular="C2" PesoBrutoVehicular="12.00" PlacaVM="XYZ988" AnioModeloVM="2025"/><cartaporte31:Seguros AseguraRespCivil="ASEGURADORA DE PRUEBA" PolizaRespCivil="POL-32"/></cartaporte31:Autotransporte>\n        <cartaporte31:TransporteMaritimo ')]),
    # CFDI 3.3.
    ('cfdi33-complemento-vacio-valido', 'ingreso33', 'Complemento vacío (también en 3.3 su xs:any es minOccurs=0)', 'valido', [],
     [('<cfdi:Complemento>', '<cfdi:Complemento/><!--'), (FIN_COMPLEMENTO, '-->')]),
    ('cfdi33-addenda-vacia', 'ingreso33', 'Addenda vacía (su xs:any pide al menos un elemento)', 'invalido', ['ESQUEMA_ELEMENTO_FALTANTE'],
     [(FIN_COMPLEMENTO + '\n', FIN_COMPLEMENTO + '\n  <cfdi:Addenda/>\n')]),
    ('cfdi33-catalogo', 'ingreso33', 'UsoCFDI fuera de c_UsoCFDI', 'invalido', ['ESQUEMA_CATALOGO'],
     [('UsoCFDI="G03"', 'UsoCFDI="Z99"')]),
    ('cfdi33-en-namespace-40', 'ingreso33', 'Receptor de cfd/4 dentro de un 3.3', 'invalido', ['ESQUEMA_ELEMENTO_NO_PERMITIDO', 'ESQUEMA_ELEMENTO_FALTANTE'],
     [('<cfdi:Receptor ', '<c4:Receptor xmlns:c4="http://www.sat.gob.mx/cfd/4" ')]),
]

MUTANTES += [
    # Pagos 1.0 (en 3.3).
    ('base-pago10', 'pago10', 'base', 'valido', [], []),
    ('pago10-monto-ausente', 'pago10', 'Pago sin Monto', 'invalido', ['ESQUEMA_ATRIBUTO_FALTANTE'],
     [(' Monto="150.00"', '')]),
    ('pago10-catalogo', 'pago10', 'MetodoDePagoDR fuera de c_MetodoPago', 'invalido', ['ESQUEMA_CATALOGO'],
     [('MetodoDePagoDR="PPD" NumParcialidad="1"', 'MetodoDePagoDR="PPX" NumParcialidad="1"')]),
    ('pago10-version', 'pago10', 'Pagos Version 2.0', 'invalido', ['ESQUEMA_VALOR_FIJO'],
     [('<pago10:Pagos Version="1.0">', '<pago10:Pagos Version="2.0">')]),
    # Comercio Exterior 2.0 y 1.1.
    ('base-cce20', 'cce20', 'base reparada', 'valido', [], []),
    ('cce20-certificado-catalogo', 'cce20', 'CertificadoOrigen 2 (enumeration de un xs:int)', 'invalido', ['ESQUEMA_CATALOGO'],
     [('CertificadoOrigen="0"', 'CertificadoOrigen="2"')]),
    ('cce20-certificado-valor-valido', 'cce20', 'CertificadoOrigen 00: enumeration de un xs:int compara valores', 'valido', [],
     [('CertificadoOrigen="0"', 'CertificadoOrigen="00"')]),
    ('cce20-totalusd-ausente', 'cce20', 'sin TotalUSD (requerido en 2.0)', 'invalido', ['ESQUEMA_ATRIBUTO_FALTANTE'],
     [(' TotalUSD="1250.5000"', '')]),
    ('cce20-mercancias-ausente', 'cce20', 'sin Mercancias', 'invalido', ['ESQUEMA_ELEMENTO_FALTANTE'],
     [('      <cce20:Mercancias>', '      <!-- <cce20:Mercancias>'), ('</cce20:Mercancias>', '</cce20:Mercancias> -->')]),
    ('base-cce11', 'cce11', 'base reparada', 'valido', [], []),
    ('cce11-tipooperacion', 'cce11', 'TipoOperacion fuera de catálogo', 'invalido', ['ESQUEMA_CATALOGO'],
     [('TipoOperacion="2"', 'TipoOperacion="3"')]),
    ('cce11-version', 'cce11', 'Version 2.0 en el namespace de 1.1', 'invalido', ['ESQUEMA_VALOR_FIJO'],
     [('Version="1.1" TipoOperacion', 'Version="2.0" TipoOperacion')]),
    # Estado de cuenta de combustibles 1.2.
    ('base-ecc12', 'ecc12', 'base reparada', 'valido', [], []),
    ('ecc12-patron-rfc', 'ecc12', 'Rfc en minúsculas', 'invalido', ['ESQUEMA_PATRON'],
     [('Rfc="XIQB891116QE4" ClaveEstacion="EST0002"', 'Rfc="xiqb891116qe4" ClaveEstacion="EST0002"')]),
    ('ecc12-catalogo', 'ecc12', 'TipoCombustible fuera de c_ClaveTipoCombustible', 'invalido', ['ESQUEMA_CATALOGO'],
     [('TipoCombustible="2"', 'TipoCombustible="99"')]),
    ('ecc12-traslados-ausente', 'ecc12', 'Concepto sin Traslados', 'invalido', ['ESQUEMA_ELEMENTO_FALTANTE'],
     [('<ecc12:Traslado Impuesto="IVA" TasaOCuota="0.160000" Importe="33.60"/>', ''),
      ('          <ecc12:Traslados>\n            \n          </ecc12:Traslados>\n', '')]),
    # Carta Porte 3.0 y 2.0.
    ('base-cartaporte30', 'cartaporte30', 'base reparada', 'valido', [], []),
    ('cartaporte30-catalogo', 'cartaporte30', 'ConfigVehicular fuera de c_ConfigAutotransporte', 'invalido', ['ESQUEMA_CATALOGO'],
     [('ConfigVehicular="T3S2"', 'ConfigVehicular="Z9"')]),
    ('cartaporte30-mercancias-ausente', 'cartaporte30', 'CartaPorte sin Mercancias', 'invalido', ['ESQUEMA_ELEMENTO_FALTANTE'],
     [('<cartaporte30:Mercancias ', '<!-- <cartaporte30:Mercancias '), ('</cartaporte30:Mercancias>', '</cartaporte30:Mercancias> -->')]),
    ('base-cartaporte20', 'cartaporte20', 'base', 'valido', [], []),
    ('cartaporte20-catalogo', 'cartaporte20', 'TranspInternac fuera de Sí/No', 'invalido', ['ESQUEMA_CATALOGO'],
     [('TranspInternac="No"', 'TranspInternac="Quizá"')]),
    ('cartaporte20-orden', 'cartaporte20', 'Mercancias antes que Ubicaciones', 'invalido', ['ESQUEMA_ORDEN'],
     [('<cartaporte20:Ubicaciones>', '<cartaporte20:Mercancias PesoBrutoTotal="10.000" UnidadPeso="KGM" NumTotalMercancias="1">'
       '<cartaporte20:Mercancia BienesTransp="43211500" Descripcion="Equipo de prueba" Cantidad="1" ClaveUnidad="H87" PesoEnKg="10.000"/>'
       '</cartaporte20:Mercancias><cartaporte20:Ubicaciones>'),
      ('</cartaporte20:Ubicaciones><cartaporte20:Mercancias', '</cartaporte20:Ubicaciones><!--'), ('</cartaporte20:Mercancias></cartaporte20:CartaPorte>', '--></cartaporte20:CartaPorte>')]),
    # Hidrocarburos y petrolíferos.
    ('base-hidroypetro', 'hidroypetro', 'base reparada', 'valido', [], []),
    ('hidroypetro-catalogo', 'hidroypetro', 'TipoPermiso fuera de c_TipoPermiso', 'invalido', ['ESQUEMA_CATALOGO'],
     [('TipoPermiso="PER01" NumeroPermiso="PL/00000/EXP/ES/2026" ClaveHYP="15101505"', 'TipoPermiso="PER99" NumeroPermiso="PL/00000/EXP/ES/2026" ClaveHYP="15101505"')]),
    ('hidroypetro-version-ausente', 'hidroypetro', 'HidroYPetro sin Version', 'invalido', ['ESQUEMA_ATRIBUTO_FALTANTE'],
     [('Version="1.0" TipoPermiso="PER01" NumeroPermiso="PL/00000/EXP/ES/2026" ClaveHYP="15101514"', 'TipoPermiso="PER01" NumeroPermiso="PL/00000/EXP/ES/2026" ClaveHYP="15101514"')]),
    # CFDI Registro Fiscal.
    ('base-registrofiscal', 'registrofiscal', 'base', 'valido', [], []),
    ('registrofiscal-patron', 'registrofiscal', 'Folio de 15 dígitos', 'invalido', ['ESQUEMA_PATRON'],
     [('Folio="0000000000000001"', 'Folio="000000000000001"')]),
    ('registrofiscal-version', 'registrofiscal', 'Version 1.1', 'invalido', ['ESQUEMA_VALOR_FIJO'],
     [('Version="1.0" Folio=', 'Version="1.1" Folio=')]),
    # Divisas.
    ('base-divisas', 'divisas', 'base', 'valido', [], []),
    ('divisas-catalogo', 'divisas', 'tipoOperacion con mayúscula', 'invalido', ['ESQUEMA_CATALOGO'],
     [('tipoOperacion="venta"', 'tipoOperacion="Venta"')]),
    ('divisas-version-ausente', 'divisas', 'Divisas sin version', 'invalido', ['ESQUEMA_ATRIBUTO_FALTANTE'],
     [(' version="1.0" tipoOperacion', ' tipoOperacion')]),
    # Donatarias 1.1.
    ('base-donat11', 'donat11', 'base', 'valido', [], []),
    ('donat11-fecha', 'donat11', 'fechaAutorizacion con mes 13', 'invalido', ['ESQUEMA_TIPO'],
     [('fechaAutorizacion="2026-01-02"', 'fechaAutorizacion="2026-13-02"')]),
    ('donat11-leyenda-ausente', 'donat11', 'sin leyenda', 'invalido', ['ESQUEMA_ATRIBUTO_FALTANTE'],
     [(' leyenda="Leyenda de prueba"', '')]),
    # Leyendas fiscales.
    ('base-leyendas', 'leyendas', 'base', 'valido', [], []),
    ('leyendas-sin-leyenda', 'leyendas', 'LeyendasFiscales sin Leyenda', 'invalido', ['ESQUEMA_ELEMENTO_FALTANTE'],
     [('<leyendasFisc:Leyenda disposicionFiscal="RESDERAUTH" norma="Artículo de prueba" textoLeyenda="Texto de prueba"/>', '')]),
    ('leyendas-texto-vacio', 'leyendas', 'textoLeyenda vacío', 'invalido', ['ESQUEMA_LONGITUD'],
     [('textoLeyenda="Texto de prueba"', 'textoLeyenda=" "')]),
    # Aerolíneas.
    ('base-aerolineas', 'aerolineas', 'base', 'valido', [], []),
    ('aerolineas-tipo', 'aerolineas', 'TUA que no es decimal', 'invalido', ['ESQUEMA_TIPO'],
     [('TUA="150.00"', 'TUA="ciento cincuenta"')]),
    ('aerolineas-sin-cargo', 'aerolineas', 'OtrosCargos sin Cargo', 'invalido', ['ESQUEMA_ELEMENTO_FALTANTE'],
     [('<aerolineas:Cargo CodigoCargo="XV" Importe="80.00"/>', '')]),
    ('aerolineas-longitud', 'aerolineas', 'CodigoCargo de 9 caracteres', 'invalido', ['ESQUEMA_LONGITUD'],
     [('CodigoCargo="XV"', 'CodigoCargo="XVXVXVXVX"')]),
    # Instituciones educativas (ComplementoConcepto).
    ('base-iedu', 'iedu', 'base', 'valido', [], []),
    ('iedu-catalogo', 'iedu', 'nivelEducativo fuera de la enumeración', 'invalido', ['ESQUEMA_CATALOGO'],
     [('nivelEducativo="Primaria"', 'nivelEducativo="Universidad"')]),
    ('iedu-curp', 'iedu', 'CURP de 17 caracteres', 'invalido', ['ESQUEMA_LONGITUD', 'ESQUEMA_PATRON'],
     [('CURP="XEXX010101HNEXXXA4"', 'CURP="XEXX010101HNEXXXA"')]),
    ('iedu-autrvoe-ausente', 'iedu', 'sin autRVOE', 'invalido', ['ESQUEMA_ATRIBUTO_FALTANTE'],
     [(' autRVOE="RVOE0001"', '')]),
    ('aerolineas-tua-ausente', 'aerolineas', 'sin TUA', 'invalido', ['ESQUEMA_ATRIBUTO_FALTANTE'],
     [(' TUA="150.00"', '')]),
    # Terceros 1.1 y venta de vehículos 1.1 (ComplementoConcepto en 3.3): xs:choice y tipos complejos con nombre.
    ('base-terceros33', 'terceros33', 'base reparada', 'valido', [], []),
    ('terceros-choice', 'terceros33', 'Parte y CuentaPredial juntas (xs:choice)', 'invalido', ['ESQUEMA_ELEMENTO_NO_PERMITIDO'],
     [('          </terceros:Parte>\n', '          </terceros:Parte>\n          <terceros:CuentaPredial numero="CP-0033"/>\n')]),
    ('terceros-choice-otra-opcion-valido', 'terceros33', 'CuentaPredial en vez de Parte', 'valido', [],
     [('<terceros:Parte cantidad="2" unidad="Pieza" noIdentificacion="PZA-01" descripcion="Parte de prueba" valorUnitario="250.000000" importe="500.000000">\n'
       '            <terceros:InformacionAduanera numero="21  47  3807  6000033" fecha="2021-06-15" aduana="PACHUCA"/>\n'
       '          </terceros:Parte>', '<terceros:CuentaPredial numero="CP-0033"/>')]),
    ('terceros-choice-vacio-valido', 'terceros33', 'ninguna opción del xs:choice (todas son minOccurs 0)', 'valido', [],
     [('<terceros:Parte cantidad="2" unidad="Pieza" noIdentificacion="PZA-01" descripcion="Parte de prueba" valorUnitario="250.000000" importe="500.000000">\n'
       '            <terceros:InformacionAduanera numero="21  47  3807  6000033" fecha="2021-06-15" aduana="PACHUCA"/>\n'
       '          </terceros:Parte>', '')]),
    ('ventavehiculos-vacio-valido', 'terceros33', 'VentaVehiculos sin hijos (el xs:choice admite cero)', 'valido', [],
     [(VV_PARTE, '')]),
    ('terceros-impuestos-ausente', 'terceros33', 'PorCuentadeTerceros sin Impuestos', 'invalido', ['ESQUEMA_ELEMENTO_FALTANTE'],
     [('<terceros:Impuestos>', '<!-- <terceros:Impuestos>'), ('</terceros:Impuestos>', '</terceros:Impuestos> -->')]),
    ('terceros-tipo-con-nombre', 'terceros33', 'InformacionFiscalTercero (extension de t_UbicacionFiscal) sin calle', 'invalido', ['ESQUEMA_ATRIBUTO_FALTANTE'],
     [('calle="Calle de prueba" ', '')]),
    ('ventavehiculos-choice', 'terceros33', 'InformacionAduanera y Parte juntas (xs:choice)', 'invalido', ['ESQUEMA_ELEMENTO_NO_PERMITIDO'],
     [(VV_PARTE, VV_ADUANA + '\n          ' + VV_PARTE)]),
    ('ventavehiculos-aduanas-valido', 'terceros33', 'dos InformacionAduanera en vez de Parte', 'valido', [],
     [(VV_PARTE, VV_ADUANA + '\n          ' + VV_ADUANA)]),
    ('ventavehiculos-version', 'terceros33', 'VentaVehiculos version 1.0 con Niv: se valida con el XSD de 1.0, que no lo tiene',
     'invalido', ['ESQUEMA_ATRIBUTO_NO_PERMITIDO'],
     [('version="1.1" ClaveVehicular', 'version="1.0" ClaveVehicular')]),
    ('ventavehiculos-version-otra', 'terceros33', 'VentaVehiculos version 2.0: ni 1.0 ni 1.1, se valida con 1.1', 'invalido',
     ['ESQUEMA_VALOR_FIJO'],
     [('version="1.1" ClaveVehicular', 'version="2.0" ClaveVehicular')]),
    ('ventavehiculos10-valido', 'terceros33', 'VentaVehiculos version 1.0 sin Niv', 'valido', [],
     [('version="1.1" ClaveVehicular="0000001" Niv="NIV-DE-PRUEBA-0001"', 'version="1.0" ClaveVehicular="0000001"')]),
]

# T10, G2: versiones viejas de complementos que ya se leían, y las que comparten namespace.
LEYENDA_DONAT10 = ('Este comprobante ampara un donativo, el cual será destinado por la donataria a los fines propios '
                   'de su objeto social. En el caso de que los bienes donados hayan sido deducidos previamente para los '
                   'efectos del impuesto sobre la renta, este donativo no es deducible. La reproducción no autorizada de '
                   'este comprobante constituye un delito en los términos de las disposiciones fiscales.')
MUTANTES += [
    ('donat10-valido', 'donat11', 'Donatarias version 1.0 con la leyenda fija de su XSD', 'valido', [],
     [('version="1.1"', 'version="1.0"'), ('leyenda="Leyenda de prueba"', 'leyenda="' + LEYENDA_DONAT10 + '"')]),
    ('donat10-leyenda-libre', 'donat11', 'Donatarias version 1.0 con otra leyenda: se valida con el XSD de 1.0, que la fija',
     'invalido', ['ESQUEMA_VALOR_FIJO'], [('version="1.1"', 'version="1.0"')]),
    ('donat-version-otra', 'donat11', 'Donatarias version 2.0: se valida con 1.1', 'invalido', ['ESQUEMA_VALOR_FIJO'],
     [('version="1.1"', 'version="2.0"')]),
    ('base-cartaporte10', 'cartaporte10', 'base', 'valido', [], []),
    ('cartaporte10-sin-pesoenkg', 'cartaporte10', 'Mercancia sin PesoEnKg', 'invalido', ['ESQUEMA_ATRIBUTO_FALTANTE'],
     [(' PesoEnKg="2000.000"', '')]),
    ('cartaporte10-figura-sin-cve', 'cartaporte10', 'FiguraTransporte sin CveTransporte', 'invalido', ['ESQUEMA_ATRIBUTO_FALTANTE'],
     [(' CveTransporte="01"', '')]),
    ('base-cce10', 'cce10', 'base', 'valido', [], []),
    ('cce10-sin-numregidtrib', 'cce10', 'Receptor sin NumRegIdTrib (requerido en 1.0)', 'invalido', ['ESQUEMA_ATRIBUTO_FALTANTE'],
     [('<cce:Receptor NumRegIdTrib="123456789"/>', '<cce:Receptor/>')]),
    ('base-ecc11', 'ecc11', 'base', 'valido', [], []),
    ('ecc11-sin-noidentificacion', 'ecc11', 'concepto sin NoIdentificacion', 'invalido', ['ESQUEMA_ATRIBUTO_FALTANTE'],
     [(' NoIdentificacion="34006"', '')]),
    ('ecc11-fecha-anterior', 'ecc11', 'Fecha antes de 2015 (minInclusive de xs:dateTime sin zona)', 'invalido', ['ESQUEMA_RANGO'],
     [('Fecha="2018-06-03T07:15:00"', 'Fecha="2014-12-31T23:59:59"')]),
    ('ecc11-fecha-limite-valido', 'ecc11', 'Fecha en el límite de 2015', 'valido', [],
     [('Fecha="2018-06-03T07:15:00"', 'Fecha="2015-01-01T00:00:00"')]),
    ('base-nomina11', 'nomina11', 'base', 'valido', [], []),
    ('nomina11-sin-curp', 'nomina11', 'Nomina sin CURP', 'invalido', ['ESQUEMA_ATRIBUTO_FALTANTE'],
     [(' CURP="XEXX010101HNEXXXA4"', '')]),
]

# T10, G3: un válido y un requerido ausente por complemento; INE 1.0/1.1 por catálogo.
MUTANTES += [
    ('base-vales', 'vales', 'base', 'valido', [], []),
    ('vales-sin-total', 'vales', 'ValesDeDespensa sin total', 'invalido', ['ESQUEMA_ATRIBUTO_FALTANTE'], [(' total="100.00"', '')]),
    ('base-ine11', 'ine11', 'base', 'valido', [], []),
    ('ine11-sin-tipoproceso', 'ine11', 'INE sin TipoProceso', 'invalido', ['ESQUEMA_ATRIBUTO_FALTANTE'], [(' TipoProceso="Campaña"', '')]),
    ('ine11-entidad-qto', 'ine11', 'ClaveEntidad QTO: solo existe en el catálogo de 1.0', 'invalido', ['ESQUEMA_CATALOGO'],
     [('ClaveEntidad="QUE"', 'ClaveEntidad="QTO"')]),
    ('base-ine10', 'ine10', 'base (QTO, catálogo de 1.0)', 'valido', [], []),
    ('ine10-entidad-que', 'ine10', 'INE 1.0 con QUE: se valida con el catálogo de 1.0, que no lo tiene', 'invalido', ['ESQUEMA_CATALOGO'],
     [('ClaveEntidad="QTO"', 'ClaveEntidad="QUE"')]),
    ('base-consumo11', 'consumo11', 'base', 'valido', [], []),
    ('consumo11-sin-tipocombustible', 'consumo11', 'concepto sin tipoCombustible (requerido en 1.1)', 'invalido', ['ESQUEMA_ATRIBUTO_FALTANTE'],
     [(' tipoCombustible="1"', '')]),
    ('base-consumo10', 'consumo10', 'base', 'valido', [], []),
    ('consumo10-tipocombustible', 'consumo10', 'tipoCombustible en 1.0, que no lo tiene', 'invalido', ['ESQUEMA_ATRIBUTO_NO_PERMITIDO'],
     [(' cantidad="40.000"', ' tipoCombustible="1" cantidad="40.000"')]),
    ('base-notarios', 'notarios', 'base', 'valido', [], []),
    ('notarios-sin-curp-notario', 'notarios', 'DatosNotario sin CURP', 'invalido', ['ESQUEMA_ATRIBUTO_FALTANTE'],
     [('<np:DatosNotario CURP="XEXX010101HNEXXXA4" ', '<np:DatosNotario ')]),
    ('base-parciales', 'parciales', 'base', 'valido', [], []),
    ('parciales-sin-numperlicoaut', 'parciales', 'parcialesconstruccion sin NumPerLicoAut', 'invalido', ['ESQUEMA_ATRIBUTO_FALTANTE'],
     [(' NumPerLicoAut="1234567"', '')]),
    ('base-vehiculousado', 'vehiculousado', 'base', 'valido', [], []),
    ('vehiculousado-sin-valor', 'vehiculousado', 'VehiculoUsado sin valor', 'invalido', ['ESQUEMA_ATRIBUTO_FALTANTE'], [(' valor="123.12"', '')]),
    ('base-tpe', 'tpe', 'base', 'valido', [], []),
    ('tpe-sin-tipotransito', 'tpe', 'TuristaPasajeroExtranjero sin tipoTransito', 'invalido', ['ESQUEMA_ATRIBUTO_FALTANTE'],
     [(' tipoTransito="Salida"', '')]),
    ('base-pagoenespecie', 'pagoenespecie', 'base', 'valido', [], []),
    ('pagoenespecie-sin-cvepic', 'pagoenespecie', 'PagoEnEspecie sin CvePIC', 'invalido', ['ESQUEMA_ATRIBUTO_FALTANTE'],
     [(' CvePIC="A&amp;Ñ091939AZ0-20891030-999"', '')]),
    ('base-acreditamiento', 'acreditamiento', 'base', 'valido', [], []),
    ('acreditamiento-sin-tar', 'acreditamiento', 'acreditamientoIEPS sin TAR', 'invalido', ['ESQUEMA_ATRIBUTO_FALTANTE'], [(' TAR="617"', '')]),
    # SPEI: xs:all (sin orden, cada uno una vez), xs:time y rango de xs:date.
    ('base-spei', 'spei', 'base', 'valido', [], []),
    ('spei-orden-libre-valido', 'spei', 'Beneficiario antes que Ordenante (xs:all no tiene orden)', 'valido', [],
     [(SPEI_ORD + '\n        ' + SPEI_BEN, SPEI_BEN + '\n        ' + SPEI_ORD)]),
    ('spei-sin-beneficiario', 'spei', 'SPEI_Tercero sin Beneficiario', 'invalido', ['ESQUEMA_ELEMENTO_FALTANTE'],
     [('\n        ' + SPEI_BEN, '')]),
    ('spei-ordenante-doble', 'spei', 'dos Ordenante (en xs:all, a lo más uno)', 'invalido', ['ESQUEMA_ELEMENTO_EXCEDENTE'],
     [(SPEI_ORD, SPEI_ORD + SPEI_ORD)]),
    ('spei-sin-sello', 'spei', 'SPEI_Tercero sin sello', 'invalido', ['ESQUEMA_ATRIBUTO_FALTANTE'],
     [(' sello="U0VMTE8gREUgUFJVRUJB"', '')]),
    ('spei-hora-invalida', 'spei', 'Hora 24:30:00 (no es xs:time)', 'invalido', ['ESQUEMA_TIPO', 'ESQUEMA_PATRON'],
     [('Hora="11:59:30"', 'Hora="24:30:00"')]),
    ('spei-hora-fraccion-valido', 'spei', 'Hora con fracción de segundo', 'valido', [],
     [('Hora="11:59:30"', 'Hora="11:59:30.25"')]),
]


# G4 (T10): decretos de vehículos, destrucción, obras de arte, PFIC e hidrocarburos. Fragmentos propios con
# RFC de prueba (los valores de catálogo salen de los ejemplos públicos de satcfdi, válidos en libxml2).
PFIC = ('<pfic:PFintegranteCoordinado xmlns:pfic="http://www.sat.gob.mx/pfic" version="1.0" '
        'RFCPF="XOJI740919U48" ClaveVehicular="A12345" Placa="ABC1234"/>')
OBRASARTE = ('<obrasarte:obrasarteantiguedades xmlns:obrasarte="http://www.sat.gob.mx/arteantiguedades" Version="1.0" '
             'TipoBien="04" OtrosTipoBien="Pieza de prueba" TituloAdquirido="05" OtrosTituloAdquirido="Otro de prueba" '
             'Subtotal="100.00" IVA="16.00" FechaAdquisicion="2015-08-10" CaracterísticasDeObraoPieza="07"/>')
DESTRUCCION = ('<destruccion:certificadodedestruccion xmlns:destruccion="http://www.sat.gob.mx/certificadodestruccion" '
               'Version="1.0" Serie="SERIE E" NumFolDesVeh="123ABC">'
               '<destruccion:VehiculoDestruido Marca="Marca" TipooClase="Clase" Año="1999" Modelo="Modelo" NIV="NIV" '
               'NumSerie="23456" NumPlacas="456789" NumMotor="7890123" NumFolTarjCir="3456"/>'
               '</destruccion:certificadodedestruccion>')
RENOVACION = ('<decreto:renovacionysustitucionvehiculos xmlns:decreto="http://www.sat.gob.mx/renovacionysustitucionvehiculos" '
              'Version="1.0" TipoDeDecreto="01">'
              '<decreto:DecretoRenovVehicular VehEnaj="02">'
              '<decreto:VehiculosUsadosEnajenadoPermAlFab PrecioVehUsado="10001.56" TipoVeh="08" Marca="Marca" '
              'TipooClase="Clase" Año="2001" Modelo="Modelo" NIV="NIV" NumSerie="987654321" NumPlacas="456789" '
              'NumMotor="7890123" NumFolTarjCir="3456" NumPedIm="56789" Aduana="Aduana" FechaRegulVeh="1999-12-31" '
              'Foliofiscal="123A"/>'
              '<decreto:VehiculoNuvoSemEnajenadoFabAlPerm Año="1990" Modelo="Modelo" NumPlacas="987654" RFC="XOJI740919U48"/>'
              '</decreto:DecretoRenovVehicular>'
              '</decreto:renovacionysustitucionvehiculos>')
GCEH = ('<gceh:GastosHidrocarburos xmlns:gceh="http://www.sat.gob.mx/GastosHidrocarburos10" Version="1.0" '
        'NumeroContrato="12345678" AreaContractual="abcdefg">'
        '<gceh:Erogacion TipoErogacion="Inversión" MontocuErogacion="100.99" Porcentaje="99.99">'
        '<gceh:DocumentoRelacionado OrigenErogacion="Nacional" FolioFiscalVinculado="12345678-1234-4234-8234-123456780001" '
        'RFCProveedor="XOJI740919U48" MontoTotalIVA="16.00" Mes="10" MontoTotalErogaciones="100.00" '
        'FechaFolioFiscalVinculado="2018-10-22"/>'
        '<gceh:Actividades ActividadRelacionada="05"><gceh:SubActividades SubActividadRelacionada="039">'
        '<gceh:Tareas TareaRelacionada="0155"/></gceh:SubActividades></gceh:Actividades>'
        '</gceh:Erogacion>'
        '</gceh:GastosHidrocarburos>')
IEEH = ('<ieeh:IngresosHidrocarburos xmlns:ieeh="http://www.sat.gob.mx/IngresosHidrocarburos10" Version="1.0" '
        'NumeroContrato="12345678" ContraprestacionPagadaOperador="1000.00" Porcentaje="80.00">'
        '<ieeh:DocumentoRelacionado FolioFiscalVinculado="12345678-1234-4234-8234-123456780002" '
        'FechaFolioFiscalVinculado="2018-10-22" Mes="10"/>'
        '</ieeh:IngresosHidrocarburos>')
BASES.update({
    'pfic': ('ingreso40-pue.xml', en_complemento(PFIC)),
    'obrasarte': ('ingreso40-pue.xml', en_complemento(OBRASARTE)),
    'destruccion': ('ingreso40-pue.xml', en_complemento(DESTRUCCION)),
    'renovacion': ('ingreso40-pue.xml', en_complemento(RENOVACION)),
    'gceh': ('ingreso40-pue.xml', en_complemento(GCEH)),
    'ieeh': ('ingreso40-pue.xml', en_complemento(IEEH)),
})
MUTANTES += [
    ('base-pfic', 'pfic', 'base reparada', 'valido', [], []),
    ('pfic-sin-placa', 'pfic', 'sin Placa', 'invalido', ['ESQUEMA_ATRIBUTO_FALTANTE'], [(' Placa="ABC1234"', '')]),
    ('base-obrasarte', 'obrasarte', 'base reparada', 'valido', [], []),
    ('obrasarte-sin-tipobien', 'obrasarte', 'sin TipoBien', 'invalido', ['ESQUEMA_ATRIBUTO_FALTANTE'], [(' TipoBien="04"', '')]),
    ('base-destruccion', 'destruccion', 'base reparada', 'valido', [], []),
    ('destruccion-sin-vehiculo', 'destruccion', 'sin VehiculoDestruido', 'invalido', ['ESQUEMA_ELEMENTO_FALTANTE'],
     [(DESTRUCCION[DESTRUCCION.index('<destruccion:VehiculoDestruido'):DESTRUCCION.index('</destruccion:certificadodedestruccion>')], '')]),
    ('base-renovacion', 'renovacion', 'base reparada', 'valido', [], []),
    ('renovacion-sin-anio', 'renovacion', 'vehículo nuevo sin Año', 'invalido', ['ESQUEMA_ATRIBUTO_FALTANTE'], [('<decreto:VehiculoNuvoSemEnajenadoFabAlPerm Año="1990" ', '<decreto:VehiculoNuvoSemEnajenadoFabAlPerm ')]),
    ('base-gceh', 'gceh', 'base reparada', 'valido', [], []),
    ('gceh-sin-numerocontrato', 'gceh', 'sin NumeroContrato', 'invalido', ['ESQUEMA_ATRIBUTO_FALTANTE'],
     [(' NumeroContrato="12345678" AreaContractual', ' AreaContractual')]),
    ('base-ieeh', 'ieeh', 'base reparada', 'valido', [], []),
    ('ieeh-sin-mes', 'ieeh', 'DocumentoRelacionado sin Mes', 'invalido', ['ESQUEMA_ATRIBUTO_FALTANTE'], [(' Mes="10"/>', '/>')]),
]


# Detallista: el único con datos en el texto de los elementos (tipos simples y xs:simpleContent),
# mixed (specialInstruction), xs:NMTOKEN, xs:all (Customs, palletInformation y
# tradeItemTaxInformation de la línea, en otro orden que el del XSD) y una secuencia repetible
# (shipTo/nameAndAddress, en dos vueltas). Fragmento propio: GLN y GTIN inventados, sin RFC.
def det(local, contenido='', attrs=''):
    return f'<detallista:{local}{attrs}>{contenido}</detallista:{local}>'


DET_ESPECIAL = det('specialInstruction', 'Nota ' + det('text', 'MIL PESOS 00/100 M.N.'), ' code="ZZZ"')
DET_ADICIONAL = det('AdditionalInformation', det('referenceIdentification', 'REF-0001', ' type="ATZ"'))
DET_ENTREGA = det('DeliveryNote', det('referenceIdentification', 'REC-0001') + det('ReferenceDate', '2026-03-02'))
DET_COMPRADOR = det('buyer', det('gln', '7500000000001') +
                    det('contactInformation', det('personOrDepartmentName', det('text', 'COMPRAS'))))
DET_ADUANA_LINEA = det('Customs', det('ReferenceDate', '2026-02-15') +
                       det('alternatePartyIdentification', '26  47  3807  6000001', ' type="TN"') +
                       det('nameAndAddress', det('name', 'ADUANA DE PRUEBA')) + det('gln', '7500000000004'))
DETALLISTA = (
    '<detallista:detallista xmlns:detallista="http://www.sat.gob.mx/detallista" type="SimpleInvoiceType" '
    'contentVersion="1.3.1" documentStructureVersion="AMC8.1" documentStatus="ORIGINAL">'
    + det('requestForPaymentIdentification', det('entityType', 'INVOICE'))
    + DET_ESPECIAL
    + det('orderIdentification', det('referenceIdentification', 'OC-0001', ' type="ON"') + det('ReferenceDate', '2026-03-01'))
    + DET_ADICIONAL + DET_ENTREGA + DET_COMPRADOR
    + det('seller', det('gln', '7500000000002') +
          det('alternatePartyIdentification', 'PROV-01', ' type="SELLER_ASSIGNED_IDENTIFIER_FOR_A_PARTY"'))
    + det('shipTo', det('gln', '7500000000003') + det('nameAndAddress', det('name', 'BODEGA DE PRUEBA') +
          det('city', 'CIUDAD DE PRUEBA') + det('name', 'ANDEN 2') + det('postalCode', '42501')))
    + det('currency', det('currencyFunction', 'BILLING_CURRENCY') + det('rateOfChange', '1.00'), ' currencyISOCode="MXN"')
    + det('paymentTerms',
          det('netPayment', det('paymentTimePeriod', det('timePeriodDue', det('value', '30'), ' timePeriod="DAYS"')),
              ' netPaymentTermsType="BASIC_NET"') +
          det('discountPayment', det('percentage', '2.00'), ' discountType="ALLOWANCE_BY_PAYMENT_ON_TIME"'),
          ' paymentTermsEvent="DATE_OF_INVOICE" PaymentTermsRelationTime="REFERENCE_AFTER"')
    + '<detallista:shipmentDetail/>'
    + det('allowanceCharge', det('specialServicesType', 'AJ') +
          det('monetaryAmountOrPercentage', det('rate', det('percentage', '3.00'), ' base="INVOICE_VALUE"')),
          ' allowanceChargeType="ALLOWANCE_GLOBAL" settlementType="BILL_BACK" sequenceNumber="1"')
    + det('lineItem',
          det('tradeItemIdentification', det('gtin', '7500000000101'))
          + det('alternateTradeItemIdentification', 'ART-0001', ' type="SUPPLIER_ASSIGNED"')
          + det('tradeItemDescriptionInformation', det('longText', 'ARTICULO DE PRUEBA'), ' language="ES"')
          + det('invoicedQuantity', '10.00', ' unitOfMeasure="PCE"')
          + det('aditionalQuantity', '1', ' QuantityType="FREE_GOODS"')
          + det('grossPrice', det('Amount', '100.00')) + det('netPrice', det('Amount', '97.00'))
          + det('AdditionalInformation', det('referenceIdentification', 'OC-0001', ' type="ON"'))
          + DET_ADUANA_LINEA
          + det('LogisticUnits', det('serialShippingContainerCode', 'SSCC-0001', ' type="SRV"'))
          + det('palletInformation', det('description', 'CAJA', ' type="BOX"') + det('palletQuantity', '2') +
                det('transport', det('methodOfPayment', 'PREPAID_BY_SELLER')))
          + det('extendedAttributes', det('lotNumber', 'LOTE-01', ' productionDate="2026-01-10"'))
          + det('allowanceCharge', det('specialServicesType', 'AA') +
                det('monetaryAmountOrPercentage', det('percentagePerUnit', '3.00') + det('ratePerUnit', det('amountPerUnit', '3.00'))),
                ' allowanceChargeType="ALLOWANCE_GLOBAL" settlementType="OFF_INVOICE"')
          + det('tradeItemTaxInformation', det('taxCategory', 'TRANSFERIDO') + det('taxTypeDescription', 'VAT') +
                det('tradeItemTaxAmount', det('taxPercentage', '16.00') + det('taxAmount', '155.20')))
          + det('totalLineAmount', det('grossAmount', det('Amount', '1000.00')) + det('netAmount', det('Amount', '970.00'))),
          ' type="SimpleInvoiceLineItemType" number="1"')
    + det('totalAmount', det('Amount', '970.00'))
    + det('TotalAllowanceCharge', det('specialServicesType', 'AJ') + det('Amount', '30.00'), ' allowanceOrChargeType="ALLOWANCE"')
    + '</detallista:detallista>')
BASES.update({
    'detallista': ('ingreso40-pue.xml', en_complemento(DETALLISTA)),
})
MUTANTES += [
    ('base-detallista', 'detallista', 'base reparada', 'valido', [], []),
    ('detallista-sin-documentstatus', 'detallista', 'sin documentStatus (atributo requerido)', 'invalido',
     ['ESQUEMA_ATRIBUTO_FALTANTE'], [(' documentStatus="ORIGINAL"', '')]),
    ('detallista-entitytype-fuera-de-catalogo', 'detallista', 'entityType (texto) fuera de la enumeración', 'invalido',
     ['ESQUEMA_CATALOGO'], [('>INVOICE<', '>FACTURA<')]),
    ('detallista-entrega-fuera-de-orden', 'detallista', 'DeliveryNote antes de AdditionalInformation', 'invalido',
     ['ESQUEMA_ORDEN'], [(DET_ADICIONAL + DET_ENTREGA, DET_ENTREGA + DET_ADICIONAL)]),
    ('detallista-specialinstruction-excedente', 'detallista', 'cinco specialInstruction (maxOccurs 4)', 'invalido',
     ['ESQUEMA_ELEMENTO_EXCEDENTE'], [(DET_ESPECIAL, DET_ESPECIAL * 5)]),
    ('detallista-sin-buyer', 'detallista', 'sin buyer (elemento requerido)', 'invalido',
     ['ESQUEMA_ELEMENTO_FALTANTE'], [(DET_COMPRADOR, '')]),
    ('detallista-gln-longitud', 'detallista', 'gln del comprador de 12 caracteres (length 13)', 'invalido',
     ['ESQUEMA_LONGITUD'], [('>7500000000001<', '>750000000001<')]),
    ('detallista-amount-no-decimal', 'detallista', 'Amount de TotalAllowanceCharge que no es xs:decimal', 'invalido',
     ['ESQUEMA_TIPO'], [('>30.00<', '>30,00<')]),
    ('detallista-referencia-fija', 'detallista', 'referenceIdentification de la orden con type="IV" (fixed y enumeración ON)',
     'invalido', ['ESQUEMA_CATALOGO', 'ESQUEMA_VALOR_FIJO'], [('type="ON">OC-0001</detallista:referenceIdentification><detallista:ReferenceDate>',
                                                              'type="IV">OC-0001</detallista:referenceIdentification><detallista:ReferenceDate>')]),
    ('detallista-texto-suelto', 'detallista', 'texto en buyer, que no es mixed', 'invalido',
     ['ESQUEMA_CONTENIDO'], [('<detallista:buyer>', '<detallista:buyer>texto')]),
    ('detallista-hijo-en-gln', 'detallista', 'elemento dentro del gln del vendedor (contenido simple)', 'invalido',
     ['ESQUEMA_ELEMENTO_NO_PERMITIDO'], [('>7500000000002<', '>7500000000002<detallista:x/><')]),
    ('detallista-aduana-sin-fecha', 'detallista', 'Customs de la línea (xs:all) sin ReferenceDate', 'invalido',
     ['ESQUEMA_ELEMENTO_FALTANTE'], [(det('ReferenceDate', '2026-02-15'), '')]),
    ('detallista-unidad-no-nmtoken', 'detallista', 'unitOfMeasure con un espacio en medio (no es xs:NMTOKEN)', 'invalido',
     ['ESQUEMA_TIPO'], [('unitOfMeasure="PCE"', 'unitOfMeasure="P CE"')]),
]


def reemplazar(texto, cambios, donde):
    for viejo, nuevo in cambios:
        n = texto.count(viejo)
        if n != 1:
            raise SystemExit(f'{donde}: {viejo!r} aparece {n} veces (se esperaba 1)')
        texto = texto.replace(viejo, nuevo)
    return texto


def main():
    esquema = validar_xsd.cargar_esquemas()
    tmp = DESTINO + '.tmp'
    shutil.rmtree(tmp, ignore_errors=True)
    os.makedirs(tmp)

    bases = {}
    for nombre, (archivo, reparaciones) in BASES.items():
        with open(os.path.join(SINTETICOS, archivo), encoding='utf-8') as f:
            bases[nombre] = reemplazar(f.read(), reparaciones, 'base ' + nombre)

    manifiesto, fallas = [], []
    for id_, base, desc, quiero, codigos, cambios in MUTANTES:
        texto = reemplazar(bases[base], cambios, id_)
        ruta = os.path.join(tmp, id_ + '.xml')
        with open(ruta, 'w', encoding='utf-8') as f:
            f.write(texto)
        v = validar_xsd.veredicto(esquema, ruta)
        if v['veredicto'] != quiero:
            fallas.append(f'{id_}: libxml2 dice {v["veredicto"]} y se buscaba {quiero}: {v["errores"][:3]}')
        manifiesto.append({'id': id_, 'base': BASES[base][0], 'mutacion': desc, 'veredicto': quiero, 'codigos': codigos})
    if fallas:
        shutil.rmtree(tmp)
        raise SystemExit('mutantes que no prueban lo que dicen:\n  ' + '\n  '.join(fallas))

    with open(os.path.join(tmp, 'MANIFIESTO.json'), 'w', encoding='utf-8') as f:
        json.dump({'generado_por': 'scripts/generar_mutantes.py', 'mutantes': manifiesto}, f, ensure_ascii=False, indent=1)
        f.write('\n')
    # Los goldens de los mutantes (esquema, cadena, modelo…) los escriben otros scripts y pruebas; se
    # conservan los de cada mutante que sigue existiendo.
    if os.path.isdir(DESTINO):
        for a in os.listdir(DESTINO):
            id_ = a.split('.', 1)[0]
            if not a.endswith('.xml') and a != 'MANIFIESTO.json' and os.path.exists(os.path.join(tmp, id_ + '.xml')):
                shutil.copy(os.path.join(DESTINO, a), tmp)
        shutil.rmtree(DESTINO)
    os.rename(tmp, DESTINO)
    print(len(MUTANTES), 'mutantes en', os.path.relpath(DESTINO, RAIZ))


if __name__ == '__main__':
    main()
