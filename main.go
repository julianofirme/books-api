package main

import "github.com/jfirme-sys/books-api/server"

func main() {
	server := server.NewServer()

	server.Run()
}