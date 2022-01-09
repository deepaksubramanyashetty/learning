package main

import "fmt"

func main() {
	c := make(chan int)
	//send
	go foo(c)

	//recieve
	bar(c)

	fmt.Println("About to exit")
}

// send

func foo(a chan<- int) {
	a <- 42
}

//recieve
func bar(a <-chan int) {
	fmt.Println(<-a)
}
