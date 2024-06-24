package services

import (
	"net/http"

	"math/rand"

	"github.com/gin-gonic/gin"
)

type MonthlySaleData struct {
	Date        string `json:"date"`
	TotalAmount int    `json:"total_amount"`
	SlipCount   []int  `json:"slip_count"`
}

func GenSaleReport(c *gin.Context) {
	saleData := MonthlySaleData{}
	c.ShouldBindJSON(&saleData)

	var result []int

	date_count := len(saleData.SlipCount)

	slipTotal := 0
	for _, num := range saleData.SlipCount {
		slipTotal += num
	}
	baseTotal := int(float32(saleData.TotalAmount) * 0.6)
	base := baseTotal / int(slipTotal)
	remain := saleData.TotalAmount - int(baseTotal)
	println(baseTotal, base, date_count)

	for i := 0; i < slipTotal; i++ {
		result = append(result, base)
	}

	tempCount := 0
	for remain > 0 {
		randomNumber := rand.Intn(81) + 20
		randomNumber = randomNumber * 5
		var randIndex int
		if tempCount != -1 {
			randIndex = tempCount
		} else {
			randIndex = rand.Intn(slipTotal)
		}
		if remain-randomNumber > 0 {
			result[randIndex] = result[randIndex] + randomNumber
		} else {
			randomNumber = remain
			result[randIndex] = result[randIndex] + randomNumber
		}
		remain = remain - randomNumber

		if tempCount > slipTotal-2 || tempCount == -1 {
			tempCount = -1
		} else {
			tempCount++
		}
	}

	println(result)

	avg := int(saleData.TotalAmount) / slipTotal * 5

	println(slipTotal, avg)
	c.JSON(http.StatusOK, gin.H{
		"status":       "OK",
		"date":         saleData.Date,
		"total_amount": saleData.TotalAmount,
	})
}
