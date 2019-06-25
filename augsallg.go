package crawly

type NewsAA struct {
	News []NewsEntryAA `xml:"url"`
	URL string
}

type NewsEntryAA struct {
	Newspaper string   `xml:"news>name"`
	URL       string   `xml:"loc"`
	Date      string   `xml:"news>publication_date"`
	Title     string   `xml:"news>title"`
	Keywords  []string `xml:"news>keywords"`
}

type SitemapsAA struct {
	Sitemaps []SitemapAA `xml:"sitemap"`
}

type SitemapAA struct {
	URL     string `xml:"loc"`
	Lastmod string `xml:"lastmod"`
}
