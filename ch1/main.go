package main

import (
	"fmt"
)

type shape interface {
	perimeter() int
	area() int
}

type rectangle struct {
	l int
	w int
}

func (r rectangle) area() int {
	return r.l * r.w
}

func (r rectangle) perimeter() int {
	return r.l * 2 + r.w * 2
}



type square struct{
	l int
}

func (s square) area() int{
	return s.l * s.l
}

func (s square) perimeter() int {
	return s.l * 4
}

func Details(g shape) {
	fmt.Println("Perimeter:", g.perimeter())
	fmt.Println("Area:", g.area())
	}
	

func main() {
	var num shape= rectangle{3, 4}
	v, ok := num.(square)
	if !ok{ 
		fmt.Println("not square")
	} else{
		fmt.Println(v)
	}

	var d shape= square{4}
	v, ok = d.(square)
	if !ok{ 
		fmt.Println("not square")
	}
	fmt.Println(v)
	Details(d)

}
