package parallelcomputing

import (
	"fmt"
	"sync"
)

func printOdd(oddChan, evenChan chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 100; i += 2 {
		select {
		case <-oddChan:
			fmt.Printf("i = %d\n", i)
			evenChan <- true
		}
	}
}

func printEven(oddChan, evenChan chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 100; i += 2 {
		select {
		case <-evenChan:
			fmt.Printf("i = %d\n", i)
			oddChan <- true
		}
	}
}

func OddEvenMonitor() {
	var wg sync.WaitGroup
	oddChan := make(chan interface{}, 1)
	evenChan := make(chan interface{}, 1)
	oddChan <- true
	wg.Add(1)
	go printOdd(oddChan, evenChan, &wg)
	wg.Add(1)
	go printEven(oddChan, evenChan, &wg)
	wg.Wait()
}
