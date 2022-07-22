// constant adalah variable yg nilainya tidak bisa diubah

package main

import "fmt"

func main() {
	const firstname string = "Achmad Maulana"
	const lastname = "Assidik"

	// akan error jika variable di ubah
	//firstname = "Achmad"
	//lastname = "Maulana Assidik"

	fmt.Println("My Name is", firstname, lastname)

	// mendeklarasikan constant secara multiple
	const (
		mywife = "Nurmaiyah"
		myboy  = "Eyza"
	)

	fmt.Println("My Wife is", mywife, "& My Boy is", myboy)
}
