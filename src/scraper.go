package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main()	{
	c := colly.NewCollector()

	//Ubah class dari inspect
	c.OnHTML("div.mw-page-container a[href^='/wiki/']", func(h *colly.HTMLElement) {
		c.Visit(h.Request.AbsoluteURL(h.Attr("href")))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL)
	})

	//Ubah link webnya disini
	c.Visit("https://id.wikipedia.org/wiki/Institut_Teknologi_Bandung")
	// c.Visit("https://id.wikipedia.org/wiki/Universitas_Indonesia")
}
