package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {
	s1 := fmt.Sprintf("I am %v year old!", 21)
	fmt.Println(s1) // I am 21 year old!
	fmt.Printf("The lawyer's name is %v", "Saul Goodman.\n")
	fmt.Printf("The foolowing sentence is either %v or %v.\n", true, false)

	fmt.Printf("His name is: %s", "Slim Shady.\n")

	fmt.Println("123" + "456")

	fmt.Printf("I am %d year old.\n", 21)

	const pi = 3.14159
	fmt.Printf("Pi is: %.3f.\n", pi)

	fmt.Printf("He said: %q.\n", "quote-quote-quote")

	n := 42
	p_to_n := &n
	fmt.Printf("%p\n", p_to_n)
	fmt.Printf("%v\n", p_to_n)
	fmt.Println(p_to_n)

	pt := Point{3, 4}
	fmt.Printf("%v\n", pt)  // {3 4}
	fmt.Printf("%+v\n", pt) // {X:3 Y:4}
	fmt.Printf("%#v\n", pt) // main.Point{X:3, Y:4}

	pointer_to_pt := &pt
	fmt.Printf("%v\n", pointer_to_pt) // &{3 4}
	fmt.Printf("%p\n", pointer_to_pt) // 0x605168c7c110
}
