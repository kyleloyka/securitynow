package securitynow

import (
	"bytes"
	"io"
	"testing"
)

func TestFeedGeneration(t *testing.T) {
	feed := NewFeed(2019, false)
	ep1, err := parseEpisode(SN001Header)
	ep700, err := parseEpisode(SN700Header)
	if err != nil {
		t.Error("Failed to parse test-episode. FeedGeneration not tested.")
		return
	}
	feed.AddEpisode(ep1)
	feed.AddEpisode(ep700)

	var out io.Writer = new(bytes.Buffer)
	err = feed.Write(out)
	if err != nil {
		t.Errorf("FeedGeneration failed with error: %v", err)
	}

	got := out.(*bytes.Buffer).String()
	if got != expectedXML {
		t.Errorf("FeedGeneration output did not match expected output")
	}

	// // Uncomment to write generated xml to a file
	// f, _ := os.OpenFile("test-feed.xml", os.O_CREATE|os.O_RDWR, 0666)
	// feed.Write(f)
}

var expectedXML = `<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0" xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd">
  <channel>
    <title>Security Now - 2019</title>
    <link>https://www.twit.tv/shows/security-now</link>
    <description>Steve Gibson, the man who coined the term spyware and created the first anti-spyware program, creator of Spinrite and ShieldsUP, discusses the hot topics in security today with Leo Laporte.</description>
    <generator>go podcast v1.3.0 (github.com/eduncan911/podcast)</generator>
    <language>en-us</language>
    <lastBuildDate>Tue, 01 Jan 2019 00:00:00 +0000</lastBuildDate>
    <managingEditor> (Steve Gibson, Leo Laporte)</managingEditor>
    <pubDate>Tue, 01 Jan 2019 00:00:00 +0000</pubDate>
    <image>
      <url>https://raw.githubusercontent.com/kyleloyka/securitynow/master/assets/sn-image.png</url>
    </image>
    <itunes:author> (Steve Gibson, Leo Laporte)</itunes:author>
    <itunes:summary><![CDATA[Steve Gibson, the man who coined the term spyware and created the first anti-spyware program, creator of Spinrite and ShieldsUP, discusses the hot topics in security today with Leo Laporte.]]></itunes:summary>
    <itunes:image href="https://raw.githubusercontent.com/kyleloyka/securitynow/master/assets/sn-image.png"></itunes:image>
    <item>
      <guid>https://media.blubrry.com/35015/www.podtrac.com/pts/redirect.mp3/cdn.twit.tv/audio/sn/sn0001/sn0001.mp3</guid>
      <title>1: As the worm turns: The first Internet worms of 2005</title>
      <link>https://media.blubrry.com/35015/www.podtrac.com/pts/redirect.mp3/cdn.twit.tv/audio/sn/sn0001/sn0001.mp3</link>
      <description>How a never-disclosed Windows vulnerability was quickly reverse-engineered from the patches to fix it and turned into more than 12 potent and damaging Internet worms in three days.  What does this mean for the future of Internet security?&#xA;&#xA;Show notes:&#xA;https://www.grc.com/sn/sn-001.txt&#xA;https://www.grc.com/sn/sn-001.pdf&#xA;https://www.grc.com/sn/sn-001-notes.pdf&#xA;&#xA;Hosts: Steve Gibson &amp; Leo Laporte</description>
      <pubDate>Fri, 19 Aug 2005 00:00:00 +0000</pubDate>
      <enclosure url="https://media.blubrry.com/35015/www.podtrac.com/pts/redirect.mp3/cdn.twit.tv/audio/sn/sn0001/sn0001.mp3" length="0" type="audio/mpeg"></enclosure>
      <itunes:author> (Steve Gibson, Leo Laporte)</itunes:author>
      <itunes:summary><![CDATA[How a never-disclosed Windows vulnerability was quickly reverse-engineered from the patches to fix it and turned into more than 12 potent and damaging Internet worms in three days.  What does this mean for the future of Internet security?

Show notes:
<a href="https://www.grc.com/sn/sn-001.txt">https://www.grc.com/sn/sn-001.txt</a>
<a href="https://www.grc.com/sn/sn-001.pdf">https://www.grc.com/sn/sn-001.pdf</a>
<a href="https://www.grc.com/sn/sn-001-notes.pdf">https://www.grc.com/sn/sn-001-notes.pdf</a>

Hosts: Steve Gibson & Leo Laporte]]></itunes:summary>
      <itunes:image href="https://raw.githubusercontent.com/kyleloyka/securitynow/master/assets/sn-image.png"></itunes:image>
    </item>
    <item>
      <guid>https://media.blubrry.com/35015/www.podtrac.com/pts/redirect.mp3/cdn.twit.tv/audio/sn/sn0700/sn0700.mp3</guid>
      <title>700: 700 &amp; Counting</title>
      <link>https://media.blubrry.com/35015/www.podtrac.com/pts/redirect.mp3/cdn.twit.tv/audio/sn/sn0700/sn0700.mp3</link>
      <description>This week we discuss Chrome getting spell check for URLs; a bunch of Linux news with reasons to be sure you&#39;re patched up; some performance enhancements, updates, additions, and deletions from Chrome and Firefox; more Facebook nonsense; a bold move planned by the Japanese government; Ubiquiti routers again in trouble; a hopeful and welcome new initiative for the Chrome browser; a piece of errata; a quick SQRL update; and some follow-up thoughts about VPN connectivity.&#xA;&#xA;Show notes:&#xA;https://www.grc.com/sn/sn-700.txt&#xA;https://www.grc.com/sn/sn-700.pdf&#xA;https://www.grc.com/sn/sn-700-notes.pdf&#xA;&#xA;Hosts: Steve Gibson &amp; Leo Laporte</description>
      <pubDate>Tue, 05 Feb 2019 00:00:00 +0000</pubDate>
      <enclosure url="https://media.blubrry.com/35015/www.podtrac.com/pts/redirect.mp3/cdn.twit.tv/audio/sn/sn0700/sn0700.mp3" length="0" type="audio/mpeg"></enclosure>
      <itunes:author> (Steve Gibson, Leo Laporte)</itunes:author>
      <itunes:summary><![CDATA[This week we discuss Chrome getting spell check for URLs; a bunch of Linux news with reasons to be sure you're patched up; some performance enhancements, updates, additions, and deletions from Chrome and Firefox; more Facebook nonsense; a bold move planned by the Japanese government; Ubiquiti routers again in trouble; a hopeful and welcome new initiative for the Chrome browser; a piece of errata; a quick SQRL update; and some follow-up thoughts about VPN connectivity.

Show notes:
<a href="https://www.grc.com/sn/sn-700.txt">https://www.grc.com/sn/sn-700.txt</a>
<a href="https://www.grc.com/sn/sn-700.pdf">https://www.grc.com/sn/sn-700.pdf</a>
<a href="https://www.grc.com/sn/sn-700-notes.pdf">https://www.grc.com/sn/sn-700-notes.pdf</a>

Hosts: Steve Gibson & Leo Laporte]]></itunes:summary>
      <itunes:image href="https://raw.githubusercontent.com/kyleloyka/securitynow/master/assets/sn-image.png"></itunes:image>
    </item>
  </channel>
</rss>`
