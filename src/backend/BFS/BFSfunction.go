package bfs
import (
	"fmt"
	"strings"
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
	for i := 0; i < len(queue); i++ {
		if(A.Title == queue[i].Title && A.Parents[len(A.Parents)-1] == queue[i].Parents[len(queue[i].Parents)-1]){
			return true
		}
	}
	
	return false
}

func printAllQueue(queue []*Node){
	fmt.Println("Simpul hidup:")
	for  i := 0; i < len(queue); i++ {
		fmt.Println("Title:", queue[i].Title)
		fmt.Println("Parents:")
		for j := 0; j < len(queue[i].Parents); j++ {
			fmt.Printf("Title: %s, Parents: %s\n", queue[i].Title, strings.Join(queue[i].Parents, ", "))
		}
		fmt.Println()
	}
}