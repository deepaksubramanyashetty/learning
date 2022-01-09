package main

import "fmt"

func main() {
	switch "Bond" {
	case "MondyPenny":
		fmt.Println("THis is money penny")
	case "Bond":
		fmt.Println("James Bond")
	default:
		fmt.Println("this is james bond")
	}
	switch_2()
	switch_3()
}

func switch_2() {
	switch "Bond" {
	case "MondyPenny", "bond":
		fmt.Println("THis is money penny")
	case "Bond":
		fmt.Println("James Bond")
	default:
		fmt.Println("this defualt")
	}
}

func switch_3() {
	switch "Bond" {
	case "MondyPenny", "bond":
		fmt.Println("THis is money penny")
	default:
		fmt.Println("this is defualt")
	}

}
