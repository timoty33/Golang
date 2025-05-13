package printfmt

import (
	f "fmt"
)

func PrintFormatado() {

	var nome string
	var estudante bool
	var idade int
	var binario int
	var tamanho float64

	nome = "Timóteo"
	estudante = true
	idade = 14
	binario = 0b1001
	tamanho = 1.70

	f.Printf("Nome >> %s, tipo - %T\n", nome, nome)                // %s -- format string
	f.Printf("Estudante >> %t, tipo - %T\n", estudante, estudante) // %t -- format bool
	f.Printf("Idade >> %d,, tipo - %T\n", idade, idade)            // %d -- format int (numero normal)
	f.Printf("Binário >> %b, tipo - %T\n", binario, binario)       // %b -- format binario
	f.Printf("Tamanho >> %.2f, tipo - %T\n", tamanho, tamanho)     // %f -- format float

}
