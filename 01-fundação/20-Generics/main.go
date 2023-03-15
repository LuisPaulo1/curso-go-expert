package main

/*func somaInteiro(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}
	return soma
}*/

type MyNumber int

type Number interface {
	~int | ~float64
}

func Soma1[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func compara[T comparable](a T, b T) bool {
	return a == b
}

func main() {
	m := map[string]int{"Wesley": 1000, "João": 2000, "Maria": 3000}
	m2 := map[string]MyNumber{"Wesley": 1000, "João": 2000, "Maria": 3000}
	println(Soma1(m))
	println(Soma1(m2))
	//println(somaInteiro(m))
	println(compara(10, 11))
}
