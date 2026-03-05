package main

import (
	"taskmanager/cmd"
	"taskmanager/db"
)

func main() {

	db.InitDB()   // initialize database

	cmd.Execute()
}