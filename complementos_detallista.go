package cfdi

import (
	"strings"

	"github.com/OrtaMarco/cfdi-go/decimal"
)

// Detallista (namespace .../detallista, XSD sin versión): factura del sector
// retail (estándar AMC 8.1). A diferencia de los demás complementos, lleva
// casi todos sus datos en el TEXTO de los elementos y no en atributos.
//
// El modelo sigue el XSD: cada elemento complejo es un objeto (puntero si
// aparece una vez, slice si se repite) y cada elemento de contenido simple es
// un campo del padre, con el nombre del SAT en snake_case. Los elementos de
// contenido simple que además llevan atributos (referenceIdentification,
// invoicedQuantity, lotNumber…) son objetos con sus atributos y el texto en
// "valor". Los montos son decimales exactos con su texto original.

// Detallista es la raíz del complemento. No tiene atributo de versión: lo
// dicta el namespace.
type Detallista struct {
	Type                            string                           `json:"type,omitempty"`
	ContentVersion                  string                           `json:"content_version,omitempty"`
	DocumentStructureVersion        string                           `json:"document_structure_version,omitempty"`
	DocumentStatus                  string                           `json:"document_status,omitempty"`
	RequestForPaymentIdentification *DetallistaRequestForPayment     `json:"request_for_payment_identification,omitempty"`
	SpecialInstruction              []DetallistaSpecialInstruction   `json:"special_instruction,omitempty"`
	OrderIdentification             *DetallistaOrderIdentification   `json:"order_identification,omitempty"`
	AdditionalInformation           *DetallistaAdditionalInformation `json:"additional_information,omitempty"`
	DeliveryNote                    *DetallistaDeliveryNote          `json:"delivery_note,omitempty"`
	Buyer                           *DetallistaBuyer                 `json:"buyer,omitempty"`
	Seller                          *DetallistaSeller                `json:"seller,omitempty"`
	ShipTo                          *DetallistaShipTo                `json:"ship_to,omitempty"`
	InvoiceCreator                  *DetallistaInvoiceCreator        `json:"invoice_creator,omitempty"`
	Customs                         []DetallistaCustoms              `json:"customs,omitempty"`
	Currency                        []DetallistaCurrency             `json:"currency,omitempty"`
	PaymentTerms                    *DetallistaPaymentTerms          `json:"payment_terms,omitempty"`
	ShipmentDetail                  *DetallistaShipmentDetail        `json:"shipment_detail,omitempty"`
	AllowanceCharge                 []DetallistaAllowanceCharge      `json:"allowance_charge,omitempty"`
	LineItem                        []DetallistaLineItem             `json:"line_item,omitempty"`
	TotalAmount                     *DetallistaMonto                 `json:"total_amount,omitempty"`
	TotalAllowanceCharge            []DetallistaTotalAllowanceCharge `json:"total_allowance_charge,omitempty"`
}

// DetallistaReferencia es un elemento de texto con atributo type
// (referenceIdentification, alternatePartyIdentification,
// alternateTradeItemIdentification, serialShippingContainerCode,
// description): Valor es el texto del elemento.
type DetallistaReferencia struct {
	Type  string `json:"type,omitempty"`
	Valor string `json:"valor,omitempty"`
}

// DetallistaMonto es un elemento que solo contiene un Amount (grossPrice,
// netPrice, grossAmount, netAmount, totalAmount).
type DetallistaMonto struct {
	Amount *decimal.Decimal `json:"amount,omitempty"`
}

type DetallistaRequestForPayment struct {
	EntityType string `json:"entity_type,omitempty"`
}

type DetallistaSpecialInstruction struct {
	Code string   `json:"code,omitempty"`
	Text []string `json:"text,omitempty"`
}

type DetallistaOrderIdentification struct {
	ReferenceIdentification []DetallistaReferencia `json:"reference_identification,omitempty"`
	ReferenceDate           Fecha                  `json:"reference_date,omitzero"`
}

type DetallistaAdditionalInformation struct {
	ReferenceIdentification []DetallistaReferencia `json:"reference_identification,omitempty"`
}

type DetallistaDeliveryNote struct {
	ReferenceIdentification []string `json:"reference_identification,omitempty"`
	ReferenceDate           Fecha    `json:"reference_date,omitzero"`
}

type DetallistaBuyer struct {
	GLN                string                        `json:"gln,omitempty"`
	ContactInformation *DetallistaContactInformation `json:"contact_information,omitempty"`
}

