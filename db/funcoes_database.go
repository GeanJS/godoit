package db

import (
	"database/sql"
	"godoit/models"
)

func SalvaTarefa(db *sql.DB, tarefa models.Tarefa) error {
	_, err := db.Exec(`INSERT INTO tarefas
		(descricao, status, criada_em, finalizada_em) 
		VALUES (?, ?, ?, ?)`,
		tarefa.Descricao, tarefa.Status, tarefa.CriadaEm, tarefa.FinalizadaEm)
	return err
}

