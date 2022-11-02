package controllers

import (
	"finance-utility/configs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IsAuthorized(c *gin.Context) {
	token := c.Request.Header.Get("apikey")
	if len(token) > 0 {
		// just send an api key
		// not authenticated yet!!
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    configs.ERROR_CODE_ERROR,
			"message": "Missing Apikey!",
		})
	}
}