type DetallistaContactInformation struct {
	PersonOrDepartmentName *DetallistaPersonOrDepartmentName `json:"person_or_department_name,omitempty"`
}

type DetallistaPersonOrDepartmentName struct {
	Text string `json:"text,omitempty"`
}

type DetallistaSeller struct {
	GLN                          string                `json:"gln,omitempty"`
	AlternatePartyIdentification *DetallistaReferencia `json:"alternate_party_identification,omitempty"`
}

type DetallistaShipTo struct {
	GLN            string                         `json:"gln,omitempty"`
	NameAndAddress *DetallistaNameAndAddressLista `json:"name_and_address,omitempty"`
}

// DetallistaNameAndAddressLista es el nameAndAddress de shipTo: en el XSD es
// una xs:sequence repetible (name, streetAddressOne, city, postalCode, todos
// opcionales), así que cada campo es una lista en orden de documento.
type DetallistaNameAndAddressLista struct {
	Name             []string `json:"name,omitempty"`
	StreetAddressOne []string `json:"street_address_one,omitempty"`
	City             []string `json:"city,omitempty"`
	PostalCode       []string `json:"postal_code,omitempty"`
}

type DetallistaInvoiceCreator struct {
	GLN                          string                    `json:"gln,omitempty"`
	AlternatePartyIdentification *DetallistaReferencia     `json:"alternate_party_identification,omitempty"`
	NameAndAddress               *DetallistaNameAndAddress `json:"name_and_address,omitempty"`
}

type DetallistaNameAndAddress struct {
	Name             string `json:"name,omitempty"`
	StreetAddressOne string `json:"street_address_one,omitempty"`
	City             string `json:"city,omitempty"`
	PostalCode       string `json:"postal_code,omitempty"`
}

// DetallistaCustoms es el Customs global (solo trae el GLN de la aduana).
type DetallistaCustoms struct {
	GLN string `json:"gln,omitempty"`
}

type DetallistaCurrency struct {
	CurrencyISOCode  string           `json:"currency_iso_code,omitempty"`
	CurrencyFunction []string         `json:"currency_function,omitempty"`
	RateOfChange     *decimal.Decimal `json:"rate_of_change,omitempty"`
}

type DetallistaPaymentTerms struct {
	PaymentTermsEvent        string                     `json:"payment_terms_event,omitempty"`
	PaymentTermsRelationTime string                     `json:"payment_terms_relation_time,omitempty"`
	NetPayment               *DetallistaNetPayment      `json:"net_payment,omitempty"`
	DiscountPayment          *DetallistaDiscountPayment `json:"discount_payment,omitempty"`
}

type DetallistaNetPayment struct {
	NetPaymentTermsType string                       `json:"net_payment_terms_type,omitempty"`
	PaymentTimePeriod   *DetallistaPaymentTimePeriod `json:"payment_time_period,omitempty"`
}

type DetallistaPaymentTimePeriod struct {
	TimePeriodDue *DetallistaTimePeriodDue `json:"time_period_due,omitempty"`
}

type DetallistaTimePeriodDue struct {
	TimePeriod string `json:"time_period,omitempty"`
	Value      string `json:"value,omitempty"`
}

type DetallistaDiscountPayment struct {
	DiscountType string `json:"discount_type,omitempty"`
	Percentage   string `json:"percentage,omitempty"`
}

// DetallistaShipmentDetail es el shipmentDetail, un complexType vacío: si
// viene, en JSON sale {}.
type DetallistaShipmentDetail struct{}

// DetallistaAllowanceCharge es el allowanceCharge global de la factura.
type DetallistaAllowanceCharge struct {
	AllowanceChargeType        string                                `json:"allowance_charge_type,omitempty"`
	SettlementType             string                                `json:"settlement_type,omitempty"`
	SequenceNumber             string                                `json:"sequence_number,omitempty"`
	SpecialServicesType        string                                `json:"special_services_type,omitempty"`
	MonetaryAmountOrPercentage *DetallistaMonetaryAmountOrPercentage `json:"monetary_amount_or_percentage,omitempty"`
}

type DetallistaMonetaryAmountOrPercentage struct {
	Rate *DetallistaRate `json:"rate,omitempty"`
}

type DetallistaRate struct {
	Base       string           `json:"base,omitempty"`
	Percentage *decimal.Decimal `json:"percentage,omitempty"`
}

type DetallistaTotalAllowanceCharge struct {
	AllowanceOrChargeType string           `json:"allowance_or_charge_type,omitempty"`
	SpecialServicesType   string           `json:"special_services_type,omitempty"`
	Amount                *decimal.Decimal `json:"amount,omitempty"`
}

