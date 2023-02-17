package main

import (
	"fmt"
)

func main() {
	// Set the number of workers
	numWorkers := 5

	// Create a buffered channel to store the data
	dataChan := make(chan string, 100)

	// Start a goroutine to write data to the channel
	go func() {
		for i := 0; i < 100; i++ {
			dataChan <- fmt.Sprintf("Data %d", i)
			//time.Sleep(time.Millisecond * 100)
		}
		close(dataChan)
	}()

	// Start the workers
	for i := 0; i < numWorkers; i++ {
		go func() {
			for data := range dataChan {
				fmt.Println(data)
			}
		}()
	}

	// Wait for the workers to finish
	//time.Sleep(time.Second * 1)
	select {}
}
