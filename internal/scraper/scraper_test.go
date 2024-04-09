package scraper

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func StartServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(`
			<!DOCTYPE html>
			<html lang="en">
			<head>
  			<meta charset=utf-8>
  			<title></title>
			</head>
			<body>
					<div class="filename">
						<span>go1.21.6.linux-amd64.tar.gz</span>
					</div>
			</body>
			</html>
		`))
	}))
}

func TestGetVersion(t *testing.T) {
	server := StartServer()
	defer server.Close()

	testsSuccess := []struct {
		Name     string
		Expected string
	}{
		{"Case latest", "go1.21.6.linux-amd64.tar.gz"},
	}

	for _, v := range testsSuccess {
		value, err := Scrape(server.URL)
		if err != nil {
			t.Errorf("error: %s", err)
		}
		if value != v.Expected {
			t.Errorf("%s: expected %s but got %s", v.Name, v.Expected, value)
		}
	}

	testsFailure := []struct {
		Name     string
		Expected string
	}{
		{"Case version not found", "Not Found"},
	}

	for _, v := range testsFailure {
		value, err := Scrape("https://go.dev/dl/123")
		if value != "" {
			t.Error("error: value should be empty")
		}
		if err.Error() != v.Expected {
			t.Errorf("%s: expected %s but got %s", v.Name, v.Expected, err.Error())
		}
	}
}
