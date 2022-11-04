package controllers

import (
	"encoding/xml"
	"finance-utility/configs"
	"finance-utility/models"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func SetInvoice(c *gin.Context) {
	var invoice models.Invoice

	if err := c.BindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    configs.ERROR_CODE_DECLINED,
			"message": "Invoice structure is not correct!",
			"error":   err,
		})
		return
	}

	/**
	* Generate XML content
	 */
	inv := models.XmlInvoice{}
	inv.CustomizationID = SetCustomizationID(invoice.Number, invoice.ContractNumber)
	inv.ID = invoice.Number
	inv.IssueDate = invoice.IssueDate
	inv.DueDate = invoice.DueDate
	inv.InvoiceTypeCode = configs.XML_INVOICE_TypeCode
	inv.Note = invoice.Comment
	inv.TaxPointDate = invoice.DeliveryDate
	inv.DocumentCurrencyCode = configs.XML_DOCUMENT_CurrencyCode
	inv.ContractDocumentReference.ID = invoice.ContractNumber

	inv.AccountingSupplierParty.Party.PostalAddress.StreetName = configs.XML_SUPPLIER_PARTY_PartyLegalEntity_RegistrationName
	inv.AccountingSupplierParty.Party.PostalAddress.CityName = configs.XML_SUPPLIER_PARTY_PostalAddress_CityName
	inv.AccountingSupplierParty.Party.PostalAddress.PostalZone = configs.XML_SUPPLIER_PARTY_PostalAddress_PostalZone
	inv.AccountingSupplierParty.Party.PostalAddress.CountrySubentity = configs.XML_SUPPLIER_PARTY_PostalAddress_CountrySubentity
	inv.AccountingSupplierParty.Party.PostalAddress.Country.IdentificationCode = configs.XML_SUPPLIER_PARTY_PostalAddress_Country_IdentificationCode

	inv.AccountingSupplierParty.Party.PartyTaxScheme.CompanyID = configs.XML_SUPPLIER_PARTY_PartyTaxScheme_CompanyID
	inv.AccountingSupplierParty.Party.PartyTaxScheme.TaxScheme.ID = configs.XML_SUPPLIER_PARTY_PartyTaxScheme_TaxScheme_ID

	inv.AccountingSupplierParty.Party.PartyLegalEntity.RegistrationName = configs.XML_SUPPLIER_PARTY_PartyLegalEntity_RegistrationName
	inv.AccountingSupplierParty.Party.PartyLegalEntity.CompanyID = configs.XML_SUPPLIER_PARTY_PartyLegalEntity_CompanyID

	inv.AccountingSupplierParty.Party.Contact.Name = configs.XML_SUPPLIER_PARTY_Contact_Name
	inv.AccountingSupplierParty.Party.Contact.Telephone = configs.XML_SUPPLIER_PARTY_Contact_Telephone
	inv.AccountingSupplierParty.Party.Contact.ElectronicMail = configs.XML_SUPPLIER_PARTY_Contact_ElectronicMail

	inv.AccountingCustomerParty.Party.PostalAddress.StreetName = invoice.Buyer.Address.Street
	inv.AccountingCustomerParty.Party.PostalAddress.CityName = invoice.Buyer.Address.City
	inv.AccountingCustomerParty.Party.PostalAddress.PostalZone = invoice.Buyer.Address.Zipcode
	inv.AccountingCustomerParty.Party.PostalAddress.CountrySubentity = invoice.Buyer.Address.County
	inv.AccountingCustomerParty.Party.PostalAddress.Country.IdentificationCode = invoice.Buyer.Address.Country

	inv.AccountingCustomerParty.Party.PartyTaxScheme.CompanyID = invoice.Buyer.Cui
	inv.AccountingCustomerParty.Party.PartyTaxScheme.TaxScheme.ID = configs.XML_CUSTOMER_PARTY_PartyTaxScheme_TaxScheme_ID

	inv.AccountingCustomerParty.Party.PartyLegalEntity.RegistrationName = invoice.Buyer.Name
	inv.AccountingCustomerParty.Party.PartyLegalEntity.CompanyID = invoice.Buyer.Cif

	inv.AccountingCustomerParty.Party.Contact.Name = invoice.Buyer.Contact.Name
	inv.AccountingCustomerParty.Party.Contact.Telephone = invoice.Buyer.Contact.Phone
	inv.AccountingCustomerParty.Party.Contact.ElectronicMail = invoice.Buyer.Contact.Email

	inv.Delivery.ActualDeliveryDate = invoice.Delivery.Date
	inv.Delivery.DeliveryLocation.Address.StreetName = invoice.Delivery.Address.Street
	inv.Delivery.DeliveryLocation.Address.CityName = invoice.Delivery.Address.City
	inv.Delivery.DeliveryLocation.Address.PostalZone = invoice.Delivery.Address.Zipcode
	inv.Delivery.DeliveryLocation.Address.CountrySubentity = invoice.Delivery.Address.County
	inv.Delivery.DeliveryLocation.Address.Country.IdentificationCode = invoice.Delivery.Address.Country

	inv.PaymentMeans.PaymentMeansCode = configs.XML_PAYMENTMeans_Code
	inv.PaymentMeans.PayeeFinancialAccount.ID = invoice.Buyer.Iban

	inv.TaxTotal.TaxAmount = getTaxAmount()
	inv.TaxTotal.TaxSubtotal.TaxableAmount = getTotalTaxableAmount()
	inv.TaxTotal.TaxSubtotal.TaxAmount = inv.TaxTotal.TaxAmount
	inv.TaxTotal.TaxSubtotal.TaxCategory.ID = configs.XML_TAXTOTAL_TaxSubtotal_TaxCategory_ID
	inv.TaxTotal.TaxSubtotal.TaxCategory.Percent = configs.XML_TAXTOTAL_TaxSubtotal_TaxCategory_Percent
	inv.TaxTotal.TaxSubtotal.TaxCategory.TaxScheme.ID = configs.XML_TAXTOTAL_TaxSubtotal_TaxCategory_TaxScheme_ID

	inv.LegalMonetaryTotal.LineExtensionAmount = invoice.Value
	inv.LegalMonetaryTotal.TaxExclusiveAmount = invoice.Value
	inv.LegalMonetaryTotal.TaxInclusiveAmount = invoice.Total
	inv.LegalMonetaryTotal.PayableAmount = invoice.Value

	inv.InvoiceLine.ID = "1"
	inv.InvoiceLine.InvoicedQuantity = "1"
	inv.InvoiceLine.LineExtensionAmount = "1"

	inv.InvoiceLine.Item.Description = invoice.Comment
	inv.InvoiceLine.Item.Name = configs.XML_INVOICELINE_Item_Name
	inv.InvoiceLine.Item.ClassifiedTaxCategory.ID = configs.XML_INVOICELINE_Item_ClassifiedTaxCategory_ID
	inv.InvoiceLine.Item.ClassifiedTaxCategory.Percent = configs.XML_INVOICELINE_Item_ClassifiedTaxCategory_Percent
	inv.InvoiceLine.Item.ClassifiedTaxCategory.TaxScheme.ID = configs.XML_INVOICELINE_Item_ClassifiedTaxCategory_TaxScheme_ID

	inv.InvoiceLine.Price.PriceAmount = invoice.Value

	filename := "tmp/" + inv.CustomizationID + ".xml"
	file, _ := os.Create(filename)

	xmlWriter := io.Writer(file)

	enc := xml.NewEncoder(xmlWriter)
	enc.Indent("  ", "    ")
	if err := enc.Encode(inv); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "File is not generated!",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "File is generated!",
		"error":   "",
	})

}

func GetInvoice(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello this is will show Invoice by ID: " + id + "  from eFactura API",
	})
}

func InvoiceSearch(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello this is SEARCH a invoice from eFactura API",
	})
}

func SetCustomizationID(invNumber string, invContractNumber string) string {
	currentTime := time.Now()
	currentDate := currentTime.Format("20060102")
	return "eFactura_" + currentDate + "_" + invNumber + "_" + invContractNumber
}

func getTaxAmount() string {
	return "0"
}

func getTotalTaxableAmount() string {
	return "0"
}
