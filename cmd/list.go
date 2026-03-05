// package cmd

// import (
// 	"taskmanager/db"

// 	"github.com/fatih/color"
// 	"github.com/spf13/cobra"
// )

// var listCmd = &cobra.Command{
// 	Use:   "list",
// 	Short: "List all tasks",

// 	Run: func(cmd *cobra.Command, args []string) {

// 		rows, err := db.DB.Query("SELECT id, title, completed FROM tasks")
// 		if err != nil {
// 			color.Red("Failed to fetch tasks")
// 			return
// 		}

// 		defer rows.Close()

// 		for rows.Next() {

// 			var id int
// 			var title string
// 			var completed bool

// 			err := rows.Scan(&id, &title, &completed)
// 			if err != nil {
// 				color.Red("Error reading task")
// 				continue
// 			}

// 			if completed {
// 				color.Green("%d | %s | DONE", id, title)
// 			} else {
// 				color.Yellow("%d | %s | PENDING", id, title)
// 			}
// 		}
// 	},
// }

// func init() {
// 	rootCmd.AddCommand(listCmd)
// }



package cmd

import (
	"taskmanager/db"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",

	Run: func(cmd *cobra.Command, args []string) {

		color.Cyan("ID | TASK | STATUS")
		color.White("----------------------")

		rows, _ := db.DB.Query("SELECT id, title, completed FROM tasks")

		for rows.Next() {

			var id int
			var title string
			var completed bool

			rows.Scan(&id, &title, &completed)

			if completed {
				color.Green("%d | %s | DONE", id, title)
			} else {
				color.Yellow("%d | %s | PENDING", id, title)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}