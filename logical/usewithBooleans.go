package logical

import "fmt"

func UsewithBooleans() {

	var isOnline bool = true
	var isMuted bool = false

	fmt.Println("Can speak:", isOnline && !isMuted)

}
