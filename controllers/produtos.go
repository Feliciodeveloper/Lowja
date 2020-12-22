package controllers

import (
	"net/http"
	"strconv"

	"LOJA-CURSO/models"
)

//Produto ...
func Produto(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscaProdutos()
	temp.ExecuteTemplate(w, "produto", produtos)
}

//Search ...
func Search(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("search")
		pesquisa := r.FormValue("pesquisa")
		produtos := models.BuscaProdutoNome(nome, pesquisa)
		temp.ExecuteTemplate(w, "produto", produtos)
	}
}

//New ...
func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "new", nil)
}

//Insert ...
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco, _ := strconv.ParseFloat(r.FormValue("preco"), 64)
		quantidade, _ := strconv.Atoi(r.FormValue("quantidade"))
		models.AdicionarProduto(nome, descricao, preco, quantidade)
	}
	http.Redirect(w, r, "/produto", 301)
}

//Delete ...
func Delete(w http.ResponseWriter, r *http.Request) {
	idproduto := r.URL.Query().Get("id")
	models.DeletaProduto(idproduto)
	http.Redirect(w, r, "/produto", 301)
}

//Edit ...
func Edit(w http.ResponseWriter, r *http.Request) {
	idproduto := r.URL.Query().Get("id")
	p := models.BuscaProdutoEsp(idproduto)
	temp.ExecuteTemplate(w, "Edit", p)
}

//Update ...
func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco, _ := strconv.ParseFloat(r.FormValue("preco"), 64)
		quantidade, _ := strconv.Atoi(r.FormValue("quantidade"))
		models.AtualizaProduto(id, nome, descricao, preco, quantidade)
	}
	http.Redirect(w, r, "/produto", 301)
}
