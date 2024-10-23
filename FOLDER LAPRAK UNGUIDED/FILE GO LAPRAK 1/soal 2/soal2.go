package main

import "fmt"

func main() {
	var (
		nama  string
		prodi = "S1-IF"
		kelas string
		nim   int
	)

	fmt.Print("masukkan nama: ")
	fmt.Scan(&nama)
	fmt.Print("masukkan kelas: ")
	fmt.Scan(&kelas)
	fmt.Print("masukkan NIM: ")
	fmt.Scan(&nim)

	fmt.Println("perkenalkan saya adalah ", nama, "salah satu mahasiswa prodi ", prodi, "dari kelas", kelas, "dengan NIM ", nim)
}
