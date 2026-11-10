#!/usr/bin/env python3
"""Genera las cadenas originales esperadas de los fixtures con un procesador XSLT real.

Uso (desde la raíz del repo; necesita lxml, que trae libxslt):

    python3 scripts/generar_cadenas.py

Si no hay lxml en el sistema, en un venv fuera del repo:

    python3 -m venv ~/tmp/cfdi-xslt-venv && ~/tmp/cfdi-xslt-venv/bin/pip install lxml
    ~/tmp/cfdi-xslt-venv/bin/python scripts/generar_cadenas.py

Corre los XSLT oficiales del SAT (copiados sin cambios en internal/gencadena/xslt) sobre cada
.xml de testdata/ (salvo los de fatales/) y escribe al lado:

    X.cadena.txt       cadena original del comprobante (cadenaoriginal_3_2, 3_3 o 4_0) o de la
                       Retenciones (retenciones.xslt 1.0 o 2.0, en esquemas/retencionpago)
    X.cadena.error     en su lugar, el motivo por el que la librería debe negarse a armarla
    X.cadena-tfd.txt   cadena original del TimbreFiscalDigital (TFD 1.0 o 1.1), si lo hay
    X.cadena-tfd.error en su lugar, el motivo

Es el oráculo independiente de las pruebas de la cadena: no usa nada del código Go. Por eso la
lista de complementos que la librería debe saber llevar a la cadena está escrita aquí a mano,
por versión de CFDI, en vez de leerse del generador: si los dos no coinciden, las pruebas
fallan.
"""
import copy
import os
import re
import sys

from lxml import etree

RAIZ = os.path.join(os.path.dirname(os.path.abspath(__file__)), '..')
TESTDATA = os.path.join(RAIZ, 'testdata')
SAT = os.path.join(RAIZ, 'internal', 'gencadena', 'xslt', 'sitio_internet')

HOJAS = {
    '3.2': 'cfd/3/cadenaoriginal_3_2/cadenaoriginal_3_2.xslt',
    '3.3': 'cfd/3/cadenaoriginal_3_3/cadenaoriginal_3_3.xslt',
    '4.0': 'cfd/4/cadenaoriginal_4_0/cadenaoriginal_4_0.xslt',
    'tfd1.0': 'timbrefiscaldigital/cadenaoriginal_TFD_1_0.xslt',
    'tfd1.1': 'cfd/TimbreFiscalDigital/cadenaoriginal_TFD_1_1.xslt',
    # Retenciones e información de pagos: www.sat.gob.mx/esquemas, al lado de sitio_internet.
    'ret1.0': '../esquemas/retencionpago/1/retenciones.xslt',
    'ret2.0': '../esquemas/retencionpago/2/retenciones.xslt',
}

NS3, NS4 = 'http://www.sat.gob.mx/cfd/3', 'http://www.sat.gob.mx/cfd/4'
RET1, RET2 = 'http://www.sat.gob.mx/esquemas/retencionpago/1', 'http://www.sat.gob.mx/esquemas/retencionpago/2'
TFD = ('http://www.sat.gob.mx/TimbreFiscalDigital', 'TimbreFiscalDigital')


def sat(ns):
    return 'http://www.sat.gob.mx/' + ns


