package crawly

import (
	"compress/gzip"
	"encoding/xml"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var KEYWORDS = [...]string{"gewerbegebiet", "industriegebiet", "investition", "investiert"}

// check for error message
func check(err error) {
	if err != nil {
		log.Println(err)
	}
}

// Gunzip unzip data from byte array
func GUnzip(zipdata io.Reader) []byte {
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
func ParseXML(xmldata *[]byte, dest interface{}) {
	log.Println("parsing xml...")
	// Try to Unsmarshal XML to Struct Slice
	check(xml.Unmarshal(*xmldata, &dest))
	log.Println("xml parsed")
}

// GetData get data from url and return the byte array
func GetData(url string) []byte {
	log.Println("getting data from URL: ", url)

	var body []byte

	// define client with timeout of 10 seconds
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := netClient.Get(url)
	defer resp.Body.Close()
	if err != nil{
		log.Println(err)
		return []byte("")
	}

	// check for gzip data, unzip if needed
	if strings.Contains(url, ".gz") {
		log.Println("content encoded with gzip")
		body = GUnzip(resp.Body)
	} else {
		log.Println("content not encoded")
		body, err = ioutil.ReadAll(resp.Body)
		check(err)
	}

	log.Println("data received")
	return body
}



