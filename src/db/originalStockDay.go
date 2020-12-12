package db

import (
	"encoding/json"
)

type original_stock_day struct {
	//tableName struct{} `sql:"original.original_stock_day"`
	tableName  struct{} `sql:"original_stock_day"`
	stock_no   string   `sql:",pk"`
	stock_date string
	//stock_json_data map[string]interface{}
}

func StockDayInsert(stockNo string, date string, stock_json_data string) string {
	db := GetConnect()
	defer db.Close()
	var aaa map[string]interface{}
	json.Unmarshal([]byte(stock_json_data), &aaa)

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

	_, err := db.Exec(`
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
