package bfs

import (
	"Tubes2_gomugomuno/Scrape"
	"fmt"
	"sync"
	"time"
	"context"
)

func printStrings(slice []string) {
	for _, s := range slice {
		fmt.Println(s)
	}
}



// // func BFS(startTitle string, endNode string) {
// func BFS(startTitle string, endNode string) ([][]string, int64, int, int){
// 	start := time.Now()
// 	firstNode := createNode(startTitle,nil)	// node pertama
// 	shortestlength:= 0
// 	numofcheckednodes := 0
// 	var queue []*Node		// simpul hidup
// 	var result [][]string		// list of answers
// 	visited := make(map[string]string)
// 	secondtolast := make(map[string]bool) 
// 	//second to last map is for efficiency of not checking the second to last parent nodes that have found in a solution
// 	//for example i went from Vector to Mathematics and i got it in the order (Vector,Euclidean_vector,Mathematics). now all the nodes i've created in the queue that has the parent euclidean vector init does not need to be checked, since we already confirm that the endnode is reachable, thus cutting the time making it more faster.  
// 	queue = append(queue, firstNode)	// masukkan node pertama ke dalam queue simpul hidup
// 	visited[firstNode.Title] = ""
// 	temptitle := Scrape.Scraper(startTitle)
// 	for _,title := range temptitle{
// 		var tempparent []string = append(firstNode.Parents, startTitle) 
// 		queue = append(queue, createNode(title,tempparent))
// 	}
	
// 	// parse for startNode and enqueue
// 	var hasFound bool = false
// 	var foundOneSol bool = false
// 	for len(queue) > 0 {		//while queue is not empty
// 		fmt.Print("Panjang queue skrg:")
// 		fmt.Println(len(queue))
// 		numofcheckednodes++
// 		currentNode := queue[0] //current branch is the start of the queue, dequeue
// 		fmt.Println(">>>>>>>>>>>>>> Skrg kita cek Node: ")
// 		currentNode.Print()
// 		fmt.Println()
// 		queue = queue[1:]
// 		// currentNode.Parents = append(currentNode.Parents,currentNode.Title)
// 		if _, ok := secondtolast[currentNode.Title]; ok {
// 			continue
// 		}
// 		if (len(currentNode.Parents)>shortestlength) && hasFound{
// 			fmt.Println("ALREADY GOT BETTER ANS STAHP⛔⛔⛔")
// 			continue
// 		}
// 		// grab only the shortest solution or when we have no solution at all.
// 		if (currentNode.Title == endNode){
// 			fmt.Println("Destination Reached!")
// 			fmt.Println("😱😱😱😱😱😱😱😱")
// 			if (shortestlength==0){	//assign when first time encountering solution
// 				shortestlength = len(currentNode.Parents)
// 			} 
// 			if (len(currentNode.Parents)==shortestlength){
// 				currentNode.Parents = append(currentNode.Parents, currentNode.Title)		// also add the endNode to the list of results, so if we have Basf with B as endNode, we get {a,s,f,b}
// 				result = append(result, currentNode.Parents)
// 				secondtolast[currentNode.Parents[len(currentNode.Parents)-2]] = true
// 				hasFound = true
// 				fmt.Println("SOLUSI LAINNN???!!🔍🔍🔍🔍🔍🔍🔍🔍🔍🔍🔍🔍🔍🔍🔍")
// 				// // ini testing 1 solusi dulu ya:")"
// 				fmt.Println("GOT BREAK🔨🔨")
// 				break
// 			}
// 		} else{
// 			// parse and enqueue/create node
// 			// dont forget to change the newly made Node Parents with the copy of currentNode.Parents + currentNode.Title
// 			fmt.Println("WAITING FOR SCRAPER🫧🫧🫧🫧🫧🫧")
// 			var temptitle []string = Scrape.Scraper(currentNode.Title)	// get a list of Node titles from parsing currentNode
// 			var tempparent []string = currentNode.Parents		// get the list of parent Node from the currentNode
// 			tempparent = append(tempparent, currentNode.Title)	// add the currentNode into the list of parent Node
// 			for i := 0; i < len(temptitle); i++ {
// 				A := createNode(temptitle[i],tempparent)
// 				if (temptitle[i]==endNode){	//ternyata waktu scraping udah ketemu jadi kt jangan scrape lagi biar hemat waktu
// 					foundOneSol = true
// 					queue = append(queue, A)
// 				}	
// 				if (!foundOneSol){
// 					lastVisited, ok := visited[A.Title]
// 					if ok && lastVisited == A.Parents[len(A.Parents)-1] {
// 						continue // Skip if already visited with the same parent
// 					}
// 					// append the newly made node into the queue
// 					queue = append(queue, A)
// 					visited[A.Title] = A.Parents[len(A.Parents)-1]
// 				}
				

