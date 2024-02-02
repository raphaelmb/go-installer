package main

import (
	"log"
	"os"
)

func main() {
	url := "https://go.dev/dl/"
	printStatus("getting version")
	v := getVersions(url)
	printStatus("done")
	f, err := download(url, v)
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
