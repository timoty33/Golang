package crud

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Struct global, visível em todo o pacote
type Model struct {
	ID       int
	Nome     string
	Consulta string
	Data     time.Time
	Marcada  time.Time
}

func Crud(database string) *sql.DB {
	db, err := sql.Open("sqlite3", database)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Banco conectado")

	query := `CREATE TABLE IF NOT EXISTS consultas (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		Nome TEXT,
		Consulta TEXT,
		Data DATETIME,
		Marcada DATETIME
	)`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Tabela criada ou já existia!!")
	}

	return db
}

// Agora a função está fora da função Crud
func Read(db *sql.DB, database string) {
	query := fmt.Sprintf("SELECT * FROM %s", database)

	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Erro ao ler a tabela: %s\n", err)
		return
	}
	defer rows.Close()

	var consultas []Model

	for rows.Next() {
		var consulta Model
		err = rows.Scan(&consulta.ID, &consulta.Nome, &consulta.Consulta, &consulta.Data, &consulta.Marcada)
		if err != nil {
			log.Fatalf("Erro ao ler as linhas: %s\n", err)
			return
		}
		consultas = append(consultas, consulta)
	}

	if err = rows.Err(); err != nil {
		log.Fatalf("Erro ao iterar as linhas: %s\n", err)
	}

	for _, consulta := range consultas {
		log.Printf("ID: %d, Nome: %s, Consulta: %s, Data: %s, Marcada: %v\n",
			consulta.ID, consulta.Nome, consulta.Consulta, consulta.Data, consulta.Marcada)
	}
}

func Create(db *sql.DB, database string, Nome string, Consulta string, Data time.Time, Marcada time.Time) {
	query := fmt.Sprintf(`INSERT INTO %s (Nome, Consulta, Data, Marcada) VALUES (?, ?, ?, ?)`, database)

	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatalf("Erro ao preparar a query: %s\n", err)
	}
	defer statement.Close()

	_, err = statement.Exec(Nome, Consulta, Data, Marcada)
	if err != nil {
		log.Fatalf("Erro ao executar: %s\n", err)
	}

	log.Println("Consulta marcada com sucesso")
}
