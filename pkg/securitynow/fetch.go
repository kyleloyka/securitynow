package securitynow

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/kyleloyka/securitynow/pkg/episode"
)

var showNotesURL = "https://www.grc.com/sn/sn-%03d.txt"
var showNotesHeaderSize int64 = 3000

// Fetch gets the metadata for a particular Security Now episode
func Fetch(episodeNumber int) (*episode.Episode, error) {
	metadataURL := fmt.Sprintf(showNotesURL, episodeNumber)

	resp, err := http.Get(metadataURL)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
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
