package main

import (
	serviceDataCollection "Golang_test/src/service/data_collection"
	"strconv"
	"time"
)

func main() {
	for i := 1; i <= 12; i++ {
		//fmt.Println(startTime.Format("20060102"))
		var mon string = strconv.Itoa(i)
		if len(mon) == 1 {
			mon = "0" + mon
		}
		serviceDataCollection.StockDayWebToOriginal("2330", "2020"+mon+"01", "2020"+mon)
		time.Sleep(time.Duration(2) * time.Second)
	}
}
