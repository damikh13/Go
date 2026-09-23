package main

import (
	"fmt"
)

type Grid[T any] struct {
	data []T
	rows int
	cols int
}

func (g *Grid[T]) At(i, j int) T     { return g.data[i*g.cols+j] }
func (g *Grid[T]) Set(i, j int, v T) { g.data[i*g.cols+j] = v }

func main() {
	var myArr [4]int = [4]int{1, 2, 3, 4}
	myArr2 := [4]int{5, 6, 7, 8}
	myArr3 := [...]int{9, 10, 11, 12}
	fmt.Println("myArr:", myArr)
	fmt.Println("myArr2:", myArr2)
	fmt.Println("myArr3:", myArr3)

	var s1 []int
	s2 := []int{1, 2, 3}
	s3 := make([]int, 5)
	s4 := make([]int, 5, 10)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)
	fmt.Println("s4:", s4)

	s5 := []int{0, 1, 2, 3, 4, 5, 6}
	s6 := s5[1:4] // [1 2 3]
	s6[0] = 1000
	fmt.Println(s6)
	fmt.Println(s5)

	s7 := s5[1:3:4]
	fmt.Println(s7)

	s8 := []int{0, 1, 2, 3, 4}
	s9 := append(s8, 5)
	fmt.Println("s8:", s8)
	fmt.Println("s9:", s9)

	s10 := s5[1:3]
	fmt.Println("s10:", s10)
	s11 := append(s10, 11, 22)
	fmt.Println("s11:", s11)
	fmt.Println("s5:", s5)

	s12 := s5[1:3:5]
	fmt.Println("s12:", s12) // [1000 2]
	s12 = append(s12, 52, 54)
	s12[0] = 500
	fmt.Println("s12:", s12) // 500 2 52 54]
	fmt.Println("s5:", s5)   // [0 1000 2 11 22 5 6]

	src := []int{0, 1, 2, 3, 4, 5}
	dst := make([]int, len(src))
	copy(src, dst)

	var emptySlice []int
	fmt.Println("emptySlice == nil:", emptySlice == nil)
	emptySlice2 := []int{}
	fmt.Println("emptySlice2 == nil:", emptySlice2 == nil)

	rows := 5
	cols := 3
	grid1 := make([][]int, rows)
	for i := range grid1 {
		grid1[i] = make([]int, cols)
	}

	counter := 1
	for i := range rows {
		for j := range cols {
			grid1[i][j] = counter
			counter++
		}
	}

	fmt.Println(grid1)
}
