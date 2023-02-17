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
	maxWorkers := 2
	gench := generator.New(maxGen)
	sumch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(maxWorkers)
	for i := 0; i < maxWorkers; i++ {
		go worker(gench, sumch, &wg, i)
	}

	go func() {
		wg.Wait()
		close(sumch)
		fmt.Println("End wg.Wait()")
	}()

	// wg.Wait()
	// close(sumch)

	var sum int
	for i := 0; i < maxGen; i++ {
		sum += <-sumch
	}

	// wg.Wait()
	// close(sumch)

	fmt.Printf("Sum: %d \n", sum)
	elapsed := time.Since(start)
	fmt.Printf("Took ===============> %s\n", elapsed)
}

func worker(inCh <-chan int, outCh chan<- int, wg *sync.WaitGroup, i int) {
	defer wg.Done()

	var sum int
	fmt.Println("start worker: ", i+1)

	for v := range inCh {
		sum += v
	}
	outCh <- sum
	fmt.Printf("end worker: %d sum: %d \n", i+1, sum)
}
