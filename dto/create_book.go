package dto

import "github.com/google/uuid"

type CreateBookRequest struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title" validate:"required"`
	Author      string    `json:"author" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Isbn        string    `json:"isbn" validate:"required"`
	Stock       int       `json:"stock" validate:"min=0"`
	Genre       string    `json:"genre" validate:"required"`
	PublishDate string    `json:"publish_date" validate:"required"`
}
