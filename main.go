package main

import (
	"net/http"

	"pt-api/services"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors" // Import CORS package
)

func main() {
	r := gin.New()

	// Apply CORS middleware
	r.Use(cors.Default())

	r.GET("/healthcheck", Healthcheck)
	r.POST("/sale-report", services.GenSaleReport)
	r.POST("/purchase-report", services.GenPurchaseReport)

	r.Run()

}

func Healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
