package service

import (
	dataCollection "Golang_test/src/data_collection"
	originalDatabase "Golang_test/src/original_database"
)

func StockDayWebToOriginal(stockNo string, date string) {
	res := dataCollection.GetStockDay(stockNo, date)
	originalDatabase.InsertStockDay(stockNo, date, res)
}

//func stockDayWebToOriginal(stockNo string, date string) {
//	res := dataCollection.GetStockDay(stockNo, date)
//	originalDatabase.InsertStockDay(stockNo, date, res)
//}
