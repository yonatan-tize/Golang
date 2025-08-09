package main

import (
	"fmt"
)

func main() {
	num := 7
	modifyValue(&num)

	fmt.Println(num)
}

func modifyValue(num *int) {
	*num = 20
}
