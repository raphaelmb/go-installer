package scraper

import (
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/raphaelmb/go-update/internal/util"
)

func Scrape(url string) (string, error) {
	c := colly.NewCollector()
	var l string
	c.OnHTML("div.filename > span", func(h *colly.HTMLElement) {
		if util.Reg(h.Text) {
			l = strings.TrimSpace(h.Text)
		}
	})
	err := c.Visit(url)
	if err != nil {
		return "", err
	}
	return l, nil
}
