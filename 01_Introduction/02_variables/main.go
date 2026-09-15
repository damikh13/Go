package main

import "fmt"

/*
default values!
:= <- available only inside functions
to know a type: Printf("%T") +
a, b = 10, 2.0
int(2.6) -> 2
const's
const full_name = first_name + " " + last_name <- вычислится во время компиялции!

bool
string

int	<- 32/64 => 1 bit for '+', '-', remaning for num
uint <- 32/64 => 2

int8, ..., int64
uint8, ..., uint64

byte = int8
rune = int32

float32
float64 <- default

complex64, complex128
*/

func main() {
	// username := "some_username"
	// password := "1201202101"

	// var another_variable string
	// var third_variable int
	//
	// another_variable = "15"
	// third_variable = 12

	// fmt.println("authorization: basic", username+":"+password)
	// fmt.println(another_variable)
	// fmt.println(third_variable)

	// var age int16
	// var name string
	// var cond bool
	// fmt.Println(age, name, cond)

	// age_ptr := &cond
	// fmt.Printf("type(%d) = %T", age_ptr, age_ptr)

	f := 12.34
	fmt.Printf("type(f) = %T\n", f)
}
