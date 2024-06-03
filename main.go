package main

import (
	"GoPay/models"
	"GoPay/routes"
	"log"
	"net/http"
)

func main() {
	models.InitDatabase()

	r := routes.SetupRouter()

	log.Println("Servidor rodando na porta 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
