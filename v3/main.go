package main

import (
	generator "ep24/internal"
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	maxGen := 100
	chr := generator.New(maxGen)
	ch := workers(chr, &wg)

	var sum int
	for i := 0; i < maxGen; i++ {
		sum += <-ch
	}

	wg.Wait()
	close(ch)
	fmt.Println("End wg.Wait()")

	fmt.Printf("Sum: %d \n", sum)
	elapsed := time.Since(start)
	fmt.Printf("Took ===============> %s\n", elapsed)
}

func workers(inch chan int, wg *sync.WaitGroup) chan int {
	chw := make(chan int)

	wg.Add(2)
	go func() {
		fmt.Println("start worker: ", 1)
		var sum int
		for v := range inch {
			chw <- v
			sum += v
		}
		defer wg.Done()
		fmt.Printf("end worker: 1 sum: %d \n", sum)
	}()

	go func() {
		fmt.Println("start worker: ", 2)
		var sum int
		for v := range inch {
			chw <- v
			sum += v
		}
		defer wg.Done()
		fmt.Printf("end worker: 2 sum: %d \n", sum)
	}()
	return chw
}