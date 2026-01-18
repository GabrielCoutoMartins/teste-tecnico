package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ConectaComBancoDeDados() *sql.DB {
	// string de conexao com o banco de dados
	conexao := "user=postgres dbname=tasks password=gabriel123 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err)
	}
	return db
}
