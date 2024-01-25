package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

var (
	fileName    string
	fullURLFile string
)

func reg(s string) bool {
	const regex = "go[1-9].([0-9][0-9]).[0-9].(linux-amd64.tar.gz)"
	r := regexp.MustCompile(regex)
	return r.Match([]byte(s))
}

func main() {
	v := getVersions()
	download(strings.TrimSpace(v))
}

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

func download(version string) {
	fullURLFile = fmt.Sprintf("https://go.dev/dl/%s", version)

	fileURL, err := url.Parse(fullURLFile)
	if err != nil {
		log.Fatal(err)
	}
	path := fileURL.Path
	segments := strings.Split(path, "/")
	fileName = segments[len(segments)-1]

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	resp, err := client.Get(fullURLFile)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Printf("Downloaded a file %s with size %d\n", fileName, size)
}
