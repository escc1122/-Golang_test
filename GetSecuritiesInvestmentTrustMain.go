package main

import (
	serviceDataCollection "Golang_test/src/service/data_collection"
	"fmt"
	"time"
)

func main() {
	// d := time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)
	// year, month, day := d.Date()
	aaaaa := time.Now()
	aaaaa = aaaaa.AddDate(0, 0, -3)
	fmt.Print(aaaaa)
	serviceDataCollection.SecuritiesInvestmentTrustToOriginal(aaaaa, 3)
}
