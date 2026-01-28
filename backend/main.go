package main

import (
	"log"

	_ "github.com/pshebel/partiburo/backend/database"
	"github.com/pshebel/partiburo/backend/server"
)

func main() {
	log.Println("version 0")
	srv := server.GetServer()
	err := srv.ListenAndServe()
	if err != nil {
		log.Println("Server failed to start:", err)
	}
}
