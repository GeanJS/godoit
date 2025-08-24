// Package utils tem a funcao para formatar a tarefa
package utils

import (
	"fmt"
	"godoit/models"
)

func FormataTarefa(tarefas []models.Tarefa)  {
	for _,t := range tarefas {
		statusStr := "✗"
		if t.Status {
			statusStr = "✓"
		}

		criada := t.CriadaEm.Time.Format("2/01/2006 15:04")
		finalizada := ""
		if t.FinalizadaEm.Valid {
			finalizada = t.FinalizadaEm.Time.Format("2/01/2006 15:04")
		}

		fmt.Printf("[%s] %-3d | %-30s | Criada: %-16s | Finalizada: %-16s\n",
			statusStr, t.ID, t.Descricao, criada, finalizada)
		}

}
