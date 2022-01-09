// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "james",
		Last:  "bond",
		Age:   22,
	}
	p2 := person{
		First: "mrs",
		Last:  "monneypeny",
		Age:   20,
	}

	people := []person{p1, p2}
	fmt.Println(people)

	br, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(br))
}

//taken our code converted into json
