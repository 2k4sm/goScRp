package scrapper

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type quote struct {
	Text   string
	Author string
	Tags   []string
}

type authordet struct {
	Author      string
	DOB         string
	Location    string
	Description string
}

// type tags struct {
// 	Tag string
// }

type quotes []quote

func scrapQuote(pathURL string) quotes {
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
		var atags []string
		e.ForEach(".tag", func(i int, h *colly.HTMLElement) {
			tag := h.Text
			atags = append(atags, tag)

		})
		allquotes = append(allquotes, quote{Text: text, Author: author, Tags: atags})
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("\nSuccessfully Visited: %v Status code: %v\n\n", r.Request.URL, r.StatusCode)
	})

	c.Visit("https://quotes.toscrape.com/" + pathURL)

	return allquotes
}

func ScrapAuthorDet(authorName string) authordet {
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

	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("\nSuccessfully Visited: %v Status code: %v\n", r.Request.URL, r.StatusCode)
	})

	c.Visit("https://quotes.toscrape.com/author/" + authorName)

	return detAuthor

}

func ScrapQuotePage(page string) quotes {
	pathURL := fmt.Sprintf("page/%s", page)
	payload := scrapQuote(pathURL)
	return payload
}

func ScrapQuoteTag(tag string) quotes {
	pathURL := fmt.Sprintf("tag/%s", tag)
	payload := scrapQuote(pathURL)
	return payload
}

func ScrapRandomQuote() quotes {
	pathURL := "random"
	payload := scrapQuote(pathURL)

	return payload
}
