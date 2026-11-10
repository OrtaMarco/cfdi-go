package cfdi

import (
	"encoding/json"
	"runtime"
	"strings"
	"testing"
)

// Detallista lleva datos en el texto de los elementos. Los públicos de satcfdi
// y nodecfdi/phpcfdi traen la mayor parte de los nodos; ninguno trae todos, así
// que detallistaSintetico arma uno completo (GLN y GTIN inventados) que se
// mete en un 3.2 sintético, como ecb32.

const (
	detCompleto40 = "publicos/SAT-CFDI-python-satcfdi/tests__cfdi_ejemplos__comprobante40__cfdv40-detallista.xml"
	detMin40      = "publicos/SAT-CFDI-python-satcfdi/tests__cfdi_ejemplos__comprobante40__cfdv40-detallista-min.xml"
	detNodecfdi   = "publicos/nodecfdi-cfdi-to-json/tests___files__detallista-example.xml"
	rutaDet       = "/Comprobante/Complemento/detallista"
)

var detallistaPublicos = []string{
	detCompleto40,
	detMin40,
	"publicos/SAT-CFDI-python-satcfdi/tests__cfdi_ejemplos__comprobante33__cfdv33-detallista.xml",
	"publicos/SAT-CFDI-python-satcfdi/tests__cfdi_ejemplos__comprobante33__cfdv33-detallista-min.xml",
	"publicos/SAT-CFDI-python-satcfdi/tests__test_cfdi__comprobante33__cfdv33-detallista.xml.xml",
	"publicos/SAT-CFDI-python-satcfdi/tests__test_cfdi__comprobante33__cfdv33-detallista-min.xml.xml",
	"publicos/SAT-CFDI-python-satcfdi/tests__test_cfdi__comprobante40__cfdv40-detallista.xml.xml",
	"publicos/SAT-CFDI-python-satcfdi/tests__test_cfdi__comprobante40__cfdv40-detallista-min.xml.xml",
	detNodecfdi,
	"publicos/phpcfdi-cfdi-to-json/tests___files__detallista-example.xml",
}

