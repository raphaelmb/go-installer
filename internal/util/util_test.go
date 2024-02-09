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

	PrintStatus("testing")
	_ = w.Close()

	os.Stdout = oldStdOut

	out, _ := io.ReadAll(r)
	expected := fmt.Sprintf("%s\n", "testing")

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
