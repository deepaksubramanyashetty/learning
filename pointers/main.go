package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a) //gives you the address stored of a

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(b)
	fmt.Println(*b) //* gives the values store at address b when you have the address

}
