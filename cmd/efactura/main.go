package main

import (
	"finance-utility/pkg/efactura/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger()) // Logger middleware

	r.GET("/", controllers.Home)
	r.NoRoute(controllers.PageNotFound)

	/*
	*  Api Specification v0.1
	 */
	r.GET("/spec", func(c *gin.Context) {
		c.File("api/eFactura_v1.0.json")
	})

	// Create Sub Router for  customised API version
	authorized := r.Group("/api/v1")
	authorized.Use(controllers.IsAuthorized)

	authorized.POST("/invoice", controllers.SetInvoice)
	authorized.GET("/invoice/:id", controllers.GetInvoice)
	authorized.POST("/invoice/search", controllers.InvoiceSearch)

	r.HandleMethodNotAllowed = true
	r.Run(":8081")
}
