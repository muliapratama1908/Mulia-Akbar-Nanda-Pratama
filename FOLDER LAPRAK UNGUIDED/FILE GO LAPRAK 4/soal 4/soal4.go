package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	fmt.Println(faktorial(n))
}

func faktorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * faktorial(n-1)
}
