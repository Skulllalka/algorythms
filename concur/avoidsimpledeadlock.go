package concur

import "fmt"

func AvoidDeadlock() {
	ch := make(chan int)

	select {
	case res := <-ch:
		fmt.Println("result is", res)
	default:
		fmt.Println("default value")
	}
}
