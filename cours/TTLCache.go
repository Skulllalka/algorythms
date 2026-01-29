package main

import (
	"context"
	"errors"
	"sync"
	"time"
)

var ErrNotFound = errors.New("key not found")

type elem struct {
	value    string
	exp_date time.Time
}

type Cache struct {
	storage map[string]elem
	mu      *sync.Mutex
	TTL     time.Duration
	done    chan struct{}
}

type ICache interface {
	Get(context.Context, string) (string, error)
	Set(context.Context, string, string) error
	Delete(context.Context, string) error
}

func New(ttl time.Duration) *Cache {
	cache := &Cache{
		storage: make(map[string]elem),
		mu:      &sync.Mutex{},
		TTL:     ttl,
		done:    make(chan struct{}),
	}
	cache.clearByTTL()
	return cache

}

func (c *Cache) Set(_ context.Context, key string, value string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	el := elem{
		value:    value,
		exp_date: time.Now().Add(c.TTL),
	}
	c.storage[key] = el
	return nil
}
func (c *Cache) Delete(_ context.Context, key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.storage, key)
	return nil
}
func (c *Cache) Get(_ context.Context, key string) (string, error) {
	c.mu.Lock()
	el, ok := c.storage[key]
	c.mu.Unlock()

	if !ok {
		return "", ErrNotFound
	}

	if el.exp_date.Before(time.Now()) {
		c.delete(key)
		return "", ErrNotFound
	}

	return el.value, nil
}

func (c *Cache) delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.storage, key)
}

func (c *Cache) clearByTTL() {
	ticker := time.NewTicker(10 * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
			case <-c.done:
				return
			}
		}
	}()

}

func (c *Cache) Stop() {
	close(c.done)
}

func (c *Cache) clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, value := range c.storage {
		if value.exp_date.Before(time.Now()) {
			delete(c.storage, key)
		}
	}

}
