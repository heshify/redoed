package models

import (
	"time"

	"github.com/google/uuid"
)

type Document struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title     string    `gorm:"not null" json:"title"`
	Content   string    `json:"content"`
	UserID    uint      `gorm:"not null;index" json:"user_id"`
	User      User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
