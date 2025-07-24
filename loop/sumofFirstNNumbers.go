package loop

import "fmt"

func SumofFirstNNumbers() {

	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum of 1 to 10:", sum)

}
