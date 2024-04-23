package bfs
import (
	"fmt"
	"strings"
	// "time"
)


type Node struct {
	Title   string   // The title of the Wikipedia link
	Parents []string // Titles of parent nodes traversed to reach this node
}

// create empty Node
func createEmptyNode(startTitle string) *Node{
	return &Node{
		Title: startTitle,
		Parents: nil,
	}
}

// Node constructor, user defined
func createNode(startTitle string, NodeParent []string) *Node{
	return &Node{
		Title: startTitle,
		Parents: NodeParent,
	}
}

// Decide if Node is visited before by check the title and the last queued Node at the Parents attribute
// Also terminate cycles conditions instantly
func isVisited(A *Node, queue []*Node) (bool){
	// fmt.Println("TES MASUK ISVISITED")
	for i := 0; i < len(queue); i++ {
		if (len(queue[i].Parents)!=0){
			if(A.Title == queue[i].Title && A.Parents[len(A.Parents)-1] == queue[i].Parents[len(queue[i].Parents)-1]){
				return true
			}
		}
	}
	
	return false
}

func printAllQueue(queue []*Node){
	fmt.Println("👽👽👽👽👽👽👽👽👽👽👽👽👽👽👽👽")
	fmt.Println("Simpul hidup:")
	for  i := 0; i < len(queue); i++ {
		fmt.Println("Title:", queue[i].Title)
		fmt.Println("Parents:")
		fmt.Print("- ")
		for j := 0; j < len(queue[i].Parents); j++ {
			fmt.Println(queue[i].Parents[j])
		}
		fmt.Println()
		// if (queue[i].Title=="Knowledge"){
		// 	fmt.Scanln()
		// }
	}
}

func (n *Node) Print() {
    fmt.Println("Title:", n.Title)
    fmt.Println("Parents:")
    if len(n.Parents) > 0 {
        fmt.Println(strings.Join(n.Parents, ", "))
    } else {
        fmt.Println("<none>")
    }
}