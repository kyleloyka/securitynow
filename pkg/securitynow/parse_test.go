package securitynow

import (
	"net/url"
	"testing"
	"time"

	"github.com/kyleloyka/securitynow/pkg/episode"
)

func TestParseEpisodeFormat(t *testing.T) {
	var tests = []struct {
		header []byte
		want   episode.Episode
	}{
		{
			header: SN001Header,
			want: episode.Episode{
				Series:   "Security Now!",
				Number:   1,
				Date:     time.Date(2005, time.August, 19, 0, 0, 0, 0, time.UTC),
				Title:    "1: As the worm turns: The first Internet worms of 2005",
				Hosts:    "Steve Gibson & Leo Laporte",
				Homepage: "http://www.GRC.com/securitynow.htm",
				Media:    *SN001Url,
				Description: "How a never-disclosed Windows vulnerability was quickly " +
					"reverse-engineered from the patches to fix it and turned into more than 12 " +
					"potent and damaging Internet worms in three days.  What does this mean for " +
					"the future of Internet security?",
			},
		},
		{
			header: SN013Header,
			want: episode.Episode{
				Series:   "Security Now!",
				Number:   13,
				Date:     time.Date(2005, time.November, 10, 0, 0, 0, 0, time.UTC),
				Title:    "13: Unbreakable WiFi Security",
				Hosts:    "Steve Gibson & Leo Laporte",
				Homepage: "http://www.GRC.com/securitynow.htm",
				Media:    *SN013Url,
				Description: "Leo and I follow up on last week's discussion of the Sony Rootkit " +
					`debacle with the distressing news of "phoning home" (spyware) behavior from` +
					" the Sony DRM software, and the rootkit's exploitation by a new malicious " +
					"backdoor Trojan.  We then return to complete our discussion of WiFi " +
					"security, demystifying the many confusing flavors of WPA encryption and " +
					"presenting several critical MUST DO tips for WPA users.",
			},
		},
		{
			header: SN700Header,
			want: episode.Episode{
				Series:   "Security Now!",
				Number:   700,
				Date:     time.Date(2019, time.February, 5, 0, 0, 0, 0, time.UTC),
				Title:    "700: 700 & Counting",
				Hosts:    "Steve Gibson & Leo Laporte",
				Homepage: "https://www.grc.com/securitynow.htm",
				Media:    *SN700Url,
				Description: "This week we discuss Chrome getting spell check for URLs; a bunch " +
					"of Linux news with reasons to be sure you're patched up; some performance " +
					"enhancements, updates, additions, and deletions from Chrome and Firefox; " +
					"more Facebook nonsense; a bold move planned by the Japanese government; " +
					"Ubiquiti routers again in trouble; a hopeful and welcome new initiative for " +
					"the Chrome browser; a piece of errata; a quick SQRL update; and some " +
					"follow-up thoughts about VPN connectivity.",
			},
		},
	}

	for _, test := range tests {
		got, err := parseEpisode(test.header)
		if err != nil {
			t.Errorf("parseEpisode(%q), parse error: %q", test.want.Title, err)
		}
		if *got != test.want {
			t.Errorf("parseEpisode(%q), actual parse doesnt match expected", test.want.Title)

			/* Refection for printing each field of got vs want */
			// fields := reflect.TypeOf(*got)
			// rgot := reflect.ValueOf(*got)
			// rwant := reflect.ValueOf(test.want)
			// num := fields.NumField()
			// for i := 0; i < num; i++ {
			// 	field := fields.Field(i)
			// 	gotval := rgot.Field(i)
			// 	wantval := rwant.Field(i)
			// 	if gotval != wantval {
			// 		log.Println(field)
			// 		log.Println(gotval)
			// 		log.Println(wantval)
			// 		log.Println()
			// 	}
			// }
		}
	}
}

func TestParseEpisodeVarryingHeaderSize(t *testing.T) {
	var tests = []struct {
		percent float64
		err     error
	}{
		{0.0, errIncompleteHeader},
		{0.01, errIncompleteHeader},
		{0.1, errIncompleteHeader},
		{0.25, errIncompleteHeader},
		{0.46, errIncompleteHeader},
		{0.47, errIncompleteHeader},
		// The Description of SN 700's Show Notes ends at 745 bytes into the file.
		// The length of the SN700Header slice is 1580 (bytes), so 745/1580 = 0.4715189873
		// For SN700Header, only byte slices with greater than 0.4715189873 of the total should
		// parse successfully
		{0.4715, errIncompleteHeader},
		{0.4715189873, errIncompleteHeader},
		{0.48, nil},
		{0.5, nil},
		{0.9, nil},
	}

	for _, test := range tests {
		_, err := parseEpisode(percentOfBytes(SN700Header, test.percent))
		if err != test.err {
			t.Errorf("parseEpisode(%q) with partial data (%f%% of header). "+
				"got: err= %v, wanted err= %v",
				"S700 & Counting", test.percent, err, test.err)
		}
	}
}

func TestOrdinalToIntStringReplacement(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"June 1st, 2019", "June 1, 2019"},
		{"June 2nd, 2019", "June 2, 2019"},
		{"June 3rd, 2019", "June 3, 2019"},
		{"June 4th, 2019", "June 4, 2019"},
		{"June 14th, 2019", "June 14, 2019"},
		{"June 21st, 2019", "June 21, 2019"},
		{"June 22nd, 2019", "June 22, 2019"},
		{"June 23rd, 2019", "June 23, 2019"},
		{"August 1st, 2019", "August 1, 2019"}, // August has substring "st" like first "1st"
	}

	for _, test := range tests {
		if got := ordinalToIntReplacement(test.input); got != test.want {
			t.Errorf("ordinalToIntReplacement(%q) = %q, want %q", test.input, got, test.want)
		}
	}
}

