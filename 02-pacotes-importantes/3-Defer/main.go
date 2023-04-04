package main

import "fmt"

func main() {
	fmt.Println("Primeira")
	defer fmt.Println("Segunda") // executada por último
	fmt.Println("Terceira")
}
