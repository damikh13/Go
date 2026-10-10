package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var value int

	increment := func() {
		mu.Lock()
		defer mu.Unlock()
		value++
		fmt.Printf("incrementing: %v\n", value)
	}
	decrement := func() {
		mu.Lock()
		defer mu.Unlock()
		value--
		fmt.Printf("decrementing: %v\n", value)
	}

	for range 5 {
		wg.Go(increment)
	}
	for range 5 {
		wg.Go(decrement)
	}

	wg.Wait()
	fmt.Printf("the result after all arithmetics: %v\n", value)
}
