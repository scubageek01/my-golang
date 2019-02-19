package main

import "fmt"

func main() {
	f := func() {
	fmt.Println("My func expression")
	}
	f()

	g := func(x int) {
		fmt.Println("Big brother is watching", x)
	}
	g(1984)
}