// 			}
// 		}
// 	}
// 	elapsed := time.Since(start).Milliseconds()
// 	return result,elapsed,shortestlength,numofcheckednodes
// }


// return the list of Nodes from startTitle created inside a queue 
func initialBFS(startTitle string) ([]*Node){
	firstNode := createNode(startTitle,nil)	// node pertama
	var queue []*Node		// simpul hidup
	temptitle := Scrape.Scraper(startTitle)
	for i := 0; i<len(temptitle); i++{
		var tempparent []string = append(firstNode.Parents, startTitle) 
		queue = append(queue, createNode(temptitle[i],tempparent))
	}
	return queue
}

	
// find a solution inside a queue (the queue has been divided by the number of threads)
func multiBFS(queue []*Node, startNode string, endNode string, shortestlengthavailable *int, resultavailable *[][]string, numofcheckednodesavail *int, elapsedavail *int64, ctx context.Context) {
	var result [][]string		// list of answers
	var hasFound bool = false
	var foundOneSol bool = false
	numofcheckednodes := 0
	shortestlength:= 0
	visited := make(map[string]string)
	secondtolast := make(map[string]bool) 
	//second to last map is for efficiency of not checking the second to last parent nodes that have found in a solution
	//for example i went from Vector to Mathematics and i got it in the order (Vector,Euclidean_vector,Mathematics). now all the nodes i've created in the queue that has the parent euclidean vector init does not need to be checked, since we already confirm that the endnode is reachable, thus cutting the time making it more faster.  
	start := time.Now()
	
	
	for len(queue) > 0 {		//while queue is not empty
		select{
		case <- ctx.Done():
			return
		default:
		numofcheckednodes++
		currentNode := queue[0] //current branch is the start of the queue, dequeue
	fmt.Println(len(currentNode.Parents))
	fmt.Println(*resultavailable)
	fmt.Println(">>>>>>>>>>>>>> Skrg kita cek Node: ")
	currentNode.Print()
	fmt.Println()
	queue = queue[1:]
	
	// stop if its checking the node with the same parents length as the shortestlengthavailable in parallelBFS
	if (len(currentNode.Parents)>=*shortestlengthavailable && *shortestlengthavailable!=0){
		fmt.Println(*shortestlengthavailable)
		fmt.Println(*resultavailable)
		fmt.Println("THREAD DIHENTIKANNN⛔⛔⛔⛔⛔⛔⛔")
		return
	}
	if _, ok := secondtolast[currentNode.Title]; ok {
		fmt.Println("THREAD DIHENTIKANNN22222⛔⛔⛔⛔⛔⛔⛔")
		continue
	}
	if (len(currentNode.Parents)>=shortestlength) && hasFound{
		fmt.Println("ALREADY GOT BETTER ANS STAHP⛔⛔⛔")
		return
	}
	// grab only the shortest solution or when we have no solution at all.
	if (currentNode.Title == endNode){
		fmt.Println("Destination Reached!")
		fmt.Println("😱😱😱😱😱😱😱😱")
		fmt.Println("😱😱😱😱😱😱😱😱")
		if (shortestlength==0){	//assign when first time encountering solution
			shortestlength = len(currentNode.Parents)
			} 
		if (len(currentNode.Parents)<=shortestlength){
			currentNode.Parents = append(currentNode.Parents, currentNode.Title)		// also add the endNode to the list of results, so if we have Basf with B as endNode, we get {a,s,f,b}
			result = append(result, currentNode.Parents)
			secondtolast[currentNode.Parents[len(currentNode.Parents)-2]] = true
		hasFound = true
		*shortestlengthavailable = shortestlength
		*resultavailable = result
		*numofcheckednodesavail = numofcheckednodes
		elapsed := time.Since(start).Milliseconds()
		*elapsedavail = elapsed
		fmt.Println("GOT BREAK🔨🔨")
		fmt.Println(elapsedavail)
		break
		}
		} else{
			var temptitle []string
			var tempparent []string
			fmt.Println(*shortestlengthavailable)
			if (!foundOneSol){
				fmt.Println("WAITING FOR SCRAPER🫧🫧🫧🫧🫧🫧")
				temptitle = Scrape.Scraper(currentNode.Title)	// get a list of Node titles from parsing currentNode
				tempparent = currentNode.Parents		// get the list of parent Node from the currentNode
				tempparent = append(tempparent, currentNode.Title)	// add the currentNode into the list of parent Node
				
			}
			for i := 0; i < len(temptitle); i++ {
				A := createNode(temptitle[i],tempparent)
				if (temptitle[i]==endNode){	//ternyata waktu scraping udah ketemu jadi kt jangan scrape lagi biar hemat waktu
					foundOneSol = true
					queue = append(queue, A)
					numofcheckednodes += len(queue)-1
					fmt.Println("WOI KETEMU DISINIIII😱😱😱😱😱😱😱😱😱😱😱😱😱😱😱😱😱😱😱😱😱😱😱😱😱😱😱😱")
					queue = queue[len(queue)-1:]
					fmt.Println("💎💎💎💎💎💎💎💎💎💎💎💎💎💎💎💎💎💎💎💎💎💎💎💎")
					break
				}	
				if (!foundOneSol){
					lastVisited, ok := visited[A.Title]
					if (ok && lastVisited == A.Parents[len(A.Parents)-1] || A.Title == startNode) {
						continue // Skip if already visited with the same parent
					}
					// append the newly made node into the queue
					queue = append(queue, A)
					visited[A.Title] = A.Parents[len(A.Parents)-1]
				}
				
				
			}
		}
	}
}
	
}

