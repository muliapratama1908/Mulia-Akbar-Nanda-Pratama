package main

import "fmt"

func main() {
	var target, donasi, totalDonasi, jumlahDonatur int
	fmt.Print("Masukkan target donasi: ")
	fmt.Scanln(&target)

	for totalDonasi < target {
		fmt.Printf("Donatur %d: Masukkan jumlah donasi: ", jumlahDonatur+1)
		fmt.Scanln(&donasi)

		totalDonasi += donasi
		jumlahDonatur++

		fmt.Printf("Donatur %d menyumbang %d. Total terkumpul: %d\n", jumlahDonatur, donasi, totalDonasi)
	}
	fmt.Printf("Target tercapai! Total donasi: %d dari %d donatur.\n", totalDonasi, jumlahDonatur)
}
