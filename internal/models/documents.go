package models

import (
	"time"

	"github.com/google/uuid"
)

type Document struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    uuid.UUID `json:"author_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
