package convert

import (
	firstDatabase "Golang_test/src/first_collation_database"
	originalDatabase "Golang_test/src/original_database"
	"encoding/json"
	"go/types"
	"strconv"
	"strings"
)

type StockJsonData struct {
	Data   [][]string
	Fields []string
}

type data struct {
	types.Array
}

//type data struct {
//	string
//}
//type dataArray []interface {}

const (
	StockDate = 0
	//Volume = 1
	Open   = 3
	High   = 4
	Low    = 5
	Close  = 6
	Change = 7
	Volume = 8
)

func StockDayOriginalToFirst(stockNo string, stockDateYearMonth string) {
	firstCollationStockDays := []firstDatabase.FirstCollationStockDay{}
	originalStockDays := originalDatabase.SelectStockDayForLastDate(stockNo, stockDateYearMonth)
	for _, originalStockDay := range originalStockDays {
		stock_no := originalStockDay.Stock_no
		Stock_json_data := originalStockDay.Stock_json_data
		aaa, _ := json.Marshal(Stock_json_data)
		var stockJsonData StockJsonData
		json.Unmarshal(aaa, &stockJsonData)
		for _, e := range stockJsonData.Data {
			open, _ := strconv.ParseFloat(e[Open], 64)
			high, _ := strconv.ParseFloat(e[High], 64)
			low, _ := strconv.ParseFloat(e[Low], 64)
			close, _ := strconv.ParseFloat(e[Close], 64)
			Change, _ := strconv.ParseFloat(e[Change], 64)
			volume, _ := strconv.ParseFloat(strings.Replace(e[Volume], ",", "", -1), 64)
			firstCollationStockDay := firstDatabase.FirstCollationStockDay{
				Stock_no:   stock_no,
				Stock_date: strings.Replace(e[StockDate], "/", "", -1),
				Open:       open,
				High:       high,
				Low:        low,
				Close:      close,
				Change:     Change,
				Volume:     volume,
			}
			firstCollationStockDays = append(firstCollationStockDays, firstCollationStockDay)
		}
		firstDatabase.InsertFirstCollationStockDay(firstCollationStockDays)

	}

}
