package main

import "fmt"

type person1 struct {
	first string
	last  string
}
type superHero struct {
	person1
	levitate bool
}

func (s superHero) speak() {
	fmt.Println(s.first, s.last,  `says "Rah rah"`)
}

func main() {
	sh1 := superHero{
		person1: person1{
			first: "Clark",
			last:  "Kent",
		},
		levitate: true,
	}
	sh2 := superHero{
		person1: person1{
			first: "Peter",
			last:  "Parker",
		},
		levitate: true,
	}
	fmt.Println(sh1.first, `can actually levitate!  It's`, sh1.levitate )
	sh1.speak()
	sh2.speak()
}
