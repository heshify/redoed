package main

import (
	"log"
	"net/http"

	"github.com/heshify/redoed/internal/db"
	"github.com/heshify/redoed/internal/router"
	"github.com/rs/cors"
)

func init() {
	db.InitDb()
}

func main() {
	r := router.NewRouter()
	port := "8080"
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Change to specific origins if needed
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	handlerWithCors := c.Handler(r)

	log.Printf("Starting server on port %s", port)
	err := http.ListenAndServe(":"+port, handlerWithCors)
	if err != nil {
		log.Fatal(err)
	}
}
