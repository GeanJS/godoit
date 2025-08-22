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
	rootCmd.AddCommand(cmdAdd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
