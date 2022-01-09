package main

import "fmt"

func main() {
	//foo empty ARGUMENT
	foo()
	//boo takes an ARGUMENT
	boo("james")
	st := woo("moneypenny")
	fmt.Println(st)

	x, y := mouse("Deepak", "shetty")
	fmt.Println(x)
	fmt.Println(y)
}

// fun (r reciever) identifier (parameters) (return(s)) { ... }
func foo() {
	fmt.Println("hello from foo")
}

func boo(s string) {
	fmt.Println("hello", s)
}

func woo(st string) string {
	return fmt.Sprint("hello from woo ", st)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, says "hello"`)
	b := true
	return a, b
}
