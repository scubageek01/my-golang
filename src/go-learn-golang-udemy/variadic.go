package main

import "fmt"

func main() {
	xi := []int{10, 2, 3, 4, 8, 102, 3, 4, 8, 10} // this is a slice of int
	res := funcAdd(xi...) // funcAdd is expecting variadic parameter of int.
	fmt.Println(res)		// by including ... after slice of int, Go will accept the arguments
}

func funcAdd(x ...int) int {
	sum := 0
	for _, a := range x {
		sum += a
	}
	return sum
}
