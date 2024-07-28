package models

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/google/uuid"
)

type Song struct {
	Id     string `json:"id"`
	Artist string `json:"artist"`
	Title  string `json:"title"`
	Image  string `json:"image"`
	Link   string `json:"link"`
}

const LINK = "https://muzkan.net/?q="

func SearchSongs(search string) ([]Song, error) {
	collector := colly.NewCollector(colly.CacheDir("./music_cache"))

	songs := make([]Song, 0, 40)
	collector.OnHTML(".files__wrapper", func(e *colly.HTMLElement) {
		e.ForEach(".file", func(i int, e *colly.HTMLElement) {
			id := uuid.NewString()
			image := e.ChildAttr("img", "data-src")
			artist := e.ChildText("h4")
			title := e.ChildText("h5")
			link := e.ChildAttr(".button", "mp3source")

			song := Song{Id: id, Artist: artist, Title: title, Image: image, Link: link}
			songs = append(songs, song)
		})
	})
	fmt.Println(LINK + search)
	err := collector.Visit(LINK + search)
	if err != nil {
		return nil, err
	}

	return songs, nil
}
