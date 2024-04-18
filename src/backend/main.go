package main

// TAMBAHAN EROR HANDLING, HANDLING GAADA WIKIPEDIA DENGAN KATA XX MISALNYA KRN TYPO
// HRS CEK KALAU NODEAWAL = NODEAKHIR EXCEPTION

import (
    // "encoding/json"
    "fmt"
    // "net/http"
    "Tubes2_gomugomuno/BFS"
    // "Tubes2_gomugomuno/IDS"
)

// ini cm prototype
func main() {
    fmt.Println("Starting Main...")
	bfs.BFS("Vector","Engineering");
    // http.HandleFunc("/wikirace", func(w http.ResponseWriter, r *http.Request) {
    //     startTitle := r.URL.Query().Get("start")
    //     endTitle := r.URL.Query().Get("end")

    //     if startTitle == "" || endTitle == "" {
    //         http.Error(w, "Missing start or end title", http.StatusBadRequest)
    //         return
    //     }

    //     result, elapsed, shortestlength, numofcheckednodes := bfs.BFS(startTitle, endTitle)

    //     response := struct {
    //         Result  [][]string `json:"result"`
    //         Elapsed int64      `json:"elapsed"`
    //         Shortest int      `json:"shortestlength"`
    //         Checked int      `json:"numofcheckednodes"`
    //     }{result, elapsed, shortestlength, numofcheckednodes}

    //     w.Header().Set("Content-Type", "application/json")
    //     json.NewEncoder(w).Encode(response)
    // })

    // fmt.Println("Server is running on :8080")
    // http.ListenAndServe(":8080", nil)
}