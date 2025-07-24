package comparison

import "fmt"

func CheckEqualityBetweenDifferentTypes() {

	var a int = 10
	var b float64 = 10.0
	fmt.Println("a == int(b):", a == int(b))
}
