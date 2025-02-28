package handlers

import (
	"errors"
	"net/http"

	"github.com/heshify/redoed/internal/models"
	"github.com/heshify/redoed/internal/repository"
	"github.com/heshify/redoed/utils"
	"gorm.io/gorm"
)

type DocumentHandler struct {
	Repo *repository.DocumentRepository
}

func NewDocumentHandler(repo *repository.DocumentRepository) *DocumentHandler {
	return &DocumentHandler{Repo: repo}
}

func (h *DocumentHandler) CreateDocument(w http.ResponseWriter, r *http.Request) {
	var newDocument models.Document

	if err := utils.ParseJSON(r, &newDocument); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.ValidateDocument(newDocument); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.Repo.CreateDocument(&newDocument); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if err := utils.WriteJSON(w, http.StatusCreated, newDocument); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *DocumentHandler) DeleteDocument(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		utils.WriteError(w, http.StatusBadRequest, errors.New("missing document ID"))
		return
	}

	if err := h.Repo.DeleteDocument(id); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, errors.New("failed to delete document"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "document deleted"})
}

func (h *DocumentHandler) UpdateDocument(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		utils.WriteError(w, http.StatusBadRequest, errors.New("missing document ID"))
		return
	}

	_, err := h.Repo.GetDocument(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.WriteError(w, http.StatusNotFound, errors.New("document not found"))
		} else {
			utils.WriteError(w, http.StatusInternalServerError, errors.New("failed to fetch document"))
		}
		return
	}

	var newDocument models.Document

	if err := utils.ParseJSON(r, &newDocument); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.ValidateDocument(newDocument); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.Repo.UpdateDocument(id, &newDocument); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, errors.New("failed to update document"))
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, newDocument); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *DocumentHandler) GetDocumentById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		utils.WriteError(w, http.StatusBadRequest, errors.New("missing document ID"))
		return
	}

	document, err := h.Repo.GetDocument(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.WriteError(w, http.StatusNotFound, errors.New("document not found"))
		} else {
			utils.WriteError(w, http.StatusInternalServerError, errors.New("failed to fetch document"))
		}
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, document); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}

func (h *DocumentHandler) GetAllDocuments(w http.ResponseWriter, r *http.Request) {
	documents, err := h.Repo.GetDocuments()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if err := utils.WriteJSON(w, http.StatusOK, documents); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}
