package crawly

import "testing"

func TestCrawly(t *testing.T) {
	url := "https://www.augsburger-allgemeine.de/sitemap/sitemap-2019-06-05-p00.xml.gz"
	parseXml(getData(url))
}