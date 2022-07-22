// Type declarations adalah cara untuk membuat type data baru dari type data yg sudah ada
// Type declarations juga bisa di bilang membuat alias dari type data yg sudah ada

package main

import "fmt"

func main() {

	type noHandphone string
	type mariedStatus bool

	var handphone noHandphone = "08561841016"
	var status mariedStatus = true

	fmt.Println(handphone)
	fmt.Println(status)

}
