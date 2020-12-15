package db

import (
	"database/sql"
	"fmt"

	"github.com/LOJA-CURSO/services"

	//Driver de conexão do banco mysql
	_ "github.com/go-sql-driver/mysql"
)

//ConectaBanco é função de conexao com o banco de dados
func ConectaBanco() *sql.DB {
	const (
		host     = "localhost"
		database = "cursogo"
		user     = "root"
		password = ""
		port     = "3305"
	)
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)
	db, err := sql.Open("mysql", connectionString)
	services.CheckErro(err)
	return db
}
