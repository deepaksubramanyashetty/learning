package main

import "fmt"

func main() {
	//channel onto it can put integer
	C := make(chan int)

	go func() {
		C <- 42
	}()

	//off of C
	fmt.Println(<-C)
}
