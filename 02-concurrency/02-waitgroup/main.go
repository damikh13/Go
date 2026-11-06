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
