package main

import "fmt"

func main() {
	salarios := map[string]int{"Teste1": 1000, "Teste2": 2000, "Teste3": 3000}
	//fmt.Println(salarios["Teste1"])
	//delete(salarios, "Teste1")
	//salarios["Teste4"] = 4000

	//sal := make(map[string]int)
	//sal1 := map[string]int{}
	//sal1["Teste5"] = 1000

	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salário é %d\n", salario)
	}
}
