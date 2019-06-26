package crawly

import (
	"log"
	"strings"
	"sync"
)

type NewsCollection struct {
	NewsEntries []NewsEntry `xml:"url"`
}

type NewsEntry struct {
	URL     string `xml:"loc"`
	Lastmod string `xml:"lastmod"`
}

type SitemapCollection struct {
	Index    string
	Sitemaps []Sitemap `xml:"sitemap"`
}

type Sitemap struct {
	URL     string `xml:"loc"`
	Lastmod string `xml:"lastmod"`
}

func NewSitemapCollection(index string) SitemapCollection {
	var newsitemaps SitemapCollection
	newsitemaps.Index = index
	return newsitemaps
}

func (smpacoll *SitemapCollection) Crawl() {
	// Vars
	var wg sync.WaitGroup

	// Crawling
	data := getData(smpacoll.Index)
	parseXml(&data, &smpacoll)

	for i := range smpacoll.Sitemaps {
		log.Println(smpacoll.Sitemaps[i].URL)
		log.Println("main: starting worker ", i)
		wg.Add(1)
		go crawlUrlSync(&wg, smpacoll.Sitemaps[i].URL)
	}

	log.Println("main: waiting for workers to finish...")
	wg.Wait()
	log.Println("main: completed")
}

func (news *NewsCollection) filterKeywords() {
	log.Println("filtering keywords...")

	var filterNews NewsCollection

	for _, n := range news.NewsEntries {
		for _, k := range KEYWORDS {
			if strings.Contains(strings.ToLower(n.URL), k) {
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
