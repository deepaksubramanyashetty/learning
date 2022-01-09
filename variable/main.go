package main

import "fmt"

var y = 43

// Declare there is VARIABLE with the IDENTIFIER "z"
//and that the VARIABLE with the IDENTIFIER "z" is of TYPE int
//ASSIGNs the ZERO VALUE of TYPE int to "z"
//false for booleans, 0 for integers, 0.0 for floats, "" for strings
//and nil for pointers, functions, interfaces, slices, channels and maps

var z int

func main() {
	// short declaration operator
	// Declare a variable and assign a VALUE (of certain type)
	x := 42
	fmt.Println(x)
	fmt.Println(y)

	foo()
	fmt.Println(z)
}

func foo() {
	fmt.Println("again", y)
}
