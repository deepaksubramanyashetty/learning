package main

import "fmt"

func main() {
	//channel onto it can put integer
	C := make(chan int)

	//on C put on number 42
	C <- 42
	//off of C
	fmt.Println(<-C)
}
