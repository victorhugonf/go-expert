package matematica

import "fmt"

func Soma[T int | float64](a, b T) T {
	return a + b
}

type Carro struct {
	Marca   string
	andando bool //minúsculo não é exportado (não é público)
}

func NewCarro(marca string) *Carro {
	return &Carro{marca, false}
}

func (c *Carro) Andar() {
	c.andando = true
}

func (c *Carro) Parar() {
	c.andando = false
}

func (c Carro) Estado() {
	estado := "parado"
	if c.andando {
		estado = "andando"
	}
	fmt.Printf("O carro %v está %v\n", c.Marca, estado)
}
