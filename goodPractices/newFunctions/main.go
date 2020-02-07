package main

import (
	"encoding/json"
	"fmt"
	"github.com/vagner-nascimento/go-studies/goodPractices/newFunctions/multitypes"
	"github.com/vagner-nascimento/go-studies/goodPractices/newFunctions/onetype"
)

func main() {
	oneTypeExample()
	multipleTypesExample()
}

func oneTypeExample() {
	v := onetype.New()
	vBytes, _ := json.Marshal(v)

	fmt.Printf("One Type New function %s\n", string(vBytes))
}

func multipleTypesExample() {
	one := multitypes.NewOne()
	two := multitypes.NewTwo()

	oneBytes, _ := json.Marshal(one)
	twoBytes, _ := json.Marshal(two)

	fmt.Printf("Multiple Types New function. Type One %s\n", string(oneBytes))
	fmt.Printf("Multiple Types New function. Type Two %s\n", string(twoBytes))
}
