package logical

import "fmt"

func LogicalAND() {

	age := 25
	hasID := true
	fmt.Println("Can enter bar:", age >= 18 && hasID)
}
