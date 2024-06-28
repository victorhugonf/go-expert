package main

import "fmt"

func main() {
	a := 10
	fmt.Printf("Endereço de memória: %p Valor: %d\n", &a, a)

	var ponteiro *int = &a
	fmt.Println(ponteiro)

	*ponteiro = 20
	fmt.Println(a)

	b := &a
	fmt.Printf("Endereço de memória: %p Valor: %d\n", b, *b)
}
