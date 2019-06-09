package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kyleloyka/securitynow/pkg/securitynow"
)

var outputFolder = "generated_feeds"

func makeOutputFolder() error {
	if err := os.Mkdir(outputFolder, 0755); err != nil {
		if !os.IsExist(err) {
			return fmt.Errorf("couldn't create ./generated_feeds directory: %s", err)
		}
	}
	return nil
}

func writeFeedToFile(feed *securitynow.Feed) {
	writeToFile := true
	file, err := os.OpenFile(fmt.Sprintf("generated_feeds/sn-%04d.xml", feed.Year),
		os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		log.Printf("Error: can't open generated_feeds/sn-%04d.xml: %v", feed.Year, err)
		writeToFile = false
	}
	if writeToFile {
		err = feed.Write(file)
		if err != nil {
			log.Printf("Error: writing feed: %v", err)
		}
	}
}
