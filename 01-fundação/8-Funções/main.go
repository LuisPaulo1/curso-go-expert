package main

import (
	"errors"
	"fmt"
)

func main() {

	valor, err := sum4(50, 10)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(valor)

	//sub(1, 2)
}

func sum1(a int, b int) int {
	return a + b
}

func sum2(a, b int) int {
	return a + b
}

func sum3(a, b int) (int, bool) {
	if a+b >= 50 {
		return a + b, true
	}
	return a + b, false
}

func sum4(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("A soma é maior que 50")
	}
	return a + b, nil
}

func sub(a int, b int) {
	println(a - b)
}
