package repository

import (
	"github.com/heshify/redoed/internal/db"
	"github.com/heshify/redoed/internal/models"
)

type DocumentRepository struct{}

func NewDocumentRepository() *DocumentRepository {
	return &DocumentRepository{}
}

func (r *DocumentRepository) CreateDocument(doc *models.Document) error {
	result := db.DB.Create(doc)
	return result.Error
}

func (r *DocumentRepository) GetDocument(id string) (models.Document, error) {
	var document models.Document
	result := db.DB.First(&document, "id = ?", id)
	return document, result.Error
}

func (r *DocumentRepository) GetDocuments() ([]models.Document, error) {
	var documents []models.Document
	result := db.DB.Find(&documents)
	return documents, result.Error
}

func (r *DocumentRepository) UpdateDocument(id string, doc *models.Document) error {
	var document models.Document
	result := db.DB.Model(&document).Where("id = ?", id).Updates(doc)
	return result.Error
}

func (r *DocumentRepository) DeleteDocument(id string) error {
	result := db.DB.Delete(&models.Document{}, "id = ?", id)
	return result.Error
}
