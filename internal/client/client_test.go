package client

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func CreateFile(t *testing.T) *os.File {
	f, err := os.Create("go1.21.6.linux-amd64.tar.gz")
	if err != nil {
		t.Errorf("error creating test file: %v\n", err)
	}
	if err := f.Truncate(1e3); err != nil {
		t.Errorf("error truncating test file: %v\n", err)
	}
	return f
}

func StartServer(t *testing.T, f *os.File) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dl, err := io.ReadAll(f)
		if err != nil {
			t.Errorf("error reading test file: %v\n", err)
		}
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		w.Write(dl)
	}))
}

func TestDownloadSuccess(t *testing.T) {
	f := CreateFile(t)
	defer f.Close()
	defer os.Remove(f.Name())
	server := StartServer(t, f)
	defer server.Close()

	value, err := Download(context.Background(), server.URL+"/", f.Name())
	if err != nil {
		t.Errorf("error downloading file: %v\n", err)
	}

	_, err = os.Stat(value)
	if err != nil {
		t.Errorf("file doest not exist: %v\n", err)
	}
}

func TestDownloadFailure(t *testing.T) {
	f := CreateFile(t)
	defer f.Close()
	defer os.Remove(f.Name())
	server := StartServer(t, f)
	defer server.Close()

	tests := []struct {
		Name      string
		Expected  string
		ServerURL string
		Filename  string
	}{
		{"Error parsing URL", "error parsing file URL", server.URL, f.Name()},
		{"Error creating file", "error creating file", server.URL + "/", ""},
		{"Error getting file", "error getting file", "/", f.Name()},
	}

	for _, v := range tests {
		_, err := Download(context.Background(), v.ServerURL, v.Filename)
		if err.Error() != v.Expected {
			t.Errorf("%s: expected '%s' but got '%s'", v.Name, v.Expected, err.Error())
		}
	}
}
