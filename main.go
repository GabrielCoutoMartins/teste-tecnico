package main

import (
	"net/http"
	"teste_tecnico/routes"
)

func main() {
	routes.CarregarRotas()
	http.ListenAndServe(":8080", nil)

}
