package configs

/*
InvoiceTypeCode
	380 - Factura
	381 - Nota de creditare
	384 - Factura corectata
	389 - Autofactura
*/

/*
eFactura XML static configuration
*/
const XML_INVOICE_TypeCode string = "300" // Factura
const XML_DOCUMENT_CurrencyCode string = "RON"

const XML_SUPPLIER_PARTY_PostalAddress_StreetName string = "Dimitrie Pompeiu nr. 9-9A"
const XML_SUPPLIER_PARTY_PostalAddress_CityName string = "SECTOR2"
const XML_SUPPLIER_PARTY_PostalAddress_PostalZone string = "020335"
const XML_SUPPLIER_PARTY_PostalAddress_CountrySubentity string = "RO-B"
const XML_SUPPLIER_PARTY_PostalAddress_Country_IdentificationCode string = "RO"

const XML_SUPPLIER_PARTY_PartyTaxScheme_CompanyID string = "RO43131360"
const XML_SUPPLIER_PARTY_PartyTaxScheme_TaxScheme string = "VAT"

const XML_SUPPLIER_PARTY_PartyLegalEntity_RegistrationName string = "NETOPIA FINANCIAL SERVICES S.R.L."
const XML_SUPPLIER_PARTY_PartyLegalEntity_CompanyID string = "43131360"

const XML_SUPPLIER_PARTY_Contact_Name string = "NETOPIA FINANCIAL SERVICES"
const XML_SUPPLIER_PARTY_Contact_Telephone string = "0700000000"
const XML_SUPPLIER_PARTY_Contact_ElectronicMail string = "finance@netopia.ro"
