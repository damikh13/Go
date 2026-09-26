package main

import (
	"fmt"
)

func SumInt(nums []int) int {
	var total int
	for _, val := range nums {
		total += val
	}
	return total
}

func SumFloat(nums []float64) float64 {
	var total float64
	for _, val := range nums {
		total += val
	}
	return total
}

func SumNums[T int | float64](nums []T) T {
	var total T
	for _, val := range nums {
		total += val
	}
	return total
}

func main() {
	numsInt := []int{1, 2, 3, 4, 5}
	fmt.Println("SumInt(numsInt) =", SumInt(numsInt))

	numsFloat := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println("SumFloat(numsFloat) =", SumFloat(numsFloat))

	fmt.Println("SumNums(numsInt) =", SumNums(numsInt))
	fmt.Println("SumNums(numsFloat) =", SumNums(numsFloat))
}
