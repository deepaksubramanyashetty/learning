// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 3, 7, 2, 8, 4, 9}
	xs := []string{"james", "A", "bond", "miss", "moneypenny"}
	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println("___________________")
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}
