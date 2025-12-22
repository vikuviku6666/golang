package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}
 

func main() {
	var wg sync.WaitGroup

	words := []string{
		"apple1",
		"apple2",
		"apple3",
		"apple4",
		"apple5",
		"apple6",
		"apple7",
		"apple8",
		"apple9",
	}

	wg.Add(len(words))

	for i, x := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg)
	}

	wg.Wait()
}