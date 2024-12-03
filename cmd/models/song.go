package models

import (
	"github.com/google/uuid"
)

type Song struct {
	Id    uuid.UUID `json:"id"`
	Title string    `json:"title"`
	Image string    `json:"image"`
	Link  string    `json:"link"`
}