// detallistaSintetico trae todos los nodos y atributos del XSD, con valores de
// sus enumeraciones.
const detallistaSintetico = `<detallista:detallista xmlns:detallista="http://www.sat.gob.mx/detallista" type="SimpleInvoiceType" contentVersion="1.3.1" documentStructureVersion="AMC8.1" documentStatus="COPY">
      <detallista:requestForPaymentIdentification>
        <detallista:entityType>CREDIT_NOTE</detallista:entityType>
      </detallista:requestForPaymentIdentification>
      <detallista:specialInstruction code="AAB">
        <detallista:text>PRIMERA</detallista:text>
        <detallista:text>SEGUNDA</detallista:text>
      </detallista:specialInstruction>
      <detallista:specialInstruction code="PUR">
        <detallista:text>UNICA</detallista:text>
      </detallista:specialInstruction>
      <detallista:orderIdentification>
        <detallista:referenceIdentification type="ON">OC-001</detallista:referenceIdentification>
        <detallista:referenceIdentification type="ON">OC-002</detallista:referenceIdentification>
        <detallista:ReferenceDate>2020-01-15</detallista:ReferenceDate>
      </detallista:orderIdentification>
      <detallista:AdditionalInformation>
        <detallista:referenceIdentification type="IV">F-1</detallista:referenceIdentification>
        <detallista:referenceIdentification type="ATZ">99</detallista:referenceIdentification>
      </detallista:AdditionalInformation>
      <detallista:DeliveryNote>
        <detallista:referenceIdentification>REC-1</detallista:referenceIdentification>
        <detallista:referenceIdentification>REC-2</detallista:referenceIdentification>
        <detallista:ReferenceDate>2020-01-16</detallista:ReferenceDate>
      </detallista:DeliveryNote>
      <detallista:buyer>
        <detallista:gln>7500000000011</detallista:gln>
        <detallista:contactInformation>
          <detallista:personOrDepartmentName>
            <detallista:text>COMPRAS</detallista:text>
          </detallista:personOrDepartmentName>
        </detallista:contactInformation>
      </detallista:buyer>
      <detallista:seller>
        <detallista:gln>7500000000028</detallista:gln>
        <detallista:alternatePartyIdentification type="IEPS_REFERENCE">PROV-7</detallista:alternatePartyIdentification>
      </detallista:seller>
      <detallista:shipTo>
        <detallista:gln>7500000000035</detallista:gln>
        <detallista:nameAndAddress>
          <detallista:name>BODEGA UNO</detallista:name>
          <detallista:streetAddressOne>CALLE 1</detallista:streetAddressOne>
          <detallista:city>CIUDAD A</detallista:city>
          <detallista:postalCode>01000</detallista:postalCode>
          <detallista:name>BODEGA DOS</detallista:name>
          <detallista:postalCode>02000</detallista:postalCode>
        </detallista:nameAndAddress>
      </detallista:shipTo>
      <detallista:InvoiceCreator>
        <detallista:gln>7500000000042</detallista:gln>
        <detallista:alternatePartyIdentification type="VA">EMISOR-3</detallista:alternatePartyIdentification>
        <detallista:nameAndAddress>
          <detallista:name>EMISOR DE PRUEBA</detallista:name>
          <detallista:streetAddressOne>CALLE 3</detallista:streetAddressOne>
          <detallista:city>CIUDAD C</detallista:city>
          <detallista:postalCode>03000</detallista:postalCode>
        </detallista:nameAndAddress>
      </detallista:InvoiceCreator>
      <detallista:Customs>
        <detallista:gln>7500000000059</detallista:gln>
      </detallista:Customs>
      <detallista:Customs>
        <detallista:gln>7500000000066</detallista:gln>
      </detallista:Customs>
      <detallista:currency currencyISOCode="MXN">
        <detallista:currencyFunction>BILLING_CURRENCY</detallista:currencyFunction>
        <detallista:currencyFunction>PRICE_CURRENCY</detallista:currencyFunction>
      </detallista:currency>
      <detallista:currency currencyISOCode="USD">
        <detallista:currencyFunction>PAYMENT_CURRENCY</detallista:currencyFunction>
        <detallista:rateOfChange>17.250000</detallista:rateOfChange>
      </detallista:currency>
      <detallista:paymentTerms paymentTermsEvent="DATE_OF_INVOICE" PaymentTermsRelationTime="REFERENCE_AFTER">
        <detallista:netPayment netPaymentTermsType="BASIC_NET">
          <detallista:paymentTimePeriod>
            <detallista:timePeriodDue timePeriod="DAYS">
              <detallista:value>30</detallista:value>
            </detallista:timePeriodDue>
          </detallista:paymentTimePeriod>
        </detallista:netPayment>
        <detallista:discountPayment discountType="ALLOWANCE_BY_PAYMENT_ON_TIME">
          <detallista:percentage>2.5</detallista:percentage>
        </detallista:discountPayment>
      </detallista:paymentTerms>
      <detallista:shipmentDetail/>
      <detallista:allowanceCharge allowanceChargeType="ALLOWANCE_GLOBAL" settlementType="BILL_BACK" sequenceNumber="1">
        <detallista:specialServicesType>AJ</detallista:specialServicesType>
        <detallista:monetaryAmountOrPercentage>
          <detallista:rate base="INVOICE_VALUE">
            <detallista:percentage>3.00</detallista:percentage>
          </detallista:rate>
        </detallista:monetaryAmountOrPercentage>
      </detallista:allowanceCharge>
      <detallista:allowanceCharge allowanceChargeType="CHARGE_GLOBAL" settlementType="OFF_INVOICE">
        <detallista:specialServicesType>FC</detallista:specialServicesType>
      </detallista:allowanceCharge>
      <detallista:lineItem type="SimpleInvoiceLineItemType" number="1">
        <detallista:tradeItemIdentification>
          <detallista:gtin>07500000000103</detallista:gtin>
        </detallista:tradeItemIdentification>
        <detallista:alternateTradeItemIdentification type="BUYER_ASSIGNED">B-1</detallista:alternateTradeItemIdentification>
        <detallista:alternateTradeItemIdentification type="SERIAL_NUMBER">S-1</detallista:alternateTradeItemIdentification>
        <detallista:tradeItemDescriptionInformation language="EN">
          <detallista:longText>ARTICULO DE PRUEBA</detallista:longText>
        </detallista:tradeItemDescriptionInformation>
        <detallista:invoicedQuantity unitOfMeasure="PCE">10.50</detallista:invoicedQuantity>
        <detallista:aditionalQuantity QuantityType="NUM_CONSUMER_UNITS">3.0</detallista:aditionalQuantity>
        <detallista:aditionalQuantity QuantityType="FREE_GOODS">1</detallista:aditionalQuantity>
        <detallista:grossPrice>
          <detallista:Amount>40.00</detallista:Amount>
        </detallista:grossPrice>
        <detallista:netPrice>
          <detallista:Amount>34.320000</detallista:Amount>
        </detallista:netPrice>
        <detallista:AdditionalInformation>
          <detallista:referenceIdentification type="ON">LIN-OC</detallista:referenceIdentification>
        </detallista:AdditionalInformation>
        <detallista:Customs>
          <detallista:gln>7500000000073</detallista:gln>
          <detallista:alternatePartyIdentification type="TN">PED-0001</detallista:alternatePartyIdentification>
          <detallista:ReferenceDate>2019-12-31</detallista:ReferenceDate>
          <detallista:nameAndAddress>
            <detallista:name>ADUANA DE PRUEBA</detallista:name>
          </detallista:nameAndAddress>
        </detallista:Customs>
        <detallista:LogisticUnits>
          <detallista:serialShippingContainerCode type="SRV">SSCC-1</detallista:serialShippingContainerCode>
        </detallista:LogisticUnits>
        <detallista:palletInformation>
          <detallista:palletQuantity>4</detallista:palletQuantity>
          <detallista:description type="PALLET_80x100">TARIMA</detallista:description>
          <detallista:transport>
            <detallista:methodOfPayment>PREPAID_BY_SELLER</detallista:methodOfPayment>
          </detallista:transport>
        </detallista:palletInformation>
        <detallista:extendedAttributes>
          <detallista:lotNumber productionDate="2019-11-30">LOTE-A</detallista:lotNumber>
          <detallista:lotNumber>LOTE-B</detallista:lotNumber>
        </detallista:extendedAttributes>
        <detallista:allowanceCharge allowanceChargeType="ALLOWANCE_GLOBAL" settlementType="OFF_INVOICE" sequenceNumber="2">
          <detallista:specialServicesType>TD</detallista:specialServicesType>
          <detallista:monetaryAmountOrPercentage>
            <detallista:percentagePerUnit>5</detallista:percentagePerUnit>
            <detallista:ratePerUnit>
              <detallista:amountPerUnit>1.5</detallista:amountPerUnit>
            </detallista:ratePerUnit>
          </detallista:monetaryAmountOrPercentage>
        </detallista:allowanceCharge>
        <detallista:allowanceCharge allowanceChargeType="CHARGE_GLOBAL">
          <detallista:monetaryAmountOrPercentage>
            <detallista:percentagePerUnit>2</detallista:percentagePerUnit>
          </detallista:monetaryAmountOrPercentage>
        </detallista:allowanceCharge>
        <detallista:tradeItemTaxInformation>
          <detallista:taxTypeDescription>VAT</detallista:taxTypeDescription>
          <detallista:referenceNumber>IVA-16</detallista:referenceNumber>
          <detallista:tradeItemTaxAmount>
            <detallista:taxPercentage>16.00</detallista:taxPercentage>
            <detallista:taxAmount>57.66</detallista:taxAmount>
          </detallista:tradeItemTaxAmount>
          <detallista:taxCategory>TRANSFERIDO</detallista:taxCategory>
        </detallista:tradeItemTaxInformation>
        <detallista:tradeItemTaxInformation>
          <detallista:taxTypeDescription>OTH</detallista:taxTypeDescription>
        </detallista:tradeItemTaxInformation>
        <detallista:totalLineAmount>
          <detallista:grossAmount>
            <detallista:Amount>420.00</detallista:Amount>
          </detallista:grossAmount>
          <detallista:netAmount>
            <detallista:Amount>360.36</detallista:Amount>
          </detallista:netAmount>
        </detallista:totalLineAmount>
      </detallista:lineItem>
      <detallista:lineItem>
        <detallista:tradeItemIdentification>
          <detallista:gtin>7500000000202</detallista:gtin>
        </detallista:tradeItemIdentification>
        <detallista:invoicedQuantity unitOfMeasure="KGM">2</detallista:invoicedQuantity>
        <detallista:totalLineAmount>
          <detallista:netAmount>
            <detallista:Amount>20.00</detallista:Amount>
          </detallista:netAmount>
        </detallista:totalLineAmount>
      </detallista:lineItem>
      <detallista:totalAmount>
        <detallista:Amount>380.36</detallista:Amount>
      </detallista:totalAmount>
      <detallista:TotalAllowanceCharge allowanceOrChargeType="ALLOWANCE">
        <detallista:specialServicesType>AJ</detallista:specialServicesType>
        <detallista:Amount>40.00</detallista:Amount>
      </detallista:TotalAllowanceCharge>
      <detallista:TotalAllowanceCharge allowanceOrChargeType="CHARGE">
        <detallista:Amount>0.50</detallista:Amount>
      </detallista:TotalAllowanceCharge>
    </detallista:detallista>
    `

const detBase = "sinteticos/t10/spei32.xml"

// leerDetallista mete el sintético en el 3.2 con los cambios (pares
// viejo, nuevo) que pida la prueba.
func leerDetallista(t *testing.T, cambios ...string) *Comprobante {
	t.Helper()
	return leerMutado(t, detBase, append([]string{`<spei:Complemento_SPEI>`, detallistaSintetico + `<spei:Complemento_SPEI>`}, cambios...)...)
}

