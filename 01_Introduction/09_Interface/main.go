package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct{ W, H float64 }

func (r Rectangle) Area() float64      { return r.W * r.H }
func (r Rectangle) Perimeter() float64 { return 2 * (r.W + r.H) }

type Circle struct{ R float64 }

func (c Circle) Area() float64      { return math.Pi * c.R * c.R }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.R }

func printShapeInfo(s Shape) {
	fmt.Printf("area = %.3f, perimeter = %.3f\n", s.Area(), s.Perimeter())
}

type Itf1 interface {
	method1() int
	method2() int
}

type Cl1 struct{}

func (c *Cl1) method1() int {
	fmt.Println("method1")
	return 1
}
func (c Cl1) method2() int {
	fmt.Println("method2")
	return 2
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d)", p.Name, p.Age)
}

type Number interface {
	~int | ~int64 | ~float64
}

func sum[T Number](nums ...T) T {
	var total T
	for _, v := range nums {
		total += v
	}
	return total
}

func main() {
	rect := Rectangle{3, 4}
	printShapeInfo(rect)
	var rect2 Shape = Rectangle{2, 4}
	printShapeInfo(rect2)
	fmt.Printf("type of rect = %T\n", rect)
	fmt.Printf("type of rect2 = %T\n", rect)
	circle := Circle{3}
	printShapeInfo(circle)

	cl := Cl1{}
	cl.method1()
	cl.method2()

	cl2 := new(Cl1)
	cl2.method1()
	cl2.method2()

	var cl3 Itf1 = new(Cl1)
	// var cl3 Itf1 = Cl1{} // error!
	cl3.method1()
	cl3.method2()

	p := Person{"John", 30}
	fmt.Println(p) // John (30)

	var circ2 Shape = Circle{4}
	if c, ok := circ2.(Circle); ok {
		fmt.Printf("type(circle = %v) = Circle\n", c)
	}

	nums := []float64{1, 2.5, 3}
	fmt.Println(sum(1, 2.5, 3))
	fmt.Println(sum(nums...))
}
