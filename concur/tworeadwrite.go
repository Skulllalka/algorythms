package concur

import (
	"fmt"
	"sync"
)

func TwoReadWrite() {
	storage := make(map[int]int, 1000)

	wg := sync.WaitGroup{}
	ops := 1000
	mu := sync.RWMutex{}

	for i := 0; i < ops; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			storage[i] = i
			mu.Unlock()
		}()
	}

	for i := 0; i < ops; i++ {
		wg.Go(func() {
			mu.RLock()
			value := storage[i]
			mu.RUnlock()
			fmt.Println(value)
		})
	}

	wg.Wait()
}
