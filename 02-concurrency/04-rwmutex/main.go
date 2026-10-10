package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeConfig struct {
	mu   sync.RWMutex
	data map[string]string
}

func (c *SafeConfig) Get(key string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.data[key]
}

func (c *SafeConfig) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = value
}

func main() {
	var wg sync.WaitGroup
	sf := &SafeConfig{
		data: make(map[string]string),
	}

	for i := range 10 {
		wg.Go(func() {
			val := sf.Get("theme")
			fmt.Printf("reader gorouitine %v: got theme = %q\n", i+1, val)
		})
	}

	for i := range 2 {
		wg.Go(func() {
			sf.Set("theme", "dark")
			fmt.Printf("write goroutine %v set theme to dark\n", i+1)
		})
	}

	wg.Wait()
}
