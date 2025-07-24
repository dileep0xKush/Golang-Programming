package main

import "fmt"

func swappingWithoutTempVariableArithmetic() {
	a, b := 5, 10
	a = a + b
	b = a - b
	a = a - b
	fmt.Println("After Swap: a =", a, ", b =", b)
}
