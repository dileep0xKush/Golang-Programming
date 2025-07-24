package logical

import "fmt"

func MultipleANDConditions() {

	score := 85
	attendance := 90
	fmt.Println("Passes class:", score >= 60 && attendance >= 75)

}
