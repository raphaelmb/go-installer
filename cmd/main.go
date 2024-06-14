package main

import (
	"fmt"
	"os"

	"github.com/raphaelmb/go-update/internal/client"
	"github.com/raphaelmb/go-update/internal/installer"
	"github.com/raphaelmb/go-update/internal/scraper"
	"github.com/raphaelmb/go-update/internal/util"
)

func main() {
	url := "https://go.dev/dl/"
	version, err := scraper.Scrape(url)
	if err != nil {
		fmt.Println("error fetching go version:", err)
		return
	}
	if util.CheckCurrentVersion(version) {
		fmt.Println("latest version already installed")
		return
	}
	fmt.Println("latest version:", util.Sanitize(version))
	f, err := client.Download(url, version)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer os.Remove(f)
	util.PrintStatus("installing...")
	err = installer.Install(f)
	if err != nil {
		fmt.Println("error installing:", err)
		return
	}
	util.PrintStatus("done")
	util.PrintStatus("remember to add /usr/local/go/bin to the PATH environment variable if needed")
}
