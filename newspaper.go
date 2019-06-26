package crawly

import (
	"log"
	"strings"
	"sync"
)

type News struct {
	NewsEntries []NewsEntry `xml:"url"`
}

type NewsEntry struct {
	URL     string `xml:"loc"`
	Lastmod string `xml:"lastmod"`
}

type Sitemaps struct {
	Index    string
	Sitemaps []Sitemap `xml:"sitemap"`
}

type Sitemap struct {
	URL     string `xml:"loc"`
	Lastmod string `xml:"lastmod"`
}

func NewSitemaps(index string) Sitemaps {
	var newsitemaps Sitemaps
	newsitemaps.Index = index
	return newsitemaps
}

func (smap *Sitemap) Crawl(wg *sync.WaitGroup) {
	defer wg.Done()

	log.Println("crawling " + smap.URL)

	var news News

	body := getData(smap.URL)
	parseXml(&body, &news)
	news.filterKeywords()

	for _, n := range news.NewsEntries {
		log.Println(n.URL)
	}
}

func (smaps *Sitemaps) Crawl() {
	// Vars
	var wg sync.WaitGroup

	// Crawling
	data := getData(smaps.Index)
	parseXml(&data, &smaps)

	for i := range smaps.Sitemaps {
		log.Println(smaps.Sitemaps[i].URL)
		log.Println("main: starting worker ", i)
		wg.Add(1)
		go smaps.Sitemaps[i].Crawl(&wg)
	}

	log.Println("main: waiting for workers to finish...")
	wg.Wait()
	log.Println("main: completed")
}

func (news *News) filterKeywords() {
	log.Println("filtering keywords...")

	var filterNews News

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
