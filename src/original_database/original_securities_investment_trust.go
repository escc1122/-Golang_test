package original_database

import (
	db "Golang_test/src/db"
	"encoding/json"
)

type OriginalSecuritiesInvestmentTrust struct {
	tableName struct{} `sql:"original.original_securities_investment_trust"`
	Date      string
	Json_data struct {
		Data [][]string
	}
}

// type StockJsonData struct {
// 	Data   [][]string
// 	Fields []string
// }

func InsertSecuritiesInvestmentTrust(date string, stockJsonData string) string {
	conn := db.GetConnect()
	defer conn.Close()
	var aaa map[string]interface{}
	json.Unmarshal([]byte(stockJsonData), &aaa)

	_, err := conn.Exec(`
	INSERT INTO original.original_securities_investment_trust(date, json_data)VALUES (?, ?)
	`, date, aaa)

	if err != nil {
		panic(err)
	}
	return ""
}

func SelectSecuritiesInvestmentTrust(start_date string, end_date string) []OriginalSecuritiesInvestmentTrust {
	conn := db.GetConnect()
	defer conn.Close()

	var originalSecuritiesInvestmentTrusts []OriginalSecuritiesInvestmentTrust

	err := conn.Model(&originalSecuritiesInvestmentTrusts).Where("date>=? and date<=?", start_date, end_date).Order("date DESC").Select()
	//return posts, err

	if err != nil {
		panic(err)
	}

	return originalSecuritiesInvestmentTrusts
}
