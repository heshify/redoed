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
	var newDocument models.Document
	err := json.NewDecoder(r.Body).Decode(&newDocument)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Println("Error decoding request body: ", err)
		return
	}
	err = h.Repo.CreateDocument(&newDocument)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to create document", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(newDocument)
	if err != nil {
		http.Error(w, "Failed to encode documents to JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *DocumentHandler) DeleteDocument(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	err := h.Repo.DeleteDocument(id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to delete document", http.StatusInternalServerError)
		return
	}
}

func (h *DocumentHandler) UpdateDocument(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var newDocument models.Document
	err := json.NewDecoder(r.Body).Decode(&newDocument)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Println("Error decoding request body: ", err)
		return
	}
	h.Repo.UpdateDocument(id, &newDocument)
}

func (h *DocumentHandler) GetDocumentById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	document, err := h.Repo.GetDocument(id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to retireve document", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(document)
	if err != nil {
		http.Error(w, "Failed to encode documents to JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *DocumentHandler) GetAllDocuments(w http.ResponseWriter, r *http.Request) {
	documents, err := h.Repo.GetDocuments()
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
