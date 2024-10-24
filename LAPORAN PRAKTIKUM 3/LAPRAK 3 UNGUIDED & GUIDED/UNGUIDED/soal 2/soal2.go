package main

import "fmt"

func main() {
	var bmi, tinggiBadan, beratBadan float64
	fmt.Print("Masukkan nilai BMI: ")
	fmt.Scanln(&beratBadan)
	fmt.Print("Masukkan tinggi badan (m): ")
	fmt.Scanln(&tinggiBadan)
	bmi = beratBadan * (tinggiBadan * tinggiBadan)
	fmt.Printf("Berat badan anda: %.f", bmi)
}
