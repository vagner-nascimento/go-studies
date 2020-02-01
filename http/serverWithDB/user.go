package main

import (
	"encoding/json"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetUserJsonString(u interface{}) string {
	userJson, _ := json.Marshal(u)
	return string(userJson)
}
