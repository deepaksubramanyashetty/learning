package main

import "fmt"

func main() {
	y := sum(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(y)
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T", x)

	sum := 0

	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are not adding ", v, "to tthe total", sum)
	}

	fmt.Println("total", sum)
	return sum
}
