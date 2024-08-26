package scrapper

import (
	"crypto/tls"
	"github.com/gocolly/colly/v2"
	"net/http"
)

func GetInstance() *colly.Collector {
	collector := colly.NewCollector(colly.CacheDir("./music_cache"))
	customTransport := &(*http.DefaultTransport.(*http.Transport))
	customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	collector.WithTransport(customTransport)

	return collector
}
