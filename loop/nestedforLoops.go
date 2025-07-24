package loop

import "fmt"

func NestedforLoops() {

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 2; j++ {
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}
}
