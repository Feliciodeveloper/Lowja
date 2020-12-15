package models

import (
	"github.com/LOJA-CURSO/db"
)

//Usuario ...
type Usuario struct {
	ID, nome, email, senha string
}

//ValidaLogin ...
func ValidaLogin(email, senha string) (Usuario, bool) {
	var u Usuario
	var qt string
	db := db.ConectaBanco()
	defer db.Close()
	db.QueryRow("SELECT count(id), id,nome,email,senha FROM usuarios WHERE email = ? AND senha = ?", email, senha).Scan(&qt, &u.ID, &u.nome, &u.email, &u.senha)
	if qt == "1" {
		return u, true
	}
	return u, false
}
