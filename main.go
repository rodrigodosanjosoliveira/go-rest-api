package main

import (
	"fmt"

	"github.com/rodrigodosanjosoliveira/go-rest-api/database"
	"github.com/rodrigodosanjosoliveira/go-rest-api/routes"
)

func main() {
	database.ConectaComBancoDeDados()

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
