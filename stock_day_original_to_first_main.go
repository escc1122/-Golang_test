package main

import (
	"Golang_test/src/service/convert"
	"strconv"
)

func main() {
	for i := 1; i <= 12; i++ {
		//fmt.Println(startTime.Format("20060102"))
		var mon string = strconv.Itoa(i)
		if len(mon) == 1 {
			mon = "0" + mon
		}
		convert.StockDayOriginalToFirst("2330", "2020"+mon)
	}
}
