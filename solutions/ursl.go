package solutions

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const timeoutLimit = 100

type Result struct {
	msg string
	err error
}

func fakeDownload(url string) Result {
	rnd := rand.Intn(100)
	time.Sleep(time.Duration(rnd) * time.Millisecond)
	if rnd > timeoutLimit {
		return Result{
			err: errors.New(fmt.Sprintf("failed to download data from %s: timeout", url)),
		}
	}
	return Result{
		msg: fmt.Sprintf("downloaded data from %s url", url),
	}
}

func download(urls []string) ([]string, error) {
	res := make([]string, len(urls))
	var err error

	ch := make(chan Result)
	wg := sync.WaitGroup{}

	for _, url := range urls {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- fakeDownload(url)
		}()
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for value := range ch {
		if value.err != nil {
			err = errors.Join(err, value.err)
		}
		res = append(res, value.msg)
	}

	return res, err
}

func UrlsMain() {
	msgs, err := download([]string{
		"https://example.com/e25e26d3-6aa3-4d79-9ab4-fc9b71103a8c.xml",
		"https://example.com/e25e26d3-6aa3-4d79-9ab4-fc9b71103a8c.xml",
		"https://example.com/e25e26d3-6aa3-4d79-9ab4-fc9b71103a8c.xml",
		"https://example.com/e25e26d3-6aa3-4d79-9ab4-fc9b71103a8c.xml",
		"https://example.com/e25e26d3-6aa3-4d79-9ab4-fc9b71103a8c.xml",
		"https://example.com/e25e26d3-6aa3-4d79-9ab4-fc9b71103a8c.xml",
		"https://example.com/e25e26d3-6aa3-4d79-9ab4-fc9b71103a8c.xml",
		"https://example.com/e25e26d3-6aa3-4d79-9ab4-fc9b71103a8c.xml",
		"https://example.com/e25e26d3-6aa3-4d79-9ab4-fc9b71103a8c.xml",
		"https://example.com/e25e26d3-6aa3-4d79-9ab4-fc9b71103a8c.xml",
	})

	if err != nil {
		panic(err)
	}
	fmt.Println(msgs)
}
