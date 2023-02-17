package main

import (
	"log"
	generator "ep24/internal"
	"sync"
	"testing"
	"io/ioutil"
)

var (
	maxGen = 100
	maxWorkers = 2
)

func BenchmarkWorkersV1(b *testing.B) {

	log.SetOutput(ioutil.Discard)
	for i := 0; i < b.N; i++ { // internal
		chr := generator.New(maxGen)
		ch := workersV1(chr)

		var sum int
		for i := 0; i < maxGen; i++ {
			sum += <-ch
		}
	}
}



func BenchmarkWorkersV2(b *testing.B) {

	for i := 0; i < b.N; i++ { // internal
		gench := generator.New(maxGen)
		sumch := make(chan int)
		var wg sync.WaitGroup

		wg.Add(maxWorkers)
		for i := 0; i < maxWorkers; i++ {
			go workerV2(gench, sumch, &wg, i)
		}

		go func() {
			wg.Wait()
			close(sumch)
			log.Println("End wg.Wait()")
		}()

		var sum int
		for i := 0; i < maxGen; i++ {
			sum += <-sumch
		}
		log.Printf("Sum: %d \n", sum)
	}
}



func BenchmarkWorkersV3(b *testing.B) {
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ { // internal
		chr := generator.New(maxGen)
		ch := workersV3(chr, &wg)

		var sum int
		for i := 0; i < maxGen; i++ {
			sum += <-ch
		}

		wg.Wait()
		close(ch)
	}
}

func workersV1(inch <-chan int) chan int {
	chw := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)
	worker := func() {
		log.Println("start worker: ", 1)
		for v := range inch {
			chw <- v
		}
		wg.Done()
		log.Printf("end worker: 1 sum: %d \n", 1)
	}

	go worker()
	go worker()
	
	go func() {
		wg.Wait()
		close(chw)
		log.Println("End wg.Wait()")
	}()

	return chw
}

func workerV2(inCh <-chan int, outCh chan<- int, wg *sync.WaitGroup, i int) {
	var sum int
	log.Println("start worker: ", i)

	for v := range inCh {
		sum += v
	}

	outCh <- sum
	wg.Done()
	log.Printf("end worker: %d sum: %d \n", i, sum)
}


func workersV3(inch chan int, wg *sync.WaitGroup) chan int {
	chw := make(chan int)

	wg.Add(2)
	go func() {
		log.Println("start worker: ", 1)
		var sum int
		for v := range inch {
			chw <- v
			sum += v
		}
		defer wg.Done()
		log.Printf("end worker: 1 sum: %d \n", sum)
	}()

	go func() {
		log.Println("start worker: ", 2)
		var sum int
		for v := range inch {
			chw <- v
			sum += v
		}
		defer wg.Done()
		log.Printf("end worker: 2 sum: %d \n", sum)
	}()
	return chw
}