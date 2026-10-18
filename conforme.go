package cfdi

import (
	"encoding/xml"
	"fmt"
)

// alcanceNS revisa, elemento por elemento, lo que encoding/xml deja pasar y
// un procesador XML conforme (libxml2, el de los XSLT y XSD del SAT)
// rechaza: atributos repetidos y prefijos sin declarar (encoding/xml deja el
// prefijo como si fuera el namespace). Lo usan ValidarEsquema y la cadena
// original.
type alcanceNS []map[string]bool // namespaces declarados en cada nivel

func (a *alcanceNS) declarado(ns string) bool {
	for i := len(*a) - 1; i >= 0; i-- {
		if (*a)[i][ns] {
			return true
		}
	}
	return ns == "" || ns == "http://www.w3.org/XML/1998/namespace"
}

// entrar revisa la etiqueta de apertura t y abre su nivel.
func (a *alcanceNS) entrar(t xml.StartElement) error {
	decls := map[string]bool{}
	vistos := map[xml.Name]bool{}
	for _, at := range t.Attr {
		if vistos[at.Name] {
			return fmt.Errorf("atributo %s repetido en %s", at.Name.Local, t.Name.Local)
		}
		vistos[at.Name] = true
		// Un namespace que no es URI válido (xmlns:schemaLocation="a b") es
		// solo una advertencia de libxml2 (WAR_NS_URI): se acepta.
		if at.Name.Space == nsXMLNSAttrs || (at.Name.Space == "" && at.Name.Local == "xmlns") {
			decls[at.Value] = true
		}
	}
	*a = append(*a, decls)
	if !a.declarado(t.Name.Space) {
		return fmt.Errorf("prefijo %q sin declarar en %s", t.Name.Space, t.Name.Local)
	}
	for _, at := range t.Attr {
		if at.Name.Space != "" && at.Name.Space != nsXMLNSAttrs && !a.declarado(at.Name.Space) {
			return fmt.Errorf("prefijo %q sin declarar en el atributo %s", at.Name.Space, at.Name.Local)
		}
	}
	return nil
}

func (a *alcanceNS) salir() { *a = (*a)[:len(*a)-1] }

// saltar consume el resto del elemento que se acaba de abrir (ya pasó por
// entrar) revisando sus descendientes, sin guardar nada, y cierra su nivel.
func (a *alcanceNS) saltar(d *xml.Decoder) error {
	for prof := 1; prof > 0; {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			if err := a.entrar(t); err != nil {
				return err
			}
			prof++
		case xml.EndElement:
			a.salir()
			prof--
		}
	}
	return nil
}
