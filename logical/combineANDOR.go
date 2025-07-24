package logical

import "fmt"

func CombineANDOR() {

	isMember := true
	hasCoupon := false
	fmt.Println("Gets discount:", isMember || hasCoupon && true)
}
