package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var numEsp int
	var numNor int

	for i := 0; i < 10; i++ {

		num := rand.Intn(10) + 1

		if num <= 5 {
			fmt.Printf("\033[1;32mO seu número é especial: %v\n\033[0;0;0m", num)
			numEsp++
		} else {
			fmt.Printf("\033[0;31mO seu número é totalmente normal: %v\n\033[0;0;0m", num)
			numNor++
		}

	}

	fmt.Println("Números especiais:", numEsp)
	fmt.Println("Números normais:", numNor)

}
