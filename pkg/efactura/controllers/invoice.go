package controllers

import (
	"finance-utility/configs"
	"finance-utility/models"
	"net/http"

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

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello this is add new SetInvoice from eFactura API",
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
