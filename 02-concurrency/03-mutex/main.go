package main

import (
	"fmt"
	"sync"
)

func main() {
	incrementsDecrementsExample()
	ticketsExample()
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

func ticketsExample() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	tickets := 500

	buyTicket := func(userID int, tickets *int) {
		mu.Lock()
		defer mu.Unlock()
		if *tickets > 0 {
			*tickets--
			fmt.Printf("user %v successfully purchased a ticket.\n", userID)
		} else {
			fmt.Printf("user %v found no ticket.\n", userID)
		}
	}

	for i := range 2000 {
		wg.Go(func() {
			buyTicket(i, &tickets)
		})
	}

	wg.Wait()
	fmt.Printf("remaining tickets: %v.\n", tickets)
}
