package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MonthlySaleData struct {
	Date        string  `json:"date"`
	TotalAmount float64 `json:"total_amount"`
	SlipCount   []int   `json:"slip_count"`
}

func GenSaleReport(c *gin.Context) {
	saleData := MonthlySaleData{}
	c.ShouldBindJSON(&saleData)
	c.JSON(http.StatusOK, gin.H{
		"status":       "OK",
		"date":         saleData.Date,
		"total_amount": saleData.TotalAmount,
	})
}