# Complementos (y de concepto) que la librería lee y cuya hoja incluye el XSLT de cada versión.
# El TFD no aporta a la cadena del comprobante: lo agrega el PAC después de sellar.
COMUNES = {
    TFD,
    (sat('nomina12'), 'Nomina'),
    (sat('implocal'), 'ImpuestosLocales'),
    (sat('ComercioExterior11'), 'ComercioExterior'),
    (sat('divisas'), 'Divisas'),
    (sat('registrofiscal'), 'CFDIRegistroFiscal'),
    (sat('leyendasFiscales'), 'LeyendasFiscales'),
    (sat('ventavehiculos'), 'VentaVehiculos'),
    (sat('aerolineas'), 'Aerolineas'),
    (sat('donat'), 'Donatarias'),
    (sat('iedu'), 'instEducativas'),
    (sat('ine'), 'INE'),
    (sat('valesdedespensa'), 'ValesDeDespensa'),
    (sat('pagoenespecie'), 'PagoEnEspecie'),
    (sat('notariospublicos'), 'NotariosPublicos'),
    (sat('servicioparcialconstruccion'), 'parcialesconstruccion'),
    (sat('vehiculousado'), 'VehiculoUsado'),
    (sat('TuristaPasajeroExtranjero'), 'TuristaPasajeroExtranjero'),
    (sat('pfic'), 'PFintegranteCoordinado'),
    (sat('renovacionysustitucionvehiculos'), 'renovacionysustitucionvehiculos'),
    (sat('certificadodestruccion'), 'certificadodedestruccion'),
    (sat('arteantiguedades'), 'obrasarteantiguedades'),
    (sat('detallista'), 'detallista'),
}
HIDROCARBUROS = {
    (sat('GastosHidrocarburos10'), 'GastosHidrocarburos'),
    (sat('IngresosHidrocarburos10'), 'IngresosHidrocarburos'),
}
SOPORTADOS = {
    '3.2': COMUNES | {
        (sat('terceros'), 'PorCuentadeTerceros'),
        (sat('ComercioExterior'), 'ComercioExterior'),
        (sat('EstadoDeCuentaCombustible'), 'EstadoDeCuentaCombustible'),
        (sat('nomina'), 'Nomina'),
        (sat('consumodecombustibles'), 'ConsumoDeCombustibles'),
        (sat('acreditamiento'), 'acreditamientoIEPS'),
        (sat('spei'), 'Complemento_SPEI'),
        (sat('ecc'), 'EstadoDeCuentaCombustible'),
        (sat('ecb'), 'EstadoDeCuentaBancario'),
    },
    '3.3': COMUNES | HIDROCARBUROS | {
        (sat('consumodecombustibles'), 'ConsumoDeCombustibles'),
        (sat('ConsumoDeCombustibles11'), 'ConsumoDeCombustibles'),
        (sat('terceros'), 'PorCuentadeTerceros'),
        (sat('Pagos'), 'Pagos'),
        (sat('CartaPorte'), 'CartaPorte'),
        (sat('CartaPorte20'), 'CartaPorte'),
        (sat('EstadoDeCuentaCombustible'), 'EstadoDeCuentaCombustible'),
        (sat('EstadoDeCuentaCombustible12'), 'EstadoDeCuentaCombustible'),
    },
    '4.0': COMUNES | HIDROCARBUROS | {
        (sat('ConsumoDeCombustibles11'), 'ConsumoDeCombustibles'),
        (sat('Pagos20'), 'Pagos'),
        (sat('CartaPorte20'), 'CartaPorte'),
        (sat('CartaPorte30'), 'CartaPorte'),
        (sat('CartaPorte31'), 'CartaPorte'),
        (sat('ComercioExterior20'), 'ComercioExterior'),
        (sat('EstadoDeCuentaCombustible12'), 'EstadoDeCuentaCombustible'),
        (sat('hidrocarburospetroliferos'), 'HidroYPetro'),
    },
}


def ret(ns):
    return RET1 + '/' + ns


# Retenciones: los complementos que incluye retenciones.xslt de cada versión (todos los que lee la
# librería). La 2.0 no incluye planesderetiro 1.0.
RET_COMUNES = {
    TFD,
    (ret('PlataformasTecnologicas10'), 'ServiciosPlataformasTecnologicas'),
    (ret('intereses'), 'Intereses'),
    (ret('pagosaextranjeros'), 'Pagosaextranjeros'),
    (ret('dividendos'), 'Dividendos'),
    (ret('enajenaciondeacciones'), 'EnajenaciondeAcciones'),
    (ret('arrendamientoenfideicomiso'), 'Arrendamientoenfideicomiso'),
    (ret('operacionesconderivados'), 'Operacionesconderivados'),
    (ret('intereseshipotecarios'), 'Intereseshipotecarios'),
    (ret('fideicomisonoempresarial'), 'Fideicomisonoempresarial'),
    (ret('planesderetiro11'), 'Planesderetiro'),
    (ret('premios'), 'Premios'),
    (ret('sectorfinanciero'), 'SectorFinanciero'),
}
SOPORTADOS['ret1.0'] = RET_COMUNES | {(ret('planesderetiro'), 'Planesderetiro')}
SOPORTADOS['ret2.0'] = RET_COMUNES

_xslt = {}


def transformar(clave, doc):
    if clave not in _xslt:
        _xslt[clave] = etree.XSLT(etree.parse(os.path.join(SAT, HOJAS[clave])))
    return bytes(_xslt[clave](doc))


def nombre(e):
    if not isinstance(e.tag, str):  # comentarios e instrucciones de proceso
        return None, None
    q = etree.QName(e)
    return q.namespace, q.localname


def version_cfdi(raiz):
    ns, local = nombre(raiz)
    if local == 'Retenciones' and ns in (RET1, RET2):
        if ns == RET2 and raiz.get('Version') == '2.0':
            return 'ret2.0'
        if ns == RET1 and raiz.get('Version') == '1.0':
            return 'ret1.0'
        return ''
    if local != 'Comprobante':
        return None
    if ns == NS4 and raiz.get('Version') == '4.0':
        return '4.0'
    if ns == NS3 and raiz.get('Version') == '3.3':
        return '3.3'
    if ns == NS3 and raiz.get('version') == '3.2':
        return '3.2'
    return ''


