package main

import "fmt"

func main() {
	var ph float64
	fmt.Print("Masukkan kadar pH: ")
	fmt.Scan(&ph)

	switch {
	case ph >= 6.5 && ph <= 8.6:
		fmt.Println("Air Layak Minum")
	case ph < 0 || ph > 14:
		fmt.Println("Nilai PH Tidak Valid. Nilai pH harus antara 0 dan 14")
	default:
		fmt.Println("Air Tidak Layak Minum")
	}
}
