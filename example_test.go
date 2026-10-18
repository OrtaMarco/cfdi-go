package cfdi_test

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/OrtaMarco/cfdi-go"
)

// Una factura 4.0 mínima, sin sello ni timbre, con el RFC de prueba del SAT.
const facturaEjemplo = `<?xml version="1.0" encoding="UTF-8"?>
<cfdi:Comprobante xmlns:cfdi="http://www.sat.gob.mx/cfd/4" Version="4.0" Serie="A" Folio="1"
    Fecha="2026-01-15T10:30:00" Sello="" NoCertificado="30001000000500003416" Certificado=""
    SubTotal="100.00" Moneda="MXN" Total="116.00" TipoDeComprobante="I" Exportacion="01"
    MetodoPago="PUE" FormaPago="03" LugarExpedicion="20000">
  <cfdi:Emisor Rfc="EKU9003173C9" Nombre="ESCUELA KEMPER URGATE" RegimenFiscal="601"/>
  <cfdi:Receptor Rfc="XAXX010101000" Nombre="PUBLICO EN GENERAL" DomicilioFiscalReceptor="20000"
      RegimenFiscalReceptor="616" UsoCFDI="S01"/>
  <cfdi:Conceptos>
    <cfdi:Concepto ClaveProdServ="01010101" Cantidad="1" ClaveUnidad="H87" Descripcion="Artículo"
        ValorUnitario="100.00" Importe="100.00" ObjetoImp="02">
      <cfdi:Impuestos>
        <cfdi:Traslados>
          <cfdi:Traslado Base="100.00" Impuesto="002" TipoFactor="Tasa" TasaOCuota="0.160000" Importe="16.00"/>
        </cfdi:Traslados>
      </cfdi:Impuestos>
    </cfdi:Concepto>
  </cfdi:Conceptos>
  <cfdi:Impuestos TotalImpuestosTrasladados="16.00">
    <cfdi:Traslados>
      <cfdi:Traslado Base="100.00" Impuesto="002" TipoFactor="Tasa" TasaOCuota="0.160000" Importe="16.00"/>
    </cfdi:Traslados>
  </cfdi:Impuestos>
</cfdi:Comprobante>`

func ExampleParse() {
	c, err := cfdi.Parse(strings.NewReader(facturaEjemplo))
	if err != nil {
		log.Fatal(cfdi.ClaseDe(err))
	}
	fmt.Println(c.Version, c.TipoDeComprobante, c.Emisor.Rfc, c.Total.Original())
	for _, a := range c.Advertencias {
		fmt.Println(a.Codigo, a.Ruta)
	}
	// Output:
	// 4.0 I EKU9003173C9 116.00
	// SIN_TIMBRE /Comprobante/Complemento
}

func ExampleComprobante_JSON() {
	c, err := cfdi.ParseBytes([]byte(facturaEjemplo))
	if err != nil {
		log.Fatal(err)
	}
	b, err := c.JSON()
	if err != nil {
		log.Fatal(err)
	}
	inicio, _, _ := strings.Cut(string(b), `"serie"`)
	fmt.Println(inicio + "…")
	// Output:
	// {"schema_version":"1","cfdi":{"version":"4.0",…
}

func ExampleValidar() {
	v := cfdi.Validar([]byte(facturaEjemplo), nil)
	fmt.Println(v.Estado)
	for _, h := range v.Hallazgos() {
		fmt.Println(h.Nivel, h.Codigo, h.Ruta)
	}
	// Output:
	// valido
}

func ExampleCadenaOriginal() {
	cadena, err := cfdi.CadenaOriginal([]byte(facturaEjemplo))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cadena[:40] + "…")
	// Output:
	// ||4.0|A|1|2026-01-15T10:30:00|03|3000100…
}

// Verificar arma la cadena original y comprueba el sello del emisor con el
// certificado que trae el XML, sin red. Este CFDI está sellado con el CSD de
// pruebas del SAT.
func ExampleVerificar() {
	xml, err := os.ReadFile("testdata/sinteticos/sellados/ingreso40-pue.xml")
	if err != nil {
		log.Fatal(err)
	}
	v := cfdi.Verificar(xml, nil)
	fmt.Println(v.Sello.Estado, v.Sello.Algoritmo, v.Sello.NoCertificado)
	// Output:
	// valido SHA256withRSA 30001000000500003416
}
