package main

import (
	ddddd "Golang_test/src/data_collection"
	//"fmt"
	//"time"

	//db_test "Golang_test/src/db"
	original_database "Golang_test/src/original_database"
)

func main() {
	//startTime := time.Now()
	//
	//for i := 0; i <3; i++ {
	//	fmt.Println(startTime.Format("20060102"))
	//	data_collection("2330",startTime.Format("20060102"))
	//	startTime = startTime.AddDate(0, 0, -1)
	//	time.Sleep(time.Duration(2)*time.Second)
	//}

	data_collection("2330", "20201101")
	//res := ddddd.Stock_day("2330", "20201207")
	//url:="https://www.twse.com.tw/exchangeReport/STOCK_DAY?response=json&date=20201208&stockNo=2330&_=1607163124419"
	//res,_ := http.Get(url)
	//defer res.Body.Close()
	//body, _ := ioutil.ReadAll(res.Body)
	//db_connect.GetConnect()
	//var aaa map[string]interface{}
	//json.Unmarshal([]byte(res), aaa)

	//db_test.StockDayInsert("2330", "20201207", res)
	//db_test.GetConnect()
	//fmt.Println(res)
}

func data_collection(stockNo string, date string) {
	res := ddddd.GetStockDay(stockNo, date)
	original_database.InsertStockDay(stockNo, date, res)
}
