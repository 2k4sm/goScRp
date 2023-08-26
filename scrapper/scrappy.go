package scrapper

import (
	"fmt"

	"github.com/gocolly/colly"
)

type quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
	//Tags   Tags `json:"tags"`
}

func scrapper(url string) {
	quotes := &quote{}
	c := colly.NewCollector(
		colly.AllowedDomains("quotes.toscrape.com"),
	)

	c.OnHTML(".text", func(e *colly.HTMLElement) {
		quotes.Text = e.Text
	})
	c.OnHTML(".author", func(e *colly.HTMLElement) {
		quotes.Author = e.Text
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("visited %v", r.Request.URL)
	})
	// c.OnHTML(".tags",func(h *colly.HTMLElement) {
	// 	tags := Tags
	// })
	c.Visit("https://quotes.toscrape.com/")

}
