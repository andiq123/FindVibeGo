package models

import "github.com/google/uuid"

type Song struct {
	Id        uuid.UUID `json:"id"`
	Artist    string    `json:"artist"`
	Title     string    `json:"title"`
	Image     string    `json:"image"`
	Link      string    `json:"link"`
	SongOrder int       `json:"songOrder"`
	UserId    uuid.UUID `json:"userId"`
}
