package loop

import "fmt"

func LoopOverMapwithrange() {

	person := map[string]string{
		"name": "Alice",
		"city": "Mumbai",
	}
	for key, value := range person {
		fmt.Printf("%s: %s\n", key, value)
	}

}
