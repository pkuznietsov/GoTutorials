package main

import (
	"fmt"
	"sync"
	"time"
)

func myFunc(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("Finished Executing Goroutine")
	wg.Done()
}

func main() {
	fmt.Println("Go WaitGroup ttrl")
	var wg sync.WaitGroup
	wg.Add(1)
	go myFunc(&wg)
	wg.Wait()
	fmt.Println("Program Finished")
}
