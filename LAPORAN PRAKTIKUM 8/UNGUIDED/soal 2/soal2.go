package main

import "fmt"

func main() {
	var angka int
	fmt.Scanln(&angka)

	if angka <= 0 {
		fmt.Println("Bilangan harus positif.")
		return
	}

	fmt.Println("Digit-digit dari bilangan tersebut:")
	for angka > 0 {
		digit := angka % 10
		fmt.Println(digit)
		angka = angka / 10
	}
}
