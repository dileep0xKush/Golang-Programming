// Ternary-like logic using if-else (Go has no ternary operator)

package condition

import "fmt"

func Ternarylikelogicusingifelse() {

	x := 5
	var result string
	if x%2 == 0 {
		result = "Even"
	} else {
		result = "Odd"
	}
	fmt.Println("Number is", result)

}
