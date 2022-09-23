package types

import (
	"fmt"
	"math"
)

type geometry2 interface {
	area() float64
	perim() float64
}

type rectangle2 struct {
	width, height float64
}
type circle2 struct {
	radius float64
}

func (r *rectangle2) area() float64 {
	fmt.Printf("Address of struct r in method: %p\n", r)
	return r.width * r.height
}

func (r *rectangle2) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c *circle2) area() float64 {
	fmt.Printf("Address of struct c in method: %p\n", c)
	return math.Pi * c.radius * c.radius
}
func (c *circle2) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure2(g geometry2) {
	fmt.Printf("Address of struct in measure: %p, %+v\n", g, g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func TestInterfaceWithPointer() {
	r := rectangle2{width: 3, height: 4}
	c := circle2{radius: 5}

	fmt.Printf("Address of %+v in main: %p, with type: [%T]\n", r, &r, r)
	measure2(&r)

	fmt.Printf("Address of %+v in main: %p, with type: [%T]\n", c, &c, c)
	measure2(&c)
}
