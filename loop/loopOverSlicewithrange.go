package loop

import "fmt"

func LoopOverSlicewithrange() {

	numbers := []int{10, 20, 30}
	for i, v := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}
}
