package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Masukkan sebuah ANGKA")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal("Error saat membaca input")
	}

	trimedInput := strings.TrimSpace(input)
	n, err := strconv.ParseInt(trimedInput, 10, 32)

	if err != nil {
		log.Fatal("Input bukan sebuah ANGKA")
	}

	fmt.Print(n)

}
