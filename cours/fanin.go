package main

import (
	"math/rand"
	"sync"
)

func fanin(chans ...<-chan int) <-chan int {
	out := make(chan int)
	wg := sync.WaitGroup{}
	for _, ch := range chans {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for val := range ch {
				out <- val
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {

}
