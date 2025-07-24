package condition

import "fmt"

func Switchwithfallthrough() {

	x := 2
	switch x {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
		fallthrough
	case 3:
		fmt.Println("Three (fallthrough)")
	default:
		fmt.Println("Other")
	}

}
