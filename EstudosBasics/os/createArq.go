package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Criando arquivo de texto!!")

	arq, err := os.Create("HelloWorld.txt")

	defer arq.Close()

	if err != nil {
		panic(err)
	}

	_, err = arq.WriteString("Hello World\n")

	if err != nil {
		panic(err)
	}

	arq.Sync()

	fmt.Println("Arquivo criado com sucesso!")
}
