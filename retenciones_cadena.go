package cfdi

import (
	"crypto"
	"fmt"
)

// MotivoNoEsRetenciones es el motivo de CadenaOriginalRetenciones (y de un
// sello no verificable de VerificarRetenciones) cuando la raíz no es un
// Retenciones de retencionpago/1 o retencionpago/2.
const MotivoNoEsRetenciones = "NO_ES_RETENCIONES"

// CadenaOriginalRetenciones arma la cadena original de un CFDI de
// Retenciones e información de pagos 1.0 o 2.0, como CadenaOriginal la del
// CFDI: con las hojas retenciones.xslt del SAT (traducidas en
// cadena_reglas.go), que incluyen los trece complementos de la 1.0
// (ServiciosPlataformasTecnologicas, Intereses, Pagosaextranjeros,
// Dividendos, EnajenaciondeAcciones, Arrendamientoenfideicomiso,
// Operacionesconderivados, Intereseshipotecarios, Fideicomisonoempresarial,
// Planesderetiro 1.0 y 1.1, Premios y SectorFinanciero). La hoja 2.0 no
// incluye Planesderetiro 1.0. Cualquier otro complemento, o uno que la hoja
// de esa versión no incluye, da *ErrorCadena con COMPLEMENTO_SIN_CADENA.
//
// La cadena del TimbreFiscalDigital de una Retenciones la arma
// CadenaOriginalTFD, igual que la de un CFDI.
func CadenaOriginalRetenciones(xml []byte) (string, error) {
	raiz, err := arbolCadenaEstricto(xml)
	if err != nil {
		return "", err
	}
	return cadenaRetenciones(raiz)
}

func cadenaRetenciones(raiz *nodoC) (string, error) {
	var h *hojaCadena
	switch {
	case raiz.nombre == nombreXML{NSRetenciones20, "Retenciones"} && raiz.attr("Version") == "2.0":
		h = hojaRetenciones20
	case raiz.nombre == nombreXML{NSRetenciones10, "Retenciones"} && raiz.attr("Version") == "1.0":
		h = hojaRetenciones10
	case raiz.nombre.local != "Retenciones" || (raiz.nombre.ns != NSRetenciones10 && raiz.nombre.ns != NSRetenciones20):
		return "", &ErrorCadena{Motivo: MotivoNoEsRetenciones, Nodo: raiz.nombre.String()}
	default:
		return "", &ErrorCadena{Motivo: MotivoVersionNoSoportada, Err: fmt.Errorf("version %q en %s", raiz.attr("Version"), raiz.nombre.ns)}
	}
	return h.evaluar(raiz)
}

// perfilRetenciones: 1.0 se llama Cert y NumCert y se sella con SHA-1 (los
// ejemplos públicos del SAT, de 2014 a 2021, verifican con SHA-1); se acepta
// también SHA-256. 2.0, SHA-256 como CFDI 4.0.
func perfilRetenciones(raiz *nodoC) perfilSello {
	p := perfilSello{cadena: cadenaRetenciones, sello: "Sello", cert: "Certificado", noCert: "NoCertificado",
		algoritmos: []crypto.Hash{crypto.SHA256}, fecha: "FechaExp", rfcEmisor: "RfcE"}
	if raiz.nombre.ns == NSRetenciones10 {
		p.cert, p.noCert, p.rfcEmisor = "Cert", "NumCert", "RFCEmisor"
		p.algoritmos = append(p.algoritmos, crypto.SHA1)
	}
	return p
}

// VerificarRetenciones es Verificar para un CFDI de Retenciones 1.0 o 2.0:
// el sello del emisor sobre CadenaOriginalRetenciones con el certificado del
// XML (SHA-256; en 1.0 también SHA-1) y que su NoCertificado (NumCert en
// 1.0) sea el número del certificado; el timbre, si certs no es nil, como en
// Verificar; y los chequeos del certificado, con FechaExp. Nunca falla.
func VerificarRetenciones(xml []byte, certs CertificadosSAT) *Verificacion {
	return verificar(xml, certs, perfilRetenciones)
}

// ValidarRetenciones valida un CFDI de Retenciones 1.0 o 2.0 contra los XSD
// del SAT (retencionpagov1/v2, sus trece complementos y el TFD 1.1): es
// ValidarEsquema en una Validacion sin Reglas. No hay reglas aritméticas de
// Retenciones: el SAT no publica para ellas matriz de errores ni guía con las
// fórmulas, y la librería no inventa reglas sin fuente. Nunca falla.
func ValidarRetenciones(xml []byte) *Validacion {
	e := ValidarEsquema(xml)
	return &Validacion{Estado: e.Estado, Esquema: e}
}