// percentOfBytes returns the first p percent of bytes
func percentOfBytes(b []byte, percent float64) []byte {
	return b[:indexAtPercent(b, percent)]
}

// indexAtPercent returns an index corresponding to some percentage of the length of b
func indexAtPercent(b []byte, percent float64) int {
	return int(float64(len(b)) * percent)
}

var SN001Url, _ = url.Parse("https://media.blubrry.com/35015/www.podtrac.com/pts/redirect.mp3/" +
	"cdn.twit.tv/audio/sn/sn0001/sn0001.mp3")
var SN001Header = []byte(`GIBSON RESEARCH CORPORATION	http://www.GRC.com/

SERIES:		Security Now!
EPISODE:	#1
DATE:		August 19, 2005
TITLE:		As the worm turns: The first Internet worms of 2005
SPEAKERS:	Steve Gibson & Leo Laporte
SOURCE FILE:	http://media.GRC.com/sn/SN-001.mp3
FILE ARCHIVE:	http://www.GRC.com/securitynow.htm

DESCRIPTION:  How a never-disclosed Windows vulnerability was quickly reverse-engineered from the patches to fix it and turned into more than 12 potent and damaging Internet worms in three days.  What does this mean for the future of Internet security?

LEO LAPORTE:  Hi, this is Leo Laporte, and I'd like to introduce a brand-new podcast to the TWiT lineup, Security Now! with Steve Gibson.  This is Episode 1 for August 18, 2005.  You all know Steve Gibson.  He, of course, appears on TWiT regularly, This Week in Tech.  We've known him for a long time.  He's been a regular on the Screensavers and Call for Help.  And, you know, he's well-known to computer users everywhere for his products.  He's very well known to consumers for SpinRite, which was the inspiration for Norton Disk Doctor and still runs rings around it.  It is the ultimate hard-drive diagnostic recovery and file-saving tool.  It's really a remarkable tool that everybody should have a copy of from GRC.com.  But he's also been a very active consumer advocate, working really hard to help folks with their security.  He first came to my attention with the Click of Death, which was - that was the Zip drive Iomega...

STEVE GIBSON:  Right.`)

var SN013Url, _ = url.Parse("https://media.blubrry.com/35015/www.podtrac.com/pts/redirect.mp3/" +
	"cdn.twit.tv/audio/sn/sn0013/sn0013.mp3")
var SN013Header = []byte(`GIBSON RESEARCH CORPORATION	http://www.GRC.com/

SERIES:		Security Now!
EPISODE:		#13
DATE:		November 10, 2005
Title:		Unbreakable WiFi Security 
SPEAKERS:	Steve Gibson & Leo Laporte
SOURCE FILE:	http://media.GRC.com/sn/SN-013.mp3
FILE ARCHIVE:	http://www.GRC.com/securitynow.htm
	
DESCRIPTION:  Leo and I follow up on last week's discussion of the Sony Rootkit debacle with the distressing news of "phoning home" (spyware) behavior from the Sony DRM software, and the rootkit's exploitation by a new malicious backdoor Trojan.  We then return to complete our discussion of WiFi security, demystifying the many confusing flavors of WPA encryption and presenting several critical MUST DO tips for WPA users.

LEO LAPORTE:  This is Security Now! with Steve Gibson, lucky Episode 13:  Wireless Encryption Part 2.

Steve Gibson is back, the hero of the hour.  Did you get a lot of calls about the Sony rootkit?
`)

var SN700Url, _ = url.Parse("https://media.blubrry.com/35015/www.podtrac.com/pts/redirect.mp3/" +
	"cdn.twit.tv/audio/sn/sn0700/sn0700.mp3")
var SN700Header = []byte(`GIBSON RESEARCH CORPORATION		https://www.GRC.com/

SERIES:		Security Now!
EPISODE:	#700
DATE:		February 5, 2019
TITLE:		700 & Counting
HOSTS:	Steve Gibson & Leo Laporte
SOURCE:	https://media.grc.com/sn/sn-700.mp3
ARCHIVE:	https://www.grc.com/securitynow.htm

DESCRIPTION:  This week we discuss Chrome getting spell check for URLs; a bunch of Linux news with reasons to be sure you're patched up; some performance enhancements, updates, additions, and deletions from Chrome and Firefox; more Facebook nonsense; a bold move planned by the Japanese government; Ubiquiti routers again in trouble; a hopeful and welcome new initiative for the Chrome browser; a piece of errata; a quick SQRL update; and some follow-up thoughts about VPN connectivity.

SHOW TEASE:  It's time for Security Now!.  Steve Gibson is here.  Lots to talk about, including new systemd vulnerabilities.  Linux users, listen up.  We'll also talk a little bit about Chrome, a new feature giving us URL spell checking, and why TLS 1.0 and 1.1 are soon to hit the highway.  It's all coming up next on Security Now!.

LEO LAPORTE:  This is Security Now! with Steve Gibson, Episode 700, recorded Tuesday, February 5th, 2019:  700 & Counting.

It's time for Security Now!, the show where we cover the latest developments in the world of security and privacy, help you understand how computing works, and have a little fun along the way with this guy right here, Steve Gibson.  He's the commander in chief of the good ship Security Now!.  Aye aye, sir.  What you pointing - that is not the logo you want.  Maybe this.

`)
