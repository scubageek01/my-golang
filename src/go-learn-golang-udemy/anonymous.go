package main

import "fmt"

func main() {
	fooAnon()
	func () {
		fmt.Println("Anonymous ran")
	}()

	func (x int) {
		fmt.Println("Anonymous ran with", x)
	}(42)

}

func fooAnon() {
	fmt.Println("fooAnon ran")
}
