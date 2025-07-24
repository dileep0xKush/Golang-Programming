package loop

import "fmt"

func LoopThroughStringCharacters() {

	word := "GoLang"
	for i, ch := range word {
		fmt.Printf("Index: %d, Character: %c\n", i, ch)
	}

}
