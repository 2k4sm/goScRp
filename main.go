package main

import (
	"fmt"
)

func main() {
	args := os.Args[1:]
	commands := []string{"-r","-t","-p","-a"}

	


}

func printGuideMsg() {
	pstr :=
		`quotesScRp is a tool to scrap quotes from (https://quotes.toscrape.com)

Usage:

	quotesScRp <command> <arguments>
	
The commands are:
	-r  --random   Scraps some random quote.
	-t  --tag      Scrap quotes using tag (ex: love,life,inspirational,humor).
	-p  --page     Scrap quotes of a specific page (Max : 10th page).
	-a  --authordet   Scrap details about a author.(ex: Albert-Einstein)
	`
	fmt.Println(pstr)
}

func printErrMsg(err string) {
	estr := fmt.Sprintf(`%s
Usage:

	quotesScRp <command> <arguments>
	
The commands are:
	-r  --random   Scraps some random quote.
	-t  --tag      Scrap quotes using tag (ex: love,life,inspirational,humor).
	-p  --page     Scrap quotes of a specific page (Max : 10th page).
	-a  --authordet   Scrap details about a author.(ex: Albert-Einstein)
	
	`, err)

	fmt.Println(estr)
}
