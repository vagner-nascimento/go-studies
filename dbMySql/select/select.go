package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	id   int
	name string
}

func main() {
	db, error := sql.Open("mysql", "root:admin@/gocourse")
	if error != nil {
		panic(error)
	}
	defer db.Close()

	rows, _ := db.Query("select id, name from users where id > ?", 3)
	defer rows.Close() // It should be closed as well

	var users = []user{}
	for rows.Next() {
		var u user

		rows.Scan(&u.id, &u.name) // It maps the rows columns into the passed reference's variable
		users = append(users, u)
	}

	fmt.Println(users)
}
