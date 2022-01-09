package main

import "fmt"

func main() {
	jb := map[string]int{
		"james":     32,
		"mis penny": 27,
	}
	fmt.Println(jb)

	fmt.Println(jb["james"])
	fmt.Println(jb["master"])

	//checking whether master exist
	vi, oki := jb["master"]
	fmt.Println(vi)
	fmt.Println(oki)

	if vi, oki := jb["master"]; oki {
		fmt.Println("THIS IS IF PRINT", vi)
	} else {
		fmt.Println("THIS IS IF PRINT")
	}

	// adding element to MAP
	//key = value
	jb["bond"] = 33

	// printing map using range
	for key, value := range jb {
		fmt.Println(key, value)
	}

	//print slice
	xi := []int{4, 5, 6, 7}

	for i, k := range xi {
		fmt.Println(i, k)
	}

	//delete from map
	// key
	delete(jb, "bond")
	fmt.Println(jb)
}
