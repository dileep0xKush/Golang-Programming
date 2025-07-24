package loop

import "fmt"

func InfiniteLoopwithExit() {

	count := 0
	for {
		fmt.Println("Looping...")
		count++
		if count == 3 {
			break
		}
	}
}
