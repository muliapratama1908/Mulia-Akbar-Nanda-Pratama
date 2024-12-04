package main

import "fmt"

func main() {
	var JK string
	var waktu, tarifparkir, totalbiaya int

	fmt.Print("masukkan jenis kendaraan: ")
	fmt.Scan(&JK)
	fmt.Print("masukkan waktu parkir: ")
	fmt.Scan(&waktu)

	switch JK {
	case "motor":
		tarifparkir = 2000
	case "mobil":
		tarifparkir = 5000
	case "truk":
		tarifparkir = 8000
	default:
		fmt.Println("Jenis kendaraan tidak valid")
		return
	}
	totalbiaya = waktu * tarifparkir
	fmt.Println("total biaya: Rp", totalbiaya)

}
