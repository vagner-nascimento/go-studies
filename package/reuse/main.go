package go_studies_area

import (
	"fmt"

	"github.com/vagner-nascimento/area"
)

func main() {
	fmt.Println(area.Circ(6.0))
	fmt.Println(area.Rect(5.0, 2.0))
	fmt.Println(area.Pi)
	// fmt.Println(area._EqTriangle(5.0, 2.0)) // It doesn't works because the function _EqTriangle is private on the area package
}
