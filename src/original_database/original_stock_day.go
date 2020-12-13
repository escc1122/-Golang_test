package original_database

import (
	db "Golang_test/src/db"
	"encoding/json"
)

type OriginalStockDay struct {
	tableName struct{} `sql:"original.original_stock_day"`
	//tableName       struct{} `sql:"original_stock_day"`
	Stock_no              string
	Stock_date            string
	Stock_date_year_month string
	Stock_json_data       map[string]interface{}
}

func InsertStockDay(stockNo string, date string, stockDateYearMonth string, stockJsonData string) string {
	conn := db.GetConnect()
	defer conn.Close()
	var aaa map[string]interface{}
	json.Unmarshal([]byte(stockJsonData), &aaa)

	//_, err := db.Model(&original_stock_day{
	//	stock_no:   stockNo,
	//	stock_date: date,
	//	//stock_json_data: aaa,
	//}).Insert()

	//test1 := &original_stock_day{
	//	stock_no:   stockNo,
	//	stock_date: date,
	//	//stock_json_data: aaa,
	//}

	_, err := conn.Exec(`
		INSERT INTO original.original_stock_day (stock_no, stock_date,stock_json_data,stock_date_year_month) VALUES (?, ?,?,?)
	`, stockNo, date, aaa, stockDateYearMonth)

	//err := db.Select(&original_stock_day{
	//	stock_no:   stockNo,
	//	//stock_date: date,
	//	//stock_json_data: aaa,
	//})

	if err != nil {
		panic(err)
	}
	return ""
}

func SelectStockDayForLastDate(stockNo string, stockDateYearMonth string) []OriginalStockDay {
	conn := db.GetConnect()
	defer conn.Close()

	var originalStockDays []OriginalStockDay

	err := conn.Model(&originalStockDays).Where("stock_no=? and stock_date_year_month=?", stockNo, stockDateYearMonth).Order("stock_date DESC").Limit(1).Select()
	//return posts, err

	if err != nil {
		panic(err)
	}

	return originalStockDays
}
