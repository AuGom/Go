package main

import "fmt"

// go supports methods defined on struct types
type rect struct {
	width, height int
}

func (r *rect) area() int { // area method
	return r.width * r.height
}

func (r rect) perim() int { // perimeter method
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area()) // calls the method using the values in r
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
