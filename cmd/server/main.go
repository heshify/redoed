package main

import (
	"log"
	"net/http"

	"github.com/heshify/redoed/internal/db"
	"github.com/heshify/redoed/internal/router"
)

func init() {
	db.InitDb()
}

func main() {
	r := router.NewRouter()
	port := "8080"

	log.Printf("Starting server on port %s", port)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal(err)
	}
}
