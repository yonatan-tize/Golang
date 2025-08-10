package main

import (
	"fmt"
)

type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r Rectangle) Perimeter() float64 {
	return r.length*2 + r.width*2
}

func main() {

	name := Rectangle{
		length: 4,
		width:  2,
	}

	fmt.Println(name.Perimeter())
	fmt.Println(name.Area())

}
