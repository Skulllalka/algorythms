package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	wg := sync.WaitGroup{}
	//var counter int64
	count := atomic.Int64{}

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count.Add(1)
			//atomic.AddInt64(&counter, 1)
		}()
	}

	wg.Wait()
	//fmt.Println(counter)
	fmt.Println(count.Load())
}
