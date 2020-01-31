package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, error := sql.Open("mysql", "root:admin@/gocourse")
	if error != nil {
		panic(error)
	}

	defer db.Close()

	stmt, _ := db.Prepare("update users set name = ? where id = ?")
	stmt.Exec("Leopold", 1)
	stmt.Exec("Yeneffer", 2)
	result, err := stmt.Exec("Dont Exists", 999) // It doesn't returns error
	rows, _ := result.RowsAffected()             // If you need to handle if none row was affected, use RowsAffected
	fmt.Printf("Rows updated: %d\n", rows)
	if rows <= 0 {
		fmt.Println("None register was updated")
	}

	if err != nil {
		log.Fatal(error)
	}

	stmt1, _ := db.Prepare("delete from users where id = ?")
	stmt1.Exec(3)

	result1, err1 := stmt1.Exec(999) // It doesn't returns error as well
	rows, _ = result1.RowsAffected()
	fmt.Printf("Deleted rows: %d", rows)

	if err1 != nil {
		log.Fatal(err1)
	}
}
