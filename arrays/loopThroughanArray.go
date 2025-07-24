package arrays

import "fmt"

func LoopThroughanArray() {

	arr := [4]int{2, 4, 6, 8}
	for i := 0; i < len(arr); i++ {
		fmt.Println("Element", i, ":", arr[i])
	}

}