// advDet devuelve las advertencias de una ruta.
func advDet(c *Comprobante, codigo, ruta string) []Advertencia {
	var out []Advertencia
	for _, a := range c.Advertencias {
		if a.Codigo == codigo && a.Ruta == ruta {
			out = append(out, a)
		}
	}
	return out
}

// soloDet comprueba que el complemento no dejó advertencias propias.
func sinAdvertenciasDet(t *testing.T, c *Comprobante) {
	t.Helper()
	for _, a := range c.Advertencias {
		if strings.HasPrefix(a.Ruta, rutaDet) {
			t.Errorf("advertencia de detallista: %+v", a)
		}
	}
}

func TestDetallistaPublicos(t *testing.T) {
	for _, p := range detallistaPublicos {
		t.Run(p, func(t *testing.T) {
			c := leerRuta(t, p)
			soportado(t, c, "detallista")
			sinAdvertenciasDet(t, c)
			if c.Detallista == nil || c.Detallista.RequestForPaymentIdentification == nil {
				t.Fatalf("Detallista = %+v", c.Detallista)
			}
			igual(t, "EntityType", c.Detallista.RequestForPaymentIdentification.EntityType, "INVOICE")
			igual(t, "DocumentStructureVersion", c.Detallista.DocumentStructureVersion, "AMC8.1")
			igual(t, "DocumentStatus", c.Detallista.DocumentStatus, "ORIGINAL")
			cumpleSchema(t, c)
		})
	}
}

func TestDetallistaCompleto40(t *testing.T) {
	c := leerRuta(t, detCompleto40)
	x := c.Detallista
	if x == nil || len(x.LineItem) != 10 || x.OrderIdentification == nil || x.AdditionalInformation == nil || x.DeliveryNote == nil ||
		x.Buyer == nil || x.Seller == nil || len(x.AllowanceCharge) != 1 || x.TotalAmount == nil || len(x.TotalAllowanceCharge) != 1 {
		t.Fatalf("Detallista = %+v", x)
	}
	igual(t, "Type", x.Type, "SimpleInvoiceType")
	igual(t, "ContentVersion", x.ContentVersion, "1.3.1")
	if len(x.SpecialInstruction) != 1 || len(x.SpecialInstruction[0].Text) != 1 {
		t.Fatalf("SpecialInstruction = %+v", x.SpecialInstruction)
	}
	igual(t, "code", x.SpecialInstruction[0].Code, "ZZZ")
	igual(t, "text", x.SpecialInstruction[0].Text[0], "TRES MIL TRESCIENTOS NOVENTA Y TRES PESOS 04/100 M.N.")

	o := x.OrderIdentification
	if len(o.ReferenceIdentification) != 1 {
		t.Fatalf("OrderIdentification = %+v", o)
	}
	igual(t, "orden type", o.ReferenceIdentification[0].Type, "ON")
	igual(t, "orden valor", o.ReferenceIdentification[0].Valor, "15749759")
	igual(t, "orden fecha", o.ReferenceDate.Original, "2017-04-30")
	if !o.ReferenceDate.Valida() {
		t.Error("ReferenceDate no se interpretó")
	}

	ai := x.AdditionalInformation.ReferenceIdentification
	if len(ai) != 2 || ai[0] != (DetallistaReferencia{Type: "IV", Valor: "A"}) || ai[1] != (DetallistaReferencia{Type: "ATZ", Valor: "27534"}) {
		t.Errorf("AdditionalInformation = %+v", ai)
	}
	if len(x.DeliveryNote.ReferenceIdentification) != 1 {
		t.Fatalf("DeliveryNote = %+v", x.DeliveryNote)
	}
	igual(t, "folio", x.DeliveryNote.ReferenceIdentification[0], "97276954")
	igual(t, "fecha folio", x.DeliveryNote.ReferenceDate.Original, "2017-05-01")

	igual(t, "buyer gln", x.Buyer.GLN, "7504000107903")
	// El texto se lee tal cual: los espacios del final se conservan.
	igual(t, "contacto", x.Buyer.ContactInformation.PersonOrDepartmentName.Text, "0390   ")
	igual(t, "seller gln", x.Seller.GLN, "7504000107903")
	if *x.Seller.AlternatePartyIdentification != (DetallistaReferencia{Type: "SELLER_ASSIGNED_IDENTIFIER_FOR_A_PARTY", Valor: "109"}) {
		t.Errorf("seller alt = %+v", x.Seller.AlternatePartyIdentification)
	}

	a := x.AllowanceCharge[0]
	igual(t, "allowanceChargeType", a.AllowanceChargeType, "ALLOWANCE_GLOBAL")
	igual(t, "settlementType", a.SettlementType, "BILL_BACK")
	igual(t, "specialServicesType", a.SpecialServicesType, "AJ")
	igual(t, "base", a.MonetaryAmountOrPercentage.Rate.Base, "INVOICE_VALUE")
	monto(t, "rate/percentage", a.MonetaryAmountOrPercentage.Rate.Percentage, "3.00")

	l := x.LineItem[0]
	igual(t, "line type", l.Type, "SimpleInvoiceLineItemType")
	igual(t, "line number", l.Number, "1")
	igual(t, "gtin", l.TradeItemIdentification.GTIN, "00636398")
	if len(l.AlternateTradeItemIdentification) != 1 || l.AlternateTradeItemIdentification[0] != (DetallistaReferencia{Type: "SUPPLIER_ASSIGNED", Valor: "00636398"}) {
		t.Errorf("alternate = %+v", l.AlternateTradeItemIdentification)
	}
	igual(t, "language", l.TradeItemDescriptionInformation.Language, "ES")
	igual(t, "longText", l.TradeItemDescriptionInformation.LongText, "cacahuate acaramelado")
	igual(t, "unitOfMeasure", l.InvoicedQuantity.UnitOfMeasure, "PCE")
	monto(t, "invoicedQuantity", l.InvoicedQuantity.Valor, "10.00")
	monto(t, "grossPrice", l.GrossPrice.Amount, "40.00")
	monto(t, "netPrice", l.NetPrice.Amount, "34.32")
	monto(t, "grossAmount", l.TotalLineAmount.GrossAmount.Amount, "400.00")
	monto(t, "netAmount", l.TotalLineAmount.NetAmount.Amount, "343.17")
	igual(t, "última línea", x.LineItem[9].Number, "10")
	igual(t, "última gtin", x.LineItem[9].TradeItemIdentification.GTIN, "37715123")
	monto(t, "última neto", x.LineItem[9].TotalLineAmount.NetAmount.Amount, "428.93")

	monto(t, "totalAmount", x.TotalAmount.Amount, "3955.00")
	igual(t, "allowanceOrChargeType", x.TotalAllowanceCharge[0].AllowanceOrChargeType, "ALLOWANCE")
	igual(t, "TotalAllowanceCharge/specialServicesType", x.TotalAllowanceCharge[0].SpecialServicesType, "AJ")
	monto(t, "TotalAllowanceCharge/Amount", x.TotalAllowanceCharge[0].Amount, "561.96")

	// Lo que este ejemplo no trae no aparece.
	if x.ShipTo != nil || x.InvoiceCreator != nil || x.PaymentTerms != nil || x.ShipmentDetail != nil || x.Currency != nil || x.Customs != nil {
		t.Errorf("nodos que no vienen: %+v", x)
	}
}

