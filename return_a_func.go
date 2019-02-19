package main

import "fmt"

func main() {
	fmt.Println(barReturnFunc()())
}

func barReturnFunc() func() int {
	return func() int {
		return 42
	}
}
