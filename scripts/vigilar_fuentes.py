#!/usr/bin/env python3
"""Vigila las fuentes oficiales de las que sale esta librería.

Las reglas, los XSD y las hojas XSLT se copian o se derivan de documentos que el SAT
cambia sin avisar. Este script compara lo que hay publicado hoy contra la foto guardada
en docs/fuentes.json y dice qué cambió. Mira tres cosas:

  documentos  URL fija (una guía, una matriz): cambia si cambian Last-Modified o el tamaño.
  indices     páginas del portal que listan los documentos: cambia si aparece o desaparece
              un enlace a un .xls/.pdf/.xsd/.xslt/.zip. Es lo único que detecta una matriz
              nueva, porque el SAT le pone la fecha al nombre del archivo y la URL vieja
              sigue respondiendo igual.
  replica     phpcfdi/resources-sat-xml, de donde se copian los XSD y las XSLT: cambia si
              desde el commit fijado en NOTICE se tocó algún archivo que aquí está copiado.

Uso:
  python3 scripts/vigilar_fuentes.py               compara e imprime un informe en Markdown
  python3 scripts/vigilar_fuentes.py --actualizar  vuelve a tomar la foto (tras revisar a mano)

Salida: 0 sin cambios y todo comprobado, 1 hay cambios, 2 sin cambios pero algo no se pudo
comprobar. Lo que no se pudo comprobar nunca se cuenta como «sin cambios».

Solo biblioteca estándar. No manda nada más que las peticiones GET/HEAD.
"""

from __future__ import annotations

import argparse
import datetime
import json
import os
import re
import sys
import urllib.error
import urllib.parse
import urllib.request

RAIZ = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
FOTO = os.path.join(RAIZ, "docs", "fuentes.json")
AGENTE = "cfdi-go-vigilancia (+https://github.com/OrtaMarco/cfdi-go)"
ESPERA = 45
ES_DOCUMENTO = re.compile(r"\.(xlsx?|pdf|xsd|xslt|zip)$", re.I)
# Carpetas de este repo que son copia de la réplica, y el prefijo que les corresponde allá.
PREFIJO_REPLICA = "resources/www.sat.gob.mx/"
COPIAS = ("internal/genxsd/xsd", "internal/gencadena/xslt")


class NoComprobado(Exception):
    pass


def pedir(url: str, metodo: str = "GET", cabeceras: dict | None = None):
    peticion = urllib.request.Request(url, method=metodo, headers={"User-Agent": AGENTE, **(cabeceras or {})})
    try:
        return urllib.request.urlopen(peticion, timeout=ESPERA)
    except urllib.error.HTTPError as e:
        raise NoComprobado(f"HTTP {e.code}") from e
    except (urllib.error.URLError, TimeoutError, OSError) as e:
        raise NoComprobado(str(getattr(e, "reason", e))) from e


def foto_documento(url: str) -> dict:
    with pedir(url, "HEAD") as r:
        return {
            "last_modified": r.headers.get("Last-Modified", ""),
            "content_length": r.headers.get("Content-Length", ""),
        }


def enlaces_de(url: str) -> list[str]:
    with pedir(url) as r:
        html = r.read().decode("utf-8", "replace")
    vistos = set()
    for href in re.findall(r"""href\s*=\s*["']([^"']+)["']""", html, flags=re.I):
        href = href.strip().split("#")[0]
        if not ES_DOCUMENTO.search(href.split("?")[0]):
            continue
        absoluto = urllib.parse.urljoin(url, href)
        # El portal mezcla http y https y mayúsculas en el host para el mismo archivo.
        partes = urllib.parse.urlsplit(absoluto)
        vistos.add(urllib.parse.urlunsplit(("", partes.netloc.lower(), partes.path, partes.query, "")).lstrip("/"))
    if not vistos:
        # Una página que antes listaba documentos y ahora ninguno es un portal caído o
        # rediseñado, no «se retiraron todos»: no se puede leer como un cambio fiable.
        raise NoComprobado("la página no lista ningún documento")
    return sorted(vistos)


def copiados_localmente() -> set[str]:
    rutas = set()
    for copia in COPIAS:
        base = os.path.join(RAIZ, copia)
        for carpeta, _, archivos in os.walk(base):
            for archivo in archivos:
                rutas.add(os.path.relpath(os.path.join(carpeta, archivo), base).replace(os.sep, "/"))
    return rutas


