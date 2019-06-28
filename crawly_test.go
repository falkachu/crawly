package crawly

import (
	"log"
	"os"
	"testing"
)

func TestCrawlMulti(t *testing.T) {

	// Vars
	smaps := NewSitemapCollection("https://www.augsburger-allgemeine.de/sitemap.xml")


	// Logger Configuration
	f, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY, 0777)
	check(err)
	defer f.Close()

	// Logger Options
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.SetOutput(f)


	// Crawling
	smaps.Crawl()
}

func TestCrawlUrl(t *testing.T) {
	// Vars
	newscoll := NewNewsCollection("https://www.augsburger-allgemeine.de/news.xml")

	// Logger Configuration
	f, err := os.OpenFile("log2.txt", os.O_CREATE|os.O_WRONLY, 0777)
	check(err)
	defer f.Close()
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.SetOutput(f)

	// Crawling
	newscoll.Crawl()
}
