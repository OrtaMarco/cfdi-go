#!/usr/bin/env python3
"""Valida los fixtures contra los XSD oficiales del SAT con libxml2 (lxml) y guarda un golden por fixture.

Uso (desde la raíz del repo; necesita lxml):

    python3 scripts/validar_xsd.py            # escribe los goldens
    python3 scripts/validar_xsd.py --revisar  # solo dice qué golden cambiaría (sale con 1 si alguno)

Es el oráculo independiente de ValidarEsquema: no usa nada del código Go. Carga con
lxml.etree.XMLSchema un esquema envoltorio que importa los XSD copiados sin cambios en
internal/genxsd/xsd (CFDI 4.0 y 3.3, TFD 1.1 y los complementos de RAICES); los catálogos y tipos de datos los resuelve libxml2 por las rutas relativas de los xs:import. Un
documento con TFD, Donatarias, VentaVehiculos o INE 1.0 se valida con un envoltorio que trae
el XSD de esa versión en lugar del de 1.1 (comparten namespace; ver VARIANTES y elegir_esquema). El
parser va con no_network y un resolver que niega cualquier URL que no sea un archivo local: si un
XSD intentara bajar algo, el script falla en vez de usar la red.

Por cada X.xml de testdata/ escribe X.esquema.json:

    {"veredicto": "valido" | "invalido" | "incompleto", "errores": [{"linea", "ruta", "mensaje"}]}

    valido      libxml2 no reporta errores
    invalido    al menos un error que no es de los de abajo, o un XML que no se puede leer
    incompleto  solo errores de "no hay declaración" en un xs:any estricto (complemento o addenda
                sin esquema en el envoltorio), o una raíz sin esquema: CFDI 3.2/3.0 (comparten el
                namespace cfd/3 con 3.3 y se distinguen por @version en minúsculas; se clasifican
                sin validar contra cfdv33.xsd) u otro documento

Las tres clases son las de ResultadoValidacion.Estado en Go; la prueba compara el veredicto por
documento. Los mensajes de libxml2 se guardan para leer las diferencias de granularidad, no se
comparan uno a uno.
"""
import json
import os
import sys

from lxml import etree

RAIZ = os.path.abspath(os.path.join(os.path.dirname(os.path.abspath(__file__)), '..'))
TESTDATA = os.path.join(RAIZ, 'testdata')
SITIO = os.path.join(RAIZ, 'internal', 'genxsd', 'xsd', 'sitio_internet')
# Retenciones vive en www.sat.gob.mx/esquemas, al lado de sitio_internet, e importa catálogos de
# sitio_internet: el resolver deja leer los dos.
XSD = os.path.dirname(SITIO)
RET = '../esquemas/retencionpago/'

