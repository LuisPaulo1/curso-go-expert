package main

import "fmt"

const a = "Hello, World!"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Teste"
	e float64 = 1.21289
	f ID      = 1
)

func main() {
	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	for i, v := range meuArray {
		fmt.Printf("O valor do índice %d é o valor %d\n", i, v)
	}

	println(len(meuArray) - 1) //len retorna o tamanho do array
	println(meuArray[len(meuArray)-1])

	fmt.Printf("O tipo de c é %T", c)
}
