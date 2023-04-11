package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	getName() string
	area() float64
	// TODO: Add Perim function
}

type Rectangle struct {
	name           string
	weight, height float64
}

type Circle struct {
	name   string
	radius float64
}

func (r Rectangle) getName() string {
	return r.name
}

func (r Rectangle) area() float64 {
	return r.height * r.weight
}

func (c *Circle) getName() string {
	return c.name
}

func (c *Circle) area() float64 {
	fmt.Printf("Inside Area %p \n", c)
	return c.radius * math.Pi
}

func Measure(geometry Geometry) {
	fmt.Println("Geometry Name - ", geometry.getName())
	fmt.Println("Geometry Area - ", geometry.area())
}

func main() {
	c := Circle{name: "Circle", radius: 2}
	r := Rectangle{name: "Rectangle", weight: 2, height: 2}

	fmt.Printf("Before Circle Measure %p \n", &c)
	Measure(&c)
	Measure(r)

	// Would Mesure(&r) compile / work ?

}
