package db

import "github.com/go-pg/pg"

const (
	//host     = "localhost"
	//port     = 5432
	user     = "admin"
	password = "123456"
	dbname   = "admin"
)

func GetConnect() *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     user,
		Password: password,
		Database: dbname,
	})
	//defer db.Close()
	return db
}
