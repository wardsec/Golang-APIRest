package main

import (
	"api/rest/database"
	"api/rest/routes"
	"fmt"
)

func main() {
	database.ConectaBanco()
	fmt.Println("Iniciando servidor com go")
	routes.HandleResquest()
}
