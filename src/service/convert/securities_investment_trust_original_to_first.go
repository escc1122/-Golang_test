package convert

import (
	firstDatabase "Golang_test/src/first_collation_database"
	originalDatabase "Golang_test/src/original_database"

	//"encoding/json"
	//"go/types"
	"fmt"
	"strconv"
	"strings"
)

// type StockJsonData struct {
// 	Data   [][]string
// 	Fields []string
// }

// type data struct {
// 	types.Array
// }

const (
	stockID           = 1
	stockName         = 2
	buyTradingVolume  = 3
	sellTradingVolume = 4
	difference        = 5
)

func SecuritiesInvestmentTrustOriginalToFirst(start_date string, end_date string) {
	firstCollationSecuritiesInvestmentTrusts := []firstDatabase.FirstCollationSecuritiesInvestmentTrust{}
	originalSecuritiesInvestmentTrusts := originalDatabase.SelectSecuritiesInvestmentTrust(start_date, end_date)
	fmt.Println(originalSecuritiesInvestmentTrusts)
	for _, originalSecuritiesInvestmentTrust := range originalSecuritiesInvestmentTrusts {
		fmt.Println("aaaaaa", originalSecuritiesInvestmentTrust)
		var datas = originalSecuritiesInvestmentTrust.Json_data.Data
		stock_date := originalSecuritiesInvestmentTrust.Date
		fmt.Println("stock_date", stock_date)
		for _, data := range datas {
			btv, _ := strconv.ParseInt(strings.Replace(data[buyTradingVolume], ",", "", -1), 10, 64)
			stv, _ := strconv.ParseInt(strings.Replace(data[sellTradingVolume], ",", "", -1), 10, 64)
			diff, _ := strconv.ParseInt(strings.Replace(data[difference], ",", "", -1), 10, 64)

			firstCollationSecuritiesInvestmentTrust := firstDatabase.FirstCollationSecuritiesInvestmentTrust{
				Stock_no:            strings.Trim(data[stockID], " "),
				Stock_date:          stock_date,
				Stock_name:          strings.Trim(data[stockName], " "),
				Buy_trading_volume:  btv,
				Sell_trading_volume: stv,
				Difference:          diff,
			}
			firstCollationSecuritiesInvestmentTrusts = append(firstCollationSecuritiesInvestmentTrusts, firstCollationSecuritiesInvestmentTrust)
		}
	}
	firstDatabase.InsertFirstCollationSecuritiesInvestmentTrust(firstCollationSecuritiesInvestmentTrusts)

}

// func StockDayOriginalToFirst(stockNo string, stockDateYearMonth string) {
// 	firstCollationStockDays := []firstDatabase.FirstCollationStockDay{}
// 	originalStockDays := originalDatabase.SelectStockDayForLastDate(stockNo, stockDateYearMonth)
// 	for _, originalStockDay := range originalStockDays {
// 		stock_no := originalStockDay.Stock_no
//

// 		aaa, _ := json.Marshal(Stock_json_data)
// 		var stockJsonData StockJsonData
// 		json.Unmarshal(aaa, &stockJsonData)
// 		for _, e := range stockJsonData.Data {
// 			open, _ := strconv.ParseFloat(e[Open], 64)
// 			high, _ := strconv.ParseFloat(e[High], 64)
// 			low, _ := strconv.ParseFloat(e[Low], 64)
// 			close, _ := strconv.ParseFloat(e[Close], 64)
// 			Change, _ := strconv.ParseFloat(e[Change], 64)
// 			volume, _ := strconv.ParseFloat(strings.Replace(e[Volume], ",", "", -1), 64)
// 			firstCollationStockDay := firstDatabase.FirstCollationStockDay{
// 				Stock_no:   stock_no,
// 				Stock_date: strings.Replace(e[StockDate], "/", "", -1),
// 				Open:       open,
// 				High:       high,
// 				Low:        low,
// 				Close:      close,
// 				Change:     Change,
// 				Volume:     volume,
// 			}
// 			firstCollationStockDays = append(firstCollationStockDays, firstCollationStockDay)
// 		}
// 		firstDatabase.InsertFirstCollationStockDay(firstCollationStockDays)

// 	}

// }
