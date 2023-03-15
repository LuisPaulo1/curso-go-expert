package main

import (
	"curso-go-expert/matematica"
	"fmt"
	"github.com/google/uuid"
)

func main() {
	soma := matematica.Soma(10, 20)
	carro := matematica.Carro{Marca: "Fiat"}

	fmt.Println(matematica.A)
	fmt.Println(carro.Marca)
	fmt.Println(carro.Andar())
	fmt.Printf("Resultado: %v\n", soma)
	fmt.Println(uuid.New())

}
