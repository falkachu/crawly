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

// Check for error message
func Check(err error) {
	if err != nil {
		log.Println(err)
	}
}

// Gunzip unzip ziped data from byte array
func Gunzip(zipdata io.Reader) []byte {
	log.Println("unzipping data...")
	zw, err := gzip.NewReader(zipdata)
	Check(err)

	unzipdata, err := ioutil.ReadAll(zw)
	Check(err)

	log.Println("data unziped")
	return unzipdata
}

// parseXML parse xml byte data into struct
func ParseXml(xmldata []byte, dest interface{}) {
	log.Println("parsing xml...")
	// Try to Unsmarshal XML to Struct Slice
	Check(xml.Unmarshal(xmldata, &dest))
	log.Println("xml parsed")
}

// GetData from URL and return the byte array
func GetData(url string) []byte {
	log.Println("getting data from url...")

	var body []byte

	resp, err := http.Get(url)
	Check(err)
	defer resp.Body.Close()

	if strings.Contains(url, ".gz") {
		log.Println("content encoded with gzip")
		body = Gunzip(resp.Body)
	} else {
		log.Println("content not encoded")
		body, err = ioutil.ReadAll(resp.Body)
		Check(err)
	}

	log.Println("data recieved")
	return body
}
