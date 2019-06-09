package main

import (
	"flag"
	"log"

	"github.com/kyleloyka/securitynow/pkg/episode"
	"github.com/kyleloyka/securitynow/pkg/securitynow"
)

func main() {
	startEp := flag.Int("startEpisode", 1, "start feed at episode number")
	endEp := flag.Int("endEpisode", 717, "end feed at episode number")
	verbose := flag.Bool("v", false, "verbose logging (default: false)")
	flag.Parse()

	err := makeOutputFolder()
	if err != nil {
		log.Fatal(err)
	}
	generateFeed(*startEp, *endEp, *verbose)
}

func createFeed(year int, verbose bool) *securitynow.Feed {
	feed := securitynow.NewFeed(year)
	if verbose {
		log.Printf("Information: created new feed for year %d", year)
	}
	return feed
}

// addToFeed adds episode to the feed. creates a new feed if feed is nil.
func addToFeed(feed *securitynow.Feed, ep *episode.Episode, verbose bool) *securitynow.Feed {
	if feed != nil && ep.Date.Year() != 1970 && ep.Date.Year() != feed.Year {
		writeFeedToFile(feed)
		feed = nil
	}
	if feed == nil {
		feed = createFeed(ep.Date.Year(), verbose)
	}
	err := feed.AddEpisode(ep)
	if err != nil {
		log.Printf("Error: adding episode to feed: %v", err)
	} else if verbose {
		log.Printf("Information: added episode %d to year %d", ep.Number, feed.Year)
	}
	return feed
}

func generateFeed(start, end int, verboseLogging bool) {
	var feed *securitynow.Feed

	for i := start; i < end+1; i++ {
		ep, err := securitynow.Fetch(i)
		if err != nil {
			if err == securitynow.ErrEpisodeNotesNotFound {
				log.Printf("Warning: episode %d show notes could not be loaded. "+
					"Automatically generating title and CDN url for episode", i)
			} else {
				log.Printf("Error: episode %d: %v\n", i, err)
				continue
			}
		}
		feed = addToFeed(feed, ep, verboseLogging)
	}
	writeFeedToFile(feed)
}
