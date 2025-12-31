package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(url string, resultCh chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 50)

	fmt.Printf("image processed: %s\n", url)
	resultCh <- url
}

func main() {
	var wg sync.WaitGroup
	resultCh := make(chan string, 10)
	startTime := time.Now()

	for i := range 10 {
		url := fmt.Sprintf("%d_image.png", i)
		wg.Add(1)
		go worker(url, resultCh, &wg)
	}
	// fan out/ fan in
	// Close resultCh when all workers are done, so ranging over it can finish.
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for result := range resultCh {
		fmt.Printf("received: %s\n", result)
	}

	fmt.Printf("it took %s\n", time.Since(startTime))
}
