package main

import "github.com/gocolly/colly/v2"

func getVersions() string {
	c := colly.NewCollector()
	l := []string{}
	c.OnHTML(".filename", func(h *colly.HTMLElement) {
		if reg(h.Text) {
			l = append(l, h.Text)
		}
	})
	c.Visit("https://go.dev/dl/")
	return l[0]
}
