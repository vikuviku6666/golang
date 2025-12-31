package main

import (
	"fmt"
	"sync"
	"time"
)



func worker(jobs chan string, resultCh chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		time.Sleep(time.Millisecond * 50)
		//fmt.Printf("image processed: %s\n", job)
		resultCh <- job
	}

	fmt.Printf("Worker shutting down\n")
}

func main() {
	jobs := []string{
	"1_image.png",
	"2_image.png",
	"3_image.png",
	"4_image.png",
	"5_image.png",
	"6_image.png",
	"7_image.png",
	"8_image.png",
	"9_image.png",
	"10_image.png",
	"11_image.png",
	"12_image.png",
	"13_image.png",
	"14_image.png",
	"15_image.png",
	"16_image.png",
	"17_image.png",
	"18_image.png",
	"19_image.png",
	"20_image.png",
	"21_image.png",
	"22_image.png",
}
	var wg sync.WaitGroup
	resultCh := make(chan string, 50)
	jobsCh := make(chan string, len(jobs))
	totalWorkers := 5
	startTime := time.Now()

	for i:=1; i <= totalWorkers; i++ {
		wg.Add(1)
		go worker(jobsCh, resultCh, &wg)
	}
	// fan out/ fan in pattern
	// worker pool pattern
	// Close resultCh when all workers are done, so ranging over it can finish.
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// send the jobs
	for i:=0; i<len(jobs); i++ {
		jobsCh <- jobs[i]
	}

	close(jobsCh)

	for result := range resultCh {
		fmt.Printf("Job completed: %s\n", result)
	}

	fmt.Printf("it took %s\n", time.Since(startTime))
}
