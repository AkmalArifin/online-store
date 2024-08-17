package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB(account, password string) {
	dataSource := account + ":" + password + "@tcp(localhost:3306)/online_store"

	var err error
	DB, err = sql.Open("mysql", dataSource)

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
}
