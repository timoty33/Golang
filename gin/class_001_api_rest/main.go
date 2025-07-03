package main

import (
	"fmt"      // Importado para usar fmt.Printf e ver mensagens no console do servidor
	"net/http" // Importado para usar códigos de status HTTP como http.StatusOK

	"github.com/gin-gonic/gin"
)

type Pessoa struct {
	Nome      string `json:"nome"`      // Campo 'Nome' que será mapeado para "nome" no JSON
	Sobrenome string `json:"sobrenome"` // Campo 'Sobrenome' que será mapeado para "sobrenome" no JSON
}

var pessoas []Pessoa
var visits int

func main() {
	visits = 0

	server := gin.Default()

	// --- Rotas Existentes ---

	server.GET("/", func(ctx *gin.Context) {
		visits++
		fmt.Printf("Visita na rota /: Total de visitas = %d\n", visits) // Para depuração
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Você entrou no meu site!",
		})
	})

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
			"visitas": visits,
		})
	})

	// --- Nova Rota POST para Salvar Pessoa (usando Parâmetros de Rota) ---

	server.POST("/salvar/:nome/:sobrenome", func(ctx *gin.Context) {
		// Recebendo parâmetros de rota
		nome := ctx.Param("nome")
		sobrenome := ctx.Param("sobrenome")

		// Criando uma nova instância de Pessoa
		novaPessoa := Pessoa{
			Nome:      nome,
			Sobrenome: sobrenome,
		}

		pessoas = append(pessoas, novaPessoa)

		fmt.Printf("Nova pessoa salva: %+v\n", novaPessoa)
		fmt.Printf("Total de pessoas salvas: %d\n", len(pessoas))
		fmt.Printf("Todas as pessoas: %+v\n", pessoas)

		// Enviando uma resposta ao cliente (IMPORTANTE!)
		ctx.JSON(http.StatusCreated, gin.H{ // Usamos http.StatusCreated (201) para indicar que um recurso foi criado
			"message": "Pessoa salva com sucesso!",
			"data":    novaPessoa,
		})
	})

	// --- Nova Rota GET para Listar Pessoas (para verificar se o POST funcionou) ---
	server.GET("/pessoas", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, pessoas) // Retorna o slice completo de pessoas
	})

	fmt.Println("Servidor Gin rodando na porta 8080...")
	server.Run(":8080")

	fmt.Println("Você pode acessar a 1° página por: http://localhost:8080/")
}
