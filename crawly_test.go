package crawly

import (
	"log"
	"os"
	"testing"
)

func TestCrawlMulti(t *testing.T) {

	// Vars
	smaps := NewSitemaps("https://www.all-in.de/sitemap.xml")


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
	url := "https://www.augsburger-allgemeine.de/news.xml"

	// Logger Configuration
	f, err := os.OpenFile("log2.txt", os.O_CREATE|os.O_WRONLY, 0777)
	check(err)
	defer f.Close()

	// Logger Options
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.SetOutput(f)

	// Crawling
	CrawlUrl(url)
}
