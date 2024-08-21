package main

import (
	"fmt"
	"strings"
)

func main() {
	// mapp
	intmap := mapp([]int{1, 2, 3, 4, 5}, func(v int) int { return v * 2 })
	fmt.Println("intmap", intmap)

	strmap := mapp([]string{"Vagner N", "Jessica G", "John B"}, func(v string) string { return "Hello " + v })
	fmt.Println("strmap", strmap)

	type person struct {
		Name string
		Age  int
		Id   string
	}

	people := []person{
		{Name: "Vagner", Age: 36, Id: "939.928.120-57"},
		{Name: "Jessica", Age: 32, Id: "070.774.710-47"},
		{Name: "Adri", Age: 26, Id: "229.835.460-04"},
		{Name: "Bern", Age: 28, Id: "974.803.110-10"},
	}
	cleanCpf := func(p person) person {
		p.Id = strings.ReplaceAll(p.Id, ".", "")
		p.Id = strings.ReplaceAll(p.Id, "-", "")

		return p
	}

	pmap := mapp(people, cleanCpf)
	fmt.Println("pmap", pmap)

	// filter
	intmap = filter([]int{1, 2, 3, 4, 5}, func(v int) bool { return v%2 == 0 })
	fmt.Println("intmap", intmap)

	pmap = filter(people, func(p person) bool { return p.Age < 30 })
	fmt.Println("pmap", pmap)
}

// mapp
func mapp[T any](list []T, mapper func(it T) T) (mapped []T) {
	for _, v := range list {
		mapped = append(mapped, mapper(v))
	}
	return
}

// filter
func filter[T any](list []T, filter func(it T) bool) (mapped []T) {
	for _, v := range list {
		if filter(v) {
			mapped = append(mapped, v)
		}
	}
	return
}
