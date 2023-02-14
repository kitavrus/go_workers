package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	maxGen := 100
	chr := generator(maxGen)
	ch := sumch(chr, &wg)

	var sum int
	for i := 0; i < maxGen; i++ {
		sum += <-ch
	}

	// go func() {
	wg.Wait()
	close(ch)
	fmt.Println("End wg.Wait()")
	// }()

	fmt.Printf("Sum: %d \n", sum)
	elapsed := time.Since(start)
	fmt.Printf("Took ===============> %s\n", elapsed)
}

func sumch(inch chan int, wg *sync.WaitGroup) chan int {
	chw := make(chan int)

	wg.Add(2)
	// wg.Add(1)
	go func() {
		fmt.Println("start worker: ", 1)
		var sum int
		for v := range inch {
			chw <- v
			sum += v
		}
		// chw <- sum
		defer wg.Done()
		// fmt.Printf("end worker: 1 sum: %d \n", 1)
		fmt.Printf("end worker: 1 sum: %d \n", sum)
	}()

	// wg.Add(1)
	go func() {
		fmt.Println("start worker: ", 2)
		var sum int
		for v := range inch {
			chw <- v
			sum += v
		}
		// chw <- sum
		defer wg.Done()
		// fmt.Printf("end worker: 2 sum: %d \n", 2)
		fmt.Printf("end worker: 2 sum: %d \n", sum)
	}()
	return chw

}

func generator(max int) chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < max; i++ {
			ch <- 1
		    time.Sleep(1 * time.Millisecond)
		}
		close(ch)
	}()

	return ch
}
