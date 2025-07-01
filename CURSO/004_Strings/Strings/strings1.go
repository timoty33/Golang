// Strings são utf-8 --> uint8, ou seja, bytes

package Strings

import "fmt"

func String1() {
	s := "a grande árvore apodrecida" // Mesmo que a frase tenha 26 caracteres, em GO, o tamanho é definido pelo tamanho na memória, que é em bytes (é --> 2 bytes), por isso a frase mede 27

	a := len(s)                  // Conta quanto s mede
	b := s[:16]                  // b representa s indo até o espaço 16
	c := s[9:]                   // c representa s começando pelo espaço 9 e indo até o final
	d := s[:9] + "alma" + s[16:] // Substitui a árvore por 'alma'

	fmt.Printf("%v - %v \n%v\n%v\n%v\n", s, a, b, c, d)

	s = s + " do pântano"

	fmt.Printf("%v - %v", s, len(s))
}
