package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	area() float64
	perimeter() float64
}

type Rect struct {
	 height, width float64
}

func (r Rect) area() float64 {
	return r.height * r.width
}

func (r Rect) perimeter() float64 {
	return 2 * r.width + 2 * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g Geometry) {
	fmt.Printf("%T %v\n", g, g)
	fmt.Printf("Area %.2f\n", g.area())
	fmt.Printf("Perimeter %.2f\n", g.perimeter())
}

func main() {
	r := Rect{width: 10, height: 10}
	c := Circle{radius: 25}

	measure(r)
	measure(c)
}
