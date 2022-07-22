// convertion adalah proses yang dilakukan untuk merubah nilai type date dari sebuah variable

package main

import "fmt"

func main() {
	name := "Achmad Maulana Assidik" //variable string
	fmt.Println(name)

	var ambilkarakter byte = name[0]           // mengambil karakter ke - 0 dari nilai variable name
	var convert string = string(ambilkarakter) // mengkonversi nilai type data dari variable ambilkarakter, awalnya type data byte atau uint8 menjadi string

	fmt.Println(ambilkarakter) // nilai karakter ke - 0 dari nilai variable name sebelum di convert
	fmt.Println(convert)       // nilai karakter ke - 0 dari nilai variable name setelah di convert

}
