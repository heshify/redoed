package router

import (
	"net/http"

	"github.com/heshify/redoed/internal/handlers"
	"github.com/heshify/redoed/internal/repository"
)

func NewRouter() *http.ServeMux {

	docRepo := repository.NewDocumentRepository()
	docHandler := handlers.NewDocumentHandler(docRepo)

	r := http.NewServeMux()
	r.HandleFunc("GET /document", docHandler.GetAllDocuments)
	r.HandleFunc("POST /document", docHandler.CreateDocument)
	return r
}
