package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetDb() (*sql.DB, error) {
	return sql.Open("mysql", "root:admin@/gocourse")
}