def complemento_sin_cadena(raiz, version):
    """El primer complemento sin soporte en el orden en que lo recorre la hoja: los de concepto
    (dentro de Conceptos) van antes que los del Complemento del comprobante."""
    ns = nombre(raiz)[0]
    contenedores = list(raiz.iter('{%s}ComplementoConcepto' % ns))
    contenedores += [c for c in raiz if nombre(c) == (ns, 'Complemento')]
    for cont in contenedores:
        for c in cont:
            if isinstance(c.tag, str) and nombre(c) not in SOPORTADOS[version]:
                return 'COMPLEMENTO_SIN_CADENA {%s}%s' % nombre(c)
    return None


def escribir(base, sufijo, cadena=None, error=None):
    for ext, contenido in (('.txt', cadena), ('.error', error)):
        p = base + sufijo + ext
        if contenido is None:
            if os.path.exists(p):
                os.remove(p)
            continue
        with open(p, 'wb') as f:
            f.write(contenido if isinstance(contenido, bytes) else (contenido + '\n').encode())


LATIN = {'iso-8859-1', 'iso8859-1', 'latin1', 'latin-1', 'l1', 'windows-1252', 'cp1252', 'iso-8859-15',
         'us-ascii', 'ascii'}


def encoding_miente(data):
    """Declara un encoding de un byte (latin1, windows-1252, ascii) pero los bytes son UTF-8 con
    caracteres no ASCII: libxml lo lee como latin1 y la cadena sale con mojibake. La librería se
    niega a armar una cadena así (no sabe qué bytes firmó el emisor)."""
    m = re.match(rb'^(?:\xef\xbb\xbf)?\s*<\?xml[^>]*?\sencoding\s*=\s*["\']([^"\']*)["\']', data)
    if not m or m.group(1).decode('ascii', 'replace').strip().lower() not in LATIN:
        return False
    try:
        data.decode('utf-8')
    except UnicodeDecodeError:
        return False
    return any(c >= 0x80 for c in data)


def parsear(p):
    """libxml rechaza un xmlns que no es URI válido (p. ej. xmlns:schemaLocation="a b"), aunque el
    XML esté bien formado; ese error solo, se relee en modo recover, que conserva el árbol."""
    parser = etree.XMLParser()  # propio: el error_log de la excepción acumula los de otros archivos
    try:
        return etree.parse(p, parser)
    except etree.XMLSyntaxError:
        if not all(x.type_name == 'WAR_NS_URI' for x in parser.error_log):
            raise
        return etree.parse(p, etree.XMLParser(recover=True))


def procesar(p):
    base = p[:-len('.xml')]
    try:
        if encoding_miente(open(p, 'rb').read()):
            raise ValueError('encoding')
        doc = parsear(p)
    except (etree.XMLSyntaxError, OSError, ValueError):  # OSError: bytes que no son del encoding declarado
        escribir(base, '.cadena', error='XML_ILEGIBLE')
        escribir(base, '.cadena-tfd', error='XML_ILEGIBLE')
        return 'xml_ilegible'
    raiz = doc.getroot()
    version = version_cfdi(raiz)
    if version is None:
        return 'no_es_comprobante'
    if version == '':
        escribir(base, '.cadena', error='VERSION_NO_SOPORTADA')
    elif motivo := complemento_sin_cadena(raiz, version):
        escribir(base, '.cadena', error=motivo)
    else:
        escribir(base, '.cadena', cadena=transformar(version, doc))

    ns = nombre(raiz)[0]
    tfd = next((t for c in raiz if nombre(c) == (ns, 'Complemento') for t in c if nombre(t) == TFD), None)
    if tfd is None:
        escribir(base, '.cadena-tfd')
    elif tfd.get('Version') == '1.1':
        escribir(base, '.cadena-tfd', cadena=transformar('tfd1.1', etree.ElementTree(copy.deepcopy(tfd))))
    elif tfd.get('version') == '1.0':
        escribir(base, '.cadena-tfd', cadena=transformar('tfd1.0', etree.ElementTree(copy.deepcopy(tfd))))
    else:
        escribir(base, '.cadena-tfd', error='VERSION_NO_SOPORTADA')
    return version or 'version_no_soportada'


def main():
    conteo = {}
    for d, _, fs in os.walk(TESTDATA):
        if 'fatales' in d.split(os.sep):
            continue
        for f in sorted(fs):
            # Los *-real* (CFDI reales que publicaron esos proyectos) no se usan para nada nuevo.
            if f.lower().endswith('.xml') and '-real' not in f:
                r = procesar(os.path.join(d, f))
                conteo[r] = conteo.get(r, 0) + 1
    print(conteo, file=sys.stderr)


main()
