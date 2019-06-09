package securitynow

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/kyleloyka/securitynow/pkg/episode"
)

// ErrEpisodeNotesNotFound is returned when the attempt to fetch shownotes returns a 404 statuscode
var ErrEpisodeNotesNotFound = errors.New("could not retrieve show notes for episode")

// Fetch gets the metadata by parsing the show notes for a particular Security Now episode.
// If the show notes cannot be retrieved, a minimal show entry is created by assuming the episode
// media url and title name. In this case, ErrEpisodeNotesNotFound will be returned along with the
// minimal episode entry.
func Fetch(episodeNumber int) (*episode.Episode, error) {
	metadataURL := fmt.Sprintf(showNotesURL, episodeNumber)

	resp, err := http.Get(metadataURL)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == http.StatusNotFound {
			ep, err := minimalParseEpisode(episodeNumber)
			if err != nil {
				return nil, fmt.Errorf("Security Now Fetch: couldn't retrieve show notes, "+
					"recovery failed: %s", err)
			}
			return ep, ErrEpisodeNotesNotFound
		}
		return nil, fmt.Errorf("Security Now Fetch: %s failed with status (%d) %s",
			metadataURL, resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()

	// try to read the whole header. If header is too large,
	// double the number of bytes read and try again
	var ep *episode.Episode
	for i := 1; i < 3; i++ {
		numBytes := showNotesHeaderSize * int64(i)
		header, err := ioutil.ReadAll(io.LimitReader(resp.Body, numBytes))
		if err != nil {
			return nil, err
		}
		ep, err = parseEpisode(header)
		if err != nil {
			if err == errIncompleteHeader {
				// resets resp.Body back to its initial state
				resp.Body = ioutil.NopCloser(io.MultiReader(bytes.NewReader(header), resp.Body))
				continue
			}
			return nil, err
		}
		break
	}

	return ep, nil
}
