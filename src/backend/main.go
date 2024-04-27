package main

import (
	"encoding/json"
	"Tubes2_gomugomuno/BFS"
	"Tubes2_gomugomuno/IDS"
	"Tubes2_gomugomuno/Scrape"
	"fmt"
	"net/http"
)

// Checks if the given Wikipedia link is valid
func isValidLink(title string) bool {
    title = Scrape.Convert(title)
	url := fmt.Sprintf("https://en.wikipedia.org/wiki/%s", title)
	resp, err := http.Head(url)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Link is valid")
		return true
	} else {
		fmt.Println("Link is invalid")
		return false
	}
}

// Print all results found
func printResult(result [][]string) {
	fmt.Println("Printing result:")
	for _, innerSlice := range result {
		fmt.Print("(")
		for i, str := range innerSlice {
			fmt.Print(str)
			// If not the last element in the slice, add a comma
			if i < len(innerSlice)-1 {
				fmt.Print(",")
			}
		}
		fmt.Println(")")
	}
}

// AlgorithmFunc defines a function signature for algorithms
type AlgorithmFunc func(startNode, endNode string) ([][]string, int64, int, int, int)

// Processing data received from the frontend
func executeAlgorithm(w http.ResponseWriter, r *http.Request, algorithm AlgorithmFunc) {
	startNode := r.URL.Query().Get("startNode")
	endNode := r.URL.Query().Get("endNode")

	if startNode == "" || endNode == "" {

		http.Error(w, "Missing start or end node", http.StatusBadRequest)
		return
	}

    if startNode == endNode{
        http.Error(w, "Start and end node cannot be the same.", http.StatusBadRequest)
		return
    }

	if !isValidLink(startNode) {
		http.Error(w, "Invalid start node", http.StatusNotFound)
		return
	}

	if !isValidLink(endNode) {
		http.Error(w, "Invalid end node", http.StatusNotFound)
		return
	}

	result, elapsed, shortestlength, numofcheckednodes, path := algorithm(startNode, endNode)
    fmt.Println("👌 ",elapsed)

	response := struct {
        Result  	[][]string `json:"result"`
        Elapsed 	int64      `json:"elapsed"`
        Shortest 	int        `json:"shortestlength"`
        Checked 	int        `json:"numofcheckednodes"`
		Path 		int        `json:"path"`
    }{result, elapsed, shortestlength, numofcheckednodes, path}

	json.NewEncoder(w).Encode(response)
}

// Connecting the backend with the frontend
func main() {
    handleAlgorithm := func(endpoint string, algorithm func(http.ResponseWriter, *http.Request)) {
        http.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
            fmt.Println("Enter", endpoint)
            w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
            w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
            w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
            algorithm(w, r)
        })
    }
    
    handleAlgorithm("/api/bfs", func(w http.ResponseWriter, r *http.Request) {
        executeAlgorithm(w, r, bfs.ParallelBFS)
    })
    
    handleAlgorithm("/api/ids", func(w http.ResponseWriter, r *http.Request) {
        executeAlgorithm(w, r, ids.IDSParalel)
    })	

    fmt.Println("Server is running on :3000")
    http.ListenAndServe(":3000", nil)
}