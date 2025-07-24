package arrays

import "fmt"

func LoopUsingrange() {

	arr := [3]string{"red", "green", "blue"}
	for i, v := range arr {
		fmt.Printf("Index %d → %s\n", i, v)
	}

}
