package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // The underscore before this import means that it will be implicitly used, but never explicit
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}

	return result
}

func main() {
	db, error := sql.Open("mysql", "root:admin@/") // It will use the mysql lib implicitly. If you remove the undercore, no errors will happen, but WILL NOT WORK
	defer db.Close()

	if error != nil {
		panic(error)
	}

	exec(db, "create database if not exists gocourse")
	exec(db, "use gocourse")
	exec(db, "drop table if exists users")
	exec(db, `create table users (
		id integer auto_increment,
		nome varchar(80),
		PRIMARY KEY (id)
	)`)
}
