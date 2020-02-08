package main

import (
	"encoding/json"
	"fmt"
)

type i interface {
	PrintMe()
}

type t struct {
	V string
}

func (o t) PrintMe() {
	fmt.Println("printing o")
}

type t1 struct {
}

func (o t1) PrintMe() {
	fmt.Println("printing s")
}

func stringfy(v interface{}) string {
	s, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return string(s)
}

func main() {
	var v i
	v = t{V: "value of v"}
	tv, ok := v.(i)
	if ok {
		fmt.Println("converted to i (interface)", stringfy(tv))
	} else {
		fmt.Println("conversion to i (interface) failed")
	}

	tv, ok = v.(t)
	if ok {
		fmt.Println("converted to t (type)", stringfy(tv))
	} else {
		fmt.Println("conversion to t (type) failed")
	}

	tv1, converted := v.(t1)
	if converted {
		fmt.Println("converted to t1 (type)", stringfy(tv1))
	} else {
		fmt.Println("conversion to t1 failed (type)")
	}
}
