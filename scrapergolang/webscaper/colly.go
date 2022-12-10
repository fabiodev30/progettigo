package webscaper

import (
	"fmt"

	"github.com/gocolly/colly"
)

func Colly(url string) {
	c := colly.NewCollector(
		colly.AllowedDomains(url),
	)
	c.OnHTML(".mw-parser-output", func(e *colly.HTMLElement) {
		links := e.ChildAttrs("a", "href")
		for _, v := range links {
			fmt.Println(v)
		}
	})
}
