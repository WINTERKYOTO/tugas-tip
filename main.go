package main

import "fmt"

func main() {
	// Input nilai tagihan restoran
	var bill float64
	fmt.Print("Masukkan nilai tagihan: ")
	fmt.Scan(&bill)

	// Hitung tip berdasarkan kondisi if/else
	var tip float64

	if bill >= 50 && bill <= 300 {
		tip = bill * 0.15
	} else {
		tip = bill * 0.20
	}

	// Hitung total nilai (tagihan + tip)
	total := bill + tip

	// Tampilkan hasil di konsol
	fmt.Printf("Tagihannya %.2f, tipnya %.2f, dan total nilainya %.2f\n", bill, tip, total)
}
