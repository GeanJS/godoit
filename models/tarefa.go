// Package models tem a struct da nossa tarfa
package models

import (
	"database/sql"
	"time"
)



type Tarefa struct {
	ID int
	Descricao string
	Status bool
	CriadaEm sql.NullTime
	FinalizadaEm sql.NullTime
}

func (t *Tarefa) CriaTarefa(descricao string) {
	t.Descricao = descricao
	t.Status = false
	t.CriadaEm = sql.NullTime{Time: time.Now(), Valid: true}
	t.FinalizadaEm = sql.NullTime{Valid: false}
}

func (t * Tarefa) Finaliza() {
	t.Status = true
	t.FinalizadaEm = sql.NullTime{
		Time: time.Now(),
		Valid: true,
	}
}
