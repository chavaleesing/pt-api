package services

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

type TbtData struct {
	TaxSale   int    `json:"tax_sale"`   // ยอดขายที่ต้องเสียภาษี
	UntaxSale int    `json:"untax_sale"` // ยอดขายที่ได้รับยกเว้น
	Date      string `json:"date"`
}

func GenPurchaseReport(c *gin.Context) {
	tbtData := TbtData{}
	c.ShouldBindJSON(&tbtData)

	startDate, _ := time.Parse("2006-01-02", tbtData.Date)

	dayCount := daysIn(startDate.Month(), startDate.Year())
	println(dayCount)

	resultTax := RandInt(tbtData.TaxSale, dayCount)
	resultUntax := RandInt(tbtData.UntaxSale, dayCount)

	file := genExcelTbt(resultTax, resultUntax, tbtData)

	var buf bytes.Buffer
	if err := file.Write(&buf); err != nil {
		c.String(http.StatusInternalServerError, "failed to write Excel file")
		return
	}
	// Set headers for file download
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename=example.xlsx")
	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", buf.Bytes())

}

func RandInt(sum int, dayCount int) []int {
	base_max := int(float64(sum) * 0.993 / float64(dayCount))
	base_min := int(int(float32(sum)*0.949) / dayCount)

	remain := sum
	var result []int
	for i := 0; i < dayCount; i++ {
		randomNumber := rand.Intn(base_max-base_min+1) + base_min
		result = append(result, randomNumber)
		remain = remain - randomNumber
	}

	advance_base := int(remain/200) + 1
	advance_base_min := int(float32(advance_base)*0.7) + 1

	numbers := []int{0, 1, 2, dayCount - 3, dayCount - 2, dayCount - 1}

	for remain > 0 {
		randomAdvNumber := rand.Intn(advance_base-advance_base_min+1) + advance_base_min
		randIndex := numbers[rand.Intn(len(numbers))]
		if randomAdvNumber < remain-randomAdvNumber {
			result[randIndex] = result[randIndex] + randomAdvNumber
			remain = remain - randomAdvNumber
		} else {
			result[randIndex] = result[randIndex] + remain
			remain = 0

		}
	}
	return result
}

func daysIn(m time.Month, year int) int {
	return time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

func genExcelTbt(resultTax []int, resultUntax []int, tbtData TbtData) *excelize.File {
	file := excelize.NewFile()
	startDate, _ := time.Parse("2006-01-02", tbtData.Date)
	//style, _ := file.NewStyle(&excelize.Style{NumFmt: 3})

	temp_count := 1
	style2, _ := file.NewStyle(&excelize.Style{NumFmt: 4})
	file.SetCellStyle("Sheet1", "C1", "C65", style2)
	file.SetCellStyle("Sheet1", "D1", "D65", style2)

	file.SetColWidth("Sheet1", "A", "A", 20)
	file.SetColWidth("Sheet1", "B", "B", 25)
	file.SetColWidth("Sheet1", "C", "C", 25)
	file.SetColWidth("Sheet1", "D", "D", 25)

	for i := range resultTax {
		file.SetCellValue("Sheet1", "A"+fmt.Sprint(temp_count), startDate.Format("2006-01-02"))
		file.SetCellValue("Sheet1", "B"+fmt.Sprint(temp_count), "ใบกำกับภาษีแบบย่อ")
		file.SetCellValue("Sheet1", "C"+fmt.Sprint(temp_count), resultTax[i])
		file.SetCellValue("Sheet1", "D"+fmt.Sprint(temp_count), float64(resultTax[i])*0.07)

		file.SetCellValue("Sheet1", "B"+fmt.Sprint(temp_count+1), "ยอดยกเว้น")
		file.SetCellValue("Sheet1", "C"+fmt.Sprint(temp_count+1), resultUntax[i])
		file.SetCellValue("Sheet1", "D"+fmt.Sprint(temp_count+1), 0)

		startDate = startDate.AddDate(0, 0, 1)
		temp_count = temp_count + 2
	}

	return file

}
