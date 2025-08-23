package db

import (
	"database/sql"
	"godoit/models"
	"log"
)

func SalvaTarefa(db *sql.DB, tarefa models.Tarefa) error {
	_, err := db.Exec(`INSERT INTO tarefas
		(descricao, status, criada_em, finalizada_em) 
		VALUES (?, ?, ?, ?)`,
		tarefa.Descricao, tarefa.Status, tarefa.CriadaEm, tarefa.FinalizadaEm)
	return err
}

func ListaTodasTarefas(db *sql.DB) ([]models.Tarefa, error) {
	rows, err := db.Query("SELECT * FROM tarefas")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var tarefas []models.Tarefa
	for rows.Next() {
		var t models.Tarefa
		if err := rows.Scan(&t.ID, &t.Descricao, &t.Status, &t.CriadaEm, &t.FinalizadaEm); err != nil {
			log.Fatal(err)
		}
		tarefas = append(tarefas, t)

	}
	return tarefas, nil
}
