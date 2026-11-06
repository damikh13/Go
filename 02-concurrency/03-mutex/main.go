package main

import (
	"fmt"
	"sync"
)

func main() {
	incrementsDecrementsExample()
}

func incrementsDecrementsExample() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var value int

	increment := func() {
		mu.Lock()
		defer mu.Unlock()
		value++
	}

	decrement := func() {
		mu.Lock()
		defer mu.Unlock()
		value--
	}

	for range 5 {
		wg.Go(increment)
		wg.Go(decrement)
	}

	wg.Wait()
	fmt.Printf("result after all increments & decrements: %v\n", value) // 0
}
