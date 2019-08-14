package crawly

import (
	"log"
	"strings"
	"sync"
	"time"
)

type NewsCollection struct {
	Url         string
	NewsEntries []NewsEntry `xml:"url"`
}

type NewsEntry struct {
	Url       string `xml:"loc"`
	Lastmod   string `xml:"lastmod"`
	Timestamp string
}

type SitemapCollection struct {
	Url      string
	Sitemaps []Sitemap `xml:"sitemap"`
}

type Sitemap struct {
	Url     string `xml:"loc"`
	Lastmod string `xml:"lastmod"`
}

// NewSitemapCollection constructor new SitemapCollection object
func NewSitemapCollection(url string) SitemapCollection {
	var smapscoll SitemapCollection
	smapscoll.Url = url
	return smapscoll
}

// NewNewsCollection constructor new NewsCollection object
func NewNewsCollection(url string) NewsCollection {
	var newscoll NewsCollection
	newscoll.Url = url
	return newscoll
}

// Crawl crawl sitemap containing multiple sitemaps
func (smapcoll *SitemapCollection) Crawl() {
	// Vars
	var wg sync.WaitGroup
	ch := make(chan int, 5)
	defer close(ch)

	// get sitemap data, parse xml to sitemap struct
	data := GetData(smapcoll.Url)
	ParseXML(&data, smapcoll)

	// crawl all sitemap urls
	for i := range smapcoll.Sitemaps {
		ch <- 1
		log.Println(smapcoll.Sitemaps[i].Url)
		log.Println("main: starting worker ", i)
		wg.Add(1)
		newscoll := NewNewsCollection(smapcoll.Sitemaps[i].Url)
		go newscoll.CrawlSync(&wg, ch)
	}
	log.Println("main: waiting for workers to finish...")
	wg.Wait()
	log.Println("main: completed")
}

// FilterKeywords filter crawled news
func (news *NewsCollection) FilterKeywords() {
	log.Println("filtering keywords...")

	var filterNews NewsCollection

	for _, n := range news.NewsEntries {
		for _, k := range KEYWORDS {
			if strings.Contains(strings.ToLower(n.Url), k) {
				filterNews.NewsEntries = append(filterNews.NewsEntries, n)
				break
			}
		}
	}

	if len(filterNews.NewsEntries) == 0 {
		log.Println("no elements found")
	}

	news.NewsEntries = filterNews.NewsEntries
}

// Crawl crawl website with news entries
func (news *NewsCollection) Crawl() {
	log.Println("crawling " + news.Url)

	// get data, parse xml to news struct, filter news
	body := GetData(news.Url)
	ParseXML(&body, news)
	news.FilterKeywords()

	// log crawled news and add timestamp
	for i := range news.NewsEntries {
		news.NewsEntries[i].Timestamp = time.Now().Format("2006-01-02 15:04:05 MST")
		log.Println(news.NewsEntries[i].Url)
	}
}

// CrawlSync like NewsCollection.Crawl() but with Waitgroup for crawling multiple websites in parallel
func (news *NewsCollection) CrawlSync(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	log.Println("crawling " + news.Url)

	// get data from url, parse xml to news struct, filter news
	body := GetData(news.Url)
	ParseXML(&body, news)
	news.FilterKeywords()

	log.Println(" entries found", len(news.NewsEntries))

	<-ch
}
