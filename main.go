package main

import (
	"fmt"

	"github.com/Kaiki2814/go-rest-api/routes"
)

func main() {
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleResquest()
}
