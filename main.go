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
	errstr := "quotesScRp Unknown Command:\nUsage:\n\t-t <tagname> Generates quotes using tag \n\t-p <pagenumber> generates quotes from the specified page \n\t-a <authorname> Prints author Information\n\t_blank Generate a random quote."
	if len(params) < 1 {
		randomQuote()
	}

	if len(params)%2 != 0 || len(params) > 4 {
		fmt.Println(errstr)
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
	}
	if tag != "" {
		tagQuote(tag)
	}

	if page != "" {
		pageQuote(page)
	}
	if authname != "" {
		authPrint(authname)
	}

}

func randomQuote() {
	fmt.Println("Generating random quote...")
	//format this later ...
	fmt.Println(scrapper.ScrapRandomQuote())
}

func tagpageQuote(tag string, page string) {
	fmt.Println("Generating quote based on tag and page...")
	fmt.Println(scrapper.ScrapQuoteTag(tag + "/page/" + page))
}

func tagQuote(tag string) {
	fmt.Printf("Generating quote based on tag %s ...", tag)
	//format this later...
	fmt.Println(scrapper.ScrapQuoteTag(tag))
}

func pageQuote(page string) {
	fmt.Printf("Generating quote based on page %s...", page)
	fmt.Println(scrapper.ScrapQuotePage(page))
}

func authPrint(authname string) {
	fmt.Printf("Generating Author details based on authorname %s....", authname)
	fmt.Println(scrapper.ScrapAuthorDet(authname))
}
