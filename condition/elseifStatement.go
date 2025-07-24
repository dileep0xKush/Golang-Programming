package condition

import "fmt"

func ElseifStatement() {

	marks := 85
	if marks >= 90 {
		fmt.Println("Grade: A")
	} else if marks >= 75 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C or below")
	}

}
