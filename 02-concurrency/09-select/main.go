package main

import (
	"fmt"
)

func gen(n int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := range n {
			ch <- i
		}
	}()

	return ch
}

func main() {
	a := gen(2)
	b := gen(1)

	for a != nil || b != nil { // at least one alive
		select {
		case v, ok := <-a:
			if !ok { // closed
				a = nil
				continue
			}

			fmt.Printf("got %v from a.\n", v)
		case v, ok := <-b:
			if !ok {
				b = nil
				continue
			}

			fmt.Printf("got %v from b.\n", v)
		}
	}
}
