#!/usr/bin/env python3
"""Regenera testdata/publicos/INVENTARIO.json a partir de los XML que hay en disco.

Uso (desde la raíz del repo):

    python3 scripts/generar_inventario.py [--origenes INVENTARIO-previo.json] [--revisar]

- Los campos de análisis (versión, tipo, complementos, BOM…) se recalculan de cada archivo.
- La procedencia (repo, commit, ruta original, licencia) no se puede deducir del archivo: se
  conserva del INVENTARIO.json actual (o del que se pase con --origenes, p. ej. el del staging
  al importar) buscando por sha256. Un XML sin procedencia conocida es un error.
- Antes de escribir, revisa los RFC (ver revisar_rfc): un archivo con certificado de producción
  y un RFC de persona física que no sea de prueba aborta. --revisar solo imprime el reporte.

El inventario no guarda RFC ni nombres: solo procedencia y forma del documento.
"""
import hashlib, json, os, re, sys
import xml.etree.ElementTree as ET

RAIZ = os.path.join(os.path.dirname(os.path.abspath(__file__)), '..')
PUBLICOS = os.path.join(RAIZ, 'testdata', 'publicos')
INVENTARIO = os.path.join(PUBLICOS, 'INVENTARIO.json')
NS_CFDI = ('http://www.sat.gob.mx/cfd/3', 'http://www.sat.gob.mx/cfd/4')

# RFC genéricos y de prueba del SAT y de las librerías de origen.
RFC_PRUEBA = {'EKU9003173C9', 'XAXX010101000', 'XEXX010101000', 'AAA010101AAA',
              'XIA190128J61', 'CACX7605101P8', 'XOJI740919U48'}
RE_RFC = re.compile(rb'(?<![A-Za-z0-9&\xd1])([A-Z&\xd1]{3,4}\d{6}[A-Z0-9]{3})(?![A-Za-z0-9])')
# NumCert es el de Retenciones 1.0.
RE_NOCERT = re.compile(rb'(?i)\b(?:NoCertificado|NumCert)="([0-9]+)"')
# Ejemplos que publica el SAT (cfdv33-*, cfdv40-*…): sus RFC se aceptan.
RE_EJEMPLO_SAT = re.compile(r'cfdv(3\d|40)-')


def ln(tag):
    return tag.split('}', 1)[1] if '}' in tag else tag


def nsof(tag):
    return tag[1:].split('}', 1)[0] if tag.startswith('{') else ''


def analizar(data):
    m = re.match(rb'^(?:\xef\xbb\xbf)?\s*<\?xml[^>]*encoding=["\']([^"\']+)["\']', data)
    info = {
        'bom': data.startswith(b'\xef\xbb\xbf'),
        'encoding_declarado': m.group(1).decode() if m else None,
        'declaracion_xml': bool(re.match(rb'^(?:\xef\xbb\xbf)?\s*<\?xml', data)),
    }
    try:
        root = ET.fromstring(data)
    except ET.ParseError as e:
        info['parse_error'] = str(e)
        mm = re.search(rb'xmlns(?::\w+)?="(http://www\.sat\.gob\.mx/cfd/[34])"', data)
        info['ns'] = mm.group(1).decode() if mm else None
        return info
    ns = nsof(root.tag)
    a = root.attrib
    comps, compsc, addenda, tfd = [], [], False, False
    for ch in root:
        if ln(ch.tag) == 'Complemento':
            for c in ch:
                comps.append({'ns': nsof(c.tag), 'elemento': ln(c.tag),
                              'version': c.attrib.get('Version') or c.attrib.get('version')})
                tfd = tfd or ln(c.tag) == 'TimbreFiscalDigital'
        addenda = addenda or ln(ch.tag) == 'Addenda'
    for cc in root.iter('{%s}ComplementoConcepto' % ns):
        compsc.extend({'ns': nsof(c.tag), 'elemento': ln(c.tag)} for c in cc)
    info.update({
        'ns': ns,
        'version': a.get('Version') or a.get('version'),
        'tipo': a.get('TipoDeComprobante') or a.get('tipoDeComprobante'),
        'complementos': comps,
        'complementos_concepto': compsc,
        'tfd': tfd,
        'addenda': addenda,
        'n_conceptos': len(list(root.iter('{%s}Concepto' % ns))),
        'no_certificado': a.get('NoCertificado') or a.get('noCertificado') or a.get('NumCert'),
    })
    return info


