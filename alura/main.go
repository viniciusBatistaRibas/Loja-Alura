package main

import (
	"net/http"

	"git.hub/vinicius/alura/routes"
	_ "go.mongodb.org/mongo-driver/mongo" //mudar a biblioteca
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
