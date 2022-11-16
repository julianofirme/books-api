package main

import (
	"github.com/jfirme-sys/books-api/database"
	"github.com/jfirme-sys/books-api/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}
