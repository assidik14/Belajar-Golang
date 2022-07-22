package main

import "fmt"

func main() {
	fmt.Println("Ini adalah type data String")
	fmt.Println("===========================")
	fmt.Print("Jumlah karakter dari kata ASSIDIK = ")
	fmt.Println(len("ASSIDIK"))
	fmt.Println("===========================")
	fmt.Println("Achmad Maulana Assidik"[1]) //Mengambil karakter ke-1 dari kata "Achmad Maulana Assidik"
	fmt.Println("Achmad Maulana Assidik"[6]) //Mengambil karakter ke-6 dari kata "Achmad Maulana Assidik"

}
