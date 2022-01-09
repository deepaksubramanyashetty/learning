// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {

	s := `[{"First":"james","Last":"bond","Age":22},{"First":"mrs","Last":"monneypeny","Age":20}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	//people := []person{}
	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Print all data", people)

	for i, v := range people {
		fmt.Println("\n Person Number", i)
		fmt.Println(v.First, v.Last, v.Age)
	}

}
