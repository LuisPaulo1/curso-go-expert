package main

import "fmt"

type Cliente struct {
	nome string
}

func (c *Cliente) andar() {
	c.nome = "Luis Paulo"
	fmt.Printf("O cliente %v andou\n", c.nome)
}

type Conta struct {
	saldo int
}

func NewConta() *Conta {
	return &Conta{saldo: 0} //cria uma struct vazia, similar ao construtor
}

func (c *Conta) simular(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
}

func main() {
	cliente := Cliente{nome: "Paulo"}
	//var cliente Cliente = Cliente{nome: "Paulo"}

	cliente.andar()
	fmt.Printf("Nome do cliente: %v\n", cliente.nome)

	conta := Conta{saldo: 100}
	conta.simular(200)
	println(conta.saldo)
}