type DetallistaLineItem struct {
	Type                             string                               `json:"type,omitempty"`
	Number                           string                               `json:"number,omitempty"`
	TradeItemIdentification          *DetallistaTradeItemIdentification   `json:"trade_item_identification,omitempty"`
	AlternateTradeItemIdentification []DetallistaReferencia               `json:"alternate_trade_item_identification,omitempty"`
	TradeItemDescriptionInformation  *DetallistaTradeItemDescription      `json:"trade_item_description_information,omitempty"`
	InvoicedQuantity                 *DetallistaInvoicedQuantity          `json:"invoiced_quantity,omitempty"`
	AditionalQuantity                []DetallistaAditionalQuantity        `json:"aditional_quantity,omitempty"` // sic, así se llama en el XSD
	GrossPrice                       *DetallistaMonto                     `json:"gross_price,omitempty"`
	NetPrice                         *DetallistaMonto                     `json:"net_price,omitempty"`
	AdditionalInformation            *DetallistaLineAdditionalInformation `json:"additional_information,omitempty"`
	Customs                          []DetallistaLineCustoms              `json:"customs,omitempty"`
	LogisticUnits                    *DetallistaLogisticUnits             `json:"logistic_units,omitempty"`
	PalletInformation                *DetallistaPalletInformation         `json:"pallet_information,omitempty"`
	ExtendedAttributes               *DetallistaExtendedAttributes        `json:"extended_attributes,omitempty"`
	AllowanceCharge                  []DetallistaLineAllowanceCharge      `json:"allowance_charge,omitempty"`
	TradeItemTaxInformation          []DetallistaTradeItemTaxInformation  `json:"trade_item_tax_information,omitempty"`
	TotalLineAmount                  *DetallistaTotalLineAmount           `json:"total_line_amount,omitempty"`
}

type DetallistaTradeItemIdentification struct {
	GTIN string `json:"gtin,omitempty"`
}

type DetallistaTradeItemDescription struct {
	Language string `json:"language,omitempty"`
	LongText string `json:"long_text,omitempty"`
}

type DetallistaInvoicedQuantity struct {
	UnitOfMeasure string           `json:"unit_of_measure,omitempty"`
	Valor         *decimal.Decimal `json:"valor,omitempty"`
}

type DetallistaAditionalQuantity struct {
	QuantityType string           `json:"quantity_type,omitempty"`
	Valor        *decimal.Decimal `json:"valor,omitempty"`
}

type DetallistaLineAdditionalInformation struct {
	ReferenceIdentification *DetallistaReferencia `json:"reference_identification,omitempty"`
}

// DetallistaLineCustoms es el Customs de un lineItem (pedimento a nivel
// detalle).
type DetallistaLineCustoms struct {
	GLN                          string                           `json:"gln,omitempty"`
	AlternatePartyIdentification *DetallistaReferencia            `json:"alternate_party_identification,omitempty"`
	ReferenceDate                Fecha                            `json:"reference_date,omitzero"`
	NameAndAddress               *DetallistaCustomsNameAndAddress `json:"name_and_address,omitempty"`
}

type DetallistaCustomsNameAndAddress struct {
	Name string `json:"name,omitempty"`
}

type DetallistaLogisticUnits struct {
	SerialShippingContainerCode *DetallistaReferencia `json:"serial_shipping_container_code,omitempty"`
}

type DetallistaPalletInformation struct {
	PalletQuantity string                `json:"pallet_quantity,omitempty"`
	Description    *DetallistaReferencia `json:"description,omitempty"`
	Transport      *DetallistaTransport  `json:"transport,omitempty"`
}

type DetallistaTransport struct {
	MethodOfPayment string `json:"method_of_payment,omitempty"`
}

type DetallistaExtendedAttributes struct {
	LotNumber []DetallistaLotNumber `json:"lot_number,omitempty"`
}

type DetallistaLotNumber struct {
	ProductionDate Fecha  `json:"production_date,omitzero"`
	Valor          string `json:"valor,omitempty"`
}

// DetallistaLineAllowanceCharge es el allowanceCharge de un lineItem: su
// settlementType es opcional y sus montos van como texto (porcentaje y
// monto por unidad), no como decimales del XSD.
type DetallistaLineAllowanceCharge struct {
	AllowanceChargeType        string                                    `json:"allowance_charge_type,omitempty"`
	SettlementType             string                                    `json:"settlement_type,omitempty"`
	SequenceNumber             string                                    `json:"sequence_number,omitempty"`
	SpecialServicesType        string                                    `json:"special_services_type,omitempty"`
	MonetaryAmountOrPercentage *DetallistaLineMonetaryAmountOrPercentage `json:"monetary_amount_or_percentage,omitempty"`
}

