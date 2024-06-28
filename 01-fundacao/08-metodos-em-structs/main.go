package main

import "fmt"

type Cliente struct {
	Nome  string
	Ativo bool
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado.\n", c.Nome)
}

func (c Cliente) Ativar() {
	c.Ativo = true
	fmt.Printf("O cliente %s foi ativado.\n", c.Nome)
}

func main() {
	victor := Cliente{
		Nome:  "Victor",
		Ativo: true,
	}

	victor.Desativar()
	victor.Ativar()
}
