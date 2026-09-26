package main

import (
	"fmt"
)

type Person struct {
	Name    string
	Age     int
	IsAdmin bool
}

func birthday(p *Person) {
	p.Age++
}

func (p *Person) birthday() {
	p.Age++
}

type MyStruct struct {
	field1 int
	Field2 int
}

func (m MyStruct) Field1() int {
	return m.field1
}

// func (m MyStruct) Field2() int {
// 	return m.Field2
// }

type Celsius float64

func (c Celsius) ToFarenheit() float64 {
	return float64(c)*5/9 + 32
}

type Address struct {
	City, State string
}

type Employee struct {
	Person
	Address
	Role string
}

func main() {
	p1 := Person{}    // type = Person
	p2 := new(Person) // type = *Person

	fmt.Printf("%+v\n", p1)
	fmt.Printf("%#v\n", p2)

	fmt.Printf("%T\n", p1)
	fmt.Printf("%T\n", p2)

	p3 := Person{Name: "Alice", Age: 25}
	p4 := Person{"Bob", 30, true}
	var p5 Person
	p5.Name = "Patrick"
	p5.Age = 44

	fmt.Printf("%+v\n", p3)
	fmt.Printf("%+v\n", p4)
	fmt.Printf("%+v\n", p5)

	fmt.Printf("p3 before birthday(&p3): %+v\n", p3)
	birthday(&p3)
	fmt.Printf("p3 after birthday(&p3): %+v\n", p3)

	fmt.Printf("p4 before p4.Birthday() %+v\n", p4)
	p4.birthday()
	fmt.Printf("p4 after p4.Birthday() %+v\n", p4)

	m1 := MyStruct{3, 4}
	fmt.Printf("%+v\n", m1)
	fmt.Println(m1.field1)
	fmt.Println(m1.Field1())
	fmt.Println(m1.Field2)

	c := Celsius(5.6)
	fmt.Printf("%v\n", c)        // 5.6
	fmt.Println(c.ToFarenheit()) // 35.1111...

	e := Employee{
		Person{"John", 30, true},
		Address{"Miami", "Florida"},
		"Admin",
	}
	fmt.Printf("%+v\n", e)
}
