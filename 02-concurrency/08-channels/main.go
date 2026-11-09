package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// unbuffChanExample()
	// buffChanExample()
	// genChanExample()
	multSendersExample()
}

func unbuffChanExample() {
	ch := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Printf("sending value %v to channel\n", 1)
		ch <- 1
	}()

	fmt.Printf("waiting for sender...\n")
	val := <-ch
	fmt.Printf("got value %v from channel\n", val)
}

func buffChanExample() {
	ch := make(chan int, 2)

	fmt.Print("sending 1 to the channel...")
	ch <- 1
	fmt.Println(" sent!")

	fmt.Print("sending 2 to the channel...")
	ch <- 2
	fmt.Println(" sent!")

	go func() {
		fmt.Print("extracting smt from the channel...")
		val := <-ch
		fmt.Printf(" got %v.\n", val)
	}()

	fmt.Print("sending 3 to the channel...")
	ch <- 3
	fmt.Println(" sent!")
}

func genChanExample() {
	gen := func(n int) <-chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			for i := range n {
				ch <- i
				fmt.Printf("put %v.\n", i)
			}
		}()
		return ch
	}

	ch := gen(3)
	for v := range ch {
		fmt.Printf("got %v.\n", v)
	}
}

func multSendersExample() {
	ch := make(chan string)
	var wg sync.WaitGroup

	for i := range 3 {
		wg.Go(func() {
			for j := range 3 {
				val := fmt.Sprintf("i = %v, j = %v", i, j)
				ch <- val
			}
		})
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for val := range ch {
		fmt.Printf("val = %q.\n", val)
	}
}
