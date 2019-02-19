package main

import "fmt"

func main() {
	a := 1984
	fmt.Println(a)
	fmt.Println(&a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a) // int != *int (they are 2 different types)
	//var b int = &a // Error because b of type int is assigned to a pointer of int
	//var b *int = &a
	b := &a
	fmt.Println(b)
	fmt.Println(*b) // *int asterisk to a type is a pointer;
					// *b asterisk is an operator to an address de-references pointer
	fmt.Println(*&a) // *&a de-referencing the pointer (&a) returns the value stored at the addr

	*b = 2019 // set the value to 2019 and store it at address (*b)
	fmt.Println(a)
}
