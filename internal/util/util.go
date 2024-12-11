package util

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func PrintStatus(s string) {
	fmt.Println(s)
}

func Reg(s string) bool {
	const regex = "go[1-9].([0-9][0-9]).[0-9].(linux-amd64.tar.gz)"
	r := regexp.MustCompile(regex)
	return r.Match([]byte(s))
}

func Sanitize(str string) string {
	parts := strings.Split(str, ".")
	return strings.Join(parts[:2+1], ".")
}

func CheckCurrentVersion(dl string) bool {
	return GetCurrentVersion() == Sanitize(dl)
}

func GetCurrentVersion() string {
	cmd := exec.Command("go", "version")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}
	v := strings.Split(string(output), " ")
	return v[2]
}
