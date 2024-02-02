package main

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

func getVersions(url string) string {
	c := colly.NewCollector()
	l := []string{}
	c.OnHTML(".filename", func(h *colly.HTMLElement) {
		if reg(h.Text) {
			l = append(l, h.Text)
		}
	})
	c.Visit(url)
	return strings.TrimSpace(l[0])
}
