package main

import "fmt"

func main() {

	f := func() {
		fmt.Println("The 1st function expression")
	}
	f()

	f2 := func(x int) {
		fmt.Println("second funtion expression", x)
	}

	f2(42)
}
