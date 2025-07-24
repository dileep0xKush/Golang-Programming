package logical

import "fmt"

func NestedLogicalExpressions() {

	a := 10
	b := 5
	c := 20

	result := a > b && (c > a || b < 0)
	fmt.Println("Result:", result)

}
