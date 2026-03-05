package cmd

import (
	"taskmanager/db"

	"github.com/spf13/cobra"
	"github.com/fatih/color"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new task",
	Run: func(cmd *cobra.Command, args []string) {

		title := args[0]

		query := "INSERT INTO tasks(title, completed) VALUES(?, ?)"

		_, err := db.DB.Exec(query, title, false)

		if err != nil {
			panic(err)
		}

		color.Green("Task Added Successfully")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}