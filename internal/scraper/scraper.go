package scraper

import (
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/raphaelmb/go-update/internal/util"
)

func GetVersions(url string, pag int) []string {
	c := colly.NewCollector()
	l := []string{}
	c.OnHTML(".filename", func(h *colly.HTMLElement) {
		if util.Reg(h.Text) {
			l = append(l, strings.TrimSpace(h.Text))
		}
	})
	c.Visit(url)
	return pagination(pag, l)
}

func pagination(n int, l []string) []string {
	l = l[1:]
	switch {
	case n == 0:
		return l[:1]
	case n > 0:
		return l[:n]
	default:
		return l[:3]
	}
}
