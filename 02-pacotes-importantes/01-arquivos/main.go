package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	//escrita
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	tamanho, err := f.WriteString("hello, world!hello, world!hello, world!hello, world!")
	//tamanho, err := f.Write([]byte("escrevendo bytes no arquivo!"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Sucesso! Tamanho: %d bytes \n", tamanho)
	f.Close()

	//leitura
	//arquivo,err := os.Open("arquivo.txt")
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	//leitura de pouco em pouco abrindo o arquivo
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	time.Sleep(5000 * time.Millisecond)

	//remoção
	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
