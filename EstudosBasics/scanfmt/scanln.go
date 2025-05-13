package scanfmt

import (
	f "fmt"
	t "time"
)

func Agendamento() {

	var nome string
	var data string
	var consulta string

	check := 2

	for check != 1 {

		f.Print("Olá qual é seu nome: ")
		f.Scanln(&nome)

		f.Print("Para que dia é a consulta: ")
		f.Scanln(&data)

		f.Print("Para que é a consulta: ")
		f.Scanln(&consulta)

		msgFinal := f.Sprintf(`
Você, %s, tem uma consulta de %s marcada para o dia: %s. 
Certo? ( 1 - Certo, 2 - Errado)
>> `, nome, consulta, data) // Mensagem de check

		f.Print(msgFinal)
		f.Scanln(&check)
	}

	f.Println("\nObrigado!!")

	t.Sleep(2 * t.Second)

}
