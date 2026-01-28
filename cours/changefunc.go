package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func unpredictableFunc() int {
	n := rand.Intn(5)
	time.Sleep(time.Duration(n) * time.Second)
	return n
}

func predictableFunc(ctx context.Context) (int, error) {
	ch := make(chan struct{})
	var result int

	go func() {
		result = unpredictableFunc()
		close(ch)
	}()
	var cancel context.CancelFunc

	if _, ok := ctx.Deadline(); !ok {
		ctx, cancel = context.WithTimeout(ctx, time.Second*4)
		defer cancel()
	}

	select {
	case <-ch:
		return result, nil
	case <-ctx.Done():
		return result, fmt.Errorf("out of time")
	}
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second)

	n, err := predictableFunc(ctx)
	if err != nil {
		fmt.Printf("time was %d with error %s", n, err)
	} else {
		fmt.Printf("time is %d", n)
	}
}
