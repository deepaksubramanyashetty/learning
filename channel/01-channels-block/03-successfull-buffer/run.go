package main

import "fmt"

func main() {
	//channel onto it can put integer and buffer channel 1
	C := make(chan int, 1)

	C <- 42

	//off of C
	fmt.Println(<-C)
}
