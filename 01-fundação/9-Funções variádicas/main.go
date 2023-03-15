package main

func main() {
	println(sum(1, 2, 3, 4, 5, 6))
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
