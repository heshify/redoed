package repository

import (
	"github.com/heshify/redoed/internal/db"
	"github.com/heshify/redoed/internal/models"
)

type DocumentRepository struct{}

func NewDocumentRepository() *DocumentRepository {
	return &DocumentRepository{}
}

func (r *DocumentRepository) Create(doc *models.Document) error {
	result := db.DB.Create(doc)
	return result.Error
}

func (r *DocumentRepository) GetAll() ([]models.Document, error) {
	var documents []models.Document
	result := db.DB.Find(&documents)
	return documents, result.Error
}
