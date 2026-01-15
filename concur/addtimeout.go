package concur

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func AddTimeOut() {
	chanForResp := make(chan resp)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
	go RPCCall(ctx, chanForResp)
	fmt.Println(<-chanForResp)
}

type resp struct {
	id  int
	err error
}

func RPCCall(ctx context.Context, ch chan<- resp) {
	select {
	case <-ctx.Done():
		ch <- resp{
			id:  0,
			err: errors.New("timeout expired"),
		}
	case <-time.After(time.Second * time.Duration(rand.Intn(5))):
		ch <- resp{
			id:  rand.Int(),
			err: nil,
		}
	}

}
