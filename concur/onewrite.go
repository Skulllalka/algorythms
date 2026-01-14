package concur

import (
	"fmt"
	"sync"
)

func OneWrite() {
	writes := 1000
	storage := make(map[int]int, writes)

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	for i := 0; i < writes; i++ {
		wg.Go(func() {
			mu.Lock()
			storage[i] = i
			mu.Unlock()
		})
	}

	wg.Wait()
	fmt.Println(storage)
}
