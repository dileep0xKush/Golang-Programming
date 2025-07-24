package logical

import "fmt"

func MultipleORConditions() {

	temp := 39
	isRaining := true
	fmt.Println("Stay home:", temp > 35 || isRaining)

}
