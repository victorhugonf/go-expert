package main

import (
	"curso-go/matematica"
	"fmt"

	"github.com/google/uuid"
)

func main() {
	s := matematica.Soma(10, 20)

	fmt.Println("Resultado: ", s)

	carro := matematica.NewCarro("Renault")
	carro.Estado()
	carro.Andar()
	carro.Estado()
	carro.Parar()
	carro.Estado()
	//carro.andando minúsculo não é exportado (não é público)

	fmt.Println(uuid.New())
}
