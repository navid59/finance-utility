package models

import "encoding/xml"

/**
*	XML invoice Model
 */

type XmlInvoice struct {
	XMLName                   xml.Name                  `xml:"Invoice"`
	CustomizationID           string                    `xml:"cbc:CustomizationID"`
	ID                        string                    `xml:"cbc:ID"`
	IssueDate                 string                    `xml:"cbc:IssueDate"`
	DueDate                   string                    `xml:"cbc:DueDate"`
	InvoiceTypeCode           string                    `xml:"cbc:InvoiceTypeCode"`
	Note                      string                    `xml:"cbc:Note"`
	TaxPointDate              string                    `xml:"cbc:TaxPointDate"`
	DocumentCurrencyCode      string                    `xml:"cbc:DocumentCurrencyCode"`
	ContractDocumentReference ContractDocumentReference `xml:"cac:ContractDocumentReference"`
	AccountingSupplierParty   AccountingSupplierParty   `xml:"cac:AccountingSupplierParty"`
	AccountingCustomerParty   AccountingCustomerParty   `xml:"cac:AccountingCustomerParty"`
	Delivery                  XmlDelivery               `xml:"cac:Delivery"`
	PaymentMeans              PaymentMeans              `xml:"cac:PaymentMeans"`
	TaxTotal                  TaxTotal                  `xml:"cac:TaxTotal"`
	LegalMonetaryTotal        LegalMonetaryTotal        `xml:"cac:LegalMonetaryTotal"`
	InvoiceLine               InvoiceLine               `xml:"cac:InvoiceLine"`
}

type ContractDocumentReference struct {
	XMLName xml.Name `xml:"cac:ContractDocumentReference"`
	ID      string   `xml:"cbc:ID"`
}

type AccountingSupplierParty struct {
	XMLName xml.Name `xml:"cac:AccountingSupplierParty"`
	Party   Party    `xml:"cac:Party"`
}

type AccountingCustomerParty struct {
	XMLName xml.Name `xml:"cac:AccountingCustomerParty"`
	Party   Party    `xml:"cac:Party"`
}

type Party struct {
	XMLName          xml.Name         `xml:"cac:Party"`
	PostalAddress    PostalAddress    `xml:"cac:PostalAddress"`
	PartyTaxScheme   PartyTaxScheme   `xml:"cac:PartyTaxScheme"`
	PartyLegalEntity PartyLegalEntity `xml:"cac:PartyLegalEntity"`
	Contact          PartyContact     `xml:"cac:Contact"`
}

type PostalAddress struct {
	XMLName          xml.Name `xml:"cac:PostalAddress"`
	StreetName       string   `xml:"cbc:StreetName"`
	CityName         string   `xml:"cbc:CityName"`
	PostalZone       string   `xml:"cbc:PostalZone"`
	CountrySubentity string   `xml:"cbc:CountrySubentity"`
	Country          Country  `xml:"cac:Country"`
}

type Country struct {
	XMLName            xml.Name `xml:"cac:Country"`
	IdentificationCode string   `xml:"cbc:IdentificationCode"`
}

type PartyTaxScheme struct {
	XMLName   xml.Name  `xml:"cac:PartyTaxScheme"`
	CompanyID string    `xml:"cbc:CompanyID"`
	TaxScheme TaxScheme `xml:"cac:TaxScheme"`
}

type TaxScheme struct {
	XMLName xml.Name `xml:"cac:TaxScheme"`
	ID      string   `xml:"cbc:ID"`
}

type PartyLegalEntity struct {
	XMLName          xml.Name `xml:"cac:PartyLegalEntity"`
	RegistrationName string   `xml:"cbc:RegistrationName"`
	CompanyID        string   `xml:"cbc:CompanyID"`
}

type PartyContact struct {
	XMLName        xml.Name `xml:"cac:Contact"`
	Name           string   `xml:"cbc:Name"`
	Telephone      string   `xml:"cbc:Telephone"`
	ElectronicMail string   `xml:"cbc:ElectronicMail"`
}

