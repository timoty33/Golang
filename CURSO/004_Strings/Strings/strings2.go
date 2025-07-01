// ALgumas string functions -->

package Strings

import (
	"fmt"
	"strings"
)

func String2() {
	s := "a grande arvore apodrecida"

	x := len(s)

	a := strings.Contains(s, "grande arvore")      // retorna true (bool)
	b := strings.Contains(s, "do pântano")         // retorna false (bool)
	c := strings.HasPrefix(s, "a grande")          // retorna true (bool)
	d := strings.HasSuffix(s, "arvore apodrecida") // retorna true (bool)

	S := strings.ToUpper(s)
	s = strings.ToLower(S)
	Ss := strings.ToTitle(s) // São quase iguais o ToUpper e o ToTitle, porém o ToTitle transcreve caracteres especiais para sua forma maiúscula.

	fmt.Printf("a -- %v \nb -- %v\nc -- %v\nd -- %v\n", a, b, c, d)
	fmt.Printf("Upper -- %v\nLower -- %v\nTitle -- %v\n", S, s, Ss)
	fmt.Printf("%v - %v", s, x)
}
