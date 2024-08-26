package models

type Song struct {
	Id     string `json:"id"`
	Artist string `json:"artist"`
	Title  string `json:"title"`
	Image  string `json:"image"`
	Link   string `json:"link"`
}
