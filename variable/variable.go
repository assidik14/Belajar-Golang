package main

import "fmt"

func main() {

	var name string

	name = "Achmad Maulana" // mendeklarasikan varible name
	fmt.Println(name)

	name = "Achmad Maulan Assidik" // mengubah isi variable name
	fmt.Println(name)

	var myboy = "Faeyza" // mendeklarasikan varible tanpa menyebutkan type datanya
	fmt.Println(myboy)

	myboy = "Muhammad Faeyza Nizam" // isi varible tetap bisa di ubah
	fmt.Println(myboy)

	mywife := "Maya" // mendeklarasikan variable tanpa kata kunci "var"
	fmt.Println(mywife)

	mywife = "Nurmaiyah" // isi varible tetap bisa di ubah
	fmt.Println(mywife)

	// ini adalah cara mendeklarasikan multiple varible
	var (
		firtname = "Achmad Maulana"
		lastname = "Assidik"
	)

	fmt.Print(firtname, " ", lastname) // ini adalah cara memanggil multiple varible

}
