package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	victor := Cliente{Nome: "Victor", Idade: 37, Ativo: true}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t \n", victor.Nome, victor.Idade, victor.Ativo)

	victor.Ativo=false

	fmt.Println(victor.Ativo)
}
