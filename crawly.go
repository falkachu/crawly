package crawly

import (
	"compress/gzip"
	"encoding/xml"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

// check for error message
func check(err error) {
	if err != nil {
		log.Println(err)
	}
}

// Gunzip unzip ziped data from byte array
func gunzip(zipdata io.Reader) []byte {
	log.Println("unzipping data...")
	zw, err := gzip.NewReader(zipdata)
	check(err)

	unzipdata, err := ioutil.ReadAll(zw)
	check(err)

	log.Println("data unziped")
	return unzipdata
}

// parseXML parse xml byte data into struct
func parseXml(xmldata *[]byte, dest interface{}) {
	log.Println("parsing xml...")
	// Try to Unsmarshal XML to Struct Slice
	check(xml.Unmarshal(*xmldata, &dest))
	log.Println("xml parsed")
}

// GetData from URL and return the byte array
func getData(url string) []byte {
	log.Println("getting data from url...")

	var body []byte

	resp, err := http.Get(url)
	check(err)
	defer resp.Body.Close()

	// check for gzip data, unzip if needed
	if strings.Contains(url, ".gz") {
		log.Println("content encoded with gzip")
		body = gunzip(resp.Body)
	} else {
		log.Println("content not encoded")
		body, err = ioutil.ReadAll(resp.Body)
		check(err)
	}

	log.Println("data received")
	return body
}

func checkKeywords(dest interface{}) {
	log.Println("checking keywords")

	switch tsrc := dest.(type) {

	case *NewsAA:
		var filterNews NewsAA

		log.Println("type identified as NewsAA")

		for _, n := range tsrc.News {
			if strings.Contains(strings.ToLower(n.URL), "gewerbegebiet") {
				filterNews.News = append(filterNews.News, n)
			}
		}
		*tsrc = filterNews

	default:
		log.Panic("unkown Type")
	}
}

func CrawlURL(wg *sync.WaitGroup, url string) {
	defer wg.Done()

	var augsburger NewsAA

	body := getData(url)
	parseXml(&body, &augsburger)
	checkKeywords(&augsburger)

	for _, n := range augsburger.News{
		log.Println(n.URL)
	}
}
