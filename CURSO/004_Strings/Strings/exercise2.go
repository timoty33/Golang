package Strings

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Exercise2() {

	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		s := scan.Text()

		palavras := len(strings.Fields(s))
		caracteres := len(strings.ReplaceAll(s, " ", ""))
		var quantidadeA int

		for i := 0; i < len(s); i++ {
			caractere := string(s[i])
			if strings.ToLower(caractere) == "a" {
				quantidadeA++
			}
		}

		fmt.Printf("Frase original: %v\nNúmero de palavras: %v\nNúmero de caracteres (sem contar os espaços): %v\nQuantidade de Aa: %v\n", s, palavras, caracteres, quantidadeA)
	}
}
