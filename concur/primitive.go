package concur

import (
	"fmt"
	"sync"
	//	"time"
)

func Counter() {
	counter := 20
	wg := sync.WaitGroup{}
	for i := 0; i <= counter; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(i * i)
		}()
	}
	wg.Wait()

}
