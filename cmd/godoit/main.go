package main

import (
	"fmt"
	"godoit/db"
	"godoit/models"
	"os"

	"github.com/spf13/cobra"
)


func main() {
	rootCmd := &cobra.Command {
		Use: "godoit",
		Short: "Ferramenta CLI para gerenciar tarefas",
		Long: "Ferramenta CLI escrita em GO para gerenciar tarefas do dia a dia",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Bem-vindo ao Godoit")
		}, 
	}
	cmdAdd := &cobra.Command{
		Use: "add",
		Short: "Adiciona uma tarefa",
		Long: "Adiciona uma tarefa no banco de dados",
		Run: func(cmd *cobra.Command, args []string) {
			conn := db.Conecta()
			defer conn.Close()
			descricao := args[0]
			t := models.Tarefa{}
			t.CriaTarefa(descricao)
			db.SalvaTarefa(conn, t)
			
		},
	}
	cmdList := &cobra.Command{
		Use: "list",
		Short: "Lista as tarefas",
		Long: "Lista todas as tarefas registradas",
		Run: func(cmd *cobra.Command, args []string) {
			conn := db.Conecta()
			defer conn.Close()
			tarefas, err := db.ListaTodasTarefas(conn)
			if err != nil {
				fmt.Println(err)
			}

			for _,t := range tarefas {
				statusStr := "✗"
				if t.Status == true {
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
		},
	}
	rootCmd.AddCommand(cmdAdd)
	rootCmd.AddCommand(cmdList)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
