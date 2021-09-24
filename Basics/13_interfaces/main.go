package main

import (
	"fmt"
	"math"
)

// Define interface
type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

type GetValue interface {
	printValue() 
}

type Person struct {
	value string
}
type Member struct {
	value string
}

func (M Member) printValue()  {
	fmt.Println("Member "+M.value)
}

func (P Person) printValue()  {
	fmt.Println("Person "+P.value)
}

func (P Person) changeValue()  {
	fmt.Println("Person Changed" + P.value)
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 10, height: 5}
	
	objects := []GetValue{
		Person{"Hello"},
		Member{"Wow"},
	}
	
	for _, value := range objects {
		value.printValue()
	}
	

	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rectangle))
}