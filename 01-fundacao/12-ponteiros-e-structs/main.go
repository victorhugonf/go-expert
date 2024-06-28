package main

import "fmt"

type Conta struct {
	Saldo int
}

func NewConta() *Conta {
	return &Conta{Saldo: 0}
}

func (c Conta) simularEmprestimo(valor int) int {
	c.Saldo += valor
	return c.Saldo
}

func (c *Conta) depositar(valor int) int {
	c.Saldo += valor
	return c.Saldo
}

func main() {
	fmt.Println("conta")
	conta := Conta{
		Saldo: 100,
	}
	fmt.Printf("Simulação %d\n", conta.simularEmprestimo(200))
	fmt.Printf("Saldo %d\n", conta.Saldo)

	fmt.Printf("Depósito %d\n", conta.depositar(200))
	fmt.Printf("Saldo %d\n", conta.Saldo)

	fmt.Println("\noutra conta")
	outraConta := NewConta()
	outraConta.depositar(50)
	fmt.Printf("Simulação %d\n", outraConta.simularEmprestimo(200))
	fmt.Printf("Saldo %d\n", outraConta.Saldo)

	fmt.Printf("Depósito %d\n", outraConta.depositar(200))
	fmt.Printf("Saldo %d\n", outraConta.Saldo)
}
