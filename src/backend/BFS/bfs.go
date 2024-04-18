package bfs

import (
	"fmt"
	// "time"
	"strings"
	// "Tubes2_gomugomuno/Scrape"
	"github.com/gocolly/colly"
)

// ini temporary ak pindahin soalnya mau eksperimen
func scraper(title string) []string {
	c := colly.NewCollector()

	var links []string
	var res []*Node
	c.OnHTML("div.mw-page-container a[href^='/wiki/']", func(h *colly.HTMLElement) {
		link := h.Request.AbsoluteURL(h.Attr("href"))
		link = strings.TrimPrefix(link, "https://en.wikipedia.org/wiki/")

		if !contains(links, link) {
			links = append(links, link)
			res = append(res,createNode(link,nil))
		}
	})

	c.Visit("https://en.wikipedia.org/wiki/" + title)

	return links
}

// Input without underscore
// Output with underscore 
func convert(input string) string {
	converted := strings.ReplaceAll(input, " ", "_")
	return converted
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}


func BFS(startTitle string, endNode string) {
// func BFS(startTitle string, endNode string) ([][]string, int64, int, int){
	// start := time.Now()
	firstNode := createNode(startTitle,nil)	// node pertama
	// shortestlength:= 0
	// numofcheckednodes := 0
	var queue []*Node		// simpul hidup
	// var result [][]string		// list of answers
	// var visitedNodes []*Node // list of visitedNodes, including the ones that had been dequeued 
	queue = append(queue, firstNode)	// masukkan node pertama ke dalam queue simpul hidup
	temptitle := scraper(startTitle)
	for _,title := range temptitle{
		var tempparent []string = append(firstNode.Parents, startTitle) 
		queue = append(queue, createNode(title,tempparent))
	}
	fmt.Println("Queue initial:")
	printAllQueue(queue)


	
	// // parse for startNode and enqueue
	// for len(queue) > 0 {		//while queue is not empty
	// 	numofcheckednodes++
	// 	visitedNodes = append(visitedNodes, queue[0])	// mark as visited
	// 	currentNode := queue[0] //current branch is the start of the queue, dequeue
	// 	queue = queue[1:]
	// 	// currentNode.Parents = append(currentNode.Parents,currentNode.Title)

	// 	// grab only the shortest solution or when we heave no solution at all.
	// 	if (currentNode.Title == endNode){
	// 		fmt.Println("Destination Reached!")
	// 		if (len(currentNode.Parents)==0){	//assign when first time encountering solution
	// 			shortestlength = len(currentNode.Parents)
	// 		} else if (len(currentNode.Parents)==shortestlength){
	// 			currentNode.Parents = append(currentNode.Parents, currentNode.Title)		// also add the endNode to the list of results, so if we have Basf with B as endNode, we get {a,s,f,b}
	// 			result = append(result, currentNode.Parents)
	// 		}
	// 	} else{
	// 		// parse and enqueue/create node
	// 		// dont forget to change the newly made Node Parents with the copy of currentNode.Parents + currentNode.Title
	// 		var temptitle []string = scrape(currentNode.Title)	// get a list of Node titles from parsing currentNode
	// 		var tempparent []string = currentNode.Parents		// get the list of parent Node from the currentNode
	// 		tempparent = append(tempparent, currentNode.Title)	// add the currentNode into the list of parent Node
	// 		for i := 0; i < len(temptitle); i++ {							// append the newly made node into the queue
	// 			A := createNode(temptitle[i],tempparent)
	// 			if (!isVisited(A,visitedNodes)){
	// 				queue = append(queue, A)
	// 				visitedNodes = append(visitedNodes, A)		// mark newlymade Nodes as visited
	// 			}
	// 		}
	// 	}
	// }
	// elapsed := time.Since(start).Milliseconds()
	// fmt.Println("Algorithm execution time:", elapsed, "ms")
	// return result,elapsed,shortestlength,numofcheckednodes
	// // should i return array instead...?
}