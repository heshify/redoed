package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", handler)
	err := http.ListenAndServe(
		":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
