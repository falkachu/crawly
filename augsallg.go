package crawly

type AllNewsAA struct {
	News []NewsEntryAA `xml:"url"`
}

type NewsEntryAA struct {
	Newspaper string   `xml:"news>name"`
	URL       string   `xml:"loc"`
	Date      string   `xml:"news>publication_date"`
	Title     string   `xml:"news>title"`
	Keywords  []string `xml:"news>keywords"`
}



