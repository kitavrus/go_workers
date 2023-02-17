package generator

import "time"

func New(max int) chan int {
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