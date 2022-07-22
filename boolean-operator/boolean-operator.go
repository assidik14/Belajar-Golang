// Booelan operator adalah perbandingan dua kondisi menggunakan logika matematika (AND, OR, NEGASI) berupa nilai true / false
// AND = &&, OR = ||, NEGASI / Kebalikan = !

package main

import "fmt"

func main() {

	var (
		nilai = 80
		absen = 70
	)

	// var lulusnilai = nilai >= 80
	// var lulusabsen = absen >= 85

	// var lulus = lulusnilai && lulusabsen
	// fmt.Println(lulus)

	var lulus = nilai >= 80 && absen >= 85 // contoh kode yang di persimpel
	fmt.Println(lulus)

}
