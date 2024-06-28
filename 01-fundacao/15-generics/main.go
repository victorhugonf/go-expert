package main

import "fmt"

type MyNumber int

type Number interface {
	~int | float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a, b T) bool {
	return a == b
}

func main() {
	m := map[string]int{"Victor": 1000, "Hugo": 500}
	fmt.Println(Soma(m))

	m2 := map[string]float64{"Victor": 123.4, "Hugo": 789.5}
	fmt.Printf("%.2f\n", Soma(m2))

	m3 := map[string]MyNumber{"Victor": 1000, "Hugo": 500}
	fmt.Println(Soma(m3))

	fmt.Println(Compara(10, 10))
}
