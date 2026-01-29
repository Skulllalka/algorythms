// package main

// import (
// 	"context"
// 	"fmt"
// 	"strconv"
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
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("sending request", request.Payload)
// 	return nil
// }

// var maxGorutines = 100

// func (c client) WithLimiter(ctx context.Context, reqs []Request) {
// 	tokens := make(chan struct{}, maxGorutines)

// 	go func() {
// 		for range maxGorutines {
// 			tokens <- struct{}{}
// 		}
// 	}()
// 	for _, req := range reqs {
// 		<-tokens
// 		go func() {
// 			defer func() {
// 				tokens <- struct{}{}
// 			}()
// 			c.SendRequest(ctx, req)
// 		}()
// 	}

// 	for range maxGorutines {
// 		<-tokens
// 	}
// }

// func main() {
// 	ctx := context.Background()

// 	c := client{}
// 	requests := make([]Request, 1000)
// 	for i := 0; i < 1000; i++ {
// 		requests[i] = Request{Payload: strconv.Itoa(i)}
// 	}

// 	c.WithLimiter(ctx, requests)

// }
