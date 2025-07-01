package Arrays_Slices

import "fmt"

func Arrays1() {
	// Criamos um array usando [quantidade]tipo{elementos}
	arr := [3]int{1, 2, 3}

	fmt.Println(arr) // [1 2 3] <-- arrays são printados com []

	// Se um array não estiver totalmente preenchido, não acontece nada
	nomes := [3]string{"Timóteo", "Sofia"}

	fmt.Println(nomes)      // [Timóteo Sofia ] <-- Espaço vazio no final
	fmt.Println(len(nomes)) // 3

	// Se quisermos preencher
	nomes[2] = "Helena"

	fmt.Println(nomes)      // [Timóteo Sofia Helena]
	fmt.Println(len(nomes)) // 3
}
