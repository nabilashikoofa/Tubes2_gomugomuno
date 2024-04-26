// package main

// import (
// 	// "encoding/json"
// 	"Tubes2_gomugomuno/BFS"
// 	"Tubes2_gomugomuno/Scrape"
// 	"fmt"
// 	"net/http"
// 	"time"
// 	// "Tubes2_gomugomuno/IDS"
// )


// func isValidLink(title string) bool {
//     title = Scrape.Convert(title)
// 	url := fmt.Sprintf("https://en.wikipedia.org/wiki/%s", title)
// 	resp, err := http.Head(url)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return false
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode == http.StatusOK {
// 		fmt.Println("Link is valid")
// 		return true
// 	} else {
// 		fmt.Println("Link is invalid")
// 		return false
// 	}
// }

// // print all results found
// func printResult(result [][]string) {
// 	fmt.Println("Printing result:")
// 	for _, innerSlice := range result {
// 		fmt.Print("(")
// 		for i, str := range innerSlice {
// 			fmt.Print(str)
// 			// If not the last element in the slice, add a comma
// 			if i < len(innerSlice)-1 {
// 				fmt.Print(",")
// 			}
// 		}
// 		fmt.Println(")")
// 	}
// }


// func main() {
//     var result [][]string
//     var elapsed int64
//     var shortestlength int
//     var numofcheckednodes int
//     fmt.Println("Starting Main...")
//     timee := time.Now()


// 	result, elapsed, shortestlength, numofcheckednodes = bfs.ParallelBFS("Vector","Mathematics") //45613 - 38262 (2 degrees)
    
// 	fmt.Println("Algorithm execution time:", elapsed, "ms")
//     printResult(result)
//     fmt.Println("")
//     fmt.Println("Shortest length: ",shortestlength)
//     fmt.Println("Nodes checked:",numofcheckednodes)
//     elapsedd := time.Since(timee).Milliseconds()
//     fmt.Println("ELAPSED REAL: ",elapsedd)
  
// }