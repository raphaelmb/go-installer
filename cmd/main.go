package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/raphaelmb/go-update/internal/client"
	"github.com/raphaelmb/go-update/internal/installer"
	"github.com/raphaelmb/go-update/internal/scraper"
	"github.com/raphaelmb/go-update/internal/util"
)

const url = "https://go.dev/dl/"

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-signals
		cancel()
	}()

	currentVersion := util.GetCurrentVersion()

	if currentVersion == "" {
		util.PrintInfo("go not installed")
	} else {
		util.PrintInfo(fmt.Sprint("current version ", currentVersion))
	}
	version, err := scraper.Scrape(url)
	if err != nil {
		util.PrintError(fmt.Sprint("error fetching go version: ", err))
		return
	}
	if util.CheckCurrentVersion(version) {
		util.PrintInfo("latest version already installed")
		return
	}
	util.PrintInfo(fmt.Sprint("latest version ", util.Sanitize(version)))
	util.PrintInfo("downloading")
	f, err := client.Download(ctx, url, version)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer os.Remove(f)
	util.PrintInfo("installing")
	err = installer.Install(f)
	if err != nil {
		util.PrintError(fmt.Sprint("error installing: ", err))
		return
	}
	util.PrintInfo("done")
	if currentVersion == "" {
		util.PrintInfo("remember to add /usr/local/go/bin to the PATH environment variable")
	}
}
