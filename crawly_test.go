package crawly

import (
	"log"
	"os"
	"testing"
)

/*
func TestCrawly(t *testing.T) {

	var smapsAA SitemapsAA
	var wg sync.WaitGroup

	f, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY, 0777)
	check(err)
	defer f.Close()

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.SetOutput(f)

	data := getData("https://www.augsburger-allgemeine.de/sitemap.xml")

	parseXml(&data, &smapsAA)

	for i, s := range smapsAA.Sitemaps{
		log.Println("main: starting worker ", i)
		wg.Add(1)
		go CrawlURL(&wg, s.URL)
	}

	log.Println("main: waiting for workers to finish...")
	wg.Wait()
	log.Println("main: completed")
}
*/

func TestCrawlyNewsAA(t *testing.T){
	// logging printed information to log.txt file
	f, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY, 0777)
	check(err)
	defer f.Close()

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.SetOutput(f)

	// function start
	log.Println("crawl started")
	CrawlURL("https://www.augsburger-allgemeine.de/news.xml")
	log.Println("crawl finished")
}
