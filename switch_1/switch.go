package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("This should not print")
	case (2 == 4):
		fmt.Println("This should not print")
	case true:
		fmt.Println("this should print") // this should print
		fallthrough
	case (2 == 2):
		fmt.Println("aslo true, does it print") // aslo true, does it print
	default:
		fmt.Println("this is default")
	}
}
