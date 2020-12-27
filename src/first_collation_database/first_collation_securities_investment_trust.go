package first_collation_database

import (
	db "Golang_test/src/db"
)

type FirstCollationSecuritiesInvestmentTrust struct {
	tableName           struct{} `sql:"first_collation.first_collation_securities_investment_trust"`
	Stock_no            string
	Stock_date          string
	Stock_name          string
	Buy_trading_volume  int64
	Sell_trading_volume int64
	Difference          int64
}

func InsertFirstCollationSecuritiesInvestmentTrust(firstCollationSecuritiesInvestmentTrusts []FirstCollationSecuritiesInvestmentTrust) string {
	conn := db.GetConnect()
	defer conn.Close()

	for _, e := range firstCollationSecuritiesInvestmentTrusts {
		_, err := conn.Model(&e).Insert()
		if err != nil {
			panic(err)
		}
	}
	return ""
}
