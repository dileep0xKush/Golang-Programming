package arrays

import "fmt"

func PassArraytoFunction() {

	arr := [3]int{1, 2, 3}
	printArray(arr)
}

func printArray(a [3]int) {
	for _, v := range a {
		fmt.Println(v)
	}
}
