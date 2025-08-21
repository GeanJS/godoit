// Package models tem a struct da nossa tarfa
package models


import "time"

type Tarefa struct {
	Descricao string
	Status bool
	CriadaEm time.Time
	FinalizadaEm time.Time
}

func (t *Tarefa) CriaTarefa(descricao string) {
	t.Descricao = descricao
	t.Status = false
	t.CriadaEm = time.Now()
	t.FinalizadaEm = time.Time{}
}


