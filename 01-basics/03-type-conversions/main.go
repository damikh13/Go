package main

import (
	"fmt"
	"strconv"
)

func main() {
	var fl float64 = -3.9
	var in int = int(fl)
	var uin uint = uint(fl)
	fmt.Printf("fl = %v, in = %v, u = %v.\n", fl, in, uin)

	// Можно с переменными, но нельзя с константами?
	// fmt.Printf("fl = %v, in = %v, u = %v.\n", 3.9, int(3.9), uint(3.9))

	var big int = 300
	var small int8 = int8(big)
	fmt.Printf("small = %v.\n", small)

	fmt.Println(string(65))       // "A"
	fmt.Println(strconv.Itoa(65)) // "65"
}
