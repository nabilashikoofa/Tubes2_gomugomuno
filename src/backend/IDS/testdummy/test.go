package main

import (
	"fmt"
	// "bufio"
	// "os"
	"strings"
	"github.com/gocolly/colly"
	"time"
)

func main() {
	// start := time.Now()
	// var input string

	// fmt.Println("Masukkan Title Wikipedia:")

	// input, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	// input = strings.TrimSpace(input)

	// converted := Convert(input)
	// fmt.Println(converted)

	// links := Scraper(converted) 
	// d := 0
	// for k, i := range links{
	// 	fmt.Println(i)
	// 	d = k
	// }
	// fmt.Println(d)
	// elapsed := time.Since(start).Milliseconds()
	// fmt.Println("Algorithm execution time:", elapsed, "ms")
	IDS("vector", "vectorman")
}

type Tree struct {
	Value string
	// Visited bool
	SubTree []string
}

// Membuat pohon kosong
func createTree(startTitle string) *Tree{
	return &Tree{
		Value: startTitle,
		// Visited: false,
		SubTree: nil,
	}
}

// Membuat anakan pohon
// func (t *Tree) AddSubtree(value string) {
// 	child := &Tree{Value: value}
// 	t.SubTree = append(t.SubTree, child)
// }


// DFS tapi dilimit di level tertentu
func DLS(root *Tree, endTitle string, path []string, iterasi int) [][]string {
	root.SubTree = Scraper(root.Value)
	// pohon kosong
	if root == nil {
		return [][]string{}
    }
	
    path = append(path, root.Value)
	// jika ketemu tujuan
    if strings.EqualFold(root.Value, endTitle) {
        return [][]string{path}
    }

	// kalo belum ketemu tapi udah sampai batas
	if iterasi == 0 {
		return [][]string{}
	}

	// untuk seluruh subtree dilakukan hal yang sama
    var paths [][]string
    for _, subtree := range root.SubTree {
		child := createTree(subtree)
		fmt.Println(child.Value)
        childPaths := DLS(child, endTitle, path, iterasi-1)
        paths = append(paths, childPaths...)
    }
	// bakal return lintasan kalo ketemu, kalo gak return array kosong
    return paths
}

// belum cek buat scrapping asli
func IDS(startTitle string, endTitle string){
	start := time.Now()
	isKetemu := false
	// panjang iterasi menunjukkan panjang lintasan
	iterasi := 0
	startUrl := Convert(startTitle)
	endUrl := Convert(endTitle)
	root := createTree(startUrl)
	for !isKetemu {
		result := DLS(root, endUrl, []string{}, iterasi)
		// kalo ketemu link yang bener
		if len(result) != 0 {
			isKetemu = true
		}
		iterasi++
		fmt.Println(result)
	}
	elapsed := time.Since(start).Milliseconds()
	fmt.Println("Algorithm execution time:", elapsed, "ms")
}

// Input with underscore
// Output an Array
func Scraper(title string) []string {
	c := colly.NewCollector()

	var links []string

	c.OnHTML("div.mw-page-container a[href^='/wiki/']", func(h *colly.HTMLElement) {
		link := h.Request.AbsoluteURL(h.Attr("href"))
		link = strings.TrimPrefix(link, "https://en.wikipedia.org/wiki/")

		if !Contains(links, link) {
			links = append(links, link)
		}
	})

	c.Visit("https://en.wikipedia.org/wiki/" + title)

	return links
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


