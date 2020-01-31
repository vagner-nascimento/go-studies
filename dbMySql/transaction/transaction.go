package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, error := sql.Open("mysql", "root:admin@/gocourse")
	if error != nil {
		panic(error)
	}

	defer db.Close()

	transaction, _ := db.Begin()                                                // Start a transaction
	stmt, _ := transaction.Prepare("insert into users(id, name) values (?, ?)") // Get statement from transaction

	stmt.Exec(202, "Charles")
	stmt.Exec(203, "John")

	//stmt.Exec(1, "Braian") // It will return a error because id 1 already exists
	_, err := stmt.Exec(204, "Braian")

	if err != nil { // If you don't handle the error, the transaction will be commit bellow. REMEMBER: Go doesn't throws exception, it should be implicit handled
		transaction.Rollback()
		log.Fatal(error)
	}

	transaction.Commit() // If Rollback was called, the Commit will not be executed
}
