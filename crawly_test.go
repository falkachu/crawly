package crawly

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestCrawly(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	var smapsAA SitemapsAA
	var newsAA NewsAA
	var filterNews NewsAA

	ParseXml(GetData("https://www.augsburger-allgemeine.de/sitemap.xml"), &smapsAA)

	for _, sm := range smapsAA.Sitemaps{
		ParseXml(GetData(sm.URL), &newsAA)
		for _, n := range newsAA.News {
			if strings.Contains(strings.ToLower(n.URL), "gewerbegebiet") {
				filterNews.News = append(filterNews.News, n)
			}
		}
	}

	for _, n := range filterNews.News {
		fmt.Println(n.URL)
	}
}