def archivos():
    for org_repo in sorted(os.listdir(PUBLICOS)):
        d = os.path.join(PUBLICOS, org_repo)
        if os.path.isdir(d):
            for f in sorted(os.listdir(d)):
                if f.lower().endswith('.xml'):
                    yield f'{org_repo}/{f}', os.path.join(d, f)


def revisar_rfc(datos):
    """datos: {ruta: bytes}. Devuelve (reporte, rechazados).

    Rechaza un archivo con certificado de producción (NoCertificado 00001…) y algún RFC de
    persona física (13 caracteres) que no sea de prueba ni salga en los ejemplos del SAT.
    """
    rfc_por_archivo = {r: {m.decode('latin1') for m in RE_RFC.findall(b)} for r, b in datos.items()}
    del_sat = {x for r, s in rfc_por_archivo.items() if RE_EJEMPLO_SAT.search(r) for x in s}
    conocidos = RFC_PRUEBA | del_sat
    desconocidos = {r: s - conocidos for r, s in rfc_por_archivo.items()}
    rechazados, prod = [], []
    for r, b in datos.items():
        certs = [c.decode() for c in RE_NOCERT.findall(b)]
        if any(c.startswith('00001') for c in certs):
            prod.append(r)
            if any(len(x) == 13 for x in desconocidos[r]):
                rechazados.append(r)
    todos = set().union(*desconocidos.values())
    reporte = {
        'archivos': len(datos),
        'rfc_desconocidos': len(todos),
        'rfc_desconocidos_pf': sum(1 for x in todos if len(x) == 13),
        'archivos_con_rfc_desconocido': sum(1 for s in desconocidos.values() if s),
        'archivos_certificado_produccion': sorted(prod),
        'rechazados': sorted(rechazados),
    }
    return reporte, rechazados


def main():
    args = sys.argv[1:]
    origenes_path = INVENTARIO
    if '--origenes' in args:
        origenes_path = args[args.index('--origenes') + 1]
    procedencia = {}
    if os.path.exists(origenes_path):
        for e in json.load(open(origenes_path, encoding='utf-8'))['archivos']:
            procedencia[e['sha256']] = e

    datos = {r: open(p, 'rb').read() for r, p in archivos()}
    reporte, rechazados = revisar_rfc(datos)
    print(json.dumps(reporte, ensure_ascii=False, indent=2), file=sys.stderr)
    if '--revisar' in args:
        return
    if rechazados:
        sys.exit('hay archivos con certificado de producción y RFC de persona física desconocido')

    inv = []
    for r, data in datos.items():
        h = hashlib.sha256(data).hexdigest()
        p = procedencia.get(h)
        if p is None:
            sys.exit(f'{r}: sin procedencia conocida (sha256 {h})')
        entrada = {'archivo': r, 'sha256': h, 'bytes': len(data)}
        for k in ('repo', 'commit', 'ruta_original', 'licencia', 'licencia_archivo'):
            entrada[k] = p[k]
        if p.get('duplicados'):
            entrada['duplicados'] = [{k: d[k] for k in ('repo', 'commit', 'ruta_original')} for d in p['duplicados']]
        entrada.update(analizar(data))
        entrada['certificado_produccion'] = (entrada.get('no_certificado') or '').startswith('00001')
        inv.append(entrada)
    with open(INVENTARIO, 'w', encoding='utf-8') as f:
        json.dump({'generado_por': 'scripts/generar_inventario.py', 'total': len(inv), 'archivos': inv},
                  f, ensure_ascii=False, indent=2)
        f.write('\n')
    print('OK', len(inv), file=sys.stderr)


main()
