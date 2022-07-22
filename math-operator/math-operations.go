package main

import "fmt"

func main() {

	var (
		a = 4
		b = 3
	)

	var result = a + b
	fmt.Println(result)

	result += 10 // operasi augmented assignments atau result = result + 10
	fmt.Println(result)

	result++ // Unary operator atau result = result + 1
	fmt.Println(result)

	result-- // Unary operator atau result = result - 1
	fmt.Println(result)

}
