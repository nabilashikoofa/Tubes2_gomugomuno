package ids

import (
	"fmt"
	"strings"
	"Tubes2_gomugomuno/Scrape"
	"time"
)

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
		hasil_scrape := Scrape.Scraper(root.Value)
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

func IDS(startTitle string, endTitle string)  ([][]string, int64, int, int, int) {
	start := time.Now()
	isKetemu := false
	iterasi := 1
	startUrl := Scrape.Convert(startTitle)
	endUrl := Scrape.Convert(endTitle)
	root := createTree(startUrl)
	result := [][]string{}
	for !isKetemu {
		fmt.Println(iterasi)
		root.displayTree()
		result = DLS(root, endUrl, []string{}, iterasi)
		if len(result) != 0 {
			isKetemu = true
		} else {
			iterasi++
		}
		fmt.Println("hasil = ", result)
	}
	// root.displayTreeWithLevel(0);
	elapsed := time.Since(start).Milliseconds()
	fmt.Println(root.getSumAll())
	fmt.Println("Algorithm execution time:", elapsed, "ms")
	return result, elapsed, iterasi, root.getSumAll(), len(result) 
}


