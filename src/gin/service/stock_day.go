package service

import (
	first_collation_database "Golang_test/src/first_collation_database"
	"strconv"
	"strings"
)

type StockDayStruct struct {
	Title []string
	Data  [][]string
}

func GetStockDay() StockDayStruct {
	var stockDayStruct StockDayStruct
	stockDayStruct.Title = []string{
		"日期",
		"股票代號",
		"股票名稱",
		"開盤價",
		"最高價",
		"最低價",
		"收盤價",
		"漲跌",
		"漲幅(%)",
		"成交量",
		"5MA",
		"20MA",
		"60MA",
		"K(9)",
		"D(9)",
		"RSI(5)",
		"RSI(10)",
		"DIF",
		"MACD",
		"DIF-MACD",
		"+DI(14)",
		"-DI(14)",
		"ADX(14)",
	}

	firstCollationStockDays := first_collation_database.GetAllFirstCollationStockDay("2330")
	var aaaa [][]string
	for _, e := range firstCollationStockDays {
		m := []string{strings.Replace(e.Stock_date, "109", "2020", 1), e.Stock_no,
			"",
			strconv.FormatFloat(e.Open, 'f', -2, 64),
			strconv.FormatFloat(e.High, 'f', -2, 64),
			strconv.FormatFloat(e.Low, 'f', -2, 64),
			strconv.FormatFloat(e.Close, 'f', -2, 64),
			strconv.FormatFloat(e.Change, 'f', -2, 64),
			"",
			strconv.FormatFloat(e.Volume, 'f', -2, 64),
			"0",
			"0",
			"0",
			"0",
			"0",
			"0",
			"0",
			"0",
			"0",
			"0",
			"0",
			"0",
			"0",
		}
		aaaa = append(aaaa, m)
	}

	stockDayStruct.Data = aaaa

	//fmt.Print(firstCollationStockDays)

	//b,_ :=json.Marshal(stockDayStruct)

	return stockDayStruct
}
