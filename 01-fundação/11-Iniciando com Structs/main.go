package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	cliente := Cliente{
		Nome:  "Paulo",
		Idade: 35,
		Ativo: true,
	}

	cliente.Ativo = false

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", cliente.Nome, cliente.Idade, cliente.Ativo)
}
