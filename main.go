package main

import (
	"fmt"
	"os"

	"github.com/2k4sm/goScRp/scrapper"
)

func main() {
	// Take the os params..
	// Guide the user on any error.
	// After parsing the params then create the querry string using the params provided.
	// Perform the functions.

	params := os.Args[1:]
	// error strings...

	errstr := "quotesScRp Unknown Command:\nSee 'quotesScRp help'"

	if len(params) < 1 {
		randomQuote()
		os.Exit(0)
	}

	if params[0] == "help" {
		helpstr := "quotesScRp is a TUI Program that scrapes quotes from https://quotes.toscrape.com \n\nUsage:\n\t-t <tagname> Generates quotes using tag \n\t-p <pagenumber> generates quotes from the specified page \n\t-a <authorname> Prints author Information\n\t_blank Generate a random quote."
		fmt.Println(helpstr)
		os.Exit(0)
	}

	if len(params)%2 != 0 || len(params) > 4 {
		fmt.Println(errstr)
		os.Exit(1)
	}

	var tag string = ""
	var page string = ""
	var authname string = ""
	for i, j := range params {
		if j == "-t" {
			tag = params[i+1]

		}
		if j == "-p" {
			page = params[i+1]
		}

		if j == "-a" {
			authname = params[i+1]
		}
	}

	if tag != "" && page != "" {
		tagpageQuote(tag, page)
		os.Exit(0)
	}
	if tag != "" {
		tagQuote(tag)
	}

	if page != "" {
		pageQuote(page)
	}
	if authname != "" {
		authPrint(authname)
		os.Exit(0)
	}

}

func randomQuote() {
	fmt.Println("Generating random quote...")

	Payload := scrapper.ScrapRandomQuote()
	data := Payload[0]
	Quote := data.Text
	Author := data.Author
	Tags := data.Tags

	fmt.Printf(Quote+"\nQuoted By:%s\n\n", Author)
	fmt.Println("Quotes tagged as", Tags)
}

func tagpageQuote(tag string, page string) {
	fmt.Println("Generating quote based on tag and page...")

	Payload := scrapper.ScrapQuoteTag(tag + "/page/" + page)
	for i, data := range Payload {
		Quote := data.Text
		Author := data.Author

		fmt.Printf("Quote:%d\t", i+1)
		fmt.Printf(Quote+"\nQuoted By:%s\n\n", Author)
	}
}

func tagQuote(tag string) {
	Payload := scrapper.ScrapQuoteTag(tag)
	for i, data := range Payload {
		Quote := data.Text
		Author := data.Author

		fmt.Printf("Quote:%d ", i+1)
		fmt.Printf(Quote+"\nQuoted By:%s\n\n", Author)
	}
}

func pageQuote(page string) {
	Payload := scrapper.ScrapQuotePage(page)
	for i, data := range Payload {
		Quote := data.Text
		Author := data.Author

		fmt.Printf("Quote:%d ", i+1)
		fmt.Printf(Quote+"\nQuoted By:%s\n\n", Author)
	}
}

func authPrint(authname string) {
	Payload := scrapper.ScrapAuthorDet(authname)

	fmt.Printf("\nAuthor Name: %s\n",Payload.Author)
	fmt.Printf("D.O.B: %s\n", Payload.DOB)
	fmt.Printf("P.O.B: %s\n\n", Payload.Location)
	fmt.Printf("Description: %s\n\n", Payload.Description)

}
