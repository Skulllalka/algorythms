package solutions

import (
	"fmt"
)

func mergeSorted(a, b <-chan int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		value1, ok1 := <-a
		value2, ok2 := <-b
		for ok1 && ok2 {
			if value1 < value2 {
				out <- value1
				value1, ok1 = <-a
			} else {
				out <- value2
				value2, ok2 = <-b
			}
		}
		for ok1 {
			out <- value1
			value1, ok1 = <-a
		}
		for ok2 {
			out <- value2
			value2, ok2 = <-b
		}
	}()

	return out
}

func fillChanA(c chan int) {
	c <- 1
	c <- 2
	c <- 4
	close(c)
}
func fillChanB(c chan int) {
	c <- -1
	c <- 4
	c <- 5
	close(c)
}

func MergeSortedMain() {
	a, b := make(chan int), make(chan int)
	go fillChanA(a)
	go fillChanB(b)

	c := mergeSorted(a, b)
	for value := range c {
		fmt.Println(value)
	}

}
