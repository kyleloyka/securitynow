// Episode defines the model for a generic podcast episode.
package episode

import (
	"net/url"
	"time"
)

type Episode struct {
	Series      string
	Number      int
	Date        time.Time
	Title       string
	Hosts       string
	Media       url.URL
	Homepage    string
	Description string
}
