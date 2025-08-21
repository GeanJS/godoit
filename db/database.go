// Package db
package db

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

func BancoDeDados() *sql.DB {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tarefas(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		descricao TEXT NOT NULL,
		status INTEGER NOT NULL,
		criada_em DATETIME,
		finalizada_em DATETIME
		)`);
	if err != nil {
		log.Fatal(err)
	}
	return db
}