def cambios_replica(repo: str, commit: str, rama: str) -> tuple[str, list[str]]:
    cabeceras = {"Accept": "application/vnd.github+json"}
    token = os.environ.get("GITHUB_TOKEN") or os.environ.get("GH_TOKEN")
    if token:
        cabeceras["Authorization"] = f"Bearer {token}"
    with pedir(f"https://api.github.com/repos/{repo}/compare/{commit}...{rama}", cabeceras=cabeceras) as r:
        datos = json.load(r)
    cabeza = datos["commits"][-1]["sha"] if datos.get("commits") else commit
    archivos = datos.get("files") or []
    if len(archivos) >= 300:
        # La API corta la lista en 300 archivos: con tantos no se sabe cuáles faltan.
        raise NoComprobado("la comparación trae 300 archivos o más y la API la trunca")
    locales = copiados_localmente()
    tocados = []
    for archivo in archivos:
        nombre = archivo["filename"]
        if nombre.startswith(PREFIJO_REPLICA) and nombre[len(PREFIJO_REPLICA):] in locales:
            tocados.append(f'{nombre[len(PREFIJO_REPLICA):]} ({archivo["status"]})')
    return cabeza, sorted(tocados)


def tomar_foto(foto: dict) -> None:
    for doc in foto["documentos"]:
        doc.update(foto_documento(doc["url"]))
        print("documento", doc["nombre"], file=sys.stderr)
    for indice in foto["indices"]:
        indice["enlaces"] = enlaces_de(indice["url"])
        print("índice", indice["nombre"], len(indice["enlaces"]), file=sys.stderr)
    foto["consultado"] = datetime.date.today().isoformat()
    with open(FOTO, "w", encoding="utf-8") as f:
        json.dump(foto, f, ensure_ascii=False, indent=2)
        f.write("\n")


def comparar(foto: dict) -> tuple[list[str], list[str], int]:
    cambios: list[str] = []
    ciegos: list[str] = []
    comprobados = 0

    for doc in foto["documentos"]:
        try:
            ahora = foto_documento(doc["url"])
        except NoComprobado as e:
            ciegos.append(f'{doc["nombre"]}: {e} — {doc["url"]}')
            continue
        comprobados += 1
        distintos = [
            f'{campo} «{doc.get(campo, "")}» → «{ahora[campo]}»'
            for campo in ("last_modified", "content_length")
            if ahora[campo] != doc.get(campo, "")
        ]
        if distintos:
            cambios.append(f'**{doc["nombre"]}** cambió ({"; ".join(distintos)}) — {doc["url"]}')

    for indice in foto["indices"]:
        try:
            ahora = set(enlaces_de(indice["url"]))
        except NoComprobado as e:
            ciegos.append(f'{indice["nombre"]}: {e} — {indice["url"]}')
            continue
        comprobados += 1
        antes = set(indice.get("enlaces", []))
        for enlace in sorted(ahora - antes):
            cambios.append(f'**{indice["nombre"]}** enlaza un documento nuevo: `{enlace}`')
        for enlace in sorted(antes - ahora):
            cambios.append(f'**{indice["nombre"]}** ya no enlaza: `{enlace}`')

    replica = foto["replica"]
    try:
        cabeza, tocados = cambios_replica(replica["repo"], replica["commit"], replica["rama"])
        comprobados += 1
        for archivo in tocados:
            cambios.append(
                f'**Réplica {replica["repo"]}** cambió un archivo copiado aquí: `{archivo}` '
                f'(de `{replica["commit"][:7]}` a `{cabeza[:7]}`)'
            )
    except NoComprobado as e:
        ciegos.append(f'Réplica {replica["repo"]}: {e}')
    except (KeyError, ValueError) as e:
        ciegos.append(f'Réplica {replica["repo"]}: respuesta inesperada ({e})')

    return cambios, ciegos, comprobados


def main() -> int:
    parser = argparse.ArgumentParser(description=__doc__.split("\n")[0])
    parser.add_argument("--actualizar", action="store_true", help="vuelve a tomar la foto en docs/fuentes.json")
    args = parser.parse_args()

    with open(FOTO, encoding="utf-8") as f:
        foto = json.load(f)

    if args.actualizar:
        tomar_foto(foto)
        return 0

    cambios, ciegos, comprobados = comparar(foto)
    total = len(foto["documentos"]) + len(foto["indices"]) + 1
    print(f'Foto del {foto["consultado"]}. Comprobadas {comprobados} de {total} fuentes.\n')
    if cambios:
        print(f"## Cambios ({len(cambios)})\n")
        print("\n".join(f"- {c}" for c in cambios))
        print()
    if ciegos:
        print(f"## No se pudo comprobar ({len(ciegos)})\n")
        print("\n".join(f"- {c}" for c in ciegos))
        print()
    if not cambios and not ciegos:
        print("Sin cambios.")
    if cambios:
        return 1
    return 2 if ciegos else 0


if __name__ == "__main__":
    sys.exit(main())
