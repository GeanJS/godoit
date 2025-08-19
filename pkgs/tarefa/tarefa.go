// Package tarefa
package tarefa

import (
	"errors"
	"time"
)

const err = "tarefa não encontrada"

type Tarefa struct{
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

func (t *Tarefa) ListaTarefa(tarefas []Tarefa, tarefa Tarefa) (Tarefa, error) {
	for _,t := range tarefas {
		if t == tarefa {
			return t, nil
		}
	}
	return tarefa, errors.New(err)
}