type XmlDelivery struct {
	XMLName            xml.Name         `xml:"cac:Delivery"`
	ActualDeliveryDate string           `xml:"cbc:ActualDeliveryDate"`
	DeliveryLocation   DeliveryLocation `xml:"cac:DeliveryLocation"`
}

type DeliveryLocation struct {
	XMLName xml.Name        `xml:"cac:DeliveryLocation"`
	Address DeliveryAddress `xml:"cac:Address"`
}

type DeliveryAddress struct {
	XMLName          xml.Name `xml:"cac:Address"`
	StreetName       string   `xml:"cbc:StreetName"`
	CityName         string   `xml:"cbc:CityName"`
	PostalZone       string   `xml:"cbc:PostalZone"`
	CountrySubentity string   `xml:"cbc:CountrySubentity"`
	Country          Country  `xml:"cac:Country"`
}

type PaymentMeans struct {
	XMLName               xml.Name              `xml:"cac:PaymentMeans"`
	PaymentMeansCode      string                `xml:"cbc:PaymentMeansCode"`
	PayeeFinancialAccount PayeeFinancialAccount `xml:"cac:PayeeFinancialAccount"`
}

type PayeeFinancialAccount struct {
	XMLName xml.Name `xml:"cac:PayeeFinancialAccount"`
	ID      string   `xml:"cbc:ID"`
}

type TaxTotal struct {
	XMLName     xml.Name    `xml:"cac:TaxTotal"`
	TaxAmount   string      `xml:"cbc:TaxAmount"`
	TaxSubtotal TaxSubtotal `xml:"cac:TaxSubtotal"`
}

type TaxSubtotal struct {
	XMLName       xml.Name    `xml:"cac:TaxSubtotal"`
	TaxableAmount string      `xml:"cbc:TaxableAmount"`
	TaxAmount     string      `xml:"cbc:TaxAmount"`
	TaxCategory   TaxCategory `xml:"cac:TaxCategory"`
}

type TaxCategory struct {
	XMLName   xml.Name  `xml:"cac:TaxCategory"`
	ID        string    `xml:"cbc:ID"`
	Percent   string    `xml:"cbc:Percent"`
	TaxScheme TaxScheme `xml:"cac:TaxScheme"`
}

type LegalMonetaryTotal struct {
	XMLName             xml.Name `xml:"cac:LegalMonetaryTotal"`
	LineExtensionAmount string   `xml:"cbc:LineExtensionAmount"`
	TaxExclusiveAmount  string   `xml:"cbc:TaxExclusiveAmount"`
	TaxInclusiveAmount  string   `xml:"cbc:TaxInclusiveAmount"`
	PayableAmount       string   `xml:"cbc:PayableAmount"`
}

type InvoiceLine struct {
	XMLName          xml.Name `xml:"cac:InvoiceLine"`
	ID               string   `xml:"cbc:ID"`
	InvoicedQuantity string   `xml:"cbc:InvoicedQuantity"`
	Item             Item     `xml:"cac:Item"`
	Price            Price    `xml:"cac:Price"`
}

type Item struct {
	XMLName               xml.Name              `xml:"cac:Item"`
	Description           string                `xml:"cbc:Description"`
	Name                  string                `xml:"cbc:Name"`
	ClassifiedTaxCategory ClassifiedTaxCategory `xml:"cac:ClassifiedTaxCategory"`
}

type ClassifiedTaxCategory struct {
	XMLName   xml.Name  `xml:"cac:ClassifiedTaxCategory"`
	ID        string    `xml:"cbc:ID"`
	Percent   string    `xml:"cbc:Percent"`
	TaxScheme TaxScheme `xml:"cac:TaxScheme"`
}

type Price struct {
	XMLName     xml.Name `xml:"cac:Price"`
	PriceAmount string   `xml:"cbc:PriceAmount"`
}
