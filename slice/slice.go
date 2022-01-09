package main

import "fmt"

func main() {
	//array
	var i [10]int
	fmt.Println(i)

	//COMPOSITE LITERAL
	x := []int{0, 5, 3, 5, 6}
	fmt.Println(x)
	// a SLICE allows you to group together VALUES of the same TYPE

	index()
}

func index() {
	x := []int{1, 2, 5, 6}
	fmt.Println(len(x))
	fmt.Println(x)
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])

	for i, v := range x {
		fmt.Println(i, v)
	}

	for i := 0; i <= 3; i++ {
		fmt.Println(i, x[i])
	}

	//append to slice
	z := []int{2, 7, 8, 9}
	fmt.Println(z)
	z = append(z, 12, 34, 54)
	fmt.Println(z)

	y := []int{64, 74, 84}
	fmt.Println(y)
	//append all y value to z
	z = append(z, y...)
	fmt.Println(z)

	//delete from a slice
	// basically appending from 0 to 1 postion with 4th position till end
	z = append(z[:2], z[4:]...)
	fmt.Println(z)

	//multi Slice
	jb := [][]int{z, y}
	fmt.Println(jb)

}
