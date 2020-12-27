package service

import (
	dataCollection "Golang_test/src/data_collection"
	originalDatabase "Golang_test/src/original_database"
	"fmt"
	"time"
)

func SecuritiesInvestmentTrustToOriginal(startTime time.Time, add int) {
	for i := 1; i <= add; i++ {
		fmt.Println("startTime = ", startTime)
		// year, month, day := startTime.Date()
		weekday := startTime.Weekday().String()
		if weekday == "Saturday" || weekday == "Sunday" {
			fmt.Println("in = ")
			startTime = startTime.AddDate(0, 0, 1)
			continue
		}
		//fmt.Println("ISOWeek = ", a)
		date := startTime.Format("20060102")
		fmt.Println("date = ", date)
		res := dataCollection.GetSecuritiesInvestmentTrust(date)
		fmt.Println("res = ", res)
		originalDatabase.InsertSecuritiesInvestmentTrust(date, res)
		startTime = startTime.AddDate(0, 0, 1)
		time.Sleep(time.Duration(2) * time.Second)
	}
}
