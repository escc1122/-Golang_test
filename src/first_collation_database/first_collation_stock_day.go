package first_collation_database

import (
	db "Golang_test/src/db"
)

type FirstCollationStockDay struct {
	//tableName struct{} `sql:"original.original_stock_day"`
	tableName      struct{} `sql:"first_collation.first_collation_stock_day"`
	Stock_no       string
	Stock_date     string
	Open           float64
	High           float64
	Low            float64
	Close          float64
	Volume         float64
	Change         float64
	Percent_change float64
	MA_5           float64
	MA_10          float64
	MA_20          float64
	MA_60          float64
	//stock_json_data map[string]interface{}
}

func InsertFirstCollationStockDay(firstCollationStockDays []FirstCollationStockDay) string {
	conn := db.GetConnect()
	defer conn.Close()

	for _, e := range firstCollationStockDays {
		_, err := conn.Model(&e).Insert()
		if err != nil {
			panic(err)
		}
	}
	return ""
}

func GetAllFirstCollationStockDay(Stock_no string) []FirstCollationStockDay {
	conn := db.GetConnect()
	defer conn.Close()
	var firstCollationStockDays []FirstCollationStockDay
	conn.Model(&firstCollationStockDays).Where("stock_no=?", Stock_no).Select()
	return firstCollationStockDays
}
