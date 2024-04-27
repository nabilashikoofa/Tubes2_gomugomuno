package ids

import (
	"fmt"
	"strings"
	"Tubes2_gomugomuno/Scrape"
	"sync"
	"time"
	"context"
)

// Struktur data Pohon
type Tree struct {
	Value string
	SubTree []*Tree
}
// Membuat Pohon tanpa anak
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
// Mencari banyak pohon 
func (t *Tree) getSumAll() int {
	sum := 1
	for _, sub := range t.SubTree {
		sumSub := sub.getSumAll()
		sum = sum + sumSub
	}
	return sum
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

func IDSParalel(startTitle string, endTitle string) ([][]string, int64, int, int, int) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
    defer cancel() // Cancel context to release resources when done

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
	// visited := make(map[string]bool)
	path := []string{startTitle}
	// visited[startTitle] = true
	sisa := numThreads
	resultavailable := &result
	fmt.Println(numThreads)
	for _, child := range childRoot {
		root.AddSubtree(child)
		// visited[child] = true
		if strings.EqualFold(child, endTitle) {
			*resultavailable = append(*resultavailable, []string{startTitle,child})
			break
		}
	}
	if numThreads > maxNumThread {
		pengali = numThreads / maxNumThread
		sisa = numThreads % maxNumThread
	}
	iterasi := 1
	for len(*resultavailable) == 0 {
		if time.Since(start).Minutes() > 5{
			break
		}
		fmt.Println("Iterasi : ",iterasi)
		for x := 0; x < pengali; x++ {
			for i := 0; i < maxNumThread; i++ {
				wg.Add(1)
				go func(akar *Tree) {
					defer wg.Done()
					DLSParalel(akar, endTitle, path, iterasi, resultavailable, ctx)
				}(root.SubTree[x*maxNumThread + i])
			}
			wg.Wait()
		}
		for r := 0; r < sisa; r++ {
			wg.Add(1)
				go func(akar *Tree) {
					defer wg.Done()
					DLSParalel(akar, endTitle, path, iterasi, resultavailable, ctx)
				}(root.SubTree[pengali*maxNumThread + r])	
		}
		wg.Wait()
		iterasi++
	}
	// root.displayTreeWithLevel(0)
	// fmt.Println(len(visited))
	fmt.Println(root.getSumAll())
	elapsed := time.Since(start).Milliseconds()
	return *resultavailable, elapsed, iterasi, root.getSumAll(), len(result)
}

func DLSParalel(root *Tree, endUrl string, path []string, iterasi int, resultavailable *[][]string, ctx context.Context) {	
	select{
	case <- ctx.Done():
		return
	default:
		if len(*resultavailable) > 0 {
			return
		}		
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
				DLSParalel(child, endUrl, path, iterasi-1, resultavailable, ctx)
				if len(*resultavailable) > 0 {
					break
				}
			}
			return
		}
	}
}
