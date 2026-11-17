package main

import (
	"context"
	"fmt"
	"time"
)

func query(jobTime, timeoutTime time.Duration) {
	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx, timeoutTime)
	defer cancel()

	done := make(chan struct{})
	go func() {
		defer close(done)
		time.Sleep(jobTime)
	}()

	select {
	case <-done:
		fmt.Println("retrieved object from API!")
	case <-ctxWithTimeout.Done():
		fmt.Println("timeout expired")
	}
}

func main() {
	query(2*time.Second, 1*time.Second) // timeout expired
	query(1*time.Second, 2*time.Second) // timeout expired
}
