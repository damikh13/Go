package main

import (
	"fmt"
	"strings"
)

func main() {
	// default
	for i := 0; i < 10; i++ {
		fmt.Print(i)
		if i != 9 {
			fmt.Print(" ")
		} else {
			fmt.Print("\n")
		}
	}
	fmt.Println(strings.Repeat("-", 50))

	// modern default
	for i := range 10 {
		fmt.Println(i)
	}
	fmt.Println(strings.Repeat("-", 50))

	// while
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	fmt.Println(strings.Repeat("-", 50))

	// while true
	cond := true
	for {
		// work
		if cond {
			break
		}
	}

	// for range
	nums := []int{1, 2, 3, 4, 5}
	for i, v := range nums { // i, _ или _, v, или <ничего>
		fmt.Println("i =", i, "v =", v)
		v = 10
	}
	for i, v := range nums {
		fmt.Println("i =", i, "v =", v)
	}

	// for map
	mp := map[string]int{"a": 1, "b": 2}
	for k, v := range mp {
		fmt.Printf("key = %q, val = %v.\n", k, v)
	}

	// for string
	str := "hello"
	for i, c := range str {
		fmt.Printf("i = %v, c = %v (%c).\n", i, c, c)
	}

	// break, continue, inner/outer!
outer:
	for i := 0; i < 5; i++ {
	middle:
		for j := 0; j < 5; j++ {
			if i == 1 {
				continue middle
			}

			if i == 3 {
				break outer
			}

			fmt.Println(i, j)
		}
	}
}
