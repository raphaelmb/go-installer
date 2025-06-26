package util

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestPrintStatus(t *testing.T) {
	oldStdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintInfo("testing")
	_ = w.Close()

	os.Stdout = oldStdOut

	out, _ := io.ReadAll(r)
	expected := fmt.Sprint(INFO+"info: "+RESET, "testing\n")

	if string(out) != expected {
		t.Errorf(`"%s" expected but got: %s`, expected, string(out))
	}
}

func TestReg(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected bool
	}{
		{"success", "go1.20.2.linux-amd64.tar.gz", true},
		{"success", "go1.21.2.linux-amd64.tar.gz", true},
		{"success", "go1.19.2.linux-amd64.tar.gz", true},
		{"fail", "go1.2a.2.linux-amd64.tar.gz", false},
		{"fail", "go1.20.2.linux-amd32.tar.gz", false},
	}

	for _, e := range tests {
		result := Reg(e.content)
		if e.expected != result {
			t.Errorf("%s: expected %t but got %t", e.name, e.expected, result)
		}
	}
}

func TestSanitize(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected string
	}{
		{"success", "go1.20.2.linux-amd64.tar.gz", "go1.20.2"},
		{"success", "go1.21.2.linux-amd64.tar.gz", "go1.21.2"},
		{"success", "go1.19.2.linux-amd64.tar.gz", "go1.19.2"},
		{"success", "go1.1.1", "go1.1.1"},
	}

	for _, e := range tests {
		result := Sanitize(e.content)
		if e.expected != result {
			t.Errorf("%s: expected %s but got %s", e.name, e.expected, result)
		}
	}

}
