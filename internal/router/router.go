package router

import (
	"net/http"

	"github.com/heshify/redoed/internal/handlers"
	"github.com/heshify/redoed/internal/repository"
)

func NewRouter() *http.ServeMux {

	docRepo := repository.NewDocumentRepository()
	docHandler := handlers.NewDocumentHandler(docRepo)

	userRepo := repository.NewUserRepository()
	userHandler := handlers.NewUserHandler(userRepo)
	r := http.NewServeMux()

	//auth routes
	r.HandleFunc("POST /api/auth/login", userHandler.Login)
	r.HandleFunc("POST /api/auth/register", userHandler.RegisterUser)
	// user routes
	r.HandleFunc("GET /api/user/{id}", userHandler.GetUser)
	r.HandleFunc("GET /api/user", userHandler.GetUsers)

	//document routes
	r.HandleFunc("POST /api/document", docHandler.CreateDocument)
	r.HandleFunc("GET /api/document", docHandler.GetDocuments)
	r.HandleFunc("GET /api/document/{id}", docHandler.GetDocument)
	r.HandleFunc("PUT /api/document/{id}", docHandler.UpdateDocument)
	r.HandleFunc("DELETE /api/document/{id}", docHandler.DeleteDocument)

	return r
}
