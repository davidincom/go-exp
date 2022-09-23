package types

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rectangle struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	fmt.Printf("Address of struct r in method: %p\n", &r)
	return r.width * r.height
}

func (r rectangle) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	fmt.Printf("Address of struct r in method: %p\n", &c)
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Printf("Address of struct in measure: %p, %+v\n", &g, g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func TestInterfaces() {
	r := rectangle{width: 3, height: 4}
	c := circle{radius: 5}

	fmt.Printf("Address of %+v in main: %p, with type: [%T]\n", r, &r, r)
	measure(r)

	fmt.Printf("Address of %+v in main: %p, with type: [%T]\n", c, &c, c)
	measure(c)
}
