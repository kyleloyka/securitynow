package securitynow

import (
	"bytes"
	"io"
	"testing"
)

func TestFeedGeneration(t *testing.T) {
	feed := NewFeed(2019)
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
	// f, _ := os.OpenFile("feed.xml", os.O_CREATE|os.O_RDWR, 0755)
	// feed.Write(f)
}

var expectedXML = `<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0" xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd">
  <channel>
    <title>Security Now! - 2019</title>
    <link>https://www.twit.tv/shows/security-now</link>
    <description>Steve Gibson, the man who coined the term spyware and created the first anti-spyware program, creator of Spinrite and ShieldsUP, discusses the hot topics in security today with Leo Laporte.</description>
    <generator>go podcast v1.3.1 (github.com/eduncan911/podcast)</generator>
    <language>en-us</language>
    <lastBuildDate>Tue, 01 Jan 2019 00:00:00 +0000</lastBuildDate>
    <pubDate>Tue, 01 Jan 2019 00:00:00 +0000</pubDate>
    <image>
      <url>https://elroycdn.twit.tv/sites/default/files/styles/twit_album_art_2048x2048/public/images/shows/security_now/album_art/audio/sn1400audio.jpg</url>
      <title>Security Now! - 2019</title>
      <link>https://www.twit.tv/shows/security-now</link>
    </image>
    <itunes:summary><![CDATA[Steve Gibson, the man who coined the term spyware and created the first anti-spyware program, creator of Spinrite and ShieldsUP, discusses the hot topics in security today with Leo Laporte.]]></itunes:summary>
    <itunes:image href="https://elroycdn.twit.tv/sites/default/files/styles/twit_album_art_2048x2048/public/images/shows/security_now/album_art/audio/sn1400audio.jpg"></itunes:image>
    <item>
      <guid>https://media.blubrry.com/35015/www.podtrac.com/pts/redirect.mp3/cdn.twit.tv/audio/sn/sn0001/sn0001.mp3</guid>
      <title>As the worm turns: The first Internet worms of 2005</title>
      <link>https://media.blubrry.com/35015/www.podtrac.com/pts/redirect.mp3/cdn.twit.tv/audio/sn/sn0001/sn0001.mp3</link>
      <description>How a never-disclosed Windows vulnerability was quickly reverse-engineered from the patches to fix it and turned into more than 12 potent and damaging Internet worms in three days.  What does this mean for the future of Internet security?&#xA;&#xA;Show notes:&#xA;https://www.grc.com/sn/sn-001.pdf&#xA;https://www.grc.com/sn/sn-001.txt&#xA;&#xA;Hosts: Steve Gibson &amp; Leo Laporte</description>
      <pubDate>Fri, 19 Aug 2005 00:00:00 +0000</pubDate>
      <enclosure url="https://media.blubrry.com/35015/www.podtrac.com/pts/redirect.mp3/cdn.twit.tv/audio/sn/sn0001/sn0001.mp3" length="0" type="audio/mpeg"></enclosure>
      <itunes:summary><![CDATA[How a never-disclosed Windows vulnerability was quickly reverse-engineered from the patches to fix it and turned into more than 12 potent and damaging Internet worms in three days.  What does this mean for the future of Internet security?

Show notes:
https://www.grc.com/sn/sn-001.pdf
https://www.grc.com/sn/sn-001.txt

Hosts: Steve Gibson & Leo Laporte]]></itunes:summary>
      <itunes:image href="https://elroycdn.twit.tv/sites/default/files/styles/twit_album_art_2048x2048/public/images/shows/security_now/album_art/audio/sn1400audio.jpg"></itunes:image>
    </item>
    <item>
      <guid>https://media.blubrry.com/35015/www.podtrac.com/pts/redirect.mp3/cdn.twit.tv/audio/sn/sn0700/sn0700.mp3</guid>
      <title>700 &amp; Counting</title>
      <link>https://media.blubrry.com/35015/www.podtrac.com/pts/redirect.mp3/cdn.twit.tv/audio/sn/sn0700/sn0700.mp3</link>
      <description>This week we discuss Chrome getting spell check for URLs; a bunch of Linux news with reasons to be sure you&#39;re patched up; some performance enhancements, updates, additions, and deletions from Chrome and Firefox; more Facebook nonsense; a bold move planned by the Japanese government; Ubiquiti routers again in trouble; a hopeful and welcome new initiative for the Chrome browser; a piece of errata; a quick SQRL update; and some follow-up thoughts about VPN connectivity.&#xA;&#xA;Show notes:&#xA;https://www.grc.com/sn/sn-700.pdf&#xA;https://www.grc.com/sn/sn-700.txt&#xA;&#xA;Hosts: Steve Gibson &amp; Leo Laporte</description>
      <pubDate>Tue, 05 Feb 2019 00:00:00 +0000</pubDate>
      <enclosure url="https://media.blubrry.com/35015/www.podtrac.com/pts/redirect.mp3/cdn.twit.tv/audio/sn/sn0700/sn0700.mp3" length="0" type="audio/mpeg"></enclosure>
      <itunes:summary><![CDATA[This week we discuss Chrome getting spell check for URLs; a bunch of Linux news with reasons to be sure you're patched up; some performance enhancements, updates, additions, and deletions from Chrome and Firefox; more Facebook nonsense; a bold move planned by the Japanese government; Ubiquiti routers again in trouble; a hopeful and welcome new initiative for the Chrome browser; a piece of errata; a quick SQRL update; and some follow-up thoughts about VPN connectivity.

Show notes:
https://www.grc.com/sn/sn-700.pdf
https://www.grc.com/sn/sn-700.txt

Hosts: Steve Gibson & Leo Laporte]]></itunes:summary>
      <itunes:image href="https://elroycdn.twit.tv/sites/default/files/styles/twit_album_art_2048x2048/public/images/shows/security_now/album_art/audio/sn1400audio.jpg"></itunes:image>
    </item>
  </channel>
</rss>`