func TestDetallistaMin40(t *testing.T) {
	c := leerRuta(t, detMin40)
	x := c.Detallista
	if x == nil || len(x.Currency) != 2 || len(x.AdditionalInformation.ReferenceIdentification) != 1 {
		t.Fatalf("Detallista = %+v", x)
	}
	igual(t, "moneda 1", x.Currency[0].CurrencyISOCode, "MXN")
	igual(t, "función 1", strings.Join(x.Currency[0].CurrencyFunction, ","), "BILLING_CURRENCY")
	if x.Currency[0].RateOfChange != nil {
		t.Errorf("RateOfChange de MXN = %v", x.Currency[0].RateOfChange)
	}
	igual(t, "moneda 2", x.Currency[1].CurrencyISOCode, "USD")
	monto(t, "rateOfChange", x.Currency[1].RateOfChange, "1.000000")
	// referenceIdentification vacío: el type viene y el texto no.
	if r := x.AdditionalInformation.ReferenceIdentification[0]; r != (DetallistaReferencia{Type: "ATZ"}) {
		t.Errorf("referencia vacía = %+v", r)
	}
	// El texto con sangría y salto de línea se lee tal cual.
	v := x.OrderIdentification.ReferenceIdentification[0].Valor
	if strings.TrimSpace(v) != "4500536101" || v == "4500536101" {
		t.Errorf("orden = %q", v)
	}
	if len(x.LineItem) != 0 {
		t.Errorf("LineItem = %+v", x.LineItem)
	}
}

func TestDetallistaNodecfdi(t *testing.T) {
	c := leerRuta(t, detNodecfdi)
	sinAdvertenciasDet(t, c)
	x := c.Detallista
	if x == nil || len(x.SpecialInstruction) != 1 || x.Seller != nil || x.ShipTo != nil {
		t.Fatalf("Detallista = %+v", x)
	}
	igual(t, "texto", strings.TrimSpace(x.SpecialInstruction[0].Text[0]), "Un mil ciento sesenta pesos 00/100 m.n.")
	igual(t, "orden", x.OrderIdentification.ReferenceIdentification[0].Valor, "3")
	igual(t, "fecha", x.OrderIdentification.ReferenceDate.Original, "2017-01-01")
	igual(t, "recibo", x.DeliveryNote.ReferenceIdentification[0], "1")
}

