package securitynow

import (
	"log"
	"math"
	"testing"
)

func TestProperDecoding(t *testing.T) {
	expectedDescription := "Steve and Leo describe the new security features Microsoft " +
		"has designed and built into their new version of Windows, Vista.  They examine the " +
		"impact of having such features built into the base product rather than offered by third " +
		"parties as add-ons.  And they carefully compare the security benefits of Vista on " +
		"64-bit versus 32 bit hardware platforms."

	resp, err := Fetch(66)
	if err != nil {
		t.Errorf("Decode test failed to retrieve show notes, %v", err)
	}
	if resp.Description != expectedDescription {
		log.Println(len(resp.Description))
		log.Println(len(expectedDescription))
		if len(resp.Description) == len(expectedDescription) {
			for idx, r := range resp.Description {
				if r != rune(expectedDescription[idx]) {
					log.Printf("line %d: got %q, want %q", idx, r, expectedDescription[idx])
					log.Printf(
						resp.Description[int(math.Max(0, float64(idx)-8)):int(math.Min(float64(len(resp.Description)), float64(idx)+8))])
				}
			}
		}
		t.Error("decoded doesnt match expected")
	}
}
