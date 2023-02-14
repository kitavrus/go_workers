package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	maxGen := 100
	chr := generator(maxGen)
	ch := sumch(chr)

	var sum int
	// Если канал не закрыт то упадем с дедлоком
	// for v := range ch {
	// 	sum += v
	// }
	//  Работает и с закрытым каналом
	for i := 0; i < maxGen; i++ {
		sum += <-ch
	}
	// sum = calculate(ch)

	fmt.Printf("Sum: %d \n", sum)
	elapsed := time.Since(start)
	fmt.Printf("Took ===============> %s\n", elapsed)
}

func sumch(inch <-chan int) chan int {
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

// Как сделать правильно?
// Используем для чтения из нескольких каналов?
// func calculate(ch chan int) int {
// 	var sum int
// 	var ok bool = true
// 	for {
// 		select {
// 		case v, ok := <-ch:
// 			if !ok {
// 				return sum
// 			}
// 			sum += v

// 		default:
// 			if !ok {
// 				return sum
// 			}
// 			//
// 		}
// 	}

// 	//  return sum
// }
