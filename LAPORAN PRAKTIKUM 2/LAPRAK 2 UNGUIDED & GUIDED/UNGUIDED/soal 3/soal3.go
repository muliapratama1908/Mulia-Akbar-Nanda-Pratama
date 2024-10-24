package main

import (
	"fmt"
)

func isLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	}
	return false
}

func main() {
	var year int
	fmt.Print("Masukkan tahun: ")
	fmt.Scan(&year)

	if isLeapYear(year) {
		fmt.Println("Tahun", year, "merupakan tahun kabisat (true)")
	} else {
		fmt.Println("Tahun", year, "bukan merupakan tahun kabisat (false)")
	}
}
