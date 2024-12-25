package main

import "fmt"

func main() {
	var angka, rataRata float64
	jumlah := 0.0
	count := 0
	fmt.Println("masukkan bilangan (9999 untuk berhenti):")

	for {
		fmt.Scan(&angka)
		if angka == 9999 {
			break
		}
		jumlah += angka
		count++
	}

	if count > 0 {
		rataRata = float64(jumlah) / float64(count)
		fmt.Printf("Rata-rata: %.2f\n", rataRata)
	} else {
		fmt.Println("Tidak ada angka yang dimasukkan.")
	}
}
