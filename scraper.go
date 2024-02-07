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
			l = append(l, strings.TrimSpace(h.Text))
		}
	})
	c.Visit(url)
	return l[0]
}