# Esquemas raíz, por namespace, relativos a sitio_internet. Si cambian, cambiar también
# internal/genxsd.Raices.
RAICES = {
    'http://www.sat.gob.mx/cfd/4': 'cfd/4/cfdv40.xsd',
    'http://www.sat.gob.mx/cfd/3': 'cfd/3/cfdv33.xsd',
    'http://www.sat.gob.mx/TimbreFiscalDigital': 'cfd/TimbreFiscalDigital/TimbreFiscalDigitalv11.xsd',
    'http://www.sat.gob.mx/Pagos20': 'cfd/Pagos/Pagos20.xsd',
    'http://www.sat.gob.mx/nomina12': 'cfd/nomina/nomina12.xsd',
    'http://www.sat.gob.mx/implocal': 'cfd/implocal/implocal.xsd',
    'http://www.sat.gob.mx/CartaPorte31': 'cfd/CartaPorte/CartaPorte31.xsd',
    'http://www.sat.gob.mx/Pagos': 'cfd/Pagos/Pagos10.xsd',
    'http://www.sat.gob.mx/CartaPorte30': 'cfd/CartaPorte/CartaPorte30.xsd',
    'http://www.sat.gob.mx/CartaPorte20': 'cfd/CartaPorte/CartaPorte20.xsd',
    'http://www.sat.gob.mx/ComercioExterior20': 'cfd/ComercioExterior20/ComercioExterior20.xsd',
    'http://www.sat.gob.mx/ComercioExterior11': 'cfd/ComercioExterior11/ComercioExterior11.xsd',
    'http://www.sat.gob.mx/hidrocarburospetroliferos': 'cfd/hidrocarburospetroliferos.xsd',
    'http://www.sat.gob.mx/registrofiscal': 'cfd/cfdiregistrofiscal/cfdiregistrofiscal.xsd',
    'http://www.sat.gob.mx/divisas': 'cfd/divisas/divisas.xsd',
    'http://www.sat.gob.mx/EstadoDeCuentaCombustible12': 'cfd/EstadoDeCuentaCombustible/ecc12.xsd',
    'http://www.sat.gob.mx/terceros': 'cfd/terceros/terceros11.xsd',
    'http://www.sat.gob.mx/ventavehiculos': 'cfd/ventavehiculos/ventavehiculos11.xsd',
    'http://www.sat.gob.mx/aerolineas': 'cfd/aerolineas/aerolineas.xsd',
    'http://www.sat.gob.mx/leyendasFiscales': 'cfd/leyendasFiscales/leyendasFisc.xsd',
    'http://www.sat.gob.mx/iedu': 'cfd/iedu/iedu.xsd',
    'http://www.sat.gob.mx/donat': 'cfd/donat/donat11.xsd',
    'http://www.sat.gob.mx/CartaPorte': 'cfd/CartaPorte/CartaPorte.xsd',
    'http://www.sat.gob.mx/ComercioExterior': 'cfd/ComercioExterior/ComercioExterior10.xsd',
    'http://www.sat.gob.mx/EstadoDeCuentaCombustible': 'cfd/EstadoDeCuentaCombustible/ecc11.xsd',
    'http://www.sat.gob.mx/nomina': 'cfd/nomina/nomina11.xsd',
    'http://www.sat.gob.mx/ine': 'cfd/ine/ine11.xsd',
    'http://www.sat.gob.mx/valesdedespensa': 'cfd/valesdedespensa/valesdedespensa.xsd',
    'http://www.sat.gob.mx/consumodecombustibles': 'cfd/consumodecombustibles/consumodecombustibles.xsd',
    'http://www.sat.gob.mx/ConsumoDeCombustibles11': 'cfd/consumodecombustibles/consumodeCombustibles11.xsd',
    'http://www.sat.gob.mx/pagoenespecie': 'cfd/pagoenespecie/pagoenespecie.xsd',
    'http://www.sat.gob.mx/notariospublicos': 'cfd/notariospublicos/notariospublicos.xsd',
    'http://www.sat.gob.mx/servicioparcialconstruccion': 'cfd/servicioparcialconstruccion/servicioparcialconstruccion.xsd',
    'http://www.sat.gob.mx/vehiculousado': 'cfd/vehiculousado/vehiculousado.xsd',
    'http://www.sat.gob.mx/TuristaPasajeroExtranjero': 'cfd/TuristaPasajeroExtranjero/TuristaPasajeroExtranjero.xsd',
    'http://www.sat.gob.mx/acreditamiento': 'cfd/acreditamiento/AcreditamientoIEPS10.xsd',
    'http://www.sat.gob.mx/spei': 'cfd/spei/spei.xsd',
    'http://www.sat.gob.mx/renovacionysustitucionvehiculos':
        'cfd/renovacionysustitucionvehiculos/renovacionysustitucionvehiculos.xsd',
    'http://www.sat.gob.mx/certificadodestruccion': 'cfd/certificadodestruccion/certificadodedestruccion.xsd',
    'http://www.sat.gob.mx/arteantiguedades': 'cfd/arteantiguedades/obrasarteantiguedades.xsd',
    'http://www.sat.gob.mx/pfic': 'cfd/pfic/pfic.xsd',
    'http://www.sat.gob.mx/GastosHidrocarburos10': 'cfd/GastosHidrocarburos10/GastosHidrocarburos10.xsd',
    'http://www.sat.gob.mx/IngresosHidrocarburos10': 'cfd/IngresosHidrocarburos10/IngresosHidrocarburos.xsd',
    'http://www.sat.gob.mx/ecb': 'cfd/ecb/ecb.xsd',
    'http://www.sat.gob.mx/ecc': 'cfd/ecc/ecc.xsd',
    'http://www.sat.gob.mx/esquemas/retencionpago/2': RET + '2/retencionpagov2.xsd',
    'http://www.sat.gob.mx/esquemas/retencionpago/1': RET + '1/retencionpagov1.xsd',
    'http://www.sat.gob.mx/esquemas/retencionpago/1/PlataformasTecnologicas10':
        RET + '1/PlataformasTecnologicas10/ServiciosPlataformasTecnologicas10.xsd',
    'http://www.sat.gob.mx/esquemas/retencionpago/1/intereses': RET + '1/intereses/intereses.xsd',
    'http://www.sat.gob.mx/esquemas/retencionpago/1/pagosaextranjeros': RET + '1/pagosaextranjeros/pagosaextranjeros.xsd',
    'http://www.sat.gob.mx/esquemas/retencionpago/1/dividendos': RET + '1/dividendos/dividendos.xsd',
    'http://www.sat.gob.mx/esquemas/retencionpago/1/enajenaciondeacciones':
        RET + '1/enajenaciondeacciones/enajenaciondeacciones.xsd',
    'http://www.sat.gob.mx/esquemas/retencionpago/1/arrendamientoenfideicomiso':
        RET + '1/arrendamientoenfideicomiso/arrendamientoenfideicomiso.xsd',
    'http://www.sat.gob.mx/esquemas/retencionpago/1/operacionesconderivados':
        RET + '1/operacionesconderivados/operacionesconderivados.xsd',
    'http://www.sat.gob.mx/esquemas/retencionpago/1/intereseshipotecarios':
        RET + '1/intereseshipotecarios/intereseshipotecarios.xsd',
    'http://www.sat.gob.mx/esquemas/retencionpago/1/fideicomisonoempresarial':
        RET + '1/fideicomisonoempresarial/fideicomisonoempresarial.xsd',
    'http://www.sat.gob.mx/esquemas/retencionpago/1/planesderetiro': RET + '1/planesderetiro/planesderetiro.xsd',
    'http://www.sat.gob.mx/esquemas/retencionpago/1/planesderetiro11': RET + '1/planesderetiro11/planesderetiro11.xsd',
    'http://www.sat.gob.mx/esquemas/retencionpago/1/premios': RET + '1/premios/premios.xsd',
    'http://www.sat.gob.mx/esquemas/retencionpago/1/sectorfinanciero': RET + '1/sectorfinanciero/sectorfinanciero.xsd',
}

