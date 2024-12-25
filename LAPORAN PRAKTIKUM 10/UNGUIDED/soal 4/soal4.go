package main

import "fmt"

func main() {
	var n int
	var bunga, pita string

	fmt.Print("N: ")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		fmt.Printf("Bunga %d: ", i)
		fmt.Scanln(&bunga)
		pita += bunga + "-"
	}

	fmt.Println("Pita: ", pita)
}
