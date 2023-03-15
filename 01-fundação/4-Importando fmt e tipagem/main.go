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
	fmt.Printf("O tipo de E é %T", e)
	fmt.Printf("O valor de E é %v", e)
	fmt.Printf("O valor de E é %.2v", e)
	fmt.Printf("O valor de E é %.3v", e)
	fmt.Printf("O valor de E é %.4v", e)
}
