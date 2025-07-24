package condition

import "fmt"

func Nestedif() {

	age := 25
	hasID := true
	if age >= 18 {
		if hasID {
			fmt.Println("Entry allowed")
		} else {
			fmt.Println("ID required")
		}
	}

}
