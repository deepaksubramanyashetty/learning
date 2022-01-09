package main

import "fmt"

func main() {
	//this one does not have name
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "james",
		last:  "bond",
		age:   32,
	}

	fmt.Println(p1)
}
