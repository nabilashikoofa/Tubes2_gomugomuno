package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"github.com/gocolly/colly"
)

func main() {
	var input string

	fmt.Println("Masukkan Title Wikipedia:")

	input, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	input = strings.TrimSpace(input)

	converted := Convert(input)
	fmt.Println(converted)

	links := Scraper(converted) 
	sort.Strings(links)
	fmt.Println("")
	fmt.Println(len(links))
	printStrings((links))

}

func printStrings(slice []string) {
	for _, s := range slice {
		fmt.Println(s)
	}
}

// Input with underscore
// Output an Array
// func Scraper(title string) []string {
// 	start := time.Now()
// 	c := colly.NewCollector()
// 	var links []string
// 	title = Convert(title)
// 	c.OnHTML("div.mw-page-container a[href^='/wiki/']", func(h *colly.HTMLElement) {
// 		link := h.Request.AbsoluteURL(h.Attr("href"))
// 		link = strings.TrimPrefix(link, "https://en.wikipedia.org/wiki/")

// 		if !strings.ContainsAny(link, ":/%") {
//             if !Contains(links, link) && link != Convert(title) {
//                 links = append(links, link)
//             }
//         }
// 	})

// 	c.Visit("https://en.wikipedia.org/wiki/" + title)
// 	elapsed := time.Since(start).Milliseconds()
// 	fmt.Print("--------------------------The amount of links visited: ")
// 	fmt.Println(len(links))
// 	fmt.Println(elapsed)
// 	return links
// }


func Scraper(title string) []string {
	c := colly.NewCollector()

    linksMap := make(map[string]bool) // Map to store links

    title = Convert(title)
    c.OnHTML("div.mw-page-container a[href^='/wiki/']", func(h *colly.HTMLElement) {
        link := h.Request.AbsoluteURL(h.Attr("href"))
        link = strings.TrimPrefix(link, "https://en.wikipedia.org/wiki/")

        if !strings.ContainsAny(link, ":/%") && link != Convert(title) {
            linksMap[link] = true // Store link in map
        }
    })

    c.Visit("https://en.wikipedia.org/wiki/" + title)
    var result []string
    for link := range linksMap {
        result = append(result, link)
    }
    return result
}

// Input without underscore
// Output with underscore 
func Convert(input string) string {
	converted := strings.ReplaceAll(input, " ", "_")
	return converted
}

