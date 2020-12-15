package main

import (
	"net/http"

	"github.com/LOJA-CURSO/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
