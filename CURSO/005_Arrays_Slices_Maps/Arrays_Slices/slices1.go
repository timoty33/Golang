package Arrays_Slices

import (
	"fmt"
)

func Slices1() {
	frutas := [3]string{"banana", "maçã", "laranja"}

	// Cria o slice começando com "macarronada"
	comidas := []string{"macarronada"}

	// Cria o slice começando com "macarronada"
	favoritos := []string{"macarronada"}

	// Adiciona todas as frutas depois
	comidas = append(comidas, frutas[:]...)

	fmt.Println(comidas)
	fmt.Println("Tamanho:", len(comidas))
	fmt.Println("Capacidade:", cap(comidas))

	// Comparando slices
	fmt.Println(comidas[0] == favoritos[0]) // true
}
