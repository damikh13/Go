package main

import (
	"fmt"
)

func sub(x int, y int) int {
	return x - y
}

func sum(x, y int) int {
	return x + y
}

func mul(nums ...int) int {
	totalProduct := 1
	for _, v := range nums {
		fmt.Printf("%v * %v = ", totalProduct, v)
		totalProduct *= v
		fmt.Printf("%v\n", totalProduct)
	}
	fmt.Printf("your total is: %v.\n", totalProduct)
	return totalProduct
}

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, fmt.Errorf("can't divide by 0")
	}
	return x / y, nil
}

func smt(summ int) (x, y int) {
	x = summ * x / 9
	y = summ - x
	return
}

func PrintValues(nums ...int) {
	for _, v := range nums {
		fmt.Println(v)
	}
}

func sortingBenchmark(sortingAlg func(nums ...int)) {
	// start of a timer
	sortingAlg()
	// end of timer
}

func changeValue(x *int) {
	*x = 50
}

func getFirstInArray[Type any](items []Type) Type {
	return items[0]
}

func outer() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

type sortingAlg func(nums ...int)

type Person struct {
	name string
}

func (p Person) printName() {
	fmt.Println(p.name)
}

func test(a, b float64) (result float64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered: %v", r)
		}
	}()

	res := a / b

	panic("aaa paniccc")

	fmt.Println("here")

	return res, err
}

func main() {
	fmt.Println(sub(5, 8))
	fmt.Println(sum(5, 8))
	nums := []int{1, 2, 3, 4, 5}
	mul(4, 8, 12)
	mul(nums...)
	div_res1, err1 := divide(5, 0)
	div_res2, err2 := divide(8, 2)
	fmt.Printf("div(%v, %v) = %v, %v\n", 5, 0, div_res1, err1)
	fmt.Printf("div(%v, %v) = %v, %v\n", 8, 2, div_res2, err2)
	fmt.Println(smt(50))
	PrintValues(1, 2, 3)

	add := func(a, b float64) float64 {
		return a + b
	}
	fmt.Println(add(10, 20))   // 30
	fmt.Println(add(10.5, 20)) // 30.5

	x := 10
	fmt.Printf("x = %v", x)
	changeValue(&x)
	fmt.Printf(", after changeValue(x) x = %v\n", x)

	func(a, b int) {
		fmt.Println(a, b)
	}(5, 6)

	out := outer()
	fmt.Println(out())
	fmt.Println(out())
	fmt.Println(out())
	fmt.Println(out())

	person := Person{"Patrick"}
	person.printName()

	test(10, 0)
}
