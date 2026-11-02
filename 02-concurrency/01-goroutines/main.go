package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

func main() {
	loopGoroutineExample()
	printSeparator('-', 40)

	goroutineSizeBookExample()
	printSeparator('-', 40)

	goroutineWgExample()
	printSeparator('-', 40)

	panicInSeparateGoroutineExample()
	printSeparator('-', 40)
}

func goroutineWgExample() {
	var wg sync.WaitGroup
	wg.Go(func() {
		fmt.Println("hello from forked function!")
	})
	wg.Wait()

	fmt.Println("main function continues")
}

func loopGoroutineExample() {
	fmt.Println("loopGoroutineExample()")
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Go(func() {
			fmt.Println(salutation)
		})
	}
	wg.Wait()
}

func goroutineSizeBookExample() {
	fmt.Println("goroutineSizeBookExample()")
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}
	var c <-chan any
	var wg sync.WaitGroup
	noop := func() { wg.Done(); <-c }
	const numGoroutines = 1e4
	wg.Add(numGoroutines)
	before := memConsumed()
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := memConsumed()
	fmt.Printf("%.3fkb\n", float64(after-before)/numGoroutines/1000)
}

func panicInSeparateGoroutineExample() {
	fmt.Println("panicInSeparateGoroutineExample()")
	var wg sync.WaitGroup
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("caught:", r)
		}
	}()

	wg.Go(func() {
		panic("boom")
	})
	wg.Wait()

	fmt.Println("main func done")
}

func printSeparator(ch rune, num int) {
	resultingString := strings.Repeat(string(ch), num)
	fmt.Println(resultingString)
}
