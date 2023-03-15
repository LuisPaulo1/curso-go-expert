package main

import "fmt"

func soma(a, b *int) int {
	*a = 50
	*b = 30
	return *a + *b
}

func main() {
	a := 10
	b := 20

	println(soma(&a, &b))

	fmt.Printf("a: %d | b: %d", a, b)
}
