package crawly

import (
	"compress/gzip"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type NewsEntryAA struct {
	Newspaper string   `xml:"news:name"`
	URL       string   `xml:"loc"`
	Date      string   `xml:"news:publication_date"`
	Title     string   `xml:"news:title"`
	Keywords  []string `xml:"news:keywords"`
}

type AllNewsAA struct {
	News []NewsEntryAA `xml:"url"`
}

func Hello() {
	fmt.Println("Hallo, ich bin eine Nachricht :)")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func XMLGUnzip(xmlzip io.Reader) []byte{
	zw, err := gzip.NewReader(xmlzip)
	checkError(err)

	xmlunzip, err := ioutil.ReadAll(zw)

	checkError(err)
	checkError(ioutil.WriteFile("news.xml", xmlunzip, 0777))

	return xmlunzip
}

func XMLtoObject(xmldata []byte) {
	var news AllNewsAA

	// Try to Unsmarshal XML to Struct Slice
	checkError(xml.Unmarshal(xmldata, &news))

	fmt.Printf("Es wurden %v Nachrichten gefunden\n", len(news.News))
}

func GetXMLData(url string) []byte{
	var body []byte

	resp, err := http.Get(url)
	checkError(err)
	defer resp.Body.Close()

	if strings.Contains(url, ".gz"){
		fmt.Printf("content encoded with gzip\n")
		body = XMLGUnzip(resp.Body)
	} else {
		fmt.Println("content not encoded")
		body, err = ioutil.ReadAll(resp.Body)
		checkError(err)
	}

	return body
}
