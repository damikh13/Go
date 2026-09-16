package main

import (
	"fmt"
	"strings"
)

func main() {
	height := 10
	fmt.Println("height =", height)
	if height > 8 {
		fmt.Println("you are super tall!")
	} else if height > 4 {
		fmt.Println("you are tall enough")
	} else {
		fmt.Println("it's okay")
	}

	fmt.Println(strings.Repeat("-", 50))

	if weight := 10; weight > 10 {
		fmt.Println("obese")
	} else {
		fmt.Println("scrawny")
	}

	// fmt.Println(weight) // error
	fmt.Println(strings.Repeat("-", 50))

	day := "Mon"
	switch day {
	case "Mon", "Tue", "Wed", "Thu", "Fri":
		fmt.Println("work day :(")
		fallthrough
	case "Sat", "Sun":
		fmt.Println("weekend day :)")
	default:
		fmt.Println("unknown")
	}

	fmt.Println(strings.Repeat("-", 50))

	switch x := 10; {
	case x > 10:
		fmt.Println("big")
	default:
		fmt.Println("small")
	}

	fmt.Println(strings.Repeat("-", 50))

	switch day := "alskdjasld"; day {
	case "Mon", "Tue", "Wed", "Thu", "Fri":
		fmt.Println("work day :(")
	case "Sat", "Sun":
		fmt.Println("weekend day :)")
	default:
		fmt.Println("unknown")
	}

	fmt.Println(day)

	fmt.Println(strings.Repeat("-", 50))
}
