package main

import "fmt"

type rectangle struct {
	length int
	width  int
}

func area(r rectangle) int {
	return r.length * r.width
}

func (r rectangle) area() int { // method read access only
	return r.length * r.width
}

func (r *rectangle) setLength(l int) { // method write access (require a pointer)
	r.length = l
}

func (r *rectangle) setWidth(w int) { // method write access (require a pointer)
	r.width = w
}

func main() {
	r := rectangle{length: 10, width: 5}
	fmt.Println("Area function: ", area(r))
	fmt.Println("Area method: ", r.area())

	r.setLength(100) // modify length
	r.setWidth(50)   // modify width
	fmt.Println("Area method: ", r.area())
}
