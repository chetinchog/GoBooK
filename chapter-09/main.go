package main

import (
	"fmt"
	"log"
	"math"
)

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func rectangleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x1, y2)
	w := distance(x1, y1, x2, y1)
	return l * w
}

func circleArea(x, y, r float64) float64 {
	return math.Pi * r * r
}

func circleAreaS(c *Circle) float64 {
	return math.Pi * math.Pow(c.r, 2)
}

// Circle struct
type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

func (c *Circle) perimeter() float64 {
	return math.Pi * c.r * 2
}

// Rectangle struct
type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.y2, r.x1)
	return l * w
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.y2, r.x1)
	return l*2 + w*2
}

// Person strcut
type Person struct {
	// Name of the Person
	Name string
}

// Talk to print Person's name
func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

// Android struct
type Android struct {
	Person
	// Model code
	Model string
}

// Shape interface
type Shape interface {
	area() float64
	perimeter() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, shape := range shapes {
		area *= shape.area()
	}
	return area
}

// MultiShape struct
type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func (m *MultiShape) perimeter() float64 {
	var perimeter float64
	for _, s := range m.shapes {
		perimeter += s.perimeter()
	}
	return perimeter
}

func main() {
	log.Println("C9!")

	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10
	var cx, cy, cr float64 = 0, 0, 5

	fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
	fmt.Println(circleArea(cx, cy, cr))

	fmt.Println("-")

	var varC Circle
	fmt.Println(varC)

	c := new(Circle)
	fmt.Println(c)
	*c = Circle{x: 0, y: 0, r: 5}
	fmt.Println(c)
	c = &Circle{0, 0, 5}
	fmt.Println(c)

	cn := Circle{0, 0, 5}
	fmt.Println(cn)

	fmt.Println("-")
	c.x = 10
	c.y = 10
	c.r = 5
	fmt.Println(circleAreaS(c))
	fmt.Println(circleAreaS(&cn))

	fmt.Println("-")
	fmt.Println(c.area())
	fmt.Println(cn.area())

	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())

	fmt.Println("-")
	a := Android{Person: Person{Name: "Pepe"}, Model: "CYT&"}
	a = Android{Person{"Pepe"}, "CYT&"}
	a.Person.Talk()
	a.Talk()

	fmt.Println(totalArea(&cn, &r))

	ms := new(MultiShape)
	ms.shapes = append(ms.shapes, &cn, &r)

	fmt.Println(ms.area())
	fmt.Println(ms.perimeter())
}
