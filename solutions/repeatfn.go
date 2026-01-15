package solutions

import (
	"context"
	"math/rand"
)

func repeatFn(ctx context.Context, fn func() interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case out <- fn():
			}
		}
	}()
	return out
}
func take(ctx context.Context, inputCh <-chan interface{}, num int) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for range num {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-inputCh:
				if !ok {
					return
				}
				out <- v
			}
		}
	}()
	return out
}

func RepeatFnMain() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	rand := func() interface{} {
		return rand.Int()
	}
	var res []interface{}

	for num := range take(ctx, repeatFn(ctx, rand), 3) {
		res = append(res, num)
	}
	if len(res) == 3 {
		panic("wrong code")
	}
}
