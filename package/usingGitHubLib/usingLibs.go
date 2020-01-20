// The commando to install a lib from GitHub is (from go path your_user/go): go get -u github.com/repo_of_lib/libname
// In this case: go get -u github.com/vagner-nascimento/gostudiesarea
// The compiled files goes to: go/pkg/linux_amd64/github.com/vagner-nascimento/gostudiesarea.a
// The source code goes to: go/src/github.com/vagner-nascimento/gostudiesarea/main.go
package main

import (
	"fmt"

	"github.com/vagner-nascimento/gostudiesarea"
	"github.com/vagner-nascimento/gostudiesarea/subpack"
)

func main() {
	fmt.Println(gostudiesarea.Circ(4.0))
	fmt.Println(gostudiesarea.Circ(6.0))
	fmt.Println(gostudiesarea.Rect(5.0, 2.0))
	fmt.Println(gostudiesarea.Pi)

	subpack.PrintBah()
}
