package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// resultsChannelExample()
	// resultsSliceAtomicIdx()
	resultsSliceIdxChannel()
}

func resultsChannelExample() {
	jobs := make(chan int)
	results := make(chan int)
	var wg sync.WaitGroup

	for range 3 {
		wg.Go(func() {
			for j := range jobs {
				results <- j * j
			}
		})
	}

	go func() {
		defer close(jobs)
		for j := range 9 {
			jobs <- j
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Printf("result: %v.\n", r)
	}
}

func resultsSliceAtomicIdx() {
	input := []int{1, 2, 3, 4, 5}
	N := len(input)
	results := make([]int, N)
	var wg sync.WaitGroup

	var atomIDX atomic.Int64
	for range N {
		wg.Go(func() {
			idx := atomIDX.Add(1) - 1
			results[idx] = input[idx] * 10
		})
	}

	wg.Wait()
	for _, r := range results {
		fmt.Println(r)
	}
}

func resultsSliceIdxChannel() {
	input := []int{1, 2, 3, 4, 5}
	N := 9
	jobs := make(chan int)
	results := make([]int, N)
	var wg sync.WaitGroup

	for range 3 {
		wg.Go(func() {
			for j := range jobs {
				results[j] = input[j] * 10
			}
		})
	}

	for j := range N {
		jobs <- j
	}
	close(jobs)

	wg.Wait()

	for i, v := range results {
		fmt.Printf("results[%v] = %v.\n", i, v)
	}
}

func process(
	ctx context.Context,
	input []int,
	workersNum int,
	fn func(context.Context, int) int,
) ([]int, error) {
	results := make([]int, len(input))
	jobs := make(chan int)
	var wg sync.WaitGroup

	for range workersNum {
		wg.Go(func() {
			for j := range jobs {
				results[j] = fn(ctx, input[j])
			}
		})
	}

feed:
	for idx := range input {
		if err := ctx.Err(); err != nil {
			break
		}

		select {
		case jobs <- idx:
		case <-ctx.Done():
			break feed
		}
	}
	close(jobs)

	wg.Wait()
	return results, ctx.Err()
}
