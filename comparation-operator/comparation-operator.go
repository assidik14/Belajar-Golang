// Comparation operator adalah cara untuk membandingkan 2 buah nilai dengan type data boolean (true / false)
// berlaku juga untuk type data string

package main

import "fmt"

func main() {

	var text1 string = "sidik"
	text2 := "sidik"
	var text3 = "assidik"

	var result bool = text1 == text2
	fmt.Println(result)

	result = text1 == text3
	fmt.Println(result)

	fmt.Println("================================")

	var (
		angka1 = 100
		angka2 = 200
	)

	fmt.Println(angka1 == angka2)
	fmt.Println(angka1 > angka2)
	fmt.Println(angka1 < angka2)
	fmt.Println(angka1 != angka2)

}
