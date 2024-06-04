package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Post struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id,omitempty"`
	Title     string    `gorm:"not null;uniqueIndex" json:"title,omitempty"`
	Content   string    `gorm:"not null" json:"content,omitempty"`
	Image     string    `gorm:"not null" json:"image,omitempty"`
	User      uuid.UUID `gorm:"type:uuid;not null" json:"user,omitempty"`
	CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at,omitempty"`
}

type CreatePostRequest struct {
	Title     string    `json:"title" binding:"required"`
	Content   string    `json:"content,omitempty" binding:"required" `
	Image     string    `json:"image,omitempty" binding:"required"`
	User      uuid.UUID `json:"user,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type UpdatePostRequest struct {
	Title     string    `json:"title" binding:"required"`
	Content   string    `json:"content,omitempty" binding:"required" `
	Image     string    `json:"image,omitempty" binding:"required"`
	User      uuid.UUID `json:"user,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
