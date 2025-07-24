package condition

import "fmt"

func Switchwithnocondition() {

	score := 90
	switch {
	case score >= 90:
		fmt.Println("Excellent")
	case score >= 75:
		fmt.Println("Very Good")
	default:
		fmt.Println("Needs improvement")
	}

}
