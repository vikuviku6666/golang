package main

import (
	"fmt"
	"runtime"
	"sync"
)


func printSomething(i int, sumCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := i + 10
	sumCh <- sum
}





func main() {
	var wg sync.WaitGroup
	sumCh := make(chan int)
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	for i:=range 10{
		wg.Add(1)
		go printSomething(i,  sumCh, &wg)
		fmt.Println(<-sumCh)
	}
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()

	fmt.Println("Yes, Printed!")
}