package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Print("x: ")
	fmt.Scan(&x)
	fmt.Print("y: ")
	fmt.Scan(&y)

	var xFaktorY bool
	if y%x == 0 {
		xFaktorY = true
	} else {
		xFaktorY = false
	}
	var yFaktorX bool
	if x%y == 0 {
		yFaktorX = true
	} else {
		yFaktorX = false
	}

	fmt.Println(xFaktorY)
	fmt.Println(yFaktorX)
}
