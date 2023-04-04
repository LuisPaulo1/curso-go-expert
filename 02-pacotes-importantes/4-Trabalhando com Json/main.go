package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"numero"` // `json:"-"` tags
	Saldo  int `json:"saldo"`
}

func main() {
	// Struct para Json - serialização
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}
	println(string(res))

	// Struct para Json - serialização
	err = json.NewEncoder(os.Stdout).Encode(conta) // imprime no console
	if err != nil {
		println(err)
	}

	// Json para Struct - desserialização
	jsonPuro := []byte(`{"numero":2-Template.Must,"saldo":200}`)
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		println(err)
	}
	println(contaX.Saldo)

}
