package main

import (
	"log"
	"net/http"
	"pt-api/services"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors" // Import CORS package
)

func main() {
	r := gin.New()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Specify allowed origin(s)
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Apply CORS middleware
	r.Use(cors.Default())

	r.Use(Logger())

	r.GET("/healthcheck", Healthcheck)
	r.POST("/sale-report", services.GenSaleReport)
	r.POST("/purchase-report", services.GenPurchaseReport)
	r.POST("/sale-report2", services.GenSaleReport2)

	r.Run(":8080")

}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Log latency
		latency := time.Since(start)
		log.Printf("[%s] %s %s %d %s", c.Request.Method, c.Request.URL.Path, c.ClientIP(), c.Writer.Status(), latency)
	}
}

func Healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
