package ids

import (
	"fmt"
	"strings"
	"Tubes2_gomugomuno/Scrape"
	"sync"
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

func IDSParalel(startTitle string, endTitle string) ([][]string, int64, int, int, int) {
	start := time.Now()
	var result [][] string 
	var wg sync.WaitGroup
	maxNumThread := 300
	pengali := 0
	startTitle = Scrape.Convert(startTitle)
	root := createTree(startTitle)
	childRoot := Scrape.Scraper(startTitle)
	numThreads := len(childRoot)
	endTitle = Scrape.Convert(endTitle)
	path := []string{startTitle}
	sisa := numThreads
	resultavailable := &result
	fmt.Println(numThreads)
	for _, child := range childRoot {
		root.AddSubtree(child)
		if strings.EqualFold(child, endTitle) {
			*resultavailable = append(*resultavailable, []string{startTitle,child})
		}
	}
	if numThreads > maxNumThread {
		pengali = numThreads / maxNumThread
		sisa = numThreads % maxNumThread
	}
	iterasi := 1
	for len(*resultavailable) == 0 {
		fmt.Println("Iterasi : ",iterasi)
		for x := 0; x < pengali; x++ {
			for i := 0; i < maxNumThread; i++ {
				wg.Add(1)
				go func(akar *Tree) {
					defer wg.Done()
					DLSParalel(akar, endTitle, path, iterasi, resultavailable)
				}(root.SubTree[x*maxNumThread + i])
			}
			wg.Wait()
		}
		for r := 0; r < sisa; r++ {
			wg.Add(1)
				go func(akar *Tree) {
					defer wg.Done()
					DLSParalel(akar, endTitle, path, iterasi, resultavailable)
				}(root.SubTree[pengali*maxNumThread + r])	
		}
		wg.Wait()
		iterasi++
	}
	elapsed := time.Since(start).Milliseconds()
	return *resultavailable, elapsed, iterasi, root.getSumAll(), len(result)
}

func DLSParalel(root *Tree, endUrl string, path []string, iterasi int, resultavailable *[][]string) {			
	if iterasi == 1 {
		hasil_scrape := Scrape.Scraper(root.Value)
		if hasil_scrape == nil {
			fmt.Println("gaada anak")
			return
		} else {
			for _, sub := range hasil_scrape {
				root.AddSubtree(sub)
				// fmt.Println("visited " + sub)
				if strings.EqualFold(sub, endUrl) {
					fmt.Println("Ketemuuu pathnya", iterasi)
					path := append(path,root.Value, sub)
					fmt.Println(path)
					// *resultavailable = append(*resultavailable, path)
					*resultavailable = [][]string{path} 
					return 
				}					
			} 
			return
		}
	} else {
		path = append(path, root.Value)
		for _, child := range root.SubTree {
			// fmt.Println(child.Value)
			DLSParalel(child, endUrl, path, iterasi-1, resultavailable)
			if len(*resultavailable) > 0 {
				break
			}
		}
		return
	}
}
