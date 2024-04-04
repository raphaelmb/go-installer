package main

import (
	"fmt"
	"log"
	"os"

	"github.com/raphaelmb/go-update/internal/client"
	"github.com/raphaelmb/go-update/internal/installer"
	"github.com/raphaelmb/go-update/internal/scraper"
	"github.com/raphaelmb/go-update/internal/util"
)

func main() {
	url := "https://go.dev/dl/"
	version := scraper.Scrape(url)
	fmt.Println("latest version:", version)
	f, err := client.Download(url, version)
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(f)
	util.PrintStatus("installing...")
	err = installer.Install(f)
	if err != nil {
		log.Fatal(err)
	}
	util.PrintStatus("done")
	util.PrintStatus("Remember to add /usr/local/go/bin to the PATH environment variable if needed")
}
