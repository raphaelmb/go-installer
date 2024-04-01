package scraper

import (
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/raphaelmb/go-update/internal/util"
)

func Scrape(url string) string {
	c := colly.NewCollector()
	l := []string{}
	c.OnHTML(".filename", func(h *colly.HTMLElement) {
		if util.Reg(h.Text) {
			l = append(l, strings.TrimSpace(h.Text))
		}
	})
	c.Visit(url)
	return l[0]
}