func TestDetallistaSintetico(t *testing.T) {
	c := leerDetallista(t)
	sinMasQue(t, c, AdvVersionAntigua)
	cumpleSchema(t, c)
	soportado(t, c, "detallista")
	x := c.Detallista
	if x == nil {
		t.Fatal("sin Detallista")
	}
	igual(t, "DocumentStatus", x.DocumentStatus, "COPY")
	igual(t, "EntityType", x.RequestForPaymentIdentification.EntityType, "CREDIT_NOTE")

	if len(x.SpecialInstruction) != 2 || len(x.SpecialInstruction[0].Text) != 2 || len(x.SpecialInstruction[1].Text) != 1 {
		t.Fatalf("SpecialInstruction = %+v", x.SpecialInstruction)
	}
	igual(t, "code 2", x.SpecialInstruction[1].Code, "PUR")
	igual(t, "text 2", x.SpecialInstruction[0].Text[1], "SEGUNDA")

	if len(x.OrderIdentification.ReferenceIdentification) != 2 {
		t.Fatalf("OrderIdentification = %+v", x.OrderIdentification)
	}
	igual(t, "OC 2", x.OrderIdentification.ReferenceIdentification[1].Valor, "OC-002")
	if got := x.DeliveryNote.ReferenceIdentification; len(got) != 2 || got[1] != "REC-2" {
		t.Errorf("DeliveryNote = %+v", got)
	}
	igual(t, "DeliveryNote fecha", x.DeliveryNote.ReferenceDate.Original, "2020-01-16")
	igual(t, "contacto", x.Buyer.ContactInformation.PersonOrDepartmentName.Text, "COMPRAS")
	igual(t, "seller type", x.Seller.AlternatePartyIdentification.Type, "IEPS_REFERENCE")

	// shipTo/nameAndAddress: dos vueltas de la secuencia, una lista por campo.
	s := x.ShipTo
	igual(t, "shipTo gln", s.GLN, "7500000000035")
	na := s.NameAndAddress
	igual(t, "names", strings.Join(na.Name, "|"), "BODEGA UNO|BODEGA DOS")
	igual(t, "calles", strings.Join(na.StreetAddressOne, "|"), "CALLE 1")
	igual(t, "ciudades", strings.Join(na.City, "|"), "CIUDAD A")
	igual(t, "cp", strings.Join(na.PostalCode, "|"), "01000|02000")

	ic := x.InvoiceCreator
	igual(t, "creator gln", ic.GLN, "7500000000042")
	igual(t, "creator type", ic.AlternatePartyIdentification.Type, "VA")
	igual(t, "creator nombre", ic.NameAndAddress.Name, "EMISOR DE PRUEBA")
	igual(t, "creator cp", ic.NameAndAddress.PostalCode, "03000")

	if len(x.Customs) != 2 {
		t.Fatalf("Customs = %+v", x.Customs)
	}
	igual(t, "Customs 2", x.Customs[1].GLN, "7500000000066")

	if len(x.Currency) != 2 || len(x.Currency[0].CurrencyFunction) != 2 {
		t.Fatalf("Currency = %+v", x.Currency)
	}
	igual(t, "función 2", x.Currency[0].CurrencyFunction[1], "PRICE_CURRENCY")
	monto(t, "rateOfChange", x.Currency[1].RateOfChange, "17.250000")

	p := x.PaymentTerms
	igual(t, "paymentTermsEvent", p.PaymentTermsEvent, "DATE_OF_INVOICE")
	igual(t, "PaymentTermsRelationTime", p.PaymentTermsRelationTime, "REFERENCE_AFTER")
	igual(t, "netPaymentTermsType", p.NetPayment.NetPaymentTermsType, "BASIC_NET")
	igual(t, "timePeriod", p.NetPayment.PaymentTimePeriod.TimePeriodDue.TimePeriod, "DAYS")
	igual(t, "value", p.NetPayment.PaymentTimePeriod.TimePeriodDue.Value, "30")
	igual(t, "discountType", p.DiscountPayment.DiscountType, "ALLOWANCE_BY_PAYMENT_ON_TIME")
	igual(t, "percentage", p.DiscountPayment.Percentage, "2.5")

	if x.ShipmentDetail == nil {
		t.Error("shipmentDetail no se leyó")
	}

	if len(x.AllowanceCharge) != 2 {
		t.Fatalf("AllowanceCharge = %+v", x.AllowanceCharge)
	}
	igual(t, "sequenceNumber", x.AllowanceCharge[0].SequenceNumber, "1")
	igual(t, "charge type", x.AllowanceCharge[1].AllowanceChargeType, "CHARGE_GLOBAL")
	igual(t, "charge servicio", x.AllowanceCharge[1].SpecialServicesType, "FC")
	if x.AllowanceCharge[1].MonetaryAmountOrPercentage != nil {
		t.Errorf("monetaryAmountOrPercentage inventado: %+v", x.AllowanceCharge[1])
	}

	if len(x.LineItem) != 2 {
		t.Fatalf("LineItem = %+v", x.LineItem)
	}
	l := x.LineItem[0]
	if len(l.AlternateTradeItemIdentification) != 2 {
		t.Fatalf("alternate = %+v", l.AlternateTradeItemIdentification)
	}
	igual(t, "alternate 2", l.AlternateTradeItemIdentification[1].Type, "SERIAL_NUMBER")
	igual(t, "language", l.TradeItemDescriptionInformation.Language, "EN")
	monto(t, "invoicedQuantity", l.InvoicedQuantity.Valor, "10.50")
	if len(l.AditionalQuantity) != 2 {
		t.Fatalf("AditionalQuantity = %+v", l.AditionalQuantity)
	}
	igual(t, "QuantityType", l.AditionalQuantity[0].QuantityType, "NUM_CONSUMER_UNITS")
	monto(t, "aditionalQuantity 1", l.AditionalQuantity[0].Valor, "3.0")
	monto(t, "aditionalQuantity 2", l.AditionalQuantity[1].Valor, "1")
	monto(t, "netPrice", l.NetPrice.Amount, "34.320000")
	igual(t, "additionalInformation", l.AdditionalInformation.ReferenceIdentification.Valor, "LIN-OC")

	if len(l.Customs) != 1 {
		t.Fatalf("Customs de línea = %+v", l.Customs)
	}
	lc := l.Customs[0]
	igual(t, "customs gln", lc.GLN, "7500000000073")
	igual(t, "pedimento", lc.AlternatePartyIdentification.Valor, "PED-0001")
	igual(t, "pedimento type", lc.AlternatePartyIdentification.Type, "TN")
	igual(t, "pedimento fecha", lc.ReferenceDate.Original, "2019-12-31")
	igual(t, "aduana", lc.NameAndAddress.Name, "ADUANA DE PRUEBA")

	igual(t, "sscc", l.LogisticUnits.SerialShippingContainerCode.Valor, "SSCC-1")
	igual(t, "sscc type", l.LogisticUnits.SerialShippingContainerCode.Type, "SRV")
	pi := l.PalletInformation
	igual(t, "palletQuantity", pi.PalletQuantity, "4")
	igual(t, "pallet description", pi.Description.Valor, "TARIMA")
	igual(t, "pallet type", pi.Description.Type, "PALLET_80x100")
	igual(t, "methodOfPayment", pi.Transport.MethodOfPayment, "PREPAID_BY_SELLER")

	lots := l.ExtendedAttributes.LotNumber
	if len(lots) != 2 {
		t.Fatalf("lotNumber = %+v", lots)
	}
	igual(t, "lote A", lots[0].Valor, "LOTE-A")
	igual(t, "productionDate", lots[0].ProductionDate.Original, "2019-11-30")
	if !lots[0].ProductionDate.Valida() || lots[1].ProductionDate.Original != "" {
		t.Errorf("fechas de lote: %+v", lots)
	}

	if len(l.AllowanceCharge) != 2 {
		t.Fatalf("allowanceCharge de línea = %+v", l.AllowanceCharge)
	}
	la := l.AllowanceCharge[0]
	igual(t, "línea settlementType", la.SettlementType, "OFF_INVOICE")
	igual(t, "línea sequenceNumber", la.SequenceNumber, "2")
	igual(t, "línea servicio", la.SpecialServicesType, "TD")
	igual(t, "percentagePerUnit", la.MonetaryAmountOrPercentage.PercentagePerUnit, "5")
	igual(t, "amountPerUnit", la.MonetaryAmountOrPercentage.RatePerUnit.AmountPerUnit, "1.5")
	if lb := l.AllowanceCharge[1]; lb.SettlementType != "" || lb.MonetaryAmountOrPercentage.RatePerUnit != nil || lb.MonetaryAmountOrPercentage.PercentagePerUnit != "2" {
		t.Errorf("allowanceCharge de línea 2 = %+v", lb)
	}

	if len(l.TradeItemTaxInformation) != 2 {
		t.Fatalf("tradeItemTaxInformation = %+v", l.TradeItemTaxInformation)
	}
	ti := l.TradeItemTaxInformation[0]
	igual(t, "taxTypeDescription", ti.TaxTypeDescription, "VAT")
	igual(t, "referenceNumber", ti.ReferenceNumber, "IVA-16")
	igual(t, "taxCategory", ti.TaxCategory, "TRANSFERIDO")
	monto(t, "taxPercentage", ti.TradeItemTaxAmount.TaxPercentage, "16.00")
	monto(t, "taxAmount", ti.TradeItemTaxAmount.TaxAmount, "57.66")
	if t2 := l.TradeItemTaxInformation[1]; t2.TaxTypeDescription != "OTH" || t2.TradeItemTaxAmount != nil {
		t.Errorf("impuesto 2 = %+v", t2)
	}
	monto(t, "grossAmount", l.TotalLineAmount.GrossAmount.Amount, "420.00")
	monto(t, "netAmount", l.TotalLineAmount.NetAmount.Amount, "360.36")

	// La segunda línea es la mínima.
	m := x.LineItem[1]
	igual(t, "min gtin", m.TradeItemIdentification.GTIN, "7500000000202")
	igual(t, "min unidad", m.InvoicedQuantity.UnitOfMeasure, "KGM")
	monto(t, "min cantidad", m.InvoicedQuantity.Valor, "2")
	if m.TotalLineAmount.GrossAmount != nil || m.GrossPrice != nil || m.PalletInformation != nil || m.ExtendedAttributes != nil {
		t.Errorf("línea mínima = %+v", m)
	}

	monto(t, "totalAmount", x.TotalAmount.Amount, "380.36")
	if len(x.TotalAllowanceCharge) != 2 {
		t.Fatalf("TotalAllowanceCharge = %+v", x.TotalAllowanceCharge)
	}
	igual(t, "CHARGE", x.TotalAllowanceCharge[1].AllowanceOrChargeType, "CHARGE")
	monto(t, "total charge", x.TotalAllowanceCharge[1].Amount, "0.50")
	igual(t, "sin servicio", x.TotalAllowanceCharge[1].SpecialServicesType, "")
}

