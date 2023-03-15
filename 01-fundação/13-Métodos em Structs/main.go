package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado", c.Nome)
}

func main() {
	cliente := Cliente{
		Nome:  "Paulo",
		Idade: 35,
		Ativo: true,
	}

	//cliente.Ativo = false
	cliente.Desativar()

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", cliente.Nome, cliente.Idade, cliente.Ativo)
}
