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
	muzkanLink := "https://muzsky.net/search/"
	return &MuzkanScrapperService{muzkanLink: muzkanLink}
}

func (m *MuzkanScrapperService) GetSongs(searchQuery string) ([]models.Song, error) {
	collector := scrapper.GetInstance()

	songs := make([]models.Song, 40)
	collector.OnHTML("tbody", func(e *colly.HTMLElement) {

		e.ForEach("tr", func(i int, e *colly.HTMLElement) {
			id := uuid.New()
			image := e.ChildAttr("img", "data-src")
			title := e.ChildText("a")
			link := e.ChildAttr("div[data-id]", "data-id")
			song := models.Song{Id: id, Title: title, Image: image, Link: link}
			songs[i] = song
		})

	})

	err := collector.Visit(m.muzkanLink + searchQuery)
	if err != nil {
		return nil, err
	}

	return songs, nil
}
