#!/usr/bin/env python3
"""Escribe monedas.go: los decimales de cada clave de c_Moneda.

Uso (desde la raíz del repo):

    python3 scripts/extraer_monedas.py RUTA/A/satcfdi/catalogs/catalogs.db

El XSD del SAT (catCFDI.xsd) solo trae las claves de c_Moneda; los decimales vienen en el Excel
del catálogo. La fuente aquí es la base de catálogos de SAT-CFDI/python-satcfdi (licencia MIT),
tabla C756_c_Moneda, que guarda por clave [descripción, decimales] serializado con pickle (solo
tuplas, listas, cadenas y enteros). La prueba TestMonedasCatalogo exige que las claves coincidan
con las del XSD.
"""
import os
import pickle
import sqlite3
import sys

RAIZ = os.path.abspath(os.path.join(os.path.dirname(os.path.abspath(__file__)), '..'))


class SoloBasicos(pickle.Unpickler):
    """No reconstruye ninguna clase: solo tipos básicos."""

    def find_class(self, module, name):
        raise pickle.UnpicklingError(f'clase no permitida: {module}.{name}')


def cargar(b):
    import io
    return SoloBasicos(io.BytesIO(b)).load()


def main():
    if len(sys.argv) != 2:
        raise SystemExit(__doc__)
    db = sqlite3.connect(f'file:{sys.argv[1]}?mode=ro', uri=True)
    filas = {cargar(k): cargar(v) for k, v in db.execute('SELECT key, value FROM C756_c_Moneda')}
    lineas = []
    for clave in sorted(filas):
        decimales = filas[clave][1]
        if not isinstance(decimales, int) or not 0 <= decimales <= 6:
            raise SystemExit(f'{clave}: decimales {decimales!r}')
        lineas.append(f'\t"{clave}": {decimales},')
    with open(os.path.join(RAIZ, 'monedas.go'), 'w', encoding='utf-8') as f:
        f.write('// Código generado por scripts/extraer_monedas.py a partir del catálogo c_Moneda\n')
        f.write('// de SAT-CFDI/python-satcfdi (catalogs.db, tabla C756_c_Moneda; licencia MIT).\n')
        f.write('// NO EDITAR: volver a correr el script.\n\n')
        f.write('package cfdi\n\n')
        f.write('// decimalesMoneda son los decimales que admite cada moneda del catálogo\n')
        f.write('// c_Moneda (columna "Decimales" del catálogo del SAT).\n')
        f.write('var decimalesMoneda = map[string]int32{\n')
        f.write('\n'.join(lineas) + '\n}\n')
    print(len(lineas), 'monedas')


if __name__ == '__main__':
    main()
