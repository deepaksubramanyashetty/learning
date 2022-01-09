package main

import "fmt"

//type person
type person struct {
	first string
	last  string
	age   int
}

func main() {
	//value to person
	p1 := person{
		first: "james",
		last:  "bond",
		age:   32,
	}

	p2 := person{
		first: "miss",
		last:  "monney penny",
		age:   27,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last)

}
