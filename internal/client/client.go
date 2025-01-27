package client

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func Download(ctx context.Context, fullUrl, version string) (string, error) {
	fullURLFile := fmt.Sprintf("%s%s", fullUrl, version)

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
	defer file.Close()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURLFile, nil)
	if err != nil {
		return "", fmt.Errorf("error making request")
	}

	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error getting file")
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			fmt.Println("\naborted by user")
			os.Remove(fileName)
			os.Exit(0)
		}
		return "", fmt.Errorf("error writing file: %w", err)
	}

	return fileName, nil
}
