package main

import (
	"net/http"

	"pt-api/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.GET("/healthcheck", Healthcheck)
	r.POST("/sale-report", services.GenSaleReport)
	r.POST("/purchase-report", services.GenPurchaseReport)

	r.Run()

}

func Healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
