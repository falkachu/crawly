package crawly

import (
	"log"
	"strings"
	"sync"
)

type NewsCollection struct {
	Url         string
	NewsEntries []NewsEntry `xml:"url"`
}

type NewsEntry struct {
	Url     string `xml:"loc"`
	Lastmod string `xml:"lastmod"`
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
func NewSitemapCollection(index string) SitemapCollection {
	var smapscoll SitemapCollection
	smapscoll.Url = index
	return smapscoll
}

// NewNewsCollection constructor new NewsCollection object
func NewNewsCollection(index string) NewsCollection{
	var newscoll NewsCollection
	newscoll.Url = index
	return newscoll
}

// Crawl crawl sitemap containing multiple sitemaps
func (smapcoll *SitemapCollection) Crawl() {
	// Vars
	var wg sync.WaitGroup

	// get sitemap data, parse xml to sitemap struct
	data := getData(smapcoll.Url)
	parseXml(&data, smapcoll)

	// crawl all sitemap urls
	for i := range smapcoll.Sitemaps {
		log.Println(smapcoll.Sitemaps[i].Url)
		log.Println("main: starting worker ", i)
		wg.Add(1)
		newscoll := NewNewsCollection(smapcoll.Sitemaps[i].Url)
		go newscoll.crawlSync(&wg)
	}

	log.Println("main: waiting for workers to finish...")
	wg.Wait()
	log.Println("main: completed")
}

// filterKeywords filter crawled news
func (news *NewsCollection) filterKeywords() {
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
func (news *NewsCollection) Crawl(){
	log.Println("crawling " + news.Url)

	// get data, parse xml to news struct, filter news
	body := getData(news.Url)
	parseXml(&body, news)
	news.filterKeywords()

	// log crawled news
	for _, n := range news.NewsEntries {
		log.Println(n.Url)
	}
}

// crawlSync like NewsCollection.Crawl() but with Waitgroup for crawling multiple websites in parallel
func (news *NewsCollection) crawlSync(wg *sync.WaitGroup){
	defer wg.Done()

	log.Println("crawling " + news.Url)

	// get data from url, parse xml to news struct, filter news
	body := getData(news.Url)
	parseXml(&body, news)
	news.filterKeywords()

	// log crawled news
	for _, n := range news.NewsEntries {
		log.Println(n.Url)
	}
}
