package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GenPurchaseReport(c *gin.Context) {
	print("WIP")
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
