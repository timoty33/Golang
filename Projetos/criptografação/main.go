package main

import (
	"fmt"
	"strconv"
	"strings"
)

func crip(message string) string {
	letrasCriptografadas := []byte(message)

	var result string

	for _, b := range letrasCriptografadas {
		result += fmt.Sprintf("%d ", b)
	}

	return strings.TrimSpace(result)
}

func descrip(message string) string {
	var result string
	var bytes []byte

	partes := strings.Split(message, " ")

	for _, parte := range partes {

		num, err := strconv.Atoi(parte)
		if err != nil {
			fmt.Printf("Erro ao converter:\t%v \n", parte)
			continue
		}

		bytes = append(bytes, byte(num))
	}

	result = string(bytes)

	return result
}

func main() {
	message := "Aldinei"

	cripMessage := crip(message)

	fmt.Println(cripMessage)

	descrip_ := descrip(cripMessage)

	fmt.Println(descrip_)
}
