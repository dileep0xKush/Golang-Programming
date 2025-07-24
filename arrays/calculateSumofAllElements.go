package arrays

import "fmt"

func CalculateSumofAllElements() {

	arr := [5]int{10, 20, 30, 40, 50}
	sum := 0
	for _, v := range arr {
		sum += v
	}
	fmt.Println("Sum:", sum)

}
