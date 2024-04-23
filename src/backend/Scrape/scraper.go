package Scrape

import (
	"github.com/gocolly/colly"
	"strings"
)

// func main() {
// 	var input string

// 	fmt.Println("Masukkan Title Wikipedia:")

// 	input, _ = bufio.NewReader(os.Stdin).ReadString('\n')
// 	input = strings.TrimSpace(input)

// 	converted := convert(input)
// 	fmt.Println(converted)

// 	links := scraper(converted) 
// 	fmt.Println(links)
// }


// Input with underscore
// Output an Array
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

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}