func ParallelBFS(startTitle string, endNode string) ([][]string, int64, int, int, int){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
    defer cancel() // Cancel context to release resources when done


	var shortestlength int = 0
	var result [][] string 
	var numofcheckednodes int = 0
	var numThreads int = 500
	var elapsed int64
	var wg sync.WaitGroup
	endNode = Scrape.Convert(endNode)
	shortestlengthavail := &shortestlength
	resultavailable := &result
	numofcheckednodesavail := &numofcheckednodes
	elapsedavail := &elapsed
	array := make([]int, numThreads)
	queue := initialBFS(startTitle)
	//divide queue into equal number of nodes per thread
	for i := 0; i<numThreads; i++ {
		array[i] = len(queue) / numThreads
	}
	for i := 0; i<len(queue)%numThreads; i++ {
		array[i]++
	}
	
	startIndex := 0

	for i := 0; i < numThreads; i++ {
		// perform go routine here
		endIndex := startIndex + array[i]
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			tes := queue[start:end]
			multiBFS(tes, startTitle, endNode, shortestlengthavail, resultavailable, numofcheckednodesavail, elapsedavail, ctx)
		}(startIndex, endIndex)
		startIndex = endIndex
	}
	wg.Wait()
	return result,elapsed,shortestlength,numofcheckednodes, len(result)
}