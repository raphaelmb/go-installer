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
	util.PrintStatus("getting version")
	pag := 0
	v := scraper.GetVersions(url, pag)
	fmt.Println(v)
	util.PrintStatus("done")
	f, err := client.Download(url, v[0])
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(f)
	util.PrintStatus("done")
	util.PrintStatus("installing...")
	err = installer.Install(f)
	if err != nil {
		log.Fatal(err)
	}
	util.PrintStatus("done")
}
