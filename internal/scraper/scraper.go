package scraper

import (
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/raphaelmb/go-update/internal/util"
)

func Scrape(url string) string {
	c := colly.NewCollector()
	var l string
	c.OnHTML("div.filename > span", func(h *colly.HTMLElement) {
		if util.Reg(h.Text) {
			l = strings.TrimSpace(h.Text)
		}
	})
	c.Visit(url)
	return l
}
