package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

func main() {
	var input string

	fmt.Println("Masukkan URL website:")

	fmt.Scanln(&input) // Kim_Soo-hyun
	url := "https://en.wikipedia.org/wiki/" + input
	
	scraper(url)
}

func scraper(website string)	{
	c := colly.NewCollector()

	var links []string

	//Ubah class dari inspect
	c.OnHTML("div.mw-page-container a[href^='/wiki/']", func(h *colly.HTMLElement) {
		link := h.Request.AbsoluteURL(h.Attr("href"))
		link = strings.TrimPrefix(link, "https://en.wikipedia.org/wiki/")

		if !contains(links, link) {
			links = append(links, link)
		}
		// yg ini buat masuk ke link baru
		// c.Visit(h.Request.AbsoluteURL(h.Attr("href")))
	})

	//Ubah link webnya disini
	c.Visit(website)

	fmt.Println(links)
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}


