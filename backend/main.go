package main

import (
	"log"

	_ "github.com/pshebel/partiburo/backend/database"
	"github.com/pshebel/partiburo/backend/server"
)

func main() {
	srv := server.GetServer()
	err := srv.ListenAndServe()
	if err != nil {
		log.Println("Server failed to start:", err)
	}
}
