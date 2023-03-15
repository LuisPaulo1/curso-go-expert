package main

func main() {

	a := 10
	var ponteiro *int = &a // recebe o endereço de memória da variável a
	var b = &a             // recebe o endereço de memória da variável a

	var c int = 30
	var d int

	println(a)  // valor de a
	println(&a) // endereço de memória da varável a

	println(ponteiro)  // endereço de memória da varável a
	println(*ponteiro) // valor que está no endereço da variável a

	println()
	println(b)  // endereço de memória da varável a
	println(*b) // retorna valor que está no endereço da variável a
	println()

	*b = 2 // altera o valor que está no endereço da variável a

	println()
	println(b)
	println(*b)
	println()

	println(c)
	println(d)
}
