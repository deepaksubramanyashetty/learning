package main

import "fmt"

func main() {
	//channel onto it can put integer
	C := make(chan int)
	cr := make(<-chan int) //recieve
	cs := make(chan<- int) //send

	fmt.Println("----------------")
	fmt.Printf("c\t%T\n", C)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)
}
