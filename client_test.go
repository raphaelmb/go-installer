package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestDownload(t *testing.T) {
	f, err := os.Create("go1.21.6.linux-amd64.tar.gz")
	if err != nil {
		t.Errorf("error creating test file: %v\n", err)
	}
	defer f.Close()
	defer os.Remove(f.Name())
	if err := f.Truncate(1e3); err != nil {
		t.Errorf("error truncating test file: %v\n", err)
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dl, err := io.ReadAll(f)
		if err != nil {
			t.Errorf("error reading test file: %v\n", err)
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		w.Write(dl)
	}))
	defer server.Close()

	value, err := download(server.URL+"/", f.Name())
	if err != nil {
		t.Errorf("error downloading file: %v\n", err)
	}

	_, err = os.Stat(value)
	if err != nil {
		t.Errorf("file doest not exist: %v\n", err)
	}
}
