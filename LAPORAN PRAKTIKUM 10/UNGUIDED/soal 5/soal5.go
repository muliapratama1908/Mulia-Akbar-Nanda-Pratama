package main

import "fmt"

func main() {
	var berat1, berat2 float64
	fmt.Scan(&berat1, &berat2)
	if berat1 < 0 || berat2 < 0 {
		fmt.Println("Berat tidak boleh negatif")
		return
	}
	totalBerat := berat1 + berat2
	if totalBerat > 150 {
		fmt.Println("Proses selesai")
		return
	}

	if berat1-berat2 >= 9 || berat2-berat1 >= 9 {
		fmt.Println("Sepeda motor Pak Andi akan oleng: true")
	} else {
		fmt.Println("Sepeda motor Pak Andi akan oleng: false")
	}
}
