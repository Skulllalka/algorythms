package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"
)

//	Варианты задания :
//	1) Ограничение количества коннектов
//	2) Ограничение количества горутин
//	3) Ограничение количества запросов в секунду

type Client interface {
	SendRequest(ctx context.Context, request Request) error
	WithLimiter(ctx context.Context, requests []Request)
}
type client struct {
}

type Request struct {
	Payload string
}

func (c client) SendRequest(ctx context.Context, request Request) error {
	time.Sleep(1 * time.Second)
	fmt.Println("sending request", request.Payload)
	return nil
}

var RPS = 10

func (c client) WithLimiter(ctx context.Context, reqs []Request) {
	ticker := time.NewTicker(time.Duration(1 * time.Second / time.Duration(RPS)))
	wg := sync.WaitGroup{}
	wg.Add(len(reqs))
	for _, req := range reqs {
		<-ticker.C

		go func() {
			defer wg.Done()
			c.SendRequest(ctx, req)
		}()
	}
	wg.Wait()
}

func main() {
	ctx := context.Background()

	c := client{}
	requests := make([]Request, 1000)
	for i := 0; i < 1000; i++ {
		requests[i] = Request{Payload: strconv.Itoa(i)}
	}

	c.WithLimiter(ctx, requests)

}
