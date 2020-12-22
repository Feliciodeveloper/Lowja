package controllers

import (
	"net/http"

	"LOJA-CURSO/models"
)

//Login ...
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		email := r.FormValue("email")
		senha := r.FormValue("senha")
		_, valida := models.ValidaLogin(email, senha)
		if valida {
			http.Redirect(w, r, "produto", 301)
		} else {
			temp.ExecuteTemplate(w, "index", true)
		}
	}
}
