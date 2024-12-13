package main

import "fmt"

func main() {
	var desimal, bilanganBulat float64
	fmt.Print("Masukkan bilangan desimal: ")
	fmt.Scan(&desimal)
	bilanganBulat = float64(int64(desimal)) + 1.0

	for {
		desimal += 0.1
		if desimal >= bilanganBulat-0.0000001 {
			break
		}
		fmt.Printf("%.1f\n", desimal)
	}
	fmt.Print(int64(bilanganBulat))
}
