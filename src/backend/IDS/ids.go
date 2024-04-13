package ids

import (
	"fmt"
	"time"
)

// Struct Pohon
type Tree struct {
	Value string
	Visited bool
	SubTree []*Tree
}

// Membuat pohon kosong
func createTree(startTitle string) *Tree{
	return &Tree{
		Value: startTitle,
		Visited: false,
		SubTree: nil,
	}
}

// Membuat anakan pohon
func (t *Tree) AddSubtree(value string) {
	child := &Tree{Value: value}
	t.SubTree = append(t.SubTree, child)
}


// DFS tapi dilimit di level tertentu
func DLS(root *Tree, endTitle string, path []string, iterasi int) [][]string {
	// pohon kosong
	if root == nil {
		return [][]string{}
    }
	
    path = append(path, root.Value)
	// jika ketemu tujuan
    if root.Value == endTitle{
        return [][]string{path}
    }

	// kalo belum ketemu tapi udah sampai batas
	if iterasi == 0 {
		return [][]string{}
	}

	// untuk seluruh subtree dilakukan hal yang sama
    var paths [][]string
    for _, child := range root.SubTree {
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
	root := createTree(startTitle)
	for !isKetemu {
		result := DLS(root, endTitle, []string{}, iterasi)
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