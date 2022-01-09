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
	case (2 == 3):
		fmt.Println("aslo true, does it print?")
	}
}
