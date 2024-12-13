package main

import "fmt"

func main() {
	var angka, jumlah int64
	fmt.Scan(&angka)
	jumlah = 0
	for angka > 0 {
		angka = angka / 10
		jumlah++
	}
	fmt.Println("banyaknya digit:", jumlah)
}
