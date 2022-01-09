package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretagent struct {
	person
	ltk bool
}

// func (r receiver) identifier(parameters) (return(s)) { code... }

func (s secretagent) speak() {
	fmt.Println("i am ", s.first, s.last)
}
func main() {
	sa1 := secretagent{
		person: person{
			"james",
			"bond",
		},
		ltk: true,
	}
	sa2 := secretagent{
		person: person{
			"miss",
			"money penny",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
}
