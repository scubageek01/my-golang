package main

import "fmt"

func main() {
	foo()
	bar("Sean")
	s1 := baz("Mai")
	fmt.Println(s1)
	x, y := multiple("Sean", "Huynh")
	fmt.Println(x)
	fmt.Println(y)

}

func foo() {
	fmt.Println("hello from foo")
}

func bar(s string) {
	fmt.Println("hello ", s)
}

func baz(s string) string {
	return fmt.Sprint("hello from baz, ", s)
}

func multiple(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, " ", `says "hello" ` )
	b := true
	return a, b

}