type DetallistaLineMonetaryAmountOrPercentage struct {
	PercentagePerUnit string                 `json:"percentage_per_unit,omitempty"`
	RatePerUnit       *DetallistaRatePerUnit `json:"rate_per_unit,omitempty"`
}

type DetallistaRatePerUnit struct {
	AmountPerUnit string `json:"amount_per_unit,omitempty"`
}

type DetallistaTradeItemTaxInformation struct {
	TaxTypeDescription string                        `json:"tax_type_description,omitempty"`
	ReferenceNumber    string                        `json:"reference_number,omitempty"`
	TradeItemTaxAmount *DetallistaTradeItemTaxAmount `json:"trade_item_tax_amount,omitempty"`
	TaxCategory        string                        `json:"tax_category,omitempty"`
}

type DetallistaTradeItemTaxAmount struct {
	TaxPercentage *decimal.Decimal `json:"tax_percentage,omitempty"`
	TaxAmount     *decimal.Decimal `json:"tax_amount,omitempty"`
}

type DetallistaTotalLineAmount struct {
	GrossAmount *DetallistaMonto `json:"gross_amount,omitempty"`
	NetAmount   *DetallistaMonto `json:"net_amount,omitempty"`
}

// --- lectura ---

// hijo devuelve el hijo único con ese nombre (NODO_DUPLICADO si viene más de
// uno) y, si es requerido y falta, advierte NODO_FALTANTE.
func (l *lector) hijo(n *nodo, ns, local, ruta string, requerido bool) *nodo {
	h := l.unico(n, ns, local, ruta)
	if h == nil && requerido {
		l.advertir(AdvNodoFaltante, ruta+"/"+local, "falta el nodo "+local)
	}
	return h
}

// textoHijo es el texto, tal cual, de un hijo de contenido simple.
func (l *lector) textoHijo(n *nodo, ns, local, ruta string, requerido bool) string {
	if h := l.hijo(n, ns, local, ruta, requerido); h != nil {
		return h.texto
	}
	return ""
}

// decHijo lee el texto de un hijo como xs:decimal. Ilegible → nil con
// DECIMAL_INVALIDO en la ruta del elemento y el texto en Valor.
func (l *lector) decHijo(n *nodo, ns, local, ruta string, requerido bool) *decimal.Decimal {
	h := l.hijo(n, ns, local, ruta, requerido)
	if h == nil {
		return nil
	}
	return l.decValor(h.texto, ruta+"/"+local)
}

// fechaHijo lee el texto de un hijo como xs:date. xs:date colapsa los
// espacios, así que se recortan antes de interpretar.
func (l *lector) fechaHijo(n *nodo, ns, local, ruta string, requerido bool) Fecha {
	h := l.hijo(n, ns, local, ruta, requerido)
	if h == nil {
		return Fecha{}
	}
	return l.fechaValor(strings.Trim(h.texto, " \t\r\n"), ruta+"/"+local, true)
}

// textosHijos es el texto de cada hijo repetible, en orden de documento.
func (l *lector) textosHijos(n *nodo, ns, local, ruta string, requerido bool) []string {
	var out []string
	for _, h := range l.hijos(n, ns, local, ruta) {
		out = append(out, h.texto)
	}
	if len(out) == 0 && requerido {
		l.advertir(AdvNodoFaltante, ruta+"/"+local, "falta el nodo "+local)
	}
	return out
}

// referencia lee un elemento de texto con atributo type obligatorio.
func (l *lector) referencia(h *nodo, ruta string) DetallistaReferencia {
	l.requerido(h, ruta, "type")
	return DetallistaReferencia{Type: l.str(h, "type"), Valor: h.texto}
}

func (l *lector) referenciaHijo(n *nodo, ns, local, ruta string, requerido bool) *DetallistaReferencia {
	h := l.hijo(n, ns, local, ruta, requerido)
	if h == nil {
		return nil
	}
	x := l.referencia(h, ruta+"/"+local)
	return &x
}

func (l *lector) referencias(n *nodo, ns, local, ruta string, requerido bool) []DetallistaReferencia {
	var out []DetallistaReferencia
	for i, h := range l.hijos(n, ns, local, ruta) {
		out = append(out, l.referencia(h, idx(ruta, local, i)))
	}
	if len(out) == 0 && requerido {
		l.advertir(AdvNodoFaltante, ruta+"/"+local, "falta el nodo "+local)
	}
	return out
}

