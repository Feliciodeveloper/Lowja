package main

import (
	"net/http"

	"LOJA-CURSO/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
