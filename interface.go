package main

import (
	"fmt"
	"math"
)

func main() {

	var c1 shape = circle{radius: 5}
	var r1 shape = rectangle{width: 5, height: 10}

	fmt.Printf("type of c1 : %T \n", c1)
	fmt.Printf("type of r1 : %T \n", r1)

	// panggil method

	fmt.Println("circle area :", c1.area())
	fmt.Println("circle perimeter :", c1.perimeter())

	fmt.Println("rectangle area :", r1.area())
	fmt.Println("rectangle perimeter :", r1.perimeter())

	print("Circle", c1)
	print("Rectangle", r1)

}

func print(t string, s shape) {
	fmt.Printf("%s area %v \n", t, s.area())
	fmt.Printf("%s perimeter %v \n", t, s.perimeter())
}

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

func (r rectangle) area() float64 {
	return r.width * r.width
}
func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}
