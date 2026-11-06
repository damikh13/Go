package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	basicPoolExample()
	goroutinesPoolExample()
}

func basicPoolExample() {
	pool := sync.Pool{
		New: func() any {
			fmt.Println("allocating new []byte slice")
			return make([]byte, 1024)
		},
	}

	obj := pool.Get().([]byte)
	fmt.Printf("got object from pool of length: %v\n", len(obj))

	pool.Put(obj)

	reusedObj := pool.Get().([]byte)
	fmt.Printf("got reused object from the pool of length: %v\n", len(reusedObj))

	pool.Put(reusedObj)
}

func goroutinesPoolExample() {
	pool := sync.Pool{
		New: func() any {
			fmt.Print(".")
			return make([]byte, 1024)
		},
	}

	var wg sync.WaitGroup
	for range 1000 {
		wg.Go(func() {
			obj := pool.Get().([]byte)
			fmt.Print("-")
			time.Sleep(100 * time.Millisecond)
			pool.Put(obj)
		})

		time.Sleep(10 * time.Millisecond)
	}

	wg.Wait()
}
