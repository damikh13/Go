package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func main() {
	wgOfTwoExample()
	printSeparator('-', 40)

	var wg sync.WaitGroup
	helloOldStyle := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("hello from %v (old)\n", id)
	}
	helloNewGen := func(id int) {
		fmt.Printf("hello from %v (new)\n", id)
	}
	numfOfCalls := 5
	for i := range numfOfCalls {
		wg.Add(1)
		helloOldStyle(&wg, i+1)
		wg.Go(func() { helloNewGen(i + 1) })
	}
	wg.Wait()
	fmt.Println("code after Wait()")
}

func wgOfTwoExample() {
	var wg sync.WaitGroup
	wg.Go(func() {
		fmt.Println("1st goroutine is doing smt that takes ~ 1 sec")
		time.Sleep(1 * time.Second)
	})
	wg.Go(func() {
		fmt.Println("2nd goroutine is doing smt that takes ~ 0.5 sec")
		time.Sleep(500 * time.Millisecond)
	})
	wg.Wait()

	fmt.Println("code after 2 wg.Wait()")
}

func printSeparator(ch rune, num int) {
	resultingString := strings.Repeat(string(ch), num)
	fmt.Println(resultingString)
}
