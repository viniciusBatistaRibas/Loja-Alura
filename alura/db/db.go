package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)



func Conecta() *sql.DB {
	conexao := "user=postgres  dbname=produtos password=example host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
