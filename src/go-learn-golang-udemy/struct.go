package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}
type secretAgent struct {
	person
	ltk bool
}

func main() {
	p1 := person{
		first: "Sean",
		last:  "Huynh",
		age:   50,
	}
	p2 := person{
		first: "Mai",
		last:  "Huynh",
		age:   48,
	}
	p3 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   27,
		},
		ltk: true,
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3.first, p3.last, p3.age, p3.ltk)

}
