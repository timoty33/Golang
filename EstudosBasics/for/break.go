package main

import "fmt"

var idade int
var peso int

func main() {

	idade, peso = 14, 50

	for i := 0; i < peso; i++ {

		if i >= idade {
			fmt.Println("Erro!!")
			break
		} else {
			fmt.Println("Peso >>", i, " O peso ainda é menor que a idade.")
		}

	}

}
