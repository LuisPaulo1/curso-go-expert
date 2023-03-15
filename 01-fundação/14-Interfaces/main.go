package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Pessoa interface {
	Desativar()
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {

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

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	cliente := Cliente{
		Nome:  "Paulo",
		Idade: 35,
		Ativo: true,
	}

	//minhaEmpresa := Empresa{}

	//Desativacao(minhaEmpresa)
	//Desativacao(&cliente) //colocar *Cliente no método Desativar()
	Desativacao(cliente)
	//cliente.Ativo = false
	//cliente.Desativar()

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", cliente.Nome, cliente.Idade, cliente.Ativo)
}
