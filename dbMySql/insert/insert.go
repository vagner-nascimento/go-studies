package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, error := sql.Open("mysql", "root:admin@/gocourse")
	if error != nil {
		panic(error)
	}

	defer db.Close()

	statement, _ := db.Prepare("insert into users(name) values (?)") // Preparing an statement with a insert command, you can pass as much values as you want
	statement.Exec("Mary")                                           // Executing the insert though the created statement
	statement.Exec("Gerald")                                         // It should be done line by line

	respose, _ := statement.Exec("Peter") // The exec returns some infos about the execution
	id, _ := respose.LastInsertId()       // It brings the id of last statement exec

	fmt.Printf("Last id inserted: %d\n", id)

	rowsAffected, _ := respose.RowsAffected() // It brings the number of affected rows of the last statement exec

	fmt.Printf("Rows inserted: %d\n", rowsAffected)
}
