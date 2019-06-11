package crawly

import (
	"compress/gzip"
	"encoding/xml"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// check for error message
func check(err error) {
	if err != nil {
		log.Println(err)
	}
}

// gUnzip unzip ziped data from byte array
func gUnzip(zipdata io.Reader) []byte{
	log.Println("unzipping data...")
	zw, err := gzip.NewReader(zipdata)
	check(err)

	unzipdata, err := ioutil.ReadAll(zw)
	check(err)

	log.Println("data unziped")
	return unzipdata
}

// parseXML parse xml byte data into struct
func parseXml(xmldata []byte, dest interface{}){
	log.Println("parsing xml...")
	// Try to Unsmarshal XML to Struct Slice
	check(xml.Unmarshal(xmldata, &dest))
	log.Println("xml parsed")
}

// getData from URL and return the byte array
func getData(url string) []byte{
	log.Println("getting data from url...")

	var body []byte

	resp, err := http.Get(url)
	check(err)
	defer resp.Body.Close()

	if strings.Contains(url, ".gz"){
		log.Println("content encoded with gzip")
		body = gUnzip(resp.Body)
	} else {
		log.Println("content not encoded")
		body, err = ioutil.ReadAll(resp.Body)
		check(err)
	}

	log.Println("data recieved")
	return body
}