// TestDetallistaJSON comprueba las claves que salen: shipment_detail como {}
// y montos y fechas como texto.
func TestDetallistaJSON(t *testing.T) {
	c := leerDetallista(t)
	b, err := json.Marshal(c.Detallista)
	if err != nil {
		t.Fatal(err)
	}
	var m map[string]any
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatal(err)
	}
	if sd, ok := m["shipment_detail"].(map[string]any); !ok || len(sd) != 0 {
		t.Errorf("shipment_detail = %#v", m["shipment_detail"])
	}
	if _, ok := m["version"]; ok {
		t.Error("detallista no tiene versión")
	}
	li := m["line_item"].([]any)[0].(map[string]any)
	igual(t, "invoiced_quantity.valor", li["invoiced_quantity"].(map[string]any)["valor"].(string), "10.50")
	igual(t, "aditional_quantity", li["aditional_quantity"].([]any)[0].(map[string]any)["quantity_type"].(string), "NUM_CONSUMER_UNITS")
	igual(t, "rate_of_change", m["currency"].([]any)[1].(map[string]any)["rate_of_change"].(string), "17.250000")
	igual(t, "reference_date", m["order_identification"].(map[string]any)["reference_date"].(string), "2020-01-15")

	// Sin shipmentDetail no sale la clave.
	s := leerDetallista(t, `<detallista:shipmentDetail/>`, ``)
	b, _ = json.Marshal(s.Detallista)
	if strings.Contains(string(b), "shipment_detail") {
		t.Errorf("shipment_detail sin el nodo: %s", b)
	}
}

