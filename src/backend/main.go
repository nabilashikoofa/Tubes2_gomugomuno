package main

import (
	"encoding/json"
	"Tubes2_gomugomuno/BFS"
	"Tubes2_gomugomuno/Scrape"
	"fmt"
	"net/http"
	// "Tubes2_gomugomuno/IDS"
)

// TAMBAHAN EROR HANDLING, HANDLING GAADA WIKIPEDIA DENGAN KATA XX MISALNYA KRN TYPO
// HRS CEK KALAU NODEAWAL = NODEAKHIR EXCEPTION

// INI BLM AKU CHECK YA
// VALIDASI LINK INPUTAN
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

// print all results found
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

type AlgorithmFunc func(startNode, endNode string) ([][]string, int64, int, int)

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

	result, elapsed, shortestlength, numofcheckednodes := algorithm(startNode, endNode)
    response := struct {
        Result  [][]string `json:"result"`
        Elapsed int64      `json:"elapsed"`
        Shortest int      `json:"shortestlength"`
        Checked int      `json:"numofcheckednodes"`
    }{result, elapsed, shortestlength, numofcheckednodes}

	json.NewEncoder(w).Encode(response)
}

// ini cm prototype
func main() {
    http.HandleFunc("/api/bfs", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Enter BFS")
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        executeAlgorithm(w, r, bfs.BFS)

    // http.HandleFunc("/api/ids", func(w http.ResponseWriter, r *http.Request) {
    //     executeAlgorithm(w, r, ids.IDS)   

    })

    fmt.Println("Server is running on :3000")
    http.ListenAndServe(":3000", nil)
}