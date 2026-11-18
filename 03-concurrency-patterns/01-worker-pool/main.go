package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// jobs := make(chan int)
	// results := make(chan int)
	// var wg sync.WaitGroup

	// for range 3 {
	// 	wg.Go(func() {
	// 		for j := range jobs {
	// 			results <- j * j
	// 		}
	// 	})
	// }

	// go func() {
	// 	defer close(jobs)
	// 	for j := range 9 {
	// 		jobs <- j
	// 	}
	// }()

	// go func() {
	// 	wg.Wait()
	// 	close(results)
	// }()

	// for r := range results {
	// 	fmt.Printf("result: %v.\n", r)
	// }
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
