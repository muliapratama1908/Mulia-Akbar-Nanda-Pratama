package main

import "fmt"

func main() {
	var x, y, hasil int
	fmt.Scan(&x, &y)

	for x >= y {
		x = x - y
		hasil++
	}
	fmt.Print(hasil)
}
