package bitwise

import "fmt"

func BitwiseANDAssignment() {

	x := 12 // 1100
	x &= 10 // 1010 → result: 1000 (8)
	fmt.Println("x &= 10 =>", x)

}
