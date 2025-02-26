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

	r.HandleFunc("POST /api/document", docHandler.CreateDocument)
	r.HandleFunc("GET /api/document", docHandler.GetAllDocuments)
	r.HandleFunc("GET /apidocument/{id}", docHandler.GetDocumentById)
	r.HandleFunc("PUT /api/document/{id}", docHandler.UpdateDocument)
	r.HandleFunc("DELETE /api/document/{id}", docHandler.DeleteDocument)
	return r
}
