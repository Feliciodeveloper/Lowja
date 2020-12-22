package models

import (
	"database/sql"
	"regexp"
	"strings"

	"LOJA-CURSO/db"
	"LOJA-CURSO/services"
)

//Produto ...
type Produto struct {
	ID         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

//BuscaProdutos é responsavel por retornar a lista de produtos
func BuscaProdutos() []Produto {
	db := db.ConectaBanco()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM produtos ORDER BY 2 ;")
	services.CheckErro(err)
	defer rows.Close()
	var produtos []Produto
	var p Produto
	for rows.Next() {
		err := rows.Scan(&p.ID, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade)
		services.CheckErro(err)
		produtos = append(produtos, p)
	}
	return produtos
}

//BuscaProdutoEsp ...
func BuscaProdutoEsp(id string) Produto {
	db := db.ConectaBanco()
	defer db.Close()
	result, err := db.Query("SELECT * FROM produtos WHERE id = ?", id)
	services.CheckErro(err)
	p := Produto{}
	for result.Next() {
		err := result.Scan(&p.ID, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade)
		services.CheckErro(err)
	}
	return p
}

//BuscaProdutoNome ...
func BuscaProdutoNome(nome, pesquisa string) []Produto {
	var err error
	var result *sql.Rows
	db := db.ConectaBanco()
	defer db.Close()
	ast, _ := regexp.MatchString(`\*`, nome)
	if ast {
		nome = strings.ReplaceAll(nome, "*", "%")
	} else if len(nome) == 0 {
		nome = "%"
	} else {
		nome = nome + "%"
	}
	if pesquisa == "1" {
		result, err = db.Query("SELECT * FROM produtos WHERE nome like ?", nome)
	} else {
		result, err = db.Query("SELECT * FROM produtos WHERE descricao like ?", nome)
	}
	services.CheckErro(err)
	p := Produto{}
	produto := []Produto{}
	for result.Next() {
		err := result.Scan(&p.ID, &p.Nome, &p.Descricao, &p.Preco, &p.Quantidade)
		services.CheckErro(err)
		produto = append(produto, p)
	}
	return produto
}

//AdicionarProduto ...
func AdicionarProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaBanco()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO produtos (nome,descricao,preco,quantidade) VALUES (?,?,?,?)")
	services.CheckErro(err)
	stmt.Exec(nome, descricao, preco, quantidade)
}

//DeletaProduto ...
func DeletaProduto(p string) {
	db := db.ConectaBanco()
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM produtos WHERE id = ?")
	services.CheckErro(err)
	stmt.Exec(p)
}

//AtualizaProduto ...
func AtualizaProduto(p, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaBanco()
	defer db.Close()
	stmt, err := db.Prepare("UPDATE produtos SET nome = ?, descricao = ?, preco = ?, quantidade = ? WHERE id = ?")
	services.CheckErro(err)
	stmt.Exec(nome, descricao, preco, quantidade, p)
}
