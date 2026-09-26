package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) birthday() {
	p.Age++
}

func main() {
	p1 := Person{"John", 30}
	fmt.Println(p1)
	p1.birthday()
	fmt.Println(p1)

	x := 10
	p := &x
	pp := &p
	**pp = 20
	fmt.Println(**pp)
	fmt.Println(x)
}
