// package main

// import (
// 	"context"
// 	"errors"
// 	"sync"
// )

// var ErrNotFound = errors.New("key not found")

// type Cache struct {
// 	storage map[string]string
// 	mu      *sync.Mutex
// }

// type ICache interface {
// 	Get(context.Context, string) (string, error)
// 	Set(context.Context, string, string) error
// 	Delete(context.Context, string) error
// }

// func New() *Cache {
// 	return &Cache{
// 		storage: make(map[string]string),
// 		mu:      &sync.Mutex{},
// 	}
// }

// func (c *Cache) Set(_ context.Context, key string, value string) error {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.storage[key] = value
// 	return nil
// }
// func (c *Cache) Delete(_ context.Context, key string) error {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	delete(c.storage, key)
// 	return nil
// }
// func (c *Cache) Get(_ context.Context, key string) (string, error) {
// 	c.mu.Lock()
// 	val, ok := c.storage[key]
// 	c.mu.Unlock()

// 	if !ok {
// 		return "", ErrNotFound
// 	}

// 	return val, nil
// }

// func main() {

// }
