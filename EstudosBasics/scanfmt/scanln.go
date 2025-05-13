package scanfmt

import (
	f "fmt"
	t "time"
)

func Agendamento() {

	var nome string
	var data string
	var consulta string

	f.Print("Olá qual é seu nome: ")
	f.Scanln(&nome)

	f.Print("Digite a data no formato (DD-MM-AAAA HH:MM): ")
	f.Scanln(&data)

	dataFormatada, _ := t.Parse("02-01-2006 15:04", data)

	f.Print("Para que é a consulta: ")
	f.Scanln(&consulta)

	f.Printf("Você, %s, tem uma consulta no dia %s, para: %s\n", nome, dataFormatada.Format("02-01-2006 15:04"), consulta)
}
