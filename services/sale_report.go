package services

import (
	"net/http"

	"math/rand"

	"github.com/gin-gonic/gin"

	"github.com/xuri/excelize/v2"

	"fmt"
	"time"
)

type MonthlySaleData struct {
	Date        string `json:"date"` // "2021-02-01"
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

	max := int(float32(base) * 2.5)

	tempCount := 0
	for remain > 0 {
		randomNumber := rand.Intn(81) + 25
		randomNumber = randomNumber * 5
		var randIndex int
		if tempCount != -1 {
			randIndex = tempCount
		} else {
			randIndex = rand.Intn(slipTotal)
		}

		if result[randIndex] > max {
			continue
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

	genExcel(result, saleData)

	println(slipTotal)
	c.JSON(http.StatusOK, gin.H{
		"status":       "OK",
		"date":         saleData.Date,
		"total_amount": saleData.TotalAmount,
	})
}

func genExcel(result []int, saleData MonthlySaleData) {
	file := excelize.NewFile()
	startDate, _ := time.Parse("2006-01-02", saleData.Date)
	style, _ := file.NewStyle(&excelize.Style{NumFmt: 3})
	headers := startDate

	tempIdResult := 0
	colNumber := 1

	for i := 0; i < len(saleData.SlipCount); i++ {
		colName := getColumnName(colNumber)
		file.SetColWidth("Sheet1", colName, colName, 15)
		file.SetCellValue("Sheet1", colName+"1", headers.Format("2006-01-02"))
		itemPerDay := saleData.SlipCount[i]
		rowNumber := 2
		for i := 0; i < itemPerDay; i++ {
			file.SetCellValue("Sheet1", colName+fmt.Sprint(rowNumber), result[tempIdResult])
			rowNumber++
			tempIdResult++
		}
		file.SetCellStyle("Sheet1", colName+"2", colName+fmt.Sprint(rowNumber), style)
		colNumber++
		headers = headers.AddDate(0, 0, 1)
	}
	file.SaveAs("sale_report.xlsx")
}

func getColumnName(col int) string {
	name := make([]byte, 0, 3) // max 16,384 columns (2022)
	const aLen = 'Z' - 'A' + 1 // alphabet length
	for ; col > 0; col /= aLen + 1 {
		name = append(name, byte('A'+(col-1)%aLen))
	}
	for i, j := 0, len(name)-1; i < j; i, j = i+1, j-1 {
		name[i], name[j] = name[j], name[i]
	}
	return string(name)
}
