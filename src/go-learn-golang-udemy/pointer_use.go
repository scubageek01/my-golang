package main

import "fmt"

func main() {
	a := 0
	fmt.Println(a)
	fmt.Println(&a)
	fooPointer(&a)
	fmt.Println(a)
	fmt.Println(&a)

}

func fooPointer(n *int) {
	*n = 42
}
