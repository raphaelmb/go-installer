package util

import (
	"fmt"
	"regexp"
	"runtime"
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
	return runtime.Version() == Sanitize(dl)
}
