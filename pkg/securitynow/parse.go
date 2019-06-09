package securitynow

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/kyleloyka/securitynow/pkg/episode"
)

var errIncompleteHeader = errors.New("Show notes header too small." +
	"Couldn't parse all episode metadata")

func parseEpisode(body []byte) (*episode.Episode, error) {
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
		fields[x[0]] = strings.Trim(strings.Replace(x[1], "\t", " ", -1), " ")
		if x[0] == "DESCRIPTION" || linesReadAfterDesc != 0 {
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
	ep.Series = fields["SERIES"]
	ep.Title = fields["TITLE"]
	ep.Description = fields["DESCRIPTION"]

	var err error
	ep.Homepage, err = findField(fields, []string{"ARCHIVE", "FILE ARCHIVE"}, "Homepage")
	if err != nil {
		return nil, err
	}

	ep.Hosts, err = findField(fields, []string{"HOSTS", "SPEAKERS"}, "Hosts")
	if err != nil {
		return nil, err
	}

	if number, err := strconv.Atoi(strings.TrimPrefix(fields["EPISODE"], "#")); err == nil {
		ep.Number = number
	} else {
		return nil, err
	}

	media := fmt.Sprintf(cdnMP3URL, ep.Number, ep.Number)
	if link, err := url.Parse(media); err == nil {
		ep.Media = *link
	} else {
		return nil, err
	}

	if date, err := time.Parse("January 2, 2006", fields["DATE"]); err == nil {
		ep.Date = date
	} else {
		return nil, err
	}

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
