package main

import "fmt"

func main() {

	var yourname string
	var age int8

	fmt.Print("Siapa nama kamu : ")
	fmt.Scan(&yourname)

	fmt.Print("Berapa umur kamu : ")
	fmt.Scan(&age)

	fmt.Println("Hello", yourname)

	if age >= 17 {
		fmt.Println("Umur kamu", age)
		fmt.Println("Kamu sudah punya KTP")
	} else {
		fmt.Println("Umur kamu", age)
		fmt.Println("Kamu belum punya KTP")
	}

}
