package main

import "fmt"

func main() {
	//for simples
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//for em range (foreach)
	numeros := []string{"um", "dois", "tres"}
	for k, v := range numeros {
		fmt.Printf("key: %v - value: %v\n", k, v)
	}

	numeros = []string{"um", "dois", "tres"}
	for k := range numeros {
		fmt.Printf("key: %v \n", k)
	}

	numeros = []string{"um", "dois", "tres"}
	for _, v := range numeros {
		fmt.Printf("value: %v\n", v)
	}

	//for como while
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	//for como loop infinito
	for {
		fmt.Print("=")
	}
}
