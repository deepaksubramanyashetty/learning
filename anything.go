package main

import "fmt"

func main() {
	fmt.Println("this is good")

	foo()

	fmt.Println("Something else after foo")

	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("I am in foo")

}

func bar() {
	fmt.Println("then we exited")
}
