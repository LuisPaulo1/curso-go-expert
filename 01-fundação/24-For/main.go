package main

func main() {

	for i := 0; i < 10; i++ {
		println(i)
	}

	numeros := []string{"um", "dois", "três"}
	for indice, valor := range numeros {
		println(indice, valor)
	}

	for _, valor := range numeros {
		println(valor)
	}

	for indice, _ := range numeros {
		println(indice)
	}

	i := 0
	for i < 10 {
		println(i)
		i++
	}

	for {
		println("Hello, World!")
	}
}
