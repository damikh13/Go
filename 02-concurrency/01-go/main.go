// package main

// import (
// 	"fmt"
// 	"sync"
// 	// "time"
// )

// func printNumbers(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := range 5 {
// 		fmt.Println("printNumers():", i)
// 		// time.Sleep(time.Millisecond * 250)
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go printNumbers(&wg)
// 	for i := range 5 {
// 		fmt.Println("main():", i)
// 		// time.Sleep(time.Millisecond * 250)
// 	}
// 	wg.Wait()
// }

package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := range 5 {
		fmt.Println("printNumbers():", i)
		time.Sleep(time.Millisecond * 250)
	}
}

func main() {
	go printNumbers()
	for i := range 5 {
		fmt.Println("main():", i)
		time.Sleep(time.Millisecond * 250)
	}
	// time.Sleep(time.Millisecond * 500)
}