// montoHijo lee un elemento que solo contiene un Amount obligatorio.
func (l *lector) montoHijo(n *nodo, ns, local, ruta string, requerido bool) *DetallistaMonto {
	h := l.hijo(n, ns, local, ruta, requerido)
	if h == nil {
		return nil
	}
	return &DetallistaMonto{Amount: l.decHijo(h, ns, "Amount", ruta+"/"+local, true)}
}

func (l *lector) detallista(n *nodo, r string) *Detallista {
	ns := n.ns
	d := &Detallista{
		Type:                     l.str(n, "type"),
		ContentVersion:           l.str(n, "contentVersion"),
		DocumentStructureVersion: l.str(n, "documentStructureVersion"),
		DocumentStatus:           l.str(n, "documentStatus"),
	}
	l.requerido(n, r, "documentStructureVersion", "documentStatus")

	if h := l.hijo(n, ns, "requestForPaymentIdentification", r, true); h != nil {
		d.RequestForPaymentIdentification = &DetallistaRequestForPayment{
			EntityType: l.textoHijo(h, ns, "entityType", r+"/requestForPaymentIdentification", true),
		}
	}

	for i, h := range l.hijos(n, ns, "specialInstruction", r) {
		ri := idx(r, "specialInstruction", i)
		l.requerido(h, ri, "code")
		// Es mixed: el texto suelto entre sus hijos no se lee.
		d.SpecialInstruction = append(d.SpecialInstruction, DetallistaSpecialInstruction{
			Code: l.str(h, "code"),
			Text: l.textosHijos(h, ns, "text", ri, true),
		})
	}

	if h := l.hijo(n, ns, "orderIdentification", r, true); h != nil {
		ro := r + "/orderIdentification"
		d.OrderIdentification = &DetallistaOrderIdentification{
			ReferenceIdentification: l.referencias(h, ns, "referenceIdentification", ro, true),
			ReferenceDate:           l.fechaHijo(h, ns, "ReferenceDate", ro, false),
		}
	}

	if h := l.hijo(n, ns, "AdditionalInformation", r, true); h != nil {
		d.AdditionalInformation = &DetallistaAdditionalInformation{
			ReferenceIdentification: l.referencias(h, ns, "referenceIdentification", r+"/AdditionalInformation", true),
		}
	}

	if h := l.hijo(n, ns, "DeliveryNote", r, false); h != nil {
		rd := r + "/DeliveryNote"
		d.DeliveryNote = &DetallistaDeliveryNote{
			ReferenceIdentification: l.textosHijos(h, ns, "referenceIdentification", rd, true),
			ReferenceDate:           l.fechaHijo(h, ns, "ReferenceDate", rd, false),
		}
	}

	if h := l.hijo(n, ns, "buyer", r, true); h != nil {
		rb := r + "/buyer"
		b := &DetallistaBuyer{GLN: l.textoHijo(h, ns, "gln", rb, true)}
		if c := l.hijo(h, ns, "contactInformation", rb, false); c != nil {
			rc := rb + "/contactInformation"
			b.ContactInformation = &DetallistaContactInformation{}
			if p := l.hijo(c, ns, "personOrDepartmentName", rc, true); p != nil {
				b.ContactInformation.PersonOrDepartmentName = &DetallistaPersonOrDepartmentName{
					Text: l.textoHijo(p, ns, "text", rc+"/personOrDepartmentName", true),
				}
			}
		}
		d.Buyer = b
	}

	if h := l.hijo(n, ns, "seller", r, false); h != nil {
		rs := r + "/seller"
		d.Seller = &DetallistaSeller{
			GLN:                          l.textoHijo(h, ns, "gln", rs, true),
			AlternatePartyIdentification: l.referenciaHijo(h, ns, "alternatePartyIdentification", rs, true),
		}
	}

	if h := l.hijo(n, ns, "shipTo", r, false); h != nil {
		rs := r + "/shipTo"
		s := &DetallistaShipTo{GLN: l.textoHijo(h, ns, "gln", rs, false)}
		if na := l.hijo(h, ns, "nameAndAddress", rs, false); na != nil {
			rna := rs + "/nameAndAddress"
			s.NameAndAddress = &DetallistaNameAndAddressLista{
				Name:             l.textosHijos(na, ns, "name", rna, false),
				StreetAddressOne: l.textosHijos(na, ns, "streetAddressOne", rna, false),
				City:             l.textosHijos(na, ns, "city", rna, false),
				PostalCode:       l.textosHijos(na, ns, "postalCode", rna, false),
			}
		}
		d.ShipTo = s
	}

	if h := l.hijo(n, ns, "InvoiceCreator", r, false); h != nil {
		rc := r + "/InvoiceCreator"
		ic := &DetallistaInvoiceCreator{
			GLN:                          l.textoHijo(h, ns, "gln", rc, false),
			AlternatePartyIdentification: l.referenciaHijo(h, ns, "alternatePartyIdentification", rc, false),
		}
		if na := l.hijo(h, ns, "nameAndAddress", rc, false); na != nil {
			rna := rc + "/nameAndAddress"
			ic.NameAndAddress = &DetallistaNameAndAddress{
				Name:             l.textoHijo(na, ns, "name", rna, false),
				StreetAddressOne: l.textoHijo(na, ns, "streetAddressOne", rna, false),
				City:             l.textoHijo(na, ns, "city", rna, false),
				PostalCode:       l.textoHijo(na, ns, "postalCode", rna, false),
			}
		}
		d.InvoiceCreator = ic
	}

	for i, h := range l.hijos(n, ns, "Customs", r) {
		d.Customs = append(d.Customs, DetallistaCustoms{GLN: l.textoHijo(h, ns, "gln", idx(r, "Customs", i), false)})
	}

	for i, h := range l.hijos(n, ns, "currency", r) {
		ri := idx(r, "currency", i)
		l.requerido(h, ri, "currencyISOCode")
		d.Currency = append(d.Currency, DetallistaCurrency{
			CurrencyISOCode:  l.str(h, "currencyISOCode"),
			CurrencyFunction: l.textosHijos(h, ns, "currencyFunction", ri, true),
			RateOfChange:     l.decHijo(h, ns, "rateOfChange", ri, false),
		})
	}

	if h := l.hijo(n, ns, "paymentTerms", r, false); h != nil {
		d.PaymentTerms = l.detallistaPaymentTerms(h, ns, r+"/paymentTerms")
	}

	if l.hijo(n, ns, "shipmentDetail", r, false) != nil {
		d.ShipmentDetail = &DetallistaShipmentDetail{}
	}

	for i, h := range l.hijos(n, ns, "allowanceCharge", r) {
		ri := idx(r, "allowanceCharge", i)
		l.requerido(h, ri, "allowanceChargeType", "settlementType")
		a := DetallistaAllowanceCharge{
			AllowanceChargeType: l.str(h, "allowanceChargeType"),
			SettlementType:      l.str(h, "settlementType"),
			SequenceNumber:      l.str(h, "sequenceNumber"),
			SpecialServicesType: l.textoHijo(h, ns, "specialServicesType", ri, false),
		}
		if m := l.hijo(h, ns, "monetaryAmountOrPercentage", ri, false); m != nil {
			rm := ri + "/monetaryAmountOrPercentage"
			a.MonetaryAmountOrPercentage = &DetallistaMonetaryAmountOrPercentage{}
			if rt := l.hijo(m, ns, "rate", rm, false); rt != nil {
				rr := rm + "/rate"
				l.requerido(rt, rr, "base")
				a.MonetaryAmountOrPercentage.Rate = &DetallistaRate{
					Base:       l.str(rt, "base"),
					Percentage: l.decHijo(rt, ns, "percentage", rr, true),
				}
			}
		}
		d.AllowanceCharge = append(d.AllowanceCharge, a)
	}

	for i, h := range l.hijos(n, ns, "lineItem", r) {
		d.LineItem = append(d.LineItem, l.detallistaLineItem(h, ns, idx(r, "lineItem", i)))
	}

	d.TotalAmount = l.montoHijo(n, ns, "totalAmount", r, false)

	for i, h := range l.hijos(n, ns, "TotalAllowanceCharge", r) {
		ri := idx(r, "TotalAllowanceCharge", i)
		l.requerido(h, ri, "allowanceOrChargeType")
		d.TotalAllowanceCharge = append(d.TotalAllowanceCharge, DetallistaTotalAllowanceCharge{
			AllowanceOrChargeType: l.str(h, "allowanceOrChargeType"),
			SpecialServicesType:   l.textoHijo(h, ns, "specialServicesType", ri, false),
			Amount:                l.decHijo(h, ns, "Amount", ri, false),
		})
	}
	return d
}

