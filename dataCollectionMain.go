package main

import (
//"fmt"
)

import (
	ddddd "Golang_test/src/data_collection"
	//db_test "Golang_test/src/db"
	db_test "Golang_test/src/db"
)

func main() {
	res := ddddd.Stock_day("2330", "20201207")
	//url:="https://www.twse.com.tw/exchangeReport/STOCK_DAY?response=json&date=20201208&stockNo=2330&_=1607163124419"
	//res,_ := http.Get(url)
	//defer res.Body.Close()
	//body, _ := ioutil.ReadAll(res.Body)
	//db_connect.GetConnect()
	//var aaa map[string]interface{}
	//json.Unmarshal([]byte(res), aaa)

	db_test.StockDayInsert("2330", "20201207", res)
	//db_test.GetConnect()
	//fmt.Println(res)
}
