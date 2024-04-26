package bfs
import (
	"fmt"
	"strings"
)


type Node struct {
	Title   string   // The title of the Wikipedia link
	Parents []string // Titles of parent nodes traversed to reach this node
}

// Node constructor, user defined
func createNode(startTitle string, NodeParent []string) *Node{
	return &Node{
		Title: startTitle,
		Parents: NodeParent,
	}
}

func printAllQueue(queue []*Node){
	fmt.Println("Simpul hidup:")
	for  i := 0; i < len(queue); i++ {
		fmt.Println("Title:", queue[i].Title)
		fmt.Println("Parents:")
		fmt.Print("- ")
		for j := 0; j < len(queue[i].Parents); j++ {
			fmt.Println(queue[i].Parents[j])
		}
		fmt.Println()
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