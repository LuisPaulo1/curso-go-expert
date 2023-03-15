package main

func main() {
	total := func() int {
		return sum(1, 2, 3)
	}()
	println(total)
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
