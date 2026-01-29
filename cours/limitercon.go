// package main

// import (
// 	"context"
// 	"fmt"
// 	"strconv"
// 	"sync"
// 	"time"
// )

// //	Варианты задания :
// //	1) Ограничение количества коннектов
// //	2) Ограничение количества горутин
// //	3) Ограничение количества запросов в секунду

// type Client interface {
// 	SendRequest(ctx context.Context, request Request) error
// 	WithLimiter(ctx context.Context, requests []Request)
// }
// type client struct {
// }

// type Request struct {
// 	Payload string
// }

// func (c client) SendRequest(ctx context.Context, request Request) error {
// 	time.Sleep(100 * time.Millisecond)
// 	fmt.Println("sending request", request.Payload)
// 	return nil
// }

// var maxConnects = 10

// func (c client) WithLimiter(ctx context.Context, ch chan Request) {
// 	wg := sync.WaitGroup{}
// 	for range maxConnects {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			for req := range ch {
// 				c.SendRequest(ctx, req)
// 			}
// 		}()
// 	}

// 	wg.Wait()
// }

// func main() {
// 	ctx := context.Background()

// 	c := client{}
// 	requests := make([]Request, 1000)
// 	for i := 0; i < 1000; i++ {
// 		requests[i] = Request{Payload: strconv.Itoa(i)}
// 	}

// 	c.WithLimiter(ctx, generate(requests))

// }

// func generate(reqs []Request) chan Request {
// 	ch := make(chan Request)
// 	go func() {
// 		for _, v := range reqs {
// 			ch <- v
// 		}
// 		close(ch)
// 	}()
// 	return ch
// }