func TestDetallistaMutantes(t *testing.T) {
	t.Run("atributo requerido", func(t *testing.T) {
		m := leerDetallista(t,
			` documentStatus="COPY"`, ``,
			`<detallista:currency currencyISOCode="USD">`, `<detallista:currency>`,
			` timePeriod="DAYS"`, ``,
			`<detallista:referenceIdentification type="IV">`, `<detallista:referenceIdentification>`,
			` QuantityType="FREE_GOODS"`, ``,
			` unitOfMeasure="KGM"`, ``,
			` allowanceOrChargeType="CHARGE"`, ``)
		for _, ruta := range []string{
			rutaDet + "/@documentStatus",
			rutaDet + "/currency[2]/@currencyISOCode",
			rutaDet + "/paymentTerms/netPayment/paymentTimePeriod/timePeriodDue/@timePeriod",
			rutaDet + "/AdditionalInformation/referenceIdentification[1]/@type",
			rutaDet + "/lineItem[1]/aditionalQuantity[2]/@QuantityType",
			rutaDet + "/lineItem[2]/invoicedQuantity/@unitOfMeasure",
			rutaDet + "/TotalAllowanceCharge[2]/@allowanceOrChargeType",
		} {
			if len(advDet(m, AdvAtributoFaltante, ruta)) != 1 {
				t.Errorf("sin ATRIBUTO_FALTANTE en %s: %+v", ruta, m.Advertencias)
			}
		}
		if cs := codigos(m); cs[AdvAtributoFaltante] != 7 || len(cs) != 2 {
			t.Errorf("advertencias = %v", m.Advertencias)
		}
		// Lo que se pudo leer se lee igual.
		igual(t, "tipo de referencia faltante", m.Detallista.AdditionalInformation.ReferenceIdentification[0].Valor, "F-1")
	})

	t.Run("atributo opcional no avisa", func(t *testing.T) {
		// El settlementType de una línea y el sequenceNumber son opcionales.
		m := leerDetallista(t, ` settlementType="OFF_INVOICE" sequenceNumber="2"`, ``)
		sinMasQue(t, m, AdvVersionAntigua)
		m = leerDetallista(t, ` settlementType="BILL_BACK" sequenceNumber="1"`, ` sequenceNumber="1"`)
		if cs := codigos(m); cs[AdvAtributoFaltante] != 1 || len(cs) != 2 {
			t.Errorf("settlementType global: %v", m.Advertencias)
		}
	})

	t.Run("nodo requerido", func(t *testing.T) {
		m := leerDetallista(t,
			`<detallista:entityType>CREDIT_NOTE</detallista:entityType>`, ``,
			`<detallista:gln>7500000000011</detallista:gln>`, ``,
			`<detallista:text>UNICA</detallista:text>`, ``,
			`<detallista:alternatePartyIdentification type="IEPS_REFERENCE">PROV-7</detallista:alternatePartyIdentification>`, ``,
			`<detallista:currencyFunction>PAYMENT_CURRENCY</detallista:currencyFunction>`, ``,
			`<detallista:gtin>7500000000202</detallista:gtin>`, ``,
			`<detallista:netAmount>
            <detallista:Amount>20.00</detallista:Amount>
          </detallista:netAmount>`, ``,
			`<detallista:lotNumber productionDate="2019-11-30">LOTE-A</detallista:lotNumber>
          <detallista:lotNumber>LOTE-B</detallista:lotNumber>`, ``,
			`<detallista:percentagePerUnit>2</detallista:percentagePerUnit>`, ``)
		for _, ruta := range []string{
			rutaDet + "/requestForPaymentIdentification/entityType",
			rutaDet + "/buyer/gln",
			rutaDet + "/specialInstruction[2]/text",
			rutaDet + "/seller/alternatePartyIdentification",
			rutaDet + "/currency[2]/currencyFunction",
			rutaDet + "/lineItem[2]/tradeItemIdentification/gtin",
			rutaDet + "/lineItem[2]/totalLineAmount/netAmount",
			rutaDet + "/lineItem[1]/extendedAttributes/lotNumber",
			rutaDet + "/lineItem[1]/allowanceCharge[2]/monetaryAmountOrPercentage/percentagePerUnit",
		} {
			if len(advDet(m, AdvNodoFaltante, ruta)) != 1 {
				t.Errorf("sin NODO_FALTANTE en %s: %+v", ruta, m.Advertencias)
			}
		}
		if cs := codigos(m); cs[AdvNodoFaltante] != 9 || len(cs) != 2 {
			t.Errorf("advertencias = %v", m.Advertencias)
		}
	})

	t.Run("nodo requerido ausente del todo", func(t *testing.T) {
		m := leerDetallista(t,
			`<detallista:buyer>`, `<!--`, `</detallista:buyer>`, `-->`,
			`<detallista:totalLineAmount>
          <detallista:netAmount>
            <detallista:Amount>20.00</detallista:Amount>
          </detallista:netAmount>
        </detallista:totalLineAmount>`, ``)
		if len(advDet(m, AdvNodoFaltante, rutaDet+"/buyer")) != 1 || len(advDet(m, AdvNodoFaltante, rutaDet+"/lineItem[2]/totalLineAmount")) != 1 {
			t.Errorf("advertencias = %v", m.Advertencias)
		}
		if m.Detallista.Buyer != nil || m.Detallista.LineItem[1].TotalLineAmount != nil {
			t.Errorf("nodos inventados: %+v", m.Detallista)
		}
	})

	t.Run("decimal ilegible", func(t *testing.T) {
		m := leerDetallista(t,
			`<detallista:Amount>40.00</detallista:Amount>`, `<detallista:Amount>cuarenta</detallista:Amount>`,
			`<detallista:taxAmount>57.66</detallista:taxAmount>`, `<detallista:taxAmount>1,5</detallista:taxAmount>`,
			`<detallista:invoicedQuantity unitOfMeasure="PCE">10.50</detallista:invoicedQuantity>`, `<detallista:invoicedQuantity unitOfMeasure="PCE">diez</detallista:invoicedQuantity>`)
		casos := []struct{ ruta, valor string }{
			{rutaDet + "/lineItem[1]/grossPrice/Amount", "cuarenta"},
			{rutaDet + "/lineItem[1]/tradeItemTaxInformation[1]/tradeItemTaxAmount/taxAmount", "1,5"},
			{rutaDet + "/lineItem[1]/invoicedQuantity", "diez"},
		}
		for _, k := range casos {
			as := advDet(m, AdvDecimalInvalido, k.ruta)
			if len(as) != 1 || as[0].Atributo != "" || as[0].Valor != k.valor {
				t.Errorf("DECIMAL_INVALIDO en %s: %+v", k.ruta, m.Advertencias)
			}
		}
		if cs := codigos(m); cs[AdvDecimalInvalido] != 3 || len(cs) != 2 {
			t.Errorf("advertencias = %v", m.Advertencias)
		}
		l := m.Detallista.LineItem[0]
		if l.GrossPrice == nil || l.GrossPrice.Amount != nil || l.InvoicedQuantity.Valor != nil || l.TradeItemTaxInformation[0].TradeItemTaxAmount.TaxAmount != nil {
			t.Errorf("decimales ilegibles deben quedar nil: %+v", l)
		}
		monto(t, "el vecino se lee", l.TradeItemTaxInformation[0].TradeItemTaxAmount.TaxPercentage, "16.00")
		monto(t, "netPrice", l.NetPrice.Amount, "34.320000")
	})

	t.Run("fecha ilegible", func(t *testing.T) {
		m := leerDetallista(t,
			`<detallista:ReferenceDate>2020-01-15</detallista:ReferenceDate>`, `<detallista:ReferenceDate>15/01/2020</detallista:ReferenceDate>`,
			`productionDate="2019-11-30"`, `productionDate="ayer"`)
		as := advDet(m, AdvFechaInvalida, rutaDet+"/orderIdentification/ReferenceDate")
		if len(as) != 1 || as[0].Atributo != "" || as[0].Valor != "15/01/2020" {
			t.Errorf("FECHA_INVALIDA del elemento: %+v", m.Advertencias)
		}
		as = advDet(m, AdvFechaInvalida, rutaDet+"/lineItem[1]/extendedAttributes/lotNumber[1]/@productionDate")
		if len(as) != 1 || as[0].Atributo != "productionDate" || as[0].Valor != "ayer" {
			t.Errorf("FECHA_INVALIDA del atributo: %+v", m.Advertencias)
		}
		if cs := codigos(m); cs[AdvFechaInvalida] != 2 || len(cs) != 2 {
			t.Errorf("advertencias = %v", m.Advertencias)
		}
		if f := m.Detallista.OrderIdentification.ReferenceDate; f.Valida() || f.Original != "15/01/2020" {
			t.Errorf("fecha ilegible = %+v", f)
		}
	})

	t.Run("fecha con hora", func(t *testing.T) {
		// xs:date no lleva hora: se lee pero avisa del formato.
		m := leerDetallista(t, `<detallista:ReferenceDate>2020-01-15</detallista:ReferenceDate>`, `<detallista:ReferenceDate>2020-01-15T10:00:00</detallista:ReferenceDate>`)
		if len(advDet(m, AdvFechaFormato, rutaDet+"/orderIdentification/ReferenceDate")) != 1 || !m.Detallista.OrderIdentification.ReferenceDate.Valida() {
			t.Errorf("advertencias = %v", m.Advertencias)
		}
	})

	t.Run("duplicado", func(t *testing.T) {
		m := leerDetallista(t,
			`<detallista:shipmentDetail/>`, `<detallista:shipmentDetail/><detallista:shipmentDetail/>`,
			`<detallista:entityType>CREDIT_NOTE</detallista:entityType>`, `<detallista:entityType>CREDIT_NOTE</detallista:entityType><detallista:entityType>INVOICE</detallista:entityType>`,
			`<detallista:gtin>7500000000202</detallista:gtin>`, `<detallista:gtin>7500000000202</detallista:gtin><detallista:gtin>1</detallista:gtin>`)
		for _, ruta := range []string{
			rutaDet + "/shipmentDetail",
			rutaDet + "/requestForPaymentIdentification/entityType",
			rutaDet + "/lineItem[2]/tradeItemIdentification/gtin",
		} {
			if len(advDet(m, AdvNodoDuplicado, ruta)) != 1 {
				t.Errorf("sin NODO_DUPLICADO en %s: %+v", ruta, m.Advertencias)
			}
		}
		if cs := codigos(m); cs[AdvNodoDuplicado] != 3 || len(cs) != 2 {
			t.Errorf("advertencias = %v", m.Advertencias)
		}
		// Se usa el primero.
		igual(t, "EntityType", m.Detallista.RequestForPaymentIdentification.EntityType, "CREDIT_NOTE")
		igual(t, "gtin", m.Detallista.LineItem[1].TradeItemIdentification.GTIN, "7500000000202")
	})

	t.Run("complemento duplicado", func(t *testing.T) {
		b := detallistaSintetico
		m := leerMutado(t, detBase, `<spei:Complemento_SPEI>`, b+b+`<spei:Complemento_SPEI>`)
		if len(advDet(m, AdvNodoDuplicado, "/Comprobante/Complemento/detallista")) != 1 {
			t.Errorf("advertencias = %v", m.Advertencias)
		}
	})

	t.Run("espacios alrededor de decimal y fecha", func(t *testing.T) {
		m := leerDetallista(t,
			`<detallista:Amount>40.00</detallista:Amount>`, "<detallista:Amount>\n          40.00  </detallista:Amount>",
			`<detallista:ReferenceDate>2020-01-15</detallista:ReferenceDate>`, "<detallista:ReferenceDate> 2020-01-15\n</detallista:ReferenceDate>",
			`<detallista:invoicedQuantity unitOfMeasure="PCE">10.50</detallista:invoicedQuantity>`, "<detallista:invoicedQuantity unitOfMeasure=\"PCE\">\t10.50 </detallista:invoicedQuantity>")
		sinMasQue(t, m, AdvVersionAntigua)
		l := m.Detallista.LineItem[0]
		monto(t, "grossPrice con espacios", l.GrossPrice.Amount, "40.00")
		monto(t, "invoicedQuantity con espacios", l.InvoicedQuantity.Valor, "10.50")
		f := m.Detallista.OrderIdentification.ReferenceDate
		if f.Original != "2020-01-15" || !f.Valida() {
			t.Errorf("fecha con espacios = %+v", f)
		}
	})

	t.Run("texto de string sin recortar", func(t *testing.T) {
		m := leerDetallista(t, `<detallista:text>COMPRAS</detallista:text>`, `<detallista:text>  COMPRAS </detallista:text>`)
		igual(t, "contacto", m.Detallista.Buyer.ContactInformation.PersonOrDepartmentName.Text, "  COMPRAS ")
	})

	t.Run("specialInstruction mixed", func(t *testing.T) {
		// El texto suelto entre los hijos no se lee ni estorba.
		m := leerDetallista(t, `<detallista:text>PRIMERA</detallista:text>`, `suelto<detallista:text>PRIMERA</detallista:text>otro`)
		sinMasQue(t, m, AdvVersionAntigua)
		igual(t, "texto", strings.Join(m.Detallista.SpecialInstruction[0].Text, "|"), "PRIMERA|SEGUNDA")
	})

	t.Run("el texto de un hijo no se confunde con el de otro", func(t *testing.T) {
		m := leerDetallista(t)
		na := m.Detallista.InvoiceCreator.NameAndAddress
		if na.Name == na.City || na.City == na.PostalCode {
			t.Errorf("campos de nameAndAddress iguales: %+v", na)
		}
		igual(t, "city", na.City, "CIUDAD C")
		igual(t, "street", na.StreetAddressOne, "CALLE 3")
	})
}

