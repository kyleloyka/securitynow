package main

import (
	"flag"
	"log"
	"time"

	"github.com/kyleloyka/securitynow/pkg/episode"
	"github.com/kyleloyka/securitynow/pkg/securitynow"
)

func main() {
	startEp := flag.Int("startEpisode", 1, "start feed at episode number")
	endEp := flag.Int("endEpisode", 794, "end feed at episode number")
	verbose := flag.Bool("v", false, "verbose logging (default: false)")
	singleFeed := flag.Bool("s", false, "specify if all episodes should be added to a single feed")
	flag.Parse()

	err := makeOutputFolder()
	if err != nil {
		log.Fatal(err)
	}
	generateFeed(*startEp, *endEp, *verbose, *singleFeed)
}

func createFeed(year int, verbose bool, singleFeed bool) *securitynow.Feed {
	feed := securitynow.NewFeed(year, singleFeed)

	if verbose {
		log.Printf("Information: created new feed for year %d", year)
	}
	return feed
}

// addToFeed adds episode to the feed. creates a new feed if feed is nil.
func addToFeed(feed *securitynow.Feed, ep *episode.Episode, verbose, singleFeed bool) *securitynow.Feed {
	if feed == nil {
		feed = createFeed(ep.Date.Year(), verbose, singleFeed)
	}
	err := feed.AddEpisode(ep)
	if err != nil {
		log.Printf("Error: adding episode to feed: %v", err)
	} else if verbose {
		log.Printf("Information: added episode %d to year %d", ep.Number, feed.Year)
	}
	writeFeedToFile(feed, singleFeed)
	return feed
}

func generateFeed(start, end int, verboseLogging, singleFeed bool) {
	var feed *securitynow.Feed
	var prevEp *episode.Episode

	for i := start; i < end+1; i++ {
		ep, err := securitynow.Fetch(i)
		if err != nil {
			if err == securitynow.ErrEpisodeNotesNotFound {
				log.Printf("Warning: episode %d show notes could not be loaded. "+
					"Automatically generating title and CDN url for episode", i)
				ep.Date = prevEp.Date.Add(time.Hour * 24 * 7)

			} else {
				log.Printf("Error: episode %d: %v\n", i, err)
				continue
			}
		}
		if !singleFeed && feed != nil && feed.Year != ep.Date.Year() {
			feed = nil
		}
		feed = addToFeed(feed, ep, verboseLogging, singleFeed)
		prevEp = ep
	}
}
