package main

import (
	"fmt"
)

func main() {
	total := func(multiplicador int) int {
		return sum(1, 3, 5, 8, 9) * multiplicador
	}(3)

	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
