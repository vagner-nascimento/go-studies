package main

import (
	"encoding/json"
	"fmt"
)

// IMPROTANT: In Go, variables and methods created with lower case means that they are private,
// on the other hand, those that are created with upper case are public.
// On the product struct, the attributes are setted as public because they will be exported to JSON
type product struct {
	ID    int      `json:"id"`
	Name  string   `json:name`
	Price float64  `json:"price`
	Tags  []string `json:"tags`
}

func structToJson(p product) string {
	json, error := json.Marshal(p)
	if error == nil {
		return string(json)
	}
	return ""
}

func jsonToStruct(data string) product {
	var p product
	json.Unmarshal([]byte(data), &p)
	return p
}

func main() {
	p1 := product{1, "Notebook", 1899.9, []string{"Promo", "Tech"}}
	p1Json := structToJson(p1)
	fmt.Println(p1Json)

	pJson := `{"id":2,"name":"Pen","price":8.90,"tags":["School","Office"]}`
	p2 := jsonToStruct(pJson)
	fmt.Println(p2)
	fmt.Println(p2.Name)
	fmt.Println(p2.Tags)
}
