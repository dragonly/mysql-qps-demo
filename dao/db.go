package dao

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	dsn := "root:pass@tcp(192.168.1.29:3306)/performance_schema?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		panic("failed to connect database")
	}
}
