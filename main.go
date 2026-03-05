package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: add/list/done/delete")
		return
	}

	command := os.Args[1]

	tasks, _ := LoadTasks()

	switch command {

	case "add":

		title := os.Args[2]

		task := Task{
			ID:        len(tasks) + 1,
			Title:     title,
			Completed: false,
		}

		tasks = append(tasks, task)

		SaveTasks(tasks)

		LogAsync("Task Added")

	case "list":

		for _, t := range tasks {
			fmt.Printf("%d | %s | %v\n", t.ID, t.Title, t.Completed)
		}

	case "done":

		id, _ := strconv.Atoi(os.Args[2])

		for i := range tasks {
			if tasks[i].ID == id {
				tasks[i].Completed = true
			}
		}

		SaveTasks(tasks)

	case "delete":

		id, _ := strconv.Atoi(os.Args[2])

		var newTasks []Task

		for _, t := range tasks {
			if t.ID != id {
				newTasks = append(newTasks, t)
			}
		}

		SaveTasks(newTasks)
	}
}