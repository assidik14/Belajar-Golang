package main

import (
	"fmt"
)

func Operasiperhitungan() {
	var operations = [5]string{
		"Penjumlahan",
		"Pengurangan",
		"Perkalian",
		"Pembagian",
		"Keluar",
	}

	for i := 0; i < len(operations); i++ {
		fmt.Printf("Operasi %d : %s\n", i, operations[i])
	}
}

func Penjumlahan(a int, b int) {
	var c = a + b
	fmt.Println("Hasilnya = ", c)
}

func Pengurangan(a int, b int) {
	var c = a - b
	fmt.Println("Hasilnya = ", c)
}

func Perkalian(a int, b int) {
	var c = a * b
	fmt.Println("Hasilnya = ", c)
}

func Pembagian(a float32, b float32) {
	var c = a / b
	fmt.Println("Hasilnya = ", c)
}

func main() {

	fmt.Println("===----- Simple Calculator -----===")
	fmt.Println("===-----------------------------===")

	// Menampilkan Operasi Perhitungan
	Operasiperhitungan()

	fmt.Println("===-----------------------------===")

	var inputoperasi int = 0
	for {

		fmt.Print("Masukkan Operasi Perhitungan : ")
		fmt.Scan(&inputoperasi)

		var x int
		var y int
		var xpem float32
		var ypem float32

		if inputoperasi == 0 {
			fmt.Print("Masukkan Nilai 1 : ")
			fmt.Scan(&x)
			fmt.Print("Masukkan Nilai 2 : ")
			fmt.Scan(&y)
			Penjumlahan(x, y)
		} else if inputoperasi == 1 {
			fmt.Print("Masukkan Nilai 1 : ")
			fmt.Scan(&x)
			fmt.Print("Masukkan Nilai 2 : ")
			fmt.Scan(&y)
			Pengurangan(x, y)
		} else if inputoperasi == 2 {
			fmt.Print("Masukkan Nilai 1 : ")
			fmt.Scan(&x)
			fmt.Print("Masukkan Nilai 2 : ")
			fmt.Scan(&y)
			Perkalian(x, y)
		} else if inputoperasi == 3 {
			fmt.Print("Masukkan Nilai 1 : ")
			fmt.Scan(&xpem)
			fmt.Print("Masukkan Nilai 2 : ")
			fmt.Scan(&ypem)
			Pembagian(xpem, ypem)
		} else if inputoperasi == 4 {
			fmt.Println("===-----------------------------===")
			fmt.Println("=====----- Aplikasi Exit -----=====")
			fmt.Println("===-----------------------------===")
			break
		} else if inputoperasi < 0 || inputoperasi > 4 {
			fmt.Println("===-----------------------------===")
			fmt.Println("Operasi yang Anda Masukkan Salah !!")
			// fmt.Println("===--- Silakan Masukan Ulang ---===")
			// Operasiperhitungan()
			fmt.Println("===-----------------------------===")
		} else {
			fmt.Println("===-----------------------------===")
			fmt.Println("Operasi yang Anda Masukkan Salah !!")
			// fmt.Println("===--- Silakan Masukan Ulang ---===")
			// Operasiperhitungan()
			fmt.Println("===-----------------------------===")
		}

		var lanjut string
		fmt.Println("=====------------------------------=====")
		fmt.Println("Ketik 'y' untuk lanjut, 'n' untuk keluar")
		fmt.Println("=====------------------------------=====")
		fmt.Print("Apakah lanjut : ")
		fmt.Scan(&lanjut)
		if lanjut == "y" {
			inputoperasi++
			fmt.Println("===-----------------------------===")
			Operasiperhitungan()
			fmt.Println("===-----------------------------===")
		} else if lanjut == "n" {
			fmt.Println("===-----------------------------===")
			fmt.Println("=====----- Aplikasi Exit -----=====")
			fmt.Println("===-----------------------------===")
			break
		} else {
			fmt.Println("===-----------------------------===")
			fmt.Println("Pilihan yang Anda Masukkan Salah !!")
			fmt.Println("===--- Silakan Masukan Ulang ---===")
			Operasiperhitungan()
			fmt.Println("===-----------------------------===")
		}

	}

}
