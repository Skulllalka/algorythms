package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	urls := []string{
		"https://google.com",
		"https://yandex.ru",
	}
	res := process(urls)
	fmt.Println(res)
}

var client http.Client
var maxCon = 10

func process(urls []string) map[int]int {
	wg := sync.WaitGroup{}
	statusCodeCount := make(map[int]int)
	ch := make(chan string)

	go func() {
		for _, url := range urls {
			ch <- url
		}
		close(ch)
	}()
	mu := &sync.Mutex{}
	processUrl := func(url string, mu *sync.Mutex) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("some err", err.Error())
		}
		mu.Lock()
		statusCodeCount[resp.StatusCode]++
		mu.Unlock()
	}

	for range maxCon {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for url := range ch {
				processUrl(url, mu)
			}
		}()
	}

	wg.Wait()
	return statusCodeCount
}
