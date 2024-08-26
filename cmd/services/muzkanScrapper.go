package services

import (
	"FindVibeGo/cmd/models"
	"FindVibeGo/cmd/scrapper"
	"github.com/gocolly/colly/v2"
	"github.com/google/uuid"
)

type MuzkanScrapperService struct {
	muzkanLink string
}

func NewMuzkanScrapperService() *MuzkanScrapperService {
	muzkanLink := "https://muzkan.net/?q="
	return &MuzkanScrapperService{muzkanLink: muzkanLink}
}

func (m *MuzkanScrapperService) GetSongs(searchQuery string) ([]models.Song, error) {
	collector := scrapper.GetInstance()

	songs := make([]models.Song, 0, 40)
	collector.OnHTML(".files__wrapper", func(e *colly.HTMLElement) {
		e.ForEach(".file", func(i int, e *colly.HTMLElement) {
			id := uuid.NewString()
			image := e.ChildAttr("img", "data-src")
			artist := e.ChildText("h4")
			title := e.ChildText("h5")
			link := e.ChildAttr(".button", "mp3source")

			song := models.Song{Id: id, Artist: artist, Title: title, Image: image, Link: link}
			songs = append(songs, song)
		})
	})
	err := collector.Visit(m.muzkanLink + searchQuery)
	if err != nil {
		return nil, err
	}

	return songs, nil
}
