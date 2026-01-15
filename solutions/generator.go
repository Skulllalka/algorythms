package solutions

import (
	"context"
	"fmt"
)

func GeneratorMain() {
	ctx := context.Background()
	pipeline := squarer(ctx, generator(ctx, 1, 2, 3))

	for v := range pipeline {
		fmt.Println(v)
	}
}

func generator(ctx context.Context, in ...int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, number := range in {
			select {
			case <-ctx.Done():
				return
			case ch <- number:
			}
		}
	}()

	return ch
}

func squarer(ctx context.Context, in <-chan int) <-chan int {
	res := make(chan int)
	go func() {
		defer close(res)
		for value := range in {
			select {
			case <-ctx.Done():
				return
			case res <- value * value:
			}
		}
	}()

	return res
}
