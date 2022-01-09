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
	fmt.Println("i am ", s.first, s.last, "- the secreteagent speak")
}

func (p person) speak() {
	fmt.Println("i am ", p.first, p.last, "- the person speak")
}

//interface
// keywork type human identifier and interface type.
//human is type of speak now
type human interface {
	speak()
}

func bar(h human) {
	fmt.Println("i am human", h)
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

	p1 := person{
		first: "Deepak",
		last:  "shetty",
	}
	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	bar(sa1)
	bar(sa2)
	bar(p1)
}
