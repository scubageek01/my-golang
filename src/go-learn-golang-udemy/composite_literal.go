package main

import "fmt"

func main() {
	// COMPOSITE LITERAL
	x := []int{10, 22, 43, 74, 5}
	y := []string{"Sean", "Mai", "Hannah", "Sydney", "Benji", "Baby"}
	fmt.Println(x[3:])
	for i, v := range y {
		fmt.Println(i, v)
	}

	for i := 1; i <= len(x); i++ {
		fmt.Println(i)
	}
	m := map[string]int{
		"Sean": 50,
		"Mai":  48,
	}
	fmt.Println(m)
	fmt.Println(m["Sean"])
}
