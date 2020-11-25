package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/kyleloyka/securitynow/pkg/securitynow"
)

var outputFolder = "docs"

func makeOutputFolder() error {
	if err := os.Mkdir(outputFolder, 0755); err != nil {
		if !os.IsExist(err) {
			return fmt.Errorf("couldn't create ./generated_feeds directory: %s", err)
		}
	}
	return nil
}

func writeFeedToFile(feed *securitynow.Feed, singleFeed bool) {
	writeToFile := true
	fullpath := filepath.Join(outputFolder, "/sn-%s.xml")

	specifier := fmt.Sprintf("%d", feed.Year)
	if singleFeed {
		specifier = "all"
	}

	file, err := os.OpenFile(fmt.Sprintf(fullpath, specifier),
		os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		log.Printf("Error: can't open "+fullpath+": %v", specifier, err)
		writeToFile = false
	}
	if writeToFile {
		err = feed.Write(file)
		if err != nil {
			log.Printf("Error: writing feed: %v", err)
		}
	}
}
