package main

import (
	"scrapergolang/webscaper"
)

func main() {
	// webscaper.GoQuery("en.wikipedia.org")
	webscaper.Colly("en.wikipedia.org")
}
