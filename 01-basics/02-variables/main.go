package main

import "fmt"

func main() {
	username := "some_username"
	password := "1201202101"

	var another_variable string
	var third_variable int

	another_variable = "15"
	third_variable = 12

	fmt.Println("authorization: basic", username+":"+password)
	fmt.Println(another_variable)
	fmt.Println(third_variable)

	var age int16
	var name string
	var cond bool
	fmt.Println(age, name, cond)

	age_ptr := &cond
	fmt.Printf("type(%d) = %T", age_ptr, age_ptr)

	f := 12.34
	fmt.Printf("type(%v) = %T\n", f, f)

	const (
		Monday = iota // 0
		Tuesday
		Wednesday
	)
	fmt.Println(Wednesday)
}
