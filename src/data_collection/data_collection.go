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

func getData(url string) string {
	fmt.Println(url)
	res, _ := http.Get(url)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return string(body)
}

/**
投信
https://www.twse.com.tw/fund/TWT44U?response=json&date=20201225&_=1608955449331
**/
func GetSecuritiesInvestmentTrust(date string) string {
	url := "https://www.twse.com.tw/fund/TWT44U?response=json&date=" + date + "&_=1608955449331"
	return getData(url)
}
