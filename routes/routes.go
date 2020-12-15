package routes

import (
	"net/http"

	"github.com/LOJA-CURSO/controllers"
)

//CarregaRotas ...
func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/produto", controllers.Produto)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/Edit", controllers.Edit)
	http.HandleFunc("/update", controllers.Update)
	http.HandleFunc("/search", controllers.Search)
	http.HandleFunc("/login", controllers.Login)
}
