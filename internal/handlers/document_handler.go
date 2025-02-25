package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/heshify/redoed/internal/models"
	"github.com/heshify/redoed/internal/repository"
)

type DocumentHandler struct {
	Repo *repository.DocumentRepository
}

func NewDocumentHandler(repo *repository.DocumentRepository) *DocumentHandler {
	return &DocumentHandler{Repo: repo}
}

func (h *DocumentHandler) CreateDocument(w http.ResponseWriter, r *http.Request) {
	doc := models.Document{Title: "First Document", Content: "This is content of my first document"}
	err := h.Repo.Create(&doc)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to create document", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(doc)
	if err != nil {
		http.Error(w, "Failed to encode documents to JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *DocumentHandler) GetAllDocuments(w http.ResponseWriter, r *http.Request) {
	documents, err := h.Repo.GetAll()
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to retireve document", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(documents)
	if err != nil {
		http.Error(w, "Failed to encode documents to JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