NS_TFD = 'http://www.sat.gob.mx/TimbreFiscalDigital'
# Variantes: versiones viejas que comparten namespace con una raíz (TFD 1.0 con 1.1, Donatarias y
# VentaVehiculos 1.0 con 1.1). libxml2 carga un esquema por namespace: un documento cuyos
# elementos {ns}local traen atributo="valor" se valida con un envoltorio que trae el XSD de la
# variante en lugar del de la raíz. Es la regla de internal/genxsd.Variantes, que en Go se aplica
# por elemento. (ns, local, atributo, valor, xsd)
VARIANTES = [
    (NS_TFD, 'TimbreFiscalDigital', 'version', '1.0', 'cfd/TimbreFiscalDigital/TimbreFiscalDigital.xsd'),
    ('http://www.sat.gob.mx/donat', 'Donatarias', 'version', '1.0', 'cfd/donat/donat.xsd'),
    ('http://www.sat.gob.mx/ventavehiculos', 'VentaVehiculos', 'version', '1.0', 'cfd/ventavehiculos/ventavehiculos.xsd'),
    ('http://www.sat.gob.mx/ine', 'INE', 'Version', '1.0', 'cfd/ine/ine10.xsd'),
]

NS3 = 'http://www.sat.gob.mx/cfd/3'
COMODIN = 'No matching global element declaration available, but demanded by the strict wildcard.'
RAIZ_SIN_ESQUEMA = 'No matching global declaration available for the validation root.'


class SoloLocal(etree.Resolver):
    """Niega todo lo que no sea un archivo local bajo XSD."""

    def resolve(self, url, pubid, context):
        ruta = url[len('file://'):] if url.startswith('file://') else url
        if '://' in ruta or not os.path.abspath(ruta).startswith(os.path.abspath(XSD)):
            raise RuntimeError('el esquema pidió algo fuera de los XSD locales: ' + url)
        return self.resolve_filename(ruta, context)


def cargar_esquema(raices=RAICES):
    imports = ''.join(
        f'<xs:import namespace="{ns}" schemaLocation="{ruta}"/>' for ns, ruta in raices.items())
    envoltorio = f'<xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema">{imports}</xs:schema>'
    parser = etree.XMLParser(no_network=True, resolve_entities=False)
    parser.resolvers.add(SoloLocal())
    doc = etree.fromstring(envoltorio.encode(), parser, base_url=os.path.join(os.path.abspath(SITIO), 'envoltorio.xsd'))
    return etree.XMLSchema(doc)


def cargar_esquemas():
    """Los envoltorios, por combinación de variantes, en un caché que se llena al pedirlos."""
    return {}


def envoltorio(esquemas, sustituir):
    clave = tuple(sorted(sustituir.items()))
    if clave not in esquemas:
        esquemas[clave] = cargar_esquema({**RAICES, **sustituir})
    return esquemas[clave]


