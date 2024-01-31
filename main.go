package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	printStatus("getting version")
	v := getVersions()
	printStatus("done")
	f, err := download(strings.TrimSpace(v))
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(f)
	printStatus("done")
	printStatus("installing...")
	err = install(f)
	if err != nil {
		log.Fatal(err)
	}
	printStatus("done")
}
