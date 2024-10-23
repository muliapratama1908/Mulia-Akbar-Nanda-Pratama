package main

import "fmt"

func main() {
	var n int
	fmt.Print("masukkan bilangan bulat positif (n): ")
	fmt.Scan(&n)
	fmt.Print("keluaran hasil penjumlahan dari 1 sampai dengan (n): ")

	hasil := 0

	for i := 0; i <= n; i++ {
		hasil += i
	}
	fmt.Println(hasil)
}
