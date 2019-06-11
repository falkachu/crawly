package crawly

import (
	"log"
	"testing"
)

func TestCrawly(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	url := "https://www.augsburger-allgemeine.de/news.xml"
	var news AllNewsAA

	parseXml(getData(url), &news)
}