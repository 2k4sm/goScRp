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

type authordet struct {
	Author      string `json:"author"`
	DOB         string `json:"dob"`
	Location    string `json:"location"`
	Description string `json:"description"`
}

type tags struct {
	Tag string `json:"tag"`
}

type quotes []quote

func scrapQuote(pathURL string) ([]byte, error) {
	var allquotes quotes

	c := colly.NewCollector(
		colly.AllowedDomains("quotes.toscrape.com"),
	)
	c.OnError(func(r *colly.Response, err error) {
		log.Printf("\nSomething went wrong: %v error Occured\n", err)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("\nMaking request to URL: %v\n", r.URL)
	})

	//Implements Quotes scrap logic.
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

	// c.OnScraped(func(r *colly.Response) {
	// 	out, _ := json.Marshal(allquotes)
	// 	fmt.Println(string(out))

	// })
	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("\nSuccessfully Visited: %v Status code: %v\n", r.Request.URL, r.StatusCode)
	})

	c.Visit("https://quotes.toscrape.com/" + pathURL)

	return json.Marshal(allquotes)
}

func ScrapAuthorDet(authorName string) ([]byte, error) {
	var detAuthor authordet

	c := colly.NewCollector(
		colly.AllowedDomains("quotes.toscrape.com"),
	)
	c.OnError(func(r *colly.Response, err error) {
		log.Printf("\nSomething went wrong: %v error Occured\n", err)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("\nMaking request to URL: %v\n", r.URL)
	})

	//Implements Author Details Scrap Logic.

	c.OnHTML(".author-details", func(h *colly.HTMLElement) {
		authorName := h.ChildText(".author-title")
		authorDesc := h.ChildText(".author-description")

		authorBdate := h.ChildText(".author-born-date")
		authorLocation := h.ChildText(".author-born-location")

		detAuthor = authordet{Author: authorName, Description: authorDesc, DOB: authorBdate, Location: authorLocation}

	})

	// c.OnScraped(func(r *colly.Response) {
	// 	out, _ := json.Marshal(detAuthor)
	// 	fmt.Println(string(out))
	// })

	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("\nSuccessfully Visited: %v Status code: %v\n", r.Request.URL, r.StatusCode)
	})

	c.Visit("https://quotes.toscrape.com/author/" + authorName)

	return json.Marshal(detAuthor)

}

func ScrapQuotePage(page string) []byte {
	pathURL := fmt.Sprintf("page/%s", page)
	payload, _ := scrapQuote(pathURL)
	return payload
}

func ScrapQuoteTag(tag string) []byte {
	pathURL := fmt.Sprintf("tag/%s", tag)
	payload, _ := scrapQuote(pathURL)
	return payload
}

func ScrapRandomQuote() []byte {
	pathURL := "random"
	payload, _ := scrapQuote(pathURL)

	return payload
}
