package main

import "fmt"

func main() {
	fmt.Println("Aplikasi menghitung luas persegi")
	var p int
	var l int
	var h int
	fmt.Print("Masukkan Panjang : ")
	fmt.Scan(&p)
	fmt.Print("Masukkan Lebar : ")
	fmt.Scan(&l)
	h = p * l
	fmt.Println("Hasilnya adalah", h)
}
