package main

import (
	"fmt"
	"regexp"
)

func printStatus(s string) {
	fmt.Println(s)
}

func reg(s string) bool {
	const regex = "go[1-9].([0-9][0-9]).[0-9].(linux-amd64.tar.gz)"
	r := regexp.MustCompile(regex)
	return r.Match([]byte(s))
}
