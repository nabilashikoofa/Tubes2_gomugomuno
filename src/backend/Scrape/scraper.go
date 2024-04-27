package Scrape

import (
	"strings"
    
	"github.com/gocolly/colly"
)

func Scraper(title string) []string {
	c := colly.NewCollector()

    linksMap := make(map[string]bool) // Map to store links
    result := []string{}
    title = Convert(title)
    c.OnHTML("div.mw-page-container a[href^='/wiki/']", func(h *colly.HTMLElement) {
        link := h.Request.AbsoluteURL(h.Attr("href"))
        link = strings.TrimPrefix(link, "https://en.wikipedia.org/wiki/")

        if !strings.ContainsAny(link, ":/%") && link != Convert(title) && !linksMap[link]{
            linksMap[link] = true // Store link in map
            result = append(result, link)
        }
    })

    c.Visit("https://en.wikipedia.org/wiki/" + title)
    return result
}

// Input without underscore
// Output with underscore 
func Convert(input string) string {
	converted := strings.ReplaceAll(input, " ", "_")
	return converted
}



