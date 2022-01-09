package main

import "fmt"

//type person which is struct
type person struct {
	first string
	last  string
	age   int
}

//type secretagent which is struct
type secretagent struct {
	//embeding type person in to other type secreteagent
	person
	ltk bool
}

func main() {
	//value to person
	sa1 := secretagent{
		//field person of: type person
		person: person{
			first: "james",
			last:  "bond",
			age:   32,
		},
		ltk: true,
	}

	p2 := person{
		first: "miss",
		last:  "monney penny",
		age:   27,
	}

	fmt.Println(sa1)
	fmt.Println(p2)

	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)
	fmt.Println(p2.first, p2.last)

}
