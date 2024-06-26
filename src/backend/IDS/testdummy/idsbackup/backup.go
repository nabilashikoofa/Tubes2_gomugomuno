package main

import (
	"fmt"
	"strings"
	"github.com/gocolly/colly"
	"time"
)

func main() {
	result,elapsed,shortestlength,numofcheckednodes := IDS("fanta cake", "dirty soda")
	fmt.Println("Algorithm execution time:", elapsed, "ms")
	fmt.Println(result)
    fmt.Println("Shortest length: ",shortestlength)
    fmt.Println("Nodes checked:",numofcheckednodes)
}

type Tree struct {
	Value string
	SubTree []*Tree
}

func createTree(startTitle string) *Tree{
	return &Tree{
		Value: startTitle,
		SubTree: nil,
	}
}

// Membuat anakan pohon
func (t *Tree) AddSubtree(value string) {
	child := &Tree{Value: value}
	t.SubTree = append(t.SubTree, child)
}

func (t *Tree) displayTree(){
	fmt.Println(t.Value)
	for _, sub := range t.SubTree {
		fmt.Print("   ")
		sub.displayTree()
	}
}

func (t *Tree) displayTreeWithLevel(level int){
	spasi := level
	for spasi > 0{
		fmt.Print("   ")
		spasi--;
	}
	fmt.Println(t.Value)
	for _, sub := range t.SubTree {
		sub.displayTreeWithLevel(level+1)
	}
}
func (t *Tree) getSumAll() int {
	sum := 1
	for _, sub := range t.SubTree {
		sumSub := sub.getSumAll()
		sum = sum + sumSub
	}
	return sum
}


// DFS tapi dilimit di level tertentu
func DLS(root *Tree, endUrl string, path []string, iterasi int) [][]string {
	fmt.Println(path)
	if strings.EqualFold(root.Value, endUrl) {
		fmt.Println("masuk atas", iterasi)
		return [][]string{append(path, root.Value)}
	} else if iterasi == 1 {
		hasil_scrape := Scraper(root.Value)
		if hasil_scrape == nil {
			fmt.Println("masuk nil")
			return [][]string{}
		} else {
			for _, sub := range hasil_scrape {
				root.AddSubtree(sub)
				fmt.Println("visited " + sub)
				if strings.EqualFold(sub, endUrl) {
					fmt.Println("masuk dalam", iterasi)
					paths := append(path,root.Value, sub)
					fmt.Println(paths)
					return [][]string{paths}
				}
			}
			// for _, url := range root.SubTree {
			// 	fmt.Println("asa " + url.Value)
			// 	if strings.EqualFold(url.Value, endUrl) {
			// 		fmt.Println("masuk dalam", iterasi)
			// 		paths := append(path,root.Value, url.Value)
			// 		fmt.Println(paths)
			// 		return [][]string{paths}
			// 	}
			// }
			return [][]string{}
		}
	} else {
		path = append(path, root.Value)
		var paths [][]string
		for _, child := range root.SubTree {
			fmt.Println(child.Value)
			childPaths := DLS(child, endUrl, path, iterasi-1)
			paths = append(paths, childPaths...)
			if len(paths) > 0 {
				break
			}
		}
		return paths
	}
}

func IDS(startTitle string, endTitle string)  ([][]string, int64, int, int) {
	start := time.Now()
	isKetemu := false
	iterasi := 1
	startUrl := Convert(startTitle)
	endUrl := Convert(endTitle)
	root := createTree(startUrl)
	result := [][]string{}
	for !isKetemu {
		fmt.Println(iterasi)
		root.displayTree()
		result = DLS(root, endUrl, []string{}, iterasi)
		if len(result) != 0 {
			isKetemu = true
		}
		iterasi++
		fmt.Println("hasil = ", result)
	}
	// root.displayTreeWithLevel(0);
	elapsed := time.Since(start).Milliseconds()
	fmt.Println(root.getSumAll())
	fmt.Println("Algorithm execution time:", elapsed, "ms")
	return result, elapsed, iterasi+1, root.getSumAll() 
}

func Scraper(title string) []string {
	c := colly.NewCollector()

	var links []string
	title = Convert(title)
	c.OnHTML("div.mw-page-container a[href^='/wiki/']", func(h *colly.HTMLElement) {
		link := h.Request.AbsoluteURL(h.Attr("href"))
		link = strings.TrimPrefix(link, "https://en.wikipedia.org/wiki/")

		if !strings.ContainsAny(link, ":/%") {
            if !Contains(links, link) && link != title {
                links = append(links, link)
            }
        }
	})

	c.Visit("https://en.wikipedia.org/wiki/" + title)

	return links
}

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


