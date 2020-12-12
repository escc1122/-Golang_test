package data_collection

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const URL = "https://www.twse.com.tw/exchangeReport/"

func GetStockDay(stockNo string, date string) string {
	var url string = URL + "/STOCK_DAY?response=json&date=" + date + "&stockNo=" + stockNo + "&_=1607163124419"
	fmt.Println(url)
	res, _ := http.Get(url)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return string(body)
	//fmt.Println(string(body))
}
