package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("C:\\Users\\DSubramanyaShett\\Documents\\goworkspace\\src\\github.com\\deepaksubramanyashetty\\Learning\\examples\\error_handling\\00-check-error\\03\\names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs))
}
