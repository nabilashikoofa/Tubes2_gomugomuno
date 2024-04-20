package bfs

import (
	"fmt"
	"time"
	"strings"
	// "Tubes2_gomugomuno/Scrape"
	"github.com/gocolly/colly"
)

func printStrings(slice []string) {
	for _, s := range slice {
		fmt.Println(s)
	}
}
// ini temporary ak pindahin soalnya mau eksperimen
func scraper(title string) []string {
	start := time.Now()
	c := colly.NewCollector()

	var links []string
	// var res []*Node
	c.OnHTML("div.mw-page-container a[href^='/wiki/']", func(h *colly.HTMLElement) {
		link := h.Request.AbsoluteURL(h.Attr("href"))
		link = strings.TrimPrefix(link, "https://en.wikipedia.org/wiki/")

		if !strings.ContainsAny(link, ":()/%") {
            if !contains(links, link) && link != title {
                links = append(links, link)
                // res = append(res, createNode(link, nil))
            }
        }
	})

	c.Visit("https://en.wikipedia.org/wiki/" + title)
	fmt.Print("--------------------------The amount of links visited: ")
	fmt.Println(len(links))

	// // printAllQueue(res)
	// printStrings((links))
	// fmt.Println()
	elapsed := time.Since(start).Milliseconds()
	fmt.Println("Its time to stop")
	fmt.Println(elapsed)
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


// func BFS(startTitle string, endNode string) {
func BFS(startTitle string, endNode string) ([][]string, int64, int, int){
	start := time.Now()
	firstNode := createNode(startTitle,nil)	// node pertama
	shortestlength:= 0
	numofcheckednodes := 0
	var queue []*Node		// simpul hidup
	var result [][]string		// list of answers
	// var visitedNodes []*Node // list of visitedNodes, including the ones that had been dequeued 
	visited := make(map[string]string)
	secondtolast := make(map[string]bool) 
	queue = append(queue, firstNode)	// masukkan node pertama ke dalam queue simpul hidup
	visited[firstNode.Title] = ""
	temptitle := scraper(startTitle)
	for _,title := range temptitle{
		var tempparent []string = append(firstNode.Parents, startTitle) 
		queue = append(queue, createNode(title,tempparent))
	}

	// // Testing first scrape on initial web
	// fmt.Println("Queue initial:")
	// printAllQueue(queue)


	
	// parse for startNode and enqueue
	var hasFound bool = false
	for len(queue) > 0 {		//while queue is not empty
		fmt.Print("Panjang queue skrg:")
		fmt.Println(len(queue))
		numofcheckednodes++
		// visitedNodes = append(visitedNodes, queue[0])	// mark as visited
		currentNode := queue[0] //current branch is the start of the queue, dequeue
		fmt.Println(">>>>>>>>>>>>>> Skrg kita cek Node: ")
		currentNode.Print()
		fmt.Println()
		queue = queue[1:]
		// currentNode.Parents = append(currentNode.Parents,currentNode.Title)
		if _, ok := secondtolast[currentNode.Title]; ok {
			continue
		}
		if (len(currentNode.Parents)>shortestlength) && hasFound{
			fmt.Println("ALREADY GOT BETTER ANS STAHP⛔⛔⛔")
			continue
		}
		// grab only the shortest solution or when we have no solution at all.
		if (currentNode.Title == endNode){
			fmt.Println("Destination Reached!")
			fmt.Println("😱😱😱😱😱😱😱😱")
			if (shortestlength==0){	//assign when first time encountering solution
				shortestlength = len(currentNode.Parents)
			} 
			if (len(currentNode.Parents)==shortestlength){
				currentNode.Parents = append(currentNode.Parents, currentNode.Title)		// also add the endNode to the list of results, so if we have Basf with B as endNode, we get {a,s,f,b}
				result = append(result, currentNode.Parents)
				secondtolast[currentNode.Parents[len(currentNode.Parents)-2]] = true
				hasFound = true
				fmt.Println("SOLUSI LAINNN???!!🔍🔍🔍🔍🔍🔍🔍🔍🔍🔍🔍🔍🔍🔍🔍")
				// // ini testing 1 solusi dulu ya:")"
				fmt.Println("GOT BREAK🔨🔨")
				break
			}
		} else{
			// parse and enqueue/create node
			// dont forget to change the newly made Node Parents with the copy of currentNode.Parents + currentNode.Title
			fmt.Println("WAITING FOR SCRAPER🫧🫧🫧🫧🫧🫧")
			var temptitle []string = scraper(currentNode.Title)	// get a list of Node titles from parsing currentNode
			var tempparent []string = currentNode.Parents		// get the list of parent Node from the currentNode
			tempparent = append(tempparent, currentNode.Title)	// add the currentNode into the list of parent Node
			for i := 0; i < len(temptitle); i++ {							// append the newly made node into the queue
				A := createNode(temptitle[i],tempparent)
				
				// if (!isVisited(A,visitedNodes)){
				// 	queue = append(queue, A)
				// 	visitedNodes = append(visitedNodes, A)		// mark newlymade Nodes as visited
				// }

				lastVisited, ok := visited[A.Title]
                if ok && lastVisited == A.Parents[len(A.Parents)-1] {
                    continue // Skip if already visited with the same parent
                }
				queue = append(queue, A)
                visited[A.Title] = A.Parents[len(A.Parents)-1]

			}
		}
	}
	elapsed := time.Since(start).Milliseconds()
	// fmt.Println("Its time to stop")
	// fmt.Println(elapsed)
	return result,elapsed,shortestlength,numofcheckednodes
// 	// should i return array instead...?
}