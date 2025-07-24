package condition

import "fmt"

func Ifwithboolean() {

	isOnline := true
	if isOnline {
		fmt.Println("User is online")
	} else {
		fmt.Println("User is offline")
	}
}
