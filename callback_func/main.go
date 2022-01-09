// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4}
	y := sum(ii...)
	fmt.Println("all number", y)

	s2 := even(sum, ii...)
	fmt.Println("event numbers", s2)

}

func sum(xi ...int) int {
	fmt.Printf("%T\n", xi)
	total := 0
	for _, v := range xi {
		total += v

	}
	return (total)
}

func even(f func(xi ...int) int, vi ...int) int {
	var ai []int
	for _, v := range vi {
		if v%2 == 0 {
			ai = append(ai, v)
		}

	}
	return f(ai...)
}
