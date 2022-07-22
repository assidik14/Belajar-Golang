package main

import "os"

func main() {
	file, err := os.Create("tesfile-GO")
	if err != nil {
		panic(err)
	}

	var isifile string = "Hello World\n"
	file.WriteString(isifile)
	file.Close()

}
