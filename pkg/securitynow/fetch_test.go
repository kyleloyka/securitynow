package securitynow

import (
	"testing"
)

func TestFetchEpisode(t *testing.T) {
	var tests = []struct {
		number int
		want   []byte
	}{
		{1, SN001Header},
		{700, SN700Header},
	}

	for _, test := range tests {
		ep, err := Fetch(test.number)
		if err != nil {
			t.Errorf("Failed to fetch Security Now episode #%d, %v", test.number, err)
			continue
		}
		trueparse, err := parseEpisode(test.want)
		if err != nil {
			t.Errorf("Failed to parse expected test case for SN #%d. Test is inconclusive",
				test.number)
			continue
		}

		if *ep != *trueparse {
			t.Errorf("Fetch(SN show #%d) did not return the expected Episode values", test.number)
		}
	}
}
