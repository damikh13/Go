package main

import (
	"context"
	"fmt"
	"sync"
	"time"
	// "time"
)

func main() {
	channelCondExample()
	fmt.Println("--------")
	simpleCondExample()
	fmt.Println("--------")

	gate := NewGate()
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	for id := range 3 {
		wg.Go(func() {
			for step := 1; ; step++ {
				gate.WaitIfPaused() // runs cond.Wait() inside for paused loop

				select {
				case <-ctx.Done():
					return
				default:
				}

				fmt.Printf("worker %d: step %d\n", id+1, step)
				time.Sleep(500 * time.Millisecond)
			}
		})
	}

	time.Sleep(1 * time.Second)
	fmt.Println(">>> PAUSE")
	gate.Pause()

	time.Sleep(1 * time.Second)

	fmt.Println(">>> RESUME")
	gate.Resume()

	time.Sleep(1 * time.Second)
	fmt.Println(">>> PAUSE again")
	gate.Pause()

	time.Sleep(1 * time.Second)
	fmt.Println(">>> RESUME again")
	gate.Resume()

	time.Sleep(1 * time.Second)
	cancel()
	wg.Wait()
}

func channelCondExample() {
	ch := make(chan struct{})
	var wg sync.WaitGroup

	for i := range 3 {
		wg.Go(func() {
			<-ch
			fmt.Printf("goroutine %v started working.\n", i)
		})
	}

	fmt.Println("sleeping for 500 ms...")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("... done.")

	close(ch)
	wg.Wait()
}

func simpleCondExample() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	ready := false

	for i := range 3 {
		wg.Go(func() {
			mu.Lock()
			for !ready {
				cond.Wait() // unlock, sleep... ...wake, lock
			}
			mu.Unlock()

			fmt.Printf("goroutine %v started working.\n", i)
		})
	}

	fmt.Println("sleeping for 500 ms...")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("... done.")

	mu.Lock()
	ready = true
	mu.Unlock()

	cond.Broadcast()
	wg.Wait()
}
