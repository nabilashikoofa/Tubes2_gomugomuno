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


// // ini cm prototype
// func main() {
//     var result [][]string
//     var elapsed int64
//     var shortestlength int
//     var numofcheckednodes int
//     fmt.Println("Starting Main...")
//     timee := time.Now()
//     // bfs.BFS("Vector","Vector_space")

// 	// result,elapsed, shortestlength,numofcheckednodes = bfs.BFS("Fanta_cake","Flour") //45613 - 38262 (2 degrees)

// 	result, elapsed, shortestlength, numofcheckednodes = bfs.ParallelBFS("Fanta_cake","Flour") //45613 - 38262 (2 degrees)
    
// 	fmt.Println("Algorithm execution time:", elapsed, "ms")
//     printResult(result)
//     fmt.Println("")
//     fmt.Println("Shortest length: ",shortestlength)
//     fmt.Println("Nodes checked:",numofcheckednodes)
//     elapsedd := time.Since(timee).Milliseconds()
//     fmt.Println("ELAPSED REAL: ",elapsedd)
//     // http.HandleFunc("/wikirace", func(w http.ResponseWriter, r *http.Request) {
//     //     startTitle := r.URL.Query().Get("start")
//     //     endTitle := r.URL.Query().Get("end")
//     //     var validend bool
//     //     var validstart bool
//     //     validend = isValidLink(endTitle)
//     //     validstart = isValidLink(startTitle)
//     //     if startTitle == "" || endTitle == "" {
//     //         http.Error(w, "Missing start or end title", http.StatusBadRequest)
//     //         return
//     //     }
//     //     if !validend{
//     //         http.Error(w, "End link did not exist", http.StatusNotFound)
//     //         return
//     //     }
//     //     if !validstart{
//     //         http.Error(w, "Start link did not exist", http.StatusNotFound)
//     //         return
//     //     }
//     //     result, elapsed, shortestlength, numofcheckednodes := bfs.BFS(startTitle, endTitle)

//     //     response := struct {
//     //         Result  [][]string `json:"result"`
//     //         Elapsed int64      `json:"elapsed"`
//     //         Shortest int      `json:"shortestlength"`
//     //         Checked int      `json:"numofcheckednodes"`
//     //     }{result, elapsed, shortestlength, numofcheckednodes}

//     //     w.Header().Set("Content-Type", "application/json")
//     //     json.NewEncoder(w).Encode(response)
//     // })

//     // fmt.Println("Server is running on :8080")
//     // http.ListenAndServe(":8080", nil)
// }