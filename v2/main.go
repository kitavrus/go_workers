package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	start := time.Now()

	maxGen := 100
	maxWorkers := 2
	gench := generator(maxGen)
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
	var sum int
	fmt.Println("start worker: ", i)

	for v := range inCh {
		sum += v
	}

	outCh <- sum
	wg.Done()
	fmt.Printf("end worker: %d sum: %d \n", i, sum)
}

func generator(max int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < max; i++ {
			ch <- 1
			time.Sleep(1 * time.Millisecond)
		}

	}()

	return ch
}
