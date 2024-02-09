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
					<div>
						<table>
							<thead></thead>
							<tbody>
								<tr>
									<td class="filename">
										<a class="download" href="/dl/go1.21.6.linux-amd64.tar.gz">go1.21.6.linux-amd64.tar.gz</a>
									</td>
								</tr>
								<tr>
									<td class="filename">
										<a class="download" href="/dl/go1.21.5.linux-amd64.tar.gz">go1.21.5.linux-amd64.tar.gz</a>
									</td>
								</tr>
								<tr>
									<td class="filename">
										<a class="download" href="/dl/go1.21.4.linux-amd64.tar.gz">go1.21.4.linux-amd64.tar.gz</a>
									</td>
								</tr>
								<tr>
									<td class="filename">
										<a class="download" href="/dl/go1.21.4.linux-amd64.tar.gz">go1.21.4.linux-amd64.tar.gz</a>
									</td>
								</tr>
								<tr>
									<td class="filename">
										<a class="download" href="/dl/go1.21.4.linux-amd64.tar.gz">go1.21.4.linux-amd64.tar.gz</a>
									</td>
								</tr>
							</tbody>
						</table>
					</div>
			</body>
			</html>
		`))
	}))
}

func TestGetVersion(t *testing.T) {
	server := StartServer()
	defer server.Close()

	tests := []struct {
		Name     string
		Expected []string
		Pag      int
	}{
		{"Case latest", []string{"go1.21.4.linux-amd64.tar.gz"}, 0},
		{"Case 2", []string{"go1.21.4.linux-amd64.tar.gz", "go1.21.4.linux-amd64.tar.gz"}, 2},
		{"Case 4", []string{"go1.21.4.linux-amd64.tar.gz", "go1.21.4.linux-amd64.tar.gz", "go1.21.4.linux-amd64.tar.gz", "go1.21.4.linux-amd64.tar.gz"}, 4},
		{"Case default", []string{"go1.21.4.linux-amd64.tar.gz", "go1.21.4.linux-amd64.tar.gz", "go1.21.4.linux-amd64.tar.gz"}, -1},
	}

	for _, v := range tests {
		value := GetVersions(server.URL, v.Pag)
		if len(value) != len(v.Expected) {
			t.Errorf("%s: expected length '%d', got '%d'", v.Name, len(v.Expected), len(value))
		}

	}

}