func (l *lector) detallistaPaymentTerms(h *nodo, ns, r string) *DetallistaPaymentTerms {
	p := &DetallistaPaymentTerms{
		PaymentTermsEvent:        l.str(h, "paymentTermsEvent"),
		PaymentTermsRelationTime: l.str(h, "PaymentTermsRelationTime"),
	}
	if np := l.hijo(h, ns, "netPayment", r, false); np != nil {
		rn := r + "/netPayment"
		l.requerido(np, rn, "netPaymentTermsType")
		p.NetPayment = &DetallistaNetPayment{NetPaymentTermsType: l.str(np, "netPaymentTermsType")}
		if pp := l.hijo(np, ns, "paymentTimePeriod", rn, false); pp != nil {
			rp := rn + "/paymentTimePeriod"
			p.NetPayment.PaymentTimePeriod = &DetallistaPaymentTimePeriod{}
			if td := l.hijo(pp, ns, "timePeriodDue", rp, true); td != nil {
				rt := rp + "/timePeriodDue"
				l.requerido(td, rt, "timePeriod")
				p.NetPayment.PaymentTimePeriod.TimePeriodDue = &DetallistaTimePeriodDue{
					TimePeriod: l.str(td, "timePeriod"),
					Value:      l.textoHijo(td, ns, "value", rt, true),
				}
			}
		}
	}
	if dp := l.hijo(h, ns, "discountPayment", r, false); dp != nil {
		rd := r + "/discountPayment"
		l.requerido(dp, rd, "discountType")
		p.DiscountPayment = &DetallistaDiscountPayment{
			DiscountType: l.str(dp, "discountType"),
			Percentage:   l.textoHijo(dp, ns, "percentage", rd, true),
		}
	}
	return p
}

