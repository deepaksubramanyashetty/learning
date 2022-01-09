package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	login1 := `password123`

	err = bcrypt.CompareHashAndPassword(bs, []byte(login1))
	if err != nil {
		fmt.Println("you cant log in")
		return
	}
	fmt.Println("you are logged in")
}
