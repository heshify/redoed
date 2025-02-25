package models

import "gorm.io/gorm"

type Document struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
}