func (l *lector) detallistaLineItem(h *nodo, ns, r string) DetallistaLineItem {
	li := DetallistaLineItem{Type: l.str(h, "type"), Number: l.str(h, "number")}

	if t := l.hijo(h, ns, "tradeItemIdentification", r, true); t != nil {
		li.TradeItemIdentification = &DetallistaTradeItemIdentification{
			GTIN: l.textoHijo(t, ns, "gtin", r+"/tradeItemIdentification", true),
		}
	}
	li.AlternateTradeItemIdentification = l.referencias(h, ns, "alternateTradeItemIdentification", r, false)
	if t := l.hijo(h, ns, "tradeItemDescriptionInformation", r, false); t != nil {
		li.TradeItemDescriptionInformation = &DetallistaTradeItemDescription{
			Language: l.str(t, "language"),
			LongText: l.textoHijo(t, ns, "longText", r+"/tradeItemDescriptionInformation", true),
		}
	}
	if q := l.hijo(h, ns, "invoicedQuantity", r, true); q != nil {
		rq := r + "/invoicedQuantity"
		l.requerido(q, rq, "unitOfMeasure")
		li.InvoicedQuantity = &DetallistaInvoicedQuantity{UnitOfMeasure: l.str(q, "unitOfMeasure"), Valor: l.decValor(q.texto, rq)}
	}
	for i, q := range l.hijos(h, ns, "aditionalQuantity", r) {
		rq := idx(r, "aditionalQuantity", i)
		l.requerido(q, rq, "QuantityType")
		li.AditionalQuantity = append(li.AditionalQuantity, DetallistaAditionalQuantity{QuantityType: l.str(q, "QuantityType"), Valor: l.decValor(q.texto, rq)})
	}
	li.GrossPrice = l.montoHijo(h, ns, "grossPrice", r, false)
	li.NetPrice = l.montoHijo(h, ns, "netPrice", r, false)
	if a := l.hijo(h, ns, "AdditionalInformation", r, false); a != nil {
		li.AdditionalInformation = &DetallistaLineAdditionalInformation{
			ReferenceIdentification: l.referenciaHijo(a, ns, "referenceIdentification", r+"/AdditionalInformation", false),
		}
	}
	for i, c := range l.hijos(h, ns, "Customs", r) {
		rc := idx(r, "Customs", i)
		cu := DetallistaLineCustoms{
			GLN:                          l.textoHijo(c, ns, "gln", rc, false),
			AlternatePartyIdentification: l.referenciaHijo(c, ns, "alternatePartyIdentification", rc, true),
			ReferenceDate:                l.fechaHijo(c, ns, "ReferenceDate", rc, true),
		}
		if na := l.hijo(c, ns, "nameAndAddress", rc, true); na != nil {
			cu.NameAndAddress = &DetallistaCustomsNameAndAddress{Name: l.textoHijo(na, ns, "name", rc+"/nameAndAddress", true)}
		}
		li.Customs = append(li.Customs, cu)
	}
	if u := l.hijo(h, ns, "LogisticUnits", r, false); u != nil {
		li.LogisticUnits = &DetallistaLogisticUnits{
			SerialShippingContainerCode: l.referenciaHijo(u, ns, "serialShippingContainerCode", r+"/LogisticUnits", true),
		}
	}
	if p := l.hijo(h, ns, "palletInformation", r, false); p != nil {
		rp := r + "/palletInformation"
		pi := &DetallistaPalletInformation{
			PalletQuantity: l.textoHijo(p, ns, "palletQuantity", rp, true),
			Description:    l.referenciaHijo(p, ns, "description", rp, true),
		}
		if t := l.hijo(p, ns, "transport", rp, true); t != nil {
			pi.Transport = &DetallistaTransport{MethodOfPayment: l.textoHijo(t, ns, "methodOfPayment", rp+"/transport", true)}
		}
		li.PalletInformation = pi
	}
	if e := l.hijo(h, ns, "extendedAttributes", r, false); e != nil {
		re := r + "/extendedAttributes"
		li.ExtendedAttributes = &DetallistaExtendedAttributes{}
		for i, ln := range l.hijos(e, ns, "lotNumber", re) {
			li.ExtendedAttributes.LotNumber = append(li.ExtendedAttributes.LotNumber, DetallistaLotNumber{
				ProductionDate: l.fecha(ln, idx(re, "lotNumber", i), "productionDate", true),
				Valor:          ln.texto,
			})
		}
		if len(li.ExtendedAttributes.LotNumber) == 0 {
			l.advertir(AdvNodoFaltante, re+"/lotNumber", "falta el nodo lotNumber")
		}
	}
	for i, a := range l.hijos(h, ns, "allowanceCharge", r) {
		ra := idx(r, "allowanceCharge", i)
		l.requerido(a, ra, "allowanceChargeType")
		ac := DetallistaLineAllowanceCharge{
			AllowanceChargeType: l.str(a, "allowanceChargeType"),
			SettlementType:      l.str(a, "settlementType"),
			SequenceNumber:      l.str(a, "sequenceNumber"),
			SpecialServicesType: l.textoHijo(a, ns, "specialServicesType", ra, false),
		}
		if m := l.hijo(a, ns, "monetaryAmountOrPercentage", ra, true); m != nil {
			rm := ra + "/monetaryAmountOrPercentage"
			ac.MonetaryAmountOrPercentage = &DetallistaLineMonetaryAmountOrPercentage{
				PercentagePerUnit: l.textoHijo(m, ns, "percentagePerUnit", rm, true),
			}
			if rt := l.hijo(m, ns, "ratePerUnit", rm, false); rt != nil {
				ac.MonetaryAmountOrPercentage.RatePerUnit = &DetallistaRatePerUnit{
					AmountPerUnit: l.textoHijo(rt, ns, "amountPerUnit", rm+"/ratePerUnit", true),
				}
			}
		}
		li.AllowanceCharge = append(li.AllowanceCharge, ac)
	}
	for i, t := range l.hijos(h, ns, "tradeItemTaxInformation", r) {
		rt := idx(r, "tradeItemTaxInformation", i)
		ti := DetallistaTradeItemTaxInformation{
			TaxTypeDescription: l.textoHijo(t, ns, "taxTypeDescription", rt, true),
			ReferenceNumber:    l.textoHijo(t, ns, "referenceNumber", rt, false),
			TaxCategory:        l.textoHijo(t, ns, "taxCategory", rt, false),
		}
		if m := l.hijo(t, ns, "tradeItemTaxAmount", rt, false); m != nil {
			rm := rt + "/tradeItemTaxAmount"
			ti.TradeItemTaxAmount = &DetallistaTradeItemTaxAmount{
				TaxPercentage: l.decHijo(m, ns, "taxPercentage", rm, true),
				TaxAmount:     l.decHijo(m, ns, "taxAmount", rm, true),
			}
		}
		li.TradeItemTaxInformation = append(li.TradeItemTaxInformation, ti)
	}
	if t := l.hijo(h, ns, "totalLineAmount", r, true); t != nil {
		rt := r + "/totalLineAmount"
		li.TotalLineAmount = &DetallistaTotalLineAmount{
			GrossAmount: l.montoHijo(t, ns, "grossAmount", rt, false),
			NetAmount:   l.montoHijo(t, ns, "netAmount", rt, true),
		}
	}
	return li
}
