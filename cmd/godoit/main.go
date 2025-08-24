package main

import (
	"fmt"
	"godoit/db"
	"godoit/models"
	"godoit/utils"
	"os"
	"strconv"
	"strings"

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
			descricao := strings.Join(args, " ")
			t := models.Tarefa{}
			t.CriaTarefa(descricao)
			db.SalvaTarefa(conn, t)
			fmt.Println("Tarefa adicionada com sucesso")
			
		},
	}
	cmdList := &cobra.Command{
		Use: "list",
		Short: "Lista as tarefas",
		Long: "Lista todas as tarefas registradas",
	}
	cmdListAll := &cobra.Command{
		Use: "all",
		Short: "lista todas as tarefas",
		Long: "lista todas as tarefas independente do status atual delas",
		Run: func(cmd *cobra.Command, args []string) {
			conn := db.Conecta()
			defer conn.Close()
			tarefas, err := db.ListaTodasTarefas(conn)
			if err != nil {
				fmt.Println(err)
			}
			utils.FormataTarefa(tarefas)
		},
	}
	cmdListDone := &cobra.Command{
		Use: "done",
		Short: "lista apenas as tarefas finalizadas",
		Long: "lista apenas as tarefas que já estejam com o status de finalizadas",
		Run: func(cmd *cobra.Command, args []string) {
			conn := db.Conecta()
			defer conn.Close()
			tarefas, err := db.ListaTarefasFinalizadas(conn)
			if err != nil {
				fmt.Println(err)
			}
			utils.FormataTarefa(tarefas)
		},
	}
	cmdListUndone := &cobra.Command{
		Use: "undone",
		Short: "lista apenas tarefas não finalizas",
		Run: func(cmd *cobra.Command, args []string) {
			conn := db.Conecta()
			defer conn.Close()

			tarefas, err := db.ListaTarefasNaoFinalizadas(conn)
			if err != nil {
				fmt.Println(err)
			}
			utils.FormataTarefa(tarefas)
		},
	}
	cmdDel := &cobra.Command{
		Use: "del",
		Short: "deleta uma tarefa",
		Long: "deleta a tarefa referente ao id recebido",
		Run: func(cmd *cobra.Command, args []string) {
			conn := db.Conecta()
			defer conn.Close()

			strinIndex := args[0]
			id, err := strconv.Atoi(strinIndex)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Printf("realmente deseja deletar a tarefa %d? (y/n)", id)
			var input string
			fmt.Scanf("%s", &input)
			if err != nil {
				fmt.Println(err)
			}
			if input == "y" || input == "yes" {
				err = db.DeletaTarefa(conn, id)
				if err == nil {
					fmt.Println(err)
				}
			}
		},
	}
	cmdUpdate := &cobra.Command{
		Use: "mark",
		Short: "altera o status de uma tarefa",
		Long: "altera o status de uma tarefa, tanto para concluida ou não concluida",
	}
	cmdDone := &cobra.Command{
		Use: "done",
		Short: "finalizada uma tarefa",
		Long: "altera o status da tarefa para finalizado e marca o horario de finalização",
		Run: func(cmd *cobra.Command, args []string) {
			conn := db.Conecta()
			defer conn.Close()

			stringIndex := args[0]
			id, err := strconv.Atoi(stringIndex)
			if err != nil {
				fmt.Println(err)
			}

			t := models.Tarefa{}
			t.Finaliza()
			err = db.AlteraTarefa(conn, id, t)
			if err != nil {
				fmt.Println(err)
			}
		},
	}
	cmdUndone := &cobra.Command{
		Use: "undone",
		Short: "retira o status de finalizado",
		Long: "altera o status de uma tarefa finalizada, e retira a hora de finalização",
		Run: func(cmd *cobra.Command, args []string) {
			conn := db.Conecta()
			defer conn.Close()

			stringIndex := args[0]
			id, err := strconv.Atoi(stringIndex)
			if err != nil {
				fmt.Println(err)
			}
			t := models.Tarefa{}
			t.DesfazFinalizacao()
			err = db.AlteraTarefa(conn, id, t)
			if err != nil {
				fmt.Println(err)
			}
		},
	}

	rootCmd.AddCommand(cmdAdd)
	rootCmd.AddCommand(cmdList)
	rootCmd.AddCommand(cmdDel)
	rootCmd.AddCommand(cmdUpdate)
	cmdList.AddCommand(cmdListAll)
	cmdList.AddCommand(cmdListDone)
	cmdList.AddCommand(cmdListUndone)
	cmdUpdate.AddCommand(cmdDone)
	cmdUpdate.AddCommand(cmdUndone)


	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
