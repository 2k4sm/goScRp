package scrapper

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
	Tags   []tags `json:"tags"`
}
type tags struct {
	Tag string `json:"tag"`
}

type quotes []quote

func Scrapper() ([]byte, error) {
	var allquotes quotes

	c := colly.NewCollector(
		colly.AllowedDomains("quotes.toscrape.com"),
	)
	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Something went wrong: %v error Occured", err)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Making request to URL: %v", r.URL)
	})

	c.OnHTML(".quote", func(e *colly.HTMLElement) {
		text := e.ChildText(".text")
		author := e.ChildText(".author")
		var atags []tags
		e.ForEach(".tag", func(i int, h *colly.HTMLElement) {
			tag := h.Text
			atags = append(atags, tags{tag})

		})
		allquotes = append(allquotes, quote{Text: text, Author: author, Tags: atags})
	})
	c.OnScraped(func(r *colly.Response) {
		out, _ := json.Marshal(allquotes)
		fmt.Println(string(out))
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("Successfully Visited: %v Status code: %v", r.Request.URL, r.StatusCode)
	})

	c.Visit("https://quotes.toscrape.com/")
	
	return json.Marshal(allquotes)
}
