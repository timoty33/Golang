package main

import (
	"agendamento/crud"
	"log"
	"time"
)

func main() {

	db := crud.Crud("data.db")
	table := "consultas"
	defer db.Close()

	crud.Create(db, table, "Timóteo", "Ortopedista", time.Now(), time.Now())
	crud.Read(db, table)

	log.Println("Fim!")
}
