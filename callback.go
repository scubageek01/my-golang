package main

import "fmt"

func main() {
	ii := []int{1, 3, 6, 2, 8}
	s := sum(ii...)
	fmt.Println("All numbers total:", s)
	s2 := even(sum, ii...)
	fmt.Println("Even numbers total:", s2)
	s3 := odd(sum, ii...)
	fmt.Println("Odd numbers total:", s3)
	s4 := odd(product, ii...)
	fmt.Println("Odd numbers product:", s4)
	s5 := even(product, ii...)
	fmt.Println("Even numbers product:", s5)
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
func product(xi ...int) int {
	total := 1
	for _, v := range xi {
		total *= v
	}
	return total
}

func even(f func(xi ...int) int, li ...int) int {
	var mi []int
	for _, v := range li {
		if v%2 == 0 {
			mi = append(mi, v)
		}
	}
	return f(mi...)
}

func odd(f func(xi ...int) int, li ...int) int {
	var mi []int
	for _, v := range li {
		if v%2 != 0 {
			mi = append(mi, v)
		}
	}
	return f(mi...)
}
