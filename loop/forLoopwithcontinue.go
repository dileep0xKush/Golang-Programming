package loop

import "fmt"

func ForLoopwithcontinue() {

	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}
