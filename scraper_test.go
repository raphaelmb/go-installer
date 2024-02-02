package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetVersion(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
							</tbody>
						</table>
					</div>
			</body>
			</html>
		`))
	}))
	defer server.Close()

	value := getVersions(server.URL)
	expected := "go1.21.6.linux-amd64.tar.gz"
	if value != expected {
		t.Errorf("Expected %s, got %s", expected, value)
	}
}
