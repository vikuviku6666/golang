package main

import (
	"fmt"
	"sync"
	"time"
)


func worker(url string, resultCh chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 10)

	fmt.Printf("image processed: %s\n", url)
	resultCh <- url
}

func main() {
	var wg sync.WaitGroup
	resultCh := make(chan string)
	startTime := time.Now()
	for i := range 100 {
		url := fmt.Sprint(i) + "_image.png"
		wg.Add(1)
		go worker(url, resultCh, &wg)
		fmt.Println(<-resultCh)
	}

	
	wg.Wait()

	close(resultCh)


	fmt.Printf("it took %s \n", time.Since(startTime))
}
