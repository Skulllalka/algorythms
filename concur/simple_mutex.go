package concur

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type myMutex struct {
	locked int64
}

func (m *myMutex) Lock() {
	for {
		if atomic.CompareAndSwapInt64(&m.locked, 0, 1) {
			return
		}
	}
}

func (m *myMutex) Unlock() {
	atomic.StoreInt64(&m.locked, 0)
}

func Smain() {
	mu := myMutex{}
	wg := &sync.WaitGroup{}
	count := 0
	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println(count)
}
