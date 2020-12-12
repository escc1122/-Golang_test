package original_database

import (
	db "Golang_test/src/db"
	"encoding/json"
)

type originalStockDay struct {
	//tableName struct{} `sql:"original.original_stock_day"`
	tableName       struct{} `sql:"original_stock_day"`
	stock_no        string   `sql:",pk"`
	stock_date      string
	stock_json_data map[string]interface{}
}

func InsertStockDay(stockNo string, date string, stockJsonData string) string {
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
		INSERT INTO original.original_stock_day (stock_no, stock_date,stock_json_data) VALUES (?, ?,?)
	`, stockNo, date, aaa)

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

func SelectStockDayForLastDate(stockNo string, date string) string {
	conn := db.GetConnect()
	defer conn.Close()

	return ""
}
