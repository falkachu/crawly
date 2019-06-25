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

var KEYWORDS = [...]string{"gewerbegebiet", "industriegebiet"}
var URLAA = "https://www.augsburger-allgemeine.de/news.xml"

// check for error message
func check(err error) {
	if err != nil {
		log.Println(err)
	}
}

// Gunzip unzip ziped data from byte array
func gUnzip(zipdata io.Reader) []byte {
	log.Println("unzipping data...")
	zw, err := gzip.NewReader(zipdata)
	check(err)

	unzipdata, err := ioutil.ReadAll(zw)
	if err != nil {
		log.Println("could not unzip data... return empty object")
		var emptydata = []byte("")
		return emptydata
	} else {
		log.Println("data unziped")
		return unzipdata
	}
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
		body = gUnzip(resp.Body)
	} else {
		log.Println("content not encoded")
		body, err = ioutil.ReadAll(resp.Body)
		check(err)
	}

	log.Println("data received")
	return body
}

func filterKeywords(dest interface{}) {
	log.Println("checking keywords...")

	switch tsrc := dest.(type) {

	case *NewsAA:
		var filterNews NewsAA

		log.Println("type identified as NewsAA")

		for _, n := range tsrc.News {
			if strings.Contains(strings.ToLower(n.URL), "usa") {
				filterNews.News = append(filterNews.News, n)
			}
		}

		if len(filterNews.News) == 0 {
			log.Println("no elements found")
		}

		*tsrc = filterNews

	default:
		log.Panic("unkown type")
	}
}

func Crawl(source string){
	switch source{

	case "augsburger":
		var newsAA NewsAA
		body := getData(URLAA)
		parseXml(&body, &newsAA)
		filterKeywords(&newsAA)

		for _, n := range newsAA.News {
			log.Println(n.Title)
		}
	}
}

func CrawlMultiURL(wg *sync.WaitGroup, url string) {
	defer wg.Done()

	var augsburger NewsAA

	body := getData(url)
	parseXml(&body, &augsburger)
	filterKeywords(&augsburger)

	for _, n := range augsburger.News {
		log.Println(n.URL)
	}
}

