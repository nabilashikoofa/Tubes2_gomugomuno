package main

import (
	"fmt"
	"sync"
)

func main() {
	numNodes := 9
	numThreads := 5

	// Calculate chunk size and remainder
	chunkSize := numNodes / numThreads
	remainder := numNodes % numThreads

	// Initialize wait group
	var wg sync.WaitGroup

	// Iterate over threads
	for i := 0; i < numThreads; i++ {
		// Increment wait group
		wg.Add(1)

		// Calculate start and end indices for each thread
		startIndex := i * chunkSize
		endIndex := (i + 1) * chunkSize

		// Adjust end index for the last thread
		if i == numThreads-1 {
			endIndex += remainder
		}

		// Goroutine for processing nodes
		go func(threadID, start, end int) {
			defer wg.Done()

			// Simulate processing nodes
			for j := start; j < end; j++ {
				fmt.Printf("Thread %d is processing node %d\n", threadID, j)
				// Process node j here
			}
		}(i+1, startIndex, endIndex)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