// TestNodoTexto fija la regla del texto en leerArbol: se guarda lo que trae
// algo que no es espacio y la indentación no.
func TestNodoTexto(t *testing.T) {
	n, _, err := leerArbol([]byte("<a xmlns=\""+NSDetallista+"\">\n  <b> x </b>\n  <c>  </c>\n  <d>1<!-- c -->2</d>\n  <e/>\n</a>"), true)
	if err != nil {
		t.Fatal(err)
	}
	igual(t, "a", n.texto, "")
	igual(t, "b", n.hijos[0].texto, " x ")
	igual(t, "c", n.hijos[1].texto, "")
	igual(t, "d", n.hijos[2].texto, "12")
	igual(t, "e", n.hijos[3].texto, "")

	// Fuera del namespace de Detallista el texto no se guarda.
	n, _, err = leerArbol([]byte("<a><b> x </b></a>"), true)
	if err != nil {
		t.Fatal(err)
	}
	igual(t, "b ajeno", n.hijos[0].texto, "")
}

// bytesAsignados es lo que reserva f en total (acumulado, no el vivo): mide el
// crecimiento sin depender del reloj.
func bytesAsignados(f func()) uint64 {
	var a, b runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&a)
	f()
	runtime.ReadMemStats(&b)
	return b.TotalAlloc - a.TotalAlloc
}

// TestLeerArbolTextoPartidoEsLineal protege la lectura de un texto partido en
// muchos trozos por comentarios: acumularlo con += recopia todo lo anterior en
// cada trozo (cuadrático). Compara el coste de N y de 4N trozos: lineal crece
// ~4 veces, cuadrático ~16. Vale para el texto que se guarda (Detallista) y
// para el que se descarta (cualquier otro elemento).
func TestLeerArbolTextoPartidoEsLineal(t *testing.T) {
	for _, c := range []struct{ nombre, ns string }{{"detallista", NSDetallista}, {"otro", "urn:otro"}} {
		t.Run(c.nombre, func(t *testing.T) {
			const n = 20000
			var chico, grande uint64
			for range 3 { // el mínimo de varias corridas quita el ruido del recolector
				a := bytesAsignados(func() { _, _, _ = leerArbol(docTrozos(c.ns, n), true) })
				b := bytesAsignados(func() { _, _, _ = leerArbol(docTrozos(c.ns, 4*n), true) })
				if chico == 0 || a < chico {
					chico = a
				}
				if grande == 0 || b < grande {
					grande = b
				}
			}
			if razon := float64(grande) / float64(chico); razon > 6 {
				t.Errorf("4x trozos asignó %.1fx más (%d vs %d bytes): crecimiento no lineal", razon, grande, chico)
			}
		})
	}
	// El texto partido se junta entero.
	r, _, err := leerArbol(docTrozos(NSDetallista, 1000), true)
	if err != nil {
		t.Fatal(err)
	}
	igual(t, "texto", r.hijos[0].texto, strings.Repeat("a", 1000))
}

// docTrozos es un elemento cuyo texto "aaa…" llega partido en n trozos.
func docTrozos(ns string, n int) []byte {
	return []byte(`<r xmlns:d="` + ns + `"><d:gln>` + strings.Repeat("a<!---->", n) + `</d:gln></r>`)
}

// TestDetallistaNMTOKENColapsado: los valores de tipo xs:NMTOKEN o enumeración
// llevan whiteSpace collapse en el XSD ("\n INVOICE\n" es válido y vale
// INVOICE), así que salen normalizados; los string libres, sin recortar.
func TestDetallistaNMTOKENColapsado(t *testing.T) {
	ws := func(tag, v string) []string {
		return []string{"<detallista:" + tag + ">" + v + "</detallista:" + tag + ">", "<detallista:" + tag + ">\n\t  " + v + " \r\n  </detallista:" + tag + ">"}
	}
	var cambios []string
	for _, c := range [][2]string{
		{"entityType", "CREDIT_NOTE"}, {"currencyFunction", "BILLING_CURRENCY"}, {"currencyFunction", "PRICE_CURRENCY"},
		{"currencyFunction", "PAYMENT_CURRENCY"}, {"methodOfPayment", "PREPAID_BY_SELLER"}, {"taxCategory", "TRANSFERIDO"},
		{"taxTypeDescription", "VAT"}, {"specialServicesType", "AJ"}, {"specialServicesType", "FC"},
		{"specialServicesType", "TD"}, {"specialServicesType", "AJ"}, // global, global 2, línea, TotalAllowanceCharge
	} {
		cambios = append(cambios, ws(c[0], c[1])...)
	}
	// Un string libre con espacios alrededor se queda como viene.
	cambios = append(cambios, "<detallista:longText>", "<detallista:longText>  ")
	c := leerDetallista(t, cambios...)
	x := c.Detallista
	igual(t, "entityType", x.RequestForPaymentIdentification.EntityType, "CREDIT_NOTE")
	igual(t, "currencyFunction", strings.Join(x.Currency[0].CurrencyFunction, "|"), "BILLING_CURRENCY|PRICE_CURRENCY")
	igual(t, "currencyFunction 2", strings.Join(x.Currency[1].CurrencyFunction, "|"), "PAYMENT_CURRENCY")
	igual(t, "allowanceCharge/specialServicesType", x.AllowanceCharge[0].SpecialServicesType, "AJ")
	igual(t, "allowanceCharge 2", x.AllowanceCharge[1].SpecialServicesType, "FC")
	igual(t, "TotalAllowanceCharge", x.TotalAllowanceCharge[0].SpecialServicesType, "AJ")
	li := x.LineItem[0]
	igual(t, "línea specialServicesType", li.AllowanceCharge[0].SpecialServicesType, "TD")
	igual(t, "methodOfPayment", li.PalletInformation.Transport.MethodOfPayment, "PREPAID_BY_SELLER")
	igual(t, "taxCategory", li.TradeItemTaxInformation[0].TaxCategory, "TRANSFERIDO")
	igual(t, "taxTypeDescription", li.TradeItemTaxInformation[0].TaxTypeDescription, "VAT")
	if got := li.TradeItemDescriptionInformation.LongText; !strings.HasPrefix(got, "  ") {
		t.Errorf("longText se recortó: %q", got)
	}
}
