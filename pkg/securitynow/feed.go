package securitynow

import (
	"fmt"
	"io"
	"time"

	"github.com/eduncan911/podcast"
	"github.com/kyleloyka/securitynow/pkg/episode"
)

// Feed is a Security Now podcast feed
type Feed struct {
	podcast.Podcast
	Year int
}

// NewFeed creates a new Security Now podcast feed
func NewFeed(year int, allYears bool) *Feed {
	feed := new(Feed)
	feed.Year = year
	create := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	modified := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	if modified.After(time.Now()) {
		modified = time.Now()
	}

	specifier := fmt.Sprintf("%d", year)
	if allYears {
		specifier = "All"
	}

	feed.Podcast = podcast.New(
		fmt.Sprintf("Security Now - %s", specifier),
		showURL,
		showSummary,
		&create, &modified,
	)

	// add some channel properties
	feed.Podcast.AddSummary(showSummary)
	feed.Podcast.AddImage(showLargeAlbumArt)
	feed.Podcast.AddAuthor(showHosts, "")
	// p.AddAtomLink(hostingpath + fmt.Sprintf("/securitynow-%d.rss", year))
	return feed
}

func toHref(url string) string {
	return fmt.Sprintf("<a href=\"%s\">%s</a>", url, url)
}

// AddEpisode adds an episode to the podcast feed
func (feed *Feed) AddEpisode(e *episode.Episode) error {
	fullDescription := fmt.Sprintf(e.Description+
		"\n\nShow notes:\n"+
		toHref(fmt.Sprintf(showTranscriptTXTURL, e.Number))+"\n"+
		toHref(fmt.Sprintf(showTranscriptPDFURL, e.Number))+"\n"+
		toHref(fmt.Sprintf(showNotesPDFURL, e.Number))+
		"\n\nHosts: %s", e.Hosts)

	plainDescription := fmt.Sprintf(e.Description+
		"\n\nShow notes:\n"+
		fmt.Sprintf(showTranscriptTXTURL, e.Number)+"\n"+
		fmt.Sprintf(showTranscriptPDFURL, e.Number)+"\n"+
		fmt.Sprintf(showNotesPDFURL, e.Number)+
		"\n\nHosts: %s", e.Hosts)

	item := podcast.Item{
		Title:       e.Title,
		Description: plainDescription,
		PubDate:     &e.Date,
	}
	item.AddImage(showLargeAlbumArt)
	item.AddSummary(fullDescription)
	item.AddEnclosure(e.Media.String(), podcast.MP3, 0)

	if _, err := feed.Podcast.AddItem(item); err != nil {
		return err
	}
	return nil
}

// Write attempts to write the podcast feed to a Writer. Also validates the feed data.
func (feed *Feed) Write(w io.Writer) error {
	// Podcast.Encode writes to an io.Writer
	if feed == nil {
		return fmt.Errorf("attempting to write nil feed")
	}
	if err := feed.Podcast.Encode(w); err != nil {
		return err
	}
	return nil
}
