package main

import (
	"fmt"
	"godoit/pkgs/tarefa"
	"time"
)

var tarefas []tarefa.Tarefa

func main() {
	t := tarefa.Tarefa{
		Descricao: "cozinhar",
		Status: false,
		CriadaEm: time.Now(),
		FinalizadaEm: time.Now(),
	}
	tarefas = append(tarefas, t)
	fmt.Println(t.ListaTarefa(tarefas, t))	
}
