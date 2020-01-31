package main

import (
	"log"
	"net/http"
)

func main() {
	//IMPORTANT: you should run it by terminal inside this folder({your_go_home}/src/github.com/vagner-nascimento/go-studies/http/static/static.go), otherwise it will return 404 on browser
	fs := http.FileServer(http.Dir("public")) // Dir starts form folder where its file was

	http.Handle("/", fs)
	log.Println("Server running at port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil)) // It gives access to http://localhost:3000
}