def elegir_esquemas(esquemas, doc):
    """Los envoltorios con que se valida el documento, cada uno con los elementos cuyos errores no
    cuentan. Si los elementos de una variante son todos de la misma versión, un solo envoltorio.
    Si mezcla versiones (VentaVehiculos 1.0 y 1.1 en dos conceptos), libxml2 no puede validarlo con
    uno: se valida con los dos y de cada uno se descartan los errores dentro de los elementos de la
    otra versión, como en Go, donde cada elemento se valida con la declaración de su versión."""
    sustituir, mezcla = {}, None
    for ns, local, atributo, valor, xsd in VARIANTES:
        es = list(doc.iter('{%s}%s' % (ns, local)))
        de_variante = [e for e in es if e.get(atributo) == valor]
        if de_variante and len(de_variante) == len(es):
            sustituir[ns] = xsd
        elif de_variante:
            if mezcla:
                raise RuntimeError('dos variantes mezcladas en el mismo documento: el oráculo no sabe validarlo')
            mezcla = (ns, xsd, de_variante, [e for e in es if e.get(atributo) != valor])
    if not mezcla:
        return [(envoltorio(esquemas, sustituir), [])]
    ns, xsd, de_variante, de_raiz = mezcla
    return [(envoltorio(esquemas, sustituir), de_variante),
            (envoltorio(esquemas, {**sustituir, ns: xsd}), de_raiz)]


def veredicto(esquemas, ruta):
    parser = etree.XMLParser(no_network=True, resolve_entities=False, huge_tree=True)
    try:
        try:
            doc = etree.parse(ruta, parser)
        except etree.XMLSyntaxError:
            # Un xmlns que no es URI válido (xmlns:schemaLocation="a b") es solo una advertencia de
            # libxml2; ese error solo se relee en modo recover, como en scripts/generar_cadenas.py.
            if not parser.error_log or not all(x.type_name == 'WAR_NS_URI' for x in parser.error_log):
                raise
            doc = etree.parse(ruta, etree.XMLParser(no_network=True, resolve_entities=False, huge_tree=True, recover=True))
    except etree.XMLSyntaxError as e:
        return {'veredicto': 'invalido', 'errores': [{'linea': e.lineno, 'ruta': '', 'mensaje': 'XML ilegible: ' + str(e)}]}
    except OSError as e:  # libxml2 lo da así con bytes que no son del encoding declarado
        return {'veredicto': 'invalido', 'errores': [{'linea': 0, 'ruta': '', 'mensaje': 'XML ilegible: ' + str(e).replace(ruta, os.path.relpath(ruta, RAIZ))}]}
    raiz = doc.getroot()
    if etree.QName(raiz).namespace == NS3 and etree.QName(raiz).localname == 'Comprobante' \
            and 'Version' not in raiz.attrib and 'version' in raiz.attrib:
        return {'veredicto': 'incompleto', 'errores': [{'linea': raiz.sourceline, 'ruta': '/',
                                                       'mensaje': 'CFDI ' + raiz.get('version') + ': fuera del alcance (sin esquema)'}]}
    errores = []
    for esquema, ignorar in elegir_esquemas(esquemas, doc):
        rutas = [doc.getpath(e) for e in ignorar]
        esquema.validate(doc)
        for e in esquema.error_log:
            err = {'linea': e.line, 'ruta': e.path or '', 'mensaje': e.message}
            if err not in errores and not any(err['ruta'] == r or err['ruta'].startswith(r + '/') for r in rutas):
                errores.append(err)
    if not errores:
        return {'veredicto': 'valido', 'errores': []}
    if all(e['mensaje'].endswith(COMODIN) or e['mensaje'].endswith(RAIZ_SIN_ESQUEMA) for e in errores):
        return {'veredicto': 'incompleto', 'errores': errores}
    return {'veredicto': 'invalido', 'errores': errores}


def fixtures():
    for base, dirs, archivos in os.walk(TESTDATA):
        dirs.sort()
        for a in sorted(archivos):
            if a.endswith('.xml'):
                yield os.path.join(base, a)


def main():
    revisar = '--revisar' in sys.argv[1:]
    esquemas = cargar_esquemas()
    cuenta, cambios = {}, 0
    for ruta in fixtures():
        v = veredicto(esquemas, ruta)
        cuenta[v['veredicto']] = cuenta.get(v['veredicto'], 0) + 1
        golden = ruta[:-len('.xml')] + '.esquema.json'
        texto = json.dumps(v, ensure_ascii=False, indent=1) + '\n'
        actual = open(golden, encoding='utf-8').read() if os.path.exists(golden) else None
        if actual != texto:
            cambios += 1
            print('cambia' if actual else 'nuevo', os.path.relpath(golden, RAIZ))
            if not revisar:
                with open(golden, 'w', encoding='utf-8') as f:
                    f.write(texto)
    print(sum(cuenta.values()), 'fixtures:', ', '.join(f'{k} {v}' for k, v in sorted(cuenta.items())))
    if revisar and cambios:
        sys.exit(1)


if __name__ == '__main__':
    main()
