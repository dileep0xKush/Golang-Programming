package arrays

import "fmt"

func FindMaximumValueinArray() {

	arr := [5]int{12, 25, 8, 41, 33}
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	fmt.Println("Max:", max)

}
