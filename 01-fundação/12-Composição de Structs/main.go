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
	//endereco Endereco // criando variável do tipo Endereco
	Endereco // utilizando Composição
}

func main() {
	cliente := Cliente{
		Nome:  "Paulo",
		Idade: 35,
		Ativo: true,
	}
	cliente.Ativo = false

	//cliente.endereco.Estado = "SP"

	cliente.Cidade = "São Paulo"
	cliente.Endereco.Cidade = "Rio de Janeiro"

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", cliente.Nome, cliente.Idade, cliente.Ativo)
}
