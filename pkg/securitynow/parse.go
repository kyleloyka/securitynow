package securitynow

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/kyleloyka/securitynow/pkg/episode"
)

var errIncompleteHeader = errors.New("Show notes header too small." +
	"Couldn't parse all episode metadata")

func parseEpisode(body []byte) (*episode.Episode, error) {
	// some show notes include non UTF-8 characters
	body = toValidUTF8(body)

	fields := make(map[string]string)
	buf := bytes.NewBuffer(body)
	scanner := bufio.NewScanner(buf)

	// want to read at least 1 line after the description to ensure that the byte slice contained
	// all of the description
	sawDescription := false
	linesReadAfterDesc := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		x := strings.SplitN(line, ":", 2)
		if len(x) <= 1 {
			continue
		}

		fieldName := strings.ToLower(x[0])
		fields[fieldName] = strings.Trim(strings.Replace(x[1], "\t", " ", -1), " ")
		if fieldName == "description" || linesReadAfterDesc != 0 {
			sawDescription = true
			linesReadAfterDesc++
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if !sawDescription || linesReadAfterDesc <= 1 {
		return nil, errIncompleteHeader
	}

	ep := new(episode.Episode)
	ep.Series = fields["series"]
	ep.Title = fields["title"]
	ep.Description = fields["description"]

	var err error
	ep.Homepage, err = findField(fields, []string{"archive", "file archive"}, "Homepage")
	if err != nil {
		return nil, err
	}

	ep.Hosts, err = findField(fields, []string{"hosts", "speakers"}, "Hosts")
	if err != nil {
		return nil, err
	}

	if number, err := strconv.Atoi(strings.TrimPrefix(fields["episode"], "#")); err == nil {
		ep.Number = number
	} else {
		return nil, err
	}

	media := generateCDNURL(ep.Number)
	if link, err := url.Parse(media); err == nil {
		ep.Media = *link
	} else {
		return nil, err
	}

	// Special handling for SN 199 date
	if len(fields["date"]) >= 3 && fields["date"][:3] == "une" {
		fields["date"] = "J" + fields["date"]
	}

	if date, err := time.Parse("January 2, 2006", ordinalToIntReplacement(fields["date"])); err == nil {
		ep.Date = date
	} else {
		return nil, err
	}

	return ep, nil
}

func minimalParseEpisode(episodeNumber int) (*episode.Episode, error) {
	ep := new(episode.Episode)
	u, err := url.Parse(generateCDNURL(episodeNumber))
	if err != nil {
		return nil, err
	}
	ep.Title = fmt.Sprintf("SN %d", episodeNumber)
	ep.Number = episodeNumber
	ep.Series = "Security Now!"
	ep.Hosts = showHosts
	ep.Homepage = showURL
	ep.Media = *u
	ep.Date = time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
	return ep, nil
}

func findField(fields map[string]string, fieldNames []string, modelName string) (string, error) {
	for _, name := range fieldNames {
		if n, ok := fields[name]; ok {
			return n, nil
		}
	}
	return "", fmt.Errorf("Could not find field for %q", modelName)
}

// GenerateCDNURL Generates the CDN mp3 url for the given episode number
func generateCDNURL(episodeNumber int) string {
	return fmt.Sprintf(cdnMP3URL, episodeNumber, episodeNumber)
}

// ordinalToIntReplacement replaces all ordinal (1st, 2nd, 3rd, 4th etc.) substrings with their
// corresponding integer values (1, 2, 3, 4, etc.)
func ordinalToIntReplacement(s string) string {
	re := regexp.MustCompile(`\b(\d{1,2})\s*([a-zA-Z]{2})\b`)
	return re.ReplaceAllString(s, "$1")
}

func toValidUTF8(enc []byte) []byte {
	out := make([]rune, 0, len(enc))
	for len(enc) > 0 {
		r, size := utf8.DecodeRune(enc)
		enc = enc[size:]
		if r == utf8.RuneError {
			continue
		}
		if r == '\x1e' {
			out = append(out, ' ')
			continue
		}
		out = append(out, r)
	}
	return []byte(string(out))
}
