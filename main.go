package main

import (
	"github.com/Bravoezz/first-fiber/db"
	"github.com/Bravoezz/first-fiber/server"
)

func main() {
	// db connection
	db.InitDbConnection()
	// db.Migrate()
	server.Start("4000")
}