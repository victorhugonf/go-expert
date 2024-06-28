package main

type Endereco struct {
	Pais string
}

type Cliente struct {
	Nome string
	Endereco
	Address Endereco
}

func main() {
	victor := Cliente{
		Nome: "Victor",
		Endereco: Endereco{
			Pais: "Brasil",
		},
		Address: Endereco{
			Pais: "Brazil",
		},
	}

	println(victor.Pais)
	println(victor.Endereco.Pais)
	println(victor.Address.Pais)
}
