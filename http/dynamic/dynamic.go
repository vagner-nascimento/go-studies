package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func currentDateTime(res http.ResponseWriter, req *http.Request) {
	/*
		This is the way that Go formats the date, using number(in other languages are used specif format strings like "dd/mm/yyyy hh:mm:ss"
		IMPORTANT: There pre defined values to do it, can't be any data.
		To get more information an valid values visit: https://programming.guide/go/format-parse-string-time-date-example.html
	*/
	nowStr := time.Now().Format("02/01/2006 03:04:05")

	fmt.Fprintf(res, "<h1>Current hour: %s<h1>", nowStr) // Write text on the passed writer(in this case, the ResponseWriter, which will write it on browser
}

/*
	This is considered dynamic because call a function that process some data and return a value which will change
*/
func main() {
	http.HandleFunc("/currentDate", currentDateTime) // It will send the res and req to the method when this address was called
	log.Println("Server running at port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil)) // None handler was passed because it was set above
}
