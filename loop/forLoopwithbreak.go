package loop

import "fmt"

func ForLoopwithbreak() {

	for i := 1; i <= 10; i++ {
		if i == 6 {
			break
		}
		fmt.Println(i)
	}
}
