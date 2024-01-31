package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/schollz/progressbar/v3"
)

func download(version string) (string, error) {
	fullURLFile := fmt.Sprintf("https://go.dev/dl/%s", version)

	fileURL, err := url.Parse(fullURLFile)
	if err != nil {
		return "", fmt.Errorf("error parsing file URL")
	}
	path := fileURL.Path
	segments := strings.Split(path, "/")
	fileName := segments[len(segments)-1]

	file, err := os.Create(fileName)
	if err != nil {
		return "", fmt.Errorf("error creating file")
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	resp, err := client.Get(fullURLFile)
	if err != nil {
		return "", fmt.Errorf("error getting file")
	}
	defer resp.Body.Close()

	bar := progressbar.DefaultBytes(resp.ContentLength, "downloading")
	_, err = io.Copy(io.MultiWriter(file, bar), resp.Body)
	if err != nil {
		return "", fmt.Errorf("error writing file")
	}
	defer file.Close()

	return fileName, nil
}
