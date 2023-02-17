package main

import (
	"fmt"
	"sync"
	"time"
	generator "ep24/internal"
)

func main() {
	start := time.Now()
	maxGen := 100
	chr := generator.New(maxGen)
	ch := workers(chr)

	var sum int
	for i := 0; i < maxGen; i++ {
		sum += <-ch
	}

	fmt.Printf("Sum: %d \n", sum)
	elapsed := time.Since(start)
	fmt.Printf("Took ===============> %s\n", elapsed)
}

func workers(inch <-chan int) chan int {
	chw := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		fmt.Println("start worker: ", 1)
		for v := range inch {
			chw <- v
		}
		wg.Done()
		fmt.Printf("end worker: 1 sum: %d \n", 1)
	}()

	go func() {
		fmt.Println("start worker: ", 2)
		for v := range inch {
			chw <- v
		}
		wg.Done()
		fmt.Printf("end worker: 2 sum: %d \n", 2)
	}()

	go func() {
		wg.Wait()
		close(chw)
		fmt.Println("End wg.Wait()")
	}()

	return chw
}
