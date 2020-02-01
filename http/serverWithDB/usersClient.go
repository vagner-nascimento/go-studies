package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	//_ "github.com/go-sql-driver/mysql"
)

// UserHandler handler the users request
func UserHandler(res http.ResponseWriter, req *http.Request) {
	/*
		It will remove everything from the string, instead the "/users/".
		In this case, will remain jus the user is(if it was informed on the URL's path)
	*/
	userId, _ := strconv.Atoi(strings.TrimPrefix(req.URL.Path, "/users/"))

	res.Header().Set("Content-Type", "application/json")

	switch {
	case req.Method == http.MethodGet && userId > 0:
		user := getUserById(userId)

		if user.Id <= 0 {
			res.WriteHeader(http.StatusNotFound)
			fmt.Fprint(res, "User not Found")
			return
		}

		fmt.Fprint(res, GetUserJsonString(user))
	case req.Method == http.MethodGet:
		users := getAllUsers()
		fmt.Fprint(res, GetUserJsonString(users))
	default:
		res.WriteHeader(http.StatusNotFound)
		fmt.Fprint(res, "Resource not found")
	}
}

func getUserById(id int) User {
	db, err := GetDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var u User
	db.QueryRow("select id, name from users where id = ?", id).Scan(&u.Id, &u.Name) // This db method returns only one row

	return u
}

func getAllUsers() []User {
	db, err := GetDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("select id, name from users")
	defer rows.Close()

	var users = []User{}
	for rows.Next() {
		var u User

		rows.Scan(&u.Id, &u.Name)
		users = append(users, u)
	}

	return users
}
