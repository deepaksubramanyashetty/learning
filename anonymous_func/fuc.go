package main

import "fmt"

func main() {
	foo()

	//anonymous function with out argument
	func() {
		fmt.Println("anonymous function")
	}()

	//anonymous function with argument
	func(x int) {
		fmt.Println("the argument to anonymous function", x)
	}(42)
}

func foo() {
	fmt.Println("i am in foo")
}
