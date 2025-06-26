package util

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

const RESET = "\033[0m"

const ERROR = "\033[31;1m"
const WARN = "\033[33;1m"
const INFO = "\033[97;1m"

func PrintInfo(s string) {
	fmt.Println(INFO + "info: " + RESET + s)
}

func PrintWarn(s string) {
	fmt.Println(WARN + "warn: " + RESET + s)
}

func PrintError(s string) {
	fmt.Println(ERROR + "error: " + RESET + s)
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
