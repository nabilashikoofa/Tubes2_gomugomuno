package bfs


import (
	"fmt"
	"time"
)


func BFS(startTitle string, endNode string) ([][]string, int64, int, int){
	start := time.Now()
	firstNode := createNode(startTitle,nil)	// node pertama
	shortestlength:= 0
	numofcheckednodes := 0
	var queue []*Node		// simpul hidup
	var result [][]string		// list of answers
	var visitedNodes []*Node // list of visitedNodes, including the ones that had been dequeued 
	queue = append(queue, firstNode)	// masukkan node pertama ke dalam queue simpul hidup
	// parse from start node
	// parse for startNode and enqueue
	for len(queue) > 0 {		//while queue is not empty
		numofcheckednodes++
		visitedNodes = append(visitedNodes, queue[0])	// mark as visited
		currentNode := queue[0] //current branch is the start of the queue, dequeue
		queue = queue[1:]
		// currentNode.Parents = append(currentNode.Parents,currentNode.Title)

		// grab only the shortest solution or when we heave no solution at all.
		if (currentNode.Title == endNode){
			fmt.Println("Destination Reached!")
			if (len(currentNode.Parents)==0){	//assign when first time encountering solution
				shortestlength = len(currentNode.Parents)
			} else if (len(currentNode.Parents)==shortestlength){
				currentNode.Parents = append(currentNode.Parents, currentNode.Title)		// also add the endNode to the list of results, so if we have Basf with B as endNode, we get {a,s,f,b}
				result = append(result, currentNode.Parents)
			}
		} else{
			// parse and enqueue/create node
			// dont forget to change the newly made Node Parents with the copy of currentNode.Parents + currentNode.Title
			var temptitle []string = scrape(currentNode.Title)	// get a list of Node titles from parsing currentNode
			var tempparent []string = currentNode.Parents		// get the list of parent Node from the currentNode
			tempparent = append(tempparent, currentNode.Title)	// add the currentNode into the list of parent Node
			for i := 0; i < len(temptitle); i++ {							// append the newly made node into the queue
				A := createNode(temptitle[i],tempparent)
				if (!isVisited(A,visitedNodes)){
					queue = append(queue, A)
					visitedNodes = append(visitedNodes, A)		// mark newlymade Nodes as visited
				}
			}
		}
	}
	elapsed := time.Since(start).Milliseconds()
	fmt.Println("Algorithm execution time:", elapsed, "ms")
	return result,elapsed,shortestlength,numofcheckednodes
	// should i return array instead...?
}

// temporary parse
func scrape(title string) []string {
	// Some parsing logic
	return []string{"ParsedTitle1", "ParsedTitle2"} // Just for illustration
}