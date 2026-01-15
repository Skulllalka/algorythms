package solutions

import (
	"fmt"
	"sync"
)

func merge(channels ...<-chan int) chan int {
	outCh := make(chan int)
	wg := sync.WaitGroup{}

	for _, channel := range channels {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for value := range channel {
				outCh <- value
			}
		}()
	}

	go func() {
		wg.Wait()
		close(outCh)
	}()

	return outCh
}

func fillChan(n int) <-chan int {
	outChan := make(chan int)
	go func() {
		defer close(outChan)
		for i := 0; i < n; i++ {
			outChan <- i
		}
	}()
	return outChan
}

func MergeMain() {
	a := fillChan(2)
	b := fillChan(4)
	c := fillChan(5)

	d := merge(a, b, c)
	for v := range d {
		fmt.Println(v)
	}
}
