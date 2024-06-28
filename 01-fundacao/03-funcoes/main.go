package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := sum(100, 20)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(valor)
}

func sum(a, b int) (int, error) {
	sum := a + b
	if sum > 50 {
		return 0, errors.New("A soma é maior que 50")
	}
	return sum, nil
}
