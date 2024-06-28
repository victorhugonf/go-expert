package main

import "fmt"

func main() {
	salarios := map[string]int{"Victor": 1000, "João": 2000, "Maria": 3000}
	fmt.Println(salarios["Victor"])

	delete(salarios, "Victor")
	fmt.Println(salarios["Victor"])

	salarios["vh"] = 5000
	fmt.Println(salarios["vh"])

	//sal := make(map[string]int)
	//sal1 := map[string]int{}
	//sal1["vh"] = 1000

	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é %d \n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salário é %d \n", salario)
	}
}
